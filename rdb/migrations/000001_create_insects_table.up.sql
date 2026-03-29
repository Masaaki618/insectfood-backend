BEGIN;
CREATE TABLE IF NOT EXISTS `insects`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '昆虫の識別子',
    `name`         VARCHAR(100)    NOT NULL COMMENT '昆虫名',
    `difficulty`   TINYINT         NOT NULL COMMENT '難易度（★1〜★5）',
    `introduction` TEXT            NOT NULL COMMENT '昆虫の説明',
    `taste`        VARCHAR(100)    NOT NULL COMMENT '味の説明',
    `texture`      VARCHAR(100)    NOT NULL COMMENT '食感の説明',
    `insect_img`   VARCHAR(255) COMMENT '画像URL',
    `created_at`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'レコードの作成日時',
    `updated_at`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'レコードの更新日時',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='昆虫情報を管理するテーブル';
COMMIT;