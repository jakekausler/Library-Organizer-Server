# Library-Organizer-2.0

Library Organizer is a complete suite of tools designed to help users manage their personal or professional libraries. It hosts a variety of features:

* Create and assign books to one or more libraries
* Grant other users permission to edit, check out from, or view your libraries
* View a filtered list of books from viewable libraries
* Add or edit books from scratch, from existing books, or from isbn
* View a physical model of each library based on bookcase and book dimensions
* Organize books based on title, series, or dewey number
* View various statistics that describe one or more libraries
* Keep track of non-owned books as a wishlist

In this document, you will find both a user guide, for normal users to learn the software, and a REST API, for developers to utilize.

* User Guide
* REST API

# User Guide

* Registering and Logging In
  * Registering an Account
  * Logging In to an Account
  * Forgotten Password
* Managing Libraries and Permissions
  * Managing Libraries
  * Granting Permissions
* Managing User Settings
  * Shelf Settings
  * Book Settings
* Searching for Books
* Adding and Editing Books
  * Add from Scratch
  * Add from ISBN
  * Add from Existing Book
  * Edit an Existing Book
  * Bulk Editing
* Checking Out and Returning Books
  * Checking out a Book
  * Returning a Book
* Exporting and Importing CSV Files
  * The CSV Format
  * Exporting
  * Importing
* Viewing a Library
  * Navigating the Library
  * Searching for Books
* Editing a Library
  * Managing Cases
  * Managing Breaks
  * Managing Library Details
* Viewing Statistics
* Value Types

## Registering and Logging In

When you first load the Library Organizer site, you will be brought to a page where you can either login to your account or register for an account. Click the appropriate button to continue to the site.

* Registering an Account
* Logging In to an Account
* Forgotten Password

### Registering an Account

If you do not have an account, you will need to register for an account. Click the register button and enter your desired username, email address, first and last name, and password. If the information provided is valid, your account will be created and you will be redirected to the main site. An empty default library is automatically created for you.

### Logging In to an Account

If you do have an account, click the log in button and enter your username and password. You will be redirected to the main site.

### Forgotten Password

If you have forgotten your password, you may click the log in button, then the forgot password button. Enter the email you used to create your account. An email with instructions on how to reset your password will be sent to you.

## Managing Libraries and Permissions

At the core of Library Organizer are libraries. Libraries hold books and shelves on which those books can be displayed. In the settings dialog (accessed by pressing the gear button on the top right of the screen) there are various ways you may manipulate your libraries.

NOTE: You must click the save button at the bottom of the settings dialog in order for your library changes to be saved.
NOTE: If you delete a library, all of its books will be deleted as well.

* Managing Libraries
* Granting Permissions

### Managing Libraries

A user may have as many libraries as he or she desires, but must have at least one. In the settings dialog, you may create, rename, and delete your libraries. To create a library, click the plus button at the bottom of the dialog. To rename a library, simply change the library name in the text input near the top of the dialog. To delete a library, click the trash can beside the library's name. Remember that if you delete a library, all of its books will be deleted as well.

### Granting Permissions

You may grant permission for other users to view, check out from, or edit your libraries. To do so, use the appropriate search box to search for a user, and click on them to add them to the corresponding permission group. You may search by name, username, or email address.

## Managing User Settings

User settings are default settings used when creating shelves. These are found in the settings dialog (accessed by pressing the gear button on the top right of the screen), in the second and third tabs, respectively.

* Shelf Settings
* Book Settings

### Shelf Settings

The possible shelf settings, with their value types, are as follows. See the Value Types section of this guide for more information.
Setting | Value Type
--- | ---
TODO | TODO

### Book Settings

The possible book settings, with their value types, are as follows. See the Value Types section of this guide for more information.
Setting | Value Type
--- | ---
TODO | TODO

## Searching for Books

