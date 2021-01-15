/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.10.128
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 192.168.10.128:3306
 Source Schema         : userdb

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 16/01/2021 00:14:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for file_info
-- ----------------------------
DROP TABLE IF EXISTS `file_info`;
CREATE TABLE `file_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `size` bigint(20) NOT NULL DEFAULT 0,
  `upload_at` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of file_info
-- ----------------------------
INSERT INTO `file_info` VALUES (1, 'log2.log', 7436, '2021-01-15 20:32:47');
INSERT INTO `file_info` VALUES (2, 'log3.log', 10571, '2021-01-15 22:47:16');
INSERT INTO `file_info` VALUES (3, 'log4.log', 57370, '2021-01-15 22:50:32');

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `p_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, '添加用户', '/usermanager/add');
INSERT INTO `permission` VALUES (2, '查找用户', '/usermanager/query');
INSERT INTO `permission` VALUES (3, '日志操作', '/loganalysis/upload');
INSERT INTO `permission` VALUES (4, '根', '/');
INSERT INTO `permission` VALUES (5, '用户列表', '/usermanager/list');
INSERT INTO `permission` VALUES (6, '修改用户', '/usermanager/modify');
INSERT INTO `permission` VALUES (7, '删除用户', '	/usermanager/del');
INSERT INTO `permission` VALUES (8, '角色管理', '/usermanager/role');
INSERT INTO `permission` VALUES (9, '数据接口', '/loganalysis/dataapi');
INSERT INTO `permission` VALUES (10, 'web展示', '/loganalysis/showweb');

-- ----------------------------
-- Table structure for recorder
-- ----------------------------
DROP TABLE IF EXISTS `recorder`;
CREATE TABLE `recorder`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `file_id` bigint(20) NOT NULL DEFAULT 0,
  `ip_addr` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `method` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT 0,
  `insert_at` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 200 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of recorder
