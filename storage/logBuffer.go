package storage

type logBuffer struct{}

//in order to fetch a logBuffer, call this function
func GetLogBuffer() *logBuffer {
	lb := &logBuffer{}
	return lb
}
