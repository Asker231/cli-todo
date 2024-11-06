package todo

type Methods interface{
	Add()
	Remove(id int)error
	Update(id int,title string)error
}
type Todos []Todo

func(t *Todos)Add(title ,body string)*Todos{
	counter := Counter()

    todo := Todo{
		Id: counter(),
		Title: title,
		Body: body,
	} 
	*t = append(*t, todo)
	return &Todos{}
}

func(t *Todos)Remove(id int) (*Todos,error) {

	return &Todos{},nil
}

func(t *Todos)Update(id int) (*Todos,error) {
    
	return &Todos{},nil
}

func Counter()func()int{
	count := 0
	return func() int {
		count = count +1
		return count
	}
}
