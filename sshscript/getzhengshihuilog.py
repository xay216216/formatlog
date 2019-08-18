#!/usr/bin/python
import sys
reload(sys)
sys.setdefaultencoding('utf8')

import subprocess
import os
import commands

cmd1 = "cd /home/xiaoayong/work/zheng_shihui"
cmd2 = "kinit -kt /etc/krb5.keytab ecom_public@DAKAQUAN.COM"
cmd3 = "hdfs dfs -get hdfs://nameservice1/songshu/track/zheng_shihui/`date -d 'yesterday' +'%Y%m%d'` "

cmd4 = "cd /home/xiaoayong/work/jinli_qingdianshang"
cmd5 = "hdfs dfs -get hdfs://nameservice1/songshu/track/jinli_qingdianshang/`date -d 'yesterday' +'%Y%m%d'`"


cmd = cmd1 + " && " + cmd2  + " && " + cmd3 + " && " + cmd4  + " && " + cmd5

subprocess.call(cmd,shell=True)
