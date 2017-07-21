package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./users"
	"github.com/gorilla/sessions"
)

//LoginHandler logs in a user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	r.ParseForm()
	key, err := users.LoginUser(db, r.Form["username"][0], r.Form["password"][0])
	if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
		return
	}
	session.Values["libraryorganizersession"] = key
	sessions.Save(r, w)
	http.Redirect(w, r, "/", 301)
	return
}

//RegisterHandler registers a user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	r.ParseForm()
	key, err := users.RegisterUser(db, r.Form["username"][0], r.Form["password"][0], r.Form["email"][0], r.Form["first"][0], r.Form["last"][0])
	if err != nil {
		logger.Printf("%+v", err)
	}
	session.Values["libraryorganizersession"] = key
	sessions.Save(r, w)
	http.Redirect(w, r, "/", 301)
	return
}

//LogoutHandler logs out a user
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	registered, key := Registered(r)
	if !registered {
		http.Redirect(w, r, "/", 301)
		return
	}
	err := users.LogoutSession(db, key)
	if err != nil {
		logger.Printf("%+v", err)
	}
	session.Values["libraryorganizersession"] = ""
	http.Redirect(w, r, "/", 301)
	return
}

//ResetPasswordHandler sends a link to reset a password
func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	r.ParseForm()
	err := users.ResetPassword(db, r.Form["email"][0])
	if err != nil {
		logger.Printf("%+v", err)
	}
	http.Redirect(w, r, "/", 301)
	return
}

//FinishResetPasswordHandler sends a link to reset a password
func FinishResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	return
}

//Registered determines whether a user is registered
func Registered(r *http.Request) (bool, string) {
	session, _ := store.Get(r, "session")
	sessionkey := session.Values["libraryorganizersession"]
	if sessionkey == nil {
		return false, ""
	}
	key := sessionkey.(string)
	if key == "" {
		return false, ""
	}
	registered, err := users.IsRegistered(db, key)
	if err != nil {
		logger.Printf("%+v", err)
		return false, ""
	}
	return registered, key
}

//GetUsernameHandler gets a username
func GetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	d, err := users.GetUsername(db, session)
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

//GetUsersHandler gets users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	registered, session := Registered(r)
	if !registered {
		logger.Printf("unauthorized")
		http.Error(w, fmt.Sprintf("Unauthorized"), http.StatusInternalServerError)
		return
	}
	d, err := users.GetUsers(db, session)
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
