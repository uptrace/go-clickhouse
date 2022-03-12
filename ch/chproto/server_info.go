package chproto

import (
	"time"
)

type ServerInfo struct {
	Name         string
	MinorVersion uint64
	MajorVersion uint64
	Revision     uint64
}

var (
	timeZoneOffset int
)

func (srv *ServerInfo) ReadFrom(rd *Reader) (err error) {
	if srv.Name, err = rd.String(); err != nil {
		return err
	}
	if srv.MajorVersion, err = rd.Uvarint(); err != nil {
		return err
	}
	if srv.MinorVersion, err = rd.Uvarint(); err != nil {
		return err
	}
	if srv.Revision, err = rd.Uvarint(); err != nil {
		return err
	}

	timezone, err := rd.String()
	if err != nil {
		return err
	}

	serverTimeZone, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	_, timeZoneOffset = time.Now().In(serverTimeZone).Zone()

	if _, err = rd.String(); err != nil { // display name
		return err
	}
	if _, err = rd.Uvarint(); err != nil { // server version patch
		return err
	}

	return nil
}

func TimeZoneOffset() int64 {
	return int64(timeZoneOffset)
}