You can open the filters menu by clicking the filter button located in the top left of the screen on the grid view. This will allow you to change which books displayed in the grid. There are several options:
* Sort Method - How the books should be primarily sorted. One of "Title", "Dewey", "Lexile", "Series", or "Author"
* Limit - Number of books to get at one time.
* Page - Page of books to get. For instance, if limit is 10 and page is 2, this will return books 11-20.
* From Dewey/To Dewey - The minimum and maximum dewey decimal values to return, inclusive. Note that FIC is the largest value, after 999.999...
* From Lexile/To Lexile - The minimum and maximum lexile numbers to return, inclusive. These can be in BR or regular format. See the Value Types section for more information.
* Search for - Search terms to find. Any book where each word entered here is found within its title, subtitle, or series will be returned.
* Read - Whether or not the book has been read by the library owner.
* Reference - Whether or not the book is considered reference by the library owner.
* Owned - Whether or not the book is owned by the library owner. If not owned, it a book is considered to be on a wish list. A book that is not owned cannot be checked out.
* Loaned - Whether or not the book is checked out from the library.
* Shipping - Whether or not the book is shipping to the library. A book being shipped cannot be checked out.
* Reading - Whether or not the book is being read by the library owner. A book being read cannot be checked out.
* Libraries - Which libraries from which to return results.

## Adding and Editing Books

You may add books to or edit books in any library you own or have permission to edit. There are various ways to add or edit books.

* The Book Editor
* Add from Scratch
* Add from ISBN
* Add from Existing Book
* Edit an Existing Book
* Bulk Editing

### The Book Editor

The book editor is a dialog that contains input fields for a book. The main fields and their values are as follows. See the Value Types section of this guide for more information.

Field | Value Type | Tab | Description
--- | --- | --- | ---
Title | String | 1 | The title of the book. Required
Subtitle | String | 1 | The subtitle of the book.
Series | String | 1 | The series the book is a part of. Suggestions will appear with possible series as you type.
Volume | Number | 1 | The series volume of the book. Required. Zero is the empty volume.
Publisher | String | 2 | The publisher of the book. Suggestions will appear with possible publishers as you type.
City | String | 2 | The city where the book was published. Suggestions will appear with possible cities as you type.
State | String | 2 | The state where the book was published. Suggestions will appear with possible states as you type.
Country | String | 2 | The country where the book was published. Suggestions will appear with possible countries as you type.
Originally Published | Year | 2 | The year in which the book was originally published. Required. 0000 is the empty year.
Edition Published | Year | 2 | The year in which this edition of the book was published. Required. 0000 is the empty year.
ISBN | ISBN | 2 | The ISBN of the book.
Edition | Non-Negative Integer | 2 | The edition of the book. Required. Zero is the empty edition.
Lexile | Lexile | 2 | The Lexile of the book. Required. 0L is the empty lexile.
Dewey | Dewey | 2 | The dewey of the book. Required. 000 is the empty dewey.
Format | String | 2 | The binding of the book.
Pages | Non-Negative Integer | 2 | The number of pages in the book. Required. Zero is the empty number of pages.
Width | Non-Negative Integer | 2 | The width of the book in millimeters. Required. Zero is the empty width.
Height | Non-Negative Integer | 2 | The height of the book in millimeters. Required. Zero is the empty height.
Depth | Non-Negative Integer | 2 | The depth of the book in millimeters. Required. Zero is the empty depth.
Weight | Non-Negative Number | 2 | The weight of the book in ounces. Required. Zero is the empty weight.
Primary Language | String | 2 | The main language in which the book is written.
Secondary Language | String | 2 | The secondary language in which the book is written.
Original Language | String | 2 | The original language in which the book is written.
Owned | Checkbox | 2 | Whether or not the book is owned. If not owned, it is on the wish list.
Read | Checkbox | 2 | Whether or not the book has been read by the library owner.
Reference | Checkbox | 2 | Whether or not the book is considered reference by the library owner.
Reading | Checkbox | 2 | Whether or not the book is being read by the library owner.
Shipping | Checkbox | 2 | Whether or not the book is being shipped.

There is also a section on the first tab for adding and editing the contributors for a book. A contributor is added by filling in its fields and pressing the plus button. To edit an existing contributor, press the pencil button and fill in its fields, then click the save button. To remove an existing contributor, press the trash button.

Field | Value Type | Description
--- | --- | ---
First Name | String | The first name of the contributor
Middle Names | String | The middle names of the contributor, separated by spaces
Last Name | String | The last name of the contributor
Role | String | The role the contributor had for the book, for instance "Author" or "Editor"

Finally, there is a section on the first tab for adding, changing, editing, and removing a book's image and spine color. An image may be added or changed in two ways: 1) By pressing the paste button and typing in an image url, or 2) By pressing the upload button and uploading an image. To remove an image, press the trash button. An image may also be cropped or rotated by pressing the pencil button.

