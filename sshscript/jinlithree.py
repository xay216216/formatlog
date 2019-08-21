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
table=file.add_sheet('jinlithree')

##创建表
spark = SparkSession \
    .builder \
    .appName("Python Spark SQL basic example") \
    .config("spark.some.config.option", "some-value") \
    .getOrCreate()

df = spark.read.format("json").schema(schema).load("file:///home/xiaoayong/work/jinli_qingdianshang/" + date_url)
#df = spark.read.format("json").schema(schema).load("hdfs://nameservice1/songshu/track/jinli_qingdianshang/" + date_url)
df.createOrReplaceTempView("jinlilog")

#灭屏拉活
MplhPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN' ").count()
MplhPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN' AND seg.custom.eplatform='1' ").count()
MplhPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN' AND seg.custom.eplatform='2' ").count()

MplhUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN' ").count()
MplhUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN' AND seg.custom.eplatform='1' ").count()
MplhUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN' AND seg.custom.eplatform='2' ").count()

#弹窗显示
TcxsPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_SHOW' ").count()
TcxsPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_SHOW' AND seg.custom.eplatform='1' ").count()
TcxsPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_SHOW' AND seg.custom.eplatform='2' ").count()

TcxsUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_SHOW' ").count()
TcxsUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_SHOW' AND seg.custom.eplatform='1' ").count()
TcxsUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_SHOW' AND seg.custom.eplatform='2' ").count()

#激活点击
JhdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK' ").count()
JhdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK' AND seg.custom.eplatform='1' ").count()
JhdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK' AND seg.custom.eplatform='2' ").count()

JhdjUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK' ").count()
JhdjUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK' AND seg.custom.eplatform='1' ").count()
JhdjUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK' AND seg.custom.eplatform='2' ").count()

if TcxsUv == 0:
 Jhdjlv=0.00
else: 
 Jhdjlv=round(JhdjUv/TcxsUv,2)
if JhdjUvTB == 0:
 JhdjlvTB=0.00
else:  
 JhdjlvTB=round(JhdjUvTB/TcxsUvTB,2)
if TcxsUvPDD == 0:
 JhdjlvPDD=0.00
else:  
 JhdjlvPDD=round(JhdjUvPDD/TcxsUvPDD,2)

 #弹窗意外关闭
TcywgbPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE' ").count()
TcywgbPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE' AND seg.custom.eplatform='1' ").count()
TcywgbPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE' AND seg.custom.eplatform='2' ").count()

TcywgbUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE' ").count()
TcywgbUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE' AND seg.custom.eplatform='1' ").count()
TcywgbUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE' AND seg.custom.eplatform='2' ").count()

if TcxsUv == 0:
 Tcywgblv=0.00
else: 
 Tcywgblv=round(TcywgbUv/TcxsUv,2)
if TcxsUvTB == 0:
 TcywgblvTB=0.00
else:  
 TcywgblvTB=round(TcywgbUvTB/TcxsUvTB,2)
if TcxsUvPDD == 0:
 TcywgblvPDD=0.00
else:  
 TcywgblvPDD=round(TcywgbUvPDD/TcxsUvPDD,2)

 #下次提醒点击
XctxdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' ").count()
XctxdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='1' ").count()
XctxdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='2' ").count()

XctxdjUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' ").count()
XctxdjUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='1' ").count()
XctxdjUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE' AND seg.custom.eplatform='2' ").count()

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

 #浮标点击
FbdjPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK' ").count()
FbdjPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK' AND seg.custom.eplatform='1' ").count()
FbdjPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK' AND seg.custom.eplatform='2' ").count()

FbdjUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK' ").count()
FbdjUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK' AND seg.custom.eplatform='1' ").count()
FbdjUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK' AND seg.custom.eplatform='2' ").count()

if XctxdjUv == 0:
 Fbdjlv=0.00
else: 
 Fbdjlv=round(FbdjPv/XctxdjUv,2)
if XctxdjUvTB == 0:
 FbdjlvTB=0.00
else:  
 FbdjlvTB=round(FbdjPvTB/XctxdjUvTB,2)
if XctxdjUvPDD == 0:
 FbdjlvPDD=0.00
else:  
 FbdjlvPDD=round(FbdjPvPDD/XctxdjUvPDD,2)

 #小弹窗激活
XtcjhPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK' ").count()
XtcjhPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK' AND seg.custom.eplatform='1' ").count()
XtcjhPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK' AND seg.custom.eplatform='2' ").count()

XtcjhUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK' ").count()
XtcjhUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK' AND seg.custom.eplatform='1' ").count()
XtcjhUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK' AND seg.custom.eplatform='2' ").count()

if FbdjUv == 0:
 Xtcjhlv=0.00
else: 
 Xtcjhlv=round(XtcjhUv/FbdjUv,2)
if FbdjUvTB == 0:
 XtcjhlvTB=0.00
else:  
 XtcjhlvTB=round(XtcjhUvTB/FbdjUvTB,2)
if FbdjUvPDD == 0:
 XtcjhlvPDD=0.00
else:  
 XtcjhlvPDD=round(XtcjhUvPDD/FbdjUvPDD,2)

#小弹窗关闭
XtcgbPv=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE' ").count()
XtcgbPvTB=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE' AND seg.custom.eplatform='1' ").count()
XtcgbPvPDD=spark.sql("SELECT * FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE' AND seg.custom.eplatform='2' ").count()

XtcgbUv=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE' ").count()
XtcgbUvTB=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE' AND seg.custom.eplatform='1' ").count()
XtcgbUvPDD=spark.sql("SELECT distinct(m1) FROM jinlilog WHERE nm='GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE' AND seg.custom.eplatform='2' ").count()

if FbdjUv == 0:
 Xtcgblv=0.00
else: 
 Xtcgblv=round(XtcgbUv/FbdjUv,2)
if FbdjUvTB == 0:
 XtcgblvTB=0.00
else:  
 XtcgblvTB=round(XtcgbUvTB/FbdjUvTB,2)
if FbdjUvPDD == 0:
 XtcgblvPDD=0.00
else:  
 XtcgblvPDD=round(XtcgbUvPDD/FbdjUvPDD,2)


data={"1":[logdate,date_url,'全部',MplhPv,MplhUv,TcxsPv,TcxsUv,JhdjPv,JhdjUv,Jhdjlv,TcywgbPv,TcywgbUv,Tcywgblv,XctxdjPv,XctxdjUv,Xctxdjlv,FbdjPv,FbdjUv,Fbdjlv,XtcjhPv,XtcjhUv,Xtcjhlv,XtcgbPv,XtcgbUv,Xtcgblv],"2":[logdate,date_url,'淘宝',MplhPvTB,MplhUvTB,TcxsPvTB,TcxsUvTB,JhdjPvTB,JhdjUvTB,JhdjlvTB,TcywgbPvTB,TcywgbUvTB,TcywgblvTB,XctxdjPvTB,XctxdjUvTB,XctxdjlvTB,FbdjPvTB,FbdjUvTB,FbdjlvTB,XtcjhPvTB,XtcjhUvTB,XtcjhlvTB,XtcgbPvTB,XtcgbUvTB,XtcgblvTB],"3":[logdate,date_url,'拼多多',MplhPvPDD,MplhUvPDD,TcxsPvPDD,TcxsUvPDD,JhdjPvPDD,JhdjUvPDD,JhdjlvPDD,TcywgbPvPDD,TcywgbUvPDD,TcywgblvPDD,XctxdjPvPDD,XctxdjUvPDD,XctxdjlvPDD,FbdjPvPDD,FbdjUvPDD,FbdjlvPDD,XtcjhPvPDD,XtcjhUvPDD,XtcjhlvPDD,XtcgbPvPDD,XtcgbUvPDD,XtcgblvPDD]}

title = ["序号","报告生成时间","日期","电商平台","灭屏拉活电商PV","灭屏拉活UV","弹窗显示PV","弹窗显示UV","激活点击PV","激活点击UV","激活率","弹窗意外关闭PV","弹窗意外关闭UV","意外关闭率","下次提醒点击PV","下次提醒点击UV","关闭率","浮标点击PV","浮标点击UV","人均浮标点击","小弹窗激活PV","小弹窗激活UV","浮标激活率","小弹窗关闭PV","小弹窗关闭UV","小窗关闭率"]
for i in range(len(title)): # 循环列
 table.write(0,i,title[i])   # 将title数组中的字段写入到0行i列中
for line in data: #　循环字典
 table.write(int(line),0,line) #　将line写入到第int(line)行，第0列中
 for i in range(len(data[line])):
  table.write(int(line),i+1,data[line][i])
 file.save('/home/xiaoayong/work/worklog/jinli/' + date_url + '/jinlithree.xls')


