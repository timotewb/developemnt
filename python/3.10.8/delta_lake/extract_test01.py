
#Import SparkSession
from pyspark.sql import SparkSession
from pyspark import SparkContext
#Create Session
sc = SparkContext('spark://X10DAi.localdomain:7077', 'test001')
#sc = SparkContext.getOrCreate()
spark = SparkSession.builder.getOrCreate()

#http://172.18.65.170:8080

newJson = '{"Name":"something","Url":"https://stackoverflow.com","Author":"jangcy","BlogEntries":100,"Caller":"jangcy"}'
df = spark.read.json(sc.parallelize([newJson]))
df.show(truncate=False)

