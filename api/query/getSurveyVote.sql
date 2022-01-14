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
  c.value AS choice_value
FROM (
    SELECT *
    FROM `surveys`
    WHERE `surveys`.`id` = ?
  ) AS s
  JOIN questions AS q ON q.survey_refer = s.id
  JOIN choices AS c ON c.question_refer = q.id