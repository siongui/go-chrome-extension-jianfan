package main

import (
	"github.com/fabioberger/chrome"
)

func main() {
	c := chrome.NewChrome()
	c.BrowserAction.OnClicked(func(tab chrome.Tab) {
		println("aaa")
	})
}
