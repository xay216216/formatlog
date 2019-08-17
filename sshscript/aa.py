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
cmd2 = "hdfs dfs -get hdfs://nameservice1/songshu/track/zheng_shihui/`date -d 'yesterday' +'%Y%m%d'` "
cmd = cmd1 + " && " + cmd2

subprocess.call(cmd,shell=True)


0 14 * * * python /home/xiaoayong/work/script/gethdfslog.py

* * * * * bash /home/xiaoayong/work/script/zhengshihuireportone.sh

10 * * * * bash /home/xiaoayong/work/script/jinlireportthree.sh

20 * * * * bash /home/xiaoayong/work/script/jinlireportfour.sh



5 14 * * * python /home/xiaoayong/work/script/getzhengshihuilog.py

0 14 * * * python /home/xiaoayong/work/script/getjinlilog.py

10 15 * * * bash /home/xiaoayong/work/script/jinlireportthree.sh

20 15 * * * bash /home/xiaoayong/work/script/jinlireportfour.sh





