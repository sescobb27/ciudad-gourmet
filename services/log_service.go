package services

import (
        "errors"
        "fmt"
        "github.com/robfig/cron"
        "log"
        "os"
        "sync"
        "time"
)

type InfoLog struct {
        InfoChan chan []byte
        log      *log.Logger
        file     *os.File
}

type WarningLog struct {
        WarningChan chan []byte
        log         *log.Logger
        file        *os.File
}

type ErrorLog struct {
        ErrorChan chan []byte
        log       *log.Logger
        file      *os.File
}

type LogFactory struct {
        mux        sync.RWMutex
        InfoLog    InfoLog
        WarningLog WarningLog
        ErrorLog   ErrorLog
        path       string
}

func createFileIfNotExist(path string) (*os.File, error) {
        var file *os.File
        _, err := os.Stat(path)
        if err != nil {
                if os.IsNotExist(err) {
                        file, err = os.Create(path)
                        if err != nil {
                                return nil, err
                        }
                } else {
                        return nil, err
                }
        }
        return file, nil
}

func formatFileName(path, prefix string) string {
        date := time.Now().Local().Format("Jan-2-2006")
        return fmt.Sprintf("%s-%s-%s", path, prefix, date)
}

func NewWriter(tag, path string) (*os.File, error) {
        var (
                file    *os.File
                err     error
        )
        switch tag {
        case "INFO":
                file, err = createFileIfNotExist(formatFileName(path, "INFO"))
        case "ERROR":
                file, err = createFileIfNotExist(formatFileName(path, "ERROR"))
        case "WARNING":
                file, err = createFileIfNotExist(formatFileName(path, "WARNING"))
        default:
                return nil, errors.New("Invalid Tag " + tag + " Expected INFO, ERROR or WARNING")
        }
        return file, err
}

func NewLogFactory(path string) (*LogFactory, error) {
        var (
                infoFile, warningFile, errorFile *os.File
                err                              error
        )
        infoFile, err = NewWriter("INFO", path)
        if err != nil {
                return nil, err
        }
        warningFile, err = NewWriter("WARNING", path)
        if err != nil {
                return nil, err
        }
        errorFile, err = NewWriter("ERROR", path)
        if err != nil {
                return nil, err
        }

        logFactory := &LogFactory{
                InfoLog:    NewInfoLog(infoFile),
                WarningLog: NewWarningLog(warningFile),
                ErrorLog:   NewErrorLog(errorFile),
                path:       path,
        }
        go logFactory.listen()
        cronJob := cron.New()
        cronJob.AddJob("@daily", logFactory)
        return logFactory, nil
}

func (l *LogFactory) listen() {
        var infoMsg, warningMsg, errorMsg []byte
        for {
                select {
                case infoMsg = <-l.InfoLog.InfoChan:
                        l.InfoLog.log.Println(string(infoMsg))
                case warningMsg = <-l.ErrorLog.ErrorChan:
                        l.WarningLog.log.Println(string(warningMsg))
                case errorMsg = <-l.WarningLog.WarningChan:
                        l.ErrorLog.log.Println(string(errorMsg))
                }
        }
}

func (l *LogFactory) Run() {
        l.mux.Lock()
        defer l.mux.Unlock()

        var (
                infoFile    *os.File
                warningFile *os.File
                errorFile   *os.File
                err         error
        )

        infoFile, err = NewWriter("INFO", l.path)
        if err != nil {
                l.ErrorLog.ErrorChan <- []byte(err.Error())
                return
        }
        warningFile, err = NewWriter("WARNING", l.path)
        if err != nil {
                l.ErrorLog.ErrorChan <- []byte(err.Error())
                return
        }
        errorFile, err = NewWriter("ERROR", l.path)
        if err != nil {
                l.ErrorLog.ErrorChan <- []byte(err.Error())
                return
        }

        l.InfoLog.file.Close()
        l.WarningLog.file.Close()
        l.ErrorLog.file.Close()

        l.InfoLog.log = NewLog("INFO: ", infoFile)
        l.WarningLog.log = NewLog("WARNING: ", warningFile)
        l.ErrorLog.log = NewLog("ERROR: ", errorFile)

}

func NewLog(tag string, file *os.File) *log.Logger {
        return log.New(file, tag, log.Ldate|log.Ltime|log.Lshortfile)
}

func NewInfoLog(file *os.File) InfoLog {
        return InfoLog{
                InfoChan: make(chan []byte, 10),
                log: log.New(file, "INFO: ",
                        log.Ldate|log.Ltime),
                file:   file,
        }
}

func NewWarningLog(file *os.File) WarningLog {
        return WarningLog{
                WarningChan: make(chan []byte, 10),
                log: log.New(file, "WARNING: ",
                        log.Ldate|log.Ltime|log.Lshortfile),
                file:   file,
        }
}

func NewErrorLog(file *os.File) ErrorLog {
        return ErrorLog{
                ErrorChan: make(chan []byte, 10),
                log: log.New(file, "ERROR: ",
                        log.Ldate|log.Ltime|log.Lshortfile),
                file:   file,
        }
}
