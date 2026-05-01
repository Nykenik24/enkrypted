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

var funFacts = []string{
	"The banner's colors are randomly chosen.",
	"Enkrypted's backend is fully written in Go.",
	"You can use the --pretty flag for friendlier logs.",
	"We recommend you use a tool such as Hoppscotch or Bruno to test the API!",
}

func OnPostStartupMessageHook(sm *fiber.PostStartupMessageData) error {
	funFacts = append(funFacts, fmt.Sprintf("This message is random, and you just got a 1/%d chance!", len(funFacts)+1))

	fmt.Println()
	fmt.Println("Welcome to \x1b[34;1mthe enkrypted REPL\x1b[0m! You can manage your\x1b[35;3menkrypted server\x1b[0m with it.")
	fmt.Printf("\x1b[1mFun fact\x1b[0m: \x1b[90;3m%s\x1b[0m\n", util.Choice(funFacts))
	fmt.Println()
	fmt.Println("Write \x1b[32mquit\x1b[0m \x1b[3mor\x1b[0m \x1b[32mq\x1b[0m to \x1b[91mclose the program\x1b[0m")
	fmt.Println("Write \x1b[32mhelp\x1b[0m to see all commands.")

	fmt.Printf("\n%s", REPL_PROMPT)
	return nil
}
