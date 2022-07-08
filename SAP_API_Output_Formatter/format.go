package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-activity-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToAppointmentCollection(raw []byte, l *logger.Logger) ([]AppointmentCollection, error) {
	pm := &responses.AppointmentCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to AppointmentCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	appointmentCollection := make([]AppointmentCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		appointmentCollection = append(appointmentCollection, AppointmentCollection{
			ObjectID:                        data.ObjectID,
			DocumentType:                    data.DocumentType,
			DocumentTypeText:                data.DocumentTypeText,
			ID:                              data.ID,
			Subject:                         data.Subject,
			LifeCycleStatusCode:             data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:         data.LifeCycleStatusCodeText,
			StartDate:                       data.StartDate,
			StartDatetimeZoneCode:           data.StartDatetimeZoneCode,
			StartDatetimeZoneCodeText:       data.StartDatetimeZoneCodeText,
			EndDate:                         data.EndDate,
			EnddatetimeZoneCode:             data.EnddatetimeZoneCode,
			EnddatetimeZoneCodeText:         data.EnddatetimeZoneCodeText,
			OrganizerUUID:                   data.OrganizerUUID,
			OrganizerPartyID:                data.OrganizerPartyID,
			Owner:                           data.Owner,
			PrimaryContact:                  data.PrimaryContact,
			OwnerPartyID:                    data.OwnerPartyID,
			FullDayIndicator:                data.FullDayIndicator,
			Category:                        data.Category,
			CategoryText:                    data.CategoryText,
			Account:                         data.Account,
			MainAccountPartyID:              data.MainAccountPartyID,
			MainContactPartyID:              data.MainContactPartyID,
			Location:                        data.Location,
			Priority:                        data.Priority,
			PriorityText:                    data.PriorityText,
			SalesTerritoryUUID:              data.SalesTerritoryUUID,
			SalesTerritoryID:                data.SalesTerritoryID,
			SalesTerritoryName:              data.SalesTerritoryName,
			SalesOrganisation:               data.SalesOrganisation,
			DistributionChannel:             data.DistributionChannel,
			DistributionChannelText:         data.DistributionChannelText,
			Division:                        data.Division,
			DivisionText:                    data.DivisionText,
			VisitTourPlanUUID:               data.VisitTourPlanUUID,
			InformationSensitivityCode:      data.InformationSensitivityCode,
			InformationSensitivityCodeText:  data.InformationSensitivityCodeText,
			AccountName:                     data.AccountName,
			AdditionalLocationName:          data.AdditionalLocationName,
			ChangedBy:                       data.ChangedBy,
			CompletionPercent:               data.CompletionPercent,
			CreatedBy:                       data.CreatedBy,
			CreatedOn:                       data.CreatedOn,
			DataOriginTypeCode:              data.DataOriginTypeCode,
			DataOriginTypeCodeText:          data.DataOriginTypeCodeText,
			GroupwareItemID:                 data.GroupwareItemID,
			OrganizerEmailURI:               data.OrganizerEmailURI,
			OrganizerName:                   data.OrganizerName,
			OrganizerPartyName:              data.OrganizerPartyName,
			OwnerName:                       data.OwnerName,
			Phone:                           data.Phone,
			PrimaryContactName:              data.PrimaryContactName,
			ReportedDateTime:                data.ReportedDateTime,
			UUID:                            data.UUID,
			AvailabilityCode:                data.AvailabilityCode,
			AvailabilityCodeText:            data.AvailabilityCodeText,
			LastChangeDateTime:              data.LastChangeDateTime,
			EntityLastChangedOn:             data.EntityLastChangedOn,
			ETag:                            data.ETag,
			VisitTypeCode:                   data.VisitTypeCode,
			VisitTypeCodeText:               data.VisitTypeCodeText,
			OccurrenceDate:                  data.OccurrenceDate,
			RecurrenceActGroupwareID:        data.RecurrenceActGroupwareID,
			RecurringBusinessActivityNodeID: data.RecurringBusinessActivityNodeID,
			CollaborationChannelCode:        data.CollaborationChannelCode,
			CollaborationChannelCodeText:    data.CollaborationChannelCodeText,
			CollaborationURL:                data.CollaborationURL,
			Draft:                           data.Draft,
			CreatedBySSI:                    data.CreatedBySSI,
			CreatedBySSIText:                data.CreatedBySSIText,
			Dataloader1KUT:                  data.Dataloader1KUT,
			ToAppointmentInvolvedParties:    data.AppointmentInvolvedParties.Deferred.URI,
			ToAppointmentOtherParties:       data.AppointmentOtherParties.Deferred.URI,
		})
	}

	return appointmentCollection, nil
}

