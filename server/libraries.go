package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./libraries"
	"github.com/gorilla/mux"
)

//GetLibrariesHandler gets libraries
func GetLibrariesHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	d, err := libraries.GetLibraries(db, session)
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
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	params := r.URL.Query()
	sortmethod := params.Get("sortmethod")
	d, err := libraries.GetCases(db, libraryid, sortmethod, session)
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
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	var editedCases libraries.EditedCases
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&editedCases)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = libraries.SaveCases(db, libraryid, editedCases)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//GetBreaksHandler gets the breaks for a library
func GetBreaksHandler(w http.ResponseWriter, r *http.Request) {
	return
}

//AddBreakHandler adds a break
func AddBreakHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	var b libraries.Break
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = libraries.AddBreak(db, libraryid, b.ValueType, b.Value, b.BreakType)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//SaveBreakHandler saves a break
func SaveBreakHandler(w http.ResponseWriter, r *http.Request) {
	return
}

//DeleteBreakHandler deletes a break
func DeleteBreakHandler(w http.ResponseWriter, r *http.Request) {
	return
}

//GetOwnedLibrariesHandler adds a break
func GetOwnedLibrariesHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	d, err := libraries.GetOwnedLibraries(db, session)
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

//SaveOwnedLibrariesHandler saves owned libraries
func SaveOwnedLibrariesHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	var ownedLibraries []libraries.OwnedLibrary
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ownedLibraries)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = libraries.SaveOwnedLibraries(db, ownedLibraries, session)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}