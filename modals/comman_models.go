package modals

type CommonModal struct {
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
	DeletedBy uint `json:"deleted_by"`
	Priority  uint `json:"priority"`
	Status    bool `json:"status"`
}
