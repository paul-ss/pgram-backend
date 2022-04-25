package main

import (
	"fmt"
	"github.com/paul-ss/pgram-backend/internal/app/server"
	postgres "github.com/paul-ss/pgram-backend/internal/pkg/database"
	"github.com/paul-ss/pgram-backend/internal/pkg/logger"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func initGlobals(initFunctions ...func() func()) (td func(), err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint("recovered: ", e))
		}
	}()

	var teardown []func()
	td = func() {
		for _, f := range teardown {
			f()
		}
	}

	for _, f := range initFunctions {
		teardown = append(teardown, f())
	}

	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("recovered: ", err)
		}
	}()

	cancel, err := initGlobals(logger.Init, postgres.Init)
	defer cancel()

	if err != nil {
		log.Error(err.Error())
		return
	}

	srv := server.NewServer()
	srv.Run()
}
