/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MariaDB
 Source Server Version : 100136
 Source Host           : localhost:3306
 Source Schema         : parking

 Target Server Type    : MariaDB
 Target Server Version : 100136
 File Encoding         : 65001

 Date: 20/06/2019 15:36:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cars
-- ----------------------------
DROP TABLE IF EXISTS `cars`;
CREATE TABLE `cars` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `slot_number` int(11) NOT NULL,
  `plat_number` varchar(255) NOT NULL,
  `color` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of cars
-- ----------------------------
BEGIN;
INSERT INTO `cars` VALUES (1, 1, 'KA-01-HH-1234', 'White');
INSERT INTO `cars` VALUES (2, 2, 'KA-01-HH-9999', 'White');
INSERT INTO `cars` VALUES (3, 3, 'KA-01-BB-0001', 'Black');
INSERT INTO `cars` VALUES (5, 5, 'KA-01-HH-2701', 'Blue');
INSERT INTO `cars` VALUES (6, 6, 'KA-01-HH-3141', 'Black');
INSERT INTO `cars` VALUES (7, 4, 'KA-01-P-333', 'White');
COMMIT;

-- ----------------------------
-- Table structure for parking_slot_detail
-- ----------------------------
DROP TABLE IF EXISTS `parking_slot_detail`;
CREATE TABLE `parking_slot_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `slot` int(11) NOT NULL,
  `status` enum('terisi','kosong') DEFAULT 'kosong',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of parking_slot_detail
-- ----------------------------
BEGIN;
INSERT INTO `parking_slot_detail` VALUES (1, 1, 'terisi');
INSERT INTO `parking_slot_detail` VALUES (2, 2, 'terisi');
INSERT INTO `parking_slot_detail` VALUES (3, 3, 'terisi');
INSERT INTO `parking_slot_detail` VALUES (4, 4, 'terisi');
INSERT INTO `parking_slot_detail` VALUES (5, 5, 'terisi');
INSERT INTO `parking_slot_detail` VALUES (6, 6, 'terisi');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
