#!/usr/bin/env python
# coding=utf-8
from pyspark import SparkContext
from xlwt import *

#import xlwt
file=Workbook(encoding = 'utf-8')
#指定file以utf-8的格式打开
table=file.add_sheet('baobiaoone')

sc=SparkContext(appName="jinlifour")
# sc.setLogLevel("ERROR")
sc.setLogLevel("ERROR")
rdd=sc.textFile("file:///d/GoWorkPath/src/formatlog/zheng_shihui/20190815")
#sc.textFile("file:///C:/spark-2.4.3-bin-hadoop2.7/data/testfile/HelloSpark.txt")
# rdd = sc.textFile('../csv/csv1/*')
numAs=rdd.filter(lambda s:'"nm":"GIONEE_OPEN_BANGO"' in s).count()
print(numAs)


#!/usr/bin/env python
# coding=utf-8
from pyspark.sql import SparkSession
from pyspark.sql import Row

spark = SparkSession \
    .builder \
    .appName("Python Spark SQL basic example") \
    .config("spark.some.config.option", "some-value") \
    .getOrCreate()

# teenagers = spark.sql("SELECT name FROM people WHERE age >= 13 AND age <= 19")

sc = spark.sparkContext

# Load a text file and convert each line to a Row.
lines = sc.textFile("file:///home/xiaoayong/work/jinli_qingdianshang/20190819")
parts = lines.map(lambda l: l.split(","))
jinlilog = parts.map(lambda p: Row(m1=p[2], nm=p[4]))

# Infer the schema, and register the DataFrame as a table.
schemaPeople = spark.createDataFrame(jinlilog)
schemaPeople.createOrReplaceTempView("jinlilog")

# SQL can be run over DataFrames that have been registered as a table.
XqyssPv = spark.sql("SELECT m1 FROM jinlilog WHERE nm ='SQG_TAOBAO_DETAIL_SEARCH' ").show()
print(XqyssPv)

XqyssPv = spark.sql("SELECT m1 FROM jinlilog WHERE nm ='SQG_TAOBAO_DETAIL_SEARCH' ").show()
print(numAs)

XqyssPv = spark.sql("SELECT m1 FROM jinlilog WHERE nm ='SQG_TAOBAO_DETAIL_SEARCH' ").show()
print(numAs)

XqyssPv = spark.sql("SELECT m1 FROM jinlilog WHERE nm ='SQG_TAOBAO_DETAIL_SEARCH' ").show()



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

spark.read.format("json").schema(schema).load("file:///home/xiaoayong/work/zheng_shihui/20190819").createOrReplaceTempView("jinlilog")