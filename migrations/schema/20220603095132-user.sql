
-- +migrate Up
CREATE TABLE IF NOT EXISTS usr (
    userId INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(60) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(10) NOT NULL);
  
-- +migrate Down
DROP TABLE IF EXISTS user;