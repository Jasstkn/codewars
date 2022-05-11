package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func convertTime(num int) (int, int, int) {
	h := num / 3600
	m := (num - h*3600) / 60
	s := num - h*3600 - m*60
	return h, m, s
}

func parseData(t string) int {
	team := strings.Split(t, "|")
	teamH, _ := strconv.Atoi(team[0])
	teamM, _ := strconv.Atoi(team[1])
	teamS, _ := strconv.Atoi(team[2])
	return teamH*3600 + teamM*60 + teamS
}

func calculateMedian(m []int) int {
	if len(m)%2 != 0 {
		return m[len(m)/2]
	}
	return (m[len(m)/2] + m[len(m)/2-1]) / 2
}

func Stati(strg string) string {
	if strg == "" {
		return ""
	}

	var sum int

	data := strings.Split(strg, ", ")
	teamTimeList := []int{}

	for _, t := range data {
		teamTimeSec := parseData(t)
		// calculate average
		sum += teamTimeSec

		// calculate median
		teamTimeList = append(teamTimeList, teamTimeSec)
	}
	// sort list
	sort.Ints(teamTimeList)

	tMax, tMin := teamTimeList[len(teamTimeList)-1], teamTimeList[0]
	rH, rM, rS := convertTime(tMax - tMin)

	average := sum / len(data)
	aH, aM, aS := convertTime(average)

	mH, mM, mS := convertTime(calculateMedian(teamTimeList))
	return fmt.Sprintf("Range: %02d|%02d|%02d Average: %02d|%02d|%02d Median: %02d|%02d|%02d", rH, rM, rS, aH, aM, aS, mH, mM, mS)
}

func main() {
	start := time.Now()
	fmt.Println(Stati("01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
