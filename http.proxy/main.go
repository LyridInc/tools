package main

import (
	"github.com/joho/godotenv"
	entry "http.proxy/entry"
)

func main() {
	godotenv.Load()

	router := entry.Initialize()
	router.Run(":3000")
}
