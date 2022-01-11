-- SQLite
SELECT s.id,
  s.user_refer,
  s.subject,
  s.description,
  s.date_start,
  s.date_end,
  q.id AS question_id,
  q.value AS question_value,
  c.id AS choice_id,
  c.value AS choice_value,
  v.id AS vote_id
FROM (
    SELECT *
    FROM `surveys`
    ORDER BY `surveys`.`id`
    LIMIT 5 OFFSET 0
  ) AS s
  JOIN questions AS q ON q.survey_refer = s.id
  JOIN choices AS c ON c.question_refer = q.id
  LEFT JOIN votes AS v ON v.choice_refer = c.id
WHERE `s`.`deleted_at` IS NULL