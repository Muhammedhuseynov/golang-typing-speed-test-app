package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type MyCustomTextGridStyle struct {
	FGColor, BGColor color.Color
}

// TextColor is the color a cell should use for the text.
func (c *MyCustomTextGridStyle) TextColor() color.Color {
	return c.FGColor
}

// BackgroundColor is the color a cell should use for the background.
func (c *MyCustomTextGridStyle) BackgroundColor() color.Color {
	return c.BGColor
}

// make custom TextGridStyle
type myColors struct {
	txtGridStyle MyCustomTextGridStyle
}

var (
	myStyle  widget.TextGridStyle
	myStyle2 widget.TextGridStyle
)

//make object(layout) on center
func makeObjCenter(obj fyne.CanvasObject) *fyne.Container {
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), obj, layout.NewSpacer())
}

//make random number
func randomNum(n int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	newInt := r.Intn(n)

	return newInt
}

//elapsed time
func timeTrack(start time.Time) time.Duration {
	elapsed := time.Since(start)
	return elapsed

}
func main() {
	txts := []string{
		"Concerns greatest margaret him absolute entrance nay. Unpacked endeavor six steepest had husbands her.",
		"present garrets limited cordial in inquiry 2. Supported me sweetness behaviour shameless excellent so arranging",
		"Up greatest am exertion or marianne. Shy occasional terminated insensible",
		"so as doubtful on striking required. Waiting we to compass assured.",
		"Extremity excellent 99 certainty, discourse sincerity no he so resembled. Joy house worse arise total boy but. Elderly",
		"engrossed suffering supposing, he recommend do 15 eagerness. Commanded no of depending extremity",
		"Sing long her way size. Waited end mutual missed myself the little sister one.",
		"If their woman could do wound on. You folly taste hoped their above are and but!",
		"on am we offices expense thought. Its hence ten smile age means. Seven chief sight far point any +",
		"Of so high into easy. Dashwoods eagerness oh extensive as discourse sportsman frankness. 1",
		"resolve garrets as. Impression was estimating surrounded SOlicitude indulgence son shy",
		"What does the Lorem Ipsum text mean?",
		"He hastened am no property exercise of. Dissimilar comparison no terminated devonshire no literature on?!",
		"02Remain lively hardly needed at do by. Two you fat downs fanny three. True mr gone most at",
		"cottage be noisier looking to we promise on. Disposal to kindness!",
	}
	rand.Seed(time.Now().UnixNano())
	randInd := randomNum(15)
	_app := app.New()
	_app.Settings().SetTheme(theme.DarkTheme())

	wind := _app.NewWindow("Typing Test")
	wind.CenterOnScreen()

	wind.SetFixedSize(true)

	wind.Resize(fyne.NewSize(1160, 270))
	label := canvas.NewText("Typing Test", color.RGBA{53, 43, 86, 1})
	label.TextSize = 50
	label.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	lblContainer := makeObjCenter(label)

	typedResultScores := widget.NewLabel("")
	typedResultScores.Hide()
	typedResultScoresContainer := makeObjCenter(typedResultScores)

	lblRandTxt := widget.NewTextGridFromString(txts[randInd])

	colors := myColors{txtGridStyle: MyCustomTextGridStyle{}}

	input := widget.NewEntry()
	input.TextStyle = fyne.TextStyle{Bold: true}
	input.PlaceHolder = "Type here"

	result := ""

	btn := widget.NewButton("RESET", func() {
		input.Text = ""
		typedResultScores.SetText(result)
		typedResultScores.Show()
		input.Enable()
		input.Refresh()

		lblRandTxt.SetText(txts[randomNum(15)])
		lblRandTxt.Refresh()
	})
	btn.Disable()
	var start_time = time.Now()

	//endTime := time.Now()
	startTimer := true
	count_corrects := 0

	input.OnChanged = func(s string) {
		trackCharInd := len(s) - 1
		if trackCharInd >= 0 {
			if startTimer {
				//fmt.Println("Timer Started:...")
				start_time = time.Now()
				//fmt.Println("Start Time: ", start_time)
			}
			startTimer = false

			userEnteredChar := string(s[trackCharInd])
			mainTxtChar := string(lblRandTxt.Text()[trackCharInd])

			colors.txtGridStyle.BGColor = nil
			if userEnteredChar != mainTxtChar {
				if mainTxtChar == " " && userEnteredChar != " " {
					//make red if space not entered
					colors.txtGridStyle.BGColor = color.RGBA{255, 0, 0, 1}
				}
				//fmt.Println("Color Changing")
				//make red if not correct character
				colors.txtGridStyle.FGColor = color.RGBA{255, 0, 0, 1}
				lblRandTxt.SetStyleRange(0, trackCharInd, 0, len(s)-1, &colors.txtGridStyle)

			} else {
				colors.txtGridStyle.FGColor = color.RGBA{27, 182, 73, 1}
				lblRandTxt.SetStyleRange(0, trackCharInd, 0, len(s)-1, &colors.txtGridStyle)
				count_corrects += 1

			}

		}

		if trackCharInd+1 == len(lblRandTxt.Text()) {
			input.Disable()
			btn.Enable()

			track := timeTrack(start_time)

			wpm := math.Round((float64(len(s)) / (track.Seconds() / 60.0)) / 5)

			accuracy := float64((float64(count_corrects) / float64(len(lblRandTxt.Text())))) * 100

			result = fmt.Sprintf("Time: %v | Acc: %.2f %s  | WPM: %v", track.Round(time.Second), accuracy, "%", wpm)
			information := dialog.NewInformation("Your Result!", result, wind)

			//information.SetOnClosed(func() {
			//	fmt.Println("Closed")
			//
			//})
			information.Show()

			startTimer = true
			count_corrects = 0

		}
	}

	forSpace := widget.NewLabel(" ")
	allLayouts := container.NewVBox(container.NewVBox(lblContainer), forSpace, lblRandTxt, input, typedResultScoresContainer, btn)
	wind.SetContent(allLayouts)
	wind.ShowAndRun()
}
