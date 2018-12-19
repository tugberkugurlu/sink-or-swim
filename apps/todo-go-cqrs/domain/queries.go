package domain

import "github.com/jetbasrawi/go.cqrs"

var inMemoryDatabase *InMemoryDatabase

type Queries interface {
	GetAllToDoItems() []*ToDoItemReadModel
	GetActiveToDoItems() []*ToDoItemReadModel
	GetCompletedToDoItems() []*ToDoItemReadModel
}

type ToDoItemReadModel struct {
	ID string
	Content string
	Completed bool
	Version int
}

type ReadModel struct {
}

func NewReadModel() *ReadModel {
	initializeInMemoryDb()
	return &ReadModel{}
}

func (m *ReadModel) GetAllToDoItems() []ToDoItemReadModel {
	var allTodos []ToDoItemReadModel
	for _, todo := range inMemoryDatabase.OrderedAll {
		allTodos = append(allTodos, *todo)
	}

	return allTodos
}

func (m *ReadModel) GetActiveToDoItems() []ToDoItemReadModel {
	var allTodos []ToDoItemReadModel
	for _, todo := range inMemoryDatabase.OrderedAll {
		if !todo.Completed {
			allTodos = append(allTodos, *todo)
		}
	}

	return allTodos
}

func (m *ReadModel) GetCompletedToDoItems() []ToDoItemReadModel {
	var allTodos []ToDoItemReadModel
	for _, todo := range inMemoryDatabase.OrderedAll {
		if todo.Completed {
			allTodos = append(allTodos, *todo)
		}
	}

	return allTodos
}

type ToDoItemsListView struct {
}

func NewToDoItemsListView() *ToDoItemsListView {
	initializeInMemoryDb()
	return &ToDoItemsListView{}
}

func (view *ToDoItemsListView) Handle(message ycq.EventMessage) {
	switch event := message.Event().(type) {

	case *ToDoItemAddedEvent:
		todo := ToDoItemReadModel{
			ID: message.AggregateID(),
			Content: event.Content,
			Completed: false,
			Version: *message.Version(),
		}

		inMemoryDatabase.OrderedAll = append(inMemoryDatabase.OrderedAll, &todo)
		inMemoryDatabase.All[message.AggregateID()] = &todo

	case *ToDoItemCompletedEvent:
		todo, ok := inMemoryDatabase.All[message.AggregateID()]
		if ok {
			todo.Completed = true
		}

	case *ToDoItemCompletionRevertedEvent:
		todo, ok := inMemoryDatabase.All[message.AggregateID()]
		if ok {
			todo.Completed = false
		}

	case *ToDoItemRemovedEvent:
		// The delete function doesn't return anything, and will do nothing if the specified key doesn't exist.
		delete(inMemoryDatabase.All, message.AggregateID())

		// Could use doublylinkedlist here instead and persist the tree node inside the map too,
		// which would make this operation also O(1)
		index := 0
		for i, todo := range inMemoryDatabase.OrderedAll {
			if todo.ID == message.AggregateID() {
				index = i
				break
			}
		}

		// https://stackoverflow.com/a/25025536/463785
		inMemoryDatabase.OrderedAll = append(inMemoryDatabase.OrderedAll[:index], inMemoryDatabase.OrderedAll[index+1:]...)
	}
}

func initializeInMemoryDb() {
	if inMemoryDatabase == nil {
		inMemoryDatabase = NewInMemoryDatabase()
	}
}

type InMemoryDatabase struct {
	All map[string]*ToDoItemReadModel
	OrderedAll []*ToDoItemReadModel
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		All: make(map[string]*ToDoItemReadModel),
	}
}
