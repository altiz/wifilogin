package models

// Data Model
type Tlogin_req struct {
	Msisdn string `gorm:"column:msisdn" json:"msisdn"`
}

type Tlogin_resp struct {
	Login  string `gorm:"column:login" json:"login"`
	Passwd string `gorm:"column:passwd" json:"passwd"`
}

// Data Model
type TData_req struct {
	User string `gorm:"column:first_name" json:"user"`
}

type TData_resp struct {
	Status int `gorm:"column:status_id" json:"status"`
}

type TVersion struct {
	BuildTime string `gorm:"column:BuildTime" json:"buildTime"`
	Commit    string `gorm:"column:Commit" json:"commit"`
	Release   string `gorm:"column:first_name" json:"release"`
}

type Debug struct {
	IsDebug int
}

/*
type TDecisions struct {
	Message      string `gorm:"column:message" json:"message"`
	decisionDate string `gorm:"column:decisionDate" json:"decisionDate"`
}

type TAssignee struct {
	Id    string `gorm:"column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
}

type TRepresentatives struct {
	Id    string `gorm:"column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
}

type TClient struct {
	Id              string           `gorm:"column:id" json:"id"`
	Name            string           `gorm:"column:name" json:"name"`
	Email           string           `gorm:"column:email" json:"email"`
	Type            string           `gorm:"column:type" json:"type"`
	Representatives TRepresentatives `gorm:"column:representative" json:"representative"`
}

type TClaimApplicant struct {
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
}

type TClaimRecipient struct {
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
}
type TLinks struct {
	Link string `gorm:"column:link" json:"link"`
	Name string `gorm:"column:name" json:"name"`
}

type TDelo struct {
	//внешний id дела <Строковое значение>
	Id string `gorm:"column:id" json:"id"`
	//внешний id категории <Строковое значение>
	ServiceId string `gorm:"column:serviceId" json:"serviceId"`
	//название категории <Строковое значение>
	ServiceName string `gorm:"column:serviceName" json:"serviceName"`
	//название дела <Строковое значение>
	IssueSubject string `gorm:"column:issueSubject" json:"issueSubject"`
	//описание дела  <Строковое значение>
	IssueDescription string `gorm:"column:issueDescription" json:"issueDescription"`
	//дата создания дела <Дата>
	Ctime string `gorm:"column:ctime" json:"ctime"`
	//
	Decisions TDecisions `gorm:"column:decisions" json:"decisions"`
	//
	Assignee TAssignee `gorm:"column:assignee" json:"assignee"`

	AssigneName  string `gorm:"column:assigneName" json:"assigneName"`
	AssigneEmail string `gorm:"column:assigneEmail" json:"assigneEmail"`

	Client TClient `gorm:"column:client" json:"client"`
	//имя заявителя <Строковое значение> - устаревшее поле, сохранённое для совместимости, его надо игнорировать и использовать объект client,
	DeclarantName string `gorm:"column:declarantName" json:"declarantName"`
	//email заявителя <Строковое значение> - устаревшее поле, сохранённое для совместимости, его надо игнорировать и использовать объект client,
	DeclarantEmail string `gorm:"column:declarantEmail" json:"declarantEmail"`
	// статус дела  <Строковое значение>,
	state string `gorm:"column:state" json:"state"`
	// крайний срок решения дела <Дата в формате ‘yyyy-mm-dd’>,
	deadline string `gorm:"column:deadline" json:"deadline"`
	//ближайший норматив <Дата и время>,
	DueDate string `gorm:"column:dueDate" json:"dueDate"`
	// ближайший норматив <Дата и время>,
	claimNumber string `gorm:"column:claimNumber" json:"claimNumber"`
	// номер претензии <Строка>
	dateOfClaim string `gorm:"column:dateOfClaim" json:"dateOfClaim"`
	// дата отправки претензии <Дата>,
	DateOfReceivedClaim string `gorm:"column:dateOfReceivedClaim" json:"dateOfReceivedClaim"`

	DateOfSendClaim string `gorm:"column:dateOfSendClaim" json:"dateOfSendClaim"`

	ClaimApplicant TClaimApplicant `gorm:"column:claimApplicant" json:"claimApplicant"`

	ClaimRecipient TClaimRecipient `gorm:"column:claimRecipient" json:"claimRecipient"`

	Links TLinks `gorm:"column:links" json:"links"`

	Fields string `gorm:"column:fields" json:"fields"`
}
*/
