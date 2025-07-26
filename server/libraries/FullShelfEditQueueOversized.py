
import requests
from pprint import pprint
import json
import svgwrite

RELOAD = True

SORT_METHOD = "Dewey:ASC--Author:ASC--SeriesHeight:ASC--Series:ASC--Volume:ASC--Title:ASC--Subtitle:ASC--Edition:ASC--Lexile:ASC--InterestLevel:ASC--AR:ASC--LearningAZ:ASC--GuidedReading:ASC--DRA:ASC--FountasPinnell:ASC--ReadingRecovery:ASC--PMReaders:ASC--Grade:ASC--Age:ASC--Publisher:ASC--LibraryOfCongress:ASC--FictionGenre:ASC||Dewey:ASC--SeriesHeight:ASC--Series:ASC--Volume:ASC--Author:ASC--Title:ASC--Subtitle:ASC--Edition:ASC--Lexile:ASC--InterestLevel:ASC--AR:ASC--LearningAZ:ASC--GuidedReading:ASC--DRA:ASC--FountasPinnell:ASC--ReadingRecovery:ASC--PMReaders:ASC--Grade:ASC--Age:ASC--Publisher:ASC--LibraryOfCongress:ASC--FictionGenre:ASC"

LETTER_WIDTH = dict([[" ",27.797],["!",27.797],["\"",35.5],["#",55.625],["$",55.625],["%",88.922],["&",66.703],["'",19.094],["(",33.313],[")",33.313],["*",38.922],["+",58.406],[",",27.797],["-",33.313],[".",27.797],["/",27.797],["0",55.625],["1",55.625],["2",55.625],["3",55.625],["4",55.625],["5",55.625],["6",55.625],["7",55.625],["8",55.625],["9",55.625],[":",27.797],[";",27.797],["<",58.406],["=",58.406],[">",58.406],["?",55.625],["@",101.516],["A",66.703],["B",66.703],["C",72.219],["D",72.219],["E",66.703],["F",61.094],["G",77.797],["H",72.219],["I",27.797],["J",50],["K",66.703],["L",55.625],["M",83.313],["N",72.219],["O",77.797],["P",66.703],["Q",77.797],["R",72.219],["S",66.703],["T",61.094],["U",72.219],["V",66.703],["W",94.391],["X",66.703],["Y",66.703],["Z",61.094],["[",27.797],["\\",27.797],["]",27.797],["^",46.938],["_",55.625],["`",33.313],["a",55.625],["b",55.625],["c",50],["d",55.625],["e",55.625],["f",27.797],["g",55.625],["h",55.625],["i",22.219],["j",22.219],["k",50],["l",22.219],["m",83.313],["n",55.625],["o",55.625],["p",55.625],["q",55.625],["r",33.313],["s",50],["t",27.797],["u",55.625],["v",50],["w",72.219],["x",50],["y",50],["z",50],["{",33.406],["|",25.984],["}",33.406],["~",58.406],["_median",55.625]])
KERNING_MODIFIER = dict([[" A",-5.531],[" T",-1.829],[" Y",-1.813],["11",-7.438],["A ",-5.531],["AT",-7.422],["AV",-7.422],["AW",-3.719],["AY",-7.422],["Av",-1.797],["Aw",-1.813],["Ay",-1.797],["F,",-11.094],["F.",-11.094],["FA",-5.531],["L ",-3.734],["LT",-7.438],["LV",-7.422],["LW",-7.438],["LY",-7.422],["Ly",-3.719],["P ",-1.813],["P,",-12.906],["P.",-12.906],["PA",-7.422],["RT",-1.813],["RV",-1.813],["RW",-1.813],["RY",-1.813],["T ",-1.829],["T,",-11.094],["T-",-5.532],["T.",-11.094],["T:",-11.094],["T;",-11.094],["TA",-7.422],["TO",-1.828],["Ta",-11.094],["Tc",-11.094],["Te",-11.094],["Ti",-3.719],["To",-11.094],["Tr",-3.72],["Ts",-11.094],["Tu",-3.719],["Tw",-5.516],["Ty",-5.516],["V,",-9.188],["V-",-5.532],["V.",-9.188],["V:",-3.719],["V;",-3.719],["VA",-7.422],["Va",-7.422],["Ve",-5.531],["Vi",-1.813],["Vo",-5.531],["Vr",-3.719],["Vu",-3.719],["Vy",-3.703],["W,",-5.532],["W-",-1.813],["W.",-5.532],["W:",-1.813],["W;",-1.813],["WA",-3.719],["Wa",-3.719],["We",-1.813],["Wo",-1.813],["Wr",-1.813],["Wu",-1.813],["Wy",-0.875],["Y ",-1.813],["Y,",-12.906],["Y-",-9.188],["Y.",-12.906],["Y:",-5.531],["Y;",-6.5],["YA",-7.422],["Ya",-7.422],["Ye",-9.187],["Yi",-3.703],["Yo",-9.187],["Yp",-7.422],["Yq",-9.187],["Yu",-5.531],["Yv",-5.516],["ff",-1.828],["r,",-5.532],["r.",-5.532],["v,",-7.422],["v.",-7.422],["w,",-5.532],["w.",-5.532],["y,",-7.422],["y.",-7.422]])

