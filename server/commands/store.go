package commands

import "go-redis-clone/server/primitives"

type Store interface {
	Delete(string) primitives.Value
	DeleteAll()
	Set(string, primitives.Value)
	Get(string) primitives.Value
	Keys() []string
}
