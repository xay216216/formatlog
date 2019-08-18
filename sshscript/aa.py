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


0 14 * * * python /home/xiaoayong/work/script/gethdfslog.py

* * * * * bash /home/xiaoayong/work/script/zhengshihuireportone.sh

10 * * * * bash /home/xiaoayong/work/script/jinlireportthree.sh

20 * * * * bash /home/xiaoayong/work/script/jinlireportfour.sh



10 01 * * * bash /home/xiaoayong/work/script/everyday.sh  1>/home/xiaoayong/work/script/everyday.log 2>&1

25 02 * * * bash /home/xiaoayong/work/script/jinlireportone.sh 1>/home/xiaoayong/work/script/jinlione.log 2>&1

10 03 * * * bash /home/xiaoayong/work/script/jinlireportthree.sh 1>/home/xiaoayong/work/script/jinlithree.log 2>&1

10 04 * * * bash /home/xiaoayong/work/script/jinlireportfour.sh 1>/home/xiaoayong/work/script/jinlifour.log 2>&1

30 09 * * * bash /home/xiaoayong/work/script/sendemail.sh  1>/home/xiaoayong/work/script/sendemail.log 2>&1



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






