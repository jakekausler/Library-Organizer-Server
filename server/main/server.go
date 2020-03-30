package main

import (
	"flag"
	libraryserver "github.com/jakekausler/Library-Organizer-2.0/server"
)

var (
	host	 = flag.String("host", "localhost", "Host of the mysql database")
	username = flag.String("username", "root", "Username of the local mysql database")
	password = flag.String("password", "", "Password of the local mysql database")
	database = flag.String("name", "library", "Name of the local mysql database")
	appport     = flag.Int("appport", 8181, "Port to run the server on")
	mysqlport     = flag.Int("mysqlport", 3306, "Port to run the server on")
)

func main() {
	flag.Parse()
	if *username == "" {
		panic("username cannot be empty!")
	}
	if *database == "" {
		panic("database cannot be empty!")
	}
	libraryserver.RunServer(*host, *username, *password, *database, *appport, *mysqlport)
}
