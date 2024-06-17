import csv
import os
import gensim
from gensim.models.phrases import Phrases, Phraser
from gensim.models import Word2Vec

new_sentences = []
modelName="w2v_bigrams.wv"

# Чтение файла и преобразование первой колонки в список предложений
with open("mqData6m.csv", "r", encoding="utf-8") as file:
    # Читаем строки из файла
    for line in file:
        # Разбиваем строку на слова
        words = line.lower().split()
        # Добавляем список слов в список sentences
        new_sentences.append(words)

model = Word2Vec(new_sentences, vector_size=100, window=5, min_count=5, workers=4, epochs=7)
model.save(modelName)
