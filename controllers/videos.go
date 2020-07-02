package controllers

import (
	"encoding/xml"
	"firstimelang/views"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

type Text struct {
	Line  string  `xml:",chardata"`
	Start float64 `xml:"start,attr"`
	Dur   float64 `xml:"dur,attr"`
	End   float64 `-`
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
		t.Text[i].End = t.Text[i].Start + t.Text[i].Dur
	}
	return nil
}

func (t *Transcript) String() string {
	var str strings.Builder
	for _, v := range t.Text {
		str.WriteString(fmt.Sprintf("%f-%f: %s\n", v.Start, v.End, v.Line))
	}

	return str.String()
}
