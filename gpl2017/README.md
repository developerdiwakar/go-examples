## This Section contains ```The Go Programming Language``` book exercises and examples


### Testing and Benchmarking
To run testing and benchmarking, save the code as a [filename]_test.go in the same directory and execute the following command in your terminal:
    
To run testing

```bash 
go test -v
```

t.Run enables running _"subtests"_, one for each table entry. These are shown separately when executing go test -v.

Running the benchmark:
```bash
go test -bench=.
```

To run both at once
```bash 
go test -v -bench=.
```