
## Go insert a value in a slice at a given index
```sh
go version
# go version go1.16.3 linux/amd64
go test -benchmem -bench . -args -n 32
# BenchmarkInsert-8    	 4260520	       290.3 ns/op	     512 B/op	       1 allocs/op
# BenchmarkInsert2-8   	 4643042	       299.4 ns/op	     512 B/op	       1 allocs/op
# BenchmarkInsert3-8   	 4161250	       298.8 ns/op	     512 B/op	       1 allocs/op
go test -benchmem -bench . -args -n 1000
# BenchmarkInsert-8    	  205935	      6107 ns/op	   16384 B/op	       1 allocs/op
# BenchmarkInsert2-8   	  197004	      6597 ns/op	   16384 B/op	       1 allocs/op
# BenchmarkInsert3-8   	  232641	      6705 ns/op	   16384 B/op	       1 allocs/op
go test -benchmem -bench . -args -n 1000000
# BenchmarkInsert-8    	     452	   2826832 ns/op	10002434 B/op	       1 allocs/op
# BenchmarkInsert2-8   	     410	   3312225 ns/op	10002435 B/op	       1 allocs/op
# BenchmarkInsert3-8   	     361	   3250836 ns/op	10002438 B/op	       1 allocs/op



make bench
go test -benchmem -bench . -args -n 32
# BenchmarkInsert-8    	 4125085  275.0 ns/op  512 B/op  1 allocs/op
# BenchmarkInsert2-8   	 3778551  316.0 ns/op  512 B/op  1 allocs/op

go test -benchmem -bench . -args -n 1000
# BenchmarkInsert-8    	  198364  5876 ns/op  16384 B/op  1 allocs/op
# BenchmarkInsert2-8   	  205197  7123 ns/op  16384 B/op  1 allocs/op

go test -benchmem -bench . -args -n 1000000
# BenchmarkInsert-8    	     643  1898436 ns/op	10002437 B/op  1 allocs/op
# BenchmarkInsert2-8   	     368  3248385 ns/op	10002436 B/op  1 allocs/op

```

1: My method (@wasmup): https://stackoverflow.com/a/61822301/8208215

2: @icza: https://stackoverflow.com/a/46130603/8208215

3: The Book (by Maximilien Andile): https://www.practical-go-lessons.com/chap-21-slices