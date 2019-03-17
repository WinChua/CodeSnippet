import requests
import time
from concurrent.futures import ThreadPoolExecutor
from seq import getflag, POP20_CC

with ThreadPoolExecutor(len(POP20_CC)) as pool:
    start = time.time()
    pool.map(getflag, POP20_CC)
    end = time.time()
    print("COST:", end - start,"S")
