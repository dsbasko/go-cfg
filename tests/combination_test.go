package tests

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	gocfg "github.com/dsbasko/go-cfg"
)

const (
	ENV     = "env"
	Flag    = "flag"
	YAML    = "yaml"
	TOML    = "toml"
	JSON    = "json"
	ENVFile = "env-file"
)

func Test_Default(t *testing.T) {
	var structPtr InStruct
	gocfg.MustReadEnv(&structPtr)
	wantStruct := stubDefault()
	assert.EqualValues(t, wantStruct, structPtr)
}

func Test_Once_Env(t *testing.T) {
	var structPtr InStruct
	wantStruct := stubEnv()
	gocfg.MustReadEnv(&structPtr)
	assert.Equal(t, wantStruct, structPtr)
}

func Test_Once_Env_Panic(t *testing.T) {
	assert.Panics(t, func() {
		gocfg.MustReadEnv("not-a-pointer")
	})
}

func Test_Once_Flag(t *testing.T) {
	var structPtr InStruct
	wantStruct := stubFlag()
	gocfg.MustReadFlag(&structPtr)
	assert.Equal(t, wantStruct, structPtr)
}

func Test_Once_Flag_Panic(t *testing.T) {
	assert.Panics(t, func() {
		gocfg.MustReadFlag("not-a-pointer")
	})
}

func Test_Once_JSON(t *testing.T) {
	var structPtr InStruct
	wantStruct := stubJSON()
	gocfg.MustReadFile("stub.json", &structPtr)
	assert.Equal(t, wantStruct, structPtr)
}

func Test_Once_File_Panic(t *testing.T) {
	assert.Panics(t, func() {
		gocfg.MustReadFile("", "not-a-pointer")
	})
}

func Test_Once_YAML(t *testing.T) {
	var structPtr InStruct
	wantStruct := stubYAML()
	gocfg.MustReadFile("stub.yaml", &structPtr)
	assert.Equal(t, wantStruct, structPtr)
}

func Test_Once_TOML(t *testing.T) {
	var structPtr InStruct
	wantStruct := stubTOML()
	gocfg.MustReadFile("stub.toml", &structPtr)
	assert.Equal(t, wantStruct, structPtr)
}

func Test_Once_ENV(t *testing.T) {
	var structPtr InStruct
	wantStruct := stubEnvFile()
	gocfg.MustReadFile("stub.env", &structPtr)
	assert.Equal(t, wantStruct, structPtr)
}

func Test_Permute_All(t *testing.T) {
	test := []string{ENV, Flag, JSON, YAML, TOML, ENVFile}
	permute := NewPermute(test)

	for _, perm := range permute.items {
		name := strings.Join(perm, ",")
		t.Run(name, func(t *testing.T) {
			var structPtr, wantStruct InStruct

			for _, p := range perm {
				switch p {
				case ENV:
					gocfg.MustReadEnv(&structPtr)
				case Flag:
					gocfg.MustReadFlag(&structPtr)
				case JSON:
					gocfg.MustReadFile("stub.json", &structPtr)
				case YAML:
					gocfg.MustReadFile("stub.yaml", &structPtr)
				case TOML:
					gocfg.MustReadFile("stub.toml", &structPtr)
				case ENVFile:
					gocfg.MustReadFile("stub.env", &structPtr)
				default:
					t.Fatalf("Unknown permutation: %s", p)
				}
			}

			switch perm[len(perm)-1] {
			case ENV:
				wantStruct = stubEnv()
			case Flag:
				wantStruct = stubFlag()
			case JSON:
				wantStruct = stubJSON()
			case YAML:
				wantStruct = stubYAML()
			case TOML:
				wantStruct = stubTOML()
			case ENVFile:
				wantStruct = stubEnvFile()
			default:
				t.Fatalf("Unknown permutation: %s", perm[len(perm)-1])
			}

			assert.EqualValues(t, wantStruct, structPtr)
		})
	}
}

func Test_Permute_Overlay_Env_flag(t *testing.T) {
	test := []string{ENV, ENV, Flag, Flag}
	permute := NewPermute(test)

	for _, perm := range permute.items {
		name := strings.Join(perm, ",")
		t.Run(name, func(t *testing.T) {
			var structPtr, wantStruct InStruct

			for _, p := range perm {
				switch p {
				case ENV:
					gocfg.MustReadEnv(&structPtr)
				case Flag:
					gocfg.MustReadFlag(&structPtr)
				default:
					t.Fatalf("Unknown permutation: %s", p)
				}
			}

			switch perm[len(perm)-1] {
			case ENV:
				wantStruct = stubEnv()
			case Flag:
				wantStruct = stubFlag()
			default:
				t.Fatalf("Unknown permutation: %s", perm[len(perm)-1])
			}

			assert.EqualValues(t, wantStruct, structPtr)
		})
	}
}

