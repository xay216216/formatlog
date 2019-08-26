#!/usr/bin/env python
# coding=utf-8
from pyspark import SparkContext
from xlwt import *
from pyspark.sql import SparkSession
from pyspark.sql import Row
from pyspark.sql.types import StructField, StructType, StringType, ArrayType, IntegerType, LongType, FloatType
import datetime

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

today = datetime.datetime.now()
# 计算偏移量
offset = datetime.timedelta(days=-1)
# 获取想要的日期的时间
logdate = (today).strftime('%Y-%m-%d')
date_url = (today + offset).strftime('%Y%m%d')

#import xlwt
file=Workbook(encoding = 'utf-8')
#指定file以utf-8的格式打开
table=file.add_sheet('jinlione')

##创建表
spark = SparkSession \
    .builder \
    .appName("Python Spark SQL basic example") \
    .config("spark.some.config.option", "some-value") \
    .getOrCreate()

#df = spark.read.format("json").schema(schema).load("file:///d/GoWorkPath/src/formatlog/jinli_qingdianshang/" + date_url)
df = spark.read.format("json").schema(schema).load("file:///home/xiaoayong/work/jinli_qingdianshang/20190822")
#df = spark.read.format("json").schema(schema).load("hdfs://nameservice1/songshu/track/jinli_qingdianshang/" + date_url)
df.createOrReplaceTempView("jinlilog")

#详情页搜索PV
XqyssPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TODAY_FIRST_OPEN' ").count()
XqyssPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TODAY_FIRST_RUN_TAOBAO' ").count()


XqyssUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TODAY_FIRST_OPEN' ").count()
XqyssUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TODAY_FIRST_RUN_TAOBAO'  ").count()

print('SQG_TODAY_FIRST_OPENPV:',XqyssPv)

print('SQG_TODAY_FIRST_RUN_TAOBAOPV:',XqyssPvTB)

print('SQG_TODAY_FIRST_OPENUV:',XqyssUv)

print('SQG_TODAY_FIRST_RUN_TAOBAOUV:',XqyssUvTB)




