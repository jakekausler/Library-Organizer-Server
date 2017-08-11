package information

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	getPublisherQuery    = "SELECT * from publishers WHERE PublisherID=?"
	getPublishersQuery   = "SELECT DISTINCT(Publisher) from publishers"
	getCitiesQuery       = "SELECT DISTINCT(City) from publishers"
	getStatesQuery       = "SELECT DISTINCT(State) from publishers"
	getCountriesQuery    = "SELECT DISTINCT(Country) from publishers"
	getSeriesQuery       = "SELECT DISTINCT(Series) from series"
	getFormatsQuery      = "SELECT DISTINCT(Format) from formats"
	getDeweysQuery       = "SELECT Number, Genre from dewey_numbers"
	getLanguagesQuery    = "SELECT DISTINCT(Langauge) from languages"
	getRolesQuery        = "SELECT DISTINCT(Role) from written_by"
	getContributorsQuery = "SELECT PersonID, Role, FirstName, MiddleNames, LastName from written_by join persons on written_by.AuthorID = persons.PersonID WHERE BookID=?"
	getTagsQuery         = "SELECT DISTINCT(Tag) from tags"
	getAwardsQuery         = "SELECT DISTINCT(Award) from awards"
	genreQuery = "SELECT genre FROM dewey_numbers WHERE number=?"

	SORTMETHOD = "Dewey:ASC--Series:ASC--Volume:ASC--Author:ASC--Title:ASC--Subtitle:ASC--Edition:ASC--Lexile:ASC--InterestLevel:ASC--AR:ASC--LearningAZ:ASC--GuidedReading:ASC--DRA:ASC--FountasPinnell:ASC--ReadingRecovery:ASC--PMReaders:ASC--Grade:ASC--Age:ASC||Dewey:ASC--Series:ASC--Volume:ASC--Author:ASC--Title:ASC--Subtitle:ASC--Edition:ASC--Lexile:ASC--InterestLevel:ASC--AR:ASC--LearningAZ:ASC--GuidedReading:ASC--DRA:ASC--FountasPinnell:ASC--ReadingRecovery:ASC--PMReaders:ASC--Grade:ASC--Age:ASC"
)

var logger = log.New(os.Stderr, "log: ", log.LstdFlags|log.Lshortfile)

type ChartInfo struct {
	Total int `json:"total"`
	Data []StatData `json:"data"`
}

//StatData is chart data
type StatData struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	ToolText string `json:"tooltext"`
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

