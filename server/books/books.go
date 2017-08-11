package books

import (
	"database/sql"
	"fmt"
	"image"
	"image/jpeg"
	"image/gif"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"../users"
	"../information"
	"github.com/jakekausler/prominentcolor"
	"github.com/go-sql-driver/mysql"
)

const (
	saveBookQuery = "UPDATE books SET Title=?, Subtitle=?, OriginallyPublished=?, EditionPublished=?, PublisherID=?, IsRead=?, IsReference=?, IsOwned=?, IsShipping=?, IsReading=?, isbn=?, Dewey=?, Pages=?, Width=?, Height=?, Depth=?, Weight=?, PrimaryLanguage=?, SecondaryLanguage=?, OriginalLanguage=?, Series=?, Volume=?, Format=?, Edition=?, ImageURL=?, LibraryId=?, Lexile=?, LexileCode=?, InterestLevel=?, AR=?, LearningAZ=?, GuidedReading=?, DRA=?, Grade=?, FountasPinnell=?, Age=?, ReadingRecovery=?, PMReaders=?, SpineColor=?, Notes=? WHERE BookId=?"
	addBookQuery  = "INSERT INTO books (Title, Subtitle, OriginallyPublished, PublisherID, IsRead, IsReference, IsOwned, IsShipping, IsReading, isbn, Dewey, Pages, Width, Height, Depth, Weight, PrimaryLanguage, SecondaryLanguage, OriginalLanguage, Series, Volume, Format, Edition, EditionPublished, LibraryId, Lexile, LexileCode, InterestLevel, AR, LearningAZ, GuidedReading, DRA, Grade, FountasPinnell, Age, ReadingRecovery, PMReaders, Notes) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	getBookQuery  = "SELECT bookid, title, subtitle, OriginallyPublished, PublisherID, isread, isreference, IsOwned, ISBN, LoaneeId, dewey, pages, width, height, depth, weight, PrimaryLanguage, SecondaryLanguage, OriginalLanguage, series, volume, format, Edition, ImageURL, IsReading, isshipping, SpineColor, CheapestNew, CheapestUsed, EditionPublished, SpineColorOverridden, Lexile,  LexileCode, InterestLevel, AR, LearningAZ, GuidedReading, DRA, Grade, FountasPinnell, Age, ReadingRecovery, PMReaders, Notes, blid, libraries.Name, library_members.usr, permissions.Permission from (select books.*, books.libraryid as blid FROM books LEFT JOIN (SELECT  PersonID, AuthorRoles.BookID, concat(COALESCE(lastname,''),COALESCE(firstname,''),COALESCE(middlenames,'')) as name FROM persons JOIN (SELECT written_by.BookID, AuthorID FROM written_by WHERE Role='Author') AS AuthorRoles ON AuthorRoles.AuthorID = persons.PersonID ORDER BY name ) AS Authors ON books.BookID=Authors.BookID) i LEFT JOIN libraries on blid=libraries.id JOIN library_members on libraries.ownerid=library_members.id JOIN permissions on permissions.userid=? AND libraries.id=permissions.libraryid WHERE BookId=?"
	deleteWrittenByQuery = "DELETE FROM written_by WHERE BookId=?"
	deleteAwardsQuery = "DELETE FROM awards WHERE BookId=?"
	deleteBookQuery = "DELETE FROM books WHERE BookId=?"
	checkoutBookQuery = "UPDATE books SET loaneeid=? WHERE bookid=?"
	checkinBookQuery = "UPDATE books SET loaneeid=-1 WHERE bookid=?"
	addContributorQuery = "REPLACE INTO written_by (BookID, AuthorID, Role) VALUES (?,?,?)"
	addAwardQuery = "REPLACE INTO awards (BookID, Award) VALUES (?,?)"
	getPersonQuery = "SELECT PersonID FROM persons WHERE FirstName=? AND MiddleNames=? AND LastName=?"
	addPersonQuery = "INSERT INTO persons (FirstName, MiddleNames, LastName) VALUES (?,?,?)"
	getPublisherQuery = "SELECT PublisherID FROM publishers WHERE Publisher=? AND City=? AND State=? AND Country=?"
	addPublisherQuery = "INSERT INTO publishers (Publisher, City, State, Country) VALUES (?,?,?,?)"
	addDeweyQuery = "REPLACE INTO dewey_numbers (Number) VALUES (?)"
	addFormatQuery = "REPLACE INTO formats (Format) VALUES (?)"
	addSeriesQuery = "REPLACE INTO series (Series) VALUES (?)"
	addLanguageQuery = "REPLACE INTO languages (Langauge) VALUES (?)"
	getLibraryMemberQuery = "SELECT firstname, lastname, usr, email from library_members WHERE id=?"
	addTagQuery = "INSERT INTO tags (bookid, tag) VALUES (?,?)"
	getTagsQuery = "SELECT tag FROM tags WHERE bookid=?"
	deleteTagsQuery = "DELETE FROM tags WHERE bookid=?"
	getExportBooksQuery = "SELECT * FROM books JOIN publishers ON books.PublisherID=publishers.PublisherID"
	getExportAuthorsQuery = "SELECT BookID, FirstName, MiddleNames, LastName, Role FROM written_by JOIN persons ON written_by.AuthorID=persons.PersonID"
	getBookForMatchQuery = "SELECT ISBN, Title, Subtitle, Series, Volume FROM books WHERE bookid=?"
	getGuessedMatchesQuery = "SELECT BookId FROM books WHERE isbn=? || (title=? && subtitle=? && series=? && volume=?)"
	addReviewQuery = "REPLACE INTO reviews (userid, bookid, review) VALUES (?,?,?)"
	addRatingQuery = "REPLACE INTO ratings (userid, bookid, rating) VALUES (?,?,?)"
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
	Publisher           information.Publisher     `json:"publisher"`
	IsRead              bool          `json:"isread"`
	IsReference         bool          `json:"isreference"`
	IsOwned             bool          `json:"isowned"`
	ISBN                string        `json:"isbn"`
	Loanee              users.User    `json:"loanee"`
	Dewey               sql.NullString	`json:"dewey"`
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
	SpineColorOverridden bool 		  `json:"spinecoloroverridden"`
	CheapestNew         float64       `json:"cheapestnew"`
	CheapestUsed        float64       `json:"cheapestused"`
	EditionPublished    string        `json:"editionpublished"`
	Contributors        []information.Contributor `json:"contributors"`
	Library             Library       		`json:"library"`
	Lexile				sql.NullInt64		`json:"lexile"`
	LexileCode			string		  		`json:"lexilecode"`
	InterestLevel		sql.NullInt64		`json:"interestlevel"`
	AR					sql.NullFloat64		`json:"ar"`
	LearningAZ			sql.NullInt64		`json:"learningaz"`
	GuidedReading		sql.NullInt64		`json:"guidedreading"`
	DRA					sql.NullInt64		`json:"dra"`
	Grade				sql.NullInt64		`json:"grade"`
	FountasPinnell		sql.NullInt64		`json:"fountaspinnell"`
	Age					sql.NullInt64		`json:"age"`
	ReadingRecovery		sql.NullInt64		`json:"readingrecovery"`
	PMReaders			sql.NullInt64		`json:"pmreaders"`
	Awards				[]string	  `json:"awards"`
	Notes				string		  `json:"notes"`
	Tags				[]string	  `json:"tags"`
}

