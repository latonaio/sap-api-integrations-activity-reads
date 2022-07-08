package responses

type ToEmailRecipients struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID             string `json:"ObjectID"`
			ParentObjectID       string `json:"ParentObjectID"`
			EMailParentID        string `json:"EMailParentID"`
			PartyID              string `json:"PartyID"`
			Name                 string `json:"Name"`
			NameLanguageCode     string `json:"NameLanguageCode"`
			NameLanguageCodeText string `json:"NameLanguageCodeText"`
			EMailID              string `json:"EMailID"`
			Phone                string `json:"Phone"`
			Address              string `json:"Address"`
			MessageToName        string `json:"MessageToName"`
			MessageToEmailID     string `json:"MessageToEmailID"`
			PartyUUID            string `json:"PartyUUID"`
			ETag                 string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
