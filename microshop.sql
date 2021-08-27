SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `sort_order` mediumint(8) unsigned NOT NULL DEFAULT '1' COMMENT '顺序',
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '名称',
  `image_url` varchar(512) NOT NULL COMMENT '图片地址',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '内容',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否展示',
  `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX order_enabled (sort_order, enabled)
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of banner
-- ----------------------------
BEGIN;
INSERT INTO `banner` VALUES (10000001, 1, '合作 谁是你的菜', 'http://yanxuan.nosdn.127.net/65091eebc48899298171c2eb6696fe27.jpg', '合作 谁是你的菜', 1,1628337345000,1628337345000);
INSERT INTO `banner` VALUES (10000002, 2, '活动 美食节', 'http://yanxuan.nosdn.127.net/bff2e49136fcef1fd829f5036e07f116.jpg', '活动 美食节', 1,1628337345000,1628337345000);
INSERT INTO `banner` VALUES (10000003, 3, '活动 母亲节', 'http://yanxuan.nosdn.127.net/8e50c65fda145e6dd1bf4fb7ee0fcecc.jpg', '活动 母亲节', 1,1628337345000,1628337345000);
COMMIT;

-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
 `name` varchar(90) NOT NULL DEFAULT '' COMMENT '名称',
 `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父类ID',
 `sort_order` tinyint(1) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
 `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否展示',
 `img_url` varchar(255) NOT NULL COMMENT '图片URL',
 `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
 `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
 PRIMARY KEY (`id`),
 KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;

BEGIN;
INSERT INTO `category` VALUES (10000001, '居家', 0, 1, 1, 'http://yanxuan.nosdn.127.net/92357337378cce650797444bc107b0f7.jpg',1628337345000,1628337345000);
INSERT INTO `category` VALUES (10000002, '餐厨', 0, 2, 1, 'http://yanxuan.nosdn.127.net/f4ff8b3d5b0767d4e578575c1fd6b921.jpg',1628337345000,1628337345000);
INSERT INTO `category` VALUES (10000003, '饮食', 0, 3, 1, 'http://yanxuan.nosdn.127.net/dd6cc8a7e996936768db5634f12447ed.jpg',1628337345000,1628337345000);
INSERT INTO `category` VALUES (10000004, '服装', 0, 4, 1, 'http://yanxuan.nosdn.127.net/003e1d1289f4f290506ac2aedbd09d35.jpg',1628337345000,1628337345000);
COMMIT;
-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int(11) unsigned NOT NULL COMMENT '商品ID',
  `category_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类',
  `goods_sn` varchar(60) NOT NULL DEFAULT '' COMMENT '商品编码',
  `name` varchar(120) NOT NULL DEFAULT '' COMMENT '商品名称',
  `goods_brief` varchar(255) NOT NULL DEFAULT '' COMMENT '商品简介',
  `goods_desc` text COMMENT '商品描述',
  `is_on_sale` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否在售',
  `sort_order` smallint(4) unsigned NOT NULL DEFAULT '100' COMMENT '排序',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `is_new` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否新品',
  `primary_pic_url` varchar(255) NOT NULL COMMENT '商品主图',
  `list_pic_url` varchar(255) NOT NULL COMMENT '商品列表图',
  `retail_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '零售价格',
  `is_limited` tinyint(1) unsigned NOT NULL COMMENT '是否限价出售',
  `is_hot` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否热门商品',
  `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `goods_sn` (`goods_sn`),
  KEY `cat_id` (`category_id`),
  KEY `sort_order` (`sort_order`)
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;

