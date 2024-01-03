package user_monitor

type UserObserver interface {
	// AddNeighbor(neighbor User)
	// RemoveNeighbor(neighbor User)
	UpdateWeather(state string)
}

type User struct {
	// Neighbors   []User
	Name        string
	UserMonitor Monitor
	Id          int
	Weather     string
}

func NewUser(name string, id int, monitor Monitor) *User {
	newUser := &User{
		Name:        name,
		Id:          id,
		UserMonitor: monitor,
	}
	newUser.UserMonitor.RegisterObserver(newUser)
	// newUser.UserMonitor.NotifyUserCreated(*newUser)
	// newUser.UserMonitor.DiscoverNeighbors()
	return newUser
}

func (u *User) UpdateWeather(weather string) {
	u.Weather = weather
}

// func (u *User) AddNeighbor(neighbor User) {
// 	u.Neighbors = append(u.Neighbors, neighbor)
// }

// func (u *User) RemoveNeighbor(neighbor User) {
// 	for i, n := range u.Neighbors {
// 		if n.Id == neighbor.Id {
// 			n.Neighbors = append(n.Neighbors[:i], n.Neighbors[i+1:]...)
// 			break
// 		}
// 	}
// }
