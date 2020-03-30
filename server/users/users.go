package users

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jakekausler/Library-Organizer-2.0/server/information"
	"golang.org/x/crypto/bcrypt"
)

const (
	getValidUserSessionQuery = "SELECT sessionkey from usersession WHERE sessionkey=? AND EXISTS (SELECT id FROM library_members where id=userid)"
	addUserQuery             = "INSERT INTO library_members (usr,pass,email,firstname,lastname) values (?,?,?,?,?)"
	addSessionQuery          = "INSERT INTO usersession (sessionkey,userid,LastSeenTime) values (?,?,NOW())"
	updateSessionTimeQuery   = "UPDATE usersession SET LastSeenTime=NOW()"
	deleteSessionQuery       = "DELETE FROM usersession WHERE sessionkey=?"
	getIDByEmailQuery        = "SELECT id FROM library_members WHERE email=?"
	getUserForCheckQuery     = "SELECT id, pass FROM library_members WHERE usr=?"
	addLibraryQuery          = "INSERT INTO libraries (name, ownerid, sortmethod) VALUES (?,?,?)"
	addPermissionQuery       = "INSERT INTO permissions (userid, libraryid, permission) VALUES (?,?,?)"
	getUsernameQuery         = "SELECT usr FROM library_members JOIN usersession ON UserID=ID WHERE sessionkey=?"
	getUserIDQuery           = "SELECT id FROM library_members JOIN usersession ON UserID=ID WHERE sessionkey=?"
	getUsersQuery            = "SELECT id, usr, firstname, lastname, email FROM library_members WHERE id != ?"

	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)

//User is a library member
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	FullName  string `json:"fullname"`
	Email     string `json:"email"`
}

//ResetPassword sends a link to reset a password
//Todo
func ResetPassword(db *sql.DB, email string) error {
	// var userid int64
	// err := db.QueryRow(getIDByEmailQuery, email).Scan(&userid)
	// if err != nil {
	// 	logger.Printf("%+v", err)
	// 	return err
	// }
	return nil
}

//IsRegistered returns whether a user session is registered and the corresponding user exists
func IsRegistered(db *sql.DB, session string) (bool, error) {
	var sessionkey string
	if err := db.QueryRow(getValidUserSessionQuery, session).Scan(&sessionkey); err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return false, err
	}
}

//MarkAsSeen marks a user session as seen
func MarkAsSeen(db *sql.DB, session string) error {
	_, err := db.Exec(updateSessionTimeQuery, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

//IsUser returns whether a username and password is valid. If they are, return the userid. If not, return -1
func IsUser(db *sql.DB, username, password string) (int64, error) {
	var id int64
	var hash string
	if err := db.QueryRow(getUserForCheckQuery, username).Scan(&id, &hash); err == nil {
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
	_, err = db.Exec(addSessionQuery, key, id)
	return key, err
}

//RegisterUser registers a user and creates a default library for him
func RegisterUser(db *sql.DB, username, password, email, first, last string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	result, err := db.Exec(addUserQuery, username, hash, email, first, last)
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
	_, err = db.Exec(addSessionQuery, key, id)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	result, err = db.Exec(addLibraryQuery, "default", id, information.SORTMETHOD)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	libID, err := result.LastInsertId()
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	_, err = db.Exec(addPermissionQuery, id, libID, 7)
	return key, err
}

//LogoutSession logs out a user
func LogoutSession(db *sql.DB, sessionkey string) error {
	_, err := db.Exec(deleteSessionQuery, sessionkey)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

//GetUsername gets a username from a session
func GetUsername(db *sql.DB, session string) (string, error) {
	var name string
	err := db.QueryRow(getUsernameQuery, session).Scan(&name)
	return name, err
}

//GetUserID gets a userid from a session
func GetUserID(db *sql.DB, session string) (string, error) {
	var name string
	err := db.QueryRow(getUserIDQuery, session).Scan(&name)
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
	rows, err := db.Query(getUsersQuery, userid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		user.FullName = user.FirstName + " " + user.LastName
		users = append(users, user)
	}
	return users, nil
}
