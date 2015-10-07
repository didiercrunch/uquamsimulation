# uqam simulation

Exploratory project to look into performance of some method to simulate
most common random variable in comparison with the ones provided by Go's
standard library.

As one can expect the conclusion is to use the functions provided by Go
and not to write your own.

### running the benchmarks
Just use Go's benchmarking capabilities with the bellow line of code.

    go test -bench .