BASE_URL = "http://library.jakekausler.com"
CASE_URL = BASE_URL + "/libraries/43/cases"
BOOK_URL = BASE_URL + "/books/bookscases?libraryid=43"
DIVIDERS_URL = BASE_URL + "/libraries/shelfdividers/{}"
SESSION = "MTY3OTMxNDY4MXxEdi1CQkFFQ180SUFBUkFCRUFBQVlfLUNBQUVHYzNSeWFXNW5EQmtBRjJ4cFluSmhjbmx2Y21kaGJtbDZaWEp6WlhOemFXOXVCbk4wY21sdVp3dzBBREpTUjBsTFQxWjZhbGhsZEhveGNqQmxjbmw1UnpWS1dqWTRXbHAyU2tVMWJESndia2R2Yms5alltaGllR05LYWxkRlJnPT186V3Uyx2gOvodfhxhH_jl3y-Gp33G7Vvx__owOuQzrkA="
COOKIES = {'session': SESSION}
CASE_MARGIN = 50
STACKED_CASE_PADDING = 15

case_data = None
book_data = None
if RELOAD:
	r = requests.get(CASE_URL, cookies=COOKIES)
	case_data = r.json()
	with open('cases.json', 'w') as f:
		json.dump(case_data, f)
	r = requests.get(BOOK_URL, cookies=COOKIES)
	book_data = r.json()
	with open('books.json', 'w') as f:
		json.dump(book_data, f)
else:
	with open('cases.json') as f:
		case_data = json.load(f)
	with open('books.json') as f:
		book_data = json.load(f)

class Case:
	def __init__(self, ID, SpacerHeight, BookMargin, Width, Shelves, AverageBookHeight, AverageBookWidth, CaseNumber):
		self.ID = ID
		self.SpacerHeight = SpacerHeight
		self.BookMargin = BookMargin
		self.Width = Width
		self.Shelves = Shelves
		self.AverageBookHeight = AverageBookHeight
		self.AverageBookWidth = AverageBookWidth
		self.CaseNumber = CaseNumber

class Shelf:
	def __init__(self, ID, CaseID, ShelfNumber, Width, Height, PaddingLeft, PaddingRight, Alignment, IsTop, DoNotUse, Books):
		self.ID = ID
		self.CaseID = CaseID
		self.ShelfNumber = ShelfNumber
		self.Width = Width
		self.Height = Height
		self.PaddingLeft = PaddingLeft
		self.PaddingRight = PaddingRight
		self.Alignment = Alignment
		self.IsTop = IsTop
		self.DoNotUse = DoNotUse
		self.Books = Books

