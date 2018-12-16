package domain

import (
	"errors"
	"github.com/jetbasrawi/go.cqrs"
	"time"
)

type ToDoItem struct {
	*ycq.AggregateBase
	content string
	completed bool
}

func NewToDoItem(id string) *ToDoItem {
	todo := &ToDoItem{
		AggregateBase: ycq.NewAggregateBase(id),
	}

	return todo
}

func (todo *ToDoItem) Create(content string) error {
	if content == "" {
		return errors.New("the content cannot be empty")
	}

	todo.Apply(ycq.NewEventMessage(todo.AggregateID(),
		&ToDoItemAddedEvent{ToDoItemID: todo.AggregateID(),
			Content: content,
			AddedAt: time.Now()},
			ycq.Int(todo.CurrentVersion())), true)

	return nil
}

func (todo *ToDoItem) ToggleCompletion() {
	var event interface{}
	if !todo.completed {
		event = ToDoItemCompletedEvent{ToDoItemID: todo.AggregateID(), CompletedAt: time.Now()}
	} else {
		event = ToDoItemCompletionRevertedEvent{ToDoItemID: todo.AggregateID(), RevertedAt: time.Now()}
	}

	todo.Apply(ycq.NewEventMessage(todo.AggregateID(), &event, ycq.Int(todo.CurrentVersion())), true)
}

func (todo *ToDoItem) Apply(message ycq.EventMessage, isNew bool) {
	if isNew {
		todo.TrackChange(message)
	}

	switch ev := message.Event().(type) {

	case *ToDoItemAddedEvent:
		todo.content = ev.Content

	case *ToDoItemCompletedEvent:
		todo.completed = true

	case *ToDoItemCompletionRevertedEvent:
		todo.completed = false

	}
}
