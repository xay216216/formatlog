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

df = spark.read.format("json").schema(schema).load("file:///D:/GoWorkPath/src/formatlog/zheng_shihui/" + date_url)
#df = spark.read.format("json").schema(schema).load("file:///home/xiaoayong/work/zheng_shihui/20190821")
#df = spark.read.format("json").schema(schema).load("hdfs://nameservice1/songshu/track/jinli_qingdianshang/" + date_url)
df.createOrReplaceTempView("jinlilog")

#详情页搜索PV
XqyssPv=spark.sql("SELECT * FROM jinlilog WHERE nm='HOME_JINGANG_PV' ").count()
XqyssPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='JINGANG_CLICK_EVERYTIME' ").count()


XqyssUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='HOME_JINGANG_PV' ").count()
XqyssUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='JINGANG_CLICK_EVERYTIME'  ").count()

print('HOME_JINGANG_PV:',XqyssPv)

print('SQG_TODAY_FIRST_RUN_TAOBAOPV:',XqyssPvTB)

print('SQG_TODAY_FIRST_OPENUV:',XqyssUv)

print('SQG_TODAY_FIRST_RUN_TAOBAOUV:',XqyssUvTB)

data={"1":[logdate,date_url,'全部',XqyssPv,XqyssPvTB,XqyssUv,XqyssUvTB]}
title = ["序号","报告生成时间","日期","电商平台","详情页搜索PV","详情页搜索UV","人均详情次数","精准商品PV"]
for i in range(len(title)): # 循环列
    table.write(0,i,title[i])   # 将title数组中的字段写入到0行i列中
for line in data: #　循环字典
    table.write(int(line),0,line) #　将line写入到第int(line)行，第0列中
    for i in range(len(data[line])):
        table.write(int(line),i+1,data[line][i])
    file.save('test.xls')




