package covidtracker

import (
	"fmt"
	"sort"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const rollAmount int = 7

type daysTo struct {
	sum  int
	days int
}

type metadata struct {
	weekdayAverages  map[int]float64
	rollingAverages  []int
	rawValues        []int
	rawTotals        []int
	daysToMultiplier []daysTo
}

type expectedChange struct {
	expected   int
	actual     int
	sixDaysAgo int
}

func GetText(index int, data []DailyCovidData) string {
	sort.Sort(byDate(data))

	return getUIForIndex(index, &data)
}

func getUIForIndex(index int, data *[]DailyCovidData) string {
	extractDeaths := func(data DailyCovidData) (int, int) { return data.NewDeaths, data.TotalDeaths }
	extractInfections := func(data DailyCovidData) (int, int) { return data.NewInfections, data.TotalInfections }

	deathMetadata := getMetadata(data, extractDeaths)
	infectionMetadata := getMetadata(data, extractInfections)

	output := fmt.Sprintln("Data for", (*data)[index].Date.Format("2006-01-02"))
	output += reportData("Deaths", index, deathMetadata, data, extractDeaths)
	output += reportDaysToSignificantValues(index, 30000, data, extractDeaths)
	output += reportData("Infections", index, infectionMetadata, data, extractInfections)
	output += reportDaysToSignificantValues(index, 1000000, data, extractInfections)

	return output
}

func reportData(title string, index int, info metadata, data *[]DailyCovidData, extractValue func(DailyCovidData) (int, int)) string {
	dailyData := (*data)[index]
	newData, totalData := extractValue(dailyData)

	output := fmt.Sprintln("\n\tTotal", title+":", formatNum(totalData))
	output += fmt.Sprintln("\tNew", title+":", formatNum(newData))
	output += fmt.Sprintln("\tChange from last week:", formatAsPercent(getChange(info.rawValues, index, rollAmount, 0)))

	if sevenDaysAgo := getFromDaysAgo(index, 7, info.rawValues); sevenDaysAgo > 0 {
		output += fmt.Sprintln("\t\t7 days ago:", formatNum(sevenDaysAgo))
	}

	expected := getExpectedValueTomorrow(index, data, info)
	output += fmt.Sprintln("\tPredicted next day:", formatNum(expected.expected))

	if expected.actual > 0 {
		output += fmt.Sprintln("\t\tActual:", formatNum(expected.actual))
	}

	if expected.sixDaysAgo > 0 {
		output += fmt.Sprintln("\t\t6 days ago:", formatNum(expected.sixDaysAgo))
	}

	output += fmt.Sprintln("\tRolling average ("+formatNum(rollAmount)+"-day):", formatNum(info.rollingAverages[index]))

	output += fmt.Sprintln("\tRolling average change from last week:", formatAsPercent(getChange(info.rollingAverages, index, rollAmount, 0)))

	if rollSevenDaysAgo := getFromDaysAgo(index, 7, info.rollingAverages); rollSevenDaysAgo > 0 {
		output += fmt.Sprintln("\t\t7 days ago:", formatNum(rollSevenDaysAgo))
	}

	return output
}

func reportDaysToSignificantValues(index int, multiplier int, data *[]DailyCovidData, extractValue func(DailyCovidData) (int, int)) string {
	output := fmt.Sprintln("\tDays to significant values:")

	value := multiplier
	prevDays := 0
	rawValues := getNewValueArray(data, extractValue)
	for i := 0; i <= index; i++ {
		ttv, err := GetTimeToValue(rawValues, value, index)
		if err != nil {
			return output
		}

		output += fmt.Sprintln("\t\t"+strconv.Itoa(ttv-prevDays), "day(s) to", formatNum(value))
		value += multiplier
		prevDays = ttv
	}

	return output
}

func formatNum(num int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", num)
}

func getExpectedValueTomorrow(index int, data *[]DailyCovidData, info metadata) expectedChange {
	currentDayOfWeek := int((*data)[index].Date.Weekday())
	nextDayOfWeek := currentDayOfWeek + 1
	if nextDayOfWeek > 6 {
		nextDayOfWeek = 0
	}

	curDayAvg, nextDayAvg := info.weekdayAverages[currentDayOfWeek], info.weekdayAverages[nextDayOfWeek]
	dayOfWeekChange := ((nextDayAvg - curDayAvg) / curDayAvg) + 1.0

	currentDay := info.rawValues[index]
	expected := float64(currentDay) * dayOfWeekChange

	actual := getFromDaysAgo(index, -1, info.rawValues)
	sixDaysAgo := getFromDaysAgo(index, 6, info.rawValues)

	return expectedChange{
		expected:   int(expected),
		actual:     actual,
		sixDaysAgo: sixDaysAgo,
	}
}

func getFromDaysAgo(index int, daysAgo int, data []int) int {
	dataFromDaysAgo := -1
	daysAgoIndex := index - daysAgo
	if daysAgoIndex > 0 && daysAgoIndex < len(data) {
		dataFromDaysAgo = data[daysAgoIndex]
	}

	return dataFromDaysAgo
}

func formatAsPercent(value float64) string {
	s := fmt.Sprintf("%.2f", value*100.0) + "%"

	if value > 0 {
		return "+" + s
	}

	return s
}

func getChange(data []int, index int, rewind int, defaultValue float64) float64 {
	previousIndex := index - rewind

	if previousIndex < 0 {
		return defaultValue
	}

	currentValue, previousValue := data[index], data[previousIndex]
	change := currentValue - previousValue

	return (float64(change) / float64(previousValue))
}

func getMetadata(data *[]DailyCovidData, extractValue func(DailyCovidData) (int, int)) metadata {
	return metadata{
		weekdayAverages: getWeekdayAverages(data, extractValue),
		rollingAverages: getRollingAverages(data, extractValue),
		rawValues:       getNewValueArray(data, extractValue),
	}
}

func getWeekdayAverages(data *[]DailyCovidData, extractValue func(DailyCovidData) (int, int)) map[int]float64 {
	inputData := make([]WeekdayValuePair, len(*data))
	for index, element := range *data {
		newData, _ := extractValue(element)
		inputData[index] = WeekdayValuePair{
			Value:   newData,
			Weekday: int(element.Date.Weekday()),
		}
	}

	return GetWeekdayAverages(inputData)
}

func getRollingAverages(data *[]DailyCovidData, extractValue func(DailyCovidData) (int, int)) []int {
	return RollAverage(getNewValueArray(data, extractValue), rollAmount)
}

func getNewValueArray(data *[]DailyCovidData, extractValue func(DailyCovidData) (int, int)) []int {
	inputData := make([]int, len(*data))
	for index, element := range *data {
		new, _ := extractValue(element)
		inputData[index] = new
	}

	return inputData
}
