package funcs

import (
	"CMS/consts"
	"CMS/miscs"
	"CMS/models"
	"flag"
	"fmt"
	"os"
)

//ParseArgs ...
func ParseArgs() {
	flag.StringVar(&consts.DBType, "db", "memory", "persistent db type (boltdb, csv, json, memory), default is memory.")
	flag.BoolVar(&consts.InitFlag, "init", false, "whether init database or not, default is false.")
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "  NOTE: if db is \"all\" combined with \"-init\", delete all database.\n")
	}
	flag.Parse()
}

//InitDB ...
func InitDB() {

	var check string

	if consts.InitFlag && consts.DBType != "memory" {
		if consts.DBType == "all" {
			os.RemoveAll(consts.DBDir)
			return
		}
		if _, err := os.Stat(consts.DBDir + "/" + consts.DBType); os.IsNotExist(err) {
			return
		}
		if ok, _ := miscs.IsDirEmpty(consts.DBDir + "/" + consts.DBType); !ok {
			for {
				fmt.Print("db has already been initialized, do you want to re-initialize? (y/n)")
				fmt.Scanln(&check)
				if check != "y" && check != "n" {
					continue
				} else if check == "n" {
					fmt.Println("abort...")
					return
				} else {
					os.RemoveAll(consts.DBDir + "/" + consts.DBType)
					return
				}
			}
		}
	} else if consts.DBType == "memory" {
		fmt.Println("no need to init when using memory")
	}

	switch consts.DBType {
	case "boltdb":
		models.MUE.PersistentStorage = new(models.BoltDB)
		models.MUE.Users = &models.Users
		models.MUE.PersistentStorage.Init()
	case "csv":
		models.MUE.PersistentStorage = new(models.CSVDB)
		models.MUE.Users = &models.Users
		models.MUE.PersistentStorage.Init()
	case "json":
		models.MUE.PersistentStorage = new(models.JSONDB)
		models.MUE.Users = &models.Users
		models.MUE.PersistentStorage.Init()
	case "memory":
		models.MUE.PersistentStorage = new(models.MemoryDB)
		models.MUE.Users = &models.Users
		models.MUE.PersistentStorage.Init()
	default:
		models.MUE.PersistentStorage = new(models.MemoryDB)
		models.MUE.Users = &models.Users
		models.MUE.PersistentStorage.Init()
	}
}
