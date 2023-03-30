package main

import (
	"challenges8/connection"
	"challenges8/routers"
)

func main() {
	var PORT = ":4000"
	connection.Connect()
	routers.StartServer().Run(PORT)
}
