-- SQLite
SELECT c.id AS choice_id,
  v.id AS vote_id,
  u.id AS user_id,
  u.birth_sex,
  u.gender_identity,
  u.birth_date,
  u.nationality
FROM (
    SELECT *
    FROM `choices`
    WHERE `choices`.`id` = 5
  ) AS c
  JOIN votes AS v ON v.choice_refer = c.id
  JOIN users AS u ON u.id = v.user_refer
ORDER BY c.id