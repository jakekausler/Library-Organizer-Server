package main

import (
    "database/sql"
    "log"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "os"
    "math"
    "image/color"

    "github.com/ajstarks/svgo"
    "github.com/jakekausler/Library-Organizer-2.0/server/libraries"
)

var (
    db     *sql.DB
    logger = log.New(os.Stdout, "log: ", log.LstdFlags|log.Lshortfile)

    username = ""
    password = ""
    host = "localhost"
    mysqlport = "3306"
    database = ""

    CaseSvgPath = "../../../web/res/caseimages"
)

func main() {
    /**********************
    ** NOT ACTUALLY USED **
    **********************/
    logger.Printf("Creating the database")
    var err error
    db, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, host, mysqlport, database))
    if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	logger.Printf("Pinging the database")

	// Test the connection
	err = db.Ping()
	for err != nil {
        logger.Printf("Could not reach the database. Sleeping for 10 seconds")
        time.Sleep(10 * time.Second)
		err = db.Ping()
	}
    /**********************
    **********************/

    // Grab all the cases
    cases, err := libraries.GetCases(db, "43", "FGiUaRYRuVCTmkPIeuNU5NDtMjkQQCrtCAOnVANHIqCyhhhX4p")
    if err != nil {
        log.Fatal(err)
    }
    
    // For each case in library
    for _, c := range cases {

        // Open case file
        cf, err := os.OpenFile(fmt.Sprintf("%s/%d.svg", CaseSvgPath, c.ID), os.O_CREATE|os.O_RDWR, 0644)
        if err != nil {
            log.Fatal(err)
        }
        cf.Truncate(0)
        cf.Seek(0, 0)

        // Get max shelf width (only books and book margins)
        // Also set default book widths and heights if none
        // Also get total shelf heights (interior only)
        maxShelfWidthBooks := 0
        totalCaseHeight := 0
        for _, s := range c.Shelves {
            width := int(s.Width)
            for _, b := range s.Books {
                if b.Width <= 0 {
                    b.Width = int64(c.AverageBookWidth)
                }
                if b.Height <= 0 {
                    b.Height = int64(c.AverageBookHeight)
                }
            }
            if width > maxShelfWidthBooks {
                maxShelfWidthBooks = width
            }
            totalCaseHeight += int(s.Height)
        }

        // Get case width
        caseWidth := int(c.SpacerHeight) * 2
        caseWidth += maxShelfWidthBooks

        // Get case height
        caseHeight := int(c.SpacerHeight) * (len(c.Shelves) + 1)
        caseHeight += totalCaseHeight

        // Start case canvas
        caseCanvas := svg.New(cf)
        caseCanvas.Start(caseWidth, caseHeight)

        // Set current y
        y := 0

        // For each shelf in case
        for _, s := range c.Shelves {

            // Set current x
            x := 0
            if s.Alignment == "right" {
                x = caseWidth - (int(s.Width) + 2 * int(c.SpacerHeight))
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
                log.Fatal(err)
            }
            sf.Truncate(0)
            sf.Seek(0, 0)

            // For each book on shelf
            for _, b := range s.Books {
                
                // Start Book Link
                caseCanvas.Link(fmt.Sprintf("http://library.jakekausler.com/books/%s", b.ID), b.ID)

                // Fix Font Color if too short
                for len(b.SpineColor) < 7 {
                    b.SpineColor += "0"
                }

                // Draw Book
                caseCanvas.Rect(x, y - int(b.Height), int(b.Width), int(b.Height), fmt.Sprintf("stroke:black;fill:%s", b.SpineColor))

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

                // End Book Link
                caseCanvas.LinkEnd()

                // Move x to start of next book
                x += int(b.Width)

            }

            // Close shelf file
            if err := sf.Close(); err != nil {
                log.Fatal(err)
            }
        }

        //End case canvas
        caseCanvas.End()

        // Close case file
        if err := cf.Close(); err != nil {
            log.Fatal(err)
        }

    }

}
