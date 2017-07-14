package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	getBooksQuery        = "SELECT * from books"
	getContributorsQuery = "SELECT PersonID, Role, FirstName, MiddleNames, LastName from written_by join persons on written_by.AuthorID = persons.PersonID WHERE BookID=?"
	getPublisherQuery    = "SELECT * from publishers WHERE PublisherID=?"
	getPublishersQuery   = "SELECT DISTINCT(Publisher) from publishers"
	getCitiesQuery       = "SELECT DISTINCT(City) from publishers"
	getStatesQuery       = "SELECT DISTINCT(State) from publishers"
	getCountriesQuery    = "SELECT DISTINCT(Country) from publishers"
	getSeriesQuery       = "SELECT DISTINCT(Series) from series"
	getFormatsQuery      = "SELECT DISTINCT(Format) from formats"
	getDeweysQuery       = "SELECT DISTINCT(Number) from dewey_numbers"
	getLanguagesQuery    = "SELECT DISTINCT(Langauge) from languages"
	getRolesQuery        = "SELECT DISTINCT(Role) from written_by"
	saveBookQuery        = "UPDATE books SET Title=?, Subtitle=?, OriginallyPublished=?, EditionPublished=?, PublisherID=?, IsRead=?, IsReference=?, IsOwned=?, IsShipping=?, IsReading=?, isbn=?, LoaneeFirst=?, LoaneeLast=?, Dewey=?, Pages=?, Width=?, Height=?, Depth=?, Weight=?, PrimaryLanguage=?, SecondaryLanguage=?, OriginalLanguage=?, Series=?, Volume=?, Format=?, Edition=?, ImageURL=? WHERE BookId=?"
	addBookQuery         = "INSERT INTO books (Title, Subtitle, OriginallyPublished, PublisherID, IsRead, IsReference, IsOwned, IsShipping, IsReading, isbn, LoaneeFirst, LoaneeLast, Dewey, Pages, Width, Height, Depth, Weight, PrimaryLanguage, SecondaryLanguage, OriginalLanguage, Series, Volume, Format, Edition, EditionPublished) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	getValidUserSession  = "SELECT sessionkey from usersession WHERE sessionkey=? AND EXISTS (SELECT id FROM library_members where id=userid)"
	getUser              = "SELECT id from library_members WHERE usr=? and pass=?"
	addUser              = "INSERT INTO library_members (usr,pass,email) values (?,?,?)"
	addSession           = "INSERT INTO usersession (sessionkey,userid,LastSeenTime) values (?,?,NOW())"
	updateSessionTime    = "UPDATE usersession SET LastSeenTime=NOW()"
	isSessionNameTaken   = "SELECT sessionkey from usersession where sessionkey=?"
	deleteSession        = "DELETE FROM usersession WHERE sessionkey=?"

	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

//IsRegistered returns whether a user session is registered and the corresponding user exists
func IsRegistered(session string) (bool, error) {
	var sessionkey string
	if err := db.QueryRow(getValidUserSession, session).Scan(&sessionkey); err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, err
	} else {
		return false, err
	}
}

//MarkAsSeen marks a user session as seen
func MarkAsSeen(session string) error {
	_, err := db.Exec(updateSessionTime, session)
	if err != nil {
		return err
	}
	return nil
}

//IsUser returns whether a username and password is valid. If they are, return the userid. If not, return -1
func IsUser(username, password string) (int64, error) {
	var id int64
	if err := db.QueryRow(getUser, username, password).Scan(&id); err == nil {
		return id, nil
	} else if err == sql.ErrNoRows {
		return -1, nil
	} else {
		return -1, err
	}
}

