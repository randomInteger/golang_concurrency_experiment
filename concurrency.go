/*
Trying to learn threads....

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
import "fmt"
import "sync"

//A simple function that computes squares and prints them
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
    //Define the max number of threads
    NUM_THREADS := 64
    //Define the range of values to square from 0 to N-1
    RANGE := 64
    //Form a waitgroup
    var wg sync.WaitGroup
    wg.Add(NUM_THREADS)

    fmt.Println("start")
    for i := 0; i < NUM_THREADS; i++ {
        go squares(RANGE, i, &wg)
    }
    wg.Wait()
}
