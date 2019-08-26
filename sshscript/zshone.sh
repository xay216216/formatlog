#!/bin/bash
export LANG="en_US.UTF-8"
log_path=/home/xiaoayong/work/zheng_shihui/`date -d "yesterday" +'%Y%m%d'`/*
domain="mail.qq.com"
email="843150233@qq.com"
maketime=`date +%Y-%m-%d" "%H":"%M`
logdate=`date -d "yesterday" +%Y-%m-%d`
echo ${log_path}
#app
AppQiDong=` cat ${log_path} | grep '"nm":"USER_START_APK_EVERYTIME"' | wc -l  `

#商品浏览总坑位数----buzhun
GoodsScan=` cat ${log_path} | grep -E ' "nm":"FIRST_PAGE_WHEN_SCAN_POINT" | "nm":"ZHUANQU_WHEN_SCAN_POINT" ' | wc -l  `

#商品卡片点击量
GoodsClick=` cat ${log_path} | grep '"nm":"EVENT_GOODS_DETAIL_PV"' | grep -E ' "from":"category" | "from":"zq" ' | wc -l  `

#商品卡片点击用户数
GoodsDetail=` cat ${log_path} | grep '"nm":"EVENT_GOODS_DETAIL_PV"' | grep -E ' "from":"category" | "from":"zq" ' | uniq | sort -t ',' -k 3,3 -u | wc -l  `

#热词使用量
HotCode=` cat ${log_path} | grep '"nm":"EVENT_SEARCH_PV"' | grep -E ' "from":"hot_search" | "from":"3" | "from":"category" | "from":"6" ' | wc -l  `

#关键词使用量
KeyCode=` cat ${log_path} | grep '"nm":"EVENT_SEARCH_PV"' | grep -E ' "from":"search_input" | "from":"7"  ' | wc -l  `

#搜索结果点击量-----buzhun
SearchResult=` cat ${log_path} | grep -E ' "nm":"EVENT_GOODS_DETAIL_PV" | "nm":"EVENT_JUMP_GET_COUPON"  ' | wc -l  `

#热词使用用户数
HotUse=` cat ${log_path} | grep '"nm":"EVENT_SEARCH_PV"' | grep -E ' "from":"hot_search" | "from":"3" | "from":"category" | "from":"6" ' | uniq | sort -t ',' -k 3,3 -u | wc -l  `

#关键词使用用户数
KeyUse=` cat ${log_path} | grep '"nm":"EVENT_SEARCH_PV"' | grep -E ' "from":"search_input" | "from":"7"  ' | uniq | sort -t ',' -k 3,3 -u | wc -l  `

#搜索结果点击用户数-----buzhun
SearchResultClick=` cat ${log_path} | grep -E ' "nm":"EVENT_GOODS_DETAIL_PV" | "nm":"EVENT_JUMP_GET_COUPON"  ' | uniq | sort -t ',' -k 3,3 -u | wc -l  `

#电商点击量
ShopClick=` cat ${log_path} | grep '"nm":"EVENT_JUMP_GET_COUPON"' | grep -E ' "clickType":"1" | "clickType":"2"  ' | wc -l  `

#电商点击用户数
ShopClickUser=` cat ${log_path} | grep '"nm":"EVENT_JUMP_GET_COUPON"' | grep -E ' "clickType":"1" | "clickType":"2"  ' | uniq | sort -t ',' -k 3,3 -u | wc -l  `

#详情页人均订单量----buzhun
DetailOrder=` cat ${log_path} | grep '"nm":"EVENT_USER_ORDER"' | wc -l  `

#详情页用户数
DetailUser=` cat ${log_path} | grep '"nm":"EVENT_GOODS_DETAIL_PV"' | wc -l  `

test=$[$AppQiDong / $HotCode]  
echo ${test}

echo  -e "报告生成时间：${maketime}" > /home/xiaoayong/work/worklog/zhengshihuireportone.csv
echo  -e "\t日期:${logdate}" >> /home/xiaoayong/work/worklog/zhengshihuireportone.csv
echo  -e "\tAPP启动量:${AppQiDong}" >> /home/xiaoayong/work/worklog/zhengshihuireportone.csv
echo  -e "\t商品卡片有效展示量UV:${GoodsScan}" >> /home/xiaoayong/work/worklog/zhengshihuireportone.csv
echo  -e "\t商品test:${test}" >> /home/xiaoayong/work/worklog/zhengshihuireportone.csv





