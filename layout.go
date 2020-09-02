package log4go

import (
	"bytes"
	"strings"
)

type sectionType int

const (
	kMsg sectionType = iota
	kLevel
	kSource
	kCategory
	kTime
	kString
)

type section struct {
	t sectionType
	v interface{}
}

type layoutConf struct {
	sections []section
}

// Known format codes:
// %T - DateTime with format string, default format is (2006-01-03 15:04:05.000)
//  	eg. %T{2006-01-03 15:04:05}
// %L - Level (FNST, FINE, DEBG, TRAC, WARN, EROR, CRIT)
// %C - Category
// %S - Source
// %M - Message
// Ignores unknown formats
// Recommended: "[%F] %L %C (%S) %M"
func createLayout(layout string) *layoutConf {
	// "[%T %D] [%C] [%L] (%S) %M"
	lc := &layoutConf{}

	r := []rune(layout)

	type stage int
	const (
		kString = iota
		kFindFlag


	)

	pre := ""
	for {
		idx := strings.IndexRune(layout, '%')
		if idx < 0 {
			if len(layout) > 0 {
				lc.sections = append(lc.sections, section{t: kString, v: layout})
				layout = ""
			}
			break
		}

		pre += layout[:idx]

	}

	st := kString
	str := ""
	for i:=0; i< len(r); i++ {
		switch r[i] {
		case '%':
		default:
			str += string(r[i])
		}
		if r[i] == '%' {
			switch st {
			case kString:
				if len(str) > 0 {

				}
			}
		}
	}

	return lc
}

func formatLog(layout *layoutConf, rec *LogRecord) string {
	if rec == nil {
		return "<nil>"
	}

	out := bytes.NewBuffer(make([]byte, 0, 64))

	for i := 0; i< len(layout.sections); i++ {
		switch layout.sections[i].t {
		case kTime:
			out.WriteString(rec.Created.Format(layout.sections[i].v.(string)))
		case kCategory:
			out.WriteString(rec.Category)
		case kLevel:
			out.WriteString(levelStrings[rec.Level])
		case kSource:
			out.WriteString(rec.Source)
		case kMsg:
			out.WriteString(rec.Message)
		case kString:
			out.WriteString(layout.sections[i].v.(string))
		}
	}

	return out.String()
}
