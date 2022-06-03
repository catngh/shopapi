
-- +migrate Up
CREATE TABLE IF NOT EXISTS cart (
  cartId INT NOT NULL AUTO_INCREMENT,
  userId INT NOT NULL,
  item INT NULL,
  PRIMARY KEY (cartId));


-- +migrate Down
DROP TABLE IF EXISTS cart;