class Book:
	def __init__(self, ID, Title, Subtitle, OriginallyPublished, Publisher, UserRead, IsReference, IsOwned, ISBN, Loanee, Dewey, Pages, Width, Height, Depth, Weight, PrimaryLanguage, SecondaryLanguage, OriginalLanguage, Series, Volume, Format, Edition, IsReading, IsShipping, ImageURL, SpineColor, SpineColorOverridden, CheapestNew, CheapestUsed, EditionPublished, Contributors, Library, Lexile, LexileCode, InterestLevel, AR, LearningAZ, GuidedReading, DRA, Grade, FountasPinnell, Age, ReadingRecovery, PMReaders, Awards, Notes, Tags, IsAnthology, IsSideways, PreviousSideways):
		self.ID = ID
		self.Title = Title
		self.Subtitle = Subtitle
		self.OriginallyPublished = OriginallyPublished
		self.Publisher = Publisher
		self.UserRead = UserRead
		self.IsReference = IsReference
		self.IsOwned = IsOwned
		self.ISBN = ISBN
		self.Loanee = Loanee
		self.Dewey = Dewey
		self.Pages = Pages
		self.Width = Width
		self.Height = Height
		self.Depth = Depth
		self.Weight = Weight
		self.PrimaryLanguage = PrimaryLanguage
		self.SecondaryLanguage = SecondaryLanguage
		self.OriginalLanguage = OriginalLanguage
		self.Series = Series
		self.Volume = Volume
		self.Format = Format
		self.Edition = Edition
		self.IsReading = IsReading
		self.IsShipping = IsShipping
		self.ImageURL = ImageURL
		self.SpineColor = SpineColor
		self.SpineColorOverridden = SpineColorOverridden
		self.CheapestNew = CheapestNew
		self.CheapestUsed = CheapestUsed
		self.EditionPublished = EditionPublished
		self.Contributors = Contributors
		self.Library = Library
		self.Lexile = Lexile
		self.LexileCode = LexileCode
		self.InterestLevel = InterestLevel
		self.AR = AR
		self.LearningAZ = LearningAZ
		self.GuidedReading = GuidedReading
		self.DRA = DRA
		self.Grade = Grade
		self.FountasPinnell = FountasPinnell
		self.Age = Age
		self.ReadingRecovery = ReadingRecovery
		self.PMReaders = PMReaders
		self.Awards = Awards
		self.Notes = Notes
		self.Tags = Tags
		self.IsAnthology = IsAnthology
		self.IsSideways = IsSideways
		self.PreviousSideways = PreviousSideways
		self.x = 0
		self.y = 0

def form_cases(case_data):
	pprint(case_data)
	cases = []
	for case in case_data:
		cases.append(Case(
			case["id"],
			case["spacerheight"],
			case["bookmargin"],
			case["width"],
			[Shelf(
				shelf["id"],
				shelf["caseid"],
				shelf["shelfnumber"],
				shelf["width"],
				shelf["height"],
				shelf["paddingleft"],
				shelf["paddingright"],
				shelf["alignment"],
				shelf["istop"],
				shelf["DoNotUse"],
				shelf["books"] if shelf["books"] else []
				) for shelf in case["shelves"]],
			case["averagebookheight"],
			case["averagebookwidth"],
			case["casenumber"]
			))
	return cases

def form_books(book_data):
	books = []
	for book in book_data:
		books.append(Book(
			book["bookid"],
			book["title"],
			book["subtitle"],
			book["originallypublished"],
			book["publisher"],
			book["userread"],
			book["isreference"],
			book["isowned"],
			book["isbn"],
			book["loanee"],
			book["dewey"],
			book["pages"],
			book["width"],
			book["height"],
			book["depth"],
			book["weight"],
			book["primarylanguage"],
			book["secondarylanguage"],
			book["originallanguage"],
			book["series"],
			book["volume"],
			book["format"],
			book["edition"],
			book["IsReading"],
			book["isshipping"],
			book["imageurl"],
			book["spinecolor"],
			book["spinecoloroverridden"],
			book["cheapestnew"],
			book["cheapestused"],
			book["editionpublished"],
			book["contributors"],
			book["library"],
			book["lexile"],
			book["lexilecode"],
			book["interestlevel"],
			book["ar"],
			book["learningaz"],
			book["guidedreading"],
			book["dra"],
			book["grade"],
			book["fountaspinnell"],
			book["age"],
			book["readingrecovery"],
			book["pmreaders"],
			book["awards"],
			book["notes"],
			book["tags"],
			book["isanthology"],
			book["issideways"],
			book["previoussideways"]))
	return books

def GetDividers(shelf_id):
	r = requests.get(DIVIDERS_URL.format(shelf_id), cookies=COOKIES)
	return [] if not r.json() else r.json()

cases = form_cases(case_data)
books = form_books(book_data)

def sort_contributors(c):
	return ('0' if c['role'] == 'Author' else '1') + c['name']['last'] + c['name']['first'] + c['name']['middles']

