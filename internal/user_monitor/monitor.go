package user_monitor

type Monitor interface {
	RegisterObserver(obs UserObserver)
	RemoveObserver(obs UserObserver)
	NotifyUserCreated(obs UserObserver)
}

type UserMonitor struct {
	Users []UserObserver
}
