package helpers

func GetMonthRu(month string) string {
	switch month {
	case "January":
		return "января"
	case "February":
		return "февраля"
	case "March":
		return "марта"
	case "April":
		return "апреля"
	case "May":
		return "мая"
	case "June":
		return "июня"
	case "July":
		return "июля"
	case "August":
		return "августа"
	case "September":
		return "сентября"
	case "October":
		return "октября"
	case "November":
		return "ноября"
	case "December":
		return "декабря"
	default:
		return ""
	}
}
