OBJ=$(patsubst %.c, %, $(wildcard *.c))

all : $(OBJ)


% : %.c
	gcc $< -o $@

clean:
	rm ${OBJ}
