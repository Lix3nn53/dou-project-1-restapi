-- SQLite
SELECT s.id,
  s.user_refer,
  s.subject,
  s.description,
  s.date_start,
  s.date_end
FROM `surveys` AS s
WHERE `s`.`id` = ?