//Library is a library
type Library struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Permissions int64  `json:"permissions"`
	Owner       string `json:"owner"`
}

//Rating is a book rating
type Rating struct {
	BookID int64 `json:"bookid"`
	UserID int64 `json:"userid"`
	Rating int64 `json:"rating"`
	Username string `json:"username"`
}

//Review is a book rating
type Review struct {
	BookID int64 `json:"bookid"`
	UserID int64 `json:"userid"`
	Review string `json:"review"`
	Username string `json:"username"`
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
		} else if book.ImageURL == "" {
			err := removeImage(book.ID)
			if err != nil {
				logger.Printf("Error while removing image: %v", err)
				return err
			}
		}

		spinecolor := book.SpineColor
		if !book.SpineColorOverridden {
			var err error
			spinecolor, err = getSpineColor("../web/res/bookimages/"+book.ID+imageType)
			if err != nil {
				spinecolor = "#000000"
			}
		}

		publisherID, err := addOrGetPublisher(db, book.Publisher.Publisher, book.Publisher.City, book.Publisher.State, book.Publisher.Country)
		if err != nil {
			logger.Printf("Error when saving publisher: %v", err)
			return err
		}
		Dewey, err := book.Dewey.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		err = addDewey(db, book.Dewey.String)
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
		Lexile, err := book.Lexile.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		LexileCode := book.LexileCode
		AR, err := book.AR.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		InterestLevel, err := book.InterestLevel.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		LearningAZ, err := book.LearningAZ.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		GuidedReading, err := book.GuidedReading.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		DRA, err := book.DRA.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		Grade, err := book.Grade.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		FountasPinnell, err := book.FountasPinnell.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		Age, err := book.Age.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		ReadingRecovery, err := book.ReadingRecovery.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		PMReaders, err := book.PMReaders.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		_, err = db.Exec(saveBookQuery, book.Title, book.Subtitle, book.OriginallyPublished, book.EditionPublished, publisherID, book.IsRead, book.IsReference, book.IsOwned, book.IsShipping, book.IsReading, book.ISBN, Dewey, book.Pages, book.Width, book.Height, book.Depth, book.Weight, book.PrimaryLanguage, book.SecondaryLanguage, book.OriginalLanguage, book.Series, book.Volume, book.Format, book.Edition, "res/bookimages/"+book.ID+imageType, book.Library.ID, Lexile, LexileCode, InterestLevel, AR, LearningAZ, GuidedReading, DRA, Grade, FountasPinnell, Age, ReadingRecovery, PMReaders, spinecolor, book.Notes, book.ID)
		if err != nil {
			logger.Printf("Error when saving book: %v", err)
			return err
		}
		err = removeAllAwards(db, book.ID)
		if err != nil {
			logger.Printf("Error when saving awards: %v", err)
			return err
		}
		for _, award := range book.Awards {
			err = addAward(db, book.ID, award)
			if err != nil {
				logger.Printf("Error when saving awards: %v", err)
				return err
			}
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
		_, err = db.Exec(deleteTagsQuery, book.ID)
		if err != nil {
			logger.Printf("Error when saving authors: %v", err)
			return err
		}
		for _, tag := range book.Tags {
			_, err = db.Exec(addTagQuery, book.ID, tag)
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
		Dewey, err := book.Dewey.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		err = addDewey(db, book.Dewey.String)
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
		Lexile, err := book.Lexile.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		LexileCode := book.LexileCode
		AR, err := book.AR.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		InterestLevel, err := book.InterestLevel.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		LearningAZ, err := book.LearningAZ.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		GuidedReading, err := book.GuidedReading.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		DRA, err := book.DRA.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		Grade, err := book.Grade.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		FountasPinnell, err := book.FountasPinnell.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		Age, err := book.Age.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		ReadingRecovery, err := book.ReadingRecovery.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		PMReaders, err := book.PMReaders.Value()
		if err != nil {
			logger.Printf("Error: %v")
			return err
		}
		res, err := db.Exec(addBookQuery, book.Title, book.Subtitle, book.OriginallyPublished, publisherID, book.IsRead, book.IsReference, book.IsOwned, book.IsShipping, book.IsReading, book.ISBN, Dewey, book.Pages, book.Width, book.Height, book.Depth, book.Weight, book.PrimaryLanguage, book.SecondaryLanguage, book.OriginalLanguage, book.Series, book.Volume, book.Format, book.Edition, book.EditionPublished, book.Library.ID, Lexile, LexileCode, InterestLevel, AR, LearningAZ, GuidedReading, DRA, Grade, FountasPinnell, Age, ReadingRecovery, PMReaders, book.Notes)
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
		spinecolor := book.SpineColor
		if !book.SpineColorOverridden {
			var err error
			spinecolor, err = getSpineColor("../web/res/bookimages/"+bookid+imageType)
			if err != nil {
				spinecolor = "#000000"
			}
		}
		imageQuery := "UPDATE books SET ImageURL='res/bookimages/" + bookid + imageType + "', SpineColor=? WHERE bookid=?"
		_, err = db.Exec(imageQuery, spinecolor, bookid)
		if err != nil {
			logger.Printf("Error when saving image: %v", err)
			return err
		}
		for _, award := range book.Awards {
			err = addAward(db, bookid, award)
			if err != nil {
				logger.Printf("Error when saving awards: %v", err)
				return err
			}
		}
		for _, contributor := range book.Contributors {
			err = addContributor(db, bookid, contributor)
			if err != nil {
				logger.Printf("Error when saving authors: %v", err)
				return err
			}
		}
		_, err = db.Exec(deleteTagsQuery, id)
		if err != nil {
			logger.Printf("Error when saving authors: %v", err)
			return err
		}
		for _, tag := range book.Tags {
			_, err = db.Exec(addTagQuery, id, tag)
			if err != nil {
				logger.Printf("Error when saving authors: %v", err)
				return err
			}
		}
	}
	return nil
}

