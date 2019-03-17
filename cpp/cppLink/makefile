LIB=-Llib -lhello
INC=-Iinclude
CC=gcc


all: lib/libhello.a lib/libhello.so static.out dynamic.out

obj/hello.o: src/lib/hello.c
	$(CC) $(INC) -c $< -o $@

lib/libhello.a: obj/hello.o
	ar crs lib/libhello.a obj/hello.o 

lib/libhello.so: src/lib/hello.c
	$(CC) $(INC) -shared -fPIC $< -o $@

static.out: src/main.c lib/libhello.a
	$(CC) $(INC) -static src/main.c $(LIB) -o $@

dynamic.out: src/main.c lib/libhello.so
	$(CC) $(INC) src/main.c $(LIB) -o $@ -Wl,-rpath lib