When a book is saved, its spine color is automatically set based on the most prominant color in the image. To override this, you may click the spine color button and choose a new color. To go back to the automatically calculated color, click the default button in the color chooser.

### Add from Scratch

To add a book from scratch, click the plus button at the top of the grid view. Any field which has a user setting will be automatically filled in.

### Add from ISBN

To add a book from isbn, click the barcode button at the top of the grid view. Enter the isbn number and press enter. A list of possible matches will appear. Click on the one you would like to add and a book editor will appear with information gathered from that isbn.

### Add from Existing Book

To add a book from an existing book, click on the book you would like to duplicate in either the library or grid view, then click on the copy button in the book view dialog. An editor menu will appear with the book's information.

### Edit an Existing Book

To edit an existing book, click on the book you would like to edit and then click the pencil button. An editor menu will appear with the book's information.

### Bulk Editing

TODO

## Checking Out and Returning Books

If you have permission to check out from a library, you may check out books (and also return them). Any book you have checked out will display as checked out to you and will not be available for others to check out.

* Checking out a Book
* Returning a Book

### Checking Out a Book

To check out a book, click on the book you would like to check out from either the grid or library view. Then click the cart button. Upon confirming that you would like to check out the book, it will now be listed as a book you have checked out.

### Returing a Book

To check out a book, click on the book you would like to check in from either the grid or library view. Then click the cart button. Upon confirming that you would like to check in the book, it will no longer be listed as a book you have checked out.

## Exporting and Importing CSV Files

You may export or import books in a csv format. There are various options for exporting books and special considerations that must be made when importing books.

* The CSV Format
* Exporting
* Importing

### The CSV Format
TODO

### Exporting
TODO

### Importing
TODO

## Viewing a Library

The library view (accessed by using the second button at the top of the screen) shows a physical representation of a library. It is based on case and book measurements.

* Navigating the Library
* Searching for Books

### Navigating the Library

To change the current library, click the filter button in the top left of the screen. To move around in the library, use the scroll bars at the right and bottom of the screen.

### Searching for Books

TODO

## Editing a Library

There are three types of information that can be edited for a library: The cases, breaks, and details.

* Managing Cases
* Managing Breaks
* Managing Library Details

### Managing Cases

A library is made of book cases. Each book case has the following fields that may be edited. See the Value Types section for more details.

Field | Value Type | Description
--- | --- | ---
Case Width | Positive Integer | This is the width of the shelf from end to end, not including the sides (so, from inside edge to inside edge)
Number of Shelves | Positive Integer | The number of shelves on this case. This is the number of locations that books may be placed, so it should include the bottom of the case.
Padding Left/Right | Positive Integer | This is the minimum amount of space that should be left on either side of a shelf to make up for measurements that were too small. The more accurate that you are in your measurements, the less padding that is needed.
Shelf Height | Positive Integer | The height, from the top of one shelf to the bottom of the next, of a shelf gap. This does not include the actual shelf height.
Spacer Size | Positive Integer | The thickness of the sides/shelves of a case. This is how thick the "outline" of the case and its shelves will be.

Cases may also be moved up or down by using the respective buttons, or deleted using the trash button of the corresponding case. To add a new shelf, click the plus button at the bottom of the dialog.

### Managing Breaks

There are two types of breaks that a library may include:
* Shelf Breaks - After a shelf break, books stop being put on the shelf the corresponding book is on and start again on the next shelf.
* Case Breaks - After a case break, books stop being put on the case the corresponding book is on and start again on the next case.

In addition to these types of breaks, breaks can be of a specific value type:
* Book - The break happens after a specific book
* Dewey - The break happens after the last book of a specific dewey
* Lexile - The break happens after the last book of a specific lexile
* Series - The break happens after the last book of a specific series
* Author - The break happens after the last book in each string of books by a specific author
* Title - The break happens after the last book whose title begins with a specific set of letters

To add a break, press the plus button at the bottom of the dialog. To remove a break, press the trash can on the corresponding break.

### Managing Details

There are two other features of a library that may be edited:

* Series - Choose series that have multiple authors, that should be sorted on their volumes first rather than their authors. For instance, normally if two books in a series have different authors, they will be sorted on their authors. If a series is on this list, however, it will be sorted by its volume instead.
* Sort Method - How this library should be primarily sorted. One of "Title", "Dewey", "Lexile", "Series", or "Author"

