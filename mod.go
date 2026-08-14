package red

import (
	"reflect"
	"sync"
)

type TaskType int

const (
	TaskType_Provide = TaskType(iota)
	TaskType_Invoke
	TaskType_Prepare
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

	instance.Use(Provide(contextProvider))

	return instance
}

func NewStandalone() *Container {
	inst := &Container{
		instances: make(map[reflect.Type]reflect.Value),
		providers: make(map[reflect.Type]reflect.Value),
	}

	inst.Use(Provide(contextProvider))

	return inst
}

func containsError(typ reflect.Type) bool {
	return typ.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem())
}
