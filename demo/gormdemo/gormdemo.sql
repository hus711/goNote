/*
 Navicat MySQL Data Transfer

 Source Server         : byron
 Source Server Type    : MySQL
 Source Server Version : 50716
 Source Host           : localhost
 Source Database       : gormdemo

 Target Server Type    : MySQL
 Target Server Version : 50716
 File Encoding         : utf-8

 Date: 11/18/2016 18:26:04 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `person`
-- ----------------------------
DROP TABLE IF EXISTS `person`;
CREATE TABLE `person` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `age` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=latin1;

-- ----------------------------
--  Records of `person`
-- ----------------------------
BEGIN;
INSERT INTO `person` VALUES ('2', '2016-11-11 15:23:48', '2016-11-11 15:23:48', '2016-11-11 15:29:37', 'hs1', '2'), ('3', '2016-11-11 15:23:48', '2016-11-11 15:23:48', '2016-11-11 15:29:37', 'hs2', '3'), ('4', '2016-11-11 15:23:48', '2016-11-11 15:23:48', '2016-11-11 15:29:37', 'hs3', '4'), ('5', '2016-11-11 15:23:48', '2016-11-11 15:23:48', '2016-11-11 15:29:37', 'hs4', '5'), ('6', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs0', '1'), ('7', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs1', '2'), ('8', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs2', '3'), ('9', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs3', '4'), ('10', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs4', '5'), ('11', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs5', '6'), ('12', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs6', '7'), ('13', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs7', '8'), ('14', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs8', '9'), ('15', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs9', '10'), ('16', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs10', '11'), ('17', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs11', '12'), ('18', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs12', '13'), ('19', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs13', '14'), ('20', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs14', '15'), ('21', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs15', '16'), ('22', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs16', '17'), ('23', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs17', '18'), ('24', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs18', '19'), ('25', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs19', '20'), ('26', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs20', '21'), ('27', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs21', '22'), ('28', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs22', '23'), ('29', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs23', '24'), ('30', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs24', '25'), ('31', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs25', '26'), ('32', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs26', '27'), ('33', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs27', '28'), ('34', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs28', '29'), ('35', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs29', '30'), ('36', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs30', '31'), ('37', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs31', '32'), ('38', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs32', '33'), ('39', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs33', '34'), ('40', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs34', '35'), ('41', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs35', '36'), ('42', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs36', '37'), ('43', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs37', '38'), ('44', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs38', '39'), ('45', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs39', '40'), ('46', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs40', '41'), ('47', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs41', '42'), ('48', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs42', '43'), ('49', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs43', '44'), ('50', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs44', '45'), ('51', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs45', '46'), ('52', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs46', '47'), ('53', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs47', '48'), ('54', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs48', '49'), ('55', '2016-11-11 15:30:27', '2016-11-11 15:30:27', null, 'hs49', '50'), ('56', null, null, null, 'huangshuai', '11');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
