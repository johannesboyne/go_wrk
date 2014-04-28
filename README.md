#GO WRK

> Go (and do some) work

That's exactly what `go_wrk` does. It is a **HTTP benchmarking tool** that is capable of generating quite some amount of load because of Go's concurrency patterns (Subroutines).

At the moment the maximum "concurrent" amount of requests seems to be `1010`. __NO__ `~200` atm because of body consuming. Will have to frame the requests.
