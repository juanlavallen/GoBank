package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
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

func (t *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
