# Library-Organizer-2.0

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
  * Login `PUT /users`
  * Register `POST /users`
  * Logout `DELETE /users`
  * Send Password Reset `PUT /users/password`
  * Finish Password Reset `GET /users/password/:token`
  * Get Username `GET /users/username`

# Books
* Get Books `GET /books`
* Save Book `PUT /books`
* Add Book `POST /books`
* Delete Book `DELETE /books`
* Checkout Book `PUT /books/checkout`
* Checkin Book `PUT /books/checkin`
* Export Books `GET /books/books`
* Export Contributors `GET /books/contributors`
* Import Books `POST /books/books`

## Get Books

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

## Save Book

* **Action**
`Saves a book.`

* **Method**
`PUT`

* **URL**
`/books`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Add Book

* **Action**
`Adds a book.`

* **Method**
`POST`

* **URL**
`/books`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Delete Book

* **Action**
`Deletes a book.`

* **Method**
`DELETE`

* **URL**
`/books`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Checkout Book

* **Action**
`Checks out a book.`

* **Method**
`PUT`

* **URL**
`/books/checkout`

* **URL Params**
`none`

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Checkin Book

* **Action**
`Checks in a book.`

* **Method**
`PUT`

* **URL**
`/books/checkin`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Export Books

* **Action**
`Exports a csv of book data.`

* **Method**
`GET`

* **URL**
`/books/books`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Export Contributors

* **Action**
`Exports a csv of contributors`

* **Method**
`GET`

* **URL**
`/books/contributors`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Import Books

* **Action**
`Imports books from a csv to a library. The column headers should match the database fields names, with the exception of publisher and contributors. Publisher should be split into four columns: "publisher", "city", "state", "country". Contributors should be in a single column, with the header "contributors". They should be written as follows: "lastName[, firstName middleName1 middleName2... middleNameN]: role", with different contributors separated by a single semicolon and omitting first and middle names as needed.`

* **Method**
`POST`

* **URL**
`/books/books`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

# Information
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

## Get Statistics

* **Action**
`Get a set of statistics.`

* **Method**
`GET`

* **URL**
`/information/statistics`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Dimensions

* **Action**
`Get a set of dimensions.`

* **Method**
`GET`

* **URL**
`/information/dimensions`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Publishers

* **Action**
`Get a the set of publishers.`

* **Method**
`GET`

* **URL**
`/information/publishers`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Cities

* **Action**
`Get the set of cities.`

* **Method**
`GET`

* **URL**
`/information/cities`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get States

* **Action**
`Get the set of states.`

* **Method**
`GET`

* **URL**
`/information/states`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Countries

* **Action**
`Get the set of countries. `

* **Method**
`GET`

* **URL**
`/information/countries`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Formats

* **Action**
`Get the set of formats.`

* **Method**
`GET`

* **URL**
`/information/formats`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Roles

* **Action**
`Get the set of contributor roles.`

* **Method**
`GET`

* **URL**
`/information/roles`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Series

* **Action**
`Get the set of series.`

* **Method**
`GET`

* **URL**
`/information/series`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Languages

* **Action**
`Get the set of languages`

* **Method**
`GET`

* **URL**
`/information/languages`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Deweys

* **Action**
`Get the set of deweys`

* **Method**
`GET`

* **URL**
`/information/deweys`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

# Libraries
* Get Libraries `GET /libaries`
* Get Owned Libraries `GET /libraries/owned`
* Save Owned Libraries `PUT /libraries/owned`
* Get Cases `GET /libraries/:libraryid/cases`
* Save Cases `PUT /libraries/:libraryid/cases`
* Get Breaks `GET /libraries/:libraryid/breaks`
* Add Break `POST /libraries/:libraryid/breaks`
* Save Break `PUT /libraries/:libraryid/breaks`
* Delete Break `DELETE /libraries/:libraryid/breaks`

## Get Libraries

* **Action**
`Get a list of libraries with which the user has permission to do something.`

* **Method**
`GET`

* **URL**
`/libraries`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Owned Libraries

* **Action**
`Get the libraries which the user owns and the user who can do something with them.`

* **Method**
`GET`

* **URL**
`/libraries/owned`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Save Owned Libraries

* **Action**
`Saves the libraries the user owns and the users who can do something with them.`

* **Method**
`PUT`

* **URL**
`/libraries/owned`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Cases

* **Action**
`Get the cases in a library.`

* **Method**
`GET`

* **URL**
`/libraries/cases`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Save Cases

* **Action**
`Save the cases in a library.`

* **Method**
`PUT`

* **URL**
`/libraries/cases`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Breaks

* **Action**
`Get the breaks for a library.`

* **Method**
`GET`

* **URL**
`/libraries/:libraryid/breaks`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Add Break

* **Action**
`Add a break to a library.`

* **Method**
`POST`

* **URL**
`/libraries/:libraryid/breaks`
``

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Save Break

* **Action**
`Save a break in a library.`

* **Method**
`PUT`

* **URL**
`/libraries/:libraryid/breaks`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Delete Break

* **Action**
`Delete a book from a library.`

* **Method**
`DELETE`

* **URL**
`/libraries/:libraryid/breaks`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

# Settings
* Get Settings `GET /settings`
* Save Settings `PUT /settings`
* Get Setting `GET /settings/:setting`

## Get Settings

* **Action**
`Get the user's settings. If nothing is set for a setting, return the default.`

* **Method**
`GET`

* **URL**
`/settings`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Save Settings

* **Action**
`Save the user's settings.`

* **Method**
`PUT`

* **URL**
`/settings`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Setting

* **Action**
`Get a single user setting.`

* **Method**
`GET`

* **URL**
`/settings/:setting`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

# Users
* Get Users `GET /users`
* Login `PUT /users`
* Register `POST /users`
* Logout `DELETE /users`
* Send Password Reset `PUT /users/password`
* Finish Password Reset `GET /users/password/:token`
* Get Username `GET /users/username`

## Get Users

* **Action**
`Get the users in a library.`

* **Method**
`GET`

* **URL**
`/users`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Login

* **Action**
`Log a user in to the system.`

* **Method**
`PUT`

* **URL**
`/users`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Register

* **Action**
`Register a user in to the system.`

* **Method**
`POST`

* **URL**
`/users`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Logout

* **Action**
`Log a user out of the system.`

* **Method**
`DELETE`

* **URL**
`/users`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Send Password Reset

* **Action**
`Send a password reset request to the user's email.`

* **Method**
`PUT`

* **URL**
`/users/reset`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Finish Password Reset

* **Action**
`Complete a password reset request.`

* **Method**
`GET`

* **URL**
`/users/reset/:token`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```

## Get Username

* **Action**
`Get the user's username.`

* **Method**
`GET`

* **URL**
`/users/username`

* **URL Params**
``

* **URL Arguments (?)**
``

* **Data Params**
``

* **Success Response**
```json

```

* **Error Response**
```json

```
