CREATE DATABASE Final_Web_Project
CREATE TABLE Users(
    User_Id INT IDENTITY(1,1) PRIMARY KEY ,
    Email VARCHAR(300),
    Username VARCHAR(300),
    Password VARCHAR(300)
)
CREATE TABLE Threads(
    Thread_Id BIGINT IDENTITY(1,1) PRIMARY KEY,
    User_Id INT FOREIGN KEY REFERENCES Users(User_Id) NOT NULL ,
    Name VARCHAR(500),
    Description VARCHAR(500),
    Date_Created VARCHAR(100)
)
CREATE TABLE Responses(
    Response_Id BIGINT IDENTITY(1,1) PRIMARY KEY,
    User_Id INT FOREIGN KEY REFERENCES Users(User_Id) NOT NULL,
    Thread_Id BIGINT FOREIGN KEY REFERENCES Threads(Thread_Id) NOT NULL,
    Description VARCHAR(500),
    Date_Created VARCHAR(100)
)
CREATE TABLE Thread_Favorites(
    Thread_Favorites_Id BIGINT IDENTITY(1,1) PRIMARY KEY,
    User_Id INT FOREIGN KEY REFERENCES Users(User_Id) NOT NULL
)
CREATE TABLE Api_Favorites(
    Api_Favorites_Id BIGINT IDENTITY(1,1) PRIMARY KEY,
    User_Id INT FOREIGN KEY REFERENCES Users(User_Id) NOT NULL,
    Stock_Id VARCHAR(500) NOT NULL,
    Api_Url VARCHAR(500) NOT NULL
)
CREATE TABLE Sessions (
    SessionId INT IDENTITY(1,1) PRIMARY KEY,
    User_Id INT FOREIGN KEY REFERENCES Users(User_Id),
    UserGuid VARCHAR(100),
    IsActive BIT
)