//GenerateSessionKey generates a random unique session key
func GenerateSessionKey() string {
	b := make([]byte, 50)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

//LoginUser logs in a user
func LoginUser(username, password string) (string, error) {
	id, err := IsUser(username, password)
	if err != nil {
		return "", err
	}
	key := GenerateSessionKey()
	_, err = db.Exec(addSession, key, id)
	return key, err
}

//RegisterUser registers a user
func RegisterUser(username, password, email string) (string, error) {
	result, err := db.Exec(addUser, username, password, email)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	key := GenerateSessionKey()
	_, err = db.Exec(addSession, key, id)
	return key, err
}

//LogoutSession logs out a user
func LogoutSession(sessionkey string) error {
	_, err := db.Exec(deleteSession, sessionkey)
	if err != nil {
		return err
	}
	return nil
}

//SaveBook saves a book
func SaveBook(book Book) error {
	if book.ID != "" {
		imageType := filepath.Ext(book.ImageURL)
		if book.ImageURL != "" && !strings.HasPrefix(book.ImageURL, "res/bookimages/") {
			err := downloadImage(book.ImageURL, "../web/res/bookimages/"+book.ID+imageType)
			if err != nil {
				log.Printf("Error while saving image: %v", err)
				return err
			}
		}
		publisherID, err := addOrGetPublisher(book.Publisher.Publisher, book.Publisher.City, book.Publisher.State, book.Publisher.Country)
		if err != nil {
			log.Printf("Error when saving publisher: %v", err)
			return err
		}
		err = addDewey(book.Dewey)
		if err != nil {
			log.Printf("Error when saving dewey: %v", err)
			return err
		}
		err = addSeries(book.Series)
		if err != nil {
			log.Printf("Error when saving Series: %v", err)
			return err
		}
		err = addFormat(book.Format)
		if err != nil {
			log.Printf("Error when saving Format: %v", err)
			return err
		}
		err = addLanguage(book.PrimaryLanguage)
		if err != nil {
			log.Printf("Error when saving PrimaryLanguage: %v", err)
			return err
		}
		err = addLanguage(book.SecondaryLanguage)
		if err != nil {
			log.Printf("Error when saving SecondaryLanguage: %v", err)
			return err
		}
		err = addLanguage(book.OriginalLanguage)
		if err != nil {
			log.Printf("Error when saving OriginalLanguage: %v", err)
			return err
		}
		_, err = db.Exec(saveBookQuery, book.Title, book.Subtitle, book.OriginallyPublished, book.EditionPublished, publisherID, book.IsRead, book.IsReference, book.IsOwned, book.IsShipping, book.IsReading, book.ISBN, book.Loanee.First, book.Loanee.Last, book.Dewey, book.Pages, book.Width, book.Height, book.Depth, book.Weight, book.PrimaryLanguage, book.SecondaryLanguage, book.OriginalLanguage, book.Series, book.Volume, book.Format, book.Edition, "res/bookimages/"+book.ID+imageType, book.ID)
		if err != nil {
			log.Printf("Error when saving book: %v", err)
			return err
		}
		err = removeAllWrittenBy(book.ID)
		if err != nil {
			log.Printf("Error when saving authors: %v", err)
			return err
		}
		for _, contributor := range book.Contributors {
			err = addContributor(book.ID, contributor)
			if err != nil {
				log.Printf("Error when saving authors: %v", err)
				return err
			}
		}
	} else {
		publisherID, err := addOrGetPublisher(book.Publisher.Publisher, book.Publisher.City, book.Publisher.State, book.Publisher.Country)
		if err != nil {
			log.Printf("Error when saving publisher: %v", err)
			return err
		}
		err = addDewey(book.Dewey)
		if err != nil {
			log.Printf("Error when saving dewey: %v", err)
			return err
		}
		err = addSeries(book.Series)
		if err != nil {
			log.Printf("Error when saving Series: %v", err)
			return err
		}
		err = addFormat(book.Format)
		if err != nil {
			log.Printf("Error when saving Format: %v", err)
			return err
		}
		err = addLanguage(book.PrimaryLanguage)
		if err != nil {
			log.Printf("Error when saving PrimaryLanguage: %v", err)
			return err
		}
		err = addLanguage(book.SecondaryLanguage)
		if err != nil {
			log.Printf("Error when saving SecondaryLanguage: %v", err)
			return err
		}
		err = addLanguage(book.OriginalLanguage)
		if err != nil {
			log.Printf("Error when saving OriginalLanguage: %v", err)
			return err
		}
		res, err := db.Exec(addBookQuery, book.Title, book.Subtitle, book.OriginallyPublished, publisherID, book.IsRead, book.IsReference, book.IsOwned, book.IsShipping, book.IsReading, book.ISBN, book.Loanee.First, book.Loanee.Last, book.Dewey, book.Pages, book.Width, book.Height, book.Depth, book.Weight, book.PrimaryLanguage, book.SecondaryLanguage, book.OriginalLanguage, book.Series, book.Volume, book.Format, book.Edition, book.EditionPublished)
		if err != nil {
			log.Printf("Error when saving book: %v", err)
			return err
		}
		id, err := res.LastInsertId()
		bookid := strconv.FormatInt(id, 10)
		imageType := filepath.Ext(book.ImageURL)
		if book.ImageURL != "" {
			err = downloadImage(book.ImageURL, "../web/res/bookimages/"+bookid+imageType)
		}
		if err != nil {
			log.Printf("Error while saving image: %v", err)
			return err
		}
		imageQuery := "UPDATE books SET ImageURL='res/bookimages/" + bookid + imageType + "' WHERE bookid=?"
		_, err = db.Exec(imageQuery, bookid)
		if err != nil {
			log.Printf("Error when saving image: %v", err)
			return err
		}
		for _, contributor := range book.Contributors {
			err = addContributor(bookid, contributor)
			if err != nil {
				log.Printf("Error when saving authors: %v", err)
				return err
			}
		}
	}
	return nil
}

func removeAllWrittenBy(bookid string) error {
	query := "DELETE FROM written_by WHERE BookId=?"
	_, err := db.Exec(query, bookid)
	if err != nil {
		log.Printf("Error when deleting written_by: %v", err)
		return err
	}
	return nil
}

//DeleteBook deletes a book
func DeleteBook(bookid string) error {
	removeAllWrittenBy(bookid)
	query := "DELETE FROM books WHERE BookId=?"
	_, err := db.Exec(query, bookid)
	if err != nil {
		log.Printf("Error when deleting book: %v", err)
		return err
	}
	return nil
}

func addContributor(bookid string, contributor Contributor) error {
	personID, err := addOrGetPerson(contributor.Name.First, contributor.Name.Middles, contributor.Name.Last)
	if err != nil {
		return err
	}
	query := "REPLACE INTO written_by (BookID, AuthorID, Role) VALUES (?,?,?)"
	_, err = db.Exec(query, bookid, personID, contributor.Role)
	if err != nil {
		return err
	}
	return nil
}

func addOrGetPerson(first, middles, last string) (string, error) {
	var id int64
	id = -1
	query := "SELECT PersonID FROM persons WHERE FirstName=? AND MiddleNames=? AND LastName=?"
	rows, err := db.Query(query, first, middles, last)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return "", err
		}
	}
	if id == -1 {
		query = "INSERT INTO persons (FirstName, MiddleNames, LastName) VALUES (?,?,?)"
		res, err := db.Exec(query, first, middles, last)
		if err != nil {
			return "", err
		}
		id, err = res.LastInsertId()
		if err != nil {
			return "", err
		}
	}
	return strconv.FormatInt(id, 10), nil
}

func addOrGetPublisher(publisher, city, state, country string) (string, error) {
	var id int64
	id = -1
	query := "SELECT PublisherID FROM publishers WHERE Publisher=? AND City=? AND State=? AND Country=?"
	rows, err := db.Query(query, publisher, city, state, country)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return "", err
		}
	}
	if id == -1 {
		query = "INSERT INTO publishers (Publisher, City, State, Country) VALUES (?,?,?,?)"
		res, err := db.Exec(query, publisher, city, state, country)
		if err != nil {
			return "", err
		}
		id, err = res.LastInsertId()
		if err != nil {
			return "", err
		}
	}
	return strconv.FormatInt(id, 10), nil
}

func addDewey(v string) error {
	query := "REPLACE INTO dewey_numbers (Number) VALUES (?)"
	_, err := db.Exec(query, v)
	if err != nil {
		return err
	}
	return nil
}

func addFormat(v string) error {
	query := "REPLACE INTO formats (Format) VALUES (?)"
	_, err := db.Exec(query, v)
	if err != nil {
		return err
	}
	return nil
}

func addSeries(v string) error {
	query := "REPLACE INTO series (Series) VALUES (?)"
	_, err := db.Exec(query, v)
	if err != nil {
		return err
	}
	return nil
}

