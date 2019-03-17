import requests
import time

POP20_CC = 'CN IN US ID BR PK NG BD RU JP MX PH VN ET EG DE IR TR CD FR'.split(" ")
BASE_URL = 'http://flupy.org/data/flags'

DEST_DIR = "downloads/"

def getflag(cc):
    url = "{}/{cc}/{cc}.gif".format(BASE_URL, cc=cc)
    resp = requests.get(url)
    with open("downloads/{cc}.gif".format(cc=cc), "wb") as f:
        f.write(resp.content)

def download():
    for cc in POP20_CC:
        getflag(cc)


def main():
    start = time.time()
    download()
    end = time.time()
    print("COST:", end - start, "s")

if __name__ == "__main__":
    main()

