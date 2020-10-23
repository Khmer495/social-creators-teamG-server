# データベース構成（MySQL）

## Prefecture

| column     | type         | pk  | not null | uq  | unsigned | fk  | 他  |
| ---------- | ------------ | --- | -------- | --- | -------- | --- | --- |
| id         | int          | o   | o        |     | o        |
| created_at | datetime     |     | o        |
| updated_at | datetime     |     | o        |
| deleted_at | datetime     |     |
| name       | varchar(255) |     | o        | 1   |
| order      | int          |     |          | 1   |

## City

| column        | type         | pk  | not null | uq  | unsigned | fk                                                 | 他  |
| ------------- | ------------ | --- | -------- | --- | -------- | -------------------------------------------------- | --- |
| id            | int          | o   | o        |     | o        |
| created_at    | datetime     |     | o        |
| updated_at    | datetime     |     | o        |
| deleted_at    | datetime     |     |
| prefecture_id | int          |     | o        |     |          | Prefecture.id, on update=cascade, ondelete=cascade |
| name          | varchar(255) |     | o        | 1   |
| order         | int          |     | o        | 2   |

## User

| column      | type         | pk  | not null | uq  | unsigned | fk  | 他  |
| ----------- | ------------ | --- | -------- | --- | -------- | --- | --- |
| id          | int          | o   | o        |     | o        |
| created_at  | datetime     |     | o        |
| updated_at  | datetime     |     | o        |
| deleted_at  | datetime     |     |
| family_name | varchar(255) |     | o        |
| first_name  | varchar(255) |     | o        |
| email       | varchar(255) |     |          |     |

## Restaurant

| column                 | type         | pk  | not null | uq  | unsigned | fk                                                  | 他                         |
| ---------------------- | ------------ | --- | -------- | --- | -------- | --------------------------------------------------- | -------------------------- |
| id                     | int          | o   | o        |     | o        |
| created_at             | datetime     |     | o        |
| updated_at             | datetime     |     | o        |
| deleted_at             | datetime     |     |          |
| name                   | varchar(255) |     | o        |
| address                | varchar(255) |     | o        |
| lat                    | float        |     | o        |     |          |                                                     | 緯度                       |
| lng                    | float        |     | o        |     |          |                                                     | 経度                       |
| prefecture_id          | int          |     | o        |     | o        | Restaurant.id, onupdate=cascade, ondeleete=restrict |
| city_id                | int          |     |          |     | o        | Prefecture.id, onupdate=cascade, ondeleete=set null |
| open_time              | time         |     |
| close_time             | time         |     |
| from_kids_welcome_time | time         |     |
| to_kids_welcome_time   | time         |     |
| menu                   | text         |     |
| is_ok_baby             | bool         |     |          |     |          |                                                     | 乳児 OK                    |
| is_ok_infant           | tinyint(1)   |     |          |     |          |                                                     | 幼児 OK                    |
| is_no_smoking          | tinyint(1)   |     |          |     |          |                                                     | 禁煙・分煙                 |
| is_ok_baby_car         | bool         |     |          |     |          |                                                     | ベビーカー入店 OK          |
| is_zashiki             | tinyint(1)   |     |          |     |          |                                                     | 座敷                       |
| is_private_room        | tinyint(1)   |     |          |     |          |                                                     | 個室                       |
| is_parking_area        | tinyint(1)   |     |          |     |          |                                                     | 駐車場                     |
| is_luggage_storage     | tinyint(1)   |     |          |     |          |                                                     | 荷物置き(席の近くに)       |
| is_nap_space           | tinyint(1)   |     |          |     |          |                                                     | お昼寝スペース             |
| is_baby_bed            | tinyint(1)   |     |          |     |          |                                                     | ベビーベッド・おむつ交換台 |
| is_kids_menu           | tinyint(1)   |     |          |     |          |                                                     | お子様メニュー             |
| is_kids_chair          | tinyint(1)   |     |          |     |          |                                                     | お子様用椅子               |
| is_kids_dish           | tinyint(1)   |     |          |     |          |                                                     | お子様用食器               |
| is_kids_toilet         | tinyint(1)   |     |          |     |          |                                                     | お子様用トイレ             |
| is_kids_toy            | tinyint(1)   |     |          |     |          |                                                     | お子様用の本・おもちゃ     |
| is_ok_baby_food        | tinyint(1)   |     |          |     |          |                                                     | 離乳食の持ち込み OK        |
| is_nursing_room        | tinyint(1)   |     |          |     |          |                                                     | 授乳室・スペース           |
| is_kids_room           | tinyint(1)   |     |          |     |          |                                                     | キッズルーム・託児施設     |
| is_allergen_label      | tinyint(1)   |     |          |     |          |                                                     | アレルギー表示・対応食     |

## Comment

| column        | type         | pk  | not null | uq  | unsigned | fk                                                 | 他  |
| ------------- | ------------ | --- | -------- | --- | -------- | -------------------------------------------------- | --- |
| id            | int          | o   |
| user_id       | int          |     | o        |     | o        | User.id, onupdate=cascade, ondeleete=cascade       |
| restaurant_id | int          |     | o        |     | o        | Restaurant.id, onupdate=cascade, ondeleete=cascade |
| created_at    | datetime     |     | o        |
| updated_at    | datetime     |     | o        |
| text          | varchar(255) |     | o        |

## Good

| column        | type     | pk  | not null | uq  | unsigned | fk                                                 | 他  |     |
| ------------- | -------- | --- | -------- | --- | -------- | -------------------------------------------------- | --- | --- |
| id            | int      | o   |
| user_id       | int      |     | o        |     | o        | User.id, onupdate=cascade, ondeleete=cascade       |
| restaurant_id | int      |     | o        |     | o        | Restaurant.id, onupdate=cascade, ondeleete=cascade |
| created_at    | datetime |     | o        |
| updated_at    | datetime |     | o        |

## Favorite

| column        | type     | pk  | not null | uq  | unsigned | fk                                                 | 他  |     |
| ------------- | -------- | --- | -------- | --- | -------- | -------------------------------------------------- | --- | --- |
| id            | int      | o   |
| user_id       | int      |     | o        |     | o        | User.id, onupdate=cascade, ondeleete=cascade       |
| restaurant_id | int      |     | o        |     | o        | Restaurant.id, onupdate=cascade, ondeleete=cascade |
| created_at    | datetime |     | o        |
| updated_at    | datetime |     | o        |
