package dto_req

type (
	MosqAuthBody struct {
		Username string
		Password string
		Topic    string
		Acc      int8
	}
)
