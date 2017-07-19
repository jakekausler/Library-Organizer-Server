package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./libraries"
)

//GetLibrariesHandler gets libraries
func GetLibrariesHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("library-organizer-session")
	if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
	}
	d, err := libraries.GetLibraries(db, cookie.Value)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//GetCasesHandler gets cases
func GetCasesHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("library-organizer-session")
	if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
	}
	if cookie == nil || cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
	}
	params := r.URL.Query()
	libraryid := params.Get("libraryid")
	sortmethod := params.Get("sortmethod")
	d, err := libraries.GetCases(db, libraryid, sortmethod, cookie.Value)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//SaveCasesHandler saves cases
func SaveCasesHandler(w http.ResponseWriter, r *http.Request) {
	if !Registered(r) {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	var editedCases EditedCases
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&editedCases)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = libraries.SaveCases(db, editedCases)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//AddBreakHandler adds a break
func AddBreakHandler(w http.ResponseWriter, r *http.Request) {
	if !Registered(r) {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	var b Break
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = libraries.AddBreak(db, b.LibraryID, b.ValueType, b.Value, b.BreakType)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}
