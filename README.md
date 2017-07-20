# Library-Organizer-2.0

# REST API
Note that this is not accurate until gorilla is set up
* Books `/books`
  * Get Books `GET /books`
  * Save Book `PUT /books`
  * Add Book `POST /books`
  * Delete Book `DELETE /books`
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
  * Get Breaks `GET /libraries/breaks`
  * Add Break `POST /libraries/breaks`
  * Save Break `PUT /libraries/breaks`
  * Delete Break `DELETE /libraries/breaks`
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
``

* **Method**
``

* **URL**
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

## Get Dimensions

* **Action**
``

* **Method**
``

* **URL**
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

## Get Publishers

* **Action**
``

* **Method**
``

* **URL**
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

## Get Cities

* **Action**
``

* **Method**
``

* **URL**
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

## Get States

* **Action**
``

* **Method**
``

* **URL**
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

## Get Countries

* **Action**
``

* **Method**
``

* **URL**
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

## Get Formats

* **Action**
``

* **Method**
``

* **URL**
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

## Get Roles

* **Action**
``

* **Method**
``

* **URL**
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

## Get Series

* **Action**
``

* **Method**
``

* **URL**
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

## Get Languages

* **Action**
``

* **Method**
``

* **URL**
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

## Get Dewys

* **Action**
``

* **Method**
``

* **URL**
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

# Libraries
* Get Libraries `GET /libaries`
* Get Owned Libraries `GET /libraries/owned`
* Save Owned Libraries `PUT /libraries/owned`
* Get Cases `GET /libraries/:libraryid/cases`
* Save Cases `PUT /libraries/:libraryid/cases`
* Get Breaks `GET /libraries/breaks`
* Add Break `POST /libraries/breaks`
* Save Break `PUT /libraries/breaks`
* Delete Break `DELETE /libraries/breaks`

## Get Libraries

* **Action**
``

* **Method**
``

* **URL**
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

## Get Owned Libraries

* **Action**
``

* **Method**
``

* **URL**
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

## Save Owned Libraries

* **Action**
``

* **Method**
``

* **URL**
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

## Get Cases

* **Action**
``

* **Method**
``

* **URL**
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

## Save Cases

* **Action**
``

* **Method**
``

* **URL**
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

## Get Breaks

* **Action**
``

* **Method**
``

* **URL**
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

## Add Break

* **Action**
``

* **Method**
``

* **URL**
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
``

* **Method**
``

* **URL**
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

## Delete Break

* **Action**
``

* **Method**
``

* **URL**
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

# Settings
* Get Settings `GET /settings`
* Save Settings `PUT /settings`
* Get Setting `GET /settings/:setting`

## Get Settings

* **Action**
``

* **Method**
``

* **URL**
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

## Save Settings

* **Action**
``

* **Method**
``

* **URL**
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

## Get Settings

* **Action**
``

* **Method**
``

* **URL**
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
``

* **Method**
``

* **URL**
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

## Login

* **Action**
``

* **Method**
``

* **URL**
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

## Register

* **Action**
``

* **Method**
``

* **URL**
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

## Logout

* **Action**
``

* **Method**
``

* **URL**
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

## Send Password Reset

* **Action**
``

* **Method**
``

* **URL**
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

## Finish Password Reset

* **Action**
``

* **Method**
``

* **URL**
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

## Get Username

* **Action**
``

* **Method**
``

* **URL**
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
