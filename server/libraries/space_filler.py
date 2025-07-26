
import random
import pygame
import json

BLACK = ( 0, 0, 0)
WHITE = ( 255, 255, 255)

RELOAD = False

MULTIPLIER = 1

SORT_METHOD = "Dewey:ASC--Author:ASC--SeriesHeight:ASC--Series:ASC--Volume:ASC--Title:ASC--Subtitle:ASC--Edition:ASC--Lexile:ASC--InterestLevel:ASC--AR:ASC--LearningAZ:ASC--GuidedReading:ASC--DRA:ASC--FountasPinnell:ASC--ReadingRecovery:ASC--PMReaders:ASC--Grade:ASC--Age:ASC--Publisher:ASC--LibraryOfCongress:ASC--FictionGenre:ASC||Dewey:ASC--SeriesHeight:ASC--Series:ASC--Volume:ASC--Author:ASC--Title:ASC--Subtitle:ASC--Edition:ASC--Lexile:ASC--InterestLevel:ASC--AR:ASC--LearningAZ:ASC--GuidedReading:ASC--DRA:ASC--FountasPinnell:ASC--ReadingRecovery:ASC--PMReaders:ASC--Grade:ASC--Age:ASC--Publisher:ASC--LibraryOfCongress:ASC--FictionGenre:ASC"

LETTER_WIDTH = dict([[" ",27.797],["!",27.797],["\"",35.5],["#",55.625],["$",55.625],["%",88.922],["&",66.703],["'",19.094],["(",33.313],[")",33.313],["*",38.922],["+",58.406],[",",27.797],["-",33.313],[".",27.797],["/",27.797],["0",55.625],["1",55.625],["2",55.625],["3",55.625],["4",55.625],["5",55.625],["6",55.625],["7",55.625],["8",55.625],["9",55.625],[":",27.797],[";",27.797],["<",58.406],["=",58.406],[">",58.406],["?",55.625],["@",101.516],["A",66.703],["B",66.703],["C",72.219],["D",72.219],["E",66.703],["F",61.094],["G",77.797],["H",72.219],["I",27.797],["J",50],["K",66.703],["L",55.625],["M",83.313],["N",72.219],["O",77.797],["P",66.703],["Q",77.797],["R",72.219],["S",66.703],["T",61.094],["U",72.219],["V",66.703],["W",94.391],["X",66.703],["Y",66.703],["Z",61.094],["[",27.797],["\\",27.797],["]",27.797],["^",46.938],["_",55.625],["`",33.313],["a",55.625],["b",55.625],["c",50],["d",55.625],["e",55.625],["f",27.797],["g",55.625],["h",55.625],["i",22.219],["j",22.219],["k",50],["l",22.219],["m",83.313],["n",55.625],["o",55.625],["p",55.625],["q",55.625],["r",33.313],["s",50],["t",27.797],["u",55.625],["v",50],["w",72.219],["x",50],["y",50],["z",50],["{",33.406],["|",25.984],["}",33.406],["~",58.406],["_median",55.625]])
KERNING_MODIFIER = dict([[" A",-5.531],[" T",-1.829],[" Y",-1.813],["11",-7.438],["A ",-5.531],["AT",-7.422],["AV",-7.422],["AW",-3.719],["AY",-7.422],["Av",-1.797],["Aw",-1.813],["Ay",-1.797],["F,",-11.094],["F.",-11.094],["FA",-5.531],["L ",-3.734],["LT",-7.438],["LV",-7.422],["LW",-7.438],["LY",-7.422],["Ly",-3.719],["P ",-1.813],["P,",-12.906],["P.",-12.906],["PA",-7.422],["RT",-1.813],["RV",-1.813],["RW",-1.813],["RY",-1.813],["T ",-1.829],["T,",-11.094],["T-",-5.532],["T.",-11.094],["T:",-11.094],["T;",-11.094],["TA",-7.422],["TO",-1.828],["Ta",-11.094],["Tc",-11.094],["Te",-11.094],["Ti",-3.719],["To",-11.094],["Tr",-3.72],["Ts",-11.094],["Tu",-3.719],["Tw",-5.516],["Ty",-5.516],["V,",-9.188],["V-",-5.532],["V.",-9.188],["V:",-3.719],["V;",-3.719],["VA",-7.422],["Va",-7.422],["Ve",-5.531],["Vi",-1.813],["Vo",-5.531],["Vr",-3.719],["Vu",-3.719],["Vy",-3.703],["W,",-5.532],["W-",-1.813],["W.",-5.532],["W:",-1.813],["W;",-1.813],["WA",-3.719],["Wa",-3.719],["We",-1.813],["Wo",-1.813],["Wr",-1.813],["Wu",-1.813],["Wy",-0.875],["Y ",-1.813],["Y,",-12.906],["Y-",-9.188],["Y.",-12.906],["Y:",-5.531],["Y;",-6.5],["YA",-7.422],["Ya",-7.422],["Ye",-9.187],["Yi",-3.703],["Yo",-9.187],["Yp",-7.422],["Yq",-9.187],["Yu",-5.531],["Yv",-5.516],["ff",-1.828],["r,",-5.532],["r.",-5.532],["v,",-7.422],["v.",-7.422],["w,",-5.532],["w.",-5.532],["y,",-7.422],["y.",-7.422]])

