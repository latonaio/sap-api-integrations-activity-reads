package responses

type TaskCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                       string `json:"ObjectID"`
			DocumentType                   string `json:"DocumentType"`
			DocumentTypeText               string `json:"DocumentTypeText"`
			ID                             string `json:"ID"`
			Subject                        string `json:"Subject"`
			Status                         string `json:"Status"`
			StatusText                     string `json:"StatusText"`
			StartDateTime                  string `json:"StartDateTime"`
			StartDatetimeZoneCode          string `json:"StartDatetimeZoneCode"`
			StartDatetimeZoneCodeText      string `json:"StartDatetimeZoneCodeText"`
			ActivityList                   string `json:"ActivityList"`
			DueDateTime                    string `json:"DueDateTime"`
			DueDatetimeZoneCode            string `json:"DueDatetimeZoneCode"`
			DueDatetimeZoneCodeText        string `json:"DueDatetimeZoneCodeText"`
			PlannedDuration                string `json:"PlannedDuration"`
			ActualDuration                 string `json:"ActualDuration"`
			CompletionDateTime             string `json:"CompletionDateTime"`
			MainEmployeeResponsiblePartyID string `json:"MainEmployeeResponsiblePartyID"`
			MainAccountPartyID             string `json:"MainAccountPartyID"`
			MainContactPartyID             string `json:"MainContactPartyID"`
			ProcessorPartyID               string `json:"ProcessorPartyID"`
			CompletionPercent              string `json:"CompletionPercent"`
			Category                       string `json:"Category"`
			CategoryText                   string `json:"CategoryText"`
			PriorityCode                   string `json:"PriorityCode"`
			PriorityCodeText               string `json:"PriorityCodeText"`
			SalesTerritoryID               string `json:"SalesTerritoryID"`
			SalesTerritoryName             string `json:"SalesTerritoryName"`
			SalesOrganisationID            string `json:"SalesOrganisationID"`
			DistributionChannelCode        string `json:"DistributionChannelCode"`
			DistributionChannelCodeText    string `json:"DistributionChannelCodeText"`
			DivisionCode                   string `json:"DivisionCode"`
			DivisionCodeText               string `json:"DivisionCodeText"`
			Account                        string `json:"Account"`
			AccountUUID                    string `json:"AccountUUID"`
			DataOriginTypeCode             string `json:"DataOriginTypeCode"`
			DataOriginTypeCodeText         string `json:"DataOriginTypeCodeText"`
			Owner                          string `json:"Owner"`
			OwnerUUID                      string `json:"OwnerUUID"`
			PrimaryContact                 string `json:"PrimaryContact"`
			PrimaryContactUUID             string `json:"PrimaryContactUUID"`
			Processor                      string `json:"Processor"`
			ProcessorEmailURI              string `json:"ProcessorEmailURI"`
			ProcessorPartyName             string `json:"ProcessorPartyName"`
			ProcessorUUID                  string `json:"ProcessorUUID"`
			ReportedDateTime               string `json:"ReportedDateTime"`
			UUID                           string `json:"UUID"`
			ChangedBy                      string `json:"ChangedBy"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
			EntityLastChangedOn            string `json:"EntityLastChangedOn"`
			ETag                           string `json:"ETag"`
			CreatedBySSI                   string `json:"CreatedBySSI"`
			CreatedBySSIText               string `json:"CreatedBySSIText"`
			Dataloader1KUT                 string `json:"dataloader1_KUT"`
		} `json:"results"`
	} `json:"d"`
}
