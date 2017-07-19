package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)
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
	http.HandleFunc("/", GetLoginPageHandler)
	http.HandleFunc("/home", GetHomePageHandler)
	http.HandleFunc("/books", GetBooksHandler)
	http.HandleFunc("/publishers", GetPublishersHandler)
	http.HandleFunc("/cities", GetCitiesHandler)
	http.HandleFunc("/states", GetStatesHandler)
	http.HandleFunc("/countries", GetCountriesHandler)
	http.HandleFunc("/series", GetSeriesHandler)
	http.HandleFunc("/formats", GetFormatsHandler)
	http.HandleFunc("/languages", GetLanguagesHandler)
	http.HandleFunc("/roles", GetRolesHandler)
	http.HandleFunc("/deweys", GetDeweysHandler)
	http.HandleFunc("/savebook", SaveBookHandler)
	http.HandleFunc("/exportbooks", ExportBooksHandler)
	http.HandleFunc("/exportauthors", ExportAuthorsHandler)
	http.HandleFunc("/import", ImportBooksHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/stats", GetStatsHandler)
	http.HandleFunc("/cases", GetCasesHandler)
	http.HandleFunc("/dimensions", GetDimensionsHandler)
	http.HandleFunc("/deletebook", DeleteBookHandler)
	http.HandleFunc("/libraries", GetLibrariesHandler)
	http.HandleFunc("/username", GetUsernameHandler)
	http.HandleFunc("/reset", ResetPasswordHandler)
	http.HandleFunc("/settings", GetSettingsHandler)
	http.HandleFunc("/updatesettings", UpdateSettingsHandler)
	http.HandleFunc("/getsetting", GetSettingHandler)
	http.HandleFunc("/savecases", SaveCasesHandler)
	http.HandleFunc("/addbreak", AddBreakHandler)
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./../web/"))))
	logger.Printf("Listening on port 8181")
	http.ListenAndServe(":8181", nil)
	logger.Printf("Closing")
}

//GetLoginPageHandler gets the login page
func GetLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if Registered(r) {
		http.Redirect(w, r, "/home", 301)
	} else {
		http.ServeFile(w, r, "../web/app/unregistered/index.html")
	}
}

//GetHomePageHandler gets the home page
func GetHomePageHandler(w http.ResponseWriter, r *http.Request) {
	if Registered(r) {
		http.ServeFile(w, r, "../web/app/main/index.html")
	} else {
		http.Redirect(w, r, "/", 301)
	}
}
