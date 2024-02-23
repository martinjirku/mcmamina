package services_test

import (
	"fmt"
	"testing"

	"jirku.sk/mcmamina/services"
)

func TestMailSend(t *testing.T) {
	service := services.NewMailService("email@emai.sk", "pwd")
	err := service.Send("email@email.sk", "ahoj svet")
	if err != nil {
		t.Error(fmt.Errorf("sending mail failed: %w", err))
	}
}
