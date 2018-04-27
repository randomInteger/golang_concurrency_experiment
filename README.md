# golang_concurrency_experiment
A very quick example of how to launch multiple independent goroutines and wait for them to finish

# build it
$go build

# run it
$./golang_concurrency_experiment <num_goroutines> <max range from 0 to compute>

# example output with very small values...
$./golang_concurrency_experiment 6 6\
goroutine: 5  square of: 0 is: 0\
goroutine: 5  square of: 1 is: 1\
goroutine: 5  square of: 2 is: 4\
goroutine: 5  square of: 3 is: 9\
goroutine: 5  square of: 4 is: 16\
goroutine: 5  square of: 5 is: 25\
goroutine: 4  square of: 0 is: 0\
goroutine: 4  square of: 1 is: 1\
goroutine: 4  square of: 2 is: 4\
goroutine: 4  square of: 3 is: 9\
goroutine: 4  square of: 4 is: 16\
goroutine: 4  square of: 5 is: 25\
goroutine: 1  square of: 0 is: 0\
goroutine: 1  square of: 1 is: 1\
goroutine: 1  square of: 2 is: 4\
goroutine: 1  square of: 3 is: 9\
goroutine: 1  square of: 4 is: 16\
goroutine: 1  square of: 5 is: 25\
goroutine: 3  square of: 0 is: 0\
goroutine: 3  square of: 1 is: 1\
goroutine: 3  square of: 2 is: 4\
goroutine: 3  square of: 3 is: 9\
goroutine: 2  square of: 0 is: 0\
goroutine: 2  square of: 1 is: 1\
goroutine: 2  square of: 2 is: 4\
goroutine: 2  square of: 3 is: 9\
goroutine: 2  square of: 4 is: 16\
goroutine: 2  square of: 5 is: 25\
goroutine: 0  square of: 0 is: 0\
goroutine: 3  square of: 4 is: 16\
goroutine: 0  square of: 1 is: 1\
goroutine: 0  square of: 2 is: 4\
goroutine: 0  square of: 3 is: 9\
goroutine: 0  square of: 4 is: 16\
goroutine: 0  square of: 5 is: 25\
goroutine: 3  square of: 5 is: 25
