package clapp

import (
	"reflect"
	"strings"
	"testing"
)

func TestSimpleApp(t *testing.T) {
	var got string

	var upper bool
	var message string

	testApp := &Command{
		Name:  "echo",
		Usage: "a simple test app that echo input args in different fashions",

		Flags: Flags{
			{
				Name:  "upper",
				Usage: "Print in uppercase",
				Var:   &upper,
			},
		},

		Args: Args{
			{
				Name:  "message",
				Usage: "Message to echo",
				Var:   &message,
			},
		},

		Execute: func() error {
			if upper {
				got = strings.ToUpper(message)
			} else {
				got = message
			}

			return nil
		},
	}

	testCases := []struct {
		in   []string
		want string
	}{
		{in: []string{"Hello"}, want: "Hello"},
		{in: []string{"--upper", "Hello"}, want: "HELLO"},
	}

	for _, tc := range testCases {
		if err := testApp.Run(tc.in); err != nil {
			t.Errorf("Test programme failed to run: %s", err)
		}
		if got != tc.want {
			t.Errorf("Test failed\nWant: %s\nGot : %s\n", tc.want, got)
		}
	}
}

func TestCmdlineFlags(t *testing.T) {
	var testFlagBool bool
	var testFlagString string
	var testFlagStrings []string

	testApp := &Command{
		Name:  "testApp",
		Usage: "A test for my minimalist cli app building lib",
		Flags: Flags{
			{
				Name:  "bool",
				Usage: "Test Boolean flag",
				Var:   &testFlagBool,
			},
			{
				Name:  "string",
				Usage: "Test String flag",
				Var:   &testFlagString,
			},
			{
				Name:  "strings",
				Usage: "Test Strings flag",
				Var:   &testFlagStrings,
			},
		},
	}

	t.Run("Simple flag parsing", func(t *testing.T) {
		testFlagBool, testFlagString = false, ""
		testApp.cmdline = []string{"--bool", "--string=testing is fun"}
		if err := testApp.parseFlags(); err != nil {
			t.Fatalf("Parsing failed: %v", err)
		}

		if testFlagString != "testing is fun" {
			t.Errorf("String flag is not recognized (string is %s)", testFlagString)
		}

		if !testFlagBool {
			t.Errorf("Boolean flag is not recognized")
		}
	})

	t.Run("unknown flag parsing", func(t *testing.T) {
		testFlagBool, testFlagString = false, ""
		testApp.cmdline = []string{"--unknown"}
		if err := testApp.parseFlags(); err == nil {
			t.Errorf("Parsing succeed despite malformed command line")
		}

		if testFlagString != "" {
			t.Errorf("String flag was modified but should not (string is %s)", testFlagString)
		}

		if testFlagBool {
			t.Errorf("Boolean flag was modified but should not")
		}
	})

	t.Run("Test end of flags", func(t *testing.T) {
		testFlagBool, testFlagString = false, ""
		testApp.cmdline = []string{"--", "--unknown"}
		if err := testApp.parseFlags(); err != nil {
			t.Errorf("Parsing failed, %v", err)
		}
	})

	t.Run("Cumulative flag parsing", func(t *testing.T) {
		testFlagStrings = []string{}
		testApp.cmdline = []string{"--strings=testing,is,fun"}
		if err := testApp.parseFlags(); err != nil {
			t.Fatalf("Parsing failed: %v", err)
		}

		if !reflect.DeepEqual(testFlagStrings, []string{"testing", "is", "fun"}) {
			t.Errorf("Strings flag is not recognized (string is %#v)", testFlagStrings)
		}
	})
}

func TestCmdlineArg(t *testing.T) {
	var testArgInt64 int64
	var testArgString string

	testApp := &Command{
		Name:  "testApp",
		Usage: "A test for my minimalist cli app building lib",
		Args: Args{
			{
				Name:  "int64",
				Usage: "Test int64 arg",
				Var:   &testArgInt64,
			},
			{
				Name:  "string",
				Usage: "Test string arg",
				Var:   &testArgString,
			},
		},
	}

	t.Run("Simple arg parsing", func(t *testing.T) {
		testArgInt64, testArgString = 0, ""
		testApp.cmdline = []string{"42", "I'm a string arg"}

		if err := testApp.parseArgs(); err != nil {
			t.Fatalf("Parsing of %#v failed: %v", testApp.cmdline, err)
		}

		if testArgString != "I'm a string arg" {
			t.Errorf("String arg is not recognized (string is %s)", testArgString)
		}

		if testArgInt64 != 42 {
			t.Errorf("Int64 arg is not recognized (int is %d)", testArgInt64)
		}
	})

	t.Run("Wrong arg number/order/type", func(t *testing.T) {
		testArgInt64, testArgString = 0, ""

		testApp.cmdline = []string{"7"}
		if err := testApp.parseArgs(); err == nil {
			t.Errorf("Parsing succeed with a malformed command line (%#v)", testApp.cmdline)
		}
		testApp.cmdline = []string{"3.14", "I'm a string arg"}
		if err := testApp.parseArgs(); err == nil {
			t.Errorf("Parsing succeed with a malformed command line (%#v)", testApp.cmdline)
		}
		testApp.cmdline = []string{"I'm a string", "7"}
		if err := testApp.parseArgs(); err == nil {
			t.Errorf("Parsing succeed with a malformed command line")
		}
	})

	t.Run("Test end of flags", func(t *testing.T) {
		testArgInt64, testArgString = 0, ""
		testApp.cmdline = []string{"--", "42", "--unknown"}

		if err := testApp.parseFlags(); err != nil {
			t.Errorf("Parsing failed, %v", err)
		}
		if err := testApp.parseArgs(); err != nil {
			t.Errorf("Parsing failed, %v", err)
		}

		if testArgString != "--unknown" {
			t.Errorf("String arg is not recognized (string is %s)", testArgString)
		}

		if testArgInt64 != 42 {
			t.Errorf("Int64 arg is not recognized (int is %d)", testArgInt64)
		}
	})
}

