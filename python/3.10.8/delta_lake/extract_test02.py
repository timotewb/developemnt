import requests
import json
from pyspark.sql import SparkSession
from pyspark import SparkContext
sc = SparkContext('spark://X10DAi.localdomain:7077', 'test001')
spark = SparkSession.builder.getOrCreate()

url = "http://192.168.144.131:8000/sql"

payload = "select * from rss_news limit 5;"
headers = {
  'Content-Type': 'application/json',
  'NS': 'many',
  'DB': 'db01',
  'Authorization': 'Basic ZXRsOmV0bA=='
}

response = requests.request("POST", url, headers=headers, data=payload)

aDict = json.loads(response.text)

print(aDict[0]['result'][0].keys())

df = spark.read.json(sc.parallelize([str(aDict[0]['result'])]))
df.show(truncate=False)