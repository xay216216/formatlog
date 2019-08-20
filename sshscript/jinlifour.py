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
table=file.add_sheet('jinlifour')

##创建表
spark = SparkSession \
    .builder \
    .appName("Python Spark SQL basic example") \
    .config("spark.some.config.option", "some-value") \
    .getOrCreate()

df = spark.read.format("json").schema(schema).load("file:///home/xiaoayong/work/jinli_qingdianshang/" + date_url)
#df = spark.read.format("json").schema(schema).load("hdfs://nameservice1/songshu/track/jinli_qingdianshang/" + date_url)
df.createOrReplaceTempView("jinlilog")

#帮购启动PV
BgqdPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_OPEN_BANGO' ").count()
BgqdPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_OPEN_BANGO' AND seg.custom.eplatform='1' ").count()
BgqdPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_OPEN_BANGO' AND seg.custom.eplatform='2' ").count()

BgqdUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_OPEN_BANGO' ").count()
BgqdUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_OPEN_BANGO' AND seg.custom.eplatform='1' ").count()
BgqdUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_OPEN_BANGO' AND seg.custom.eplatform='2' ").count()

#弹窗显示PV
TcxsPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_SHOW' ").count()
TcxsPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_SHOW' AND seg.custom.eplatform='1' ").count()
TcxsPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_SHOW' AND seg.custom.eplatform='2' ").count()

TcxsUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_SHOW' ").count()
TcxsUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_SHOW' AND seg.custom.eplatform='1' ").count()
TcxsUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_SHOW' AND seg.custom.eplatform='2' ").count()

#弹窗激活PV
TcjhPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK' ").count()
TcjhPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK' AND seg.custom.eplatform='1' ").count()
TcjhPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK' AND seg.custom.eplatform='2' ").count()

TcjhUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK' ").count()
TcjhUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK' AND seg.custom.eplatform='1' ").count()
TcjhUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK' AND seg.custom.eplatform='2' ").count()

if TcxsUv == 0:
 Tcjhlv=0.00
else: 
 Tcjhlv=round(TcjhUv/TcxsUv,2)
if TcxsUvTB == 0:
 TcjhlvTB=0.00
else: 
 TcjhlvTB=round(TcjhUvTB/TcxsUvTB,2)
if TcjhlvPDD == 0:
 JzspdjlvPDD=0.00
else: 
 TcjhlvPDD=round(TcjhUvPDD/TcxsUvPDD,2)

#弹窗到后台
TcdhtPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND' ").count()
TcdhtPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND' AND seg.custom.eplatform='1' ").count()
TcdhtPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND' AND seg.custom.eplatform='2' ").count()

TcdhtUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND' ").count()
TcdhtUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND' AND seg.custom.eplatform='1' ").count()
TcdhtUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND' AND seg.custom.eplatform='2' ").count()

if TcxsUv == 0:
 Tcdhtlv=0.00
else: 
 Tcdhtlv=round(TcdhtUv/TcxsUv,2)
if TcxsUvTB == 0:
 TcdhtlvTB=0.00
else: 
 TcdhtlvTB=round(TcdhtUvTB/TcxsUvTB,2)
if TcxsUvPDD == 0:
 TcdhtlvPDD=0.00
else: 
 TcdhtlvPDD=round(TcdhtUvPDD/TcxsUvPDD,2)

#下次提醒点击
XctxdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' ").count()
XctxdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='1' ").count()
XctxdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='2' ").count()

XctxdjUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' ").count()
XctxdjUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='1' ").count()
XctxdjUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='2' ").count()

if TcxsUv == 0:
 Xctxdjlv=0.00
else: 
 Xctxdjlv=round(XctxdjUv/TcxsUv,2)
if TcxsUvTB == 0:
 XctxdjlvTB=0.00
else: 
 XctxdjlvTB=round(XctxdjUvTB/TcxsUvTB,2)
if TcxsUvPDD == 0:
 XctxdjlvPDD=0.00
else: 
 XctxdjlvPDD=round(XctxdjUvPDD/TcxsUvPDD,2)

#底部tab展示PV
DbtabzsPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_TAB_SHOW' ").count()
DbtabzsPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_TAB_SHOW' AND seg.custom.eplatform='1' ").count()
DbtabzsPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_TAB_SHOW' AND seg.custom.eplatform='2' ").count()

DbtabzsUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_TAB_SHOW' ").count()
DbtabzsUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_TAB_SHOW' AND seg.custom.eplatform='1' ").count()
DbtabzsUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_TAB_SHOW' AND seg.custom.eplatform='2' ").count()

JianfaUv=TcxsUv-$TcdhtUv
JianfalvTB=TcxsPvTB-$TcdhtUvTB
JianfalvPDD=TcxsPvPDD-$TcdhtUvPDD

if JianfaUv == 0:
 Dbtabzslv=0.00
else: 
 Dbtabzslv=round(DbtabzsPv/JianfaUv,2)
if JianfaUvTB == 0:
 DbtabzslvTB=0.00
else: 
 DbtabzslvTB=round(DbtabzsPvTB/JianfaUvTB,2)
if JianfaUvPDD == 0:
 DbtabzslvPDD=0.00
else: 
 DbtabzslvPDD=round(DbtabzsPvPDD/JianfaUvPDD,2)

#底部tab点击PV
DbtabdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' ").count()
DbtabdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='1' ").count()
DbtabdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='2' ").count()

DbtabdjUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' ").count()
DbtabdjUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='1' ").count()
DbtabdjUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='2' ").count()

if DbtabzsUv == 0:
 Dbtabdjlv=0.00
else: 
 Dbtabdjlv=round(DbtabdjUv/DbtabzsUv,2)
if DbtabzsUvTB == 0:
 DbtabdjlvTB=0.00
else: 
 DbtabdjlvTB=round(DbtabdjUvTB/DbtabzsUvTB,2)
if DbtabzsUvPDD == 0:
 DbtabdjlvPDD=0.00
else: 
 DbtabdjlvPDD=round(DbtabdjUvPDD/DbtabzsUvPDD,2)


#助手页开启点击PV
ZsykqdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_HELPER_SWITCH_PAGE_OPEN_BTN' ").count()
ZsykqdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_HELPER_SWITCH_PAGE_OPEN_BTN' AND seg.custom.eplatform='1' ").count()
ZsykqdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='SQG_HELPER_SWITCH_PAGE_OPEN_BTN' AND seg.custom.eplatform='2' ").count()

ZsykqdjUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_HELPER_SWITCH_PAGE_OPEN_BTN' ").count()
ZsykqdjUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_HELPER_SWITCH_PAGE_OPEN_BTN' AND seg.custom.eplatform='1' ").count()
ZsykqdjUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='SQG_HELPER_SWITCH_PAGE_OPEN_BTN' AND seg.custom.eplatform='2' ").count()

if DbtabdjUv == 0:
 Zsykqdjlv=0.00
else: 
 Zsykqdjlv=round(ZsykqdjUv/DbtabdjUv,2)
if DbtabdjUvTB == 0:
 ZsykqdjlvTB=0.00
else: 
 ZsykqdjlvTB=round(ZsykqdjUvTB/DbtabdjUvTB,2)
if DbtabdjUvPDD == 0:
 ZsykqdjlvPDD=0.00
else: 
 ZsykqdjlvPDD=round(ZsykqdjUvPDD/DbtabdjUvPDD,2)


if DbtabdjUv == 0:
 Zsykqlv=0.00
else: 
 Zsykqlv=round(ZsykqdjUv/DbtabdjUv,2)
if DbtabdjUvTB == 0:
 ZsykqlvTB=0.00
else: 
 ZsykqlvTB=round(ZsykqdjUvTB/DbtabdjUvTB,2)
if DbtabdjUvPDD == 0:
 ZsykqlvPDD=0.00
else: 
 ZsykqlvPDD=round(ZsykqdjUvPDD/DbtabdjUvPDD,2)


data={"1":[logdate,date_url,'全部',BgqdPv,BgqdUv,TcxsPv,TcxsUv,TcjhPv,TcjhUv,Tcjhlv,TcdhtPv,TcdhtUv,Tcdhtlv,XctxdjPv,XctxdjUv,Xctxdjlv,DbtabzsPv,DbtabzsUv,Dbtabzslv,DbtabdjPv,DbtabdjUv,Dbtabdjlv,ZsykqdjPv,ZsykqdjUv,Zsykqdjlv,Zsykqlv],"2":[logdate,date_url,'淘宝',BgqdPvTB,BgqdUvTB,TcxsPvTB,TcxsUvTB,TcjhPvTB,TcjhUvTB,TcjhlvTB,TcdhtPvTB,TcdhtUvTB,TcdhtlvTB,XctxdjPvTB,XctxdjUvTB,XctxdjlvTB,DbtabzsPvTB,DbtabzsUvTB,DbtabzslvTB,DbtabdjPvTB,DbtabdjUvTB,DbtabdjlvTB,ZsykqdjPvTB,ZsykqdjUvTB,ZsykqdjlvTB,ZsykqlvTB],"3":[logdate,date_url,'拼多多',BgqdPvPDD,BgqdUvPDD,TcxsPvPDD,TcxsUvPDD,TcjhPvPDD,TcjhUvPDD,TcjhlvPDD,TcdhtPvPDD,TcdhtUvPDD,TcdhtlvPDD,XctxdjPvPDD,XctxdjUvPDD,XctxdjlvPDD,DbtabzsPvPDD,DbtabzsUvPDD,DbtabzslvPDD,DbtabdjPvPDD,DbtabdjUvPDD,DbtabdjlvPDD,ZsykqdjPvPDD,ZsykqdjUvPDD,ZsykqdjlvPDD,ZsykqlvPDD]}

title = ["序号","报告生成时间","日期","电商平台","帮购启动PV","帮购启动UV","弹窗显示PV","弹窗显示UV","弹窗激活PV","弹窗激活UV","激活率","弹窗到后台PV","弹窗到后台Uv","弹窗到后台率","下次提醒点击PV","下次提醒点击UV","关闭率","底部tab展示PV","底部tab展示UV","tab展示率","底部tab点击PV","底部tab点击UV","tab点击率","助手页开启点击PV","助手页开启点击UV","助手页开启点击率","助手页开启率"]
for i in range(len(title)): # 循环列
 table.write(0,i,title[i])   # 将title数组中的字段写入到0行i列中
for line in data: #　循环字典
 table.write(int(line),0,line) #　将line写入到第int(line)行，第0列中
 for i in range(len(data[line])):
  table.write(int(line),i+1,data[line][i])
 file.save('/home/xiaoayong/work/worklog/jinli/' + date_url + '/jinlifour.xls')


