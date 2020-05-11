/*
 Navicat Premium Data Transfer

 Source Server         : 本地测试8
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : localhost:3306
 Source Schema         : helloworld

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 09/05/2020 18:41:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for channel_roles
-- ----------------------------
DROP TABLE IF EXISTS `channel_roles`;
CREATE TABLE `channel_roles`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `channel_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `power` smallint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '能力\r\n1',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务器角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for channel_users
-- ----------------------------
DROP TABLE IF EXISTS `channel_users`;
CREATE TABLE `channel_users`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `channel_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `user_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '加入事件',
  `open_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开放ID',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `union_unique_channel_id_open_id`(`channel_id`, `open_id`) USING BTREE COMMENT '每隔频道的openid保持唯一',
  UNIQUE INDEX `union_unique_channel_id_user_id`(`channel_id`, `user_id`) USING BTREE COMMENT '每个频道用户保持唯一',
  INDEX `fk_server_channel_users_servers_1`(`channel_id`) USING BTREE,
  INDEX `fk_server_channel_users_users_1`(`user_id`) USING BTREE,
  CONSTRAINT `fk_channel_users_channels_1` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_server_channel_users_users_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务器用户列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for channels
-- ----------------------------
DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `server_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '服务器id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `avatar` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图标',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `create_user_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建用户',
  `owner_user_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拥有者',
  `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '1 是否经过验证\r\n2 是否公开\r\n4 是否可商用\r\n8 是否关闭 - 任何人不可访问',
  `parent_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID，0则表示主频道',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_channels_users_1`(`create_user_id`) USING BTREE,
  INDEX `fk_channels_users_2`(`owner_user_id`) USING BTREE,
  CONSTRAINT `fk_channels_users_1` FOREIGN KEY (`create_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_channels_users_2` FOREIGN KEY (`owner_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '频道列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for server_roles
-- ----------------------------
DROP TABLE IF EXISTS `server_roles`;
CREATE TABLE `server_roles`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `server_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `power` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '能力',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_server_roles_servers_1`(`server_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务器角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for server_users
-- ----------------------------
DROP TABLE IF EXISTS `server_users`;
CREATE TABLE `server_users`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `server_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `user_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '加入事件',
  `open_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开放ID',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `union_unique_server_id_open_id`(`server_id`, `open_id`) USING BTREE COMMENT '服务器的openid保持唯一',
  UNIQUE INDEX `union_unique_server_id_user_id`(`server_id`, `user_id`) USING BTREE COMMENT '每隔服务器和用户保持唯一',
  INDEX `fk_server_users_servers_1`(`server_id`) USING BTREE,
  INDEX `fk_server_users_users_1`(`user_id`) USING BTREE,
  CONSTRAINT `fk_server_users_servers_1` FOREIGN KEY (`server_id`) REFERENCES `servers` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_server_users_users_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务器用户列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for servers
-- ----------------------------
DROP TABLE IF EXISTS `servers`;
CREATE TABLE `servers`  (
  `id` bigint(0) UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0.0.1',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '区域',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务器列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_accounts
-- ----------------------------
DROP TABLE IF EXISTS `user_accounts`;
CREATE TABLE `user_accounts`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `type` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '账户类型\r\n0 邮箱\r\n1 手机号码\r\n2',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_user_accounts_users_1`(`user_id`) USING BTREE,
  CONSTRAINT `fk_user_accounts_users_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户账户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_contact_black
-- ----------------------------
DROP TABLE IF EXISTS `user_contact_black`;
CREATE TABLE `user_contact_black`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(0) UNSIGNED NOT NULL COMMENT '用户ID',
  `contact_id` bigint(0) UNSIGNED NOT NULL COMMENT '对方id',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_contact_request
-- ----------------------------
DROP TABLE IF EXISTS `user_contact_request`;
CREATE TABLE `user_contact_request`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(0) UNSIGNED NOT NULL COMMENT '请求方',
  `contact_id` bigint(0) UNSIGNED NOT NULL COMMENT '被请求方',
  `user_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '描述内容',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '时间',
  `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态\r\n0 新请求\r\n1 通过\r\n2 拒绝\r\n3 忽略\r\n4 需求内容',
  `contact_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '被请求方返回说明',
  `update_time` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '好友请求' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_contacts
-- ----------------------------
DROP TABLE IF EXISTS `user_contacts`;
CREATE TABLE `user_contacts`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `contact_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '联系人的user_id',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_user_contacts_users_3`(`user_id`) USING BTREE,
  INDEX `fk_user_contacts_users_4`(`contact_id`) USING BTREE,
  CONSTRAINT `fk_user_contacts_users_3` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_user_contacts_users_4` FOREIGN KEY (`contact_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '联系人' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_passwords
-- ----------------------------
DROP TABLE IF EXISTS `user_passwords`;
CREATE TABLE `user_passwords`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `type` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 登陆密码\r\n1 支付密码',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_user_passwords_users_1`(`user_id`) USING BTREE,
  CONSTRAINT `fk_user_passwords_users_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户密码表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_tokens
-- ----------------------------
DROP TABLE IF EXISTS `user_tokens`;
CREATE TABLE `user_tokens`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'token令牌',
  `expire_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '过期时间戳',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间戳',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户令牌' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `language` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '语言',
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

-- 变更字段名称 2020年5月11日09:57:10
ALTER TABLE `user_contacts`
    CHANGE COLUMN `remarks` `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注';

-- 黑名单表增加备注字段 2020年5月11日09:59:17
ALTER TABLE `user_contact_black`
    ADD COLUMN `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '拉黑描述';
