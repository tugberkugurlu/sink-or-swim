package domain

type AddToDoItemCommand struct {
	ID string
	Content string
}

type RemoveToDoItemCommand struct {
	ID string
	Version int
}

type ToggleToDoItemCompletionCommand struct {
	ID string
	Version int
}
