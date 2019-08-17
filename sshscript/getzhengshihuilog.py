#!/usr/bin/python
import sys
reload(sys)
sys.setdefaultencoding('utf8')

import subprocess
import os
import commands

cmd1 = "cd /home/xiaoayong/work/worklog/jinli"
cmd2 = "mkdir `date -d 'yesterday' +'%Y%m%d'` "
cmd3 = "cd /home/xiaoayong/work/zheng_shihui"
cmd4 = "kinit -kt /etc/krb5.keytab ecom_public@DAKAQUAN.COM"
cmd5 = "hdfs dfs -get hdfs://nameservice1/songshu/track/zheng_shihui/`date -d 'yesterday' +'%Y%m%d'` "

cmd = cmd1 + " && " + cmd2  + " && " + cmd3 + " && " + cmd4  + " && " + cmd5

subprocess.call(cmd,shell=True)
