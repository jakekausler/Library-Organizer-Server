package books

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	saveBookQuery = "UPDATE books SET Title=?, Subtitle=?, OriginallyPublished=?, EditionPublished=?, PublisherID=?, IsRead=?, IsReference=?, IsOwned=?, IsShipping=?, IsReading=?, isbn=?, LoaneeFirst=?, LoaneeLast=?, Dewey=?, Pages=?, Width=?, Height=?, Depth=?, Weight=?, PrimaryLanguage=?, SecondaryLanguage=?, OriginalLanguage=?, Series=?, Volume=?, Format=?, Edition=?, ImageURL=?, LibraryId=? WHERE BookId=?"
	addBookQuery  = "INSERT INTO books (Title, Subtitle, OriginallyPublished, PublisherID, IsRead, IsReference, IsOwned, IsShipping, IsReading, isbn, LoaneeFirst, LoaneeLast, Dewey, Pages, Width, Height, Depth, Weight, PrimaryLanguage, SecondaryLanguage, OriginalLanguage, Series, Volume, Format, Edition, EditionPublished, LibraryId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
)

var logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)

//BookIds is a list of book ids
type BookIds struct {
	BookIds []string `json:"bookids"`
}

//BookSet is a collection of books, along with the number of pages of data there are without a limit imposed
type BookSet struct {
	Books         []Book `json:"books"`
	NumberOfBooks int64  `json:"numbooks"`
}

//Book is a book
type Book struct {
	ID                  string        `json:"bookid"`
	Title               string        `json:"title"`
	Subtitle            string        `json:"subtitle"`
	OriginallyPublished string        `json:"originallypublished"`
	Publisher           Publisher     `json:"publisher"`
	IsRead              bool          `json:"isread"`
	IsReference         bool          `json:"isreference"`
	IsOwned             bool          `json:"isowned"`
	ISBN                string        `json:"isbn"`
	Loanee              Name          `json:"loanee"`
	Dewey               string        `json:"dewey"`
	Pages               int64         `json:"pages"`
	Width               int64         `json:"width"`
	Height              int64         `json:"height"`
	Depth               int64         `json:"depth"`
	Weight              float64       `json:"weight"`
	PrimaryLanguage     string        `json:"primarylanguage"`
	SecondaryLanguage   string        `json:"secondarylanguage"`
	OriginalLanguage    string        `json:"originallanguage"`
	Series              string        `json:"series"`
	Volume              float64       `json:"volume"`
	Format              string        `json:"format"`
	Edition             int64         `json:"edition"`
	IsReading           bool          `json:"isreading"`
	IsShipping          bool          `json:"isshipping"`
	ImageURL            string        `json:"imageurl"`
	SpineColor          string        `json:"spinecolor"`
	CheapestNew         float64       `json:"cheapestnew"`
	CheapestUsed        float64       `json:"cheapestused"`
	EditionPublished    string        `json:"editionpublished"`
	Contributors        []Contributor `json:"contributors"`
	Library             Library       `json:"library"`
}

