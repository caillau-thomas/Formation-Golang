package main

import (
	"errors"
	"fmt"
)

func sayByeFunc(name string) (string, error) {
	if name == "" {
		return "", errors.New("\n\n\bErreur, aucun nom spécifié ! ")
	}

	byeMsg := fmt.Sprintf("\n\n%v s'en va! Bonne soirée", name)
	return byeMsg, nil
}
