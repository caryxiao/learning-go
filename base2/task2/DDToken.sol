// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
contract DDToken is IERC20 {
    address manager;
    uint public constant MAX_SUPPLY = 100000000000;
    uint private _totalSupply;
    string private _name = "Demo Data Token";
    string private _symbol = "DDT";
    mapping(address => uint) private _balances;
    mapping(address => mapping(address => uint)) private _allowances;

    constructor() {
        manager = msg.sender;
        _mint(msg.sender, 1000000);
    }

    function name() public view returns (string memory) {
        return _name;
    }

    function symbol() public view returns (string memory) {
        return _symbol;
    }

    function decimals() public pure returns (uint8) {
        return 6;
    }

    modifier isManager {
        require(msg.sender == manager, "Only owner can call this function");
        _;
    }

    function totalSupply() external view returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) external view returns (uint256) {
        return _balances[account];
    }


    function transfer(address to, uint256 value) external returns (bool) {
        return _trans(msg.sender, to, value);
    }

    function _trans(address from, address to, uint256 value) internal returns (bool) {
        require(_balances[from] >= value, "Insufficient balance!");
        _balances[from] -= value;
        _balances[to] += value;
        emit Transfer(from, to, value);
        return true;
    }

    function allowance(address owner, address spender) external view returns (uint256) {
        return _allowances[owner][spender];
    }


    function approve(address spender, uint256 value) external returns (bool) {
        require(spender != address(0));
        require(msg.sender != spender, "You cannot authorize yourself.");
        _allowances[msg.sender][spender] = value;
        emit Approval(msg.sender, spender, value);
        return true;
    }

    function transferFrom(address from, address to, uint256 value) external returns (bool) {
        require(_balances[from] >= value, "Insufficient balance!");
        require(_allowances[from][msg.sender] >= value, "Unauthorized or insufficient authorization limit!");
        _allowances[from][msg.sender] -= value;
        return _trans(from, to, value);
    }

    function mint(address to, uint256 value) external isManager returns (bool) {
        _mint(to, value);
        return true;
    }

    function _mint(address to, uint256 value) internal {
        require(_totalSupply + value <= MAX_SUPPLY, string(abi.encodePacked("The total supply must be less than or equal ", MAX_SUPPLY)));
        _totalSupply += value;
        _balances[to] += value;
        emit Transfer(address(0), to, value);
    }
}