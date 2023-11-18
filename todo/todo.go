package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   string
	CompletedAt string
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now().Format(time.RFC822),
		CompletedAt: "",
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(idx int) error {
	ls := *t
	if idx <= 0 || idx > len(ls) {
		return errors.New("invalid index")
	}

	ls[idx-1].CompletedAt = time.Now().Format(time.RFC822)
	ls[idx-1].Done = true

	return nil
}

func (t *Todos) Delete(idx int) error {
	ls := *t
	if idx <= 0 || idx > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:idx-1], ls[idx:]...)

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

func (t *Todos) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "N"},
			{Align: simpletable.AlignCenter, Text: "Задание"},
			{Align: simpletable.AlignCenter, Text: "Готово?"},
			{Align: simpletable.AlignRight, Text: "Начало"},
			{Align: simpletable.AlignRight, Text: "Окончание"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("нет")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("да")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt},
			{Text: item.CompletedAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() (total int) {
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}
	return
}
