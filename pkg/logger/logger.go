package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var (
	file       *os.File
	currentDay string
)

func InitLogger(logDir string) error {
	// Inisialisasi log pertama kali
	return createLogFile(logDir)
}

func createLogFile(logDir string) error {
	// Tentukan nama file log berdasarkan tanggal saat ini
	today := time.Now().Format("2006-01-02")
	logFilePath := filepath.Join(logDir, fmt.Sprintf("APP-%s.log", today))

	// Simpan tanggal saat ini
	currentDay = today

	// Buat folder logs jika belum ada
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// Buka atau buat file log
	var err error
	file, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Set output log ke file dan terminal (stdout)
	multiWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multiWriter)

	log.Printf("Logger initialized, logging to %s\n", logFilePath)
	return nil
}

// Check if date has changed and rotate log file
func RotateLogFileIfNeeded(logDir string) error {
	today := time.Now().Format("2006-01-02")
	if today != currentDay {
		// Tanggal sudah berubah, buat file log baru
		if err := createLogFile(logDir); err != nil {
			return err
		}
	}
	return nil
}

// CloseLogger menutup file log
func CloseLogger() {
	if file != nil {
		log.Println("Application shutting down")
		file.Close()
	}
}

// SetupCloseHandler menangani sinyal penghentian aplikasi
func SetupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Received shutdown signal")
		CloseLogger()
		os.Exit(0)
	}()
}
