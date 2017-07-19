package main

type EditedCases struct {
	EditedCases []EditedCase `json:"editedcases"`
	ToRemoveIds []int64 `json:"toremoveids"`
	LibraryID int64 `json:"libraryid"`
}

type EditedCase struct {
	ID                int64       `json:"id"`
	SpacerHeight      int64       `json:"spacerheight"`
	PaddingLeft       int64       `json:"paddingleft"`
	PaddingRight      int64       `json:"paddingright"`
	Width             int64       `json:"width"`
	ShelfHeight int64 `json:"shelfheight"`
	NumberOfShelves int64 `json:"numberofshelves"`
	CaseNumber	int64 `json:"casenumber"`
}

//Setting is a setting
type Setting struct {
	Name string `json:"name"`
	Value string `json:"value"`
	ValueType string `json:"valuetype"`
	Group string `json:"group"`
	PossibleValues []string `json:"possiblevalues"`
}

//Library is a library
type Library struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Permissions int64 `json:"permissions"`
	Owner string `json:"owner"`
}

//Bookcase is a bookcase
type Bookcase struct {
	ID                int64       `json:"id"`
	SpacerHeight      int64       `json:"spacerheight"`
	PaddingLeft       int64       `json:"paddingleft"`
	PaddingRight      int64       `json:"paddingright"`
	BookMargin        int64       `json:"bookmargin"`
	Width             int64       `json:"width"`
	Shelves           []Bookshelf `json:"shelves"`
	AverageBookHeight float64     `json:"averagebookheight"`
	AverageBookWidth  float64     `json:"averagebookwidth"`
	Library           Library     `json:"library"`
	CaseNumber	int64 `json:"casenumber"`
}

//Bookshelf is a shelf on a bookcase
type Bookshelf struct {
	ID     int64  `json:"id"`
	Height int64  `json:"height"`
	Books  []Book `json:"books"`
}

//StatChart is data for a chart
type StatChart struct {
	Chart StatChartInfo `json:"chart"`
	Data  []StatData    `json:"data"`
}

//StatChartInfo is chart metadata
type StatChartInfo struct {
	Caption           string `json:"caption"`
	FormatNumberScale string `json:"formatNumberScale"`
	NumberSuffix      string `json:"numberSuffix"`
	Decimals          string `json:"decimals"`
}

//StatData is chart data
type StatData struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	ToolText string `json:"tooltext"`
}

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
