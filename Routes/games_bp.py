from flask import Blueprint

from Controllers.gamesController import index

games_bp = Blueprint('user_bp', __name__)

games_bp.route('/', methods=['GET'])(index)
