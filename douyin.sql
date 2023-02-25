/*
 Navicat Premium Data Transfer

 Source Server         : douyin
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : 43.139.147.169:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 04/02/2023 15:57:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `comment_uuid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '评论uuid',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '评论作者id',
  `video_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '评论视频id',
  `contents` varchar(255) NOT NULL DEFAULT '' COMMENT '评论内容',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '自设创建时间(unix)',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除',
  PRIMARY KEY (`id`),
  KEY `fk_user_comment` (`user_id`),
  KEY `fk_video_comment` (`video_id`),
  KEY `fk_uuid_comment` (`comment_uuid`),
  CONSTRAINT `fk_user_comment` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_video_comment` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`),
  UNIQUE KEY `uni_comment_uuid` (`comment_uuid`) COMMENT 'uuid需要唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `video_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '视频id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '软删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_favorite` (`user_id`,`video_id`,`deleted_at`),
  KEY `fk_user_favorite` (`user_id`),
  KEY `fk_video_favorite` (`video_id`),
  CONSTRAINT `fk_user_favorite` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_video_favorite` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='点赞表';

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '消息id',
  `message_uuid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '消息uuid',
  `to_user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '该消息接收者的id',
  `from_user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '该消息发送者的id',
  `contents` varchar(255) NOT NULL DEFAULT '' COMMENT '消息内容',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '自设创建时间(unix)',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除',
  PRIMARY KEY (`id`),
  KEY `fk_user_message_to` (`to_user_id`),
  KEY `fk_user_message_from` (`from_user_id`),
  KEY `fk_uuid_message` (`message_uuid`),
  CONSTRAINT `fk_user_message_from` FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_user_message_to` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`),
  UNIQUE KEY `uni_message_uuid` (`message_uuid`) COMMENT 'uuid需要唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息表';

-- ----------------------------
-- Table structure for relation
-- ----------------------------
DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `to_user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '关注目标的用户id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_relation` (`user_id`,`to_user_id`,`deleted_at`),
  KEY `fk_user_relation` (`user_id`),
  KEY `fk_user_relation_to` (`to_user_id`),
  CONSTRAINT `fk_user_relation` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_user_relation_to` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='关注表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码，已加密',
  `follow_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '关注人数',
  `follower_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '粉丝人数',
  `work_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '作品数',
  `favorite_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '点赞视频数',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_name` (`name`) COMMENT '用户名称需要唯一'
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'user表主键',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '视频标题',
  `play_url` varchar(128) NOT NULL DEFAULT '' COMMENT '视频地址',
  `cover_url` varchar(128) NOT NULL DEFAULT '' COMMENT '封面地址',
  `favorite_count` int(15) unsigned NOT NULL DEFAULT '0' COMMENT '获赞数量',
  `comment_count` int(15) unsigned NOT NULL DEFAULT '0' COMMENT '评论数量',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除',
  PRIMARY KEY (`id`),
  KEY `fk_user_video` (`user_id`),
  CONSTRAINT `fk_user_video` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

SET FOREIGN_KEY_CHECKS = 1;