-- ----------------------------
INSERT INTO `recorder` VALUES (1, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (2, 1, '10.106.37.191', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (3, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (4, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (5, 1, '10.106.37.191', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (6, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (7, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (8, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (9, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (10, 1, '10.106.37.191', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (11, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (12, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (13, 1, '10.106.37.191', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (14, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (15, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (16, 1, '10.10.14.66', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (17, 1, '10.53.1.73', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (18, 1, '10.106.37.191', 'GET', 200, '2021-01-15 20:32:47');
INSERT INTO `recorder` VALUES (19, 2, '10.103.20.155', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (20, 2, '10.103.20.155', 'GET', 304, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (21, 2, '10.103.20.155', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (22, 2, '10.103.20.155', 'GET', 304, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (23, 2, '10.103.20.155', 'GET', 304, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (24, 2, '10.103.20.155', 'GET', 304, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (25, 2, '10.103.20.155', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (26, 2, '10.103.20.155', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (27, 2, '10.103.20.155', 'GET', 502, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (28, 2, '10.103.20.155', 'GET', 302, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (29, 2, '10.103.20.155', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (30, 2, '10.103.20.155', 'GET', 304, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (31, 2, '10.64.36.168', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (32, 2, '10.103.20.194', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (33, 2, '10.53.1.73', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (34, 2, '10.106.38.116', 'GET', 200, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (35, 2, '10.10.100.39', 'POST', 201, '2021-01-15 22:47:16');
INSERT INTO `recorder` VALUES (36, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:32');
INSERT INTO `recorder` VALUES (37, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:32');
INSERT INTO `recorder` VALUES (38, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:32');
INSERT INTO `recorder` VALUES (39, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:32');
INSERT INTO `recorder` VALUES (40, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:32');
INSERT INTO `recorder` VALUES (41, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:32');
INSERT INTO `recorder` VALUES (42, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (43, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (44, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (45, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (46, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (47, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (48, 3, '10.64.36.168', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (49, 3, '10.103.20.155', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (50, 3, '10.103.20.194', 'POST', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (51, 3, '10.103.20.194', 'POST', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (52, 3, '10.106.38.199', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (53, 3, '10.103.20.254', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (54, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (55, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (56, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (57, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (58, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (59, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (60, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (61, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (62, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (63, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (64, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (65, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (66, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (67, 3, '10.64.36.168', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (68, 3, '10.64.36.55', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (69, 3, '10.103.20.155', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (70, 3, '10.103.20.194', 'POST', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (71, 3, '10.103.20.194', 'POST', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (72, 3, '10.103.20.194', 'POST', 302, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (73, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (74, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (75, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (76, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (77, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (78, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (79, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (80, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (81, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (82, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (83, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (84, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (85, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (86, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (87, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (88, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (89, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (90, 3, '10.106.38.199', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (91, 3, '10.103.20.254', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (92, 3, '10.64.36.168', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (93, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (94, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (95, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (96, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (97, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (98, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (99, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (100, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (101, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (102, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (103, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (104, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (105, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (106, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (107, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (108, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (109, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (110, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (111, 3, '10.103.20.194', 'GET', 302, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (112, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (113, 3, '10.103.20.155', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (114, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (115, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (116, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (117, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (118, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (119, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (120, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (121, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (122, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (123, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (124, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (125, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (126, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (127, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (128, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (129, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (130, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (131, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (132, 3, '10.106.38.199', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (133, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (134, 3, '10.64.36.168', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (135, 3, '10.103.20.254', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (136, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (137, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (138, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (139, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (140, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (141, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (142, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (143, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (144, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (145, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (146, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (147, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (148, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (149, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (150, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (151, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (152, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (153, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (154, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (155, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (156, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (157, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (158, 3, '10.103.20.194', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (159, 3, '10.103.20.155', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (160, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (161, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (162, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (163, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (164, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (165, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (166, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (167, 3, '10.103.20.194', 'GET', 304, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (168, 3, '10.103.21.126', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (169, 3, '10.106.37.191', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (170, 3, '10.64.36.168', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (171, 3, '10.106.38.199', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (172, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (173, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (174, 3, '10.103.20.254', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (175, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (176, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (177, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (178, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (179, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (180, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (181, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (182, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (183, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (184, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (185, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (186, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (187, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (188, 3, '10.103.20.155', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (189, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (190, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (191, 3, '10.64.36.168', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (192, 3, '10.64.33.16', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (193, 3, '10.106.38.199', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (194, 3, '10.103.20.254', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (195, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (196, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (197, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (198, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');
INSERT INTO `recorder` VALUES (199, 3, '10.103.12.197', 'GET', 200, '2021-01-15 22:50:33');

-- ----------------------------
-- Table structure for role_info
-- ----------------------------
DROP TABLE IF EXISTS `role_info`;
CREATE TABLE `role_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_info
-- ----------------------------
INSERT INTO `role_info` VALUES (1, '超级管理员');
INSERT INTO `role_info` VALUES (2, '管理员');
INSERT INTO `role_info` VALUES (3, '操作员');

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL DEFAULT 0,
  `permission_id` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO `role_permission` VALUES (1, 1, 1);
INSERT INTO `role_permission` VALUES (2, 1, 2);
INSERT INTO `role_permission` VALUES (3, 1, 3);
INSERT INTO `role_permission` VALUES (4, 2, 1);
INSERT INTO `role_permission` VALUES (5, 2, 2);
INSERT INTO `role_permission` VALUES (6, 3, 3);
INSERT INTO `role_permission` VALUES (7, 3, 4);
INSERT INTO `role_permission` VALUES (8, 1, 4);
INSERT INTO `role_permission` VALUES (9, 2, 4);
INSERT INTO `role_permission` VALUES (10, 1, 5);
INSERT INTO `role_permission` VALUES (11, 2, 5);
INSERT INTO `role_permission` VALUES (13, 1, 6);
INSERT INTO `role_permission` VALUES (14, 1, 7);
INSERT INTO `role_permission` VALUES (15, 2, 6);
INSERT INTO `role_permission` VALUES (16, 1, 8);
INSERT INTO `role_permission` VALUES (17, 2, 3);
INSERT INTO `role_permission` VALUES (18, 3, 9);
INSERT INTO `role_permission` VALUES (19, 3, 10);

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT 0,
  `role_id` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 5, 1);
INSERT INTO `user_role` VALUES (2, 6, 2);
INSERT INTO `user_role` VALUES (3, 7, 3);

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `sex` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `addr` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `tel` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `birthday` date NOT NULL,
  `passwd` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of userinfo
-- ----------------------------
INSERT INTO `userinfo` VALUES (5, 'admin', '0', 'pek', '110', '2020-10-10', '$2a$10$DNSCTjeOO7mVzDKvJeq2vOZ3GqCc7B7Kg2D01cQxx4VY4LoR/l1I2', '2021-01-12 23:26:55', '2021-01-12 23:26:55', NULL);
INSERT INTO `userinfo` VALUES (6, 'user1', '1', 'sdsd', '110', '2020-10-10', '$2a$10$zLfI8F7BRlIeUpRwDlnR1OEvXFhGGhb6Id8zpZhIU5npsYO9H/WwK', '2021-01-12 23:27:30', '2021-01-14 00:12:32', NULL);
INSERT INTO `userinfo` VALUES (7, 'user3', '0', 'sdsd', '110', '2020-12-16', '$2a$10$gP6OY/0K8Lr2gtoHp1gE6uKwqrNq2EOIBvu1dc09SBYGZ.Uo.8.hi', '2021-01-14 00:10:43', '2021-01-14 00:10:43', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (6, 'user1', 0, 'pek', '110', '1990-10-01 00:00:00', '111111', '2020-12-10 01:14:21', '2020-12-10 01:14:21');
INSERT INTO `users` VALUES (9, 'admin', 1, 'pek', '110', '2020-12-16 00:00:00', 'admin', '2020-12-16 00:35:41', '2020-12-16 00:35:41');

SET FOREIGN_KEY_CHECKS = 1;
