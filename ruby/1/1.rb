def problem_1
  (0...1000).inject do |sum,i|
    if i % 3 == 0 || i % 5 == 0
      sum + i
    else
      sum
    end
  end
end

require 'minitest/autorun'
class TestProblem1 < Minitest::Test
  def test_that_answer_is_correct
    assert_equal(233168, problem_1)
  end
end
