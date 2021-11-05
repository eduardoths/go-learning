package migrations

type Users struct {
	Base
	Name             string
	Email            string
	PasswordHash     string
	SentMessages     []Messages `gorm:"foreignKey:SenderID"`
	ReceivedMessages []Messages `gorm:"foreignKey:ReceiverID"`
}
