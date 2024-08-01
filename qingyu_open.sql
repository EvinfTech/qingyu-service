/*
 Navicat Premium Data Transfer

 Source Server         : .localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80100
 Source Host           : localhost:3306
 Source Schema         : qingyu_open

 Target Server Type    : MySQL
 Target Server Version : 80100
 File Encoding         : 65001

 Date: 01/08/2024 10:32:34
*/
CREATE DATABASE qingyu_open CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE qingyu_open;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account_record
-- ----------------------------
DROP TABLE IF EXISTS `account_record`;
CREATE TABLE `account_record`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `ouid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `change_type` tinyint UNSIGNED NOT NULL COMMENT '变更类型',
  `actual_money` int UNSIGNED NOT NULL COMMENT '实际金额',
  `present_money` int UNSIGNED NOT NULL COMMENT '赠送金额',
  `new_actual` int UNSIGNED NOT NULL COMMENT '新实际金额',
  `new_present` int UNSIGNED NOT NULL COMMENT '新赠送金额',
  `related_order_no` varchar(31) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关联订单号',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `creator_id` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of account_record
-- ----------------------------

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT ' 自增id',
  `user_ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户唯一标识',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `role_id` int NOT NULL COMMENT '用户角色',
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `del` int NOT NULL COMMENT '删除状态 0:正常  1:已删除',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `shop_id` int NOT NULL COMMENT '门店id',
  `level` int NOT NULL COMMENT '等级',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'd69d68737b6fe254c7a15994838c51f0de1966f7fea89cf67b925231b6715f11', 'admin', '', 'admin1', 1, '管理员', 0, '2023-12-04 14:30:11', '2024-05-11 16:15:31', 0, 0);

-- ----------------------------
-- Table structure for feedback
-- ----------------------------
DROP TABLE IF EXISTS `feedback`;
CREATE TABLE `feedback`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ouid',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `user_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户头像',
  `photo` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '上传图片',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '信息',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of feedback
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `pid` int NOT NULL COMMENT '父id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `p_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '父名称',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `grade` int NOT NULL COMMENT '菜单等级',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路径',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标',
  `hidden` int NOT NULL COMMENT '是否隐藏',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中文标题',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视图路径',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------

-- ----------------------------
-- Table structure for menu_role
-- ----------------------------
DROP TABLE IF EXISTS `menu_role`;
CREATE TABLE `menu_role`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` int NOT NULL COMMENT '角色id',
  `menu_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单id',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu_role
-- ----------------------------

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `order_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
  `order_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单名称',
  `original_price` bigint NOT NULL COMMENT '原价',
  `coupons` bigint NOT NULL COMMENT '优惠卷价格',
  `money` bigint NOT NULL COMMENT '实付价格',
  `pay_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付方式',
  `shop_id` int NOT NULL COMMENT '门店id',
  `shop_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门店名称',
  `shop_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门店头像',
  `user_ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ouid',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
  `user_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户头像',
  `pre_pay` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '预支付信息',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单状态',
  `gmt_success` datetime NULL DEFAULT NULL COMMENT '支付成功时间',
  `gmt_refund` datetime NULL DEFAULT NULL COMMENT '退款成功时间',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `site_detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订场信息',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `type` int NOT NULL COMMENT '订单类型(线上、线下)',
  `user_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户手机号',
  `supervise_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '结算状态',
  `check_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单核验码',
  `check_qr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '核验二维码',
  `reserve_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '预约姓名',
  `reserve_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '预约手机号',
  `reserve_date` date NOT NULL COMMENT '预约时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order
-- ----------------------------

-- ----------------------------
-- Table structure for pay
-- ----------------------------
DROP TABLE IF EXISTS `pay`;
CREATE TABLE `pay`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `order_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单编号',
  `pay_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付类型',
  `total_money` bigint NOT NULL COMMENT '订单总额',
  `money` bigint NOT NULL COMMENT '实付金额',
  `status` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付状态',
  `response_context` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '第三方返回的信息',
  `pay_form` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '从谁转',
  `pay_to` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '转到谁',
  `user_ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ouid',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
  `user_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户头像',
  `gmt_response` datetime NOT NULL COMMENT '第三方返回时间',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `out_order_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '外部订单号',
  `order_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单名称',
  `order_id` bigint NOT NULL COMMENT '订单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay
