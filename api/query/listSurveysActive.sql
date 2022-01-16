-- SQLite
SELECT s.id,
  s.user_refer,
  s.subject,
  s.description,
  s.date_start,
  s.date_end
FROM `surveys` AS s
WHERE `s`.`deleted_at` IS NULL
  AND date('now') BETWEEN `s`.`date_start` AND `s`.`date_end`
ORDER BY `s`.`id`
LIMIT 5 OFFSET 0