BASE_URL = "http://library.jakekausler.com"
CASE_URL = BASE_URL + "/libraries/43/cases"
BOOK_URL = BASE_URL + "/books/bookscases?libraryid=43"
DIVIDERS_URL = BASE_URL + "/libraries/shelfdividers/{}"
SESSION = "MTY3MTAyNTUyN3xEdi1CQkFFQ180SUFBUkFCRUFBQVlfLUNBQUVHYzNSeWFXNW5EQmtBRjJ4cFluSmhjbmx2Y21kaGJtbDZaWEp6WlhOemFXOXVCbk4wY21sdVp3dzBBREp2TjJ0aVlVSlZaVTlTTmtzMFJGVTBSM0J5UzJOVlZYQkRRVWRZWW1oMVlXODNOWEJIU2pZeWN6UnZTMUY2Ykd4WU5nPT18ONm9FUFp-36C_SZNF-r5WBmH1v8cTS8iBh2bux-qXcM="
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

	def __str__(self):
		return "Book idx {} (previousSideways is {})".format(self.ID, self.PreviousSideways.ID if self.PreviousSideways else '-')

def form_cases(case_data):
	width = 0
	height = 0
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

# books = []
# for i in range(0,100):
# 	books.append(Book(i, random.randrange(1,5), random.randrange(5,10)))

def insertSideBook(shelf, b, sidewaysLength, previousSideways, x, y):
	b.x = x
	b.y = y
	b.IsSideways = True
	b.PreviousSideways = previousSideways
	shelf.Books.append(b)
	print("Side\n\t{}\n\tX/Y: {}, {}\n\tW/H: {}, {}".format(str(b), x, y, b.Width, b.Height))
	previousSideways = b
	if y == 0:
		sidewaysLength = b.Height
	y += b.Width
	return previousSideways, sidewaysLength, x, y

def determineSideways(previousSideways, lengthOnSideways, b, x, y, shelfWidth, shelfHeight, books, i):
	if previousSideways and lengthOnSideways == 0:
		if b.Height == previousSideways.Height and b.Width + y <= shelfHeight and b.Height + x <= shelfWidth:
			doneChecking = False
			shouldBeSideways = True
			checkingFutureSideways = True
			sidewaysLength = 0
			checkIndex = i+1
			newY = y+b.Height
			while not doneChecking and checkIndex < len(books):
				newBook = books[checkIndex]
				if checkingFutureSideways and newBook.Height == b.Height and newBook.Width + newY <= shelfHeight and newBook.Height + x <= shelfWidth:
					newY += newBook.Width
				else:
					checkingFutureSideways = False
					if newBook.Width + sidewaysLength <= b.Height:
						if newBook.Height + newY <= shelfHeight:
							sidewaysLength += newBook.Width
						else:
							doneChecking = True
							shouldBeSideways = b.Height <= newY
					else:
						doneChecking = True
				checkIndex += 1
			return shouldBeSideways
	return False

