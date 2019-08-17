#!/bin/bash
export LANG="en_US.UTF-8"
log_path=/home/xiaoayong/work/jinli_qingdianshang/`date -d "yesterday" +'%Y%m%d'`/*
domain="mail.qq.com"
email="843150233@qq.com"
maketime=`date +%Y-%m-%d" "%H":"%M`
logdate=`date -d "yesterday" +%Y-%m-%d`
#帮购启动PV
BgqdPv=` cat ${log_path} | grep '"nm":"GIONEE_OPEN_BANGO"' | wc -l  `
##BgqdPvTB=` cat ${log_path} | grep '"nm":"GIONEE_OPEN_BANGO"' | grep '"eplatform":"1"' | wc -l  `
##BgqdPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_OPEN_BANGO"' | grep '"eplatform":"2"' | wc -l  `
BgqdUv=` cat ${log_path} | grep '"nm":"GIONEE_OPEN_BANGO"' | uniq | sort -t ',' -k 3,3 -u | wc -l `
##BgqdUvTB=` cat ${log_path} | grep '"nm":"GIONEE_OPEN_BANGO"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l `
##BgqdUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_OPEN_BANGO"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l `

#弹窗显示PV
TcxsPv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_SHOW"' | wc -l  `
##TcxsPvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_SHOW"' | grep '"eplatform":"1"' | wc -l  `
##TcxsPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_SHOW"' | grep '"eplatform":"2"' | wc -l  `
TcxsUv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_SHOW"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##TcxsUvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_SHOW"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##TcxsUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_SHOW"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `

#弹窗激活PV
TcjhPv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK"' | wc -l  `
##TcjhPvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK"' | grep '"eplatform":"1"' | wc -l  `
##TcjhPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK"' | grep '"eplatform":"2"' | wc -l  `
TcjhUv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##TcjhUvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##TcjhUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_ACTIVATE_CLICK"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Tcjhlv=$[$TcjhUv/$TcxsUv]
##TcjhlvTB=$[$TcjhUvTB/$TcxsUvTB]
##TcjhlvPDD=$[$TcjhUvPDD/$TcxsUvPDD]

#弹窗到后台
TcdhtPv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND"' | wc -l  `
##TcdhtPvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND"' | grep '"eplatform":"1"' | wc -l  `
##TcdhtPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND"' | grep '"eplatform":"2"' | wc -l  `
TcdhtUv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##TcdhtUvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##TcdhtUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_COUSTOM_BACKGROUND"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Tcdhtlv=$[$TcdhtUv/$TcxsUv]
##TcdhtlvTB=$[$TcdhtUvTB/$TcxsUvTB]
##TcdhtlvPDD=$[$TcdhtUvPDD/$TcxsUvPDD]

#下次提醒点击
XctxdjPv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | wc -l  `
##XctxdjPvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | grep '"eplatform":"1"' | wc -l  `
##XctxdjPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | grep '"eplatform":"2"' | wc -l  `
XctxdjUv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##XctxdjUvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##XctxdjUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Xctxdjlv=$[$XctxdjUv/$TcxsUv]
##XctxdjlvTB=$[$XctxdjUvTB/$TcxsUvTB]
##XctxdjlvPDD=$[$XctxdjUvPDD/$TcxsUvPDD]

#底部tab展示PV
DbtabzsPv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_SHOW"' | wc -l  `
##DbtabzsPvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_SHOW"' | grep '"eplatform":"1"' | wc -l  `
##DbtabzsPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_SHOW"' | grep '"eplatform":"2"' | wc -l  `
DbtabzsUv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_SHOW"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##DbtabzsUvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_SHOW"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##DbtabzsUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_SHOW"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Jianfalv=$[$TcxsPv-$TcdhtUv]
##JianfalvTB=$[$TcxsPvTB-$TcdhtUvTB]
##JianfalvPDD=$[$TcxsPvPDD-$TcdhtUvPDD]
Dbtabzslv=$[$DbtabzsPv/$JianfaUv]
##DbtabzslvTB=$[$DbtabzsPvTB/$JianfaUvTB]
##DbtabzslvPDD=$[$DbtabzsPvPDD/$JianfaUvPDD]

#底部tab点击PV
DbtabdjPv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_CLICK"' | wc -l  `
##DbtabdjPvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_CLICK"' | grep '"eplatform":"1"' | wc -l  `
##DbtabdjPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_CLICK"' | grep '"eplatform":"2"' | wc -l  `
DbtabdjUv=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_CLICK"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##DbtabdjUvTB=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_CLICK"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##DbtabdjUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_BANGO_HELPER_TAB_CLICK"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Dbtabdjlv=$[$DbtabdjUv/$DbtabzsUv]
##DbtabdjlvTB=$[$DbtabdjUvTB/$DbtabzsUvTB]
##DbtabdjlvPDD=$[$DbtabdjUvPDD/$DbtabzsUvPDD]

#助手页开启点击PV
ZsykqdjPv=` cat ${log_path} | grep '"nm":"SQG_HELPER_SWITCH_PAGE_OPEN_BTN"' | wc -l  `
##ZsykqdjPvTB=` cat ${log_path} | grep '"nm":"SQG_HELPER_SWITCH_PAGE_OPEN_BTN"' | grep '"eplatform":"1"' | wc -l  `
##ZsykqdjPvPDD=` cat ${log_path} | grep '"nm":"SQG_HELPER_SWITCH_PAGE_OPEN_BTN"' | grep '"eplatform":"2"' | wc -l  `
ZsykqdjUv=` cat ${log_path} | grep '"nm":"SQG_HELPER_SWITCH_PAGE_OPEN_BTN"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##ZsykqdjUvTB=` cat ${log_path} | grep '"nm":"SQG_HELPER_SWITCH_PAGE_OPEN_BTN"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
##ZsykqdjUvPDD=` cat ${log_path} | grep '"nm":"SQG_HELPER_SWITCH_PAGE_OPEN_BTN"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Zsykqdjlv=$[$ZsykqdjUv/$DbtabdjUv]
##ZsykqdjlvTB=$[$ZsykqdjUvTB/$DbtabdjUvTB]
##ZsykqdjlvPDD=$[$ZsykqdjUvPDD/$DbtabdjUvPDD]
Zsykqlv=$[$ZsykqdjUv/$FbdjUv]
##ZsykqlvTB=$[$ZsykqdjUvTB/$FbdjUvTB]
##ZsykqlvPDD=$[$ZsykqdjUvPDD/$FbdjUvPDD]


echo  -e "报告生成时间：  ${maketime}" > /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "日期：  ${logdate}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "电商平台：  全部" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "帮购启动PV：  ${BgqdPv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "帮购启动UV：  ${BgqdUv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "弹窗显示PV：  ${TcxsPv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "弹窗显示UV：  ${TcxsUv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "弹窗激活PV：  ${TcjhPv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "弹窗激活UV：  ${TcjhUv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "激活率：  ${Tcjhlv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "弹窗到后台PV：  ${TcdhtPv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "弹窗到后台Uv：  ${TcdhtUv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "弹窗到后台率：  ${Tcdhtlv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "下次提醒点击PV：  ${XctxdjPv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "下次提醒点击UV：  ${XctxdjUv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "关闭率：  ${Xctxdjlv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "底部tab展示PV：  ${DbtabzsPv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "底部tab展示UV：  ${DbtabzsUv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "tab展示率：  ${Dbtabzslv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "底部tab点击PV：  ${DbtabdjPv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "底部tab点击UV：  ${DbtabdjUv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "tab点击率：  ${Dbtabdjlv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "助手页开启点击PV：  ${ZsykqdjPv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "助手页开启点击UV：  ${ZsykqdjUv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "助手页开启点击率：  ${Zsykqdjlv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv
echo  -e "助手页开启率：  ${Zsykqlv}" >> /home/xiaoayong/work/worklog/`date -d "yesterday" +'%Y%m%d'`/jinlifour.csv




