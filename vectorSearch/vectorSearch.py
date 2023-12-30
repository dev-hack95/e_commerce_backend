import json
from fastapi import FastAPI
from sentence_transformers import SentenceTransformer

server = FastAPI()
model = SentenceTransformer('sentence-transformers/all-MiniLM-L6-v2')

@server.get("/v1/getVector")
def getVector(productName: str):
    embeddings = model.encode(productName).tolist()
    return {"product_embeddings": embeddings, "shape": len(embeddings)}