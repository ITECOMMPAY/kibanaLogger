package kibanaLogger

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"
)

var logFile *os.File

func init() {
	logTo := flag.String("kibanaLog", "kibana.log", "Kibana logs")
	f, err := os.OpenFile(
		*logTo,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	logFile = f

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(logFile)
	log.SetFlags(0)
}

type logData struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Hostname  string `json:"hostname"`
}

// Info logging kibana message with status OK
func Info(message string) {
	writeLog("OK", message)
}

// Warning logging kibana message with status WRN
func Warning(message string) {
	writeLog("WRN", message)
}

// Error logging kibana message with status ERR
func Error(message string) {
	writeLog("ERR", message)
}

// Close log file
func Close() {
	_ = logFile.Sync()
	err := logFile.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func writeLog(status string, message string) {
	hostname, err := os.Hostname()

	if err != nil {
		log.Fatalf("can't get hostname")
	}

	logData := &logData{
		Status:    status,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339Nano),
		Hostname:  hostname,
	}

	logString, err := json.Marshal(logData)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(logString))
}
