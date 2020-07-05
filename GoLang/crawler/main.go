package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const BIRTHS = `<h2>.+Births.+</h2>`
const DEATHS = `<h2>.+Deaths.+</h2>`
const HOLIDAYS = `<h2>.+Holidays.+</h2>`
//const COUNTRIES = map[string]int{}
const DELIMITTER = "&#8211"
const CSV_DELIMITTER = ";"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
func generateUrl(month string, day int) string {
	rootURL := "http://en.wikipedia.org/wiki/"
	URL := rootURL + month + "_" + strconv.Itoa(day)
	return URL
}

func generateURLBody(URL string) *string {
	res, err := http.Get(URL)
	checkError(err)
	defer res.Body.Close()
	bodyByte, err := ioutil.ReadAll(res.Body)
	checkError(err)
	bodyString := string(bodyByte)
	return &(bodyString)
}

func generateLabelStrings(URLBody *string, label string) []string{
	var firstLabel string
	var nextLabel string
	switch label {
	case "BIRTHS":
		firstLabel = BIRTHS
		nextLabel = DEATHS
	case "DEATHS":
		firstLabel = DEATHS
		nextLabel = HOLIDAYS
	}
	firstReg := regexp.MustCompile(firstLabel)
	nextReg := regexp.MustCompile(nextLabel)
	firstLabelIndex := firstReg.FindStringIndex(*URLBody)
	nextLabelIndex := nextReg.FindStringIndex(*URLBody)
	labelStrings := strings.ToLower((*URLBody)[firstLabelIndex[1]+1:nextLabelIndex[0]])
	return strings.Split(labelStrings, "\n")
}

type person struct {
	year int
	month string
	day int
	fullName string
	firstName string
	lastName string
	dob string
	country string
	description string
	weekday string
}

func sparseName(fullName string) (string, string) {
	invalidNameSuffix := map[string]int{"i":0, "ii":0, "iii":0, "iv":0, "v":0, "vi":0, "vii":0,
		"viii":0, "ix":0, "x":0, "xi":0, "xiii":0}
	var (
		firstName string
		lastName string
	)
	fullName = strings.TrimRight(fullName, " ")
	fullNameList := strings.Split(fullName, " ")
	if len(fullNameList) >= 2 {
		firstName = fullNameList[0]
		for i := len(fullNameList)-1; i >= 0 && fullNameList[i] != " " ; i-- {
			_, ok := invalidNameSuffix[fullNameList[i]]
			if !ok {
				lastName = fullNameList[i]
				break
			}
		}
	}
	return firstName, lastName
}

func extractYear(singlePersonInfo string) int {
	yearLabel := `[0-9]{4}`
	yearReg := regexp.MustCompile(yearLabel)
	yearIndex := yearReg.FindStringIndex(singlePersonInfo)
	if len(yearIndex) != 0 {
		year, err := strconv.Atoi(singlePersonInfo[yearIndex[0]:(yearIndex[0]+4)])
		checkError(err)
		return year
	}
	return 0
}

func extractName(singlePersonInfo string) string {
	titleLabel := `title="`
	titleReg := regexp.MustCompile(titleLabel)
	nameIndex := titleReg.FindStringIndex(singlePersonInfo)
	if len(nameIndex) >= 2 {
		startIndex := nameIndex[1]
		var endIndex int
		for i := startIndex; string(singlePersonInfo[i]) != "\"" && string(singlePersonInfo[i]) != "("; i++ {
			endIndex = i+1
		}
		fullName := singlePersonInfo[startIndex:endIndex]
		fullName = strings.Split(fullName, ",")[0]
		fullName = strings.Split(fullName, ";")[0]
		fullName = strings.TrimSpace(fullName)
		return fullName
	}
	return ""
}

func extractDescription(singlePersonInfo string) string {
	singlePersonDescription := strings.Split(singlePersonInfo, ",")[len(strings.Split(singlePersonInfo, ","))-1]
	var endIndex int
	for i := 0; string(singlePersonDescription[i]) != "<"; i++ {
		endIndex = i+1
	}
	return singlePersonDescription[:endIndex]
}

func checkWeekday(year int, month string, day int) string {
	const longForm = "January 2, 2006"
	date, _ := time.Parse(longForm, month + " " + strconv.Itoa(day) + ", " + strconv.Itoa(year))
	return date.Weekday().String()
}

func handleAllPersons(singlePerson *string, p *person) {
	singlePersonInfo := strings.Split(*singlePerson, DELIMITTER)
	if len(singlePersonInfo) == 2 {
		year := extractYear(singlePersonInfo[0])
		fullName := extractName(singlePersonInfo[1])
		firstName, lastName := sparseName(fullName)
		description := extractDescription(singlePersonInfo[1])
		weekday := checkWeekday(year, p.month, p.day)
		p.year = year
		p.fullName = fullName
		p.firstName = firstName
		p.lastName = lastName
		p.dob = strconv.Itoa(year) + "-" + p.month + "-" + strconv.Itoa(p.day)
		p.description = strings.TrimSpace(description)
		p.weekday = weekday
	}
}

func combinePersonInfo(p *person) []string {
	combinedPersonInfoSlice := []string{}
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, p.dob)
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, p.fullName)
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, p.firstName)
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, p.lastName)
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, strconv.Itoa(p.year))
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, p.month)
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, strconv.Itoa(p.day))
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, p.description)
	combinedPersonInfoSlice = append(combinedPersonInfoSlice, p.weekday)
	return combinedPersonInfoSlice
}

func main() {
	//dates := map[string]int{"January": 31, "February": 29, "March": 31,
	//						"April": 30, "May": 31, "June": 30,
	//						"July": 31, "August": 31, "September": 30,
	//						"October": 31, "November": 30, "December": 31}
	dates := map[string]int{"April":2}
	startDate := 1
	file, err := os.OpenFile("test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	defer file.Close()
	for month, days := range dates {
		for day := startDate; day <= days; day++ {
			URL := generateUrl(month, day)
			URLBody := generateURLBody(URL)
			labelStrings := generateLabelStrings(URLBody, "BIRTHS")
			for _, singlePerson := range labelStrings {
				p := *(new(person))
				p.day = day
				p.month = month
				handleAllPersons(&singlePerson, &p)
				if len(strconv.Itoa(p.year)) != 4  || len(singlePerson) < 10 {
					continue
				}
				combinedPersonInfo := combinePersonInfo(&p)
				fmt.Println(combinedPersonInfo)
				csvWriter := csv.NewWriter(file)
				csvWriter.Write(combinedPersonInfo)
				csvWriter.Flush()
			}
		}
	}
}
