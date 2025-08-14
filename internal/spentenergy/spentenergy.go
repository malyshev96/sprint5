package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных параметров
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("параметры должны быть больше 0")
	}

	//Расчет калорий
	cal := (MeanSpeed(steps, height, duration) * weight * duration.Minutes()) / float64(minInH) * walkingCaloriesCoefficient

	return cal, nil
}


func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных параметров
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("параметры должны быть больше 0")
	}

	//Расчет калорий
	cal := (MeanSpeed(steps, height, duration) * weight * duration.Minutes()) / float64(minInH)

	return cal, nil
}


func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверка длительности на ноль
	if duration <= 0 {
		return 0
	}

	// Скорость
	speed := Distance(steps, height) / duration.Hours()

	return speed
}


func Distance(steps int, height float64) float64 {
	// Длина шага
	stepLength := height * stepLengthCoefficient

	// Дистанция
	dist := stepLength * float64(steps) / float64(mInKm)

	return dist
}

