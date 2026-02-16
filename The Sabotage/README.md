***Go Concurrency: ErrGroup Practice***
This project is part of a deep-dive into Go's concurrency patterns. It demonstrates how to use the golang.org/x/sync/errgroup package to manage multiple goroutines, handle errors, and synchronize tasks.

***Project Overview***
In this specific task (Project #2: The Sabotage), we simulate a real-world scenario where multiple workers perform tasks simultaneously. One worker is designed to fail to demonstrate how errgroup catches and propagates errors back to the main thread.

***Features***
Concurrent Execution: Tasks run in parallel using goroutines managed by errgroup.Group.
Error Propagation: If any goroutine returns an error, the Wait() method captures it.
Closure Safety: Demonstrates the correct way to capture loop variables (id := i) to avoid common concurrency bugs.
📋 Prerequisites
Before running the code, ensure you have Go installed and the errgroup package downloaded:

***Code Structure***
The main logic follows this flow:
Initialize an errgroup.Group.
Spawn 3 workers using the g.Go() method.
Worker 2 is programmed to return a fmt.Errorf after a short delay.
The main function calls g.Wait() to block until all workers finish or an error occurs.
Results are printed based on whether an error was returned.

***What I Learned***
The difference between sync.WaitGroup (waiting only) and errgroup.Group (waiting + error handling).
How to use fmt.Errorf to create dynamic error messages.
Why returning nil is necessary when a task is successful.
What's Next?