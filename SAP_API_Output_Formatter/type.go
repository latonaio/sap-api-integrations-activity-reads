package sap_api_output_formatter

type AppointmentCollection struct {
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
	ToAppointmentInvolvedParties    string `json:"AppointmentInvolvedParties"`
	ToAppointmentOtherParties       string `json:"AppointmentOtherParties"`
}

type EmailCollection struct {
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
	Dataloader1KUT             string `json:"Dataloader1KUT"`
	ToEMailSenderParty         string `json:"EMailSenderParty"`
	ToEMailRecipients          string `json:"EMailToRecipients"`
}

type AppointmentInvolvedParties struct {
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
}

type AppointmentOtherParties struct {
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
}

type EmailRecipients struct {
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
}

type EmailSenderParty struct {
	ObjectID       string `json:"ObjectID"`
	ParentObjectID string `json:"ParentObjectID"`
	EMailID        string `json:"EMailID"`
	PartyID        string `json:"PartyID"`
	PartyName      string `json:"PartyName"`
	EMailValue     string `json:"EMailValue"`
	ETag           string `json:"ETag"`
}

type TaskCollection struct {
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
}
