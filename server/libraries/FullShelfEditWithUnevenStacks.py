
import argparse
import os
import requests
import math
from pprint import pprint
import json
import svgwrite

USE_STACKS = False
USE_TOP_FOR_STACK = True

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
    def __init__(self, ID, Title, Subtitle, OriginallyPublished, Publisher, UserRead, IsReference, IsOwned, ISBN, Loanee, Dewey, Pages, Width, Height, Depth, Weight, PrimaryLanguage, SecondaryLanguage, OriginalLanguage, Series, Volume, Format, Edition, IsReading, IsShipping, ImageURL, SpineColor, SpineColorOverridden, CheapestNew, CheapestUsed, EditionPublished, Contributors, Library, Lexile, LexileCode, InterestLevel, AR, LearningAZ, GuidedReading, DRA, Grade, FountasPinnell, Age, ReadingRecovery, PMReaders, Awards, Notes, Tags, IsAnthology, IsSideways, PreviousSideways, SortDewey, KeepWithNext):
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
        self.SortDewey = SortDewey
        self.KeepWithNext = KeepWithNext
        self.x = 0
        self.y = 0
        self.Checked = False
        self.LastInStack = False
        self.SeriesHeight = 0

def form_cases(case_data):
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
            book["previoussideways"],
            book["sortdewey"],
            book["keepwithnext"]))
    return books

def GetDividers(shelf_id):
    r = requests.get(DIVIDERS_URL.format(shelf_id), cookies=COOKIES)
    return [] if not r.json() else r.json()

def sort_contributors(c):
    return ('0' if c['role'] == 'Author' else '1') + c['name']['last'] + c['name']['first'] + c['name']['middles']

def check_next_contributor_equality(books, i):
    return (
        (len(books[i].Contributors) == 0 and (books[i+1].Contributors) == 0) or
        (len(books[i].Contributors) > 0 and len(books[i+1].Contributors) > 0 and books[i].Contributors[0] == books[i+1].Contributors[0]))

def avg(l):
    return sum(l)/len(l)

def sort_books(books, deweyMergeLevel=-1, useSortDeweys=True, useSortHeights=True):
    # books.sort(key=lambda book: book.Height, reverse=True)
    deweySplits = []
    currentSplit = [books[0]]
    for b in books[1:]:
        # if b.Series.startswith("Reader's Digest"):
        #     print(b.Title)
        #     b.Dewey = "zzzDoNotUse"
        #     b.SortDewey = "zzzDoNotUse"
        try:
            dewey = ""
            oldDewey = ""
            if deweyMergeLevel == -1:
                dewey = b.SortDewey if useSortDeweys else b.Dewey
                oldDewey = currentSplit[-1].SortDewey if useSortDeweys else currentSplit[-1].Dewey
            else:
                dewey = math.floor(float(b.SortDewey if useSortDeweys else b.Dewey)/deweyMergeLevel)*deweyMergeLevel
                oldDewey = math.floor(float(currentSplit[-1].SortDewey if useSortDeweys else currentSplit[-1].Dewey)/deweyMergeLevel)*deweyMergeLevel
            if oldDewey != dewey:
                deweySplits.append(currentSplit)
                currentSplit = []
        except Exception as e:
            if ((currentSplit[-1].SortDewey != b.SortDewey) if useSortDeweys else (currentSplit[-1].Dewey != b.Dewey)):
                deweySplits.append(currentSplit)
                currentSplit = []
        currentSplit.append(b)
        #if currentSplit[-1].Dewey != b.Dewey:
        #    deweySplits.append(currentSplit)
        #    currentSplit = []
        #currentSplit.append(b)
    deweySplits.append(currentSplit)
    for d, ds in enumerate(deweySplits):
        #if ds[0].SortDewey in ["aFIC", "bGEO"]:
        #    continue
        series = []
        for i, b in enumerate(ds):
            if b.Series:
                if i == 0 or not series:
                    series.append(b)
                elif b.Series == ds[i-1].Series:
                    series.append(b)
                else:
                    seriesHeight = avg([s.Height for s in series])
                    for j in range(i-len(series), i):
                        ds[j].SeriesHeight = seriesHeight
                    series = [b]
            elif series:
                seriesHeight = avg([s.Height for s in series])
                for j in range(i-len(series), i):
                    ds[j].SeriesHeight = seriesHeight
                series = []
        if series:
            seriesHeight = avg([s.Height for s in series])
            for j in range(len(ds)-len(series), len(ds)):
                ds[j].SeriesHeight = seriesHeight
        deweySplits[d] = sorted(ds, key=lambda book: (sorted_contributors(book) if ds[0].SortDewey == 'aFIC' else [], (-book.SeriesHeight if book.SeriesHeight else -book.Height) if useSortHeights else 0))
    return [b for ds in deweySplits for b in ds if b.Notes == "Remove"]

