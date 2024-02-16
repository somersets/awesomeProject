package userChatService

import (
	"awesomeProject/domain"
	"fmt"
	"strconv"
)

// Hub maintains the set of active Clients and broadcasts messages to the
// Clients.
type Hub struct {
	// Registered Clients.
	Clients map[string]map[*Client]bool

	// Inbound messages from the Clients.
	Broadcast chan Message

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	Unregister chan *Client
}

type Message struct {
	Type      string                               `json:"type"`
	Sender    domain.SenderAndRecipientInfoMessage `json:"sender"`
	Recipient domain.SenderAndRecipientInfoMessage `json:"recipient"`
	Content   string                               `json:"content"`
	ID        string                               `json:"id"`
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]map[*Client]bool),
	}
}

// Run example usage
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.RegisterNewClient(client)
		case client := <-h.Unregister:
			h.RemoveClient(client)
		case message := <-h.Broadcast:
			h.HandleMessage(message)
		}
	}
}

func (h *Hub) HandleMessage(message Message) {
	//Check if the message is a type of "message"
	if message.Type == "message" {
		sclients := h.Clients[strconv.Itoa(message.Sender.ID)]
		for client := range sclients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.Clients[strconv.Itoa(message.Sender.ID)], client)
			}
		}

		clients := h.Clients[strconv.Itoa(message.Recipient.ID)]
		for client := range clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.Clients[strconv.Itoa(message.Recipient.ID)], client)
			}
		}
	}

	//Check if the message is a type of "notification"
	/*if message.Type == "notification" {
		fmt.Println("Notification: ", message.Content)
		Clients := h.Clients[message.Recipient]
		for client := range Clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.Clients[message.Recipient], client)
			}
		}
	}*/

}

// RegisterNewClient function check if room exists and if not create it and add client to it
func (h *Hub) RegisterNewClient(client *Client) {
	fmt.Println(client.ID)
	connections := h.Clients[client.ID]
	if connections == nil {
		connections = make(map[*Client]bool)
		h.Clients[client.ID] = connections
	}
	h.Clients[client.ID][client] = true

	fmt.Println("Size of Clients: ", len(h.Clients[client.ID]))
}

// RemoveClient function to remove client from room
func (h *Hub) RemoveClient(client *Client) {
	if _, ok := h.Clients[client.ID]; ok {
		delete(h.Clients[client.ID], client)
		close(client.send)
		fmt.Println("Removed client")
	}
}
