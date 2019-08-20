# -*- coding: utf-8 -*-

import datetime
import os
from spark_core import SparkCore
import sparktool
from sparktool import log, config
import warnings

warnings.filterwarnings("ignore")
'''
信息流统计
'''

def mkSql(sql_str):
    global keyDict
    for k in keyDict:
        rp = "{%s}" % k
        f = sql_str.find(rp)
        if f != -1:
            sql_str = sql_str.replace(rp,keyDict[k])
    return sql_str

def mkSqlBat(objArr):
    sqlArr = []
    for s in objArr:
        sql = mkSql(s)
        sqlArr.append(sql)
    return sqlArr

def load_data(spark, day, path):
    begin = day - datetime.timedelta(days=6)
    end = day
    date = begin
    dateArr = []
    while date <= end:
        filedate = date.strftime("%Y%m%d") + "/IS_TRACK_*"
        dateArr.append(filedate)
        date += datetime.timedelta(days=1)
    dateStr = ",".join(dateArr)
    fileTrackPath = "{0}{{{1}}}".format(path, dateStr)
    spark.load_parquet_table(fileTrackPath, "temp1")

keyDict = {
    "pid":"if(pid is not null, pid, '')",
    "dv":"if(dv is not null, dv, '')",
    "action":"if(action is not null, action, '')",
    "event":"if(event is not null, event, '')",
    "app_vname":"if(app_vname is not null, app_vname, '')",
}

