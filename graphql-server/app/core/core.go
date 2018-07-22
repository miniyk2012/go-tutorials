// Package core implements the core business logic
// This is where all the uses cases, entities live according to Clean Architecture
// All code in pkg core should not depend on other packages of this application.
// All other packages of this application can and are designed to depend on pkg core.

package core

// GetButtonClickCounts maps a button name to its click count.
func GetButtonClickCounts() map[string]int32 {
	return map[string]int32 {
		"button 1": 1,
		"button 2": 2,
		"button 3": 3,
	}
}

type ServiceGetButtonClickCounts func() map[string]int32
