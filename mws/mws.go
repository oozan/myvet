package mws

import (
	"encoding/json"
	"log"
	"time"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	msgRcvd func(client string, message string)
}

// NewHub creates a new websocket encapsulation hub.
func NewHub(msgRcvd func(client string, message string)) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		msgRcvd:    msgRcvd,
	}
}

// Run the hub.
func (h *Hub) Run() {
	// ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case client := <-h.register:
			log.Println("Registered client")
			log.Println(client)

			for client := range h.Clients {
				select {
				case client.Send <- []byte(client.ID):
					log.Println("sending message")
					log.Println(string([]byte(client.ID)))
				default:
					log.Println("client broken")
					close(client.Send)
					delete(h.Clients, client)
				}
			}

		case client := <-h.unregister:
			if _, ok := h.Clients[client]; ok {
				log.Println("client unregistered itself")
				delete(h.Clients, client)
				close(client.Send)
				pl, err := json.Marshal(ConnUser{ConnectionID: client.ID, Connected: client.Connected})
				if err != nil {
					log.Fatal(err)
				}
				log.Println("------------------------------------------")
				log.Println(string(pl))
				announcement, err := json.Marshal(string(pl))
				if err != nil {
					log.Fatal(err)
				}
				log.Println("broadcasting: " + string(announcement))
				for client := range h.Clients {
					select {
					case client.Send <- announcement:
						log.Println("sending message")
						log.Println(string(announcement))
					default:
						log.Println("client borken")
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}

		case message := <-h.broadcast:
			log.Println("got a message to bcast")
			for client := range h.Clients {
				select {
				case client.Send <- message:
					log.Println("sending message")
					log.Println(string(message))
				default:
					log.Println("client borken")
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

// GetConnectedUsers returns a slice of ConnUsers representing all connected users.
func (h *Hub) GetConnectedUsers() []ConnUser {
	var retArr []ConnUser
	for client := range h.Clients {
		retArr = append(retArr, ConnUser{ConnectionID: client.ID, Connected: client.Connected})
	}
	return retArr
}

// ConnUser represents one connection to the websocket by user.
type ConnUser struct {
	ConnectionID string            `json:"connectionId"`
	Connected    time.Time         `json:"connected"`
	Attrs        map[string]string `json:"attrs"`
}