def sorted_contributors(book):
    s = sorted([(c['name']['last'], c['name']['first'], c['name']['middles']) for c in filter(lambda c: c['role'] == 'Author', book.Contributors)])
    return s[0] if s else ()
def fill_cases(cases, books):
    # sort_books(books)
    oversized = []
    index = 0
    for case in cases:
        index = fill_case(case, books, index, oversized)
    return oversized

def check_book(cases, books, currentCase, currentShelf, x, shelfIsTop, shelfWidth, shelfHeight, i, dividers, currentDivider, useStacks=USE_STACKS, useTopForStack=USE_TOP_FOR_STACK):
    shouldBreakShelf = False
    shouldBreakDivider = False
    # If the book's width pushes the book over the edge, we should go to the next shelf before checking. We will break the shelf and then keep checking
    if books[i].Width + x > shelfWidth:
        shouldBreakShelf = True
        return shouldBreakShelf, shouldBreakDivider
    # If the book's width pushes the book past the next divider, we should reset past it before checking. We will break to after the divider then keep checking
    if currentDivider < len(dividers) and books[i].Width + x > dividers[currentDivider]['distancefromleft']:
        shouldBreakDivider = True
        return shouldBreakShelf, shouldBreakDivider
    books[i].x = x
    books[i].y = 0
    # If the book's height pushes the book over the edge sideways, there's no point checking. It should be placed normally
    if books[i].Height + x > shelfWidth:
        books[i].Checked = True
        books[i].LastInStack = True
        return shouldBreakShelf, shouldBreakDivider
    # If the book's height pushes the book past the next divider sideways, there's no point checking. It should be placed normally
    if currentDivider < len(dividers) and books[i].Height + x > dividers[currentDivider]['distancefromleft']:
        books[i].Checked = True
        books[i].LastInStack = True
        return shouldBreakShelf, shouldBreakDivider
    # If the book is on the top shelf, we don't want it to be sideways. It should be placed normally
    if not useTopForStack and cases[currentCase].Shelves[currentShelf].IsTop:
        books[i].Checked = True
        books[i].LastInStack = True
        return shouldBreakShelf, shouldBreakDivider
    if useStacks:
        stack = [books[i]]
        stackSidewaysInfo = [(True, None, x, 0)] # IsSideways, PreviousSideways, x, y
        currentY = books[i].Width
        checkingIdx = i + 1
        # print()
        # print("Starting stack with {} at x: {}".format(books[i].Title, x))
        checkingIdx, _ = form_stack(books, stack, stackSidewaysInfo, stack[-1], x, currentY, shelfIsTop, shelfHeight, checkingIdx)
        stackWidth = 0
        for b in stack:
            stackWidth += b.Width
        # print()
        # print("Stack Width: {}".format(stack[0].Height))
        # print("Width of all books in stack: {}".format(stackWidth))
        # if the width of all the books in the entire stack is less than the height of the first book (i.e. the width of the stack), it makes sense to keep it
        if stackWidth >= stack[0].Height:
            for idx, b in enumerate(stack):
                books[idx+i].IsSideways = stackSidewaysInfo[idx][0]
                books[idx+i].PreviousSideways = stackSidewaysInfo[idx][1]
                books[idx+i].x = stackSidewaysInfo[idx][2]
                books[idx+i].y = stackSidewaysInfo[idx][3]
                books[idx+i].Checked = True
                # print(b.Title)
                # print("\tIsSideways: {}".format(b.IsSideways))
                # print("\tPreviousSideways: {}".format(b.PreviousSideways.Title if b.PreviousSideways else b.PreviousSideways))
                # print("\tX: {}".format(b.x))
                # print("\tY: {}".format(b.y))
            books[(len(stack)-1)+i].LastInStack = True
            # print("This was a good stack")
        else:
            books[i].LastInStack = True
            books[i].Checked = True
    else:
        books[i].LastInStack = True
        books[i].Checked = True
    # print("This was not a good stack")
    # print("Returning from stack")
    return shouldBreakShelf, shouldBreakDivider

