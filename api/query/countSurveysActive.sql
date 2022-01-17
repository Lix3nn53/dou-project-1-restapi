-- SQLite
SELECT count(1)
FROM `surveys` AS s
WHERE `s`.`deleted_at` IS NULL
  AND date('now') BETWEEN `s`.`date_start` AND `s`.`date_end`