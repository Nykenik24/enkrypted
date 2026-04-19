package app

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/handlers"
	"github.com/Nykenik24/enkrypted/internal/server"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "web/home.html")
}

func Start() {
	flag.Parse()

	srv := server.NewServer()
	go srv.Run()

	http.HandleFunc("/", serveHome)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("WS: %s", r.RemoteAddr)
		c := ws.ServeWebsockets(srv, w, r)
		c.AddHandler(event.PingEvent, handlers.PingEV)
		c.AddHandler(event.MessageEvent, handlers.MessageEV)
	})

	httpServer := &http.Server{
		Addr:         *addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("server listening on %s", *addr)
	log.Fatal(httpServer.ListenAndServe())
}