func addLanguage(v string) error {
	query := "REPLACE INTO languages (Langauge) VALUES (?)"
	_, err := db.Exec(query, v)
	if err != nil {
		return err
	}
	return nil
}

func downloadImage(url, fileLocation string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	file, err := os.Create(fileLocation)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

//GetBooks gets all books
//todo include authors in filter
func GetBooks(sortMethod, isread, isreference, isowned, isloaned, isreading, isshipping, text, page, numberToGet, fromDewey, toDewey string) ([]Book, int64, error) {
	var order string
	if sortMethod == "title" {
		order = "Title2, minname"
	} else if sortMethod == "series" {
		order = "if(Series2='' or Series2 is null,1,0), Series2, Volume, minname, Title2"
	} else {
		order = "Dewey, minname, Series2, Volume, Title2"
	}
	titlechange := "CASE WHEN Title LIKE 'The %%' THEN TRIM(SUBSTR(Title from 4)) ELSE CASE WHEN Title LIKE 'An %%' THEN TRIM(SUBSTR(Title from 3)) ELSE CASE WHEN Title LIKE 'A %%' THEN TRIM(SUBSTR(Title from 2)) ELSE Title END END END AS Title2"
	serieschange := "CASE WHEN Series LIKE 'The %%' THEN TRIM(SUBSTR(Series from 4)) ELSE CASE WHEN Series LIKE 'An %%' THEN TRIM(SUBSTR(Series from 3)) ELSE CASE WHEN Series LIKE 'A %%' THEN TRIM(SUBSTR(Series from 2)) ELSE Series END END END AS Series2"
	authors := "(SELECT  PersonID, AuthorRoles.BookID, concat(COALESCE(lastname,''),COALESCE(firstname,''),COALESCE(middlenames,'')) as name FROM persons JOIN (SELECT written_by.BookID, AuthorID FROM written_by WHERE Role='Author') AS AuthorRoles ON AuthorRoles.AuthorID = persons.PersonID ORDER BY name ) AS Authors"
	read := ""
	if isread == "yes" {
		read = "isread=1"
	} else if isread == "no" {
		read = "isread=0"
	}
	reference := ""
	if isreference == "yes" {
		reference = "isreference=1"
	} else if isreference == "no" {
		reference = "isreference=0"
	}
	owned := ""
	if isowned == "yes" {
		owned = "isowned=1"
	} else if isowned == "no" {
		owned = "isowned=0"
	}
	loaned := ""
	if isloaned == "yes" {
		isloaned = "LoaneeFirst IS NOT NULL OR LoaneeLast IS NOT NULL"
	} else if isread == "no" {
		isloaned = "LoaneeFirst IS NULL AND LoaneeLast IS NULL"
	}
	reading := ""
	if isreading == "yes" {
		reading = "isreading=1"
	} else if isreading == "no" {
		reading = "isreading=0"
	}
	shipping := ""
	if isshipping == "yes" {
		shipping = "isshipping=1"
	} else if isshipping == "no" {
		shipping = "isshipping=0"
	}
	startDewey := "Dewey >= '" + formatDewey(fromDewey) + "'"
	endDewey := "Dewey <= '" + formatDewey(toDewey) + "'"
	filter := "WHERE "
	if read != "" {
		filter = filter + read
	}
	if reference != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + reference
	}
	if owned != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + owned
	}
	if loaned != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + loaned
	}
	if reading != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + reading
	}
	if shipping != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + shipping
	}

	if startDewey != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startDewey
	}
	if endDewey != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endDewey
	}
	filterText := formFilterText(text)
	if filterText != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + filterText
	}
	if filter == "WHERE " || filter == "WHERE" {
		filter = ""
	}
	pag, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return nil, 0, err
	}
	ntg, err := strconv.ParseInt(numberToGet, 10, 64)
	if err != nil {
		return nil, 0, err
	}
	query := "SELECT bookid, title, subtitle, OriginallyPublished, PublisherID, isread, isreference, IsOwned, ISBN, LoaneeFirst, LoaneeLast, dewey, pages, width, height, depth, weight, PrimaryLanguage, SecondaryLanguage, OriginalLanguage, series, volume, format, Edition, ImageURL, IsReading, isshipping, SpineColor, CheapestNew, CheapestUsed, EditionPublished from (select books.*, " + titlechange + ", " + serieschange + ", min(name) as minname FROM books LEFT JOIN " + authors + " ON books.BookID = Authors.BookID " + filter + " GROUP BY books.BookID) i ORDER BY " + order
	if numberToGet != "-1" {
		query += " LIMIT " + numberToGet + " OFFSET " + strconv.FormatInt(((pag-1)*ntg), 10)
	}
	pageQuery := "SELECT count(bookid) from (select books.bookid, " + titlechange + ", " + serieschange + ", min(name) as minname FROM books LEFT JOIN " + authors + " ON books.BookID = Authors.BookID " + filter + " GROUP BY books.BookID) i"

	b := Book{}
	var books = make([]Book, 0)

	var PublisherID sql.NullString
	var IsRead int64
	var IsReference int64
	var IsOwned int64
	var IsShipping int64
	var IsReading int64
	var LoaneeFirst sql.NullString
	var LoaneeLast sql.NullString
	var Title sql.NullString
	var Subtitle sql.NullString
	var OriginallyPublished mysql.NullTime
	var Dewey sql.NullString
	var ISBN sql.NullString
	var PrimaryLanguage sql.NullString
	var SecondaryLanguage sql.NullString
	var OriginalLanguage sql.NullString
	var Series sql.NullString
	var Format sql.NullString
	var ImageURL sql.NullString
	var SpineColor sql.NullString
	var EditionPublished mysql.NullTime

	var p Publisher
	var c []Contributor

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying books: %v", err)
		return nil, 0, err
	}
	for rows.Next() {
		if err := rows.Scan(&b.ID, &Title, &Subtitle, &OriginallyPublished, &PublisherID, &IsRead, &IsReference, &IsOwned, &ISBN, &LoaneeFirst, &LoaneeLast, &Dewey, &b.Pages, &b.Width, &b.Height, &b.Depth, &b.Weight, &PrimaryLanguage, &SecondaryLanguage, &OriginalLanguage, &Series, &b.Volume, &Format, &b.Edition, &ImageURL, &IsReading, &IsShipping, &SpineColor, &b.CheapestNew, &b.CheapestUsed, &EditionPublished); err != nil {
			log.Printf("Error scanning books: %v", err)
			return nil, 0, err
		}
		if PublisherID.Valid {
			p, err = GetPublisher(PublisherID.String)
			if err != nil {
				log.Printf("Error getting publisher: %v", err)
				return nil, 0, err
			}
		}
		b.Publisher = p
		c, err = GetContributors(b.ID)
		if err != nil {
			log.Printf("Error getting contributors: %v", err)
			return nil, 0, err
		}
		b.Contributors = c
		b.Loanee = Name{
			First: "",
			Last:  "",
		}
		if LoaneeFirst.Valid {
			b.Loanee.First = LoaneeFirst.String
		}
		if LoaneeLast.Valid {
			b.Loanee.Last = LoaneeLast.String
		}
		b.IsOwned = IsOwned == 1
		b.IsReference = IsReference == 1
		b.IsReading = IsReading == 1
		b.IsRead = IsRead == 1
		b.IsShipping = IsShipping == 1
		b.Title = ""
		if Title.Valid {
			b.Title = Title.String
		}
		b.Subtitle = ""
		if Subtitle.Valid {
			b.Subtitle = Subtitle.String
		}
		b.OriginallyPublished = "0000"
		if OriginallyPublished.Valid {
			b.OriginallyPublished = strconv.Itoa(OriginallyPublished.Time.Year())
		}
		b.Dewey = ""
		if Dewey.Valid {
			b.Dewey = Dewey.String
		}
		b.ISBN = ""
		if ISBN.Valid {
			b.ISBN = ISBN.String
		}
		b.PrimaryLanguage = ""
		if PrimaryLanguage.Valid {
			b.PrimaryLanguage = PrimaryLanguage.String
		}
		b.SecondaryLanguage = ""
		if SecondaryLanguage.Valid {
			b.SecondaryLanguage = SecondaryLanguage.String
		}
		b.OriginalLanguage = ""
		if OriginalLanguage.Valid {
			b.OriginalLanguage = OriginalLanguage.String
		}
		b.Series = ""
		if Series.Valid {
			b.Series = Series.String
		}
		b.Format = ""
		if Format.Valid {
			b.Format = Format.String
		}
		b.ImageURL = ""
		if ImageURL.Valid {
			b.ImageURL = ImageURL.String
		}
		b.SpineColor = ""
		if SpineColor.Valid {
			b.SpineColor = SpineColor.String
		}
		b.EditionPublished = "0000"
		if EditionPublished.Valid {
			b.EditionPublished = strconv.Itoa(EditionPublished.Time.Year())
		}

		books = append(books, b)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error in rows: %v", err)
		return nil, 0, err
	}

	var numberOfBooks int64
	err = db.QueryRow(pageQuery).Scan(&numberOfBooks)
	if err != nil {
		return nil, 0, err
	}

	return books, numberOfBooks, rows.Close()
}

