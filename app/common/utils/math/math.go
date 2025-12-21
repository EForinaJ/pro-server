package utils_math

func FindMinMax(numbers []float64) (min float64, max float64) {
	if len(numbers) == 0 {
		// 如果切片为空，返回0值
		return 0, 0
	}

	// 初始化min和max为切片的第一个元素
	min = numbers[0]
	max = numbers[0]

	// 遍历切片，更新min和max
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}
