package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var schools = map[int]string{12668: "广东技术师范大学天河学院", 10571: "广东医科大学", 10574: "华南师范大学", 10576: "韶关学院", 10577: "惠州学院", 10578: "韩山师范学院", 10579: "岭南师范学院", 10580: "肇庆学院", 10582: "嘉应学院", 10585: "广州体育学院", 10586: "广州美术学院", 10588: "广东技术师范大学", 10592: "广东财经大学", 10822: "广东白云学院", 11106: "广州航海学院", 11347: "仲恺农业工程学院", 11349: "五邑大学", 11540: "广东金融学院", 11545: "电子科技大学中山学院", 11656: "广东石油化工学院", 11819: "东莞理工学院", 11847: "佛山科学技术学院", 12059: "广东培正学院", 12619: "中山大学南方学院", 12621: "广东财经大学华商学院", 12622: "广东海洋大学寸金学院", 12623: "华南农业大学珠江学院", 13177: "北京师范大学珠海分校", 13656: "广东工业大学华立学院", 13657: "广州大学松田学院", 13667: "广州商学院", 13675: "北京理工大学珠海学院", 13684: "吉林大学珠海学院", 13714: "广州工商学院", 13717: "广州科技职业技术大学", 13719: "广东科技学院", 13720: "广东理工学院", 13721: "广东工商职业技术大学", 13844: "东莞理工学院城市学院", 13902: "中山大学新华学院", 14278: "广东第二师范学院"}

func main() {
	fmt.Printf("学校列表：\r\n")
	for k, v := range schools {
		fmt.Printf("编号：%d\t学校：%s\r\n", k, v)
	}
	fmt.Printf("注意：若无上述列表无需要查找的目标学校请自行在官网寻找学校编号\r\n作者：NetTunnel 使用本程序使用 Golang 编写\r\n程序免费发布并开源于 Github 使用本程序及源代码请遵循 GPLv3 协议, 如果发现倒卖请向相关平台举报\r\n项目地址：https://github.com/NetTunnel/www.eeagd.edu.cn_bcks_Spider\r\n程序仅用于交流测试禁止用于商业用途请于下载后24小时内删除，否则后果自负\r\n查询完毕后会回到此界面，如要退出程序请点击右上角的关闭按钮\r\n")
	Select()
}

func Select() {
	fmt.Printf("请输入要查找的学校编号(直接回车查询3月份新增预报名人数,输入1查询上述列表全部学校3月份已审核的报名人数)：")
	schoolId, typeId := "", ""
	_, _ = fmt.Scanln(&schoolId)
	schoolId = strings.Replace(schoolId, "\n", "", -1)
	fmt.Print("查询中，请稍等......\r\n")
	if schoolId == "" {
		fmt.Printf("此院校 3 月份新增预报名人数为：%d \r\n", batchVerify("3", true))
	} else if schoolId == "1" {
		orgitotal, nowtotal, orgi, now := 0, 0, 0, 0
		for k,v := range schools {
			orgi = batchVerify(fmt.Sprintf("%d", k)+"0", false)
			now = batchVerify(fmt.Sprintf("%d", k)+"1", false)
			orgitotal += orgi
			nowtotal += now
			fmt.Printf("%s\t1月份原有报名人数：%d\t3 月份新增已审核人数为：%d\t合计：%d \r\n", v, orgi, now, orgi+now)
		}
		fmt.Printf("上述列表全部学校 1月份原有报名人数：%d , 3 月份新增已审核人数为：%d \r\n", orgitotal, nowtotal)
	} else {
		fmt.Printf("查询模式1为穷举可筛选出改志愿后的人数（注意没做延时访问机制有可能访问异常慎用, 出现问题请后果自负），否则为二分法效率高速度快但不可筛选出改志愿后的人数\r\n请选择查询模式：")
		_, _ = fmt.Scanln(&typeId)
		typeId = strings.Replace(typeId, "\n", "", -1)
		if typeId != "1" {
			first, second, total :=  batchVerify(schoolId+"0", false), batchVerify(schoolId+"1", false), 0
			total = first + second
			fmt.Printf("此院校 1 月报考人数为：%d , 3 月份新增报考人数为：%d , 总计：%d\r\n", first, second, total)
		} else {
			total, sum := batchVerify(schoolId+"0", false), 0
			for i := 0; i <= total; i++ {
				if ! verifyUserId(fmt.Sprintf("%s%04d", schoolId+"0", i)) {
					sum++
				}
			}
			fmt.Printf("此院校 1 月报考人数为：%d , 3 月份跳车人数为：%d , 剩余1月报名人数：%d\r\n", total, sum, total-sum)
		}
	}
	fmt.Printf("本程序已开源项目地址：https://github.com/NetTunnel/www.eeagd.edu.cn_bcks_Spider 使用本程序及其源代码请遵循 GPLv3 协议\r\n如要继续查询请按回车键，否则请点击右上角关闭按钮退出程序")
	_, _ = fmt.Scanf("%s")
	Select()
}

func batchVerify(head string, kslx0 bool) int {
	var padding, userId, start, end = "%04d", 0, 0, 9999
	if kslx0 {
		padding, userId, start, end = "%05d", 0, 0, 99999
	}
	for start <= end {
		userId = (start + end) >> 1
		if verifyUserId(head + fmt.Sprintf(padding, userId)) {
			start = userId + 1
		} else if !verifyUserId(head + fmt.Sprintf(padding, userId)) {
			if verifyUserId(head + fmt.Sprintf(padding, userId+1)) { //验证一下不存在的准考号的下一个是否存在一定程度上防止位于二分点的考生号跳车影响准确性
				start = userId + 1
			} else {
				end = userId - 1
			}
		}
	}
	if userId > 0 {
		if verifyUserId(head + fmt.Sprintf(padding, userId)) {
			return userId + 1
		} else {
			return userId
		}
	} else {
		return 0
	}
}

func verifyUserId(UserId string) bool {
	path, kslx0, key := "ksretrievepwd.jsp", "0", "reSetPwd"
	if len(UserId) <= 6 {
		path, kslx0, key = "retrievepwd.jsp", "1", "mmtswt"
	}
	resp, _ := http.Post("https://www.eeagd.edu.cn/bcks/ybmks/"+path, "application/x-www-form-urlencoded", strings.NewReader("kslx0="+kslx0+"&userid="+UserId))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return strings.Contains(string(body), key)
}
