package domain

import (
	"github.com/jetbasrawi/go.cqrs"
	"log"
)

type ToDoItemRepository interface {
	Load(string) (*ToDoItem, error)
	Save(ycq.AggregateRoot, *int) error
}

type ToDoItemCommandHandlers struct {
	repo ToDoItemRepository
}

func NewToDoItemCommandHandlers(repo ToDoItemRepository) *ToDoItemCommandHandlers {
	return &ToDoItemCommandHandlers{
		repo: repo,
	}
}

func (h *ToDoItemCommandHandlers) Handle(message ycq.CommandMessage) error {
	var todo *ToDoItem

	switch cmd := message.Command().(type) {

	// TODO: Doesn't check the number of uncompleted todo items
	case *AddToDoItemCommand:
		todo = NewToDoItem(message.AggregateID())
		if err := todo.Create(cmd.Content); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}

		return h.repo.Save(todo, ycq.Int(todo.OriginalVersion()))

	case *ToggleToDoItemCompletionCommand:
		todo, _ = h.repo.Load(message.AggregateID())
		todo.ToggleCompletion()

		return h.repo.Save(todo, ycq.Int(todo.OriginalVersion()))

	// TODO: Doesn't handle RemoveToDoItemCommand

	default:
		log.Fatalf("ToDoItemCommandHandlers has received a command that it is does not know how to handle, %#v", cmd)
	}

	return nil
}
