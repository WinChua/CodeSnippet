import socket
import time
while True:
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.connect(("localhost", 12345))
    sock.send(b"hello")
    data = sock.recv(4096)
    print(data)
    sock.close()
    time.sleep(1)

