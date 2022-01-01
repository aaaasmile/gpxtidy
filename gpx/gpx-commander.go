package gpx

import (
	"fmt"
	"log"
	"path"
)

type Commander struct {
	Dir        string
	SourcePath string
	TargetPath string
	Cmd        string
	sourcePts  []TrkPt
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

	return nil
}