-- ----------------------------

-- ----------------------------
-- Table structure for recharge_rule
-- ----------------------------
DROP TABLE IF EXISTS `recharge_rule`;
CREATE TABLE `recharge_rule`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `origin_money` int UNSIGNED NOT NULL COMMENT '原金额',
  `present_money` int UNSIGNED NOT NULL COMMENT '赠送金额',
  `creator_id` int UNSIGNED NOT NULL,
  `mender_id` int UNSIGNED NOT NULL,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of recharge_rule
-- ----------------------------
INSERT INTO `recharge_rule` VALUES (1, 10000, 1000, 1, 1, '2024-07-30 11:44:00', '2024-07-30 11:44:00', 0);
INSERT INTO `recharge_rule` VALUES (2, 20000, 2500, 1, 1, '2024-07-30 11:44:07', '2024-07-30 11:44:07', 0);
INSERT INTO `recharge_rule` VALUES (3, 50000, 10000, 1, 1, '2024-07-30 11:44:16', '2024-07-30 11:44:16', 0);
INSERT INTO `recharge_rule` VALUES (4, 100000, 50000, 1, 1, '2024-07-30 11:44:31', '2024-07-30 11:44:31', 0);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `del` int NOT NULL COMMENT '删除状态  0:正常 1:删除',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `level` int NOT NULL COMMENT '等级',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------

