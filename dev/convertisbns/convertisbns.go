package main

import (
	"flag"
	"log"
	"database/sql"
	"os"
	"strconv"

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
	query := "SELECT bookid, isbn FROM books"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var id int64
		var isbn string
		err = rows.Scan(&id, &isbn)
		if err != nil {
			panic(err.Error())
		}
		if len(isbn) == 10 {
			if isValid(isbn) {
				isbn = isbn10to13(isbn)
				query = "UPDATE books SET isbn=? WHERE bookid=?"
				_, err = db.Exec(query, isbn, id)
				if err != nil {
					log.Fatalf("Error: %+v", err)
				}
			} else {
				logger.Printf("Deleting invalid isbn: %v", isbn)
				query = "UPDATE books SET isbn=? WHERE bookid=?"
				_, err = db.Exec(query, "", id)
				if err != nil {
					log.Fatalf("Error: %+v", err)
				}				
			}
		}
	}
	logger.Printf("Closing")
}

func isValid(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	var sum int
	var multiply int = 10
	for i, v := range isbn {
		digitString := string(v)
		if i == 9 && digitString == "X" {
			digitString = "10"
		}
		digit, err := strconv.Atoi(digitString)
		if err != nil {
			panic(err)
		} else {
			sum = sum + (multiply * digit)
			multiply--
		}
	}
	return sum%11 == 0
}

func isbn10to13(isbn10 string) string {
	Isbn13Prefix := "978"
	isbn := Isbn13Prefix + string(isbn10[0:9])
	return isbn + CalculateCheckSum(isbn)
}

func CalculateCheckSum(isbnNoCheckSum string) string {
	isbnLen := len(isbnNoCheckSum)
	var sum = 0
	// ISBN 10
	if isbnLen == 9 {
		for i := 0; i < 9; i++ {
			toInt, _ := strconv.Atoi(string(isbnNoCheckSum[i]))
			sum += (10 - i) * toInt
		}
		countResult := 11 - sum%11
		if countResult == 10 {
			return "X"
		} else if countResult == 11 {
			return "0"
		} else {
			return strconv.Itoa(countResult)
		}
		// ISBN 13
	} else {
		for i := 0; i < 12; i++ {
			toInt, _ := strconv.Atoi(string(isbnNoCheckSum[i]))
			if i%2 == 0 {
				sum += toInt
			} else {

				sum += toInt * 3
			}
		}
		countResult := 10 - sum%10
		if countResult == 10 {
			return "0"
		} else {
			return strconv.Itoa(countResult)
		}
	}
}