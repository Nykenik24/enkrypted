package enkrypt

import (
	"flag"
	"log/slog"
	"os"
	"strings"

	"github.com/Nykenik24/enkrypted/internal/config"
	"github.com/Nykenik24/enkrypted/internal/db"
	"github.com/Nykenik24/enkrypted/internal/routes"
	"github.com/gofiber/fiber/v3"
)

const addr = ":8080"

var log *slog.Logger

var prettyLog = flag.Bool("pretty", false, "use pretty logging instead of the usual JSON logging")

func initLogger() {
	levelStr := os.Getenv("LOG_LEVEL")
	if levelStr == "" {
		levelStr = "info"
	}

	var level slog.Level
	switch strings.ToLower(levelStr) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Configure the global default logger
	if *prettyLog {
		logger := slog.New(&PrettyHandler{
			h: slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
				Level: level,
			}),
		})

		slog.SetDefault(logger)
	} else {
		slog.SetDefault(slog.New(
			slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
				Level: level,
			}),
		))
	}
}

func Start() {
	initLogger()

	app := fiber.New()

	app.Hooks().OnPreStartupMessage(config.OnPreStartupMessageHook)

	db := db.GetInstance().Database
	database, err := db.DB()
	if err != nil {
		slog.Error("could not get db instance", "error", err)
		panic(1)
	}
	defer database.Close()

	routes.RegisterAll(app)

	err = app.Listen(addr)
	if err != nil {
		slog.Error("error binding server to address", "address", addr, "error", err)
	}
}
