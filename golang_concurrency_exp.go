/*
Trying to learn goroutines....

Observations:

1)  Use a sync.WaitGroup if you need your main routine to wait for all
goroutines to finish, as they are truly independent and the main routine
can easily exit before the goroutines have finished doing their work.

2)  Pass in a pointer to the WaitGroup to the func that will be called
as a goroutine

3)  Observe that the threads are subject to the whims of a scheduler and the
output is not sequential between threads.

Author:  c.gleeson 2018
*/
package main
import "os"
import "fmt"
import "log"
import "sync"
import "strconv"

//A simple goroutine that computes squares and prints them
//Calls:  wg.Done() when finished so the main routine can resume.
//Takes:  N is the range(0, N), and thread int is the thread number for example
//purposes.
func squares(n int, thread_num int, wg *sync.WaitGroup) {
    for i := 0; i < n; i ++ {
        sq := i * i
        fmt.Println("Thread:", thread_num, " square of:", i , "is:", sq)
    }
    defer wg.Done()
}

func main(){
    //grab args 0=program path 1=first arg 2=2nd arg...
    args := os.Args[:]
    //Expect the following usage
    //Arg1:  number of threads
    //Arg2:  N where we compute all squares in range(0,N)
    //Define the max number of threads
    NUM_THREADS, err := strconv.Atoi(args[1])
    if err != nil {
        log.Fatal(err)
    }
    //Define the range of values to square from 0 to N-1
    RANGE, err := strconv.Atoi(args[2])
    if err != nil {
        log.Fatal(err)
    }
    //Form a waitgroup
    var wg sync.WaitGroup

    for i := 0; i < NUM_THREADS; i++ {
        wg.Add(1) //Populate the waitgroup with the correct num of goroutines
        go squares(RANGE, i, &wg)
    }
    wg.Wait()
}
