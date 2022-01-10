package main

import (
	"DuckAGoGo/translate"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"bufio"
	"os"
	"log"
)

//Parse the JSON results so they are easily accesible to manipulate.
type Results struct {
	Abstract string `json:"Abstract"`
	Source string `json:"AbstractSource"`
	AbstractText string `json:"AbstractText"`
	AbstractURL string `json:"AbstractURL"`
}


//This function formats the query to be URL friendly by removing the spaces
func validateAndFormatQuery(inputQuery *string){
	terminateProgramString := ".."
	if string(*inputQuery) == terminateProgramString {	
		log.Fatalf("Bye")	
	}
	if strings.EqualFold(string(*inputQuery), "translate") {
		translate.Translate()
		log.Fatalf("Bye")
	}
	*inputQuery = strings.ReplaceAll(*inputQuery, " ", "%20")

}

func main() {
	var searchResult Results
	var query string
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("DuckAGoGo > ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading

	query = scanner.Text()

	validateAndFormatQuery(&query)

	for(len(query) < 1){
		fmt.Println("Invalid input, try again")
		scanner.Scan()
		query = scanner.Text()
		validateAndFormatQuery(&query)
	}

	url := "https://duckduckgo-duckduckgo-zero-click-info.p.rapidapi.com/?q="+query+"&no_html=1&no_redirect=1&skip_disambig=1&format=json"
		
	//fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
		

	req.Header.Add("x-rapidapi-host", "duckduckgo-duckduckgo-zero-click-info.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", " ") //<------ API key, contact me if you want to run.
		

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
		
	body, _ := ioutil.ReadAll(res.Body)
		
	if err := json.Unmarshal([]byte(body), &searchResult); err != nil{
			
		fmt.Println(err)
			
	}

	fmt.Println(searchResult.Source)
	fmt.Println( searchResult.Abstract)


}
