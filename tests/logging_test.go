package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	log := logrus.New()
	log.Println("Hello World")
}

func TestOsExit(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i == 5 {
			os.Exit(1)
		}
		fmt.Println(i)
	}
}

func TestLevel(t *testing.T) {
	logger := logrus.New()
	// jika ingin menampilkan informasi log level dari level Trace
	logger.SetLevel(logrus.TraceLevel)	

	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("../application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.Info("Informasi Log Yang Tersimpan di application.log")
	logger.Debug("Informasi Log Yang Tersimpan di application.log")
	logger.Warn("Informasi Log Yang Tersimpan di application.log")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetFormatter(&logrus.TextFormatter{})

	logger.Info("Membuat data logging")
}

func TestAddAnotherField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Fauzan").Info("Yang login saat ini")

	logger.WithField("username", "Susi").WithField("role", "admin").Info("Yang login saat ini")

	// jika menggunakan field yang banyak

	logger.WithFields(logrus.Fields{
		"username": "Susi",
		"role": "Admin",
	}).Infof("Yang login saat ini")
}

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "Fauzan")
	entry.Info("Yang login")
}

type SampleHook struct {

}

func(s SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func(s SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook ", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("Log Info") // tidak akan keluar karena levelnya info
	// agar keluar, harus levelnya error atau warning sesuai dengan struct SampleHook
	logger.Error("Log Error")
	logger.Warn("Log Warning")
}

func TestSingleton(t *testing.T) {
	// tanpa menggunakan object langsung (membuat objek baru)
	logrus1 := logrus.New()
	logrus1.SetFormatter(&logrus.JSONFormatter{})
	logrus1.Error("Error dari logrus1")

	logrus2 := logrus.New()
	logrus2.SetFormatter(&logrus.TextFormatter{})
	logrus2.Error("Error dari logrus2")

	// menggunakan objek langsung/singleton
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Error("Error dari global")
}