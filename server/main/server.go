package main

import (
	"flag"
	libraryserver "github.com/jakekausler/Library-Organizer-2.0/server"
)

var (
	username = flag.String("username", "root", "Username of the local mysql database")
	password = flag.String("password", "", "Password of the local mysql database")
	database = flag.String("name", "library", "Name of the local mysql database")
	port     = flag.Int("port", 8181, "Port to run the server on")
)

func main() {
	flag.Parse()
	if *username == "" {
		panic("username cannot be empty!")
	}
	if *database == "" {
		panic("database cannot be empty!")
	}
	libraryserver.RunServer(*username, *password, *database, *port)
}
