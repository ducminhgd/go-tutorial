# Go tutorial

A collection of Go tutorials for myself.

## Result

- 2021-09-30: I used WSL2 to run.

    ```bash
    → go1.17.1 test -bench=.
    goos: linux
    goarch: amd64
    pkg: go-tutorial/cmd/cache
    cpu: Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz
    BenchmarkWriteMemcache-4                      10         100472420 ns/op
    BenchmarkReadMemcache-4                       10         100689330 ns/op
    BenchmarkWriteMemcacheGoroutine-4         652153              4958 ns/op
    BenchmarkReadMemcacheGoroutine-4         1000000              3686 ns/op
    BenchmarkWriteRedis-4                       1652            683277 ns/op
    BenchmarkReadRedis-4                        1827            746675 ns/op
    BenchmarkWriteKeyDb-4                       1678            716956 ns/op
    BenchmarkReadKeyDb-4                        1707            702617 ns/op
    BenchmarkWriteGoCache-4                   548565              3070 ns/op
    BenchmarkReadGoCache-4                  15934350                71.62 ns/op
    BenchmarkWriteGoCacheGoroutine-4         1000000              4790 ns/op
    BenchmarkReadGoCacheGoRoutine-4          3935263               313.0 ns/op
    BenchmarkWriteRistretto-4                    732           1880123 ns/op
    BenchmarkReadRistretto-4                     100          18708744 ns/op
    BenchmarkWriteBigcache-4                       1        4617171600 ns/op
    BenchmarkReadBigcache-4                 11408173               104.6 ns/op
    PASS
    ok      go-tutorial/cmd/cache   50.058s
    ```

    ```bash
    goos: linux
    goarch: amd64
    pkg: go-tutorial/cmd/litedb
    cpu: Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz
    BenchmarkWriteSQLite-4                               151           7949415 ns/op
    BenchmarkReadSQLite-4                               8212            137751 ns/op
    BenchmarkWriteShareMemSQLite-4                       135           7731625 ns/op
    BenchmarkReadShareMemSQLite-4                       8472            137281 ns/op
    BenchmarkWriteShareMemSQLiteGoroutine-4           133040              8507 ns/op
    BenchmarkReadShareMemSQLiteGoroutine-4           1000000              8075 ns/op
    BenchmarkWriteMemSQLite-4                         141237              8991 ns/op
    BenchmarkReadMemSQLite-4                            3776            593021 ns/op
    BenchmarkWriteRQLite-4                               151           6873940 ns/op
    BenchmarkReadRQLite-4                                193           6421101 ns/op
    BenchmarkWriteRQLiteGoroutine-4                   774082            131756 ns/op
    BenchmarkReadRQLiteGoroutine-4                    799147             58185 ns/op
    PASS
    ok      go-tutorial/cmd/litedb  17.752s
    ```

## Known issues

1. Some cache tests cannot run with Goroutine:
    - Connection timeout:Redis, KeyDB
    - Run too long: Ristretto, Bigcache.