package domain

import "time"

type ToDoItemAddedEvent struct {
	ToDoItemID int64
	Content string
	AddedAt time.Time
}

type ToDoItemRemovedEvent struct {
	ToDoItemID int64
	RemovedAt time.Time
}

type ToDoItemCompletedEvent struct {
	ToDoItemID int64
	CompletedAt time.Time
}

type ToDoItemCompletionRevertedEvent struct {
	ToDoItemID int64
	RevertedAt time.Time
}
