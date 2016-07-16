# i3gobar

A super simple package for making your own custom i3bar status line.


## Why does this exist?

Becasue when configuration files get too complicated sometimes it's easier to
just write your own program configured the way you want. That is the conclusion
I reached after trying to configure various bars just the way I like them.

This little library makes it super easy to make your own stats bar that does
exactly what you want it to.

## How do I use this thing?

See [main.go](example/main.go) for a complete example. It's that easy. Define
your array of what functions you want adding blocks to your status in the order
you want them, and define your own functions which simply return a block down
the provided channel. Each function runs asynchronously and handles it's own
timing so each block can refresh it's data as frequently as you want.

Some functions that I thought other people might like, like Load, CPU and Memory
I built into the library, so if you like the way they look, just use them. If
you don't, it's easy to build your own.
