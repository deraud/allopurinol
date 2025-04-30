package main

import (
	_ "github.com/joho/godotenv/autoload"
	"healthcare-allopurinol/server"
)

func main() {
	server.StartServer()
}
