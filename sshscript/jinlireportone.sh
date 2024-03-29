#!/bin/bash
export LANG="en_US.UTF-8"
date_url=`date -d "yesterday" +'%Y%m%d'`
log_path=/home/xiaoayong/work/jinli_qingdianshang/${date_url}/*
domain="mail.qq.com"
email="843150233@qq.com"
maketime=`date +%Y-%m-%d" "%H":"%M`
logdate=`date -d "yesterday" +%Y-%m-%d`
#详情页搜索PV
XqyssPv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_DETAIL_SEARCH"' | wc -l  `
XqyssPvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_DETAIL_SEARCH"' | grep '"eplatform":"1"' | wc -l  `
XqyssPvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_DETAIL_SEARCH"' | grep '"eplatform":"2"' | wc -l  `
XqyssUv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_DETAIL_SEARCH"' | uniq | sort -t ',' -k 3,3 -u | wc -l `
XqyssUvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_DETAIL_SEARCH"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l `
XqyssUvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_DETAIL_SEARCH"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l `
Xqysslv=$(printf "%.3f" `echo "scale=3;$XqyssPv/$XqyssUv"|bc`)
XqysslvTB=$(printf "%.3f" `echo "scale=3;$XqyssPvTB/$XqyssUvTB"|bc`)
XqysslvPDD=$(printf "%.3f" `echo "scale=3;$XqyssPvPDD/$XqyssUvPDD"|bc`)

#精准商品PV
JzspPv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT"' | grep '"match":"1"' | wc -l  `
JzspPvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT"' | grep '"eplatform":"1"' | wc -l  `
JzspPvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT"' | grep '"eplatform":"2"' | wc -l  `
JzspUv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT"' | grep '"match":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
JzspUvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
JzspUvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `

#精准商品点击PV
JzspdjPv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK"' | grep '"match":"1"' | wc -l  `
JzspdjPvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK"' | grep '"eplatform":"1"' | wc -l  `
JzspdjPvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK"' | grep '"eplatform":"2"' | wc -l  `
JzspdjUv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK"' | grep '"match":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
JzspdjUvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
JzspdjUvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Jzspdjlv=$(printf "%.3f" `echo "scale=3;$JzspdjPv/$JzspPv"|bc`)
JzspdjlvTB=$(printf "%.3f" `echo "scale=3;$JzspdjPvTB/$JzspPvTB"|bc`)
JzspdjlvPDD=$(printf "%.3f" `echo "scale=3;$JzspdjPvPDD/$JzspPvPDD"|bc`)

#相似商品PV
SsspPv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT"' | grep '"match":"2"' | wc -l  `
SsspPvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT"' | grep '"eplatform":"1"' | wc -l  `
SsspPvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT"' | grep '"eplatform":"2"' | wc -l  `
SsspUv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT"' | grep '"match":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
SsspUvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
SsspUvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `


#相似商品点击PV
SsspdjPv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK"' | grep '"match":"2"' | wc -l  `
SsspdjPvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK"' | grep '"eplatform":"1"' | wc -l  `
SsspdjPvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK"' | grep '"eplatform":"2"' | wc -l  `
SsspdjUv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK"' | grep '"match":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
SsspdjUvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK"' | grep '"eplatform":"1"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
SsspdjUvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK"' | grep '"eplatform":"2"' | uniq | sort -t ',' -k 3,3 -u | wc -l  `
Ssspdjlv=$(printf "%.3f" `echo "scale=3;$SsspdjPv/$SsspPv"|bc`)
SsspdjlvTB=$(printf "%.3f" `echo "scale=3;$SsspdjPvTB/$SsspPvTB"|bc`)
SsspdjlvPDD=$(printf "%.3f" `echo "scale=3;$SsspdjPvPDD/$SsspPvPDD"|bc`)

#无结果PV
WjgPv=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_NO_RESULT"' | wc -l  `
WjgPvTB=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_NO_RESULT"' | grep '"eplatform":"1"' | wc -l  `
WjgPvPDD=` cat ${log_path} | grep '"nm":"SQG_TAOBAO_SEARCH_NO_RESULT"' | grep '"eplatform":"2"' | wc -l  `

#结果总点击PV
JgzdjPv=` cat ${log_path} | grep -E ' "nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK" | "nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK" '  | wc -l  `
JgzdjPvTB=` cat ${log_path} | grep -E ' "nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK" | "nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK" ' | grep '"eplatform":"1"' | wc -l  `
JgzdjPvPDD=` cat ${log_path} | grep -E ' "nm":"SQG_TAOBAO_SEARCH_ACCURATE_RESULT_CLICK" | "nm":"SQG_TAOBAO_SEARCH_SIMILAR_RESULT_CLICK" ' | grep '"eplatform":"2"' | wc -l  `


#下单PV
XdPv=0
XdPvTB=0
XdPvPDD=0
Xdlv=0
XdlvTB=0
XdlvPDD=0



echo  -e "报告生成时间\t日期\t电商平台\t详情页搜索PV\t详情页搜索UV\t人均详情次数\t精准商品PV\t精准商品点击PV\t精准商品点击率\t相似商品PV\t相似商品点击PV\t相似商品点击率\t无结果PV\t结果总点击PV\t下单PV\t下单率\t精准商品UV\t精准商品点击UV\t相似商品UV\t相似商品点击UV" > /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.xls

echo  -e "${maketime}\t${logdate}\t全部\t${XqyssPv}\t${XqyssUv}\t${Xqysslv}\t${JzspPv}\t${JzspdjPv}\t${Jzspdjlv}\t${SsspPv}\t${SsspdjPv}\t${Ssspdjlv}\t${WjgPv}\t${JgzdjPv}\t${XdPv}\t${Xdlv}\t${JzspUv}\t${JzspdjUv}\t${SsspUv}\t${SsspdjUv}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.xls

echo  -e "${maketime}\t${logdate}\t淘宝\t${XqyssPvTB}\t${XqyssUvTB}\t${XqysslvTB}\t${JzspPvTB}\t${JzspdjPvTB}\t${JzspdjlvTB}\t${SsspPvTB}\t${SsspdjPvTB}\t${SsspdjlvTB}\t${WjgPvTB}\t${JgzdjPvTB}\t${XdPvTB}\t${XdlvTB}\t${JzspUvTB}\t${JzspdjUvTB}\t${SsspUvTB}\t${SsspdjUvTB}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.xls

echo  -e "${maketime}\t${logdate}\t拼多多\t${XqyssPvPDD}\t${XqyssUvPDD}\t${XqysslvPDD}\t${JzspPvPDD}\t${JzspdjPvPDD}\t${JzspdjlvPDD}\t${SsspPvPDD}\t${SsspdjPvPDD}\t${SsspdjlvPDD}\t${WjgPvPDD}\t${JgzdjPvPDD}\t${XdPvPDD}\t${XdlvPDD}\t${JzspUvPDD}\t${JzspdjUvPDD}\t${SsspUvPDD}\t${SsspdjUvPDD}" >> /home/xiaoayong/work/worklog/jinli/${date_url}/jinlione.xls
