class List
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
  def each_element
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
      self.each_element do |element|
        count += 1
      end
      count
    else
      0
    end
  end

  # Deletes the first instance of value from list
  def delete(value)
    self.each_element do |element|
      if element.next_element && element.next_element.value == value
        element.next_element = element.next_element.next_element
      end
    end
  end

  def pop
    if self.head
      value = self.head.value
      self.head = self.head.next_element
      value
    end
  end
end

require 'minitest/autorun'
class TestList < Minitest::Test
  def test_pushing_an_item
    stack = List.new
    stack.push 1

    assert_equal(1, stack.length)
  end

  def test_pushing_and_popping_two_items
    stack = List.new
    stack.push 1
    stack.push 456

    assert_equal(2, stack.length)
    assert_equal(456, stack.pop)
    assert_equal(1, stack.pop)
  end

  def test_popping_an_item
    stack = List.new
    stack.push 123

    assert_equal(123, stack.pop)
    assert_equal(0, stack.length)
  end

  def test_popping_last_item
    stack = List.new
    stack.push 123

    assert_equal(123, stack.pop)
    assert_equal(0, stack.length)
    assert_equal(nil, stack.pop)
    assert_equal(nil, stack.head)
  end

  def test_empty_length
    stack = List.new
    assert_equal(0, stack.length)
  end

  def test_each_element
    stack = List.new
    stack.push 1
    stack.push 456

    values = []
    stack.each_element do |element|
      values << element.value
    end
    assert_equal([456, 1], values)
  end

  def test_delete_with_a_simple_case
    stack = List.new
    stack.push 1
    stack.push 456
    stack.push 1024

    stack.delete(456)
    assert_equal(1024, stack.pop)
    assert_equal(1, stack.pop)
  end
end

