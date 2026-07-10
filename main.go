package main

import (
    "flag"
    "fmt"
    "os"
    "os/signal"
    "syscall"
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

    workMinutes := flag.Int("work", 25, "durasi sesi fokus (menit)")
    breakMinutes := flag.Int("break", 5, "durasi sesi istirahat (menit)")
    totalSessions := flag.Int("sessions", 4, "jumlah sesi fokus")
    flag.Parse()

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-sigChan
        fmt.Println("\n\n👋 Sesi dihentikan. Sampai jumpa lagi!")
        os.Exit(0)
    }()

    fmt.Println("🐧 Penguin Pomodoro CLI")
    fmt.Println("----------------------")

    for session := 1; session <= *totalSessions; session++ {

        fmt.Printf("\n🔥 Sesi %d/%d\n", session, *totalSessions)

        countdown(*workMinutes, "Fokus")

        if session < *totalSessions {
            countdown(*breakMinutes, "istirahat")
        }
    }

    fmt.Println("\n🎉 Satu sesi Pomodoro selesai!")
}
