package handlers

import (
	"net/http"
	_ "wifilogin/cmd/wifilogin/docs" // docs is generated by Swag CLI, you have to import it.
	"wifilogin/cmd/wifilogin/models"

	logs "github.com/sirupsen/logrus"

	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	// swagger embed files

	_ "github.com/godror/godror"
)

/*type id struct {
	id int
}

type login struct {
	login  string
	passwd string
}*/

// set-msisdn godoc
// @Summary Сохраняет данные по делам в биллинговой системе
// @Description Передача данных по делам из Jeffit во внешнюю систему
// @Accept  json
// @Produce json
// Param [param_name] [param_type] [data_type] [required/mandatory] [description]
// @Param q body models.TDelo false "name search by q"
// @Success 200 {object} models.TData_resp
// @Router //set-msisdn/ [post]
func Setmsisdn_test(c *gin.Context) {

	var req models.Tlogin_req
	var resp models.Tlogin_resp
	logs.SetFormatter(&logs.JSONFormatter{})

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(req)

	db, err := sql.Open("godror", `user="wifiservice" password="wifi" connectString="e-scan:1521/irbis" poolMaxSessions=2000 poolIncrement=15 standaloneConnection=1`)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(0)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("OK")
	rows, err := db.Query("select wifi_02_login_seq_test.nextval id from dual")
	if err != nil {
		panic(err)
	}
	id := id{}
	rows.Next()
	defer rows.Close()
	err1 := rows.Scan(&id.id)
	if err1 != nil {
		fmt.Println(err)
	}
	rows.Close()
	fmt.Println(id)
	fmt.Println("OK")
	rows, err2 := db.Query("SELECT login, passwd FROM wifi_02_login_test where id = " + strconv.Itoa(id.id))
	if err2 != nil {
		panic(err)
	}

	logins := login{}
	rows.Next()
	err3 := rows.Scan(&logins.login, &logins.passwd)
	if err3 != nil {
		fmt.Println(err)
	}
	rows.Close()
	db.Exec("update wifi_02_login_test set  bdate = sysdate,  msisdn = :1 where   id = :2", req.Msisdn, id.id)
	db.Close()
	fmt.Println("OK")
	defer db.Close()
	fmt.Println(logins)

	resp.Login = logins.login
	resp.Passwd = logins.passwd
	c.JSON(http.StatusOK, resp)
}
