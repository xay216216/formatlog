package script

import (
	"bufio"
	_ "encoding/json"
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type M1uid struct {
	M1  string
	UID string
	SEX string
}

type Uid struct {
	UID  string
	SEXY int
}

var (
	countLogFileName = "countlog.xls"
	dateUrlA         = ""
	logData          map[string]M1uid
	aa               string
	uidArray         = map[string]string{
		"1": "1", "3": "1", "4": "1", "2061": "1", "7": "0", "8": "1", "9": "1", "11": "1", "2683": "1", "14": "1", "2913": "1", "17": "0", "19": "1", "21": "0", "23": "1", "2832": "1", "26": "1", "27": "2", "28": "1", "29": "0", "30": "1", "32": "0", "33": "2", "34": "1", "36": "2", "37": "1", "39": "1", "40": "2", "41": "0", "2852": "1", "43": "1", "44": "0", "45": "0", "46": "2", "47": "2", "48": "2", "49": "1", "50": "0", "51": "0", "52": "0", "53": "1", "54": "1", "55": "0", "56": "0", "59": "1", "61": "1", "63": "2", "64": "2", "65": "2", "66": "0", "68": "1", "70": "2", "71": "2", "72": "2", "74": "2", "76": "2", "77": "2", "78": "1", "79": "0", "80": "1", "81": "0", "82": "2", "83": "2", "84": "1", "85": "0", "86": "1", "87": "0", "89": "2", "90": "2", "91": "0", "92": "2", "93": "2", "94": "2", "95": "1", "96": "2", "97": "1", "98": "1", "99": "1", "100": "0", "101": "2", "102": "1", "104": "0", "105": "1", "106": "2", "107": "1", "108": "2", "109": "0", "112": "1", "114": "1", "116": "0", "117": "1", "118": "2", "119": "1", "120": "1", "122": "1", "123": "2", "124": "2", "126": "1", "127": "1", "128": "2", "129": "1", "130": "2", "131": "0", "132": "1", "134": "0", "135": "1", "136": "1", "137": "2", "140": "1", "142": "0", "143": "0", "144": "2", "146": "0", "147": "1", "148": "2", "149": "1", "150": "1", "151": "0", "153": "0", "152": "1", "154": "1", "157": "1", "158": "2", "160": "1", "161": "2", "162": "1", "163": "0", "165": "1", "166": "0", "167": "2", "169": "2", "168": "1", "170": "2", "171": "1", "173": "1", "174": "0", "175": "1", "176": "1", "177": "0", "180": "1", "181": "1", "178": "1", "182": "2", "183": "1", "184": "1", "186": "0", "187": "0", "188": "0", "189": "2", "190": "2", "192": "2", "193": "0", "196": "0", "197": "1", "199": "0", "200": "1", "201": "1", "202": "1", "203": "1", "204": "2", "205": "1", "206": "0", "207": "2", "208": "2", "209": "2", "210": "2", "211": "1", "213": "0", "214": "1", "216": "2", "217": "0", "218": "2", "219": "2", "220": "1", "221": "1", "222": "2", "223": "1", "224": "1", "225": "1", "226": "1", "229": "2", "231": "2", "232": "0", "233": "2", "235": "2", "237": "1", "238": "2", "239": "0", "240": "0", "241": "1", "242": "1", "243": "2", "244": "0", "245": "0", "246": "1", "247": "1", "248": "1", "249": "2", "250": "0", "251": "1", "252": "1", "253": "0", "254": "2", "256": "2", "257": "1", "258": "2", "259": "0", "260": "1", "261": "2", "263": "0", "264": "0", "265": "1", "266": "2", "268": "0", "270": "1", "272": "0", "273": "1", "274": "2", "275": "0", "276": "0", "277": "1", "278": "0", "279": "1", "280": "0", "281": "1", "282": "2", "283": "1", "284": "1", "285": "0", "286": "2", "287": "1", "288": "2", "290": "1", "291": "1", "292": "2", "293": "1", "294": "0", "296": "2", "297": "1", "299": "1", "300": "1", "301": "1", "302": "0", "303": "1", "304": "2", "305": "2", "307": "2", "308": "2", "309": "1", "310": "1", "311": "1", "312": "2", "313": "2", "315": "2", "317": "1", "319": "1", "320": "2", "321": "2", "322": "0", "323": "1", "324": "0", "325": "1", "326": "1", "328": "1", "329": "0", "330": "1", "331": "2", "332": "0", "334": "1", "335": "1", "337": "0", "338": "1", "339": "1", "340": "2", "341": "1", "342": "0", "343": "0", "344": "1", "345": "0", "346": "1", "347": "0", "348": "0", "349": "2", "350": "2", "351": "1", "352": "1", "354": "0", "355": "1", "356": "1", "357": "1", "358": "1", "360": "1", "361": "0", "362": "1", "363": "2", "364": "1", "365": "2", "366": "0", "368": "1", "369": "2", "370": "0", "372": "1", "373": "0", "375": "0", "377": "1", "378": "1", "380": "0", "381": "1", "382": "0", "384": "1", "386": "1", "387": "0", "388": "2", "389": "0", "390": "0", "391": "2", "392": "2", "393": "1", "394": "1", "395": "2", "396": "1", "398": "0", "399": "1", "400": "2", "401": "0", "402": "0", "403": "2", "404": "1", "405": "2", "406": "0", "407": "1", "412": "1", "413": "1", "415": "0", "416": "1", "417": "1", "418": "1", "419": "0", "421": "1", "422": "2", "423": "1", "424": "1", "426": "0", "427": "1", "428": "1", "429": "2", "430": "0", "431": "1", "433": "2", "434": "0", "436": "1", "437": "1", "438": "2", "439": "2", "440": "1", "441": "1", "443": "1", "444": "1", "445": "0", "447": "2", "448": "2", "449": "0", "450": "2", "452": "1", "453": "1", "454": "0", "455": "0", "456": "1", "458": "2", "459": "1", "460": "2", "462": "2", "463": "1", "464": "0", "465": "2", "466": "0", "467": "2", "469": "0", "470": "1", "472": "1", "473": "1", "474": "2", "475": "2", "476": "1", "477": "0", "478": "1", "479": "1", "480": "2", "481": "2", "482": "1", "483": "1", "484": "1", "485": "1", "486": "2", "487": "0", "488": "1", "489": "1", "490": "2", "491": "1", "492": "1", "493": "0", "494": "0", "495": "1", "496": "2", "497": "1", "498": "1", "499": "2", "500": "1", "501": "1", "502": "1", "503": "2", "504": "1", "505": "2", "506": "2", "507": "1", "508": "0", "509": "0", "510": "2", "511": "2", "512": "2", "513": "1", "514": "1", "515": "1", "516": "0", "517": "2", "518": "2", "520": "0", "521": "1", "522": "1", "523": "1", "524": "2", "525": "2", "526": "1", "527": "1", "528": "1", "529": "2", "530": "1", "534": "1", "535": "0", "536": "0", "537": "1", "538": "1", "539": "1", "540": "2", "541": "1", "542": "0", "543": "2", "544": "1", "545": "1", "546": "1", "547": "1", "548": "2", "549": "0", "550": "2", "551": "1", "552": "2", "554": "2", "555": "1", "556": "1", "558": "1", "560": "2", "562": "0", "563": "2", "564": "0", "565": "0", "566": "1", "567": "2", "569": "2", "570": "1", "571": "2", "572": "2", "574": "0", "576": "2", "577": "2", "578": "0", "579": "2", "580": "1", "581": "1", "582": "2", "585": "0", "586": "1", "587": "0", "588": "1", "590": "1", "591": "1", "592": "2", "593": "1", "594": "2", "595": "2", "596": "0", "598": "2", "600": "1", "601": "1", "602": "0", "604": "2", "605": "0", "606": "0", "607": "0", "608": "1", "609": "1", "610": "1", "611": "1", "612": "2", "613": "0", "614": "2", "615": "1", "616": "1", "618": "0", "620": "1", "621": "0", "622": "1", "623": "1", "624": "1", "625": "2", "626": "1", "627": "0", "628": "2", "630": "2", "632": "1", "633": "2", "634": "2", "635": "2", "636": "2", "637": "1", "638": "0", "639": "0", "640": "1", "641": "1", "642": "1", "643": "1", "644": "2", "645": "0", "646": "1", "647": "0", "648": "1", "649": "1", "650": "1", "651": "1", "652": "2", "653": "2", "654": "1", "655": "1", "656": "0", "657": "0", "659": "1", "660": "0", "661": "2", "662": "1", "663": "2", "664": "1", "665": "2", "666": "1", "668": "0", "669": "1", "670": "1", "671": "2", "672": "0", "673": "1", "674": "1", "676": "2", "678": "2", "679": "0", "681": "2", "682": "2", "683": "1", "684": "2", "685": "0", "686": "0", "688": "0", "689": "1", "690": "1", "691": "0", "692": "1", "693": "1", "694": "2", "695": "2", "696": "2", "697": "2", "698": "0", "700": "1", "701": "2", "702": "1", "703": "0", "704": "0", "705": "1", "706": "1", "708": "1", "709": "0", "710": "1", "711": "0", "712": "1", "713": "1", "714": "2", "715": "1", "716": "2", "718": "2", "719": "1", "720": "2", "723": "1", "724": "1", "725": "1", "726": "1", "727": "2", "728": "1", "729": "0", "730": "0", "731": "2", "732": "0", "733": "1", "734": "1", "735": "1", "736": "0", "737": "1", "738": "2", "739": "0", "740": "1", "741": "2", "742": "1", "743": "1", "744": "2", "745": "2", "746": "1", "747": "0", "749": "2", "750": "1", "751": "1", "752": "1", "2943": "1", "754": "0", "755": "1", "756": "1", "757": "1", "758": "1", "759": "1", "760": "1", "761": "1", "762": "2", "763": "1", "764": "1", "766": "1", "767": "2", "768": "2", "769": "1", "770": "1", "772": "1", "773": "1", "775": "1", "776": "1", "779": "1", "780": "1", "781": "2", "782": "2", "783": "2", "784": "2", "785": "2", "786": "2", "787": "1", "788": "2", "789": "2", "790": "0", "791": "1", "792": "2", "793": "1", "794": "0", "795": "0", "796": "2", "797": "1", "798": "0", "799": "1", "800": "1", "801": "1", "802": "1", "803": "1", "804": "1", "805": "0", "806": "0", "808": "1", "809": "0", "810": "2", "811": "1", "812": "2", "813": "2", "814": "1", "815": "2", "817": "2", "818": "1", "819": "2", "820": "1", "821": "2", "824": "2", "825": "1", "826": "1", "827": "1", "828": "2", "829": "1", "830": "0", "831": "1", "833": "0", "834": "1", "835": "1", "836": "1", "837": "1", "838": "1", "839": "0", "840": "2", "841": "2", "843": "0", "844": "1", "845": "1", "846": "0", "847": "2", "848": "2", "849": "1", "850": "1", "851": "1", "854": "1", "855": "2", "856": "0", "857": "1", "859": "2", "861": "1", "863": "1", "864": "1", "865": "2", "866": "0", "867": "2", "868": "1", "869": "2", "870": "1", "871": "1", "873": "0", "874": "1", "875": "2", "876": "0", "877": "2", "878": "1", "879": "1", "880": "1", "881": "1", "882": "2", "883": "0", "884": "1", "885": "2", "886": "1", "888": "1", "889": "1", "890": "0", "891": "0", "892": "1", "893": "1", "894": "1", "895": "2", "896": "1", "899": "0", "902": "1", "904": "1", "905": "1", "906": "1", "907": "2", "908": "1", "910": "1", "911": "2", "913": "0", "914": "0", "917": "1", "918": "1", "919": "1", "920": "1", "921": "1", "922": "1", "923": "1", "925": "2", "926": "0", "927": "0", "928": "2", "929": "1", "930": "1", "931": "1", "933": "2", "934": "1", "936": "2", "938": "1", "939": "0", "940": "0", "942": "1", "943": "2", "944": "2", "945": "1", "946": "1", "947": "1", "948": "1", "949": "2", "950": "2", "951": "2", "952": "0", "953": "2", "954": "2", "955": "1", "957": "2", "958": "1", "959": "1", "961": "2", "962": "1", "963": "0", "964": "0", "965": "1", "966": "2", "967": "1", "970": "1", "971": "2", "972": "1", "973": "1", "974": "1", "975": "1", "977": "1", "978": "0", "979": "1", "980": "1", "982": "2", "983": "1", "984": "1", "985": "1", "987": "2", "990": "2", "991": "1", "992": "2", "993": "2", "995": "1", "996": "0", "997": "1", "998": "0", "999": "0", "1000": "0", "1001": "0", "1002": "1", "1003": "2", "1004": "2", "1005": "2", "1006": "1", "1007": "1", "1008": "2", "1009": "1", "1011": "1", "1012": "2", "1016": "0", "1014": "1", "1015": "0", "1017": "0", "1019": "2", "1020": "2", "1021": "2", "1022": "0", "1023": "1", "1025": "2", "1027": "1", "1029": "0", "1030": "1", "1031": "2", "1032": "1", "1033": "1", "1034": "1", "1035": "1", "1036": "1", "1037": "0", "1038": "2", "1039": "2", "1040": "1", "1042": "1", "1043": "0", "1044": "0", "1045": "1", "1046": "1", "1047": "2", "1048": "1", "1049": "0", "1050": "0", "1051": "1", "1052": "0", "1053": "2", "1054": "1", "1055": "0", "1056": "1", "1057": "0", "1058": "1", "1060": "0", "1061": "1", "1062": "2", "1063": "2", "1064": "1", "1065": "1", "1066": "1", "1067": "0", "1068": "1", "1069": "2", "1070": "1", "1071": "2", "1074": "0", "1075": "1", "1076": "1", "1077": "2", "1078": "1", "1079": "0", "1081": "1", "1082": "1", "1083": "0", "1084": "2", "1085": "1", "1087": "1", "1089": "1", "1090": "2", "1091": "1", "1092": "1", "1093": "2", "1094": "0", "1095": "2", "1096": "1", "1097": "0", "1098": "1", "1100": "2", "1102": "1", "1104": "2", "1106": "0", "1107": "0", "1108": "1", "1109": "2", "1110": "0", "1111": "1", "1112": "2", "1113": "2", "1115": "1", "1116": "2", "1117": "0", "1119": "1", "1120": "1", "1121": "1", "1122": "1", "1123": "2", "1124": "1", "1125": "1", "1127": "0", "1129": "1", "1130": "2", "1131": "1", "1132": "1", "1133": "2", "1134": "1", "1136": "2", "1138": "2", "1139": "1", "1140": "1", "1141": "0", "1142": "1", "1143": "2", "1144": "2", "1145": "1", "1146": "1", "1147": "1", "1148": "2", "1149": "1", "1150": "1", "1151": "1", "1152": "2", "1153": "1", "1154": "2", "1155": "2", "1156": "1", "1157": "2", "1158": "2", "1159": "2", "1160": "1", "1162": "1", "1163": "0", "1164": "1", "1166": "1", "1168": "2", "1169": "1", "1170": "2", "1171": "1", "1172": "2", "1173": "1", "1175": "1", "1176": "2", "1177": "2", "1178": "2", "1179": "1", "1180": "2", "1181": "1", "1182": "1", "1183": "2", "1184": "1", "1185": "1", "1186": "2", "1187": "1", "1188": "0", "1189": "0", "1191": "2", "1192": "1", "1194": "0", "1195": "2", "1196": "2", "1199": "1", "1200": "2", "1201": "1", "1202": "1", "1203": "0", "1204": "1", "1205": "0", "1206": "1", "1207": "2", "1208": "0", "1210": "1", "1211": "2", "1212": "1", "1215": "0", "1216": "2", "1217": "1", "1218": "1", "1219": "1", "1221": "0", "1222": "2", "1223": "1", "1224": "0", "1226": "2", "1227": "2", "1229": "2", "1230": "1", "1231": "0", "1232": "1", "1233": "2", "1234": "0", "1235": "2", "1236": "0", "1237": "1", "1238": "1", "1239": "1", "1240": "0", "1241": "1", "1243": "2", "1244": "1", "1245": "1", "1246": "2", "1247": "2", "1248": "2", "1249": "2", "1251": "2", "1252": "1", "1253": "0", "1254": "0", "1255": "0", "1256": "0", "1257": "1", "1258": "1", "1259": "2", "1260": "1", "1261": "1", "1263": "1", "1264": "2", "1265": "1", "1266": "1", "1267": "1", "1269": "0", "1270": "1", "1271": "2", "1272": "2", "1273": "0", "1274": "1", "1275": "1", "1276": "2", "1277": "1", "1278": "1", "1282": "1", "1283": "2", "1286": "1", "1288": "2", "1289": "1", "1290": "1", "1291": "0", "1292": "1", "1293": "2", "1294": "2", "1295": "0", "1296": "2", "1297": "2", "1298": "1", "1301": "1", "1302": "0", "1303": "1", "1304": "1", "1305": "2", "1306": "1", "1307": "1", "1308": "0", "1309": "2", "1311": "1", "1312": "2", "1313": "0", "1314": "0", "1315": "1", "1316": "2", "1317": "0", "1318": "1", "1319": "2", "1320": "2", "1322": "1", "1323": "1", "1324": "2", "1325": "1", "1326": "1", "1327": "0", "1328": "2", "1329": "1", "1330": "2", "1331": "0", "1332": "0", "1333": "0", "1336": "1", "1337": "0", "1338": "0", "1339": "0", "1340": "1", "1341": "0", "1342": "1", "1343": "0", "1346": "1", "1347": "1", "1348": "2", "1349": "2", "1350": "2", "1351": "2", "1352": "1", "1353": "2", "1354": "2", "1355": "0", "1356": "1", "1357": "0", "1358": "0", "1359": "1", "1360": "2", "1361": "2", "1362": "2", "1365": "1", "1368": "1", "1369": "2", "1370": "2", "1371": "1", "1372": "0", "1373": "1", "1374": "0", "1375": "1", "1376": "2", "1377": "0", "1379": "1", "1380": "1", "1381": "2", "1382": "1", "1383": "1", "1384": "2", "1385": "2", "1386": "1", "1387": "0", "1388": "1", "1390": "1", "1391": "2", "1392": "1", "1393": "2", "1394": "2", "1395": "1", "1396": "0", "1397": "1", "1398": "1", "1399": "1", "1400": "1", "1401": "0", "1402": "1", "1403": "1", "1404": "1", "1405": "1", "1406": "2", "1407": "1", "1409": "0", "1410": "1", "1411": "1", "1412": "1", "1413": "1", "1415": "2", "1416": "2", "1417": "1", "1418": "2", "1419": "0", "1421": "1", "1422": "2", "1423": "2", "1425": "2", "1426": "1", "1427": "1", "1428": "1", "1429": "2", "1433": "0", "1434": "0", "1435": "2", "1436": "1", "1437": "0", "1438": "1", "1439": "2", "1440": "1", "1441": "1", "1443": "1", "1444": "1", "1445": "0", "1446": "1", "1447": "0", "1449": "1", "1450": "0", "1451": "0", "1452": "0", "1453": "2", "1454": "0", "1455": "0", "1456": "0", "1457": "2", "1458": "2", "1459": "2", "1461": "2", "1462": "1", "1463": "1", "1464": "1", "1465": "2", "1466": "1", "1467": "2", "1468": "2", "1469": "1", "1470": "2", "1471": "0", "1473": "0", "1475": "2", "1476": "2", "1477": "0", "1478": "0", "1479": "0", "1480": "2", "1481": "2", "1482": "1", "1483": "2", "1484": "1", "1485": "0", "1486": "0", "1487": "2", "1489": "2", "1490": "2", "1491": "2", "1492": "0", "1493": "2", "1494": "1", "1497": "1", "1498": "2", "1499": "0", "1500": "2", "1502": "0", "1503": "1", "1504": "0", "1506": "0", "1507": "2", "1508": "1", "1509": "2", "1510": "0", "1511": "1", "1512": "2", "1514": "0", "1515": "2", "1516": "1", "1517": "2", "1518": "0", "1519": "1", "1520": "1", "1521": "2", "1522": "0", "1523": "2", "1525": "2", "1526": "2", "1527": "0", "1528": "2", "1529": "2", "1530": "1", "1531": "1", "1532": "0", "1533": "0", "1534": "2", "1535": "1", "1536": "2", "1537": "1", "1538": "2", "1539": "1", "1540": "1", "1541": "1", "1542": "0", "1543": "1", "1544": "2", "1545": "0", "1546": "1", "1547": "0", "1548": "1", "1549": "1", "1550": "0", "1552": "1", "1553": "2", "1554": "1", "1555": "1", "1556": "1", "1557": "2", "1558": "1", "1559": "2", "1560": "1", "1561": "1", "1565": "0", "1566": "2", "1563": "1", "1567": "0", "1568": "1", "1569": "0", "1570": "1", "1571": "2", "1572": "0", "1573": "0", "1574": "0", "1575": "1", "1576": "1", "1577": "0", "1578": "2", "1579": "2", "1580": "1", "1581": "1", "1582": "2", "1583": "1", "1584": "1", "1585": "2", "1586": "0", "1587": "1", "1588": "2", "1589": "2", "1590": "1", "1592": "0", "1593": "0", "1594": "2", "1595": "1", "1597": "2", "1598": "1", "1599": "1", "1600": "0", "1601": "2", "1602": "2", "1603": "1", "1604": "1", "1605": "1", "1606": "2", "1607": "2", "1608": "0", "1609": "0", "1610": "2", "1611": "2", "1613": "1", "1614": "1", "1615": "1", "1616": "1", "1617": "1", "1618": "2", "1619": "1", "1620": "1", "1622": "2", "1624": "1", "1625": "2", "1626": "2", "1628": "1", "1629": "1", "1630": "0", "1631": "1", "1632": "2", "1633": "2", "1634": "2", "1635": "0", "1636": "1", "1637": "1", "1638": "0", "1639": "1", "1640": "2", "1641": "2", "1642": "1", "1643": "1", "1644": "2", "1646": "1", "9165": "2", "1651": "1", "1652": "1", "1653": "2", "1654": "1", "1655": "1", "1657": "2", "1658": "2", "1659": "1", "1660": "1", "1661": "0", "1663": "2", "1664": "1", "1665": "1", "1666": "0", "1668": "2", "1669": "2", "1672": "1", "1673": "1", "1674": "2", "1675": "1", "1676": "1", "1678": "2", "1679": "1", "1680": "1", "374": "2", "1682": "1", "1684": "0", "1685": "0", "1687": "0", "1688": "2", "1689": "2", "1690": "0", "1691": "1", "1692": "2", "1693": "2", "1694": "0", "1695": "2", "1696": "1", "1697": "1", "1699": "1", "1700": "2", "1701": "1", "1703": "1", "1704": "0", "1705": "1", "1706": "1", "1707": "1", "1708": "1", "1711": "0", "1712": "0", "1713": "0", "1714": "2", "1715": "0", "1716": "2", "1717": "2", "1718": "1", "1719": "1", "1720": "0", "1721": "2", "1722": "2", "1723": "2", "1724": "1", "1725": "1", "1727": "1", "1729": "0", "1730": "1", "1731": "1", "1732": "1", "1733": "0", "1734": "1", "1735": "2", "1738": "2", "1740": "0", "1741": "0", "1742": "0", "1743": "1", "1744": "2", "1745": "1", "1746": "1", "1747": "1", "1748": "0", "1750": "1", "1751": "2", "1752": "1", "1753": "1", "1754": "2", "1755": "2", "1756": "1", "1757": "2", "1758": "1", "1759": "1", "1760": "0", "1761": "1", "1762": "1", "1763": "2", "1764": "2", "1765": "0", "1766": "1", "1767": "1", "1769": "0", "1770": "1", "1771": "0", "1772": "2", "1773": "1", "1774": "1", "1775": "1", "1776": "1", "1778": "0", "1780": "2", "1781": "0", "1782": "0", "1783": "1", "1784": "1", "1785": "2", "1786": "1", "1787": "1", "1788": "2", "1789": "1", "1790": "0", "1791": "0", "1792": "1", "1793": "2", "1794": "0", "1795": "0", "1796": "1", "1797": "1", "1798": "1", "1799": "1", "1801": "0", "1802": "2", "1804": "1", "1805": "1", "1806": "1", "1807": "1", "1808": "1", "1809": "2", "1811": "1", "1813": "2", "1814": "2", "1816": "2", "1817": "1", "1819": "1", "1820": "1", "1822": "2", "1824": "1", "2855": "1", "1826": "1", "1829": "2", "1830": "1", "1831": "2", "1832": "0", "1833": "0", "1835": "1", "1836": "2", "1837": "1", "1839": "2", "1840": "0", "1841": "2", "1842": "1", "1843": "2", "1844": "0", "1845": "0", "1847": "2", "1848": "0", "1849": "1", "1850": "1", "1851": "2", "1852": "2", "1853": "1", "1854": "2", "1855": "1", "1856": "2", "1857": "2", "1859": "1", "1860": "1", "1861": "1", "1862": "1", "1863": "1", "1864": "2", "1865": "0", "1866": "0", "1868": "1", "1869": "1", "1870": "1", "1871": "1", "1872": "2", "1873": "1", "1874": "1", "1875": "2", "1876": "2", "1877": "0", "1878": "1", "1879": "0", "1880": "2", "1881": "1", "1882": "2", "1883": "2", "1884": "1", "60": "1", "1888": "1", "1889": "0", "2081": "0", "1891": "2", "1892": "0", "1893": "1", "1894": "0", "1895": "1", "1896": "2", "1898": "1", "1899": "2", "1900": "2", "1901": "2", "1902": "1", "1903": "2", "1904": "2", "1906": "1", "1907": "2", "1908": "1", "1909": "1", "1910": "2", "1911": "0", "1912": "2", "1913": "2", "1914": "1", "1915": "0", "1916": "0", "1917": "0", "1918": "1", "1919": "0", "1920": "2", "1921": "0", "1922": "0", "1923": "0", "1924": "2", "1925": "0", "1927": "0", "1928": "1", "1929": "1", "1930": "1", "1931": "1", "1932": "2", "1933": "0", "1934": "1", "1935": "1", "1936": "0", "1937": "0", "1938": "2", "1939": "1", "1940": "0", "1942": "1", "1943": "1", "1944": "0", "1945": "0", "1946": "0", "1947": "0", "1948": "1", "1950": "1", "1951": "1", "1952": "2", "1953": "0", "1954": "1", "1955": "0", "1956": "2", "1957": "1", "1958": "2", "1960": "2", "1961": "1", "1962": "2", "1963": "2", "1964": "2", "1965": "2", "1967": "1", "1968": "2", "1969": "0", "1970": "1", "1971": "0", "1972": "0", "1973": "0", "1974": "1", "1975": "0", "1976": "2", "1977": "1", "1978": "0", "1979": "1", "1980": "1", "1981": "1", "1982": "2", "1983": "1", "1984": "2", "1985": "1", "1986": "0", "1987": "1", "1989": "2", "1990": "2", "1992": "1", "1993": "0", "1995": "1", "1996": "2", "1997": "1", "1998": "0", "1999": "0", "2000": "1", "2001": "1", "2003": "2", "2004": "2", "2006": "2", "2007": "0", "2008": "2", "2011": "0", "2013": "2", "2015": "1", "2016": "0", "2018": "2", "2019": "1", "2020": "0", "2022": "2", "2023": "2", "2024": "2", "2025": "2", "2026": "1", "2028": "1", "2027": "0", "2029": "1", "2031": "2", "2032": "0", "2033": "1", "2034": "0", "2035": "1", "2036": "1", "2037": "2", "2038": "2", "2039": "0", "2040": "0", "2041": "0", "2043": "0", "2044": "0", "2045": "2", "2046": "0", "2047": "0", "2048": "1", "2049": "0", "2050": "2", "2052": "1", "2053": "0", "2054": "0", "2055": "2", "2057": "1", "2058": "1", "2059": "1", "2060": "1", "2062": "1", "2063": "2", "2064": "0", "2065": "2", "2066": "1", "2067": "1", "1867": "0", "2069": "0", "4757": "1", "2071": "0", "2073": "2", "2076": "2", "1710": "0", "2077": "1", "2078": "1", "2080": "2", "2082": "0", "2083": "0", "2084": "0", "2085": "2", "2087": "2", "2088": "1", "2089": "1", "2091": "2", "2092": "2", "2093": "0", "2094": "0", "2095": "0", "2096": "1", "2097": "1", "2098": "2", "2100": "2", "2102": "1", "2103": "2", "2104": "2", "2107": "2", "2108": "0", "2110": "1", "2111": "2", "2112": "1", "2113": "2", "2114": "1", "2116": "1", "2117": "0", "2118": "1", "2119": "1", "2120": "1", "2121": "0", "2122": "1", "2123": "1", "2124": "0", "2127": "1", "2128": "0", "2129": "0", "2130": "1", "2131": "1", "2132": "0", "2133": "1", "2134": "2", "2135": "1", "2136": "1", "2137": "1", "2138": "2", "2139": "0", "2140": "1", "2142": "2", "2143": "2", "2144": "2", "2145": "0", "2146": "1", "2147": "2", "2148": "1", "2149": "1", "2150": "2", "2151": "2", "2695": "1", "2153": "1", "2154": "1", "2155": "1", "2156": "1", "2157": "2", "2159": "2", "2160": "1", "2161": "1", "2162": "2", "2163": "2", "2165": "2", "2166": "1", "2167": "1", "2168": "2", "2169": "1", "2170": "1", "2171": "0", "2172": "2", "2173": "2", "2174": "2", "2175": "2", "2176": "1", "2177": "2", "2178": "1", "2179": "1", "2180": "1", "2181": "1", "2182": "1", "2183": "2", "2184": "2", "2186": "2", "2187": "0", "2188": "1", "2189": "1", "2190": "2", "2191": "1", "2192": "1", "2193": "1", "2194": "0", "2195": "1", "2196": "1", "2197": "1", "2198": "0", "2200": "1", "2201": "2", "2202": "2", "2203": "1", "2204": "1", "2205": "1", "2206": "1", "2207": "0", "2208": "1", "2209": "1", "2210": "1", "2212": "0", "2213": "2", "2214": "2", "2215": "1", "2216": "1", "2218": "2", "2219": "1", "2220": "2", "2222": "2", "2223": "1", "2225": "2", "2227": "0", "2228": "2", "2229": "0", "2232": "0", "2233": "0", "2235": "1", "2236": "1", "2237": "1", "2238": "1", "2239": "1", "2241": "0", "2243": "1", "2244": "0", "2245": "1", "2246": "2", "2247": "0", "2248": "1", "2249": "1", "2250": "0", "2251": "0", "2252": "2", "2253": "1", "2255": "1", "2256": "0", "2257": "0", "2258": "1", "2259": "0", "2261": "0", "2262": "1", "2263": "1", "2264": "2", "2265": "1", "2266": "0", "2267": "0", "2269": "1", "2270": "1", "2271": "2", "2272": "1", "2273": "1", "2274": "1", "2276": "0", "2277": "0", "2278": "2", "2279": "2", "2280": "1", "2281": "1", "2282": "1", "2283": "0", "2284": "0", "2285": "1", "2286": "0", "2287": "0", "2288": "0", "2290": "2", "2291": "1", "2294": "2", "2295": "1", "2296": "1", "2297": "0", "2298": "2", "2299": "1", "2300": "1", "2301": "2", "2302": "1", "2304": "2", "2305": "2", "2306": "2", "2307": "1", "2308": "2", "2309": "1", "2310": "0", "2311": "0", "2312": "1", "2313": "2", "2314": "1", "57": "0", "2316": "2", "2318": "1", "2319": "0", "2320": "1", "2321": "0", "2322": "1", "2323": "0", "2326": "1", "2327": "1", "2328": "1", "2330": "2", "2331": "1", "2332": "1", "2333": "1", "2334": "1", "2335": "0", "2336": "0", "2337": "0", "2338": "1", "2341": "1", "2342": "0", "2343": "1", "2344": "2", "2345": "0", "2346": "2", "2347": "0", "2349": "2", "2350": "2", "2351": "2", "2352": "2", "2353": "1", "2354": "0", "2356": "1", "2357": "0", "2358": "2", "2359": "1", "2360": "1", "2361": "0", "2362": "1", "2363": "0", "2364": "2", "2365": "2", "2366": "2", "2367": "0", "2368": "2", "2369": "2", "2370": "1", "2372": "2", "2373": "1", "2374": "1", "2375": "1", "2377": "2", "2378": "1", "2379": "2", "2380": "1", "2381": "1", "2382": "1", "2383": "1", "2384": "1", "2386": "0", "2387": "0", "2389": "1", "2390": "2", "2391": "0", "2392": "0", "2393": "2", "2394": "0", "2395": "1", "2396": "1", "2399": "1", "2400": "0", "2401": "0", "2402": "1", "2403": "2", "2404": "1", "2405": "0", "2407": "2", "2409": "0", "2410": "0", "2411": "0", "2412": "1", "2413": "2", "2414": "1", "2415": "1", "2417": "1", "2418": "1", "2419": "1", "2420": "2", "2421": "2", "2423": "1", "2424": "1", "2425": "0", "2426": "1", "2427": "1", "2428": "1", "2430": "0", "2431": "1", "2432": "0", "2433": "1", "2434": "2", "2435": "2", "2436": "0", "2438": "2", "2439": "1", "2440": "0", "2441": "0", "2442": "0", "2443": "0", "2444": "0", "2445": "1", "2446": "1", "2447": "1", "2448": "1", "2449": "1", "2450": "2", "2451": "1", "2452": "1", "2454": "2", "2455": "0", "2456": "0", "2457": "1", "2458": "2", "2459": "1", "2460": "1", "2461": "1", "2463": "2", "2464": "2", "2465": "2", "2466": "1", "2467": "1", "2468": "2", "2469": "1", "2470": "0", "2472": "2", "2473": "0", "2474": "2", "2475": "0", "2476": "1", "2477": "1", "2478": "0", "2479": "2", "2480": "2", "2481": "2", "2482": "1", "2484": "1", "2486": "2", "2488": "1", "2489": "1", "2490": "1", "2491": "1", "2492": "0", "2493": "2", "2494": "2", "2495": "1", "2496": "1", "2497": "1", "2499": "1", "2500": "0", "2501": "0", "2502": "2", "2504": "2", "2505": "1", "2506": "2", "2507": "0", "2508": "1", "2510": "0", "2511": "2", "2513": "1", "2514": "1", "2515": "0", "2516": "1", "2517": "0", "2518": "0", "2519": "0", "2520": "0", "2521": "2", "2522": "2", "2523": "1", "2524": "0", "2525": "0", "2530": "1", "2531": "1", "2532": "0", "2533": "0", "2534": "2", "2535": "0", "2536": "2", "2537": "1", "2538": "2", "2539": "2", "2540": "2", "2541": "0", "2542": "0", "2543": "1", "2544": "0", "2545": "1", "2546": "2", "2547": "1", "2548": "1", "2549": "0", "2550": "1", "2551": "2", "2552": "0", "2553": "1", "2555": "2", "2557": "1", "2558": "2", "2559": "2", "2561": "1", "2562": "2", "2563": "2", "2566": "1", "2567": "1", "2568": "2", "2569": "2", "2570": "1", "2571": "0", "2572": "2", "2574": "0", "2575": "2", "2576": "2", "2577": "0", "2578": "1", "2579": "0", "2580": "1", "2581": "1", "2582": "0", "2583": "2", "2584": "1", "2585": "0", "2586": "1", "2587": "0", "2588": "0", "2589": "2", "2590": "1", "2591": "1", "2592": "2", "2593": "1", "2594": "1", "2512": "1", "2596": "2", "2597": "0", "2599": "2", "2600": "1", "2601": "2", "2602": "1", "2603": "1", "2605": "1", "2606": "1", "2607": "1", "2608": "1", "2610": "1", "2611": "1", "2612": "0", "2613": "2", "2614": "2", "2615": "1", "2616": "2", "2617": "1", "2618": "2", "2619": "1", "2620": "2", "2621": "2", "2622": "2", "2623": "0", "2624": "2", "2625": "0", "2627": "2", "2628": "0", "2630": "2", "2631": "2", "2632": "1", "2633": "2", "2634": "2", "2635": "2", "2636": "1", "2637": "1", "2638": "1", "2640": "1", "2642": "1", "2643": "1", "2644": "0", "2645": "1", "2646": "2", "2647": "0", "2648": "1", "2649": "0", "2650": "2", "2651": "1", "2652": "0", "2654": "2", "2656": "1", "2657": "2", "2658": "2", "2659": "1", "2660": "1", "2661": "0", "2662": "0", "2663": "2", "2664": "0", "2875": "1", "2667": "2", "2668": "0", "2675": "0", "2670": "2", "2671": "0", "2672": "0", "2673": "2", "2676": "2", "2677": "2", "2679": "0", "2680": "0", "2260": "2", "2684": "2", "2685": "2", "2691": "2", "2692": "2", "2694": "1", "2697": "0", "2699": "1", "2700": "2", "2701": "2", "2703": "0", "2704": "0", "2705": "2", "2706": "2", "2708": "2", "2710": "1", "1505": "1", "2713": "2", "2714": "2", "2715": "1", "2716": "2", "2718": "0", "2724": "1", "2729": "2", "2730": "1", "2732": "0", "2734": "2", "2739": "2", "2743": "2", "2744": "0", "2748": "2", "2749": "2", "2752": "1", "2755": "1", "2756": "1", "2754": "2", "2757": "2", "1414": "2", "2758": "2", "2760": "1", "2765": "1", "2768": "0", "2769": "0", "2918": "2", "2777": "0", "2781": "1", "2782": "1", "2783": "2", "2786": "0", "2787": "1", "2788": "1", "2802": "1", "2813": "2", "2822": "1", "2824": "0", "2825": "0", "2831": "0", "2834": "0", "2835": "0", "2829": "1", "2839": "2", "2840": "1", "2856": "1", "2764": "1", "2862": "1", "2863": "0", "2864": "2", "2867": "1", "2870": "0", "2871": "1", "2874": "2", "2876": "2", "2879": "1", "2883": "2", "2884": "0", "2885": "2", "2886": "0", "2892": "1", "2893": "1", "2898": "1", "2902": "2", "2906": "0", "2907": "2", "4590": "1", "2915": "2", "2916": "0", "2917": "0", "2919": "0", "2920": "2", "2837": "2", "2930": "2", "2932": "1", "2933": "2", "2935": "2", "2936": "0", "2940": "1", "2946": "2", "2959": "0", "4706": "2", "2965": "2", "4692": "1", "3041": "0", "3046": "1", "3052": "0", "3092": "2", "3102": "2", "3130": "2", "3158": "2", "3241": "1", "3259": "1", "3270": "2", "3347": "1", "3359": "0", "3382": "0", "3454": "0", "3516": "2", "3527": "2", "3539": "2", "3569": "2", "3618": "1", "3638": "2", "3751": "2", "3624": "2", "3823": "1", "4544": "1", "3936": "1", "4557": "2", "4480": "2", "4580": "2", "4583": "2", "4683": "2", "4772": "2", "4827": "1", "4845": "2", "4870": "1", "4893": "2", "4899": "1", "4909": "0", "4914": "0", "4917": "2", "4924": "1", "4925": "2", "4938": "2", "4970": "1", "5000": "2", "5002": "0", "4878": "2", "5025": "2", "584": "1", "5047": "1", "3504": "1", "5164": "2", "5774": "2", "5793": "2", "5831": "0", "6117": "2", "6137": "1", "6187": "1", "6221": "2", "6294": "2", "6460": "0", "6629": "1", "6827": "2", "6905": "1", "7031": "1", "7296": "0", "7524": "1", "7577": "1", "7666": "1", "7691": "2", "7705": "0", "7714": "1", "7929": "1", "7986": "2", "8242": "1", "8290": "2", "4733": "1", "8418": "2", "8480": "0", "8452": "1", "8575": "2", "8348": "2", "8781": "2", "8914": "2", "9037": "1", "9071": "1", "9136": "2", "9140": "0", "9163": "1", "9170": "1", "9183": "2", "9189": "1", "8514": "2", "9218": "2", "9234": "2", "9237": "2", "9240": "2", "9289": "2", "9308": "1", "9320": "1", "9479": "0", "9481": "1", "9514": "0", "9531": "2", "9533": "2", "9558": "1", "9565": "1", "9592": "2", "9636": "2", "9658": "2", "9669": "1", "9679": "2", "9694": "1", "9697": "2", "9757": "1", "9858": "2", "9861": "1", "9862": "2",
	}
)

// 执行脚本
func GetUidM1() {
	log.Info("开始执行脚本---GetUidM1")
	//获取数据
	url := logReadUrl + dateUrlA
	log.Info("logUrl:", url)
	logData = make(map[string]M1uid)
	readDirectoryLog(url)
	writeM1Pid()
	beego.Info("脚本执行完毕---GetUidM1")
	os.Exit(2)
}

//读取日志
func readDirectoryLog(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
Loop:
	for _, fi := range files {
		if fi.IsDir() {
			aa = fi.Name()
			bb, error := strconv.Atoi(aa)
			if error != nil {
				log.Info("字符串转换成整数失败")
			}
			if bb < 20190428 {
				continue Loop
			}
			log.Info(aa)
			readDirectoryLog(dir + fi.Name() + "\\")
		} else {
			//countContentTitle := "ml" + "\t" + "nm" + "\t" + "uid"
			//WriteLog(countLogFileName, countContentTitle)
			filename := fi.Name()

			//读取具体内容
			openUrlfile := dir + "/" + filename
			logFile, err := os.OpenFile(openUrlfile, os.O_RDWR, 0666)
			if err != nil {
				log.Info("Open file error!")
				return
			}
			defer logFile.Close()

			stat, err := logFile.Stat()
			if err != nil {
				panic(err)
			}

			var size = stat.Size()
			_ = size
			//log.Info("file size=", size)

			buf := bufio.NewReader(logFile)
			var lineNum = 0
			for {
				line, err := buf.ReadString('\n')
				line = strings.TrimSpace(line)
				lineNum++
				//log.Info(line)
				if len(line) != 0 {
					m1 := gjson.Get(line, "m1").String()
					uid := gjson.Get(line, "seg.custom.uid").String()
					uidTwo := gjson.Get(line, "seg.uid").String()
					//countContent = m1 + "\t" + uid
					if len(uid) != 0 {
						if _, ok := uidArray[uid]; ok {
							log.Info(aa)
							//把存在uid的写进map Nmpid中
							sex := uidArray[uid]
							mapString := m1 + uid + "_" + strconv.Itoa(lineNum)
							logData[mapString] = M1uid{m1, uid, sex}
							//WriteLog(countLogFileName, countContent)
						}

					}
					if len(uidTwo) != 0 {
						if _, ok := uidArray[uidTwo]; ok {
							log.Info(aa)
							//把存在uid的写进map Nmpid中
							sexTwo := uidArray[uidTwo]
							mapString := m1 + uidTwo + "_" + strconv.Itoa(lineNum)
							logData[mapString] = M1uid{m1, uidTwo, sexTwo}
							//WriteLog(countLogFileName, countContent)
						}
					}
				}
				if err != nil {
					if err == io.EOF {
						//log.Info("File read ok!")
						break
					} else {
						//log.Info("Read file error!", err)
						return
					}
				}
			}
		}
	}
}

func writeM1Pid() {
	countContentTitle := "ml" + "\t" + "uid" + "\t" + "sex"
	WriteLog(countLogFileName, countContentTitle)
	for key, value := range logData {
		log.Info("key:"+key, value)
		countContent = value.M1 + "\t" + value.UID + "\t" + value.SEX
		WriteLog(countLogFileName, countContent)
	}
}
