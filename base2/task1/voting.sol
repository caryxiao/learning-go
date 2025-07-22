// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract voting {
    mapping(address candidate => uint voteCount) private voteMap;
    mapping(address candidate => uint voteVersion) private  voteVersionMap;
    mapping(address voter => uint voteVersion) private voterMap;
    uint public currentVersion;
    address public owner;
    constructor(){
        owner = msg.sender;
        // 初始化的时候投票版本为1
        currentVersion = 1;
    }

    modifier isOwner() {
        require(owner == msg.sender, "Caller is not owner");
        _;
    }

    function vote(address candidate) public returns (string memory) {
        // 判断当前版本用户是否已经投票, 已经投票过的是不能再次进行投票的
        require(voterMap[msg.sender] != currentVersion, "This user has voted.");
        voterMap[msg.sender] = currentVersion;
        // 如果候选人当前版本被投票过，则直接使用投票的历史票数相加，否则历史本次投票作为新一轮投票的开始，本次投票的票数归零，本次投票数为1
        if (voteVersionMap[candidate] == currentVersion) {
            voteMap[candidate] += 1;
        } else {
            voteMap[candidate] = 1;
            voteVersionMap[candidate] = currentVersion;
        }

        return "Voted successful!";
    }

    function getVotes(address candidate) public view returns(uint) {
        if(voteVersionMap[candidate] == currentVersion){
            return voteMap[candidate];
        }
        return 0;
    }

    // 只有管理员可以重置所有的票数和投票状态
    function resetVotes() public isOwner returns (string memory) {
        currentVersion += 1;
        return "success";
    }
}