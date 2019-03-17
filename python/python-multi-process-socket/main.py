import os
import socket

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
server.bind(("localhost", 12345))
server.listen(5)

for i in range(4):
    i = os.fork()
    if i != 0:
        break

pid = os.getpid()
while True:
    cli, addr = server.accept()
    data = cli.recv(4096)
    cli.send(str(pid).encode("ascii"))
    cli.close()
