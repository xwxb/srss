package platform

import (
	"github.com/getlantern/systray"
	// "github.com/xwxb/srss/server"
)

func Start() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(Icon)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")

	menuOpen := systray.AddMenuItem("Open", "")
	systray.AddSeparator()
	menuQuit := systray.AddMenuItem("Quit", "")
	// Sets the icon of a menu item. Only available on Mac and Windows.
	menuQuit.SetIcon(Icon)

	go func() {
		for {
			select {
			case <-menuOpen.ClickedCh:
				// Open(s.GetAddr())
			case <-menuQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	// clean up here
}