func formatDewey(dewey string) string {
	retval := ""
	i, err := strconv.ParseFloat(dewey, 64)
	if err != nil {
		if dewey == "FIC" {
			retval = dewey
		}
	} else {
		if i < 10 {
			dewey = "00" + dewey
		} else if i < 100 {
			dewey = "0" + dewey
		}
		retval = dewey
	}
	return retval
}

func formFilterText(text string) string {
	s := ""
	filters := strings.Split(text, " ")
	for _, filter := range filters {
		if filter != "" {
			s = s + "(Title LIKE '%" + filter + "%' OR Subtitle LIKE '%" + filter + "%' OR Series LIKE '%" + filter + "%') AND "
		}
	}
	if strings.HasSuffix(s, " AND ") {
		s = s[:len(s)-5]
	}
	return s
}

//GetContributors gets all contributors for a book id
func GetContributors(id string) ([]Contributor, error) {
	c := Contributor{}
	var contributors = make([]Contributor, 0)

	var Role sql.NullString
	var First sql.NullString
	var Middles sql.NullString
	var Last sql.NullString

	rows, err := db.Query(getContributorsQuery, id)
	if err != nil {
		log.Printf("Error querying contributors: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&c.ID, &Role, &First, &Middles, &Last); err != nil {
			log.Printf("Error scanning contributors: %v", err)
			return nil, err
		}
		c.Role = ""
		if Role.Valid {
			c.Role = Role.String
		}
		c.Name.First = ""
		if First.Valid {
			c.Name.First = First.String
		}
		c.Name.Middles = ""
		if Middles.Valid {
			c.Name.Middles = Middles.String
		}
		c.Name.Last = ""
		if Last.Valid {
			c.Name.Last = Last.String
		}
		contributors = append(contributors, c)
	}

	return contributors, rows.Close()
}

//GetPublisher gets the publisher for a publisher id
func GetPublisher(id string) (Publisher, error) {
	p := Publisher{}
	var Publisher sql.NullString
	var City sql.NullString
	var State sql.NullString
	var Country sql.NullString
	var ParentCompany sql.NullString

	err := db.QueryRow(getPublisherQuery, id).Scan(&p.ID, &Publisher, &City, &State, &Country, &ParentCompany)
	p.Publisher = ""
	if Publisher.Valid {
		p.Publisher = Publisher.String
	}
	p.City = ""
	if City.Valid {
		p.City = City.String
	}
	p.State = ""
	if State.Valid {
		p.State = State.String
	}
	p.Country = ""
	if Country.Valid {
		p.Country = Country.String
	}
	p.ParentCompany = ""
	if ParentCompany.Valid {
		p.ParentCompany = ParentCompany.String
	}
	if err != nil {
		log.Printf("Error scanning publisher for id %v: %v", id, err)
		return p, err
	}

	return p, nil
}

