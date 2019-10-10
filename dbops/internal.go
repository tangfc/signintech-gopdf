package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"go-test/pdf/defs"
	"log"
)

// 获取用户信息
func GetMemberInfo() ([]*defs.MemberInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, username, nickname, realname, CASE sex WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '保密' END AS sex, signature FROM tb_member")
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	var data []*defs.MemberInfo
	for rows.Next() {
		var (
			id        int
			nickname  string
			username  string
			realname  string
			sex       string
			signature string
		)
		err := rows.Scan(&id, &username, &nickname, &realname, &sex, &signature);
		if err != nil {
			log.Printf("error: %s", err.Error())
			return nil, err
		}
		data = append(data, &defs.MemberInfo{id, nickname, username, realname, sex, signature})
	}
	return data, nil
}
