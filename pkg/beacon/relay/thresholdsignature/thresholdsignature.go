package thresholdsignature

import (
	"fmt"
	"time"

	"github.com/dfinity/go-dfinity-crypto/bls"
	"github.com/keep-network/keep-core/pkg/chain"
	"github.com/keep-network/keep-core/pkg/net"
	"github.com/keep-network/keep-core/pkg/thresholdgroup"
)

const (
	setupBlocks     = 15
	signatureBlocks = 10
)

// Init initializes a given broadcast channel to be able to perform distributed
// key generation interactions.
func Init(channel net.BroadcastChannel) {
	channel.RegisterUnmarshaler(
		func() net.TaggedUnmarshaler { return &SignatureShareMessage{} })
}

// Execute triggers the threshold signature process for the given bytes. After
// the process has completed, it returns either the threshold signature's final
// bytes, or an error.
func Execute(
	bytes []byte,
	blockCounter chain.BlockCounter,
	channel net.BroadcastChannel,
	member *thresholdgroup.Member,
) ([]byte, error) {
	// Use an unbuffered channel to serialize message processing.
	recvChan := make(chan net.Message)
	handler := net.HandleMessageFunc{
		Type: fmt.Sprintf("relay/signature/%s", string(time.Now().UTC().UnixNano())),
		Handler: func(msg net.Message) error {
			recvChan <- msg
			return nil
		},
	}

	channel.Recv(handler)
	defer channel.UnregisterRecv(handler.Type)

	seenShares := make(map[bls.ID][]byte)

	fmt.Printf(
		"[member:%v, state:signing] Waiting for other group members to enter signing state...\n",
		member.MemberID(),
	)

	err := blockCounter.WaitForBlocks(setupBlocks)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to wait %d blocks entering threshold setup: [%v]",
			setupBlocks,
			err,
		)
	}

	fmt.Printf(
		"[member:%v] Sending signature share...\n",
		member.MemberID(),
	)

	err = blockCounter.WaitForBlocks(1)
	if err != nil {
		return nil, fmt.Errorf("failed to wait 1 block: [%v]", err)
	}

	err = sendSignatureShare(bytes, channel, member)
	if err != nil {
		return nil, err
	}

	blockWaiter, err := blockCounter.BlockWaiter(signatureBlocks)
	if err != nil {
		return nil, err
	}

	fmt.Printf("[member:%v] Receiving other group signature share.\n", member.ID)

	for {
		select {
		case msg := <-recvChan:
			fmt.Printf(
				"[member:%v, state:signing] Processing message.\n",
				member.MemberID(),
			)

			switch signatureShareMsg := msg.Payload().(type) {
			case *SignatureShareMessage:
				if senderID, ok := msg.ProtocolSenderID().(*bls.ID); ok {
					seenShares[*senderID] = signatureShareMsg.ShareBytes
				}
			}

		case <-blockWaiter:
			signature, err := member.CompleteSignature(seenShares)
			if err != nil {
				return nil, fmt.Errorf(
					"[member:%v] failed to complete signature inside active period [%v]",
					member.MemberID(),
					signatureBlocks,
				)
			}

			return signature.Serialize(), nil

		}
	}
}

func sendSignatureShare(
	bytes []byte,
	channel net.BroadcastChannel,
	member *thresholdgroup.Member,
) error {
	share := member.SignatureShare(string(bytes))
	return channel.Send(&SignatureShareMessage{&member.BlsID, share})
}