def fill_cases(cases, books):
	currentCase = 0
	currentShelf = 0
	x = 0
	y = 0
	sidewaysLength = 0
	previousSideways = None
	lengthOnSideways = 0
	shelfWidth = cases[currentCase].Shelves[currentShelf].Width - cases[currentCase].Shelves[currentShelf].PaddingLeft - cases[currentCase].Shelves[currentShelf].PaddingRight
	shelfHeight = cases[currentCase].Shelves[currentShelf].Height - STACKED_CASE_PADDING
	print(shelfWidth, shelfHeight)
	for i, b in enumerate(books):
		if b.Height > 300:
			continue
		if determineSideways(previousSideways, lengthOnSideways, b, x, y, shelfWidth, shelfHeight, books, i):
			previousSideways, sidewaysLength, x, y = insertSideBook(cases[currentCase].Shelves[currentShelf], b, sidewaysLength, previousSideways, x, y)
		else:
			foundRoom = False
			if previousSideways and y + b.Height <= shelfHeight and x + lengthOnSideways + b.Width <= x + previousSideways.Height:
				foundRoom = True
			while not foundRoom and previousSideways:
				# print("Testing tall book idx {} with width {} and height {}. Current x: {}. Current y: {}. Current Length on Sideways: {}. Previous book idx {} has a height of {}".format(b.Priority, b.Width, b.Height, x, y, lengthOnSideways, previousSideways.Priority, previousSideways.Height))
				# print("\t{} + {} <= {} is {}".format(y, b.Height, shelfHeight, y + b.Height <= shelfHeight))
				# print("\t{} + {} + {} <= {} + {} is {}".format(x, lengthOnSideways, b.Width, x, previousSideways.Height, x + lengthOnSideways + b.Width <= x + previousSideways.Height))
				if y + b.Height <= shelfHeight and x + lengthOnSideways + b.Width <= x + previousSideways.Height:
					foundRoom = True
				if not foundRoom:
					lengthOnSideways = previousSideways.Height
					y -= previousSideways.Width
					previousSideways = previousSideways.PreviousSideways
			if previousSideways and y + b.Height <= shelfHeight and x + lengthOnSideways + b.Width <= x + previousSideways.Height:
				b.x = x+lengthOnSideways
				b.y = y
				b.IsSideways = False
				b.PreviousSideways = previousSideways
				cases[currentCase].Shelves[currentShelf].Books.append(b)
				print("Tall\n\t{}\n\tX/Y: {}, {}\n\tW/H: {}, {}".format(str(b), x, y, b.Width, b.Height))
				lengthOnSideways += b.Width
			else:
				print("\nStarting new stack\n")
				y = 0
				x += sidewaysLength
				previousSideways = None
				lengthOnSideways = 0
				if b.Height + x > shelfWidth:
					if b.Width + x <= shelfWidth:
						b.x = x
						b.y = y
						b.IsSideways = False
						b.PreviousSideways = previousSideways
						cases[currentCase].Shelves[currentShelf].Books.append(b)
						print("Tall\n\t{}\n\tX/Y: {}, {}\n\tW/H: {}, {}".format(str(b), x, y, b.Width, b.Height))
						x += b.Width
					else:
						print("\nStarting new shelf\n")
						currentShelf += 1
						if currentShelf >= len(cases[currentCase].Shelves):
							currentShelf = 0
							currentCase += 1
							if currentCase >= len(cases):
								return
							print(shelfWidth, shelfHeight)
						shelfWidth = cases[currentCase].Shelves[currentShelf].Width - cases[currentCase].Shelves[currentShelf].PaddingLeft - cases[currentCase].Shelves[currentShelf].PaddingRight
						shelfHeight = cases[currentCase].Shelves[currentShelf].Height - STACKED_CASE_PADDING
						x = 0
						previousSideways, sidewaysLength, x, y = insertSideBook(cases[currentCase].Shelves[currentShelf], b, sidewaysLength, previousSideways, x, y)
				else:
					previousSideways, sidewaysLength, x, y = insertSideBook(cases[currentCase].Shelves[currentShelf], b, sidewaysLength, previousSideways, x, y)

def pygameRect(book, cases, caseNumber, shelfNumber):
	x = CASE_MARGIN + caseNumber * (CASE_MARGIN+CASE_MARGIN) + sum(max(shelf.Width for shelf in case.Shelves) for case in cases[:caseNumber]) + cases[caseNumber].Shelves[shelfNumber].PaddingLeft
	y = CASE_MARGIN + sum(shelf.Height for shelf in cases[caseNumber].Shelves[:shelfNumber+1]) + (shelfNumber) * cases[caseNumber].SpacerHeight
	return pygame.Rect(
		int((x + book.x)*MULTIPLIER),
		int((y - book.y - (book.Width if book.IsSideways else book.Height))*MULTIPLIER),
		int((book.Height if book.IsSideways else book.Width)*MULTIPLIER),
		int((book.Width if book.IsSideways else book.Height)*MULTIPLIER)
		)

