def fib(input = [1,2])
  if input[-1] <= 4_000_000
    fib(input << input[-1]+input[-2])
  else
    input
  end
end

def sum_of_even_fibbs(input = [1,2],sum=2)
  if input[-1] <= 4_000_000
    new = input[-1]+input[-2]
    sum = (new%2==0)?sum+new:sum
    sum_of_even_fibbs(input << input[-1]+input[-2],sum)
  else
    sum
  end
end

puts sum_of_even_fibbs