func Test_Permute_Overlay_JSON(t *testing.T) {
	test := []string{ENV, Flag, JSON, JSON}
	permute := NewPermute(test)

	for _, perm := range permute.items {
		name := strings.Join(perm, ",")
		t.Run(name, func(t *testing.T) {
			var structPtr, wantStruct InStruct

			for _, p := range perm {
				switch p {
				case ENV:
					gocfg.MustReadEnv(&structPtr)
				case Flag:
					gocfg.MustReadFlag(&structPtr)
				case JSON:
					gocfg.MustReadFile("stub.json", &structPtr)
				}
			}

			switch perm[len(perm)-1] {
			case ENV:
				wantStruct = stubEnv()
			case Flag:
				wantStruct = stubFlag()
			case JSON:
				wantStruct = stubJSON()
			case YAML:
				wantStruct = stubYAML()
			case TOML:
				wantStruct = stubTOML()
			case ENVFile:
				wantStruct = stubEnvFile()
			default:
				t.Fatalf("Unknown permutation: %s", perm[len(perm)-1])
			}

			assert.EqualValues(t, wantStruct, structPtr)
		})
	}
}

func Test_Permute_Overlay_YAML(t *testing.T) {
	test := []string{ENV, ENV, Flag, Flag}
	permute := NewPermute(test)

	for _, perm := range permute.items {
		name := strings.Join(perm, ",")
		t.Run(name, func(t *testing.T) {
			var structPtr, wantStruct InStruct

			for _, p := range perm {
				switch p {
				case ENV:
					gocfg.MustReadEnv(&structPtr)
				case Flag:
					gocfg.MustReadFlag(&structPtr)
				case YAML:
					gocfg.MustReadFile("stub.yaml", &structPtr)
				default:
					t.Fatalf("Unknown permutation: %s", p)
				}
			}

			switch perm[len(perm)-1] {
			case ENV:
				wantStruct = stubEnv()
			case Flag:
				wantStruct = stubFlag()
			case YAML:
				wantStruct = stubYAML()
			default:
				t.Fatalf("Unknown permutation: %s", perm[len(perm)-1])
			}

			assert.EqualValues(t, wantStruct, structPtr)
		})
	}
}

func Test_Permute_Overlay_TOML(t *testing.T) {
	test := []string{ENV, Flag, TOML, TOML}
	permute := NewPermute(test)

	for _, perm := range permute.items {
		name := strings.Join(perm, ",")
		t.Run(name, func(t *testing.T) {
			var structPtr, wantStruct InStruct

			for _, p := range perm {
				switch p {
				case ENV:
					gocfg.MustReadEnv(&structPtr)
				case Flag:
					gocfg.MustReadFlag(&structPtr)
				case TOML:
					gocfg.MustReadFile("stub.toml", &structPtr)
				default:
					t.Fatalf("Unknown permutation: %s", p)
				}
			}

			switch perm[len(perm)-1] {
			case ENV:
				wantStruct = stubEnv()
			case Flag:
				wantStruct = stubFlag()
			case TOML:
				wantStruct = stubTOML()
			default:
				t.Fatalf("Unknown permutation: %s", perm[len(perm)-1])
			}

			assert.EqualValues(t, wantStruct, structPtr)
		})
	}
}

func Test_Permute_Overlay_Env(t *testing.T) {
	test := []string{ENV, Flag, ENV, ENV}
	permute := NewPermute(test)

	for _, perm := range permute.items {
		name := strings.Join(perm, ",")
		t.Run(name, func(t *testing.T) {
			var structPtr, wantStruct InStruct

			for _, p := range perm {
				switch p {
				case ENV:
					gocfg.MustReadEnv(&structPtr)
				case Flag:
					gocfg.MustReadFlag(&structPtr)
				case ENVFile:
					gocfg.MustReadFile("stub.env", &structPtr)
				default:
					t.Fatalf("Unknown permutation: %s", p)
				}
			}

			switch perm[len(perm)-1] {
			case ENV:
				wantStruct = stubEnv()
			case Flag:
				wantStruct = stubFlag()
			case ENVFile:
				wantStruct = stubEnvFile()
			default:
				t.Fatalf("Unknown permutation: %s", perm[len(perm)-1])
			}

			assert.EqualValues(t, wantStruct, structPtr)
		})
	}
}

type Permute struct {
	items      [][]string
	wantStruct []InStruct
}

func NewPermute(data []string) *Permute {
	permuteItems := permuteFn(data)
	permute := Permute{items: permuteItems}

	permute.wantStruct = make([]InStruct, len(permuteItems))
	for i, perm := range permuteItems {
		lastItem := perm[len(perm)-1]
		switch lastItem {
		case ENV:
			permute.wantStruct[i] = stubEnv()
		case Flag:
			permute.wantStruct[i] = stubFlag()
		case JSON:
			permute.wantStruct[i] = stubJSON()
		case YAML:
			permute.wantStruct[i] = stubYAML()
		case TOML:
			permute.wantStruct[i] = stubTOML()
		case ENVFile:
			permute.wantStruct[i] = stubEnvFile()

		default:
			log.Fatalf("Unknown permutation: %s", lastItem)
		}
	}

	return &permute
}

func permuteFn(data []string) [][]string {
	if len(data) == 1 {
		return [][]string{data}
	}

	var permutations [][]string
	for i, v := range data {
		temp := make([]string, len(data))
		copy(temp, data)
		temp = append(temp[:i], temp[i+1:]...)
		for _, perm := range permuteFn(temp) {
			permutations = append(permutations, append([]string{v}, perm...))
		}
	}
	return permutations
}