def drawCases(screen, cases):
	for cn, case in enumerate(cases):
		for sn, shelf in enumerate(case.Shelves):
			for book in shelf.Books:
				# print(entry.Book)
				pygame.draw.rect(screen, BLACK, pygameRect(book, cases, cn, sn), 1)
			pygame.draw.line(screen, (0,255,0), (
				(CASE_MARGIN+CASE_MARGIN*cn+shelf.PaddingLeft)*MULTIPLIER,
				(CASE_MARGIN+sum(s.Height+case.SpacerHeight for s in case.Shelves[:sn]))*MULTIPLIER
				),(
				(CASE_MARGIN+CASE_MARGIN*cn+shelf.PaddingLeft)*MULTIPLIER,
				(CASE_MARGIN+sum(s.Height+case.SpacerHeight for s in case.Shelves[:sn])+shelf.Height)*MULTIPLIER),
				1)
			pygame.draw.line(screen, (0,255,0), (
				(CASE_MARGIN+CASE_MARGIN*cn+shelf.Width-shelf.PaddingRight)*MULTIPLIER,
				(CASE_MARGIN+sum(s.Height+case.SpacerHeight for s in case.Shelves[:sn]))*MULTIPLIER
				),(
				(CASE_MARGIN+CASE_MARGIN*cn+shelf.Width-shelf.PaddingRight)*MULTIPLIER,
				(CASE_MARGIN+sum(s.Height+case.SpacerHeight for s in case.Shelves[:sn])+shelf.Height)*MULTIPLIER),
				1)
			pygame.draw.line(screen, (0,255,0), (
				(CASE_MARGIN+CASE_MARGIN*cn)*MULTIPLIER,
				(CASE_MARGIN+sum(s.Height+case.SpacerHeight for s in case.Shelves[:sn])+STACKED_CASE_PADDING)*MULTIPLIER
				), (
				(CASE_MARGIN+CASE_MARGIN*cn+shelf.Width)*MULTIPLIER,
				(CASE_MARGIN+sum(s.Height+case.SpacerHeight for s in case.Shelves[:sn])+STACKED_CASE_PADDING)*MULTIPLIER),
				1)
			pygame.draw.rect(screen, (0,0,255), (
				(CASE_MARGIN+CASE_MARGIN*cn)*MULTIPLIER,
				(CASE_MARGIN+sum(s.Height+case.SpacerHeight for s in case.Shelves[:sn]))*MULTIPLIER,
				(shelf.Width)*MULTIPLIER,
				(shelf.Height)*MULTIPLIER),
				1)
		pygame.draw.rect(screen, (255,0,0), (
			(CASE_MARGIN+CASE_MARGIN*cn)*MULTIPLIER,
			CASE_MARGIN*MULTIPLIER,
			(max(shelf.Width for shelf in case.Shelves))*MULTIPLIER,
			(sum(shelf.Height for shelf in case.Shelves)+case.SpacerHeight*len(case.Shelves))*MULTIPLIER),
			2)

def getSize(cases):
	width = 0
	height = 0
	for case in cases:
		if height < sum(s.Height for s in case.Shelves) + (len(case.Shelves)+1) * case.SpacerHeight:
			height = sum(s.Height for s in case.Shelves) + (len(case.Shelves)+1) * case.SpacerHeight
		width += max(s.Width for s in case.Shelves)
	return int((width + len(cases) * CASE_MARGIN*2)*MULTIPLIER), int((height + CASE_MARGIN*2)*MULTIPLIER)
	
if __name__ == '__main__':
	cases = form_cases(case_data)
	books = form_books(book_data)
	fill_cases(cases[:1], books)
	stop = False
	size = getSize(cases[:1])
	print(size)
	screen = pygame.display.set_mode(size)

	screen.fill(WHITE)
	drawCases(screen, cases[:1])
	while not stop:
		for event in pygame.event.get():
			if event.type == pygame.QUIT:
				stop = True
		pygame.display.flip()