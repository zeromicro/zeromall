SET NAMES utf8mb4;
SET
  FOREIGN_KEY_CHECKS = 0;

-- 选择 db:
USE
  app_sms;

-- ----------------------------
-- Table structure for sms_provider 短信服务提供商
-- ----------------------------
DROP TABLE IF EXISTS `sms_provider`;
CREATE TABLE `sms_provider`
(
  `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `desc`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',

  `provider`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信服务提供商',
  `app_id`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信服务提供商',
  `app_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信服务提供商',
  `sign`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信服务提供商',

  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='短信服务提供商';

-- ----------------------------
-- Table structure for sms_template 短信模板
-- ----------------------------
DROP TABLE IF EXISTS `sms_template`;
CREATE TABLE `sms_template`
(
  `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,

  `sn`         int(11)                                                       NOT NULL DEFAULT '0' COMMENT '短信模板编号',
  `template`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信模板',
  `desc`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',
  `status`     int(11)                                                       NOT NULL DEFAULT '0' COMMENT '状态 0:禁用 1:启用',

  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='短信模板';


-- ----------------------------
-- Table structure for sms_task 短信定时任务
-- ----------------------------
DROP TABLE IF EXISTS `sms_task`;
CREATE TABLE `sms_task`
(
  `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
  `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,

  `task`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '任务',
  `params`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '任务参数',
  `status`     int(11)                                                       NOT NULL DEFAULT '0' COMMENT '任务状态 0:禁用 1:启用',
  `cron`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '执行计划表',

  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='短信定时任务';

