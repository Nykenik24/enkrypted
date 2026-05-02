package config

import "flag"

const VERSION = "v0.0.1"

var ADDR = flag.String("address", ":8080", "the address where the server is hosted")
var REPL_PROMPT = "\x1b[92menkrypted>\x1b[0m "

func InitConfig() {
	if ADDR == nil {
		*ADDR = ":8080"
	}
}
