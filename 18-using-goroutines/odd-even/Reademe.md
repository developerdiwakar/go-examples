# Find all Odd and Even numbers between 1 to `N` numbers

## Proposed Solution

1. **Separate Distribution:**

    Use a distributor goroutine that reads numbers from the `for` loop and sends them to separate `oddChan` and `evenChan` channels based on their parity.
    This ensures that each processing goroutine receives exactly the numbers it needs to process.

2. **Processing Goroutines:**

    Two separate goroutines (`processOdds` and `processEvens`) will read from their respective channels and collect numbers into slices.

3. **Synchronization with WaitGroup:**

    A `sync.WaitGroup` will ensure that the main function waits for all processing to complete before printing the results.

4. **Proper Channel Closure:**

    All channels will be properly closed to signal completion and avoid deadlocks.

5. **Simplified and Clear Structure:**

    The code will be straightforward, making it easier to understand and maintain.
