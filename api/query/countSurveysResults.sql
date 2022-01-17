-- SQLite
SELECT count(1)
FROM `surveys`
WHERE `surveys`.`deleted_at` IS NULL
  AND NOT date('now') BETWEEN `surveys`.`date_start` AND `surveys`.`date_end`