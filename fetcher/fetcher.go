package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Fetch(url string) ([]byte, error) {
	//resp, err := http.Get(url)
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("wrong http request: %s", err.Error())
		return nil, err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Println("resp Error")
		}
		fmt.Println(string(all))
		return nil, fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Microsecond)
	return all, nil
}
