-- SQLite
SELECT u.id,
  u.id_number,
  u.gender_identity,
  u.birth_date,
  u.nationality,
  u.birth_sex
FROM `users` AS u
WHERE `u`.`id` = 1