/*
 Navicat Premium Data Transfer

 Source Server         : OVH-NVME-R1-DS
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Schema         : short

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 20/12/2020 01:04:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for shorts
-- ----------------------------
DROP TABLE IF EXISTS `shorts`;
CREATE TABLE `shorts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `long` varchar(1024) NOT NULL,
  `short` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;