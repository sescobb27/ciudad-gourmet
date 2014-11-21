package log

import (
    "github.com/stretchr/testify/assert"
    "os"
    "sync"
    "testing"
    "time"
    "unsafe"
)

const (
    msg = "hello"
)

var (
    now = time.Now().Local().Format("Jan-2-2006")
)

func TestCreateFileIfNotExist(t *testing.T) {
    t.Parallel()
    path := "./TestCreateFileIfNotExist"
    file, err := createFileIfNotExist(path)
    assert.NoError(t, err)
    file.Close()
    assertFileExist(t, path)
}

func TestFormatFileName(t *testing.T) {
    t.Parallel()
    path := "./TestFormatFileName"
    formatedFileName := formatFileName(path, "TEST")
    assert.Equal(t, path+"-TEST-"+now, formatedFileName)
}

func TestNewFile(t *testing.T) {
    t.Parallel()
    path := "./TestNewFile"
    infoFilePath := path + "-INFO-" + now
    errorFilePath := path + "-ERROR-" + now
    warningFilePath := path + "-WARNING-" + now
    // testing if the NewFile method creates a file with the pattern
    // /PATH/TO/FILE-INFO-Month-Day-Year
    // then if PASS remove the file
    infoFile, err := NewFile("INFO", path)
    assert.NoError(t, err)
    infoFile.Close()
    assertFileExist(t, infoFilePath)

    // testing if the NewFile method creates a file with the pattern
    // /PATH/TO/FILE-ERROR-Month-Day-Year
    // then if PASS remove the file
    errorFile, err := NewFile("ERROR", path)
    assert.NoError(t, err)
    errorFile.Close()
    assertFileExist(t, errorFilePath)

    // testing if the NewFile method creates a file with the pattern
    // /PATH/TO/FILE-WARNING-Month-Day-Year
    // then if PASS remove the file
    warningFile, err := NewFile("WARNING", path)
    assert.NoError(t, err)
    warningFile.Close()
    assertFileExist(t, warningFilePath)
}

func TestNewLogFactory(t *testing.T) {
    t.Parallel()
    path := "./TestNewLogFactory"
    logFactory, err := NewLogFactory(path)
    assert.NoError(t, err)
    const (
        msg = "hello"
    )

    infoFilePath := path + "-INFO-" + now
    errorFilePath := path + "-ERROR-" + now
    warningFilePath := path + "-WARNING-" + now
    // testing if the info log file is not empty after a send msg
    // then if PASS, delete the file
    logFactory.Info(msg)
    assertFileIsNotEmpty(t, infoFilePath)

    // testing if the error log file is not empty after a send msg
    // then if PASS, delete the file
    logFactory.Error(msg)
    assertFileIsNotEmpty(t, errorFilePath)

    // testing if the warning log file is not empty after a send msg
    // then if PASS, delete the file
    logFactory.Warning(msg)
    assertFileIsNotEmpty(t, warningFilePath)
}

func TestCronJobRun(t *testing.T) {
    t.Parallel()
    path := "./TestCronJobRun"
    infoFilePath := path + "-INFO-" + now
    errorFilePath := path + "-ERROR-" + now
    warningFilePath := path + "-WARNING-" + now

    logFactory, err := NewLogFactory(path)
    assert.NoError(t, err)

    oldInfoLog := unsafe.Pointer(logFactory.InfoLog.log)
    oldErrorLog := unsafe.Pointer(logFactory.ErrorLog.log)
    oldWarningLog := unsafe.Pointer(logFactory.WarningLog.log)

    logFactory.Info(msg)
    logFactory.Error(msg)
    logFactory.Warning(msg)

    logFactory.Run()
    newInfoLog := unsafe.Pointer(&logFactory.InfoLog.log)
    newErrorLog := unsafe.Pointer(&logFactory.ErrorLog.log)
    newWarningLog := unsafe.Pointer(&logFactory.WarningLog.log)
    assert.NotEqual(t, oldInfoLog, newInfoLog)
    assert.NotEqual(t, oldErrorLog, newErrorLog)
    assert.NotEqual(t, oldWarningLog, newWarningLog)

    logFactory.Info(msg)
    logFactory.Error(msg)
    logFactory.Warning(msg)

    assertFileIsNotEmpty(t, infoFilePath)
    assertFileIsNotEmpty(t, errorFilePath)
    assertFileIsNotEmpty(t, warningFilePath)
}

func TestMultipleGoRoutinesWritting(t *testing.T) {
    t.Parallel()
    path := "./TestMultipleGoRoutinesWritting"
    infoFilePath := path + "-INFO-" + now
    errorFilePath := path + "-ERROR-" + now
    warningFilePath := path + "-WARNING-" + now

    logFactory, err := NewLogFactory(path)
    assert.NoError(t, err)
    var wg sync.WaitGroup
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func() {
            logFactory.Info(msg)
            wg.Done()
        }()
        wg.Add(1)
        go func() {
            logFactory.Error(msg)
            wg.Done()
        }()
        wg.Add(1)
        go func() {
            logFactory.Warning(msg)
            wg.Done()
        }()
    }
    wg.Wait()
    assertFileIsNotEmpty(t, infoFilePath)
    assertFileIsNotEmpty(t, errorFilePath)
    assertFileIsNotEmpty(t, warningFilePath)
}

func assertFileExist(t *testing.T, path string) {
    _, err := os.Stat(path)
    assert.NoError(t, err)
    err = os.Remove(path)
    if !os.IsNotExist(err) {
        assert.NoError(t, err)
    }
}

func assertFileIsNotEmpty(t *testing.T, path string) {
    var (
        file *os.File
        err  error
    )
    buffer := make([]byte, 1024)
    file, err = os.Open(path)
    file.Read(buffer)
    assert.NotEmpty(t, buffer)
    err = os.Remove(path)
    if !os.IsNotExist(err) {
        assert.NoError(t, err)
    }
}
