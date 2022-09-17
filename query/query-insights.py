import cohere
import numpy as np
import re
import pandas as pd
from tqdm import tqdm
from datasets import load_dataset
import umap
import altair as alt
from sklearn.metrics.pairwise import cosine_similarity
from annoy import AnnoyIndex
import warnings
warnings.filterwarnings('ignore')
pd.set_option('display.max_colwidth', None)
import os
import argparse

parser = argparse.ArgumentParser()
parser.add_argument("-q", "--query", help="Query for processing.")
args = parser.parse_args()

api_key = None
with open('../secrets/cohere') as f:
    api_key = f.readlines()
output = './output'
co = cohere.Client(api_key)

# Get dataset
dataset = load_dataset("../data/set", split="train")
df = pd.DataFrame(dataset)[:1000] # INCREASE LATER
df.head(10)

embeds = co.embed(texts=list(df['text']),
                  model="large",
                  truncate="LEFT").embeddings

# Check the dimensions of the embeddings
embeds = np.array(embeds)
embeds.shape

# 3.2. Find the neighbors of a user query
query = args.query

query_embed = co.embed(texts=[query],
                  model="large",
                  truncate="LEFT").embeddings
similar_item_ids = search_index.get_nns_by_vector(query_embed[0],10,
                                                include_distances=True)
results = pd.DataFrame(data={'texts': df.iloc[similar_item_ids[0]]['text'], 
                             'distance': similar_item_ids[1]})

# capture results from output
#print(results)
results.to_pickle(output) # serialize pandas dataframe with pickle

