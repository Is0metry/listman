package list

type Database interface {
	GetList(name string) (*List, error)
	AddItem(name string, item string) error
	RemoveItem(name string, item int) error
}