func removeAllWrittenBy(db *sql.DB, bookid string) error {
	_, err := db.Exec(deleteWrittenByQuery, bookid)
	if err != nil {
		logger.Printf("Error when deleting written_by: %v", err)
		return err
	}
	return nil
}

func removeAllAwards(db *sql.DB, bookid string) error {
	_, err := db.Exec(deleteAwardsQuery, bookid)
	if err != nil {
		logger.Printf("Error when deleting award: %v", err)
		return err
	}
	return nil
}

//DeleteBook deletes a book
func DeleteBook(db *sql.DB, bookid string) error {
	removeAllWrittenBy(db, bookid)
	removeAllAwards(db, bookid)
	_, err := db.Exec(deleteBookQuery, bookid)
	if err != nil {
		logger.Printf("Error when deleting book: %v", err)
		return err
	}
	return nil
}

//CheckoutBook checks out a book
func CheckoutBook(db *sql.DB, session string, bookid int) error {
	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	_, err = db.Exec(checkoutBookQuery, userid, bookid)
	return err
}

//CheckinBook checks in a book
func CheckinBook(db *sql.DB, bookid int) error {
	_, err := db.Exec(checkinBookQuery, bookid)
	return err
}

func addContributor(db *sql.DB, bookid string, contributor information.Contributor) error {
	personID, err := addOrGetPerson(db, contributor.Name.First, contributor.Name.Middles, contributor.Name.Last)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	_, err = db.Exec(addContributorQuery, bookid, personID, contributor.Role)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addAward(db *sql.DB, bookid, award string) error {
	_, err := db.Exec(addAwardQuery, bookid, award)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addOrGetPerson(db *sql.DB, first, middles, last string) (string, error) {
	var id int64
	id = -1
	rows, err := db.Query(getPersonQuery, first, middles, last)
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
		res, err := db.Exec(addPersonQuery, first, middles, last)
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
	rows, err := db.Query(getPublisherQuery, publisher, city, state, country)
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
		res, err := db.Exec(addPublisherQuery, publisher, city, state, country)
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
	_, err := db.Exec(addDeweyQuery, v)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addFormat(db *sql.DB, v string) error {
	_, err := db.Exec(addFormatQuery, v)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addSeries(db *sql.DB, v string) error {
	_, err := db.Exec(addSeriesQuery, v)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	return nil
}

func addLanguage(db *sql.DB, v string) error {
	_, err := db.Exec(addLanguageQuery, v)
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
	img, extension, err := image.Decode(response.Body)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	file, err := os.Create(fileLocation)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	defer file.Close()
	switch extension {
	case "jpg", "jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{})
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
	case "gif":
		err = gif.Encode(file, img, &gif.Options{})
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
	case "png":
		err = png.Encode(file, img)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return err
		}
	}
	return nil
}

func getSpineColor(imageLocation string) (string, error) {
	img, err := loadImage(imageLocation)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	var spinecolor string
	cols, err := prominentcolor.Kmeans(img)
	col := cols[0].Color
	if err != nil {
		spinecolor = "#000000"
	} else {
		spinecolor = fmt.Sprintf("#%X%X%X", col.R, col.G, col.B)
	}
	return spinecolor, nil
}

func loadImage(fileInput string) (image.Image, error) {
	f, err := os.Open(fileInput)
	defer f.Close()
	if err != nil {
		log.Println("File not found:", fileInput)
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func removeImage(bookid string) error {
	fs, err := filepath.Glob("../web/res/bookimages/"+bookid+"*")
	if err != nil {
		logger.Printf("Error removing image: %v", err)
		return err
	}
	for _, f := range fs {
		err = os.Remove(f)
		if err != nil {
			logger.Printf("Error removing image: %v", err)
			return err
		}
	}
	return nil
}

//GetBooks gets all books
//todo include authors in filter
func GetBooks(db *sql.DB, sortMethod, isread, isreference, isowned, isloaned, isreading, isshipping, text, page, numberToGet, fromDewey, toDewey, fromLexile, toLexile, fromInterestLevel, toInterestLevel, fromAR, toAR, fromLearningAZ, toLearningAZ, fromGuidedReading, toGuidedReading, fromDRA, toDRA, fromGrade, toGrade, fromFountasPinnell, toFountasPinnell, fromAge, toAge, fromReadingRecovery, toReadingRecovery, fromPMReaders, toPMReaders, libraryids, isbn, session string, authorseries []string, caseQuery bool) ([]Book, int64, error) {
	text = strings.Replace(text, "'", "\\'", -1)
	if libraryids == "" {
		return nil, 0, nil
	}
	var order string
	if strings.Contains(sortMethod, "||") {
		order = ""
		sortMethod = strings.ToLower(sortMethod)
		sortMethod = strings.Replace(sortMethod, "author", "minname", -1)
		sortMethod = strings.Replace(sortMethod, "title", "title2", -1)
		sortMethod = strings.Replace(sortMethod, "series", "series2", -1)
		sortMethod = strings.Replace(sortMethod, "volume", "LPAD(Volume, 8, '0')", -1)
		orders := strings.Split(sortMethod, "||")
		normalOrders := strings.Split(orders[0], "--")
		specialOrders := strings.Split(orders[1], "--")
		if authorseries != nil {
			var normal [][]string
			var special [][]string
			for _, o := range normalOrders {
				normal = append(normal, strings.Split(o, ":"))
			}
			for _, o := range specialOrders {
				special = append(special, strings.Split(o, ":"))
			}
			caseWhenThen := "(CASE WHEN Series IN('"+strings.Join(authorseries, "','")+"') THEN "
			var splitOrders []string
			for i := range normalOrders {
				splitOrders = append(splitOrders, caseWhenThen + special[i][0] + " ELSE " + normal[i][0] + " END) ")
			}
			order = strings.Join(splitOrders, ",")
		} else {
			order = strings.Replace(strings.Join(normalOrders, ","), ";", " ", -1)
		}
	} else {
		if sortMethod == "title" {
			order = "Title2, minname"
		} else if sortMethod == "series" {
			order = "if(Series2='' or Series2 is null,1,0), Series2, Volume, minname, Title2"
		} else {
			if authorseries != nil {
				order = "Dewey, CASE WHEN Series IN('"+strings.Join(authorseries, "','")+"') THEN Series2 ELSE minname END, CASE WHEN Series IN('"+strings.Join(authorseries, "','")+"') THEN LPAD(Volume, 8, '0') ELSE series2 END, CASE WHEN Series IN('"+strings.Join(authorseries, "','")+"') THEN minname ELSE LPAD(Volume, 8, '0') END, Title2, Subtitle2, Edition"
			} else {
				order = "Dewey, minname, Series2, Volume, title2, Subtitle2, edition"
			}
		}
	}
	titlechange := "CASE WHEN Title LIKE 'The %%' THEN TRIM(SUBSTR(Title from 4)) ELSE CASE WHEN Title LIKE 'An %%' THEN TRIM(SUBSTR(Title from 3)) ELSE CASE WHEN Title LIKE 'A %%' THEN TRIM(SUBSTR(Title from 2)) ELSE Title END END END AS Title2"
	subtitlechange := "CASE WHEN Subtitle LIKE 'The %%' THEN TRIM(SUBSTR(Subtitle from 4)) ELSE CASE WHEN Title LIKE 'An %%' THEN TRIM(SUBSTR(Subtitle from 3)) ELSE CASE WHEN Title LIKE 'A %%' THEN TRIM(SUBSTR(Subtitle from 2)) ELSE Subtitle END END END AS Subtitle2"
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
		loaned = "LoaneeId!=-1"
	} else if isloaned == "no" {
		loaned = "LoaneeId=-1"
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
	startDewey := ""
	if fromDewey != "" {
		startDewey = "Dewey >= '" + formatDewey(fromDewey) + "'"
	}
	endDewey := ""
	if toDewey != "" {
		endDewey = "Dewey <= '" + formatDewey(toDewey) + "'"
	}
	startLexile := ""
	if fromLexile != "" {
		startLexile = "Lexile >= '" + fromLexile + "'"
	}
	endLexile := ""
	if toLexile != "" {
		endLexile = "Lexile <= '" + toLexile + "'"
	}
	startInterestLevel := ""
	if fromInterestLevel != "" {
		startInterestLevel = "InterestLevel >= '" + fromInterestLevel + "'"
	}
	endInterestLevel := ""
	if toInterestLevel != "" {
		endInterestLevel = "InterestLevel <= '" + toInterestLevel + "'"
	}
	startAR := ""
	if fromAR != "" {
		startAR = "AR >= '" + fromAR + "'"
	}
	endAR := ""
	if toAR != "" {
		endAR = "AR <= '" + toAR + "'"
	}
	startLearningAZ := ""
	if fromLearningAZ != "" {
		startLearningAZ = "LearningAZ >= '" + fromLearningAZ + "'"
	}
	endLearningAZ := ""
	if toLearningAZ != "" {
		endLearningAZ = "LearningAZ <= '" + toLearningAZ + "'"
	}
	startGuidedReading := ""
	if fromGuidedReading != "" {
		startGuidedReading = "GuidedReading >= '" + fromGuidedReading + "'"
	}
	endGuidedReading := ""
	if toGuidedReading != "" {
		endGuidedReading = "GuidedReading <= '" + toGuidedReading + "'"
	}
	startDRA := ""
	if fromDRA != "" {
		startDRA = "DRA >= '" + fromDRA + "'"
	}
	endDRA := ""
	if toDRA != "" {
		endDRA = "DRA <= '" + toDRA + "'"
	}
	startGrade := ""
	if fromGrade != "" {
		startGrade = "Grade >= '" + fromGrade + "'"
	}
	endGrade := ""
	if toGrade != "" {
		endGrade = "Grade <= '" + toGrade + "'"
	}
	startFountasPinnell := ""
	if fromFountasPinnell != "" {
		startFountasPinnell = "FountasPinnell >= '" + fromFountasPinnell + "'"
	}
	endFountasPinnell := ""
	if toFountasPinnell != "" {
		endFountasPinnell = "FountasPinnell <= '" + toFountasPinnell + "'"
	}
	startAge := ""
	if fromAge != "" {
		startAge = "Age >= '" + fromAge + "'"
	}
	endAge := ""
	if toAge != "" {
		endAge = "Age <= '" + toAge + "'"
	}
	startReadingRecovery := ""
	if fromReadingRecovery != "" {
		startReadingRecovery = "ReadingRecovery >= '" + fromReadingRecovery + "'"
	}
	endReadingRecovery := ""
	if toReadingRecovery != "" {
		endReadingRecovery = "ReadingRecovery <= '" + toReadingRecovery + "'"
	}
	startPMReaders := ""
	if fromPMReaders != "" {
		startPMReaders = "Lexile >= '" + fromPMReaders + "'"
	}
	endPMReaders := ""
	if toPMReaders != "" {
		endPMReaders = "PMReaders <= '" + toPMReaders + "'"
	}
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

	if startLexile != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startLexile
	}
	if endLexile != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endLexile
	}

	if startInterestLevel != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startInterestLevel
	}
	if endInterestLevel != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endInterestLevel
	}

	if startAR != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startAR
	}
	if endAR != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endAR
	}

	if startLearningAZ != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startLearningAZ
	}
	if endLearningAZ != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endLearningAZ
	}

	if startGuidedReading != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startGuidedReading
	}
	if endGuidedReading != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endGuidedReading
	}

	if startDRA != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startDRA
	}
	if endDRA != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endDRA
	}

	if startGrade != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startGrade
	}
	if endGrade != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endGrade
	}

	if startFountasPinnell != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startFountasPinnell
	}
	if endFountasPinnell != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endFountasPinnell
	}

	if startAge != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startAge
	}
	if endAge != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endAge
	}

	if startReadingRecovery != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startReadingRecovery
	}
	if endReadingRecovery != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endReadingRecovery
	}

	if startPMReaders != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + startPMReaders
	}
	if endPMReaders != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + endPMReaders
	}

	if isbn != "" {
		if filter != "WHERE " {
			filter = filter + " AND "
		}
		filter = filter + "ISBN=" + isbn
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
	fields := "bookid, title, subtitle, ImageURL, blid, libraries.Name, library_members.usr, permissions.Permission"
	if caseQuery {
		fields += ", SpineColor, width, height, dewey, series"
	}
	query := "SELECT " + fields + " FROM (select books.*, books.libraryid as blid, " + titlechange + ", " + subtitlechange + ", " + serieschange + ", min(name) as minname FROM books LEFT JOIN " + authors + " ON books.BookID = Authors.BookID " + filter + " GROUP BY books.BookID) i LEFT JOIN libraries ON blid=libraries.id JOIN library_members on libraries.ownerid=library_members.id JOIN permissions on permissions.userid=? and libraries.id=permissions.libraryid ORDER BY " + order
	if numberToGet != "-1" {
		query += " LIMIT " + numberToGet + " OFFSET " + strconv.FormatInt(((pag-1)*ntg), 10)
	}
	pageQuery := "SELECT count(bookid) FROM (select books.bookid, " + titlechange + ", " + subtitlechange + ", " + serieschange + ", min(name) as minname FROM books LEFT JOIN " + authors + " ON books.BookID = Authors.BookID " + filter + " GROUP BY books.BookID) i"

	b := Book{}
	var books = make([]Book, 0)

	var Title sql.NullString
	var Subtitle sql.NullString
	var ImageURL sql.NullString
	var SpineColor sql.NullString
	var Dewey sql.NullString
	var Series sql.NullString

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
		if caseQuery {
			if err := rows.Scan(&b.ID, &Title, &Subtitle, &ImageURL, &b.Library.ID, &b.Library.Name, &b.Library.Owner, &b.Library.Permissions, &SpineColor, &b.Width, &b.Height, &Dewey, &Series); err != nil {
				logger.Printf("Error scanning books: %v", err)
				return nil, 0, err
			}
		} else {
			if err := rows.Scan(&b.ID, &Title, &Subtitle, &ImageURL, &b.Library.ID, &b.Library.Name, &b.Library.Owner, &b.Library.Permissions); err != nil {
				logger.Printf("Error scanning books: %v", err)
				return nil, 0, err
			}
		}
		b.Title = ""
		if Title.Valid {
			b.Title = Title.String
		}
		b.Subtitle = ""
		if Subtitle.Valid {
			b.Subtitle = Subtitle.String
		}
		b.ImageURL = ""
		if ImageURL.Valid {
			b.ImageURL = ImageURL.String
		}
		if caseQuery {
			b.SpineColor = ""
			if SpineColor.Valid {
				b.SpineColor = SpineColor.String
			}
			b.Dewey = Dewey
			b.Series = ""
			if Series.Valid {
				b.Series = Series.String
			}
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
			s = s + "(Title LIKE '%%" + filter + "%%' OR Subtitle LIKE '%%" + filter + "%%' OR Series LIKE '%%" + filter + "%%') AND "
		}
	}
	if strings.HasSuffix(s, " AND ") {
		s = s[:len(s)-5]
	}
	return s
}

