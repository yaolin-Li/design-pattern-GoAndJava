package classadapter

type TFCard interface {
	readTF() string
	writeTF(msg string)
}