class SimpleStack
  def initialize
    @data = []
  end

  def push(item)
    @data << item
  end
  
  def length
    @data.length
  end

  def pop
    @data.pop
  end
end

require 'minitest/autorun'
class TestSimpleStack < Minitest::Test
  def stack_class
    SimpleStack
  end

  def test_pushing_an_item
    stack = stack_class.new
    stack.push 1

    assert_equal(1, stack.length)
  end

  def test_pushing_and_popping_two_items
    stack = stack_class.new
    stack.push 1
    stack.push 456

    assert_equal(2, stack.length)
    assert_equal(456, stack.pop)
    assert_equal(1, stack.pop)
  end

  def test_popping_an_item
    stack = stack_class.new
    stack.push 123

    assert_equal(123, stack.pop)
    assert_equal(0, stack.length)
  end

  def test_empty_length
    stack = stack_class.new
    assert_equal(0, stack.length)
  end
end

