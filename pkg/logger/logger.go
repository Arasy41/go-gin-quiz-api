package logger

import (
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	file *os.File
)

func InitLogger(logFilePath string) error {
	// Buat folder logs jika belum ada
	logDir := filepath.Dir(logFilePath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// Buka atau buat file log
	var err error
	file, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Set output log ke file
	multiWriter := io.MultiWriter(file, os.Stdout)
	os.Stdout = file
	os.Stderr = file

	// Set output log ke multiWriter
	log.SetOutput(multiWriter)
	log.Println("Logger initialized, logging to file and terminal")
	return nil
}

// CloseLogger menutup file log dan mencatat pesan saat aplikasi berhenti
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
