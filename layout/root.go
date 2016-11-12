package layout

import (
	"reflect"
	"sort"

	log "github.com/Sirupsen/logrus"
	wlc "github.com/mikkeloscar/go-wlc"
)

type Root struct {
	outputs map[string]*Output
	focused *Output
}

func NewRoot() *Root {
	return &Root{
		outputs: make(map[string]*Output),
		focused: nil,
	}
}

func (r *Root) Type() ContainerType {
	return CRoot
}

func (r *Root) Geometry() *wlc.Geometry {
	return nil
}

func (r *Root) Children() []Container {
	containers := make([]Container, 0, len(r.outputs))
	for _, o := range r.outputs {
		containers = append(containers, o)
	}
	return containers
}

func (r *Root) Floating() []Container {
	return nil
}

func (r *Root) Focused() Container {
	return r.focused
}

func (r *Root) Parent() Container {
	return nil
}

func (r *Root) Fullscreen() Container {
	if r.focused != nil {
		return r.focused.Fullscreen()
	}
	return nil
}

func (r *Root) Visible() bool {
	return true
}

func (r *Root) AddChild(output Container) {
	switch o := output.(type) {
	case *Output:
		// TODO: focus new output?
		log.Debugf("Added output %s", o.Name())
		r.outputs[o.Name()] = o

		// if there is only one output then focus it.
		if len(r.outputs) == 1 {
			r.focused = o
		}
	default:
		log.Errorf("Failed to add output, invalid container type: %s", reflect.TypeOf(output))
	}
}

// SortedWorkspaces returns an aggregated sorted list of workspaces on all
// outputs.
func (r *Root) SortedWorkspaces() []*Workspace {
	var ws []*Workspace
	for _, output := range r.outputs {
		ws = append(ws, output.workspaces...)
	}

	sort.Sort(Workspaces(ws))
	return ws
}
