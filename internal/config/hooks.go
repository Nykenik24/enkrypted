package config

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/Nykenik24/enkrypted/internal/util"
	"github.com/gofiber/fiber/v3"
)

var os = util.GetRuntimeOS()

func randint(i, j int) int {
	if i < j {
		panic(errors.New("max less than min in randint"))
	}
	return rand.IntN(i-j) + j
}

func rainbow(text string) string {
	rainbowed := ""
	for _, c := range text {
		rainbowed = rainbowed + fmt.Sprintf("\x1b[%dm", randint(38, 31)) + string(c)
	}
	return rainbowed + "\x1b[0m"
}

func OnPreStartupMessageHook(sm *fiber.PreStartupMessageData) error {
	prefix := ""
	suffix := "\n------"
	banner := `                     __                                        __                      __
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
                                         $$$$$$/  $$/                                     `

	if os == util.Unix {
		banner = rainbow(banner)
	}

	sm.BannerHeader = prefix + banner + suffix

	sm.AddInfo("version", "Server version", "\t"+VERSION, 999)

	return nil
}
