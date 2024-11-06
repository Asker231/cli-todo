package internal


	type Operation interface{
		AddTodo(todo Todo)
		RemoveTodo(id int)
	}

	type Input struct{
		input any
	}

	func NewInput(input any)*Input{
		return &Input{
			input: input,
		}
	}