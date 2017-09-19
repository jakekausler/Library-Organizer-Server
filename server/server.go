package main

import "flag"

var (
	username = flag.String("username", "root", "Username of the local mysql database")
	password = flag.String("password", "", "Password of the local mysql database")
	database = flag.String("name", "library", "Name of the local mysql database")
)

func main() {
	flag.Parse()
	if *username == "" {
		panic("username cannot be empty!")
	}
	if *database == "" {
		panic("database cannot be empty!")
	}
	RunServer(*username, *password, *database)
}
