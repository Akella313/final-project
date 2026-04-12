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
	ruleMonth   = "m"
	ruleWeek    = "w"
)

// если не указано правило для повторения задачи, её надо удалять
// если в правиле повторения задачи указаны дни, задачу надо перенести на указанное количество дней, но не больше 400
// если в правиле указан год, задачу надо перенести на календарный год

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// проверка на то, что правило существует
	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// парсинг даты в time.Time.
	date, err := time.ParseInLocation(dateLayout, dstart, now.Location())
	if err != nil {
		return "", err
	}

	// заготовка, чтобы определять правило
	// в каком виде будет поступать строка repeat?
	// если в случае месяца или недели, в числах будет пробел после запятой, то заготовка развалится
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
		date = date.AddDate(0, 0, days)
		for !date.After(now) {
			date = date.AddDate(0, 0, days)
		}
		return date.Format(dateLayout), nil
	case ruleYear:
		// прописано правило для проверки правила года
		if len(parts) != 1 {
			return "", errors.New("invalid repeat")
		}
		// прописать логику переноса года
		date = date.AddDate(1, 0, 0)
		for !date.After(now) {
			date = date.AddDate(1, 0, 0)
		}
		return date.Format(dateLayout), nil
	case ruleWeek:
		if len(parts) != 2 {
			return "", errors.New("invalid repeat")
		}
		weekDays := make([]int, 0, 7)
		weekParts := strings.Split(parts[1], ",")
		for _, v := range weekParts {
			weekDay, err := strconv.Atoi(v)
			if err != nil {
				return "", errors.New("error parsing week days")
			}
			if weekDay < 1 || weekDay > 7 {
				return "", errors.New("invalid weekday")
			}
			weekDays = append(weekDays, weekDay)
		}
		for i := 0; i < 7; i++ {
			date = date.AddDate(0, 0, 1)

			currentWeekday := int(date.Weekday())
			if currentWeekday == 0 {
				currentWeekday = 7
			}

			for _, allowedDays := range weekDays {
				if allowedDays == currentWeekday {
					return date.Format(dateLayout), nil
				}
			}
		}
		return "", errors.New("failed to find next week day")
		// незаконченный блок для правила месяца
	/*
		case ruleMonth:
			if len(parts) < 2 || len(parts) > 3 {
				return "", errors.New("invalid repeat")
			}
			monthDays := make([]int, 0)
			months := make([]int, 0, 12)
			for _, v := range strings.Split(parts[1], ",") {
				monthDay, err := strconv.Atoi(v)
				if err != nil {
					return "", errors.New("error parsing month days")
				}
				if monthDay != -1 && monthDay != -2 {
					if monthDay < 1 || monthDay > 31 {
						return "", errors.New("invalid monthday")
					}
				}
				monthDays = append(monthDays, monthDay)
			}
			if len(parts) == 3 {
				for _, v := range strings.Split(parts[2], ",") {
					m, err := strconv.Atoi(v)
					if err != nil {
						return "", errors.New("error parsing months")
					}
					if m < 1 || m > 12 {
						return "", errors.New("invalid month")
					}
					months = append(months, m)
				}
			}

	*/
	default:
		return "", errors.New("invalid repeat rule")
	}

}
