package libraries

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
	"fmt"
	"math"
	"image/color"

	"github.com/jakekausler/Library-Organizer-2.0/server/books"
	"github.com/jakekausler/Library-Organizer-2.0/server/information"
	"github.com/jakekausler/Library-Organizer-2.0/server/users"

	"github.com/ajstarks/svgo"
)

const (
	getLibrariesQuery = "SELECT libraries.id, name, permission, usr FROM libraries JOIN permissions ON libraries.id=permissions.libraryid join library_members on libraries.ownerid=library_members.id WHERE permissions.permission & 1 and permissions.userid=(SELECT id from library_members join usersession on library_members.id=usersession.userid WHERE sessionkey=?)"
	getLibraryQuery = "SELECT libraries.id, name, permission, usr FROM libraries JOIN permissions ON libraries.id=permissions.libraryid join library_members on libraries.ownerid=library_members.id WHERE libraries.id=? AND permissions.permission & 1 and permissions.userid=(SELECT id from library_members join usersession on library_members.id=usersession.userid WHERE sessionkey=?)"
	getBreaksQuery    = "SELECT breaktype, valuetype, value FROM breaks WHERE libraryid=?"
	//getCasesQuery                    = "SELECT CaseId, Width, SpacerHeight, PaddingLeft, PaddingRight, BookMargin, CaseNumber, NumberOfShelves, ShelfHeight FROM bookcases WHERE libraryid=? ORDER BY CaseNumber"
	getCasesQuery                    = "SELECT caseid, casenumber, paddingleft, paddingright, bookmargin, spacerheight FROM bookcases WHERE libraryid=? ORDER BY casenumber"
	getCaseIdsQuery                  = "SELECT caseid FROM bookcases WHERE libraryid=? ORDER BY casenumber"
	getShelfIdsQuery                 = "SELECT id FROM shelves WHERE caseid=? ORDER BY ShelfNumber"
	getShelvesQuery                  = "SELECT id, shelfnumber, caseid, width, height, alignment FROM shelves WHERE shelves.caseid=? ORDER BY shelfnumber"
	addCaseQuery                     = "INSERT INTO bookcases (casenumber, width, spacerheight, paddingleft, paddingright, libraryid, numberofshelves, shelfheight) VALUES (?,?,?,?,?,?,?,?)"
	updateCaseQuery                  = "UPDATE bookcases SET casenumber=?, width=?, spacerheight=?, paddingleft=?, paddingright=?, libraryid=?, numberOfShelves=?, shelfheight=? WHERE caseid=?"
	deleteBreaksQuery                = "DELETE FROM breaks WHERE libraryid=?"
	addBreakQuery                    = "INSERT INTO breaks (libraryid, breaktype, valuetype, value) VALUES (?,?,?,?)"
	getOwnedLibrariesQuery           = "SELECT libraries.id, libraries.name FROM libraries WHERE libraries.ownerid=?"
	getLibraryMembersPermissionQuery = "SELECT library_members.id, library_members.usr, library_members.firstname, library_members.lastname, library_members.email, library_members.iconurl, permissions.permission FROM libraries JOIN permissions ON libraries.id=permissions.libraryid JOIN library_members ON permissions.userid=library_members.id WHERE libraries.id=? AND permissions.userid != ?"
	deleteLibrariesQuery             = "DELETE FROM libraries WHERE ownerid=?"
	addLibraryQuery                  = "INSERT INTO libraries (name, ownerid, sortmethod) VALUES (?,?)"
	deletePermissionsQuery           = "DELETE FROM permissions WHERE libraryid=?"
	updateBooksLibraryQuery          = "UPDATE books SET libraryid=? WHERE libraryid=?"
	updateCasesLibraryQuery          = "UPDATE bookcases SET libraryid=? WHERE libraryid=?"
	updateBreaksLibraryQuery         = "UPDATE breaks SET libraryid=? WHERE libraryid=?"
	updateSortsLibraryQuery          = "UPDATE series_author_sorts SET libraryid=? WHERE libraryid=?"
	addPermissionQuery               = "INSERT INTO permissions (userid, libraryid, permission) VALUES (?,?,?)"
	getBooksLibrariesQuery           = "SELECT DISTINCT(libraryid) FROM books"
	getPermissionsLibrariesQuery     = "SELECT DISTINCT(libraryid) FROM permissions"
	getCasesLibrariesQuery           = "SELECT DISTINCT(libraryid) FROM bookcases"
	getBreaksLibrariesQuery          = "SELECT DISTINCT(libraryid) FROM breaks"
	getSeriesLibrariesQuery          = "SELECT DISTINCT(libraryid) FROM series_author_sorts"
	getLibraryIdsQuery               = "SELECT id FROM libraries"
	deleteBooksLibraryQuery          = "DELETE FROM books WHERE libraryid=?"
	deletePermissionsLibraryQuery    = "DELETE FROM permissions WHERE libraryid=?"
	deleteBreaksLibraryQuery         = "DELETE FROM breaks WHERE libraryid=?"
	deleteCasesLibraryQuery          = "DELETE FROM bookcases WHERE libraryid=?"
	deleteSeriesLibraryQuery         = "DELETE FROM series_author_sorts WHERE libraryid=?"
	getAuthorBasedSeriesQuery        = "SELECT series FROM series_author_sorts WHERE libraryid=?"
	updateAuthorBasedSeriesQuery     = "DELETE FROM series_author_sorts WHERE libraryid=?"
	addAuthorBasedSeriesQuery        = "INSERT INTO series_author_sorts (libraryid, series) VALUES (?,?)"
	getSortMethodQuery               = "SELECT sortmethod FROM libraries WHERE id=?"
	updateSortMethodQuery            = "UPDATE libraries SET sortmethod=? WHERE id=?"

	SVG_PATH = "/caseimages"
	CASE_MARGIN = 50
)

var logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)

//Break is a case break
type Break struct {
	LibraryID string `json:"libaryid"`
	ValueType string `json:"valuetype"`
	Value     string `json:"value"`
	BreakType string `json:"breaktype"`
}

//EditedCases is a set of edited cases and ids to remove
type EditedCases struct {
	EditedCases []EditedCase `json:"editedcases"`
	ToRemoveIds []int64      `json:"toremoveids"`
	LibraryID   int64        `json:"libraryid"`
}

//EditedCase is an edited case
type EditedCase struct {
	ID              int64 `json:"id"`
	SpacerHeight    int64 `json:"spacerheight"`
	PaddingLeft     int64 `json:"paddingleft"`
	PaddingRight    int64 `json:"paddingright"`
	Width           int64 `json:"width"`
	ShelfHeight     int64 `json:"shelfheight"`
	NumberOfShelves int64 `json:"numberofshelves"`
	CaseNumber      int64 `json:"casenumber"`
}

//Bookcase is a bookcase
type Bookcase struct {
	ID                int64         `json:"id"`
	SpacerHeight      int64         `json:"spacerheight"`
	PaddingLeft       int64         `json:"paddingleft"`
	PaddingRight      int64         `json:"paddingright"`
	BookMargin        int64         `json:"bookmargin"`
	Width             int64         `json:"width"`
	Shelves           []Bookshelf   `json:"shelves"`
	AverageBookHeight float64       `json:"averagebookheight"`
	AverageBookWidth  float64       `json:"averagebookwidth"`
	Library           books.Library `json:"library"`
	CaseNumber        int64         `json:"casenumber"`
}

//Bookshelf is a shelf on a bookcase
type Bookshelf struct {
	ID          int64        `json:"id"`
	CaseID      int64        `json:"caseid"`
	ShelfNumber int64        `json:"shelfnumber"`
	Width       int64        `json:"width"`
	Height      int64        `json:"height"`
	Alignment   string       `json:"alignment"`
	Books       []books.Book `json:"books"`
}

//OwnedLibrary is a library owned by the user and the users that have permission to view them
type OwnedLibrary struct {
	ID    int64                `json:"id"`
	Name  string               `json:"name"`
	Users []UserWithPermission `json:"user"`
}

