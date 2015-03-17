def fib(input = [1,2])
  if input[-1] <= 4_000_000
    fib(input << input[-1]+input[-2])
  else
    input
  end
end

def problem_2(input = [1,2],sum=2)
  if input[-1] <= 4_000_000
    new = input[-1] + input[-2]
    sum = (new % 2 == 0) ? sum+new : sum
    problem_2(input << input[-1]+input[-2],sum)
  else
    sum
  end
end

require 'minitest/autorun'
class TestProblem2 < Minitest::Test
  def test_that_answer_is_correct
    assert_equal(4613732, problem_2)
  end
end
