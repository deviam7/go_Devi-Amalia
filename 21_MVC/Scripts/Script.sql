CREATE TABLE `crud_go`.`users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `crud_go`.`user` (`name`, `email`, `password`) VALUES ('Devi Amalia', 'deviam@example.com', 'devi123');
