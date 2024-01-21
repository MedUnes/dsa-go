package sort

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
