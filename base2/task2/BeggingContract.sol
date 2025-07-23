// SPDX-License-Identifier: MIT
// 已部署的测试地址: https://sepolia.etherscan.io/address/0x7f02e21e770b5a4afb8a0c42a44092a973fe42d6
pragma solidity ^0.8.10;
contract BeggingContract {
    address payable private owner;
    // 捐赠排行
    mapping(address donor => uint amount) private donation_list;
    address[3] private donation_top3;
    // 提取所有的ETH
    event Withdrawed(address indexed donor, uint amount);
    event Donated(address indexed donor, uint amount);
    // Top3 排行事件
    event DonantedTop3(address indexed donor, uint top);

    constructor() {
        owner = payable(msg.sender);
    }

    modifier isOwner {
        require(msg.sender == owner, "You are not the owner");
        _;
    }

    modifier canDonate {
        require(onlyInCNTime(), "Only allowed from 09:00 to 24:00 China time");
        _;
    }

    function donate() external payable canDonate {
        require(msg.value > 0, "Send ETH with the call");
        donation_list[msg.sender] += msg.value;
        _updateDonationTop3(msg.sender);
        emit Donated(msg.sender, msg.value);
    }

    function withdraw() external isOwner payable  {
        require(address(this).balance > 0, "balance is zero.");
        uint amount = address(this).balance;
        owner.transfer(amount);
        emit Withdrawed(msg.sender, amount);
    }

    function getDonation() external view returns (uint) {
        return donation_list[msg.sender];
    }

    function _updateDonationTop3(address donor) internal {
        uint i = 0;
        while (i < donation_top3.length && donation_list[donor] <= donation_list[donation_top3[i]]) {
            i++;
        }

        if (i >= donation_top3.length) {
            return;
        }

        for (uint j = donation_top3.length - 1; j > i; j--) {
            donation_top3[j] = donation_top3[j - 1];
        }

        donation_top3[i] = donor;
        emit DonantedTop3(donor, i + 1);
    }

    function getDonationTop3() external view returns (address[3] memory) {
        return donation_top3;
    }

    function onlyInCNTime() public view returns (bool) {
        // 中国时间 = UTC + 8 小时
        uint256 hour = ((block.timestamp + 8 * 3600) / 3600) % 24;
        return hour >= 9 && hour < 24;
    }
}