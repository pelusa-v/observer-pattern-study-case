package user_monitor

type Monitor interface {
	RegisterObserver(obs UserObserver)
	RemoveObserver(obs UserObserver)
	// NotifyUserCreated(u User)
	// DiscoverNeighbors()
}

type WeatherMonitor struct {
	Users   []UserObserver
	Weather string
}

func NewUserMonitor() WeatherMonitor {
	return WeatherMonitor{}
}

func (mon *WeatherMonitor) RegisterObserver(obs UserObserver) {
	mon.Users = append(mon.Users, obs)
}

func (mon *WeatherMonitor) RemoveObserver(obs UserObserver) {
	for i, o := range mon.Users {
		if o == obs {
			mon.Users = append(mon.Users[:i], mon.Users[i+1:]...)
			break
		}
	}
}

func (mon *WeatherMonitor) SetWeather(weather string) {
	mon.Weather = weather
	for _, o := range mon.Users {
		o.UpdateWeather(weather)
	}
}

// func (mon *UserMonitor) NotifyUserCreated(u User) {
// 	for _, o := range mon.Users {
// 		if o.(*User).Id != u.Id {
// 			o.AddNeighbor(u)
// 		}
// 	}
// }

// func (mon *UserMonitor) DiscoverNeighbors() {
// 	totalUsers := mon.Users
// 	for _, u1 := range totalUsers {
// 		for _, u2 := range totalUsers {
// 			if u1.(*User).Id != u2.(*User).Id {
// 				u1.AddNeighbor(*u2.(*User))
// 			}
// 		}
// 	}
// }
