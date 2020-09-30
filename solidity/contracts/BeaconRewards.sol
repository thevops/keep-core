/**
▓▓▌ ▓▓ ▐▓▓ ▓▓▓▓▓▓▓▓▓▓▌▐▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▄
▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▌▐▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
  ▓▓▓▓▓▓    ▓▓▓▓▓▓▓▀    ▐▓▓▓▓▓▓    ▐▓▓▓▓▓   ▓▓▓▓▓▓     ▓▓▓▓▓   ▐▓▓▓▓▓▌   ▐▓▓▓▓▓▓
  ▓▓▓▓▓▓▄▄▓▓▓▓▓▓▓▀      ▐▓▓▓▓▓▓▄▄▄▄         ▓▓▓▓▓▓▄▄▄▄         ▐▓▓▓▓▓▌   ▐▓▓▓▓▓▓
  ▓▓▓▓▓▓▓▓▓▓▓▓▓▀        ▐▓▓▓▓▓▓▓▓▓▓         ▓▓▓▓▓▓▓▓▓▓▌        ▐▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
  ▓▓▓▓▓▓▀▀▓▓▓▓▓▓▄       ▐▓▓▓▓▓▓▀▀▀▀         ▓▓▓▓▓▓▀▀▀▀         ▐▓▓▓▓▓▓▓▓▓▓▓▓▓▓▀
  ▓▓▓▓▓▓   ▀▓▓▓▓▓▓▄     ▐▓▓▓▓▓▓     ▓▓▓▓▓   ▓▓▓▓▓▓     ▓▓▓▓▓   ▐▓▓▓▓▓▌
▓▓▓▓▓▓▓▓▓▓ █▓▓▓▓▓▓▓▓▓ ▐▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓ ▐▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓▓▓▓▓▓▓▓▓

                           Trust math, not hardware.
*/

pragma solidity ^0.5.17;

import "./Rewards.sol";
import "./KeepRandomBeaconOperator.sol";
import "./TokenStaking.sol";

contract BeaconRewards is Rewards {

    // 
    // Random Beacon Rewards constants
    // 
    uint256[] internal beaconIntervalWeights = [
        4, 8, 10, 12, 15, 15,
        15, 15, 15, 15, 15, 15,
        15, 15, 15, 15, 15, 15,
        15, 15, 15, 15, 15, 15
    ];

    uint256 internal constant beaconTermLength = 30 days;

    uint256 internal constant minimumBeaconGroupsPerInterval = 2;
    //
    //

    KeepRandomBeaconOperator operatorContract;
    TokenStaking tokenStaking;

    constructor (
        address _token,
        uint256 _firstIntervalStart, // TODO: can we accept this parameter as the last one? should we make it a constant in this contract?
        address _operatorContract,
        address _stakingContract
    ) public Rewards(
        _token,
        _firstIntervalStart,
        beaconIntervalWeights,
        beaconTermLength,
        minimumBeaconGroupsPerInterval
    ) {
        operatorContract = KeepRandomBeaconOperator(_operatorContract);
        tokenStaking = TokenStaking(_stakingContract);
    }

    function receiveReward(uint256 i) public {
        receiveReward(bytes32(i));
    }

    function eligibleForReward(uint256 i) public view returns (bool) {
        return eligibleForReward(bytes32(i));
    }

    function reportTermination(uint256 i) public {
        reportTermination(bytes32(i));
    }

    function isTerminated(uint256 i) public view returns (bool) {
        return eligibleButTerminated(bytes32(i));
    }

    function _getKeepCount() internal view returns (uint256) {
        return operatorContract.getNumberOfCreatedGroups();
    }

    function _getKeepAtIndex(uint256 i) internal view returns (bytes32) {
        return bytes32(i);
    }

    function _getCreationTime(bytes32 groupIndexBytes) internal view returns (uint256) {
        return operatorContract.getGroupRegistrationTime(uint256(groupIndexBytes));
    }

    function _isClosed(bytes32 groupIndexBytes) internal view returns (bool) {
        if (_isTerminated(groupIndexBytes)) { return false; }
        bytes memory groupPubkey = operatorContract.getGroupPublicKey(
            uint256(groupIndexBytes)
        );
        return operatorContract.isStaleGroup(groupPubkey);
    }

    function _isTerminated(bytes32 groupIndexBytes) internal view returns (bool) {
        return operatorContract.isGroupTerminated(uint256(groupIndexBytes));
    }

    function _recognizedByFactory(bytes32 groupIndexBytes) internal view returns (bool) {
        return _getKeepCount() > uint256(groupIndexBytes);
    }

    function _distributeReward(bytes32 groupIndexBytes, uint256 _value) internal {
        bytes memory groupPubkey = operatorContract.getGroupPublicKey(
            uint256(groupIndexBytes)
        );
        address[] memory members = operatorContract.getGroupMembers(groupPubkey);

        uint256 memberCount = members.length;
        uint256 dividend = _value.div(memberCount);

        // Only pay other members if dividend is nonzero.
        if(dividend > 0) {
            for (uint256 i = 0; i < memberCount - 1; i++) {
                token.safeTransfer(
                    tokenStaking.beneficiaryOf(members[i]),
                    dividend
                );
            }
        }

        // Transfer of dividend for the last member. Remainder might be equal to
        // zero in case of even distribution or some small number.
        uint256 remainder = _value.mod(memberCount);
        token.safeTransfer(
            tokenStaking.beneficiaryOf(members[memberCount - 1]),
            dividend.add(remainder)
        );
    }
}
