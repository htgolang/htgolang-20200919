/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.10.128
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 192.168.10.128:3306
 Source Schema         : usermanager

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 20/01/2021 00:12:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `sex` tinyint(1) NOT NULL DEFAULT 1,
  `addr` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `tel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `birthday` datetime(0) NOT NULL,
  `passwd` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (8, 'admin', 1, 'pek', '110', '1990-10-01 00:00:00', 'admin', '2021-01-17 23:45:19', '2021-01-17 23:45:19');
INSERT INTO `users` VALUES (9, 'user3', 1, 'pek', '110', '1990-10-01 00:00:00', '1111', '2021-01-17 23:46:22', '2021-01-17 23:46:22');

SET FOREIGN_KEY_CHECKS = 1;
