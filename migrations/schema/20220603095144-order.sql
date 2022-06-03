
-- +migrate Up
CREATE TABLE IF NOT EXISTS`order`(
  orderId INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  cartId INT NOT NULL,
  subTotal FLOAT NOT NULL,
  timeCreated VARCHAR(45) NOT NULL);

-- +migrate Down
DROP TABLE IF EXISTS `order`;
