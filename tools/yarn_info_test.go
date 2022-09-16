package tools

import (
	"crypto/tls"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestTaskNum(t *testing.T) {
	f, _ := os.Open("/Users/wangyang/Downloads/preprocess_task.csv")
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, _ := csvReader.ReadAll()

	taskMap := make(map[string]int)
	for _, line := range data {
		appName := line[2]
		if strings.HasSuffix(appName, "_af") {
			taskMap[appName] += 1
		}
	}
	fmt.Printf("taskMap len= %d,\n taskMap=%v\n", len(taskMap), taskMap)
}

var interestPreprcessMap = map[string]bool{
	"DV_Trend_Preproc_delta_nl_supervised_keywords_51_20220912_af": true,
	"DV_Trend_Preproc_delta_nl_keywords_51_20220912_af":            true,
}

var tfidfPreprocessMap = map[string]bool{
	"DV_Trend_Preproc_delta_nl_supervised_keywords_0_20220912": true,
	"DV_Trend_Preproc_delta_nl_keywords_0_20220912":            true,
}

func TestPreprocessYarnInfo(t *testing.T) {
	f, _ := os.Open("/Users/wangyang/Downloads/preprocess_task.csv")
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, _ := csvReader.ReadAll()
	for _, line := range data {
		appId, appName := line[0], line[2]
		//if strings.HasSuffix(appName, "_af") &&
		//	!interestPreprcessMap[appName] {
		//	//if strings.HasSuffix(appName, "_af") {
		//	fmt.Println(appId)
		//}
		//interest preprocess
		if interestPreprcessMap[appName] {
			fmt.Println(appId)
		}
		//tfidf preprocess
		//if tfidfPreprocessMap[appName]{
		//	fmt.Println(appId)
		//}
	}
}

var interestAggMap = map[string]bool{
	"Davinci_Pref_Aggr_20220912_nl_supervised_keywords_51_af": true,
	"Davinci_Pref_Aggr_20220912_nl_keywords_51_af":            true,
}

var tfidfAggMap = map[string]bool{
	"Davinci_Pref_Aggr_20220912_nl_supervised_keywords_tfidf_511_af": true,
	"Davinci_Pref_Aggr_20220912_nl_keywords_tfidf_511_af":            true,
}

func TestAggYarnInfo(t *testing.T) {
	f, _ := os.Open("/Users/wangyang/Downloads/agg_task.csv")
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, _ := csvReader.ReadAll()
	for _, line := range data {
		appId, appName := line[0], line[2]
		//if strings.HasSuffix(appName, "_af") &&
		//	!interestAggMap[appName] &&
		//	!tfidfAggMap[appName] {
		//	fmt.Println(appId)
		//}
		// interest agg application
		if interestAggMap[appName] {
			fmt.Println(appId)
		}
		//tfidf agg application
		//if tfidfAggMap[appName]{
		//	fmt.Println(appId)
		//}
	}
}

func TestYarnInfo(t *testing.T) {
	f, _ := os.Open("/Users/wangyang/Desktop/application_ids")
	br, _ := io.ReadAll(f)
	totalMUsed, totolCpuTime := int64(0), int64(0)
	for _, line := range strings.Split(string(br), "\n") {
		line = strings.TrimSpace(line)
		fmt.Println(line)
		dbImpPre := "http://n17-07-04:8088/cluster/app/"
		//preProcessPre := "http://a31-01.fn.lati.osa:8088/cluster/app/"
		cost, mused, cputime, err := GetYarnInfo(dbImpPre + line)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("cost=%s,mused=%d(MB),cputime=%d(s)\n", cost, mused, cputime)
		totalMUsed += mused
		totolCpuTime += cputime
	}
	fmt.Printf("totalMUsed=%d(MB),totalCpuTime=%d(s)\n", totalMUsed, totolCpuTime)
}

//GetYarnInfo 获取yarn页面的信息
//return cost 任务消耗时间
//mused 使用内存总量（mb-seconds）
//cputime cpu使用时间（vcore-seconds）
func GetYarnInfo(url string) (cost string, mused int64, cputime int64, err error) {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:4444", nil, proxy.Direct)
	httpTransport := &http.Transport{
		Dial:            dialer.Dial,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: httpTransport,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("出错了", err)
		return
	}
	defer resp.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	text, _ := doc.Find("th:contains('Elapsed:')").Next().Html()
	mr, _ := doc.Find("th:contains('Aggregate Resource Allocation:')").Next().Html()
	//fmt.Println(strings.TrimSpace(text))
	//fmt.Println(strings.TrimSpace(mr))
	arr := strings.Split(mr, ",")
	if len(arr) != 2 {
		return "", 0, 0, errors.New("unexpected error")
	}
	musedStr := strings.ReplaceAll(strings.TrimSpace(arr[0]), "MB-seconds", "")
	mused, err = strconv.ParseInt(strings.TrimSpace(musedStr), 10, 64)
	if err != nil {
		return "", 0, 0, err
	}
	cputimeStr := strings.ReplaceAll(strings.TrimSpace(arr[1]), "vcore-seconds", "")
	cputime, err = strconv.ParseInt(strings.TrimSpace(cputimeStr), 10, 64)
	if err != nil {
		return "", 0, 0, err
	}
	return strings.TrimSpace(text), mused, cputime, nil
}
