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
 - Go doesn't give you a way to ensure interface compliance for the interface implementor as Go embraces structural typing. However, it may not be a big deal considering the below two cases:
   - You probably use the interface implementation to assign to a variable or pass it as an argument to a method or function. This provides automatic compile-time checking
   - If you don't do the above (in cases where you rely on DI for example), you may have test on this to ensure it conforms the interface signature
   
   In other words, [in Go, you don’t declare that a type implements an interface. Instead, a type just implements the methods in the interface. This lack of ceremony makes interfaces feel very simple and informal](https://blog.carbonfive.com/2012/09/23/structural-typing-compile-time-duck-typing/).
 - In Go, idiomatic way to code defensively is to handle the defensive case 
   first so that all the happy path is written w/o indentation. For example,
 
   ```
	events, ok := repo.current[id]
	if !ok {
		return nil, &ycq.ErrAggregateNotFound{}
	}
	
	// do the rest here ...
   ```

   This is good since we check for the defensive case first and terminate the function.
