# Strategy Design Patterns
**Strategy** is a behavioral design pattern that turns a set of behaviors into objects and makes them interchangeable inside original context object.

The original object, called context, holds a reference to a strategy object. The context delegates executing the behavior to the linked strategy object. In order to change the way the context performs its work, other objects may replace the currently linked strategy object with another one.


## Understanding the Problem
We aim to create a flexible payment system in Go that can handle multiple payment methods. The system should be extensible to accommodate new payment methods without significantly affecting existing code.

## Design Pattern: Strategy Pattern
The Strategy pattern is ideal for this scenario. It allows us to define a family of algorithms, encapsulate each one, and make them interchangeable.

## Explanation
1. _PaymentMethod interface:_ Defines the common behavior for all payment methods.
2. _PaymentMethod interface:_ Extended to include CalculateFee method for fee calculation.
3. _ProcessPayment function:_ Calculates the total amount including fees, processes the payment, and potentially adds fee information to the response.
4. _CreditCard, UpiCard, CreditCardPayment, UpiPayment structs:_ Concrete implementations of payment methods.
5. _PaymentContext struct:_ Holds the current payment method and provides a Pay method.
6. _NewPaymentContext function:_ Creates a new PaymentContext instance.


## Advantages of Strategy Pattern
- _Open-closed principle:_ New payment methods can be added without modifying existing code.
- _Flexibility:_ Different payment methods can be easily swapped.
- _Testability:_ Each payment method can be tested independently.


## Additional Considerations
- _Error handling:_ Implement proper error handling for payment failures.
- _Security:_ Ensure secure handling of payment information.
- _Dependency injection:_ Consider using a dependency injection framework for managing dependencies.
- _Payment gateway integration:_ Integrate with real-world payment gateways using their respective SDKs.
- Currency conversion: Handle different currencies if applicable.
- Payment status tracking: Track payment status and provide updates.
- _Transaction management:_ Implement transaction management to ensure data consistency.

_Source: refactoring.guru, gemini_ 