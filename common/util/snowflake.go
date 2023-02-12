package util

import (
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func GenSnowFlake(machineId uint16) (uint64, error) {
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		return machineId, nil
	}
	sf = sonyflake.NewSonyflake(st)

	id, err := sf.NextID()
	if err != nil {
		return 0, errors.Wrap(err, " + sonyflake not created")
	}
	return id, nil
}
