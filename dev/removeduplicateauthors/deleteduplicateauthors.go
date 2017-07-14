package main

import (
	"flag"
	"log"
	"database/sql"
	"os"
	"strings"

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

	query := "SELECT GROUP_CONCAT(PersonID) ids FROM persons GROUP BY FirstName,MiddleNames,LastName HAVING COUNT(*) > 1"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var idstring string
		err = rows.Scan(&idstring)
		if err != nil {
			panic(err.Error())
		}
		ids := strings.Split(idstring, ",")
		keep := ids[0]
		throw := strings.Join(ids[1:], ",")
		query = "DELETE FROM persons WHERE personid in ("+throw+")"
		logger.Printf(query)
		_, err = db.Exec(query)
		if err != nil {
			panic(err.Error())
		}
		query = "UPDATE Written_by SET AuthorID=? WHERE AUTHORID IN ("+throw+")"
		logger.Printf(query)
		_, err = db.Exec(query, keep)
		if err != nil {
			panic(err.Error())
		}
	}

	logger.Printf("Closing")
}