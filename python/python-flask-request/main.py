from gevent.wsgi import WSGIServer
from flask import Flask, request
import logging

app = Flask(__name__)
app.logger.setLevel(logging.ERROR)

@app.route("/")
def index():
    return "Hello, World\n"


if __name__ == "__main__":
    httpserver = WSGIServer(("", 5000), app, log=None)
    httpserver.serve_forever()
