WITH RECURSIVE nums(a) AS (
  VALUES (0)
UNION ALL
  SELECT a+1 AS a FROM nums
  WHERE a < 1000
)
SELECT sum(a) FROM nums WHERE a % 3 = 0 OR a % 5 = 0;
