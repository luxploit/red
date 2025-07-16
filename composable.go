package red

import "sync"

func (c *Container) Use(steps ...Task) *Container {
	c.tasks = append(c.tasks, steps...)
	return c
}

func (c *Container) Run() error {
	wg := sync.WaitGroup{}

	//! This is terrible but such is life

	invTasks := 0
	for _, task := range c.tasks {
		if task.typ == TaskType_Invoke {
			invTasks++
		}
	}
	invErrs := make(chan error, invTasks)

	for _, task := range c.tasks {

		switch task.typ {
		case TaskType_Provide:
			if err := task.fn(c); err != nil {
				return err
			}
		case TaskType_Invoke:
			wg.Add(1)
			go func(task func(*Container) error) {
				defer wg.Done()
				if err := task(c); err != nil {
					invErrs <- err
				}
			}(task.fn)
		}
	}

	wg.Wait()
	close(invErrs)

	for err := range invErrs {
		if err != nil {
			return err
		}
	}

	return nil
}

func Provide(provider any) Task {
	return Task{
		fn: func(c *Container) error {
			return c.provide(provider)
		},
		typ: TaskType_Provide,
	}
}

func Invoke(fn any) Task {
	return Task{
		fn: func(c *Container) error {
			return c.invoke(fn)
		},
		typ: TaskType_Invoke,
	}
}
