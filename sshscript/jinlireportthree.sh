#!/bin/bash
export LANG="en_US.UTF-8"
date_url=`date -d "yesterday" +'%Y%m%d'`
log_path=/home/xiaoayong/work/jinli_qingdianshang/${date_url}/*
domain="mail.qq.com"
email="843150233@qq.com"
maketime=`date +%Y-%m-%d" "%H":"%M`
logdate=`date -d "yesterday" +%Y-%m-%d`
#灭屏拉活
MplhPv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN"' | wc -l  `
MplhPvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN"' | grep '"eplatform":"1"' | wc -l  `
MplhPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN"' | grep '"eplatform":"2"' | wc -l  `
MplhUv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN"' | uniq | sort -t ',' -k 3,3 -u | wc -l `
MplhUvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l `
MplhUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_OFF_SCREEN"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l `

#弹窗显示
TcxsPv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_SHOW"' | wc -l  `
TcxsPvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_SHOW"' | grep '"eplatform":"1"' | wc -l  `
TcxsPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_SHOW"' | grep '"eplatform":"2"' | wc -l  `
TcxsUv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_SHOW"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
TcxsUvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_SHOW"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
TcxsUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_SHOW"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `

#激活点击
JhdjPv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK"' | wc -l  `
JhdjPvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK"' | grep '"eplatform":"1"' | wc -l  `
JhdjPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK"' | grep '"eplatform":"2"' | wc -l  `
JhdjUv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
JhdjUvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
JhdjUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_PERMISSION_AGREE_CLICK"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Jhdjlv=$(printf "%.3f" `echo "scale=3;$JhdjUv/$TcxsUv"|bc`)
JhdjlvTB=$(printf "%.3f" `echo "scale=3;$JhdjUvTB/$TcxsUvTB"|bc`)
JhdjlvPDD=$(printf "%.3f" `echo "scale=3;$JhdjUvPDD/$TcxsUvPDD"|bc`)
#弹窗意外关闭
TcywgbPv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE"' | wc -l  `
TcywgbPvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE"' | grep '"eplatform":"1"' | wc -l  `
TcywgbPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE"' | grep '"eplatform":"2"' | wc -l  `
TcywgbUv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
TcywgbUvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
TcywgbUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_ACCIDENT_CLOSE"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Tcywgblv=$(printf "%.3f" `echo "scale=3;$TcywgbUv/$TcxsUv"|bc`)
TcywgblvTB=$(printf "%.3f" `echo "scale=3;$TcywgbUvTB/$TcxsUvTB"|bc`)
TcywgblvPDD=$(printf "%.3f" `echo "scale=3;$TcywgbUvPDD/$TcxsUvPDD"|bc`)

#下次提醒点击
XctxdjPv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | wc -l  `
XctxdjPvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | grep '"eplatform":"1"' | wc -l  `
XctxdjPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | grep '"eplatform":"2"' | wc -l  `
XctxdjUv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
XctxdjUvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
XctxdjUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_HELPER_DIALOG_NEXT_REMINDER_CLOSE"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Xctxdjlv=$(printf "%.3f" `echo "scale=3;$XctxdjUv/$TcxsUv"|bc`)
XctxdjlvTB=$(printf "%.3f" `echo "scale=3;$XctxdjUvTB/$TcxsUvTB"|bc`)
XctxdjlvPDD=$(printf "%.3f" `echo "scale=3;$XctxdjUvPDD/$TcxsUvPDD"|bc`)

#浮标点击
FbdjPv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK"' | wc -l  `
FbdjPvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK"' | grep '"eplatform":"1"' | wc -l  `
FbdjPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK"' | grep '"eplatform":"2"' | wc -l  `
FbdjUv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
FbdjUvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
FbdjUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_SMALL_SHAKE_BIG_COUPON_OPEN_CLICK"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Fbdjlv=$(printf "%.3f" `echo "scale=3;$FbdjPv/$XctxdjUv"|bc`)
FbdjlvTB=$(printf "%.3f" `echo "scale=3;$FbdjPvTB/$XctxdjUvTB"|bc`)
FbdjlvPDD=$(printf "%.3f" `echo "scale=3;$FbdjPvPDD/$XctxdjUvPDD"|bc`)

#小弹窗激活
XtcjhPv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK"' | wc -l  `
XtcjhPvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK"' | grep '"eplatform":"1"' | wc -l  `
XtcjhPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK"' | grep '"eplatform":"2"' | wc -l  `
XtcjhUv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
XtcjhUvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
XtcjhUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_OPEN_CLICK"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Xtcjhlv=$(printf "%.3f" `echo "scale=3;$XtcjhUv/$FbdjUv"|bc`)
XtcjhlvTB=$(printf "%.3f" `echo "scale=3;$XtcjhUvTB/$FbdjUvTB"|bc`)
XtcjhlvPDD=$(printf "%.3f" `echo "scale=3;$XtcjhUvPDD/$FbdjUvPDD"|bc`)

#小弹窗关闭
XtcgbPv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE"' | wc -l  `
XtcgbPvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE"' | grep '"eplatform":"1"' | wc -l  `
XtcgbPvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE"' | grep '"eplatform":"2"' | wc -l  `
XtcgbUv=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
XtcgbUvTB=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
XtcgbUvPDD=` cat ${log_path} | grep '"nm":"GIONEE_TAOBAO_FROM_SHAKE_BIG_COUPON_OPEN_HELPER_DIALOG_CUSTOM_CLOSE"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Xtcgblv=$(printf "%.3f" `echo "scale=3;$XtcgbUv/$FbdjUv"|bc`)
XtcgblvTB=$(printf "%.3f" `echo "scale=3;$XtcgbUvTB/$FbdjUvTB"|bc`)
XtcgblvPDD=$(printf "%.3f" `echo "scale=3;$XtcgbUvPDD/$FbdjUvPDD"|bc`)

echo -e "报告生成时间\t日期\t电商平台\t灭屏拉活电商PV\t灭屏拉活UV\t弹窗显示PV\t弹窗显示UV\t激活点击PV\t激活点击UV\t激活率\t弹窗意外关闭PV\t弹窗意外关闭UV\t意外关闭率\t下次提醒点击PV\t下次提醒点击UV\t关闭率\t浮标点击PV\t浮标点击UV\t人均浮标点击\t小弹窗激活PV\t小弹窗激活UV\t浮标激活率\t小弹窗关闭PV\t小弹窗关闭UV\t小窗关闭率" > /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.xls

echo -e "${maketime}\t${logdate}\t全部\t${MplhPv}\t${MplhUv}\t${TcxsPv}\t${TcxsUv}\t${JhdjPv}\t${JhdjUv}\t${Jhdjlv}\t${TcywgbPv}\t${TcywgbUv}\t${Tcywgblv}\t${XctxdjPv}\t${XctxdjUv}\t${Xctxdjlv}\t${FbdjPv}\t${FbdjUv}\t${Fbdjlv}\t${XtcjhPv}\t${XtcjhUv}\t${Xtcjhlv}\t${XtcgbPv}\t${XtcgbUv}\t${Xtcgblv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.xls

echo -e "${maketime}\t${logdate}\t淘宝\t${MplhPvTB}\t${MplhUvTB}\t${TcxsPvTB}\t${TcxsUvTB}\t${JhdjPvTB}\t${JhdjUvTB}\t${JhdjlvTB}\t${TcywgbPvTB}\t${TcywgbUvTB}\t${TcywgblvTB}\t${XctxdjPvTB}\t${XctxdjUvTB}\t${XctxdjlvTB}\t${FbdjPvTB}\t${FbdjUvTB}\t${FbdjlvTB}\t${XtcjhPvTB}\t${XtcjhUvTB}\t${XtcjhlvTB}\t${XtcgbPvTB}\t${XtcgbUvTB}\t${XtcgblvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.xls

echo -e "${maketime}\t${logdate}\t拼多多\t${MplhPvPDD}\t${MplhUvPDD}\t${TcxsPvPDD}\t${TcxsUvPDD}\t${JhdjPvPDD}\t${JhdjUvPDD}\t${JhdjlvPDD}\t${TcywgbPvPDD}\t${TcywgbUvPDD}\t${TcywgblvPDD}\t${XctxdjPvPDD}\t${XctxdjUvPDD}\t${XctxdjlvPDD}\t${FbdjPvPDD}\t${FbdjUvPDD}\t${FbdjlvPDD}\t${XtcjhPvPDD}\t${XtcjhUvPDD}\t${XtcjhlvPDD}\t${XtcgbPvPDD}\t${XtcgbUvPDD}\t${XtcgblvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.xls


