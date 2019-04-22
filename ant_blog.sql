ALTER TABLE `tokens` DROP FOREIGN KEY `fk_tokens`;
ALTER TABLE `article` DROP FOREIGN KEY `fk_article`;
ALTER TABLE `article` DROP FOREIGN KEY `fk_article_1`;
ALTER TABLE `article` DROP FOREIGN KEY `fk_article_2`;

DROP INDEX `artilceIdIndex` ON `article`;

DROP TABLE `article`;
DROP TABLE `tag`;
DROP TABLE `category`;
DROP TABLE `user`;
DROP TABLE `tokens`;

CREATE TABLE `article` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `title` varchar(255) NOT NULL,
                         `userId` int(11) NOT NULL,
                         `content` longtext NOT NULL,
                         `publishAt` datetime NOT NULL,
                         `viewCount` int(11) NULL DEFAULT 0,
                         `categoryId` int(11) NOT NULL,
                         `tagId` int(11) NOT NULL,
                         PRIMARY KEY (`id`) ,
                         UNIQUE INDEX `artilceIdIndex` (`id` ASC)
) DEFAULT CHARSET=utf8;

CREATE TABLE `tag` (
                     `id` int(11) NOT NULL AUTO_INCREMENT,
                     `tagName` varchar(255) NOT NULL,
                     PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;

CREATE TABLE `category` (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `categoryName` varchar(255) NOT NULL,
                          PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
                      `id` int(11) NOT NULL AUTO_INCREMENT,
                      `username` varchar(255) NOT NULL,
                      `password` varchar(255) NULL,
                      PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;

CREATE TABLE `tokens` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `userId` int(11) NULL,
                        `token` varchar(255) NULL,
                        `expired` datetime NULL,
                        `loginAt` datetime NULL,
                        `ip` varchar(255) NULL,
                        `userAgent` varchar(255) NULL,
                        PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;


ALTER TABLE `tokens` ADD CONSTRAINT `fk_tokens` FOREIGN KEY (`userId`) REFERENCES `user` (`id`);
ALTER TABLE `article` ADD CONSTRAINT `fk_article` FOREIGN KEY (`userId`) REFERENCES `user` (`id`);
ALTER TABLE `article` ADD CONSTRAINT `fk_article_1` FOREIGN KEY (`categoryId`) REFERENCES `category` (`id`);
ALTER TABLE `article` ADD CONSTRAINT `fk_article_2` FOREIGN KEY (`tagId`) REFERENCES `tag` (`id`);

