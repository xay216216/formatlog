#!/usr/bin/env python
# coding=utf-8
from pyspark.sql import SparkSession
from pyspark.sql import Row
from pyspark.sql.types import StructField, StructType, StringType, ArrayType, IntegerType, LongType, FloatType

schema = StructType([
    StructField("ba",StringType(),True),
    StructField("ip",StringType(),True),
    StructField("m1",StringType(),True),
    StructField("mo",StringType(),True),
    StructField("nm",StringType(),True),
    StructField("p",StringType(),True),
    StructField("seg",StructType([
        StructField("custom",StructType([
            StructField("cid",StringType(),True),
            StructField("appv",StringType(),True),
            StructField("categoryId",StringType(),True),
            StructField("itemid",StringType(),True),
            StructField("m1",StringType(),True),
            StructField("pageIndex",StringType(),True),
            StructField("pid",StringType(),True),
            StructField("qd",StringType(),True),
            StructField("match",StringType(),True),
            StructField("uid",StringType(),True),
            StructField("eplatform",StringType(),True),
            StructField("clickType",StringType(),True),
            StructField("index",StringType(),True),
            StructField("tid",StringType(),True),
            StructField("position",StringType(),True),
            StructField("type",StringType(),True),
            StructField("name",StringType(),True),
        ]),True),
        StructField("keyWords",StringType(),True),
    ]),True),
    StructField("tm",LongType(),True),
    StructField("v",StringType(),True),
])

spark = SparkSession \
    .builder \
    .appName("Python Spark SQL basic example") \
    .config("spark.some.config.option", "some-value") \
    .getOrCreate()

df = spark.read.format("json").schema(schema).load("file:///home/xiaoayong/work/zheng_shihui/20190819")
df.createOrReplaceTempView("jinlilog")

XqyssPv=spark.sql("SELECT * FROM jinlilog WHERE nm='HOME_JINGANG_PV' ").count()
print(XqyssPv)

# SQL can be run over DataFrames that have been registered as a table.
XqyssPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='JINGANG_CLICK_EVERYTIME' AND seg.custom.name='大额神券'").count()
print(XqyssPvTB)

XqyssUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='HOME_JINGANG_PV' ").count()
print(XqyssUv)


