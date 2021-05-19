-- use paxful-test;

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(25),
    last_name  VARCHAR(30),
    middle_name  VARCHAR(30),
    email  VARCHAR(50)
);