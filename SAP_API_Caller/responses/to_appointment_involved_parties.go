package responses

type ToAppointmentInvolvedParties struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID             string `json:"ObjectID"`
			ParentObjectID       string `json:"ParentObjectID"`
			ID                   string `json:"ID"`
			PartyID              string `json:"PartyID"`
			RoleCode             string `json:"RoleCode"`
			RoleCodeText         string `json:"RoleCodeText"`
			RoleCategoryCode     string `json:"RoleCategoryCode"`
			RoleCategoryCodeText string `json:"RoleCategoryCodeText"`
			PartyTypeCode        string `json:"PartyTypeCode"`
			PartyTypeCodeText    string `json:"PartyTypeCodeText"`
			MainIndicator        bool   `json:"MainIndicator"`
			Address              string `json:"Address"`
			EMail                string `json:"EMail"`
			Name                 string `json:"Name"`
			PartyEmailURI        string `json:"PartyEmailURI"`
			PartyUUID            string `json:"PartyUUID"`
			PartyName            string `json:"PartyName"`
			Phone                string `json:"Phone"`
			ETag                 string `json:"ETag"`
			StartDate            string `json:"StartDate"`
		} `json:"results"`
	} `json:"d"`
}
