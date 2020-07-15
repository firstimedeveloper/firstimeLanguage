package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/firstimedeveloper/firstimeLanguage/views"

	"github.com/pkg/errors"
)

func NewVideo() *Video {
	return &Video{
		NewView:  views.NewView("bootstrap", "videos/new"),
		ShowView: views.NewView("bootstrap", "videos/show"),
	}
}

type Video struct {
	Transcript Transcript
	NewView    *views.View
	ShowView   *views.View
}

// New is used to render the form where the end user can enter a link
// to get the subtitles for that video.
//
// GET /new
func (v *Video) New(w http.ResponseWriter, r *http.Request) {
	v.NewView.Render(w, r, nil)
}

type VideoForm struct {
	Link string `schema:"link"`
}

// Show is used to show the end user the appropriate information after
// they enter the form with an appropriate link.
//
// POST /new
func (v *Video) Show(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form VideoForm
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		v.NewView.Render(w, r, vd)
		return
	}
	err := v.Transcript.parseSubtitles(form.Link)
	if err != nil {
		vd.SetAlert(err)
		v.NewView.Render(w, r, vd)
		log.Println(err)
		return
	}
	//fmt.Fprintf(w, "Retrieving video subtitles from %s", form.Link)
	vd.Yield = v.Transcript.Text
	v.ShowView.Render(w, r, vd)
}

type Time string

func (t Time) String() string {
	temp, err := strconv.ParseFloat(string(t), 64)
	if err != nil {
		panic(err)
	}
	if temp >= 60 {
		min := math.Floor(temp / 60)
		return fmt.Sprintf("%1.f:%.2f", min, temp-60*min)
	}
	return fmt.Sprintf("%.2f", temp)
}

type Text struct {
	Line  string `xml:",chardata"`
	Start Time   `xml:"start,attr"`
	Dur   Time   `xml:"dur,attr"`
	End   Time   `-`
}

type Transcript struct {
	Text []Text `xml:"text"`
}

func (t *Transcript) parseSubtitles(link string) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.Errorf("Response status: %s", resp.Status)
	}

	subXML, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	subXML = []byte(strings.ReplaceAll(string(subXML), "\n", " "))

	var sub Transcript
	if err := xml.Unmarshal(subXML, &sub); err != nil {
		return err
	}
	t.Text = sub.Text
	for i := range t.Text {
		tempStart, err := strconv.ParseFloat(string(t.Text[i].Start), 64)
		if err != nil {
			return errors.Errorf("Unable to parse tempStart: %v", err)
		}
		tempDur, err := strconv.ParseFloat(string(t.Text[i].Dur), 64)
		if err != nil {
			return errors.Errorf("Unable to parse tempStart: %v", err)
		}

		num := tempStart + tempDur
		t.Text[i].End = Time(strconv.FormatFloat(num, 'f', 2, 64))
	}
	return nil
}

func (t *Transcript) String() string {
	var str strings.Builder
	for _, v := range t.Text {
		str.WriteString(fmt.Sprintf("%s-%s: %s\n", v.Start, v.End, v.Line))
	}

	return str.String()
}
