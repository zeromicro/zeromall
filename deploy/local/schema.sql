-- ---------------------------------------------------------------------------
-- ref: https://mariadb.com/kb/en/setting-character-sets-and-collations/
-- ---------------------------------------------------------------------------


-- ---------------------------------------------------------------------------
-- create a database;
-- ---------------------------------------------------------------------------

-- create database if not exists `test`;

-- 数据库创建: 用户数据库
DROP DATABASE IF EXISTS `user`;
CREATE DATABASE `user` DEFAULT CHARACTER SET = `utf8mb4` DEFAULT COLLATE `utf8mb4_unicode_ci`;


-- fix for deleted_at = 0
select @@sql_mode;
set sql_mode = '';


-- ---------------------------------------------------------------------------
-- create tables;
-- ---------------------------------------------------------------------------


-- 用户元信息表:
CREATE TABLE IF NOT EXISTS `user_meta`
(
    `id`            int(11) unsigned                   NOT NULL AUTO_INCREMENT COMMENT '自增主键(pk)',
    `user_id`       int(11) unsigned                   NOT NULL DEFAULT '0' COMMENT '用户 ID',
    `password`      varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '密码',
    `status`        int(11)                            NOT NULL DEFAULT '1' COMMENT '账号状态: -1, 主动停用, 1:正常, 2, 未验证, -2, 账号封禁.',
    -- 昵称:
    `unique_name`   varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '方案1: 用户唯一昵称',
    `nick_name`     varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '方案2: 用户昵称+编号',
    `nick_no`       int(11) unsigned                   NOT NULL DEFAULT '0' COMMENT '方案2: 昵称编号',
    -- 注册信息:
    `register`      varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '注册信息',
    `register_type` int(11)                            NOT NULL DEFAULT '1' COMMENT '注册方式: mobile/email',
    --
    `mobile_no`     varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '用户唯一昵称',
    `mobile_code`   varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '用户昵称',
    `email`         varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '邮箱',
    -- 默认:
    `ver`           int(11)                            NOT NULL DEFAULT '1' COMMENT '版本号，用于乐观锁',
    `created_at`    timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at`    timestamp                          NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '删除时间',
    --
    --
    --
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`),
    UNIQUE KEY `uk_register` (`register`),
    --
    KEY `idx_register_type` (`register_type`),
    --
    KEY `idx_created_at` (`created_at`),
    KEY `idx_updated_at` (`updated_at`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户元信息表';


-- 用户鉴权方式表:
CREATE TABLE IF NOT EXISTS `user_auth`
(
    `id`         int(11) unsigned                   NOT NULL AUTO_INCREMENT COMMENT '自增主键(pk)',
    `user_id`    int(11) unsigned                   NOT NULL DEFAULT '0' COMMENT '用户 ID',
    `status`     int(11)                            NOT NULL DEFAULT '-1' COMMENT '登录认证方式可用状态: -1, 无效, 1:有效',
    `label`      varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '备注信息',
    -- 认证信息:
    `auth_type`  int(11) unsigned                   NOT NULL DEFAULT '0' COMMENT '注册方式: mobile/email/QQ',
    `auth_meta1` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '认证信息',
    `auth_meta2` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '认证信息',
    `auth_meta3` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '认证信息',
    `auth_meta4` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '认证信息',
    -- 默认:
    `ver`        int(11)                            NOT NULL DEFAULT '1' COMMENT '版本号，用于乐观锁',
    `created_at` timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at` timestamp                          NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '删除时间',
    --
    --
    --
    PRIMARY KEY (`id`),
    --
    KEY `idx_user_id` (`user_id`),
    KEY `idx_auth_type` (`auth_type`),
    KEY `idx_auth_meta1` (`auth_meta1`),
    KEY `idx_auth_meta2` (`auth_meta2`),
    KEY `idx_auth_meta3` (`auth_meta3`),
    KEY `idx_auth_meta4` (`auth_meta4`),
    --
    KEY `idx_created_at` (`created_at`),
    KEY `idx_updated_at` (`updated_at`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户鉴权方式表';


-- 用户详情信息:
CREATE TABLE IF NOT EXISTS `user_profile`
(
    `id`          int(11) unsigned                   NOT NULL AUTO_INCREMENT COMMENT '自增主键(pk)',
    `user_id`     int(11) unsigned                   NOT NULL DEFAULT '0' COMMENT '用户 ID',
    -- 身份信息:
    `real_name`   varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '真实姓名',
    `first_name`  varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '姓名',
    `middle_name` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '姓名',
    `last_name`   varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '姓名',
    --
    `age`         int(11) unsigned                   NOT NULL DEFAULT '0' COMMENT '年龄',
    `sex`         varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '性别',
    `avatar`      varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '头像',
    -- 默认:
    `ver`         int(11)                            NOT NULL DEFAULT '1' COMMENT '版本号，用于乐观锁',
    `created_at`  timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at`  timestamp                          NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '删除时间',
    --
    --
    --
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`),
    --
    KEY `idx_created_at` (`created_at`),
    KEY `idx_updated_at` (`updated_at`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户鉴权方式表';




