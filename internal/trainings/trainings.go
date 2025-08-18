package trainings

import(
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"fmt"
	"strings"
	"time"
	"strconv"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)
type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// Разделяем строку
	sepData := strings.Split(datastring, ",")

	// Проверяем правильность разделения
	if len(sepData) != 3 {
		return fmt.Errorf("incorrect input data:%s", datastring)
	}

	// Преобразуем количество шагов в число
	steps, err := strconv.Atoi(sepData[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования числа шагов: %s, ошибка:%w", sepData[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов 0 или меньше: %s", sepData[0])
	}
	t.Steps = steps

	// Сохраняем тип тренировки
	t.TrainingType = sepData[1]

	//Парсинг времени
	duration, err := time.ParseDuration(sepData[2])
	if err != nil {
		return fmt.Errorf("ошибка преобразования продолжительности тренировки: %s, ошибка:%w", sepData[2], err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность тренировки 0 или меньше: %s", sepData[2])
	}
	t.Duration = duration
	return nil
}


func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var cal float64
    var err error

	//калории по типу тренировки
	switch t.TrainingType {
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)	
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
			return "", fmt.Errorf("ошибка в вычислении калорий:%w", err)
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, speed, cal), nil
}