func ConvertToEmailCollection(raw []byte, l *logger.Logger) ([]EmailCollection, error) {
	pm := &responses.EmailCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to EmailCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	emailCollection := make([]EmailCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		emailCollection = append(emailCollection, EmailCollection{
			ObjectID:                   data.ObjectID,
			ID:                         data.ID,
			Subject:                    data.Subject,
			Status:                     data.Status,
			StatusText:                 data.StatusText,
			Category:                   data.Category,
			CategoryText:               data.CategoryText,
			PriorityCode:               data.PriorityCode,
			PriorityCodeText:           data.PriorityCodeText,
			DateTime:                   data.DateTime,
			DateTimetimeZoneCode:       data.DateTimetimeZoneCode,
			DateTimetimeZoneCodeText:   data.DateTimetimeZoneCodeText,
			TransmissionStatusCode:     data.TransmissionStatusCode,
			TransmissionStatusCodeText: data.TransmissionStatusCodeText,
			OwnerPartyID:               data.OwnerPartyID,
			Account:                    data.Account,
			MainAccountPartyID:         data.MainAccountPartyID,
			PrimaryContact:             data.PrimaryContact,
			MainContactPartyID:         data.MainContactPartyID,
			Direction:                  data.Direction,
			DirectionText:              data.DirectionText,
			SalesTerritoryID:           data.SalesTerritoryID,
			SalesTerritoryName:         data.SalesTerritoryName,
			Campaign:                   data.Campaign,
			ResponseOption:             data.ResponseOption,
			SentimentTypeCode:          data.SentimentTypeCode,
			SentimentTypeCodeText:      data.SentimentTypeCodeText,
			DocumentType:               data.DocumentType,
			DocumentTypeText:           data.DocumentTypeText,
			TicketID:                   data.TicketID,
			EntityLastChangedOn:        data.EntityLastChangedOn,
			ETag:                       data.ETag,
			DataOriginTypeCode:         data.DataOriginTypeCode,
			DataOriginTypeCodeText:     data.DataOriginTypeCodeText,
			AccountName:                data.AccountName,
			GroupwareItemID:            data.GroupwareItemID,
			MessageFromEmailURI:        data.MessageFromEmailURI,
			MessageFromName:            data.MessageFromName,
			MessageFromPartyID:         data.MessageFromPartyID,
			MessageFromPartyName:       data.MessageFromPartyName,
			MessageFromPartyUUID:       data.MessageFromPartyUUID,
			OwnerName:                  data.OwnerName,
			OwnerUUID:                  data.OwnerUUID,
			PrimaryContactName:         data.PrimaryContactName,
			UUID:                       data.UUID,
			LastChangeDateTime:         data.LastChangeDateTime,
			CreatedBySSI:               data.CreatedBySSI,
			CreatedBySSIText:           data.CreatedBySSIText,
			Dataloader1KUT:             data.Dataloader1KUT,
			ToEMailSenderParty:         data.EMailSenderParty.Deferred.URI,
			ToEMailRecipients:          data.EMailToRecipients.Deferred.URI,
		})
	}
	return emailCollection, nil
}

