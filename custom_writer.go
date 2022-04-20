package main

import (
	"errors"

	"fyne.io/fyne/v2"
)

type customWriter struct {
	w fyne.URIWriteCloser
}

func (e customWriter) Write(p []byte) (int, error) {
	logPath := fyne.CurrentApp().Preferences().StringWithFallback("logFile", "app.log")

	if e.w == nil {
		opFile, err := fyne.CurrentApp().Storage().Save(logPath)
		if err != nil {
			nFile, nErr := fyne.CurrentApp().Storage().Create(logPath)
			if nErr != nil {
				fyne.LogError("Cannot log error", nErr)
				return 0, errors.New("Cannot Log")
			}
			e.w = nFile
		} else {
			e.w = opFile
		}
		defer e.w.Close()
	}
	e.w.Write(p)

	return len(p), nil
}

func NewCustomWriter() *customWriter {
	p := new(customWriter)
	return p
}
