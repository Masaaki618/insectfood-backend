ALTER TABLE `insects` ADD COLUMN `deleted_at` DATETIME(3) NULL COMMENT '論理削除日時';
ALTER TABLE `radar_charts` ADD COLUMN `deleted_at` DATETIME(3) NULL COMMENT '論理削除日時';
ALTER TABLE `questions` ADD COLUMN `deleted_at` DATETIME(3) NULL COMMENT '論理削除日時';

CREATE INDEX `idx_insects_deleted_at` ON `insects` (`deleted_at`);
CREATE INDEX `idx_radar_charts_deleted_at` ON `radar_charts` (`deleted_at`);
CREATE INDEX `idx_questions_deleted_at` ON `questions` (`deleted_at`);
