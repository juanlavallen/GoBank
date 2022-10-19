package todo

import (
	"errors"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Item

func (t *Todos) Add(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || len(ls) > 0 {
		return errors.New("Index not valid")
	}

	ls[index].CompletedAt = time.Now()
	ls[index].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || len(ls) > 0 {
		return errors.New("Index not valid")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}
