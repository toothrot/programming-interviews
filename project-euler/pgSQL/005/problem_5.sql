\timing
WITH RECURSIVE ints(n) AS (
  values (20)
union all
  select n + 20 from ints
)


select *
from ints
where ints.n NOT IN (
  SELECT ints.n
  FROM generate_series(20, 1, -1) gs(i)
  where n % i != 0
)
limit 1

