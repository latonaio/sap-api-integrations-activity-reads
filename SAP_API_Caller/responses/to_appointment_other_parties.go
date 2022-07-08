package responses

type ToAppointmentOtherParties struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID             string `json:"ObjectID"`
			ParentObjectID       string `json:"ParentObjectID"`
			AppointmentID        string `json:"AppointmentID"`
			PartyUUID            string `json:"PartyUUID"`
			PartyID              string `json:"PartyID"`
			RoleCategoryCode     string `json:"RoleCategoryCode"`
			RoleCategoryCodeText string `json:"RoleCategoryCodeText"`
			RoleCode             string `json:"RoleCode"`
			RoleCodeText         string `json:"RoleCodeText"`
			ETag                 string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
