package services

import (
        "github.com/stretchr/testify/assert"
        "os"
        "testing"
        "time"
)

const (
        path    = "./log.test"
        msg     = "hello"
)

var (
        now             = time.Now().Local().Format("Jan-2-2006")
        infoFilePath    = path + "-INFO-" + now
        errorFilePath   = path + "-ERROR-" + now
        warningFilePath = path + "-WARNING-" + now
)

func TestCreateFileIfNotExist(t *testing.T) {
        t.Parallel()
        file, err := createFileIfNotExist(path)
        assert.NoError(t, err)
        file.Close()
        assertFileExist(t, path)
}

func TestFormatFileName(t *testing.T) {
        t.Parallel()
        formatedFileName := formatFileName(path, "TEST")
        assert.Equal(t, path+"-TEST-"+now, formatedFileName)
}

func TestNewWriter(t *testing.T) {
        // testing if the NewWriter method creates a file with the pattern
        // /PATH/TO/FILE-INFO-Month-Day-Year
        // then if PASS remove the file
        infoFile, err := NewWriter("INFO", path)
        assert.NoError(t, err)
        infoFile.Close()
        assertFileExist(t, infoFilePath)

        // testing if the NewWriter method creates a file with the pattern
        // /PATH/TO/FILE-ERROR-Month-Day-Year
        // then if PASS remove the file
        errorFile, err := NewWriter("ERROR", path)
        assert.NoError(t, err)
        errorFile.Close()
        assertFileExist(t, errorFilePath)

        // testing if the NewWriter method creates a file with the pattern
        // /PATH/TO/FILE-WARNING-Month-Day-Year
        // then if PASS remove the file
        warningFile, err := NewWriter("WARNING", path)
        assert.NoError(t, err)
        warningFile.Close()
        assertFileExist(t, warningFilePath)
}

func TestNewLogFactory(t *testing.T) {
        logFactory, err := NewLogFactory(path)
        assert.NoError(t, err)
        const (
                msg = "hello"
        )

        buffer := make([]byte, len(msg))
        // testing if the info log file is not empty after a send msg
        // then if PASS, delete the file
        logFactory.InfoLog.InfoChan <- []byte(msg)
        assertFileIsNotEmpty(t, infoFilePath, buffer)

        // testing if the error log file is not empty after a send msg
        // then if PASS, delete the file
        logFactory.ErrorLog.ErrorChan <- []byte(msg)
        assertFileIsNotEmpty(t, errorFilePath, buffer)

        // testing if the warning log file is not empty after a send msg
        // then if PASS, delete the file
        logFactory.WarningLog.WarningChan <- []byte(msg)
        assertFileIsNotEmpty(t, warningFilePath, buffer)
}

func TestCronJobRun(t *testing.T) {

        logFactory, err := NewLogFactory(path)
        assert.NoError(t, err)

        oldInfoLog := logFactory.InfoLog
        oldErrorLog := logFactory.ErrorLog
        oldWarningLog := logFactory.WarningLog

        logFactory.InfoLog.InfoChan <- []byte(msg)
        logFactory.ErrorLog.ErrorChan <- []byte(msg)
        logFactory.WarningLog.WarningChan <- []byte(msg)

        logFactory.Run()

        assert.NotEqual(t, oldInfoLog, logFactory.InfoLog)
        assert.NotEqual(t, oldErrorLog, logFactory.ErrorLog)
        assert.NotEqual(t, oldWarningLog, logFactory.WarningLog)

        logFactory.InfoLog.InfoChan <- []byte(msg)
        logFactory.ErrorLog.ErrorChan <- []byte(msg)
        logFactory.WarningLog.WarningChan <- []byte(msg)

        buffer := make([]byte, len(msg))
        assertFileIsNotEmpty(t, infoFilePath, buffer)
        assertFileIsNotEmpty(t, errorFilePath, buffer)
        assertFileIsNotEmpty(t, warningFilePath, buffer)
}

func assertFileExist(t *testing.T, path string) {
        _, err := os.Stat(path)
        assert.NoError(t, err)
        err = os.Remove(path)
        assert.NoError(t, err)
}

func assertFileIsNotEmpty(t *testing.T, path string, buffer []byte) {
        var (
                file    *os.File
                err     error
        )
        file, err = os.Open(path)
        file.Read(buffer)
        assert.NotEmpty(t, buffer)
        err = os.Remove(path)
        assert.NoError(t, err)
}
