//Андрей родился 26 ноября 1993 года. Посчитайте количество дней до его 100-летия — относительно сегодняшнего дня.
package main

import (
	"fmt"
	"time"
)

func main() {
	birthday := time.Date(1993, time.November, 26, 0, 0, 0, 0, time.Local)
	date := birthday.AddDate(100, 0, 0)
	duration := time.Until(date)
	days := int(duration.Hours() / 24)
	fmt.Println(days)
}
