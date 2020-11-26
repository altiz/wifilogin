package irbis

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/godror/godror"
)

type Config struct {
	DataSource string `properties:"datasource,default="`
	Login      string `properties:"login,default="`
	Password   string `properties:"password,default="`
}

type Session struct {
	cfg Config
	db  *sqlx.DB
}

func NewSession(cfg Config) *Session {
	return &Session{cfg: cfg}
}

func (s *Session) Open() error {
	if s.cfg.Login == "" {
		return errors.New("Ошибка при подключнии к БД, не указан login ")
	}
	if s.cfg.Password == "" {
		return errors.New("Ошибка при подключнии к БД, не указан password ")
	}
	if s.cfg.DataSource == "" {
		return errors.New("Ошибка при подключнии к БД, не указан datasource ")
	}
	connString := fmt.Sprintf("%s/%s@%s", s.cfg.Login, s.cfg.Password, s.cfg.DataSource)
	db, err := sqlx.Connect("godror", connString)
	if err != nil {
		return errors.New(fmt.Sprintf("Ошибка при подключнии к БД %v@%v, %v", s.cfg.Login, s.cfg.DataSource, err))
	}
	s.db = db
	return nil
}

func (s *Session) Close() error {
	return s.db.Close()
}

func (s *Session) GetTime() (string, error) {
	if s.db == nil {
		return "", errors.New("Соединение с БД не установлено")
	}
	rows, err := s.db.Query("select sysdate from dual")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {
		rows.Scan(&thedate)
	}

	return thedate, nil
}

type ReciveAccount struct {
	AccountID    int64 `db:"ACCOUNT_ID"`
	ClientTypeID int64 `db:"CLIENT_TYPE_ID"`
}

const reciveAccountListSql = `SELECT ` +
	`    TCL.ACC ACCOUNT_ID, ` +
	`    TC.CLIENTTYPE_ID CLIENT_TYPE_ID ` +
	`FROM ` +
	`    BILLING.TCLAIM TCL, ` +
	`    BILLING.TACCOUNT TA, ` +
	`    BILLING.TCLIENT TC ` +
	`WHERE TCL.PASSEDTOLEGAL BETWEEN :BEGIN_PERIOD AND SYSDATE ` +
	`      AND TCL.ACC = TA.OBJECT_NO ` +
	`      AND TA.CLIENT_ID = TC.OBJECT_NO ` +
	`ORDER BY TCL.PASSEDTOLEGAL `

func (s *Session) GetReciveAccountList(beginPeriod time.Time) ([]ReciveAccount, error) {
	if s.db == nil {
		return nil, errors.New("Соединение с БД не установлено")
	}

	rows, err := s.db.Queryx(reciveAccountListSql, beginPeriod)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := ReciveAccount{}
	rl := []ReciveAccount{}
	for rows.Next() {
		err := rows.StructScan(&r)
		if err != nil {
			log.Fatal(err)
		}
		rl = append(rl, r)
		// log.Println(r)
	}

	return rl, nil
}

type NatAccount struct {
	AccountID              int64           `db:"ACCOUNT_ID"`
	ClientName             sql.NullString  `db:"CLIENT_NAME"`
	ClientTypeID           sql.NullInt64   `db:"CLIENTTYPE_ID"`
	ClientType             sql.NullString  `db:"CLIENTTYPE"`
	Address                sql.NullString  `db:"ADDRESS"`
	AccountNumb            sql.NullString  `db:"ACCOUNT_NUMB"`
	CompanyZoneID          sql.NullInt64   `db:"COMPANYZONE_ID"`
	CompanyZoneName        sql.NullString  `db:"COMPANYZONE_NAME"`
	BranchName             sql.NullString  `db:"BRANCH_NAME"`
	ProviderName           sql.NullString  `db:"PROVIDER_NAME"`
	AccountServiceTypeName sql.NullString  `db:"ACCOUNTSERVICETYPE_NAME"`
	EquipDebt              sql.NullFloat64 `db:"EQUIPDEBT"`
	ContractEnd            sql.NullTime    `db:"CONTRACT_END"`
	Debt                   sql.NullFloat64 `db:"DEBT"`
	DebtStart              sql.NullTime    `db:"DEBTSTART"`
	PassedToLegal          sql.NullTime    `db:"PASSEDTOLEGAL"`
	AddressConnect         sql.NullString  `db:"ADDRESS_CONNECT"`

	DocumentType       sql.NullString `db:"DOCUMENT_TYPE"`
	DocumentNumber     sql.NullString `db:"DOCUMENT_NUMBER"`
	DocumentGiven      sql.NullString `db:"DOCUMENT_GIVEN"`
	DocumentIssuerCode sql.NullString `db:"DOCUMENT_ISSUER_CODE"`
	DocumentGivenDate  sql.NullTime   `db:"DOCUMENT_GIVENDATE"`
}

