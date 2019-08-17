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



