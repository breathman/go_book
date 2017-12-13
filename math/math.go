package math

//Найти среднее значение числового массива
func Average(arr []float64) float64 {
	total := 0.0
	for _, i := range arr {
		total = i
	}
	return total / float64(len(arr))
}