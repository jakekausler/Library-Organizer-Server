package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./information"
	"github.com/gorilla/mux"
)

//GetStatsHandler gets statistics
func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("Unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	d, err := information.GetStats(db, params.Get("type"), params.Get("libraryids"))
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

//GetDimensionsHandler gets dimensions
func GetDimensionsHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("Unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	libraryids := params.Get("libraryids")
	d, err := information.GetDimensions(db, libraryids)
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

//GetPublishersHandler gets publishers
func GetPublishersHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetPublishers(db, str)
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

//GetCitiesHandler gets cities
func GetCitiesHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetCities(db, str)
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

//GetStatesHandler gets states
func GetStatesHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetStates(db, str)
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

//GetCountriesHandler gets countries
func GetCountriesHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetCountries(db, str)
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

//GetSeriesHandler gets series
func GetSeriesHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetSeries(db, str)
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

//GetFormatsHandler gets formats
func GetFormatsHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetFormats(db, str)
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

//GetLanguagesHandler gets languages
func GetLanguagesHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetLanguages(db, str)
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

//GetRolesHandler gets roles
func GetRolesHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetRoles(db, str)
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

//GetDeweysHandler gets deweys
func GetDeweysHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetDeweys(db, str)
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

//GetGenreHandler gets deweys
func GetGenreHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	dewey := mux.Vars(r)["dewey"]
	d, err := information.GetGenre(db, dewey)
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

//GetTagsHandler gets tags
func GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetTags(db, str)
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

//GetAwardsHandler gets tags
func GetAwardsHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Registered(r); !ok {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query()
	str := params.Get("str")
	d, err := information.GetAwards(db, str)
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