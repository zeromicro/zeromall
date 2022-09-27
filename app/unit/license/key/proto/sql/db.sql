-- 数据库创建: 用户数据库

-- del:
DROP
  DATABASE IF EXISTS `mall_license`;

-- new:
CREATE
  DATABASE `mall_license` DEFAULT CHARACTER SET = `utf8mb4` DEFAULT COLLATE `utf8mb4_unicode_ci`;

-- fix for deleted_at = 0
select @@sql_mode;
set
  sql_mode = '';
