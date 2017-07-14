package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"database/sql"
	"encoding/csv"
	"bytes"
	"io"
	"strconv"
	"mime/multipart"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

//RunServer runs the library server
func RunServer(username, password, database string) {
	fmt.Printf("Creating the database")

	var err error
	// Create sql.DB
	db, err = sql.Open("mysql", "root:@/library?parseTime=true")
	// db, _ = sql.Open("mysql", "root:@/library?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Printf("Pinging the database")
	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	http.Handle("./web/", http.StripPrefix("./web/", http.FileServer(http.Dir("./web"))))
	http.HandleFunc("/", getLoginPage)
	http.HandleFunc("/home", getHomePage)
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/publishers", getPublishers)
	http.HandleFunc("/cities", getCities)
	http.HandleFunc("/states", getStates)
	http.HandleFunc("/countries", getCountries)
	http.HandleFunc("/series", getSeries)
	http.HandleFunc("/formats", getFormats)
	http.HandleFunc("/languages", getLanguages)
	http.HandleFunc("/roles", getRoles)
	http.HandleFunc("/deweys", getDeweys)
	http.HandleFunc("/savebook", saveBook)
	http.HandleFunc("/exportbooks", exportBooks)
	http.HandleFunc("/exportauthors", exportAuthors)
	http.HandleFunc("/import", importLibrary)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/stats", getStats)
	http.HandleFunc("/cases", getCases)
	http.HandleFunc("/dimensions", getDimensions)
	http.HandleFunc("/deletebook", deleteBook)
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./../web/"))))
	log.Printf("Listening on port 8181")
	http.ListenAndServe(":8181", nil)
	log.Printf("Closing")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("Unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	sortMethod := params.Get("sortmethod")
	isread := params.Get("isread")
	isreference := params.Get("isreference")
	isowned := params.Get("isowned")
	isloaned := params.Get("isloaned")
	isreading := params.Get("isreading")
	isshipping := params.Get("isshipping")
	text := params.Get("text")
	page := params.Get("page")
	numberToGet := params.Get("numbertoget")
	fromDewey := params.Get("fromdewey")
	toDewey := params.Get("todewey")
	books, numberOfBooks, err := GetBooks(sortMethod, isread, isreference, isowned, isloaned, isreading, isshipping, text, page, numberToGet, fromDewey, toDewey)
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(BookSet{
		Books: books,
		NumberOfBooks: numberOfBooks,
	})
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func saveBook(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	var b Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = SaveBook(b)
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

func getPublishers(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetPublishers();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getCities(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetCities();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getStates(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetStates();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetCountries();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getSeries(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetSeries();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getFormats(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetFormats();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getLanguages(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetLanguages();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getRoles(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetRoles();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getDeweys(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	d, err := GetDeweys();
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func exportBooks(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	l, err := GetBooksForExport()
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)
	err = writer.WriteAll(l)
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	writer.Flush()
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=books.csv")
	w.Write(b.Bytes())
}

func exportAuthors(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	l, err := GetAuthorsForExport()
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)
	err = writer.WriteAll(l)
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	writer.Flush()
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=authors.csv")
	w.Write(b.Bytes())
}

func importLibrary(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	var (
		err	error
	)
	defer func() {
		if nil != err {
			fmt.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
		}
	}()
	// parse request with maximum memory of _24Kilobits
	const _24K = (1 << 20) * 24
	if err = r.ParseMultipartForm(_24K); nil != err {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("starting to write")
	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {
			// open uploaded
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				fmt.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			// open destination
			var outfile *os.File
			if outfile, err = os.Create("./tmp/" + hdr.Filename); nil != err {
				fmt.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			// 32K buffer copy
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				fmt.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			reader, err := os.Open("./tmp/"+hdr.Filename)
			if err != nil {
				fmt.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			c_reader := csv.NewReader(reader)
			if err != nil {
				fmt.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			records, err := c_reader.ReadAll()
			if err != nil {
				fmt.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			ImportLibrary(records)
			w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}
	return

}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	r.ParseForm()
	key, err := LoginUser(r.Form["username"][0], r.Form["password"][0])
	if err != nil {
		fmt.Printf("%+v", err)
	}
	http.SetCookie(w, &http.Cookie{Name:"library-organizer-session",Value:key})
	http.Redirect(w, r, "/", 301)
	return
}

//todo - Allow registering after set up for multiple members
func register (w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	// 	http.Redirect(w, r, "/", 301)
	// 	return
	// }
	// r.ParseForm()
	// key, err := RegisterUser(r.Form["username"][0], r.Form["password"][0], r.Form["email"][0])
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// }
	// http.SetCookie(w, &http.Cookie{Name:"library-organizer-session",Value:key})
	http.Redirect(w, r, "/", 301)
	return
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	cookie, err := r.Cookie("library-organizer-session")
	if err != nil {
		fmt.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
	}
	err = LogoutSession(cookie.Value)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	http.Redirect(w, r, "/", 301)
	return
}

func registered(r *http.Request) bool {
	cookie, err := r.Cookie("library-organizer-session")
	if err != nil {
		fmt.Printf("%+v", err)
		return false
	}
	if cookie.Value == "" {
		return false
	}
	registered, err := IsRegistered(cookie.Value)
	if  err != nil {
		fmt.Printf("%+v", err)
		return false
	}
	return registered
}

func getLoginPage(w http.ResponseWriter, r *http.Request) {
	if registered(r) {
		http.Redirect(w, r, "/home", 301)
	} else {
		http.ServeFile(w, r, "../web/app/unregistered/index.html")
	}
}

func getHomePage(w http.ResponseWriter, r *http.Request) {
	if registered(r) {
		http.ServeFile(w, r, "../web/app/main/index.html")
	} else {
		http.Redirect(w, r, "/", 301)
	}
}

func getStats(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	d, err := GetStats(params.Get("type"))
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getCases(w http.ResponseWriter, r *http.Request) {
	d, err := GetCases()
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func getDimensions(w http.ResponseWriter, r *http.Request) {
	d, err := GetDimensions()
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d);
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(data)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	if !registered(r) {
		fmt.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	var i int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&i)
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	id := strconv.Itoa(i)
	err = DeleteBook(id)
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}