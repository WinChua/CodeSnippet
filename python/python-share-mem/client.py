import mmap
import contextlib
import time

while True:
    with open("test.data", "r") as f:
        with contextlib.closing(mmap.mmap(f.fileno(), 1024, access=mmap.ACCESS_READ)) as m:
            s = m.read(1024).replace(b"\x00", b'')
            print(s)

    time.sleep(1)
