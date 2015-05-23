\timing
with primes AS (
  select n
  from generate_series(3,(|/600851475143)::integer, 2) gs(n)
  where not exists (
    select 1
    from generate_series(2,(|/gs.n)::integer) ints(i)
    where n % i = 0
  )
  AND 600851475143 % gs.n = 0
)

select max(n) from primes;
