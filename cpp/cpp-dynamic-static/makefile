all: obj/hello.o lib/libhello.a lib/libhello.so main_static main_dynamic

obj/hello.o: src/hello.c
	gcc -I include -o obj/hello.o src/hello.c -c

lib/libhello.a: obj/hello.o
	ar rcs lib/libhello.a obj/hello.o

lib/libhello.so: src/hello.c
	gcc -fPIC -shared src/hello.c -I include src/hello.c

main_static: src/main.c
	gcc src/main.c -static -I include -Llib -lhello -omain_static

main_dynamic: src/main.c
	gcc src/main.c -I include -Llib -lhello -Wl,-rpath=`pwd`/lib -omain_dynamic

dlopen_example: dlopen_example.c
	gcc dlopen_example.c -o dlopen_example -Wl,-rpath=`pwd`/lib -ldl
