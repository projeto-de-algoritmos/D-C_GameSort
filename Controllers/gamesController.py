from flask import render_template, redirect, url_for, request, abort
import pandas as pd

from Models.games import Game
from Controllers.mergeSort import mergeSort

def load_data():
    df = pd.read_csv('vgsales.csv')[['Name', 'Year', 'Global_Sales']]

    alist = []

    for index, row in df.iterrows():
        alist.append(Game(row['Name'], row['Year'], row['Global_Sales']))
    return alist

def index():
    alist = load_data()
    mergeSort(alist,True,'name')
    return str({'name': alist[0].getName(), 'year': alist[0].getYear(), 'global_sales': alist[0].getSales() })
