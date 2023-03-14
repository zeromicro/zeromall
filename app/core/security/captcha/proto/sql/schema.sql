SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- 选择 db:
USE
app_captcha;

-- ----------------------------
-- Table structure for captcha_provider: 验证码提供商
-- ----------------------------
DROP TABLE IF EXISTS `captcha_provider`;
CREATE TABLE `user`
(
  `id`          bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,

  -- 验证码服务提供商
  `provider_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '验证码提供商',
  `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '验证码提供商名称',
  `desc`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '验证码提供商描述',
  `status`      tinyint(8) NOT NULL DEFAULT '0' COMMENT '验证码提供商状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',

  PRIMARY KEY (`id`),
  KEY           `idx_updated_at` (`updated_at`),
  KEY           `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_provider_id` (`provider_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户信息表';


-- ----------------------------
-- Table structure for captcha_biz: 使用验证码服务的业务
-- ----------------------------
DROP TABLE IF EXISTS `captcha_biz`;
CREATE TABLE `captcha_biz`
(
  `id`               bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at`       datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`       datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`       datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,

  -- 产品业务id
  `biz_id`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务ID',
  `name`             varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务名称',
  `desc`             varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务描述',
  `status`           tinyint(8) NOT NULL DEFAULT '0' COMMENT '业务状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',

  -- provider_id
  `provider_id`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '验证码提供商',
  -- 向服务商申请的密钥对
  `biz_apply_key`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '向服务商申请的业务id',
  `biz_apply_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '申请的密钥',

  PRIMARY KEY (`id`),
  KEY                `idx_updated_at` (`updated_at`),
  KEY                `idx_created_at` (`created_at`),

  UNIQUE KEY `uk_biz_id` (`biz_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='使用验证码的业务信息表';



SET
FOREIGN_KEY_CHECKS = 1;
