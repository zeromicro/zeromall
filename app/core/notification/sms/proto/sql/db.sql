-- 数据库创建: 短信服务

-- del:
DROP
  DATABASE IF EXISTS `app_sms`;

-- new:
CREATE
  DATABASE `app_sms` DEFAULT CHARACTER SET = `utf8mb4` DEFAULT COLLATE `utf8mb4_unicode_ci`;

-- fix for deleted_at = 0
select @@sql_mode;
set
  sql_mode = '';