BEGIN;
INSERT INTO `goods` VALUES (10000001, 10000001, '10000001', '母亲节礼物-舒适安睡组合', '安心舒适是最好的礼物', '不错的商品', 1, 1,0,1, 'http://yanxuan.nosdn.127.net/6f3e94fa4b44341bda5a73224d605896.jpg', 'http://yanxuan.nosdn.127.net/1f67b1970ee20fd572b7202da0ff705d.png', 2598.00, 1,1,1628337345000,1628337345000);
INSERT INTO `goods` VALUES (10000002, 10000001, '10000002', '可水洗舒柔丝羽绒枕', '超细纤维，蓬松轻盈回弹', '不错的商品', 1, 2,0,1, 'http://yanxuan.nosdn.127.net/3f2cc7a7e4472aa40c997e70efe6aeed.jpg', 'http://yanxuan.nosdn.127.net/a196b367f23ccfd8205b6da647c62b84.png', 1512.00, 1,1,1628337345000,1628337345000);
INSERT INTO `goods` VALUES (10000003, 10000002, '10000003', '100年传世珐琅锅 全家系列', '特质铸铁，大容量全家共享', '不错的商品', 1, 1,0,1, 'http://yanxuan.nosdn.127.net/9c9f47d3c321b96ad9c8d658ff4249e1.jpg', 'http://yanxuan.nosdn.127.net/c39d54c06a71b4b61b6092a0d31f2335.png', 398, 1,1,1628337345000,1628337345000);
INSERT INTO `goods` VALUES (10000004, 10000002, '10000004', '铸铁珐琅牛排煎锅', '沥油隔水，煎出外焦里嫩', '不错的商品', 1, 2,0,1, 'http://yanxuan.nosdn.127.net/61a56890eada2439eb4ff0a613cf8347.jpg', 'http://yanxuan.nosdn.127.net/619e46411ccd62e5c0f16692ee1a85a0.png', 149, 1,1,1628337345000,1628337345000);
INSERT INTO `goods` VALUES (10000005, 10000003, '10000005', '粽情乡思端午粽礼盒 640克', '五种口味，寄情端午', '不错的商品', 1, 1,0,1, 'http://yanxuan.nosdn.127.net/e34581a51939ceaf69ac71fe19c1cc45.jpg', 'http://yanxuan.nosdn.127.net/d1fd69cee4990f4de1109baef30efeeb.png', 68, 1,1,1628337345000,1628337345000);
INSERT INTO `goods` VALUES (10000006, 10000003, '10000006', '妙曲奇遇记曲奇礼盒 520克', '六种口味，酥香脆爽', '不错的商品', 1, 2,0,1, 'http://yanxuan.nosdn.127.net/b5b25363daed8cc56f8b455564538fd5.png', 'http://yanxuan.nosdn.127.net/8d228f767b136a67aaf2cbbf6deb46fa.png',73, 1,1,1628337345000,1628337345000);
INSERT INTO `goods` VALUES (10000007, 10000004, '10000007', '新生彩棉初衣礼盒（婴童）', '来自天然彩棉的礼物', '不错的商品', 1, 1,0,1, 'http://yanxuan.nosdn.127.net/d820f03d67e68071d30c922ea87eb023.png', 'http://yanxuan.nosdn.127.net/9aab9a0bf4fef8fe3dc8c732bc22d4b7.png',199, 1,1,1628337345000,1628337345000);
INSERT INTO `goods` VALUES (10000008, 10000004, '10000008', '格纹棉质褶皱娃娃裙（婴童）', '彼得潘领 内搭短裤', '不错的商品', 1, 2,0,1, 'http://yanxuan.nosdn.127.net/cdd7640a18e30a9477e361070aa1f8d5.jpg', 'http://yanxuan.nosdn.127.net/f82995ccb2a2f6beddd4ad794f5da2a1.png', 159, 1,1,1628337345000,1628337345000);
COMMIT;
-- Table structure for goods_gallery
-- ----------------------------
DROP TABLE IF EXISTS `goods_gallery`;
CREATE TABLE `goods_gallery` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `img_url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片URL',
  `img_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '图片描述',
  `img_type` tinyint(1) unsigned NOT NULL COMMENT '图片类型',
  `sort_order` int(11) unsigned NOT NULL DEFAULT '5' COMMENT '排序',
  `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `goods_id` (`goods_id`)
) ENGINE=InnoDB  AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;
BEGIN;
INSERT INTO `goods_gallery` VALUES (10000001, 10000001, 'http://yanxuan.nosdn.127.net/355efbcc32981aa3b7869ca07ee47dac.jpg', '不错', 1,  1, 1628337345000,1628337345000);
INSERT INTO `goods_gallery` VALUES (10000002, 10000001, 'http://yanxuan.nosdn.127.net/43e283df216881037b70d8b34f8846d3.jpg', '不错', 1,  2, 1628337345000,1628337345000);
INSERT INTO `goods_gallery` VALUES (10000003, 10000001, 'http://yanxuan.nosdn.127.net/43e283df216881037b70d8b34f8846d3.jpg', '不错', 2,  1, 1628337345000,1628337345000);
INSERT INTO `goods_gallery` VALUES (10000004, 10000001, 'http://yanxuan.nosdn.127.net/43e283df216881037b70d8b34f8846d3.jpg', '不错', 2, 2, 1628337345000,1628337345000);
INSERT INTO `goods_gallery` VALUES (10000005, 10000002, 'http://yanxuan.nosdn.127.net/c2f88baff6d3d9c954bf437649d26954.jpg', '不错', 1,  1, 1628337345000,1628337345000);
INSERT INTO `goods_gallery` VALUES (10000006, 10000002, 'http://yanxuan.nosdn.127.net/36176eb5337c5048cf4403b145f43bc4.jpg', '不错', 2,  1, 1628337345000,1628337345000);
COMMIT;

-- ----------------------------
-- Table structure for goods_attribute
-- ----------------------------
DROP TABLE IF EXISTS `goods_attribute`;
CREATE TABLE `goods_attribute` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `goods_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '属性名',
    `value` text NOT NULL DEFAULT '' COMMENT '属性值',
    `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `goods_id` (`goods_id`),
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;
BEGIN;
INSERT INTO `goods_attribute` VALUES (10000001, 10000001, '产地', '四川', 1628337345000,1628337345000);
INSERT INTO `goods_attribute` VALUES (10000002, 10000001, '材质', '丝绸', 1628337345000,1628337345000);
INSERT INTO `goods_attribute` VALUES (10000003, 10000002, '商家', '百货', 1628337345000,1628337345000);
INSERT INTO `goods_attribute` VALUES (10000004, 10000002, '颜色', '红色', 1628337345000,1628337345000);
INSERT INTO `goods_attribute` VALUES (10000005, 10000003, '商家', '百货', 1628337345000,1628337345000);
INSERT INTO `goods_attribute` VALUES (10000006, 10000003, '颜色', '红色', 1628337345000,1628337345000);
INSERT INTO `goods_attribute` VALUES (10000007, 10000004, '商家', '百货', 1628337345000,1628337345000);
INSERT INTO `goods_attribute` VALUES (10000008, 10000004, '颜色', '红色', 1628337345000,1628337345000);
COMMIT;
-- ----------------------------
-- Table structure for nideshop_product
-- ----------------------------
DROP TABLE IF EXISTS `goods_product`;
CREATE TABLE `goods_product` (
 `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
 `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
 `goods_number` mediumint(8) unsigned NOT NULL DEFAULT '100' COMMENT '商品数量',
 `retail_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '零售价',
 `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
 `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
 PRIMARY KEY (`id`),
 KEY `goods_id` (`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;

BEGIN;
INSERT INTO `goods_product` VALUES (10000001, 10000001, 100, 2598.00, 1628337345000,1628337345000);
INSERT INTO `goods_product` VALUES (10000002, 10000002, 100, 1512.00, 1628337345000,1628337345000);
INSERT INTO `goods_product` VALUES (10000003, 10000003, 100, 398, 1628337345000,1628337345000);
INSERT INTO `goods_product` VALUES (10000004, 10000004, 100, 149, 1628337345000,1628337345000);
INSERT INTO `goods_product` VALUES (10000005, 10000005, 100, 68, 1628337345000,1628337345000);
INSERT INTO `goods_product` VALUES (10000006, 10000006, 100, 73, 1628337345000,1628337345000);
INSERT INTO `goods_product` VALUES (10000007, 10000007, 100, 199, 1628337345000,1628337345000);
INSERT INTO `goods_product` VALUES (10000008, 10000008, 100, 159, 1628337345000,1628337345000);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
 `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
 `username` varchar(60) NOT NULL DEFAULT '' COMMENT '用户名',
 `password` varchar(32) NOT NULL DEFAULT '' COMMENT '用户密码',
 `gender` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
 `birthday` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户生日',
 `register_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '注册时间',
 `last_login_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '最近登录时间',
 `last_login_ip` varchar(255) NOT NULL DEFAULT '' COMMENT '登录IP',
 `nickname` varchar(60) NOT NULL COMMENT '用户昵称',
 `mobile` varchar(20) NOT NULL COMMENT '用户电话',
 `register_ip` varchar(255) NOT NULL DEFAULT '' COMMENT '注册IP',
 `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
 PRIMARY KEY (`id`),
 UNIQUE KEY `user_name` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
    `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
    `province_name` varchar(50) NOT NULL DEFAULT '' COMMENT '省',
    `city_name` varchar(50) NOT NULL DEFAULT '' COMMENT '市',
    `district_name` varchar(50) NOT NULL DEFAULT '' COMMENT '区',
    `address` varchar(120) NOT NULL DEFAULT '' COMMENT '详细地址',
    `mobile` varchar(60) NOT NULL DEFAULT '' COMMENT '电话',
    `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否默认地址',
    `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
 `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
 `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
 `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
 `goods_name` varchar(120) NOT NULL DEFAULT '' COMMENT '商品名称',
 `retail_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品价格',
 `number` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
 `checked` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否选中',
 `list_pic_url` varchar(255) NOT NULL DEFAULT '' COMMENT '商品图片',
 `goods_brief` varchar(255) NOT NULL DEFAULT '' COMMENT '商品简介',
 `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
 `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `order_sn` varchar(20) NOT NULL DEFAULT '' COMMENT '订单编号',
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `order_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态',
  `pay_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '支付状态',
  `consignee` varchar(60) NOT NULL DEFAULT '' COMMENT '下单人姓名',
  `province_name` varchar(50) NOT NULL DEFAULT '' COMMENT '省',
  `city_name` varchar(50) NOT NULL DEFAULT '' COMMENT '市',
  `district_name` varchar(50) NOT NULL DEFAULT '' COMMENT '区',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '详细地址',
  `mobile` varchar(60) NOT NULL DEFAULT '' COMMENT '电话',
  `pay_name` varchar(120) NOT NULL DEFAULT '' COMMENT '支付名称',
  `pay_id` tinyint(3) NOT NULL DEFAULT '0' COMMENT '支付ID',
  `actual_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '实际需要支付的金额',
  `order_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单总价',
  `confirm_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '收货时间',
  `pay_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '支付时间',
  `freight_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '配送费用',
  `callback_status` enum('true','false') DEFAULT 'true' COMMENT '状态',
  `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_sn` (`order_sn`),
  KEY `user_id` (`user_id`),
  KEY `order_status` (`order_status`),
  KEY `pay_status` (`pay_status`),
  KEY `pay_id` (`pay_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for order_goods
-- ----------------------------
DROP TABLE IF EXISTS `order_goods`;
CREATE TABLE `order_goods` (
    `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单明细ID',
    `order_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '订单ID',
    `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
    `goods_name` varchar(120) NOT NULL DEFAULT '' COMMENT '商品名称',
    `goods_brief` varchar(255) NOT NULL DEFAULT '' COMMENT '商品简介',
    `number` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '数量',
    `retail_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '单价',
    `list_pic_url` varchar(255) NOT NULL DEFAULT '' COMMENT '商品图片',
    `create_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `update_time` bigint(16) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `order_id` (`order_id`),
    KEY `goods_id` (`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4;