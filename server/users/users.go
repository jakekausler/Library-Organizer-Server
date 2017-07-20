package users

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	getValidUserSession = "SELECT sessionkey from usersession WHERE sessionkey=? AND EXISTS (SELECT id FROM library_members where id=userid)"
	addUser             = "INSERT INTO library_members (usr,pass,email) values (?,?,?)"
	addSession          = "INSERT INTO usersession (sessionkey,userid,LastSeenTime) values (?,?,NOW())"
	updateSessionTime   = "UPDATE usersession SET LastSeenTime=NOW()"
	isSessionNameTaken  = "SELECT sessionkey from usersession where sessionkey=?"
	deleteSession       = "DELETE FROM usersession WHERE sessionkey=?"
	charset             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)

//User is a library member
type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
}

//ResetPassword sends a link to reset a password
//Todo
func ResetPassword(db *sql.DB, email string) error {
	// query := "SELECT id FROM library_members WHERE email=?"
	// var userid int64
	// err := db.QueryRow(query, email).Scan(&userid)
	// if err != nil {
	// 	logger.Printf("%+v", err)
	// 	return err
	// }
	return nil
}

//IsRegistered returns whether a user session is registered and the corresponding user exists
func IsRegistered(db *sql.DB, session string) (bool, error) {
	var sessionkey string
	if err := db.QueryRow(getValidUserSession, session).Scan(&sessionkey); err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return false, err
	}
}

//MarkAsSeen marks a user session as seen
func MarkAsSeen(db *sql.DB, session string) error {
	_, err := db.Exec(updateSessionTime, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

//IsUser returns whether a username and password is valid. If they are, return the userid. If not, return -1
func IsUser(db *sql.DB, username, password string) (int64, error) {
	query := "SELECT id, pass FROM library_members WHERE usr=?"
	var id int64
	var hash string
	if err := db.QueryRow(query, username).Scan(&id, &hash); err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
			return -1, err
		}
		return id, nil
	} else if err == sql.ErrNoRows {
		return -1, nil
	} else {
		return -1, err
	}
}

//GenerateSessionKey generates a random unique session key
func GenerateSessionKey(db *sql.DB) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, 50)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

//LoginUser logs in a user
func LoginUser(db *sql.DB, username, password string) (string, error) {
	id, err := IsUser(db, username, password)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	key := GenerateSessionKey(db)
	_, err = db.Exec(addSession, key, id)
	return key, err
}

//RegisterUser registers a user and creates a default library for him
func RegisterUser(db *sql.DB, username, password, email string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	result, err := db.Exec(addUser, username, hash, email)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	key := GenerateSessionKey(db)
	_, err = db.Exec(addSession, key, id)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	query := "INSERT INTO libraries (name, ownerid) VALUES (?,?)"
	result, err = db.Exec(query, "default", id)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	libID, err := result.LastInsertId()
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	query = "INSERT INTO permissions (userid, libraryid, permission) VALUES (?,?,7)"
	_, err = db.Exec(query, id, libID)
	return key, err
}

//LogoutSession logs out a user
func LogoutSession(db *sql.DB, sessionkey string) error {
	_, err := db.Exec(deleteSession, sessionkey)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

//GetUsername gets a username from a session
func GetUsername(db *sql.DB, session string) (string, error) {
	var name string
	query := "SELECT usr FROM library_members JOIN usersession ON UserID=ID WHERE sessionkey=?"
	err := db.QueryRow(query, session).Scan(&name)
	return name, err
}

//GetUserID gets a userid from a session
func GetUserID(db *sql.DB, session string) (string, error) {
	var name string
	query := "SELECT id FROM library_members JOIN usersession ON UserID=ID WHERE sessionkey=?"
	err := db.QueryRow(query, session).Scan(&name)
	return name, err
}

//GetUsers gets users
func GetUsers(db *sql.DB, session string) ([]User, error) {
	userid, err := GetUserID(db, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	var users []User
	query := "SELECT id, usr FROM library_members WHERE id != ?"
	rows, err := db.Query(query, userid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}