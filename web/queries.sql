CREATE TABLE libraries (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    ownerid int NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE shelves
ADD libraryid int NOT NULL;

ALTER TABLE books
ADD libraryid int NOT NULL;

CREATE TABLE permissions (
    userid int NOT NULL,
    libraryid int NOT NULL,
    permission int NOT NULL
);

CREATE TABLE breaks (
    libraryid int NOT NULL,
    breaktype varchar(255) NOT NULL,
    value varchar(255) NOT NULL,
    PRIMARY KEY (libaryid, breaktype)
);

CREATE TABLE librarysettings (
    libraryid int NOT NULL,
    setting varchar(255) NOT NULL,
    value varchar(255),
    valuetype varchar(255),
    PRIMARY KEY (libaryid, setting)
);

CREATE TABLE librarysettingspossiblevalues (
    setting varchar(255) NOT NULL,
    possiblevalue varchar(255) NOT NULL
);