## Viewing Statistics

The statistics menu can be accessed by clicking the third button on the top of the screen. To change which libraries are being used to calculate the statistics, click the filter button in the top left of the screen. Only owned books are used in the calculations. There are several types of statistics, grouped into various categories:

* General - Information on books
  * Counts - Number of books that are read, reference, shipping, loaned, or being read
  * Size - Volume that is read, reference, shipping, loaned, or being read
  * Pages - Number of pages that are read, reference, shipping, loaned, or being read
* Publishers - Information on publishers and publication locations
  * Parent Company - Number of books published by each parent company
  * Top Publishers - Number of books published by the top publishers
  * Top Locations - Number of books published in the top locations
* Series - Number of books in each series
* Language - Information on languages
  * Primary Languages - Number of books written primarily in each language
  * Secondary Languages - Number of books written secondarily in each language
  * Original Languages - Number of books written originally in each language
* Deweys - Number of books in each overall dewey genre
* Formats - Number of books with each binding type
* Contributors - Information on contributors and roles
  * Top Contributors - Number of books written by the top contributors
  * Roles - Number of contributors with each role
* Dates - Information on publication dates
  * Original Publication Date - Number of books originally published in each decade
  * Edition Publication Date - Number of books whose edition was published in each decade
* Dimensions - Information on book dimensions
  * Volume - Total volume of the books
  * Width - Total, average, minimum, and maximum widths of the books
  * Height - Total, average, minimum, and maximum heights of the books
  * Depth - Total, average, minimum, and maximum depths of the books
  * Weight - Total, average, minimum, and maximum weights of the books
  * Pages - Total, average, minimum, and maximum pages of the books
* Lexiles - Number of books in each lexile grade level

## Value Types

There are several fields that require a specific value type. The different value types are described here:

* String - A generic text field
* Number - A positive or negative number, or zero
* Positive Integer - The field must be an integer that is greater than zero
* Non-Negative Integer - The field must be an integer that is at least zero
* Positive Number - The field must be an number that is greater than zero
* Non-Negative Number - The field must be an number that is at least zero
* ISBN - The field must be a valid isbn-10 or -13
* Dewey - The field must be a valid dewey decimal number (of the form "#[.#]" where the number on the left has three digits and the optional number after the decimal can contain any number greater than one of digits) or "FIC"
* Lexile - The field must be a valid lexile number (of the form "[BR]#L", where BR is optional and the number can contain any number of digits). Note that if a lexile begins with BR is means it is negative, so the BR100L is the integer equivalent of -100.
* Year - The field must be a four digit year

# REST API

* Books `/books`
  * Get Books `GET /books`
  * Save Book `PUT /books`
  * Add Book `POST /books`
  * Delete Book `DELETE /books`
  * Checkout Book `PUT /books/checkout`
  * Checkin Book `PUT /books/checkin`
  * Export Books `GET /books/books`
  * Export contributors `GET /books/contributors`
  * Import Books `POST /books/books`
  * Get Ratings `GET /books/:bookid/ratings`
  * Get Reviews `GET /books/:bookid/reviews`
  * Add Rating `POST /books/:bookid/ratings`
  * Add Review `POST /books/:bookid/reviews`

* Information `/information`
  * Get Statistics `GET /information/statistics`
  * Get Dimensions `GET /information/dimensions`
  * Get Publishers `GET /information/publishers`
  * Get Cities `GET /information/cities`
  * Get States `GET /information/states`
  * Get Countries `GET /information/countries`
  * Get Formats `GET /information/formats`
  * Get Roles `GET /information/roles`
  * Get Series `GET /information/series`
  * Get Languages `GET /information/languages`
  * Get Deweys `GET /information/deweys`

* Libraries `/libraries`
  * Get Libraries `GET /libaries`
  * Get Owned Libraries `GET /libraries/owned`
  * Save Owned Libraries `PUT /libraries/owned`
  * Get Cases `GET /libraries/:libraryid/cases`
  * Save Cases `PUT /libraries/:libraryid/cases`
  * Get Breaks `GET /libraries/:libraryid/breaks`
  * Update Breaks `PUT /libraries/:libraryid/breaks`
  * Get Author Based Series `GET /libraries/:libraryid/series`
  * Update Author Based Series `PUT /libraries/:libraryid/series`
  * Get Sort Method `GET /libraries/:libraryid/sort`
  * Update Sort Method `PUT /libraries/:libraryid/sort`

