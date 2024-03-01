package tests

import (
	"os"
)

type InStructParent struct {
	FldNested InStructNested `json:"nested" yaml:"nested" toml:"nested"`
}
type InStructNested struct {
	FldBool bool `env:"BOOL" flag:"bool" default:"false" json:"bool" yaml:"bool" toml:"bool"`
}
type InStruct struct {
	FldString  string         `env:"STR" flag:"str" default:"default-string" json:"str" yaml:"str" toml:"str"`
	FldInt     int            `env:"INT" flag:"int" default:"-42" json:"int" yaml:"int" toml:"int"`
	FldInt8    int8           `env:"INT_8" flag:"int-8" default:"-8" json:"int_8" yaml:"int_8" toml:"int_8"`
	FldInt16   int16          `env:"INT_16" flag:"int-16" default:"-16" json:"int_16" yaml:"int_16" toml:"int_16"`
	FldInt32   int32          `env:"INT_32" flag:"int-32" default:"-32" json:"int_32" yaml:"int_32" toml:"int_32"`
	FldInt64   int64          `env:"INT_64" flag:"int-64" default:"-64" json:"int_64" yaml:"int_64" toml:"int_64"`
	FldUint    uint           `env:"UINT" flag:"uint" default:"42" json:"uint" yaml:"uint" toml:"uint"`
	FldUint8   uint8          `env:"UINT_8" flag:"uint-8" default:"8" json:"uint_8" yaml:"uint_8" toml:"uint_8"`
	FldUint16  uint16         `env:"UINT_16" flag:"uint-16" default:"32" json:"uint_16" yaml:"uint_16" toml:"uint_16"`
	FldUint32  uint16         `env:"UINT_32" flag:"uint-32" default:"32" json:"uint_32" yaml:"uint_32" toml:"uint_32"`
	FldUint64  uint64         `env:"UINT_64" flag:"uint-64" default:"64" json:"uint_64" yaml:"uint_64" toml:"uint_64"`
	FldFloat32 float32        `env:"FLOAT_32" flag:"float-32" default:"32.32" json:"float_32" yaml:"float_32" toml:"float_32"`
	FldFloat64 float64        `env:"FLOAT_64" flag:"float-64" default:"64.64" json:"float_64" yaml:"float_64" toml:"float_64"`
	FldParent  InStructParent `json:"parent" yaml:"parent" toml:"parent"`
}

func stubDefault() InStruct {
	return InStruct{
		FldString:  "default-string",
		FldInt:     -42,
		FldInt8:    -8,
		FldInt16:   -16,
		FldInt32:   -32,
		FldInt64:   -64,
		FldUint:    42,
		FldUint8:   8,
		FldUint16:  32,
		FldUint32:  32,
		FldUint64:  64,
		FldFloat32: 32.32,
		FldFloat64: 64.64,
		FldParent: InStructParent{
			FldNested: InStructNested{
				FldBool: false,
			},
		},
	}
}

func stubEnv() InStruct {
	_ = os.Setenv("STR", "env-string")
	_ = os.Setenv("INT", "-142")
	_ = os.Setenv("INT_8", "-18")
	_ = os.Setenv("INT_16", "-116")
	_ = os.Setenv("INT_32", "-132")
	_ = os.Setenv("INT_64", "-164")
	_ = os.Setenv("UINT", "142")
	_ = os.Setenv("UINT_8", "18")
	_ = os.Setenv("UINT_16", "132")
	_ = os.Setenv("UINT_32", "132")
	_ = os.Setenv("UINT_64", "164")
	_ = os.Setenv("FLOAT_32", "132.32")
	_ = os.Setenv("FLOAT_64", "164.64")
	_ = os.Setenv("BOOL", "true")

	return InStruct{
		FldString:  "env-string",
		FldInt:     -142,
		FldInt8:    -18,
		FldInt16:   -116,
		FldInt32:   -132,
		FldInt64:   -164,
		FldUint:    142,
		FldUint8:   18,
		FldUint16:  132,
		FldUint32:  132,
		FldUint64:  164,
		FldFloat32: 132.32,
		FldFloat64: 164.64,
		FldParent: InStructParent{
			FldNested: InStructNested{
				FldBool: true,
			},
		},
	}
}

func stubFlag() InStruct {
	os.Args = append(
		os.Args,
		"--str", "flag-string",
		"--int", "-242",
		"--int-8", "-28",
		"--int-16", "-216",
		"--int-32", "-232",
		"--int-64", "-264",
		"--uint", "242",
		"--uint-8", "28",
		"--uint-16", "216",
		"--uint-32", "232",
		"--uint-64", "264",
		"--float-32", "232.32",
		"--float-64", "264.64",
		"--bool", "true",
	)

	return InStruct{
		FldString:  "flag-string",
		FldInt:     -242,
		FldInt8:    -28,
		FldInt16:   -216,
		FldInt32:   -232,
		FldInt64:   -264,
		FldUint:    242,
		FldUint8:   28,
		FldUint16:  216,
		FldUint32:  232,
		FldUint64:  264,
		FldFloat32: 232.32,
		FldFloat64: 264.64,
		FldParent: InStructParent{
			FldNested: InStructNested{
				FldBool: true,
			},
		},
	}
}

func stubJSON() InStruct {
	return InStruct{
		FldString:  "json-string",
		FldInt:     -342,
		FldInt8:    -38,
		FldInt16:   -316,
		FldInt32:   -332,
		FldInt64:   -364,
		FldUint:    342,
		FldUint8:   38,
		FldUint16:  316,
		FldUint32:  332,
		FldUint64:  364,
		FldFloat32: 332.32,
		FldFloat64: 364.64,
		FldParent: InStructParent{
			FldNested: InStructNested{
				FldBool: true,
			},
		},
	}
}

func stubYAML() InStruct {
	return InStruct{
		FldString:  "yaml-string",
		FldInt:     -442,
		FldInt8:    -48,
		FldInt16:   -416,
		FldInt32:   -432,
		FldInt64:   -464,
		FldUint:    442,
		FldUint8:   48,
		FldUint16:  416,
		FldUint32:  432,
		FldUint64:  464,
		FldFloat32: 432.32,
		FldFloat64: 464.64,
		FldParent: InStructParent{
			FldNested: InStructNested{
				FldBool: true,
			},
		},
	}
}

func stubTOML() InStruct {
	return InStruct{
		FldString:  "toml-string",
		FldInt:     -542,
		FldInt8:    -58,
		FldInt16:   -516,
		FldInt32:   -532,
		FldInt64:   -564,
		FldUint:    542,
		FldUint8:   58,
		FldUint16:  516,
		FldUint32:  532,
		FldUint64:  564,
		FldFloat32: 532.32,
		FldFloat64: 564.64,
		FldParent: InStructParent{
			FldNested: InStructNested{
				FldBool: true,
			},
		},
	}
}

func stubEnvFile() InStruct {
	return InStruct{
		FldString:  "env-file-string",
		FldInt:     -642,
		FldInt8:    -68,
		FldInt16:   -616,
		FldInt32:   -632,
		FldInt64:   -664,
		FldUint:    642,
		FldUint8:   68,
		FldUint16:  616,
		FldUint32:  632,
		FldUint64:  664,
		FldFloat32: 632.32,
		FldFloat64: 664.64,
		FldParent: InStructParent{
			FldNested: InStructNested{
				FldBool: true,
			},
		},
	}
}
