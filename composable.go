package red

func (c *Container) Use(steps ...Task) *Container {
	c.tasks = append(c.tasks, steps...)
	return c
}

func (c *Container) Run() error {

	dieCh := make(chan error)

	for _, task := range c.tasks {

		switch task.typ {
		case TaskType_Provide:
			if err := task.fn(c); err != nil {
				return err
			}
		case TaskType_Invoke:
			go func(task func(*Container) error) {
				if err := task(c); err != nil {
					dieCh <- err
				}
			}(task.fn)
		}
	}

	if err, ok := <-dieCh; err != nil && ok {
		return err
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
