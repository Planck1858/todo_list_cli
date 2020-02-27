package models

type ToDo struct {
	Index     int    `json:"index"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type List struct {
	Tasks []ToDo `json:"tasks"`
}

/// List ///
func NewList() *List {
	var list List
	list.Tasks = make([]ToDo, 0)
	return &list
}

func (l *List) ClearList() {
	l.Tasks = nil
}

func (l List) GetList() []ToDo {
	return l.Tasks
}

func (l *List) updateList(index int) {
	arr := l.GetList()
	for i := index-1; i < len(l.GetList()); i++ {
		arr[i].Index = i+1
	}

	l.Tasks = arr
}

/// Tasks ///
func (l *List) NewTask(text string) {
	var newIndex int

	if len(l.GetList()) == 0 {
		newIndex = 1
	} else {
		newIndex = l.GetList()[len(l.GetList())-1].Index + 1
	}

	var newToDo = ToDo{newIndex, text, false}
	l.Tasks = append(l.GetList(), newToDo)
}

func (l *List) DeleteTask(index int) {
	l.Tasks = append(l.Tasks[:index], l.Tasks[index+1:]...)
	l.updateList(index)
}

func (l *List) ChangeTask(index int, status bool) {
	l.Tasks[index].Completed = status
}