const getNatAccountInfoSql = `SELECT ` +
	`       TA.OBJECT_NO ACCOUNT_ID, ` +
	`       TCL.CLIENT_NAME CLIENT_NAME, ` +
	`       TCT.OBJECT_NO CLIENTTYPE_ID, ` +
	`       TCT.CLIENTTYPE_NAME CLIENTTYPE, ` +
	`       TDT.NAME DOCUMENT_TYPE, ` +
	`       TCP.DOCUMENT_NUMBER DOCUMENT_NUMBER, ` +
	`       TCP.DOCUMENT_GIVEN DOCUMENT_GIVEN, ` +
	`       TCP.DOCUMENT_ISSUER_CODE DOCUMENT_ISSUER_CODE, ` +
	`       TCP.DOCUMENT_GIVENDATE DOCUMENT_GIVENDATE, ` +
	`       ADR.OBJECT_NAME ADDRESS, ` +
	`       TA.ACCOUNT_NUMB ACCOUNT_NUMB, ` +
	`       TCZ.OBJECT_NO COMPANYZONE_ID, ` +
	`       TCZ.COMPANYZONE_NAME COMPANYZONE_NAME, ` +
	`       TCB.BRANCH_NAME BRANCH_NAME, ` +
	`       TP.PROVIDER_SHORTNAME PROVIDER_NAME, ` +
	`       TAS.ACCOUNTSERVICETYPE_NAME ACCOUNTSERVICETYPE_NAME, ` +
	`       TC.EQUIPDEBT EQUIPDEBT, ` +
	`       (SELECT MAX(CONTRACT_END) FROM BILLING.TCONTRACTCOMMON TCC WHERE TCC.ACCOUNT_ID = TA.OBJECT_NO) CONTRACT_END, ` +
	`       TC.DEBT DEBT, ` +
	`       TC.DEBTSTART DEBTSTART, ` +
	`       TC.PASSEDTOLEGAL PASSEDTOLEGAL, ` +
	`       JEFFIT_GET_CONNECT_ADDRESS(TA.OBJECT_NO) ADDRESS_CONNECT ` +
	`FROM ` +
	`  BILLING.TCLAIM TC, ` +
	`  BILLING.TACCOUNT TA, ` +
	`  BILLING.TCOMPANYBRANCH TCB, ` +
	`  BILLING.TCOMPANYZONE TCZ, ` +
	`  BILLING.TACCOUNTSERVICETYPE TAS, ` +
	`  BILLING.TPROVIDER TP, ` +
	`  BILLING.TCLIENT TCL ` +
	`  LEFT JOIN BILLING.TCIVILPERSON TCP ON TCL.OBJECT_NO = TCP.OBJECT_NO ` +
	`  LEFT JOIN BILLING.TDOCUMENTTYPE TDT ON TCP.DOCTYPE_ID = TDT.OBJECT_NO ` +
	`  LEFT JOIN BILLING.TADDRESS ADR ON  ADR.OBJECT_NO = TCP.DOCUMENTADDRESS_ID ` +
	`  LEFT JOIN BILLING.TCLIENTTYPE TCT ON TCL.CLIENTTYPE_ID = TCT.OBJECT_NO ` +
	`WHERE TC.OBJECT_NO = TA.OBJECT_NO ` +
	`      AND TA.CLIENT_ID = TCL.OBJECT_NO ` +
	`      AND TA.COMPANYBRANCH_ID = TCB.OBJECT_NO ` +
	`      AND TCB.COMPANYZONE_ID = TCZ.OBJECT_NO ` +
	`      AND TA.ACCOUNTSERVICETYPE_ID = TAS.OBJECT_NO ` +
	`      AND TA.PROVIDER_ID = TP.OBJECT_NO ` +
	`      AND TA.OBJECT_NO = :ACCOUNT_ID `