//UserWithPermission is a user and permission
type UserWithPermission struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	FullName   string `json:"fullname"`
	Email      string `json:"email"`
	IconURL    string `json:"iconurl"`
	Permission int64  `json:"permission"`
}

/*
//GetCases gets cases
func GetCases(db *sql.DB, libraryid, session string) ([]Bookcase, error) {
	authorseries, err := GetAuthorBasedSeries(db, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	sortMethod, err := GetLibrarySort(db, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	books, _, err := books.GetBooks(db, sortMethod, "both", "both", "yes", "both", "both", "both", "", "1", "-1", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", libraryid, "", "both", session, authorseries, true)
	breaks, err := GetLibraryBreaks(db, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	rows, err := db.Query(getCasesQuery, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	dim, err := information.GetDimensions(db, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	var cases []Bookcase
	for rows.Next() {
		var id, width, spacerHeight, paddingLeft, paddingRight, bookMargin, caseNumber, numberOfShelves, shelfHeight int64
		err = rows.Scan(&id, &width, &spacerHeight, &paddingLeft, &paddingRight, &bookMargin, &caseNumber, &numberOfShelves, &shelfHeight)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		avgWidth := 1
		avgHeight := 25
		if dim["averagewidth"] > 0 {
			avgWidth = int(dim["averagewidth"])
		}
		//Leave room for text
		if dim["averageheight"] > 25 {
			avgHeight = int(dim["averageheight"])
		}
		bookcase := Bookcase{
			ID:                id,
			Width:             width,
			SpacerHeight:      spacerHeight,
			PaddingLeft:       paddingLeft,
			PaddingRight:      paddingRight,
			BookMargin:        bookMargin,
			AverageBookWidth:  float64(avgWidth),
			AverageBookHeight: float64(avgHeight),
			CaseNumber:        caseNumber,
		}
		for i := 0; i < int(numberOfShelves); i++ {
			bookcase.Shelves = append(bookcase.Shelves, Bookshelf{
				ID:     0,
				Height: shelfHeight,
			})
		}
		cases = append(cases, bookcase)
	}
	index := 0
	x := 0
	for c, bookcase := range cases {
		breakcase := false
		for s := range bookcase.Shelves {
			if breakcase {
				break
			}
			x = int(bookcase.PaddingLeft)
			useWidth := int(bookcase.AverageBookWidth)
			if index < len(books) && books[index].Width > 0 {
				useWidth = int(books[index].Width)
			}
			breakshelf := false
			for index < len(books) && useWidth+x <= int(bookcase.Width)-int(bookcase.PaddingRight) {
				if breakshelf || breakcase {
					break
				}
				cases[c].Shelves[s].Books = append(cases[c].Shelves[s].Books, books[index])
				x += useWidth
				index++
				useWidth = int(bookcase.AverageBookWidth)
				if index < len(books) && books[index].Width != 0 {
					useWidth = int(books[index].Width)
				}
				breakInner := false
				for _, b := range breaks {
					if breakInner {
						break
					}
					switch b.ValueType {
					case "ID":
						if b.Value == books[index].ID {
							breakInner = true
							if b.BreakType == "CASE" {
								breakcase = true
							} else {
								breakshelf = true
							}
							break
						}
					case "DEWEY":
                        if index != 0 && books[index-1].Dewey < b.Value && books[index].Dewey >= b.Value {
							breakInner = true
							if b.BreakType == "CASE" {
								breakcase = true
							} else {
								breakshelf = true
							}
							break
						}
					case "SERIES":
						if index != 0 && books[index-1].Series < b.Value && books[index].Series >= b.Value {
							breakInner = true
							if b.BreakType == "CASE" {
								breakcase = true
							} else {
								breakshelf = true
							}
							break
						}
					}
				}
			}
		}
	}
	return cases, nil
}
*/

