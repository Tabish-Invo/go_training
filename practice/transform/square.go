package transform

func SquareSlice(s []int) []int {

	squareS := make([]int, len(s))

	for index, value := range s {
		squareS[index] = value * value
	}

	return squareS
}
