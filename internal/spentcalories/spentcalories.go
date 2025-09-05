package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	// lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 3 {
		return 0, "", 0, fmt.Errorf("неверное количество данных")
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("неверный формат: %w", err)
	}
	trainingType := dataSlice[1]
	duration, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("неверный формат: %w", err)
	}
	return steps, trainingType, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	distance := distance(steps, height)
	if duration <= 0 {
		return 0
	}
	return distance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	switch trainingType {
	case "Ходьба":
		dist := distance(steps, height)
		meanSpeed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf(`Тип тренировки: %s
						Длительность: %v ч.
						Дистанция: %.2f км.
						Скорость: %.2f км/ч
						Сожгли калорий: %.2f`, trainingType, duration.Hours(), dist, meanSpeed, calories), nil
	case "Бег":
		dist := distance(steps, height)
		meanSpeed := meanSpeed(steps, height, duration)
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf(`Тип тренировки: %s
							Длительность: %v ч.
							Дистанция: %.2f км.
							Скорость: %.2f км/ч
							Сожгли калорий: %.2f`, trainingType, duration.Hours(), dist, meanSpeed, calories), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность тренировки должна быть больше нуля")
	} else if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля")
	} else if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	} else if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность тренировки должна быть больше нуля")
	} else if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля")
	} else if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	} else if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / 60
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
