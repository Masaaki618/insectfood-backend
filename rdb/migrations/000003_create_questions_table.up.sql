BEGIN;
CREATE TABLE IF NOT EXISTS `questions`
(
    `id`         BIGINT UNSIGNED                     NOT NULL AUTO_INCREMENT COMMENT '質問の識別子',
    `body`       VARCHAR(255)                        NOT NULL COMMENT '質問文',
    `category`   ENUM ('visual','physical','mental') NOT NULL COMMENT 'カテゴリ（visual: 視覚耐性 / physical: 身体耐性 / mental: 精神耐性）',
    `created_at` DATETIME                            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'レコードの作成日時',
    `updated_at` DATETIME                            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'レコードの更新日時',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='診断質問を管理するテーブル';
COMMIT;