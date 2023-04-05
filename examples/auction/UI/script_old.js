// script.js
document.addEventListener("DOMContentLoaded", function () {
    const createAuctionButton = document.getElementById("create-auction-btn");
    if (createAuctionButton) {
      createAuctionButton.addEventListener("click", createAuction);
    }
    displayAuctions();
  });
  
  function createAuction() {
    // Get input values
    const nftType = document.getElementById("nft-type").value;
    const smartContractAddress = document.getElementById("smart-contract-address").value;
    const conclusionTime = document.getElementById("conclusion-time").value;
  
    // Create auction object
    const auction = {
      id: Date.now(),
      nftType: nftType,
      smartContractAddress: smartContractAddress,
      conclusionTime: new Date(conclusionTime).getTime(),
      bids: [],
    };
  
    // Store the auction in localStorage
    let auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    auctions.push(auction);
    localStorage.setItem("auctions", JSON.stringify(auctions));
  
    // Redirect to the index page
    window.location.href = "index.html";
  }
  

  function displayAuctions() {
    const auctionsContainer = document.getElementById("auctions-container");
    if (!auctionsContainer) return;
  
    const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    auctionsContainer.innerHTML = "";
  
    auctions.forEach((auction) => {
      const isActive = Date.now() < auction.conclusionTime;
      const topBid = auction.bids.length ? Math.max(...auction.bids.map((bid) => bid.amount)) : "N/A";
  
      const auctionItem = document.createElement("div");
      auctionItem.className = "auction-item";
      auctionItem.innerHTML = `
        <div class="auction-info">
          <p><strong>Auction #:</strong> ${auction.id}</p>
          <p><strong>Number of Bids:</strong> ${auction.bids.length}</p>
          <p><strong>Top Bid:</strong> ${topBid}</p>
          <p><strong>Active:</strong> ${isActive ? "Yes" : "No"}</p>
        </div>
        <a href="view_auction.html?id=${auction.id}" class="view-btn">View</a>
      `;
      auctionsContainer.appendChild(auctionItem);
    });
  }
  

  function loadAuction() {
    const params = new URLSearchParams(window.location.search);
    const auctionId = params.get("id");
  
    if (!auctionId) return;
  
    const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    const auction = auctions.find((auction) => auction.id.toString() === auctionId);
  
    if (!auction) return;
  
    document.getElementById("nft-type").innerText = auction.nftType;
    document.getElementById("smart-contract-address").innerText = auction.smartContractAddress;
  }
  
  // Call loadAuction function when the DOM is fully loaded
  document.addEventListener("DOMContentLoaded", loadAuction);
