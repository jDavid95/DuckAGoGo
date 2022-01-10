package translate

import (
	"fmt"
	"bufio"
	"os"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Translate() {
	var translatedInput map[string]interface{}
	var srcText, srcLanguage, targetLanguage string
	var APIkey = " " //<------- API key contact me if you'd like to run

	
	//User enters the source language of the text
	scanSrcLang := bufio.NewScanner(os.Stdin)
	fmt.Print("Translate , Enter source language > ")
	scanSrcLang.Scan()
	srcLanguage = scanSrcLang.Text()

	//User enters the target language
	scanTargetLang := bufio.NewScanner(os.Stdin)
	fmt.Print("Translate , Enter target language > ")
	scanTargetLang.Scan()
	targetLanguage = scanTargetLang.Text()

	//User enters the text to be translated
	scanTextToTranslate := bufio.NewScanner(os.Stdin)
	fmt.Print("Translate , Enter source text > ")
	scanTextToTranslate.Scan()
	srcText = scanTextToTranslate.Text()

	
	
	postBody, _ := json.Marshal(map[string]string{
		"q": srcText,
		"source": srcLanguage,
		"target": targetLanguage,
		"format": "text",
		"api_key": APIkey,
	})
	responseBody := bytes.NewBuffer(postBody)
	
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://libretranslate.com/translate", "application/json", responseBody)
	
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	
	defer resp.Body.Close()
	
	//Read the response body
	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal([]byte(body), &translatedInput); err != nil{
			
		fmt.Println(err)
		fmt.Println("no")
			
	}

	fmt.Println(translatedInput["translatedText"])
}
