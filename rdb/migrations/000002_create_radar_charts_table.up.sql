BEGIN;
CREATE TABLE IF NOT EXISTS `radar_charts`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'レーダーチャートの識別子',
    `insect_id`    BIGINT UNSIGNED NOT NULL COMMENT '昆虫の識別子（insectsテーブルのFK）',
    `umami_score`  TINYINT         NOT NULL COMMENT '旨味スコア（1〜5）',
    `bitter_score` TINYINT         NOT NULL COMMENT '苦味スコア（1〜5）',
    `egu_score`    TINYINT         NOT NULL COMMENT 'エグ味スコア（1〜5）',
    `flavor_score` TINYINT         NOT NULL COMMENT '風味スコア（1〜5）',
    `kimo_score`   TINYINT         NOT NULL COMMENT 'キモみスコア（1〜5）',
    `created_at`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'レコードの作成日時',
    `updated_at`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'レコードの更新日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_radar_charts_insect_id` FOREIGN KEY (`insect_id`) REFERENCES `insects` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='昆虫のレーダーチャート用スコアを管理するテーブル';
COMMIT;