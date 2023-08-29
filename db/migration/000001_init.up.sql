-- 创建 user 表
CREATE TABLE
	`user` (
					 `id` bigint unsigned NOT NULL COMMENT '主键',
					 `nickname` varchar(255) NOT NULL COMMENT '用户昵称',
					 `email` varchar(255) NOT NULL COMMENT '用户邮箱',
					 `avatar` varchar(255) COMMENT '用户头像',
					 `password` varchar(255) NOT NULL COMMENT '加密后的用户密码',
					 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
					 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					 `deleted_at` datetime,
					 PRIMARY KEY (`id`),
					 UNIQUE KEY `idx_email` (`email`),
					 INDEX `idx_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';


-- 创建 tag 表
CREATE TABLE
	`tag` (
					`id` bigint unsigned NOT NULL COMMENT '主键',
					`name` varchar(255) NOT NULL COMMENT '标签名称',
					`friendly_url` varchar(255) NOT NULL COMMENT '标签的friendly_url，SEO用',
					`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
					`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					`deleted_at` datetime,
					PRIMARY KEY (`id`),
					UNIQUE KEY `idx_name` (`name`),
					UNIQUE KEY `idx_friendly_url` (`friendly_url`),
					INDEX `idx_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '标签表';


-- 创建 article 表
CREATE TABLE
	`article` (
							`id` bigint unsigned NOT NULL COMMENT '主键',
							`title` varchar(255) NOT NULL COMMENT '文章标题',
							`description` text NOT NULL COMMENT '文章描述',
							`cover` varchar(255) COMMENT '文章封面',
							`content` text NOT NULL COMMENT '文章内容',
							`friendly_url` varchar(255) NOT NULL COMMENT '文章的friendly_url，SEO用',
							`is_top` boolean NOT NULL DEFAULT FALSE COMMENT '文章是否置顶',
							`top_priority` tinyint(10) NOT NULL DEFAULT 0 COMMENT '文章置顶优先级',
							`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
							`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
							`deleted_at` datetime,
							PRIMARY KEY (`id`),
							UNIQUE KEY `idx_title` (`title`),
							UNIQUE KEY `idx_friendly_url` (`friendly_url`),
							INDEX `idx_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '文章表';


-- 创建 article 和 tag 关联表
CREATE TABLE
	`article_tag` (
									`article_id` bigint unsigned NOT NULL COMMENT '文章id',
									`tag_id` bigint unsigned NOT NULL COMMENT '标签id',
									PRIMARY KEY (`article_id`, `tag_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '文章和标签关联表';