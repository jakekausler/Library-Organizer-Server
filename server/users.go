package libraryserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jakekausler/Library-Organizer-2.0/server/users"
)

//LoginHandler logs in a user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}
	r.ParseForm()
	username, u_ok := r.Form["username"]
	password, p_ok := r.Form["password"]
	if !u_ok || !p_ok {
		decoder := json.NewDecoder(r.Body)
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err := decoder.Decode(&user)
		if err != nil {
			logger.Printf("%+v", err)
			http.Redirect(w, r, "/", 301)
			return
		}
		key, err := users.LoginUser(db, user.Username, user.Password)
		if err != nil {
			logger.Printf("%+v", err)
			http.Redirect(w, r, "/", 301)
			return
		}
		session.Values["libraryorganizersession"] = key
	} else {
		key, err := users.LoginUser(db, username[0], password[0])
		if err != nil {
			logger.Printf("%+v", err)
			http.Redirect(w, r, "/", 301)
			return
		}
		session.Values["libraryorganizersession"] = key
	}
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
	registered, key := Registered(r)
	if !registered {
		session.Values["libraryorganizersession"] = ""
		sessions.Save(r, w)
		http.Redirect(w, r, "/", 301)
		return
	}
	err := users.LogoutSession(db, key)
	if err != nil {
		logger.Printf("%+v", err)
	}
	session.Values["libraryorganizersession"] = ""
	sessions.Save(r, w)
	http.Redirect(w, r, "/", 301)
	return
}

//IsLoggedIn determines if a user is logged in
func IsLoggedIn(r *http.Request) bool {
	registered, key := Registered(r)
	if !registered {
		return false
	}
	return key != ""
}

//IsLoggedInHandler determines if a user is logged in
func IsLoggedInHandler(w http.ResponseWriter, r *http.Request) {
	if IsLoggedIn(r) {
		w.Write([]byte("true"))
		return
	}
	w.Write([]byte("false"))
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

//CanWrite determines if the current user has write access to a library
func CanWrite(session string, libraryid string) bool {
	return users.CanWrite(db, session, libraryid)
}

//Registered determines whether a user is registered
func Registered(r *http.Request) (bool, string) {
	//TODO: Undo these comments
	// return true, ""
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
