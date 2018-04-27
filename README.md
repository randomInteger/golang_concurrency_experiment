# golang_concurrency_experiment
A very quick example of how to launch multiple independent goroutines and wait for them to finish

goroutine:  a lightweight thread managed by the Go runtime.

# install golang
https://golang.org/dl/

# build your go app
$go build

# run it
$./golang_concurrency_experiment <num_goroutines> <max range from 0 to compute>

# example of runtime on a MBP15/i7 computing all squares up to 32,000 on 16 threads
$time ./golang_concurrency_experiment 16 32000\
...\
goroutine: 12  square of: 31994 is: 1023616036\
goroutine: 12  square of: 31995 is: 1023680025\
goroutine: 12  square of: 31996 is: 1023744016\
goroutine: 12  square of: 31997 is: 1023808009\
goroutine: 12  square of: 31998 is: 1023872004\
goroutine: 12  square of: 31999 is: 1023936001\
\
real	0m2.658s\
user	0m1.413s\
sys	0m1.618s

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
