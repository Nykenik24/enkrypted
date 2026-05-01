package config

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/util"
	"github.com/gofiber/fiber/v3"
)

func rainbow(text string) string {
	rainbowed := ""
	for _, c := range text {
		rainbowed = rainbowed + fmt.Sprintf("\x1b[%dm", util.RandInt(38, 31)) + string(c)
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

	banner = rainbow(banner)
	sm.BannerHeader = prefix + banner + suffix

	sm.AddInfo("version", "Server version", "\t"+VERSION, 999)
	sm.AddInfo("os", "OS", "\t\t"+util.GetRuntimeOS().String(), 998)

	return nil
}

func OnPostStartupMessageHook(sm *fiber.PostStartupMessageData) error {
	fmt.Println()
	fmt.Println("Welcome to \x1b[34;1mthe enkrypted REPL\x1b[0m! You can manage \x1b[35;3myour enkrypted server\x1b[0m with it.")
	fmt.Println("Write \x1b[32mquit\x1b[0m to \x1b[91mclose the program\x1b[0m")
	fmt.Println("Write \x1b[32mhelp\x1b[0m to see all commands.")

	fmt.Printf("\n%s", REPL_PROMPT)
	return nil
}
