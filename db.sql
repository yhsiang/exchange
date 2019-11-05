CREATE DATABASE IF NOT EXISTS `sbi_exchange`;

USE `sbi_exchange`;

CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` varchar(36) NOT NULL DEFAULT '',
  `symbol` varchar(10) NOT NULL DEFAULT '',
  `side` varchar(10) NOT NULL DEFAULT '',
  `price` varchar(40) NOT NULL DEFAULT '',
  `quantity` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;