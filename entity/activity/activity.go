package entity

type Activity struct {
	CodeActivity Code
}

func NewActivity(c Code) *Activity {
	a := &Activity{
		CodeActivity: c,
	}

	return a
}
