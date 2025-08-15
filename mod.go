package red

import (
	"reflect"
	"sync"
)

type TaskType int

const (
	TaskType_Provide = TaskType(iota)
	TaskType_Invoke
)

type Task struct {
	typ TaskType
	fn  func(*Container) error
}

type Container struct {
	mu        sync.RWMutex
	im        sync.RWMutex
	instances map[reflect.Type]reflect.Value
	providers map[reflect.Type]reflect.Value
	tasks     []Task
}

var instance *Container

func New() *Container {
	instance = &Container{
		instances: make(map[reflect.Type]reflect.Value),
		providers: make(map[reflect.Type]reflect.Value),
	}

	return instance
}

func NewStandalone() *Container {
	return &Container{
		instances: make(map[reflect.Type]reflect.Value),
		providers: make(map[reflect.Type]reflect.Value),
	}
}

func containsError(typ reflect.Type) bool {
	return typ.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem())
}
