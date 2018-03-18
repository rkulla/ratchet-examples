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


INSERT INTO `users` (`id`, `name`)
VALUES
	(123, 'Alex'),
	(456, 'John');

use dstDB

DROP TABLE IF EXISTS `users2`;
CREATE TABLE `users2` (
  `user_id` int(11) NOT NULL,
  `some_new_field` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
