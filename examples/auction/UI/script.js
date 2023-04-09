document.addEventListener("DOMContentLoaded", () => {
  //remove  || window.location.pathname === "/"
  if (window.location.pathname.includes("index.html")) {
    displayAuctions();
  } else if (window.location.pathname.includes("view_auction.html")) {
    loadAuction();
  }
  /*
  if (window.location.pathname.includes("login.html")) {
    alert("login page");
    setupLoginPage();
  }
  */
});

document.getElementById("make-bid-form").addEventListener("submit", (event) => {
  
  event.preventDefault(); // Prevent the default form submission behavior
  //window.alert("Catch submit?");
  addBid(); // Call the handleMakeBid function
  //window.alert("Adding ?")
});

document.getElementById("client-option").addEventListener("change", function () {
  const selectedOption = this.value;
  const displayText = document.getElementById("client-addr");

  const auctionId = new URLSearchParams(window.location.search).get("id");
  const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
  const auction = auctions.find((auction) => auction.id === parseInt(auctionId));
  window.alert(auction.assetId);


  if (selectedOption === "eth") {
    displayText.innerHTML = auction.EthAddr;
  } else if (selectedOption === "quo") {
    displayText.innerHTML = auction.QuorumAddr;
  } else {
    displayText.innerHTML = "";
  }
});

let bidderAddress = '';

//addAsset("asset1")
  function resetAuctions() {
    localStorage.removeItem('auctions');
    auctions = [];
    displayAuctions();
  }
  
  function createAuction() {
    //window.alert("create Auction");
    const assetId = document.getElementById("asset-id").value;


    //const blockType = document.getElementById("blockchain-type").value;
    //const nftType = document.getElementById("nft-type").value;
    //const blockAddress = document.getElementById("blockchain-address").value;
    const conclusionTime = document.getElementById("conclusion-time").value;
    const description = document.getElementById("auction-description").value;
    const checkedBoxes = document.querySelectorAll('.sup-BC:checked');
    const supportBlock = Array.from(checkedBoxes).map(checkbox => checkbox.value);

    //const response = startAuction(assetId);
    //const data = await response.json();
    //window.alert(response.data.owner);

    let auctions = JSON.parse(localStorage.getItem("auctions")) || [];  
    const auction = {
      id: auctions.length + 1,
      assetId: assetId,
      conclusionTime: new Date(conclusionTime).getTime(),
      description: description,
      supportBlock: supportBlock,
      //EthAddr: response.Data.EthAddr,
      //QuorumAddr: response.Data.QuorumAddr,
      topBid: 0,
      bids: [],
    };
  
    auctions.push(auction);
    localStorage.setItem("auctions", JSON.stringify(auctions));
  

    const response = addAsset(assetId);
    window.alert("Asset added by " + response.Owner);
    window.location.href = "index.html";

    //window.alert("Adding asset?");

    displayAuctions();


    // Testing addAsset

  }

  function loadFile() {
    const fileInput = document.getElementById("fileInput");
    const file = fileInput.files[0];
    const reader = new FileReader();
  
    reader.onload = function (event) {
      const contents = event.target.result;
      const json = JSON.parse(contents);
  
      bidderAddress = json.address;
      document.getElementById("bidder-address").textContent = "0x" + bidderAddress;
    };
  
    reader.readAsText(file);
  }

  function addBid() {
    //window.alert("Adding bid");
    const auctionId = new URLSearchParams(window.location.search).get("id");
    const auctions = JSON.parse(localStorage.getItem("auctions")) || [];
    const auction = auctions.find((auction) => auction.id === parseInt(auctionId));
    const index = auctions.findIndex((auction) => auction.id === parseInt(auctionId));
    //
    const bidder = "0x" + bidderAddress;
    const amount = parseInt(document.getElementById('bid-value').value,10);

    if (!bidder || !amount) {
      alert('Please load key file and fill in Bid Amount fields.');
      return;
    }
    const bid = {
      id: auction.bids.length + 1,
      bidder: bidder,
      amount: amount,
      timestamp: new Date().toISOString().replace('T', ' ').substr(0, 19)
    };

    //localStorage.setItem(auctionId, JSON.stringify(auction));
    //window.alert(auction.topBid+"?"+bid.amount);
    if (auction.topBid < bid.amount){
      auctions[index].topBid = bid.amount;
      auctions[index].topBidder = bidder;
    }
    //auctions[index].topBid = bid.amount;
    auctions[index].bids.push(bid);
    localStorage.setItem("auctions", JSON.stringify(auctions));
    document.getElementById("top-bid").textContent = auction.topBid;

    displayBid(bid);
    //document.getElementById("top-bid").textContent = auction.bids.length > 0 ? amount : "N/A";
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

          <div class="auction-item" data-active=${isActive}>
            <h4>Auction #${auction.id}</h4>
            <p>Number of Bids: ${auction.bids.length}</p>
            <p>Top Bid: ${auction.topBid}</p>
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
    //window.alert(document.getElementById("auction-name"));
    document.getElementById("asset-id").textContent = auction.assetId;
    document.getElementById("auction-id").textContent = auction.id;
    //document.getElementById("nft-type").textContent = auction.nftType;
    //document.getElementById("blockchain-address").textContent = auction.blockAddress;
    document.getElementById("auction-description").textContent = auction.description;
    document.getElementById("top-bid").textContent = auction.topBid;

    document.getElementById("conclusion-time").textContent = new Date(auction.conclusionTime).toLocaleString();
    document.getElementById("active-status").textContent = Date.now() < auction.conclusionTime ? "true" : "false";
    
    for (const bid of auction.bids) {
      displayBid(bid);
    }
    /*
    // Check if the user is logged in and their user type
    const userType = sessionStorage.getItem("userType");

    if (userType === "bidder" && auction.active) {
      document.getElementById("bid-section").style.display = "block";
    } else {
      document.getElementById("bid-section").style.display = "none";
    }
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
    //const nftName = document.getElementById("name").value;

    
  }
// ---------------------------------------------------


async function addAsset(assetID) {
  //window.alert("Try add asset");
  try {
      const response = await fetch(`http://localhost:6789/api/add-asset?asset_id=${assetID}`);

      console.log("Response object:", response);
      if (!response.ok) {
        throw new Error(`HTTP error ${response.status}`);
      }
      const data = await response.json();
      window.alert("Asset added:"+data.Owner);
      return data;
      //window.alert(data.owner);
  } catch (error) {
      window.alert(error);
      console.error("Error adding asset:", error);
  }
}

/*
async function startAuction(assetID) {
  //window.alert("Start Auction?");
  try {
      const response = await fetch(`http://localhost:8080/api/start-auction?asset_id=${assetID}`);
      const data = await response.json();
      window.alert("startAuction: "+data.EthAddr);
      return data;

  } catch (error) {
      window.alert(error);
      console.error("Error adding asset:", error);
  }
}
*/
async function startAuction(assetID) {
  const TIMEOUT_MS = 10000; // 10 seconds
  const controller = new AbortController();
  const { signal } = controller;
  setTimeout(() => controller.abort(), TIMEOUT_MS);

  try {
    const response = await fetch(`http://localhost:6789/api/start-auction?asset_id=${assetID}`, {
      method: 'GET',
      signal,
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    const data = await response.json();
    window.alert("startAuction: " + data.EthAddr);
    return data;
  } catch (error) {
    if (error.name === 'AbortError') {
      window.alert('Request timed out');
    } else {
      window.alert("other error:"+error);
      console.error("Error adding asset:", error);
    }
  }
}


