WITH primes AS (
  SELECT n 
  FROM generate_series(3,(|/600851475143)::integer, 2) i(n)
  WHERE NOT EXISTS(
    SELECT 1 
    FROM generate_series(2,(|/n)::integer) ints(i)
    WHERE n % i = 0
  )
)

SELECT max(n)
FROM primes
WHERE 600851475143 % n = 0
  AND n < (|/600851475143) 
;
