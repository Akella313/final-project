package api

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	dateLayout  = "20060102"
	maxInterval = 400
	ruleDay     = "d"
	ruleYear    = "y"
	ruleWeek    = "w"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	date, err := time.ParseInLocation(dateLayout, dstart, now.Location())
	if err != nil {
		return "", err
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", errors.New("repeat is empty")
	}

	switch parts[0] {
	case ruleDay:
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
		date = date.AddDate(0, 0, days)
		for !date.After(now) {
			date = date.AddDate(0, 0, days)
		}
		return date.Format(dateLayout), nil
	case ruleYear:
		if len(parts) != 1 {
			return "", errors.New("invalid repeat")
		}
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
	default:
		return "", errors.New("invalid repeat rule")
	}

}
