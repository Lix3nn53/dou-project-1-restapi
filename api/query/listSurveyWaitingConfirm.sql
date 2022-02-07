-- SQLite
SELECT s.id, s.user_refer, s.subject, s.description, s.date_start, s.date_end 
FROM `surveys` AS s WHERE `s`.`deleted_at` IS NULL 
AND `s`.`confirm_status` = 'waiting' ORDER BY `s`.`id` LIMIT 10 OFFSET 0