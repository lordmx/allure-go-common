package beans

import (
	"time"
)

type Step struct {
	Parent *Step `xml:"-"`

	Status      string        `xml:"status,attr"`
	Start       int64         `xml:"start,attr"`
	Stop        int64         `xml:"stop,attr"`
	Name        string        `xml:"name"`
	Steps       []*Step       `xml:"steps>step"`
	Attachments []*Attachment `xml:"attachments>attachment"`
}

func NewStep(name string, start time.Time) *Step {
	test := new(Step)
	test.Name = name

	if !start.IsZero() {
		test.Start = start.UnixNano() / int64(time.Millisecond)
	} else {
		test.Start = time.Now().UnixNano() / int64(time.Millisecond)
	}

	return test
}

func (s *Step) End(status string, end time.Time) {
	if !end.IsZero() {
		s.Stop = end.UnixNano() / int64(time.Millisecond)
	} else {
		s.Stop = time.Now().UnixNano() / int64(time.Millisecond)
	}
	s.Status = status
}

func (s *Step) AddStep(step *Step) {
	if step != nil {
		s.Steps = append(s.Steps, step)
	}
}
