# uuid-convertor
  
This is a convenience package used to convert UUIDs and or GUID to their numeric representations.
This package is useful for feature flagging systems that depend on numeric types

Note: If used with advertising ids, anonymous ids (00000000-0000-0000-0000-000000000000) will always return 0

## Benchmarks

goos: darwin
goarch: amd64
pkg: github.com/marcsantiago/uuid-converto

|BenchMark Name| Operations | ns/op | B/op| allocs/op|
|--|--|--|--|--|
|BenchmarkUUIDToIntString-8|1000000  |1003  |288  |7  |
|BenchmarkUUIDToUInt-8|2000000  |600  |144  |4  |