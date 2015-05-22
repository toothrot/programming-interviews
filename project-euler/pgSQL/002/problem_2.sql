WITH recursive fib(a, b) AS (
  VALUES (0, 1)
UNION ALL
  SELECT greatest(a, b), a+b AS a FROM fib
  WHERE b < 4000000
)
SELECT SUM(a) FROM fib WHERE a % 2 = 0;
