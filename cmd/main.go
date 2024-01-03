package main

import (
	"fmt"

	"github.com/pelusa-v/observer-pattern-study-case.git/internal/user_monitor"
)

func main() {
	monitor := user_monitor.NewUserMonitor()
	u1 := user_monitor.NewUser("Bob", 0, &monitor)
	u2 := user_monitor.NewUser("Carl", 1, &monitor)
	u3 := user_monitor.NewUser("Charlie", 2, &monitor)
	u4 := user_monitor.NewUser("Tom", 3, &monitor)
	u5 := user_monitor.NewUser("Roy", 4, &monitor)
	u6 := user_monitor.NewUser("Tiny", 5, &monitor)

	monitor.SetWeather("Verano")

	fmt.Println(u1.Name + " -> " + u1.Weather)
	fmt.Println(u2.Name + " -> " + u2.Weather)
	fmt.Println(u3.Name + " -> " + u3.Weather)
	fmt.Println(u4.Name + " -> " + u4.Weather)
	fmt.Println(u5.Name + " -> " + u5.Weather)
	fmt.Println(u6.Name + " -> " + u6.Weather)

	fmt.Println("----------------------------")

	monitor.RemoveObserver(u5)
	monitor.RemoveObserver(u1)
	monitor.SetWeather("Invierno")

	fmt.Println(u1.Name + " -> " + u1.Weather)
	fmt.Println(u2.Name + " -> " + u2.Weather)
	fmt.Println(u3.Name + " -> " + u3.Weather)
	fmt.Println(u4.Name + " -> " + u4.Weather)
	fmt.Println(u5.Name + " -> " + u5.Weather)
	fmt.Println(u6.Name + " -> " + u6.Weather)
}
