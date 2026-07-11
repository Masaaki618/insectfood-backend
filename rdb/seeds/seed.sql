-- 昆虫10種の初期データ
INSERT INTO `insects` (`name`, `difficulty`, `introduction`, `taste`, `texture`, `insect_img`)
VALUES ('コオロギパウダー', 1, '粉末状に加工されたコオロギ。料理に混ぜやすく昆虫食入門に最適。', 'ナッツ系', 'サラサラ',
        'https://storage.googleapis.com/insectfood-images/cricket_powder.png'),
       ('ミールワーム', 1, 'チャイロコメノゴミムシダマシの幼虫。たんぱく質が豊富でクセが少ない。', '淡白', 'サクサク',
        'https://storage.googleapis.com/insectfood-images/mealworm.png'),
       ('シルクワーム', 2, 'カイコの蛹。東アジアで古くから食べられてきた伝統食材。', 'クリーミー', 'もちもち',
        'https://storage.googleapis.com/insectfood-images/silkworm.png'),
       ('ハチの子', 2, 'ミツバチやスズメバチの幼虫・蛹。日本各地で昔から食べられている。', '甘い', 'プリプリ',
        'https://storage.googleapis.com/insectfood-images/hachi_no_ko.png'),
       ('イナゴの佃煮', 3, '稲を食べるバッタの一種。甘辛い佃煮は日本の伝統的な昆虫食。', '甘辛', 'カリカリ',
        'https://storage.googleapis.com/insectfood-images/inago.png'),
       ('カイコ', 3, 'シルクワームの成虫前の幼虫。繭を取った後に食用として利用される。', '淡白', 'ふわふわ',
        'https://storage.googleapis.com/insectfood-images/kaiko.png'),
       ('タガメ', 4, '水生昆虫の王様とも呼ばれる大型の水生カメムシ。東南アジアの珍味。', 'フルーティー', 'パリパリ',
        'https://storage.googleapis.com/insectfood-images/tagame.png'),
       ('ゲンゴロウ', 4, '水中に生息する甲虫。独特の苦味がクセになると言われる。', '苦味あり', 'カリカリ',
        'https://storage.googleapis.com/insectfood-images/gengoro.png'),
       ('サソリ', 5, '毒は加熱で無害化される。中国では串焼きが有名な観光グルメ。', 'エビ系', 'パリパリ',
        'https://storage.googleapis.com/insectfood-images/sasori.png'),
       ('タランチュラ', 5, '世界最大級のクモ。カンボジアでは揚げ物として一般的に食べられる。', '鶏肉系', 'もちもち',
        'https://storage.googleapis.com/insectfood-images/tarantula.png');

-- レーダーチャート10種の初期データ
INSERT INTO `radar_charts` (`insect_id`, `umami_score`, `bitter_score`, `egu_score`, `flavor_score`, `kimo_score`)
VALUES (1, 3, 1, 1, 3, 1),  -- コオロギパウダー
       (2, 3, 1, 1, 2, 2),  -- ミールワーム
       (3, 4, 1, 2, 3, 2),  -- シルクワーム
       (4, 4, 1, 1, 4, 2),  -- ハチの子
       (5, 3, 2, 2, 3, 2),  -- イナゴの佃煮
       (6, 3, 1, 2, 2, 3),  -- カイコ
       (7, 2, 2, 3, 4, 4),  -- タガメ
       (8, 2, 4, 3, 2, 4),  -- ゲンゴロウ
       (9, 3, 2, 2, 3, 5),  -- サソリ
       (10, 3, 2, 3, 3, 5); -- タランチュラ

-- 質問12問の初期データ
INSERT INTO `questions` (`body`, `category`)
VALUES ('虫が出てくる映像や画像を抵抗なく見られる', 'visual'),
       ('昆虫の形がそのまま残っている食べ物でも食べられそう', 'visual'),
       ('グロテスクな見た目のものでも食べ物なら挑戦できる', 'visual'),
       ('虫の標本や写真を見ても気分が悪くならない', 'visual'),
       ('虫を手で触ったことがある', 'physical'),
       ('生き物をそのまま食べることに抵抗がない', 'physical'),
       ('屋台などで正体不明の食べ物を食べたことがある', 'physical'),
       ('見た目が気になっても一口食べてみることができる', 'physical'),
       ('食べ物の見た目より味や栄養を優先できる', 'mental'),
       ('周りが驚くようなことに挑戦するのが好き', 'mental'),
       ('新しい食体験にワクワクする方だ', 'mental'),
       ('話のネタのためなら多少無理をして食べられる', 'mental');
