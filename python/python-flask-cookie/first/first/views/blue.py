from . import app

from flask import Blueprint

blue = Blueprint("simpel_page", __name__)

@blue.route("/", methods=["GET", "POST"])
def index():
    return "Hello, from blue"
