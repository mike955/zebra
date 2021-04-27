package data

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func InitSf(machineId uint16) {
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		return machineId, nil
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

type FlakeData struct {
	logger *logrus.Entry
}

func NewFlakeData(logger *logrus.Entry) *FlakeData {
	return &FlakeData{
		logger: logger,
	}
}

func (s *FlakeData) SetLogger(logger *logrus.Entry) {
	s.logger = logger
}

func (s *FlakeData) New(ctx context.Context) (id uint64, err error) {
	id, err = sf.NextID()
	return
}
