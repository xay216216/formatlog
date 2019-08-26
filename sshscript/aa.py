#!/usr/bin/python
import sys
#reload(sys)
#sys.setdefaultencoding('utf8')
 
import subprocess
import os
import commands
 
cmd1 = "cd /d/GoWorkPath/src/formatlog/jinli_qingdianshang"
cmd2 = "cd 20190812/ "

cmd = cmd1 + " && " + cmd2 
 
 
subprocess.call(cmd,shell=True)

#cmdss = 'cat * | grep  GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN | wc -l '

#status, output = commands.getstatusoutput(cmdss)

output = os.popen('cat * | grep  GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN | wc -l ')
print output.read()


#process2 = subprocess.Popen(cmdount, shell=True,stdout=subprocess.PIPE, stderr=subprocess.PIPE)


#!/usr/bin/python
import sys
reload(sys)
sys.setdefaultencoding('utf8')

import subprocess
import os
import commands

cmd1 = "cd /home/xiaoayong/work/zheng_shihui"
cmd2 = "cd /home/xiaoayong/work/zheng_shihui"
cmd3 = "hdfs dfs -get hdfs://nameservice1/songshu/track/zheng_shihui/`date -d 'yesterday' +'%Y%m%d'` "
cmd = cmd1 + " && " + cmd2

subprocess.call(cmd,shell=True)

`date --date='1 days ago' +%Y%m%d`


10 01 * * * bash /home/xiaoayong/work/script/everyday.sh  1>/home/xiaoayong/work/script/everyday.log 2>&1

20 02 * * * /data/server/pinedd/spark-2.4.3/bin/spark-submit /home/xiaoayong/work/script/jinlione.py 1>/home/xiaoayong/work/script/jinlione.log 2>&1

20 03 * * * /data/server/pinedd/spark-2.4.3/bin/spark-submit /home/xiaoayong/work/script/jinlithree.py 1>/home/xiaoayong/work/script/jinlithree.log 2>&1

20 04 * * * /data/server/pinedd/spark-2.4.3/bin/spark-submit /home/xiaoayong/work/script/jinlifour.py 1>/home/xiaoayong/work/script/jinlifour.log 2>&1

50 09 * * * bash /home/xiaoayong/work/script/sendemail.sh  1>/home/xiaoayong/work/script/sendemail.log 2>&1

55 09 * * * bash /home/xiaoayong/work/script/sendemailtwo.sh  1>/home/xiaoayong/work/script/sendemailtwo.log 2>&1


/data/server/pinedd/spark-2.4.3/bin/spark-submit /home/xiaoayong/work/script/tmp.py

* * * * * bash /home/xiaopeng9/t.sh 1>/tmp/xxxx.log 2>&1

 export HADOOP_HOME=/data/server/pinedd/hadoop-3.1.2/bin

 export PATH=$PATH:$

 vim ~/.bash_profile

 source ~/.bash_profile


 vim ~/.bashrc 

 source ~/.bashrc


 /usr/bin/sendemail --help

 /usr/bin/sendemail

-f  619341326@qq.com 发件人邮箱地址

-t  test@qq.com  收件人邮箱

-s  smtp.qq.com  发件人邮箱的smtp服务器

-u  '标题'  邮件的主题

-o message-content-type=html  邮件内容的格式为html，也可以是text

-o message-charset=utf8  邮件内容编码

-xu 619341326@qq.com  发件人账号

-xp 123456  发件人密码

-m  '邮件内容'  邮件的内容

/opt/spark-2.1.0/bin/spark-submit bbb.py

/opt/spark-2.1.0/bin/spark-submit bbb.py
/d/spark-2.4.3/bin/spark-submit wordCount.py

spark-submit D:/GoWorkPath/src/formatlog/sshscript/wordCount.py

安装xlwt
一 安装pip
curl -O https://bootstrap.pypa.io/get-pip.py
python get-pip.py

二 安装xlrd--读 xlwt--写
xlrd：是python从excel读数据的第三方控件；
xlwt：是python从excel写数据的第三方控件；
xlutils：是python使用xlrd、xlwt的工具箱。若安装不成功，可能原因是需要安装setuptools。
pip  install xlrd
pip  install xlwt
pip  install xlutils







