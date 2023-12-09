obu:
	go build -o bin/obu cmd/main.go
	./bin/obu

receiver:
	go build -o bin/receiver -buildvcs=false ./obu_reciever
	./bin/receiver

calculator:
	go build -o bin/calculator -buildvcs=false ./distance_calculator
	./bin/calculator