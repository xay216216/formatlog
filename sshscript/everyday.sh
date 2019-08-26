#！/usr/bin/bash

#.  /etc/profile
source /etc/profile      
source ~/.bash_profile
export LANG="en_US.UTF-8"
#PATH=/etc:/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin
#zhengshihui
cd /home/xiaoayong/work/worklog/zhengshihui
rm -rf *
mkdir `date -d 'yesterday' +'%Y%m%d'`
cd /home/xiaoayong/work/zheng_shihui
kinit -kt /etc/krb5.keytab ecom_public@DAKAQUAN.COM
hdfs dfs -get hdfs://nameservice1/songshu/track/zheng_shihui/`date -d 'yesterday' +'%Y%m%d'`
#jinli
cd /home/xiaoayong/work/worklog/jinli
rm -rf *
mkdir `date -d 'yesterday' +'%Y%m%d'`
cd /home/xiaoayong/work/jinli_qingdianshang
kinit -kt /etc/krb5.keytab ecom_public@DAKAQUAN.COM
hdfs dfs -get hdfs://nameservice1/songshu/track/jinli_qingdianshang/`date -d 'yesterday' +'%Y%m%d'`

##source那句很重要  才能执行hdfs命令