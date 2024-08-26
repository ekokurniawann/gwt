package params


type EmployeeSingleView struct {
	ID        int    `json:"id"`
	NIP       string `json:"nip"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
