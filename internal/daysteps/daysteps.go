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
		return fmt.Errorf("incorrect input data:%s", datastring)
	}

	// Преобразуем количество шагов в число
	steps, err := strconv.Atoi(sepData[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования числа шагов: %s, Ошибка:%w", sepData[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов 0 или меньше: %s", sepData[0])
	}
	ds.Steps = steps

	//Парсинг времени
	duration, err := time.ParseDuration(sepData[1])
	if err != nil{
		return fmt.Errorf("ошибка преобразования времени ходьбы: %s, Ошибка:%w", sepData[1], err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность тренировки 0 или меньше: %s", sepData[1])
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	//Проверка на ноль
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", fmt.Errorf("данные должны быть больше 0. Шаги:%d, Продолжительность:%s, Вес:%f, Рост:%f", ds.Steps, ds.Duration, ds.Weight, ds.Height)
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
