package main

import (
	"flag"
    "os"
    "strconv"
	libraryserver "github.com/jakekausler/Library-Organizer-2.0/server"
)

const (
    HostEnv = "HOST"
    UsernameEnv = "USERNAME"
    PasswordEnv = "PASSWORD"
    DatabaseEnv = "DATABASE"
    AppPortEnv = "APPPORT"
    MySqlPortEnv = "MYSQLPORT"
)

func main() {
    appportValue, err := strconv.Atoi(os.Getenv(AppPortEnv))
    if err != nil {
        panic("APPPORT Environment variable must be an integer.")
    }
    mysqlportValue, err := strconv.Atoi(os.Getenv(MySqlPortEnv))
    if err != nil {
        panic("MYSQLPORT Environment variable must be an integer.")
    }

	flag.Parse()

    host	 := flag.String("host", os.Getenv(HostEnv), "Host of the mysql database")
	username := flag.String("username", os.Getenv(UsernameEnv), "Username of the local mysql database")
	password := flag.String("password", os.Getenv(PasswordEnv), "Password of the local mysql database")
	database := flag.String("name", os.Getenv(DatabaseEnv), "Name of the local mysql database")
	appport     := flag.Int("appport", appportValue, "Port to run the server on")
	mysqlport     := flag.Int("mysqlport", mysqlportValue, "Port to run the server on")

	if *username == "" {
		panic("username cannot be empty!")
	}
	if *database == "" {
		panic("database cannot be empty!")
	}
	libraryserver.RunServer(*host, *username, *password, *database, *appport, *mysqlport)
}
