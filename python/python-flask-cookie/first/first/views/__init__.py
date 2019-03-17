from first import app
from .index import *
from .blue import blue


app.register_blueprint(blue, url_prefix="/blue")
