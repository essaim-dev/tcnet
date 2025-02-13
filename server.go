package tcnet

import (
	"errors"
	"fmt"
	"net"
	"time"
)

type Server struct {
	nodeID NodeID
	mode   NodeType

	startAt time.Time

	broadcastListener *net.UDPConn
	timeListener      *net.UDPConn
	nodeListener      *net.UDPConn
}

func NewServer() (*Server, error) {
	broadcastListener, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 60000,
	})
	if err != nil {
		return nil, fmt.Errorf("could not open broadcast listener: %w", err)
	}

	timeListener, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 60001,
	})
	if err != nil {
		return nil, fmt.Errorf("could not open time listener: %w", err)
	}

	nodeListener, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 65023,
	})
	if err != nil {
		return nil, fmt.Errorf("could not open time listener: %w", err)
	}

	return &Server{
		startAt:           time.Now(),
		broadcastListener: broadcastListener,
		timeListener:      timeListener,
		nodeListener:      nodeListener,
	}, nil
}

func (s *Server) Close() (err error) {
	err = errors.Join(s.broadcastListener.Close())
	err = errors.Join(s.timeListener.Close())
	err = errors.Join(s.nodeListener.Close())

	return
}