def form_stack(books, stack, stackSidewaysInfo, previousSideways, currentX, currentY, shelfIsTop, shelfHeight, checkingIdx):
    innerX = 0
    # keep going until we either return to the previous level because a book doesn't fit or we run out of books
    while checkingIdx < len(books):
        # if this book can't be placed normally because of the width on the previous sideways, return and try on the next lower level
        if books[checkingIdx].Width + innerX > previousSideways.Height:
            # print("{} ({}) can't be placed normally because of its width. Returning to the lower level ({}) (CurrentX: {}, CurrentY: {}, InnerX: {}, BookWidth: {}, BookHeight: {}, PreviousSidewaysHeight: {}, ShelfHeight: {}). Returing to previous call.".format(books[checkingIdx].Title, checkingIdx, previousSideways.Title, currentX, currentY+innerX, innerX, books[checkingIdx].Width, books[checkingIdx].Height, previousSideways.Height, shelfHeight))
            return checkingIdx, currentX + previousSideways.Height
        # if this book can't be placed sideways on the previous sideways because it is taller, try normally
        elif books[checkingIdx].Height + innerX > previousSideways.Height:
            # print("{} ({}) can't be placed sideways because of it's height. (CurrentX: {}, CurrentY: {}, InnerX: {}, BookWidth: {}, BookHeight: {}, PreviousSidewaysHeight: {}, ShelfHeight: {}). Trying normally.".format(books[checkingIdx].Title, checkingIdx, currentX+innerX, currentY, innerX, books[checkingIdx].Width, books[checkingIdx].Height, previousSideways.Height, shelfHeight))
            # if this book can't be placed normally because of the height on the previous sideways, return and try on the lower level
            if books[checkingIdx].Height + currentY > shelfHeight:
                # print("{} ({}) can't be placed normally because of it's height. Returning to lower level ({}). (CurrentX: {}, CurrentY: {}, InnerX: {}, BookWidth: {}, BookHeight: {}, PreviousSidewaysHeight: {}, ShelfHeight: {}). Returing to previous call.".format(books[checkingIdx].Title, checkingIdx, previousSideways.Title, currentX+innerX, currentY, innerX, books[checkingIdx].Width, books[checkingIdx].Height, previousSideways.Height, shelfHeight))
                return checkingIdx, currentX + previousSideways.Height
            # if it can be placed normally, place it and advance the innerX past it
            else:
                # print("{} ({}) CAN be placed normally. (X: {}, Y: {}, Width: {}, Height: {}, PreviousSideways: {})".format(books[checkingIdx].Title, checkingIdx, currentX+innerX, currentY, books[checkingIdx].Width, books[checkingIdx].Height, previousSideways.Title))
                stack.append(books[checkingIdx])
                stackSidewaysInfo.append((False, previousSideways, currentX + innerX, currentY))
                innerX += stack[-1].Width
                checkingIdx += 1
        # if it can't be placed sideways because it will overfill the shelf, return and try on the lower level
        elif currentY + books[checkingIdx].Width > (MAX_TOP_STACK_SIZE if shelfIsTop else shelfHeight):
            # print("{} ({}) can't be placed sideways because of it's width. (CurrentX: {}, CurrentY: {}, InnerX: {}, BookWidth: {}, BookHeight: {}, PreviousSidewaysHeight: {}, ShelfHeight: {}). Trying normally.".format(books[checkingIdx].Title, checkingIdx, currentX+innerX, currentY, innerX, books[checkingIdx].Width, books[checkingIdx].Height, previousSideways.Height, shelfHeight))
            return checkingIdx, currentX + previousSideways.Height
        # if it can be placed sideways, add it to the stack and form the stack on top of it. When the stack on top of it is done, move past it
        else:
            # print("{} ({}) CAN be placed sideways. (X: {}, Y: {}, Width: {}, Height: {}, PreviousSideways: {})".format(books[checkingIdx].Title, checkingIdx, currentX+innerX, currentY, books[checkingIdx].Width, books[checkingIdx].Height, previousSideways.Title))
            stack.append(books[checkingIdx])
            stackSidewaysInfo.append((True, previousSideways, currentX + innerX, currentY))
            checkingIdx += 1
            checkingIdx, ix = form_stack(books, stack, stackSidewaysInfo, stack[-1], currentX + innerX, currentY+stack[-1].Width, shelfIsTop, shelfHeight, checkingIdx)
            innerX += ix - currentX
    return checkingIdx, currentX