//GetCases gets cases
func GetCases(db *sql.DB, libraryid, session string) ([]Bookcase, error) {
	authorseries, err := GetAuthorBasedSeries(db, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	sortMethod, err := GetLibrarySort(db, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	books, _, err := books.GetBooksCases(db, sortMethod, "both", "both", "yes", "both", "both", "both", "", "1", "-1", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", libraryid, "", "both", session, authorseries)
	breaks, err := GetLibraryBreaks(db, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	rows, err := db.Query(getCasesQuery, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	dim, err := information.GetDimensions(db, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	var cases []Bookcase
	for rows.Next() {
		var id, caseNumber, paddingLeft, paddingRight, bookMargin, spacerHeight int64
		err = rows.Scan(&id, &caseNumber, &paddingLeft, &paddingRight, &bookMargin, &spacerHeight)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		avgWidth := 1
		avgHeight := 25
		if dim["averagewidth"] > 0 {
			avgWidth = int(dim["averagewidth"])
		}
		//Leave Room for Text
		if dim["averageheight"] > 25 {
			avgHeight = int(dim["averageheight"])
		}
		bookcase := Bookcase{
			ID:                id,
			SpacerHeight:      spacerHeight,
			PaddingLeft:       paddingLeft,
			PaddingRight:      paddingRight,
			BookMargin:        bookMargin,
			AverageBookWidth:  float64(avgWidth),
			AverageBookHeight: float64(avgHeight),
			CaseNumber:        caseNumber,
		}
		shelfRows, err := db.Query(getShelvesQuery, id)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		for shelfRows.Next() {
			var shelfid, caseid, shelfnumber, width, height int64
			var alignment string
			err = shelfRows.Scan(&shelfid, &caseid, &shelfnumber, &width, &height, &alignment)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return nil, err
			}
			bookcase.Shelves = append(bookcase.Shelves, Bookshelf{
				ID:          shelfid,
				CaseID:      caseid,
				ShelfNumber: shelfnumber,
				Width:       width,
				Height:      height,
				Alignment:   alignment,
			})
		}
		cases = append(cases, bookcase)
	}
	index := 0
	x := 0
	for c, bookcase := range cases {
		breakcase := false
		for s := range bookcase.Shelves {
			if breakcase {
				break
			}
			x = int(bookcase.PaddingLeft)
			useWidth := int(bookcase.AverageBookWidth)
			if index < len(books) && books[index].Width > 0 {
				useWidth = int(books[index].Width)
			}
			breakshelf := false
			for index < len(books) && useWidth+x <= int(cases[c].Shelves[s].Width)-int(bookcase.PaddingRight) {
				if breakshelf || breakcase {
					break
				}
				cases[c].Shelves[s].Books = append(cases[c].Shelves[s].Books, books[index])
				x += useWidth
				index++
				useWidth = int(bookcase.AverageBookWidth)
				if index < len(books) && books[index].Width != 0 {
					useWidth = int(books[index].Width)
				}
				breakInner := false
				for _, b := range breaks {
					if breakInner {
						break
					}
					switch b.ValueType {
					case "ID":
						if b.Value == books[index].ID {
							breakInner = true
							if b.BreakType == "CASE" {
								breakcase = true
							} else {
								breakshelf = true
							}
							break
						}
					case "DEWEY":
						if index < len(books)-1 && books[index-1].Dewey < b.Value && books[index].Dewey >= b.Value {
							breakInner = true
							if b.BreakType == "CASE" {
								breakcase = true
							} else {
								breakshelf = true
							}
							break
						}
					case "SERIES":
						if index != len(books)-2 && books[index-1].Series < b.Value && books[index].Series >= b.Value {
							breakInner = true
							if b.BreakType == "CASE" {
								breakcase = true
							} else {
								breakshelf = true
							}
							break
						}
					}
				}
			}
		}
	}
	return cases, nil
}

//GetLibraries gets the libraries available to a user
func GetLibraries(db *sql.DB, session string) ([]books.Library, error) {
	var libraries []books.Library
	rows, err := db.Query(getLibrariesQuery, session)
	if err != nil {
		logger.Printf("%+v", err)
		return nil, err
	}
	for rows.Next() {
		var l books.Library
		if err := rows.Scan(&l.ID, &l.Name, &l.Permissions, &l.Owner); err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		libraries = append(libraries, l)
	}
	return libraries, nil
}

//GetLibrary gets the library by its id
func GetLibrary(db *sql.DB, session string, libraryid string) (books.Library, error) {
	var library books.Library
	row := db.QueryRow(getLibraryQuery, libraryid, session)
	if err := row.Scan(&library.ID, &library.Name, &library.Permissions, &library.Owner); err != nil {
		logger.Printf("Error: %+v", err)
		return library, err
	}
	return library, nil
}

//SaveCases saves book cases
func SaveCases(db *sql.DB, libraryid string, cases EditedCases) error {
	for _, c := range cases.EditedCases {
		id := c.ID
		caseNumber := c.CaseNumber
		width := c.Width
		spacerHeight := c.SpacerHeight
		paddingLeft := c.PaddingLeft
		paddingRight := c.PaddingRight
		shelfHeight := c.ShelfHeight
		numberOfShelves := c.NumberOfShelves
		if id == 0 {
			_, err := db.Exec(addCaseQuery, caseNumber, width, spacerHeight, paddingLeft, paddingRight, libraryid, numberOfShelves, shelfHeight)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		} else {
			_, err := db.Exec(updateCaseQuery, caseNumber, width, spacerHeight, paddingLeft, paddingRight, libraryid, numberOfShelves, shelfHeight, id)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		}
	}
	var removedCases []string
	for _, id := range cases.ToRemoveIds {
		i := strconv.FormatInt(id, 10)
		removedCases = append(removedCases, i)
	}
	if len(removedCases) > 0 {
		query := "DELETE FROM bookcases WHERE CaseId IN (" + strings.Join(removedCases, ",") + ")"
		_, err := db.Exec(query)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
	}
	return nil
}

//UpdateBreaks updates shelf breaks
func UpdateBreaks(db *sql.DB, libraryid string, breaks []Break) error {
	_, err := db.Exec(deleteBreaksQuery, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	for _, b := range breaks {
		_, err = db.Exec(addBreakQuery, libraryid, b.BreakType, b.ValueType, b.Value)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
	}
	return nil
}

//GetLibraryBreaks gets shelf breaks
func GetLibraryBreaks(db *sql.DB, libraryid string) ([]Break, error) {
	var breaks []Break
	rows, err := db.Query(getBreaksQuery, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	for rows.Next() {
		var b Break
		b.LibraryID = libraryid
		err = rows.Scan(&b.BreakType, &b.ValueType, &b.Value)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		breaks = append(breaks, b)
	}
	return breaks, nil
}

//GetOwnedLibraries gets owned libraries and the people who have permission to do things with them
func GetOwnedLibraries(db *sql.DB, session string) ([]OwnedLibrary, error) {
	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	var libraries []OwnedLibrary
	rows, err := db.Query(getOwnedLibrariesQuery, userid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	for rows.Next() {
		var library OwnedLibrary
		err = rows.Scan(&library.ID, &library.Name)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		innerRows, err := db.Query(getLibraryMembersPermissionQuery, library.ID, userid)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		for innerRows.Next() {
			var user UserWithPermission
			err := innerRows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.IconURL, &user.Permission)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return nil, err
			}
			user.FullName = user.FirstName + " " + user.LastName
			library.Users = append(library.Users, user)
		}
		libraries = append(libraries, library)
	}
	return libraries, nil
}

//SaveOwnedLibraries saves owned libraries and the people who have permission to do things with them
func SaveOwnedLibraries(db *sql.DB, ownedLibraries []OwnedLibrary, session string) error {
	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	_, err = db.Exec(deleteLibrariesQuery, userid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	for _, ownedLibrary := range ownedLibraries {
		oldID := ownedLibrary.ID
		res, err := db.Exec(addLibraryQuery, ownedLibrary.Name, userid, information.SORTMETHOD)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		ownedLibrary.ID, err = res.LastInsertId()
		_, err = db.Exec(deletePermissionsQuery, ownedLibrary.ID)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		_, err = db.Exec(deletePermissionsQuery, oldID)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		if oldID != -1 {
			_, err = db.Exec(updateBooksLibraryQuery, ownedLibrary.ID, oldID)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
			_, err = db.Exec(updateCasesLibraryQuery, ownedLibrary.ID, oldID)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
			_, err = db.Exec(updateBreaksLibraryQuery, ownedLibrary.ID, oldID)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
			_, err = db.Exec(updateSortsLibraryQuery, ownedLibrary.ID, oldID)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		}
		_, err = db.Exec(addPermissionQuery, userid, ownedLibrary.ID, 7)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		for _, user := range ownedLibrary.Users {
			_, err = db.Exec(addPermissionQuery, user.ID, ownedLibrary.ID, user.Permission)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		}
	}
	rows, err := db.Query(getBooksLibrariesQuery)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	var booklibraryids []int64
	for rows.Next() {
		var lid int64
		err := rows.Scan(&lid)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		booklibraryids = append(booklibraryids, lid)
	}
	rows, err = db.Query(getPermissionsLibrariesQuery)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	var permissionslibraryids []int64
	for rows.Next() {
		var lid int64
		err := rows.Scan(&lid)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		permissionslibraryids = append(permissionslibraryids, lid)
	}
	rows, err = db.Query(getCasesLibrariesQuery)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	var bookcaseslibraryids []int64
	for rows.Next() {
		var lid int64
		err := rows.Scan(&lid)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		bookcaseslibraryids = append(bookcaseslibraryids, lid)
	}
	rows, err = db.Query(getBreaksLibrariesQuery)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	var breakslibraryids []int64
	for rows.Next() {
		var lid int64
		err := rows.Scan(&lid)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		breakslibraryids = append(breakslibraryids, lid)
	}
	rows, err = db.Query(getSeriesLibrariesQuery)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	var serieslibraryids []int64
	for rows.Next() {
		var lid int64
		err := rows.Scan(&lid)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		serieslibraryids = append(serieslibraryids, lid)
	}
	rows, err = db.Query(getLibraryIdsQuery)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	var actuallibraryids []int64
	for rows.Next() {
		var lid int64
		err := rows.Scan(&lid)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
		actuallibraryids = append(actuallibraryids, lid)
	}
	for _, id := range booklibraryids {
		if !contains(actuallibraryids, id) {
			_, err := db.Exec(deleteBooksLibraryQuery, id)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		}
	}
	for _, id := range permissionslibraryids {
		if !contains(actuallibraryids, id) {
			_, err := db.Exec(deletePermissionsLibraryQuery, id)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		}
	}
	for _, id := range breakslibraryids {
		if !contains(actuallibraryids, id) {
			_, err := db.Exec(deleteBreaksLibraryQuery, id)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		}
	}
	for _, id := range bookcaseslibraryids {
		if !contains(actuallibraryids, id) {
			_, err := db.Exec(deleteCasesLibraryQuery, id)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		}
	}
	for _, id := range serieslibraryids {
		if !contains(actuallibraryids, id) {
			_, err := db.Exec(deleteSeriesLibraryQuery, id)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return err
			}
		}
	}
	return nil
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//GetAuthorBasedSeries gets series that are sorted by author then title, instead of volume
func GetAuthorBasedSeries(db *sql.DB, libraryid string) ([]string, error) {
	var series []string
	rows, err := db.Query(getAuthorBasedSeriesQuery, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		series = append(series, s)
	}
	return series, nil
}

//UpdateAuthorBasedSeries updates series that are sorted by author then title, instead of volume
func UpdateAuthorBasedSeries(db *sql.DB, libraryid string, series []string) error {
	_, err := db.Exec(updateAuthorBasedSeriesQuery, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	for _, s := range series {
		_, err = db.Exec(addAuthorBasedSeriesQuery, libraryid, s)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
	}
	return nil
}

//SearchLocation is a location in a library
type SearchLocation struct {
	Case  int    `json:"case"`
	Shelf int    `json:"shelf"`
	Book  int    `json:"book"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

//GetLibrarySort gets the sort method of a library
func GetLibrarySort(db *sql.DB, libraryid string) (string, error) {
	var method string
	err := db.QueryRow(getSortMethodQuery, libraryid).Scan(&method)
	return method, err
}

//UpdateLibrarySort updates the sort method of a library
func UpdateLibrarySort(db *sql.DB, libraryid string, method string) error {
	_, err := db.Exec(updateSortMethodQuery, method, libraryid)
	return err
}

//SearchShelves gets the locations of books on shelves by a search
func SearchShelves(db *sql.DB, libraryid string, session string, text string) ([]SearchLocation, error) {
	var locations []SearchLocation
	cases, err := GetCases(db, libraryid, session)
	if err != nil {
		return locations, err
	}
	for cidx, c := range cases {
		for sidx, s := range c.Shelves {
			for bidx, b := range s.Books {
				if IsMatch(b, strings.ToLower(text)) {
					locations = append(locations, SearchLocation{
						Case:  cidx,
						Shelf: sidx,
						Book:  bidx,
						ID:    b.ID,
						Title: b.Title,
					})
				}
			}
		}
	}
	return locations, nil
}

//IsMatch determines if a book matches a search string
func IsMatch(book books.Book, query string) bool {
	return strings.Contains(strings.ToLower(book.Title), query) || strings.Contains(strings.ToLower(book.Subtitle), query) || strings.Contains(strings.ToLower(book.Series), query)
}

//GetCaseIDs gets a list of the case ids for a library in case order
func GetCaseIDs(db *sql.DB, libraryid string, session string) ([]int, error) {
	var values []int
	rows, err := db.Query(getCaseIdsQuery, libraryid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	for rows.Next() {
		var v int
		err = rows.Scan(&v)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		values = append(values, v)
	}
	return values, nil
}

//GetShelfIDs gets a list of the shelf ids for a case in shelf order {
func GetShelfIDs(db *sql.DB, libraryid string, session string, caseid string) ([]int, error) {
	var values []int
	rows, err := db.Query(getShelfIdsQuery, caseid)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	for rows.Next() {
		var v int
		err = rows.Scan(&v)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return nil, err
		}
		values = append(values, v)
	}
	return values, nil
}

//RefreshCases refreshes case and shelf images {
func RefreshCases(db *sql.DB, libraryid string, session string, resRoot string) error {
	// Grab all the cases
    cases, err := GetCases(db, libraryid, session)
    if err != nil {
    	logger.Printf("Error: %s", err)
        return err
    }

    CaseSvgPath := resRoot + SVG_PATH

    maxCaseHeight := 0
    for _, c := range cases {
    	height := int(c.SpacerHeight) * (len(c.Shelves) + 1)
    	for _, s := range(c.Shelves) {
    		height += int(s.Height)
    	}
    	if maxCaseHeight < height {
    		maxCaseHeight = height
    	}
    }
    maxCaseHeight += CASE_MARGIN * 2

    // For each case in library
    for _, c := range cases {

        // Open case file
        cf, err := os.OpenFile(fmt.Sprintf("%s/%d.svg", CaseSvgPath, c.ID), os.O_CREATE|os.O_RDWR, 0644)
        if err != nil {
        	logger.Printf("Error: %s", err)
            return err
        }
        cf.Truncate(0)
        cf.Seek(0, 0)

        // Get max shelf width (only books and book margins)
        // Also get total shelf heights (interior only)
        maxShelfWidthBooks := 0
        totalCaseHeight := 0
        for _, s := range c.Shelves {
            width := int(s.Width)
            if width > maxShelfWidthBooks {
                maxShelfWidthBooks = width
            }
            totalCaseHeight += int(s.Height)
        }

        // Get case width
        caseWidth := int(c.SpacerHeight) * 2
        caseWidth += maxShelfWidthBooks + CASE_MARGIN * 2

        // Get case height
        caseHeight := int(c.SpacerHeight) * (len(c.Shelves) + 1)
        caseHeight += totalCaseHeight

        // Start case canvas
        caseCanvas := svg.New(cf)
        caseCanvas.Startview(caseWidth, maxCaseHeight, 0, 0, caseWidth, maxCaseHeight)

        // Set current y
        y := (maxCaseHeight - caseHeight)

        // For each shelf in case
        for _, s := range c.Shelves {

            // Set current x
            x := CASE_MARGIN
            if s.Alignment == "right" {
                x = caseWidth - (int(s.Width) + 2 * int(c.SpacerHeight)) - CASE_MARGIN
            }

            //Add the left and right shelf borders
            caseCanvas.Rect(x, y, int(c.SpacerHeight), int(s.Height + c.SpacerHeight))
            caseCanvas.Rect(x + int(c.SpacerHeight) + int(s.Width), y, int(c.SpacerHeight), int(s.Height + c.SpacerHeight))

            // Add the top border
            caseCanvas.Rect(x + int(c.SpacerHeight), y, int(s.Width), int(c.SpacerHeight))

            // Update the current y to bottom of current shelf
            y += int(c.SpacerHeight + s.Height)

            // Add the bottom border if last shelf
            caseCanvas.Rect(x, y, int(s.Width) + int(c.SpacerHeight) * 2, int(c.SpacerHeight))

            // Update current x to inside of shelf
            x += int(c.SpacerHeight + c.PaddingLeft)

            // Open shelf file
            if _, err := os.Stat(fmt.Sprintf("%s/%d/", CaseSvgPath, c.ID)); os.IsNotExist(err) {
                os.MkdirAll(fmt.Sprintf("%s/%d/", CaseSvgPath, c.ID), 0700)
            }
            sf, err := os.OpenFile(fmt.Sprintf("%s/%d/%d.svg", CaseSvgPath, c.ID, s.ID), os.O_CREATE|os.O_RDWR, 0644)
            if err != nil {
            	logger.Printf("Error: %s", err)
                return err
            }
            sf.Truncate(0)
            sf.Seek(0, 0)

            // For each book on shelf
            for _, b := range s.Books {
                
                // Fix Font Color if too short
                for len(b.SpineColor) < 7 {
                    b.SpineColor += "0"
                }

                // Set default width/height if neccessary
                if b.Width <= 0 {
                	b.Width = int64(c.AverageBookWidth)
                }
                if b.Height <= 0 {
                	b.Height = int64(c.AverageBookHeight)
                }

                // Draw Book
                caseCanvas.Rect(x, y - int(b.Height), int(b.Width), int(b.Height), fmt.Sprintf("id='book-%s'", b.ID), "class='bookcase-book'", "stroke='black'", fmt.Sprintf("fill='%s'", b.SpineColor))

                // Get Font Size
                fontSize := int(math.Min(18, float64(b.Width) / 4 * 3))

                // Get Font Color
                var fontColor string
                var c color.RGBA
                c.A = 0xff
                _, err = fmt.Sscanf(b.SpineColor, "#%02x%02x%02x", &c.R, &c.G, &c.B)
                if err != nil {
                    c.A = 0x00
                }
                if c.A == 0xff {
                    o := math.Round(float64(int(c.R) * 299 + int(c.G) * 587 + int(c.B) * 114) / 1000)
                    if o > 125 {
                        fontColor = "black"
                    } else {
                        fontColor = "white"
                    }
                } else {
                    fontColor = "white"
                }

                // Draw Title
                caseCanvas.TranslateRotate(x + int(b.Width) / 2, y, -90)
                caseCanvas.Path(fmt.Sprintf("M 2 0 L %d 0", int(b.Height) - 4), fmt.Sprintf(`id="PATH%s"`, b.ID))
                // caseCanvas.Use(0, 0, fmt.Sprintf("#PATH%s", b.ID))
                caseCanvas.Textpath(b.Title, fmt.Sprintf("#PATH%s", b.ID), fmt.Sprintf("dominant-baseline:middle;font-family:Arial;font-size:%dpx;fill:%s", fontSize, fontColor))
                caseCanvas.Gend()

                // Move x to start of next book
                x += int(b.Width)

            }

            // Close shelf file
            if err := sf.Close(); err != nil {
            	logger.Printf("Error: %s", err)
                return err
            }
        }

        //End case canvas
        caseCanvas.End()

        // Close case file
        if err := cf.Close(); err != nil {
        	logger.Printf("Error: %s", err)
            return err
        }

    }

    return nil
}