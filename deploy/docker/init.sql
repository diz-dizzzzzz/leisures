-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS acupofcoffee DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE acupofcoffee;

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `username` varchar(50) NOT NULL,
    `password` varchar(100) NOT NULL,
    `email` varchar(100) NOT NULL,
    `nickname` varchar(50) DEFAULT '',
    `avatar` varchar(255) DEFAULT '',
    `phone` varchar(20) DEFAULT '',
    `status` tinyint DEFAULT 1 COMMENT '状态 1:正常 0:禁用',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_users_username` (`username`),
    UNIQUE KEY `idx_users_email` (`email`),
    KEY `idx_users_phone` (`phone`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