def fill_cases_with_lookahead(cases, books):
    useSortHeights=False
    books = sort_books(books, useSortDeweys=False, useSortHeights=useSortHeights)
    currentCase = 0
    currentShelf = 0
    x = 0
    shelfIsTop = cases[currentCase].Shelves[currentShelf].IsTop == 1
    shelfWidth = cases[currentCase].Shelves[currentShelf].Width - cases[currentCase].Shelves[currentShelf].PaddingLeft - cases[currentCase].Shelves[currentShelf].PaddingRight
    shelfHeight = cases[currentCase].Shelves[currentShelf].Height
    oversized = []
    stay_togethers = []
    dividers = GetDividers(cases[currentCase].Shelves[currentShelf].ID)
    currentDivider = 0
    i = 0
    while i < len(books):
        # Ignore oversized, for now
        if books[i].Height >= shelfHeight:
            oversized.append(books.pop(i))
            continue
        # Don't use stay_togethers when using stacks
        # If there are any stay_together sets, try to put them in the current location, if they fit
        # If this is the first book of a stay_together set, check if it and the rest in the set fit
        # If it fits, place it normally
        # If it does not fit, check if there will be a future shelf where it will fit.
        # If so, add it as a stay_together set
        # If it will not fit in the future, place it normally.
        # Probably do this inside check_book below

        # If this book has not been checked, we check it. If it has been checked, it already has the correct parameters
        # print("About to check book: {}. Checked: {}".format(books[i].Title, books[i].Checked))
        if not books[i].Checked:
            shouldBreakShelf, shouldBreakDivider = check_book(cases, books, currentCase, currentShelf, x, shelfIsTop, shelfWidth, shelfHeight, i, dividers, currentDivider)
            if shouldBreakShelf:
                x = 0
                currentShelf += 1
                if currentShelf >= len(cases[currentCase].Shelves):
                    currentShelf = 0
                    currentCase += 1
                    if currentCase >= len(cases):
                        break
                while cases[currentCase].Shelves[currentShelf].DoNotUse:
                    currentShelf += 1
                    if currentShelf >= len(cases[currentCase].Shelves):
                        currentShelf = 0
                        currentCase += 1
                        if currentCase >= len(cases):
                            break
                shelfIsTop = cases[currentCase].Shelves[currentShelf].IsTop == 1
                shelfWidth = cases[currentCase].Shelves[currentShelf].Width - cases[currentCase].Shelves[currentShelf].PaddingLeft - cases[currentCase].Shelves[currentShelf].PaddingRight
                shelfHeight = cases[currentCase].Shelves[currentShelf].Height - STACKED_CASE_PADDING
                dividers = GetDividers(cases[currentCase].Shelves[currentShelf].ID)
                currentDivider = 0
                while oversized:
                    books.insert(i, oversized.pop(len(oversized)-1))
                continue
            if shouldBreakDivider:
                x = dividers[currentDivider]['distancefromleft'] + dividers[currentDivider]['width'] + cases[currentCase].BookMargin
                currentDivider += 1
                continue
        cases[currentCase].Shelves[currentShelf].Books.append(books[i])
        # print(books[i].LastInStack)
        # If this book is the last in the stack, find the first in the stack (the longest height, or width if it is a solo book (normally placed)) and advance x by that much
        if books[i].LastInStack:
            checking = books[i]
            while checking.PreviousSideways:
                checking = checking.PreviousSideways
            x += checking.Height if checking.IsSideways else checking.Width
        i += 1
    oversizedWidth = 20
    x = 0
    for o in oversized:
        o.x = x
        o.y = 0
        x += o.Width
        oversizedWidth += o.Width
    cases.append(Case(
        -1,
        cases[0].SpacerHeight,
        cases[0].BookMargin,
        oversizedWidth + 20,
        [Shelf(
            -1,
            -1,
            0,
            oversizedWidth + 20,
            300,
            cases[0].Shelves[0].PaddingLeft,
            cases[0].Shelves[0].PaddingRight,
            "Left",
            0,
            0,
            oversized
        )],
        cases[0].AverageBookHeight,
        cases[0].AverageBookWidth,
        len(cases) + 1
        ))


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
    caseCanvas = svgwrite.Drawing(os.path.join(CaseSvgPath, '{}.svg').format(c.ID), size=(caseWidth, maxCaseHeight), viewBox=('0 0 {} {}'.format(caseWidth, maxCaseHeight)))
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
        text_x = x + b.x+b.Width/2
        text_y = y - b.y - 4
        rotation = -90
        caseCanvas.add(caseCanvas.text(title, (text_x, text_y), class_='bookcase-book-text', transform='rotate({},{},{})'.format(rotation, text_x, text_y), dominant_baseline="middle", font_family="Arial", font_size="{}px".format(fontSize), fill=fontColor, id="PATH{}".format(b.ID)))
    else:
        text_x = b.x
        text_y = y - b.y - b.Width/2
        caseCanvas.add(caseCanvas.text(title, (x + text_x, text_y), class_='bookcase-book-text', dominant_baseline="middle", font_family="Arial", font_size="{}px".format(fontSize), fill=fontColor))
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
    if b.IsSideways:
        caseCanvas.add(caseCanvas.rect((x + b.x, y - b.y - b.Width), (b.Height, b.Width), id='book-{}'.format(b.ID), class_='bookcase-book', stroke='black', fill='{}'.format(b.SpineColor)))
    else:
        caseCanvas.add(caseCanvas.rect((x + b.x, y - b.y - b.Height), (b.Width, b.Height), id='book-{}'.format(b.ID), class_='bookcase-book', stroke='black', fill='{}'.format(b.SpineColor)))

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
    # x = get_next_x(c, s, b, x, bidx)
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

