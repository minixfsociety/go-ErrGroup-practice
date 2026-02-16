***Go Concurrency: Project — The Domino Effect***
This project demonstrates the powerful synergy between errgroup and context in Go. It shows how to implement a "fail-fast" mechanism where one error automatically cancels all other running tasks.

***Concept: Fail-Fast with Context***
In real-world applications (like network scanners or scrapers), if one task fails critically, we often want to stop all other concurrent tasks to save resources. This project simulates that behavior.

***Key Features***
errgroup.WithContext: Initializes a group that manages a context.Context tied to the group's error state.
Automatic Cancellation: As soon as one worker returns a non-nil error, the context is canceled.
Graceful Shutdown: Workers use a select statement to "listen" for the cancellation signal (ctx.Done()) and stop their execution immediately.
***How It Works***
Worker 1 & 3: Designed to perform a long-running task (5 seconds).
Worker 2: Designed to fail after only 1 second.
The Magic: When Worker 2 fails, Workers 1 and 3 receive the ctx.Done() signal and exit instantly, rather than waiting for the full 5 seconds.
***Lessons Learned***
How errgroup acts as a "manager" that hits the "Big Red Button" (cancel) when an error occurs.
The importance of ctx.Err() to report why a goroutine was terminated.
Handling multiple flow paths in a goroutine using the select statement.