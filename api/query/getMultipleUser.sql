-- SQLite
SELECT u.id,
  u.id_number,
  u.gender_identity,
  u.birth_date,
  u.is_resident,
  u.birth_sex
FROM `users` AS u
WHERE `u`.`id` = 1