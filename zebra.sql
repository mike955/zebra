drop TABLE if exists `zebra`.`accounts`;
CREATE TABLE `zebra`.`accounts` (
  `id` bigint(32) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(512) not NULL COMMENT '用户名',
  `age` int(32) NOT NULL COMMENT "年龄",
  `age_id` bigint(32) NOT NULL COMMENT "age id",
  `password` binary(64) NOT NULL COMMENT "密码",
  `salt` binary(64) NOT NULL COMMENT "盐",
  `level` tinyint(64) not NULL COMMENT '用户等级,0:普通',
  `qq` varchar(32) DEFAULT 0 COMMENT 'qq 号',
  `wechat` varchar(32) DEFAULT NULL COMMENT '微信号',
  `cellphone` varchar(32) DEFAULT NULL COMMENT '手机号',
  `email` varchar(32) DEFAULT NULL COMMENT '邮箱号',
  `state` tinyint(32) not NULL default 0 COMMENT '状态,0:正常',
  `last_login_time` datetime not NULL COMMENT '上次登陆时间',
  `is_deleted` tinyint(4) DEFAULT 0 COMMENT '是否删除,0:未删除,1:删除',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号表';

DROP TABLE IF EXISTS `zebra`.`ages`;
CREATE TABLE `zebra`.`ages` (
    `id` bigint(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `age` int(32) NOT NULL COMMENT '类别',
    `is_deleted` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '删除标志,0:正常,1:删除',
    `created_at` DATETIME NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL COMMENT '最后更新时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='年龄表';

