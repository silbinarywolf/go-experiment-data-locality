# Golang Experiment with Cache Locality

This is a set of benchmarks aimed at allowing me to see play with how data locality can affect performance. I may have gotten things wrong or not accounted for Golang optimizations, so do not assume this benchmark to be accurate until you confirm the results for yourself!

No idea why `BenchmarkEntitiesSmallListNoAlloc10` is coming up slower than the version with allocation right now. Needs more investigation.

```
const entitySize = 1024

$ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/silbinarywolf/go-experiment-data-locality
BenchmarkSlowCountEntitiesOnFireQuery10-4          30000             50733 ns/op
BenchmarkCountEntitiesOnFire10-4                 2000000               637 ns/op
BenchmarkEntitiesSlowSmallList10-4                300000              3499 ns/op
BenchmarkEntitiesSmallList10-4                   1000000              1028 ns/op
BenchmarkEntitiesSmallListNoAlloc10-4            1000000              1267 ns/op
BenchmarkEntitiesListAllOnFire10-4                500000              3893 ns/op
BenchmarkEntitiesListAllOnFireNoAlloc10-4        2000000               794 ns/op
PASS
ok      github.com/silbinarywolf/go-experiment-data-locality    12.059s

$ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/silbinarywolf/go-experiment-data-locality
BenchmarkSlowCountEntitiesOnFireQuery10-4          30000             50733 ns/op
BenchmarkCountEntitiesOnFire10-4                 2000000               637 ns/op
BenchmarkEntitiesSlowSmallList10-4                300000              3499 ns/op
BenchmarkEntitiesSmallList10-4                   1000000              1028 ns/op
BenchmarkEntitiesSmallListNoAlloc10-4            1000000              1267 ns/op
BenchmarkEntitiesListAllOnFire10-4                500000              3893 ns/op
BenchmarkEntitiesListAllOnFireNoAlloc10-4        2000000               794 ns/op
PASS
ok      github.com/silbinarywolf/go-experiment-data-locality    12.059s

```