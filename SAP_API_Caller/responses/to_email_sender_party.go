package responses

type ToEmailSenderParty struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID       string `json:"ObjectID"`
			ParentObjectID string `json:"ParentObjectID"`
			EMailID        string `json:"EMailID"`
			PartyID        string `json:"PartyID"`
			PartyName      string `json:"PartyName"`
			EMailValue     string `json:"EMailValue"`
			ETag           string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
