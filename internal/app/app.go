package enkrypt

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/Nykenik24/enkrypted/internal/config"
	"github.com/Nykenik24/enkrypted/internal/db"
	"github.com/Nykenik24/enkrypted/internal/routes"
	"github.com/gofiber/fiber/v3"
)

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

	logger := slog.New(&PrettyHandler{
		h: slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: level,
		}),
	})

	slog.SetDefault(logger)
}

type App struct {
	fiber *fiber.App
	repl  *REPL
}

func NewInstance() *App {
	return &App{
		fiber: fiber.New(),
		repl:  NewREPL(""),
	}
}

func (a *App) Start() {
	config.InitConfig()
	a.repl.Prompt = config.REPL_PROMPT
	a.repl.RegisterDefault()

	initLogger()

	a.fiber.Hooks().OnPreStartupMessage(config.OnPreStartupMessageHook)
	a.fiber.Hooks().OnPostStartupMessage(config.OnPostStartupMessageHook)

	db := db.GetInstance().Database
	database, err := db.DB()
	if err != nil {
		slog.Error("could not get db instance", "error", err)
		panic(1)
	}
	defer database.Close()

	routes.RegisterAll(a.fiber)

	go func() {
		if err := a.fiber.Listen(*config.ADDR); err != nil {
			slog.Error("error binding server to address", "address", config.ADDR, "error", err)
			os.Exit(1)
		}
	}()

	if err := a.repl.Run(a); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
