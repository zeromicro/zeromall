-- import this sql:
--     方法1: mysql -u dev -p dev dev < ./infra-db-mysql-init.sql;
--     方法2: mysql> source /home/to/path/infra-db-mysql-init.sql;
--     方法3: IDE 插件连接 db, 导入此 sql

-- create the databases:
-- CREATE DATABASE IF NOT EXISTS `dev`;
-- CREATE DATABASE IF NOT EXISTS `test`;
-- USE `dev`;


-- create the users for each database:
-- CREATE USER 'test'@'%' IDENTIFIED BY 'test';
-- GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON `test`.* TO 'test'@'%';

-- FLUSH PRIVILEGES;


-- crate table:
-- 电池兑换商品订单表:
CREATE TABLE IF NOT EXISTS `order_exchange` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键(pk)',
	`mid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'up主的mid',
	`order_no` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'UP主兑换订单号(uk)',
	`merchant_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
	`product_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
	`product_type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '商品类型: 0=未定义, 1=虚拟商品, 2=实物商品',
	`product_sku` int(11) unsigned  NOT NULL DEFAULT '0' COMMENT '商品SKU',
	`product_title` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '商品标题',
	`product_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
	`product_price` decimal(10,2) NOT NULL DEFAULT '0.0' COMMENT '商品单价',
	`product_discount` decimal(10,2) NOT NULL DEFAULT '1.0' COMMENT '商品折扣',
	`cost` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '消耗电池总数',
	`status` int(11) NOT NULL DEFAULT '0' COMMENT '订单状态: -2=退款, -1=失败(结束), 0=未定义, 1=成功(结束), 2=pending',
	`ver` int(11) NOT NULL DEFAULT '1' COMMENT '版本号，用于乐观锁',
	`ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	PRIMARY KEY (`id`),
	UNIQUE KEY `uk_order_no` (`order_no`),
	KEY `ix_mid` (`mid`),
	KEY `ix_ctime` (`ctime`),
	KEY `ix_mtime` (`mtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='电池兑换商品订单表';


-- drop table:
-- DROP TABLE IF EXISTS `order_exchange`;

-- show table:
SHOW CREATE TABLE `order_exchange`;