-- ----------------------------
-- Table structure for shop
-- ----------------------------
DROP TABLE IF EXISTS `shop`;
CREATE TABLE `shop`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门店名称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `photo` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门店照片',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门店地址',
  `longitude` decimal(10, 6) NOT NULL COMMENT '经度',
  `latitude` decimal(10, 6) NOT NULL COMMENT '维度',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门店标签',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门店电话',
  `work_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '营业时间',
  `site_count` int NOT NULL COMMENT '场地数量',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `bottom_price` int NOT NULL COMMENT '最低价格',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `facility` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '场馆设施',
  `serve` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '场馆服务',
  `aggregate_amount` bigint NOT NULL COMMENT '总金额',
  `balance` bigint NOT NULL COMMENT '余额',
  `agreement` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '协议',
  `about_us` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '关于我们',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of shop
-- ----------------------------
INSERT INTO `shop` VALUES (1, '轻羽·羽毛球运动馆', 'avatar/17144414040.png', '[\"photo/17144415850.jpg\",\"photo/17144415870.jpeg\",\"photo/17144415900.jpeg\",\"photo/17144415950.jpg\"]', '山东省青岛市黄岛区开发区长江中路1577号', 120.192519, 35.975585, '[]', '13855556666', '08:00~22:00', 4, '2023-11-13 16:09:17', '2024-07-04 17:47:33', 3000, '这是一个场馆简介', '设施:轻钢龙骨搭建的框架\\n橡胶地板', '服务:24小时停车免费，24小时热水，淋浴免费', 2726000, 448000, '感谢您选择使用轻羽·羽毛球场馆的手机应用程序。我们十分重视您的隐私保护，并承诺严格遵守相关法律法规，保护您的个人信息安全。在您使用我们的应用程序时，请务必仔细阅读以下隐私协议条款：\n\n1. 信息收集：我们会收集您的个人信息，包括但不限于姓名、电话号码、电子邮件地址等，以便为您提供更好的服务和支持。\n\n2. 信息使用：我们会合理使用您的个人信息，包括但不限于与您联系、提供服务、改进产品、发送通知等用途。\n\n3. 信息保护：我们会采取合理的安全措施，保护您的个人信息不受未经授权的访问、使用或泄露。\n\n4. 信息共享：除非得到您的明确同意或法律要求，我们不会向第三方共享您的个人信息。\n\n5. 隐私权利：您有权随时访问、更正、删除您的个人信息，并有权撤销之前的同意。\n\n6. 隐私政策更新：我们保留随时更新隐私政策的权利，更新后的政策将在应用程序中公布。\n\n通过使用轻羽·羽毛球场馆的手机应用程序，即表示您同意接受本隐私协议的条款和条件。如果您有任何疑问或意见，请联系我们。感谢您的信任与支持！', '轻羽·羽毛球场馆是一家位于市中心的专业羽毛球场馆，拥有先进的设施和优质的服务。我们致力于为广大羽毛球爱好者提供一个优质的场地，让大家在轻松愉快的环境中享受运动乐趣。\n\n我们场馆拥有多个羽毛球场地，每个场地都配备了专业级的羽毛球场地和设施，确保运动员可以尽情发挥自己的技术水平。\n\n除了提供场地租赁服务，我们还提供羽毛球培训课程，由资深教练团队为您提供专业的指导和培训，帮助您提升技术水平。\n\n我们注重细节，每一个环节都力求完美，让每一位顾客都能感受到我们的用心和关爱。无论您是初学者还是职业选手，我们都会尽心尽力地为您提供最好的服务。\n\n欢迎来到轻羽·羽毛球场馆，让我们一起享受运动的快乐！');

-- ----------------------------
-- Table structure for shop_audit
-- ----------------------------
DROP TABLE IF EXISTS `shop_audit`;
CREATE TABLE `shop_audit`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '审核状态 Y:审核成功  N:审核中',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名字',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门店地址',
  `longitude` decimal(10, 6) NOT NULL COMMENT '经度',
  `latitude` decimal(10, 6) NOT NULL COMMENT '维度',
  `site_count` int NOT NULL COMMENT '场地数量',
  `gate_photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '门口照片',
  `indoor_photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '室内照片',
  `business_license` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '营业执照',
  `id_card_front` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '身份证正面',
  `id_card_back` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '身份证反面',
  `user_ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ouid',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '联系人名称',
  `user_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '联系人电话',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '审核状态 Y:审核成功  N:审核中',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `shop_id` int NOT NULL COMMENT '船舰门店后的门店id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of shop_audit
-- ----------------------------
INSERT INTO `shop_audit` VALUES (1, '店铺名字', '人民路888号', 11.111000, 33.330000, 3, 'http://oss.sportguider.com/file/common/170070438173.jpeg', 'http://oss.sportguider.com/file/common/170070438173.jpeg', 'http://oss.sportguider.com/file/common/170070438173.jpeg', 'http://oss.sportguider.com/file/common/170070438173.jpeg', 'http://oss.sportguider.com/file/common/170070438173.jpeg', 'aa', '联系人名称', '18888888888', '这个店铺周卡888，季卡999，年卡1000000', 'Y', '2023-12-19 09:36:55', '2023-12-19 09:36:55', 1);

-- ----------------------------
-- Table structure for shop_bill
-- ----------------------------
DROP TABLE IF EXISTS `shop_bill`;
CREATE TABLE `shop_bill`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `transaction_serial_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易流水号',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '流水类型 T提现 J结算',
  `money` bigint NOT NULL COMMENT '流水金额',
  `balance` bigint NOT NULL COMMENT '余额',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of shop_bill
-- ----------------------------

