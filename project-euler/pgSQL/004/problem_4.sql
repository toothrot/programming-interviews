select max(a * b)
from generate_series(100,999) gs1(a), generate_series(100,999) gs2(b)
where (a * b)::varchar(255) = reverse((a * b)::varchar(255))
