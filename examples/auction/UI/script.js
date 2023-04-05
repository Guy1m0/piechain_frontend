document.addEventListener("DOMContentLoaded", () => {
    // Initialize event listeners
    initEventListeners();
  
    // Display auctions on the dashboard
    displayAuctions();
  });
  
  function initEventListeners() {
    const createAuctionForm = document.getElementById("create-auction-form");
    if (createAuctionForm) {
      createAuctionForm.addEventListener("submit", (event) => {
        event.preventDefault();
        createAuction();
      });
    }
  
    const viewAuctionPage = document.getElementById("view-auction-page");
    if (viewAuctionPage) {
      const auctionId = new URLSearchParams(window.location.search).get("id");
      loadAuction(auctionId);
    }
  }

  function resetAuctions() {
    localStorage.removeItem('auctions');
    auctions = [];
    displayAuctions();
  }
  
  function createAuction() {
    const nftType = document.getElementById("nft-type").value;
    const smartContractAddress = document.getElementById("smart-contract-address").value;
    const conclusionTime = document.getElementById("conclusion-time").value;
  
    const auction = {
      id: Date.now(),
      nftType: nftType,
      smartContractAddress: smartContractAddress,
      conclusionTime: new Date(conclusionTime).getTime(),
      bids: [],
    };
  
    let auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    auctions.push(auction);
    localStorage.setItem("auctions", JSON.stringify(auctions));
  
    window.location.href = "index.html";
    displayAuctions();
  }
  
  function displayAuctions() {
    const auctionsContainer = document.getElementById("auction-list");
    const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
  
    let auctionsHtml = '';
    auctions.forEach((auction) => {
      auctionsHtml += `
        <div class="auction-item">
          <h4>Auction #${auction.id}</h4>
          <p>Number of Bids: ${auction.bids.length}</p>
          <p>Top Bid: ${auction.bids.length > 0 ? auction.bids[auction.bids.length - 1].value : 'N/A'}</p>
          <p>Active: ${Date.now() < auction.conclusionTime ? 'true' : 'false'}</p>
          <a href="view_auction.html?id=${auction.id}" class="view-button">View</a>
        </div>
      `;
    });

    auctionsContainer.innerHTML = auctionsHtml;
  }
  
  function loadAuction(auctionId) {
    const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    const auction = auctions.find((auction) => auction.id === parseInt(auctionId));
  
    if (!auction) {
      alert("Auction not found!");
      window.location.href = "index.html";
      return;
    }
  
    document.getElementById("auction-id").textContent = auction.id;
    document.getElementById("nft-type").textContent = auction.nftType;
    document.getElementById("smart-contract-address").textContent = auction.smartContractAddress;
    document.getElementById("conclusion-time").textContent = new Date(auction.conclusionTime).toLocaleString();
    document.getElementById("number-of-bids").textContent = auction.bids.length;
    document.getElementById("top-bid").textContent = auction.bids.length > 0 ? auction.bids[auction.bids.length - 1].value : "N/A";
    document.getElementById("active-status").textContent = Date.now() < auction.conclusionTime ? "true" : "false";
  
    const otherBidsContainer = document.getElementById("other-bids");
    let otherBidsHtml = '';
    auction.bids.forEach((bid, index) => {
      otherBidsHtml += `
        <div class="bid-item">
          <p>Bid #${index + 1}: ${bid.value} (Timestamp: ${new Date(bid.timestamp).toLocaleString()})</p>
        </div>
      `;
    });
  
    otherBidsContainer.innerHTML = otherBidsHtml;
  }
  