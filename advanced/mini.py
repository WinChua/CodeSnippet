import click
import json
from flask import Flask

app = Flask(__name__)

@app.route("/getjson")
def getjson():
    data = {"title": "Hello",
            "next": 5000*1000*1000}
    return json.dumps(data)


@click.command()
@click.option("--host", default="0.0.0.0", help="host")
@click.option("--port", default=12345, help="port")
def run(host, port):
    app.run(host, port)


if __name__ == "__main__":
    run()
