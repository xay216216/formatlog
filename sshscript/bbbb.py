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
sc.setLogLevel("INFO")
rdd=sc.textFile("file:///home/xiaoayong/work/jinli_qingdianshang/20190818")
# rdd = sc.textFile('../csv/csv1/*')
#print(rdd.count())
numAs=rdd.filter(lambda s:'"nm":"GIONEE_OPEN_BANGO"' in s).count()
numAsa=rdd.filter(lambda s:'"nm":"GIONEE_BANGO_HELPER_DIALOG_SHOW"' in s).count()
numAsb=rdd.filter(lambda s:'"nm":"GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK"' in s).count()
numAsc=rdd.filter(lambda s:'"nm":"GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND"' in s).count()
data={"1":["张三",numAs,numAsa,100],"2":["李四",numAsb,numAsc,95],"3":["王五",60,66,68]}

title = ["学号","姓名","语文成绩","数学成绩","英语成绩","总分","平均分"]
for i in range(len(title)): # 循环列
 table.write(0,i,title[i])   # 将title数组中的字段写入到0行i列中
for line in data: #　循环字典
 print('line:',line)
 table.write(int(line),0,line) #　将line写入到第int(line)行，第0列中
 summ = data[line][2] + data[line][3] # 成绩总分
 table.write(int(line),5,summ) # 总分
 table.write(int(line),6,summ/3) # 平均分
 for i in range(len(data[line])):
  table.write(int(line),i+1,data[line][i])
 file.save('baobiaoone.xls')


