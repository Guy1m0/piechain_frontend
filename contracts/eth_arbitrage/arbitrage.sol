// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.4.22;

contract Arbitrage {
    address public token1Address;
    address public token2Address;
    address public amm1Address;
    address public amm2Address;

    address public lender;
    address public arbitrageur;
    address public exchange;

    uint256 public loan;
    uint256 public intrest;

    bytes32 public loanHash;

    uint8 public status;

    constructor(
        address token1_,
        address token2_,
        address amm1_,
        address amm2_
    ) {
        token1Address = token1_;
        token2Address = token2_;
        amm1Address = amm1_;
        amm2Address = amm2_;
    }

    function setup(
        address lender_,
        address arbitrageur_,
        address exchange_,
        uint256 loan_,
        uint256 intrest_,
        bytes32 loanHash_
    ) public {
        lender = lender_;
        arbitrageur = arbitrageur_;
        exchange = exchange_;
        loanHash = loanHash_;
        loan = loan_;
        intrest = intrest_;
        status = 1;
    }

    function execute(
        uint8 vl,
        bytes32 rl,
        bytes32 sl,
        uint8 va,
        bytes32 ra,
        bytes32 sa
    ) public {
        address lender_ = ecrecover(loanHash, vl, rl, sl);
        address arbitrageur_ = ecrecover(loanHash, va, ra, sa);

        require(lender == lender_, "invalid lender signature");
        require(arbitrageur == arbitrageur_, "invalid arbitrageur signature");

        AMM amm1 = AMM(amm1Address);
        AMM amm2 = AMM(amm2Address);
        ERC20 token1 = ERC20(token1Address);
        ERC20 token2 = ERC20(token2Address);

        token1.approve(amm1Address, loan);
        uint256 t2amount = amm1.exchange1to2(loan);

        token2.approve(amm2Address, t2amount);
        uint256 t1amount = amm2.exchange2to1(t2amount);

        require(t1amount > loan + intrest, "not enough profit");

        token1.transfer(exchange, loan + intrest);
        token1.transfer(arbitrageur, t1amount - loan - intrest);

        status = 2;
    }

    function rollback() public {
        require(
            msg.sender == exchange || msg.sender == arbitrageur,
            "don't have approval"
        );
        ERC20 token1 = ERC20(token1Address);
        token1.transfer(exchange, loan);
        status = 10;
    }
}

contract AMM {
    uint256 public rate1;
    uint256 public rate2;
    address public token1Address;
    address public token2Address;

    constructor(address token1_, address token2_) {
        token1Address = token1_;
        token2Address = token2_;
    }

    function setRate(uint256 rate1_, uint256 rate2_) public {
        rate1 = rate1_;
        rate2 = rate2_;
    }

    function exchange1to2(uint256 amount) public returns (uint256) {
        ERC20 token1 = ERC20(token1Address);
        ERC20 token2 = ERC20(token2Address);

        uint256 result = (amount * rate2) / rate1;
        token1.transferFrom(msg.sender, address(this), amount);
        token2.transfer(msg.sender, result);
        return result;
    }

    function exchange2to1(uint256 amount) public returns (uint256) {
        ERC20 token1 = ERC20(token1Address);
        ERC20 token2 = ERC20(token2Address);

        uint256 result = (amount * rate1) / rate2;
        token2.transferFrom(msg.sender, address(this), amount);
        token1.transfer(msg.sender, result);
        return result;
    }
}

contract ERC20 {
    mapping(address => uint256) private _balances;

    mapping(address => mapping(address => uint256)) private _allowances;

    uint256 private _totalSupply;

    constructor(uint256 initialSupply) {
        _mint(msg.sender, initialSupply);
    }

    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    function transfer(address recipient, uint256 amount) public returns (bool) {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    function allowance(address owner, address spender)
        public
        view
        returns (uint256)
    {
        return _allowances[owner][spender];
    }

    function approve(address spender, uint256 amount) public returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public returns (bool) {
        _transfer(sender, recipient, amount);

        uint256 currentAllowance = _allowances[sender][msg.sender];
        require(
            currentAllowance >= amount,
            "ERC20: transfer amount exceeds allowance"
        );
        _approve(sender, msg.sender, currentAllowance - amount);

        return true;
    }

    function _transfer(
        address sender,
        address recipient,
        uint256 amount
    ) internal {
        require(sender != address(0), "ERC20: transfer from the zero address");
        require(recipient != address(0), "ERC20: transfer to the zero address");

        uint256 senderBalance = _balances[sender];
        require(
            senderBalance >= amount,
            "ERC20: transfer amount exceeds balance"
        );
        _balances[sender] = senderBalance - amount;
        _balances[recipient] += amount;
    }

    function _mint(address account, uint256 amount) internal {
        require(account != address(0), "ERC20: mint to the zero address");

        _totalSupply += amount;
        _balances[account] += amount;
    }

    function _approve(
        address owner,
        address spender,
        uint256 amount
    ) internal {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
    }
}
