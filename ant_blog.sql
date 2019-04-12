-- ALTER TABLE `article` DROP FOREIGN KEY `fk_articles_articles_3`;
-- ALTER TABLE `tokens` DROP FOREIGN KEY `fk_tokens`;
-- ALTER TABLE `tag` DROP FOREIGN KEY `fk_tag`;
-- ALTER TABLE `category` DROP FOREIGN KEY `fk_category`;
--
-- DROP TABLE `article`;
-- DROP TABLE `tag`;
-- DROP TABLE `category`;
-- DROP TABLE `user`;
-- DROP TABLE `tokens`;

CREATE TABLE `article` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`title` varchar(255) NOT NULL,
`userId` int(11) NOT NULL,
`content` longtext NOT NULL,
`publishAt` datetime NOT NULL,
`viewCount` int(11) NULL DEFAULT 0,
PRIMARY KEY (`id`) 
);

CREATE TABLE `tag` (
`id` int(11) NOT NULL,
`tagName` varchar(255) NOT NULL,
`articleId` int(11) NULL,
PRIMARY KEY (`id`) 
);

CREATE TABLE `category` (
`id` int(11) NOT NULL,
`categoryName` varchar(255) NOT NULL,
`articleId` int(11) NULL,
PRIMARY KEY (`id`) 
);

CREATE TABLE `user` (
`id` int(11) NOT NULL,
`username` varchar(255) NOT NULL,
`password` varchar(255) NULL,
PRIMARY KEY (`id`) 
);

CREATE TABLE `tokens` (
`id` int(11) NOT NULL,
`userId` int(11) NULL,
`token` varchar(255) NULL,
`expired` datetime NULL,
`loginAt` datetime NULL,
`ip` varchar(255) NULL,
`userAgent` varchar(255) NULL,
PRIMARY KEY (`id`) 
);


ALTER TABLE `article` ADD CONSTRAINT `fk_articles_articles_3` FOREIGN KEY (`userId`) REFERENCES `user` (`id`);
ALTER TABLE `tokens` ADD CONSTRAINT `fk_tokens` FOREIGN KEY (`userId`) REFERENCES `user` (`id`);
ALTER TABLE `tag` ADD CONSTRAINT `fk_tag` FOREIGN KEY (`articleId`) REFERENCES `article` (`id`);
ALTER TABLE `category` ADD CONSTRAINT `fk_category` FOREIGN KEY (`articleId`) REFERENCES `article` (`id`);

