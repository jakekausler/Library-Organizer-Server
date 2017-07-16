package main

import (
	"flag"
	"log"
	"database/sql"
	"os"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)

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

var (
	db *sql.DB
	logger = log.New(os.Stderr, "log: ", log.LstdFlags | log.Lshortfile)
)

//RunServer runs the library server
func RunServer(username, password, database string) {
	logger.Printf("Creating the database")

	var err error
	// Create sql.DB
	db, err = sql.Open("mysql", "root:@/library?parseTime=true")
	// db, _ = sql.Open("mysql", "root:@/library?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	logger.Printf("Pinging the database")
	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	query := "SELECT id, pass FROM library_members WHERE id IN (1,2,3)"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var id int64
		var password string
		err = rows.Scan(&id, &password)
		if err != nil {
			panic(err.Error())
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Error: %+v", err)
		}
		query = "UPDATE library_members SET pass=? WHERE id=?"
		_, err = db.Exec(query, hash, id)
		if err != nil {
			log.Fatalf("Error: %+v", err)
		}
	}

	logger.Printf("Closing")
}