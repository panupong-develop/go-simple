https://www.youtube.com/watch?v=5Z8skvm4g64

import sync


# Wait group
var wg sync.WaitGroup

1. read through the program
2. if we see goroutine syntax
    go xxx --> spawn a new thread to store the state for cycle loop (non blocking io)
        wg.Add(1)               +1
        wg.Add(1)               +1
3. execute as sequencetial + wait the routine are done all
            defer wg.Done()     -1
            defer wg.Done()     -1

    wg.Wait() == 0?          // block io in the main thread;
                             // I won't let you go if my children haven't come


# Channels
is an iterable shared memory pipeline storage that we can pass thing across the thread

c := make(chan in, size)

go func(c chan<- any) { // lambda
    c <- passValue
}(c)

outputValue := <-c          // blockio in the main thread

Caution!
    if we didn't sent data to fulfill the channel size, it will block the main thread

# Mutex
mutual exclusive = I only allowed 1 person to access me at one time (Thread safety)

object := Object{data: make(map[string]int)}

const personNumber := 100
for i := 0; i < personNumber; i++ {
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        object.RandomReadWriteWith(i)       // Race condition here, it's too crownded
    }
}
wg.Wait()

// solution
type Object struct {
    mux     sync.Mutex
    data    map[string]int
}

func (o *Object) RandomReadWriteWith(num int) {
    o.mux.Lock()                            // Block other people
    defer o.mux.Unlock()                    // Unblock when I finished
    o.data[randomIndex] = num
}


# Worker Pool
a manager to handle set of goroutine to manage with the resource that we have
so that it won't let too much children go buy snacks more than our budgets
and also manage QUEUE of result that will be ordered accordingly to the jobs.

Manager is manage(
            (1) 'Jobs' through channel until it fulfillment,
            (2) 'Results' through channel until it fulfillment,
            (3) by specify 'NumWorker'
        )


// Define Job channels and create jobs
const NUM_JOBS = 3
jobsQueue := make(chan number, NUM_JOBS)

jobsQueue <- MyJob1
jobsQueue <- MyJob2
jobsQueue <- MyJob3

// Define action
                // get jobs  -> doing -> throw the result
function doJob(jobsQueue <-chan struct, resultsCh chan<- struct) {
    for jobStruct := range jobsQueue {
        resultsCh <- something(jobStruct)
    }
}

// Worker do job
const NUM_WORKERS = 2
for worker := 0; worker < NUM_WORKERS; worker++ {
    go doJob(jobsQueue, resultsCh)
}

// Get the result
for result := range resultsCh {
    fmt.Println("Received:", result)
}