package enkrypt

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/Nykenik24/enkrypted/internal/handlers"
	"github.com/Nykenik24/enkrypted/internal/server"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var addr = flag.String("addr", ":8080", "http service address")
var passwd = flag.String("password", "hunter2", "password")

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "web/index.html")
}

type App struct {
	Handlers map[string]ws.EventHandler
	Server   *server.Server
}

var default_handlers = map[string]ws.EventHandler{
	handlers.MessageEventKind.String(): &handlers.MessageHandler{},

	handlers.IdentifyEventKind.String():   &handlers.IdentifyHandler{},
	handlers.ConnectEventKind.String():    &handlers.ConnectHandler{},
	handlers.GetClientsEventKind.String(): &handlers.GetClientsHandler{},

	handlers.JoinRoomEventKind.String():  &handlers.JoinRoomHandler{},
	handlers.LeaveRoomEventKind.String(): &handlers.LeaveRoomHandler{},

	handlers.CreateRoomEventKind.String(): &handlers.CreateRoomHandler{},
	handlers.RemoveRoomEventKind.String(): &handlers.RemoveRoomHandler{},
	handlers.GetRoomsEventKind.String():   &handlers.GetRoomsHandler{},
}

func NewApp() *App {
	return &App{
		Handlers: default_handlers,
	}
}

func (app *App) ServeHTTP() {
	hub := app.Server.Hub

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("WS: %s", r.RemoteAddr)

		c := ws.ServeWebsockets(hub, w, r)

		if c == nil {
			log.Println("ws upgrade failed")
			return
		}

		connect := handlers.NewConnectEvent(c.GetID())

		hub.Emit(c, handlers.Base(connect))
	})

	httpServer := &http.Server{
		Addr:         *addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("server listening on %s", *addr)
	log.Fatal(httpServer.ListenAndServe())

	// http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	// http.HandleFunc("/", serveHome)
}

func (app *App) StartHub() {
	flag.Parse()

	app.Server = server.NewServer(server.Config(*passwd))
	hub := app.Server.Hub

	hub.AddHandlers(app.Handlers)

	go hub.Run()
}
