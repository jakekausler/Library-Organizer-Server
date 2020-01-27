package settings

import (
	"database/sql"
	"log"
	"os"

	"github.com/jakekausler/Library-Organizer-3.0/server/users"
)

const (
	getValueQuery          = "SELECT IF (EXISTS (SELECT value FROM librarysettings WHERE userid=? AND setting=?), (SELECT value FROM librarysettings WHERE userid=? AND setting=? LIMIT 1), (SELECT value FROM librarysettings WHERE userid=0 AND setting=? LIMIT 1))"
	getValueTypeQuery      = "SELECT IF (EXISTS (SELECT valuetype FROM librarysettings WHERE userid=? AND setting=?), (SELECT valuetype FROM librarysettings WHERE userid=? AND setting=? LIMIT 1), (SELECT valuetype FROM librarysettings WHERE userid=0 AND setting=? LIMIT 1))"
	getSettingGroupQuery   = "SELECT IF (EXISTS (SELECT settinggroup FROM librarysettings WHERE userid=? AND setting=?), (SELECT settinggroup FROM librarysettings WHERE userid=? AND setting=? LIMIT 1), (SELECT settinggroup FROM librarysettings WHERE userid=0 AND setting=? LIMIT 1))"
	getPossibleValuesQuery = "SELECT possiblevalue FROM librarysettingspossiblevalues WHERE setting=? ORDER BY possiblevalue"
	updateSettingQuery     = "REPLACE INTO librarysettings (userid, setting, value, valuetype, settinggroup) VALUES (?,?,?,?,?)"
)

var logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)

//Setting is a setting
type Setting struct {
	Name           string   `json:"name"`
	Value          string   `json:"value"`
	ValueType      string   `json:"valuetype"`
	Group          string   `json:"group"`
	PossibleValues []string `json:"possiblevalues"`
}

//GetSettings gets user settings
func GetSettings(db *sql.DB, session string) ([]Setting, error) {
	var settings []Setting
	userid, err := users.GetUserID(db, session)
	query := "SELECT setting from librarysettings group by setting"
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	for rows.Next() {
		var setting Setting
		var value sql.NullString
		var valuetype sql.NullString
		var group sql.NullString
		rows.Scan(&setting.Name)
		err = db.QueryRow(getValueQuery, userid, setting.Name, userid, setting.Name, setting.Name).Scan(&value)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		err = db.QueryRow(getValueTypeQuery, userid, setting.Name, userid, setting.Name, setting.Name).Scan(&valuetype)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		err = db.QueryRow(getSettingGroupQuery, userid, setting.Name, userid, setting.Name, setting.Name).Scan(&group)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		setting.Value = value.String
		setting.ValueType = valuetype.String
		setting.Group = group.String
		innerrows, err := db.Query(getPossibleValuesQuery, setting.Name)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		for innerrows.Next() {
			var value string
			err = innerrows.Scan(&value)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return nil, err
			}
			setting.PossibleValues = append(setting.PossibleValues, value)
		}
		settings = append(settings, setting)
	}
	return settings, nil
}

//UpdateSettings updates user settings
func UpdateSettings(db *sql.DB, settings []Setting, session string) error {
	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	for _, setting := range settings {
		_, err = db.Exec(updateSettingQuery, userid, setting.Name, setting.Value, setting.ValueType, setting.Group)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
	}
	return nil
}

//GetSetting gets a user setting
func GetSetting(db *sql.DB, name, session string) (string, error) {
	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	var value string
	err = db.QueryRow(getValueQuery, userid, name, userid, name, name).Scan(&value)
	return value, err
}
