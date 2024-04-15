package utils

import (
	"errors"

	log "github.com/gilbarco-ai/event-bus/common/log"
)

func ValidateList(input interface{}, args ...interface{}) (error, bool) {
	found := false
	if len(args) == 0 {
		log.Logger.Error("Invalid number of arguments in ValidateList function")
		return errors.New("Invalid number of arguments in ValidateList function"), found
	}
	if value, ok := input.(string); ok {
		if len(value) == 0 {
			log.Logger.Error("Invalid input in ValidateList function")
			return errors.New("Invalid input in ValidateList function"), found
		}
		for _, arg := range args {
			if arg == input {
				found = true
				break
			}
		}
	} else if value, ok := input.(bool); ok {
		if value != true && value != false {
			log.Logger.Error("Invalid input in ValidateList function")
			return errors.New("Invalid input in ValidateList function"), found
		} else {
			found = true
		}
	}

	return nil, found
}

//==============================================   Set   =======================================================
type StringSet map[string]struct{}

func AddToSet(set StringSet, value string, typeSet string) {
	if _, exists := set[value]; exists {
		msg := "Value " + value + " already exists in " + typeSet
		log.Logger.Error(msg)
		panic(msg)
	}

	set[value] = struct{}{}
}

func IsInSet(set StringSet, value string) bool {
	_, exists := set[value]
	return exists
}
