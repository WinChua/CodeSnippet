PWD=$(shell pwd)



ministart:
	for p in 12340 12341;\
	do\
		python mini.py --port $$p &\
	done

ministop:
	kill `lsof -i :12340,12341 | tail -n 2 | cut -d   -f 3 `


gobuild:
	go build && ./$(shell basename $(PWD))

gorace:
	go run -race main.go
