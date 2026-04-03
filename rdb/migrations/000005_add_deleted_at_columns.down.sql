DROP INDEX `idx_insects_deleted_at` ON `insects`;
DROP INDEX `idx_radar_charts_deleted_at` ON `radar_charts`;
DROP INDEX `idx_questions_deleted_at` ON `questions`;

ALTER TABLE `insects` DROP COLUMN `deleted_at`;
ALTER TABLE `radar_charts` DROP COLUMN `deleted_at`;
ALTER TABLE `questions` DROP COLUMN `deleted_at`;
