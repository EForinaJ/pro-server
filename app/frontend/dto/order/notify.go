package dto_order

type Notify struct {
	Id           string   `p:"string"`
	CreateTime   string   `p:"create_time"`
	ResourceType string   `p:"resource_type"`
	EventType    string   `p:"event_type"`
	Summary      string   `p:"summary"`
	Resource     Resource `p:"resource"`
}

type Resource struct {
	OriginalType   string `p:"original_type"`
	Algorithm      string `p:"algorithm"`
	Ciphertext     string `p:"ciphertext"`
	AssociatedData string `p:"associated_data"`
	Nonce          string `p:"nonce"`
}
