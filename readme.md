# uqam simulation

Exploratory project to look into performance of some method to simulate
most common random variable in comparison with the ones provided by Go's
standard library.

As one can expect the conclusion is to use the functions provided by Go
and not to write your own.

### running the benchmarks
Just use Go's benchmarking capabilities with the bellow line of code.

    go test -bench .

### results

    didier@miaou ~/didiercrunch/uqamsimulation ±master » go test -bench .                                                         
    testing: warning: no tests to run
    PASS
    BenchmarkMersenneTwister-8      50000000                25.1 ns/op
    BenchmarkStdUniforme01-8        200000000                9.96 ns/op
    BenchmarkAESRandom-8            50000000                25.2 ns/op
    BenchmarkStdExpoVariate-8       100000000               14.8 ns/op
    BenchmarkExpoVariate-8          50000000                32.8 ns/op
    BenchmarkNormalBoxMuller-8      30000000                42.2 ns/op
    BenchmarkStdNormal-8            100000000               21.4 ns/op
    BenchmarkNormalMarsaglia-8      30000000                40.6 ns/op
    ok      github.com/didiercrunch/uqamsimulation  13.489s
