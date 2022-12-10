package main
import "fmt"

func main() {
	arrF := []float32{0.1,0.2,0.3}
	result := Sum(arrF)
	fmt.Printf("Sum is %f\n", result)

	b := []int{1,2,3,4}
	sum, av := SumAndAverage(b)
	fmt.Printf("Sum is %d, Average is %f\n", sum, av)

	min := minSlice(b)
	max := maxSlice(b)
	fmt.Printf("min is %d, max is %d\n", min, max)

	s := b[1:1]
	fmt.Printf("len is %d\n", len(s))

}

func Sum(arrF []float32) (sum float32) {
	sum = 0
	for _, val := range arrF {
		sum += val
	}
	fmt.Printf("Sum is %f\n", sum)
	return
}

func SumAndAverage(arr []int) (sum int, average float32) {
	sum = 0
	for _, val := range arr {
		sum += val
	}
	average = float32(sum / len(arr))
	return
}

func minSlice(slice []int) (min int) {
	for _, val := range slice {
		if val < min {
			min = val
		}
	}
	return
} 

func maxSlice(slice []int) (max int) {
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return
} 