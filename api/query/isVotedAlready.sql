-- SQLite
SELECT v.id AS vote_id
FROM (
    SELECT *
    FROM `surveys`
    WHERE `surveys`.`id` = ?
  ) AS s
  JOIN questions AS q ON q.survey_refer = s.id
  JOIN choices AS c ON c.question_refer = q.id
  LEFT JOIN votes AS v ON v.choice_refer = c.id
WHERE v.user_refer = '1'
ORDER BY c.id