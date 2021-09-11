from flask import render_template, redirect, url_for, request, abort

from Models.games import Game
from Controllers.mergeSort import mergeSort

def index():
    new_game = Game('LOL', 2013, 0)
    return f'Hello world! {new_game.name}'
