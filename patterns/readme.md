# Design Patterns

The 23 Gang of Four patterns from the book “Design Patterns: Elements of Reusable Object-Oriented Software” (1994) are 
generally considered the foundation for all other patterns.

I've created this page as it's been a while since I've gone over these design patterns. Go is not an Object-Oriented
language, so I'm interested to see how these patterns can be implemented in Go as part of my Go learning.

There is a very good website on these patterns at https://refactoring.guru/design-patterns

## Object Creation Patterns

Deals with the creation of objects.

- Abstract Factory: Creates an instance of several families of classes.
- Builder: Separates object construction from its representation.
- Factory Method: Creates an instance of several derived classes.
- Prototype: A fully initialized instance to be copied or cloned.
- Singleton: A class of which only a single instance can exist.

## Structural Patterns

Deals with class structure such as Inheritance and Composition.

- Adapter: Match interfaces of different classes.
- Bridge: Separates an object’s interface from its implementation.
- Composite: A tree structure of simple and composite objects.
- Decorator: Add responsibilities to objects dynamically.
- Facade: A single class that represents an entire subsystem.
- Flyweight: A fine-grained instance used for efficient sharing.
- Proxy: An object representing another object.

## Behavioral Patterns

Patterns for better interaction between objects, how to provide lose coupling, and flexibility to extend easily in 
the future.

- Chain of Resp: A way of passing a request between a chain of objects.
- Command: Encapsulate a command request as an object.
- Interpreter: A way to include language elements in a program.
- Iterator: Sequentially access the elements of a collection.
- Mediator: Defines simplified communication between classes.
- Memento: Capture and restore an object's internal state.
- Observer: A way of notifying change to a number of classes.
- State: Alter an object's behavior when its state changes.
- Strategy: Encapsulates an algorithm inside a class.
- Template Method: Defer the exact steps of an algorithm to a subclass.
- Visitor: Defines a new operation to a class without change requirements.