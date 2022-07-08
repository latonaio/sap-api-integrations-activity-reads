package responses

type AppointmentCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                        string `json:"ObjectID"`
			DocumentType                    string `json:"DocumentType"`
			DocumentTypeText                string `json:"DocumentTypeText"`
			ID                              string `json:"ID"`
			Subject                         string `json:"Subject"`
			LifeCycleStatusCode             string `json:"LifeCycleStatusCode"`
			LifeCycleStatusCodeText         string `json:"LifeCycleStatusCodeText"`
			StartDate                       string `json:"StartDate"`
			StartDatetimeZoneCode           string `json:"StartDatetimeZoneCode"`
			StartDatetimeZoneCodeText       string `json:"StartDatetimeZoneCodeText"`
			EndDate                         string `json:"EndDate"`
			EnddatetimeZoneCode             string `json:"EnddatetimeZoneCode"`
			EnddatetimeZoneCodeText         string `json:"EnddatetimeZoneCodeText"`
			OrganizerUUID                   string `json:"OrganizerUUID"`
			OrganizerPartyID                string `json:"OrganizerPartyID"`
			Owner                           string `json:"Owner"`
			PrimaryContact                  string `json:"PrimaryContact"`
			OwnerPartyID                    string `json:"OwnerPartyID"`
			FullDayIndicator                bool   `json:"FullDayIndicator"`
			Category                        string `json:"Category"`
			CategoryText                    string `json:"CategoryText"`
			Account                         string `json:"Account"`
			MainAccountPartyID              string `json:"MainAccountPartyID"`
			MainContactPartyID              string `json:"MainContactPartyID"`
			Location                        string `json:"Location"`
			Priority                        string `json:"Priority"`
			PriorityText                    string `json:"PriorityText"`
			SalesTerritoryUUID              string `json:"SalesTerritoryUUID"`
			SalesTerritoryID                string `json:"SalesTerritoryID"`
			SalesTerritoryName              string `json:"SalesTerritoryName"`
			SalesOrganisation               string `json:"SalesOrganisation"`
			DistributionChannel             string `json:"DistributionChannel"`
			DistributionChannelText         string `json:"DistributionChannelText"`
			Division                        string `json:"Division"`
			DivisionText                    string `json:"DivisionText"`
			VisitTourPlanUUID               string `json:"VisitTourPlanUUID"`
			InformationSensitivityCode      string `json:"InformationSensitivityCode"`
			InformationSensitivityCodeText  string `json:"InformationSensitivityCodeText"`
			AccountName                     string `json:"AccountName"`
			AdditionalLocationName          string `json:"AdditionalLocationName"`
			ChangedBy                       string `json:"ChangedBy"`
			CompletionPercent               string `json:"CompletionPercent"`
			CreatedBy                       string `json:"CreatedBy"`
			CreatedOn                       string `json:"CreatedOn"`
			DataOriginTypeCode              string `json:"DataOriginTypeCode"`
			DataOriginTypeCodeText          string `json:"DataOriginTypeCodeText"`
			GroupwareItemID                 string `json:"GroupwareItemID"`
			OrganizerEmailURI               string `json:"OrganizerEmailURI"`
			OrganizerName                   string `json:"OrganizerName"`
			OrganizerPartyName              string `json:"OrganizerPartyName"`
			OwnerName                       string `json:"OwnerName"`
			Phone                           string `json:"Phone"`
			PrimaryContactName              string `json:"PrimaryContactName"`
			ReportedDateTime                string `json:"ReportedDateTime"`
			UUID                            string `json:"UUID"`
			AvailabilityCode                string `json:"AvailabilityCode"`
			AvailabilityCodeText            string `json:"AvailabilityCodeText"`
			LastChangeDateTime              string `json:"LastChangeDateTime"`
			EntityLastChangedOn             string `json:"EntityLastChangedOn"`
			ETag                            string `json:"ETag"`
			VisitTypeCode                   string `json:"VisitTypeCode"`
			VisitTypeCodeText               string `json:"VisitTypeCodeText"`
			OccurrenceDate                  string `json:"OccurrenceDate"`
			RecurrenceActGroupwareID        string `json:"RecurrenceActGroupwareID"`
			RecurringBusinessActivityNodeID string `json:"RecurringBusinessActivityNodeID"`
			CollaborationChannelCode        string `json:"CollaborationChannelCode"`
			CollaborationChannelCodeText    string `json:"CollaborationChannelCodeText"`
			CollaborationURL                string `json:"CollaborationURL"`
			Draft                           bool   `json:"Draft"`
			CreatedBySSI                    string `json:"CreatedBySSI"`
			CreatedBySSIText                string `json:"CreatedBySSIText"`
			Dataloader1KUT                  string `json:"dataloader1_KUT"`
			AppointmentInvolvedParties      struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"AppointmentInvolvedParties"`
			AppointmentOtherParties struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"AppointmentOtherParties"`
		} `json:"results"`
	} `json:"d"`
}
