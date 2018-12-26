CC=gcc
SRC=$(wildcard *.c)
OBJ=$(patsubst %.c, %, $(SRC))

## $(OBJ): $(SRC)
## 	$(CC) -o $@ $< 
all: $(OBJ)

## this is implicit rule which couldn't be the default rule
%: %.c  
	$(CC) -pthread -o $@ $<
clean:
	rm $(OBJ)

test: $(OBJ)
	@echo $<
	@./$<


.PHONY: clean test
