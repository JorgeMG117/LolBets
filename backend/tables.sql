CREATE TABLE League (
    Id      INT AUTO_INCREMENT PRIMARY KEY,
    Name    VARCHAR(50)
);

CREATE TABLE Team (
    Id      INT AUTO_INCREMENT PRIMARY KEY,
    Name    VARCHAR(50)
);

CREATE TABLE Game (
    Id          INT AUTO_INCREMENT PRIMARY KEY,
    Team_1      INT NOT NULL,
    Team_2      INT NOT NULL,
    League      INT NOT NULL,
    /*Date        date,*/
    CHECK ( Team_1 <> Team_2 ),
    FOREIGN KEY (League) REFERENCES League(Id),
    FOREIGN KEY (Team_1) REFERENCES Team(Id),
    FOREIGN KEY (Team_2) REFERENCES Team(Id)
);
