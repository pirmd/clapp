package clapp

import (
	"bytes"
	"testing"

	"github.com/pirmd/style"
	"github.com/pirmd/verify"
)

func TestManpage(t *testing.T) {
	testApp := buildTestApp()

	out := new(bytes.Buffer)

	//Ensure that manpage date will remain the same whenever the test is run
	manDate = "2019-07-12"
	PrintManpage(out, testApp, style.NewMan())

	if failure := verify.MatchGolden(t.Name(), out.String()); failure != nil {
		t.Errorf("Manpage message is incorrectly formatted.\n %v", failure)
	}
}
