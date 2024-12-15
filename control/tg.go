package control

import (
	"fmt"
	"time"
)

func Tg() {
	t := time.Now()
	fmt.Printf(t.Format("2006.01.02") + " " + t.Format("15:04:05") + " " + "Телеграм сервис недоступен")
	fmt.Println()
}
