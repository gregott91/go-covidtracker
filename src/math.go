package covidtracker

import "errors"

type WeekdayValuePair struct {
	Weekday int
	Value   int
}

func GetTimeToValue(values []int, value int, maxIndex int) (int, error) {
	totalSum := 0
	for i := 0; i <= maxIndex; i++ {
		element := values[i]
		totalSum += element

		if totalSum > value {
			return i, nil
		}
	}

	return -1, errors.New("Value never reached")
}

func GetWeekdayAverages(values []WeekdayValuePair) map[int]float64 {
	sums := make(map[int]int)
	totalSum := 0
	for _, element := range values {
		sums[element.Weekday] += element.Value
		totalSum += element.Value
	}

	results := make(map[int]float64)
	for i := 0; i < 7; i++ {
		results[i] = (float64(sums[i]) / float64(totalSum))
	}

	return results
}

func RollAverage(values []int, average int) []int {
	result := make([]int, len(values))

	for index := range values {
		result[index] = getRolledAverage(values, average, index)
	}

	return result
}

func getRolledAverage(values []int, average int, index int) int {
	look := int(average / 2)

	lookStart, lookEnd := look, look
	valuesAfter := len(values) - index

	if index < lookStart {
		lookStart = index
		lookEnd += (look - lookStart)
	} else if valuesAfter <= lookEnd {
		lookEnd = valuesAfter - 1
		lookStart += (look - lookEnd)
	}

	count := 0
	sum := 0
	for i := index - lookStart; i <= (index + lookEnd); i++ {
		count++
		sum += values[i]
	}

	return int(sum / count)
}