func (s *Session) GetNatAccountInfo(id int64) (*NatAccount, error) {
	if s.db == nil {
		return nil, errors.New("Соединение с БД не установлено")
	}

	rows, err := s.db.Queryx(getNatAccountInfoSql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := NatAccount{}
	for rows.Next() {
		err := rows.StructScan(&r)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &r, nil
}

type JurAccount struct {
	AccountID              int64           `db:"ACCOUNT_ID"`
	ClientName             sql.NullString  `db:"CLIENT_NAME"`
	ClientTypeID           sql.NullInt64   `db:"CLIENTTYPE_ID"`
	ClientType             sql.NullString  `db:"CLIENTTYPE"`
	Address                sql.NullString  `db:"ADDRESS"`
	AccountNumb            sql.NullString  `db:"ACCOUNT_NUMB"`
	CompanyZoneID          sql.NullInt64   `db:"COMPANYZONE_ID"`
	CompanyZoneName        sql.NullString  `db:"COMPANYZONE_NAME"`
	BranchName             sql.NullString  `db:"BRANCH_NAME"`
	ProviderName           sql.NullString  `db:"PROVIDER_NAME"`
	AccountServiceTypeName sql.NullString  `db:"ACCOUNTSERVICETYPE_NAME"`
	EquipDebt              sql.NullFloat64 `db:"EQUIPDEBT"`
	ContractEnd            sql.NullTime    `db:"CONTRACT_END"`
	Debt                   sql.NullFloat64 `db:"DEBT"`
	DebtStart              sql.NullTime    `db:"DEBTSTART"`
	PassedToLegal          sql.NullTime    `db:"PASSEDTOLEGAL"`
	AddressConnect         sql.NullString  `db:"ADDRESS_CONNECT"`

	INN  sql.NullString `db:"INN"`
	OGRN sql.NullString `db:"OGRN"`
}

const getJurAccountInfoSql = `SELECT ` +
	`  TA.OBJECT_NO                             ACCOUNT_ID, ` +
	`  TCL.CLIENT_NAME                          CLIENT_NAME, ` +
	`  TCT.OBJECT_NO                            CLIENTTYPE_ID, ` +
	`  TCT.CLIENTTYPE_NAME                      CLIENTTYPE, ` +
	`  TLP.INN                                  INN, ` +
	`  TLP.OGRN                                 OGRN, ` +
	`  ADR.OBJECT_NAME                          ADDRESS, ` +
	`  TA.ACCOUNT_NUMB                          ACCOUNT_NUMB, ` +
	`  TCZ.OBJECT_NO                            COMPANYZONE_ID, ` +
	`  TCZ.COMPANYZONE_NAME                     COMPANYZONE_NAME, ` +
	`  TCB.BRANCH_NAME                          BRANCH_NAME, ` +
	`  TP.PROVIDER_SHORTNAME                    PROVIDER_NAME, ` +
	`  TAS.ACCOUNTSERVICETYPE_NAME              ACCOUNTSERVICETYPE_NAME, ` +
	`  TC.EQUIPDEBT                             EQUIPDEBT, ` +
	`  (SELECT MAX(CONTRACT_END) ` +
	`   FROM BILLING.TCONTRACTCOMMON TCC ` +
	`   WHERE TCC.ACCOUNT_ID = TA.OBJECT_NO)    CONTRACT_END, ` +
	`  TC.DEBT                                  DEBT, ` +
	`  TC.DEBTSTART                             DEBTSTART, ` +
	`  TC.PASSEDTOLEGAL                         PASSEDTOLEGAL, ` +
	`  JEFFIT_GET_CONNECT_ADDRESS(TA.OBJECT_NO) ADDRESS_CONNECT ` +
	`FROM ` +
	`  BILLING.TCLAIM TC, ` +
	`  BILLING.TACCOUNT TA, ` +
	`  BILLING.TCOMPANYBRANCH TCB, ` +
	`  BILLING.TCOMPANYZONE TCZ, ` +
	`  BILLING.TACCOUNTSERVICETYPE TAS, ` +
	`  BILLING.TPROVIDER TP, ` +
	`  BILLING.TCLIENT TCL ` +
	`  LEFT JOIN BILLING.TLEGALPERSON TLP ON TCL.OBJECT_NO = TLP.OBJECT_NO ` +
	`  LEFT JOIN BILLING.TADDRESS ADR ON ADR.OBJECT_NO = TLP.ADDRESS_ID ` +
	`  LEFT JOIN BILLING.TCLIENTTYPE TCT ON TCL.CLIENTTYPE_ID = TCT.OBJECT_NO ` +
	`WHERE TC.OBJECT_NO = TA.OBJECT_NO ` +
	`      AND TA.CLIENT_ID = TCL.OBJECT_NO ` +
	`      AND TA.COMPANYBRANCH_ID = TCB.OBJECT_NO ` +
	`      AND TCB.COMPANYZONE_ID = TCZ.OBJECT_NO ` +
	`      AND TA.ACCOUNTSERVICETYPE_ID = TAS.OBJECT_NO ` +
	`      AND TA.PROVIDER_ID = TP.OBJECT_NO ` +
	`      AND TA.OBJECT_NO = :ACCOUNT_ID `

func (s *Session) GetJurAccountInfo(id int64) (*JurAccount, error) {
	if s.db == nil {
		return nil, errors.New("Соединение с БД не установлено")
	}

	rows, err := s.db.Queryx(getJurAccountInfoSql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := JurAccount{}
	for rows.Next() {
		err := rows.StructScan(&r)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &r, nil
}

//AccountExecutionInfo - информация л/с по исполнительному производству
type AccountExecutionInfo struct {
	AccountID         int64          `db:"ACCOUNT_ID"`
	ExecuteActionDate sql.NullTime   `db:"EXECUTE_ACTION_DATE"` // дата
	ExecuteAction     sql.NullString `db:"EXECUTE_ACTION"`      // действие
}

const selectExecutionChangeListSQL = `SELECT TC.ACC ACCOUNT_ID, ` +
	`    TC.EXECUTEACTIONDATE EXECUTE_ACTION_DATE, ` +
	`    EX.NAME EXECUTE_ACTION ` +
	`FROM BILLING.TCLAIM TC ` +
	`    LEFT JOIN BILLING.VDEBT_EXECUTEACTION EX ON TC.EXECUTEACTION = EX.ID ` +
	`WHERE TC.ACC IN ( ` +
	`        SELECT ACCOUNT_ID ` +
	`        FROM EXECUTEACTION_UPDATE_TRIGGER ` +
	`        WHERE UPDATE_DATE BETWEEN :BEGIN_PERIOD AND SYSDATE ` +
	`	) `

func (s *Session) GetExecutionChangeList(beginPeriod time.Time) ([]AccountExecutionInfo, error) {
	if s.db == nil {
		return nil, errors.New("Соединение с БД не установлено")
	}

	rows, err := s.db.Queryx(selectExecutionChangeListSQL, beginPeriod)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ei := AccountExecutionInfo{}
	eil := []AccountExecutionInfo{}
	for rows.Next() {
		err := rows.StructScan(&ei)
		if err != nil {
			log.Fatal(err)
		}
		eil = append(eil, ei)
	}
	return eil, nil
}

//AccountPaymentInfo - информация л/с по платежам
type AccountPaymentInfo struct {
	AccountID            int64           `db:"ACCOUNT_ID"`
	PrincipalDebt        sql.NullFloat64 `db:"PRINCIPALDEBT"`         //Оплата основного долга
	StateDutySum         sql.NullFloat64 `db:"STATEDUTY_SUM"`         //Оплата госпошлины
	PenaltySum           sql.NullFloat64 `db:"PENALTY_SUM"`           //Оплата неустойки
	ClaimAmountRemainder sql.NullFloat64 `db:"CLAIMAMOUNT_REMAINDER"` //Остаток долга по иску
	UpdateDate           sql.NullTime    `db:"UPDATE_DATE"`           //Дата последнего изменения
}

const selectPaymentChangeListSQL = `SELECT ACCOUNT_ID, ` +
	`    PRINCIPALDEBT, ` +
	`    STATEDUTY_SUM, ` +
	`    PENALTY_SUM, ` +
	`    CLAIMAMOUNT_REMAINDER, ` +
	`    UPDATE_DATE ` +
	`FROM PAYMENT_UPDATE ` +
	`WHERE UPDATE_DATE BETWEEN :BEGIN_PERIOD AND SYSDATE `

func (s *Session) GetPaymentChangeList(beginPeriod time.Time) ([]AccountPaymentInfo, error) {
	if s.db == nil {
		return nil, errors.New("Соединение с БД не установлено")
	}

	rows, err := s.db.Queryx(selectPaymentChangeListSQL, beginPeriod)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pi := AccountPaymentInfo{}
	pil := []AccountPaymentInfo{}
	for rows.Next() {
		err := rows.StructScan(&pi)
		if err != nil {
			log.Fatal(err)
		}
		pil = append(pil, pi)
	}
	return pil, nil
}
