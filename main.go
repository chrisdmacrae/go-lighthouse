package main

import (
	"log/slog"

	// vlc "github.com/adrg/libvlc-go/v3"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	// systray.SetIcon()
	systray.SetTitle("I'm alive!")
	systray.SetTooltip("Look at me, I'm a tooltip!")
}

func onExit() {
	// Cleaning stuff here.
}

func player() {
	slog.Info("Starting GoLighthouse")

	// // Setup VLC
	// if err := vlc.Init("--no-video", "--quiet"); err != nil {
	// 	log.Fatal(err)
	// }
	// defer vlc.Release()

	// // Create a new player
	// player, err := vlc.NewPlayer()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	player.Stop()
	// 	player.Release()
	// }()

	// // Add a media file from path or from URL.
	// // Set player media from path:
	// // media, err := player.LoadMediaFromPath("localpath/test.mp4")
	// // Set player media from URL:
	// media, err := player.LoadMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer media.Release()

	// // Retrieve player event manager.
	// manager, err := player.EventManager()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Register the media end reached event with the event manager.
	// quit := make(chan struct{})
	// eventCallback := func(event vlc.Event, userData interface{}) {
	// 	close(quit)
	// }

	// eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer manager.Detach(eventID)

	// // Start playing the media.
	// err = player.Play()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// <-quit
}