* Settings `/settings`
  * Get Settings `GET /settings`
  * Save Settings `PUT /settings`
  * Get Setting `GET /settings/:setting`

* Users `/users`
  * Get Users `GET /users`
  * Login `POST /users/login`
  * Register `POST /users`
  * Logout `POST /users/logout`
  * Send Password Reset `PUT /users/password`
  * Finish Password Reset `GET /users/password/:token`
  * Get Username `GET /users/username`

## Books
* Get Books `GET /books`
* Save Book `PUT /books`
* Add Book `POST /books`
* Delete Book `DELETE /books`
* Checkout Book `PUT /books/checkout`
* Checkin Book `PUT /books/checkin`
* Export Books `GET /books/books`
* Export Contributors `GET /books/contributors`
* Import Books `POST /books/books`
* Get Ratings `GET /books/:bookid/ratings`
* Get Reviews `GET /books/:bookid/reviews`
* Add Rating `POST /books/:bookid/ratings`
* Add Review `POST /books/:bookid/reviews`

### Get Books

* **Action**
`Returns an array of books.`

* **Method**
`GET`

* **URL**
`/books`

* **URL Params**
`none`

* **URL Arguments (?)**
  * `sortmethod`: Primary method of sorting. One of `dewey`, `series`, `lexile`, `author`, or `title`.
  * `numbertoget`: Number of books for pagination purposes.
  * `page`: Page of books to get.
  * `fromdewey`: Starting dewey of the results, inclusive.
  * `todewey`: Ending dewey of the results, inclusive
  * `text`: Filter text, used for searching titles, subtitles, and series
  * `isread`: Whether or not the book is read by the library owner. One of `yes`, `no`, or `both`.
  * `isreference`: Whether or not the book is considered reference by the library owner. One of `yes`, `no`, or `both`.
  * `isowned`: Whether or not the book is owned by the library owner. One of `yes`, `no`, or `both`.
  * `isreading`: Whether or not the book is being read by the library owner. One of `yes`, `no`, or `both`.
  * `isshipping`: Whether or not the book is shipping to the library owner. One of `yes`, `no`, or `both`.
  * `isloaned`: Whether or not the book is checked out. One of `yes`, `no`, or `both`.
  * `libraryids`: A comma separated list of library ids from which to include books.

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Save Book

* **Action**
`Saves a book.`

* **Method**
`PUT`

* **URL**
`/books`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
```json
{
    "bookid": "12345",
    "title": "Sample Title",
    "subtitle": "Sample Subtitle",
    "originallypublished": "2006-01-02T15:04:05Z07:00",
    "publisher": {
       "id": "12345",
       "publisher": "Sample Publishing House",
       "city": "New York",
       "state": "New York",
       "country": "USA",
       "parentcompany": "Sample Parent Publishing House"
    },
    "isread": true,
    "isreference": true,
    "isowned": true,
    "isbn": "123456789012X",
    "loanee": {
        "id": 123,
        "username": "testuser",
        "first": "Jane",
        "last": "Smith",
        "fullname": "Jane Smith",
        "email": "jsmith@test.com",
        "iconurl": "/res/userimages/123.jpg"
    },
    "dewey": "FIC",
    "pages": 0,
    "width": 0,
    "height": 0,
    "depth": 0,
    "weight": 0.0,
    "primarylanguage": "English",
    "secondarylanguage": "Klingon",
    "originallanguage": "English",
    "series": "Sample Series",
    "volume": 0.0,
    "format": "Paperback",
    "edition": 1,
    "isreading": true,
    "isshipping": true,
    "imageurl": "/res/bookimages/12345.jpg",
    "spinecolor": "##123456",
    "cheapestnew": 0.01,
    "cheapestused": 0.01,
    "editionpublished": "2006-01-02T15:04:05Z07:00",
    "contributors": [{
        "id": "1234",
        "name": {
            "first": "John",
            "middles": "J;Kimily",
            "last": "Smith"
        },
        "role": "Author"
    }],
    "library": {
        "id": 12345,
        "name": "default",
        "permissions": 7,
        "owner": "123"
    }
}
```

* **Success Response**
```json

```

* **Error Response**
```json

```

