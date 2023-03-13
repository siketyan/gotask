# ⏰ gotask

Easy asynchronous task runner for Go, inspired by Promise API of ECMAScript and the Task API of .NET.


## Prerequisites

- Go 1.20+


## Installation

```go
import "github.com/siketyan/gotask"
```


## Usages

### Running a Task synchronously

A task is the basically a closure, so you can run it by just calling `Do` receiver function.
Running synchronously means the run blocks the current thread.

```go
task := gotask.NewTask(func (ctx context.Context) error {
    time.Sleep(1 * time.Second)
	fmt.Println("Task has been finished!")

	return nil
})

err := task.Do(ctx)
require.NoError(t, err)
```

### Running a Task asynchronously

Using an asynchronous run, it never block the current thread and tasks can be run in parallel.
`DoAsync` takes a write channel to write the resolved value onto.

```go
task := gotask.NewTask(func (ctx context.Context) error {
    time.Sleep(1 * time.Second)
    fmt.Println("Task has been finished!")

    return nil
})

errChan := make(chan error)
task.DoAsync(ctx, errChan)
fmt.Println("Task has been started.")

err := <-errChan
require.NoError(t, err)
```

### Running tasks in parallel

Using `Parallel` function, you can run many tasks in parallel.
`Parallel` does **NOT** run any tasks, but only creates a new task from the child tasks.
You need to run the new task explicitly to run the child tasks and resolve the values.

Note that `Parallel` requires the child tasks to return a `Result` struct.
It holds both value and error in one struct, instead of writing `(T, error)` tuples.

> **Warning**
> The order of the tasks and the returned values are not the same.

```go
task := gotask.Parallel(
    gotask.NewTask(func(ctx context.Context) Result[string, error] {
        time.Sleep(1000 * time.Millisecond)

        return gotask.ResultOk[string, error]("The first task is resolved!")
    }),
    gotask.NewTask(func(ctx context.Context) Result[string, error] {
        time.Sleep(900 * time.Millisecond)

        return gotask.ResultOk[string, error]("The second task is resolved!")
    }),
)

values := task.Do(ctx).Unwrap()
assert.Equal(t, []string{
    "The second task is resolved!",
	"The first task is resolved!",
}, values)
```

In the example above, the parallel task will be resolved in 1000 ms.

When any child task returned an error, the run will be canceled immediately and returns the error.

```go
task := gotask.Parallel(
    gotask.NewTask(func(ctx context.Context) Result[string, error] {
        time.Sleep(900 * time.Millisecond)

        return gotask.ResultOk[string, error]("The first task is resolved!")
    }),
    gotask.NewTask(func(ctx context.Context) Result[string, error] {
        time.Sleep(500 * time.Millisecond)

        return gotask.ResultErr[string, error](errors.New("the second task occurred an error"))
    }),
)

err := task.Do(ctx).UnwrapErr()
assert.Error(t, err)
```

The parallel task will return an error after 500 ms, canceling the first child task.

### Running tasks in parallel (Settled)

`ParallelSettled` is basically the same as `Parallel`, but it never returns after any errors are occurred.
After run all the child tasks, it returns all the value resolved.

Also it is compatible with the tasks that does not return `Result`.

```go
task := gotask.ParallelSettled(
    gotask.NewTask(func(ctx context.Context) Result[string, error] {
        time.Sleep(900 * time.Millisecond)

        return gotask.ResultOk[string, error]("The first task is resolved!")
    }),
    gotask.NewTask(func(ctx context.Context) Result[string, error] {
        time.Sleep(500 * time.Millisecond)

        return gotask.ResultErr[string, error](errors.New("the second task occurred an error"))
    }),
)

results := task.Do(ctx)
assert.Error(t, results[0].UnwrapErr())
assert.Equal(t, "The first task is resolved!", results[1].Unwrap())
```

### Running task in parallel (Race)

`Race` creates a new task that runs the child task, and returns the first value resolved by any task.


```go
task := gotask.Race(
    gotask.NewTask(func(ctx context.Context) Result[string, error] {
        time.Sleep(1000 * time.Millisecond)

        return gotask.ResultOk[string, error]("The first task is resolved!")
    }),
    gotask.NewTask(func(ctx context.Context) Result[string, error] {
        time.Sleep(900 * time.Millisecond)

        return gotask.ResultOk[string, error]("The second task is resolved!")
    }),
)

value := task.Do(ctx).Unwrap()
assert.Equal(t, "The second task is resolved!", value)
```