func TestCumulativeArg(t *testing.T) {
	var testArgInt64 int64
	var testArgStrings []string

	testApp := &Command{
		Name:  "testApp",
		Usage: "A test for my minimalist cli app building lib",
		Args: Args{
			{
				Name:  "int64",
				Usage: "Test int64 arg",
				Var:   &testArgInt64,
			},
			{
				Name:  "string",
				Usage: "Test strings arg",
				Var:   &testArgStrings,
			},
		},
	}

	testApp.cmdline = []string{"42", "I'm a string arg", "I'm another string arg"}
	if err := testApp.parseArgs(); err != nil {
		t.Fatalf("Parsing of %#v failed: %v", testApp.cmdline, err)
	}

	if !reflect.DeepEqual(testArgStrings, testApp.cmdline[1:]) {
		t.Errorf("Strings arg is not recognized (string is %s, instead of %v)", testArgStrings, testApp.cmdline[1:])
	}

	if testArgInt64 != 42 {
		t.Errorf("Int64 arg is not recognized (int is %d)", testArgInt64)
	}
}

func TestCmdlineNoArgs(t *testing.T) {
	var testArgInt64 int64

	testApp := &Command{
		Name:  "testApp",
		Usage: "A test for my minimalist cli app building lib",
		Args: Args{
			{
				Name:     "int64",
				Usage:    "Test int64 arg",
				Var:      &testArgInt64,
				Optional: true,
			},
		},
	}

	t.Run("Simple arg parsing", func(t *testing.T) {
		testArgInt64 = 0
		testApp.cmdline = []string{"42"}

		if err := testApp.parseArgs(); err != nil {
			t.Fatalf("Parsing of %#v failed: %v", testApp.cmdline, err)
		}

		if testArgInt64 != 42 {
			t.Errorf("Int64 arg is not recognized (int is %d)", testArgInt64)
		}
	})

	t.Run("Wrong arg number/order/type", func(t *testing.T) {
		testArgInt64 = 7

		testApp.cmdline = []string{}
		if err := testApp.parseArgs(); err != nil {
			t.Fatalf("Parsing of %#v without args failed: %v", testApp.cmdline, err)
		}

		if testArgInt64 != 7 {
			t.Errorf("Parsing empty args list failed: arg does not fallback to its default value (int is %d)", testArgInt64)
		}
	})
}

func TestSubCommandsParsing(t *testing.T) {
	var testAppFlag bool
	var testArgString string
	var testCmdArg string
	var testResult string

	testApp := &Command{
		Name:  "testApp",
		Usage: "A test for my minimalist cli app building lib",
		Flags: Flags{
			{
				Name:  "bool",
				Usage: "Test bool flag",
				Var:   &testAppFlag,
			},
		},
		Args: Args{
			{
				Name:  "string",
				Usage: "Test String arg",
				Var:   &testArgString,
			},
		},
		Execute: func() error {
			testResult += testArgString
			return nil
		},
		SubCommands: Commands{
			{
				Name:  "test",
				Usage: "Test a sub-command",
				Args: Args{
					{
						Name:  "test_arg",
						Usage: "Test string arg of sub-command test",
						Var:   &testCmdArg,
					},
				},
				Execute: func() error {
					testResult += testCmdArg
					return nil
				},
			},
		},
	}

	t.Run("Without Subcommand", func(t *testing.T) {
		testResult, testAppFlag, testArgString, testCmdArg = "", false, "", ""
		cmdline := []string{"--bool", "I'm the test string of test app"}

		if err := testApp.Run(cmdline); err != nil {
			t.Fatalf("Run of %#v failed: %v", cmdline, err)
		}

		if testAppFlag != true {
			t.Errorf("Boolean flag is incorrectly set")
		}

		if testResult != "I'm the test string of test app" {
			t.Errorf("Test sub-command doe snot run as expected (get %s)", testResult)
		}
	})

	t.Run("With Subcommand", func(t *testing.T) {
		testResult, testAppFlag, testArgString, testCmdArg = "", false, "", ""
		cmdline := []string{"--bool", "test", "I'm the test string of test sub-command"}

		if err := testApp.Run(cmdline); err != nil {
			t.Fatalf("Run of %#v failed: %v", cmdline, err)
		}

		if testAppFlag != true {
			t.Errorf("Boolean flag is incorrectly set")
		}

		if testResult != "I'm the test string of test sub-command" {
			t.Errorf("Test sub-command doe snot run as expected (get %s)", testResult)
		}
	})

	t.Run("Wrong arg number/order/type", func(t *testing.T) {
		testResult, testAppFlag, testArgString, testCmdArg = "", false, "", ""
		cmdline := []string{"--bool"}
		if err := testApp.Run(cmdline); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", cmdline)
		}

		cmdline = []string{"test"}
		if err := testApp.Run(cmdline); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", cmdline)
		}

		cmdline = []string{"test", "--bool", "I'm the test string of test sub-command"}
		if err := testApp.Run(cmdline); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", cmdline)
		}
	})
}
