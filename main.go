package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/rs/cors"
)

type message struct {
	Message string `json:"message"`
	Group   string `json:"group"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		//so.Join("chat")
		so.On("send:message", func(msg string) {
			log.Println(msg)
			var m message
			byt := []byte(msg)
			if json.Unmarshal(byt, &m); err != nil {
				log.Println(err)
				return
			}
			so.BroadcastTo(m.Group, "recieve:message", m.Message)
			return
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})

		so.On("subscribe", func(group string) string {
			so.Join(group)
			return "Successfully Joined"
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	mux.Handle("/socket.io/", server)
	// provide default cors to the mux
	handler := cors.Default().Handler(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	// decorate existing handler with cors functionality set in c
	handler = c.Handler(handler)

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
