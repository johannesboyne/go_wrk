#GO WRK

> Go (and do some) work

That's exactly what `go_wrk` does. It is a **HTTP benchmarking tool** that is capable of generating quite some amount of load because of Go's concurrency patterns (Subroutines).

Be sure you have enough `maxfiles` and `max sockets` etc. checkout: [http://b.oldhu.com/2012/07/19/increase-tcp-max-connections-on-mac-os-x/](http://b.oldhu.com/2012/07/19/increase-tcp-max-connections-on-mac-os-x/) for information to increase this on OS X or [http://stackoverflow.com/questions/410616/increasing-the-maximum-number-of-tcp-ip-connections-in-linux](http://stackoverflow.com/questions/410616/increasing-the-maximum-number-of-tcp-ip-connections-in-linux) for linux.
