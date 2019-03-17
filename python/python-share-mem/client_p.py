import mmap
import contextlib
import time

while True:
    with contextlib.closing(mmap.mmap(-1, 1024, tagname="test", access=mmap.ACCESS_READ)) as m:
        s = m.read(1024).replace(b"\x00", b'')
        print(s)

    time.sleep(1)
