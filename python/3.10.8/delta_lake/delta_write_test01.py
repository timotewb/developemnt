import pyspark
from pyspark.sql import SparkSession
from pyspark import SparkContext
from delta import *

# builder = pyspark.sql.SparkSession.builder.appName("test001") \
#     .config("spark.sql.extensions", "io.delta.sql.DeltaSparkSessionExtension")

builder = pyspark.sql.SparkSession.builder.appName("test001") \
    .master("local") \
    .config("spark.jars.packages", "io.delta:delta-core_2.12:0.8.0") \
    .config("spark.sql.extensions", "io.delta.sql.DeltaSparkSessionExtension") \
    .config("spark.sql.catalog.spark_catalog", "org.apache.spark.sql.delta.catalog.DeltaCatalog")

#Import SparkSession
# from pyspark import SparkContext
#Create Session
# sc = SparkContext('spark://X10DAi.localdomain:7077', 'test001')
#spark = SparkSession.builder.getOrCreate()

spark = configure_spark_with_delta_pip(builder).getOrCreate()
sc = SparkContext.getOrCreate()




#http://172.18.65.170:8080

newJson = '{"Name":"something1","Url":"https://stackoverflow.com","Author":"jangcy","BlogEntries":100,"Caller":"jangcy"}'
df = spark.read.json(sc.parallelize([newJson]))
df.show(truncate=False)

df.write.format("delta").mode("append").save("delta_write_test01.parquet")

df1 = spark.read.format("delta").load("delta_write_test01.parquet")
df1.show()


#https://docs.delta.io/latest/quick-start.html#read-data&language-python