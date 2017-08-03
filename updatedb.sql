ALTER TABLE books
ADD COLUMN lexilecode varchar(2) NULL,
ADD COLUMN interestlevel decimal(10, 2) NULL,
ADD COLUMN ar decimal(10, 2) NULL,
ADD COLUMN learningaz decimal(10, 2) NULL,
ADD COLUMN guidedreading decimal(10, 2) NULL,
ADD COLUMN dra decimal(10, 2) NULL,
ADD COLUMN grade decimal(10, 2) NULL,
ADD COLUMN fountaspinnell decimal(10, 2) NULL,
ADD COLUMN age decimal(10, 2) NULL,
ADD COLUMN readingrecovery decimal(10, 2) NULL,
ADD COLUMN pmreaders decimal(10, 2) NULL,
MODIFY COLUMN lexile decimal(10, 2) NULL;

UPDATE books
SET lexile=NULL WHERE lexile=0;

CREATE TABLE awards (
	BookID int NOT NULL,
	Award varchar(255) NOT NULL,
	PRIMARY KEY (BookID, Award)
);

UPDATE librarysettings
SET lexile="";

INSERT INTO librarysettings
(setting, value, valuetype, userid)
VALUES
(Interest Level, null, nonnegativefloat, 0),
(AR, null, nonnegativefloat, 0),
(Learning AZ, null, nonnegativefloat, 0),
(Guided Reading, null, nonnegativefloat, 0),
(DRA, null, nonnegativefloat, 0),
(Grade, null, nonnegativefloat, 0),
(Fountas Spinnell, null, nonnegativefloat, 0),
(Age, null, nonnegativefloat, 0),
(Reading Recovery, null, nonnegativefloat, 0),
(PM Readers, null, nonnegativefloat, 0);