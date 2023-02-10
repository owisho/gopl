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

//var interestPreprcessMap = map[string]bool{
//	"DV_Trend_Preproc_delta_nl_supervised_keywords_51_20220912_af": true,
//	"DV_Trend_Preproc_delta_nl_keywords_51_20220912_af":            true,
//}

var interestPreprocessFeatures = []string{"nl_supervised_keywords_51", "nl_keywords_51"}

//var tfidfPreprocessMap = map[string]bool{
//	"DV_Trend_Preproc_delta_nl_supervised_keywords_0_20220912": true,
//	"DV_Trend_Preproc_delta_nl_keywords_0_20220912":            true,
//}

var tfidfPreprocessFeatures = []string{"nl_supervised_keywords_0", "nl_keywords_0"}

func TestPreprocessYarnInfo(t *testing.T) {
	f, _ := os.Open("/Users/wangyang/Desktop/preprocess_task.csv")
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, _ := csvReader.ReadAll()
	for _, line := range records {
		appId, appName := line[0], line[2]
		//过滤title
		if appName == "Name" {
			continue
		}
		//过滤市场
		if !strings.HasSuffix(appName, "_af") {
			continue
		}
		//过滤时间
		if !strings.Contains(appName, "_20230204") {
			continue
		}
		//过滤interest feature
		var toContinue = false
		for _, feature := range interestPreprocessFeatures {
			if strings.Contains(appName, feature) {
				toContinue = true
				break
			}
		}
		if toContinue {
			continue
		}
		//打印最终结果
		println(appId)
	}
}

//var interestAggMap = map[string]bool{
//	"Davinci_Pref_Aggr_20220912_nl_supervised_keywords_51_af": true,
//	"Davinci_Pref_Aggr_20220912_nl_keywords_51_af":            true,
//}

var interestAggFeatures = []string{"nl_supervised_keywords_51", "nl_keywords_51"}

//var tfidfAggMap = map[string]bool{
//	"Davinci_Pref_Aggr_20220912_nl_supervised_keywords_tfidf_511_af": true,
//	"Davinci_Pref_Aggr_20220912_nl_keywords_tfidf_511_af":            true,
//}

var tfidfAggFeatures = []string{"nl_supervised_keywords_tfidf_511", "nl_keywords_tfidf_511"}

func TestAggYarnInfo(t *testing.T) {
	f, _ := os.Open("/Users/wangyang/Desktop/aggr_task.csv")
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, _ := csvReader.ReadAll()
	for _, line := range data {
		appId, appName := line[0], line[2]
		//过滤title
		if appName == "Name" {
			continue
		}
		//过滤市场
		if !strings.HasSuffix(appName, "_af") {
			continue
		}
		//过滤时间
		if !strings.Contains(appName, "_20230204") {
			continue
		}
		//过滤interest feature
		var toContinue = false
		for _, feature := range interestAggFeatures {
			if strings.Contains(appName, feature) {
				toContinue = true
				break
			}
		}
		if toContinue {
			continue
		}
		//过滤tfidf feature
		for _, feature := range tfidfAggFeatures {
			if strings.Contains(appName, feature) {
				toContinue = true
				break
			}
		}
		if toContinue {
			continue
		}
		//打印最终结果
		println(appId)
	}
}

func TestYarnInfo(t *testing.T) {
	f, _ := os.Open("/Users/wangyang/Desktop/application_ids")
	br, _ := io.ReadAll(f)
	totalMUsed, totolCpuTime := int64(0), int64(0)
	for _, line := range strings.Split(string(br), "\n") {
		line = strings.TrimSpace(line)
		fmt.Println(line)
		yarnsite := "http://n17-07-04:8088/cluster/app/" //非洲 yarn 地址
		//yarnsite := "http://a31-01.fn.lati.osa:8088/cluster/app/" //us yarn 地址
		cost, mused, cputime, err := GetYarnInfo(yarnsite + line)
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
