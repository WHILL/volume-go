package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/itchyny/volume-go"
)

func run() error {
	if len(os.Args) <= 1 {
		return errors.New("no arg")
	}
	switch os.Args[1] {
	case "get":
		if len(os.Args) == 2 {
			return getVolume()
		}
		return errors.New("invalid argument for volume get")
	case "set":
		if len(os.Args) == 3 {
			return setVolume(os.Args[2])
		}
		return errors.New("invalid argument for volume get")
	case "mute":
		if len(os.Args) == 2 {
			return volume.Mute()
		}
		return errors.New("invalid argument for volume mute")
	case "unmute":
		if len(os.Args) == 2 {
			return volume.Unmute()
		}
		return errors.New("invalid argument for volume unmute")
	}
	return fmt.Errorf("invalid argument for volume: %+v", os.Args[1:])
}

func getVolume() error {
	vol, err := volume.GetVolume()
	if err != nil {
		return err
	}
	fmt.Println(vol)
	return nil
}

func setVolume(volStr string) error {
	vol, err := strconv.Atoi(volStr)
	if err != nil {
		return err
	}
	return volume.SetVolume(vol)
}
