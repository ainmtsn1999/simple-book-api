package main

import "github.com/ainmtsn1999/simple-book-api/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
