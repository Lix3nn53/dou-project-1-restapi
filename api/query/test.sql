-- SQLite
-- delete
DELETE FROM votes;
-- update id
UPDATE votes
SET id = 1
WHERE id ISNULL;
-- select choices
SELECT q.id as question_id,
  c.id AS choice_id
FROM (
    SELECT *
    FROM `surveys`
    WHERE `surveys`.`id` = 3
  ) AS s
  JOIN questions AS q ON q.survey_refer = s.id
  JOIN choices AS c ON c.question_refer = q.id
ORDER BY c.id;
-- time
SELECT DATE("now");
-- expire
UPDATE surveys
SET date_end = "2022-01-01 00:00:00+00:00"
WHERE id = 2;