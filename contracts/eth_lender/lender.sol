// abigen --sol Lender.sol  --pkg eth_lender --type Lender --out lender_gen.go

pragma solidity >=0.4.22 <0.6;

contract Lender {
    address public tokenAddress;
    address payable public lender;
    address payable public arbitrage;
    address payable public exchange;
    uint256 public loan;
    uint256 public intrest;
    bytes32 public loanHash;

    uint256 public status;

    constructor(address tokenAddress_) public {
        tokenAddress = tokenAddress_;
    }

    function setup(
        address payable lender_, 
        address payable arbitrage_, 
        address payable exchange_,
        uint256 loan_, uint256 intrest_, bytes32 loanHash_
    ) public {
        lender = lender_;
        arbitrage = arbitrage_;
        exchange = exchange_;
        loan = loan_;
        intrest = intrest_;
        loanHash = loanHash_;
    }

    function initialize(
        uint8 vl, bytes32 rl, bytes32 sl, 
        uint8 va, bytes32 ra, bytes32 sa
    ) public {
        // require(lender == ecrecover(loanHash, vl, rl, sl), "Invalid Lender Signature");
        // require(arbitrage == ecrecover(loanHash, va, ra, sa), "Invalid Arbitrage Signature");
        ERC20 token = ERC20(tokenAddress);
        if (token.transferFrom(lender, exchange, loan)) {
            status = 1;
        }
    }

    function endLoan(bool success) public {
        uint256 refund = loan;
        if (success) {
            refund += intrest;
        }
        ERC20 token = ERC20(tokenAddress);
        if (token.transferFrom(exchange, lender, refund)) {
            status = 2;
        }
    }

}

contract ERC20 {
    mapping(address => uint256) private _balances;

    mapping(address => mapping(address => uint256)) private _allowances;

    uint256 private _totalSupply;

    constructor(uint256 initialSupply) public {
        _mint(msg.sender, initialSupply);
    }

    function totalSupply() public view  returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) public view  returns (uint256) {
        return _balances[account];
    }

    function transfer(address recipient, uint256 loan) public  returns (bool) {
        _transfer(msg.sender, recipient, loan);
        return true;
    }

    function allowance(address owner, address spender) public view  returns (uint256) {
        return _allowances[owner][spender];
    }

    function approve(address spender, uint256 loan) public  returns (bool) {
        _approve(msg.sender, spender, loan);
        return true;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 loan
    ) public  returns (bool) {
        _transfer(sender, recipient, loan);

        uint256 currentAllowance = _allowances[sender][msg.sender];
        require(currentAllowance >= loan, "ERC20: transfer loan exceeds allowance");
        _approve(sender, msg.sender, currentAllowance - loan);

        return true;
    }

    function _transfer(
        address sender,
        address recipient,
        uint256 loan
    ) internal {
        require(sender != address(0), "ERC20: transfer from the zero address");
        require(recipient != address(0), "ERC20: transfer to the zero address");

        uint256 senderBalance = _balances[sender];
        require(senderBalance >= loan, "ERC20: transfer loan exceeds balance");
        _balances[sender] = senderBalance - loan;
        _balances[recipient] += loan;
    }

    function _mint(address account, uint256 loan) internal {
        require(account != address(0), "ERC20: mint to the zero address");

        _totalSupply += loan;
        _balances[account] += loan;
    }

    function _approve(
        address owner,
        address spender,
        uint256 loan
    ) internal {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = loan;
    }

}
