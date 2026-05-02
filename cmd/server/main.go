package main

import (
	"flag"

	enkrypt "github.com/Nykenik24/enkrypted/internal/app"
)

func main() {
	flag.Parse()

	app := enkrypt.NewInstance()
	app.Start()
}
