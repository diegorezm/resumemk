build:
	go build -o bin/resumemk

run: build
	./bin/resumemk

test:
	./bin/resumemk -f exemple.md -o resume -type 'pdf' 

clean:
	rm -f bin/resumemk

