package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

//const inputPath = "input2.txt"

const inputPath = "input.txt"

var (
	TimeRE  = regexp.MustCompile("\\[(.*)\\]")
	GuardRE = regexp.MustCompile("#(.[^\\s]+)")
)

func getGuard(line string) (int, error) {
	matches := GuardRE.FindStringSubmatch(line)
	if len(matches) != 2 {
		return -1, fmt.Errorf("[getguard] invalid syntax: %s", line)
	}

	return strconv.Atoi(matches[1])
}

func getTime(line string) (time.Time, error) {
	matches := TimeRE.FindStringSubmatch(line)
	if len(matches) != 2 {
		return time.Now(), fmt.Errorf("[gettime] invalid syntax: %s", line)
	}

	return time.Parse("2006-01-02 15:04", matches[1])
}

type TimeInput struct {
	T    time.Time
	Line string
}

type TimeInputs []TimeInput

func (ts TimeInputs) Len() int {
	return len(ts)
}

func (ts TimeInputs) Less(i, j int) bool {
	return ts[i].T.Before(ts[j].T)
}

func (ts TimeInputs) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func readerToInputs(input io.Reader) (TimeInputs, error) {
	var inputs TimeInputs

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		t, err := getTime(line)
		if err != nil {
			return inputs, err
		}

		ti := TimeInput{
			T:    t,
			Line: line,
		}

		inputs = append(inputs, ti)
	}

	sort.Sort(inputs)

	return inputs, nil
}

func solvePart1(inputs []TimeInput) (result int, err error) {
	// Map of guard ID to the minute to the count
	guardMap := make(map[int]map[int]int)

	var curGuard int
	var lastSleepMinute int
	for _, input := range inputs {

		curMinute := input.T.Minute()

		if strings.Contains(input.Line, "falls asleep") {
			if _, ok := guardMap[curGuard][curMinute]; !ok {
				guardMap[curGuard][curMinute] = 0
			}

			lastSleepMinute = curMinute
		} else if strings.Contains(input.Line, "wakes up") {
			// Need to add entries from when he fell asleep to when
			// he woke up
			minutesAsleep := input.T.Minute() - lastSleepMinute
			for i := minutesAsleep; i > 0; i-- {
				if _, ok := guardMap[curGuard][input.T.Minute()-i]; !ok {
					guardMap[curGuard][input.T.Minute()-i] = 0
				}

				guardMap[curGuard][input.T.Minute()-i] += 1
			}
		} else {
			// Initialize new guard starting
			curGuard, err = getGuard(input.Line)
			if err != nil {
				return -1, err
			}

			if _, ok := guardMap[curGuard]; !ok {
				guardMap[curGuard] = make(map[int]int)
			}
		}
	}

	// Now we have our inputs all entered into our guardMap; let's figure
	// out which guard needs a dose of 5 hour energy
	var sleepiestGuard, highestSleepCount int
	for guardKey, minuteMap := range guardMap {

		var curSleepCount int
		for _, count := range minuteMap {
			curSleepCount += count
		}

		if curSleepCount > highestSleepCount {
			highestSleepCount = curSleepCount
			sleepiestGuard = guardKey
		}

	}

	fmt.Println("sleepiest guard", sleepiestGuard)
	fmt.Println("min slept", highestSleepCount)

	// Now we can find out which minute the guard slept the longest (we
	// could have done this above.. but #lazyadvent)
	var sleepiestMinute, minuteCount int
	for minute, count := range guardMap[sleepiestGuard] {
		if count > minuteCount {
			minuteCount = count
			sleepiestMinute = minute
		}
	}

	return sleepiestMinute * sleepiestGuard, nil
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	inputs, err := readerToInputs(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(solvePart1(inputs))
}
