/*
 Navicat MySQL Data Transfer

 Source Server         : byron
 Source Server Type    : MySQL
 Source Server Version : 50716
 Source Host           : localhost
 Source Database       : movie

 Target Server Type    : MySQL
 Target Server Version : 50716
 File Encoding         : utf-8

 Date: 11/18/2016 18:25:45 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `info`
-- ----------------------------
DROP TABLE IF EXISTS `info`;
CREATE TABLE `info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `imid` varchar(45) NOT NULL COMMENT '信息的id，删除查询等操作使用',
  `title` varchar(500) DEFAULT NULL COMMENT '标题',
  `year` varchar(45) DEFAULT NULL COMMENT '年份',
  `runtime` varchar(45) DEFAULT NULL COMMENT '时长，单位秒',
  `actors` varchar(500) DEFAULT NULL COMMENT '演员列表',
  `plot` varchar(500) DEFAULT NULL COMMENT '描述',
  `language` varchar(45) DEFAULT NULL COMMENT '语言',
  `country` varchar(45) DEFAULT NULL COMMENT '国家',
  `poster` varchar(200) DEFAULT NULL COMMENT '海报的URL',
  `type` varchar(45) DEFAULT NULL COMMENT '类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `imid_UNIQUE` (`imid`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

-- ----------------------------
--  Records of `info`
-- ----------------------------
BEGIN;
INSERT INTO `info` VALUES ('5', 'tt0068646', 'The Godfather', '1972', '175 min', 'Marlon Brando, Al Pacino, James Caan, Richard S. Castellano', 'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.', 'English, Italian, Latin', 'USA', 'https://images-na.ssl-images-amazon.com/images/M/MV5BNTUxOTdjMDMtMWY1MC00MjkxLTgxYTMtYTM1MjU5ZTJlNTZjXkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg', 'movie'), ('6', 'tt1187043', '3 Idiots', '2009', '170 min', 'Aamir Khan, Madhavan, Sharman Joshi, Kareena Kapoor', 'N/A', 'Hindi, English', 'India', 'https://images-na.ssl-images-amazon.com/images/M/MV5BZWRlNDdkNzItMzhlZC00YTdmLWIwNjktYjY5NjQ1ZmQ3N2FkXkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_SX300.jpg', 'movie'), ('7', 'tt0251413', 'Star Wars', '1983', 'N/A', 'Harrison Ford, Alec Guinness, Mark Hamill, James Earl Jones', 'N/A', 'English', 'USA', 'https://images-na.ssl-images-amazon.com/images/M/MV5BMWJhYWQ3ZTEtYTVkOS00ZmNlLWIxZjYtODZjNTlhMjMzNGM2XkEyXkFqcGdeQXVyNzg5OTk2OA@@._V1_SX300.jpg', 'game');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
