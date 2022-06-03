
-- +migrate Up
CREATE TABLE IF NOT EXISTS product (
  productId INT NOT NULL AUTO_INCREMENT,
  vendorId INT NOT NULL,
  name VARCHAR(200) NOT NULL,
  price FLOAT NOT NULL,
  PRIMARY KEY (productId),
  INDEX userId_idx (vendorId ASC));
-- +migrate Down
DROP TABLE IF EXISTS product;
