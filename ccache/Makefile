all:
	go test -i -v
	go test -i -v -race
	go test -v
	go test -v -race

bench: all
	go test -v -bench=. -count=3 -cpu 1,2,4,8 | tee bench.txt
	benchstat bench.txt

bench-block: all
	go test -v -bench=. -blockprofile=block.pprof
	go tool pprof -svg -lines -output=block.svg ccache.test block.pprof
