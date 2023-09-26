package models

type UserModel struct {
	BaseModel
	UserName    string            `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Notes       []NoteModel       `json:"notes,omitempty" bson:"notes,omitempty"`
	Attachments []AttachmentModel `json:"attachments,omitempty" bson:"attachments,omitempty"`
}
