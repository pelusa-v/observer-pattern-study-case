package user_monitor

type UserObserver interface {
	AddNeighbor(neighbor User)
	RemoveNeighbor(neighbor User)
}

type User struct {
	Neighbors   []User
	Name        string
	UserMonitor Monitor
	Id          int
}

func (u *User) AddNeighbor(neighbor User) {
	u.Neighbors = append(u.Neighbors, neighbor)
}

func (u *User) RemoveNeighbor(neighbor User) {
	for i, n := range u.Neighbors {
		if n.Id == neighbor.Id {
			n.Neighbors = append(n.Neighbors[:i], n.Neighbors[i+1:]...)
			break
		}
	}
}
