all:
	go test -i -v
	go test -i -v -race
	go test -v
	go test -v -race

bench: all
	go test -v -bench=. -count=3 -cpu 1 | tee bench.txt
	benchstat bench.txt

bench-cpu: all
	go test -v -bench=/^Slice,20 -cpuprofile=Slice-20-cpu.pprof
	go test -v -bench=/^Map,20 -cpuprofile=Map-20-cpu.pprof

	go tool pprof -svg -unit=ms -output=Slice-20-cpu.svg cache.test Slice-20-cpu.pprof
	go tool pprof -svg -unit=ms -output=Map-20-cpu.svg cache.test Map-20-cpu.pprof
