BINARY_NAME=angus

build:
	go build -o $(BINARY_NAME)

run:
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

all: build