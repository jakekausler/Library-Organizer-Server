
CREATE TABLE librarysettings (
    libraryid int NOT NULL,
    setting varchar(255) NOT NULL,
    value varchar(255),
    valuetype varchar(255),
    PRIMARY KEY (libraryid, setting)
);

CREATE TABLE librarysettingspossiblevalues (
    setting varchar(255) NOT NULL,
    possiblevalue varchar(255) NOT NULL
);
