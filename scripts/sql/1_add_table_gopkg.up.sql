CREATE TABLE gopkg_package (
  id BIGINT(20) NOT NULL AUTO_INCREMENT,
  `name` CHAR(50) NOT NULL COMMENT '包名称',
  source VARCHAR(100) NOT NULL COMMENT '包源码地址',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '描述',
  updated_at DATETIME NOT NULL COMMENT '更新时间',
  deleted_at datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  created_at DATETIME NOT NULL COMMENT '创建时间',
  `version` int(20) NOT NULL DEFAULT 0 COMMENT '版本 乐观锁',
  PRIMARY KEY (id),
  INDEX `idx_deletedAt`(`deleted_at`) USING BTREE COMMENT '删除索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;