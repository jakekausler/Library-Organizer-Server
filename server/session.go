package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"
)

var (
	db     *sql.DB
	logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)
	store = sessions.NewCookieStore([]byte("Session"))
)

//RunServer runs the library server
func RunServer(username, password, database string) {
	logger.Printf("Creating the database")
	var err error
	// Create sql.DB
	db, err = sql.Open("mysql", "root:@/library?parseTime=true")
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
	logger.Printf("Opening the log")
	logFile, err := os.OpenFile("../logs/server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", RouteByAuthenticate)
	r.HandleFunc("/home", GetHomePageHandler)
	r.HandleFunc("/unregistered", GetUnregisteredPageHandler)
	r.HandleFunc("/books", GetBooksHandler).Methods("GET")
	r.HandleFunc("/books", AddBookHandler).Methods("POST")
	r.HandleFunc("/books", SaveBookHandler).Methods("PUT")
	r.HandleFunc("/books", DeleteBookHandler).Methods("DELETE")
	r.HandleFunc("/books/checkout", CheckoutBookHandler).Methods("PUT")
	r.HandleFunc("/books/checkin", CheckinBookHandler).Methods("PUT")
	r.HandleFunc("/books/books", ExportBooksHandler).Methods("GET")
	r.HandleFunc("/books/contributors", ExportAuthorsHandler).Methods("GET")
	r.HandleFunc("/books/books", ImportBooksHandler).Methods("POST")
	r.HandleFunc("/books/{bookid}", GetBookHandler).Methods("GET")
	r.HandleFunc("/books/{bookid}/ratings", GetRatingsHandler).Methods("GET")
	r.HandleFunc("/books/{bookid}/ratings", AddRatingHandler).Methods("PUT")
	r.HandleFunc("/books/{bookid}/reviews", GetReviewsHandler).Methods("GET")
	r.HandleFunc("/books/{bookid}/reviews", AddReviewHandler).Methods("PUT")
	r.HandleFunc("/information/statistics", GetStatsHandler).Methods("GET")
	r.HandleFunc("/information/dimensions", GetDimensionsHandler).Methods("GET")
	r.HandleFunc("/information/publishers", GetPublishersHandler).Methods("GET")
	r.HandleFunc("/information/cities", GetCitiesHandler).Methods("GET")
	r.HandleFunc("/information/states", GetStatesHandler).Methods("GET")
	r.HandleFunc("/information/countries", GetCountriesHandler).Methods("GET")
	r.HandleFunc("/information/formats", GetFormatsHandler).Methods("GET")
	r.HandleFunc("/information/roles", GetRolesHandler).Methods("GET")
	r.HandleFunc("/information/series", GetSeriesHandler).Methods("GET")
	r.HandleFunc("/information/languages", GetLanguagesHandler).Methods("GET")
	r.HandleFunc("/information/deweys", GetDeweysHandler).Methods("GET")
	r.HandleFunc("/information/deweys/{{dewey}}", GetGenreHandler).Methods("GET")
	r.HandleFunc("/information/tags", GetTagsHandler).Methods("GET")
	r.HandleFunc("/information/awards", GetAwardsHandler).Methods("GET")
	r.HandleFunc("/libraries", GetLibrariesHandler).Methods("GET")
	r.HandleFunc("/libraries/owned", GetOwnedLibrariesHandler).Methods("GET")
	r.HandleFunc("/libraries/owned", SaveOwnedLibrariesHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/breaks", GetBreaksHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/breaks", UpdateBreaksHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/cases", GetCasesHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/cases", SaveCasesHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/series", GetAuthorBasedSeriesHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/series", UpdateAuthorBasedSeriesHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/sort", GetLibrarySortHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/sort", UpdateLibrarySortHandler).Methods("PUT")
	r.HandleFunc("/settings", GetSettingsHandler).Methods("GET")
	r.HandleFunc("/settings", SaveSettingsHandler).Methods("PUT")
	r.HandleFunc("/settings/{setting}", GetSettingHandler).Methods("GET")
	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/login", LoginHandler).Methods("POST")
	r.HandleFunc("/users", RegisterHandler).Methods("POST")
	r.HandleFunc("/users/logout", LogoutHandler).Methods("POST")
	r.HandleFunc("/users/reset", ResetPasswordHandler).Methods("PUT")
	r.HandleFunc("/users/reset/{token}", FinishResetPasswordHandler).Methods("GET")
	r.HandleFunc("/users/username", GetUsernameHandler).Methods("GET")
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./../web/"))))
	logger.Printf("Listening on port 8181")
	loggedRouter := handlers.CombinedLoggingHandler(logFile, r)
	http.ListenAndServe(":8181", handlers.CompressHandler(loggedRouter))
	// http.ListenAndServe(":8181", nil)
	logger.Printf("Closing")
}

//RouteByAuthenticate determines which page to load
func RouteByAuthenticate(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		http.Redirect(w, r, "/unregistered", 301)
	} else {
		http.Redirect(w, r, "/home", 301)
	}
}

//GetHomePageHandler gets the home page
func GetHomePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/app/main/index.html")
}

//GetUnregisteredPageHandler gets the home page
func GetUnregisteredPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/app/unregistered/unregistered.html")
}
