CREATE database IF NOT EXISTS `devbook`;
USE `devbook`;

drop table if exists `devbook`.`posts`;
drop table if exists `devbook`.`followers`;
drop table if exists `devbook`.`users`;

CREATE TABLE `devbook`.`users` (
                                   `id` int NOT NULL AUTO_INCREMENT primary key,
                                   `name` varchar(255) NOT NULL,
                                   `nick` varchar(255) NOT NULL,
                                   `password` varchar(255) NOT NULL,
                                   `email` varchar(255) NOT NULL,
                                   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE table followers (
                          user_id int NOT NULL,
                          foreign key (user_id) references users(id) on delete cascade,
                          follower_id int NOT NULL,
                            foreign key (follower_id) references users(id) on delete cascade,
                          created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
                            primary key (user_id, follower_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into users (name, nick, password, email)
values
    ('John Doe', 'johndoe', '$2a$10$QMoM5cwzfflL71gDhrMBSO7XRF8Djjgt/uJULEByOxQh3kxwAx5gm', 'john@email.com'),
    ('Jane Doe', 'janedoe', '$2a$10$QMoM5cwzfflL71gDhrMBSO7XRF8Djjgt/uJULEByOxQh3kxwAx5gm', 'jane@email.com'),
    ('John2 Smith', 'johnsmith2', '$2a$10$QMoM5cwzfflL71gDhrMBSO7XRF8Djjgt/uJULEByOxQh3kxwAx5gm', 'john2@email.com');

insert into followers (user_id, follower_id)
values
    (1, 2),
    (3, 1),
    (1, 3);


CREATE TABLE `devbook`.`posts` (
                                  `id` int NOT NULL AUTO_INCREMENT primary key,
                                  `title` varchar(255) NOT NULL,
                                  `content` text NOT NULL,
                                    `author_id` int NOT NULL,
                                    foreign key (author_id) references users(id) on delete cascade,
                                    `likes` int NOT NULL DEFAULT 0,
                                    `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;