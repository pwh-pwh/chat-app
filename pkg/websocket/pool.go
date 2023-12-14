package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Client     map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Client:     make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Client[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Client))
			for client, _ := range pool.Client {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "New User Joined...",
				})
			}
		case client := <-pool.Unregister:
			delete(pool.Client, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Client))
			for client, _ := range pool.Client {
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "User Disconnected...",
				})
			}
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Client {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}

		}
	}
}
