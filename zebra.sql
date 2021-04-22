drop TABLE if exists `zebra`.`accounts`;
CREATE TABLE `zebra`.`accounts` (
  `id` bigint(32) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(512) not NULL COMMENT '用户名',
  `password` binary(64) NOT NULL COMMENT "密码",
  `salt` binary(64) NOT NULL COMMENT "盐",
  `age` int(32) NOT NULL COMMENT "年龄",
  `age_id` bigint(32) NOT NULL COMMENT "年龄 id",
  `email` varchar(32) NOT NULL COMMENT "邮箱",
  `email_id` bigint(32) NOT NULL COMMENT "邮箱 id",
  `bank` varchar(32) NOT NULL COMMENT "银行账户余额,$",
  `bank_id` bigint(32) NOT NULL COMMENT "银行账户余额 id",
  `cellphone` varchar(32) DEFAULT NULL COMMENT '手机号',
  `cellphone_id` varchar(32) DEFAULT NULL COMMENT '手机号 id',
  `is_deleted` tinyint(4) DEFAULT 0 COMMENT '是否删除,0:未删除,1:删除',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`),

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号表';

DROP TABLE IF EXISTS `zebra`.`ages`;
CREATE TABLE `zebra`.`ages` (
  `id` bigint(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `age` int(32) NOT NULL COMMENT '类别',
  `is_deleted` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '删除标志,0:正常,1:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_idx_age` (`age`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='年龄表';

DROP TABLE IF EXISTS `zebra`.`emails`;
CREATE TABLE `zebra`.`emails` (
  `id` bigint(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `email` varchar(128) NOT NULL COMMENT '邮箱',
  `is_deleted` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '删除标志,0:正常,1:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_idx_email` (`email`(32))
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邮箱表';

DROP TABLE IF EXISTS `zebra`.`banks`;
CREATE TABLE `zebra`.`banks` (
  `id` bigint(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `bank` bigint(32) NOT NULL COMMENT 'money count,$',
  `is_deleted` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '删除标志,0:正常,1:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_idx_bank` (`bank`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='银行账户表';

DROP TABLE IF EXISTS `zebra`.`cellphones`;
CREATE TABLE `zebra`.`cellphones` (
  `id` bigint(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `cellphone` bigint(32) NOT NULL COMMENT '手机号',
  `is_deleted` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '删除标志,0:正常,1:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_idx_cellphone` (`cellphone`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='手机号表';