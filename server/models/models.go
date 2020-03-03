package models

import "time"

type ToDo struct {
	Index     int       `json:"index"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	Created   time.Time `json:"created"`
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

	if index == 0 {
		for i := index; i < len(l.GetList()); i++ {
			arr[i].Index = i + 1
		}

		l.Tasks = arr
	} else {
		for i := index - 1; i < len(l.GetList()); i++ {
			arr[i].Index = i + 1
		}

		l.Tasks = arr
	}
}

/// Tasks ///
func (l *List) NewTask(text string) {
	var newIndex int

	if len(l.GetList()) == 0 {
		newIndex = 1
	} else {
		newIndex = l.GetList()[len(l.GetList())-1].Index + 1
	}

	var newToDo = ToDo{newIndex, text,
		false, time.Now()}
	l.Tasks = append(l.GetList(), newToDo)
}

func (l *List) DeleteTask(index int) {
	//if index == 0 {
	//	l.Tasks = append(l.Tasks[index:], l.Tasks[index+1:]...)
	//	l.updateList(index)
	//} else {
	l.Tasks = append(l.Tasks[:index], l.Tasks[index+1:]...)
	l.updateList(index)
	//}
}

func (l *List) ChangeTask(index int) {
	b := l.Tasks[index].Completed
	l.Tasks[index].Completed = !b
}