selectCommon = (
    "count(if(t2.nm=1,m1,NULL)) show_pv,count(distinct(if(t2.nm=1,m1,NULL))) show_uv,count(if(t2.nm=2,m1,NULL)) click_pv,count(distinct(if(t2.nm=2,m1,NULL))) click_uv,count(if(t2.nm=3,m1,NULL)) begin_download_pv,count(distinct(if(t2.nm=3,m1,NULL))) begin_download_uv,count(if(t2.nm=4,m1,NULL)) finish_download_pv,count(distinct(if(t2.nm=4,m1,NULL))) finish_download_uv,count(if(t2.nm=5,m1,NULL)) finish_install_pv,count(distinct(if(t2.nm=5,m1,NULL))) finish_install_uv"
)
sqlDict = {
    'is_user': [
        "select distinct {0}, if(pid is not null, pid, '') pid, uid uid, now() update_time from req",
        "select {0}, 'month' time, if(pid is not null, pid, '') pid, count(distinct(uid)) month_active_user from month_req group by if(pid is not null, pid, '')"
    ],
    'is_req': [
        "select {0}, if(pid is not null, pid, '') pid, if(realSource is not null, realSource, '') real_source, count(distinct(bid)) pv, count(distinct(uid)) uv from req group by if(pid is not null, pid, ''), if(realSource is not null, realSource, '')",
        "select {0}, if(realSource is not null, realSource, '') real_source, count(distinct(bid)) pv, count(distinct(uid)) uv from req group by if(realSource is not null, realSource, '')",
        "select {0}, if(t1.pid is not null, t1.pid, '') pid, if(t1.exid is not null, t1.exid, '') expid, count(distinct(t1.bid)) pv, count(distinct(t1.uid)) uv,%s from (select pid, uid, bid, explode(split(expId, ',')) exid from (select distinct pid, expId, uid, bid from req) temp) t1 left join (select distinct nm, seg.bid bid, seg.adId adid, m1 from hy ) t2 on t1.bid = t2.bid  group by if(t1.pid is not null, t1.pid, ''), if(t1.exid is not null, t1.exid, '')" % selectCommon,
        "select {0}, if(t1.pid is not null, t1.pid, '') pid, count(distinct(t1.bid)) pv, count(distinct(t1.uid)) uv, %s from (select distinct pid, uid, bid from req) t1 left join (select distinct nm, seg.bid bid, seg.adId adid, m1 from hy ) t2 on t1.bid = t2.bid  group by if(t1.pid is not null, t1.pid, '')" % selectCommon,
        "select {0}, if(pid is not null, pid, '') pid, if(app_vname is not null, app_vname, '') app_vname, count(distinct(bid)) pv, count(distinct(uid)) uv from req group by if(pid is not null, pid, ''), if(app_vname is not null, app_vname, '')",
        "select {0}, '#All' pid, count(distinct(bid)) pv, count(distinct(uid)) uv from req",
    ],
    'is_req_cost': [
        "select {0}, if(realSource is not null, realSource, '') real_source, count(if(cSize = 0, cSize, null)) con_fail_pv, count(if(adSize = 0, adSize, null)) ad_noback_pv, sum(if(cCost is not null, cCost, 0)) con_cost, avg(if(cCost is not null, cCost, 0)) con_avg_cost, sum(if(adCost is not null, adCost, 0)) ad_cost, avg(if(adCost is not null, adCost, 0)) ad_avg_cost, count(if(cCost > 500, cCost, null)) `req>500ms`, count(if(adCost > 1000, adCost, null)) `req>1s` from req where adOnly = '0' group by if(realSource is not null, realSource, '')",
    ],
    'is_content': [
        "select {0},  if(pid is not null, pid, '') pid,  count(if(event = '34', event, null)) news_show_pv,  count(distinct(if(event = '34', uuid, null))) news_show_uv,  count(if(event = '1', event, null)) news_click_pv,  count(distinct(if(event = '1', uuid, null))) news_click_uv,  count(if(event = '32', event, null)) active_pv,  count(distinct(if(event = '32', uuid, null))) active_uv,  count(distinct(if(event = '34' and full_screen = '1', uuid, null))) fs_expose_uv,  count(distinct(if(event = '34' and full_screen = '0', uuid, null))) bs_expose_uv,  count(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and totalNum > 0, event, null)) con_re_req_pv,  count(distinct(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and totalNum > 0, uuid, null))) con_re_req_uv,  count(if((event = '19' or event = '24' or event = '25') and totalNum > 0, event, null)) con_auto_req_pv, count(distinct(if((event = '19' or event = '24' or event = '25') and totalNum > 0, uuid, null))) con_auto_req_uv, count(if(event = '28', event, null)) ad_show_pv, count(distinct(if(event = '28', uuid, null))) ad_show_uv, count(if(event = '27', event, null)) ad_click_pv, count(distinct(if(event = '27', uuid, null))) ad_click_uv, count(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and adNum > 0, event, null)) ad_re_req_pv, count(distinct(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and adNum > 0, uuid, null))) ad_re_req_uv,  count(if((event = '19' or event = '24' or event = '25') and adNum > 0, event, null)) ad_auto_req_pv,  count(distinct(if((event = '19' or event = '24' or event = '25') and adNum > 0, uuid, null))) ad_auto_req_uv,  count(distinct(if( (event = '34' and full_screen = '1') or event = '1', uuid, null))) fs_expose_or_click_uv,  sum(if(event = '9' or event = '10' or event = '11' or event = '19' or event = '20' or event = '21' or event = '22' or event = '23' or event = '24', totalNum, 0)) pull_numbers,  count(distinct(if(event = '9' or event = '10' or event = '11' or event = '19' or event = '20' or event = '21' or event = '22' or event = '23' or event = '24', uuid, null))) pull_uv  from track  group by if(pid is not null, pid, '')",
        "select {0}, if(pid is not null, pid, '') pid, if(expid is not null, expid, '') expid, count(if(event = '34', event, null)) news_show_pv, count(distinct(if(event = '34', uuid, null))) news_show_uv, count(if(event = '1', event, null)) news_click_pv, count(distinct(if(event = '1', uuid, null))) news_click_uv, count(if(event = '32', event, null)) active_pv, count(distinct(if(event = '32', uuid, null))) active_uv, count(distinct(if(event = '34' and full_screen = '1', uuid, null))) fs_expose_uv, count(distinct(if(event = '34' and full_screen = '0', uuid, null))) bs_expose_uv, count(distinct(if( (event = '34' and full_screen = '1') or event = '1', uuid, null))) fs_expose_or_click_uv, sum(if(event = '9' or event = '10' or event = '11' or event = '19' or event = '20' or event = '21' or event = '22' or event = '23' or event = '24', totalNum, 0)) pull_numbers, count(distinct(if(event = '9' or event = '10' or event = '11' or event = '19' or event = '20' or event = '21' or event = '22' or event = '23' or event = '24', uuid, null))) pull_uv from (select pid, uuid, event, full_screen, totalNum, explode(split(exp_ids, ',')) expid from track) t  group by if(pid is not null, pid, ''), if(expid is not null, expid, '')",
        "select {0}, if(pid is not null, pid, '') pid,  if(app_vname is not null, app_vname, '') app_vname, count(if(event = '34', event, null)) news_show_pv,  count(distinct(if(event = '34', uuid, null))) news_show_uv,  count(if(event = '1', event, null)) news_click_pv,  count(distinct(if(event = '1', uuid, null))) news_click_uv,  count(if(event = '32', event, null)) active_pv,  count(distinct(if(event = '32', uuid, null))) active_uv,  count(distinct(if(event = '34' and full_screen = '1', uuid, null))) fs_expose_uv,  count(distinct(if(event = '34' and full_screen = '0', uuid, null))) bs_expose_uv,  count(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and totalNum > 0, event, null)) con_re_req_pv,  count(distinct(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and totalNum > 0, uuid, null))) con_re_req_uv,  count(if((event = '19' or event = '24' or event = '25') and totalNum > 0, event, null)) con_auto_req_pv, count(distinct(if((event = '19' or event = '24' or event = '25') and totalNum > 0, uuid, null))) con_auto_req_uv, count(if(event = '28', event, null)) ad_show_pv, count(distinct(if(event = '28', uuid, null))) ad_show_uv, count(if(event = '27', event, null)) ad_click_pv, count(distinct(if(event = '27', uuid, null))) ad_click_uv, count(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and adNum > 0, event, null)) ad_re_req_pv, count(distinct(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and adNum > 0, uuid, null))) ad_re_req_uv,  count(if((event = '19' or event = '24' or event = '25') and adNum > 0, event, null)) ad_auto_req_pv,  count(distinct(if((event = '19' or event = '24' or event = '25') and adNum > 0, uuid, null))) ad_auto_req_uv,  count(distinct(if( (event = '34' and full_screen = '1') or event = '1', uuid, null))) fs_expose_or_click_uv,  sum(if(event = '9' or event = '10' or event = '11' or event = '19' or event = '20' or event = '21' or event = '22' or event = '23' or event = '24', totalNum, 0)) pull_numbers,  count(distinct(if(event = '9' or event = '10' or event = '11' or event = '19' or event = '20' or event = '21' or event = '22' or event = '23' or event = '24', uuid, null))) pull_uv  from track group by if(pid is not null, pid, ''), if(app_vname is not null, app_vname, '')",
        "select {0}, '#All' pid, count(if(event = '34', event, null)) news_show_pv,  count(distinct(if(event = '34', uuid, null))) news_show_uv,  count(if(event = '1', event, null)) news_click_pv,  count(distinct(if(event = '1', uuid, null))) news_click_uv,  count(if(event = '32', event, null)) active_pv,  count(distinct(if(event = '32', uuid, null))) active_uv,  count(distinct(if(event = '34' and full_screen = '1', uuid, null))) fs_expose_uv,  count(distinct(if(event = '34' and full_screen = '0', uuid, null))) bs_expose_uv,  count(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and totalNum > 0, event, null)) con_re_req_pv,  count(distinct(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and totalNum > 0, uuid, null))) con_re_req_uv,  count(if((event = '19' or event = '24' or event = '25') and totalNum > 0, event, null)) con_auto_req_pv, count(distinct(if((event = '19' or event = '24' or event = '25') and totalNum > 0, uuid, null))) con_auto_req_uv, count(if(event = '28', event, null)) ad_show_pv, count(distinct(if(event = '28', uuid, null))) ad_show_uv, count(if(event = '27', event, null)) ad_click_pv, count(distinct(if(event = '27', uuid, null))) ad_click_uv, count(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and adNum > 0, event, null)) ad_re_req_pv, count(distinct(if((event = '9' or event = '10' or event = '11' or event = '20' or event = '21' or event = '22' or event = '23' or event = '26' or event = '33') and adNum > 0, uuid, null))) ad_re_req_uv,  count(if((event = '19' or event = '24' or event = '25') and adNum > 0, event, null)) ad_auto_req_pv,  count(distinct(if((event = '19' or event = '24' or event = '25') and adNum > 0, uuid, null))) ad_auto_req_uv,  count(distinct(if( (event = '34' and full_screen = '1') or event = '1', uuid, null))) fs_expose_or_click_uv,  sum(if(event = '9' or event = '10' or event = '11' or event = '19' or event = '20' or event = '21' or event = '22' or event = '23' or event = '24', totalNum, 0)) pull_numbers,  count(distinct(if(event = '9' or event = '10' or event = '11' or event = '19' or event = '20' or event = '21' or event = '22' or event = '23' or event = '24', uuid, null))) pull_uv  from track"
    ],
    'is_track': mkSqlBat([
        #pid,app_vname,dv,event,action
        "select {0}, {pid} pid, {dv} dv, {action} action, {event} event, {app_vname} app_vname, count(*) pv, count(distinct(uuid)) uv from track group by {pid}, {dv}, {action}, {event}, {app_vname}",
        #pid,event,action
        "select {0}, {pid} pid, {action} action, {event} event, count(*) pv, count(distinct(uuid)) uv from track group by {pid}, {action}, {event}",
        #pid,app_vname,event,action
        "select {0}, {pid} pid, {action} action, {event} event, {app_vname} app_vname, count(*) pv, count(distinct(uuid)) uv from track group by {pid}, {action}, {event}, {app_vname}",
        #pid,dv,event,action
        "select {0}, {pid} pid, {dv} dv, {action} action, {event} event, count(*) pv, count(distinct(uuid)) uv from track group by {pid}, {dv}, {action}, {event}",
        #pid,app_vname
        "select {0}, {pid} pid, {app_vname} app_vname, count(*) pv, count(distinct(uuid)) uv from track group by {pid}, {app_vname}",
        #pid,dv
        "select {0}, {pid} pid, {dv} dv, count(*) pv, count(distinct(uuid)) uv from track group by {pid}, {dv}",
        #pid,dv,event
        "select {0}, {pid} pid, {dv} dv, {event} event, count(*) pv, count(distinct(uuid)) uv from track group by {pid}, {dv}, {event}",
        #pid,app_vname,event
        "select {0}, {pid} pid, {event} event, {app_vname} app_vname, count(*) pv, count(distinct(uuid)) uv from track group by {pid}, {event}, {app_vname}",
    ]),
    "is_retain": [
        "select '{0}' date, 'fs_expose' nm , if(t1.pid is not null, t1.pid, '') pid, 'active' retain_type, '{1}' retain_days, count(t1.uuid) today_uv, count(t2.uuid) retain_uv from (select pid, uuid from temp where event = '34' and full_screen = '1' group by pid, uuid) t1 left join (select pid, uuid from track where event = '34' and full_screen = '1' group by pid, uuid) t2 on t1.pid = t2.pid and t1.uuid = t2.uuid group by if(t1.pid is not null, t1.pid, '')",
        "select '{0}' date, 'fs_expose' nm , if(t1.pid is not null, t1.pid, '') pid, if(t1.expid is not null, t1.expid, '') expid, 'active' retain_type, '{1}' retain_days, count(t1.uuid) today_uv, count(t2.uuid) retain_uv from (select pid, expid, uuid from (select pid, uuid, event, full_screen, explode(split(exp_ids, ',')) expid from temp) temp where event = '34' and full_screen = '1' group by pid, uuid, expid) t1 left join (select pid, expid, uuid from (select pid, uuid, event, full_screen, explode(split(exp_ids, ',')) expid from track) track where event = '34' and full_screen = '1' group by pid, uuid, expid) t2 on t1.pid = t2.pid and t1.expid = t2.expid and t1.uuid = t2.uuid group by if(t1.pid is not null, t1.pid, ''), if(t1.expid is not null, t1.expid, '')",
        "select '{0}' date, 'active_uv' nm , if(t1.pid is not null, t1.pid, '') pid, 'active' retain_type, '{1}' retain_days, count(t1.uuid) today_uv, count(t2.uuid) retain_uv from (select pid, uuid from temp where event = '32' group by pid, uuid) t1 left join (select pid, uuid from track where event = '32' group by pid, uuid ) t2 on t1.pid = t2.pid and t1.uuid = t2.uuid group by if(t1.pid is not null, t1.pid, '')",
        "select '{0}' date, 'active_uv' nm , if(t1.pid is not null, t1.pid, '') pid, if(t1.expid is not null, t1.expid, '') expid, 'active' retain_type, '{1}' retain_days, count(t1.uuid) today_uv, count(t2.uuid) retain_uv from (select pid, expid, uuid from (select pid, uuid, explode(split(exp_ids, ',')) expid from temp where event = '32' ) temp group by pid, uuid, expid) t1 left join (select pid, expid, uuid from (select pid, uuid, explode(split(exp_ids, ',')) expid from track where event = '32') track group by pid, uuid, expid) t2 on t1.pid = t2.pid and t1.expid = t2.expid and t1.uuid = t2.uuid group by if(t1.pid is not null, t1.pid, ''), if(t1.expid is not null, t1.expid, '')",
    ]
}

