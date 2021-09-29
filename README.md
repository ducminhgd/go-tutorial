# Go tutorial

A collection of Go tutorials for myself.

## Result

- 2021-09-29
    ```bash
    goos: linux
    goarch: amd64
    pkg: go-tutorial/cmd/cache
    cpu: Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz
    BenchmarkWriteMemcache-4              10         101207120 ns/op
    BenchmarkReadMemcache-4               10         101249020 ns/op
    BenchmarkWriteRedis-4               1803            646960 ns/op
    BenchmarkReadRedis-4                1872            599227 ns/op
    PASS
    ok      go-tutorial/cmd/cache   5.647s
    ```