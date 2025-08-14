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
		return fmt.Errorf("неправильно введены данные")
	}

	// Преобразуем количество шагов в число
	steps, errSteps := strconv.Atoi(sepData[0])
	if errSteps != nil || steps <= 0 {
		return fmt.Errorf("ошибка преобразования числа шагов: %s", sepData[0])
	}
	t.Steps = steps

	// Сохраняем тип тренировки
	t.TrainingType = sepData[1]

	//Парсинг времени
	duration, errTime := time.ParseDuration(sepData[2])
	if errTime != nil || duration <= 0 {
		return fmt.Errorf("ошибка преобразования продолжительности тренировки: %s", sepData[2])
	}
	t.Duration = duration
	return nil
}


func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	//калории по типу тренировки
	switch t.TrainingType {
	case "Ходьба":
		cal, errCal := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if errCal != nil {
			return "", fmt.Errorf("%w", errCal)
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, speed, cal), nil
	case "Бег":
		cal, errCal := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if errCal != nil {
			return "", fmt.Errorf("%w", errCal)
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, speed, cal), nil		
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