### Add Book

* **Action**
`Adds a book.`

* **Method**
`POST`

* **URL**
`/books`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
```json
{
    "bookid": "12345",
    "title": "Sample Title",
    "subtitle": "Sample Subtitle",
    "originallypublished": "2006-01-02T15:04:05Z07:00",
    "publisher": {
       "id": "12345",
       "publisher": "Sample Publishing House",
       "city": "New York",
       "state": "New York",
       "country": "USA",
       "parentcompany": "Sample Parent Publishing House"
    },
    "isread": true,
    "isreference": true,
    "isowned": true,
    "isbn": "123456789012X",
    "loanee": {
        "id": 123,
        "username": "testuser",
        "first": "Jane",
        "last": "Smith",
        "fullname": "Jane Smith",
        "email": "jsmith@test.com",
        "iconurl": "/res/userimages/123.jpg"
    },
    "dewey": "FIC",
    "pages": 0,
    "width": 0,
    "height": 0,
    "depth": 0,
    "weight": 0.0,
    "primarylanguage": "English",
    "secondarylanguage": "Klingon",
    "originallanguage": "English",
    "series": "Sample Series",
    "volume": 0.0,
    "format": "Paperback",
    "edition": 1,
    "isreading": true,
    "isshipping": true,
    "imageurl": "/res/bookimages/12345.jpg",
    "spinecolor": "##123456",
    "cheapestnew": 0.01,
    "cheapestused": 0.01,
    "editionpublished": "2006-01-02T15:04:05Z07:00",
    "contributors": [{
        "id": "1234",
        "name": {
            "first": "John",
            "middles": "J;Kimily",
            "last": "Smith"
        },
        "role": "Author"
    }],
    "library": {
        "id": 12345,
        "name": "default",
        "permissions": 7,
        "owner": "123"
    }
}
```

* **Success Response**
```json

```

* **Error Response**
```json

```

### Delete Book

* **Action**
`Deletes a book.`

* **Method**
`DELETE`

* **URL**
`/books`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
Book Id:
`12345`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Checkout Book

* **Action**
`Checks out a book.`

* **Method**
`PUT`

* **URL**
`/books/checkout`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
Book Id:
`12345`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Checkin Book

* **Action**
`Checks in a book.`

* **Method**
`PUT`

* **URL**
`/books/checkin`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
Book Id:
`12345`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Export Books

* **Action**
`Exports a csv of book data.`

* **Method**
`GET`

* **URL**
`/books/books`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Export Contributors

* **Action**
`Exports a csv of contributors`

* **Method**
`GET`

* **URL**
`/books/contributors`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Import Books

* **Action**
`Imports books from a csv to a library. The column headers should match the database fields names, with the exception of publisher and contributors. Publisher should be split into four columns: "publisher", "city", "state", "country". Contributors should be in a single column, with the header "contributors". They should be written as follows: "lastName[, firstName middleName1 middleName2... middleNameN]: role", with different contributors separated by a single semicolon and omitting first and middle names as needed.`

* **Method**
`POST`

* **URL**
`/books/books`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
Form Data:
`csv file`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Ratings

* **Action**
`Get the ratings for a book and its best guest matches.`

* **Method**
`GET`

* **URL**
`/books/:bookid/ratings`

* **URL Params**
`:bookid` Id of the book to get ratings from

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Add Rating

* **Action**
`Add or replace a rating for a book.`

* **Method**
`PUT`

* **URL**
`/books/:bookid/ratings`

* **URL Params**
`:bookid` Id of the book to get ratings from

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Reviews

* **Action**
`Get the reviews for a book and its best guest matches.`

* **Method**
`GET`

* **URL**
`/books/:bookid/reviews`

* **URL Params**
`:bookid` Id of the book to get reviews from

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Add Review

* **Action**
`Add or replace a review for a book.`

* **Method**
`PUT`

* **URL**
`/books/:bookid/reviews`

* **URL Params**
`:bookid` Id of the book to get reviews from

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

## Information
* Get Statistics `GET /information/statistics`
* Get Dimensions `GET /information/dimensions`
* Get Publishers `GET /information/publishers`
* Get Cities `GET /information/cities`
* Get States `GET /information/states`
* Get Countries `GET /information/countries`
* Get Formats `GET /information/formats`
* Get Roles `GET /information/roles`
* Get Series `GET /information/series`
* Get Languages `GET /information/languages`
* Get Deweys `GET /information/deweys`

