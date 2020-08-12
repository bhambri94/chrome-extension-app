package chrome

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ChromeExtensionResponse struct {
	Draw            int             `json:"draw"`
	RecordsTotal    int             `json:"recordsTotal"`
	RecordsFiltered int             `json:"recordsFiltered"`
	Data            [][]interface{} `json:"data"`
}

func GetChromeExtensionWithFilters(category string, users string, ratings string, reviewsCount string) [][]interface{} {
	var finalValues [][]interface{}
	epochTime := 1597038427648
	start := 0
	i := 0
	recordsTotal := 1
	firstApiCall := true
	for i < recordsTotal {
		length := 100
		url := "https://chromeflix.net/API/ajax-paging.php?draw=2&columns%255B0%255D%255Bdata%255D=0&columns%255B0%255D%255Bname%255D=&columns%255B0%255D%255Bsearchable%255D=true&columns%255B0%255D%255Borderable%255D=true&columns%255B0%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B0%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B1%255D%255Bdata%255D=1&columns%255B1%255D%255Bname%255D=&columns%255B1%255D%255Bsearchable%255D=true&columns%255B1%255D%255Borderable%255D=true&columns%255B1%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B1%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B2%255D%255Bdata%255D=2&columns%255B2%255D%255Bname%255D=&columns%255B2%255D%255Bsearchable%255D=true&columns%255B2%255D%255Borderable%255D=true&columns%255B2%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B2%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B3%255D%255Bdata%255D=3&columns%255B3%255D%255Bname%255D=&columns%255B3%255D%255Bsearchable%255D=true&columns%255B3%255D%255Borderable%255D=true&columns%255B3%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B3%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B4%255D%255Bdata%255D=4&columns%255B4%255D%255Bname%255D=&columns%255B4%255D%255Bsearchable%255D=true&columns%255B4%255D%255Borderable%255D=true&columns%255B4%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B4%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B5%255D%255Bdata%255D=5&columns%255B5%255D%255Bname%255D=&columns%255B5%255D%255Bsearchable%255D=true&columns%255B5%255D%255Borderable%255D=true&columns%255B5%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B5%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B6%255D%255Bdata%255D=6&columns%255B6%255D%255Bname%255D=&columns%255B6%255D%255Bsearchable%255D=true&columns%255B6%255D%255Borderable%255D=true&columns%255B6%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B6%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B7%255D%255Bdata%255D=7&columns%255B7%255D%255Bname%255D=&columns%255B7%255D%255Bsearchable%255D=true&columns%255B7%255D%255Borderable%255D=true&columns%255B7%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B7%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B8%255D%255Bdata%255D=8&columns%255B8%255D%255Bname%255D=&columns%255B8%255D%255Bsearchable%255D=true&columns%255B8%255D%255Borderable%255D=true&columns%255B8%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B8%255D%255Bsearch%255D%255Bregex%255D=false&order%255B0%255D%255Bcolumn%255D=0&order%255B0%255D%255Bdir%255D=asc&start=" + strconv.Itoa(start) + "&length=" + strconv.Itoa(length) + "&search%255Bvalue%255D=&search%255Bregex%255D=false&_=" + strconv.Itoa(epochTime)
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("authority", "chromeflix.net")
		req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
		req.Header.Add("x-requested-with", "XMLHttpRequest")
		req.Header.Add("sec-fetch-site", "same-origin")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("referer", "https://chromeflix.net/chrome-extension-search-engine/")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cookie", "__cfduid=d269c76a918082b4515e6ae9592ba302c1597037948; _ga=GA1.2.1544050276.1597037957; _gid=GA1.2.1406445235.1597037957; _fbp=fb.1.1597037956707.1473496876; signup-email=ss%40gmail.com; _gat_gtag_UA_159669744_2=1 max_users=; max_reviews=; max_score=; min_users="+users+"; min_reviews="+reviewsCount+"; min_score="+ratings)

		res, err := client.Do(req)
		defer res.Body.Close()
		start = start + 100

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}

		var chromeExtension ChromeExtensionResponse
		err = json.Unmarshal(body, &chromeExtension)
		if err != nil {
			fmt.Println("whoops:", err)
		}
		if firstApiCall {
			recordsTotal = (chromeExtension.RecordsTotal) / 100
			firstApiCall = false
		}
		j := 0
		usersInt, _ := strconv.Atoi(users)
		ratingsFloat, _ := strconv.ParseFloat(ratings, 4)
		reviewsCountInt, _ := strconv.Atoi(reviewsCount)
		for j < 100 {
			Value := fmt.Sprintf("%v", chromeExtension.Data[j][4])
			ValueUsersInt, _ := strconv.Atoi(Value)
			Value = fmt.Sprintf("%v", chromeExtension.Data[j][5])
			valueRatingsFloat, _ := strconv.ParseFloat(Value, 4)
			Value = fmt.Sprintf("%v", chromeExtension.Data[j][6])
			valueReviewsCountInt, _ := strconv.Atoi(Value)
			if category == chromeExtension.Data[j][1] && usersInt <= ValueUsersInt && ratingsFloat <= valueRatingsFloat && reviewsCountInt <= valueReviewsCountInt {
				finalValues = append(finalValues, chromeExtension.Data[j])
				// fmt.Println(finalValues)
			}
			j++
		}
		// time.Sleep(2 * time.Second)
		i++
	}
	return finalValues
}

