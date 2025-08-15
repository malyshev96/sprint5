package daysteps

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разделяем строку
	sepData := strings.Split(datastring, ",")

	// Проверяем правильность разделения
	if len(sepData) != 2 {
		return fmt.Errorf("не удалось обработать данные")
	}

	// Преобразуем количество шагов в число
	steps, errSteps := strconv.Atoi(sepData[0])
	if errSteps != nil || steps <= 0 {
		return fmt.Errorf("ошибка преобразования числа шагов: %s", sepData[0])
	}
	ds.Steps = steps

	//Парсинг времени
	duration, errTime := time.ParseDuration(sepData[1])
	if errTime != nil || duration <= 0 {
		return fmt.Errorf("ошибка преобразования времени ходьбы: %s", sepData[1])
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	//Проверка на ноль
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", fmt.Errorf("данные должны быть больше 0")
	}

	//Вычисление дистанции и калорий
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err:= spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	//Формируем строку
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cal), nil
}
