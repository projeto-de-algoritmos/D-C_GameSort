from flask import Flask

from Routes.games_bp import games_bp

app = Flask(__name__)

app.register_blueprint(games_bp, url_prefix='/')

app.run(debug=True)
