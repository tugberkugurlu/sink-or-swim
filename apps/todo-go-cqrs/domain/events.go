package domain

import "time"

type ToDoItemAddedEvent struct {
	ToDoItemID string
	Content string
	AddedAt time.Time
}

type ToDoItemRemovedEvent struct {
	ToDoItemID string
	RemovedAt time.Time
}

type ToDoItemCompletedEvent struct {
	ToDoItemID string
	CompletedAt time.Time
}

type ToDoItemCompletionRevertedEvent struct {
	ToDoItemID string
	RevertedAt time.Time
}
