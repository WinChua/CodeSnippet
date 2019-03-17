from flask import render_template, make_response, request
from . import app

@app.route("/", methods=["GET", "POST"])
def index():
    print(request.cookies)
    resp = make_response(render_template("index.html"))
    resp.set_cookie("UserID", "WinChua")
    return resp
