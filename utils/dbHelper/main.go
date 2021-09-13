package dbhelper

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type commandControllerFake struct{}

func (cc *commandControllerFake) Rollback()            {}
func (cc *commandControllerFake) Commit()              {}
func (cc *commandControllerFake) Done(err error) error { return err }
func FakeCC() CommandController {
	return new(commandControllerFake)
}

type CommandController interface {
	Rollback()
	Commit()
	Done(err error) error
}

type commandController struct {
	tx     *gorm.DB
	called bool
	cancel context.CancelFunc
}

func (cc *commandController) Rollback() {
	if !cc.called {
		fmt.Println("Rollback", cc.tx.Rollback().Error)
		cc.cancel()
		cc.called = true
	}
}
func (cc *commandController) Commit() {
	if !cc.called {
		fmt.Println("Commit", cc.tx.Commit().Error)
		cc.cancel()
		cc.called = true
	}
}
func (cc *commandController) Done(err error) error {
	if err != nil {
		cc.Rollback()
		return err
	}
	cc.Commit()
	return nil
}

const timeout = time.Second * 5

func Transaction(tx *gorm.DB, fc func(tx *gorm.DB) error) (CommandController, error) {
	tx = tx.Begin()

	err := fc(tx)
	if err == nil {
		err = tx.Error
	}

	ctx, cancel := context.WithCancel(context.Background())
	commandController := &commandController{
		tx:     tx,
		called: false,
		cancel: cancel,
	}
	go func() {
		select {
		case <-ctx.Done():
			break
		case <-time.After(timeout):
			commandController.Commit()
		}
	}()

	if err != nil {
		commandController.Rollback()
	}

	return commandController, err
}
