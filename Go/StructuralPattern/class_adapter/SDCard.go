package classadapter

type SDCard interface {
	readSD() string
	writeSD(msg string)
}