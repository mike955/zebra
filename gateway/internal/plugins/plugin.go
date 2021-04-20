package plugins

// type Handler func() error

type Plugin interface {
	Handler() error
	Name() string
}
