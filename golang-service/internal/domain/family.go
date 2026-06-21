package domain

type Family struct {
	ID         int64  `json:"fl_id"`
	CustomerID int64  `json:"cst_id"`
	Relation   string `json:"fl_relation"`
	Name       string `json:"fl_name"`
	DOB        string `json:"fl_dob"`
}
