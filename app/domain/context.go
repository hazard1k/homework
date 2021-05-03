package domain

type Context interface {
	Connection() Connection
	PresentersFactory() Presenters
}
