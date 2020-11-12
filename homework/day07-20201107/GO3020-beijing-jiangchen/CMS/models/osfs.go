package models

import (
	"os"
	"time"

	"github.com/spf13/afero"
)

//OsFs ...
type OsFs struct{}

//Name ...
func (OsFs) Name() string { return "OsFs" }

//Create ...
func (OsFs) Create(name string) (afero.File, error) { return os.Create(name) }

//Mkdir ...
func (OsFs) Mkdir(name string, perm os.FileMode) error { return os.Mkdir(name, perm) }

//MkdirAll ...
func (OsFs) MkdirAll(path string, perm os.FileMode) error { return os.MkdirAll(path, perm) }

//Open ...
func (OsFs) Open(name string) (afero.File, error) { return os.Open(name) }

//OpenFile ...
func (OsFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	return os.OpenFile(name, flag, perm)
}

//Remove ...
func (OsFs) Remove(name string) error { return os.Remove(name) }

//RemoveAll ...
func (OsFs) RemoveAll(path string) error { return os.RemoveAll(path) }

//Rename ...
func (OsFs) Rename(oldname, newname string) error { return os.Rename(oldname, newname) }

//Stat ...
func (OsFs) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }

//Chmod ...
func (OsFs) Chmod(name string, mode os.FileMode) error { return os.Chmod(name, mode) }

//Chtimes ...
func (OsFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}
