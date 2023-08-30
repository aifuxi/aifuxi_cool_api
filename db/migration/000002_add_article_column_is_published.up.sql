-- article 表，新增 is_published 字段
ALTER TABLE `article` 
  ADD COLUMN `is_published` boolean NOT NULL DEFAULT FALSE COMMENT '文章是否发布';