def check_next_contributor_equality(books, i):
	return (
		(len(books[i].Contributors) == 0 and (books[i+1].Contributors) == 0) or
		(len(books[i].Contributors) > 0 and len(books[i+1].Contributors) > 0 and books[i].Contributors[0] == books[i+1].Contributors[0]))

def sort_books(books):
	sort(books)

def fill_cases(cases, books):
	# sort_books(books)
	oversized = []
	index = 0
	for case in cases:
		index = fill_case(case, books, index, oversized)
	return oversized

def fill_case(case, books, index, oversized):
	for shelf in case.Shelves:
		index = fill_shelf(case, shelf, books, index, oversized)
	return index

def fill_shelf(case, shelf, books, index, oversized):
	if shelf.DoNotUse or index >= len(books):
		return index
	currentBook = books[index]
	x = shelf.PaddingLeft
	useWidth = get_use_width(case, currentBook)
	dividers = GetDividers(shelf.ID)
	currentDivider = 0
	while index < len(books) and useWidth+x <= shelf.Width-shelf.PaddingRight:
		index, x, currentDivider = fill_book(case, shelf, books, index, currentBook, useWidth, x, oversized, dividers, currentDivider)
		if index < len(books):
			currentBook = books[index]
			useWidth = get_use_width(case, currentBook)
	return index

def fill_book(case, shelf, books, index, currentBook, useWidth, x, oversized, dividers, currentDivider):
	if currentBook.Height > shelf.Height:
		oversized.append(currentBook)
		return index + 1, x, currentDivider
	if not currentBook.IsSideways and oversized:
		for i, b in enumerate(oversized):
			if b.Height <= shelf.Height and x + b.Width <= shelf.Width-shelf.PaddingRight:
				books.insert(index, b)
				currentBook = books[index]
				useWidth = get_use_width(case, currentBook)
				del oversized[i]
				break
	x, currentDivider = check_fill_dividers(case, x, useWidth, dividers, currentDivider)
	y = check_sideways(case, shelf, books, index, currentBook, useWidth, x, dividers, currentDivider)
	useWidth = add_book(shelf, books, index, currentBook, useWidth, x, y)
	return index + 1, x + useWidth, currentDivider

def add_book(shelf, books, index, currentBook, useWidth, x, y):
	currentBook.x = x
	currentBook.y = y
	if currentBook.IsSideways:
		if index < len(books)-1 and not books[index+1].PreviousSideways:
			shelf.Books.append(currentBook)
			useWidth = currentBook.Height
		else:
			shelf.Books.append(currentBook)
			useWidth = 0
	else:
		shelf.Books.append(currentBook)
	return useWidth

def check_sideways(case, shelf, books, index, currentBook, useWidth, x, dividers, currentDivider):
	y = 0
	if (not currentBook.IsSideways and
		 index < len(books) - 1 and
		 shelf.IsTop == 0 and
		 currentBook.Height == books[index+1].Height and
		 currentBook.Height + x <= shelf.Width - shelf.PaddingRight and
		 (len(dividers) == 0 or currentDivider >= len(dividers) or x+currentBook.Height <= dividers[currentDivider]['distancefromleft'])):
		maxStackHeight = shelf.Height - STACKED_CASE_PADDING
		stackHeight = currentBook.Width
		checkIndex = index + 1
		while checkIndex < len(books) and books[checkIndex].Height == currentBook.Height and stackHeight + books[checkIndex].Width < maxStackHeight:
			stackHeight += books[checkIndex].Width
			checkIndex += 1
		if stackHeight > currentBook.Height:
			for i in range(index, checkIndex):
				books[i].IsSideways = True
				if i > index:
					y += books[i-1].Width
					books[i].PreviousSideways = books[i-1]
	return y

def check_fill_dividers(case, x, useWidth, dividers, currentDivider):
	if currentDivider < len(dividers):
		if x + useWidth > dividers[currentDivider]['distancefromleft']:
			x = dividers[currentDivider]['distancefromleft'] + dividers[currentDivider]['width'] + case.BookMargin
			currentDivider += 1
	return x, currentDivider

def get_use_width(case, currentBook):
	useWidth = case.AverageBookWidth
	if currentBook.Width > 0:
		useWidth = currentBook.Width
	return useWidth + case.BookMargin

fill_cases(cases, books)