parquetPath = config['hdfs_data_path'] + "is/"
hyParquetPath = config['hdfs_data_path'] + "hy/"
opt = sparktool.get_opt()
if opt['env'] == 'ali':
    parquetPath = config['oss_parquet_path'] + "is/"
    hyParquetPath = config['oss_parquet_path'] + "hy/"
log("begin spark. opt: {0}".format(opt))
date = opt['begin']
hourtype = opt['hourtype']
type = "is"
while date < opt['end']:
    if not hourtype and not sparktool.check_hdfs_file_count("{0}{1}".format(parquetPath, date.strftime("%Y%m%d")), "is"):
        date += datetime.timedelta(days=1)
        continue
    filedate = date.strftime("%Y%m%d")
    filehour = date.strftime("%Y%m%d%H")
    spark = SparkCore()
    if not hourtype:
        fileReqParquet = "{0}{1}/IS_SERVER_*".format(parquetPath, filedate)
        fileTrackParquet = "{0}{1}/IS_TRACK_*".format(parquetPath, filedate)
        fileHyTrackParquet = "{0}{1}/HY_TRACK_*".format(hyParquetPath, filedate)
    else:
        fileReqParquet = "{0}{1}/IS_SERVER_{2}.parquet".format(parquetPath,filedate,filehour)
        fileTrackParquet = "{0}{1}/IS_TRACK_{2}.parquet".format(parquetPath,filedate,filehour)
        fileHyTrackParquet = "{0}{1}/HY_TRACK_{2}.parquet".format(hyParquetPath, filedate,filehour)
    spark.load_parquet_table(fileReqParquet, "req")
    spark.load_parquet_table(fileTrackParquet, "track")
    spark.load_parquet_table(fileHyTrackParquet, "hy")
    if len(opt['tables']) > 0:
        keys = opt['tables']
    else:
        if hourtype:
            keys = ['is_req:1.2.3.4', 'is_req_cost', 'is_content:0.1']
        else:
            keys = ['is_req:0.2.3.4.5', 'is_content', 'is_track', 'is_retain']
    if 'is_user' in keys and date.strftime('%d') == '01':
        monthStr = (date + datetime.timedelta(days=-1)).strftime("%Y%m")
        fileMonthReq = "{0}{1}*/IS_SERVER_{1}*".format(parquetPath, monthStr)
        spark.load_parquet_table(fileMonthReq, "month_req")
    dateStr = date.strftime('%Y-%m-%d')
    hourStr = date.strftime("%H")
    for tb in keys:
        log("{0} begin".format(tb))
        tbArr = tb.split(":")
        tb = tbArr[0]
        index = []
        if len(tbArr) > 1:
            index = map(lambda x: int(x), tbArr[1].split("."))
        if opt["overwrite"]:
            if hourtype:
                sql = "delete from {0} where date='{1}' and hour = {2}".format(tb, dateStr, hourStr)
            else:
                sql = "delete from {0} where date='{1}' and hour is null".format(tb, dateStr)
            sparktool.execute_db_sql(sql, "commit")
        for i in range(0, len(sqlDict[tb])):
            if len(index) > 0 and i not in index:
                continue
            s = sqlDict[tb][i]
            if hourtype:
                sqlStr = "'{0}' date,'{1}' hour".format(dateStr, hourStr)
            else:
                sqlStr = "'{0}' date".format(dateStr)
            if tb == "is_user" and i == 1:
                if date.strftime('%d') == '01':
                    spark.exe_spark_sql_to_db(type, s.format(sqlStr), tb)
                    #monthUserDf = spark.execute_spark_sql(s.format(sqlStr))
                    #spark.write_db(monthUserDf, tb)
                continue
            if tb == 'is_retain':
                for j in [1, 7]:
                    lastdate = date - datetime.timedelta(days=j)
                    retaindate = lastdate.strftime("%Y-%m-%d")
                    fileTrackPath = parquetPath + lastdate.strftime("%Y%m%d") + "/IS_TRACK_*"
                    delIndex = ['fs_expose']
                    if opt["overwrite"]:
                        sql = "delete from is_retain where date = '{0}' and nm = '{1}' and retain_days = '{2}'".format(retaindate,delIndex[i],str(j)+"day")
                        sparktool.execute_db_sql(sql, "commit")
                    spark.load_parquet_table(fileTrackPath, "temp")
                    #tempSql = s
                    #if j == 7:
                    #    load_data(spark, date, parquetPath)
                    #    tempSql = tempSql.replace('track', 'temp1')
                    #res = spark.exe_spark_sql_to_db(type, s.format(retaindate, str(j) + "day"), tb,j,True)
                    retainDf = spark.execute_spark_sql(s.format(retaindate, str(j) + "day"))
                    spark.write_db(retainDf, tb, j)
                continue
            df = spark.execute_spark_sql(s.format(sqlStr))
            userPath = "/user/report/is/"
            if tb == "is_user" and i == 0:
                df = spark.execute_spark_sql(s.format(sqlStr))
                allUserParquet = userPath + "new_user.parquet"
                spark.create_table_by_df(df, "newuser")
                if date == datetime.datetime.strptime("20190509", '%Y%m%d'):
                    df.write.mode('append').parquet(allUserParquet)
                else:
                    spark.load_parquet_table(allUserParquet, "alluser")
                    spark.execute_spark_sql("select n.date, n.pid, n.uid, n.update_time from newuser n left join alluser a on n.uid = a.uid and n.pid = a.pid where a.uid is null").write.mode('append').parquet(allUserParquet)
                spark.load_parquet_table(allUserParquet, "alluser")
                spark.exe_spark_sql_to_db(type, "select {0}, pid pid, count(*) day_new_user from alluser where date = '{1}' group by pid".format(sqlStr, date.strftime("%Y-%m-%d")), tb)
                #newUserDf = spark.execute_spark_sql("select {0}, pid pid, count(*) day_new_user from alluser where date = '{1}' group by pid".format(sqlStr, date.strftime("%Y-%m-%d")))
                #spark.write_db(newUserDf, tb)
                # merge file per month
                if date.strftime('%d') == '01':
                    df = spark.spark.read.parquet(allUserParquet)
                    df.write.parquet(userPath + "new_user_tmp.parquet")
                    os.system(config['hdfs_bin_path'] + " dfs -rm -r -f {0}".format(allUserParquet))
                    os.system(config['hdfs_bin_path'] + " dfs -mv {0}new_user_tmp.parquet {0}new_user.parquet".format(userPath))
                continue
            #res = spark.exe_spark_sql_to_db(type, s.format(sqlStr), tb)
            spark.write_db(df, tb)
        log("{0} done".format(tb))
    if hourtype:
        date += datetime.timedelta(hours=1)
    else:
        date += datetime.timedelta(days=1)

log("end spark. opt: {0}".format(opt))