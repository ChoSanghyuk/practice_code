package requestbody

type User struct {
	Name string `json:"이름" validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
	Age  int    `json:"나이" validate:"required,teener"`       // Required field, and client needs to implement our 'teener' tag format which we'll see later
}
