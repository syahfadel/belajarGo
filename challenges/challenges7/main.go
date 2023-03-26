package main

import "challenges7/routers"

func main() {
	var PORT = ":4000"

	routers.StartServer().Run(PORT)
}
