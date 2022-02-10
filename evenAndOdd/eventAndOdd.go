package evenAndOdd

type oddAndEven []int

func startOddAndEven() oddAndEven {
	oddAndEven := oddAndEven{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return oddAndEven
}

func (o oddAndEven) getOdds() oddAndEven {
	result := oddAndEven{}
	for _, val := range o {
		if val%2 == 0 {
			result = append(result, val)
		}
	}
	return result
}

func (o oddAndEven) getEven() oddAndEven {
	result := oddAndEven{}
	for _, val := range o {
		if val%2 != 0 {
			result = append(result, val)
		}
	}
	return result
}