def draw_cases(cases, directory):
    CaseSvgPath = directory

    maxCaseHeight = get_max_case_height(cases)

    for c in cases:
        draw_case(CaseSvgPath, c, maxCaseHeight)

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('--directory', default="libraries/testsvgs")
    parser.add_argument('--session', default="MTY4NTA0MzY5OXxEdi1CQkFFQ180SUFBUkFCRUFBQVlfLUNBQUVHYzNSeWFXNW5EQmtBRjJ4cFluSmhjbmx2Y21kaGJtbDZaWEp6WlhOemFXOXVCbk4wY21sdVp3dzBBREpCUVVWR1VtaHZPRkYwT0ZsR09ESTBXSGd4YWtoc09HZDZNSFkzTkhaTVpGcHNNRU42TUUxeVUyMUxWVTB5YkZGcVl3PT18IXzq-2ZDiEF4oLJF-yheob-OBHZpDD0gHsW8hKIQXyg=")
    parser.add_argument('--libraryid', default="43")
    args = parser.parse_args()
    directory = args.directory

    RELOAD = True

    SORT_METHOD = "Dewey:ASC--Author:ASC--SeriesHeight:ASC--Series:ASC--Volume:ASC--Title:ASC--Subtitle:ASC--Edition:ASC--Lexile:ASC--InterestLevel:ASC--AR:ASC--LearningAZ:ASC--GuidedReading:ASC--DRA:ASC--FountasPinnell:ASC--ReadingRecovery:ASC--PMReaders:ASC--Grade:ASC--Age:ASC--Publisher:ASC--LibraryOfCongress:ASC--FictionGenre:ASC||Dewey:ASC--SeriesHeight:ASC--Series:ASC--Volume:ASC--Author:ASC--Title:ASC--Subtitle:ASC--Edition:ASC--Lexile:ASC--InterestLevel:ASC--AR:ASC--LearningAZ:ASC--GuidedReading:ASC--DRA:ASC--FountasPinnell:ASC--ReadingRecovery:ASC--PMReaders:ASC--Grade:ASC--Age:ASC--Publisher:ASC--LibraryOfCongress:ASC--FictionGenre:ASC"

    LETTER_WIDTH = dict([[" ",27.797],["!",27.797],["\"",35.5],["#",55.625],["$",55.625],["%",88.922],["&",66.703],["'",19.094],["(",33.313],[")",33.313],["*",38.922],["+",58.406],[",",27.797],["-",33.313],[".",27.797],["/",27.797],["0",55.625],["1",55.625],["2",55.625],["3",55.625],["4",55.625],["5",55.625],["6",55.625],["7",55.625],["8",55.625],["9",55.625],[":",27.797],[";",27.797],["<",58.406],["=",58.406],[">",58.406],["?",55.625],["@",101.516],["A",66.703],["B",66.703],["C",72.219],["D",72.219],["E",66.703],["F",61.094],["G",77.797],["H",72.219],["I",27.797],["J",50],["K",66.703],["L",55.625],["M",83.313],["N",72.219],["O",77.797],["P",66.703],["Q",77.797],["R",72.219],["S",66.703],["T",61.094],["U",72.219],["V",66.703],["W",94.391],["X",66.703],["Y",66.703],["Z",61.094],["[",27.797],["\\",27.797],["]",27.797],["^",46.938],["_",55.625],["`",33.313],["a",55.625],["b",55.625],["c",50],["d",55.625],["e",55.625],["f",27.797],["g",55.625],["h",55.625],["i",22.219],["j",22.219],["k",50],["l",22.219],["m",83.313],["n",55.625],["o",55.625],["p",55.625],["q",55.625],["r",33.313],["s",50],["t",27.797],["u",55.625],["v",50],["w",72.219],["x",50],["y",50],["z",50],["{",33.406],["|",25.984],["}",33.406],["~",58.406],["_median",55.625]])
    KERNING_MODIFIER = dict([[" A",-5.531],[" T",-1.829],[" Y",-1.813],["11",-7.438],["A ",-5.531],["AT",-7.422],["AV",-7.422],["AW",-3.719],["AY",-7.422],["Av",-1.797],["Aw",-1.813],["Ay",-1.797],["F,",-11.094],["F.",-11.094],["FA",-5.531],["L ",-3.734],["LT",-7.438],["LV",-7.422],["LW",-7.438],["LY",-7.422],["Ly",-3.719],["P ",-1.813],["P,",-12.906],["P.",-12.906],["PA",-7.422],["RT",-1.813],["RV",-1.813],["RW",-1.813],["RY",-1.813],["T ",-1.829],["T,",-11.094],["T-",-5.532],["T.",-11.094],["T:",-11.094],["T;",-11.094],["TA",-7.422],["TO",-1.828],["Ta",-11.094],["Tc",-11.094],["Te",-11.094],["Ti",-3.719],["To",-11.094],["Tr",-3.72],["Ts",-11.094],["Tu",-3.719],["Tw",-5.516],["Ty",-5.516],["V,",-9.188],["V-",-5.532],["V.",-9.188],["V:",-3.719],["V;",-3.719],["VA",-7.422],["Va",-7.422],["Ve",-5.531],["Vi",-1.813],["Vo",-5.531],["Vr",-3.719],["Vu",-3.719],["Vy",-3.703],["W,",-5.532],["W-",-1.813],["W.",-5.532],["W:",-1.813],["W;",-1.813],["WA",-3.719],["Wa",-3.719],["We",-1.813],["Wo",-1.813],["Wr",-1.813],["Wu",-1.813],["Wy",-0.875],["Y ",-1.813],["Y,",-12.906],["Y-",-9.188],["Y.",-12.906],["Y:",-5.531],["Y;",-6.5],["YA",-7.422],["Ya",-7.422],["Ye",-9.187],["Yi",-3.703],["Yo",-9.187],["Yp",-7.422],["Yq",-9.187],["Yu",-5.531],["Yv",-5.516],["ff",-1.828],["r,",-5.532],["r.",-5.532],["v,",-7.422],["v.",-7.422],["w,",-5.532],["w.",-5.532],["y,",-7.422],["y.",-7.422]])

    BASE_URL = "http://library.jakekausler.com"
    CASE_URL = BASE_URL + "/libraries/{}/cases".format(args.libraryid)
    BOOK_URL = BASE_URL + "/books/bookscases?libraryid={}".format(args.libraryid)
    DIVIDERS_URL = BASE_URL + "/libraries/shelfdividers/{}"
    COOKIES = {'session': args.session}
    CASE_MARGIN = 50
    STACKED_CASE_PADDING = 15
    MAX_TOP_STACK_SIZE = 200

    case_data = None
    book_data = None

    print("Getting Data")
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

    cases = form_cases(case_data)
    books = form_books(book_data)
    print("Filling Cases")
    fill_cases_with_lookahead(cases, books)
    print("Drawing Cases")
    draw_cases(cases, directory)
    print("Done Creating Cases")