//GetBooksForExport selects all books as strings
func GetBooksForExport(db *sql.DB) ([][]string, error) {
	rows, err := db.Query(getExportBooksQuery)
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
	rows, err := db.Query(getExportAuthorsQuery)
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

//GetBookForMatch gets a book by its id
func GetBookForMatch(db *sql.DB, id string) (Book, error) {
	var book Book
	err := db.QueryRow(getBookForMatchQuery, id).Scan(&book.ISBN, &book.Title, &book.Subtitle, &book.Series, &book.Volume)
	return book, err
}

//GetGuessMatchedBooks gets ids of books that are similar enough to a book that they could be the same
func GetGuessMatchedBooks(db *sql.DB, id string) ([]int64, error) {
	book, err := GetBookForMatch(db, id)
	if err != nil {
		logger.Printf("Error: %v", err)
		return nil, err
	}
	rows, err := db.Query(getGuessedMatchesQuery, book.ISBN, book.Title, book.Subtitle, book.Series, book.Volume)
	if err != nil {
		logger.Printf("Error: %v", err)
		return nil, err
	}
	var ids []int64
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			logger.Printf("Error: %v", err)
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

//GetBookReviews gets book reviews for a book and its guessed matches
func GetBookReviews(db *sql.DB, id string) ([]Review, error) {
	matches, err := GetGuessMatchedBooks(db, id)
	if err != nil {
		logger.Printf("Error: %v", err)
		return nil, err
	}
	var reviews []Review
	query := "SELECT bookid, userid, review, usr FROM reviews JOIN library_members ON userid=id WHERE bookid IN ('"+intArrToStr(matches, "','")+"')"
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error: %v", err)
		return nil, err
	}
	for rows.Next() {
		var r Review
		// var rawReview []uint8
		err = rows.Scan(&r.BookID, &r.UserID, &r.Review, &r.Username)
		if err != nil {
			logger.Printf("Error: %v", err)
			return nil, err
		}
		// r.Review = uint8ToString(rawReview)
		reviews = append(reviews, r)
	}
	return reviews, nil
}

func uint8ToString(t []uint8) string {
	b := make([]byte, len(t))
	for i, v := range t {
		b[i] = byte(v)
	}
	return string(b)
}

//AddBookReview adds or replaces a review for a book
func AddBookReview(db *sql.DB, id, session string, review string) error {
	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	_, err = db.Exec(addReviewQuery, userid, id, review)
	return err
}

//GetBookRatings gets book ratings for a book and its guessed matches
func GetBookRatings(db *sql.DB, id string) ([]Rating, error) {
	matches, err := GetGuessMatchedBooks(db, id)
	if err != nil {
		logger.Printf("Error: %v", err)
		return nil, err
	}
	var ratings []Rating
	query := "SELECT bookid, userid, rating, usr FROM ratings JOIN library_members ON userid=id WHERE bookid IN ('"+intArrToStr(matches, "','")+"')"
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error: %v", err)
		return nil, err
	}
	for rows.Next() {
		var r Rating
		err = rows.Scan(&r.BookID, &r.UserID, &r.Rating, &r.Username)
		if err != nil {
			logger.Printf("Error: %v", err)
			return nil, err
		}
		ratings = append(ratings, r)
	}
	return ratings, nil
}

