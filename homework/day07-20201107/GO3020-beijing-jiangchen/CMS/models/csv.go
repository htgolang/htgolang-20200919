package models

import (
	"CMS/consts"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

//CSVDB ...
type CSVDB struct {
}

//Check ...
//CSVDB Check
func (csvdb *CSVDB) Check() (err error) {
	return nil
}

//Init ...
//CSVDB Initialization
func (csvdb *CSVDB) Init() (err error) {
	// check if file exists
	info, err := os.Stat(consts.DBDir + "/" + consts.DBType + "/Users.csv")
	if err != nil {
		fmt.Printf("csvfile %s does not exists, create a new file\n", consts.DBDir+"/"+consts.DBType+"/Users.csv")
	}
	os.MkdirAll(consts.DBDir+"/"+consts.DBType, 0755)

	if os.IsNotExist(err) || info.IsDir() {
		csvFile, err := os.Create(consts.DBDir + "/" + consts.DBType + "/Users.csv")
		defer csvFile.Close()
		if err != nil {
			log.Fatal(err)
		}

		writer := csv.NewWriter(csvFile)
		for _, user := range Users {
			line := []string{strconv.Itoa(user.UserID), user.Name, user.Tel, user.Address, user.Birthday.Format("2006-01-02"), string(user.Password[:])}
			err := writer.Write(line)
			if err != nil {
				return err
			}
		}
		writer.Flush()
	}
	return nil
}

//SyncFromDBToMemory ...
// Sync Data from CSVDB to Memory
func (csvdb *CSVDB) SyncFromDBToMemory() (err error) {
	csvFile, err := os.Open(consts.DBDir + "/" + consts.DBType + "/Users.csv")
	defer csvFile.Close()
	if err != nil {
		return err
	}

	// empty Users slice before Sync
	Users = nil

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	for _, record := range records {
		*MUE.Users = append(*MUE.Users, User{UserID: func() (ret int) { ret, _ = strconv.Atoi(record[0]); return }(), Name: record[1], Tel: record[2], Address: record[3], Birthday: func() (ret time.Time) { ret, _ = time.Parse("2006-01-02", record[4]); return }(), Password: func() (ret [16]byte) { copy(ret[:], ([]byte(record[5]))[:16]); return }()})
	}
	return nil
}

//SyncFromMemoryToDB ...
// Sync Data from Memory to CSVDB
func (csvdb *CSVDB) SyncFromMemoryToDB() (err error) {
	csvFile, err := os.Create(consts.DBDir + "/" + consts.DBType + "/Users.csv")
	defer csvFile.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(csvFile)
	for idx, user := range Users {
		if idx < len(Users)-5 {
			continue
		}
		line := []string{strconv.Itoa(user.UserID), user.Name, user.Tel, user.Address, user.Birthday.Format("2006-01-02"), string(user.Password[:])}
		err := writer.Write(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

//InsertToDB ...
// Insert Data into CSVDB directly
func (csvdb *CSVDB) InsertToDB(element *User) (err error) {
	// csvFile, err := os.OpenFile(consts.DBDir+"/"+consts.DBType+"/Users.csv", os.O_WRONLY|os.O_CREATE, 0644)
	csvFile, err := os.Create(consts.DBDir + "/" + consts.DBType + "/Users.csv")
	defer csvFile.Close()
	if err != nil {
		return err
	}
	if len(Users) > 5 {
		*MUE.Users = (*MUE.Users)[1:]
	}
	writer := csv.NewWriter(csvFile)
	// err = writer.Write([]string{strconv.Itoa(element.UserID), element.Name, element.Tel, element.Address, element.Birthday.Format("2006-01-02"), string(element.Password[:])})
	// if err != nil {
	// 	return err
	// }
	for idx, user := range Users {
		if idx < len(Users)-5 {
			continue
		}
		line := []string{strconv.Itoa(user.UserID), user.Name, user.Tel, user.Address, user.Birthday.Format("2006-01-02"), string(user.Password[:])}
		err := writer.Write(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

//DeleteFromDB ...
// Delete Data from CSVDB
// in fact, it's identical to SyncFromMemoryToDB ......
func (csvdb *CSVDB) DeleteFromDB(id int) (err error) {
	csvFile, err := os.Create(consts.DBDir + "/" + consts.DBType + "/Users.csv")
	defer csvFile.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(csvFile)
	for _, user := range Users {
		line := []string{strconv.Itoa(user.UserID), user.Name, user.Tel, user.Address, user.Birthday.Format("2006-01-02"), string(user.Password[:])}
		err := writer.Write(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

//ModifyFromDB ...
// Modify Data to CSVDB
// in fact, it's identical to SyncFromMemoryToDB ......
func (csvdb *CSVDB) ModifyFromDB(element *User) (err error) {
	csvFile, err := os.Create(consts.DBDir + "/" + consts.DBType + "/Users.csv")
	defer csvFile.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(csvFile)
	for _, user := range Users {
		line := []string{strconv.Itoa(user.UserID), user.Name, user.Tel, user.Address, user.Birthday.Format("2006-01-02"), string(user.Password[:])}
		err := writer.Write(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

//GetNonAdmin ...
// Get Non-Admin data from CSVDB
func (csvdb *CSVDB) GetNonAdmin() (err error) {
	csvFile, err := os.Open(consts.DBDir + "/" + consts.DBType + "/Users.csv")
	defer csvFile.Close()
	if err != nil {
		return err
	}

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	for _, record := range records {
		if record[1] == "admin" {
			continue
		} else {
			*MUE.Users = append(*MUE.Users, User{UserID: func() (ret int) { ret, _ = strconv.Atoi(record[0]); return }(), Name: record[1], Tel: record[2], Address: record[3], Birthday: func() (ret time.Time) { ret, _ = time.Parse("2006-01-02", record[4]); return }(), Password: func() (ret [16]byte) { copy(ret[:], ([]byte(record[5]))[:16]); return }()})
		}
	}
	return nil
}
