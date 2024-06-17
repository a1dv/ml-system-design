import csv
import os
import gensim
from gensim.models.phrases import Phrases, Phraser
from gensim.models import Word2Vec
from flask import Flask, request, jsonify

app = Flask(__name__)

# Загружаем модель Word2Vec
model = gensim.models.Word2Vec.load('w2v_bigrams.wv')

@app.route('/api/v1/similarity', methods=['GET'])
def similarity():
    word1 = request.args.get('word1')
    word2 = request.args.get('word2')
    if not word1 or not word2:
        return jsonify({'error': 'Missing word1 or word2 parameter'}), 400
    try:
        similarity = model.wv.similarity(word1, word2)
        similarity = float(similarity)
        return jsonify({'similarity': similarity})
    except KeyError:
        return jsonify({'error': 'Word not found in vocabulary'}), 400

@app.route('/api/v1/frequency', methods=['GET'])
def frequency():
    word = request.args.get('word')
    if not word:
        return jsonify({'error': 'Missing word parameter'}), 400
    try:
        frequency = model.wv.get_vecattr(word, 'count')
        frequency = int(frequency)
        return jsonify({'frequency': frequency})
    except KeyError:
        return jsonify({'error': 'Word not found in vocabulary'}), 400

@app.route('/api/v1/synonyms', methods=['GET'])
def synonyms():
    word = request.args.get('word')
    if not word:
        return jsonify({'error': 'Missing word parameter'}), 400
    try:
        synonyms = getSynonyms(word)
        return jsonify({'synonyms': synonyms})
    except KeyError:
        return jsonify({'error': 'Word not found in vocabulary'}), 400

@app.route('/api/v1/spellchecks', methods=['GET'])
def spellchecks():
    word = request.args.get('word')
    if not word:
        return jsonify({'error': 'Missing word parameter'}), 400
    try:
        spellchecks = getSpellchecks(word)
        return jsonify({'spellchecks': spellchecks})
    except KeyError:
        return jsonify({'error': 'Word not found in vocabulary'}), 400

if __name__ == '__main__':
    app.run(debug=True)
