package main

import (
	"github.com/ahmetberke/christmas-raffle/configs"
	"github.com/ahmetberke/christmas-raffle/pkg/mailer"
	"github.com/ahmetberke/christmas-raffle/pkg/raffle"
)

func init() {
	configs.ReadConfigFile("./configs/", "config")
}

func main() {
	// Creating new raffle
	raf := raffle.NewRaffle()
	for _, np := range configs.Manager.Participants {
		// Adding participants from config manager
		raf.AddParticipant(np.Name, np.Email)
	}
	// Christmas Raffle is drawing
	raf.Draw()
	// Print Draws results
	raf.Results.Print()
	// Sending email result to participants by privately
	for _, rel := range raf.Results.Relations {
		mail := mailer.NewMail([]string{rel.Who.Mail}, "HAPPY NEW YEARS 🎄✨🎅")
		mail.Send(configs.Manager.MailSettings.BodyFilePath, rel)
	}
}
