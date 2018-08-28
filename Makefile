build:
	go build .

setup:
	./setup.sh

run:
	make build
	./icecream

clean:
	rm icecream

