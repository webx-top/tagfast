package c

type Coscms struct {
	Url   string `xorm:"not null default '' VARCHAR(255)" valid:"Requied" form_widget:"text"`
	Email string `xorm:"not null default '' VARCHAR(100)" valid:"Max(100)" form_widget:"radio"`
}
