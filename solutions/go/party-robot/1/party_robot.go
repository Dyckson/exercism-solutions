package partyrobot

import (
    "fmt"
)

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	const template = `Welcome to my party, %s!
You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.
You will be sitting next to %s.`
	return fmt.Sprintf(template, name, table, direction, distance, neighbor)
}