//Publisher is a publisher
type Publisher struct {
	ID            string `json:"id"`
	Publisher     string `json:"publisher"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ParentCompany string `json:"parentcompany"`
}

//Name is a name
type Name struct {
	First   string `json:"first"`
	Middles string `json:"middles"`
	Last    string `json:"last"`
}

//Contributor is a contributor
type Contributor struct {
	ID   string `json:"id"`
	Name Name   `json:"name"`
	Role string `json:"role"`
}

//Library is a library
type Library struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Permissions int64  `json:"permissions"`
	Owner       string `json:"owner"`
}

//SaveBook saves a book
func SaveBook(db *sql.DB, book Book) error {
	if book.ID != "" {
		imageType := filepath.Ext(book.ImageURL)
		if book.ImageURL != "" && !strings.HasPrefix(book.ImageURL, "res/bookimages/") {
			err := downloadImage(book.ImageURL, "../web/res/bookimages/"+book.ID+imageType)
			if err != nil {
				logger.Printf("Error while saving image: %v", err)
				return err
			}
		}
		publisherID, err := addOrGetPublisher(db, book.Publisher.Publisher, book.Publisher.City, book.Publisher.State, book.Publisher.Country)
		if err != nil {
			logger.Printf("Error when saving publisher: %v", err)
			return err
		}
		err = addDewey(db, book.Dewey)
		if err != nil {
			logger.Printf("Error when saving dewey: %v", err)
			return err
		}
		err = addSeries(db, book.Series)
		if err != nil {
			logger.Printf("Error when saving Series: %v", err)
			return err
		}
		err = addFormat(db, book.Format)
		if err != nil {
			logger.Printf("Error when saving Format: %v", err)
			return err
		}
		err = addLanguage(db, book.PrimaryLanguage)
		if err != nil {
			logger.Printf("Error when saving PrimaryLanguage: %v", err)
			return err
		}
		err = addLanguage(db, book.SecondaryLanguage)
		if err != nil {
			logger.Printf("Error when saving SecondaryLanguage: %v", err)
			return err
		}
		err = addLanguage(db, book.OriginalLanguage)
		if err != nil {
			logger.Printf("Error when saving OriginalLanguage: %v", err)
			return err
		}
		_, err = db.Exec(saveBookQuery, book.Title, book.Subtitle, book.OriginallyPublished, book.EditionPublished, publisherID, book.IsRead, book.IsReference, book.IsOwned, book.IsShipping, book.IsReading, book.ISBN, book.Loanee.First, book.Loanee.Last, book.Dewey, book.Pages, book.Width, book.Height, book.Depth, book.Weight, book.PrimaryLanguage, book.SecondaryLanguage, book.OriginalLanguage, book.Series, book.Volume, book.Format, book.Edition, "res/bookimages/"+book.ID+imageType, book.Library.ID, book.ID)
		if err != nil {
			logger.Printf("Error when saving book: %v", err)
			return err
		}
		err = removeAllWrittenBy(db, book.ID)
		if err != nil {
			logger.Printf("Error when saving authors: %v", err)
			return err
		}
		for _, contributor := range book.Contributors {
			err = addContributor(db, book.ID, contributor)
			if err != nil {
				logger.Printf("Error when saving authors: %v", err)
				return err
			}
		}
	} else {
		publisherID, err := addOrGetPublisher(db, book.Publisher.Publisher, book.Publisher.City, book.Publisher.State, book.Publisher.Country)
		if err != nil {
			logger.Printf("Error when saving publisher: %v", err)
			return err
		}
		err = addDewey(db, book.Dewey)
		if err != nil {
			logger.Printf("Error when saving dewey: %v", err)
			return err
		}
		err = addSeries(db, book.Series)
		if err != nil {
			logger.Printf("Error when saving Series: %v", err)
			return err
		}
		err = addFormat(db, book.Format)
		if err != nil {
			logger.Printf("Error when saving Format: %v", err)
			return err
		}
		err = addLanguage(db, book.PrimaryLanguage)
		if err != nil {
			logger.Printf("Error when saving PrimaryLanguage: %v", err)
			return err
		}
		err = addLanguage(db, book.SecondaryLanguage)
		if err != nil {
			logger.Printf("Error when saving SecondaryLanguage: %v", err)
			return err
		}
		err = addLanguage(db, book.OriginalLanguage)
		if err != nil {
			logger.Printf("Error when saving OriginalLanguage: %v", err)
			return err
		}
		res, err := db.Exec(addBookQuery, book.Title, book.Subtitle, book.OriginallyPublished, publisherID, book.IsRead, book.IsReference, book.IsOwned, book.IsShipping, book.IsReading, book.ISBN, book.Loanee.First, book.Loanee.Last, book.Dewey, book.Pages, book.Width, book.Height, book.Depth, book.Weight, book.PrimaryLanguage, book.SecondaryLanguage, book.OriginalLanguage, book.Series, book.Volume, book.Format, book.Edition, book.EditionPublished, book.Library.ID)
		if err != nil {
			logger.Printf("Error when saving book: %v", err)
			return err
		}
		id, err := res.LastInsertId()
		bookid := strconv.FormatInt(id, 10)
		imageType := filepath.Ext(book.ImageURL)
		if book.ImageURL != "" {
			err = downloadImage(book.ImageURL, "../web/res/bookimages/"+bookid+imageType)
		}
		if err != nil {
			logger.Printf("Error while saving image: %v", err)
			return err
		}
		imageQuery := "UPDATE books SET ImageURL='res/bookimages/" + bookid + imageType + "' WHERE bookid=?"
		_, err = db.Exec(imageQuery, bookid)
		if err != nil {
			logger.Printf("Error when saving image: %v", err)
			return err
		}
		for _, contributor := range book.Contributors {
			err = addContributor(bookid, contributor)
			if err != nil {
				logger.Printf("Error when saving authors: %v", err)
				return err
			}
		}
	}
	return nil
}

func removeAllWrittenBy(db *sql.DB, bookid string) error {
	query := "DELETE FROM written_by WHERE BookId=?"
	_, err := db.Exec(query, bookid)
	if err != nil {
		logger.Printf("Error when deleting written_by: %v", err)
		return err
	}
	return nil
}

//DeleteBook deletes a book
func DeleteBook(db *sql.DB, bookid string) error {
	removeAllWrittenBy(bookid)
	query := "DELETE FROM books WHERE BookId=?"
	_, err := db.Exec(query, bookid)
	if err != nil {
		logger.Printf("Error when deleting book: %v", err)
		return err
	}
	return nil
}

func addContributor(db *sql.DB, bookid string, contributor Contributor) error {
	personID, err := addOrGetPerson(db, contributor.Name.First, contributor.Name.Middles, contributor.Name.Last)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	query := "REPLACE INTO written_by (BookID, AuthorID, Role) VALUES (?,?,?)"
	_, err = db.Exec(query, bookid, personID, contributor.Role)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addOrGetPerson(db *sql.DB, first, middles, last string) (string, error) {
	var id int64
	id = -1
	query := "SELECT PersonID FROM persons WHERE FirstName=? AND MiddleNames=? AND LastName=?"
	rows, err := db.Query(query, first, middles, last)
	if err != nil {
		logger.Printf("Error: %+v", err)
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
			logger.Printf("Error: %+v", err)
			return "", err
		}
		id, err = res.LastInsertId()
		if err != nil {
			logger.Printf("Error: %+v", err)
			return "", err
		}
	}
	return strconv.FormatInt(id, 10), nil
}

func addOrGetPublisher(db *sql.DB, publisher, city, state, country string) (string, error) {
	var id int64
	id = -1
	query := "SELECT PublisherID FROM publishers WHERE Publisher=? AND City=? AND State=? AND Country=?"
	rows, err := db.Query(query, publisher, city, state, country)
	if err != nil {
		logger.Printf("Error: %+v", err)
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
			logger.Printf("Error: %+v", err)
			return "", err
		}
		id, err = res.LastInsertId()
		if err != nil {
			logger.Printf("Error: %+v", err)
			return "", err
		}
	}
	return strconv.FormatInt(id, 10), nil
}

func addDewey(db *sql.DB, v string) error {
	query := "REPLACE INTO dewey_numbers (Number) VALUES (?)"
	_, err := db.Exec(query, v)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addFormat(db *sql.DB, v string) error {
	query := "REPLACE INTO formats (Format) VALUES (?)"
	_, err := db.Exec(query, v)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addSeries(db *sql.DB, v string) error {
	query := "REPLACE INTO series (Series) VALUES (?)"
	_, err := db.Exec(query, v)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addLanguage(db *sql.DB, v string) error {
	query := "REPLACE INTO languages (Langauge) VALUES (?)"
	_, err := db.Exec(query, v)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func downloadImage(url, fileLocation string) error {
	response, err := http.Get(url)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	defer response.Body.Close()
	file, err := os.Create(fileLocation)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	file.Close()
	return nil
}

//GetBooks gets all books
//todo include authors in filter
func GetBooks(db *sql.DB, sortMethod, isread, isreference, isowned, isloaned, isreading, isshipping, text, page, numberToGet, fromDewey, toDewey, libraryids, session string) ([]Book, int64, error) {
	if libraryids == "" {
		return nil, 0, nil
	}
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
	if filter != "WHERE " {
		filter = filter + " AND "
	}
	filter = filter + "libraryid IN (" + libraryids + ")"
	if filter == "WHERE " || filter == "WHERE" {
		filter = ""
	}
	pag, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, 0, err
	}
	ntg, err := strconv.ParseInt(numberToGet, 10, 64)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, 0, err
	}
	query := "SELECT bookid, title, subtitle, OriginallyPublished, PublisherID, isread, isreference, IsOwned, ISBN, LoaneeFirst, LoaneeLast, dewey, pages, width, height, depth, weight, PrimaryLanguage, SecondaryLanguage, OriginalLanguage, series, volume, format, Edition, ImageURL, IsReading, isshipping, SpineColor, CheapestNew, CheapestUsed, EditionPublished, blid, libraries.Name, library_members.usr, permissions.Permission from (select books.*, books.libraryid as blid, " + titlechange + ", " + serieschange + ", min(name) as minname FROM books LEFT JOIN " + authors + " ON books.BookID = Authors.BookID " + filter + " GROUP BY books.BookID) i LEFT JOIN libraries ON blid=libraries.id JOIN library_members on libraries.ownerid=library_members.id JOIN permissions on permissions.userid=? and libraries.id=permissions.libraryid ORDER BY " + order
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

	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error getting username: %v", err)
		return nil, 0, err
	}
	rows, err := db.Query(query, userid)
	if err != nil {
		logger.Printf("Error querying books: %v", err)
		return nil, 0, err
	}
	for rows.Next() {
		if err := rows.Scan(&b.ID, &Title, &Subtitle, &OriginallyPublished, &PublisherID, &IsRead, &IsReference, &IsOwned, &ISBN, &LoaneeFirst, &LoaneeLast, &Dewey, &b.Pages, &b.Width, &b.Height, &b.Depth, &b.Weight, &PrimaryLanguage, &SecondaryLanguage, &OriginalLanguage, &Series, &b.Volume, &Format, &b.Edition, &ImageURL, &IsReading, &IsShipping, &SpineColor, &b.CheapestNew, &b.CheapestUsed, &EditionPublished, &b.Library.ID, &b.Library.Name, &b.Library.Owner, &b.Library.Permissions); err != nil {
			logger.Printf("Error scanning books: %v", err)
			return nil, 0, err
		}
		if PublisherID.Valid {
			p, err = information.GetPublisher(PublisherID.String)
			if err != nil {
				logger.Printf("Error getting publisher: %v", err)
				return nil, 0, err
			}
		}
		b.Publisher = p
		c, err = users.GetContributors(b.ID)
		if err != nil {
			logger.Printf("Error getting contributors: %v", err)
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
		logger.Printf("Error in rows: %v", err)
		return nil, 0, err
	}

	var numberOfBooks int64
	err = db.QueryRow(pageQuery).Scan(&numberOfBooks)
	if err != nil {
		logger.Printf("Error: %+v", err)
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

//GetBooksForExport selects all books as strings
func GetBooksForExport(db *sql.DB) ([][]string, error) {
	query := "select * from books join publishers on books.PublisherID=publishers.PublisherID"
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error exporting books: %v", err)
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		logger.Printf("Error exporting books: %v", err)
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
			logger.Printf("Error exporting books: %v", err)
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
func GetAuthorsForExport(db *sql.DB) ([][]string, error) {
	query := "SELECT BookID, FirstName, MiddleNames, LastName, Role from written_by JOIN persons on written_by.AuthorID=persons.PersonID"
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error exporting books: %v", err)
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		logger.Printf("Error exporting books: %v", err)
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
			logger.Printf("Error exporting books: %v", err)
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

//ImportBooks imports records from a csv file
//todo finish function
func ImportBooks(db *sql.DB, records [][]string) error {
	logger.Printf("Importing...")
	return nil
}
