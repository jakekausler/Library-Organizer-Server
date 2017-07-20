package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./settings"
	"github.com/gorilla/mux"
)

//GetSettingsHandler gets settings
func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("libraryorganizersession")
	if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
	}
	d, err := settings.GetSettings(db, cookie.Value)
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

//GetSettingHandler gets a setting
func GetSettingHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("libraryorganizersession")
	if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
		return
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
		return
	}
	name := mux.Vars(r)["setting"]
	d, err := settings.GetSetting(db, name, cookie.Value)
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

//SaveSettingsHandler updates settings
func SaveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("libraryorganizersession")
	if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
	}
	var s []settings.Setting
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&s)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = settings.UpdateSettings(db, s, cookie.Value)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}
