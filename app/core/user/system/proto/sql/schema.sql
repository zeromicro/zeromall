SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 选择 db:
USE
  mall_license;

-- ----------------------------
-- Table structure for license_key 注册密钥表 - 库存(预先生成, 待分配)
-- ----------------------------
DROP TABLE IF EXISTS `license_key_volume`;
CREATE TABLE `license_key_volume`
(
  `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`     tinyint(1)                                                    NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `public_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公钥',
  `secret_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '私钥',

  PRIMARY KEY (`id`),
  KEY `idx_updated_at` (`updated_at`),
  KEY `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_public_key` (`public_key`),
  UNIQUE KEY `uk_secret_key` (`secret_key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='注册码库存表';


-- ----------------------------
-- Table structure for 产品+密钥关系表
-- ----------------------------
DROP TABLE IF EXISTS `license_key_assign`;
CREATE TABLE `license_key_assign`
(
  `id`          bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`      tinyint(1)                                                    NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `public_key`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公钥',
  `product_id`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '产品id',
  `shop_id`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '店铺id',
  `consumer_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '消费者id',
  `order_no`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单号',


  PRIMARY KEY (`id`),
  KEY `idx_updated_at` (`updated_at`),
  KEY `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_public_key` (`public_key`),
  UNIQUE KEY `uk_license_consumer` (`public_key`, `consumer_id`),
  KEY `idx_consumer_id` (`consumer_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_shop_id` (`shop_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='注册码售卖记录';


-- ----------------------------
-- Table structure for 产品+密钥关系表
-- ----------------------------
DROP TABLE IF EXISTS `license_key_attr`;
CREATE TABLE `license_key_attr`
(
  `id`           bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at`   datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`   datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`   datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`       tinyint(1)                                                    NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `public_key`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公钥',
  `expired`      tinyint(1)                                                    NOT NULL DEFAULT '0' COMMENT '有效期： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `permission`   tinyint(1)                                                    NOT NULL DEFAULT '0' COMMENT '权限: 完整, 部分',
  `lifecycle`    tinyint(1)                                                    NOT NULL DEFAULT '0' COMMENT '付费方式: 按月续费, 永久有效, 年付费, 订阅： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `device_limit` tinyint(1)                                                    NOT NULL DEFAULT '0' COMMENT '设备限制数: 0, 不限制, >0, 限制数目',


  PRIMARY KEY (`id`),
  KEY `idx_updated_at` (`updated_at`),
  KEY `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_public_key` (`public_key`)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户授权表';



SET FOREIGN_KEY_CHECKS = 1;
