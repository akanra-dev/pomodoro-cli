package main

import (
    "fmt"
    "time"
)

func countdown(minutes int, activity string) {
    fmt.Printf("\n🍅 %s dimulai (%d menit)\n", activity, minutes)

    totalSeconds := minutes * 60

    for totalSeconds > 0 {
        mins := totalSeconds / 60
        secs := totalSeconds % 60

        fmt.Printf("\r⏳ %02d:%02d\r", mins, secs)

        time.Sleep(time.Second)

        totalSeconds--
    }

    fmt.Printf("\n✅ %s selesai!\n", activity)
}

func main() {

    fmt.Println("🐧 Penguin Pomodoro CLI")
    fmt.Println("----------------------")

    for session := 1; session <= 4; session++ {

        fmt.Printf("\n🔥 Sesi %d/4\n", session)

        countdown(25, "Fokus")
        countdown(5, "Istirahat")

        if session < 4 {
            countdown(5, "Istirahat")
        }
    }

    fmt.Println("\n🎉 Satu sesi Pomodoro selesai!")
}
