all:
	go test -i -v
	go test -i -v -race
	go test -v
	go test -v -race

bench: all
	go test -v -bench=. -count=5

bench-block: all
	go test -v -bench=. -blockprofile=block.pprof
	go tool pprof -svg -lines -output=block.svg ccache.test block.pprof
