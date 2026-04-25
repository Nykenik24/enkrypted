package main

import enkrypt "github.com/Nykenik24/enkrypted/internal/app"

func main() {
	app := enkrypt.NewApp()

	app.StartHub()
	app.ServeHTTP()
}
