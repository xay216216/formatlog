#!/bin/bash
#.  /etc/profile
source /etc/profile      
source ~/.bash_profile
export LANG="en_US.UTF-8"

pyspark
#帮购启动PV
#spark.read.format("json").load("hdfs://nameservice1/songshu/track/jinli_qingdianshang/` date -d 'yesterday' +'%Y%m%d' ` ").createOrReplaceTempView("ds")
spark.read.format("json").load("/home/xiaoayong/work/jinli_qingdianshang/20190818").createOrReplaceTempView("ds")

XqyssUv=` spark.sql("SELECT * FROM ds where nm='GIONEE_OPEN_BANGO' ").count() `

echo ${XqyssUv}


/opt/spark-2.1.0/bin/spark-submit test.py


weiyuqin@baice100.com  




from pyspark import SparkContext

sc = SparkContext(appName="jinlifour")
# sc.setLogLevel("ERROR")
sc.setLogLevel("INFO")
rdd = sc.textFile("file:///home/xiaoayong/work/jinli_qingdianshang/20190818")
# rdd = sc.textFile('../csv/csv1/*')
numAs=rdd.filter(lambda s:'GIONEE_OPEN_BANGO' in s).count()
print(numAs)


from xlwt import *
#需要xlwt库的支持
#import xlwt
file = Workbook(encoding = 'utf-8')
#指定file以utf-8的格式打开
table = file.add_sheet('data')
#指定打开的文件名
  
data = {
    "1":["张三",150,120,100],
    "2":["李四",90,99,95],
    "3":["王五",60,66,68]
    }
#字典数据
  
ldata = []
num = [a for a in data]
#for循环指定取出key值存入num中
num.sort()
#字典数据取出后无需，需要先排序
  
for x in num:
#for循环将data字典中的键和值分批的保存在ldata中
  t = [int(x)]
  for a in data[x]:
    t.append(a)
  ldata.append(t)
  
for i,p in enumerate(ldata):
#将数据写入文件,i是enumerate()函数返回的序号数
  for j,q in enumerate(p):
    # print i,j,q
    table.write(i,j,q)
file.save('data.xlsx')