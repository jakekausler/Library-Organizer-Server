package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./settings"
)

//GetSettingsHandler gets settings
func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("library-organizer-session")
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
	cookie, err := r.Cookie("library-organizer-session")
	if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
	}
	b, err := ioutil.ReadAll(r.Body)
	name := string(b)
	defer r.Body.Close()
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

//UpdateSettingsHandler updates settings
func UpdateSettingsHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("library-organizer-session")
	if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
	}
	var settings []Setting
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&settings)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = settings.UpdateSettings(db, settings, cookie.Value)
	if err != nil {
		logger.Printf("%+v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	return
}
