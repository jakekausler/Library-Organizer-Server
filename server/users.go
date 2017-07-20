package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"./users"
)

//LoginHandler logs in a user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	r.ParseForm()
	key, err := users.LoginUser(db, r.Form["username"][0], r.Form["password"][0])
	if err != nil {
		logger.Printf("%+v", err)
	}
	http.SetCookie(w, &http.Cookie{Name: "libraryorganizersession", Value: key, Expires: time.Now().Add(14*24*time.Hour)})
	http.Redirect(w, r, "/", 301)
	return
}

//RegisterHandler registers a user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	r.ParseForm()
	key, err := users.RegisterUser(db, r.Form["username"][0], r.Form["password"][0], r.Form["email"][0])
	if err != nil {
		logger.Printf("%+v", err)
	}
	http.SetCookie(w, &http.Cookie{Name: "libraryorganizersession", Value: key})
	http.Redirect(w, r, "/", 301)
	return
}

//LogoutHandler logs out a user
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	cookie, err := r.Cookie("libraryorganizersession")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", 301)
		return
	} else if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
		return
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
		return
	}
	err = users.LogoutSession(db, cookie.Value)
	if err != nil {
		logger.Printf("%+v", err)
	}
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
func Registered(r *http.Request) bool {
	cookie, err := r.Cookie("libraryorganizersession")
	if err == http.ErrNoCookie {
		return false
	} else if err != nil {
		logger.Printf("%+v", err)
		return false
	}
	if cookie.Value == "" {
		return false
	}
	registered, err := users.IsRegistered(db, cookie.Value)
	if err != nil {
		logger.Printf("%+v", err)
		return false
	}
	return registered
}

//GetUsernameHandler gets a username
func GetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("libraryorganizersession")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", 301)
		return
	} else if err != nil {
		logger.Printf("%+v", err)
		http.Redirect(w, r, "/", 301)
		return
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", 301)
		return
	}
	d, err := users.GetUsername(db, cookie.Value)
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
	d, err := users.GetUsers(db, cookie.Value)
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
