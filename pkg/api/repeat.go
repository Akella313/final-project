package api

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// определение основных переменных в константы
const (
	dateLayout  = "20060102"
	maxInterval = 400
	ruleDay     = "d"
	ruleYear    = "y"
)

// если не указано правило для повторения задачи, её надо удалять
// если в правиле повторения задачи указаны дни, задачу надо перенести на указанное количество дней, но не больше 400
// если в правиле указан год, задачу надо перенести на календарный год

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// проверка на то, что правило существует
	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	// парсинг даты в time.Time.
	date, err := time.Parse(dateLayout, dstart)
	if err != nil {
		return "", err
	}

	// заготовка, чтобы определять правило
	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", errors.New("repeat is empty")
	}
	// определение типа правила через switch, по первому элементу массива
	switch parts[0] {
	case ruleDay:
		// прописаны правила для проверки правила дня и парсинг количества дней в int.
		if len(parts) != 2 {
			return "", errors.New("invalid repeat")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", errors.New("error parsing days")
		}
		if days <= 0 || days > maxInterval {
			return "", errors.New("invalid amount days")
		}
		// прописать логику для переноса даты
	case ruleYear:
		// прописано правило для проверки правила года
		if len(parts) != 1 {
			return "", errors.New("invalid repeat")
		}
		// прописать логику переноса года
	default:
		return "", errors.New("invalid repeat rule")
	}
}