//GetPublishers gets all publishers
func GetPublishers() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getPublishersQuery)
	if err != nil {
		log.Printf("Error querying publishers: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning publishers: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetCities gets all cities
func GetCities() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getCitiesQuery)
	if err != nil {
		log.Printf("Error querying cities: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning cities: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetStates gets all states
func GetStates() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getStatesQuery)
	if err != nil {
		log.Printf("Error querying states: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning states: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetCountries gets all countries
func GetCountries() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getCountriesQuery)
	if err != nil {
		log.Printf("Error querying countries: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning countries: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetSeries gets all series
func GetSeries() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getSeriesQuery)
	if err != nil {
		log.Printf("Error querying series: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning series: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetFormats gets all formats
func GetFormats() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getFormatsQuery)
	if err != nil {
		log.Printf("Error querying formats: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning formats: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetLanguages gets all languages
func GetLanguages() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getLanguagesQuery)
	if err != nil {
		log.Printf("Error querying languages: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning languages: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetRoles gets all roles
func GetRoles() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getRolesQuery)
	if err != nil {
		log.Printf("Error querying roles: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning roles: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetDeweys gets all deweys
func GetDeweys() ([]string, error) {
	var s string
	var r = make([]string, 0)
	rows, err := db.Query(getDeweysQuery)
	if err != nil {
		log.Printf("Error querying deweys: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Printf("Error scanning deweys: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetBooksForExport selects all books as strings
func GetBooksForExport() ([][]string, error) {
	query := "select * from books join publishers on books.PublisherID=publishers.PublisherID"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error exporting books: %v", err)
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Printf("Error exporting books: %v", err)
		return nil, err
	}
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	var retval [][]string
	retval = append(retval, cols)
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			log.Printf("Error exporting books: %v", err)
			return nil, err
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
		toAppend := make([]string, len(result))
		copy(toAppend, result)
		retval = append(retval, toAppend)
	}
	return retval, nil
}

//GetAuthorsForExport selects all books as strings
func GetAuthorsForExport() ([][]string, error) {
	query := "SELECT BookID, FirstName, MiddleNames, LastName, Role from written_by JOIN persons on written_by.AuthorID=persons.PersonID"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error exporting books: %v", err)
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Printf("Error exporting books: %v", err)
		return nil, err
	}
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	var retval [][]string
	retval = append(retval, cols)
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			log.Printf("Error exporting books: %v", err)
			return nil, err
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
		toAppend := make([]string, len(result))
		copy(toAppend, result)
		retval = append(retval, toAppend)
	}
	return retval, nil
}

//ImportLibrary imports records from a csv file
//todo finish function
func ImportLibrary(records [][]string) error {
	log.Printf("Importing...")
	return nil
}

//GetStats gets statistics by type
func GetStats(t string) (StatChart, error) {
	var chart StatChart
	var data []StatData
	var query string
	switch t {
	case "generalbycounts":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		var read int64
		var reading int64
		var reference int64
		var loaned int64
		var shipping int64
		var toread int64
		query = `SELECT * FROM (
				(SELECT count(*) as total from books WHERE isowned=1) AS t,
				(SELECT count(*) as rea from books WHERE isread=1 and isowned=1) AS red,
				(SELECT count(*) as reading from books WHERE isreading=1 and isowned=1) AS rng,
				(SELECT count(*) as toread from books WHERE isread=0 and isreference=0 and isreading=0 and isowned=1) AS trd,
				(SELECT count(*) as reference from books WHERE isreference=1 and isowned=1) AS ref,
				(SELECT count(*) as loaned from books WHERE loaneelast is not null and loaneelast !='' and isowned=1) AS loa,
				(SELECT count(*) as shipping from books WHERE isshipping=1 and isowned=1) AS shi)`
		err := db.QueryRow(query).Scan(&total, &read, &reading, &toread, &reference, &loaned, &shipping)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Count (Total: " + strconv.FormatInt(total, 10) + ")"
		data = append(data, StatData{
			Label:    "Read",
			Value:    strconv.FormatInt(read, 10),
			ToolText: fmt.Sprintf("%.2f%%", float64(read)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Reading",
			Value:    strconv.FormatInt(reading, 10),
			ToolText: fmt.Sprintf("%.2f%%", float64(reading)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Reference",
			Value:    strconv.FormatInt(reference, 10),
			ToolText: fmt.Sprintf("%.2f%%", float64(reference)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "To Read",
			Value:    strconv.FormatInt(toread, 10),
			ToolText: fmt.Sprintf("%.2f%%", float64(toread)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Shipping",
			Value:    strconv.FormatInt(shipping, 10),
			ToolText: fmt.Sprintf("%.2f%%", float64(shipping)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Loaned",
			Value:    strconv.FormatInt(loaned, 10),
			ToolText: fmt.Sprintf("%.2f%%", float64(loaned)/float64(total)*100),
		})
	case "generalbysize":
		chart.Chart.NumberSuffix = " mm³"
		chart.Chart.Decimals = "0"
		var total int64
		var read int64
		var reading int64
		var reference int64
		var loaned int64
		var shipping int64
		var toread int64
		query = `SELECT * FROM (
				(SELECT SUM(width*height*depth) as total from books WHERE isowned=1) AS t,
				(SELECT SUM(width*height*depth) as rea from books WHERE isread=1 and isowned=1) AS red,
				(SELECT SUM(width*height*depth) as reading from books WHERE isreading=1 and isowned=1) AS rng,
				(SELECT SUM(width*height*depth) as toread from books WHERE isread=0 and isreference=0 and isreading=0 and isowned=1) AS trd,
				(SELECT SUM(width*height*depth) as reference from books WHERE isreference=1 and isowned=1) AS ref,
				(SELECT SUM(width*height*depth) as loaned from books WHERE loaneelast is not null and isowned=1) AS loa,
				(SELECT SUM(width*height*depth) as shipping from books WHERE isshipping=1 and isowned=1) AS shi)`
		err := db.QueryRow(query).Scan(&total, &read, &reading, &toread, &reference, &loaned, &shipping)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Size"
		data = append(data, StatData{
			Label:    "Read",
			Value:    fmt.Sprintf("%.2f", float64(read)),
			ToolText: fmt.Sprintf("%.2f%%", float64(read)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Reading",
			Value:    fmt.Sprintf("%.2f", float64(reading)),
			ToolText: fmt.Sprintf("%.2f%%", float64(reading)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Reference",
			Value:    fmt.Sprintf("%.2f", float64(reference)),
			ToolText: fmt.Sprintf("%.2f%%", float64(reference)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "To Read",
			Value:    fmt.Sprintf("%.2f", float64(toread)),
			ToolText: fmt.Sprintf("%.2f%%", float64(toread)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Shipping",
			Value:    fmt.Sprintf("%.2f", float64(shipping)),
			ToolText: fmt.Sprintf("%.2f%%", float64(shipping)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Loaned",
			Value:    fmt.Sprintf("%.2f", float64(loaned)),
			ToolText: fmt.Sprintf("%.2f%%", float64(loaned)/float64(total)*100),
		})
	case "generalbypages":
		chart.Chart.NumberSuffix = " pages"
		chart.Chart.Decimals = "0"
		var total int64
		var read int64
		var reading int64
		var reference int64
		var loaned int64
		var shipping int64
		var toread int64
		query = `SELECT * FROM (
				(SELECT SUM(pages) as total from books WHERE isowned=1) AS t,
				(SELECT SUM(pages) as rea from books WHERE isread=1 and isowned=1) AS red,
				(SELECT SUM(pages) as reading from books WHERE isreading=1 and isowned=1) AS rng,
				(SELECT SUM(pages) as toread from books WHERE isread=0 and isreference=0 and isreading=0 and isowned=1) AS trd,
				(SELECT SUM(pages) as reference from books WHERE isreference=1 and isowned=1) AS ref,
				(SELECT SUM(pages) as loaned from books WHERE loaneelast is not null and isowned=1) AS loa,
				(SELECT SUM(pages) as shipping from books WHERE isshipping=1 and isowned=1) AS shi)`
		err := db.QueryRow(query).Scan(&total, &read, &reading, &toread, &reference, &loaned, &shipping)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Pages"
		data = append(data, StatData{
			Label:    "Read",
			Value:    fmt.Sprintf("%.2f", float64(read)),
			ToolText: fmt.Sprintf("%.2f%%", float64(read)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Reading",
			Value:    fmt.Sprintf("%.2f", float64(reading)),
			ToolText: fmt.Sprintf("%.2f%%", float64(reading)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Reference",
			Value:    fmt.Sprintf("%.2f", float64(reference)),
			ToolText: fmt.Sprintf("%.2f%%", float64(reference)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "To Read",
			Value:    fmt.Sprintf("%.2f", float64(toread)),
			ToolText: fmt.Sprintf("%.2f%%", float64(toread)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Shipping",
			Value:    fmt.Sprintf("%.2f", float64(shipping)),
			ToolText: fmt.Sprintf("%.2f%%", float64(shipping)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Loaned",
			Value:    fmt.Sprintf("%.2f", float64(loaned)),
			ToolText: fmt.Sprintf("%.2f%%", float64(loaned)/float64(total)*100),
		})
	case "publishersbooksperparent":
		return chart, nil
	case "publisherstopchildren":
		return chart, nil
	case "publisherslocations":
		return chart, nil
	case "publisherstoplocations":
		return chart, nil
	case "series":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1`
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Series"
		series, err := GetSeries()
		if err != nil {
			return chart, err
		}
		for _, s := range series {
			var count int64
			seriesquery := `SELECT COUNT(*) FROM books WHERE series=? AND IsOwned=1`
			err := db.QueryRow(seriesquery, s).Scan(&count)
			if err != nil {
				return chart, err
			}
			if count > 0 && s != "" {
				data = append(data, StatData{
					Label:    s,
					Value:    fmt.Sprintf("%d", count),
					ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
				})
			}
		}
	case "languagesprimary":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1`
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Primary Language"
		languages, err := GetLanguages()
		if err != nil {
			return chart, err
		}
		for _, language := range languages {
			var count int64
			languagequery := `SELECT COUNT(*) FROM books WHERE PrimaryLanguage=? AND IsOwned=1`
			err := db.QueryRow(languagequery, language).Scan(&count)
			if err != nil {
				return chart, err
			}
			if count > 0 && language != "" {
				data = append(data, StatData{
					Label:    language,
					Value:    fmt.Sprintf("%d", count),
					ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
				})
			}
		}
	case "languagessecondary":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1`
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Secondary Language"
		languages, err := GetLanguages()
		if err != nil {
			return chart, err
		}
		for _, language := range languages {
			var count int64
			languagequery := `SELECT COUNT(*) FROM books WHERE SecondaryLanguage=? AND IsOwned=1`
			err := db.QueryRow(languagequery, language).Scan(&count)
			if err != nil {
				return chart, err
			}
			if count > 0 && language != "" {
				data = append(data, StatData{
					Label:    language,
					Value:    fmt.Sprintf("%d", count),
					ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
				})
			}
		}
	case "languagesoriginal":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1`
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Original Language"
		languages, err := GetLanguages()
		if err != nil {
			return chart, err
		}
		for _, language := range languages {
			var count int64
			languagequery := `SELECT COUNT(*) FROM books WHERE OriginalLanguage=? AND IsOwned=1`
			err := db.QueryRow(languagequery, language).Scan(&count)
			if err != nil {
				return chart, err
			}
			if count > 0 && language != "" {
				data = append(data, StatData{
					Label:    language,
					Value:    fmt.Sprintf("%d", count),
					ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
				})
			}
		}
	case "deweys":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		var d0 int64
		var d1 int64
		var d2 int64
		var d3 int64
		var d4 int64
		var d5 int64
		var d6 int64
		var d7 int64
		var d8 int64
		var d9 int64
		var df int64
		query = `SELECT * FROM (
				(SELECT count(*) as total from books WHERE isowned=1) AS t,
				(SELECT count(*) as d0 from books where dewey<100 and dewey >= 0 and dewey != 'fic' and IsOwned=1) as dewey0,
				(SELECT count(*) as d1 from books where dewey<200 and dewey >= 100 and dewey != 'fic' and IsOwned=1) as dewey1,
				(SELECT count(*) as d2 from books where dewey<300 and dewey >= 200 and dewey != 'fic' and IsOwned=1) as dewey2,
				(SELECT count(*) as d3 from books where dewey<400 and dewey >= 300 and dewey != 'fic' and IsOwned=1) as dewey3,
				(SELECT count(*) as d4 from books where dewey<500 and dewey >= 400 and dewey != 'fic' and IsOwned=1) as dewey4,
				(SELECT count(*) as d5 from books where dewey<600 and dewey >= 500 and dewey != 'fic' and IsOwned=1) as dewey5,
				(SELECT count(*) as d6 from books where dewey<700 and dewey >= 600 and dewey != 'fic' and IsOwned=1) as dewey6,
				(SELECT count(*) as d7 from books where dewey<800 and dewey >= 700 and dewey != 'fic' and IsOwned=1) as dewey7,
				(SELECT count(*) as d8 from books where dewey<900 and dewey >= 800 and dewey != 'fic' and IsOwned=1) as dewey8,
				(SELECT count(*) as d9 from books where dewey<1000 and dewey >= 900 and dewey != 'fic' and IsOwned=1) as dewey9,
				(SELECT count(*) as fic from books where dewey='FIC' and IsOwned=1) as deweyfic)`
		err := db.QueryRow(query).Scan(&total, &d0, &d1, &d2, &d3, &d4, &d5, &d6, &d7, &d8, &d9, &df)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Category"
		data = append(data, StatData{
			Label:    "Information Sciences",
			Value:    fmt.Sprintf("%d", d0),
			ToolText: fmt.Sprintf("%.2f%%", float64(d0)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Philosophy and Psychology",
			Value:    fmt.Sprintf("%d", d1),
			ToolText: fmt.Sprintf("%.2f%%", float64(d1)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Religion",
			Value:    fmt.Sprintf("%d", d2),
			ToolText: fmt.Sprintf("%.2f%%", float64(d2)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Social Sciences",
			Value:    fmt.Sprintf("%d", d3),
			ToolText: fmt.Sprintf("%.2f%%", float64(d3)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Language",
			Value:    fmt.Sprintf("%d", d4),
			ToolText: fmt.Sprintf("%.2f%%", float64(d4)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Mathematics and Science",
			Value:    fmt.Sprintf("%d", d5),
			ToolText: fmt.Sprintf("%.2f%%", float64(d5)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Technology",
			Value:    fmt.Sprintf("%d", d6),
			ToolText: fmt.Sprintf("%.2f%%", float64(d6)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Arts",
			Value:    fmt.Sprintf("%d", d7),
			ToolText: fmt.Sprintf("%.2f%%", float64(d7)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Literature",
			Value:    fmt.Sprintf("%d", d8),
			ToolText: fmt.Sprintf("%.2f%%", float64(d8)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Geography and History",
			Value:    fmt.Sprintf("%d", d9),
			ToolText: fmt.Sprintf("%.2f%%", float64(d9)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Fiction",
			Value:    fmt.Sprintf("%d", df),
			ToolText: fmt.Sprintf("%.2f%%", float64(df)/float64(total)*100),
		})
	case "formats":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1`
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			return chart, err
		}
		chart.Chart.Caption = "Books By Format"
		formats, err := GetFormats()
		if err != nil {
			return chart, err
		}
		for _, format := range formats {
			var count int64
			formatquery := `SELECT COUNT(*) FROM books WHERE format=? AND IsOwned=1`
			err := db.QueryRow(formatquery, format).Scan(&count)
			if err != nil {
				return chart, err
			}
			if count > 0 && format != "" {
				data = append(data, StatData{
					Label:    format,
					Value:    fmt.Sprintf("%d", count),
					ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
				})
			}
		}
	case "contributorstop":
		return chart, nil
	case "contributorsperrole":
		return chart, nil
	case "datesoriginal":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1`
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			return chart, err
		}
		query = `Select OriginallyPublished from books where OriginallyPublished != '0000-00-00' AND isowned=1`
		rows, err := db.Query(query)
		if err != nil {
			return chart, err
		}
		var dates []int
		for rows.Next() {
			var date time.Time
			err := rows.Scan(&date)
			if err != nil {
				return chart, err
			}
			dates = append(dates, date.Year())
		}
		sort.Ints(dates)
		decadeCounts := make(map[string]int)
		for _, date := range dates {
			decade := math.Floor(float64(date)/10) * 10
			if decade > 1000 {
				key := fmt.Sprintf("%.0f", decade) + "-" + fmt.Sprintf("%.0f", decade+10)
				if _, ok := decadeCounts[key]; ok {
					decadeCounts[key]++
				} else {
					decadeCounts[key] = 1
				}
			}
		}
		decades := []string{}
		for d := range decadeCounts {
			decades = append(decades, d)
		}
		sort.Slice(decades, func(i, j int) bool {
			d1, err := strconv.ParseInt(decades[i][0:4], 10, 64)
			if err != nil {
				return false
			}
			d2, err := strconv.ParseInt(decades[j][0:4], 10, 64)
			if err != nil {
				return false
			}
			return d1 < d2
		})
		for _, decade := range decades {
			data = append(data, StatData{
				Label:    decade,
				Value:    fmt.Sprintf("%d", decadeCounts[decade]),
				ToolText: fmt.Sprintf("%.2f%%", float64(decadeCounts[decade])/float64(total)*100),
			})
		}
		chart.Chart.Caption = "Books By Original Publication Date"
	case "datespublication":
		chart.Chart.FormatNumberScale = "0"
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1`
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			return chart, err
		}
		query = `Select EditionPublished from books where EditionPublished != '0000-00-00' AND isowned=1`
		rows, err := db.Query(query)
		if err != nil {
			return chart, err
		}
		var dates []int
		for rows.Next() {
			var date time.Time
			err := rows.Scan(&date)
			if err != nil {
				return chart, err
			}
			dates = append(dates, date.Year())
		}
		sort.Ints(dates)
		decadeCounts := make(map[string]int)
		for _, date := range dates {
			decade := math.Floor(float64(date)/10) * 10
			if decade > 1000 {
				key := fmt.Sprintf("%.0f", decade) + "-" + fmt.Sprintf("%.0f", decade+10)
				if _, ok := decadeCounts[key]; ok {
					decadeCounts[key]++
				} else {
					decadeCounts[key] = 1
				}
			}
		}
		decades := []string{}
		for d := range decadeCounts {
			decades = append(decades, d)
		}
		sort.Slice(decades, func(i, j int) bool {
			d1, err := strconv.ParseInt(decades[i][0:4], 10, 64)
			if err != nil {
				return false
			}
			d2, err := strconv.ParseInt(decades[j][0:4], 10, 64)
			if err != nil {
				return false
			}
			return d1 < d2
		})
		for _, decade := range decades {
			data = append(data, StatData{
				Label:    decade,
				Value:    fmt.Sprintf("%d", decadeCounts[decade]),
				ToolText: fmt.Sprintf("%.2f%%", float64(decadeCounts[decade])/float64(total)*100),
			})
		}
		chart.Chart.Caption = "Books By Edition Publication Date"
	}
	chart.Data = data
	return chart, nil
}

//GetDimensions gets dimensions
func GetDimensions() (map[string]float64, error) {
	dimensions := make(map[string]float64)
	var totalwidth float64
	var averagewidth float64
	var minimumwidth float64
	var maximumwidth float64
	var totalheight float64
	var averageheight float64
	var minimumheight float64
	var maximumheight float64
	var totaldepth float64
	var averagedepth float64
	var minimumdepth float64
	var maximumdepth float64
	var totalweight float64
	var averageweight float64
	var minimumweight float64
	var maximumweight float64
	var totalpages float64
	var averagepages float64
	var minimumpages float64
	var maximumpages float64
	var volume float64
	query := `SELECT * FROM (
				(SELECT SUM(Width) As TotalWidth, AVG(Width) As AvgWidth, MIN(Width) AS MinWidth, MAX(Width) AS MaxWidth FROM books WHERE Width>0 AND IsOwned=1) AS w,
				(SELECT SUM(Height) As TotalHeight, AVG(Height) As AvgHeight, MIN(Height) AS MinHeight, MAX(Height) AS MaxHeight FROM books WHERE Height>0 AND IsOwned=1) AS h,
				(SELECT SUM(Depth) As TotalDepth, AVG(Depth) As AvgDepth, MIN(Depth) AS MinDepth, MAX(Depth) AS MaxDepth FROM books WHERE Depth>0 AND IsOwned=1) AS d,
				(SELECT SUM(Weight) As TotalWeight, AVG(Weight) As AvgWeight, MIN(Weight) AS MinWeight, MAX(Weight) AS MaxWeight FROM books WHERE Weight>0 AND IsOwned=1) AS we,
				(SELECT SUM(Pages) As TotalPages, AVG(Pages) As AvgPages, MIN(Pages) AS MinPages, MAX(Pages) AS MaxPages FROM books WHERE pages>0 AND IsOwned=1) AS p,
				(SELECT SUM(Width*Height*Depth) as Volume FROM books WHERE IsOwned=1) as v)`
	err := db.QueryRow(query).Scan(&totalwidth, &averagewidth, &minimumwidth, &maximumwidth, &totalheight, &averageheight, &minimumheight, &maximumheight, &totaldepth, &averagedepth, &minimumdepth, &maximumdepth, &totalweight, &averageweight, &minimumweight, &maximumweight, &totalpages, &averagepages, &minimumpages, &maximumpages, &volume)
	if err != nil {
		return nil, err
	}
	dimensions["totalwidth"] = totalwidth
	dimensions["averagewidth"] = averagewidth
	dimensions["minimumwidth"] = minimumwidth
	dimensions["maximumwidth"] = maximumwidth
	dimensions["totalheight"] = totalheight
	dimensions["averageheight"] = averageheight
	dimensions["minimumheight"] = minimumheight
	dimensions["maximumheight"] = maximumheight
	dimensions["totaldepth"] = totaldepth
	dimensions["averagedepth"] = averagedepth
	dimensions["minimumdepth"] = minimumdepth
	dimensions["maximumdepth"] = maximumdepth
	dimensions["totalweight"] = totalweight
	dimensions["averageweight"] = averageweight
	dimensions["minimumweight"] = minimumweight
	dimensions["maximumweight"] = maximumweight
	dimensions["totalpages"] = totalpages
	dimensions["averagepages"] = averagepages
	dimensions["minimumpages"] = minimumpages
	dimensions["maximumpages"] = maximumpages
	dimensions["volume"] = volume
	return dimensions, nil
}

//GetCases gets cases
func GetCases() ([]Bookcase, error) {
	books, _, err := GetBooks("dewey", "both", "both", "yes", "both", "both", "both", "", "1", "-1", "0", "FIC")
	query := "SELECT CaseId, Width, SpacerHeight, PaddingLeft, PaddingRight, BookMargin FROM bookcases ORDER BY CaseNumber"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	dim, err := GetDimensions()
	if err != nil {
		return nil, err
	}
	var cases []Bookcase
	for rows.Next() {
		var id, width, spacerHeight, paddingLeft, paddingRight, bookMargin int64
		err = rows.Scan(&id, &width, &spacerHeight, &paddingLeft, &paddingRight, &bookMargin)
		if err != nil {
			return nil, err
		}
		bookcase := Bookcase{
			ID:                id,
			Width:             width,
			SpacerHeight:      spacerHeight,
			PaddingLeft:       paddingLeft,
			PaddingRight:      paddingRight,
			BookMargin:        bookMargin,
			AverageBookWidth:  dim["averagewidth"],
			AverageBookHeight: dim["averageheight"],
		}
		shelfquery := "SELECT ShelfId, Height FROM shelves WHERE CaseId=? ORDER BY ShelfNumber"
		shelfrows := db.Query(shelfquery, id)
		for shelfrows.Next() {
			var shelfid, height int64
			bookcase.Shelves = append(bookcase.Shelves, Bookshelf{
				ID:     shelfid,
				Height: shelfHeight,
			})
		}
		cases = append(cases, bookcase)
	}
	index := 0
	x := 0
	for c, bookcase := range cases {
		for s := range bookcase.Shelves {
			x = int(bookcase.PaddingLeft)
			useWidth := int(dim["averagewidth"])
			if index < len(books) && books[index].Width != 0 {
				useWidth = int(books[index].Width)
			}
			for index < len(books) && useWidth+x <= int(bookcase.Width) {
				cases[c].Shelves[s].Books = append(cases[c].Shelves[s].Books, books[index])
				x += useWidth
				index++
				useWidth = int(dim["averagewidth"])
				if index < len(books) && books[index].Width != 0 {
					useWidth = int(books[index].Width)
				}
			}
		}
	}
	return cases, nil
}
