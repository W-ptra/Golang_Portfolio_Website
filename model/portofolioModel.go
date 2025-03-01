package model

type Portfolio struct{
	Title 		string		`json:"title"`
	Description string		`json:"description"`
	Thumbnail 	string		`json:"thumbnail"`
	Github 		string		`json:"github"`
	Youtube 	string 		`json:"youtube"`
	Demo 		string 		`json:"demo"`
	Winner		string 		`json:"winner"`
	WinnerLink  string		`json:"winner_link"`
	Skills 		[]string	`json:"skills"`
}