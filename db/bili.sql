/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : demo

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 29/09/2020 17:38:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bili
-- ----------------------------
DROP TABLE IF EXISTS `bili`;
CREATE TABLE `bili` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `author_id` varchar(255) DEFAULT NULL,
  `video_page` varchar(255) DEFAULT NULL,
  `video_url` varchar(255) DEFAULT '',
  `pic` varchar(500) DEFAULT NULL,
  `tags` varchar(500) DEFAULT NULL,
  `description` text,
  `created` datetime DEFAULT NULL,
  `Updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `video_page` (`video_page`)
) ENGINE=InnoDB AUTO_INCREMENT=401 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;