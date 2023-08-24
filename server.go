package serverfix

import (
	"github.com/gardusig/serverfix/internal"
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type Server struct {
	application         quickfix.Application
	logFactory          quickfix.LogFactory
	messageStoreFactory quickfix.MessageStoreFactory
	settings            *quickfix.Settings

	acceptor *quickfix.Acceptor
}

func NewServer(filepath string, app quickfix.Application) (*Server, error) {
	settings, err := internal.GetSettingsFromFile(filepath)
	if err != nil {
		return nil, err
	}
	server := Server{
		application:         app,
		logFactory:          quickfix.NewScreenLogFactory(),
		messageStoreFactory: quickfix.NewMemoryStoreFactory(),
		settings:            settings,
	}
	return &server, nil
}

func (s Server) Start() error {
	logrus.Debug("Starting FIX server...")
	acceptor, err := quickfix.NewAcceptor(
		s.application,
		s.messageStoreFactory,
		s.settings,
		s.logFactory,
	)
	if err != nil {
		logrus.Debug("Failed to create fix initiator, reason: ", err.Error())
		return err
	}
	s.acceptor = acceptor
	return s.acceptor.Start()
}

func (s Server) Stop() {
	s.acceptor.Stop()
}
