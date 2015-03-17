class ListStack
  attr_accessor :head

  class ListElement
    attr_accessor :value, :next_element
    def initialize(value: , next_element: )
      self.value = value
      self.next_element = next_element
    end
  end

  def push(value)
    old_head = self.head
    self.head = ListElement.new(value: value, next_element: old_head)
  end

  # Needs more tests
  def each
    current_head = head
    yield current_head
    
    while current_head.next_element != nil
      current_head = current_head.next_element
      yield current_head
    end

    current_head
  end

  def length
    if head
      count = 0
      self.each do |element|
        count += 1
      end
      count
    else
      0
    end
  end

  def pop
    former_head = self.head
    self.head = former_head.next_element
    former_head.value
  end
end

require 'minitest/autorun'
class TestListStack < Minitest::Test
  def stack_class
    ListStack
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

  def test_each
    stack = stack_class.new
    stack.push 1
    stack.push 456

    values = []
    stack.each do |element|
      values << element.value
    end
    assert_equal([456, 1], values)
  end
end
