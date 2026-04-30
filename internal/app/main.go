package app

import (
	"log"

	"github.com/Nykenik24/enkrypted/internal/db"
	"github.com/Nykenik24/enkrypted/internal/routes"
	"github.com/gofiber/fiber/v3"
)

const VERSION = "v0.0.1"

func Start() {
	app := fiber.New()

	app.Hooks().OnPreStartupMessage(func(sm *fiber.PreStartupMessageData) error {
		sm.BannerHeader = `                     __                                        __                      __
                    /  |                                      /  |                    /  |
  ______   _______  $$ |   __   ______   __    __   ______   _$$ |_     ______    ____$$ |
 /      \ /       \ $$ |  /  | /      \ /  |  /  | /      \ / $$   |   /      \  /    $$ |
/$$$$$$  |$$$$$$$  |$$ |_/$$/ /$$$$$$  |$$ |  $$ |/$$$$$$  |$$$$$$/   /$$$$$$  |/$$$$$$$ |
$$    $$ |$$ |  $$ |$$   $$<  $$ |  $$/ $$ |  $$ |$$ |  $$ |  $$ | __ $$    $$ |$$ |  $$ |
$$$$$$$$/ $$ |  $$ |$$$$$$  \ $$ |      $$ \__$$ |$$ |__$$ |  $$ |/  |$$$$$$$$/ $$ \__$$ |
$$       |$$ |  $$ |$$ | $$  |$$ |      $$    $$ |$$    $$/   $$  $$/ $$       |$$    $$ |
 $$$$$$$/ $$/   $$/ $$/   $$/ $$/        $$$$$$$ |$$$$$$$/     $$$$/   $$$$$$$/  $$$$$$$/
                                        /  \__$$ |$$ |
                                        $$    $$/ $$ |
                                         $$$$$$/  $$/                                     ` + VERSION + "\n------"

		return nil
	})

	db := db.GetInstance().Database
	database, err := db.DB()
	if err != nil {
		log.Fatalf("could not get db instance: %s", err.Error())
	}
	defer database.Close()

	routes.RegisterAll(app)

	log.Fatal(app.Listen(":8080"))
}
