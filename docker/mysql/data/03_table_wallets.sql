-- use paxful-test;

CREATE TABLE wallets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    users_id INTEGER NOT NULL,
    balance FLOAT,
    FOREIGN KEY (users_id) REFERENCES users(id) ON DELETE CASCADE
);