package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"./books"
)

//GetBooksHandler gets books from a filter query
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
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
	libraryids := params.Get("libraryids")
	bs, numberOfBooks, err := books.GetBooks(db, sortMethod, isread, isreference, isowned, isloaned, isreading, isshipping, text, page, numberToGet, fromDewey, toDewey, libraryids, session)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(books.BookSet{
		Books:         bs,
		NumberOfBooks: numberOfBooks,
	})
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//AddBookHandler saves a book
func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	var b books.Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = books.SaveBook(db, b)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//SaveBookHandler saves a book
func SaveBookHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	var b books.Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = books.SaveBook(db, b)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//DeleteBookHandler deletes a book
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	var i int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&i)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	id := strconv.Itoa(i)
	err = books.DeleteBook(db, id)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//CheckoutBookHandler checks out a book
func CheckoutBookHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	var i int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&i)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = books.CheckoutBook(db, session, i)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//CheckinBookHandler checks out a book
func CheckinBookHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	var i int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&i)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = books.CheckinBook(db, i)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//ImportBooksHandler imports books
func ImportBooksHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	var (
		err error
	)
	defer func() {
		if nil != err {
			logger.Printf("%+v", err)
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			return
		}
	}()
	// parse request with maximum memory of _24Kilobits
	const _24K = (1 << 20) * 24
	if err = r.ParseMultipartForm(_24K); nil != err {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	logger.Printf("starting to write")
	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {
			// open uploaded
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				logger.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			// open destination
			var outfile *os.File
			if outfile, err = os.Create("./tmp/" + hdr.Filename); nil != err {
				logger.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			// 32K buffer copy
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				logger.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			reader, err := os.Open("./tmp/" + hdr.Filename)
			if err != nil {
				logger.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			cReader := csv.NewReader(reader)
			if err != nil {
				logger.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			records, err := cReader.ReadAll()
			if err != nil {
				logger.Printf("%+v", err)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}
			books.ImportBooks(db, records)
			w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}
	return
}

//ExportBooksHandler exports books
func ExportBooksHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	l, err := books.GetBooksForExport(db)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)
	err = writer.WriteAll(l)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	writer.Flush()
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=books.csv")
	w.Write(b.Bytes())
}

//ExportAuthorsHandler exports contributors
func ExportAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	l, err := books.GetAuthorsForExport(db)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)
	err = writer.WriteAll(l)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	writer.Flush()
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=authors.csv")
	w.Write(b.Bytes())
}
