from flask import Flask
from flask_sqlalchemy import SQLAlchemy 
from config import Config

db = SQLAlchemy()

def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)
    db.init_app(app)

    from .book_routes import book_routes as book_blueprint
    app.register_blueprint(book_blueprint)

    with app.app_context():
        db.create_all()
        return app