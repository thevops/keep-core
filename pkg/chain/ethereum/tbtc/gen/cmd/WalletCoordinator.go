// Code generated - DO NOT EDIT.
// This file is a generated command and any manual changes will be lost.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	chainutil "github.com/keep-network/keep-common/pkg/chain/ethereum/ethutil"
	"github.com/keep-network/keep-common/pkg/cmd"
	"github.com/keep-network/keep-common/pkg/utils/decode"
	"github.com/keep-network/keep-core/pkg/chain/ethereum/tbtc/gen/abi"
	"github.com/keep-network/keep-core/pkg/chain/ethereum/tbtc/gen/contract"

	"github.com/spf13/cobra"
)

var WalletCoordinatorCommand *cobra.Command

var walletCoordinatorDescription = `The wallet-coordinator command allows calling the WalletCoordinator contract on an
	Ethereum network. It has subcommands corresponding to each contract method,
	which respectively each take parameters based on the contract method's
	parameters.

	Subcommands will submit a non-mutating call to the network and output the
	result.

	All subcommands can be called against a specific block by passing the
	-b/--block flag.

	Subcommands for mutating methods may be submitted as a mutating transaction
	by passing the -s/--submit flag. In this mode, this command will terminate
	successfully once the transaction has been submitted, but will not wait for
	the transaction to be included in a block. They return the transaction hash.

	Calls that require ether to be paid will get 0 ether by default, which can
	be changed by passing the -v/--value flag.`

func init() {
	WalletCoordinatorCommand := &cobra.Command{
		Use:   "wallet-coordinator",
		Short: `Provides access to the WalletCoordinator contract.`,
		Long:  walletCoordinatorDescription,
	}

	WalletCoordinatorCommand.AddCommand(
		wcBridgeCommand(),
		wcDepositMinAgeCommand(),
		wcDepositRefundSafetyMarginCommand(),
		wcDepositSweepMaxSizeCommand(),
		wcDepositSweepProposalSubmissionGasOffsetCommand(),
		wcDepositSweepProposalValidityCommand(),
		wcIsProposalSubmitterCommand(),
		wcOwnerCommand(),
		wcReimbursementPoolCommand(),
		wcWalletRegistryCommand(),
		wcAddProposalSubmitterCommand(),
		wcInitializeCommand(),
		wcRemoveProposalSubmitterCommand(),
		wcRenounceOwnershipCommand(),
		wcSubmitDepositSweepProposalCommand(),
		wcSubmitDepositSweepProposalWithReimbursementCommand(),
		wcTransferOwnershipCommand(),
		wcUpdateDepositMinAgeCommand(),
		wcUpdateDepositRefundSafetyMarginCommand(),
		wcUpdateDepositSweepMaxSizeCommand(),
		wcUpdateDepositSweepProposalSubmissionGasOffsetCommand(),
		wcUpdateDepositSweepProposalValidityCommand(),
		wcUpdateReimbursementPoolCommand(),
	)

	ModuleCommand.AddCommand(WalletCoordinatorCommand)
}

/// ------------------- Const methods -------------------

func wcBridgeCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "bridge",
		Short:                 "Calls the view method bridge on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcBridge,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcBridge(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.BridgeAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcDepositMinAgeCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "deposit-min-age",
		Short:                 "Calls the view method depositMinAge on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcDepositMinAge,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcDepositMinAge(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.DepositMinAgeAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcDepositRefundSafetyMarginCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "deposit-refund-safety-margin",
		Short:                 "Calls the view method depositRefundSafetyMargin on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcDepositRefundSafetyMargin,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcDepositRefundSafetyMargin(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.DepositRefundSafetyMarginAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcDepositSweepMaxSizeCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "deposit-sweep-max-size",
		Short:                 "Calls the view method depositSweepMaxSize on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcDepositSweepMaxSize,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcDepositSweepMaxSize(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.DepositSweepMaxSizeAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcDepositSweepProposalSubmissionGasOffsetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "deposit-sweep-proposal-submission-gas-offset",
		Short:                 "Calls the view method depositSweepProposalSubmissionGasOffset on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcDepositSweepProposalSubmissionGasOffset,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcDepositSweepProposalSubmissionGasOffset(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.DepositSweepProposalSubmissionGasOffsetAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcDepositSweepProposalValidityCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "deposit-sweep-proposal-validity",
		Short:                 "Calls the view method depositSweepProposalValidity on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcDepositSweepProposalValidity,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcDepositSweepProposalValidity(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.DepositSweepProposalValidityAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcIsProposalSubmitterCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "is-proposal-submitter [arg0]",
		Short:                 "Calls the view method isProposalSubmitter on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcIsProposalSubmitter,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcIsProposalSubmitter(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg0, err := chainutil.AddressFromHex(args[0])
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg0, a address, from passed value %v",
			args[0],
		)
	}

	result, err := contract.IsProposalSubmitterAtBlock(
		arg0,
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcOwnerCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "owner",
		Short:                 "Calls the view method owner on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcOwner,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcOwner(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.OwnerAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcReimbursementPoolCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "reimbursement-pool",
		Short:                 "Calls the view method reimbursementPool on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcReimbursementPool,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcReimbursementPool(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.ReimbursementPoolAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

func wcWalletRegistryCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "wallet-registry",
		Short:                 "Calls the view method walletRegistry on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcWalletRegistry,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	cmd.InitConstFlags(c)

	return c
}

func wcWalletRegistry(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	result, err := contract.WalletRegistryAtBlock(
		cmd.BlockFlagValue.Int,
	)

	if err != nil {
		return err
	}

	cmd.PrintOutput(result)

	return nil
}

/// ------------------- Non-const methods -------------------

func wcAddProposalSubmitterCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "add-proposal-submitter [arg_proposalSubmitter]",
		Short:                 "Calls the nonpayable method addProposalSubmitter on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcAddProposalSubmitter,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcAddProposalSubmitter(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg_proposalSubmitter, err := chainutil.AddressFromHex(args[0])
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg_proposalSubmitter, a address, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.AddProposalSubmitter(
			arg_proposalSubmitter,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallAddProposalSubmitter(
			arg_proposalSubmitter,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcInitializeCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "initialize [arg__bridge]",
		Short:                 "Calls the nonpayable method initialize on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcInitialize,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcInitialize(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg__bridge, err := chainutil.AddressFromHex(args[0])
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg__bridge, a address, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.Initialize(
			arg__bridge,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallInitialize(
			arg__bridge,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcRemoveProposalSubmitterCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "remove-proposal-submitter [arg_proposalSubmitter]",
		Short:                 "Calls the nonpayable method removeProposalSubmitter on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcRemoveProposalSubmitter,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcRemoveProposalSubmitter(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg_proposalSubmitter, err := chainutil.AddressFromHex(args[0])
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg_proposalSubmitter, a address, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.RemoveProposalSubmitter(
			arg_proposalSubmitter,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallRemoveProposalSubmitter(
			arg_proposalSubmitter,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcRenounceOwnershipCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "renounce-ownership",
		Short:                 "Calls the nonpayable method renounceOwnership on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(0),
		RunE:                  wcRenounceOwnership,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcRenounceOwnership(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.RenounceOwnership()
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallRenounceOwnership(
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcSubmitDepositSweepProposalCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "submit-deposit-sweep-proposal [arg_proposal_json] [arg_walletMemberContext_json]",
		Short:                 "Calls the nonpayable method submitDepositSweepProposal on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(2),
		RunE:                  wcSubmitDepositSweepProposal,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcSubmitDepositSweepProposal(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg_proposal_json := abi.WalletCoordinatorDepositSweepProposal{}
	if err := json.Unmarshal([]byte(args[0]), &arg_proposal_json); err != nil {
		return fmt.Errorf("failed to unmarshal arg_proposal_json to abi.WalletCoordinatorDepositSweepProposal: %w", err)
	}

	arg_walletMemberContext_json := abi.WalletCoordinatorWalletMemberContext{}
	if err := json.Unmarshal([]byte(args[1]), &arg_walletMemberContext_json); err != nil {
		return fmt.Errorf("failed to unmarshal arg_walletMemberContext_json to abi.WalletCoordinatorWalletMemberContext: %w", err)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.SubmitDepositSweepProposal(
			arg_proposal_json,
			arg_walletMemberContext_json,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallSubmitDepositSweepProposal(
			arg_proposal_json,
			arg_walletMemberContext_json,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcSubmitDepositSweepProposalWithReimbursementCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "submit-deposit-sweep-proposal-with-reimbursement [arg_proposal_json] [arg_walletMemberContext_json]",
		Short:                 "Calls the nonpayable method submitDepositSweepProposalWithReimbursement on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(2),
		RunE:                  wcSubmitDepositSweepProposalWithReimbursement,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcSubmitDepositSweepProposalWithReimbursement(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg_proposal_json := abi.WalletCoordinatorDepositSweepProposal{}
	if err := json.Unmarshal([]byte(args[0]), &arg_proposal_json); err != nil {
		return fmt.Errorf("failed to unmarshal arg_proposal_json to abi.WalletCoordinatorDepositSweepProposal: %w", err)
	}

	arg_walletMemberContext_json := abi.WalletCoordinatorWalletMemberContext{}
	if err := json.Unmarshal([]byte(args[1]), &arg_walletMemberContext_json); err != nil {
		return fmt.Errorf("failed to unmarshal arg_walletMemberContext_json to abi.WalletCoordinatorWalletMemberContext: %w", err)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.SubmitDepositSweepProposalWithReimbursement(
			arg_proposal_json,
			arg_walletMemberContext_json,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallSubmitDepositSweepProposalWithReimbursement(
			arg_proposal_json,
			arg_walletMemberContext_json,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcTransferOwnershipCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "transfer-ownership [arg_newOwner]",
		Short:                 "Calls the nonpayable method transferOwnership on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcTransferOwnership,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcTransferOwnership(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg_newOwner, err := chainutil.AddressFromHex(args[0])
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg_newOwner, a address, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.TransferOwnership(
			arg_newOwner,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallTransferOwnership(
			arg_newOwner,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcUpdateDepositMinAgeCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "update-deposit-min-age [arg__depositMinAge]",
		Short:                 "Calls the nonpayable method updateDepositMinAge on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcUpdateDepositMinAge,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcUpdateDepositMinAge(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg__depositMinAge, err := decode.ParseUint[uint32](args[0], 32)
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg__depositMinAge, a uint32, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.UpdateDepositMinAge(
			arg__depositMinAge,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallUpdateDepositMinAge(
			arg__depositMinAge,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcUpdateDepositRefundSafetyMarginCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "update-deposit-refund-safety-margin [arg__depositRefundSafetyMargin]",
		Short:                 "Calls the nonpayable method updateDepositRefundSafetyMargin on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcUpdateDepositRefundSafetyMargin,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcUpdateDepositRefundSafetyMargin(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg__depositRefundSafetyMargin, err := decode.ParseUint[uint32](args[0], 32)
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg__depositRefundSafetyMargin, a uint32, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.UpdateDepositRefundSafetyMargin(
			arg__depositRefundSafetyMargin,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallUpdateDepositRefundSafetyMargin(
			arg__depositRefundSafetyMargin,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcUpdateDepositSweepMaxSizeCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "update-deposit-sweep-max-size [arg__depositSweepMaxSize]",
		Short:                 "Calls the nonpayable method updateDepositSweepMaxSize on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcUpdateDepositSweepMaxSize,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcUpdateDepositSweepMaxSize(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg__depositSweepMaxSize, err := decode.ParseUint[uint16](args[0], 16)
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg__depositSweepMaxSize, a uint16, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.UpdateDepositSweepMaxSize(
			arg__depositSweepMaxSize,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallUpdateDepositSweepMaxSize(
			arg__depositSweepMaxSize,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcUpdateDepositSweepProposalSubmissionGasOffsetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "update-deposit-sweep-proposal-submission-gas-offset [arg__depositSweepProposalSubmissionGasOffset]",
		Short:                 "Calls the nonpayable method updateDepositSweepProposalSubmissionGasOffset on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcUpdateDepositSweepProposalSubmissionGasOffset,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcUpdateDepositSweepProposalSubmissionGasOffset(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg__depositSweepProposalSubmissionGasOffset, err := decode.ParseUint[uint32](args[0], 32)
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg__depositSweepProposalSubmissionGasOffset, a uint32, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.UpdateDepositSweepProposalSubmissionGasOffset(
			arg__depositSweepProposalSubmissionGasOffset,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallUpdateDepositSweepProposalSubmissionGasOffset(
			arg__depositSweepProposalSubmissionGasOffset,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcUpdateDepositSweepProposalValidityCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "update-deposit-sweep-proposal-validity [arg__depositSweepProposalValidity]",
		Short:                 "Calls the nonpayable method updateDepositSweepProposalValidity on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcUpdateDepositSweepProposalValidity,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcUpdateDepositSweepProposalValidity(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg__depositSweepProposalValidity, err := decode.ParseUint[uint32](args[0], 32)
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg__depositSweepProposalValidity, a uint32, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.UpdateDepositSweepProposalValidity(
			arg__depositSweepProposalValidity,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallUpdateDepositSweepProposalValidity(
			arg__depositSweepProposalValidity,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

func wcUpdateReimbursementPoolCommand() *cobra.Command {
	c := &cobra.Command{
		Use:                   "update-reimbursement-pool [arg__reimbursementPool]",
		Short:                 "Calls the nonpayable method updateReimbursementPool on the WalletCoordinator contract.",
		Args:                  cmd.ArgCountChecker(1),
		RunE:                  wcUpdateReimbursementPool,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
	}

	c.PreRunE = cmd.NonConstArgsChecker
	cmd.InitNonConstFlags(c)

	return c
}

func wcUpdateReimbursementPool(c *cobra.Command, args []string) error {
	contract, err := initializeWalletCoordinator(c)
	if err != nil {
		return err
	}

	arg__reimbursementPool, err := chainutil.AddressFromHex(args[0])
	if err != nil {
		return fmt.Errorf(
			"couldn't parse parameter arg__reimbursementPool, a address, from passed value %v",
			args[0],
		)
	}

	var (
		transaction *types.Transaction
	)

	if shouldSubmit, _ := c.Flags().GetBool(cmd.SubmitFlag); shouldSubmit {
		// Do a regular submission. Take payable into account.
		transaction, err = contract.UpdateReimbursementPool(
			arg__reimbursementPool,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput(transaction.Hash())
	} else {
		// Do a call.
		err = contract.CallUpdateReimbursementPool(
			arg__reimbursementPool,
			cmd.BlockFlagValue.Int,
		)
		if err != nil {
			return err
		}

		cmd.PrintOutput("success")

		cmd.PrintOutput(
			"the transaction was not submitted to the chain; " +
				"please add the `--submit` flag",
		)
	}

	return nil
}

/// ------------------- Initialization -------------------

func initializeWalletCoordinator(c *cobra.Command) (*contract.WalletCoordinator, error) {
	cfg := *ModuleCommand.GetConfig()

	client, err := ethclient.Dial(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to host chain node: [%v]", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf(
			"failed to resolve host chain id: [%v]",
			err,
		)
	}

	key, err := chainutil.DecryptKeyFile(
		cfg.Account.KeyFile,
		cfg.Account.KeyFilePassword,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read KeyFile: %s: [%v]",
			cfg.Account.KeyFile,
			err,
		)
	}

	miningWaiter := chainutil.NewMiningWaiter(client, cfg)

	blockCounter, err := chainutil.NewBlockCounter(client)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create block counter: [%v]",
			err,
		)
	}

	address, err := cfg.ContractAddress("WalletCoordinator")
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get %s address: [%w]",
			"WalletCoordinator",
			err,
		)
	}

	return contract.NewWalletCoordinator(
		address,
		chainID,
		key,
		client,
		chainutil.NewNonceManager(client, key.Address),
		miningWaiter,
		blockCounter,
		&sync.Mutex{},
	)
}
