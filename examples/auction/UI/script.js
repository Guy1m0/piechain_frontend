document.addEventListener("DOMContentLoaded", () => {
  if (window.location.pathname.includes("index.html") || window.location.pathname === "/") {
    displayAuctions();
  } else if (window.location.pathname.includes("view_auction.html")) {
    loadAuction();
  }
});

document.getElementById("make-bid-form").addEventListener("submit", (event) => {
  
  event.preventDefault(); // Prevent the default form submission behavior
  //window.alert("Catch submit?");
  addBid(); // Call the handleMakeBid function
  //window.alert("Adding ?")
});

  function resetAuctions() {
    localStorage.removeItem('auctions');
    auctions = [];
    displayAuctions();
  }
  
  function createAuction() {
    const nftName = document.getElementById("name").value;
    const blockType = document.getElementById("blockchain-type").value;
    const nftType = document.getElementById("nft-type").value;
    const blockAddress = document.getElementById("blockchain-address").value;
    const conclusionTime = document.getElementById("conclusion-time").value;
    const description = document.getElementById("auction-description").value;

    const checkedBoxes = document.querySelectorAll('.sup-BC:checked');
    const supportBlock = Array.from(checkedBoxes).map(checkbox => checkbox.value);

    let auctions = JSON.parse(localStorage.getItem("auctions")) || [];  
    const auction = {
      id: auctions.length + 1,
      nftName: nftName,
      blockType: blockType,
      nftType: nftType,
      blockAddress: blockAddress,
      conclusionTime: new Date(conclusionTime).getTime(),
      description: description,
      supportBlock: supportBlock,
      bids: [],
    };
  

    auctions.push(auction);
    localStorage.setItem("auctions", JSON.stringify(auctions));
  
    window.location.href = "index.html";
    displayAuctions();
  }

  function addBid() {
    //window.alert("Adding bid");
    const auctionId = new URLSearchParams(window.location.search).get("id");
    const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    const auction = auctions.find((auction) => auction.id === parseInt(auctionId));
    const index = auctions.findIndex((auction) => auction.id === parseInt(auctionId));
    //
    const bidder = document.getElementById('bidder-address').value;
    const amount = document.getElementById('bid-value').value;

    if (!bidder || !amount) {
      alert('Please fill in both the Bidder Address and Bid Amount fields.');
      return;
    }
    const bid = {
      id: auction.bids.length + 1,
      bidder: bidder,
      amount: amount,
      timestamp: new Date().toISOString().replace('T', ' ').substr(0, 19)
    };
  
    auctions[index].bids.push(bid);
    localStorage.setItem("auctions", JSON.stringify(auctions));
    //localStorage.setItem(auctionId, JSON.stringify(auction));
  
    displayBid(bid);
    document.getElementById("top-bid").textContent = auction.bids.length > 0 ? amount : "N/A";
  }
  
  
  function displayAuctions() {
    const auctionsContainer = document.getElementById("auction-list");
    const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    // window.alert(auctions);
    let auctionsHtml = '';
    auctions.forEach((auction) => {
    // Determine if the auction is active or not
    const isActive = (new Date(auction.conclusionTime) > new Date()) ? true : false;
    // auctionsHtml.style.backgroundColor = isActive ? '#3d3d3d' : '#6b6b6b';

    // <form onsubmit="event.preventDefault(); loadAuction(${auction.id});">
      auctionsHtml += `

          <div class="auction-item">
            <h4>Auction #${auction.id}</h4>
            <p>Number of Bids: ${auction.bids.length}</p>
            <p>Top Bid: ${auction.bids.length > 0 ? auction.bids[auction.bids.length - 1].value : 'N/A'}</p>
            <p>Active: ${isActive}</p>
            <a href="view_auction.html?id=${auction.id}"><button class="view-btn" id="view-auction-page">View</button></a>
          </div>

      `;
    });

    auctionsContainer.innerHTML = auctionsHtml;
  }
  
  function loadAuction() {
    //window.alert("loadAuction");
    //window.location.href = "view_auction.html?id=" + auctionId;
    const auctionId = new URLSearchParams(window.location.search).get("id");
    const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    const auction = auctions.find((auction) => auction.id === parseInt(auctionId));

    //window.alert(auctionId);
    if (!auction) {
      alert("Auction not found!");
      window.location.href = "index.html";
      return;
    }
    /*
      const auction = {
      id: auctions.length + 1,
      nftName: nftName,
      blockType: blockType,
      nftType: nftType,
      blockAddress: blockAddress,
      conclusionTime: new Date(conclusionTime).getTime(),
      description: description,
      supportBlock: supportBlock,
      bids: [],
    }
    
    */
    
    //window.alert(document.getElementById("auction-name"));
    document.getElementById("auction-name").textContent = auction.nftName;
    document.getElementById("auction-id").textContent = auction.id;
    document.getElementById("nft-type").textContent = auction.nftType;
    document.getElementById("blockchain-address").textContent = auction.blockAddress;
    document.getElementById("auction-description").textContent = auction.description;
    document.getElementById("top-bid").textContent = auction.bids.length > 0 ? auction.bids[auction.bids.length - 1].value : "N/A";

    document.getElementById("conclusion-time").textContent = new Date(auction.conclusionTime).toLocaleString();
    document.getElementById("active-status").textContent = Date.now() < auction.conclusionTime ? "true" : "false";
    
    for (const bid of auction.bids) {
      displayBid(bid);
    }
    /*
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
    */
  }

  function displayBid(bid) {
    const bidElement = document.createElement('div');
    bidElement.className = 'bid';
  
    bidElement.innerHTML = `
      <p>Bid #${bid.id}</p>
      <p>Bidder: ${bid.bidder}</p>
      <p>Amount: ${bid.amount} ETH</p>
      <p>Timestamp: ${bid.timestamp}</p>
    `;
  
    document.getElementById('bid-list').appendChild(bidElement);
    
  }

  