package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(idx int) error {
	ls := *t
	if idx <= 0 || idx > len(ls) {
		return errors.New("invalid index")
	}

	ls[idx-1].CompletedAt = time.Now()
	ls[idx-1].Done = true

	return nil
}

func (t *Todos) Delete(idx int) error {
	ls := *t
	if idx <= 0 || idx > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[idx-1:], ls[idx:]...)

	return nil
}

func (t *Todos) Load(fname string) error {
	file, err := os.ReadFile(fname)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	return json.Unmarshal(file, t)
}

func (t *Todos) Store(fname string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(fname, data, 0644)
}
