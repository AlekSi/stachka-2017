all:
	go test -i -v
	go test -i -v -race
	go test -v
	go test -v -race

bench: all
	go test -v -bench=. -count=5
