mycp: cp.cpp
	gcc cp.cpp -omycp

mywho: who.cpp
	gcc who.cpp -omywho

mywho2: who2.cpp
	gcc who2.cpp -omywho2

myls: myls.cpp
	gcc myls.cpp -omyls

mycat: mycat.cpp
	gcc mycat.cpp -omycat

all: mycp mywho mywho2 myls mycat

test:
	./mywho > mywho.txt
	./mywho2 > mywho2.txt
	./mycp mywho.txt mycp.txt
	./mycat mycp.txt

.PHONY: all test
