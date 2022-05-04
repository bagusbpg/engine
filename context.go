package context

import . "context"

const contextKeyName = "WithInfoContextData"

type withInfoBuilder struct {
	parent Context
	info   map[string]string
}

func WithInfoBuilder(parent Context) *withInfoBuilder {
	return &withInfoBuilder{
		parent: parent,
		info:   make(map[string]string),
	}
}
