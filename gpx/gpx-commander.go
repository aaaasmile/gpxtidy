package gpx

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type Commander struct {
	Dir        string
	SourcePath string
	TargetPath string
	Cmd        string
	AbsTraget  bool
	sourcePts  []TrkPt
	gpxTime    string
	trackName  string
}

func NewCommander(cmd, dir, source, target string) *Commander {
	res := Commander{
		Dir:        dir,
		SourcePath: source,
		TargetPath: target,
		Cmd:        cmd,
	}
	return &res
}

func (c *Commander) Run() error {
	switch c.Cmd {
	case "remext":
		return c.doRemoveExtension()
	default:
		return fmt.Errorf("err: command not recognized: %s", c.Cmd)
	}
}

func (c *Commander) doRemoveExtension() error {
	err := c.parseSource()
	if err != nil {
		return err
	}

	err = c.writeTarget()
	if err != nil {
		return err
	}

	return nil
}

func (c *Commander) parseSource() error {
	path := path.Join(c.Dir, c.SourcePath)
	log.Println("Processing source file: ", path)

	dt, err := FromFile(path)
	if err != nil {
		return err
	}
	ptlen := len(dt.Trk.TrkSeg.TrkPts)
	log.Println("Gpx: ", dt.Version, dt.Creator)
	log.Printf("Track name: %s. Number of points %d", dt.Trk.Name, ptlen)

	if ptlen == 0 {
		return fmt.Errorf("expect track points but nothing was recognized")
	}

	c.sourcePts = make([]TrkPt, 0)
	c.sourcePts = append(c.sourcePts, dt.Trk.TrkSeg.TrkPts...)
	c.trackName = dt.Trk.Name
	c.gpxTime = dt.Metadata.Time

	return nil
}

type UnitInfo struct {
	ID                 string
	SourceIsEmpty      bool
	TargetHasStateAttr bool
	NoteGen            string
	TargetStateAttr    string
	Source             string
	Target             string
	TransObjTarget     string
	TransHasObjTarget  bool
	NoteDev            string
}

func (c *Commander) writeTarget() error {
	ctx := struct {
		GpxTime   string
		Trackname string
		TrkPts    []TrkPt
	}{
		GpxTime:   c.gpxTime,
		Trackname: c.trackName,
		TrkPts:    c.sourcePts,
	}

	templFileName := "templates/track.gpx_templ"

	pathOut := c.TargetPath
	if !c.AbsTraget {
		pathOut = path.Join(c.Dir, c.TargetPath)
	}

	f, err := os.Create(pathOut)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Println("Prepare to execute template ", templFileName)
	tmplXliff := template.Must(template.New("Main").ParseFiles(templFileName))

	var partHeader, partPlainContent, partFooter bytes.Buffer
	if err := tmplXliff.ExecuteTemplate(&partPlainContent, "body", ctx); err != nil {
		return err
	}

	if err := tmplXliff.ExecuteTemplate(&partHeader, "mainheader", ctx); err != nil {
		return err
	}
	if err := tmplXliff.ExecuteTemplate(&partFooter, "mainfooter", ctx); err != nil {
		return err
	}
	f.Write(partHeader.Bytes())
	f.Write(partPlainContent.Bytes())
	_, err = f.Seek(-2, os.SEEK_END) // eat spaces at the end of body before write the footer
	if err != nil {
		return err
	}
	f.Write(partFooter.Bytes())

	log.Println("File written: ", pathOut)

	return nil
}

type vContentMapping struct {
	Decoded string
	Encoded string
}

func escapeXmlString(s string) string {
	var vContentMappings = []vContentMapping{
		{Decoded: `&`, Encoded: `&amp;`},
		{Decoded: `<`, Encoded: `&lt;`},
		{Decoded: `>`, Encoded: `&gt;`},
	}

	content := s
	//fmt.Println("** s", s)
	for _, mapping := range vContentMappings {
		content = strings.Replace(content, mapping.Decoded, mapping.Encoded, -1)
		//fmt.Println("** content: ", content, mapping.Encoded, mapping.Decoded)
	}
	return content
}