func ConvertToAppointmentInvolvedParties(raw []byte, l *logger.Logger) ([]AppointmentInvolvedParties, error) {
	pm := &responses.ToAppointmentInvolvedParties{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToAppointmentInvolvedParties. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	appointmentInvolvedParties := make([]AppointmentInvolvedParties, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		appointmentInvolvedParties = append(appointmentInvolvedParties, AppointmentInvolvedParties{
			ObjectID:             data.ObjectID,
			ParentObjectID:       data.ParentObjectID,
			ID:                   data.ID,
			PartyID:              data.PartyID,
			RoleCode:             data.RoleCode,
			RoleCodeText:         data.RoleCodeText,
			RoleCategoryCode:     data.RoleCategoryCode,
			RoleCategoryCodeText: data.RoleCategoryCodeText,
			PartyTypeCode:        data.PartyTypeCode,
			PartyTypeCodeText:    data.PartyTypeCodeText,
			MainIndicator:        data.MainIndicator,
			Address:              data.Address,
			EMail:                data.EMail,
			Name:                 data.Name,
			PartyEmailURI:        data.PartyEmailURI,
			PartyUUID:            data.PartyUUID,
			PartyName:            data.PartyName,
			Phone:                data.Phone,
			ETag:                 data.ETag,
			StartDate:            data.StartDate,
		})
	}

	return appointmentInvolvedParties, nil
}

func ConvertToAppointmentOtherParties(raw []byte, l *logger.Logger) ([]AppointmentOtherParties, error) {
	pm := &responses.ToAppointmentOtherParties{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to AppointmentOtherParties. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	appointmentOtherParties := make([]AppointmentOtherParties, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		appointmentOtherParties = append(appointmentOtherParties, AppointmentOtherParties{
			ObjectID:             data.ObjectID,
			ParentObjectID:       data.ParentObjectID,
			AppointmentID:        data.AppointmentID,
			PartyUUID:            data.PartyUUID,
			PartyID:              data.PartyID,
			RoleCategoryCode:     data.RoleCategoryCode,
			RoleCategoryCodeText: data.RoleCategoryCodeText,
			RoleCode:             data.RoleCode,
			RoleCodeText:         data.RoleCodeText,
			ETag:                 data.ETag,
		})
	}

	return appointmentOtherParties, nil
}

func ConvertToEmailRecipients(raw []byte, l *logger.Logger) ([]EmailRecipients, error) {
	pm := &responses.ToEmailRecipients{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to EmailRecipients. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	emailRecipients := make([]EmailRecipients, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		emailRecipients = append(emailRecipients, EmailRecipients{
			ObjectID:             data.ObjectID,
			ParentObjectID:       data.ParentObjectID,
			EMailParentID:        data.EMailParentID,
			PartyID:              data.PartyID,
			Name:                 data.Name,
			NameLanguageCode:     data.NameLanguageCode,
			NameLanguageCodeText: data.NameLanguageCodeText,
			EMailID:              data.EMailID,
			Phone:                data.Phone,
			Address:              data.Address,
			MessageToName:        data.MessageToName,
			MessageToEmailID:     data.MessageToEmailID,
			PartyUUID:            data.PartyUUID,
			ETag:                 data.ETag,
		})
	}

	return emailRecipients, nil
}

func ConvertToEmailSenderParty(raw []byte, l *logger.Logger) ([]EmailSenderParty, error) {
	pm := &responses.ToEmailSenderParty{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to EmailSenderParty. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	emailSenderParty := make([]EmailSenderParty, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		emailSenderParty = append(emailSenderParty, EmailSenderParty{
			ObjectID:       data.ObjectID,
			ParentObjectID: data.ParentObjectID,
			EMailID:        data.EMailID,
			PartyID:        data.PartyID,
			PartyName:      data.PartyName,
			EMailValue:     data.EMailValue,
			ETag:           data.ETag,
		})
	}

	return emailSenderParty, nil
}

func ConvertToTaskCollection(raw []byte, l *logger.Logger) ([]TaskCollection, error) {
	pm := &responses.TaskCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to TaskCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	taskCollection := make([]TaskCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		taskCollection = append(taskCollection, TaskCollection{
			ObjectID:                       data.ObjectID,
			DocumentType:                   data.DocumentType,
			DocumentTypeText:               data.DocumentTypeText,
			ID:                             data.ID,
			Subject:                        data.Subject,
			Status:                         data.Status,
			StatusText:                     data.StatusText,
			StartDateTime:                  data.StartDateTime,
			StartDatetimeZoneCode:          data.StartDatetimeZoneCode,
			StartDatetimeZoneCodeText:      data.StartDatetimeZoneCodeText,
			ActivityList:                   data.ActivityList,
			DueDateTime:                    data.DueDateTime,
			DueDatetimeZoneCode:            data.DueDatetimeZoneCode,
			DueDatetimeZoneCodeText:        data.DueDatetimeZoneCodeText,
			PlannedDuration:                data.PlannedDuration,
			ActualDuration:                 data.ActualDuration,
			CompletionDateTime:             data.CompletionDateTime,
			MainEmployeeResponsiblePartyID: data.MainEmployeeResponsiblePartyID,
			MainAccountPartyID:             data.MainAccountPartyID,
			MainContactPartyID:             data.MainContactPartyID,
			ProcessorPartyID:               data.ProcessorPartyID,
			CompletionPercent:              data.CompletionPercent,
			Category:                       data.Category,
			CategoryText:                   data.CategoryText,
			PriorityCode:                   data.PriorityCode,
			PriorityCodeText:               data.PriorityCodeText,
			SalesTerritoryID:               data.SalesTerritoryID,
			SalesTerritoryName:             data.SalesTerritoryName,
			SalesOrganisationID:            data.SalesOrganisationID,
			DistributionChannelCode:        data.DistributionChannelCode,
			DistributionChannelCodeText:    data.DistributionChannelCodeText,
			DivisionCode:                   data.DivisionCode,
			DivisionCodeText:               data.DivisionCodeText,
			Account:                        data.Account,
			AccountUUID:                    data.AccountUUID,
			DataOriginTypeCode:             data.DataOriginTypeCode,
			DataOriginTypeCodeText:         data.DataOriginTypeCodeText,
			Owner:                          data.Owner,
			OwnerUUID:                      data.OwnerUUID,
			PrimaryContact:                 data.PrimaryContact,
			PrimaryContactUUID:             data.PrimaryContactUUID,
			Processor:                      data.Processor,
			ProcessorEmailURI:              data.ProcessorEmailURI,
			ProcessorPartyName:             data.ProcessorPartyName,
			ProcessorUUID:                  data.ProcessorUUID,
			ReportedDateTime:               data.ReportedDateTime,
			UUID:                           data.UUID,
			ChangedBy:                      data.ChangedBy,
			LastChangeDateTime:             data.LastChangeDateTime,
			EntityLastChangedOn:            data.EntityLastChangedOn,
			ETag:                           data.ETag,
			CreatedBySSI:                   data.CreatedBySSI,
			CreatedBySSIText:               data.CreatedBySSIText,
			Dataloader1KUT:                 data.Dataloader1KUT,
		})
	}
	return taskCollection, nil
}
