package gpx

import "fmt"

type Commander struct {
	Dir        string
	SourcePath string
	TargetPath string
	Cmd        string
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
	return nil
}
