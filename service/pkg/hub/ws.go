package hub

import (
	log "github.com/sirupsen/logrus"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func New() *Hub {
	return &Hub{
		Broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			log.Debug("Client registered")
			h.clients[client] = true
		case client := <-h.unregister:
			log.Debug("Client unregistered")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			log.Debug("Message to be broadcasted")
			for client := range h.clients {
				select {
				case client.send <- message:
					log.Debug("Sending message to the client")
				default:
					log.Debug("Removing client from the hub")
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
