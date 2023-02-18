/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : bluebell

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 24/08/2022 17:10:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comment_id` bigint NOT NULL COMMENT '评论id',
  `content` longtext COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint NOT NULL COMMENT '作者',
  `like_count` int unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `comment_count` int unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_comment_id` (`comment_id`),
  KEY `idx_author_Id` (`author_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
INSERT INTO `comment` VALUES (3, 161392999815712768, '广东昨日新增2例本土确诊病例', 152408122030297088, 99, 10, '2021-06-22 16:38:10', '2021-06-22 16:38:09');
INSERT INTO `comment` VALUES (4, 161393385943339008, '广东昨日新增2例本土确诊病例', 152408122030297088, 0, 2, '2021-06-23 22:33:01', '2021-06-23 22:33:01');
INSERT INTO `comment` VALUES (7, 161466478581780480, '广东昨日新增2例本土确诊病例112222', 152408122030297088, 0, 0, '2021-06-22 21:30:08', '2021-06-22 21:30:08');
INSERT INTO `comment` VALUES (8, 161825635361099776, '神舟十二号航天员首次天地通话', 152408122030297088, 0, 0, '2021-06-23 21:17:18', '2021-06-23 21:17:18');
INSERT INTO `comment` VALUES (9, 161826788941500416, '此次神舟十二号任务，给航天员们配备了120余种航天食品，食谱按照航天员喜好定制，荤素搭配、营养均衡。祝航天员顺利完成任务，平安凯旋！', 152408122030297088, 0, 0, '2021-06-23 21:21:53', '2021-06-23 21:21:53');
INSERT INTO `comment` VALUES (10, 161872179791663104, '航天员的牛年公仔挂进核心舱', 152408122030297088, 0, 1, '2021-06-24 00:22:15', '2021-06-24 00:22:15');
COMMIT;

-- ----------------------------
-- Table structure for comment_rel
-- ----------------------------
DROP TABLE IF EXISTS `comment_rel`;
CREATE TABLE `comment_rel` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comment_id` bigint unsigned NOT NULL COMMENT '评论ID',
  `parent_id` bigint NOT NULL COMMENT '父ID',
  `level` int NOT NULL COMMENT '层次1:第一层2:第二层',
  `post_id` bigint NOT NULL COMMENT '帖子评论的id',
  `reply_author_id` bigint NOT NULL COMMENT '回复评论作者的id',
  `reply_comment_id` bigint NOT NULL DEFAULT '0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_post_id_level` (`post_id`,`level`),
  KEY `idx_parent_id_level` (`parent_id`,`level`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of comment_rel
-- ----------------------------
BEGIN;
INSERT INTO `comment_rel` VALUES (1, 161392999815712768, 23, 1, 152800143215693824, 152408122030297088, 0, '2021-06-22 16:38:09', '2021-06-22 16:38:09');
INSERT INTO `comment_rel` VALUES (2, 161393385943339008, 23, 1, 152800143215693824, 152408122030297088, 0, '2021-06-22 16:39:41', '2021-06-22 16:39:41');
INSERT INTO `comment_rel` VALUES (3, 161466478581780480, 161393385943339008, 2, 152800143215693824, 152408122030297088, 161393385943339008, '2021-06-22 21:30:08', '2021-06-22 21:30:08');
INSERT INTO `comment_rel` VALUES (4, 161825635361099776, 161466478581780480, 2, 152800143215693824, 152408122030297088, 161466478581780480, '2021-06-23 21:17:18', '2021-06-23 21:17:18');
INSERT INTO `comment_rel` VALUES (5, 161826788941500416, 0, 1, 152800143215693824, 152408122030297088, 0, '2021-06-23 21:21:53', '2021-06-23 21:21:53');
INSERT INTO `comment_rel` VALUES (6, 161872179791663104, 161825635361099776, 2, 152800143215693824, 152408122030297088, 161825635361099776, '2021-06-24 00:22:15', '2021-06-24 00:22:15');
COMMIT;

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` int NOT NULL AUTO_INCREMENT,
  `community_id` int unsigned NOT NULL,
  `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_community_id` (`community_id`),
  UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of community
-- ----------------------------
BEGIN;
INSERT INTO `community` VALUES (1, 1, 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO `community` VALUES (2, 2, 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO `community` VALUES (3, 3, 'PUBG', '大吉大利，今晚吃鸡。', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO `community` VALUES (4, 4, 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');
COMMIT;

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint NOT NULL COMMENT '帖子id',
  `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint NOT NULL COMMENT '作者的用户id',
  `community_id` bigint NOT NULL COMMENT '所属社区',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_post_id` (`post_id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of post
-- ----------------------------
BEGIN;
INSERT INTO `post` VALUES (12, 314641777497739264, 'mac os 命令切换独显 核显', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 2, 1, '2022-08-19 13:54:05', '2022-08-19 13:54:05');
INSERT INTO `post` VALUES (13, 314670397456912384, 'mac os 命令切换独显 核显1', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 2, 1, '2022-08-19 15:47:49', '2022-08-19 15:47:49');
INSERT INTO `post` VALUES (14, 314670418512318464, 'mac os 命令切换独显 核显12', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 2, 1, '2022-08-19 15:47:54', '2022-08-19 15:47:54');
INSERT INTO `post` VALUES (15, 314670430050848768, 'mac os 命令切换独显 核显123', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 2, 1, '2022-08-19 15:47:56', '2022-08-19 15:47:56');
INSERT INTO `post` VALUES (16, 314670448195407872, 'mac os 命令切换独显 核显1234', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 2, 1, '2022-08-19 15:48:01', '2022-08-19 15:48:01');
INSERT INTO `post` VALUES (17, 314670467401125888, 'mac os 命令切换独显 核显12345', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 2, 1, '2022-08-19 15:48:05', '2022-08-19 15:48:05');
INSERT INTO `post` VALUES (18, 314672992992890880, 'mac os 命令切换独显 核显6', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 1, 1, '2022-08-19 15:58:07', '2022-08-19 15:58:07');
INSERT INTO `post` VALUES (19, 314673018808832000, 'mac os 命令切换独显 核显7', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 1, 1, '2022-08-19 15:58:14', '2022-08-19 15:58:14');
INSERT INTO `post` VALUES (20, 314673042997383168, 'mac os 命令切换独显 核显8', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 1, 1, '2022-08-19 15:58:19', '2022-08-19 15:58:19');
INSERT INTO `post` VALUES (21, 314673072340733952, 'mac os 命令切换独显 核显9', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 3, 1, '2022-08-19 15:58:26', '2022-08-19 15:58:26');
INSERT INTO `post` VALUES (22, 314673093064790016, 'mac os 命令切换独显 核显10', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 3, 1, '2022-08-19 15:58:31', '2022-08-19 15:58:31');
INSERT INTO `post` VALUES (23, 314673111431647232, 'mac os 命令切换独显 核显11', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 3, 1, '2022-08-19 15:58:36', '2022-08-19 15:58:36');
INSERT INTO `post` VALUES (24, 314673128053673984, 'mac os 命令切换独显 核显12', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 3, 1, '2022-08-19 15:58:40', '2022-08-19 15:58:40');
INSERT INTO `post` VALUES (25, 314673166410584064, 'mac os 命令切换独显 核显13', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 4, 1, '2022-08-19 15:58:49', '2022-08-19 15:58:49');
INSERT INTO `post` VALUES (26, 314673182294413312, 'mac os 命令切换独显 核显15', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 4, 1, '2022-08-19 15:58:53', '2022-08-19 15:58:53');
INSERT INTO `post` VALUES (27, 314673209561583616, 'mac os 命令切换独显 核显14', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 4, 1, '2022-08-19 15:58:59', '2022-08-19 15:58:59');
INSERT INTO `post` VALUES (28, 314673233750134784, 'mac os 命令切换独显 核显16', ' 可以使用mac自带的 pmset 电源管理命令，在终端输入以下命令 sudo pmset -a GPUSwitch 0 sudo pmset -c GPUSwitch 1 命令含义：- c 表示 charger 为电源模式，-b 表示 battery 为电池模式0 表示用核显，1 表示用独显，2表示自动此命令可能会有不隐定的情况，重启或是玩游戏时会重新启用独立显卡。除此方法外，可以下载下面的软件，可在菜单栏随时进行切换。', 152408122030297088, 4, 1, '2022-08-19 15:59:05', '2022-08-19 15:59:05');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `gender` tinyint NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 28018727488323585, 'q1mi', '313233343536639a9119599647d841b1bef6ce5ea293', '100345@qq.com', 0, '2020-07-12 07:01:03', '2021-06-23 19:29:23');
INSERT INTO `user` VALUES (2, 4183532125556736, '七米', '313233639a9119599647d841b1bef6ce5ea293', '100345@qq.com', 0, '2020-07-12 13:03:51', '2021-06-23 19:29:23');
INSERT INTO `user` VALUES (3, 152347055732297728, 'admin11', 'e10adc3949ba59abbe56e057f20f883e', '100345@qq.com', 0, '2021-05-28 17:32:48', '2021-06-23 19:29:23');
INSERT INTO `user` VALUES (4, 152408122030297088, 'admin113', 'e10adc3949ba59abbe56e057f20f883e', '100345@qq.com', 0, '2021-05-28 21:35:27', '2021-06-23 19:29:23');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
