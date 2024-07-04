package main

func main() {

}

type DynamicArray struct {
	array    []int
	size     int
	capacity int
}

func (d *DynamicArray) get(i int) int {
	return d.array[i]
}

func (d *DynamicArray) set(i int, n int) {
	d.array[i] = n
}

func (d *DynamicArray) pushback(n int) {

	d.array = append(d.array, n)
	d.size += 1
}

func (d *DynamicArray) popback() int {
	lastElement := d.array[len(d.array)-1]
	d.array = d.array[:len(d.array)-1]
	d.size -= 1
	return lastElement
}

func (d *DynamicArray) resize() {
	d.capacity *= 2
	newArray := make([]int, d.capacity)
	for i := range d.array {
		newArray[i] = d.array[i]
	}
	d.array = newArray
}

func (d *DynamicArray) getSize() int {
	// return number of elements in the array
	return d.size
}

func (d *DynamicArray) getCapacity() int {
	return d.capacity
}
