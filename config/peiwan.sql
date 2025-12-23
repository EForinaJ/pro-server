/*
Navicat MySQL Data Transfer

Source Server         : shrimp
Source Server Version : 50744
Source Host           : localhost:3306
Source Database       : peiwan

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2025-12-24 03:18:21
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_balance
-- ----------------------------
DROP TABLE IF EXISTS `sys_balance`;
CREATE TABLE `sys_balance` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `related` varchar(255) DEFAULT NULL,
  `after` decimal(10,2) DEFAULT NULL,
  `before` decimal(10,2) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `mode` tinyint(4) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_balance
-- ----------------------------
INSERT INTO `sys_balance` VALUES ('63', '23', '系统增加余额', '300.00', '1300.00', '1000.00', '1', '1', 'asd ', '2025-12-08 16:29:52');
INSERT INTO `sys_balance` VALUES ('65', '23', 'CZ202512901541330020729', '1400.00', '1450.00', '50.00', '1', '2', '用户充值余额', '2025-12-08 17:54:14');

-- ----------------------------
-- Table structure for sys_casbin
-- ----------------------------
DROP TABLE IF EXISTS `sys_casbin`;
CREATE TABLE `sys_casbin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `p_type` varchar(64) DEFAULT NULL,
  `v0` varchar(256) DEFAULT NULL,
  `v1` varchar(256) DEFAULT NULL,
  `v2` varchar(256) DEFAULT NULL,
  `v3` varchar(256) DEFAULT NULL,
  `v4` varchar(256) DEFAULT NULL,
  `v5` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18062 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='管理员_casbin权限表';

-- ----------------------------
-- Records of sys_casbin
-- ----------------------------
INSERT INTO `sys_casbin` VALUES ('1', 'p', 'R_TEST', '/menu/all', '', '', '', '');
INSERT INTO `sys_casbin` VALUES ('2', 'p', 'R_TEST', '/menu/edit', '', '', '', '');
INSERT INTO `sys_casbin` VALUES ('25', 'g', '4', 'R_TEST', '', '', '', '');
INSERT INTO `sys_casbin` VALUES ('26', 'g', '5', 'R_TEST', '', '', '', '');
INSERT INTO `sys_casbin` VALUES ('27', 'g', '4', 'R_TEST_TWO', '', '', '', '');
INSERT INTO `sys_casbin` VALUES ('28', 'g', '5', 'R_TEST_TWO', '', '', '', '');
INSERT INTO `sys_casbin` VALUES ('11694', 'g', '1', 'ADMIN', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17919', 'p', 'ADMIN', '/menu/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17920', 'p', 'ADMIN', '/menu', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17921', 'p', 'ADMIN', '/menu/all', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17922', 'p', 'ADMIN', '/menu/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17923', 'p', 'ADMIN', '/menu/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17924', 'p', 'ADMIN', '/role', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17925', 'p', 'ADMIN', '/role/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17926', 'p', 'ADMIN', '/role/all/menu', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17927', 'p', 'ADMIN', '/role/all/permission', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17928', 'p', 'ADMIN', '/role/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17929', 'p', 'ADMIN', '/role/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17930', 'p', 'ADMIN', '/role/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17931', 'p', 'ADMIN', '/manage', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17932', 'p', 'ADMIN', '/manage/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17933', 'p', 'ADMIN', '/manage/all/role', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17934', 'p', 'ADMIN', '/manage/role', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17935', 'p', 'ADMIN', '/manage/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17936', 'p', 'ADMIN', '/manage/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17937', 'p', 'ADMIN', '/manage/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17938', 'p', 'ADMIN', '/permission', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17939', 'p', 'ADMIN', '/permission/all', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17940', 'p', 'ADMIN', '/permission/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17941', 'p', 'ADMIN', '/permission/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17942', 'p', 'ADMIN', '/permission/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17943', 'p', 'ADMIN', '/system', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17944', 'p', 'ADMIN', '/system/file', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17945', 'p', 'ADMIN', '/upload', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17946', 'p', 'ADMIN', '/upload/file', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17947', 'p', 'ADMIN', '/upload/chunkIdentifier', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17948', 'p', 'ADMIN', '/upload/chunk', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17949', 'p', 'ADMIN', '/upload/chunk/merge', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17950', 'p', 'ADMIN', '/system/base', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17951', 'p', 'ADMIN', '/system/email', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17952', 'p', 'ADMIN', '/role/permissions', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17953', 'p', 'ADMIN', '/role/menus', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17954', 'p', 'ADMIN', '/role/edit/menus', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17955', 'p', 'ADMIN', '/role/edit/permissions', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17956', 'p', 'ADMIN', '/media', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17957', 'p', 'ADMIN', '/media/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17958', 'p', 'ADMIN', '/media/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17959', 'p', 'ADMIN', '/user', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17960', 'p', 'ADMIN', '/user/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17961', 'p', 'ADMIN', '/level', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17962', 'p', 'ADMIN', '/level/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17963', 'p', 'ADMIN', '/level/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17964', 'p', 'ADMIN', '/level/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17965', 'p', 'ADMIN', '/level/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17966', 'p', 'ADMIN', '/vip', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17967', 'p', 'ADMIN', '/vip/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17968', 'p', 'ADMIN', '/vip/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17969', 'p', 'ADMIN', '/vip/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17970', 'p', 'ADMIN', '/vip/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17971', 'p', 'ADMIN', '/game', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17972', 'p', 'ADMIN', '/game/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17973', 'p', 'ADMIN', '/game/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17974', 'p', 'ADMIN', '/game/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17975', 'p', 'ADMIN', '/game/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17976', 'p', 'ADMIN', '/title', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17977', 'p', 'ADMIN', '/title/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17978', 'p', 'ADMIN', '/title/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17979', 'p', 'ADMIN', '/title/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17980', 'p', 'ADMIN', '/title/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17981', 'p', 'ADMIN', '/product', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17982', 'p', 'ADMIN', '/product/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17983', 'p', 'ADMIN', '/product/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17984', 'p', 'ADMIN', '/product/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17985', 'p', 'ADMIN', '/product/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17986', 'p', 'ADMIN', '/witkey', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17987', 'p', 'ADMIN', '/witkey/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17988', 'p', 'ADMIN', '/witkey/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17989', 'p', 'ADMIN', '/witkey/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17990', 'p', 'ADMIN', '/witkey/punish/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17991', 'p', 'ADMIN', '/witkey/punish/revoke', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17992', 'p', 'ADMIN', '/withdraw', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17993', 'p', 'ADMIN', '/withdraw/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17994', 'p', 'ADMIN', '/withdraw/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17995', 'p', 'ADMIN', '/withdraw/apply', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17996', 'p', 'ADMIN', '/withdraw/detail', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17997', 'p', 'ADMIN', '/system/withdraw', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17998', 'p', 'ADMIN', '/order', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('17999', 'p', 'ADMIN', '/order/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18000', 'p', 'ADMIN', '/order/detail', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18001', 'p', 'ADMIN', '/order/refund', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18002', 'p', 'ADMIN', '/order/add/discount', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18003', 'p', 'ADMIN', '/dict', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18004', 'p', 'ADMIN', '/dict/type/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18005', 'p', 'ADMIN', '/dict/type/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18006', 'p', 'ADMIN', '/dict/type/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18007', 'p', 'ADMIN', '/dict/type/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18008', 'p', 'ADMIN', '/dict/data/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18009', 'p', 'ADMIN', '/dict/data/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18010', 'p', 'ADMIN', '/dict/data/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18011', 'p', 'ADMIN', '/dict/data/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18012', 'p', 'ADMIN', '/order/logs', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18013', 'p', 'ADMIN', '/user/detail', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18014', 'p', 'ADMIN', '/user/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18015', 'p', 'ADMIN', '/user/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18016', 'p', 'ADMIN', '/user/change/balance', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18017', 'p', 'ADMIN', '/user/balance/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18018', 'p', 'ADMIN', '/user/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18019', 'p', 'ADMIN', '/witkey/detail', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18020', 'p', 'ADMIN', '/witkey/punish', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18021', 'p', 'ADMIN', '/user/logs', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18022', 'p', 'ADMIN', '/system/wechat/mini/program', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18023', 'p', 'ADMIN', '/system/user', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18024', 'p', 'ADMIN', '/order/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18025', 'p', 'ADMIN', '/order/paid', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18026', 'p', 'ADMIN', '/user/recharge', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18027', 'p', 'ADMIN', '/user/recharge/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18028', 'p', 'ADMIN', '/order/cancel', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18029', 'p', 'ADMIN', '/witkey/logs', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18030', 'p', 'ADMIN', '/witkey/commission/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18031', 'p', 'ADMIN', '/witkey/change/commission', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18032', 'p', 'ADMIN', '/witkey/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18033', 'p', 'ADMIN', '/recharge', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18034', 'p', 'ADMIN', '/recharge/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18035', 'p', 'ADMIN', '/recharge/revoke', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18036', 'p', 'ADMIN', '/system/contact', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18037', 'p', 'ADMIN', '/game/all', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18038', 'p', 'ADMIN', '/user/select/options', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18039', 'p', 'ADMIN', '/title/all', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18040', 'p', 'ADMIN', '/dict/data/all', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18041', 'p', 'ADMIN', '/user/add/witkey', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18042', 'p', 'ADMIN', '/user/edit/witkey', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18043', 'p', 'ADMIN', '/user/change/deposit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18044', 'p', 'ADMIN', '/user/witkey/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18045', 'p', 'ADMIN', '/user/delete/witkey', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18046', 'p', 'ADMIN', '/bill', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18047', 'p', 'ADMIN', '/bill/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18048', 'p', 'ADMIN', '/level/all', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18049', 'p', 'ADMIN', '/system/wechat/pay', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18050', 'p', 'ADMIN', '/system/user/agreement', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18051', 'p', 'ADMIN', '/system/privacy/policy', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18052', 'p', 'ADMIN', '/system/about/us', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18053', 'p', 'ADMIN', '/category', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18054', 'p', 'ADMIN', '/category/list', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18055', 'p', 'ADMIN', '/category/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18056', 'p', 'ADMIN', '/category/edit', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18057', 'p', 'ADMIN', '/category/delete', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18058', 'p', 'ADMIN', '/distribute', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18059', 'p', 'ADMIN', '/distribute/create', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18060', 'p', 'ADMIN', '/distribute/cancel', null, null, null, null);
INSERT INTO `sys_casbin` VALUES ('18061', 'p', 'ADMIN', '/distribute/list', null, null, null, null);

-- ----------------------------
-- Table structure for sys_category
-- ----------------------------
DROP TABLE IF EXISTS `sys_category`;
CREATE TABLE `sys_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `game_id` bigint(20) DEFAULT NULL,
  `pid` bigint(20) DEFAULT '0' COMMENT '顶级分类',
  `name` varchar(50) DEFAULT NULL COMMENT '分类名称',
  `pic` varchar(255) DEFAULT NULL COMMENT '分类背景图',
  `sort` int(10) DEFAULT NULL COMMENT '分类排序',
  `description` varchar(255) DEFAULT NULL COMMENT '分类描述',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime DEFAULT NULL COMMENT '更新日期',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_category
-- ----------------------------
INSERT INTO `sys_category` VALUES ('38', '38', '0', '三角洲行动', 'https://q8.itc.cn/q_70/images03/20241017/b3ff223261b14b18a3ed2cdadec148f7.png', '1', 'gsdgf', '2025-01-19 17:50:20', '2025-12-14 16:31:22');
INSERT INTO `sys_category` VALUES ('39', '39', '0', 'APEX', '/public/uploads/2025-12-14/dey0pr0nv2ywcbtf2w.jpg', '2', 'fsdf', '2025-04-05 09:19:16', '2025-12-14 16:31:19');

-- ----------------------------
-- Table structure for sys_comment
-- ----------------------------
DROP TABLE IF EXISTS `sys_comment`;
CREATE TABLE `sys_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `order_id` bigint(20) DEFAULT NULL,
  `product_id` bigint(20) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `reply_user_id` bigint(20) DEFAULT NULL,
  `content` text,
  `rating` tinyint(4) DEFAULT NULL,
  `images` text,
  `is_top` tinyint(4) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `ip` varchar(64) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`),
  KEY `idx_order` (`order_id`),
  KEY `idx_product` (`product_id`),
  KEY `idx_parent` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_comment
-- ----------------------------
INSERT INTO `sys_comment` VALUES ('1', '15', '7', '28', null, null, 'sdfgasdfasdf', '5', '[\"/public/uploads/2025-11-04/de003c89q2cgky9pyu.jpg\"]', null, '1', '::1', '2025-11-04 15:03:51', '2025-11-04 15:03:51');

-- ----------------------------
-- Table structure for sys_commission
-- ----------------------------
DROP TABLE IF EXISTS `sys_commission`;
CREATE TABLE `sys_commission` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `witkey_id` bigint(20) NOT NULL,
  `related` varchar(255) DEFAULT NULL,
  `after` decimal(10,2) DEFAULT NULL,
  `before` decimal(10,2) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `mode` tinyint(4) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`witkey_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_commission
-- ----------------------------
INSERT INTO `sys_commission` VALUES ('66', '17', '系统减少余额', null, '-100.00', '100.00', '2', '1', 'hf sdf ', '2025-12-11 10:47:33');
INSERT INTO `sys_commission` VALUES ('67', '17', '系统增加余额', '0.00', '100.00', '100.00', '1', '1', 'gfh', '2025-12-11 10:49:07');

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `id` int(5) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `name` varchar(100) DEFAULT '' COMMENT '参数名称',
  `key` varchar(100) DEFAULT '' COMMENT '参数键名',
  `value` text COMMENT '参数键值',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='参数配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES ('8', '文件设置', 'FileSetting', '{\"engine\":\"local\",\"path\":\"uploads\",\"fileSize\":50,\"fileType\":[\".crt\"],\"imageSize\":2,\"imageType\":[\".png\",\".jpg\",\".jpeg\"]}', '2020-05-30 22:47:11', '2024-10-05 20:19:51', '文件设置');
INSERT INTO `sys_config` VALUES ('9', '支付宝配置', 'AlyPayOptions', '{\"appId\":\"2016102700770129\",\"privateKey\":\"MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCjkmnBSEjwpvBFLUF279MLV1bSjjnnokhocdEgTYx4Qy6zhtpUfyV/LwRQ31ddUtoTYhCedOO5nd95vMrAOSU2++ZdXWVjUfw1er4fFaMlNp8hQkKOgp8bDLO05x3/o9rG0oOWqWreyUGoPoWRQhnzJt6kYbJtunx3d8hBN59jM0PHiDODDsgiu8sq4muk/BjayF8FKypkp2LQCNjpxbUPIcbjpLsRzrJ4eH+Om+10cMHpIDLl6JZITwq3JF6l39cS61N4xNwOOtE/1KCo3x9/y8T2XF3PQvfHjXHeNuwQNIvey828a7KFbTgV6O09lOCTvjP1HmA1zvpHD6MizaofAgMBAAECggEAfX9WGG9HVywd6FVihshWGbt257EriCage1Hn62rUPWj+KctrM60hrcT7ALl6pCVvH7P7oDd6iO0xioto27Z8cQUvp66CnYNHiBiWSe8l7uVLjg7yVbiuLei+8CrqfzrOHgFh6HQvhKLQ9y6Q9/PJSR8nbuNuLHYKDPcf87mjVk3Cz6D03V9B1avfaH5IiLvlZwrAgVk7FoZEoBnhZhIg7yME4DUitPykTGg1ur/Gofe+xrPqONOnvZW8wNa/l2nKIPQwBaiUVcSV7gVBRRkw/GADMV59titx1dyzmR0b61GLexjI2WxXRLnAe1+qXrbOcbsCKIgFjQELbCz8MYFYgQKBgQDM7m/rzIcsJkmHXFGZdVRIdFzx73SUKMSu3srCs+7urUG9HNupf16ykMHv1p7rA0tBt6lzTAUgpf0F959JE1fUl+tyy2gUDyqOrP1FzF91wbertTXcrBcv9wMD5mvENSHNz9CDQ3Z7ZHtpPSnk2AvGhXzDSc6pnEkBzUIIVjgQfwKBgQDMVXSsnESsUwzz9qfzWxT9qsELIMM91xyn/w7FWcvIZvd33TzpPP0aZ69hylmaUPpSp4wwL5+zqPOvfI03y5En7ZzohMBTebU4H4m3aK4GHF0w+6ft+bt04ZAyFc7lyf/w+nEniPgQHSQqMu64FbK+GrYxrAXzSFhHfH/b35uWYQKBgGYpK3RSdsRkpd0sAaXN3uFr2PXnGKfPlxVSDaR4jNFBX/dFzp+11mCQV44X4QtpjffJ9lh6+kdnWDbEVgzY7r0VqxOEIXN2iBGuXWiRVLlghA6+fIZw5/JKYp9sHCcpEZwHUHxPgl5LHla9XgguR9iErUixn6vgNGkIiTWcvcBDAoGAfUGNShppBnHKqOp0vfsBfRZlS9sDlC7/RARYG6YWA30LChE2u4tFZCBXJE0UbEJjkLNgflFTRqC08MgbES7ahm1kGCz4cLNU4ViD5UhoFRriDZrWsEy8GsQCzpELyVTwbdo37xJJbidO+gdKytGSRnK9aOmYpC+e3gN1pWUHTUECgYA0zjkcEHfYzxH/VChkT9SzuI4WIJ+NSQ5ztljaZlDPg8rskUCB/kV6tRXW5oVBmgm4VyHl5bOl8PkoUC06Zhe6jiZ5cSUnfvCi7w2ysc1y5Nypc1P/qBu0bGK0nLRdFHbuv4IZsCrcvaR86QFeK14xO+yJ3kkOyXLd4UTvXYG/xg==\",\"appPublicKey\":\"http://beethorn.com/public/uploads/2022-07-03/cl6446516yqk8s03bl.crt\",\"alyCertPublicKey\":\"http://beethorn.com/public/uploads/2022-07-03/cl6448p2rn6cs8ppzr.crt\",\"alyRootCert\":\"http://beethorn.com/public/uploads/2022-07-03/cl6449uckjbwua78na.crt\"}', '2020-06-13 15:32:17', '2022-07-03 23:05:54', '支付宝配置');
INSERT INTO `sys_config` VALUES ('10', '提现设置', 'WithdrawSetting', '{\"symbol\":\"金币\",\"minWithdraw\":50,\"withDrawPercent\":1}', '2020-06-13 16:37:21', '2024-10-05 20:54:56', '支付设置');
INSERT INTO `sys_config` VALUES ('12', '用户设置', 'UserSetting', '{\"avatar\":\"/public/uploads/2025-12-05/depm9pfjzbs4wyxjpn.jpg\",\"levelId\":\"1\",\"checkInExperience\":\"500\"}', '2020-10-25 17:25:12', '2024-10-07 21:57:26', '用户设置');
INSERT INTO `sys_config` VALUES ('24', '通知设置', 'NoticeSetting', '{\"register\":\"{siteTitle}欢迎您的加入\",\"create\":\"您发布《{title}》,已经审核: {reason}\",\"remove\":\"您发布的内容{title}，被删除了，{reason}\",\"groupCreate\":\"您创建《{title}》,已经审核: {reason}\",\"groupRemove\":\"您创建{title}，被删除了: {reason}\",\"report\":\"您举报的内容，已处理，{reason}\",\"userProhibit\":\"{reason}\",\"verify\":\"您实名认证已审核，{reason}\",\"cash\":\"您编号为{code}提现申请已经打款，\"}', '2021-05-16 02:07:09', '2021-06-21 00:44:05', '通知设置');
INSERT INTO `sys_config` VALUES ('27', '阿里云oss配置', 'AlyOssOption', '{\"roleArn\":\"acs:ram::1986051699684916:role/ramosstest\",\"roleSessionName\":\"beetorn\",\"regionId\":\"cn-shenzhen\",\"endpoint\":\"oss-cn-shenzhen.aliyuncs.com\",\"accessKeyId\":\"LTAI5t8RPwZu2SL4Rm7wKZSe\",\"accessKeySecret\":\"d4r3CXZJr1DIusaX4R2JTaViXAh42P\",\"bucketName\":\"beethorn\"}', '2021-07-02 21:38:09', '2022-06-27 18:41:49', '阿里云oss配置');
INSERT INTO `sys_config` VALUES ('28', '七牛oss配置', 'QiNiuOssOption', '{\"endpoint\":\"21312\",\"accessKeyId\":\"312312\",\"accessKeySecret\":\"3123\",\"bucketName\":\"123\",\"address\":\"213123\"}', '2021-07-02 22:17:48', '2021-07-02 22:17:48', '七牛oss配置');
INSERT INTO `sys_config` VALUES ('29', '阿里云短信配置', 'AlySmsOptions', '{\"id\":\"LTAIyOnBE3wGtyAA\",\"secret\":\"PlwvkrExsup3XccmMxaonmJe4HshK3\",\"publicKey\":null,\"appPublicKey\":null,\"alyCertPublicKey\":null,\"alyRootCert\":null}', '2021-11-28 17:05:20', '2021-11-28 17:05:27', '阿里云短信配置');
INSERT INTO `sys_config` VALUES ('32', '积分设置', 'IntegralSetting', '{\"moneyToIntegral\":50,\"signIn\":\"100-1000\",\"register\":200,\"likeFavorite\":\"2-50\",\"follow\":\"2-50\",\"report\":\"2-50\",\"recommend\":\"2-50\",\"comment\":\"2-50\",\"content\":\"2-50\",\"like\":\"2-50\",\"favorite\":\"2-50\"}', '2024-03-19 13:21:21', '2024-10-06 11:02:36', '积分设置');
INSERT INTO `sys_config` VALUES ('34', '经验设置', 'ExperienceSetting', '{\"signIn\":\"50-500\",\"register\":200,\"like\":\"2-50\",\"favorite\":\"2-50\",\"follow\":\"2-50\",\"report\":\"2-50\",\"recommend\":\"2-50\",\"comment\":\"2-50\",\"content\":\"2-50\"}', '2024-09-25 06:30:17', '2024-10-06 10:54:10', '');
INSERT INTO `sys_config` VALUES ('36', '邮箱设置', 'EmailSetting', '{\"host\":\"fdgsdfgsdfgsdfg\",\"port\":1231,\"user\":\"sdfsadfsadf\",\"email\":\"sadfsdf\",\"pass\":\"sdf\",\"minSendTime\":2234234,\"ipMaxSendCount\":123,\"codeEndTime\":123}', null, null, null);
INSERT INTO `sys_config` VALUES ('38', '联系配置', 'ContactSetting', '{\"platform\":\"KOOK\",\"number\":\"245331255\",\"icon\":\"/resource/public/uploads/2025-09-15/dcspvcg1drmkert7a5.jpg\",\"onlineTime\":\"12:00~24:00\",\"wechat\":\"/resource/public/uploads/2025-09-15/dcspvcg1drmkert7a5.jpg\"}', null, null, null);
INSERT INTO `sys_config` VALUES ('39', '微信小程序设置', 'WechatMiniProgramSetting', '{\"appId\":\"wxd135bf2f0bd7cc8f\",\"secret\":\"9d7dc2c24ae036da82bea9c73d2d62d7\"}', null, null, null);
INSERT INTO `sys_config` VALUES ('40', '微信支付设置', 'WechatPaySetting', '{\"mchId\":\"1732225763\",\"mchKey\":\"e1997b9f1bed1b73e45c4d6100a7b1fe\",\"serialNo\":\"61EFB70335A38B5BD50FF014E58B16343D570366\",\"apiclientCert\":\"/public/uploads/2025-11-12/de6s099hs8u8h63vep.pem\",\"apiclientKey\":\"/public/uploads/2025-11-12/de6s0af1f87sefa8to.pem\",\"pubKey\":\"/public/uploads/2025-11-13/de75b24q9cfwvio3ub.pem\",\"pubId\":\"PUB_KEY_ID_0117322257632025111200291500003005\"}', null, null, null);
INSERT INTO `sys_config` VALUES ('42', '用户协议', 'UserAgreement', '<p>黄金矿工回家看sdgsdfg</p>', '2025-12-05 11:02:54', '2025-12-05 11:07:23', null);
INSERT INTO `sys_config` VALUES ('44', '基本设置', 'BaseSetting', '{\"title\":\"Pro玩家\",\"domain\":\"prowanjia.com\",\"logo\":\"/public/uploads/2025-12-05/deq8fmqgeqic4msvl4.jpg\",\"icon\":\"/public/uploads/2025-12-05/deq8fnfgrgosrsbqdy.jpg\",\"description\":\"dfsdf\",\"icp\":\"sdfsdfs\"}', '2025-12-05 11:05:26', null, null);
INSERT INTO `sys_config` VALUES ('45', '隐私协议', 'PrivacyPolicy', '<p>山豆根发射点给</p>', '2025-12-05 12:13:37', null, null);
INSERT INTO `sys_config` VALUES ('46', '关于我们', 'AboutUs', '<p>sadfsadfasdfsa</p>', '2025-12-06 15:59:57', null, null);

-- ----------------------------
-- Table structure for sys_credential
-- ----------------------------
DROP TABLE IF EXISTS `sys_credential`;
CREATE TABLE `sys_credential` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `credential_type` tinyint(4) DEFAULT NULL,
  `credential_name` varchar(100) DEFAULT NULL,
  `credential_number` varchar(100) DEFAULT NULL,
  `credential_front_image` varchar(500) DEFAULT NULL,
  `credential_against_image` varchar(500) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_credential
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(50) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `key` varchar(100) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES ('10', 'PRODUCT_TYPE', '护航', 'escort', '1', '2025-10-11 16:32:46', '2025-10-11 16:32:46');
INSERT INTO `sys_dict_data` VALUES ('11', 'PRODUCT_TYPE', '代练', 'power_leveling', '2', '2025-10-11 16:33:31', '2025-10-11 16:33:31');
INSERT INTO `sys_dict_data` VALUES ('12', 'RECHARGE_TYPE', '支付宝', 'alypay', '1', '2025-10-27 12:06:02', '2025-10-27 12:06:02');
INSERT INTO `sys_dict_data` VALUES ('13', 'RECHARGE_TYPE', '微信', 'wechat', '2', '2025-10-27 12:06:23', '2025-10-27 12:06:23');
INSERT INTO `sys_dict_data` VALUES ('14', 'RECHARGE_TYPE', '人工充值', 'manual', '3', '2025-10-27 12:07:26', '2025-11-03 20:28:42');
INSERT INTO `sys_dict_data` VALUES ('15', 'PAY_TYPE', '支付宝', 'alyPay', '1', '2025-10-27 12:24:38', '2025-10-27 12:24:38');
INSERT INTO `sys_dict_data` VALUES ('16', 'PAY_TYPE', '微信', 'wechat', '2', '2025-10-27 12:24:51', '2025-10-27 12:24:51');
INSERT INTO `sys_dict_data` VALUES ('17', 'RECHARGE_MONEY', '20', '20', '20', '2025-10-28 11:10:30', '2025-10-28 11:10:30');
INSERT INTO `sys_dict_data` VALUES ('18', 'RECHARGE_MONEY', '50', '50', '50', '2025-10-28 11:10:36', '2025-10-28 11:10:36');
INSERT INTO `sys_dict_data` VALUES ('19', 'RECHARGE_MONEY', '100', '100', '100', '2025-10-28 11:10:43', '2025-10-28 11:10:43');
INSERT INTO `sys_dict_data` VALUES ('20', 'RECHARGE_MONEY', '200', '200', '200', '2025-10-28 11:10:49', '2025-10-28 11:10:49');
INSERT INTO `sys_dict_data` VALUES ('21', 'RECHARGE_MONEY', '500', '500', '500', '2025-10-28 11:10:57', '2025-10-28 11:10:57');
INSERT INTO `sys_dict_data` VALUES ('22', 'RECHARGE_MONEY', '1000', '1000', '1000', '2025-10-28 11:11:06', '2025-10-28 11:11:06');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `code` varchar(100) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES ('1', '商品类型', 'PRODUCT_TYPE', '2025-04-06 07:09:11', '2025-10-11 16:31:53');
INSERT INTO `sys_dict_type` VALUES ('2', '充值类型', 'RECHARGE_TYPE', '2025-10-27 12:05:40', '2025-10-27 12:05:40');
INSERT INTO `sys_dict_type` VALUES ('3', '支付类型', 'PAY_TYPE', '2025-10-27 12:24:19', '2025-10-27 12:24:19');
INSERT INTO `sys_dict_type` VALUES ('4', '充值金额', 'RECHARGE_MONEY', '2025-10-28 11:08:57', '2025-10-28 11:08:57');

-- ----------------------------
-- Table structure for sys_distribute
-- ----------------------------
DROP TABLE IF EXISTS `sys_distribute`;
CREATE TABLE `sys_distribute` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) DEFAULT NULL,
  `witkey_id` bigint(20) DEFAULT NULL,
  `manage_id` bigint(20) DEFAULT NULL,
  `is_cancel` tinyint(4) DEFAULT NULL,
  `reason` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order` (`order_id`),
  KEY `idx_witkey` (`witkey_id`),
  KEY `idx_manage` (`manage_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_distribute
-- ----------------------------
INSERT INTO `sys_distribute` VALUES ('3', '4', '17', '1', '2', 'ghdfghfgh', '2025-12-22 18:48:06');
INSERT INTO `sys_distribute` VALUES ('4', '4', '16', '1', '1', null, '2025-12-23 08:33:24');

-- ----------------------------
-- Table structure for sys_game
-- ----------------------------
DROP TABLE IF EXISTS `sys_game`;
CREATE TABLE `sys_game` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) DEFAULT '0' COMMENT '顶级分类',
  `name` varchar(50) DEFAULT NULL COMMENT '分类名称',
  `pic` varchar(255) DEFAULT NULL COMMENT '分类背景图',
  `sort` int(10) DEFAULT NULL COMMENT '分类排序',
  `description` varchar(255) DEFAULT NULL COMMENT '分类描述',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime DEFAULT NULL COMMENT '更新日期',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_game
-- ----------------------------
INSERT INTO `sys_game` VALUES ('38', '0', '三角洲行动', 'https://q8.itc.cn/q_70/images03/20241017/b3ff223261b14b18a3ed2cdadec148f7.png', null, 'gsdgf', '2025-01-19 17:50:20', '2025-10-12 19:42:16');
INSERT INTO `sys_game` VALUES ('39', '0', 'APEX', '/public/uploads/2025-12-16/deyyn4mebsd0acktsm.jpg', null, 'fsdf', '2025-04-05 09:19:16', '2025-12-15 17:18:59');
INSERT INTO `sys_game` VALUES ('40', '0', '测试游戏', '/public/uploads/2025-12-16/deyyqyxho32gc2059e.png', null, '撒打发士大夫', '2025-12-15 17:23:59', '2025-12-15 17:23:59');

-- ----------------------------
-- Table structure for sys_level
-- ----------------------------
DROP TABLE IF EXISTS `sys_level`;
CREATE TABLE `sys_level` (
  `id` bigint(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `level` int(11) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `experience` int(11) DEFAULT NULL,
  `benefits` text,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_level
-- ----------------------------
INSERT INTO `sys_level` VALUES ('1', 'Lv1', '1', 'https://demo.zibll.com/wp-content/themes/zibll/img/user-level-1.png', '200', '[]', '2024-12-23 08:08:49', '2024-12-23 08:08:58');
INSERT INTO `sys_level` VALUES ('2', 'Lv2', '2', 'https://demo.zibll.com/wp-content/themes/zibll/img/user-level-1.png', '400', '[]', '2024-12-23 08:08:51', '2024-12-23 08:09:00');
INSERT INTO `sys_level` VALUES ('3', 'Lv3', '3', 'https://demo.zibll.com/wp-content/themes/zibll/img/user-level-1.png', '600', '[]', '2024-12-23 08:08:53', '2024-12-23 08:09:02');
INSERT INTO `sys_level` VALUES ('4', 'Lv4', '4', 'https://demo.zibll.com/wp-content/themes/zibll/img/user-level-1.png', '800', '[]', '2024-12-23 08:08:55', '2024-12-23 08:09:05');
INSERT INTO `sys_level` VALUES ('5', 'Lv5', '5', 'https://demo.zibll.com/wp-content/themes/zibll/img/user-level-1.png', '1000', null, '2025-11-05 12:35:58', '2025-11-05 13:36:43');

-- ----------------------------
-- Table structure for sys_manage
-- ----------------------------
DROP TABLE IF EXISTS `sys_manage`;
CREATE TABLE `sys_manage` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `name` varchar(30) NOT NULL COMMENT '用户昵称',
  `phone` varchar(11) DEFAULT '' COMMENT '手机号码',
  `sex` tinyint(4) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `salt` char(10) DEFAULT NULL COMMENT '密码盐',
  `status` tinyint(4) DEFAULT '0' COMMENT '帐号状态（1停用,2正常）',
  `qq` varchar(50) DEFAULT NULL,
  `wechat` varchar(500) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `delete_time` datetime DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_manage
-- ----------------------------
INSERT INTO `sys_manage` VALUES ('1', 'Eforinaj', '17777978384', '1', 'https://static.7b2.com/wp-content/uploads/2024/10/anime-7687171_1280.jpg', 'bd623e944177ee599da3ebda94e588ea', '0LIXaK', '2', null, null, '2021-05-10 05:50:02', '2025-09-14 17:04:56', '', null);
INSERT INTO `sys_manage` VALUES ('4', '小河', '17772978384', '3', 'https://static.7b2.com/wp-content/uploads/2024/10/anime-7687171_1280.jpg', 'dsfgsdfg', null, '2', null, null, '2024-12-18 04:02:46', '2024-12-18 04:02:46', null, null);
INSERT INTO `sys_manage` VALUES ('5', '小海', '17747978384', '3', 'https://static.7b2.com/wp-content/uploads/2024/10/anime-7687171_1280.jpg', 'dfgsdfgsfg', null, '2', null, null, '2024-12-18 04:28:24', '2024-12-18 04:28:24', null, null);

-- ----------------------------
-- Table structure for sys_manage_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_manage_login_log`;
CREATE TABLE `sys_manage_login_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `manage_id` bigint(20) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_manage` (`manage_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_manage_login_log
-- ----------------------------
INSERT INTO `sys_manage_login_log` VALUES ('46', '1', '::1', '2025-11-03 20:56:27');
INSERT INTO `sys_manage_login_log` VALUES ('47', '1', '::1', '2025-11-05 11:19:51');
INSERT INTO `sys_manage_login_log` VALUES ('48', '1', '::1', '2025-11-11 22:00:35');
INSERT INTO `sys_manage_login_log` VALUES ('49', '1', '::1', '2025-11-12 23:51:49');
INSERT INTO `sys_manage_login_log` VALUES ('50', '1', '::1', '2025-11-17 17:35:30');
INSERT INTO `sys_manage_login_log` VALUES ('51', '1', '::1', '2025-11-24 10:08:44');
INSERT INTO `sys_manage_login_log` VALUES ('52', '1', '::1', '2025-11-25 10:10:02');
INSERT INTO `sys_manage_login_log` VALUES ('53', '1', '::1', '2025-11-28 08:48:46');
INSERT INTO `sys_manage_login_log` VALUES ('54', '1', '::1', '2025-11-28 08:50:40');
INSERT INTO `sys_manage_login_log` VALUES ('55', '1', '::1', '2025-11-28 10:10:09');
INSERT INTO `sys_manage_login_log` VALUES ('56', '1', '::1', '2025-11-28 11:07:32');
INSERT INTO `sys_manage_login_log` VALUES ('57', '1', '::1', '2025-11-28 11:32:25');
INSERT INTO `sys_manage_login_log` VALUES ('58', '1', '::1', '2025-11-28 11:33:05');
INSERT INTO `sys_manage_login_log` VALUES ('59', '1', '::1', '2025-11-28 11:42:40');
INSERT INTO `sys_manage_login_log` VALUES ('60', '1', '::1', '2025-11-28 12:57:24');
INSERT INTO `sys_manage_login_log` VALUES ('61', '1', '::1', '2025-11-28 18:52:00');
INSERT INTO `sys_manage_login_log` VALUES ('62', '1', '::1', '2025-11-29 18:58:47');
INSERT INTO `sys_manage_login_log` VALUES ('63', '1', '::1', '2025-11-30 06:13:04');
INSERT INTO `sys_manage_login_log` VALUES ('64', '1', '::1', '2025-11-30 07:31:32');
INSERT INTO `sys_manage_login_log` VALUES ('65', '1', '::1', '2025-12-01 08:21:09');
INSERT INTO `sys_manage_login_log` VALUES ('66', '1', '::1', '2025-12-02 08:29:55');
INSERT INTO `sys_manage_login_log` VALUES ('67', '1', '::1', '2025-12-03 08:03:49');
INSERT INTO `sys_manage_login_log` VALUES ('68', '1', '::1', '2025-12-04 08:09:43');
INSERT INTO `sys_manage_login_log` VALUES ('69', '1', '::1', '2025-12-05 08:31:17');
INSERT INTO `sys_manage_login_log` VALUES ('70', '1', '::1', '2025-12-06 15:51:30');
INSERT INTO `sys_manage_login_log` VALUES ('71', '1', '::1', '2025-12-07 16:08:07');
INSERT INTO `sys_manage_login_log` VALUES ('72', '1', '::1', '2025-12-08 16:08:47');
INSERT INTO `sys_manage_login_log` VALUES ('73', '1', '::1', '2025-12-09 16:11:23');
INSERT INTO `sys_manage_login_log` VALUES ('74', '1', '::1', '2025-12-10 16:30:29');
INSERT INTO `sys_manage_login_log` VALUES ('75', '1', '::1', '2025-12-12 06:46:43');
INSERT INTO `sys_manage_login_log` VALUES ('76', '1', '::1', '2025-12-13 13:38:25');
INSERT INTO `sys_manage_login_log` VALUES ('77', '1', '::1', '2025-12-14 13:53:45');
INSERT INTO `sys_manage_login_log` VALUES ('78', '1', '::1', '2025-12-15 14:22:51');
INSERT INTO `sys_manage_login_log` VALUES ('79', '1', '::1', '2025-12-16 14:42:03');
INSERT INTO `sys_manage_login_log` VALUES ('80', '1', '::1', '2025-12-17 14:53:34');
INSERT INTO `sys_manage_login_log` VALUES ('81', '1', '::1', '2025-12-19 07:13:28');
INSERT INTO `sys_manage_login_log` VALUES ('82', '1', '::1', '2025-12-20 08:23:50');
INSERT INTO `sys_manage_login_log` VALUES ('83', '1', '::1', '2025-12-21 09:03:20');
INSERT INTO `sys_manage_login_log` VALUES ('84', '1', '::1', '2025-12-22 11:11:22');
INSERT INTO `sys_manage_login_log` VALUES ('85', '1', '::1', '2025-12-23 16:44:37');

-- ----------------------------
-- Table structure for sys_manage_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_manage_role`;
CREATE TABLE `sys_manage_role` (
  `manage_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`manage_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_manage_role
-- ----------------------------
INSERT INTO `sys_manage_role` VALUES ('1', '1');
INSERT INTO `sys_manage_role` VALUES ('4', '2');
INSERT INTO `sys_manage_role` VALUES ('4', '4');
INSERT INTO `sys_manage_role` VALUES ('5', '2');
INSERT INTO `sys_manage_role` VALUES ('5', '4');

-- ----------------------------
-- Table structure for sys_media
-- ----------------------------
DROP TABLE IF EXISTS `sys_media`;
CREATE TABLE `sys_media` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `path` varchar(255) DEFAULT NULL COMMENT '存放路径',
  `name` varchar(255) DEFAULT NULL COMMENT '文件名称',
  `or_name` varchar(255) DEFAULT NULL COMMENT '原始文件名称',
  `size` varchar(50) NOT NULL COMMENT '文件大小',
  `ext` varchar(10) DEFAULT NULL COMMENT '文件后缀',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `delete_time` datetime DEFAULT NULL,
  `media_type` varchar(20) DEFAULT NULL COMMENT '文件类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=410 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_media
-- ----------------------------
INSERT INTO `sys_media` VALUES ('189', '/resource/uploads/2025-09-22/dcz8cyg7jh0czzjrkm.jpg', 'dcz8cyg7jh0czzjrkm.jpg', 'soybean.jpg', '114241', '.jpg', '2025-09-22 09:43:54', null, 'jpg');
INSERT INTO `sys_media` VALUES ('190', '/public/uploads/2025-10-10/dddzzafadbiomsgjji.png', 'dddzzafadbiomsgjji.png', 'QQ20250531-142850.png', '133417', '.png', '2025-10-09 18:19:53', null, 'png');
INSERT INTO `sys_media` VALUES ('191', '/public/uploads/2025-10-10/dde00fplpydchcel5v.png', 'dde00fplpydchcel5v.png', '462d2442e61c8c123dd9d3bafd37d9e3be0e98577d766e-0nRsrq.png', '8222318', '.png', '2025-10-09 18:21:23', null, 'png');
INSERT INTO `sys_media` VALUES ('192', '/public/uploads/2025-10-10/dde0xak52xbgpw7l9p.png', 'dde0xak52xbgpw7l9p.png', '462d2442e61c8c123dd9d3bafd37d9e3be0e98577d766e-0nRsrq.png', '8222318', '.png', '2025-10-09 19:04:18', null, 'png');
INSERT INTO `sys_media` VALUES ('193', '/public/uploads/2025-10-10/dde0xbuiipzkslkfcj.png', 'dde0xbuiipzkslkfcj.png', 'QQ20250531-142850.png', '133417', '.png', '2025-10-09 19:04:21', null, 'png');
INSERT INTO `sys_media` VALUES ('194', '/public/uploads/2025-10-19/ddm80csjjai88vec5k.png', 'ddm80csjjai88vec5k.png', 'QQ20250531-142850.png', '133417', '.png', '2025-10-19 10:18:45', null, 'png');
INSERT INTO `sys_media` VALUES ('195', '/public/uploads/2025-10-19/ddm80xend5monpdue8.png', 'ddm80xend5monpdue8.png', 'QQ20250531-142850.png', '133417', '.png', '2025-10-19 10:19:30', null, 'png');
INSERT INTO `sys_media` VALUES ('196', '/public/uploads/2025-10-19/ddm856774qwwvhwfwr.png', 'ddm856774qwwvhwfwr.png', 'QQ20250531-142850.png', '133417', '.png', '2025-10-19 10:25:03', null, 'png');
INSERT INTO `sys_media` VALUES ('197', '/public/uploads/2025-10-23/ddpipr9chvnwabnbk8.jpg', 'ddpipr9chvnwabnbk8.jpg', 'ddpipr9chvnwabnbk8.jpg', '3161', '.jpg', '2025-10-23 07:20:07', null, 'jpg');
INSERT INTO `sys_media` VALUES ('198', '/public/uploads/2025-10-23/ddpiptw17qfgwldgp1.jpg', 'ddpiptw17qfgwldgp1.jpg', 'ddpiptw17qfgwldgp1.jpg', '3161', '.jpg', '2025-10-23 07:20:13', null, 'jpg');
INSERT INTO `sys_media` VALUES ('199', '/public/uploads/2025-10-23/ddpiq3neguoklxgukp.jpg', 'ddpiq3neguoklxgukp.jpg', 'ddpiq3neguoklxgukp.jpg', '3161', '.jpg', '2025-10-23 07:20:34', null, 'jpg');
INSERT INTO `sys_media` VALUES ('200', '/public/uploads/2025-11-03/ddyfl4z1mmr8h1k2jf.jpg', 'ddyfl4z1mmr8h1k2jf.jpg', 'ddyfl4z1mmr8h1k2jf.jpg', '114241', '.jpg', '2025-11-02 18:47:01', null, 'jpg');
INSERT INTO `sys_media` VALUES ('201', '/public/uploads/2025-11-03/ddyfogy8ukfsnyewu4.jpg', 'ddyfogy8ukfsnyewu4.jpg', 'ddyfogy8ukfsnyewu4.jpg', '114241', '.jpg', '2025-11-02 18:51:22', null, 'jpg');
INSERT INTO `sys_media` VALUES ('202', '/public/uploads/2025-11-03/ddyfoj3uxnmgkezvgz.jpg', 'ddyfoj3uxnmgkezvgz.jpg', 'ddyfoj3uxnmgkezvgz.jpg', '114241', '.jpg', '2025-11-02 18:51:27', null, 'jpg');
INSERT INTO `sys_media` VALUES ('203', '/public/uploads/2025-11-03/ddyfokj4nie8jmh3c9.jpg', 'ddyfokj4nie8jmh3c9.jpg', 'ddyfokj4nie8jmh3c9.jpg', '114241', '.jpg', '2025-11-02 18:51:30', null, 'jpg');
INSERT INTO `sys_media` VALUES ('204', '/public/uploads/2025-11-03/ddyfpbvquaw0rmymuy.jpg', 'ddyfpbvquaw0rmymuy.jpg', 'ddyfpbvquaw0rmymuy.jpg', '114241', '.jpg', '2025-11-02 18:52:30', null, 'jpg');
INSERT INTO `sys_media` VALUES ('205', '/public/uploads/2025-11-03/ddyfpg7a9d6oelcm4s.jpg', 'ddyfpg7a9d6oelcm4s.jpg', 'ddyfpg7a9d6oelcm4s.jpg', '114241', '.jpg', '2025-11-02 18:52:39', null, 'jpg');
INSERT INTO `sys_media` VALUES ('206', '/public/uploads/2025-11-03/ddyfpgr04gwkd6jeg1.jpg', 'ddyfpgr04gwkd6jeg1.jpg', 'ddyfpgr04gwkd6jeg1.jpg', '114241', '.jpg', '2025-11-02 18:52:40', null, 'jpg');
INSERT INTO `sys_media` VALUES ('207', '/public/uploads/2025-11-03/ddyfphbc923gtxcmxo.jpg', 'ddyfphbc923gtxcmxo.jpg', 'ddyfphbc923gtxcmxo.jpg', '114241', '.jpg', '2025-11-02 18:52:42', null, 'jpg');
INSERT INTO `sys_media` VALUES ('208', '/public/uploads/2025-11-04/ddzagebvlpg0vj3kiy.jpg', 'ddzagebvlpg0vj3kiy.jpg', 'ddzagebvlpg0vj3kiy.jpg', '114241', '.jpg', '2025-11-03 18:58:24', null, 'jpg');
INSERT INTO `sys_media` VALUES ('209', '/public/uploads/2025-11-04/ddzai4ni0fzgh8vbkt.jpg', 'ddzai4ni0fzgh8vbkt.jpg', 'ddzai4ni0fzgh8vbkt.jpg', '114241', '.jpg', '2025-11-03 19:00:40', null, 'jpg');
INSERT INTO `sys_media` VALUES ('210', '/public/uploads/2025-11-04/ddzainenn5fstr9028.jpg', 'ddzainenn5fstr9028.jpg', 'ddzainenn5fstr9028.jpg', '114241', '.jpg', '2025-11-03 19:01:21', null, 'jpg');
INSERT INTO `sys_media` VALUES ('211', '/public/uploads/2025-11-04/ddzajhrhrocgv76jzy.jpg', 'ddzajhrhrocgv76jzy.jpg', 'ddzajhrhrocgv76jzy.jpg', '114241', '.jpg', '2025-11-03 19:02:27', null, 'jpg');
INSERT INTO `sys_media` VALUES ('212', '/public/uploads/2025-11-04/ddzajpm6j95oxi7yf3.jpg', 'ddzajpm6j95oxi7yf3.jpg', 'ddzajpm6j95oxi7yf3.jpg', '4731', '.jpg', '2025-11-03 19:02:44', null, 'jpg');
INSERT INTO `sys_media` VALUES ('213', '/public/uploads/2025-11-04/ddzzy58rgfuov97lim.jpg', 'ddzzy58rgfuov97lim.jpg', 'ddzzy58rgfuov97lim.jpg', '114241', '.jpg', '2025-11-04 14:57:03', null, 'jpg');
INSERT INTO `sys_media` VALUES ('214', '/public/uploads/2025-11-04/ddzzy5w0uq4sqjolgu.jpg', 'ddzzy5w0uq4sqjolgu.jpg', 'ddzzy5w0uq4sqjolgu.jpg', '114241', '.jpg', '2025-11-04 14:57:04', null, 'jpg');
INSERT INTO `sys_media` VALUES ('215', '/public/uploads/2025-11-04/ddzzy6i53xcshf5ivz.jpg', 'ddzzy6i53xcshf5ivz.jpg', 'ddzzy6i53xcshf5ivz.jpg', '114241', '.jpg', '2025-11-04 14:57:06', null, 'jpg');
INSERT INTO `sys_media` VALUES ('216', '/public/uploads/2025-11-04/de003c89q2cgky9pyu.jpg', 'de003c89q2cgky9pyu.jpg', 'de003c89q2cgky9pyu.jpg', '114241', '.jpg', '2025-11-04 15:03:50', null, 'jpg');
INSERT INTO `sys_media` VALUES ('217', '/public/uploads/2025-11-05/de0rj84xqjn4pmwptk.jpg', 'de0rj84xqjn4pmwptk.jpg', 'soybean.jpg', '114241', '.jpg', '2025-11-05 12:34:05', null, 'jpg');
INSERT INTO `sys_media` VALUES ('218', '/public/uploads/2025-11-05/de0rjs9bgpfsusdzmc.jpg', 'de0rjs9bgpfsusdzmc.jpg', 'soybean.jpg', '114241', '.jpg', '2025-11-05 12:34:49', null, 'jpg');
INSERT INTO `sys_media` VALUES ('219', '/public/uploads/2025-11-05/de0rknf6sdl05earij.jpg', 'de0rknf6sdl05earij.jpg', 'soybean.jpg', '114241', '.jpg', '2025-11-05 12:35:56', null, 'jpg');
INSERT INTO `sys_media` VALUES ('220', '/public/uploads/2025-11-10/de511j8ol0wwiaogxc.jpg', 'de511j8ol0wwiaogxc.jpg', 'de511j8ol0wwiaogxc.jpg', '58805', '.jpg', '2025-11-10 12:51:49', null, 'jpg');
INSERT INTO `sys_media` VALUES ('221', '/public/uploads/2025-11-10/de51n01a255gaqwtjo.jpg', 'de51n01a255gaqwtjo.jpg', 'de51n01a255gaqwtjo.jpg', '58805', '.jpg', '2025-11-10 13:19:51', null, 'jpg');
INSERT INTO `sys_media` VALUES ('222', '/public/uploads/2025-11-10/de53p3tt55lsl8htke.jpg', 'de53p3tt55lsl8htke.jpg', 'de53p3tt55lsl8htke.jpg', '58805', '.jpg', '2025-11-10 14:56:39', null, 'jpg');
INSERT INTO `sys_media` VALUES ('223', '/public/uploads/2025-11-10/de5497yf4230owpdr1.png', 'de5497yf4230owpdr1.png', 'de5497yf4230owpdr1.png', '21119', '.png', '2025-11-10 15:22:55', null, 'png');
INSERT INTO `sys_media` VALUES ('224', '/public/uploads/2025-11-10/de549zi7cfnk4fyfk9.png', 'de549zi7cfnk4fyfk9.png', 'de549zi7cfnk4fyfk9.png', '21119', '.png', '2025-11-10 15:23:55', null, 'png');
INSERT INTO `sys_media` VALUES ('225', '/public/uploads/2025-11-10/de54ahf6jqugr3gggm.jpg', 'de54ahf6jqugr3gggm.jpg', 'de54ahf6jqugr3gggm.jpg', '58805', '.jpg', '2025-11-10 15:24:34', null, 'jpg');
INSERT INTO `sys_media` VALUES ('226', '/public/uploads/2025-11-10/de54dz3jg8icrt8t2i.jpg', 'de54dz3jg8icrt8t2i.jpg', 'de54dz3jg8icrt8t2i.jpg', '58805', '.jpg', '2025-11-10 15:29:07', null, 'jpg');
INSERT INTO `sys_media` VALUES ('227', '/public/uploads/2025-11-10/de54e9hcupkcuvlsk5.png', 'de54e9hcupkcuvlsk5.png', 'de54e9hcupkcuvlsk5.png', '21119', '.png', '2025-11-10 15:29:30', null, 'png');
INSERT INTO `sys_media` VALUES ('228', '/public/uploads/2025-11-10/de54eao374xosjt8zz.jpg', 'de54eao374xosjt8zz.jpg', 'de54eao374xosjt8zz.jpg', '58805', '.jpg', '2025-11-10 15:29:33', null, 'jpg');
INSERT INTO `sys_media` VALUES ('229', '/public/uploads/2025-11-10/de54epwzvmzkxveshd.jpg', 'de54epwzvmzkxveshd.jpg', 'de54epwzvmzkxveshd.jpg', '58805', '.jpg', '2025-11-10 15:30:06', null, 'jpg');
INSERT INTO `sys_media` VALUES ('230', '/public/uploads/2025-11-10/de54f0k6dgccm2uwnc.jpg', 'de54f0k6dgccm2uwnc.jpg', 'de54f0k6dgccm2uwnc.jpg', '58805', '.jpg', '2025-11-10 15:30:29', null, 'jpg');
INSERT INTO `sys_media` VALUES ('231', '/public/uploads/2025-11-10/de54gsz5ppu0cbtoqk.png', 'de54gsz5ppu0cbtoqk.png', 'de54gsz5ppu0cbtoqk.png', '21119', '.png', '2025-11-10 15:32:49', null, 'png');
INSERT INTO `sys_media` VALUES ('232', '/public/uploads/2025-11-10/de54hekdysp8wbnfdd.png', 'de54hekdysp8wbnfdd.png', 'de54hekdysp8wbnfdd.png', '21119', '.png', '2025-11-10 15:33:36', null, 'png');
INSERT INTO `sys_media` VALUES ('233', '/public/uploads/2025-11-10/de54i1b3v0dcuaycj0.jpg', 'de54i1b3v0dcuaycj0.jpg', 'de54i1b3v0dcuaycj0.jpg', '58805', '.jpg', '2025-11-10 15:34:26', null, 'jpg');
INSERT INTO `sys_media` VALUES ('234', '/public/uploads/2025-11-10/de54jyju8brgyx47sm.jpg', 'de54jyju8brgyx47sm.jpg', 'de54jyju8brgyx47sm.jpg', '58805', '.jpg', '2025-11-10 15:36:56', null, 'jpg');
INSERT INTO `sys_media` VALUES ('235', '/public/uploads/2025-11-10/de54k3nvq1kkulp714.jpg', 'de54k3nvq1kkulp714.jpg', 'de54k3nvq1kkulp714.jpg', '58805', '.jpg', '2025-11-10 15:37:08', null, 'jpg');
INSERT INTO `sys_media` VALUES ('236', '/public/uploads/2025-11-10/de54kb3byvigwvnxbs.jpg', 'de54kb3byvigwvnxbs.jpg', 'de54kb3byvigwvnxbs.jpg', '58805', '.jpg', '2025-11-10 15:37:24', null, 'jpg');
INSERT INTO `sys_media` VALUES ('237', '/public/uploads/2025-11-10/de54lkgvdzy4zuxhyz.jpg', 'de54lkgvdzy4zuxhyz.jpg', 'de54lkgvdzy4zuxhyz.jpg', '58805', '.jpg', '2025-11-10 15:39:03', null, 'jpg');
INSERT INTO `sys_media` VALUES ('238', '/public/uploads/2025-11-10/de54ng4z885krk0nc5.jpg', 'de54ng4z885krk0nc5.jpg', 'de54ng4z885krk0nc5.jpg', '58805', '.jpg', '2025-11-10 15:41:30', null, 'jpg');
INSERT INTO `sys_media` VALUES ('239', '/public/uploads/2025-11-10/de54oagfown0raxi4l.jpg', 'de54oagfown0raxi4l.jpg', 'de54oagfown0raxi4l.jpg', '58805', '.jpg', '2025-11-10 15:42:36', null, 'jpg');
INSERT INTO `sys_media` VALUES ('240', '/public/uploads/2025-11-10/de54ochmxqncewneoy.jpg', 'de54ochmxqncewneoy.jpg', 'de54ochmxqncewneoy.jpg', '58805', '.jpg', '2025-11-10 15:42:40', null, 'jpg');
INSERT INTO `sys_media` VALUES ('241', '/public/uploads/2025-11-10/de54q6u4tyv41rdrkk.jpg', 'de54q6u4tyv41rdrkk.jpg', 'de54q6u4tyv41rdrkk.jpg', '58805', '.jpg', '2025-11-10 15:45:05', null, 'jpg');
INSERT INTO `sys_media` VALUES ('242', '/public/uploads/2025-11-10/de54qbo6ad40bvvq7z.jpg', 'de54qbo6ad40bvvq7z.jpg', 'de54qbo6ad40bvvq7z.jpg', '58805', '.jpg', '2025-11-10 15:45:15', null, 'jpg');
INSERT INTO `sys_media` VALUES ('243', '/public/uploads/2025-11-10/de54qk2n6uh0kugscu.jpg', 'de54qk2n6uh0kugscu.jpg', 'de54qk2n6uh0kugscu.jpg', '58805', '.jpg', '2025-11-10 15:45:33', null, 'jpg');
INSERT INTO `sys_media` VALUES ('244', '/public/uploads/2025-11-10/de54qvrfu2x0ixw8ur.jpg', 'de54qvrfu2x0ixw8ur.jpg', 'de54qvrfu2x0ixw8ur.jpg', '58805', '.jpg', '2025-11-10 15:45:59', null, 'jpg');
INSERT INTO `sys_media` VALUES ('245', '/public/uploads/2025-11-10/de54rfanlkbcrs4mni.jpg', 'de54rfanlkbcrs4mni.jpg', 'de54rfanlkbcrs4mni.jpg', '58805', '.jpg', '2025-11-10 15:46:41', null, 'jpg');
INSERT INTO `sys_media` VALUES ('246', '/public/uploads/2025-11-10/de54t3p6v26wud7xar.jpg', 'de54t3p6v26wud7xar.jpg', 'de54t3p6v26wud7xar.jpg', '58805', '.jpg', '2025-11-10 15:48:53', null, 'jpg');
INSERT INTO `sys_media` VALUES ('247', '/public/uploads/2025-11-10/de54u0gkfrpgubvw2x.jpg', 'de54u0gkfrpgubvw2x.jpg', 'de54u0gkfrpgubvw2x.jpg', '58805', '.jpg', '2025-11-10 15:50:04', null, 'jpg');
INSERT INTO `sys_media` VALUES ('248', '/public/uploads/2025-11-10/de54udm6g728j2onti.jpg', 'de54udm6g728j2onti.jpg', 'de54udm6g728j2onti.jpg', '58805', '.jpg', '2025-11-10 15:50:33', null, 'jpg');
INSERT INTO `sys_media` VALUES ('249', '/public/uploads/2025-11-10/de54uknr28e0kakgsl.jpg', 'de54uknr28e0kakgsl.jpg', 'de54uknr28e0kakgsl.jpg', '58805', '.jpg', '2025-11-10 15:50:48', null, 'jpg');
INSERT INTO `sys_media` VALUES ('250', '/public/uploads/2025-11-10/de54uwyhh0poto5mpo.jpg', 'de54uwyhh0poto5mpo.jpg', 'de54uwyhh0poto5mpo.jpg', '58805', '.jpg', '2025-11-10 15:51:15', null, 'jpg');
INSERT INTO `sys_media` VALUES ('251', '/public/uploads/2025-11-10/de54uzxxwmu4u1iemx.jpg', 'de54uzxxwmu4u1iemx.jpg', 'de54uzxxwmu4u1iemx.jpg', '58805', '.jpg', '2025-11-10 15:51:21', null, 'jpg');
INSERT INTO `sys_media` VALUES ('252', '/public/uploads/2025-11-10/de54vefqbgg04bzmfs.jpg', 'de54vefqbgg04bzmfs.jpg', 'de54vefqbgg04bzmfs.jpg', '58805', '.jpg', '2025-11-10 15:51:53', null, 'jpg');
INSERT INTO `sys_media` VALUES ('253', '/public/uploads/2025-11-10/de54vtsv8h3wvenkru.jpg', 'de54vtsv8h3wvenkru.jpg', 'de54vtsv8h3wvenkru.jpg', '58805', '.jpg', '2025-11-10 15:52:26', null, 'jpg');
INSERT INTO `sys_media` VALUES ('254', '/public/uploads/2025-11-10/de54wuy0rbnwto4owq.jpg', 'de54wuy0rbnwto4owq.jpg', 'de54wuy0rbnwto4owq.jpg', '58805', '.jpg', '2025-11-10 15:53:47', null, 'jpg');
INSERT INTO `sys_media` VALUES ('255', '/public/uploads/2025-11-10/de54xagk6f90vxl5tw.jpg', 'de54xagk6f90vxl5tw.jpg', 'de54xagk6f90vxl5tw.jpg', '58805', '.jpg', '2025-11-10 15:54:21', null, 'jpg');
INSERT INTO `sys_media` VALUES ('256', '/public/uploads/2025-11-10/de54yfl41oygo5rnlf.jpg', 'de54yfl41oygo5rnlf.jpg', 'de54yfl41oygo5rnlf.jpg', '58805', '.jpg', '2025-11-10 15:55:51', null, 'jpg');
INSERT INTO `sys_media` VALUES ('257', '/public/uploads/2025-11-10/de54yn1ylyjs2kiuac.jpg', 'de54yn1ylyjs2kiuac.jpg', 'de54yn1ylyjs2kiuac.jpg', '58805', '.jpg', '2025-11-10 15:56:07', null, 'jpg');
INSERT INTO `sys_media` VALUES ('258', '/public/uploads/2025-11-10/de54yx8kxhe0me1qkn.jpg', 'de54yx8kxhe0me1qkn.jpg', 'de54yx8kxhe0me1qkn.jpg', '58805', '.jpg', '2025-11-10 15:56:29', null, 'jpg');
INSERT INTO `sys_media` VALUES ('259', '/public/uploads/2025-11-10/de551m32fm30dduigb.jpg', 'de551m32fm30dduigb.jpg', 'de551m32fm30dduigb.jpg', '58805', '.jpg', '2025-11-10 16:00:00', null, 'jpg');
INSERT INTO `sys_media` VALUES ('260', '/public/uploads/2025-11-11/de551rsvspt4n4pifo.jpg', 'de551rsvspt4n4pifo.jpg', 'de551rsvspt4n4pifo.jpg', '58805', '.jpg', '2025-11-10 16:00:12', null, 'jpg');
INSERT INTO `sys_media` VALUES ('261', '/public/uploads/2025-11-11/de552rk2i2j88bkaw4.jpg', 'de552rk2i2j88bkaw4.jpg', 'de552rk2i2j88bkaw4.jpg', '58805', '.jpg', '2025-11-10 16:01:30', null, 'jpg');
INSERT INTO `sys_media` VALUES ('262', '/public/uploads/2025-11-11/de553g5bnvsgrnbtn2.jpg', 'de553g5bnvsgrnbtn2.jpg', 'de553g5bnvsgrnbtn2.jpg', '58805', '.jpg', '2025-11-10 16:02:24', null, 'jpg');
INSERT INTO `sys_media` VALUES ('263', '/public/uploads/2025-11-11/de55crbxfvp4tom5tx.jpg', 'de55crbxfvp4tom5tx.jpg', 'de55crbxfvp4tom5tx.jpg', '58805', '.jpg', '2025-11-10 16:14:33', null, 'jpg');
INSERT INTO `sys_media` VALUES ('264', '/public/uploads/2025-11-11/de55jy3vvmrgrg8ffn.jpg', 'de55jy3vvmrgrg8ffn.jpg', 'de55jy3vvmrgrg8ffn.jpg', '58805', '.jpg', '2025-11-10 16:23:57', null, 'jpg');
INSERT INTO `sys_media` VALUES ('265', '/public/uploads/2025-11-11/de55m0ng0x0ojk21kv.jpg', 'de55m0ng0x0ojk21kv.jpg', 'de55m0ng0x0ojk21kv.jpg', '58805', '.jpg', '2025-11-10 16:26:39', null, 'jpg');
INSERT INTO `sys_media` VALUES ('266', '/public/uploads/2025-11-11/de55mwhkcmqct6jwax.jpg', 'de55mwhkcmqct6jwax.jpg', 'de55mwhkcmqct6jwax.jpg', '58805', '.jpg', '2025-11-10 16:27:48', null, 'jpg');
INSERT INTO `sys_media` VALUES ('267', '/public/uploads/2025-11-11/de55nef6kovkr4sncd.jpg', 'de55nef6kovkr4sncd.jpg', 'de55nef6kovkr4sncd.jpg', '58805', '.jpg', '2025-11-10 16:28:27', null, 'jpg');
INSERT INTO `sys_media` VALUES ('268', '/public/uploads/2025-11-11/de55nkwbktwgvknavb.jpg', 'de55nkwbktwgvknavb.jpg', 'de55nkwbktwgvknavb.jpg', '58805', '.jpg', '2025-11-10 16:28:41', null, 'jpg');
INSERT INTO `sys_media` VALUES ('269', '/public/uploads/2025-11-11/de55o72s7be0jowhty.jpg', 'de55o72s7be0jowhty.jpg', 'de55o72s7be0jowhty.jpg', '58805', '.jpg', '2025-11-10 16:29:30', null, 'jpg');
INSERT INTO `sys_media` VALUES ('270', '/public/uploads/2025-11-11/de55ogapohj8i7thje.jpg', 'de55ogapohj8i7thje.jpg', 'de55ogapohj8i7thje.jpg', '58805', '.jpg', '2025-11-10 16:29:50', null, 'jpg');
INSERT INTO `sys_media` VALUES ('271', '/public/uploads/2025-11-11/de56uxo2yz1sedsmcp.jpg', 'de56uxo2yz1sedsmcp.jpg', 'de56uxo2yz1sedsmcp.jpg', '58805', '.jpg', '2025-11-10 17:25:19', null, 'jpg');
INSERT INTO `sys_media` VALUES ('272', '/public/uploads/2025-11-11/de570zr2g37gfiemhk.jpg', 'de570zr2g37gfiemhk.jpg', 'de570zr2g37gfiemhk.jpg', '58805', '.jpg', '2025-11-10 17:33:13', null, 'jpg');
INSERT INTO `sys_media` VALUES ('273', '/public/uploads/2025-11-11/de573mr2utwgmsasuc.jpg', 'de573mr2utwgmsasuc.jpg', 'de573mr2utwgmsasuc.jpg', '58805', '.jpg', '2025-11-10 17:36:40', null, 'jpg');
INSERT INTO `sys_media` VALUES ('274', '/public/uploads/2025-11-11/de574jc4q2b83enkhj.jpg', 'de574jc4q2b83enkhj.jpg', 'de574jc4q2b83enkhj.jpg', '58805', '.jpg', '2025-11-10 17:37:51', null, 'jpg');
INSERT INTO `sys_media` VALUES ('275', '/public/uploads/2025-11-11/de57698nqx5gqeqmhb.jpg', 'de57698nqx5gqeqmhb.jpg', 'de57698nqx5gqeqmhb.jpg', '58805', '.jpg', '2025-11-10 17:40:06', null, 'jpg');
INSERT INTO `sys_media` VALUES ('276', '/public/uploads/2025-11-11/de576fnqazm4gvrx2l.jpg', 'de576fnqazm4gvrx2l.jpg', 'de576fnqazm4gvrx2l.jpg', '58805', '.jpg', '2025-11-10 17:40:20', null, 'jpg');
INSERT INTO `sys_media` VALUES ('277', '/public/uploads/2025-11-11/de576u9szovsp7luys.jpg', 'de576u9szovsp7luys.jpg', 'de576u9szovsp7luys.jpg', '58805', '.jpg', '2025-11-10 17:40:52', null, 'jpg');
INSERT INTO `sys_media` VALUES ('278', '/public/uploads/2025-11-11/de5782o0xtdwp3vubx.jpg', 'de5782o0xtdwp3vubx.jpg', 'de5782o0xtdwp3vubx.jpg', '58805', '.jpg', '2025-11-10 17:42:28', null, 'jpg');
INSERT INTO `sys_media` VALUES ('279', '/public/uploads/2025-11-11/de578jpq2m2w0yojug.jpg', 'de578jpq2m2w0yojug.jpg', 'de578jpq2m2w0yojug.jpg', '58805', '.jpg', '2025-11-10 17:43:05', null, 'jpg');
INSERT INTO `sys_media` VALUES ('280', '/public/uploads/2025-11-11/de578sdfip7w2keixb.jpg', 'de578sdfip7w2keixb.jpg', 'de578sdfip7w2keixb.jpg', '58805', '.jpg', '2025-11-10 17:43:24', null, 'jpg');
INSERT INTO `sys_media` VALUES ('281', '/public/uploads/2025-11-11/de578xnod7s0isczvq.jpg', 'de578xnod7s0isczvq.jpg', 'de578xnod7s0isczvq.jpg', '58805', '.jpg', '2025-11-10 17:43:36', null, 'jpg');
INSERT INTO `sys_media` VALUES ('282', '/public/uploads/2025-11-11/de579xrwrsogfk9klq.jpg', 'de579xrwrsogfk9klq.jpg', 'de579xrwrsogfk9klq.jpg', '58805', '.jpg', '2025-11-10 17:44:54', null, 'jpg');
INSERT INTO `sys_media` VALUES ('283', '/public/uploads/2025-11-11/de57afxqa05wcbtpvf.png', 'de57afxqa05wcbtpvf.png', 'de57afxqa05wcbtpvf.png', '21119', '.png', '2025-11-10 17:45:34', null, 'png');
INSERT INTO `sys_media` VALUES ('284', '/public/uploads/2025-11-11/de57ai3f82lodyf7fc.jpg', 'de57ai3f82lodyf7fc.jpg', 'de57ai3f82lodyf7fc.jpg', '58805', '.jpg', '2025-11-10 17:45:39', null, 'jpg');
INSERT INTO `sys_media` VALUES ('285', '/public/uploads/2025-11-11/de57aj5dlcvk5y3cgl.png', 'de57aj5dlcvk5y3cgl.png', 'de57aj5dlcvk5y3cgl.png', '560', '.png', '2025-11-10 17:45:41', null, 'png');
INSERT INTO `sys_media` VALUES ('286', '/public/uploads/2025-11-11/de57ceqclw90bg28bd.jpg', 'de57ceqclw90bg28bd.jpg', 'de57ceqclw90bg28bd.jpg', '58805', '.jpg', '2025-11-10 17:48:08', null, 'jpg');
INSERT INTO `sys_media` VALUES ('287', '/public/uploads/2025-11-11/de57dd845210tf1jxu.jpg', 'de57dd845210tf1jxu.jpg', 'de57dd845210tf1jxu.jpg', '58805', '.jpg', '2025-11-10 17:49:23', null, 'jpg');
INSERT INTO `sys_media` VALUES ('288', '/public/uploads/2025-11-11/de57e19tc28gokcfi6.jpg', 'de57e19tc28gokcfi6.jpg', 'de57e19tc28gokcfi6.jpg', '58805', '.jpg', '2025-11-10 17:50:16', null, 'jpg');
INSERT INTO `sys_media` VALUES ('289', '/public/uploads/2025-11-11/de57fck6a0yoq2ifax.jpg', 'de57fck6a0yoq2ifax.jpg', 'de57fck6a0yoq2ifax.jpg', '58805', '.jpg', '2025-11-10 17:51:58', null, 'jpg');
INSERT INTO `sys_media` VALUES ('290', '/public/uploads/2025-11-11/de57fdvm3n2wdeev7n.jpg', 'de57fdvm3n2wdeev7n.jpg', 'de57fdvm3n2wdeev7n.jpg', '58805', '.jpg', '2025-11-10 17:52:01', null, 'jpg');
INSERT INTO `sys_media` VALUES ('291', '/public/uploads/2025-11-11/de57ftg7mjeozyinhb.jpg', 'de57ftg7mjeozyinhb.jpg', 'de57ftg7mjeozyinhb.jpg', '58805', '.jpg', '2025-11-10 17:52:35', null, 'jpg');
INSERT INTO `sys_media` VALUES ('292', '/public/uploads/2025-11-11/de57g1chhws8tuclaz.jpg', 'de57g1chhws8tuclaz.jpg', 'de57g1chhws8tuclaz.jpg', '58805', '.jpg', '2025-11-10 17:52:52', null, 'jpg');
INSERT INTO `sys_media` VALUES ('293', '/public/uploads/2025-11-11/de57gid7getsjgpkvf.jpg', 'de57gid7getsjgpkvf.jpg', 'de57gid7getsjgpkvf.jpg', '58805', '.jpg', '2025-11-10 17:53:29', null, 'jpg');
INSERT INTO `sys_media` VALUES ('294', '/public/uploads/2025-11-11/de57hbh6io4o7lm8wi.jpg', 'de57hbh6io4o7lm8wi.jpg', 'de57hbh6io4o7lm8wi.jpg', '58805', '.jpg', '2025-11-10 17:54:33', null, 'jpg');
INSERT INTO `sys_media` VALUES ('295', '/public/uploads/2025-11-11/de57hgmdsvcw7qafhm.jpg', 'de57hgmdsvcw7qafhm.jpg', 'de57hgmdsvcw7qafhm.jpg', '58805', '.jpg', '2025-11-10 17:54:44', null, 'jpg');
INSERT INTO `sys_media` VALUES ('296', '/public/uploads/2025-11-11/de57ia1pnbj4hjf1o6.jpg', 'de57ia1pnbj4hjf1o6.jpg', 'de57ia1pnbj4hjf1o6.jpg', '58805', '.jpg', '2025-11-10 17:55:48', null, 'jpg');
INSERT INTO `sys_media` VALUES ('297', '/public/uploads/2025-11-11/de57s1c8dgfs7xar0o.jpg', 'de57s1c8dgfs7xar0o.jpg', 'de57s1c8dgfs7xar0o.jpg', '58805', '.jpg', '2025-11-10 18:08:33', null, 'jpg');
INSERT INTO `sys_media` VALUES ('298', '/public/uploads/2025-11-11/de581ongws88j2aejr.png', 'de581ongws88j2aejr.png', 'de581ongws88j2aejr.png', '21119', '.png', '2025-11-10 18:21:09', null, 'png');
INSERT INTO `sys_media` VALUES ('299', '/public/uploads/2025-11-11/de5e7gyb7p04dwagiq.jpg', 'de5e7gyb7p04dwagiq.jpg', 'de5e7gyb7p04dwagiq.jpg', '58805', '.jpg', '2025-11-10 23:10:49', null, 'jpg');
INSERT INTO `sys_media` VALUES ('300', '/public/uploads/2025-11-11/de5e8jtmwpjknckfxl.jpg', 'de5e8jtmwpjknckfxl.jpg', 'de5e8jtmwpjknckfxl.jpg', '58805', '.jpg', '2025-11-10 23:12:13', null, 'jpg');
INSERT INTO `sys_media` VALUES ('301', '/public/uploads/2025-11-11/de5efi867848dsozfu.jpg', 'de5efi867848dsozfu.jpg', 'de5efi867848dsozfu.jpg', '58805', '.jpg', '2025-11-10 23:21:19', null, 'jpg');
INSERT INTO `sys_media` VALUES ('302', '/public/uploads/2025-11-11/de5efpbd8wrsmnftzy.jpg', 'de5efpbd8wrsmnftzy.jpg', 'de5efpbd8wrsmnftzy.jpg', '58805', '.jpg', '2025-11-10 23:21:34', null, 'jpg');
INSERT INTO `sys_media` VALUES ('303', '/public/uploads/2025-11-11/de5efynza6ugbtgxqa.jpg', 'de5efynza6ugbtgxqa.jpg', 'de5efynza6ugbtgxqa.jpg', '58805', '.jpg', '2025-11-10 23:21:54', null, 'jpg');
INSERT INTO `sys_media` VALUES ('304', '/public/uploads/2025-11-11/de5ei8ghvhkcqorz3f.jpg', 'de5ei8ghvhkcqorz3f.jpg', 'de5ei8ghvhkcqorz3f.jpg', '58805', '.jpg', '2025-11-10 23:24:52', null, 'jpg');
INSERT INTO `sys_media` VALUES ('305', '/public/uploads/2025-11-11/de5ejiq34e9ghqo8sz.png', 'de5ejiq34e9ghqo8sz.png', 'de5ejiq34e9ghqo8sz.png', '21119', '.png', '2025-11-10 23:26:33', null, 'png');
INSERT INTO `sys_media` VALUES ('306', '/public/uploads/2025-11-12/de6rwx9x4ghg962ku2.pem', 'de6rwx9x4ghg962ku2.pem', 'apiclient_cert.pem', '1537', '.pem', '2025-11-12 14:07:58', null, 'pem');
INSERT INTO `sys_media` VALUES ('307', '/public/uploads/2025-11-12/de6rx015jpr0uwqzmw.pem', 'de6rx015jpr0uwqzmw.pem', 'apiclient_key.pem', '1704', '.pem', '2025-11-12 14:08:04', null, 'pem');
INSERT INTO `sys_media` VALUES ('308', '/public/uploads/2025-11-12/de6s099hs8u8h63vep.pem', 'de6s099hs8u8h63vep.pem', 'apiclient_cert.pem', '1537', '.pem', '2025-11-12 14:12:19', null, 'pem');
INSERT INTO `sys_media` VALUES ('309', '/public/uploads/2025-11-12/de6s0af1f87sefa8to.pem', 'de6s0af1f87sefa8to.pem', 'apiclient_key.pem', '1704', '.pem', '2025-11-12 14:12:22', null, 'pem');
INSERT INTO `sys_media` VALUES ('310', '/public/uploads/2025-11-13/de759nqhwufohf4gyv.pem', 'de759nqhwufohf4gyv.pem', 'pub_key.pem', '451', '.pem', '2025-11-13 00:35:50', null, 'pem');
INSERT INTO `sys_media` VALUES ('311', '/public/uploads/2025-11-13/de75b24q9cfwvio3ub.pem', 'de75b24q9cfwvio3ub.pem', 'pub_key.pem', '451', '.pem', '2025-11-13 00:37:40', null, 'pem');
INSERT INTO `sys_media` VALUES ('312', '/public/uploads/2025-12-03/denyv7ep56nkuytjmz.png', 'denyv7ep56nkuytjmz.png', 'vben-logo.png', '361181', '.png', '2025-12-02 19:10:10', null, 'png');
INSERT INTO `sys_media` VALUES ('313', '/public/uploads/2025-12-03/denywlpqp7zky2atcv.png', 'denywlpqp7zky2atcv.png', 'vben-logo.png', '361181', '.png', '2025-12-02 19:12:00', null, 'png');
INSERT INTO `sys_media` VALUES ('314', '/public/uploads/2025-12-03/denyww3n8tq8scmmrr.png', 'denyww3n8tq8scmmrr.png', 'vben-logo.png', '361181', '.png', '2025-12-02 19:12:22', null, 'png');
INSERT INTO `sys_media` VALUES ('315', '/public/uploads/2025-12-03/denyxjt819m00lcc41.png', 'denyxjt819m00lcc41.png', 'vben-logo.png', '361181', '.png', '2025-12-02 19:13:14', null, 'png');
INSERT INTO `sys_media` VALUES ('316', '/public/uploads/2025-12-03/denyyj52n9aolxhzoq.png', 'denyyj52n9aolxhzoq.png', 'vben-logo.png', '361181', '.png', '2025-12-02 19:14:31', null, 'png');
INSERT INTO `sys_media` VALUES ('317', '/public/uploads/2025-12-03/denyzyi5w2ws6yjtnp.jpg', 'denyzyi5w2ws6yjtnp.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-02 19:16:23', null, 'jpg');
INSERT INTO `sys_media` VALUES ('318', '/public/uploads/2025-12-03/deod61pz6fbcgeezmf.jpg', 'deod61pz6fbcgeezmf.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 06:22:35', null, 'jpg');
INSERT INTO `sys_media` VALUES ('319', '/public/uploads/2025-12-03/deod7sa6ymmcocaex5.jpg', 'deod7sa6ymmcocaex5.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 06:24:52', null, 'jpg');
INSERT INTO `sys_media` VALUES ('320', '/public/uploads/2025-12-03/deod8g0cl9ogokffep.jpg', 'deod8g0cl9ogokffep.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 06:25:43', null, 'jpg');
INSERT INTO `sys_media` VALUES ('321', '/public/uploads/2025-12-03/deodh0jfxjd0k2svfr.jpg', 'deodh0jfxjd0k2svfr.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 06:36:55', null, 'jpg');
INSERT INTO `sys_media` VALUES ('322', '/public/uploads/2025-12-03/deodo1pvkci0teflua.jpg', 'deodo1pvkci0teflua.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 06:46:06', null, 'jpg');
INSERT INTO `sys_media` VALUES ('323', '/public/uploads/2025-12-03/deodoy7nq1wox5em0d.jpg', 'deodoy7nq1wox5em0d.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 06:47:17', null, 'jpg');
INSERT INTO `sys_media` VALUES ('324', '/public/uploads/2025-12-03/deodq5105mmktsweyh.jpg', 'deodq5105mmktsweyh.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 06:48:50', null, 'jpg');
INSERT INTO `sys_media` VALUES ('325', '/public/uploads/2025-12-03/deogrj0nzkigie9bgv.jpg', 'deogrj0nzkigie9bgv.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 09:11:42', null, 'jpg');
INSERT INTO `sys_media` VALUES ('326', '/public/uploads/2025-12-03/deohucxkdfsgxdq7by.jpg', 'deohucxkdfsgxdq7by.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 10:02:25', null, 'jpg');
INSERT INTO `sys_media` VALUES ('327', '/public/uploads/2025-12-04/deorhlt5qokkckruat.jpg', 'deorhlt5qokkckruat.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 17:35:57', null, 'jpg');
INSERT INTO `sys_media` VALUES ('328', '/public/uploads/2025-12-04/deoriny49hesvspq58.jpg', 'deoriny49hesvspq58.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 17:37:20', null, 'jpg');
INSERT INTO `sys_media` VALUES ('329', '/public/uploads/2025-12-04/deorrfv9b4yg5tfkhx.jpg', 'deorrfv9b4yg5tfkhx.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 17:48:47', null, 'jpg');
INSERT INTO `sys_media` VALUES ('330', '/public/uploads/2025-12-04/deorrvk2a1qkblzh3z.jpg', 'deorrvk2a1qkblzh3z.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 17:49:22', null, 'jpg');
INSERT INTO `sys_media` VALUES ('331', '/public/uploads/2025-12-04/deorsxp7rdpoyhctsd.jpg', 'deorsxp7rdpoyhctsd.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 17:50:45', null, 'jpg');
INSERT INTO `sys_media` VALUES ('332', '/public/uploads/2025-12-04/deos05hjckvwueau0f.jpg', 'deos05hjckvwueau0f.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 18:00:10', null, 'jpg');
INSERT INTO `sys_media` VALUES ('333', '/public/uploads/2025-12-04/deos0u7rz0jsvg5pkj.jpg', 'deos0u7rz0jsvg5pkj.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-03 18:01:04', null, 'jpg');
INSERT INTO `sys_media` VALUES ('334', '/public/uploads/2025-12-05/depm38fmplnoxyiioe.jpg', 'depm38fmplnoxyiioe.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:34:45', null, 'jpg');
INSERT INTO `sys_media` VALUES ('335', '/public/uploads/2025-12-05/depm39kp692wvfn2i6.jpg', 'depm39kp692wvfn2i6.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:34:47', null, 'jpg');
INSERT INTO `sys_media` VALUES ('336', '/public/uploads/2025-12-05/depm4li8yc40hq4lvu.jpg', 'depm4li8yc40hq4lvu.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:36:32', null, 'jpg');
INSERT INTO `sys_media` VALUES ('337', '/public/uploads/2025-12-05/depm4mgyfdbggrtmnb.jpg', 'depm4mgyfdbggrtmnb.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:36:34', null, 'jpg');
INSERT INTO `sys_media` VALUES ('338', '/public/uploads/2025-12-05/depm55ooy5wwdg5nfj.jpg', 'depm55ooy5wwdg5nfj.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:37:16', null, 'jpg');
INSERT INTO `sys_media` VALUES ('339', '/public/uploads/2025-12-05/depm568901awlvssft.jpg', 'depm568901awlvssft.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:37:17', null, 'jpg');
INSERT INTO `sys_media` VALUES ('340', '/public/uploads/2025-12-05/depm5zdt9cj4w3wizq.jpg', 'depm5zdt9cj4w3wizq.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:38:20', null, 'jpg');
INSERT INTO `sys_media` VALUES ('341', '/public/uploads/2025-12-05/depm5zxeblpwbzeq5k.jpg', 'depm5zxeblpwbzeq5k.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:38:21', null, 'jpg');
INSERT INTO `sys_media` VALUES ('342', '/public/uploads/2025-12-05/depm9pfjzbs4wyxjpn.jpg', 'depm9pfjzbs4wyxjpn.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-04 17:43:12', null, 'jpg');
INSERT INTO `sys_media` VALUES ('343', '/public/uploads/2025-12-05/deq7c91cjjhggjm4u8.jpg', 'deq7c91cjjhggjm4u8.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:13:55', null, 'jpg');
INSERT INTO `sys_media` VALUES ('344', '/public/uploads/2025-12-05/deq7czf8631kn1dqi0.jpg', 'deq7czf8631kn1dqi0.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:14:52', null, 'jpg');
INSERT INTO `sys_media` VALUES ('345', '/public/uploads/2025-12-05/deq7fxp5zktsnigyty.jpg', 'deq7fxp5zktsnigyty.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:18:44', null, 'jpg');
INSERT INTO `sys_media` VALUES ('346', '/public/uploads/2025-12-05/deq7gr6f4du42dh9lk.jpg', 'deq7gr6f4du42dh9lk.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:19:48', null, 'jpg');
INSERT INTO `sys_media` VALUES ('347', '/public/uploads/2025-12-05/deq7gs1el6igxanxmb.jpg', 'deq7gs1el6igxanxmb.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:19:50', null, 'jpg');
INSERT INTO `sys_media` VALUES ('348', '/public/uploads/2025-12-05/deq7gtaoiskkvbcsb2.jpg', 'deq7gtaoiskkvbcsb2.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:19:52', null, 'jpg');
INSERT INTO `sys_media` VALUES ('349', '/public/uploads/2025-12-05/deq7hd2ldjxwmz9bt4.jpg', 'deq7hd2ldjxwmz9bt4.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:20:35', null, 'jpg');
INSERT INTO `sys_media` VALUES ('350', '/public/uploads/2025-12-05/deq7iad10e0oyoffx2.jpg', 'deq7iad10e0oyoffx2.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:21:48', null, 'jpg');
INSERT INTO `sys_media` VALUES ('351', '/public/uploads/2025-12-05/deq7ihmevsukksx4yx.jpg', 'deq7ihmevsukksx4yx.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:22:04', null, 'jpg');
INSERT INTO `sys_media` VALUES ('352', '/public/uploads/2025-12-05/deq7it8ggx78wvjylx.jpg', 'deq7it8ggx78wvjylx.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:22:29', null, 'jpg');
INSERT INTO `sys_media` VALUES ('353', '/public/uploads/2025-12-05/deq7k4v40rlshbkloz.jpg', 'deq7k4v40rlshbkloz.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:24:13', null, 'jpg');
INSERT INTO `sys_media` VALUES ('354', '/public/uploads/2025-12-05/deq7kpe7l9ekkrdoof.jpg', 'deq7kpe7l9ekkrdoof.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:24:57', null, 'jpg');
INSERT INTO `sys_media` VALUES ('355', '/public/uploads/2025-12-05/deq7n007emykzxhtwy.jpg', 'deq7n007emykzxhtwy.jpg', 'soybean.jpg', '114241', '.jpg', '2025-12-05 10:27:57', null, 'jpg');
INSERT INTO `sys_media` VALUES ('356', '/public/uploads/2025-12-05/deq8fmqgeqic4msvl4.jpg', 'deq8fmqgeqic4msvl4.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-05 11:05:21', null, 'jpg');
INSERT INTO `sys_media` VALUES ('357', '/public/uploads/2025-12-05/deq8fnfgrgosrsbqdy.jpg', 'deq8fnfgrgosrsbqdy.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-05 11:05:22', null, 'jpg');
INSERT INTO `sys_media` VALUES ('358', '/public/uploads/2025-12-07/dercb839or8ofxcrzl.jpg', 'dercb839or8ofxcrzl.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-06 18:20:20', null, 'jpg');
INSERT INTO `sys_media` VALUES ('359', '/public/uploads/2025-12-07/dercn51scozgssm7jo.jpg', 'dercn51scozgssm7jo.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-06 18:35:54', null, 'jpg');
INSERT INTO `sys_media` VALUES ('360', '/public/uploads/2025-12-07/dercp7u8zz7weahf0s.jpg', 'dercp7u8zz7weahf0s.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-06 18:38:36', null, 'jpg');
INSERT INTO `sys_media` VALUES ('361', '/public/uploads/2025-12-10/dettg83yh8cwzfrofb.jpg', 'dettg83yh8cwzfrofb.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:11:31', null, 'jpg');
INSERT INTO `sys_media` VALUES ('362', '/public/uploads/2025-12-10/dettg9myrs8gb9d7yo.png', 'dettg9myrs8gb9d7yo.png', 'Cancel.png', '6717', '.png', '2025-12-09 16:11:34', null, 'png');
INSERT INTO `sys_media` VALUES ('363', '/public/uploads/2025-12-10/detti6c48ayssh6spi.jpg', 'detti6c48ayssh6spi.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:14:03', null, 'jpg');
INSERT INTO `sys_media` VALUES ('364', '/public/uploads/2025-12-10/dettjtiw1d3kdhcjtb.jpg', 'dettjtiw1d3kdhcjtb.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:16:12', null, 'jpg');
INSERT INTO `sys_media` VALUES ('365', '/public/uploads/2025-12-10/dettkbd9dxkg95vfdj.jpg', 'dettkbd9dxkg95vfdj.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:16:51', null, 'jpg');
INSERT INTO `sys_media` VALUES ('366', '/public/uploads/2025-12-10/dettl51hvox8qwx6ay.jpg', 'dettl51hvox8qwx6ay.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:17:56', null, 'jpg');
INSERT INTO `sys_media` VALUES ('367', '/public/uploads/2025-12-10/dettlpgr3eb85wq4ic.jpg', 'dettlpgr3eb85wq4ic.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:18:40', null, 'jpg');
INSERT INTO `sys_media` VALUES ('368', '/public/uploads/2025-12-10/dettnelartv08rrrjs.jpg', 'dettnelartv08rrrjs.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:20:53', null, 'jpg');
INSERT INTO `sys_media` VALUES ('369', '/public/uploads/2025-12-10/detto7uapgvwcz6mmo.jpg', 'detto7uapgvwcz6mmo.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:21:57', null, 'jpg');
INSERT INTO `sys_media` VALUES ('370', '/public/uploads/2025-12-10/detto8inlzpgbfuyri.png', 'detto8inlzpgbfuyri.png', 'Completed.png', '7842', '.png', '2025-12-09 16:21:58', null, 'png');
INSERT INTO `sys_media` VALUES ('371', '/public/uploads/2025-12-10/dettoj987px0ri4jqu.jpg', 'dettoj987px0ri4jqu.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:22:22', null, 'jpg');
INSERT INTO `sys_media` VALUES ('372', '/public/uploads/2025-12-10/dettok09pi7weovpyd.png', 'dettok09pi7weovpyd.png', 'Cancel.png', '6717', '.png', '2025-12-09 16:22:23', null, 'png');
INSERT INTO `sys_media` VALUES ('373', '/public/uploads/2025-12-10/dettoqrw8shsswahwq.jpg', 'dettoqrw8shsswahwq.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:22:38', null, 'jpg');
INSERT INTO `sys_media` VALUES ('374', '/public/uploads/2025-12-10/dettpol8mj6kh7gypi.jpg', 'dettpol8mj6kh7gypi.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:23:52', null, 'jpg');
INSERT INTO `sys_media` VALUES ('375', '/public/uploads/2025-12-10/dettpuqvat2cvo8ahq.png', 'dettpuqvat2cvo8ahq.png', 'Cancel.png', '6717', '.png', '2025-12-09 16:24:05', null, 'png');
INSERT INTO `sys_media` VALUES ('376', '/public/uploads/2025-12-10/dettpve5pq0od72c7l.jpg', 'dettpve5pq0od72c7l.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:24:06', null, 'jpg');
INSERT INTO `sys_media` VALUES ('377', '/public/uploads/2025-12-10/dettpw9yamy47nybr1.png', 'dettpw9yamy47nybr1.png', 'WaitStart.png', '8290', '.png', '2025-12-09 16:24:08', null, 'png');
INSERT INTO `sys_media` VALUES ('378', '/public/uploads/2025-12-10/detu6olf8rbctxcvhi.jpg', 'detu6olf8rbctxcvhi.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 16:46:04', null, 'jpg');
INSERT INTO `sys_media` VALUES ('379', '/public/uploads/2025-12-10/detuxbljjrmcrwphyx.jpg', 'detuxbljjrmcrwphyx.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 17:20:51', null, 'jpg');
INSERT INTO `sys_media` VALUES ('380', '/public/uploads/2025-12-10/detuxc8mpkq8zohowd.jpg', 'detuxc8mpkq8zohowd.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 17:20:53', null, 'jpg');
INSERT INTO `sys_media` VALUES ('381', '/public/uploads/2025-12-10/detuxcsrb6d4trvvdj.jpg', 'detuxcsrb6d4trvvdj.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 17:20:54', null, 'jpg');
INSERT INTO `sys_media` VALUES ('382', '/public/uploads/2025-12-10/detuxdfgttj09i8xtc.jpg', 'detuxdfgttj09i8xtc.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 17:20:55', null, 'jpg');
INSERT INTO `sys_media` VALUES ('383', '/public/uploads/2025-12-10/detuxe8jn4toogstjp.jpg', 'detuxe8jn4toogstjp.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 17:20:57', null, 'jpg');
INSERT INTO `sys_media` VALUES ('384', '/public/uploads/2025-12-10/detvoxa092scabqy3v.jpg', 'detvoxa092scabqy3v.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 17:56:54', null, 'jpg');
INSERT INTO `sys_media` VALUES ('385', '/public/uploads/2025-12-10/detvs9he4tpwgoz2nb.jpg', 'detvs9he4tpwgoz2nb.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 18:01:16', null, 'jpg');
INSERT INTO `sys_media` VALUES ('386', '/public/uploads/2025-12-10/detwluvqr0vg8uhm7j.png', 'detwluvqr0vg8uhm7j.png', 'WaitStart.png', '8290', '.png', '2025-12-09 18:39:55', null, 'png');
INSERT INTO `sys_media` VALUES ('387', '/public/uploads/2025-12-10/detwm1wvs904ljh1mj.jpg', 'detwm1wvs904ljh1mj.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-09 18:40:11', null, 'jpg');
INSERT INTO `sys_media` VALUES ('388', '/public/uploads/2025-12-13/dex7pgfsu2mkahz8mz.jpg', 'dex7pgfsu2mkahz8mz.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-13 15:59:49', null, 'jpg');
INSERT INTO `sys_media` VALUES ('389', '/public/uploads/2025-12-14/dex7pte8n6qcclvdgz.png', 'dex7pte8n6qcclvdgz.png', 'page-bg.png', '19993', '.png', '2025-12-13 16:00:18', null, 'png');
INSERT INTO `sys_media` VALUES ('390', '/public/uploads/2025-12-14/dex8ftbwr6r0ggergq.png', 'dex8ftbwr6r0ggergq.png', 'page-bg.png', '19993', '.png', '2025-12-13 16:34:15', null, 'png');
INSERT INTO `sys_media` VALUES ('391', '/public/uploads/2025-12-14/dex8g9yvtp0sjoydjj.jpg', 'dex8g9yvtp0sjoydjj.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-13 16:34:51', null, 'jpg');
INSERT INTO `sys_media` VALUES ('392', '/public/uploads/2025-12-14/dey0pr0nv2ywcbtf2w.jpg', 'dey0pr0nv2ywcbtf2w.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 14:43:45', null, 'jpg');
INSERT INTO `sys_media` VALUES ('393', '/public/uploads/2025-12-15/dey59i5yf8xctrdwqg.jpg', 'dey59i5yf8xctrdwqg.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:17:37', null, 'jpg');
INSERT INTO `sys_media` VALUES ('394', '/public/uploads/2025-12-15/dey5mhngdil40rcszs.jpg', 'dey5mhngdil40rcszs.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:34:35', null, 'jpg');
INSERT INTO `sys_media` VALUES ('395', '/public/uploads/2025-12-15/dey5mxawkcp0z7fvvo.png', 'dey5mxawkcp0z7fvvo.png', 'WaitStart.png', '8290', '.png', '2025-12-14 18:35:09', null, 'png');
INSERT INTO `sys_media` VALUES ('396', '/public/uploads/2025-12-15/dey5obr1c73ozywh0h.jpg', 'dey5obr1c73ozywh0h.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:36:59', null, 'jpg');
INSERT INTO `sys_media` VALUES ('397', '/public/uploads/2025-12-15/dey5p0dve4zwlorjye.png', 'dey5p0dve4zwlorjye.png', 'WaitStart.png', '8290', '.png', '2025-12-14 18:37:52', null, 'png');
INSERT INTO `sys_media` VALUES ('398', '/public/uploads/2025-12-15/dey5p2mb8qogmvgg4r.jpg', 'dey5p2mb8qogmvgg4r.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:37:57', null, 'jpg');
INSERT INTO `sys_media` VALUES ('399', '/public/uploads/2025-12-15/dey5padkfqswediman.jpg', 'dey5padkfqswediman.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:38:14', null, 'jpg');
INSERT INTO `sys_media` VALUES ('400', '/public/uploads/2025-12-15/dey5pwtclsvgyb5n59.jpg', 'dey5pwtclsvgyb5n59.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:39:03', null, 'jpg');
INSERT INTO `sys_media` VALUES ('401', '/public/uploads/2025-12-15/dey5q88g65lolww3re.jpg', 'dey5q88g65lolww3re.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:39:28', null, 'jpg');
INSERT INTO `sys_media` VALUES ('402', '/public/uploads/2025-12-15/dey5qijlko88znkhku.jpg', 'dey5qijlko88znkhku.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:39:50', null, 'jpg');
INSERT INTO `sys_media` VALUES ('403', '/public/uploads/2025-12-15/dey5qor52wtoy3g3sv.jpg', 'dey5qor52wtoy3g3sv.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:40:04', null, 'jpg');
INSERT INTO `sys_media` VALUES ('404', '/public/uploads/2025-12-15/dey5v2tfcf9giwn4lf.jpg', 'dey5v2tfcf9giwn4lf.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:45:48', null, 'jpg');
INSERT INTO `sys_media` VALUES ('405', '/public/uploads/2025-12-15/dey5w3cvcfngcbindx.jpg', 'dey5w3cvcfngcbindx.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 18:47:07', null, 'jpg');
INSERT INTO `sys_media` VALUES ('406', '/public/uploads/2025-12-15/dey6bmw56us4qmplpp.jpg', 'dey6bmw56us4qmplpp.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-14 19:07:25', null, 'jpg');
INSERT INTO `sys_media` VALUES ('407', '/public/uploads/2025-12-16/deyyn4mebsd0acktsm.jpg', 'deyyn4mebsd0acktsm.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-15 17:18:57', null, 'jpg');
INSERT INTO `sys_media` VALUES ('408', '/public/uploads/2025-12-16/deyyqyxho32gc2059e.png', 'deyyqyxho32gc2059e.png', 'huojian.png', '13505', '.png', '2025-12-15 17:23:58', null, 'png');
INSERT INTO `sys_media` VALUES ('409', '/public/uploads/2025-12-16/deyza071nnicwhlmt9.jpg', 'deyza071nnicwhlmt9.jpg', 'kook.jpg', '4731', '.jpg', '2025-12-15 17:48:50', null, 'jpg');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `p_id` bigint(20) DEFAULT '0' COMMENT '父菜单ID',
  `menu_type` char(10) DEFAULT '1' COMMENT '菜单类型（1目录 2菜单）',
  `name` varchar(50) DEFAULT '#' COMMENT '请求地址',
  `path` varchar(50) DEFAULT NULL,
  `label` varchar(50) DEFAULT NULL COMMENT '菜单名称',
  `component` varchar(100) DEFAULT NULL COMMENT '组件地址',
  `icon` varchar(100) DEFAULT NULL,
  `is_enable` tinyint(4) DEFAULT '2',
  `sort` int(10) DEFAULT NULL,
  `is_menu` tinyint(4) DEFAULT '2',
  `keep_alive` tinyint(4) DEFAULT '1',
  `is_hide` tinyint(4) DEFAULT '2' COMMENT '菜单状态（2显示 1隐藏）',
  `is_hide_tab` tinyint(4) DEFAULT '1',
  `link` varchar(200) DEFAULT NULL COMMENT '跳转',
  `is_iframe` tinyint(4) DEFAULT '1',
  `show_badge` tinyint(4) DEFAULT '1',
  `show_text_badge` tinyint(4) DEFAULT '1',
  `fixed_tab` tinyint(4) DEFAULT '1',
  `active_path` varchar(50) DEFAULT NULL,
  `is_full_page` tinyint(4) DEFAULT '1',
  `auth_name` varchar(50) DEFAULT NULL,
  `auth_label` varchar(50) DEFAULT NULL,
  `auth_icon` varchar(50) DEFAULT NULL,
  `auth_sort` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8 COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES ('1', '0', 'menu', '系统管理', '/system', 'System', '/index/index', 'solar:display-bold', '2', '2', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('2', '1', 'menu', '菜单管理', 'system_menu', 'SystemMenu', '/system/menu', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('3', '2', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '创建', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('4', '2', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('5', '2', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('7', '1', 'menu', '权限管理', 'system_permission', 'SystemPermission', '/system/permission', '', '2', '2', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('8', '7', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('9', '7', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('10', '7', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('11', '1', 'menu', '角色管理', 'system_role', 'SystemRole', '/system/role', '', '2', '3', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('12', '11', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('13', '11', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('14', '11', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('15', '11', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '设置菜单', 'setMenu', '', '4');
INSERT INTO `sys_menu` VALUES ('16', '11', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '设置权限', 'setPermission', '', '5');
INSERT INTO `sys_menu` VALUES ('17', '1', 'menu', '员工管理', 'system_manage', 'SystemManage', '/system/manage', '', '2', '4', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('18', '17', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('19', '17', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('20', '17', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('21', '0', 'menu', '配置中心', '/config', 'Config', '/index/index', 'solar:database-bold', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('22', '21', 'menu', '站点设置', 'config_site', 'ConfigSite', '/config/site', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('23', '21', 'menu', '微信配置', 'config_wechat', 'ConfigWechat', '/config/wechat', '', '2', '3', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('24', '21', 'menu', '协议政策', 'config_agreement', 'ConfigAgreement', '/config/agreement', '', '2', '2', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('25', '0', 'menu', '用户管理', '/user', 'User', '/index/index', 'solar:user-hands-bold', '2', '4', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('26', '25', 'menu', '用户列表', 'user_list', 'UserList', '/user/list', '', '2', '5', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('27', '26', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('28', '26', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('29', '26', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('30', '26', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '查看', 'view', '', '4');
INSERT INTO `sys_menu` VALUES ('31', '26', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '变更余额', 'changeBalance', '', '5');
INSERT INTO `sys_menu` VALUES ('32', '26', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', ' 预存充值', 'recharge', '', '6');
INSERT INTO `sys_menu` VALUES ('33', '25', 'menu', '威客列表', 'user_witkey', 'UserWitkey', '/user/witkey', '', '2', '4', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('34', '33', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('35', '33', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('36', '33', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('37', '33', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '查看', 'view', '', '4');
INSERT INTO `sys_menu` VALUES ('38', '33', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '变更佣金', 'changeCommission', '', '5');
INSERT INTO `sys_menu` VALUES ('39', '25', 'menu', '等级列表', 'user_level', 'UserLevel', '/user/level', '', '2', '3', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('40', '39', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('41', '39', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('42', '39', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('43', '0', 'menu', '商品管理', '/product', 'Product', '/index/index', 'solar:backpack-bold', '2', '5', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('44', '43', 'menu', '商品列表', 'product_list', 'ProductList', '/product/list', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('45', '44', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('46', '44', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('47', '44', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('48', '43', 'menu', '商品分类', 'product_category', 'ProductCategory', '/product/category', '', '2', '2', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('49', '48', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('50', '48', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('51', '48', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('52', '0', 'menu', '运营管理', '/operation', 'Operation', '/index/index', 'solar:gameboy-bold', '2', '6', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('53', '52', 'menu', '游戏列表', 'operation_game', 'OperationGame', '/operation/game', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('54', '53', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('55', '53', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('56', '53', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('57', '52', 'menu', '头衔列表', 'operation_title', 'OperationTitle', '/operation/title', '', '2', '2', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('58', '57', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('59', '57', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '编辑', 'edit', '', '2');
INSERT INTO `sys_menu` VALUES ('60', '57', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '删除', 'delete', '', '3');
INSERT INTO `sys_menu` VALUES ('61', '0', 'menu', '订单管理', '/order', 'Order', '/index/index', 'solar:reorder-bold', '2', '7', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('62', '61', 'menu', '订单列表', 'order_list', 'Order_list', '/order/list', '', '2', '3', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('65', '62', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '查看订单', 'view', '', '1');
INSERT INTO `sys_menu` VALUES ('66', '62', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '添加优惠', 'addDiscount', '', '2');
INSERT INTO `sys_menu` VALUES ('67', '62', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '确认付款', 'paid', '', '3');
INSERT INTO `sys_menu` VALUES ('68', '62', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '关闭订单', 'cancel', '', '4');
INSERT INTO `sys_menu` VALUES ('70', '62', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '退款订单', 'refund', '', '6');
INSERT INTO `sys_menu` VALUES ('72', '61', 'menu', '派单记录', 'order_distribute', 'OrderDistribute', '/order/distribute', '', '2', '2', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('73', '72', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '新增', 'create', '', '1');
INSERT INTO `sys_menu` VALUES ('74', '72', 'button', '', '', '', '', '', '1', '0', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '取消', 'cancel', '', '2');
INSERT INTO `sys_menu` VALUES ('77', '72', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '查看', 'view', '', '5');
INSERT INTO `sys_menu` VALUES ('78', '62', 'button', '', '', '', '', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '删除订单', 'delete', '', '7');
INSERT INTO `sys_menu` VALUES ('79', '61', 'menu', '报单结算', 'order_settlement', 'Order_Settlement', '/order/settlement', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('80', '0', 'menu', '财务管理', '/finance', 'Finance', '/index/index', 'solar:hand-money-bold', '2', '3', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('81', '80', 'menu', '提现申请', 'finance_withdraw', 'FinanceWithdraw', '/finance/withdraw', '', '2', '1', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('82', '80', 'menu', '充值记录', 'finance_recharge', 'FinanceRecharge', '/finance/recharge', '', '2', '2', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('84', '80', 'menu', '资金流水', 'finance_capital', 'FinanceCapital', '/finance/capital', '', '2', '5', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('87', '80', 'menu', '交易记录', 'finance_payment', 'FinancePayment', '/finance/payment', '', '1', '3', '2', '1', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');
INSERT INTO `sys_menu` VALUES ('88', '80', 'menu', '退款记录', 'finance_refund', 'FinanceRefund', '/finance/refund', '', '2', '4', '2', '2', '1', '1', '', '1', '1', '1', '1', '', '1', '', '', '', '1');

-- ----------------------------
-- Table structure for sys_order
-- ----------------------------
DROP TABLE IF EXISTS `sys_order`;
CREATE TABLE `sys_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `product_id` bigint(20) DEFAULT NULL,
  `specifications` varchar(255) DEFAULT NULL,
  `witkey_count` int(10) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `unit` char(10) DEFAULT NULL,
  `unit_price` decimal(10,2) DEFAULT NULL,
  `total_amount` decimal(10,2) DEFAULT NULL COMMENT '订单总额',
  `discount_amount` decimal(10,2) DEFAULT NULL COMMENT '优惠金额',
  `actual_amount` decimal(10,2) DEFAULT NULL COMMENT '实付金额',
  `pay_amount` decimal(10,2) DEFAULT NULL,
  `requirements` varchar(255) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `pay_mode` tinyint(4) DEFAULT NULL,
  `pay_status` tinyint(4) DEFAULT NULL,
  `pay_time` datetime DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `finish_time` datetime DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_code` (`code`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE,
  KEY `idx_product` (`product_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_order
-- ----------------------------
INSERT INTO `sys_order` VALUES ('4', '2025526023314104952', '19', '28', '地图:巴克什', '2', '1', '次', '105.00', '105.00', '100.00', '5.00', '5.00', '无', '2', '4', '2', '2025-12-20 17:44:09', null, null, '2025-09-05 18:33:15', null, null, null);

-- ----------------------------
-- Table structure for sys_payment
-- ----------------------------
DROP TABLE IF EXISTS `sys_payment`;
CREATE TABLE `sys_payment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `related_id` bigint(20) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `mode` tinyint(4) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_payment
-- ----------------------------
INSERT INTO `sys_payment` VALUES ('1', '4', 'PM2025122101440823626262', '1', '5.00', '4', '2025-12-20 17:44:09');

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `pid` bigint(20) DEFAULT NULL,
  `permission` text COMMENT '权限标识',
  `name` varchar(50) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL COMMENT '菜单名称',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=244 DEFAULT CHARSET=utf8 COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
INSERT INTO `sys_permission` VALUES ('1', '2', '/menu/create', '菜单创建', '创建菜单', '2024-12-15 08:15:22', '2024-12-15 09:17:39');
INSERT INTO `sys_permission` VALUES ('2', '0', '/menu', '菜单模块', '', '2024-12-15 11:21:48', '2024-12-21 11:32:41');
INSERT INTO `sys_permission` VALUES ('3', '2', '/menu/all', '全部菜单', '', '2024-12-15 12:44:22', '2025-12-02 10:51:41');
INSERT INTO `sys_permission` VALUES ('4', '2', '/menu/edit', '菜单删除', '', '2024-12-15 12:44:41', '2024-12-15 12:44:41');
INSERT INTO `sys_permission` VALUES ('5', '2', '/menu/delete', '菜单删除', '', '2024-12-15 12:44:59', '2024-12-15 12:44:59');
INSERT INTO `sys_permission` VALUES ('13', '0', '/role', '角色模块', '', '2024-12-18 10:05:00', '2024-12-18 10:08:47');
INSERT INTO `sys_permission` VALUES ('14', '13', '/role/list', '所有角色', '', '2024-12-18 10:05:19', '2025-12-02 10:54:17');
INSERT INTO `sys_permission` VALUES ('15', '13', '/role/all/menu', '角色所有菜单', '角色修改下的所有权限和菜单', '2024-12-18 10:06:40', '2025-04-03 18:23:16');
INSERT INTO `sys_permission` VALUES ('16', '13', '/role/all/permission', '角色所有的接口权限', '', '2024-12-18 10:07:01', '2025-04-03 18:23:08');
INSERT INTO `sys_permission` VALUES ('17', '13', '/role/create', '角色创建', '', '2024-12-18 10:07:24', '2024-12-18 10:07:24');
INSERT INTO `sys_permission` VALUES ('18', '13', '/role/edit', '角色修改', '', '2024-12-18 10:07:35', '2024-12-18 10:07:35');
INSERT INTO `sys_permission` VALUES ('19', '13', '/role/delete', '角色删除', '', '2024-12-18 10:07:46', '2024-12-18 10:07:46');
INSERT INTO `sys_permission` VALUES ('20', '0', '/manage', '管理模块', '', '2024-12-18 10:08:35', '2024-12-18 10:08:35');
INSERT INTO `sys_permission` VALUES ('21', '20', '/manage/delete', '管理删除', '', '2024-12-18 10:09:31', '2024-12-18 10:09:31');
INSERT INTO `sys_permission` VALUES ('22', '20', '/manage/all/role', '管理下所有角色', '', '2024-12-18 10:10:05', '2025-04-11 10:48:21');
INSERT INTO `sys_permission` VALUES ('23', '20', '/manage/role', '管理拥有的角色', '', '2024-12-18 10:10:19', '2024-12-18 10:10:19');
INSERT INTO `sys_permission` VALUES ('24', '20', '/manage/list', '管理列表', '', '2024-12-18 10:10:31', '2024-12-18 10:10:31');
INSERT INTO `sys_permission` VALUES ('25', '20', '/manage/create', '管理创建', '', '2024-12-18 10:10:43', '2024-12-18 10:10:43');
INSERT INTO `sys_permission` VALUES ('26', '20', '/manage/edit', '管理修改', '', '2024-12-18 10:10:58', '2024-12-18 10:10:58');
INSERT INTO `sys_permission` VALUES ('27', '0', '/permission', '权限模块', '', '2024-12-18 10:11:41', '2024-12-18 10:11:41');
INSERT INTO `sys_permission` VALUES ('28', '27', '/permission/all', '所有权限', '', '2024-12-18 10:12:11', '2024-12-18 10:12:11');
INSERT INTO `sys_permission` VALUES ('29', '27', '/permission/create', '权限创建', '', '2024-12-18 10:12:20', '2024-12-18 10:12:20');
INSERT INTO `sys_permission` VALUES ('30', '27', '/permission/edit', '权限修改', '', '2024-12-18 10:12:32', '2024-12-18 10:12:32');
INSERT INTO `sys_permission` VALUES ('31', '27', '/permission/delete', '权限删除', '', '2024-12-18 10:12:42', '2024-12-18 10:12:42');
INSERT INTO `sys_permission` VALUES ('32', '0', '/system', '系统模块', '', '2024-12-18 10:34:38', '2024-12-19 08:31:22');
INSERT INTO `sys_permission` VALUES ('33', '32', '/system/file', '文件设置', '', '2024-12-18 10:34:56', '2025-02-10 21:28:08');
INSERT INTO `sys_permission` VALUES ('34', '0', '/upload', '上传模块', '', '2024-12-18 10:35:29', '2024-12-19 08:31:28');
INSERT INTO `sys_permission` VALUES ('35', '34', '/upload/file', '小文件上传', '', '2024-12-18 10:36:17', '2025-12-02 19:09:55');
INSERT INTO `sys_permission` VALUES ('36', '34', '/upload/chunkIdentifier', '获取切片标识', '', '2024-12-18 10:36:31', '2024-12-18 10:36:31');
INSERT INTO `sys_permission` VALUES ('37', '34', '/upload/chunk', '上传切片', '', '2024-12-18 10:36:49', '2024-12-18 10:36:49');
INSERT INTO `sys_permission` VALUES ('38', '34', '/upload/chunk/merge', '上传切片合并', '', '2024-12-18 10:37:07', '2024-12-18 10:37:07');
INSERT INTO `sys_permission` VALUES ('42', '32', '/system/base', '基础设置', '', '2024-12-19 11:12:22', '2025-02-10 21:28:05');
INSERT INTO `sys_permission` VALUES ('43', '32', '/system/email', '邮箱设置', '', '2024-12-19 15:51:09', '2025-02-10 21:28:02');
INSERT INTO `sys_permission` VALUES ('45', '13', '/role/permissions', '角色拥有的接口权限', '', '2024-12-21 16:06:05', '2024-12-21 16:06:07');
INSERT INTO `sys_permission` VALUES ('46', '13', '/role/menus', '角色拥有的菜单', '', '2024-12-21 16:06:41', '2024-12-21 16:06:43');
INSERT INTO `sys_permission` VALUES ('47', '13', '/role/edit/menus', '修改角色拥有的菜单', '', '2024-12-21 16:38:30', '2025-04-03 18:23:56');
INSERT INTO `sys_permission` VALUES ('48', '13', '/role/edit/permissions', '修改角色拥有的接口权限', '', '2024-12-21 17:16:06', '2025-04-03 18:23:44');
INSERT INTO `sys_permission` VALUES ('51', '0', '/media', '附件媒体', '', '2024-12-22 13:30:31', '2024-12-22 13:30:31');
INSERT INTO `sys_permission` VALUES ('52', '51', '/media/list', '获取列表', '', '2024-12-22 13:30:49', '2024-12-22 13:31:10');
INSERT INTO `sys_permission` VALUES ('53', '51', '/media/delete', '附件删除', '', '2024-12-22 15:41:20', '2024-12-22 15:41:20');
INSERT INTO `sys_permission` VALUES ('54', '0', '/user', '用户模块', '', '2024-12-22 16:51:56', '2024-12-22 16:51:56');
INSERT INTO `sys_permission` VALUES ('55', '54', '/user/list', '用户列表', '', '2024-12-22 16:52:16', '2024-12-22 16:52:16');
INSERT INTO `sys_permission` VALUES ('56', '0', '/level', '等级模块', '', '2024-12-23 00:07:27', '2025-11-05 12:19:22');
INSERT INTO `sys_permission` VALUES ('57', '56', '/level/list', '等级列表', '', '2024-12-23 00:07:40', '2025-11-05 12:19:27');
INSERT INTO `sys_permission` VALUES ('58', '56', '/level/create', '等级添加', '', '2024-12-23 00:38:47', '2025-11-05 12:19:31');
INSERT INTO `sys_permission` VALUES ('59', '56', '/level/edit', '等级修改', '', '2024-12-23 01:06:42', '2025-11-05 12:19:35');
INSERT INTO `sys_permission` VALUES ('60', '56', '/level/delete', '等级删除', '', '2024-12-23 01:18:43', '2025-11-05 12:19:39');
INSERT INTO `sys_permission` VALUES ('61', '0', '/vip', '会员模块', '', '2024-12-23 01:47:25', '2024-12-23 01:47:25');
INSERT INTO `sys_permission` VALUES ('62', '61', '/vip/list', '会员列表', '', '2024-12-23 01:47:38', '2024-12-23 01:47:38');
INSERT INTO `sys_permission` VALUES ('63', '61', '/vip/create', '会员创建', '', '2024-12-23 01:48:02', '2024-12-23 01:48:02');
INSERT INTO `sys_permission` VALUES ('64', '61', '/vip/edit', '会员修改', '', '2024-12-23 01:48:32', '2024-12-23 01:48:32');
INSERT INTO `sys_permission` VALUES ('65', '61', '/vip/delete', '会员删除', '', '2024-12-23 01:48:47', '2024-12-23 01:48:47');
INSERT INTO `sys_permission` VALUES ('66', '0', '/game', '游戏模块', '', '2025-01-16 06:19:34', '2025-10-19 14:51:52');
INSERT INTO `sys_permission` VALUES ('67', '66', '/game/list', '游戏列表', '', '2025-01-16 06:19:51', '2025-10-19 14:52:05');
INSERT INTO `sys_permission` VALUES ('68', '66', '/game/create', '游戏创建', '', '2025-01-16 07:45:45', '2025-10-19 14:52:11');
INSERT INTO `sys_permission` VALUES ('69', '66', '/game/edit', '游戏编辑', '', '2025-01-16 10:03:46', '2025-10-19 14:52:17');
INSERT INTO `sys_permission` VALUES ('70', '66', '/game/delete', '游戏删除', '', '2025-01-16 10:12:22', '2025-10-19 14:52:24');
INSERT INTO `sys_permission` VALUES ('71', '0', '/title', '头衔模块', '', '2025-01-20 18:59:35', '2025-01-20 18:59:35');
INSERT INTO `sys_permission` VALUES ('72', '71', '/title/list', '头衔列表', '', '2025-01-20 19:00:01', '2025-01-20 19:00:01');
INSERT INTO `sys_permission` VALUES ('73', '71', '/title/create', '头衔创建', '', '2025-01-20 19:00:22', '2025-01-20 19:00:22');
INSERT INTO `sys_permission` VALUES ('74', '71', '/title/edit', '头衔编辑', '', '2025-01-20 19:00:38', '2025-01-20 19:00:38');
INSERT INTO `sys_permission` VALUES ('75', '71', '/title/delete', '头衔删除', '', '2025-01-20 19:00:53', '2025-01-20 19:00:53');
INSERT INTO `sys_permission` VALUES ('76', '0', '/product', '商品模块', '', '2025-01-27 16:25:28', '2025-10-11 16:20:50');
INSERT INTO `sys_permission` VALUES ('77', '76', '/product/create', '商品创建', '', '2025-01-27 16:25:45', '2025-10-11 16:20:53');
INSERT INTO `sys_permission` VALUES ('78', '76', '/product/list', '商品列表', '', '2025-01-28 18:36:01', '2025-10-11 16:20:56');
INSERT INTO `sys_permission` VALUES ('79', '76', '/product/edit', '商品编辑', '', '2025-01-29 03:32:24', '2025-10-11 16:20:59');
INSERT INTO `sys_permission` VALUES ('80', '76', '/product/delete', '商品删除', '', '2025-01-29 10:14:25', '2025-10-11 16:21:02');
INSERT INTO `sys_permission` VALUES ('86', '0', '/witkey', '威客模块', '', '2025-02-07 16:53:06', '2025-09-04 15:56:58');
INSERT INTO `sys_permission` VALUES ('87', '86', '/witkey/create', '创建申请', '', '2025-02-07 16:53:27', '2025-09-04 16:37:18');
INSERT INTO `sys_permission` VALUES ('88', '86', '/witkey/list', '威客列表', '', '2025-02-07 16:53:46', '2025-09-04 15:57:10');
INSERT INTO `sys_permission` VALUES ('89', '86', '/witkey/edit', '内容修改', '', '2025-02-07 16:54:12', '2025-09-04 17:24:38');
INSERT INTO `sys_permission` VALUES ('90', '86', '/witkey/punish/list', '处罚内容列表', '', '2025-02-07 16:54:37', '2025-09-04 18:16:54');
INSERT INTO `sys_permission` VALUES ('91', '86', '/witkey/punish/revoke', '处罚内容测回', '', '2025-02-07 16:55:07', '2025-09-04 18:32:35');
INSERT INTO `sys_permission` VALUES ('92', '0', '/withdraw', '提现模块', '', '2025-02-10 11:19:29', '2025-04-16 09:19:55');
INSERT INTO `sys_permission` VALUES ('93', '92', '/withdraw/list', '提现列表', '', '2025-02-10 11:19:43', '2025-04-16 09:20:05');
INSERT INTO `sys_permission` VALUES ('94', '92', '/withdraw/delete', '提现删除', '', '2025-02-10 11:20:02', '2025-04-16 09:20:12');
INSERT INTO `sys_permission` VALUES ('95', '92', '/withdraw/apply', '提现审核', '', '2025-02-10 11:20:15', '2025-04-16 09:20:18');
INSERT INTO `sys_permission` VALUES ('96', '92', '/withdraw/detail', '提现详情', '', '2025-02-10 15:45:40', '2025-04-16 09:20:25');
INSERT INTO `sys_permission` VALUES ('97', '32', '/system/withdraw', '提现设置', '', '2025-02-10 21:27:19', '2025-12-04 17:53:07');
INSERT INTO `sys_permission` VALUES ('100', '0', '/order', '订单模块', '', '2025-02-13 17:26:05', '2025-02-13 17:26:05');
INSERT INTO `sys_permission` VALUES ('101', '100', '/order/list', '订单列表', '', '2025-02-13 17:26:17', '2025-02-13 17:26:17');
INSERT INTO `sys_permission` VALUES ('103', '100', '/order/detail', '订单详情', '', '2025-02-18 10:56:41', '2025-02-18 10:56:41');
INSERT INTO `sys_permission` VALUES ('104', '100', '/order/refund', '订单退款', '', '2025-02-19 13:28:05', '2025-10-17 10:56:49');
INSERT INTO `sys_permission` VALUES ('105', '100', '/order/add/discount', '订单添加优惠', '', '2025-02-19 17:14:36', '2025-12-16 14:42:41');
INSERT INTO `sys_permission` VALUES ('124', '0', '/dict', '字典模块', '', '2025-04-05 18:41:23', '2025-04-06 07:07:06');
INSERT INTO `sys_permission` VALUES ('125', '124', '/dict/type/list', '字典类型列表', '', '2025-04-05 18:41:39', '2025-04-06 07:07:19');
INSERT INTO `sys_permission` VALUES ('126', '124', '/dict/type/create', '字典类型增加', '', '2025-04-06 07:06:54', '2025-04-06 07:06:54');
INSERT INTO `sys_permission` VALUES ('127', '124', '/dict/type/edit', '字典类型编辑', '', '2025-04-06 07:34:44', '2025-04-06 07:35:26');
INSERT INTO `sys_permission` VALUES ('128', '124', '/dict/type/delete', '字典类型删除', '', '2025-04-06 07:35:02', '2025-04-06 07:35:30');
INSERT INTO `sys_permission` VALUES ('129', '124', '/dict/data/list', '字典数据列表', '', '2025-04-06 09:22:57', '2025-04-06 09:22:57');
INSERT INTO `sys_permission` VALUES ('130', '124', '/dict/data/create', '字典数据添加', '', '2025-04-06 09:23:07', '2025-04-06 09:23:07');
INSERT INTO `sys_permission` VALUES ('131', '124', '/dict/data/edit', '字典数据编辑', '', '2025-04-06 09:23:26', '2025-04-06 09:23:26');
INSERT INTO `sys_permission` VALUES ('132', '124', '/dict/data/delete', '字典数据删除', '', '2025-04-06 09:23:38', '2025-04-06 09:23:38');
INSERT INTO `sys_permission` VALUES ('134', '100', '/order/logs', '订单操作日志', '', '2025-04-08 18:54:26', '2025-04-08 18:54:26');
INSERT INTO `sys_permission` VALUES ('138', '54', '/user/detail', '用户信息', '', '2025-04-14 10:26:53', '2025-04-14 10:26:53');
INSERT INTO `sys_permission` VALUES ('139', '54', '/user/create', '用户创建', '', '2025-04-14 12:59:08', '2025-04-14 12:59:08');
INSERT INTO `sys_permission` VALUES ('140', '54', '/user/edit', '用户编辑', '', '2025-04-14 14:00:02', '2025-04-14 14:00:02');
INSERT INTO `sys_permission` VALUES ('141', '54', '/user/change/balance', '余额变更', '', '2025-04-14 18:32:33', '2025-10-19 07:50:00');
INSERT INTO `sys_permission` VALUES ('142', '54', '/user/balance/list', '余额记录', '', '2025-04-14 18:48:01', '2025-12-08 18:00:34');
INSERT INTO `sys_permission` VALUES ('143', '54', '/user/delete', '用户删除', '', '2025-04-14 19:14:57', '2025-04-14 19:14:57');
INSERT INTO `sys_permission` VALUES ('144', '86', '/witkey/detail', '威客信息', '', '2025-04-15 14:26:18', '2025-09-04 16:49:21');
INSERT INTO `sys_permission` VALUES ('145', '86', '/witkey/punish', '添加处罚', '', '2025-04-15 15:01:16', '2025-09-04 18:07:01');
INSERT INTO `sys_permission` VALUES ('146', '54', '/user/logs', '用户操作日志', '', '2025-04-16 08:43:15', '2025-04-16 08:43:15');
INSERT INTO `sys_permission` VALUES ('151', '32', '/system/wechat/mini/program', '微信小程序配置', '', '2025-05-03 09:56:03', '2025-11-11 22:23:57');
INSERT INTO `sys_permission` VALUES ('153', '32', '/system/user', '用户设置', '', '2025-05-03 11:20:48', '2025-05-03 11:20:48');
INSERT INTO `sys_permission` VALUES ('165', '100', '/order/delete', '订单删除', '', '2025-05-25 18:29:36', '2025-05-25 18:29:36');
INSERT INTO `sys_permission` VALUES ('166', '100', '/order/paid', '订单确认收款', '', '2025-05-25 18:53:39', '2025-12-16 16:25:25');
INSERT INTO `sys_permission` VALUES ('172', '54', '/user/recharge', '预存充值', '', '2025-05-27 18:49:48', '2025-09-05 18:10:40');
INSERT INTO `sys_permission` VALUES ('173', '54', '/user/recharge/list', '预存列表', '', '2025-05-27 19:09:39', '2025-09-05 18:10:52');
INSERT INTO `sys_permission` VALUES ('175', '100', '/order/cancel', '关闭订单', '', '2025-08-11 18:29:09', '2025-12-17 13:31:07');
INSERT INTO `sys_permission` VALUES ('185', '86', '/witkey/logs', '操作日志', '', '2025-09-04 17:35:42', '2025-09-04 17:35:42');
INSERT INTO `sys_permission` VALUES ('186', '86', '/witkey/commission/list', '威客资金日志', '', '2025-09-04 17:41:13', '2025-12-11 11:15:59');
INSERT INTO `sys_permission` VALUES ('187', '86', '/witkey/change/commission', '变更佣金', '', '2025-09-04 18:46:41', '2025-10-10 09:07:26');
INSERT INTO `sys_permission` VALUES ('188', '86', '/witkey/delete', '威客删除', '', '2025-09-04 18:52:28', '2025-09-04 18:52:28');
INSERT INTO `sys_permission` VALUES ('190', '0', '/recharge', '预存充值', '', '2025-09-06 09:38:04', '2025-09-06 09:38:04');
INSERT INTO `sys_permission` VALUES ('191', '190', '/recharge/list', '充值列表', '', '2025-09-06 09:38:23', '2025-09-06 09:38:23');
INSERT INTO `sys_permission` VALUES ('192', '190', '/recharge/revoke', '撤回充值', '', '2025-09-06 09:54:52', '2025-09-06 09:54:52');
INSERT INTO `sys_permission` VALUES ('193', '32', '/system/contact', '联系设置', '', '2025-09-14 17:57:22', '2025-09-14 17:57:22');
INSERT INTO `sys_permission` VALUES ('202', '66', '/game/all', '所有游戏', '', '2025-10-09 17:44:07', '2025-10-19 14:52:30');
INSERT INTO `sys_permission` VALUES ('203', '54', '/user/select/options', '用户选择选项列表', '', '2025-10-09 18:09:09', '2025-10-09 18:09:09');
INSERT INTO `sys_permission` VALUES ('204', '71', '/title/all', '所有头衔', '', '2025-10-09 19:21:48', '2025-10-09 19:21:48');
INSERT INTO `sys_permission` VALUES ('206', '124', '/dict/data/all', '获取字典所有数据', '', '2025-10-11 16:46:02', '2025-10-11 16:46:02');
INSERT INTO `sys_permission` VALUES ('208', '54', '/user/add/witkey', '添加威客信息', '', '2025-10-19 09:30:03', '2025-10-19 09:30:03');
INSERT INTO `sys_permission` VALUES ('209', '54', '/user/edit/witkey', '修改威客信息', '', '2025-10-19 09:30:21', '2025-10-19 09:30:21');
INSERT INTO `sys_permission` VALUES ('210', '54', '/user/change/deposit', '变更用户押金', '', '2025-10-19 09:31:01', '2025-10-19 09:31:01');
INSERT INTO `sys_permission` VALUES ('211', '54', '/user/witkey/edit', '获取用户威客编辑信息', '', '2025-10-19 10:42:05', '2025-10-19 10:42:05');
INSERT INTO `sys_permission` VALUES ('212', '54', '/user/delete/witkey', '删除威客信息', '', '2025-10-19 12:03:37', '2025-10-19 12:03:37');
INSERT INTO `sys_permission` VALUES ('218', '0', '/bill', '账单模块', '', '2025-10-30 18:49:26', '2025-10-30 18:49:26');
INSERT INTO `sys_permission` VALUES ('219', '218', '/bill/list', '获取账单列表', '', '2025-10-30 18:49:40', '2025-10-30 18:49:40');
INSERT INTO `sys_permission` VALUES ('220', '56', '/level/all', '所有等级', '', '2025-11-05 13:47:22', '2025-11-05 13:47:22');
INSERT INTO `sys_permission` VALUES ('221', '32', '/system/wechat/pay', '微信支付配置', '', '2025-11-12 14:04:03', '2025-11-12 14:04:03');
INSERT INTO `sys_permission` VALUES ('231', '32', '/system/user/agreement', '用户协议', '', '2025-12-05 10:59:39', '2025-12-05 10:59:39');
INSERT INTO `sys_permission` VALUES ('232', '32', '/system/privacy/policy', '隐私协议', '', '2025-12-05 12:13:21', '2025-12-05 12:13:21');
INSERT INTO `sys_permission` VALUES ('233', '32', '/system/about/us', '关于我们', '', '2025-12-06 15:59:39', '2025-12-06 15:59:39');
INSERT INTO `sys_permission` VALUES ('234', '0', '/category', '分类模块', '', '2025-12-14 14:01:26', '2025-12-14 14:01:26');
INSERT INTO `sys_permission` VALUES ('235', '234', '/category/list', '分类列表', '', '2025-12-14 14:01:44', '2025-12-14 14:01:44');
INSERT INTO `sys_permission` VALUES ('236', '234', '/category/create', '分类新增', '', '2025-12-14 14:02:08', '2025-12-14 14:02:08');
INSERT INTO `sys_permission` VALUES ('237', '234', '/category/edit', '分类编辑', '', '2025-12-14 14:02:24', '2025-12-14 14:02:24');
INSERT INTO `sys_permission` VALUES ('238', '234', '/category/delete', '分类删除', '', '2025-12-14 14:02:37', '2025-12-14 14:02:37');
INSERT INTO `sys_permission` VALUES ('240', '0', '/distribute', '派单模块', '', '2025-12-22 14:28:04', '2025-12-22 14:28:04');
INSERT INTO `sys_permission` VALUES ('241', '240', '/distribute/create', '新增派单', '', '2025-12-22 14:28:16', '2025-12-22 14:28:16');
INSERT INTO `sys_permission` VALUES ('242', '240', '/distribute/cancel', '取消派单', '', '2025-12-22 14:28:45', '2025-12-22 14:28:45');
INSERT INTO `sys_permission` VALUES ('243', '240', '/distribute/list', '派单列表', '', '2025-12-22 18:46:57', '2025-12-22 18:46:57');

-- ----------------------------
-- Table structure for sys_product
-- ----------------------------
DROP TABLE IF EXISTS `sys_product`;
CREATE TABLE `sys_product` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pic` varchar(255) DEFAULT NULL,
  `type` tinyint(4) DEFAULT '1' COMMENT '1护航，2代肝',
  `category_id` bigint(20) DEFAULT NULL,
  `game_id` bigint(20) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `content` text,
  `witkey_count` int(10) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `unit` char(10) DEFAULT NULL,
  `sales_count` bigint(20) DEFAULT NULL,
  `rate` decimal(3,2) DEFAULT '5.00',
  `sku_type` tinyint(4) DEFAULT '2',
  `status` tinyint(4) DEFAULT '2' COMMENT '状态：1-禁用，2-启用',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`) USING BTREE,
  KEY `idx_name` (`name`) USING BTREE,
  KEY `idx_category` (`category_id`) USING BTREE,
  KEY `idx_game` (`game_id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_product
-- ----------------------------
INSERT INTO `sys_product` VALUES ('28', 'https://static.7b2.com/wp-content/uploads/2024/10/anime-7687171_1280.jpg?x-oss-process=image/resize,m_fill,h_388,w_516/sharpen,120/format,webp', '1', '38', '38', '护航趣味单', '测试商品护航内容', '<p>sdfasdf</p>', '2', '1000.00', 'ci', '1000', '4.00', '2', '2', '2025-10-13 14:06:49', '2025-12-15 13:42:58');
INSERT INTO `sys_product` VALUES ('29', '/public/uploads/2025-12-15/dey6bmw56us4qmplpp.jpg', '1', '38', '38', '测试商品', '撒旦发射点风格', '<p>豆腐干岁的法国</p>', '2', '1000.00', '次', '1000', '5.00', '2', '2', '2025-12-14 19:07:27', '2025-12-15 12:10:12');

-- ----------------------------
-- Table structure for sys_product_sku
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_sku`;
CREATE TABLE `sys_product_sku` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) NOT NULL COMMENT '商品ID',
  `code` varchar(100) DEFAULT NULL,
  `price` decimal(10,2) NOT NULL COMMENT '价格',
  `original_price` decimal(10,2) DEFAULT NULL,
  `status` tinyint(4) DEFAULT '2' COMMENT '状态：1-禁用，2-启用',
  `stock` int(11) DEFAULT NULL,
  `created_time` datetime NOT NULL,
  `updated_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sku_code` (`code`) USING BTREE,
  KEY `idx_sku_product` (`product_id`) USING BTREE,
  KEY `idx_sku_status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COMMENT='商品SKU表';

-- ----------------------------
-- Records of sys_product_sku
-- ----------------------------
INSERT INTO `sys_product_sku` VALUES ('40', '28', '地图:不限', '100.00', '150.00', '2', '999', '2025-10-13 14:06:49', '2025-10-13 14:06:49');
INSERT INTO `sys_product_sku` VALUES ('41', '28', '地图:航天', '105.00', '150.00', '2', '999', '2025-10-13 14:06:49', '2025-10-13 14:06:49');
INSERT INTO `sys_product_sku` VALUES ('42', '28', '地图:巴克什', '105.00', '150.00', '2', '999', '2025-10-13 14:06:49', '2025-10-13 14:06:49');

-- ----------------------------
-- Table structure for sys_product_sku_spec_relations
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_sku_spec_relations`;
CREATE TABLE `sys_product_sku_spec_relations` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sku_id` bigint(20) DEFAULT NULL,
  `spec_attr_id` bigint(20) DEFAULT NULL,
  `spec_value_id` bigint(20) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_sku_spec` (`sku_id`,`spec_attr_id`) USING BTREE,
  KEY `idx_sku_spec_relations_attr` (`spec_attr_id`) USING BTREE,
  KEY `idx_sku_spec_relations_sku` (`sku_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_product_sku_spec_relations
-- ----------------------------
INSERT INTO `sys_product_sku_spec_relations` VALUES ('75', '40', '56', '100', '2025-10-13 14:06:49');
INSERT INTO `sys_product_sku_spec_relations` VALUES ('76', '41', '56', '101', '2025-10-13 14:06:49');
INSERT INTO `sys_product_sku_spec_relations` VALUES ('77', '42', '56', '102', '2025-10-13 14:06:49');

-- ----------------------------
-- Table structure for sys_product_spec_attrs
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_spec_attrs`;
CREATE TABLE `sys_product_spec_attrs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_spec_attr_product` (`product_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_product_spec_attrs
-- ----------------------------
INSERT INTO `sys_product_spec_attrs` VALUES ('56', '28', '地图', '2025-10-13 14:06:49');

-- ----------------------------
-- Table structure for sys_product_spec_values
-- ----------------------------
DROP TABLE IF EXISTS `sys_product_spec_values`;
CREATE TABLE `sys_product_spec_values` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `spec_atts_id` bigint(20) DEFAULT NULL,
  `value` varchar(100) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_spec_values_attr` (`spec_atts_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_product_spec_values
-- ----------------------------
INSERT INTO `sys_product_spec_values` VALUES ('100', '56', '不限', null, '2025-10-13 14:06:49');
INSERT INTO `sys_product_spec_values` VALUES ('101', '56', '航天', null, '2025-10-13 14:06:49');
INSERT INTO `sys_product_spec_values` VALUES ('102', '56', '巴克什', null, '2025-10-13 14:06:49');

-- ----------------------------
-- Table structure for sys_punish
-- ----------------------------
DROP TABLE IF EXISTS `sys_punish`;
CREATE TABLE `sys_punish` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) DEFAULT NULL,
  `witkey_id` bigint(20) DEFAULT NULL,
  `manage_id` bigint(20) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `content` text,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_punish
-- ----------------------------
INSERT INTO `sys_punish` VALUES ('2', '202552802505410826036', '10', '1', '100.00', '0', '2025-05-27 18:50:54', '2025-05-27 18:50:54');
INSERT INTO `sys_punish` VALUES ('4', '202552803213610281170', '10', '1', '123.00', '飞过的痕迹地方各级', '2025-05-27 19:21:37', '2025-05-27 19:21:37');

-- ----------------------------
-- Table structure for sys_recharge
-- ----------------------------
DROP TABLE IF EXISTS `sys_recharge`;
CREATE TABLE `sys_recharge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `pay_type` tinyint(4) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_recharge
-- ----------------------------
INSERT INTO `sys_recharge` VALUES ('72', '23', 'CZ202512901541330020729', '50.00', '3', '2', '2025-12-08 17:54:14', null, null);

-- ----------------------------
-- Table structure for sys_refund
-- ----------------------------
DROP TABLE IF EXISTS `sys_refund`;
CREATE TABLE `sys_refund` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) DEFAULT NULL,
  `manage_id` bigint(20) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `mode` tinyint(4) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_refund
-- ----------------------------
INSERT INTO `sys_refund` VALUES ('1', '4', '1', '1.00', '4', '2025-12-20 17:07:18');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(30) NOT NULL COMMENT '角色名称',
  `code` varchar(50) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` tinyint(4) NOT NULL COMMENT '角色状态（2正常 1停用）',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '超级管理员', 'ADMIN', '后台所有权限管理', '2', '2024-12-13 16:55:02', '2025-04-16 15:25:25', null);
INSERT INTO `sys_role` VALUES ('2', '客服', 'R_TEST', 'sdfasdfsasdfsd', '2', '2024-12-14 18:09:00', '2025-11-03 21:06:50', null);
INSERT INTO `sys_role` VALUES ('4', '售后', 'R_TEST_TWO', 'fgsdfgsdfg ', '2', '2024-12-16 09:21:23', '2025-11-03 21:06:56', null);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL,
  `menu_id` bigint(20) NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES ('1', '1');
INSERT INTO `sys_role_menu` VALUES ('1', '2');
INSERT INTO `sys_role_menu` VALUES ('1', '3');
INSERT INTO `sys_role_menu` VALUES ('1', '4');
INSERT INTO `sys_role_menu` VALUES ('1', '5');
INSERT INTO `sys_role_menu` VALUES ('1', '7');
INSERT INTO `sys_role_menu` VALUES ('1', '8');
INSERT INTO `sys_role_menu` VALUES ('1', '9');
INSERT INTO `sys_role_menu` VALUES ('1', '10');
INSERT INTO `sys_role_menu` VALUES ('1', '11');
INSERT INTO `sys_role_menu` VALUES ('1', '12');
INSERT INTO `sys_role_menu` VALUES ('1', '13');
INSERT INTO `sys_role_menu` VALUES ('1', '14');
INSERT INTO `sys_role_menu` VALUES ('1', '15');
INSERT INTO `sys_role_menu` VALUES ('1', '16');
INSERT INTO `sys_role_menu` VALUES ('1', '17');
INSERT INTO `sys_role_menu` VALUES ('1', '18');
INSERT INTO `sys_role_menu` VALUES ('1', '19');
INSERT INTO `sys_role_menu` VALUES ('1', '20');
INSERT INTO `sys_role_menu` VALUES ('1', '21');
INSERT INTO `sys_role_menu` VALUES ('1', '22');
INSERT INTO `sys_role_menu` VALUES ('1', '23');
INSERT INTO `sys_role_menu` VALUES ('1', '24');
INSERT INTO `sys_role_menu` VALUES ('1', '25');
INSERT INTO `sys_role_menu` VALUES ('1', '26');
INSERT INTO `sys_role_menu` VALUES ('1', '27');
INSERT INTO `sys_role_menu` VALUES ('1', '28');
INSERT INTO `sys_role_menu` VALUES ('1', '29');
INSERT INTO `sys_role_menu` VALUES ('1', '30');
INSERT INTO `sys_role_menu` VALUES ('1', '31');
INSERT INTO `sys_role_menu` VALUES ('1', '32');
INSERT INTO `sys_role_menu` VALUES ('1', '33');
INSERT INTO `sys_role_menu` VALUES ('1', '34');
INSERT INTO `sys_role_menu` VALUES ('1', '35');
INSERT INTO `sys_role_menu` VALUES ('1', '36');
INSERT INTO `sys_role_menu` VALUES ('1', '37');
INSERT INTO `sys_role_menu` VALUES ('1', '38');
INSERT INTO `sys_role_menu` VALUES ('1', '39');
INSERT INTO `sys_role_menu` VALUES ('1', '40');
INSERT INTO `sys_role_menu` VALUES ('1', '41');
INSERT INTO `sys_role_menu` VALUES ('1', '42');
INSERT INTO `sys_role_menu` VALUES ('1', '43');
INSERT INTO `sys_role_menu` VALUES ('1', '44');
INSERT INTO `sys_role_menu` VALUES ('1', '45');
INSERT INTO `sys_role_menu` VALUES ('1', '46');
INSERT INTO `sys_role_menu` VALUES ('1', '47');
INSERT INTO `sys_role_menu` VALUES ('1', '48');
INSERT INTO `sys_role_menu` VALUES ('1', '49');
INSERT INTO `sys_role_menu` VALUES ('1', '50');
INSERT INTO `sys_role_menu` VALUES ('1', '51');
INSERT INTO `sys_role_menu` VALUES ('1', '52');
INSERT INTO `sys_role_menu` VALUES ('1', '53');
INSERT INTO `sys_role_menu` VALUES ('1', '54');
INSERT INTO `sys_role_menu` VALUES ('1', '55');
INSERT INTO `sys_role_menu` VALUES ('1', '56');
INSERT INTO `sys_role_menu` VALUES ('1', '57');
INSERT INTO `sys_role_menu` VALUES ('1', '58');
INSERT INTO `sys_role_menu` VALUES ('1', '59');
INSERT INTO `sys_role_menu` VALUES ('1', '60');
INSERT INTO `sys_role_menu` VALUES ('1', '61');
INSERT INTO `sys_role_menu` VALUES ('1', '62');
INSERT INTO `sys_role_menu` VALUES ('1', '65');
INSERT INTO `sys_role_menu` VALUES ('1', '66');
INSERT INTO `sys_role_menu` VALUES ('1', '67');
INSERT INTO `sys_role_menu` VALUES ('1', '68');
INSERT INTO `sys_role_menu` VALUES ('1', '70');
INSERT INTO `sys_role_menu` VALUES ('1', '72');
INSERT INTO `sys_role_menu` VALUES ('1', '73');
INSERT INTO `sys_role_menu` VALUES ('1', '74');
INSERT INTO `sys_role_menu` VALUES ('1', '77');
INSERT INTO `sys_role_menu` VALUES ('1', '78');
INSERT INTO `sys_role_menu` VALUES ('1', '79');
INSERT INTO `sys_role_menu` VALUES ('1', '80');
INSERT INTO `sys_role_menu` VALUES ('1', '81');
INSERT INTO `sys_role_menu` VALUES ('1', '82');
INSERT INTO `sys_role_menu` VALUES ('1', '84');
INSERT INTO `sys_role_menu` VALUES ('1', '87');
INSERT INTO `sys_role_menu` VALUES ('1', '88');
INSERT INTO `sys_role_menu` VALUES ('4', '1');
INSERT INTO `sys_role_menu` VALUES ('4', '2');
INSERT INTO `sys_role_menu` VALUES ('4', '3');
INSERT INTO `sys_role_menu` VALUES ('4', '4');
INSERT INTO `sys_role_menu` VALUES ('4', '5');
INSERT INTO `sys_role_menu` VALUES ('4', '7');
INSERT INTO `sys_role_menu` VALUES ('4', '8');
INSERT INTO `sys_role_menu` VALUES ('4', '9');
INSERT INTO `sys_role_menu` VALUES ('4', '10');
INSERT INTO `sys_role_menu` VALUES ('4', '11');
INSERT INTO `sys_role_menu` VALUES ('4', '12');
INSERT INTO `sys_role_menu` VALUES ('4', '13');
INSERT INTO `sys_role_menu` VALUES ('4', '14');
INSERT INTO `sys_role_menu` VALUES ('4', '15');
INSERT INTO `sys_role_menu` VALUES ('4', '16');

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission` (
  `permission_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_permission
-- ----------------------------
INSERT INTO `sys_role_permission` VALUES ('3', '2');
INSERT INTO `sys_role_permission` VALUES ('4', '2');
INSERT INTO `sys_role_permission` VALUES ('1', '12');
INSERT INTO `sys_role_permission` VALUES ('240', '1');
INSERT INTO `sys_role_permission` VALUES ('243', '1');
INSERT INTO `sys_role_permission` VALUES ('242', '1');
INSERT INTO `sys_role_permission` VALUES ('241', '1');
INSERT INTO `sys_role_permission` VALUES ('234', '1');
INSERT INTO `sys_role_permission` VALUES ('238', '1');
INSERT INTO `sys_role_permission` VALUES ('237', '1');
INSERT INTO `sys_role_permission` VALUES ('236', '1');
INSERT INTO `sys_role_permission` VALUES ('235', '1');
INSERT INTO `sys_role_permission` VALUES ('218', '1');
INSERT INTO `sys_role_permission` VALUES ('219', '1');
INSERT INTO `sys_role_permission` VALUES ('190', '1');
INSERT INTO `sys_role_permission` VALUES ('192', '1');
INSERT INTO `sys_role_permission` VALUES ('191', '1');
INSERT INTO `sys_role_permission` VALUES ('124', '1');
INSERT INTO `sys_role_permission` VALUES ('206', '1');
INSERT INTO `sys_role_permission` VALUES ('132', '1');
INSERT INTO `sys_role_permission` VALUES ('131', '1');
INSERT INTO `sys_role_permission` VALUES ('130', '1');
INSERT INTO `sys_role_permission` VALUES ('129', '1');
INSERT INTO `sys_role_permission` VALUES ('128', '1');
INSERT INTO `sys_role_permission` VALUES ('127', '1');
INSERT INTO `sys_role_permission` VALUES ('126', '1');
INSERT INTO `sys_role_permission` VALUES ('125', '1');
INSERT INTO `sys_role_permission` VALUES ('100', '1');
INSERT INTO `sys_role_permission` VALUES ('175', '1');
INSERT INTO `sys_role_permission` VALUES ('166', '1');
INSERT INTO `sys_role_permission` VALUES ('165', '1');
INSERT INTO `sys_role_permission` VALUES ('134', '1');
INSERT INTO `sys_role_permission` VALUES ('105', '1');
INSERT INTO `sys_role_permission` VALUES ('104', '1');
INSERT INTO `sys_role_permission` VALUES ('103', '1');
INSERT INTO `sys_role_permission` VALUES ('101', '1');
INSERT INTO `sys_role_permission` VALUES ('92', '1');
INSERT INTO `sys_role_permission` VALUES ('96', '1');
INSERT INTO `sys_role_permission` VALUES ('95', '1');
INSERT INTO `sys_role_permission` VALUES ('94', '1');
INSERT INTO `sys_role_permission` VALUES ('93', '1');
INSERT INTO `sys_role_permission` VALUES ('86', '1');
INSERT INTO `sys_role_permission` VALUES ('188', '1');
INSERT INTO `sys_role_permission` VALUES ('187', '1');
INSERT INTO `sys_role_permission` VALUES ('186', '1');
INSERT INTO `sys_role_permission` VALUES ('185', '1');
INSERT INTO `sys_role_permission` VALUES ('145', '1');
INSERT INTO `sys_role_permission` VALUES ('144', '1');
INSERT INTO `sys_role_permission` VALUES ('91', '1');
INSERT INTO `sys_role_permission` VALUES ('90', '1');
INSERT INTO `sys_role_permission` VALUES ('89', '1');
INSERT INTO `sys_role_permission` VALUES ('88', '1');
INSERT INTO `sys_role_permission` VALUES ('87', '1');
INSERT INTO `sys_role_permission` VALUES ('76', '1');
INSERT INTO `sys_role_permission` VALUES ('80', '1');
INSERT INTO `sys_role_permission` VALUES ('79', '1');
INSERT INTO `sys_role_permission` VALUES ('78', '1');
INSERT INTO `sys_role_permission` VALUES ('77', '1');
INSERT INTO `sys_role_permission` VALUES ('71', '1');
INSERT INTO `sys_role_permission` VALUES ('204', '1');
INSERT INTO `sys_role_permission` VALUES ('75', '1');
INSERT INTO `sys_role_permission` VALUES ('74', '1');
INSERT INTO `sys_role_permission` VALUES ('73', '1');
INSERT INTO `sys_role_permission` VALUES ('72', '1');
INSERT INTO `sys_role_permission` VALUES ('66', '1');
INSERT INTO `sys_role_permission` VALUES ('202', '1');
INSERT INTO `sys_role_permission` VALUES ('70', '1');
INSERT INTO `sys_role_permission` VALUES ('69', '1');
INSERT INTO `sys_role_permission` VALUES ('68', '1');
INSERT INTO `sys_role_permission` VALUES ('67', '1');
INSERT INTO `sys_role_permission` VALUES ('61', '1');
INSERT INTO `sys_role_permission` VALUES ('65', '1');
INSERT INTO `sys_role_permission` VALUES ('64', '1');
INSERT INTO `sys_role_permission` VALUES ('63', '1');
INSERT INTO `sys_role_permission` VALUES ('62', '1');
INSERT INTO `sys_role_permission` VALUES ('56', '1');
INSERT INTO `sys_role_permission` VALUES ('220', '1');
INSERT INTO `sys_role_permission` VALUES ('60', '1');
INSERT INTO `sys_role_permission` VALUES ('59', '1');
INSERT INTO `sys_role_permission` VALUES ('58', '1');
INSERT INTO `sys_role_permission` VALUES ('57', '1');
INSERT INTO `sys_role_permission` VALUES ('54', '1');
INSERT INTO `sys_role_permission` VALUES ('212', '1');
INSERT INTO `sys_role_permission` VALUES ('211', '1');
INSERT INTO `sys_role_permission` VALUES ('210', '1');
INSERT INTO `sys_role_permission` VALUES ('209', '1');
INSERT INTO `sys_role_permission` VALUES ('208', '1');
INSERT INTO `sys_role_permission` VALUES ('203', '1');
INSERT INTO `sys_role_permission` VALUES ('173', '1');
INSERT INTO `sys_role_permission` VALUES ('172', '1');
INSERT INTO `sys_role_permission` VALUES ('146', '1');
INSERT INTO `sys_role_permission` VALUES ('143', '1');
INSERT INTO `sys_role_permission` VALUES ('142', '1');
INSERT INTO `sys_role_permission` VALUES ('141', '1');
INSERT INTO `sys_role_permission` VALUES ('140', '1');
INSERT INTO `sys_role_permission` VALUES ('139', '1');
INSERT INTO `sys_role_permission` VALUES ('138', '1');
INSERT INTO `sys_role_permission` VALUES ('55', '1');
INSERT INTO `sys_role_permission` VALUES ('51', '1');
INSERT INTO `sys_role_permission` VALUES ('53', '1');
INSERT INTO `sys_role_permission` VALUES ('52', '1');
INSERT INTO `sys_role_permission` VALUES ('34', '1');
INSERT INTO `sys_role_permission` VALUES ('38', '1');
INSERT INTO `sys_role_permission` VALUES ('37', '1');
INSERT INTO `sys_role_permission` VALUES ('36', '1');
INSERT INTO `sys_role_permission` VALUES ('35', '1');
INSERT INTO `sys_role_permission` VALUES ('32', '1');
INSERT INTO `sys_role_permission` VALUES ('233', '1');
INSERT INTO `sys_role_permission` VALUES ('232', '1');
INSERT INTO `sys_role_permission` VALUES ('231', '1');
INSERT INTO `sys_role_permission` VALUES ('221', '1');
INSERT INTO `sys_role_permission` VALUES ('193', '1');
INSERT INTO `sys_role_permission` VALUES ('153', '1');
INSERT INTO `sys_role_permission` VALUES ('151', '1');
INSERT INTO `sys_role_permission` VALUES ('97', '1');
INSERT INTO `sys_role_permission` VALUES ('43', '1');
INSERT INTO `sys_role_permission` VALUES ('42', '1');
INSERT INTO `sys_role_permission` VALUES ('33', '1');
INSERT INTO `sys_role_permission` VALUES ('27', '1');
INSERT INTO `sys_role_permission` VALUES ('31', '1');
INSERT INTO `sys_role_permission` VALUES ('30', '1');
INSERT INTO `sys_role_permission` VALUES ('29', '1');
INSERT INTO `sys_role_permission` VALUES ('28', '1');
INSERT INTO `sys_role_permission` VALUES ('20', '1');
INSERT INTO `sys_role_permission` VALUES ('26', '1');
INSERT INTO `sys_role_permission` VALUES ('25', '1');
INSERT INTO `sys_role_permission` VALUES ('24', '1');
INSERT INTO `sys_role_permission` VALUES ('23', '1');
INSERT INTO `sys_role_permission` VALUES ('22', '1');
INSERT INTO `sys_role_permission` VALUES ('21', '1');
INSERT INTO `sys_role_permission` VALUES ('13', '1');
INSERT INTO `sys_role_permission` VALUES ('48', '1');
INSERT INTO `sys_role_permission` VALUES ('47', '1');
INSERT INTO `sys_role_permission` VALUES ('46', '1');
INSERT INTO `sys_role_permission` VALUES ('45', '1');
INSERT INTO `sys_role_permission` VALUES ('19', '1');
INSERT INTO `sys_role_permission` VALUES ('18', '1');
INSERT INTO `sys_role_permission` VALUES ('17', '1');
INSERT INTO `sys_role_permission` VALUES ('16', '1');
INSERT INTO `sys_role_permission` VALUES ('15', '1');
INSERT INTO `sys_role_permission` VALUES ('14', '1');
INSERT INTO `sys_role_permission` VALUES ('2', '1');
INSERT INTO `sys_role_permission` VALUES ('5', '1');
INSERT INTO `sys_role_permission` VALUES ('4', '1');
INSERT INTO `sys_role_permission` VALUES ('3', '1');
INSERT INTO `sys_role_permission` VALUES ('1', '1');

-- ----------------------------
-- Table structure for sys_settlement
-- ----------------------------
DROP TABLE IF EXISTS `sys_settlement`;
CREATE TABLE `sys_settlement` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) DEFAULT NULL,
  `witkey_id` bigint(20) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `amount` decimal(10,2) DEFAULT NULL,
  `commission` decimal(10,2) DEFAULT NULL,
  `deduction` decimal(10,2) DEFAULT NULL,
  `service_charge` decimal(10,2) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `settlement_time` datetime DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order` (`order_id`),
  KEY `idx_witkey` (`witkey_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_settlement
-- ----------------------------

-- ----------------------------
-- Table structure for sys_title
-- ----------------------------
DROP TABLE IF EXISTS `sys_title`;
CREATE TABLE `sys_title` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `game_id` bigint(20) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `service_percent` int(10) DEFAULT NULL,
  `pic` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_category` (`game_id`) USING BTREE,
  KEY `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_title
-- ----------------------------
INSERT INTO `sys_title` VALUES ('2', '38', '魔王', '100', '/public/uploads/2025-12-16/deyza071nnicwhlmt9.jpg', 'sdfsdf', '2025-02-04 13:04:02', '2025-12-15 17:49:00');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `level_id` bigint(20) DEFAULT NULL,
  `wx_union_id` varchar(255) DEFAULT NULL,
  `wx_open_id` varchar(255) DEFAULT NULL,
  `name` varchar(30) NOT NULL COMMENT '用户昵称',
  `phone` varchar(11) DEFAULT '' COMMENT '手机号码',
  `sex` tinyint(4) DEFAULT '3' COMMENT '用户性别（1男 2女 3未知）',
  `address` varchar(255) DEFAULT NULL,
  `birthday` datetime DEFAULT NULL,
  `avatar` varchar(255) DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `salt` char(10) DEFAULT NULL COMMENT '密码盐',
  `cover` varchar(255) DEFAULT NULL,
  `experience` bigint(20) DEFAULT NULL,
  `balance` decimal(10,2) DEFAULT NULL COMMENT '余额',
  `commission` decimal(10,2) DEFAULT NULL,
  `deposit` decimal(10,2) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` tinyint(4) DEFAULT '2' COMMENT '帐号状态（1停用,2正常）',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '软删除',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_name` (`name`) USING BTREE,
  KEY `idx_user_phone` (`phone`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('2', '1', null, null, '测试用户', '13677997053', '3', '北京', '2024-10-19 11:33:54', 'https://static.7b2.com/wp-content/uploads/2025/10/ice-cream-9864294_1920.jpg?x-oss-process=image/resize,m_fill,h_388,w_516/sharpen,120/format,webp', '40bd2624a8d37be3b24dc58df12cd159', 'YCzLID', '/public/uploads/2024-10-11/d4sxkleq7ugkzto4gf.png', null, '123.00', null, null, '1豆腐干岁的法国', '2', '2024-10-19 11:34:20', '2024-10-19 11:34:22', null, null);
INSERT INTO `sys_user` VALUES ('10', '1', null, 'oP9gA7eZ80V9amevfRTY1XcnfLt4', '解开了回家看了', '', '1', null, '1999-12-31 16:00:00', 'https://static.7b2.com/wp-content/uploads/2025/10/ice-cream-9864294_1920.jpg?x-oss-process=image/resize,m_fill,h_388,w_516/sharpen,120/format,webp', '', null, null, null, '1175.00', null, null, null, '2', '2025-09-07 11:23:09', '2025-05-02 11:23:09', null, null);
INSERT INTO `sys_user` VALUES ('12', '1', null, null, '测试客户', '12343244123', '3', null, '1989-12-31 16:00:00', 'https://static.7b2.com/wp-content/uploads/2025/10/ice-cream-9864294_1920.jpg?x-oss-process=image/resize,m_fill,h_388,w_516/sharpen,120/format,webp', '4d32108157feb69fdc2f45d28fa6c3be', 'T2g5uF', '/resource/public/uploads/2025-05-03/d9mhidr0mch0by8k4x.png', null, null, null, null, null, '2', '2025-09-07 11:27:03', '2025-09-05 11:27:03', null, null);
INSERT INTO `sys_user` VALUES ('13', '1', null, null, '豆腐干士大夫', '1231241212', '3', '北京', '1989-12-31 16:00:00', 'https://static.7b2.com/wp-content/uploads/2025/10/ice-cream-9864294_1920.jpg?x-oss-process=image/resize,m_fill,h_388,w_516/sharpen,120/format,webp', '', 'DtzJc3', '/resource/public/uploads/2025-05-03/d9mhidr0mch0by8k4x.png', null, '2344.00', null, null, '1豆腐干岁的法国1豆腐干岁的法国1豆腐干岁的法国1豆腐干岁的法国', '2', '2025-09-07 12:49:25', '2025-09-05 12:50:55', null, null);
INSERT INTO `sys_user` VALUES ('14', '1', null, null, '测试用户', '17777978987', '3', null, '1989-12-31 16:00:00', 'https://static.7b2.com/wp-content/uploads/2025/10/ice-cream-9864294_1920.jpg?x-oss-process=image/resize,m_fill,h_388,w_516/sharpen,120/format,webp', 'e2c9c8bc141feeab84585ed46c50a697', 'XHDnC8', '/resource/public/uploads/2025-05-03/d9mhidr0mch0by8k4x.png', null, null, null, null, null, '2', '2025-10-07 16:46:14', '2025-10-07 16:46:14', null, null);
INSERT INTO `sys_user` VALUES ('15', '5', null, null, '的风格', '17777942184', '1', '撒旦飞洒地方 ', '2006-12-16 16:00:00', 'https://static.7b2.com/wp-content/uploads/2025/09/sunrise-580379_1280.jpg', '3b129dd4ba21b25b8d4b010af9240b3f', 'W3njk0', '/resource/public/uploads/2025-05-03/d9mhidr0mch0by8k4x.png', '5351', '6831.00', null, null, 'sdfsadfgasdfasdf', '2', '2025-10-07 16:53:21', '2025-11-04 14:45:20', null, null);
INSERT INTO `sys_user` VALUES ('16', '3', null, null, '小二黑d很卷', '17777978384', '1', '', '2012-11-10 02:14:12', 'https://static.7b2.com/wp-content/uploads/2025/09/sunrise-580379_1280.jpg', 'a0735a1d7aab1693b86445e3a802bd80', 'EFRTTv', '/resource/public/uploads/2025-05-03/d9mhidr0mch0by8k4x.png', '600', null, null, null, 'vhfghb', '2', '2025-11-05 15:15:22', '2025-11-11 15:35:12', null, null);
INSERT INTO `sys_user` VALUES ('19', '1', null, 'o-qVh135NhWmum-8O9ThDfPT_NI4', '新用户8l04Xz', '19977990737', '3', null, '1999-12-31 16:00:00', '/resource/public/uploads/2025-05-03/d9mhid210av4yd9ak5.png', '24c6be3682dfdcae76cd84c28a5be818', 'nrzHJ3', '/resource/public/uploads/2025-05-03/d9mhidr0mch0by8k4x.png', '200', '65.00', null, null, null, '2', '2025-11-12 19:17:23', '2025-11-12 19:17:23', null, null);
INSERT INTO `sys_user` VALUES ('20', '1', null, null, '测试用户', '13097777860', '1', '', '1989-12-31 16:00:00', '/public/uploads/2025-12-05/depm9pfjzbs4wyxjpn.jpg', '6f732b824ff0020b3cbe98b788cc9960', 'abk16e', null, '200', null, null, null, '撒打发士大夫士大夫', '2', '2025-12-06 18:35:58', '2025-12-06 18:35:58', null, null);
INSERT INTO `sys_user` VALUES ('21', '1', null, null, '测智商', '13411132345', '2', '', '1989-12-31 16:00:00', '/public/uploads/2025-12-05/depm9pfjzbs4wyxjpn.jpg', '416fcb0e1569dc3ae6f4706471ceb250', 'AorVvT', null, '200', null, null, null, '撒打发士大夫', '2', '2025-12-06 18:37:55', '2025-12-06 18:37:55', null, null);
INSERT INTO `sys_user` VALUES ('23', '4', null, null, ' 策划告诉', '13097787864', '3', '[\"13\",\"1301\"]', '2001-05-15 16:00:00', '/public/uploads/2025-12-05/depm9pfjzbs4wyxjpn.jpg', '3af09e15078ad5c4c6a715e01ec6ac0e', 'hyEWGu', null, '200', '1450.00', null, null, '', '2', '2025-12-07 09:03:35', '2025-12-07 09:53:10', null, null);

-- ----------------------------
-- Table structure for sys_user_bill
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_bill`;
CREATE TABLE `sys_user_bill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `related_id` bigint(20) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `mode` tinyint(4) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=97 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_bill
-- ----------------------------
INSERT INTO `sys_user_bill` VALUES ('77', '19', '50', 'BL2025111321451068345453', '1', '0.01', '1', '2025-11-13 13:45:10');
INSERT INTO `sys_user_bill` VALUES ('78', '19', '51', 'BL2025111321471464905767', '1', '0.01', '1', '2025-11-13 13:47:14');
INSERT INTO `sys_user_bill` VALUES ('79', '19', '53', 'BL2025111402122528393347', '1', '0.01', '1', '2025-11-13 18:12:25');
INSERT INTO `sys_user_bill` VALUES ('80', '19', '54', 'BL2025111402220269779716', '1', '0.01', '1', '2025-11-13 18:22:03');
INSERT INTO `sys_user_bill` VALUES ('81', '19', '55', 'BL2025111402235429733093', '1', '0.01', '1', '2025-11-13 18:23:54');
INSERT INTO `sys_user_bill` VALUES ('82', '19', '56', 'BL2025111402245311068576', '1', '0.01', '1', '2025-11-13 18:24:54');
INSERT INTO `sys_user_bill` VALUES ('83', '19', '4', '202552602331410495411', '3', '5.00', '2', '2025-11-24 03:47:36');
INSERT INTO `sys_user_bill` VALUES ('84', '19', '4', '202552602331410495411', '3', '5.00', '2', '2025-11-25 17:09:10');
INSERT INTO `sys_user_bill` VALUES ('85', '19', '4', '202552602331410495411', '3', '5.00', '2', '2025-11-25 17:14:42');
INSERT INTO `sys_user_bill` VALUES ('86', '19', '4', '202552602331410495411', '3', '5.00', '2', '2025-11-25 17:25:33');
INSERT INTO `sys_user_bill` VALUES ('87', '19', '4', 'BL2025112602424255498152', '1', '0.10', '2', '2025-11-25 18:42:43');
INSERT INTO `sys_user_bill` VALUES ('88', '19', '4', 'BL202511260245336696318', '1', '0.10', '2', '2025-11-25 18:45:33');
INSERT INTO `sys_user_bill` VALUES ('89', '23', '70', 'BL202512818093511310010', '1', '100.00', '1', '2025-12-08 10:09:36');
INSERT INTO `sys_user_bill` VALUES ('90', '23', '71', 'BL202512901532297847091', '1', '100.00', '1', '2025-12-08 17:53:23');
INSERT INTO `sys_user_bill` VALUES ('91', '23', '72', 'BL20251290154133702627', '1', '50.00', '1', '2025-12-08 17:54:14');
INSERT INTO `sys_user_bill` VALUES ('92', '19', '4', 'BL2025121702374344507381', '3', '5.00', '2', '2025-12-16 18:37:44');
INSERT INTO `sys_user_bill` VALUES ('93', '19', '4', 'BL2025121721533963503814', '3', '5.00', '2', '2025-12-17 13:53:39');
INSERT INTO `sys_user_bill` VALUES ('94', '19', '4', 'BL2025122101071733497848', '2', '1.00', '1', '2025-12-20 17:07:18');
INSERT INTO `sys_user_bill` VALUES ('96', '19', '4', 'BL2025122101440814856277', '3', '5.00', '2', '2025-12-20 17:44:09');

-- ----------------------------
-- Table structure for sys_user_card
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_card`;
CREATE TABLE `sys_user_card` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `number` varchar(100) DEFAULT NULL,
  `nick_name` varchar(50) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_witkey` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_card
-- ----------------------------
INSERT INTO `sys_user_card` VALUES ('21', '15', '2', '士大夫', 's地方', '撒旦飞洒', '2025-10-26 07:56:25', '2025-10-26 07:56:25');

-- ----------------------------
-- Table structure for sys_user_experience
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_experience`;
CREATE TABLE `sys_user_experience` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `experience` bigint(20) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_experience
-- ----------------------------
INSERT INTO `sys_user_experience` VALUES ('1', '15', '373', '1', '2025-11-06 00:21:58');
INSERT INTO `sys_user_experience` VALUES ('2', '15', '406', '1', '2025-11-06 00:22:31');
INSERT INTO `sys_user_experience` VALUES ('3', '15', '462', '1', '2025-11-06 00:22:39');
INSERT INTO `sys_user_experience` VALUES ('4', '15', '102', '1', '2025-11-06 00:22:50');
INSERT INTO `sys_user_experience` VALUES ('5', '15', '372', '1', '2025-11-06 13:23:49');
INSERT INTO `sys_user_experience` VALUES ('6', '15', '153', '1', '2025-11-06 13:24:29');
INSERT INTO `sys_user_experience` VALUES ('7', '15', '270', '1', '2025-11-06 13:25:14');
INSERT INTO `sys_user_experience` VALUES ('8', '15', '144', '1', '2025-11-06 13:25:27');
INSERT INTO `sys_user_experience` VALUES ('10', '15', '38', '1', '2025-11-06 13:35:37');
INSERT INTO `sys_user_experience` VALUES ('11', '15', '101', '1', '2025-11-06 13:36:12');
INSERT INTO `sys_user_experience` VALUES ('12', '15', '97', '1', '2025-11-06 13:39:54');
INSERT INTO `sys_user_experience` VALUES ('13', '15', '488', '1', '2025-11-06 13:40:21');
INSERT INTO `sys_user_experience` VALUES ('14', '15', '454', '1', '2025-11-06 13:43:13');
INSERT INTO `sys_user_experience` VALUES ('15', '15', '162', '1', '2025-11-06 13:43:14');
INSERT INTO `sys_user_experience` VALUES ('16', '15', '104', '1', '2025-11-06 13:43:36');
INSERT INTO `sys_user_experience` VALUES ('17', '15', '460', '1', '2025-11-06 13:49:56');
INSERT INTO `sys_user_experience` VALUES ('18', '15', '347', '1', '2025-11-06 14:00:47');
INSERT INTO `sys_user_experience` VALUES ('19', '15', '258', '1', '2025-11-06 23:01:57');

-- ----------------------------
-- Table structure for sys_user_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_login_log`;
CREATE TABLE `sys_user_login_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=113 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_login_log
-- ----------------------------
INSERT INTO `sys_user_login_log` VALUES ('32', '15', '::1', '2025-10-23 14:33:05');
INSERT INTO `sys_user_login_log` VALUES ('33', '15', '::1', '2025-10-23 17:52:29');
INSERT INTO `sys_user_login_log` VALUES ('34', '15', '::1', '2025-10-23 17:52:58');
INSERT INTO `sys_user_login_log` VALUES ('35', '15', '::1', '2025-10-23 18:44:07');
INSERT INTO `sys_user_login_log` VALUES ('36', '15', '::1', '2025-10-25 07:09:53');
INSERT INTO `sys_user_login_log` VALUES ('37', '15', '::1', '2025-10-25 07:10:15');
INSERT INTO `sys_user_login_log` VALUES ('38', '15', '::1', '2025-10-26 07:46:10');
INSERT INTO `sys_user_login_log` VALUES ('39', '15', '::1', '2025-10-27 08:09:08');
INSERT INTO `sys_user_login_log` VALUES ('40', '15', '::1', '2025-10-28 11:13:53');
INSERT INTO `sys_user_login_log` VALUES ('41', '15', '::1', '2025-10-29 13:33:12');
INSERT INTO `sys_user_login_log` VALUES ('42', '15', '::1', '2025-10-30 22:31:40');
INSERT INTO `sys_user_login_log` VALUES ('43', '15', '::1', '2025-11-01 11:21:33');
INSERT INTO `sys_user_login_log` VALUES ('44', '15', '::1', '2025-11-02 11:25:15');
INSERT INTO `sys_user_login_log` VALUES ('45', '15', '::1', '2025-11-03 11:35:51');
INSERT INTO `sys_user_login_log` VALUES ('46', '15', '::1', '2025-11-04 14:45:41');
INSERT INTO `sys_user_login_log` VALUES ('47', '15', '::1', '2025-11-05 18:21:27');
INSERT INTO `sys_user_login_log` VALUES ('48', '15', '::1', '2025-11-06 13:22:48');
INSERT INTO `sys_user_login_log` VALUES ('49', '16', '::1', '2025-11-08 20:04:23');
INSERT INTO `sys_user_login_log` VALUES ('50', '16', '::1', '2025-11-08 20:07:21');
INSERT INTO `sys_user_login_log` VALUES ('51', '16', '::1', '2025-11-08 20:07:23');
INSERT INTO `sys_user_login_log` VALUES ('52', '16', '::1', '2025-11-08 20:07:25');
INSERT INTO `sys_user_login_log` VALUES ('53', '16', '::1', '2025-11-08 20:07:26');
INSERT INTO `sys_user_login_log` VALUES ('54', '16', '::1', '2025-11-08 20:07:44');
INSERT INTO `sys_user_login_log` VALUES ('55', '16', '::1', '2025-11-08 20:12:36');
INSERT INTO `sys_user_login_log` VALUES ('56', '16', '::1', '2025-11-08 20:12:37');
INSERT INTO `sys_user_login_log` VALUES ('57', '16', '::1', '2025-11-08 20:12:38');
INSERT INTO `sys_user_login_log` VALUES ('58', '16', '::1', '2025-11-08 20:12:55');
INSERT INTO `sys_user_login_log` VALUES ('59', '16', '::1', '2025-11-08 20:12:56');
INSERT INTO `sys_user_login_log` VALUES ('60', '16', '::1', '2025-11-08 20:12:58');
INSERT INTO `sys_user_login_log` VALUES ('61', '16', '::1', '2025-11-08 20:13:24');
INSERT INTO `sys_user_login_log` VALUES ('62', '16', '::1', '2025-11-08 20:13:27');
INSERT INTO `sys_user_login_log` VALUES ('63', '16', '::1', '2025-11-08 23:00:33');
INSERT INTO `sys_user_login_log` VALUES ('64', '16', '::1', '2025-11-08 23:00:38');
INSERT INTO `sys_user_login_log` VALUES ('65', '16', '::1', '2025-11-08 23:00:41');
INSERT INTO `sys_user_login_log` VALUES ('66', '16', '::1', '2025-11-08 23:00:44');
INSERT INTO `sys_user_login_log` VALUES ('67', '16', '::1', '2025-11-08 23:00:45');
INSERT INTO `sys_user_login_log` VALUES ('68', '16', '::1', '2025-11-08 23:01:21');
INSERT INTO `sys_user_login_log` VALUES ('69', '16', '::1', '2025-11-08 23:01:22');
INSERT INTO `sys_user_login_log` VALUES ('70', '16', '::1', '2025-11-08 23:01:38');
INSERT INTO `sys_user_login_log` VALUES ('71', '16', '::1', '2025-11-08 23:04:48');
INSERT INTO `sys_user_login_log` VALUES ('72', '16', '::1', '2025-11-08 23:06:38');
INSERT INTO `sys_user_login_log` VALUES ('73', '16', '::1', '2025-11-08 23:08:02');
INSERT INTO `sys_user_login_log` VALUES ('74', '16', '::1', '2025-11-08 23:09:31');
INSERT INTO `sys_user_login_log` VALUES ('75', '16', '::1', '2025-11-08 23:20:05');
INSERT INTO `sys_user_login_log` VALUES ('76', '16', '::1', '2025-11-08 23:20:17');
INSERT INTO `sys_user_login_log` VALUES ('77', '16', '::1', '2025-11-08 23:21:35');
INSERT INTO `sys_user_login_log` VALUES ('78', '16', '::1', '2025-11-08 23:56:52');
INSERT INTO `sys_user_login_log` VALUES ('79', '16', '::1', '2025-11-10 10:29:12');
INSERT INTO `sys_user_login_log` VALUES ('80', '16', '::1', '2025-11-11 10:29:53');
INSERT INTO `sys_user_login_log` VALUES ('81', '16', '::1', '2025-11-11 10:30:12');
INSERT INTO `sys_user_login_log` VALUES ('82', '16', '::1', '2025-11-11 10:31:01');
INSERT INTO `sys_user_login_log` VALUES ('83', '16', '::1', '2025-11-11 10:32:03');
INSERT INTO `sys_user_login_log` VALUES ('84', '16', '::1', '2025-11-11 10:32:42');
INSERT INTO `sys_user_login_log` VALUES ('85', '16', '::1', '2025-11-11 10:33:28');
INSERT INTO `sys_user_login_log` VALUES ('86', '16', '::1', '2025-11-11 10:33:54');
INSERT INTO `sys_user_login_log` VALUES ('87', '16', '::1', '2025-11-11 10:33:59');
INSERT INTO `sys_user_login_log` VALUES ('88', '16', '::1', '2025-11-11 10:39:20');
INSERT INTO `sys_user_login_log` VALUES ('89', '16', '::1', '2025-11-11 11:09:06');
INSERT INTO `sys_user_login_log` VALUES ('90', '16', '::1', '2025-11-11 11:10:23');
INSERT INTO `sys_user_login_log` VALUES ('91', '16', '::1', '2025-11-11 11:10:42');
INSERT INTO `sys_user_login_log` VALUES ('92', '16', '::1', '2025-11-11 11:12:36');
INSERT INTO `sys_user_login_log` VALUES ('93', '16', '::1', '2025-11-11 11:13:17');
INSERT INTO `sys_user_login_log` VALUES ('94', '16', '::1', '2025-11-11 11:14:20');
INSERT INTO `sys_user_login_log` VALUES ('95', '16', '::1', '2025-11-11 11:15:52');
INSERT INTO `sys_user_login_log` VALUES ('96', '16', '::1', '2025-11-11 11:16:19');
INSERT INTO `sys_user_login_log` VALUES ('97', '16', '::1', '2025-11-11 11:17:11');
INSERT INTO `sys_user_login_log` VALUES ('98', '16', '::1', '2025-11-11 11:27:56');
INSERT INTO `sys_user_login_log` VALUES ('99', '16', '::1', '2025-11-11 12:06:56');
INSERT INTO `sys_user_login_log` VALUES ('100', '16', '::1', '2025-11-11 14:54:04');
INSERT INTO `sys_user_login_log` VALUES ('101', '17', '::1', '2025-11-11 23:17:11');
INSERT INTO `sys_user_login_log` VALUES ('102', '18', '::1', '2025-11-11 23:23:11');
INSERT INTO `sys_user_login_log` VALUES ('103', '19', '::1', '2025-11-12 19:18:32');
INSERT INTO `sys_user_login_log` VALUES ('104', '19', '::1', '2025-11-12 19:18:58');
INSERT INTO `sys_user_login_log` VALUES ('105', '19', '::1', '2025-11-12 19:19:34');
INSERT INTO `sys_user_login_log` VALUES ('106', '19', '::1', '2025-11-12 19:46:53');
INSERT INTO `sys_user_login_log` VALUES ('107', '19', '::1', '2025-11-12 19:46:54');
INSERT INTO `sys_user_login_log` VALUES ('108', '19', '::1', '2025-11-12 19:54:34');
INSERT INTO `sys_user_login_log` VALUES ('109', '19', '::1', '2025-11-12 20:19:55');
INSERT INTO `sys_user_login_log` VALUES ('110', '19', '::1', '2025-11-12 20:31:42');
INSERT INTO `sys_user_login_log` VALUES ('111', '19', '::1', '2025-11-12 20:32:49');
INSERT INTO `sys_user_login_log` VALUES ('112', '19', '::1', '2025-11-12 20:33:29');

-- ----------------------------
-- Table structure for sys_vip
-- ----------------------------
DROP TABLE IF EXISTS `sys_vip`;
CREATE TABLE `sys_vip` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '会员ID',
  `name` varchar(50) NOT NULL COMMENT '会员名称',
  `level` int(11) NOT NULL COMMENT '会员等级',
  `icon` varchar(255) DEFAULT NULL COMMENT '会员图标',
  `price` decimal(10,2) NOT NULL COMMENT '会员价格',
  `duration` int(11) NOT NULL COMMENT '有效期(天)',
  `benefits` varchar(500) DEFAULT NULL COMMENT '会员权益说明',
  `discount` decimal(10,2) DEFAULT '1.00' COMMENT '折扣率(0-1)',
  `points_rate` decimal(10,2) DEFAULT '1.00' COMMENT '积分倍率',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_level` (`level`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户会员表';

-- ----------------------------
-- Records of sys_vip
-- ----------------------------
INSERT INTO `sys_vip` VALUES ('1', '测试会员士大夫', '1', '士大夫', '10.00', '30', '士大夫是', '20.00', '20.00', '2024-12-23 02:19:08', '2024-12-23 02:25:54', null);

-- ----------------------------
-- Table structure for sys_withdraw
-- ----------------------------
DROP TABLE IF EXISTS `sys_withdraw`;
CREATE TABLE `sys_withdraw` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) DEFAULT NULL,
  `manage_id` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `money` decimal(10,2) DEFAULT NULL,
  `settled_amount` decimal(10,2) DEFAULT NULL,
  `service_fee` decimal(10,2) DEFAULT NULL,
  `receipt_files` text,
  `receipt_num` varchar(255) DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL,
  `number` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_withdraw
-- ----------------------------
INSERT INTO `sys_withdraw` VALUES ('14', '20251027025616443453', null, '15', '100.00', '99.00', '1.00', null, null, '2', 's地方', '士大夫', '1', null, '2025-10-26 18:56:16', null);
INSERT INTO `sys_withdraw` VALUES ('15', '20251027025617438815', null, '15', '100.00', '99.00', '1.00', null, null, '2', 's地方', '士大夫', '1', null, '2025-10-26 18:56:17', null);
INSERT INTO `sys_withdraw` VALUES ('16', '20251027030338938145', null, '15', '100.00', '99.00', '1.00', null, null, '2', 's地方', '士大夫', '1', null, '2025-10-26 19:03:39', null);

-- ----------------------------
-- Table structure for sys_witkey
-- ----------------------------
DROP TABLE IF EXISTS `sys_witkey`;
CREATE TABLE `sys_witkey` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_id` bigint(20) DEFAULT NULL,
  `title_id` bigint(20) DEFAULT NULL,
  `game_id` bigint(20) DEFAULT NULL,
  `album` mediumtext,
  `rate` int(5) DEFAULT '5',
  `commission` decimal(10,2) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_witkey_category` (`game_id`) USING BTREE,
  KEY `idx_witkey_title` (`title_id`) USING BTREE,
  KEY `idx_witkey_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_witkey
-- ----------------------------
INSERT INTO `sys_witkey` VALUES ('13', '15', '2', '38', '[]', '5', null, null, '2025-10-09 16:34:10', '2025-10-19 13:24:38');
INSERT INTO `sys_witkey` VALUES ('14', '12', '2', '38', '[]', '5', null, null, '2025-10-09 18:30:11', '2025-10-19 11:49:39');
INSERT INTO `sys_witkey` VALUES ('16', '14', '2', '38', '[\"/public/uploads/2025-10-19/ddm80xend5monpdue8.png\"]', '5', null, null, '2025-10-19 10:23:28', '2025-10-19 10:23:28');
INSERT INTO `sys_witkey` VALUES ('17', '13', '2', '38', '[\"/public/uploads/2025-12-10/detwm1wvs904ljh1mj.jpg\"]', '3', '100.00', null, '2025-10-19 10:25:04', '2025-12-09 18:40:11');
