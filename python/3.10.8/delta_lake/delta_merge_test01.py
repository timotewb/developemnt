from pyspark.sql import SparkSession
from pyspark import SparkContext
from delta import *

# builder = pyspark.sql.SparkSession.builder.appName("test001") \
#     .config("spark.sql.extensions", "io.delta.sql.DeltaSparkSessionExtension")

builder = SparkSession.builder.appName("test001") \
    .master("local") \
    .config("spark.jars.packages", "io.delta:delta-core_2.12:0.8.0") \
    .config("spark.sql.extensions", "io.delta.sql.DeltaSparkSessionExtension") \
    .config("spark.sql.catalog.spark_catalog", "org.apache.spark.sql.delta.catalog.DeltaCatalog")

spark = configure_spark_with_delta_pip(builder).getOrCreate()
sc = SparkContext.getOrCreate()