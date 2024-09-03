package main

import (
	"log/slog"

	"github.com/google/uuid"
)

func main() {
	slog.Info("UUID sortable")

	u1, _ := uuid.NewV7()
	//u1 := uuid.New()
	u2, _ := uuid.NewV7()
	//u2 := uuid.New()
	//u3, _ := uuid.NewV7()
	u3 := uuid.New()

	us1 := u1.String()
	us2 := u2.String()
	us3 := u3.String()

	slog.Info("strings", "u1", us1, "u2", us2, "u3", us3)

	slog.Info("compare", "u1 > u2", us1 > us2)
}
