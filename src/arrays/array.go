package arrays

func Sum(numbers []int) int{
	sum := 0
	//i := 0
	//if (i != 0) {
	//	sum += 1
	//}
	//a := "chen"
	//fmt.Fprintln(len(a))
	//
	for _, number := range numbers {
		sum += number
	}
	return sum
}
