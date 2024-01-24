package sort

func Bubble(array []int) []int {
	length := len(array)
	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < length-1; i++ {
			if array[i+1] < array[i] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
		}
	}
	return array
}
func Selection(array []int) []int {
	length := len(array)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func Insertion(array []int) []int {
	length := len(array)

	for i := 0; i < length; i++ {
		x := array[i]
		j := i
		for ; j >= 0 && array[j] > array[i]; j-- {
			array[j] = array[j-1]
		}
		array[j] = x
	}
	return array
}

/**
 * The complexity of Quicksort with this scheme degrades to O(n2)
 * when the array is already in order, due to the partition being the worst possible one.
 */
func QuickLomuto(array []int, lo, hi int) []int {

	if lo >= hi || lo < 0 {

		return array
	}
	array, p := partition(array, lo, hi)
	array = QuickLomuto(array, lo, p-1)
	array = QuickLomuto(array, p+1, hi)

	return array

}

func partition(array []int, lo, hi int) ([]int, int) {

	i := lo - 1
	pivot := array[lo]
	for j := lo; j < hi; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	i++
	array[i], array[hi] = array[hi], array[i]

	return array, i
}
