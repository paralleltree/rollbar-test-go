package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/rollbar/rollbar-go"
	"golang.org/x/xerrors"
)

func main() {
	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	rollbar.SetCodeVersion(os.Getenv("CODE_VERSION"))
	rollbar.SetEnvironment("production") // defaults to "development"

	err := doSomething1()
	if err != nil {
		rollbar.Error(err)
	}

	rollbar.Wait()
}

func doSomething1() error {
	if err := doSomething2(); err != nil {
		return fmt.Errorf("doSomething1: %w", err)
	}
	return nil
}
func doSomething2() error {
	if err := somethingWentWrong(); err != nil {
		return fmt.Errorf("doSomething2: %w", err)
	}
	return nil
}

func somethingWentWrong() error {
	err := errors.New("internal error: param1: b")
	return xerrors.Errorf("oops: %w", err)
}
