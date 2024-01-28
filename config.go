package main

import (
    "runtime"
    "github.com/charmbracelet/lipgloss"
)

type OperatingSystemConfig struct {
    taskTitleCutoff int
    borderStyle     lipgloss.Border
    columnHeight    int
    headerPadding   int
}

func initializeConfig() *OperatingSystemConfig {
    if runtime.GOOS == "linux" {
        return NewArchConfig()
    }
    return NewWindowsConfig()

}

func NewArchConfig() *OperatingSystemConfig {
    config := OperatingSystemConfig {
        taskTitleCutoff: 30,
        borderStyle: lipgloss.DoubleBorder(),
        columnHeight: 15,
        headerPadding: 0,
    }
    return &config
}

func NewWindowsConfig() *OperatingSystemConfig {
    config := OperatingSystemConfig {
        taskTitleCutoff: 696969,
        borderStyle: lipgloss.NormalBorder(),
        columnHeight: 15,
        headerPadding: 0,
    }
    return &config
}

