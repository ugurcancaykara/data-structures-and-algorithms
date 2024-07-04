package main

import "fmt"

func main() {

}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		capacity: 2,
		size:     0,
		array:    make([]int, 2),
	}
}

type DynamicArray struct {
	array          []int
	size, capacity int
}

func (d *DynamicArray) get(i int) int {
	if i < d.size {
		return d.array[i]
	}
	return -1
}

func (d *DynamicArray) set(i int, n int) {
	if i < d.size {
		d.array[i] = n
	}
}

func (d *DynamicArray) pushback(n int) {
	if d.size == d.capacity {
		d.resize()
	}
	d.array[d.size] = n
	d.size += 1
}

func (d *DynamicArray) popback() int {
	if d.size > 0 {
		lastElement := d.array[d.size-1]
		d.array = d.array[:d.size-1]
		d.size -= 1
		return lastElement
	}
	return -1
}

func (d *DynamicArray) resize() {
	d.capacity *= 2
	newArray := make([]int, d.capacity)
	copy(newArray, d.array)
	d.array = newArray
}

func (d *DynamicArray) getSize() int {
	// return number of elements in the array
	return d.size
}

func (d *DynamicArray) getCapacity() int {
	return d.capacity
}

func (d *DynamicArray) Print() {
	for i := 0; i < d.size; i++ {
		fmt.Println(d.array[i])
	}
}