### Get Statistics

* **Action**
`Get a set of statistics.`

* **Method**
`GET`

* **URL**
`/information/statistics`

* **URL Params**
`none`

* **URL Arguments (?)**
  * `type`: One of `generalbycounts`, `generalbypages`, `generalbysize`, `publishersbooksperparent`, `publisherstopchildren`, `publisherstoplocations`, `series`, `languagesprimary`, `languagessecondary`, `languagesoriginal`, `deweys`, `formats`, `contributorstop`, `contributorsperrole`, `datesoriginal`, `datespublication`
  * `libraryids`: A comma separated list of library ids from which the statistics are calculated

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Dimensions

* **Action**
`Get a set of dimensions.`

* **Method**
`GET`

* **URL**
`/information/dimensions`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
Library Ids:
`123,124,125`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Publishers

* **Action**
`Get a the set of publishers.`

* **Method**
`GET`

* **URL**
`/information/publishers`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Cities

* **Action**
`Get the set of cities.`

* **Method**
`GET`

* **URL**
`/information/cities`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get States

* **Action**
`Get the set of states.`

* **Method**
`GET`

* **URL**
`/information/states`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Countries

* **Action**
`Get the set of countries. `

* **Method**
`GET`

* **URL**
`/information/countries`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Formats

* **Action**
`Get the set of formats.`

* **Method**
`GET`

* **URL**
`/information/formats`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Roles

* **Action**
`Get the set of contributor roles.`

* **Method**
`GET`

* **URL**
`/information/roles`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Series

* **Action**
`Get the set of series.`

* **Method**
`GET`

* **URL**
`/information/series`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Languages

* **Action**
`Get the set of languages`

* **Method**
`GET`

* **URL**
`/information/languages`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Deweys

* **Action**
`Get the set of deweys`

* **Method**
`GET`

* **URL**
`/information/deweys`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

## Libraries
* Get Libraries `GET /libaries`
* Get Owned Libraries `GET /libraries/owned`
* Save Owned Libraries `PUT /libraries/owned`
* Get Cases `GET /libraries/:libraryid/cases`
* Save Cases `PUT /libraries/:libraryid/cases`
* Get Breaks `GET /libraries/:libraryid/breaks`
* Add Break `POST /libraries/:libraryid/breaks`
* Update Breaks `PUT /libraries/:libraryid/breaks`
* Delete Break `DELETE /libraries/:libraryid/breaks`
* Get Author Based Series `GET /libraries/:libraryid/series`
* Update Author Based Series `PUT /libraries/:libraryid/series`
* Get Sort Method `GET /libraries/:libraryid/sort`
* Update Sort Method `PUT /libraries/:libraryid/sort`

### Get Libraries

* **Action**
`Get a list of libraries with which the user has permission to do something.`

* **Method**
`GET`

* **URL**
`/libraries`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Owned Libraries

* **Action**
`Get the libraries which the user owns and the user who can do something with them.`

* **Method**
`GET`

* **URL**
`/libraries/owned`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Save Owned Libraries

* **Action**
`Saves the libraries the user owns and the users who can do something with them.`

* **Method**
`PUT`

* **URL**
`/libraries/owned`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
```json
[{
    "id": 123,
    "name": "default",
    "user": [{
        "id": 1234,
        "username": "janedoe",
        "firstname": "Jane",
        "lastname": "Doe",
        "fullname": "Jane Doe",
        "email": "jdoe@test.com",
        "iconurl": "/res/usericons/1234.jpg",
        "permission": "7"
    }]
}]
```

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Cases

* **Action**
`Get the cases in a library.`

* **Method**
`GET`

* **URL**
`/libraries/:libraryid/cases`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
  * `sortmethod` One of `DEWEY`, `SERIES`, or `TITLE`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Save Cases

* **Action**
`Save the cases in a library.`

* **Method**
`PUT`

* **URL**
`/libraries/:libraryid/cases`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
```json
{
    "editedcases": [{
        "id": 1233,
        "casenumber": 1,
        "numberofshelves": 5,
        "shelfheight": 300,
        "width": 500,
        "paddingright": 10,
        "paddingleft": 10,
        "spacerheight": 12
    }],
    "toremoveids": [
        1234,
        1235
    ]
}
```

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Breaks

