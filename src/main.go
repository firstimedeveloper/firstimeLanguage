package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/xml"
)

type Transcript struct {
	XMLName  xml.Name `xml:"transcript"`
	Chardata string   `xml:",chardata"`
	Text     []struct {
		Text  string `xml:",chardata"`
		Start string `xml:"start,attr"`
		Dur   string `xml:"dur,attr"`
	} `xml:"text"`
} 
/*
func (t Transcript) String() string {
	
	return fmt.Sprintf("%s %s %s\n", t.Text.Start, t.Text.Dur, t.Text.Text)
}*/

func main() {
	resp, err := http.Get("http://video.google.com/timedtext?lang=de&v=dL5oGKNlR6I")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Response status:", resp.Status)
	}
	
	subtitle, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	var sub Transcript
	xml.Unmarshal(subtitle, &sub)
	for _,v := range sub.Text {
		fmt.Printf(v.Text)
	}
	
}
//http://video.google.com/timedtext?lang=de&v=dL5oGKNlR6I