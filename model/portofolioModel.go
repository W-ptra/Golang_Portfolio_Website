package model

type Portfolio struct{
	Title 		string		`json:"title"`
	Description string		`json:"description"`
	Thumbnail 	string		`json:"thumbnail"`
	Github 		string		`json:"github"`
	Youtube 	string 		`json:"youtube"`
	Skills 		[]string	`json:"skills"`
}