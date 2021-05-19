-- use paxful-test;

CREATE TABLE transactions (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    idempotence_key VARCHAR(50),
    expired DATETIME,
    attempt INTEGER
);