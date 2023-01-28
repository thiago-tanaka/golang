CREATE database IF NOT EXISTS `devbook`;
USE `devbook`;

drop table if exists `devbook`.`users`;

CREATE TABLE `devbook`.`users` (
                                   `id` int NOT NULL AUTO_INCREMENT primary key,
                                   `name` varchar(255) NOT NULL,
                                   `nick` varchar(255) NOT NULL,
                                   `password` varchar(255) NOT NULL,
                                   `email` varchar(255) NOT NULL,
                                   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
