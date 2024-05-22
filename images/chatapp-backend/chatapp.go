package main

import (
	"context"       // For managing lifecycle: cancellation, deadlines, blah-blah.
	"encoding/json" // For encoding and decoding messages to JSON.
	"flag"          // For parsing command-line arguments.
	"log"           // For logging messages and errors.
	"net/http"      // Provides HTTP client and server implementations.
	"time"          // For timestamping messages.

	"github.com/go-redis/redis/v8" // Redis client for Go.
	"github.com/gorilla/websocket" // WebSocket protocol implementation.
)

// Configuration variables.
// var addr = flag.String("addr", "0.0.0.0:8080", "HTTP service address")
var addr = flag.String("addr", "0.0.0.0:14222", "HTTP service address")
var ctx = context.Background()

// WebSocket connection upgrade configuration.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin.
	},
}

var redisAddr = flag.String("redis-addr", "redis:6379", "Redis server address")

// Initialize Redis client.
var rdb = redis.NewClient(&redis.Options{
	Addr:     *redisAddr, // Address of the Redis server.
	Password: "",         // No password is set.
	DB:       0,          // Use default DB.
})

// Message struct represents the structure of a chat message.
type Message struct {
	Name    string `json:"name"`    // Sender's name.
	Email   string `json:"email"`   // Sender's email.
	Time    string `json:"time"`    // Time message was sent.
	Topic   string `json:"topic"`   // Topic of the message.
	Content string `json:"content"` // Content of the message.
}

// Clients tracks all connected WebSocket clients.
var clients = make(map[*websocket.Conn]bool)

func main() {
	flag.Parse()
	// http.HandleFunc("/ws", handleConnections) // Handle WebSocket connections.
	http.HandleFunc("/websocket", handleConnections) // Updated endpoint
	go subscribeToMessages()                         // Subscribe to Redis messages.

	// Start the WebSocket server.
	log.Printf("WebSocket server started on %s\n", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// handleConnections upgrades HTTP to WebSocket and handles incoming messages.
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Register client.
	clients[ws] = true

	// Read messages from WebSocket and publish to Redis.
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws) // Remove client on error.
			break
		}
		msg.Time = time.Now().Format(time.RFC3339) // Timestamp message.
		publishMessage(msg)                        // Publish message via Redis.
	}
}

// publishMessage marshals the message to JSON and publishes it to Redis.
func publishMessage(msg Message) {
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return
	}

	// Publish message to Redis channel "chat_messages".
	err = rdb.Publish(ctx, "chat_messages", msgJSON).Err()
	if err != nil {
		log.Printf("Error publishing message to Redis: %v", err)
	}
}

// subscribeToMessages subscribes to the Redis channel and broadcasts messages.
func subscribeToMessages() {
	pubsub := rdb.Subscribe(ctx, "chat_messages") // Subscribe to Redis channel.
	defer pubsub.Close()
	ch := pubsub.Channel()

	// Listen for messages from Redis and broadcast to WebSocket clients.
	for msg := range ch {
		var message Message
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}
		broadcastMessage(message) // Broadcast message to clients.
	}
}

// broadcastMessage sends messages to all connected WebSocket clients.
func broadcastMessage(msg Message) {
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("Websocket error: %s", err)
			client.Close()          // Close connection on error.
			delete(clients, client) // Remove client from the tracking map.
		}
	}
}
