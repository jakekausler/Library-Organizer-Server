package libraryserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jakekausler/Library-Organizer-2.0/server/libraries"
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

//GetLibraryHandler gets a library by its id
func GetLibraryHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	d, err := libraries.GetLibrary(db, session, libraryid)
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
	// params := r.URL.Query()
	r.ParseForm()
	_, includeBooks := r.Form["includeBooks"]
	d, _, err := libraries.GetCases(db, libraryid, session, includeBooks, true, false, true)
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

//RefreshCasesHandler refreshes case and shelf images
func RefreshCasesHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	err := libraries.RefreshCases(db, libraryid, session, resRoot)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//GetBreaksHandler gets the breaks for a library
func GetBreaksHandler(w http.ResponseWriter, r *http.Request) {
	registered, _ := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	d, err := libraries.GetLibraryBreaks(db, libraryid)
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

//UpdateBreaksHandler adds a break
func UpdateBreaksHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	var b []libraries.Break
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = libraries.UpdateBreaks(db, libraryid, b)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
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

//GetAuthorBasedSeriesHandler gets series that are sorted by author then title, instead of volume
func GetAuthorBasedSeriesHandler(w http.ResponseWriter, r *http.Request) {
	registered, _ := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	d, err := libraries.GetAuthorBasedSeries(db, libraryid)
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

//UpdateAuthorBasedSeriesHandler updates series that are sorted by author then title, instead of volume
func UpdateAuthorBasedSeriesHandler(w http.ResponseWriter, r *http.Request) {
	registered, _ := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	var series []string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&series)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = libraries.UpdateAuthorBasedSeries(db, libraryid, series)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//GetLibrarySortHandler gets series that are sorted by author then title, instead of volume
func GetLibrarySortHandler(w http.ResponseWriter, r *http.Request) {
	registered, _ := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	d, err := libraries.GetLibrarySort(db, libraryid)
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

//UpdateLibrarySortHandler updates series that are sorted by author then title, instead of volume
func UpdateLibrarySortHandler(w http.ResponseWriter, r *http.Request) {
	registered, _ := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	var method string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&method)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = libraries.UpdateLibrarySort(db, libraryid, method)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}

//GetLibrarySearchHandler gets titles by search
func GetLibrarySearchHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	params := r.URL.Query()
	text := params.Get("text")
	searchusingtitle := params.Get("searchusingtitle") == "true"
	searchusingsubtitle := params.Get("searchusingsubtitle") == "true"
	searchusingseries := params.Get("searchusingseries") == "true"
	searchusingauthor := params.Get("searchusingauthor") == "true"
	d, err := libraries.SearchShelves(db, libraryid, session, text, searchusingtitle, searchusingsubtitle, searchusingseries, searchusingauthor)
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

//GetCaseIDsHandler gets cases
func GetCaseIDsHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	d, err := libraries.GetCaseIDs(db, libraryid, session)
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

//GetCaseIDsHandler gets cases
func GetShelfIDsHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	libraryid := mux.Vars(r)["libraryid"]
	caseid := mux.Vars(r)["caseid"]
	d, err := libraries.GetShelfIDs(db, libraryid, session, caseid)
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
