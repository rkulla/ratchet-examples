-- Note: Before running this make sure you understand it. Use at own risk.
-- Usage:
--   mysql> source schema.sql

DROP DATABASE IF EXISTS `srcDB`;
CREATE DATABASE `srcDB`;
DROP DATABASE IF EXISTS `dstDB`;
CREATE DATABASE `dstDB`;

use srcDB

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `addresses`;
CREATE TABLE `addresses` (
  `id` int(11) NOT NULL,
  `city` varchar(100) NOT NULL DEFAULT '',
  `state` char(2) NOT NULL DEFAULT '',
  KEY `fk_id` (`id`),
  CONSTRAINT `fk_id` FOREIGN KEY (`id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `users` (`id`, `name`)
VALUES
	(123, 'Alex'),
	(456, 'John'),
	(789, 'Jane');

INSERT INTO `addresses` (`id`, `city`, `state`)
VALUES
	(123, 'Austin', 'TX'),
	(456, 'Los Angeles', 'CA'),
	(789, 'San Diego', 'CA');

use dstDB

DROP TABLE IF EXISTS `users2`;
CREATE TABLE `users2` (
  `user_id` int(11) NOT NULL,
  `city` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