//AddBookRating adds or replaces a rating for a book
func AddBookRating(db *sql.DB, id, session string, rating int) error {
	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	_, err = db.Exec(addRatingQuery, userid, id, rating)
	return err
}

func intArrToStr(arr []int64, del string) string {
	var retval []string
	for _, i := range arr {
		retval = append(retval, strconv.Itoa(int(i)))
	}
	return strings.Join(retval, del)
}

//GetBook gets a book by its id
func GetBook(db *sql.DB, session, bookid string) (Book, error) {
	b := Book{}

	var PublisherID sql.NullString
	var IsRead int64
	var IsReference int64
	var IsOwned int64
	var IsShipping int64
	var IsReading int64
	var LoaneeId int64
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
	var spinecoloroverridden int64

	var p information.Publisher
	var c []information.Contributor

	userid, err := users.GetUserID(db, session)
	if err != nil {
		logger.Printf("Error getting username: %v", err)
		return b, err
	}
	err = db.QueryRow(getBookQuery, userid, bookid).Scan(&b.ID, &Title, &Subtitle, &OriginallyPublished, &PublisherID, &IsRead, &IsReference, &IsOwned, &ISBN, &LoaneeId, &Dewey, &b.Pages, &b.Width, &b.Height, &b.Depth, &b.Weight, &PrimaryLanguage, &SecondaryLanguage, &OriginalLanguage, &Series, &b.Volume, &Format, &b.Edition, &ImageURL, &IsReading, &IsShipping, &SpineColor, &b.CheapestNew, &b.CheapestUsed, &EditionPublished, &spinecoloroverridden, &b.Lexile, &b.LexileCode, &b.InterestLevel, &b.AR, &b.LearningAZ, &b.GuidedReading, &b.DRA, &b.Grade, &b.FountasPinnell, &b.Age, &b.ReadingRecovery, &b.PMReaders, &b.Notes, &b.Library.ID, &b.Library.Name, &b.Library.Owner, &b.Library.Permissions)
	if err != nil {
		logger.Printf("Error getting username: %v", err)
		return b, err
	}
	if PublisherID.Valid {
		p, err = information.GetPublisher(db, PublisherID.String)
		if err != nil {
			logger.Printf("Error getting publisher: %v", err)
			return b, err
		}
	}
	b.Publisher = p
	c, err = information.GetContributors(db, b.ID)
	if err != nil {
		logger.Printf("Error getting contributors: %v", err)
		return b, err
	}
	b.Contributors = c
	b.Loanee = users.User{
		ID: LoaneeId,
	}
	if b.Loanee.ID != -1 {
		err = db.QueryRow(getLibraryMemberQuery, b.Loanee.ID).Scan(&b.Loanee.FirstName, &b.Loanee.LastName, &b.Loanee.Username, &b.Loanee.Email)
		if err != nil {
			logger.Printf("Error getting loanee: %v", err)
			return b, err
		}
	}
	b.IsOwned = IsOwned == 1
	b.IsReference = IsReference == 1
	b.IsReading = IsReading == 1
	b.IsRead = IsRead == 1
	b.IsShipping = IsShipping == 1
	b.SpineColorOverridden = spinecoloroverridden == 1
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
	b.Dewey = Dewey
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
	innerRows, err := db.Query(getTagsQuery, b.ID)
	if err != nil {
		logger.Printf("Error: %v", err)
		return b, err
	}
	for innerRows.Next() {
		var tag string
		err := innerRows.Scan(&tag)
		if err != nil {
			logger.Printf("Error: %v", err)
			return b, err
		}
		b.Tags = append(b.Tags, tag)
	}
	return b, err
}