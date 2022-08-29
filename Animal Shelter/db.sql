create database animal_shelter

drop table shelters;
create table shelters (
    `Id` int auto_increment,
    `Name` varchar(255) not null default "",
    `Zip` int(4) not null default 0000,
    `City` varchar(255) not null default "",
    `Street` varchar(255) not null default "",
    `StreetNumber` int not null default 0,
    `Email` varchar(255) not null default "",
    `PhoneNumber` int(9) default null,
    `Webpage` varchar(255) default null,

    primary key(`Id`)
);

drop table species;
create table species (
    `Id` int auto_increment,
    `Name` varchar(255),

    primary key(`Id`)
);

drop table animals;
create table animals (
    `Id` int auto_increment,
    `ShelterId` int not null,
    `SpeciesId` int not null,
    `Name` varchar(255) not null default "névtelen",
    `Breed` varchar(255) not null default "keverék",
    `Gender` enum("hím", "nőstény") not null default "hím",
    `Age` int(3) not null default 0,
    `Description` varchar(255) not null default "",
    `Height` int(3) not null default 0 comment "cm",
    `Weight` decimal(6, 2) not null default 0.0 comment "kg",

    primary key(`Id`),
    foreign key(`ShelterId`) references shelters(`Id`),
    foreign key(`SpeciesId`) references species(`Id`)
);

drop table suggested_species;
create table suggested_species (
    `Name` varchar(255)
);


insert into species(`Name`) values 
    ("kutya"),
    ("macska"),
    ("nyúl"),
    ("hörcsög"),
    ("papagáj")
;