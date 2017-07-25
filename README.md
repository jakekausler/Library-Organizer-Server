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
* Viewing Statistics

## Registering and Logging In

* Registering an Account
* Logging In to an Account
* Forgotten Password

### Registering an Account

### Logging In to an Account

### Forgotten Password

## Managing Libraries and Permissions

* Managing Libraries
* Granting Permissions

### Managing Libraries

### Granting Permissions

## Managing User Settings

* Shelf Settings
* Book Settings

### Shelf Settings

### Book Settings

## Searching for Books

## Adding and Editing Books

* Add from Scratch
* Add from ISBN
* Add from Existing Book
* Edit an Existing Book
* Bulk Editing

### Add from Scratch

### Add from ISBN

### Add from Existing Book

### Edit an Existing Book

### Bulk Editing

## Checking Out and Returning Books

* Checking out a Book
* Returning a Book

### Checking Out a Book

### Returing a Book

## Exporting and Importing CSV Files

* The CSV Format
* Exporting
* Importing

### The CSV Format

### Exporting

### Importing

## Viewing a Library

* Navigating the Library
* Searching for Books

### Navigating the Library

### Searching for Books

## Editing a Library

* Managing Cases
* Managing Breaks

### Managing Cases

### Managing Breaks

## Viewing Statistics

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
  * Add Break `POST /libraries/:libraryid/breaks`
  * Save Break `PUT /libraries/:libraryid/breaks`
  * Delete Break `DELETE /libraries/:libraryid/breaks`

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
  * `sortmethod`: Primary method of sorting. One of `dewey`, `series`, or `title`.
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
* Save Break `PUT /libraries/:libraryid/breaks`
* Delete Break `DELETE /libraries/:libraryid/breaks`

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

### Add Break

* **Action**
`Add a break to a library.`

* **Method**
`POST`

* **URL**
`/libraries/:libraryid/breaks`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
`none`

* **Data Params**
```json
{
    "libraryid": 1234,
    "valuetype": "ID",
    "value": "12345",
    "breaktype": "SHELF"
}
```

* **Success Response**
```json

```

* **Error Response**
```json

```

### Save Break

* **Action**
`Save a break in a library.`

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
{
    "libraryid": 1234,
    "valuetype": "ID",
    "value": "12345",
    "breaktype": "SHELF"
}
```

* **Success Response**
```json

```

* **Error Response**
```json

```

### Delete Break

* **Action**
`Delete a book from a library.`

* **Method**
`DELETE`

* **URL**
`/libraries/:libraryid/breaks`

* **URL Params**
`:libraryid` The id of the selected library

* **URL Arguments (?)**
`none`

* **Data Params**
```json
{
    "libraryid": 1234,
    "valuetype": "ID",
    "value": "12345",
    "breaktype": "SHELF"
}
```

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
