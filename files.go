package yamltypes

import (
	"fmt"
	"os"
)

// Upon unmarshalling, verifies that the path is valid
type ExistingPath string

// Upon unmarshalling, verifies that the path refers to an existing file.
type ExistingFile string

// Upon unmarshalling, verifies that the path refers to an existing directory.
type ExistingDir string

// Upon unmarshalling, verifies that the path refers to an executable file.
type ExecutableFile string

type OpenFile os.File

type CreateFile os.File

func checkStat(f func(s os.FileInfo) error) func(p string) error {
	return func(p string) error {
		if stats, err := os.Stat(p); err == nil {
			return f(stats)
		} else {
			return err
		}
	}
}

func (f *OpenFile) UnmarshalYAML(u func(interface{}) error) error {
	var s string
	if err := u(&s); err != nil {
		return err
	}

	file, err := os.Open(s)
	if err != nil {
		return err
	}

	*f = OpenFile(*file)
	return nil
}

func (f *CreateFile) UnmarshalYAML(u func(interface{}) error) error {
	var s string
	if err := u(&s); err != nil {
		return err
	}

	file, err := os.Create(s)
	if err != nil {
		return err
	}

	*f = CreateFile(*file)
	return nil
}

func (f *ExistingPath) UnmarshalYAML(u func(interface{}) error) error {
	return unmarshalStringAndValidate((*string)(f), u, checkStat(func(stat os.FileInfo) error {
		return nil
	}))
}

func (f *ExistingFile) UnmarshalYAML(u func(interface{}) error) error {
	return unmarshalStringAndValidate((*string)(f), u, checkStat(func(stat os.FileInfo) error {
		if stat.IsDir() {
			return fmt.Errorf("%s is not a file, but a directory.", stat.Name())
		} else {
			return nil
		}
	}))
}

func (f *ExistingDir) UnmarshalYAML(u func(interface{}) error) error {
	return unmarshalStringAndValidate((*string)(f), u, checkStat(func(stat os.FileInfo) error {
		if !stat.IsDir() {
			return fmt.Errorf("%s is not a directory.", stat.Name())
		} else {
			return nil
		}
	}))
}

func (f *ExecutableFile) UnmarshalYAML(u func(interface{}) error) error {
	return unmarshalStringAndValidate((*string)(f), u, checkStat(func(stat os.FileInfo) error {
		if stat.IsDir() || stat.Mode().Perm()&0111 == 0 {
			return fmt.Errorf("%s is not an executable file.", stat.Name())
		} else {
			return nil
		}
	}))
}
