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

df = spark.read.format("json").schema(schema).load("file:///d/GoWorkPath/src/formatlog/jinli_qingdianshang/" + date_url)
#df = spark.read.format("json").schema(schema).load("file:///home/xiaoayong/work/jinli_qingdianshang/" + date_url)
#df = spark.read.format("json").schema(schema).load("hdfs://nameservice1/songshu/track/jinli_qingdianshang/" + date_url)
df.createOrReplaceTempView("jinlilog")

#详情页搜索PV
XqyssPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_DETAIL_SEARCH' ").count()
XqyssPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_DETAIL_SEARCH' AND seg.custom.eplatform='1' ").count()
XqyssPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_DETAIL_SEARCH' AND seg.custom.eplatform='2' ").count()

XqyssUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_DETAIL_SEARCH' ").count()
XqyssUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_DETAIL_SEARCH' AND seg.custom.eplatform='1' ").count()
XqyssUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_DETAIL_SEARCH' AND seg.custom.eplatform='2' ").count()

if XqyssUv == 0:
 Xqysslv=0.00
else: 
 Xqysslv=round(XqyssPv/XqyssUv,2)
if XqyssUvTB == 0:
 XqysslvTB=0.00
else:  
 XqysslvTB=round(XqyssPvTB/XqyssUvTB,2)
if XqyssUvPDD == 0:
 XqysslvPDD=0.00
else:  
 XqysslvPDD=round(XqyssPvPDD/XqyssUvPDD,2)

#精准商品PV
JzspPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT' ").count()
JzspPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT' AND seg.custom.eplatform='1' ").count()
JzspPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT' AND seg.custom.eplatform='2' ").count()

JzspUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT' ").count()
JzspUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT' AND seg.custom.eplatform='1' ").count()
JzspUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT' AND seg.custom.eplatform='2' ").count()

#精准商品点击PV
JzspdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' ").count()
JzspdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' AND seg.custom.eplatform='1' ").count()
JzspdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' AND seg.custom.eplatform='2' ").count()

JzspdjUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' ").count()
JzspdjUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' AND seg.custom.eplatform='1' ").count()
JzspdjUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' AND seg.custom.eplatform='2' ").count()

if JzspPv == 0:
 Jzspdjlv=0.00
else: 
 Jzspdjlv=round(JzspdjPv/JzspPv,2)
if JzspPvTB == 0:
 JzspdjlvTB=0.00
else: 
 JzspdjlvTB=round(JzspdjPvTB/JzspPvTB,2)
if JzspPvPDD == 0:
 JzspdjlvPDD=0.00
else: 
 JzspdjlvPDD=round(JzspdjPvPDD/JzspPvPDD,2)

#相似商品PV
SsspPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT' ").count()
SsspPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT' AND seg.custom.eplatform='1' ").count()
SsspPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT' AND seg.custom.eplatform='2' ").count()

SsspUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT' ").count()
SsspUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT' AND seg.custom.eplatform='1' ").count()
SsspUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT' AND seg.custom.eplatform='2' ").count()

#相似商品点击PV
SsspdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK' ").count()
SsspdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK' AND seg.custom.eplatform='1' ").count()
SsspdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK' AND seg.custom.eplatform='2' ").count()

SsspdjUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK' ").count()
SsspdjUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK' AND seg.custom.eplatform='1' ").count()
SsspdjUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK' AND seg.custom.eplatform='2' ").count()

if SsspPv == 0:
 Ssspdjlv=0.00
else: 
 Ssspdjlv=round(SsspdjPv/SsspPv,2)
if SsspPvTB == 0:
 SsspdjlvTB=0.00
else: 
 SsspdjlvTB=round(SsspdjPvTB/SsspPvTB,2)
if SsspPvPDD == 0:
 SsspdjlvPDD=0.00
else: 
 SsspdjlvPDD=round(SsspdjPvPDD/SsspPvPDD,2)

#无结果PV
WjgPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_NO_RESULT' ").count()
WjgPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_NO_RESULT' AND seg.custom.eplatform='1' ").count()
WjgPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_NO_RESULT' AND seg.custom.eplatform='2' ").count()

#结果总点击PV
JgzdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' OR nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK' ").count()
JgzdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE (nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' OR nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK') AND seg.custom.eplatform='1' ").count()
JgzdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE (nm='SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK' OR nm='SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK') AND seg.custom.eplatform='2' ").count()

#下单PV
XdPv=0
XdPvTB=0
XdPvPDD=0
Xdlv=0
XdlvTB=0
XdlvPDD=0


data={"1":[logdate,date_url,'全部',XqyssPv,XqyssUv,Xqysslv,JzspPv,JzspdjPv,Jzspdjlv,SsspPv,SsspdjPv,Ssspdjlv,WjgPv,JgzdjPv,XdPv,Xdlv,JzspUv,JzspdjUv,SsspUv,SsspdjUv],"2":[logdate,date_url,'淘宝',XqyssPvTB,XqyssUvTB,XqysslvTB,JzspPvTB,JzspdjPvTB,JzspdjlvTB,SsspPvTB,SsspdjPvTB,SsspdjlvTB,WjgPvTB,JgzdjPvTB,XdPvTB,XdlvTB,JzspUvTB,JzspdjUvTB,SsspUvTB,SsspdjUvTB],"3":[logdate,date_url,'拼多多',XqyssPvPDD,XqyssUvPDD,XqysslvPDD,JzspPvPDD,JzspdjPvPDD,JzspdjlvPDD,SsspPvPDD,SsspdjPvPDD,SsspdjlvPDD,WjgPvPDD,JgzdjPvPDD,XdPvPDD,XdlvPDD,JzspUvPDD,JzspdjUvPDD,SsspUvPDD,SsspdjUvPDD]}

title = ["序号","报告生成时间","日期","电商平台","详情页搜索PV","详情页搜索UV","人均详情次数","精准商品PV","精准商品点击PV","精准商品点击率","相似商品PV","相似商品点击PV","相似商品点击率","无结果PV","结果总点击PV","下单PV","下单率","精准商品UV","精准商品点击UV","相似商品UV","相似商品点击UV"]
for i in range(len(title)): # 循环列
 table.write(0,i,title[i])   # 将title数组中的字段写入到0行i列中
for line in data: #　循环字典
 table.write(int(line),0,line) #　将line写入到第int(line)行，第0列中
 for i in range(len(data[line])):
  table.write(int(line),i+1,data[line][i])
 file.save('/home/xiaoayong/work/worklog/jinli/' + date_url + '/jinlione.xls')


