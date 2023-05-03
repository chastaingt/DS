package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (m *MaxHeap) insert(key int) {
	m.array = append(m.array, key)
	m.maxHeapifyUp(len(m.array) - 1)
}

func (m *MaxHeap) Extract() int {
	extracted := m.array[0]
	lastIndex := len(m.array) - 1

	if len(m.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	m.array[0] = m.array[lastIndex]
	m.array = m.array[:lastIndex]

	m.maxHeapifyDown(0)

	return extracted
}

func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.array[parent(index)] < m.array[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(m.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.array[l] > m.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.array[index] < m.array[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (m *MaxHeap) swap(index1, index2 int) {
	m.array[index1], m.array[index2] = m.array[index2], m.array[index1]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
