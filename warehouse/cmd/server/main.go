package main

import "../../server"
import "os"

func main() {
	instance := server.New("", os.Getenv("PORT"))
	instance.Start()
}
