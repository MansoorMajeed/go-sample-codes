// https://github.com/Sirupsen/logrus
package main

import (
    log "github.com/Sirupsen/logrus"
    "os"
    "fmt"
)

func main(){
    var filename string = "logfile.log"
    f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
    Formatter := new(log.TextFormatter)
    Formatter.TimestampFormat = "02-01-2006 15:04:05"
    Formatter.FullTimestamp = true
    log.SetFormatter(Formatter)
    if err != nil {
        // Cannot open log file. Logging to stderr
        fmt.Println(err)
    }else{
        log.SetOutput(f)
    }

    log.Info("Some info. Earth is not flat")
    log.Warning("This is a warning")
    log.Error("Not fatal. An error. Won't stop execution")
    log.Fatal("MAYDAY MAYDAY MAYDAY")
    log.Panic("Do not panic")
}
