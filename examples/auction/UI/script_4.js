displayAuctions();

// Load auctions from localStorage or set an empty array if none exists
function getAuctions() {
    return JSON.parse(localStorage.getItem("auctions")) || [];
  }
  
  // Save the auctions array to localStorage
  function saveAuctions(auctions) {
    localStorage.setItem("auctions", JSON.stringify(auctions));
  }
  
  // Create a new auction
  function createAuction(auction) {
    const auctions = getAuctions();
    auctions.push(auction);
    saveAuctions(auctions);
  }
  
  // Load an auction by ID
  function loadAuction(id) {
    const auctions = getAuctions();
    return auctions.find(auction => auction.id === id);
  }
  
  // Make a bid on an auction
  function makeBid(id, bid) {
    const auctions = getAuctions();
    const auction = auctions.find(auction => auction.id === id);
  
    if (auction) {
      auction.bids.push(bid);
      saveAuctions(auctions);
    }
  }
  
  // Display auctions on the dashboard
  function displayAuctions() {
    const auctions = getAuctions();
    const auctionList = document.getElementById("auction-list");
    auctionList.innerHTML = '';
  
    auctions.forEach((auction) => {
      const auctionElement = document.createElement("div");
      auctionElement.className = "auction-item";
      auctionElement.innerHTML = `
        <h3>Auction #${auction.id}: ${auction.name}</h3>
        <p>Number of Bids: ${auction.bids.length}</p>
        <p>Top Bid: ${auction.bids.length ? Math.max(...auction.bids.map(bid => bid.value)) : 'N/A'}</p>
        <p>Active: ${auction.active ? 'Yes' : 'No'}</p>
        <button class="view-auction" data-auction-id="${auction.id}">View</button>
      `;
      auctionList.appendChild(auctionElement);
    });
  
    document.querySelectorAll(".view-auction").forEach((viewBtn) => {
      viewBtn.addEventListener("click", (e) => {
        const auctionId = e.target.dataset.auctionId;
        window.location.href = `view_auction.html?id=${auctionId}`;
      });
    });
  }
  

  