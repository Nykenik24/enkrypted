package app

import (
	"flag"
	"log"
	"net/http"
	"time"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
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
	http.ServeFile(w, r, "web/index.html")
}

var default_handlers = map[string]ws.EventHandler{
	builtin_ev.MessageEventKind.String(): &handlers.MessageHandler{},

	builtin_ev.IdentifyEventKind.String(): &handlers.IdentifyHandler{},
	builtin_ev.ConnectEventKind.String():  &handlers.ConnectHandler{},

	builtin_ev.JoinRoomEventKind.String():  &handlers.JoinRoomHandler{},
	builtin_ev.LeaveRoomEventKind.String(): &handlers.LeaveRoomHandler{},

	builtin_ev.CreateRoomEventKind.String(): &handlers.CreateRoomHandler{},
	builtin_ev.RemoveRoomEventKind.String(): &handlers.RemoveRoomHandler{},
}

func Start() {
	flag.Parse()

	srv := server.NewServer()
	hub := srv.Hub

	hub.AddHandlers(default_handlers)

	go hub.Run()

	// http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	// http.HandleFunc("/", serveHome)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("WS: %s", r.RemoteAddr)

		c := ws.ServeWebsockets(hub, w, r)

		if c == nil {
			log.Println("ws upgrade failed")
			return
		}

		connect := builtin_ev.NewConnectEvent(c.GetID())

		hub.Emit(c, builtin_ev.Generic(connect))
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
