#!/usr/bin/env python
#-*- coding: UTF-8 -*-
#import os,sys
#reload(sys)
import sys
import importlib
importlib.reload(sys)
#sys.setdefaultencoding('utf8')   //python3 没有了
import getopt
import smtplib
from email.mime.text import MIMEText
from email.utils import formataddr
from  subprocess import *
def sendqqmail(username,password,mailfrom,mailto,subject,content):
    # 邮箱的服务地址
    gserver = 'smtp.qq.com'
    gport = 587
    try:
        msg = MIMEText(unicode(content).encode('utf-8'))
        msg['from'] = mailfrom
        msg['to'] = mailto
        msg['Reply-To'] = mailfrom
        msg['Subject'] = subject
        smtp = smtplib.SMTP(gserver, gport)
        smtp.connect(gserver,gport)
        smtp.set_debuglevel(0)
        smtp.ehlo()
        smtp.login(username,password)
        smtp.sendmail(mailfrom, mailto, msg.as_string())
        smtp.close()
    except Exception as err:
        print ('Exception: ', err)
        ##print "Send mail failed. Error: %s" % err
        
def main():
    to=sys.argv[1]
    subject=sys.argv[2]
    content=sys.argv[3]
    #定义QQ邮箱的账号和密码，你需要修改成你自己的账号和密码
    sendqqmail('843150233@qq.com','xay216216qhmj','843150233@qq.com',to,subject,content)
if __name__ == "__main__":
    main()

#####脚本使用说明######
#1. 首先定义好脚本中的邮箱账号和密码
#2. 脚本执行命令为：python mail.py 目标邮箱 "邮件主题" "邮件内容"