-- ----------------------------
-- Table structure for shop_error
-- ----------------------------
DROP TABLE IF EXISTS `shop_error`;
CREATE TABLE `shop_error`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `shop_id` int NOT NULL COMMENT '门店id',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '错误描述',
  `photos` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '照片',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `user_ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ouid',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of shop_error
-- ----------------------------

-- ----------------------------
-- Table structure for site
-- ----------------------------
DROP TABLE IF EXISTS `site`;
CREATE TABLE `site`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '场地名称',
  `shop_id` int NOT NULL COMMENT '场地id',
  `time_enum` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '时间枚举标签',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '状态',
  `free_price` int NOT NULL COMMENT '空闲时价格',
  `busy_price` int NOT NULL COMMENT '忙时价格',
  `weekend_busy` int NOT NULL COMMENT '周末忙时标记1:是忙时 0：不是忙时',
  `holiday_busy` int NOT NULL COMMENT '法定节假日忙时标记1:忙时 0:闲时',
  `data_busy` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '特殊忙时日期',
  `busy_time_enum` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '忙时时间段标签',
  `del` int NOT NULL COMMENT '删除标记 0:正常  1:删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of site
-- ----------------------------
INSERT INTO `site` VALUES (1, '1号场地', 1, '[9,10,11,12,13,14,15,16,17,18,19,20,21,22]', '2024-04-28 16:45:32', '2024-07-04 15:14:14', 'Y', 3000, 5000, 1, 1, '[\"2023-11-12\",\"2023-11-13\",\"2023-11-12\",\"2023-11-13\",\"2023-11-12\",\"2023-11-13\",\"2023-11-12\",\"2023-11-13\"]', '[17,18,19,20,21,22,23]', 0);
INSERT INTO `site` VALUES (2, '2号场地', 1, '[7,8,9,10,11,12,13,14,15,16,17,18,19,20]', '2024-04-28 16:25:06', '2024-05-16 11:34:37', 'Y', 3000, 5000, 1, 0, '[\"2023-11-12\",\"2023-11-13\"]', '[8,9,10,11,12,13,14,15,16,17,18,19,20]', 0);
INSERT INTO `site` VALUES (3, '3号场地', 1, '[8,9,10,11,12,13,14,15,16,17,18,19,20,21,22]', '2024-04-28 16:25:42', '2024-05-11 10:12:46', 'Y', 3000, 5000, 1, 1, '[\"2023-11-12\",\"2023-11-13\"]', '[1,2,3,4,5]', 0);
INSERT INTO `site` VALUES (4, '5号场地', 1, '[8,9,10,11,12,13,14,15,16,17,18,19,20,21,22]', '2024-04-28 15:45:54', '2024-05-11 10:12:55', 'Y', 9000, 18000, 1, 1, '[\"2023-11-12\",\"2023-11-13\"]', '[1,2,3,4,5]', 0);
INSERT INTO `site` VALUES (5, '6号场地', 1, '[8,9,10,11,12,13,14,15,16,17,18,19,20,21,22]', '2024-04-28 15:45:28', '2024-05-11 10:13:13', 'Y', 8000, 10000, 1, 1, '', '[2]', 0);

-- ----------------------------
-- Table structure for site_use
-- ----------------------------
DROP TABLE IF EXISTS `site_use`;
CREATE TABLE `site_use`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `site_id` int NOT NULL COMMENT '场地id',
  `shop_id` int NOT NULL COMMENT '门店id',
  `use_date` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '占用日期',
  `use_time_enum` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '占用时间标签',
  `user_ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ouid',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
  `user_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户头像',
  `user_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户手机号',
  `operator_ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作人员的ouid',
  `order_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单编号',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '占用状态',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `type` int NOT NULL COMMENT '订单类型(线上、线下)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of site_use
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `ouid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ouid',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户头像',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户手机号',
  `birthday` bigint NOT NULL COMMENT '生日',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户状态',
  `sex` bigint NOT NULL COMMENT '性别',
  `total_count` bigint NOT NULL COMMENT '统计参加次数',
  `total_length` bigint NOT NULL COMMENT '统计运动次数',
  `wx_openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '微信openid（不同微信程序，openid不同）',
  `wx_unionid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '微信unionid（不同微信程序，unionid相同）',
  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_update` datetime NOT NULL COMMENT '更新时间',
  `introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '个人简介',
  `vip_type` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '会员类型，1普通会员；2充值会员',
  `actual_balance` int UNSIGNED NOT NULL COMMENT '实际余额',
  `present_balance` int UNSIGNED NOT NULL COMMENT '赠送余额',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '14e3ce7d8d8c587c78516c6cc09b723659e58fdda75c3bc0d002abb0179100c1', '演示账户', 'avatar/17169692970.png', '18888888888', 20240505, 'Y', 2, 0, 0, '', '', '2024-04-23 15:39:36', '2024-05-29 15:55:05', '', 2, 0, 1800);

SET FOREIGN_KEY_CHECKS = 1;
