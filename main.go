package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/dustin/go-humanize"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"strconv"
	"strings"
)

// Removed switch -- you've already got one built in to tview. Instead of nul, just give it an action
func main() {
	app := tview.NewApplication()
	list := tview.NewList().AddItem("Select an item from below...", "", 'm', func() {
		app.Stop()
		fmt.Println("Restarting Menu,")
		displayBox("press any key to Continue ...")
	})
	list.AddItem("Get available memory.", "", 'a', func() {
		app.Stop()
		// This is the main thing I wanted to demonstrate, Script allows you to simply use the built in commands easily.
		mem, err := script.Exec("sysctl -n hw.memsize").String()
		errorHandle(err, "get available memory")
		mem = strings.TrimSpace(mem)
		// convert string to uint64
		m, err := strconv.ParseUint(mem, 10, 64)
		errorHandle(err, "convert string to int")
		availMem := fmt.Sprintf("Available Memory: %s", humanize.Bytes(m))
		displayBox(availMem)
	})
	list.AddItem("Get available disk space.", "", 'b', func() {
		app.Stop()
		dsk, err := script.Exec("df -h /").Column(4).String()
		errorHandle(err, "get disk space")
		displayBox(dsk)
	})
	list.AddItem("Get the current logged in user.", "", 'c', func() {
		app.Stop()
		dsk, err := script.Exec("who").String()
		errorHandle(err, "who")
		displayBox(dsk)
	})
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
		os.Exit(1)
	})
	// WHy not add a border...
	list.SetBorder(true)
	// This will loop forever, everytime the app is stopped by one of the functions, this will restart it, until they
	// select quit, then the entire program exits.
	for {
		err := app.SetRoot(list, true).SetFocus(list).Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}

func errorHandle(err error, str string) {
	if err != nil {
		fmt.Println("error:", str)
		fmt.Println(err)
	}
}

// show the results from script in a window. wait for a keypress and return.
func displayBox(txt string) {
	app := tview.NewApplication()
	txt = fmt.Sprintf("%s\n\nPress any key to continue...", txt)
	textView := tview.NewTextView().
		SetText(txt).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	textView.SetBorder(true)
	textView.SetTextAlign(0)
	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		app.Stop()
		return event
	})
	textView.GetInputCapture()
	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
