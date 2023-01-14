CREATE TABLE User (
    Id      INT AUTO_INCREMENT PRIMARY KEY,
    Name    VARCHAR(50) UNIQUE,
    Coins   INT
);

CREATE TABLE League (
    Id      INT AUTO_INCREMENT PRIMARY KEY,
    Name    VARCHAR(50),
    Slug    VARCHAR(50),
    Image   VARCHAR(50)
);

CREATE TABLE Team (
    Id      INT AUTO_INCREMENT PRIMARY KEY,
    Name    VARCHAR(50),
    Code    VARCHAR(50),
    Image   VARCHAR(50)
);

CREATE TABLE Bet (
    Id      INT AUTO_INCREMENT PRIMARY KEY,
    GameId  INT,
    UserId  INT,
    Value   INT,
    Team    TINYINT
);

CREATE TABLE Game (
    Id          INT AUTO_INCREMENT PRIMARY KEY,
    Team_1      INT NOT NULL,
    Team_2      INT NOT NULL,
    League      INT NOT NULL,
    Time        TIMESTAMP  DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CHECK ( Team_1 <> Team_2 ),
    FOREIGN KEY (League) REFERENCES League(Id),
    FOREIGN KEY (Team_1) REFERENCES Team(Id),
    FOREIGN KEY (Team_2) REFERENCES Team(Id)
);
