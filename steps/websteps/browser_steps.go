package websteps

import (
	"fmt"
	"time"

	"github.com/cbush06/kosher/steps/steputils"
)

func iMaximizeTheWindow(s *steputils.StepUtils) func() error {
	return func() error {
		if err := s.Page.Size(s.GetMaxWindowSize()); err != nil {
			return fmt.Errorf("error encountered resizing window: %s", err)
		}
		return nil
	}
}

func iWaitSeconds() func(int) error {
	return func(seconds int) error {
		time.Sleep(time.Duration(seconds) * time.Second)
		return nil
	}
}

func iTakeAScreenshot(s *steputils.StepUtils) func() error {
	return func() error {
		now := time.Now()
		fileName := "screenshot_" + now.Format("02Jan2006-150405.000.png")
		filePath, _ := s.Settings.FileSystem.ResultsDir.RealPath(fileName)
		if err := s.Page.Screenshot(filePath); err != nil {
			return fmt.Errorf("error encountered while taking screenshot: %s", err)
		}
		return nil
	}
}
