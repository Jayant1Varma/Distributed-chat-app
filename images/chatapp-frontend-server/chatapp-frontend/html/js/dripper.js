// Initialization, called once upon page load:
openChannel = function() {
    
    var loc = window.location, new_uri;
    if (loc.protocol === "https:") {
        new_uri = "wss:";
    } else {
        new_uri = "ws:";
    }
    // Backend service is named 'chatapp-backend' in Kubernetes
    // and is exposed on port 14222
    // new_uri += "//chatapp-backend:14222/websocket"; //old code, might need update for dynamic localhost name. 

    // Use window.location.hostname which will be dynamic based on where the frontend is being accessed from
    new_uri += "//" + window.location.hostname + ":14222/websocket";

    websocket = new WebSocket(new_uri);

    websocket.onopen = onOpen;
    websocket.onmessage = onMessage;
    websocket.onerror = onError;
    websocket.onclose = onClose;
};

// Function to handle form submission without navigating to another page
postForm = function(oFormElement) {
    var message = {
        name: document.getElementById('name').value,
        email: document.getElementById('email').value,
        topic: document.getElementById('topic').value,
        content: document.getElementById('content').value,
        // Time will be set server-side
    };
    websocket.send(JSON.stringify(message));
    document.getElementById('content').value = ''; // Clear the message input field
    return false; // Prevent the default form submission
};

// Make it so ctrl-Enter can send a message
onKeyDown = function(e) {
    if (e.ctrlKey && (e.keyCode == 13 || e.keyCode == 77)) { // ctrl-Enter or ctrl-M
        postForm(document.getElementById('messageForm'));
        return false;
    }
    return true;
};

// Called every time a new message shows up from the server
onMessage = function(m) {
    var message = JSON.parse(m.data);
    var newDiv = document.createElement("div");
    newDiv.className = "message";

    var newB = document.createElement("b");
    newB.title = message.email; // Display email as a tooltip
    newB.textContent = `${message.name} (${message.topic}) [${message.time}]`;

    var newContent = document.createElement("div");
    newContent.textContent = message.content;

    newDiv.appendChild(newB);
    newDiv.appendChild(newContent);

    var bottom = document.getElementsByClassName("bottom")[0];
    document.body.insertBefore(newDiv, bottom); // Insert new message above the "bottom" div
};

// WebSocket connection opened
onOpen = function() {
    console.log("WebSocket connection opened.");
};

// WebSocket error handler
onError = function(e) {
    console.error("WebSocket encountered an error:", e);
};

// WebSocket connection closed
onClose = function() {
    console.log("WebSocket connection closed. Attempting to reconnect...");
    setTimeout(openChannel, 5000); // Try to reconnect after 5 seconds
};

window.addEventListener('DOMContentLoaded', (event) => {
    openChannel(); // Open WebSocket connection when the DOM is fully loaded
});









