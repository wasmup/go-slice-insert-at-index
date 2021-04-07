
## Go insert a value in a slice at a given index
```sh
go version
# go version go1.16.3 linux/amd64
make bench
go test -benchmem -bench . -args -n 32
# BenchmarkInsert-8    	 4125085	       275.0 ns/op	     512 B/op	       1 allocs/op
# BenchmarkInsert2-8   	 3778551	       316.0 ns/op	     512 B/op	       1 allocs/op

go test -benchmem -bench . -args -n 1000
# BenchmarkInsert-8    	  198364	      5876 ns/op	   16384 B/op	       1 allocs/op
# BenchmarkInsert2-8   	  205197	      7123 ns/op	   16384 B/op	       1 allocs/op

go test -benchmem -bench . -args -n 1000000
# BenchmarkInsert-8    	     643	   1898436 ns/op	10002437 B/op	       1 allocs/op
# BenchmarkInsert2-8   	     368	   3248385 ns/op	10002436 B/op	       1 allocs/op

```