func GetChromeExtension() [][]interface{} {
	var finalValues [][]interface{}
	epochTime := 1597038427648
	start := 0
	i := 0
	recordsTotal := 1
	firstApiCall := true
	for i < recordsTotal {
		length := 100
		url := "https://chromeflix.net/API/ajax-paging.php?draw=2&columns%255B0%255D%255Bdata%255D=0&columns%255B0%255D%255Bname%255D=&columns%255B0%255D%255Bsearchable%255D=true&columns%255B0%255D%255Borderable%255D=true&columns%255B0%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B0%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B1%255D%255Bdata%255D=1&columns%255B1%255D%255Bname%255D=&columns%255B1%255D%255Bsearchable%255D=true&columns%255B1%255D%255Borderable%255D=true&columns%255B1%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B1%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B2%255D%255Bdata%255D=2&columns%255B2%255D%255Bname%255D=&columns%255B2%255D%255Bsearchable%255D=true&columns%255B2%255D%255Borderable%255D=true&columns%255B2%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B2%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B3%255D%255Bdata%255D=3&columns%255B3%255D%255Bname%255D=&columns%255B3%255D%255Bsearchable%255D=true&columns%255B3%255D%255Borderable%255D=true&columns%255B3%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B3%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B4%255D%255Bdata%255D=4&columns%255B4%255D%255Bname%255D=&columns%255B4%255D%255Bsearchable%255D=true&columns%255B4%255D%255Borderable%255D=true&columns%255B4%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B4%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B5%255D%255Bdata%255D=5&columns%255B5%255D%255Bname%255D=&columns%255B5%255D%255Bsearchable%255D=true&columns%255B5%255D%255Borderable%255D=true&columns%255B5%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B5%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B6%255D%255Bdata%255D=6&columns%255B6%255D%255Bname%255D=&columns%255B6%255D%255Bsearchable%255D=true&columns%255B6%255D%255Borderable%255D=true&columns%255B6%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B6%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B7%255D%255Bdata%255D=7&columns%255B7%255D%255Bname%255D=&columns%255B7%255D%255Bsearchable%255D=true&columns%255B7%255D%255Borderable%255D=true&columns%255B7%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B7%255D%255Bsearch%255D%255Bregex%255D=false&columns%255B8%255D%255Bdata%255D=8&columns%255B8%255D%255Bname%255D=&columns%255B8%255D%255Bsearchable%255D=true&columns%255B8%255D%255Borderable%255D=true&columns%255B8%255D%255Bsearch%255D%255Bvalue%255D=&columns%255B8%255D%255Bsearch%255D%255Bregex%255D=false&order%255B0%255D%255Bcolumn%255D=0&order%255B0%255D%255Bdir%255D=asc&start=" + strconv.Itoa(start) + "&length=" + strconv.Itoa(length) + "&search%255Bvalue%255D=&search%255Bregex%255D=false&_=" + strconv.Itoa(epochTime)
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("authority", "chromeflix.net")
		req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
		req.Header.Add("x-requested-with", "XMLHttpRequest")
		req.Header.Add("sec-fetch-site", "same-origin")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("referer", "https://chromeflix.net/chrome-extension-search-engine/")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cookie", "__cfduid=d269c76a918082b4515e6ae9592ba302c1597037948; _ga=GA1.2.1544050276.1597037957; _gid=GA1.2.1406445235.1597037957; _fbp=fb.1.1597037956707.1473496876; signup-email=ss%40gmail.com; _gat_gtag_UA_159669744_2=1")

		res, err := client.Do(req)
		defer res.Body.Close()
		start = start + 100

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}

		var chromeExtension ChromeExtensionResponse
		err = json.Unmarshal(body, &chromeExtension)
		if err != nil {
			fmt.Println("whoops:", err)
		}
		if firstApiCall {
			recordsTotal = (chromeExtension.RecordsTotal) / 100
			firstApiCall = false
		}
		j := 0
		for j < 100 {
			finalValues = append(finalValues, chromeExtension.Data[j])

			j++
		}
		// time.Sleep(2 * time.Second)
		i++
	}
	return finalValues
}
