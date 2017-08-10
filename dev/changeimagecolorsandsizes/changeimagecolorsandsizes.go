package main

import (
	"flag"
	"fmt"
	"log"
	"database/sql"
	"os"
	"image"
	"image/jpeg"
	"image/gif"
	"image/png"

	"github.com/jakekausler/prominentcolor"
	"github.com/nfnt/resize"
	_ "github.com/go-sql-driver/mysql"
)

var (
	username = flag.String("username", "root", "Username of the local mysql database")
	password = flag.String("password", "", "Password of the local mysql database")
	database = flag.String("name", "library", "Name of the local mysql database")
)

func main() {
	flag.Parse()
	if *username == "" {
		panic("username cannot be empty!")
	}
	if *database == "" {
		panic("database cannot be empty!")
	}
	RunServer(*username, *password, *database)
}	

var (
	db *sql.DB
	logger = log.New(os.Stderr, "log: ", log.LstdFlags | log.Lshortfile)
)

//RunServer runs the library server
func RunServer(username, password, database string) {
	logger.Printf("Creating the database")
	var err error
	// Create sql.DB
	db, err = sql.Open("mysql", "root:@/library?parseTime=true")
	// db, _ = sql.Open("mysql", "root:@/library?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	logger.Printf("Pinging the database")
	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT bookid, ImageURL from books"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var id int64
		var url string
		err = rows.Scan(&id, &url)
		if err != nil {
			panic(err.Error())
		}
		// err = resizeImage("../../web/"+url)
		// if err != nil && err.Error() == "The system cannot find the file specified." {
		// 	logger.Printf("Skipping %v. File does not exist", id)
		// 	continue
		// } else if err != nil {
		// 	panic(err.Error())
		// }
		spinecolor, err := getSpineColor("../../web/"+url)
		if err != nil {
			logger.Printf("Skipping %v", id)
			continue
		}
		query = "UPDATE books SET SpineColor=? WHERE bookid=?"
		_, err = db.Exec(query, spinecolor, id)
		if err != nil {
			panic(err.Error())
		}
		logger.Printf("Finished %v", id)
	}
	logger.Printf("Closing")
}

func resizeImage(fileLocation string) error {
	img, filetype, err := loadImage(fileLocation)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	img = resize.Thumbnail(400, 400, img, resize.Lanczos3)
	file, err := os.Create(fileLocation)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return err
	}
	defer file.Close()
	switch filetype {
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
	logger.Printf(imageLocation)
	img, _, err := loadImage(imageLocation)
	if err != nil {
		logger.Printf("Error: %+v", err)
		return "", err
	}
	var spinecolor string
	cols, err := prominentcolor.Kmeans(img)
	if err != nil {
		spinecolor = "#000000"
	} else {
		col := cols[0].Color
		spinecolor = fmt.Sprintf("#%X%X%X", col.R, col.G, col.B)
	}
	return spinecolor, nil
}

func loadImage(fileInput string) (image.Image, string, error) {
	f, err := os.Open(fileInput)
	defer f.Close()
	if err != nil {
		log.Println("File not found:", fileInput)
		return nil, "", err
	}
	img, filetype, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}

	return img, filetype, nil
}