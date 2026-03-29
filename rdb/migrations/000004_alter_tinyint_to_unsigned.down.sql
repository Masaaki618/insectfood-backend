BEGIN;

ALTER TABLE `insects`
    MODIFY COLUMN `difficulty` TINYINT NOT NULL COMMENT '難易度（★1〜★5）';

ALTER TABLE `radar_charts`
    MODIFY COLUMN `umami_score`  TINYINT NOT NULL COMMENT '旨味スコア（1〜5）',
    MODIFY COLUMN `bitter_score` TINYINT NOT NULL COMMENT '苦味スコア（1〜5）',
    MODIFY COLUMN `egu_score`    TINYINT NOT NULL COMMENT 'エグ味スコア（1〜5）',
    MODIFY COLUMN `flavor_score` TINYINT NOT NULL COMMENT '風味スコア（1〜5）',
    MODIFY COLUMN `kimo_score`   TINYINT NOT NULL COMMENT 'キモみスコア（1〜5）';

COMMIT;
