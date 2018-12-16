# Todo GO CQRS

This is inspired by [TodoMVC](https://github.com/tastejs/todomvc). This will have a Todo application implemented using:

 - Go as the programming language
 - [CQRS (Command-query responsibility segregation)](http://www.cqrs.nu) and [Hexagonal Architecture](http://alistair.cockburn.us/Hexagonal+architecture) as the architectural principals (yes, [you can use both and they both address different decoupling issues](https://softwareengineering.stackexchange.com/a/361683/22417))

## System

### Events

 - ToDoItemAddedEvent
 - ToDoItemRemovedEvent
 - ToDoItemCompletedEvent
 - ToDoItemCompletionRevertedEvent

### Commands

 - AddToDoItemCommand (emits `ToDoItemAddedEvent` event)
 - RemoveToDoItemCommand (emits `ToDoItemRemovedEvent` event)
 - ToggleToDoItemCompletionCommand (emits either `ToDoItemCompletedEvent` or `ToDoItemCompletionRevertedEvent` event)

### Queries

 - AllToDoItemsQuery
 - ActiveToDoItemsQuery
 - CompletedToDoItemsQuery

### Sagas

 - CompletedItemCleanerSaga: Sends `RemoveToDoItemCommand` for each completed item found by `CompletedToDoItemsQuery`. This is implemented as a saga since this bulk operation goes across consistency boundaries. For more information, see ["How can I update a set of aggregates with a single command?"](http://www.cqrs.nu/faq) section in CQRS FAQ provided by Edument.

## Side Learning

 - Go type embedding: https://travix.io/type-embedding-in-go-ba40dd4264df
 - Type switches: https://yourbasic.org/golang/type-assertion-switch/
 - How do I do a literal *int64 in Go?: https://stackoverflow.com/questions/30716354/how-do-i-do-a-literal-int64-in-go