//Publisher is a publisher
type Publisher struct {
	ID            string `json:"id"`
	Publisher     string `json:"publisher"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ParentCompany string `json:"parentcompany"`
}

//Dewey is a dewey
type Dewey struct {
	Dewey string `json:"dewey"`
	Genre string `json:"genre"`
}

//GetStats gets statistics by type
func GetStats(db *sql.DB, t, libraryids string) (ChartInfo, error) {
	var data []StatData
	totalCount := 0
	if libraryids == "" {
		return ChartInfo{}, nil
	}
	var query string
	inlibrary := "AND libraryid IN (" + libraryids + ")"
	switch t {
	case "generalbycounts":
		var total int64
		var read int64
		var reading int64
		var reference int64
		var loaned int64
		var shipping int64
		var toread int64
		query = `SELECT * FROM (
				(SELECT count(*) as total from books WHERE isowned=1 ` + inlibrary + `) AS t,
				(SELECT count(*) as rea from books WHERE isread=1 and isowned=1 ` + inlibrary + `) AS red,
				(SELECT count(*) as reading from books WHERE isreading=1 and isowned=1 ` + inlibrary + `) AS rng,
				(SELECT count(*) as toread from books WHERE isread=0 and isreference=0 and isreading=0 and isowned=1 ` + inlibrary + `) AS trd,
				(SELECT count(*) as reference from books WHERE isreference=1 and isowned=1 ` + inlibrary + `) AS ref,
				(SELECT count(*) as loaned from books WHERE loaneeid!=-1 and isowned=1 ` + inlibrary + `) AS loa,
				(SELECT count(*) as shipping from books WHERE isshipping=1 and isowned=1 ` + inlibrary + `) AS shi)`
		err := db.QueryRow(query).Scan(&total, &read, &reading, &toread, &reference, &loaned, &shipping)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		totalCount = int(total)
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
		var total sql.NullInt64
		var read sql.NullInt64
		var reading sql.NullInt64
		var reference sql.NullInt64
		var loaned sql.NullInt64
		var shipping sql.NullInt64
		var toread sql.NullInt64
		query = `SELECT * FROM (
				(SELECT SUM(width*height*depth) as total from books WHERE isowned=1 ` + inlibrary + `) AS t,
				(SELECT SUM(width*height*depth) as rea from books WHERE isread=1 and isowned=1 ` + inlibrary + `) AS red,
				(SELECT SUM(width*height*depth) as reading from books WHERE isreading=1 and isowned=1 ` + inlibrary + `) AS rng,
				(SELECT SUM(width*height*depth) as toread from books WHERE isread=0 and isreference=0 and isreading=0 and isowned=1 ` + inlibrary + `) AS trd,
				(SELECT SUM(width*height*depth) as reference from books WHERE isreference=1 and isowned=1 ` + inlibrary + `) AS ref,
				(SELECT SUM(width*height*depth) as loaned from books WHERE loaneeid!=-1 and isowned=1 ` + inlibrary + `) AS loa,
				(SELECT SUM(width*height*depth) as shipping from books WHERE isshipping=1 and isowned=1 ` + inlibrary + `) AS shi)`
		err := db.QueryRow(query).Scan(&total, &read, &reading, &toread, &reference, &loaned, &shipping)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		data = append(data, StatData{
			Label:    "Read",
			Value:    fmt.Sprintf("%.2f", float64(read.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(read.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "Reading",
			Value:    fmt.Sprintf("%.2f", float64(reading.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(reading.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "Reference",
			Value:    fmt.Sprintf("%.2f", float64(reference.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(reference.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "To Read",
			Value:    fmt.Sprintf("%.2f", float64(toread.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(toread.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "Shipping",
			Value:    fmt.Sprintf("%.2f", float64(shipping.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(shipping.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "Loaned",
			Value:    fmt.Sprintf("%.2f", float64(loaned.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(loaned.Int64)/float64(total.Int64)*100),
		})
	case "generalbypages":
		var total sql.NullInt64
		var read sql.NullInt64
		var reading sql.NullInt64
		var reference sql.NullInt64
		var loaned sql.NullInt64
		var shipping sql.NullInt64
		var toread sql.NullInt64
		query = `SELECT * FROM (
				(SELECT SUM(pages) as total from books WHERE isowned=1 ` + inlibrary + `) AS t,
				(SELECT SUM(pages) as rea from books WHERE isread=1 and isowned=1 ` + inlibrary + `) AS red,
				(SELECT SUM(pages) as reading from books WHERE isreading=1 and isowned=1 ` + inlibrary + `) AS rng,
				(SELECT SUM(pages) as toread from books WHERE isread=0 and isreference=0 and isreading=0 and isowned=1 ` + inlibrary + `) AS trd,
				(SELECT SUM(pages) as reference from books WHERE isreference=1 and isowned=1 ` + inlibrary + `) AS ref,
				(SELECT SUM(pages) as loaned from books WHERE loaneeid!=-1 and isowned=1 ` + inlibrary + `) AS loa,
				(SELECT SUM(pages) as shipping from books WHERE isshipping=1 and isowned=1 ` + inlibrary + `) AS shi)`
		err := db.QueryRow(query).Scan(&total, &read, &reading, &toread, &reference, &loaned, &shipping)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		data = append(data, StatData{
			Label:    "Read",
			Value:    fmt.Sprintf("%.2f", float64(read.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(read.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "Reading",
			Value:    fmt.Sprintf("%.2f", float64(reading.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(reading.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "Reference",
			Value:    fmt.Sprintf("%.2f", float64(reference.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(reference.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "To Read",
			Value:    fmt.Sprintf("%.2f", float64(toread.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(toread.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "Shipping",
			Value:    fmt.Sprintf("%.2f", float64(shipping.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(shipping.Int64)/float64(total.Int64)*100),
		})
		data = append(data, StatData{
			Label:    "Loaned",
			Value:    fmt.Sprintf("%.2f", float64(loaned.Int64)),
			ToolText: fmt.Sprintf("%.2f%%", float64(loaned.Int64)/float64(total.Int64)*100),
		})
	case "publishersbooksperparent":
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		query := `SELECT COUNT(parentcompany) AS number, ParentCompany FROM publishers JOIN books ON publishers.PublisherID = books.PublisherID WHERE isowned = 1 and parentcompany != "" ` + inlibrary + ` GROUP BY parentcompany ORDER BY number DESC`
		rows, err := db.Query(query)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for rows.Next() {
			var count int64
			var company string
			err = rows.Scan(&count, &company)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
			}
			data = append(data, StatData{
				Label:    company,
				Value:    fmt.Sprintf("%.2f", float64(count)),
				ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
			})
		}
	case "publisherstopchildren":
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		query := `SELECT COUNT(publisher) AS number, Publisher FROM publishers JOIN books ON publishers.PublisherID = books.PublisherID WHERE isowned = 1 and publisher != "" ` + inlibrary + ` GROUP BY publisher ORDER BY number DESC LIMIT 30`
		rows, err := db.Query(query)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for rows.Next() {
			var count int64
			var company string
			err = rows.Scan(&count, &company)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
			}
			data = append(data, StatData{
				Label:    company,
				Value:    fmt.Sprintf("%.2f", float64(count)),
				ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
			})
		}
	case "publisherstoplocations":
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		query := `SELECT COUNT(*) AS number, city, state, country FROM publishers JOIN books ON publishers.PublisherID = books.PublisherID WHERE isowned = 1 and (city != "" or state != "" or country != "") ` + inlibrary + ` GROUP BY city, state, country ORDER BY number DESC LIMIT 30`
		rows, err := db.Query(query)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for rows.Next() {
			var count int64
			var city, state, country string
			err = rows.Scan(&count, &city, &state, &country)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
			}
			label := city + ", "
			if state != "" {
				label = label + state
			} else {
				label = label + country
			}
			data = append(data, StatData{
				Label:    label,
				Value:    fmt.Sprintf("%.2f", float64(count)),
				ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
			})
		}
	case "series":
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		series, err := GetSeries(db, "")
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for _, s := range series {
			var count int64
			seriesquery := `SELECT COUNT(*) FROM books WHERE series=? AND IsOwned=1 ` + inlibrary
			err := db.QueryRow(seriesquery, s).Scan(&count)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
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
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		languages, err := GetLanguages(db, "")
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for _, language := range languages {
			var count int64
			languagequery := `SELECT COUNT(*) FROM books WHERE PrimaryLanguage=? AND IsOwned=1 ` + inlibrary
			err := db.QueryRow(languagequery, language).Scan(&count)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
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
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		languages, err := GetLanguages(db, "")
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for _, language := range languages {
			var count int64
			languagequery := `SELECT COUNT(*) FROM books WHERE SecondaryLanguage=? AND IsOwned=1 ` + inlibrary
			err := db.QueryRow(languagequery, language).Scan(&count)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
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
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		languages, err := GetLanguages(db, "")
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for _, language := range languages {
			var count int64
			languagequery := `SELECT COUNT(*) FROM books WHERE OriginalLanguage=? AND IsOwned=1 ` + inlibrary
			err := db.QueryRow(languagequery, language).Scan(&count)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
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
				(SELECT count(*) as total from books WHERE isowned=1 ` + inlibrary + `) AS t,
				(SELECT count(*) as d0 from books where dewey<100 and dewey >= 0 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey0,
				(SELECT count(*) as d1 from books where dewey<200 and dewey >= 100 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey1,
				(SELECT count(*) as d2 from books where dewey<300 and dewey >= 200 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey2,
				(SELECT count(*) as d3 from books where dewey<400 and dewey >= 300 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey3,
				(SELECT count(*) as d4 from books where dewey<500 and dewey >= 400 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey4,
				(SELECT count(*) as d5 from books where dewey<600 and dewey >= 500 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey5,
				(SELECT count(*) as d6 from books where dewey<700 and dewey >= 600 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey6,
				(SELECT count(*) as d7 from books where dewey<800 and dewey >= 700 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey7,
				(SELECT count(*) as d8 from books where dewey<900 and dewey >= 800 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey8,
				(SELECT count(*) as d9 from books where dewey<1000 and dewey >= 900 and dewey != 'fic' and IsOwned=1 ` + inlibrary + `) as dewey9,
				(SELECT count(*) as fic from books where dewey='FIC' and IsOwned=1 ` + inlibrary + `) as deweyfic)`
		err := db.QueryRow(query).Scan(&total, &d0, &d1, &d2, &d3, &d4, &d5, &d6, &d7, &d8, &d9, &df)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
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
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		formats, err := GetFormats(db, "")
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for _, format := range formats {
			var count int64
			formatquery := `SELECT COUNT(*) FROM books WHERE format=? AND IsOwned=1 ` + inlibrary
			err := db.QueryRow(formatquery, format).Scan(&count)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
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
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		query := `SELECT COUNT(authorid) AS BooksWritten, FirstName, MiddleNames, LastName, Role FROM persons JOIN written_by ON written_by.authorid = persons.personid JOIN books ON books.BookID = written_by.BookID WHERE isowned = 1 and lastname != "" ` + inlibrary + ` GROUP BY AUTHORID ORDER BY BooksWritten DESC LIMIT 30`
		rows, err := db.Query(query)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for rows.Next() {
			var count int64
			var fn, mn, ln, role string
			err = rows.Scan(&count, &fn, &mn, &ln, &role)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
			}
			name := ln + " (" + role + ")"
			if mn != "" {
				name = strings.Join(strings.Split(mn, ";"), " ") + " " + name
			}
			if fn != "" {
				name = fn + " " + name
			}
			data = append(data, StatData{
				Label:    name,
				Value:    fmt.Sprintf("%.2f", float64(count)),
				ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
			})
		}
	case "contributorsperrole":
		query := `SELECT COUNT(role) AS InRole, Role FROM persons JOIN written_by ON written_by.authorid = persons.personid JOIN books ON books.BookID = written_by.BookID WHERE isowned = 1 and role != "" ` + inlibrary + ` GROUP BY Role ORDER BY InRole DESC`
		rows, err := db.Query(query)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for rows.Next() {
			var count int64
			var role string
			err = rows.Scan(&count, &role)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
			}
			data = append(data, StatData{
				Label: role,
				Value: fmt.Sprintf("%.2f", float64(count)),
			})
		}
	case "datesoriginal":
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		query = `Select OriginallyPublished from books where OriginallyPublished != '0000-00-00' AND isowned=1 ` + inlibrary
		rows, err := db.Query(query)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		var dates []int
		for rows.Next() {
			var date time.Time
			err := rows.Scan(&date)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
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
				logger.Printf("Error: %+v", err)
				return false
			}
			d2, err := strconv.ParseInt(decades[j][0:4], 10, 64)
			if err != nil {
				logger.Printf("Error: %+v", err)
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
	case "datespublication":
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		query = `Select EditionPublished from books where EditionPublished != '0000-00-00' AND isowned=1 ` + inlibrary
		rows, err := db.Query(query)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		var dates []int
		for rows.Next() {
			var date time.Time
			err := rows.Scan(&date)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
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
				logger.Printf("Error: %+v", err)
				return false
			}
			d2, err := strconv.ParseInt(decades[j][0:4], 10, 64)
			if err != nil {
				logger.Printf("Error: %+v", err)
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
	case "lexile":
		var total int64
		var l0 int64
		var l1 int64
		var l2 int64
		var l3 int64
		var l4 int64
		var l5 int64
		var l6 int64
		var l7 int64
		var l8 int64
		var l9 int64
		var l10 int64
		var l11 int64
		var l12 int64
		query = `SELECT * FROM (
				(SELECT count(*) as total from books WHERE isowned=1 ` + inlibrary + `) AS t,
				(SELECT count(*) as l0 from books where lexile<=189 and IsOwned=1 ` + inlibrary + `) as l0,
				(SELECT count(*) as l1 from books where lexile<=530 and lexile >= 190 and IsOwned=1 ` + inlibrary + `) as l1,
				(SELECT count(*) as l2 from books where lexile<=650 and lexile >= 420 and IsOwned=1 ` + inlibrary + `) as l2,
				(SELECT count(*) as l3 from books where lexile<=820 and lexile >= 520 and IsOwned=1 ` + inlibrary + `) as l3,
				(SELECT count(*) as l4 from books where lexile<=940 and lexile >= 740 and IsOwned=1 ` + inlibrary + `) as l4,
				(SELECT count(*) as l5 from books where lexile<=1010 and lexile >= 830 and IsOwned=1 ` + inlibrary + `) as l5,
				(SELECT count(*) as l6 from books where lexile<=1070 and lexile >= 925 and IsOwned=1 ` + inlibrary + `) as l6,
				(SELECT count(*) as l7 from books where lexile<=1120 and lexile >= 970 and IsOwned=1 ` + inlibrary + `) as l7,
				(SELECT count(*) as l8 from books where lexile<=1185 and lexile >= 800 and IsOwned=1 ` + inlibrary + `) as l8,
				(SELECT count(*) as l9 from books where lexile<=1260 and lexile >= 1010 and IsOwned=1 ` + inlibrary + `) as l9,
				(SELECT count(*) as l10 from books where lexile<=1335 and lexile >= 1050 and IsOwned=1 ` + inlibrary + `) as l10,
				(SELECT count(*) as l11 from books where lexile<=1385 and lexile >= 1080 and IsOwned=1 ` + inlibrary + `) as l11,
				(SELECT count(*) as l12 from books where lexile >= 1386 and IsOwned=1 ` + inlibrary + `) as l12)`
		err := db.QueryRow(query).Scan(&total, &l0, &l1, &l2, &l3, &l4, &l5, &l6, &l7, &l8, &l9, &l10, &l11, &l12)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		data = append(data, StatData{
			Label:    "Pre Grade 1\n(Less than 190L)",
			Value:    fmt.Sprintf("%d", l0),
			ToolText: fmt.Sprintf("%.2f%%", float64(l0)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 1\n(190L-530L)",
			Value:    fmt.Sprintf("%d", l1),
			ToolText: fmt.Sprintf("%.2f%%", float64(l1)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 2\n(420L-650L)",
			Value:    fmt.Sprintf("%d", l2),
			ToolText: fmt.Sprintf("%.2f%%", float64(l2)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 3\n(520L-820L)",
			Value:    fmt.Sprintf("%d", l3),
			ToolText: fmt.Sprintf("%.2f%%", float64(l3)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 4\n(740L-940L)",
			Value:    fmt.Sprintf("%d", l4),
			ToolText: fmt.Sprintf("%.2f%%", float64(l4)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 5\n(830L-1010L)",
			Value:    fmt.Sprintf("%d", l5),
			ToolText: fmt.Sprintf("%.2f%%", float64(l5)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 6\n(925L-1070L)",
			Value:    fmt.Sprintf("%d", l6),
			ToolText: fmt.Sprintf("%.2f%%", float64(l6)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 7\n(970L-1120L)",
			Value:    fmt.Sprintf("%d", l7),
			ToolText: fmt.Sprintf("%.2f%%", float64(l7)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 8\n(1010L-1185L)",
			Value:    fmt.Sprintf("%d", l8),
			ToolText: fmt.Sprintf("%.2f%%", float64(l8)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 9\n(1050L-1260L)",
			Value:    fmt.Sprintf("%d", l9),
			ToolText: fmt.Sprintf("%.2f%%", float64(l9)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 10\n(1080L-1335L)",
			Value:    fmt.Sprintf("%d", l10),
			ToolText: fmt.Sprintf("%.2f%%", float64(l10)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Grade 11-12\n(1185L-1385L)",
			Value:    fmt.Sprintf("%d", l11),
			ToolText: fmt.Sprintf("%.2f%%", float64(l11)/float64(total)*100),
		})
		data = append(data, StatData{
			Label:    "Post Grade 12\n(Greater than 1385L)",
			Value:    fmt.Sprintf("%d", l12),
			ToolText: fmt.Sprintf("%.2f%%", float64(l12)/float64(total)*100),
		})
	case "tag":
		var total int64
		totalquery := `SELECT count(*) FROM books WHERE isowned=1 ` + inlibrary
		err := db.QueryRow(totalquery).Scan(&total)
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		tags, err := GetTags(db, "")
		if err != nil {
			logger.Printf("Error: %+v", err)
			return ChartInfo{}, err
		}
		for _, tag := range tags {
			var count int64
			tagsQuery := `SELECT COUNT(*) FROM tags JOIN books ON tags.BookID=books.BookID WHERE tag=? AND IsOwned=1 ` + inlibrary
			err := db.QueryRow(tagsQuery, tag).Scan(&count)
			if err != nil {
				logger.Printf("Error: %+v", err)
				return ChartInfo{}, err
			}
			if count > 0 && tag != "" {
				data = append(data, StatData{
					Label:    tag,
					Value:    fmt.Sprintf("%d", count),
					ToolText: fmt.Sprintf("%.2f%%", float64(count)/float64(total)*100),
				})
			}
		}
	}
	return ChartInfo{
		Total: totalCount,
		Data: data,
	}, nil
}

//GetDimensions gets dimensions
func GetDimensions(db *sql.DB, libraryids string) (map[string]float64, error) {
	dimensions := make(map[string]float64)
	if libraryids == "" {
		return dimensions, nil
	}
	var totalwidth sql.NullFloat64
	var averagewidth sql.NullFloat64
	var minimumwidth sql.NullFloat64
	var maximumwidth sql.NullFloat64
	var totalheight sql.NullFloat64
	var averageheight sql.NullFloat64
	var minimumheight sql.NullFloat64
	var maximumheight sql.NullFloat64
	var totaldepth sql.NullFloat64
	var averagedepth sql.NullFloat64
	var minimumdepth sql.NullFloat64
	var maximumdepth sql.NullFloat64
	var totalweight sql.NullFloat64
	var averageweight sql.NullFloat64
	var minimumweight sql.NullFloat64
	var maximumweight sql.NullFloat64
	var totalpages sql.NullFloat64
	var averagepages sql.NullFloat64
	var minimumpages sql.NullFloat64
	var maximumpages sql.NullFloat64
	var volume sql.NullFloat64
	inlibrary := "AND libraryid IN (" + libraryids + ")"
	query := `SELECT * FROM (
				(SELECT SUM(Width) As TotalWidth, AVG(Width) As AvgWidth, MIN(Width) AS MinWidth, MAX(Width) AS MaxWidth FROM books WHERE Width>0 AND IsOwned=1 ` + inlibrary + `) AS w,
				(SELECT SUM(Height) As TotalHeight, AVG(Height) As AvgHeight, MIN(Height) AS MinHeight, MAX(Height) AS MaxHeight FROM books WHERE Height>0 AND IsOwned=1 ` + inlibrary + `) AS h,
				(SELECT SUM(Depth) As TotalDepth, AVG(Depth) As AvgDepth, MIN(Depth) AS MinDepth, MAX(Depth) AS MaxDepth FROM books WHERE Depth>0 AND IsOwned=1 ` + inlibrary + `) AS d,
				(SELECT SUM(Weight) As TotalWeight, AVG(Weight) As AvgWeight, MIN(Weight) AS MinWeight, MAX(Weight) AS MaxWeight FROM books WHERE Weight>0 AND IsOwned=1 ` + inlibrary + `) AS we,
				(SELECT SUM(Pages) As TotalPages, AVG(Pages) As AvgPages, MIN(Pages) AS MinPages, MAX(Pages) AS MaxPages FROM books WHERE pages>0 AND IsOwned=1 ` + inlibrary + `) AS p,
				(SELECT SUM(Width*Height*Depth) as Volume FROM books WHERE IsOwned=1 ` + inlibrary + `) as v)`
	err := db.QueryRow(query).Scan(&totalwidth, &averagewidth, &minimumwidth, &maximumwidth, &totalheight, &averageheight, &minimumheight, &maximumheight, &totaldepth, &averagedepth, &minimumdepth, &maximumdepth, &totalweight, &averageweight, &minimumweight, &maximumweight, &totalpages, &averagepages, &minimumpages, &maximumpages, &volume)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return nil, err
	}
	dimensions["totalwidth"] = totalwidth.Float64
	dimensions["averagewidth"] = averagewidth.Float64
	dimensions["minimumwidth"] = minimumwidth.Float64
	dimensions["maximumwidth"] = maximumwidth.Float64
	dimensions["totalheight"] = totalheight.Float64
	dimensions["averageheight"] = averageheight.Float64
	dimensions["minimumheight"] = minimumheight.Float64
	dimensions["maximumheight"] = maximumheight.Float64
	dimensions["totaldepth"] = totaldepth.Float64
	dimensions["averagedepth"] = averagedepth.Float64
	dimensions["minimumdepth"] = minimumdepth.Float64
	dimensions["maximumdepth"] = maximumdepth.Float64
	dimensions["totalweight"] = totalweight.Float64
	dimensions["averageweight"] = averageweight.Float64
	dimensions["minimumweight"] = minimumweight.Float64
	dimensions["maximumweight"] = maximumweight.Float64
	dimensions["totalpages"] = totalpages.Float64
	dimensions["averagepages"] = averagepages.Float64
	dimensions["minimumpages"] = minimumpages.Float64
	dimensions["maximumpages"] = maximumpages.Float64
	dimensions["volume"] = volume.Float64
	return dimensions, nil
}

//GetContributors gets all contributors for a book id
func GetContributors(db *sql.DB, id string) ([]Contributor, error) {
	c := Contributor{}
	var contributors = make([]Contributor, 0)

	var Role sql.NullString
	var First sql.NullString
	var Middles sql.NullString
	var Last sql.NullString

	rows, err := db.Query(getContributorsQuery, id)
	if err != nil {
		logger.Printf("Error querying contributors: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&c.ID, &Role, &First, &Middles, &Last); err != nil {
			logger.Printf("Error scanning contributors: %v", err)
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
func GetPublisher(db *sql.DB, id string) (Publisher, error) {
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
		logger.Printf("Error scanning publisher for id %v: %v", id, err)
		return p, err
	}

	return p, nil
}

//GetPublishers gets all publishers
func GetPublishers(db *sql.DB, queryString string) ([]string, error) {
	var s string
	var r = make([]string, 0)
	query := getPublishersQuery
	if queryString != "" {
		query += " WHERE Publisher LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying publishers: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("Error scanning publishers: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetCities gets all cities
func GetCities(db *sql.DB, queryString string) ([]string, error) {
	var s string
	var r = make([]string, 0)
	query := getCitiesQuery
	if queryString != "" {
		query += " WHERE City LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying cities: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("Error scanning cities: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetStates gets all states
func GetStates(db *sql.DB, queryString string) ([]string, error) {
	var s string
	var r = make([]string, 0)
	query := getStatesQuery
	if queryString != "" {
		query += " WHERE State LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying states: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("Error scanning states: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetCountries gets all countries
func GetCountries(db *sql.DB, queryString string) ([]string, error) {
	var s string
	var r = make([]string, 0)
	query := getCountriesQuery
	if queryString != "" {
		query += " WHERE Country LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying countries: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("Error scanning countries: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetSeries gets all series
func GetSeries(db *sql.DB, queryString string) ([]string, error) {
	var s string
	var r = make([]string, 0)
	query := getSeriesQuery
	if queryString != "" {
		query += " WHERE Series LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying series: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("Error scanning series: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetFormats gets all formats
func GetFormats(db *sql.DB, queryString string) ([]string, error) {
	var s string
	var r = make([]string, 0)
	query := getFormatsQuery
	if queryString != "" {
		query += " WHERE Format LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying formats: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("Error scanning formats: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetLanguages gets all languages
func GetLanguages(db *sql.DB, queryString string) ([]string, error) {
	var s string
	var r = make([]string, 0)
	query := getLanguagesQuery
	if queryString != "" {
		query += " WHERE Language LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying languages: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("Error scanning languages: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetRoles gets all roles
func GetRoles(db *sql.DB, queryString string) ([]string, error) {
	var s string
	var r = make([]string, 0)
	query := getRolesQuery
	if queryString != "" {
		query += " WHERE Role LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying roles: %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Printf("Error scanning roles: %v", err)
			return nil, err
		}
		r = append(r, s)
	}
	return r, nil
}

//GetDeweys gets all deweys
func GetDeweys(db *sql.DB, queryString string) ([]Dewey, error) {
	var r = make([]Dewey, 0)
	query := getDeweysQuery
	if queryString != "" {
		query += " WHERE Number LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying deweys: %v", err)
		return nil, err
	}
	for rows.Next() {
		var d Dewey
		var s sql.NullString
		if err := rows.Scan(&d.Dewey, &s); err != nil {
			logger.Printf("Error scanning deweys: %v", err)
			return nil, err
		}
		d.Genre = s.String
		r = append(r, d)
	}
	return r, nil
}

//GetGenre gets the genre of a dewey
func GetGenre(db *sql.DB, dewey string) (string, error) {
	var genre sql.NullString
	err := db.QueryRow(genreQuery, dewey).Scan(&genre)
	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		logger.Printf("Error scanning genre: %v", err)
		return "", err
	}
	return genre.String, nil
}

//GetTags gets all tags
func GetTags(db *sql.DB, queryString string) ([]string, error) {
	var r []string
	query := getTagsQuery
	if queryString != "" {
		query += " WHERE Tag LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying tags: %v", err)
		return nil, err
	}
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			logger.Printf("Error scanning tags: %v", err)
			return nil, err
		}
		r = append(r, tag)
	}
	return r, nil
}

//GetAwards gets all tags
func GetAwards(db *sql.DB, queryString string) ([]string, error) {
	var r []string
	query := getAwardsQuery
	if queryString != "" {
		query += " WHERE Award LIKE '%%" + queryString + "%%'"
	}
	rows, err := db.Query(query)
	if err != nil {
		logger.Printf("Error querying awards: %v", err)
		return nil, err
	}
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			logger.Printf("Error scanning awards: %v", err)
			return nil, err
		}
		r = append(r, tag)
	}
	return r, nil
}