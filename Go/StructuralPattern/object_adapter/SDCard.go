package objectadapter

type SDCard interface {
	readSD() string
	writeSD(msg string)
}