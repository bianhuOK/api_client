-- 创建表
CREATE TABLE `sequence_table` (
  `seq_name` varchar(64) NOT NULL COMMENT '序列名称',
  `current_value` bigint NOT NULL COMMENT '当前值',
  `step` int NOT NULL COMMENT '步长',
  `max_value` bigint NOT NULL COMMENT '最大值',
  `preload_threshold` float NOT NULL COMMENT '预加载阈值',
  `version` int NOT NULL DEFAULT 1 COMMENT '版本号，用于乐观锁',
  `create_at` int NOT NULL COMMENT '创建时间(秒级时间戳)',
  `modified_at` int NOT NULL COMMENT '修改时间(秒级时间戳)',
  `creator` varchar(64) NOT NULL DEFAULT 'system' COMMENT '创建者',
  `modifier` varchar(64) NOT NULL DEFAULT 'system' COMMENT '修改者',
  PRIMARY KEY (`seq_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入测试数据
INSERT INTO `sequence_table` (
    `seq_name`, 
    `current_value`, 
    `step`, 
    `max_value`, 
    `preload_threshold`,
    `version`, 
    `create_at`, 
    `modified_at`, 
    `creator`, 
    `modifier`
) VALUES (
    'test_seq',
    1000,
    1000,
    9999999999,
    0.8,
    1,
    1737283732,
    1737283732,
    'bianhuOK',
    'bianhuOK'
);