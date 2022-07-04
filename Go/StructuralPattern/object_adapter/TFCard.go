package objectadapter

type TFCard interface {
	readTF() string
	writeTF(msg string)
}