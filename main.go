package main

import (
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	myWindow := a.NewWindow("Password generator")

	myWindow.Resize(fyne.NewSize(400, 400))

	title := canvas.NewText("Password generator", theme.ForegroundColor())

	// input box
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter password length")

	// label to show password
	text := canvas.NewText("", theme.ForegroundColor())

	// button to generate password
	generateButton := widget.NewButton("Generate", func() {
		// Convert input to integer
		passwordLength, err := strconv.Atoi(input.Text)
		if err != nil {
			text.Text = "Please enter a valid number"
			text.Refresh()
			return
		}

		// Generate password of the specified length
		text.Text = PasswordGenerator(passwordLength)
		text.Refresh()
	})

	// show content
	myWindow.SetContent(
		container.NewVBox(
			title,
			input,
			text,
			generateButton,
		),
	)

	myWindow.ShowAndRun()
}

func PasswordGenerator(passwordLength int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz" // lowercase
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // uppercase
	Numbers := "0123456789"                   // numbers
	specialChar := "!@#$%^&*()_-+={}[/?]"     // specialchar

	// variable for storing password
	password := ""

	// Initialize the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for n := 0; n < passwordLength; n++ {
		// NOW RANDOM CHARACTER
		randNum := rng.Intn(4)

		switch randNum {
		case 0:
			randCharNum := rng.Intn(len(lowerCase))
			password += string(lowerCase[randCharNum])
		case 1:
			randCharNum := rng.Intn(len(upperCase))
			password += string(upperCase[randCharNum])
		case 2:
			randCharNum := rng.Intn(len(Numbers))
			password += string(Numbers[randCharNum])
		case 3:
			randCharNum := rng.Intn(len(specialChar))
			password += string(specialChar[randCharNum])
		}
	}

	return password
}
