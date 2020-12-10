package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"zhao/db"
)

func StoreToFile(name, path string, file multipart.File) error {
	localpath := filepath.Join("file", path)
	os.MkdirAll(localpath, os.ModePerm)

	fileinfs, err := ioutil.ReadDir(localpath)
	if err != nil {
		return fmt.Errorf("%v read %s false", err, path)
	}
	for _, fileinfo := range fileinfs {
		if fileinfo.Name() == name {
			return fmt.Errorf("%s%s 文件名重复", path, name)
		}
	}

	dfile, err := os.Create(filepath.Join(localpath, name))
	defer dfile.Close()
	defer file.Close()
	if err != nil {
		return fmt.Errorf("%v 打开目标文件路径失败", err)
	}
	_, err = io.Copy(dfile, file)
	if err != nil {
		return fmt.Errorf("%s 文件传输失败", err)
	}

	return nil
}

func StoreToMysql(name string, file multipart.File) error {
	sql2 := fmt.Sprintf(`
		create table if not exists %s (
			id bigint primary key auto_increment,
			ip varchar(15) not null default '',
			code varchar(3) not null default '900',
			meth varchar(10) not null default '',
			key idx_ip(ip),
			key idx_code(code)
		)engine innodb charset utf8mb4`, name)
	_, err := db.DB.Exec(sql2)
	if err != nil {
		fmt.Println(err, "创建表错误")
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		ip := fields[0]
		code := fields[8]
		meth := strings.Trim(fields[5], `"`)
		sql := fmt.Sprintf(`insert into %s(ip, code, meth) values(?,?,?)`, name)
		result, err := db.DB.Exec(sql, ip, code, meth)
		if err != nil {
			return fmt.Errorf("%v 数据库exec错误 sql:%v", err, sql)
		}

		n, err := result.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return err
		}
		if n != 1 {
			return fmt.Errorf("修改了%d行", n)
		}
	}

	sql := "insert into info(name) values(?)"
	fmt.Println(sql)
	result, err := db.DB.Exec(sql, name)
	if err != nil {
		fmt.Println(err, "db Exec false -- sql", sql)
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if n != 1 {
		return fmt.Errorf("修改了%d行", n)
	}

	return nil
}

func QueryFileSlice() []string {
	fileSlice := make([]string, 0, 10)
	sql := `select name from info`
	rows, err := db.DB.Query(sql)
	if err != nil {
		fmt.Printf("%v, query info efalse sql :%s\n", err, sql)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Printf("%v 查询文件名错误", err)
			return nil
		}
		fileSlice = append(fileSlice, name)
	}
	return fileSlice
}

func QueryField(tablename string, field string, channel chan [][]string) [][]string {
	res := make([][]string, 0, 10)
	sql := fmt.Sprintf(`
	select %s ,count(*) as number 
	from %s 
	group by %s
	order by number desc
	limit 10`, field, tablename, field)
	rows, err := db.DB.Query(sql)
	if err != nil {
		fmt.Printf("%v query %s false sql : %s", err, field, sql)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name  string
			count int
		)
		err := rows.Scan(&name, &count)
		if err != nil {
			fmt.Printf("%v, row scan err sql: %s", err, sql)
			return nil
		}
		res = append(res, []string{name, strconv.Itoa(count)})
	}
	channel <- res
	return res
}
