package gpx

import (
	"encoding/xml"
	"io/ioutil"
)

type TrkPt struct {
	Lat     string `xml:"lat,attr"`
	Lon     string `xml:"lon,attr"`
	Element string `xml:"ele"`
	Time    string `xml:"time"`
}

type TrkSeg struct {
	TrkPts []TrkPt `xml:"trkpt"`
}

type Trk struct {
	Name   string `xml:"name"`
	TrkSeg TrkSeg `xml:"trkseg"`
}

type Metadata struct {
	Time string `xml:"time"`
}

type Document struct {
	Version  string   `xml:"version,attr"`
	Creator  string   `xml:"creator,attr"`
	Metadata Metadata `xml:"metadata"`
	Trk      Trk      `xml:"trk"`
}

func FromFile(path string) (Document, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Document{}, err
	}

	var document Document
	if err := xml.Unmarshal(data, &document); err != nil {
		return Document{}, err
	}

	return document, nil
}
