# concurrent_lift_simulation

A very basic concurrent lift simulation in order
to better understand goroutines and channels in Go.

This is a very basic simulation which fills a buffered
channel with randomised floor calls. These are then passed
on to a dispatcher which finds the nearest available lift
and sends a channel message to get it to respond.

This is meant to be an infinite system that continues to
run, that is why the channels are not closed at any point in the code.

![Lift Simulation](https://github.com/rhodriguerrier/concurrent_lift_simulation/blob/main/lift_concurrency_system_example.png?raw=true)