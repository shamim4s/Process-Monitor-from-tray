package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io/fs"
    "os"
    "os/exec"
    "path/filepath"
    "sort"
    "strings"
    "time"

    "github.com/getlantern/systray"
    "golang.org/x/sys/windows"
)

// Base configuration
var (
    basePath        = `C:\Windows\filebrowser`
    processName     = "filebrowser.exe"
    batFileName     = "start.bat"
    logFileName     = `filebrowser.log`
    iconFolderName  = `img\icons\animation`
    iconPrefix      = "anim"
    iconExtension   = ".ico"
    checkInterval   = 5 * time.Second
    animationDelay  = 500 * time.Millisecond
    debugPrint      = true
)

var iconFrames [][]byte
var iconIndex int
var showLogCmd *exec.Cmd

func main() {
    systray.Run(onReady, onExit)
}

func onReady() {
    loadIcons()

    mShowLog := systray.AddMenuItem("Show Log", "Open log in console")
    mQuit := systray.AddMenuItem("Exit", "Exit the application")

    go startMonitoring()
    go animateIcon()

    go func() {
        for {
            select {
            case <-mShowLog.ClickedCh:
                showLogWindow()
            case <-mQuit.ClickedCh:
                cleanupAndExit()
            }
        }
    }()
}

func onExit() {
    cleanupAndExit()
}

func loadIcons() {
    iconPath := filepath.Join(basePath, iconFolderName)
    files := []string{}

    _ = filepath.WalkDir(iconPath, func(path string, d fs.DirEntry, err error) error {
        if err == nil && strings.HasPrefix(filepath.Base(path), iconPrefix) && strings.HasSuffix(path, iconExtension) {
            files = append(files, path)
        }
        return nil
    })

    sort.Strings(files)

    for _, file := range files {
        data, err := os.ReadFile(file)
        if err == nil {
            iconFrames = append(iconFrames, data)
        }
    }

    if len(iconFrames) == 0 {
        fmt.Println("No icon frames found. Exiting.")
        os.Exit(1)
    }
}

func animateIcon() {
    for {
        systray.SetIcon(iconFrames[iconIndex])
        iconIndex = (iconIndex + 1) % len(iconFrames)
        time.Sleep(animationDelay)
    }
}

func startMonitoring() {
    for {
        running := isProcessRunning(processName)
        if running {
            logDebug("Server is running.")
        } else {
            logDebug("Server NOT found. Starting BAT file...")
            runBatFile(filepath.Join(basePath, batFileName))
        }

        tooltip := readLastLogLines(filepath.Join(basePath, logFileName), 10)
        systray.SetTooltip(tooltip)

        time.Sleep(checkInterval)
    }
}

func isProcessRunning(name string) bool {
    cmd := exec.Command("tasklist")
    var out bytes.Buffer
    cmd.Stdout = &out
    if err := cmd.Run(); err != nil {
        logDebug("tasklist command failed: " + err.Error())
        return false
    }
    return strings.Contains(out.String(), name)
}

func runBatFile(path string) {
    cmd := exec.Command("cmd", "/C", path)
    cmd.SysProcAttr = &windows.SysProcAttr{HideWindow: false}

    err := cmd.Start()
    if err != nil {
        logDebug("Failed to start BAT file: " + err.Error())
    } else {
        logDebug("Server file started... ")
    }
}

func showLogWindow() {
    if showLogCmd == nil || showLogCmd.Process == nil {
        showLogCmd = exec.Command("cmd", "/C", "start", "cmd", "/K", "type "+filepath.Join(basePath, logFileName))
        _ = showLogCmd.Start()
    }
}

func cleanupAndExit() {
    if showLogCmd != nil && showLogCmd.Process != nil {
        _ = showLogCmd.Process.Kill()
    }
    systray.Quit()
    os.Exit(0)
}

func readLastLogLines(path string, lines int) string {
    file, err := os.Open(path)
    if err != nil {
        return "Log not found"
    }
    defer file.Close()

    var logLines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        logLines = append(logLines, scanner.Text())
    }

    if len(logLines) > lines {
        logLines = logLines[len(logLines)-lines:]
    }

    return strings.Join(logLines, "\n")
}

func logDebug(msg string) {
    if debugPrint {
        fmt.Println("[DEBUG]", msg)
    }
}
