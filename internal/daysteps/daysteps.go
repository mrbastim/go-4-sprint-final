package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mrbastim/go-4-sprint-final/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, time.Duration(0), err
	}
	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, time.Duration(0), err
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return "error"
	} else if steps < 0 {
		return ""
	}
	distance := stepLength * float64(steps) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return err.Error()
	}
	return `Количество шагов: ` + strconv.Itoa(steps) + `.
			Дистанция составила ` + fmt.Sprintf("%.2f", distance) + ` км.
			Вы сожгли ` + fmt.Sprintf("%.2f", calories) + ` ккал.`
}
