build:
	go build -o bin/resumemk

run: build
	./bin/resumemk

test:
	go test -v

clean:
	rm -f bin/resumemk

