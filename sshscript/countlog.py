#!/usr/bin/python
import sys
reload(sys)
sys.setdefaultencoding('utf8')
 
import subprocess
import os
import commands
import xlwt
 
cmd1 = "cd /home/xiaoayong/work/hdfslog"
cmd2 = "hdfs dfs -get hdfs://nameservice1/songshu/track/zheng_shihui/`date -d 'yesterday' +'%Y%m%d'` "
cmd = cmd1 + " && " + cmd2 
 
subprocess.call(cmd,shell=True)

cmdc = "cd `date -d 'yesterday' +'%Y%m%d'`  "
cmdc1 = "cat * | grep ' 'nm':'HOME_JINGANG_PV' ' | wc -l "
cmdount = cmdc + " && " + cmdc1 

def run_cmd(cmdount):
result_str = ''
process = subprocess.Popen(cmd, shell=True,stdout=subprocess.PIPE, stderr=subprocess.PIPE)
result_f = process.stdout
result_astr = result_f.read().strip()

wbk = xlwt.Workbook(encoding = 'ascii') 
sheet = wbk.add_sheet("wordCount")#Excel单元格名字 

sheet.write(i, 1, label = result_str) 
sheet.write(i, 0, label = result_str) 







# coding: UTF-8
import sys
reload(sys)
sys.setdefaultencoding('utf8')
 
import subprocess
import os
import commands
import xlwt
 
cmd1 = "cd /home/xiaoayong/work/hdfslog"
cmd2 = "hdfs dfs -get hdfs://nameservice1/songshu/track/zheng_shihui/`date -d 'yesterday' +'%Y%m%d'` "
cmd = cmd1 + " && " + cmd2 
 
subprocess.call(cmd,shell=True)

cmdc = "cd `date -d 'yesterday' +'%Y%m%d'`  "
cmdc1 = "cat * | grep ' 'nm':'HOME_JINGANG_PV' ' | wc -l "
cmdount = cmdc + " && " + cmdc1 

def run_cmd(cmdount):
result_str = ''
process = subprocess.Popen(cmd, shell=True,stdout=subprocess.PIPE, stderr=subprocess.PIPE)
result_f = process.stdout
result_str = result_f.read().strip()

wbk = xlwt.Workbook(encoding = 'ascii') 
sheet = wbk.add_sheet("wordCount")#Excel单元格名字 

sheet.write(i, 1, label = result_str) 
sheet.write(i, 0, label = result_str) 
wbk.save('wordCount.xls') #保存为 wordCount.xls文件 


sc=SparkContext(appName="jinlione")
# sc.setLogLevel("ERROR")
sc.setLogLevel("INFO")
rdd=sc.textFile("hdfs://nameservice1/songshu/track/jinli_qingdianshang/20190818")
