package responses

type EmailCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                   string `json:"ObjectID"`
			ID                         string `json:"ID"`
			Subject                    string `json:"Subject"`
			Status                     string `json:"Status"`
			StatusText                 string `json:"StatusText"`
			Category                   string `json:"Category"`
			CategoryText               string `json:"CategoryText"`
			PriorityCode               string `json:"PriorityCode"`
			PriorityCodeText           string `json:"PriorityCodeText"`
			DateTime                   string `json:"DateTime"`
			DateTimetimeZoneCode       string `json:"DateTimetimeZoneCode"`
			DateTimetimeZoneCodeText   string `json:"DateTimetimeZoneCodeText"`
			TransmissionStatusCode     string `json:"TransmissionStatusCode"`
			TransmissionStatusCodeText string `json:"TransmissionStatusCodeText"`
			OwnerPartyID               string `json:"OwnerPartyID"`
			Account                    string `json:"Account"`
			MainAccountPartyID         string `json:"MainAccountPartyID"`
			PrimaryContact             string `json:"PrimaryContact"`
			MainContactPartyID         string `json:"MainContactPartyID"`
			Direction                  string `json:"Direction"`
			DirectionText              string `json:"DirectionText"`
			SalesTerritoryID           string `json:"SalesTerritoryID"`
			SalesTerritoryName         string `json:"SalesTerritoryName"`
			Campaign                   string `json:"Campaign"`
			ResponseOption             string `json:"ResponseOption"`
			SentimentTypeCode          string `json:"SentimentTypeCode"`
			SentimentTypeCodeText      string `json:"SentimentTypeCodeText"`
			DocumentType               string `json:"DocumentType"`
			DocumentTypeText           string `json:"DocumentTypeText"`
			TicketID                   string `json:"TicketID"`
			EntityLastChangedOn        string `json:"EntityLastChangedOn"`
			ETag                       string `json:"ETag"`
			DataOriginTypeCode         string `json:"DataOriginTypeCode"`
			DataOriginTypeCodeText     string `json:"DataOriginTypeCodeText"`
			AccountName                string `json:"AccountName"`
			GroupwareItemID            string `json:"GroupwareItemID"`
			MessageFromEmailURI        string `json:"MessageFromEmailURI"`
			MessageFromName            string `json:"MessageFromName"`
			MessageFromPartyID         string `json:"MessageFromPartyID"`
			MessageFromPartyName       string `json:"MessageFromPartyName"`
			MessageFromPartyUUID       string `json:"MessageFromPartyUUID"`
			OwnerName                  string `json:"OwnerName"`
			OwnerUUID                  string `json:"OwnerUUID"`
			PrimaryContactName         string `json:"PrimaryContactName"`
			UUID                       string `json:"UUID"`
			LastChangeDateTime         string `json:"LastChangeDateTime"`
			CreatedBySSI               string `json:"CreatedBySSI"`
			CreatedBySSIText           string `json:"CreatedBySSIText"`
			Dataloader1KUT             string `json:"dataloader1_KUT"`
			EMailSenderParty           struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"EMailSenderParty"`
			EMailToRecipients struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"EMailToRecipients"`
		} `json:"results"`
	} `json:"d"`
}