def get_max_case_height(cases):
	maxCaseHeight = 0
	for c in cases:
		height = c.SpacerHeight * (len(c.Shelves) + 1)
		for s in c.Shelves:
			height += s.Height
		if maxCaseHeight < height:
			maxCaseHeight = height
	maxCaseHeight += CASE_MARGIN * 2
	return maxCaseHeight

def get_case_dimensions(c):
	maxShelfWidthBooks = 0
	totalCaseHeight = 0
	for s in c.Shelves:
		width = s.Width
		if width > maxShelfWidthBooks:
			maxShelfWidthBooks = width
		totalCaseHeight += s.Height
	caseWidth = c.SpacerHeight * 2 + maxShelfWidthBooks + CASE_MARGIN*2
	caseHeight = c.SpacerHeight * (len(c.Shelves) + 1) + totalCaseHeight
	return maxShelfWidthBooks, totalCaseHeight, caseWidth, caseHeight

def draw_case(CaseSvgPath, c, maxCaseHeight):
	maxShelfWidthBooks, totalCaseHeight, caseWidth, caseHeight = get_case_dimensions(c)
	caseCanvas = svgwrite.Drawing('{}/{}.svg'.format(CaseSvgPath, c.ID), size=(caseWidth, maxCaseHeight))
	y = maxCaseHeight - caseHeight
	for s in c.Shelves:
		y = draw_shelf(c, s, y, caseCanvas, caseWidth)
	caseCanvas.save()

def set_initial_x(c, s, caseWidth):
	x = CASE_MARGIN
	if s.Alignment == "right":
		x = caseWidth - (s.Width + 2*c.SpacerHeight) - CASE_MARGIN
	return x

def draw_top_and_sides(c, s, x, y, caseCanvas):
	if s.IsTop != 1:
		# Add the left and right shelf borders
		caseCanvas.add(caseCanvas.rect((x, y), (c.SpacerHeight, s.Height+c.SpacerHeight)))
		caseCanvas.add(caseCanvas.rect((x+c.SpacerHeight+s.Width, y), (c.SpacerHeight, s.Height+c.SpacerHeight)))

		# Add the top border
		caseCanvas.add(caseCanvas.rect((x+c.SpacerHeight, y), (s.Width, c.SpacerHeight)))

def add_dividers(c, s, x, y, caseCanvas, dividers):
	for d in dividers:
		if d['height'] == 0:
			if d['imageurl'] == "":
				caseCanvas.add(caseCanvas.rect((CASE_MARGIN+c.SpacerHeight+d['distancefromleft'], y-s.Height), (d['width'], s.Height)))
			else:
				caseCanvas.add(caseCanvas.image(BASE_URL + d['imageurl'], (CASE_MARGIN+c.SpacerHeight+d['distancefromleft'], y-s.Height), (d['width'], s.Height)))
		else:
			if d['imageurl'] == "":
				caseCanvas.add((caseCanvas.rect(CASE_MARGIN+c.SpacerHeight+d['distancefromleft'], y-d['height']), (d['width'], d['height'])))
			else:
				caseCanvas.add(caseCanvas.image(BASE_URL + d['imageurl'], (CASE_MARGIN+c.SpacerHeight+d['distancefromleft'], y-d['height']), (d['width'], d['height'])))

def add_bottom_border(c, s, x, y, caseCanvas):
	caseCanvas.add(caseCanvas.rect((x, y), (s.Width+c.SpacerHeight*2, c.SpacerHeight)))

def get_font_color(b):
	fontColor = ""
	c = b.SpineColor[1:]
	col = tuple(int(c[i:i+2], 16) for i in (0, 2, 4))
	o = round((col[0]*299+col[1]*587+col[2]*114) / 1000)
	if o > 125:
		fontColor = "black"
	else:
		fontColor = "white"
	return fontColor

def measure_text(st, sz):
	width = 0
	for i, c in enumerate(st):
		width += LETTER_WIDTH[c] if c in LETTER_WIDTH else LETTER_WIDTH['_median']
		if i < len(st)-1:
			width += KERNING_MODIFIER[c + st[i+1]] if c + st[i+1] in KERNING_MODIFIER else 0
	return width * sz / 100

