#!/bin/bash

#收件箱
EMAIL_RECIVER="843150233@qq.com"
#抄送
EMAIL_CC="843150231@qq.com"
#发送者邮箱
EMAIL_SENDER=xiaoayong@baice100.com
#邮箱用户名
EMAIL_USERNAME=xiaoayong@baice100.com
#邮箱密码
#使用qq邮箱进行发送需要注意：首先需要开启：POP3/SMTP服务，其次发送邮件的密码需要使用在开启POP3/SMTP服务时候腾讯提供的第三方客户端登陆码。
EMAIL_PASSWORD=Bckj@8888

#附件路径
FILE1_PATH="/home/xiaoayong/work/worklog/jinli/`date -d "yesterday" +'%Y%m%d'`/jinlione.excel /home/xiaoayong/work/worklog/jinli/`date -d "yesterday" +'%Y%m%d'`/jinlithree.excel /home/xiaoayong/work/worklog/jinli/`date -d "yesterday" +'%Y%m%d'`/jinlifour.excel /home/xiaoayong/work/worklog/zhengshihui/`date -d "yesterday" +'%Y%m%d'`/reportthree.xls"

#smtp服务器地址
EMAIL_SMTPHOST=smtp.exmail.qq.com
EMAIL_PORT=465

#EMAIL_SMTPHOST=smtp.qq.com
#EMAIL_PORT=587

EMAIL_TITLE="报表"
EMAIL_CONTENT="谢谢!"

sendEmail -f ${EMAIL_SENDER} -t ${EMAIL_RECIVER} -cc ${EMAIL_CC} -s ${EMAIL_SMTPHOST}  -u ${EMAIL_TITLE} -xu ${EMAIL_USERNAME} -xp ${EMAIL_PASSWORD} -m ${EMAIL_CONTENT} -a ${FILE1_PATH} -o message-charset=utf-8 -o tls=no












##说明   yum install -y sendemail  https://blog.csdn.net/boonya/article/details/89457436  yum install perl-Net-SSLeay perl-IO-Socket-SSL -y
