import mmap
import contextlib
import time

with open("test.data", "w") as f:
    f.write("\x00" * 1024)

with open("test.data", "r+") as f:
    with contextlib.closing(mmap.mmap(f.fileno(), 1024, access=mmap.ACCESS_WRITE)) as m:
        for i in range(1, 10001):
            m.seek(0)
            s = b"msg " + str(i).encode("utf8")
            s.rjust(1024, b'\x00')
            m.write(s)
            m.flush()
            time.sleep(1)
