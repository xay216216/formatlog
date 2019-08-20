#!/bin/bash
export LANG="en_US.UTF-8"
date_url=`date -d "yesterday" +'%Y%m%d'`
log_path=/home/xiaoayong/work/jinli_qingdianshang/${date_url}/*
domain="mail.qq.com"
email="843150233@qq.com"
maketime=`date +%Y-%m-%d" "%H":"%M`
logdate=`date -d "yesterday" +%Y-%m-%d`

# date --date='1 days ago' +%Y%m%d
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

echo -e "概况\n报告生成时间：${maketime}\n日期:${logdate}\n电商平台:全部\n\t灭屏拉活电商PV:${MplhPv}\n\n灭屏拉活UV\n${MplhUv}\n\n弹窗显示PV\n${TcxsPv}\n\n弹窗显示UV\n${TcxsUv}n\n激活点击PV\n${JhdjPv}\n\n激活点击UV\n${JhdjUv}\n\n激活率\n${Jhdjlv}\n\n弹窗意外关闭PV\n${TcywgbPv}\n\n弹窗意外关闭UV\n${TcywgbUv}\n\n意外关闭率\n${Tcywgblv}\n\n下次提醒点击PV\n${XctxdjPv}\n\n 下次提醒点击UV\n${XctxdjUv}\n\n关闭率\n${Xctxdjlv}\n\n浮标点击PV\n${FbdjPv}\n\n浮标点击UV\n${FbdjUv}\n\n人均浮标点击\n${Fbdjlv}\n\n小弹窗激活PV\n${XtcjhPv}\n\n小弹窗激活UV\n${XtcjhUv}\n\n浮标激活率\n${Xtcjhlv}\n\n小弹窗关闭PV\n${XtcgbPv}\n\n小弹窗关闭UV\n${XtcgbUv}\n\n小窗关闭率\n${Xtcgblv}\n\n电商平台:淘宝\n\n灭屏拉活电商PV:${MplhPvTB}\n\n灭屏拉活UV\n${MplhUvTB}\n\n弹窗显示PV\n${TcxsPvTB}\n\n弹窗显示UV\n${TcxsUvTB}\n\n激活点击PV\n${JhdjPvTB}\n\n激活点击UV\n${JhdjUvTB}\n\n激活率\n${JhdjlvTB}\n\n弹窗意外关闭PV\n${TcywgbPvTB}\n\n弹窗意外关闭UV\n${TcywgbUvTB}\n\n意外关闭率\n${TcywgblvTB}\n\n下次提醒点击PV\n${XctxdjPvTB}\n\n下次提醒点击UV\n${XctxdjUvTB}\n\n关闭率\n${XctxdjlvTB}\n\n浮标点击PV\n${FbdjPvTB}\n\n浮标点击UV\n${FbdjUvTB}\n\n人均浮标点击\n${FbdjlvTB}\n\n小弹窗激活PV\n${XtcjhPvTB}\n\n小弹窗激活UV\n${XtcjhUvTB}\n\n浮标激活率\n${XtcjhlvTB}\n\n小弹窗关闭PV\n${XtcgbPvTB}\n\n小弹窗关闭UV\n${XtcgbUvTB}\n\n小窗关闭率\n${XtcgblvTB}\n\n电商平台:平多多\n\n灭屏拉活电商PV:${MplhPvPDD}\n\n灭屏拉活UV\n${MplhUvPDD}\n\n弹窗显示PV\n${TcxsPvPDD}\n\n弹窗显示UV\n${TcxsUvPDD}\n\n激活点击PV\n${JhdjPvPDD}\n\n激活点击UV\n${JhdjUvPDD}\n\n激活率\n${JhdjlvPDD}\n\n弹窗意外关闭PV\n${TcywgbPvPDD}\n\n弹窗意外关闭UV\n${TcywgbUvPDD}\n\n意外关闭率\n${TcywgblvPDD}\n\n下次提醒点击PV\n${XctxdjPvPDD}\n\n下次提醒点击UV\n${XctxdjUvPDD}\n\n关闭率\n${XctxdjlvPDD}\n\n浮标点击PV\n${FbdjPvPDD}\n\n浮标点击UV\n${FbdjUvPDD}\n\n人均浮标点击\n${FbdjlvPDD}\n\n小弹窗激活PV\n${XtcjhPvPDD}\n\n小弹窗激活UV\n${XtcjhUvPDD}\n\n浮标激活率\n${XtcjhlvPDD}\n\n小弹窗关闭PV\n${XtcgbPvPDD}\n\n小弹窗关闭UV\n${XtcgbUvPDD}\n\n小窗关闭率\n${XtcgblvPDD}" | mail -s "$domain $logdate log statistics" ${email}　

##报告三
echo  -e "报告生成时间：  ${maketime}" > /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "日期：  ${logdate}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "电商平台：  全部" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "灭屏拉活电商PV：  ${MplhPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "灭屏拉活UV：  ${MplhUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗显示PV：  ${TcxsPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗显示UV：  ${TcxsUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活点击PV：  ${JhdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活点击UV：  ${JhdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活率：  ${Jhdjlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗意外关闭PV：  ${TcywgbPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗意外关闭UV：  ${TcywgbUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "意外关闭率：  ${Tcywgblv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "下次提醒点击PV：  ${XctxdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "下次提醒点击UV：  ${XctxdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "关闭率：  ${Xctxdjlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标点击PV：  ${FbdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标点击UV：  ${FbdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "人均浮标点击：  ${Fbdjlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗激活PV：  ${XtcjhPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗激活UV：  ${XtcjhUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标激活率：  ${Xtcjhlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗关闭PV：  ${XtcgbPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗关闭UV：  ${XtcgbUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小窗关闭率：  ${Xtcgblv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "电商平台：  淘宝" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "灭屏拉活电商PV：  ${MplhPvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "灭屏拉活UV：  ${MplhUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗显示PV：  ${TcxsPvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗显示UV：  ${TcxsUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活点击PV：  ${JhdjPvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活点击UV：  ${JhdjUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活率：  ${JhdjlvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗意外关闭PV：  ${TcywgbPvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗意外关闭UV：  ${TcywgbUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "意外关闭率：  ${TcywgblvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "下次提醒点击PV：  ${XctxdjPvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "下次提醒点击UV：  ${XctxdjUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "关闭率：  ${XctxdjlvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标点击PV：  ${FbdjPvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标点击UV：  ${FbdjUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "人均浮标点击：  ${FbdjlvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗激活PV：  ${XtcjhPvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗激活UV：  ${XtcjhUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标激活率：  ${XtcjhlvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗关闭PV：  ${XtcgbPvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗关闭UV：  ${XtcgbUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小窗关闭率：  ${XtcgblvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "电商平台：  平多多" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "灭屏拉活电商PV：  ${MplhPvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "灭屏拉活UV：  ${MplhUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗显示PV：  ${TcxsPvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗显示UV：  ${TcxsUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活点击PV：  ${JhdjPvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活点击UV：  ${JhdjUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "激活率：  ${JhdjlvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗意外关闭PV：  ${TcywgbPvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "弹窗意外关闭UV：  ${TcywgbUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "意外关闭率：  ${TcywgblvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "下次提醒点击PV：  ${XctxdjPvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "下次提醒点击UV：  ${XctxdjUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "关闭率：  ${XctxdjlvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标点击PV：  ${FbdjPvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标点击UV：  ${FbdjUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "人均浮标点击：  ${FbdjlvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗激活PV：  ${XtcjhPvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗激活UV：  ${XtcjhUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "浮标激活率：  ${XtcjhlvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗关闭PV：  ${XtcgbPvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小弹窗关闭UV：  ${XtcgbUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv
echo  -e "小窗关闭率：  ${XtcgblvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlithree.csv





##报告四
echo  -e "报告生成时间：  ${maketime}" > /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "日期：  ${logdate}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "电商平台：  全部" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "帮购启动PV：  ${BgqdPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "帮购启动UV：  ${BgqdUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "弹窗显示PV：  ${TcxsPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "弹窗显示UV：  ${TcxsUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "弹窗激活PV：  ${TcjhPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "弹窗激活UV：  ${TcjhUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "激活率：  ${Tcjhlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "弹窗到后台PV：  ${TcdhtPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "弹窗到后台Uv：  ${TcdhtUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "弹窗到后台率：  ${Tcdhtlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "下次提醒点击PV：  ${XctxdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "下次提醒点击UV：  ${XctxdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "关闭率：  ${Xctxdjlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "底部tab展示PV：  ${DbtabzsPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "底部tab展示UV：  ${DbtabzsUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "tab展示率：  ${Dbtabzslv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "底部tab点击PV：  ${DbtabdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "底部tab点击UV：  ${DbtabdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "tab点击率：  ${Dbtabdjlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "助手页开启点击PV：  ${ZsykqdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "助手页开启点击UV：  ${ZsykqdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "助手页开启点击率：  ${Zsykqdjlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv
echo  -e "助手页开启率：  ${Zsykqlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlifour.csv

##报告一
echo  -e "报告生成时间：  ${maketime}" > /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "日期：  ${logdate}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "电商平台：  全部" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "详情页搜索PV：  ${XqyssPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "详情页搜索UV：  ${XqyssUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "人均详情次数：  ${Xqysslv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "精准商品PV：  ${JzspPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "精准商品点击PV：  ${JzspdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "精准商品点击率：  ${Jzspdjlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "相似商品PV：  ${SsspPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "相似商品点击PV：  ${SsspdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "相似商品点击率   ：  ${Ssspdjlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "无结果PV：  ${WjgPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "结果总点击PV：  ${JgzdjPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "下单PV：  ${XdPv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "下单率：  ${Xdlv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "精准商品UV：  ${JzspUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "精准商品点击UV：  ${JzspdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "相似商品UV：  ${SsspUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
echo  -e "相似商品点击UV：  ${SsspdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.csv
