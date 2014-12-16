package log

import (
    "fmt"
    "github.com/robfig/cron"
    "log"
    "os"
    "time"
)

type InfoLog struct {
    InfoChan   chan string
    UpdateChan chan *os.File
    log        *log.Logger
    file       *os.File
}

type WarningLog struct {
    WarningChan chan string
    UpdateChan  chan *os.File
    log         *log.Logger
    file        *os.File
}

type ErrorLog struct {
    ErrorChan  chan string
    UpdateChan chan *os.File
    log        *log.Logger
    file       *os.File
}

type LogFactory struct {
    InfoLog    InfoLog
    WarningLog WarningLog
    ErrorLog   ErrorLog
    path       string
}

const (
    INFO        = "INFO"
    ERROR       = "ERROR"
    WARNING     = "WARNING"
    INFO_TAG    = "INFO: "
    ERROR_TAG   = "ERROR: "
    WARNING_TAG = "WARNING: "
)

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
    } else {
        file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
        if err != nil {
            return nil, err
        }
    }
    return file, nil
}

func formatFileName(path, prefix string) string {
    date := time.Now().Local().Format("Jan-2-2006")
    return fmt.Sprintf("%s-%s-%s", path, prefix, date)
}

func NewFile(tag, path string) (*os.File, error) {
    return createFileIfNotExist(formatFileName(path, tag))
}

func NewLogFactory(path string) (*LogFactory, error) {
    var (
        infoFile, warningFile, errorFile *os.File
        err                              error
    )
    infoFile, err = NewFile(INFO, path)
    if err != nil {
        return nil, err
    }
    warningFile, err = NewFile(WARNING, path)
    if err != nil {
        return nil, err
    }
    errorFile, err = NewFile(ERROR, path)
    if err != nil {
        return nil, err
    }

    logFactory := &LogFactory{
        InfoLog:    NewInfoLog(infoFile),
        WarningLog: NewWarningLog(warningFile),
        ErrorLog:   NewErrorLog(errorFile),
        path:       path,
    }
    logFactory.listen()
    cronJob := cron.New()
    cronJob.AddJob("@daily", logFactory)
    return logFactory, nil
}

func (l *LogFactory) listen() {
    go l.InfoLog.listen()
    go l.ErrorLog.listen()
    go l.WarningLog.listen()
}

func (i *InfoLog) listen() {
    for {
        select {
        case infoMsg := <-i.InfoChan:
            i.log.Println(infoMsg)
        case file := <-i.UpdateChan:
            i.file.Close()
            i.log = NewLog(INFO_TAG, file)
            i.file = file
        }
    }
}

func (e *ErrorLog) listen() {
    for {
        select {
        case errorMsg := <-e.ErrorChan:
            e.log.Println(errorMsg)
        case file := <-e.UpdateChan:
            e.file.Close()
            e.log = NewLog(ERROR_TAG, file)
            e.file = file
        }
    }
}

func (w *WarningLog) listen() {
    for {
        select {
        case warningMsg := <-w.WarningChan:
            w.log.Println(warningMsg)
        case file := <-w.UpdateChan:
            w.file.Close()
            w.log = NewLog(WARNING_TAG, file)
            w.file = file
        }
    }
}

func (l *LogFactory) Error(msg string) {
    l.ErrorLog.ErrorChan <- msg
}

func (l *LogFactory) Info(msg string) {
    l.InfoLog.InfoChan <- msg
}

func (l *LogFactory) Warning(msg string) {
    l.WarningLog.WarningChan <- msg
}

func (l *LogFactory) Run() {
    var (
        infoFile    *os.File
        warningFile *os.File
        errorFile   *os.File
        err         error
    )
    infoFile, err = NewFile(INFO, l.path)
    if err != nil {
        l.ErrorLog.ErrorChan <- err.Error()
        return
    }
    warningFile, err = NewFile(WARNING, l.path)
    if err != nil {
        l.ErrorLog.ErrorChan <- err.Error()
        return
    }
    errorFile, err = NewFile(ERROR, l.path)
    if err != nil {
        l.ErrorLog.ErrorChan <- err.Error()
        return
    }
    l.InfoLog.UpdateChan <- infoFile
    l.WarningLog.UpdateChan <- warningFile
    l.ErrorLog.UpdateChan <- errorFile
}

func NewLog(tag string, file *os.File) *log.Logger {
    logger := log.New(file, tag, log.Ldate|log.Ltime)
    if tag != INFO_TAG {
        logger.SetFlags(log.Lshortfile)
    }
    return logger
}

func NewInfoLog(file *os.File) InfoLog {
    return InfoLog{
        InfoChan:   make(chan string, 10),
        UpdateChan: make(chan *os.File),
        log:        NewLog(INFO_TAG, file),
        file:       file,
    }
}

func NewWarningLog(file *os.File) WarningLog {
    return WarningLog{
        WarningChan: make(chan string, 10),
        UpdateChan:  make(chan *os.File),
        log:         NewLog(WARNING_TAG, file),
        file:        file,
    }
}

func NewErrorLog(file *os.File) ErrorLog {
    return ErrorLog{
        ErrorChan:  make(chan string, 10),
        UpdateChan: make(chan *os.File),
        log:        NewLog(ERROR_TAG, file),
        file:       file,
    }
}

var (
    Log *LogFactory
    err error
)

func init() {
    Log, err = NewLogFactory("./ciudad-gourmet.log")
    if err != nil {
        panic(err)
    }
}
