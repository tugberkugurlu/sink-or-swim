package domain

import "github.com/jetbasrawi/go.cqrs"

type InMemoryRepo struct {
	current map[string][]ycq.EventMessage
	publisher ycq.EventBus
}

func NewInMemoryRepo(eventBus ycq.EventBus) *InMemoryRepo {
	return &InMemoryRepo{
		current: make(map[string][]ycq.EventMessage),
		publisher: eventBus,
	}
}

func (repo *InMemoryRepo) Load(id string) (*ToDoItem, error) {
	events, ok := repo.current[id]
	if !ok {
		return nil, &ycq.ErrAggregateNotFound{}
	}

	todoItem := NewToDoItem(id)

	for _, event := range events {
		todoItem.Apply(event, false)
		todoItem.IncrementVersion() // Q: why do we increment the version here?
	}

	return todoItem, nil
}

func (repo *InMemoryRepo) Save(aggregate ycq.AggregateRoot) error {
	//TODO: Look at the expected version

	_, ok := repo.current[aggregate.AggregateID()]
	if !ok {
		return &ycq.ErrAggregateNotFound{}
	}

	for _, event := range aggregate.GetChanges() {
		repo.current[aggregate.AggregateID()] = append(repo.current[aggregate.AggregateID()], event)
		repo.publisher.PublishEvent(event)
	}

	return nil
}
