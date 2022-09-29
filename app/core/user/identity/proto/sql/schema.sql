SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- 选择 db:
USE
mall_user;

-- ----------------------------
-- Table structure for user_identity: 用户基础信息表
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
  `id`             bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`         tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `user_id`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',
  `username`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `email`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile_no`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `mobile_country` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号国家码',

  PRIMARY KEY (`id`),
  KEY              `idx_updated_at` (`updated_at`),
  KEY              `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户信息表';


-- ----------------------------
-- Table structure for user_authn: 用户登录方式
-- ----------------------------
DROP TABLE IF EXISTS `user_authn`;
CREATE TABLE `user_authn`
(
  `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`     tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `unique_via` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '唯一可登录信息: 可以是: 用户名, 手机号<国家码-手机号>, 邮箱, 等',
  `auth_type`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录方式：1=用户名, 2=邮箱, 3=短信 4=2fa, 5=oauth',
  `user_id`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',
  `password`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',

  PRIMARY KEY (`id`),
  KEY          `idx_updated_at` (`updated_at`),
  KEY          `idx_created_at` (`created_at`),
  KEY          `idx_auth_type` (`auth_type`),

  UNIQUE KEY `uk_unique_via` (`unique_via`),
  UNIQUE KEY `uk_unique_via_type_id` (`unique_via`, `auth_type`, `user_id`)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户登录方式表';


-- ----------------------------
-- Table structure for user_authn: 用户登录方式
-- ----------------------------
DROP TABLE IF EXISTS `user_authn_oauth`;
CREATE TABLE `user_authn_oauth`
(
  `id`             bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`         tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `oauth_token`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '唯一可登录信息: 可以是: 用户名, 手机号<国家码-手机号>, 邮箱, 等',
  `oauth_provider` tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录方式：1=google, 2=twitter, 3=facebook, 4=apple, 5=...',
  `user_id`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',

  PRIMARY KEY (`id`),
  KEY              `idx_updated_at` (`updated_at`),
  KEY              `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_oauth_token_provider` (`oauth_token`, `oauth_provider`, `user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户登录方式表';


-- ----------------------------
-- Table structure for user_authn: 用户登录方式
-- ----------------------------
DROP TABLE IF EXISTS `user_authn_log`;
CREATE TABLE `user_authn_log`
(
  `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`     tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `auth_type`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录方式：1=用户名, 2=邮箱, 3=短信 4=2fa, 5=oauth',
  `platform`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录方式：1=android, macos, ios, windows, pc, mobile',
  `user_id`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',
  `ip`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'ip',
  `device`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '设备信息',

  PRIMARY KEY (`id`),
  KEY          `idx_updated_at` (`updated_at`),
  KEY          `idx_created_at` (`created_at`)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户登录历史表';


-- ----------------------------
-- Table structure for user_identity: 用户基础信息表
-- ----------------------------
DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `user_profile`
(
  `id`          bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`      tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `user_id`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',
  `nickname`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',

  `real_name`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `first_name`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `middle_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `last_name`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `sex`         tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `birth`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',

  `address`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `avatar`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像URL',


  PRIMARY KEY (`id`),
  KEY           `idx_updated_at` (`updated_at`),
  KEY           `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户信息表';


-- ----------------------------
-- Table structure for user_kyc: 用户实名认证信息
-- ----------------------------
DROP TABLE IF EXISTS `user_kyc`;
CREATE TABLE `user_kyc`
(
  `id`               bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at`       datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`       datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`       datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status`           tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
  `desc`             varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  -- 唯一键
  `user_id`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',

  -- 实名信息
  `real_name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `first_name`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '姓',
  `middle_name`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '中间名',
  `last_name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名',

  -- 年龄, 性别, 身份证号 证件
  `gender`           tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',
  `birthday`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '生日',
  `language`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '语言',

  -- 证件信息: 证件类型, 证件编号
  `certificate_no`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证件编号',
  `certificate_meta` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证件信息',
  `certificate_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证件类型',
  `certificate_img1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证件图片类型',
  `certificate_img2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证件图片类型',
  `certificate_img3` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证件图片类型',

  -- 地址信息:
  `country`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '国家',
  `province`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '省份',
  `city`             varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '城市',
  `street`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '街道',
  `address`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地址',

  PRIMARY KEY (`id`),
  KEY                `idx_updated_at` (`updated_at`),
  KEY                `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户实名信息表';


SET
FOREIGN_KEY_CHECKS = 1;
