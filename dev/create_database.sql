-- phpMyAdmin SQL Dump
-- version 4.6.6
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 11, 2017 at 03:17 AM
-- Server version: 10.1.13-MariaDB
-- PHP Version: 5.6.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `library`
--

-- --------------------------------------------------------

--
-- Table structure for table `bookcases`
--

CREATE TABLE `bookcases` (
  `CaseId` int(11) NOT NULL,
  `CaseNumber` int(1) NOT NULL,
  `Width` int(11) NOT NULL DEFAULT '500',
  `ShelfHeight` int(11) NOT NULL DEFAULT '350',
  `NumShelves` int(11) NOT NULL DEFAULT '1',
  `SpacerHeight` int(11) NOT NULL DEFAULT '12',
  `PaddingLeft` int(11) NOT NULL DEFAULT '0',
  `PaddingRight` int(11) NOT NULL DEFAULT '25',
  `BookMargin` int(11) NOT NULL DEFAULT '2'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `books`
--

CREATE TABLE `books` (
  `BookID` int(11) NOT NULL,
  `Title` varchar(255) NOT NULL,
  `Subtitle` varchar(255) DEFAULT NULL,
  `OriginallyPublished` date DEFAULT NULL,
  `PublisherID` int(11) DEFAULT NULL,
  `IsRead` tinyint(1) NOT NULL DEFAULT '0',
  `IsReference` tinyint(1) NOT NULL DEFAULT '0',
  `IsOwned` tinyint(1) NOT NULL DEFAULT '0',
  `ISBN` varchar(255) DEFAULT NULL,
  `LoaneeFirst` varchar(30) DEFAULT NULL,
  `LoaneeLast` varchar(30) DEFAULT NULL,
  `Dewey` varchar(255) DEFAULT NULL,
  `Pages` int(11) DEFAULT '0',
  `Width` decimal(10,0) DEFAULT '0',
  `Height` decimal(10,0) DEFAULT '0',
  `Depth` decimal(10,0) DEFAULT '0',
  `Weight` decimal(10,2) DEFAULT '0.00',
  `PrimaryLanguage` varchar(255) DEFAULT 'English',
  `SecondaryLanguage` varchar(255) DEFAULT NULL,
  `OriginalLanguage` varchar(255) DEFAULT 'English',
  `Series` varchar(255) DEFAULT NULL,
  `Volume` decimal(10,2) DEFAULT '0.00',
  `Format` varchar(255) DEFAULT NULL,
  `Edition` int(11) DEFAULT '1',
  `ImageURL` varchar(255) DEFAULT NULL,
  `IsReading` tinyint(4) DEFAULT '0',
  `IsShipping` tinyint(4) DEFAULT '0',
  `SpineColor` text,
  `CheapestNew` decimal(10,2) NOT NULL DEFAULT '99999999.99',
  `CheapestUsed` decimal(10,2) NOT NULL DEFAULT '99999999.99',
  `EditionPublished` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `dewey_numbers`
--

CREATE TABLE `dewey_numbers` (
  `Number` varchar(255) DEFAULT NULL,
  `Genre` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `formats`
--

CREATE TABLE `formats` (
  `Format` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `languages`
--

CREATE TABLE `languages` (
  `Langauge` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `library_members`
--

CREATE TABLE `library_members` (
  `id` int(11) NOT NULL,
  `usr` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `pass` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `regIP` varchar(15) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `dt` datetime NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `persons`
--

CREATE TABLE `persons` (
  `PersonID` int(11) NOT NULL,
  `FirstName` varchar(255) DEFAULT NULL,
  `MiddleNames` varchar(255) DEFAULT NULL,
  `LastName` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `publishers`
--

CREATE TABLE `publishers` (
  `PublisherID` int(11) NOT NULL,
  `Publisher` varchar(255) DEFAULT NULL,
  `City` varchar(255) DEFAULT NULL,
  `State` varchar(100) DEFAULT NULL,
  `Country` varchar(100) DEFAULT NULL,
  `ParentCompany` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `read_books`
--

CREATE TABLE `read_books` (
  `UserID` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `BookID` int(11) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `Role` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `rollback`
--

CREATE TABLE `rollback` (
  `id` int(11) NOT NULL,
  `type` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `series`
--

CREATE TABLE `series` (
  `Series` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `usersession`
--

CREATE TABLE `usersession` (
  `SessionKey` varchar(50) NOT NULL,
  `UserID` int(11) NOT NULL,
  `LastSeenTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `written_by`
--

CREATE TABLE `written_by` (
  `BookID` int(11) NOT NULL,
  `AuthorID` int(11) NOT NULL,
  `Role` varchar(255) NOT NULL DEFAULT 'Author'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `bookcases`
--
ALTER TABLE `bookcases`
  ADD PRIMARY KEY (`CaseId`);

--
-- Indexes for table `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`BookID`),
  ADD UNIQUE KEY `BookID` (`BookID`),
  ADD UNIQUE KEY `BookID_2` (`BookID`);

--
-- Indexes for table `dewey_numbers`
--
ALTER TABLE `dewey_numbers`
  ADD UNIQUE KEY `UQ__dewey_nu__78A1A19DCE629647` (`Number`);

--
-- Indexes for table `formats`
--
ALTER TABLE `formats`
  ADD PRIMARY KEY (`Format`),
  ADD UNIQUE KEY `UQ__formats__FB054B2EFC6A87EB` (`Format`);

--
-- Indexes for table `languages`
--
ALTER TABLE `languages`
  ADD PRIMARY KEY (`Langauge`),
  ADD UNIQUE KEY `UQ__language__AD43140AFA46337A` (`Langauge`);

--
-- Indexes for table `persons`
--
ALTER TABLE `persons`
  ADD PRIMARY KEY (`PersonID`),
  ADD UNIQUE KEY `PersonID` (`PersonID`);

--
-- Indexes for table `publishers`
--
ALTER TABLE `publishers`
  ADD PRIMARY KEY (`PublisherID`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`Role`);

--
-- Indexes for table `series`
--
ALTER TABLE `series`
  ADD PRIMARY KEY (`Series`),
  ADD UNIQUE KEY `UQ__series__1A00001F177C8503` (`Series`);

--
-- Indexes for table `usersession`
--
ALTER TABLE `usersession`
  ADD PRIMARY KEY (`SessionKey`);

--
-- Indexes for table `written_by`
--
ALTER TABLE `written_by`
  ADD PRIMARY KEY (`BookID`,`AuthorID`,`Role`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `bookcases`
--
ALTER TABLE `bookcases`
  MODIFY `CaseId` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;
--
-- AUTO_INCREMENT for table `books`
--
ALTER TABLE `books`
  MODIFY `BookID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20159;
--
-- AUTO_INCREMENT for table `persons`
--
ALTER TABLE `persons`
  MODIFY `PersonID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14948;
--
-- AUTO_INCREMENT for table `publishers`
--
ALTER TABLE `publishers`
  MODIFY `PublisherID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19401;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
