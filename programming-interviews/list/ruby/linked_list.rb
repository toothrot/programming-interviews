class List
  attr_accessor :head

  class Element
    attr_accessor :value, :next_element
    def initialize(value: , next_element: )
      self.value = value
      self.next_element = next_element
    end
  end

  def push(value)
    old_head = self.head
    self.head = Element.new(value: value, next_element: old_head)
  end

  def each_element(&block)
    current_head = head
    yield current_head if block
    
    while current_head.next_element != nil
      current_head = current_head.next_element
      yield current_head if block
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
    if self.head && self.head.value == value
      self.head = self.head.next_element
    else
      self.each_element do |element|
        if element.next_element && element.next_element.value == value
          element.next_element = element.next_element.next_element
        end
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

  def insert_after(target_value, value)
    self.each_element do |element|
      if element.value == target_value
        new_element = Element.new(value: value, next_element: element.next_element)
        element.next_element = new_element
      end
    end
  end

  def append(value)
    last_element = self.each_element
    last_element.next_element = Element.new(value: value, next_element: nil)
  end

  def nth_to_last(n)
    nth_to_last = nil
    index = 0

    self.each_element do |element|
      nth_to_last ||= element

      if index >= n + 1
        nth_to_last = nth_to_last.next_element
      end

      index += 1
    end

    if n > index
      nth_to_last = nil
    end

    if nth_to_last
      nth_to_last.value
    end
  end
end

require 'minitest/autorun'
class TestList < Minitest::Test
  def test_pushing_an_item
    list = List.new
    list.push 1

    assert_equal(1, list.length)
  end

  def test_pushing_and_popping_two_items
    list = List.new
    list.push 1
    list.push 456

    assert_equal(2, list.length)
    assert_equal(456, list.pop)
    assert_equal(1, list.pop)
  end

  def test_popping_an_item
    list = List.new
    list.push 123

    assert_equal(123, list.pop)
    assert_equal(0, list.length)
  end

  def test_popping_last_item
    list = List.new
    list.push 123

    assert_equal(123, list.pop)
    assert_equal(0, list.length)
    assert_equal(nil, list.pop)
    assert_equal(nil, list.head)
  end

  def test_empty_length
    list = List.new
    assert_equal(0, list.length)
  end

  def test_each_element
    list = List.new
    list.push 1
    list.push 456

    values = []
    list.each_element do |element|
      values << element.value
    end
    assert_equal([456, 1], values)
  end

  def test_delete_with_a_simple_case
    list = List.new
    list.push 1
    list.push 456
    list.push 1024

    list.delete(456)
    assert_equal(1024, list.pop)
    assert_equal(1, list.pop)
  end

  def test_deleting_head_of_list
    list = List.new
    list.push 1
    list.push 456
    list.push 1024

    list.delete(1024)
    assert_equal(456, list.pop)
    assert_equal(1, list.pop)
  end

  def test_deleting_only_element
    list = List.new
    list.push 1024

    list.delete(1024)
    assert_equal(0, list.length)
  end

  def test_deleting_final_element
    list = List.new
    list.push 1
    list.push 456
    list.push 1024

    list.delete(1)
    assert_equal(1024, list.pop)
    assert_equal(456, list.pop)
  end

  def test_insert_after
    list = List.new
    list.push 1
    list.push 456
    list.push 1024

    list.insert_after(456, 'abc')
    assert_equal(1024, list.pop)
    assert_equal(456, list.pop)
    assert_equal('abc', list.pop)
    assert_equal(1, list.pop)
  end

  def test_append
    list = List.new
    list.push 1
    list.push 456
    list.append(1024)

    assert_equal(456, list.pop)
    assert_equal(1, list.pop)
    assert_equal(1024, list.pop)
    assert_equal(nil, list.pop)
  end

  def test_nth_to_last
    list = List.new
    list.push 1
    list.push 2
    list.push 3
    list.push 4
    list.push 5

    assert_equal(1, list.nth_to_last(0))
    assert_equal(2, list.nth_to_last(1))
    assert_equal(3, list.nth_to_last(2))
    assert_equal(4, list.nth_to_last(3))
    assert_equal(5, list.nth_to_last(4))
    assert_equal(nil, list.nth_to_last(6))
  end
end
