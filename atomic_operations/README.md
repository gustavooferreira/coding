# README

Test atomic operations performance.

Compares Atomic values with mutexes.

Run the following command:

```make
go test -v -bench . -run ^$ ./...
```

The value we are passing to the argument `-run` necessary in order to avoid running any existing tests, as they will run by default.

## GoLand

### Running a Benchmark

When running a benchmark in GoLand, these are the commands it ran:

```
/opt/homebrew/opt/go/libexec/bin/go test -c -o /private/var/folders/x4/lp0kwd7s1p17vq7kc4vxgkrm0000gp/T/GoLand/___BenchmarkIncrement_in_atomicops_ops.test atomicops/ops #gosetup
/private/var/folders/x4/lp0kwd7s1p17vq7kc4vxgkrm0000gp/T/GoLand/___BenchmarkIncrement_in_atomicops_ops.test -test.v -test.paniconexit0 -test.bench ^\QBenchmarkIncrement\E$ -test.run ^$
```

Essentially, these were the commands ran:

```
go test -c -o pkg.test atomicops/ops
./pkg.test -test.v -test.paniconexit0 -test.bench ^\QBenchmarkIncrement\E$ -test.run ^$
```

Without the double step creating the test binary, we could mimic the above by doing:

```
go test -v -bench "^\QBenchmarkIncrement\E$" -run "^$" atomicops/ops
```

The regex `\Q` and `\E` are respectively the start and end of a literal string in a regex literal; they instruct the
regex engine to not interpret the text in-between those two "markers" as regexes.

---

### Running a Test

When running a test in GoLand, these are the commands it ran:

```
/opt/homebrew/opt/go/libexec/bin/go test -c -o /private/var/folders/x4/lp0kwd7s1p17vq7kc4vxgkrm0000gp/T/GoLand/___TestIncrementOperation_in_atomicops_ops.test atomicops/ops #gosetup
/opt/homebrew/opt/go/libexec/bin/go tool test2json -t /private/var/folders/x4/lp0kwd7s1p17vq7kc4vxgkrm0000gp/T/GoLand/___TestIncrementOperation_in_atomicops_ops.test -test.v -test.paniconexit0 -test.run ^\QTestIncrementOperation\E$
```

Essentially, these were the commands ran:

```
go test -c -o pkg.test atomicops/ops
go tool test2json -t pkg.test -test.v -test.paniconexit0 -test.run ^\QTestIncrementOperation\E$
```

Without the double step creating the test binary, we could mimic the above by doing:

```
go test -v -json -run "^\QTestIncrementOperation\E$" atomicops/ops
```

## More information

Run the following commands to get more information:

```
go help test
go help testflag
go help packages
```
