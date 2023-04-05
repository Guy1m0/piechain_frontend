// The auctions array will store all the auction data
let auctions = JSON.parse(localStorage.getItem('auctions')) || [];

// Initialize the dashboard by displaying the auctions
displayAuctions();


document.getElementById('create-auction-btn').addEventListener('click', () => {
  const auctionName = document.getElementById('auction-name').value;
  const conclusionTime = document.getElementById('conclusion-time-input').value;
  const blockchain = document.getElementById('blockchain-select').value;
  const contractAddress = document.getElementById('contract-address-input').value;
  const nftDetails = document.getElementById('nft-details').value;
  
  const newAuction = {
    auctionName,
    conclusionTime,
    blockchain,
    contractAddress,
    nftDetails,
    createdTime: new Date(),
    bids: [],
    active: true,
  };
  window.alert("Create!");
  createAuction(newAuction);
  window.location.href = 'index.html';
});

function displayAuctions() {
  // Retrieve the auction data from localStorage
  const auctions = JSON.parse(localStorage.getItem('auctions')) || [];

  // Find the auction-list element in the index.html file
  const auctionList = document.getElementById('auction-list');

  // Clear the auction list before adding new items
  auctionList.innerHTML = '';

  // Loop through the auctions array and create auction items
  auctions.forEach((auction, index) => {
    const auctionItem = document.createElement('div');
    auctionItem.className = 'auction-item';

    // Determine if the auction is active or not
    const isActive = (new Date(auction.conclusionTime) > new Date()) ? true : false;
    auctionItem.style.backgroundColor = isActive ? '#3d3d3d' : '#6b6b6b';

    auctionItem.innerHTML = `
      <h3>Auction #${index + 1}</h3>
      <p>Number of Bids: ${auction.bids.length}</p>
      <p>Top Bid: ${auction.bids.length ? Math.max(...auction.bids.map(bid => bid.value)) : 'N/A'}</p>
      <p>Active: ${isActive}</p>
      <button onclick="viewAuction(${index})" class="view-button">View</button>
    `;

    auctionList.appendChild(auctionItem);
  });
}



function resetAuctions() {
  localStorage.removeItem('auctions');
  auctions = [];
  displayAuctions();
}

function createAuction(auction) {
  auctions.push(auction);
  localStorage.setItem('auctions', JSON.stringify(auctions));
  displayAuctions();
}

function addExistingAuction() {
  // Add your logic for adding an existing auction
}

function viewAuction(index) {
  localStorage.setItem('currentAuctionIndex', index);
  window.location.href = 'view_auction.html';
}

// Function to load auction data for the view_auction.html page
function loadAuction() {
  const currentAuctionIndex = localStorage.getItem('currentAuctionIndex');
  const auction = auctions[currentAuctionIndex];

  document.getElementById('auction-title').innerText = `Auction #${parseInt(currentAuctionIndex) + 1}`;
  document.getElementById('number-of-bids').innerText = `Number of Bids: ${auction.bids.length}`;
  document.getElementById('top-bid').innerText = `Top Bid: ${auction.bids.length ? Math.max(...auction.bids.map(bid => bid.value)) : 'N/A'}`;
  document.getElementById('created-time').innerText = `Created Time: ${new Date(auction.createdTime).toLocaleString()}`;
  document.getElementById('conclusion-time').innerText = `Conclusion Time: ${new Date(auction.conclusionTime).toLocaleString()}`;
}




