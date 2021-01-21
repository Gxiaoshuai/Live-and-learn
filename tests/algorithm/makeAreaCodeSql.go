package main

import (
	"fmt"
	"strings"
)

func main (){
	
	
//var formate string ="INSERT INTO \"public\".\"callers\" (\"caller_sn\",\"user_sn\",\"user_name\",\"team_sn\",\"team_name\",\"created_at\",\"project_sn\",\"source\",\"caller_group\",\"caller_range\",\"note\",\"area_code\")VALUES('%sd2f18d39f1e216g','dc0e8029f1e216g','200000111','webank','TEAM|webank','2020-08-28 11:37:11.891576','%s','200000111','0755-107106','075533944187,075533944198,075533944203,075533944208,075533944210,075533944214,075533944215,075533944218,075533944235,075533944241','深圳联通机器人','0755');"
	var callerList = `107106,075533944187,075533944198,075533944203,075533944208,075533944210,02022009771,057728784320,02285579269,09314266602,07718067552,04723992125,05923284557,09517875851,045187844443,089836697347,055167157379,02388294932,02388294933,02022009778,02022009780
107107,076023707145,075763285200,07523037585,02388294942,02987947125,02160387587,076933556659,057426957480,02566970994,051066121692,055166195364,053189604957,01080909159,057728784313,02022009773,057728784270,075533944214,075533944215,075533944218,075533944235,075533944241
107108,057728784312,02022009781,051287830664,075533944165,075533944166
107109,02285579267,02285579268,09314266598,09314266599,07718067542,07718067544,04723992115,04723992124,05923284384,05923284415,09517875549,09517875567,045187844439,089836697319,089836697327,076023705341,076023705559,075722193921,075722193922,07523037583,07523037584,02388294939,02388294941,02987946916,02987946932,02160387584,02160387586,076933554778,076933555791,057728784318,057728784319,057426957478,057426957479,02566970970,02566970993,051066121557,051066121690,053266014634,053266014636,055166195362,055166195363,02022004815,02022009747,02022009770,075533614859,075533614867,075533614979,075533614788,075533614834,075533614835`
	strArr := strings.Split(callerList,"\n")
	var callerMap map[string][]string = make(map[string][]string)
	var tempArr []string
	var areaCode string
	for _,val := range strArr{
		tempArr = strings.Split(val,",")
		callerGroupId := tempArr[0]
		for i := 1;i <  len(tempArr);i++{
			areaCode = tempArr[i][0:len(tempArr[i])-8]
			callerMap[callerGroupId+"_"+areaCode] = append(callerMap[callerGroupId+"_"+areaCode], tempArr[i])
		}
	}
	for key,val := range callerMap{
		tempArr = strings.Split(key,"_")
		callPoolId,areaCode := tempArr[0],tempArr[1]
		fmt.Printf("INSERT INTO \"public\".\"callers\" (\"caller_sn\",\"user_sn\",\"user_name\",\"team_sn\",\"team_name\",\"created_at\",\"project_sn\",\"source\",\"caller_group\",\"caller_range\",\"note\",\"area_code\")VALUES('%sd2f18d39f1e221g','dc0e8029f1e216g','200000111','webank','TEAM|webank','2020-08-28 11:37:11.891576','%s','200000111','%s-%s','%s','深圳联通机器人','%s');",callPoolId,callPoolId,areaCode,callPoolId,strings.Join(val,","),areaCode)
		if areaCode == "0755"{
			fmt.Printf("INSERT INTO \"public\".\"callers\" (\"caller_sn\",\"user_sn\",\"user_name\",\"team_sn\",\"team_name\",\"created_at\",\"project_sn\",\"source\",\"caller_group\",\"caller_range\",\"note\",\"area_code\")VALUES('%sd2f18d39f1e221g','dc0e8029f1e216g','200000111','webank','TEAM|webank','2020-08-28 11:37:11.891576','%s','200000111','*-%s','%s','深圳联通机器人','%s');",callPoolId,callPoolId,callPoolId,strings.Join(val,","),areaCode)
		}
		fmt.Println("")
	}

}