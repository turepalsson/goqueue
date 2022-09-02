# goqueue
Toying around with Go, trying to implement a heap queue.

## Making a Queue

A queue is simply a slice of whatever type you want.

```
myQueue := make([]int, 0)
```

Ordering is determined by a function that compares two elements, and returns `true` if the first is “less than” (goes before in the queue) the second one. This function is passed as an argument to the “add” and “pop” functions. Notice that you should never use different `less` functions with a single queue, or strange things will happen.

```
func less(x, y T) boolean
```

## Adding Elements

```
queue = Qadd(queue, newitem, less)
```

Notice that just like with `append`, you need to assign the return value back to your variable, since it is actually a slice, which *will* change!

## Removing elements

```
queue = Qpop(queue, less)
```



