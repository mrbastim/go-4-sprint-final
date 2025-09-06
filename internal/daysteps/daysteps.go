package daysteps

import (
	"fmt"
	"log"
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
	if strings.Count(data, ",") != 1 {
		return 0, 0, fmt.Errorf("invalid format")
	}
	left, right, _ := strings.Cut(data, ",")
	if left == "" || strings.TrimSpace(left) != left {
		return 0, 0, fmt.Errorf("invalid steps")
	}
	if right == "" || strings.TrimSpace(right) != right {
		return 0, 0, fmt.Errorf("invalid duration")
	}
	dataSlice := []string{left, right}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, err
	} else if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, 0, err
	} else if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	} else if steps <= 0 {
		log.Println("количество шагов должно быть больше нуля")
		return ""
	}
	distance := stepLength * float64(steps) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	return `Количество шагов: ` + strconv.Itoa(steps) + `.
Дистанция составила ` + fmt.Sprintf("%.2f", distance) + ` км.
Вы сожгли ` + fmt.Sprintf("%.2f", calories) + ` ккал.
`
}
