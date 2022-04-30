package utils

import (
	"project_1/models"
	"runtime"
	"sort"
	"strings"
)

//简单的词频统计任务
func CountTestBase(fileText string) []*models.WordFrequency {

	//根据CPU核数新开协程
	newRountineCount := runtime.NumCPU()*2 - 1
	runtime.GOMAXPROCS(newRountineCount + 1)
	//切分文件
	parts := splitFileText(fileText, newRountineCount)

	var ch chan map[string]int = make(chan map[string]int, newRountineCount)
	for i := 0; i < newRountineCount; i++ {
		go countTest(parts[i], ch)
	}

	//主线程接收数据
	var totalWordsMap map[string]int = make(map[string]int, 0)
	completeCount := 0
	for {
		receiveData := <-ch
		for k, v := range receiveData {
			totalWordsMap[strings.ToLower(k)] += v
		}
		completeCount++

		if newRountineCount == completeCount {
			break
		}
	}

	//添加进slice，并排序
	//var list []models.WordFrequency = make([]models.WordFrequency, 0)
	list := make(WordCountBeanList, 0)
	for k, v := range totalWordsMap {
		list = append(list, NewWordCountBean(k, v))
	}
	sort.Sort(list)
	if len(list) > 10 {
		list = list[:10]
	}
	return list
}

func countTest(text string, ch chan map[string]int) {
	var wordMap map[string]int = make(map[string]int, 0)

	//按字母读取，除26个字母（大小写）之外的所有字符均认为是分隔符
	startIndex := 0
	letterStart := false
	for i, v := range text {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			if !letterStart {
				letterStart = true
				startIndex = i
			}
		} else {
			if letterStart {
				wordMap[text[startIndex:i]]++
				letterStart = false
			}
		}
	}

	//最后一个单词
	if letterStart {
		wordMap[text[startIndex:]]++
	}
	ch <- wordMap
}

//将全文分成n段
func splitFileText(fileText string, n int) []string {
	length := len(fileText)
	parts := make([]string, n)

	lastPostion := 0
	for i := 0; i < n-1; i++ {
		position := length / n * (i + 1)
		for string(fileText[position]) != " " {
			position++
		}

		parts[i] = fileText[lastPostion:position]
		lastPostion = position
	}

	//最后一段
	parts[n-1] = fileText[lastPostion:]
	return parts
}

func NewWordCountBean(word string, count int) *models.WordFrequency {
	return &models.WordFrequency{word, count}
}

type WordCountBeanList []*models.WordFrequency

func (list WordCountBeanList) Len() int {
	return len(list)
}

func (list WordCountBeanList) Less(i, j int) bool {
	if list[i].Count > list[j].Count {
		return true
	} else if list[i].Count < list[j].Count {
		return false
	} else {
		return list[i].Count < list[j].Count
	}
}

func (list WordCountBeanList) Swap(i, j int) {
	var temp *models.WordFrequency = list[i]
	list[i] = list[j]
	list[j] = temp
}
