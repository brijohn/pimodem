package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-systemd/journal"
	"github.com/rs/zerolog"
	"io"
	"log/syslog"
	"sort"
	"strconv"
	"sync"
)

var loggingBufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 100))
	},
}

type syslogWriter struct {
	w *syslog.Writer
}

func needsQuote(s string) bool {
	for i := range s {
		if s[i] < 0x20 || s[i] > 0x7e || s[i] == ' ' || s[i] == '\\' || s[i] == '"' {
			return true
		}
	}
	return false
}

func formatMessageFromEvent(buf *bytes.Buffer, event map[string]interface{}) {
	fmt.Fprintf(buf, "[%s] %s", event[zerolog.LevelFieldName], event[zerolog.MessageFieldName])
	fields := make([]string, 0, len(event))
	for field := range event {
		switch field {
		case zerolog.LevelFieldName, zerolog.TimestampFieldName, zerolog.MessageFieldName:
			continue
		}
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		fmt.Fprintf(buf, " %s=", field)
		switch value := event[field].(type) {
		case string:
			if needsQuote(value) {
				buf.WriteString(strconv.Quote(value))
			} else {
				buf.WriteString(value)
			}
		case json.Number:
			fmt.Fprint(buf, value)
		default:
			b, err := json.Marshal(value)
			if err != nil {
				fmt.Fprintf(buf, "[error: %v]", err)
			} else {
				fmt.Fprint(buf, string(b))
			}
		}
	}
	buf.WriteByte('\n')
}

func NewSyslogWriter() io.Writer {
	w, _ := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "")
	return syslogWriter{w}
}

func (s syslogWriter) syslogLevelFunction(level string) func(m string) error {
	lvl, _ := zerolog.ParseLevel(level)
	switch lvl {
	case zerolog.DebugLevel:
		return s.w.Debug
	case zerolog.InfoLevel:
		return s.w.Info
	case zerolog.WarnLevel:
		return s.w.Warning
	case zerolog.ErrorLevel:
		return s.w.Err
	case zerolog.FatalLevel:
		return s.w.Emerg
	case zerolog.PanicLevel:
		return s.w.Crit
	case zerolog.NoLevel:
		return s.w.Info
	default:
		panic("Invalid Level")
	}
	return nil
}

func (s syslogWriter) Write(p []byte) (n int, err error) {
	var event map[string]interface{}
	syslogWrite := s.syslogLevelFunction("")
	d := json.NewDecoder(bytes.NewReader(p))
	d.UseNumber()
	err = d.Decode(&event)
	if err != nil {
		return
	}
	buf := loggingBufPool.Get().(*bytes.Buffer)
	defer loggingBufPool.Put(buf)
	formatMessageFromEvent(buf, event)
	if l, ok := event[zerolog.LevelFieldName].(string); ok {
		syslogWrite = s.syslogLevelFunction(l)
	}
	n = len(buf.Bytes())
	err = syslogWrite(buf.String())
	buf.Reset()
	return
}

type journalWriter struct {
}

func NewJournalWriter() io.Writer {
	return journalWriter{}
}

func levelToJournalPriority(level string) journal.Priority {
	lvl, _ := zerolog.ParseLevel(level)

	switch lvl {
	case zerolog.DebugLevel:
		return journal.PriDebug
	case zerolog.InfoLevel:
		return journal.PriInfo
	case zerolog.WarnLevel:
		return journal.PriWarning
	case zerolog.ErrorLevel:
		return journal.PriErr
	case zerolog.FatalLevel:
		return journal.PriCrit
	case zerolog.PanicLevel:
		return journal.PriEmerg
	case zerolog.NoLevel:
		return journal.PriNotice
	}
	panic("Invalid Level")
}

func (j journalWriter) Write(p []byte) (n int, err error) {
	if !journal.Enabled() {
		err = fmt.Errorf("Cannot connect to journalD!!")
		return
	}
	var event map[string]interface{}
	priority := journal.PriNotice
	args := make(map[string]string, 0)
	d := json.NewDecoder(bytes.NewReader(p))
	d.UseNumber()
	err = d.Decode(&event)
	if err != nil {
		return
	}
	buf := loggingBufPool.Get().(*bytes.Buffer)
	defer loggingBufPool.Put(buf)
	formatMessageFromEvent(buf, event)
	if l, ok := event[zerolog.LevelFieldName].(string); ok {
		priority = levelToJournalPriority(l)
	}
	args["JSON"] = string(p)
	n = len(buf.Bytes())
	err = journal.Send(buf.String(), priority, args)
	buf.Reset()
	return
}
