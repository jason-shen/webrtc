package sctp

import (
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
)

type HMACAlgorithm uint16

const (
	HMACResv1  HMACAlgorithm = 0
	HMACSHA128               = 1
	HMACResv2  HMACAlgorithm = 2
	HMACSHA256 HMACAlgorithm = 3
)

func (c HMACAlgorithm) String() string {
	switch c {
	case HMACResv1:
		return "HMAC Reserved (0x00)"
	case HMACSHA128:
		return "HMAC SHA-128"
	case HMACResv2:
		return "HMAC Reserved (0x02)"
	case HMACSHA256:
		return "HMAC SHA-256"
	default:
		return fmt.Sprintf("Unknown HMAC Algorithm type: %d", c)
	}
}

type ParamRequestedHMACAlgorithm struct {
	Raw                 []byte
	AvailableAlgorithms []HMACAlgorithm
}

func (r *ParamRequestedHMACAlgorithm) Marshal() ([]byte, error) {
	return nil, errors.New("Not implemented")
}

func (r *ParamRequestedHMACAlgorithm) Unmarshal(raw []byte) (Param, error) {
	r.Raw = raw

	i := 0
	for i < len(raw) {
		a := HMACAlgorithm(binary.BigEndian.Uint16(raw[i:]))
		switch a {
		case HMACSHA128:
			fallthrough
		case HMACSHA256:
			r.AvailableAlgorithms = append(r.AvailableAlgorithms, a)
		default:
			return nil, errors.Errorf("Invalid algorithm type '%v'", a)
		}

		i += 2
	}

	return r, nil
}