* **Action**
`Get the breaks for a library.`

* **Method**
`GET`

* **URL**
`/libraries/:libraryid/breaks`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Update Breaks

* **Action**
`Update the breaks in a library.`

* **Method**
`PUT`

* **URL**
`/libraries/:libraryid/breaks`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
`none`

* **Data Params**
```json
[{
    "libraryid": 1234,
    "valuetype": "ID",
    "value": "12345",
    "breaktype": "SHELF"
}]
```

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Author Based Series

* **Action**
`Get the series that should be sorted primarily on volume, instead of author.`

* **Method**
`GET`

* **URL**
`/libraries/:libraryid/series`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Update Author Based Series

* **Action**
`Set the series that should be sorted primarily on volume, instead of author.`

* **Method**
`PUT`

* **URL**
`/libraries/:libraryid/series`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Sort Method

* **Action**
`Get the sort method of a library.`

* **Method**
`GET`

* **URL**
`/libraries/:libraryid/sort`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Update Sort Method

* **Action**
`Set the sort method of a library. One of "dewey", "title", or "series"`

* **Method**
`PUT`

* **URL**
`/libraries/:libraryid/sort`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

## Settings
* Get Settings `GET /settings`
* Save Settings `PUT /settings`
* Get Setting `GET /settings/:setting`

### Get Settings

* **Action**
`Get the user's settings. If nothing is set for a setting, return the default.`

* **Method**
`GET`

* **URL**
`/settings`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json
[{
    "name": "width",
    "value": "500",
    "valuetype": "nonnegativeinteger",
    "group": "BOOK",
    "possiblevalues": []
}]
```

* **Error Response**
```json

```

### Save Settings

* **Action**
`Save the user's settings.`

* **Method**
`PUT`

* **URL**
`/settings`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
```json
    "group": "Book",
    "name": "IsOwned",
    "value": "true",
    "valuetype": "select",
    "possiblevalues": [
        "true",
        "false"
    ]
```

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Setting

* **Action**
`Get a single user setting.`

* **Method**
`GET`

* **URL**
`/settings/:setting`

* **URL Params**
`none`

* **URL Arguments (?)**
``

* **Data Params**
Setting Name
`Width`

* **Success Response**
`settingvalue`

* **Error Response**
```json

```

## Users
* Get Users `GET /users`
* Login `POST /users/login`
* Register `POST /users`
* Logout `POST /users/logout`
* Send Password Reset `PUT /users/password`
* Finish Password Reset `GET /users/password/:token`
* Get Username `GET /users/username`

### Get Users

* **Action**
`Get the users in a library.`

* **Method**
`GET`

* **URL**
`/users`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json
    [{
        "id": 1234,
        "username": "janedoe",
        "first": "Jane",
        "last": "Doe",
        "fullname": "Jane Doe",
        "email": "jdoe@test.com",
        "iconurl": "/res/usericons/1234.jpg"
    }]
```

* **Error Response**
```json

```

### Login

* **Action**
`Log a user in to the system.`

* **Method**
`POST`

* **URL**
`/users/login`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
Form Data
  * `username`
  * `password`

* **Success Response**
Redirect to main page

* **Error Response**
```json

```

### Register

* **Action**
`Register a user in to the system.`

* **Method**
`POST`

* **URL**
`/users`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
  * `username`
  * `email`
  * `firstname`
  * `lastname`
  * `password`


* **Success Response**
Redirect to main page

* **Error Response**
```json

```

### Logout

* **Action**
`Log a user out of the system.`

* **Method**
`POST`

* **URL**
`/users/logout`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
Redirect to unregistered page

* **Error Response**
```json

```

### Send Password Reset

* **Action**
`Send a password reset request to the user's email.`

* **Method**
`PUT`

* **URL**
`/users/reset`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`email`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Finish Password Reset

* **Action**
`Complete a password reset request.`

* **Method**
`GET`

* **URL**
`/users/reset/:token`

* **URL Params**
`:token` The token to lookup for the reset

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
```json

```

* **Error Response**
```json

```

### Get Username

* **Action**
`Get the user's username.`

* **Method**
`GET`

* **URL**
`/users/username`

* **URL Params**
`none`

* **URL Arguments (?)**
`none`

* **Data Params**
`none`

* **Success Response**
`username`

* **Error Response**
```json

```
