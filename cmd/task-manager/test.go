package main

import (
    "fmt"
    "strings"

    "github.com/charmbracelet/lipgloss"
)

var (
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FFFFFF")).
        PaddingBottom(1)

    headerStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#888888"))

    doneStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#04B575")).
        Bold(true)

    pendingStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#E5A000")).
        Bold(true)

    idStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#777777"))

    titleCellStyle = lipgloss.NewStyle().
        Width(30)
)

func RenderTasks(tasks []Task) {
    fmt.Println(titleStyle.Render("TASK MANAGER"))
    fmt.Println()

    header := lipgloss.JoinHorizontal(
        lipgloss.Top,
        headerStyle.Render("ID"),
        headerStyle.Render("TITLE"),
        headerStyle.Render("STATUS"),
    )

    fmt.Println(header)
    fmt.Println(strings.Repeat("-", 60))

    for _, task := range tasks {
        status := pendingStyle.Render("PENDING")

        if task.Status == "completed" {
            status = doneStyle.Render("DONE")
        }

        row := lipgloss.JoinHorizontal(
            lipgloss.Top,
            idStyle.Width(6).Render(
                fmt.Sprintf("%d", task.ID),
            ),
            titleCellStyle.Render(task.Title),
            status,
        )

        fmt.Println(row)
    }

    fmt.Println()
    fmt.Printf("Total tasks: %d\n", len(tasks))
}
