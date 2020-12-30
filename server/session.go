package libraryserver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
    "time"
)

//Env Variables
const (
	LogEnv = "LIBRARY_LOG_ROOT"
	ResEnv = "LIBRARY_RESOURCE_ROOT"
	AppEnv = "LIBRARY_APP_ROOT"
)

var (
	db     *sql.DB
	logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)
	store  = sessions.NewCookieStore([]byte("Session"))

	logRoot = os.Getenv(LogEnv)
	resRoot = os.Getenv(ResEnv)
	appRoot = os.Getenv(AppEnv)
)

//RunServer runs the library server
func RunServer(host, username, password, database string, appport, mysqlport int) {
	logger.Printf("Creating the database")
	var err error
	// Create sql.DB
	db, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, host, mysqlport, database))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	logger.Printf("Pinging the database")
	// Test the connection
	err = db.Ping()
	for err != nil {
        logger.Printf("Could not reach the database. Sleeping for 10 seconds")
        time.Sleep(10 * time.Second)
		err = db.Ping()
	}
	logger.Printf("Opening the log")
	logFile, err := os.OpenFile(fmt.Sprintf("%v/server.log", logRoot), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", RouteByAuthenticate)
	r.HandleFunc("/home", RouteByAuthenticate)
	r.HandleFunc("/unregistered", GetUnregisteredPageHandler)
	r.HandleFunc("/books", GetBooksHandler).Methods("GET")
	r.HandleFunc("/books", AddBookHandler).Methods("POST")
	r.HandleFunc("/books", SaveBookHandler).Methods("PUT")
	r.HandleFunc("/books/checkout", CheckoutBookHandler).Methods("PUT")
	r.HandleFunc("/books/checkin", CheckinBookHandler).Methods("PUT")
	r.HandleFunc("/books/books", ExportBooksHandler).Methods("GET")
	r.HandleFunc("/books/contributors", ExportAuthorsHandler).Methods("GET")
	r.HandleFunc("/books/books", ImportBooksHandler).Methods("POST")
	r.HandleFunc("/books/{bookid}", DeleteBookHandler).Methods("DELETE")
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
	r.HandleFunc("/info/stats", GetStatsHandler2).Methods("GET")
	r.HandleFunc("/libraries", GetLibrariesHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}", GetLibraryHandler).Methods("GET")
	r.HandleFunc("/libraries/owned", GetOwnedLibrariesHandler).Methods("GET")
	r.HandleFunc("/libraries/owned", SaveOwnedLibrariesHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/breaks", GetBreaksHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/breaks", UpdateBreaksHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/cases/ids", GetCaseIDsHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/cases/{caseid}/shelves/ids", GetShelfIDsHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/cases", GetCasesHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/cases", SaveCasesHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/cases", RefreshCasesHandler).Methods("POST")
	r.HandleFunc("/libraries/{libraryid}/series", GetAuthorBasedSeriesHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/series", UpdateAuthorBasedSeriesHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/sort", GetLibrarySortHandler).Methods("GET")
	r.HandleFunc("/libraries/{libraryid}/sort", UpdateLibrarySortHandler).Methods("PUT")
	r.HandleFunc("/libraries/{libraryid}/search", GetLibrarySearchHandler).Methods("GET")
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
	r.HandleFunc("/imagelist", GetImages).Methods("GET")
	r.HandleFunc("/information/locationcounts", GetPublisherLocationCounts).Methods("GET")
	r.HandleFunc("/bookimages/{bookid}", GetBookImage).Methods("GET")
	r.HandleFunc("/caseimages/{caseid}", GetCaseImage).Methods("GET")
	r.HandleFunc("/shelfimages/{caseid}/{shelfid}", GetShelfImage).Methods("GET")
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir(appRoot)))) //+"/../"))))
	logger.Printf("Listening on port %v", appport)
	loggedRouter := handlers.CombinedLoggingHandler(logFile, r)
	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", appport), handlers.CompressHandler(loggedRouter)))
	// http.ListenAndServe(":8181", nil)
}

//RouteByAuthenticate determines which page to load
func RouteByAuthenticate(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		http.Redirect(w, r, "/unregistered", 301)
	} else {
		http.ServeFile(w, r, fmt.Sprintf("%v/index.html", appRoot))
	}
}

//GetUnregisteredPageHandler gets the home page
func GetUnregisteredPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("%v/index.html", appRoot))
}

//GetBookImage gets an image
func GetBookImage(w http.ResponseWriter, r *http.Request) {
	bookid := mux.Vars(r)["bookid"]
	params := r.URL.Query()
	size := params.Get("size")
	http.ServeFile(w, r, fmt.Sprintf("%v/bookimages/%v/%v.jpg", resRoot, size, bookid))
}

//GetCaseImage gets an image
func GetCaseImage(w http.ResponseWriter, r *http.Request) {
	caseid := mux.Vars(r)["caseid"]
	http.ServeFile(w, r, fmt.Sprintf("%v/caseimages/%v.svg", resRoot, caseid))
}

//GetShelfImage gets an image
func GetShelfImage(w http.ResponseWriter, r *http.Request) {
	caseid := mux.Vars(r)["caseid"]
	shelfid := mux.Vars(r)["shelfid"]
	http.ServeFile(w, r, fmt.Sprintf("%v/caseimages/%v/%v.svg", resRoot, caseid, shelfid))
}
