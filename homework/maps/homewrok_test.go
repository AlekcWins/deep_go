package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	root *node
	size int
}

type node struct {
	key   int
	value int
	left  *node
	right *node
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	m.root = insertNode(m.root, key, value, &m.size)
}

func insertNode(n *node, key, value int, size *int) *node {
	if n == nil {
		*size++
		return &node{key: key, value: value}
	}

	if key < n.key {
		n.left = insertNode(n.left, key, value, size)
	} else if key > n.key {
		n.right = insertNode(n.right, key, value, size)
	} else {
		n.value = value
	}
	return n
}

func (m *OrderedMap) Erase(key int) {
	m.root = eraseNode(m.root, key, &m.size)
}

func eraseNode(n *node, key int, size *int) *node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = eraseNode(n.left, key, size)
	} else if key > n.key {
		n.right = eraseNode(n.right, key, size)
	} else {
		*size--
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		} else {
			minRight := n.right
			for minRight.left != nil {
				minRight = minRight.left
			}
			n.key = minRight.key
			n.value = minRight.value
			n.right = eraseNode(n.right, minRight.key, size)
		}
	}
	return n
}

func (m *OrderedMap) Contains(key int) bool {
	return containsNode(m.root, key)
}

func containsNode(n *node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return containsNode(n.left, key)
	} else if key > n.key {
		return containsNode(n.right, key)
	}
	return true
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	inOrder(m.root, action)
}

func inOrder(n *node, action func(int, int)) {
	if n == nil {
		return
	}
	inOrder(n.left, action)
	action(n.key, n.value)
	inOrder(n.right, action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}

func TestEraseNonExistentKeys(t *testing.T) {
	m := NewOrderedMap()
	m.Insert(10, 100)
	m.Insert(20, 200)
	m.Insert(30, 300)

	initialSize := m.Size()

	// Удаляем несколько раз несуществующие ключи
	nonExistentKeys := []int{42, 99, -1}
	for _, key := range nonExistentKeys {
		m.Erase(key)
	}
	// проверяем что не произошло лишнее срабатывание size--
	if m.Size() != initialSize {
		t.Errorf("Size changed after deleting non-existent keys. Expected %d, got %d", initialSize, m.Size())
	}

	// Проверяем, что существующие элементы остались
	existingKeys := []int{10, 20, 30}
	for _, key := range existingKeys {
		if !m.Contains(key) {
			t.Errorf("Key %d should still exist in the map", key)
		}
	}
}