def draw_title(b, x, y, caseCanvas, previousSidewaysHeight):
	fontSize = min(18, b.Width/4*3)
	fontColor = get_font_color(b)
	title = b.Title
	while measure_text(title, fontSize) > b.Height - 4:
		title = title[:-1]
	if not b.IsSideways:
		text_x = x+b.Width/2
		text_y = y - 4
		rotation = -90
		caseCanvas.add(caseCanvas.text(title, (text_x, text_y), transform='rotate({},{},{})'.format(rotation, text_x, text_y), dominant_baseline="middle", font_family="Arial", font_size="{}px".format(fontSize), fill=fontColor))
	else:
		text_x = x
		text_y = y - previousSidewaysHeight - b.Width/2
		caseCanvas.add(caseCanvas.text(title, (text_x, text_y), dominant_baseline="middle", font_family="Arial", font_size="{}px".format(fontSize), fill=fontColor))
		pass

def get_next_x(c, s, b, x, bidx):
	if not b.IsSideways:
		x += b.Width + c.BookMargin
	elif b.IsSideways and bidx < len(s.Books) - 1 and not s.Books[bidx+1].PreviousSideways:
		x += b.Height + c.BookMargin
	return x

def fix_spine_color(b):
	b.SpineColor = b.SpineColor.replace("-", "")
	while len(b.SpineColor) < 7:
		b.SpineColor += "0"

def set_book_width_and_height(c, b):
	if b.Width <= 0:
		b.Width = c.AverageBookWidth
	if b.Height <= 0:
		b.Height = c.AverageBookHeight

def draw_book_box(c, b, x, y, caseCanvas):
	previousSidewaysHeight = 0
	if b.IsSideways:
		previousSideways = b.PreviousSideways
		while previousSideways != None:
			previousSidewaysHeight += previousSideways.Width
			previousSideways = previousSideways.PreviousSideways
		caseCanvas.add(caseCanvas.rect((x, y-b.Width-previousSidewaysHeight), (b.Height, b.Width), id='book-{}'.format(b.ID), class_="bookcase-book", stroke='black', fill='{}'.format(b.SpineColor)))
	else:
		caseCanvas.add(caseCanvas.rect((x, y-b.Height), (b.Width, b.Height), id='book-{}'.format(b.ID), class_='bookcase-book', stroke='black', fill='{}'.format(b.SpineColor)))
	return previousSidewaysHeight

def checkDividers(c, b, x, dividers, currentDivider):
	if currentDivider < len(dividers):
		if x+b.Width > CASE_MARGIN+c.SpacerHeight+dividers[currentDivider]['distancefromleft']:
			x = CASE_MARGIN+c.SpacerHeight+dividers[currentDivider]['distancefromleft'] + dividers[currentDivider]['width'] + c.BookMargin
			currentDivider += 1
	return x, currentDivider

def draw_book(c, s, b, bidx, x, y, caseCanvas, dividers, currentDivider):
	fix_spine_color(b)
	set_book_width_and_height(c, b)
	x, currentDivider = checkDividers(c, b, x, dividers, currentDivider)
	previousSidewaysHeight = draw_book_box(c, b, x, y, caseCanvas)
	draw_title(b, x, y, caseCanvas, previousSidewaysHeight)
	# Move x to start of next book (if not in a sideways stack, or if in the last book of a sideways stack)
	x = get_next_x(c, s, b, x, bidx)
	return x, currentDivider

def draw_shelf(c, s, y, caseCanvas, caseWidth):
	dividers = GetDividers(s.ID)
	x = set_initial_x(c, s, caseWidth)
	draw_top_and_sides(c, s, x, y, caseCanvas)
	# Update the current y to bottom of current shelf
	y += c.SpacerHeight + s.Height
	add_dividers(c, s, x, y, caseCanvas, dividers)
	# Add the bottom border
	add_bottom_border(c, s, x, y, caseCanvas)
	# Update current x to inside of shelf
	x += c.SpacerHeight + s.PaddingLeft
	currentDivider = 0
	for bidx, b in enumerate(s.Books):
		x, currentDivider = draw_book(c, s, b, bidx, x, y, caseCanvas, dividers, currentDivider)
	return y

def draw_cases(cases):
	CaseSvgPath = './testsvgs'

	maxCaseHeight = get_max_case_height(cases)

	for c in cases:
		draw_case(CaseSvgPath, c, maxCaseHeight)

draw_cases(cases)