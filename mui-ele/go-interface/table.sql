CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(45) DEFAULT NULL,
  `password` char(32) DEFAULT NULL,
  `phone` char(11) DEFAULT NULL,
  `headimg` text,
  `content` text,
  `class` int(11) DEFAULT NULL,
  `cTime` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

CREATE TABLE `ele`.`feedback` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `phone` VARCHAR(45) NULL,
  `content` VARCHAR(450) NULL,
  `cTime` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));
