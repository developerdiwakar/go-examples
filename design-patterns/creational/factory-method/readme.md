# Factory Method Design Pattern
The Factory Method pattern is a creational pattern that provides an interface for creating objects in a superclass but allows subclasses to alter the type of objects that will be created. The main goal is to delegate the responsibility of object instantiation to subclasses.


## Example Scenario: Logger Factory
Imagine you need to create different types of loggers (like file-based or console-based) based on a configuration. The Factory Method pattern allows the instantiation of these loggers without exposing the creation logic.

In this example, the ```LoggerFactory``` interface defines the factory method ```CreateLogger()```. Each concrete factory (like ```ConsoleLoggerFactory``` and ```FileLoggerFactory```) implements this method to produce different loggers.