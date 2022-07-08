package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-activity-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetActivity(iD, subject, category string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "AppointmentCollection":
			func() {
				c.AppointmentCollection(iD)
				wg.Done()
			}()
		case "EmailCollection":
			func() {
				c.EmailCollection(subject)
				wg.Done()
			}()
		case "TaskCollection":
			func() {
				c.TaskCollection(category)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) AppointmentCollection(iD string) {
	appointmentCollectionData, err := c.callActivitySrvAPIRequirementAppointmentCollection("AppointmentCollection", iD)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(appointmentCollectionData)

	appointmentInvolvedPartiesData, err := c.callAppointmentInvolvedParties(appointmentCollectionData[0].ToAppointmentInvolvedParties)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(appointmentInvolvedPartiesData)

	appointmentOtherPartiesData, err := c.callAppointmentOtherParties(appointmentCollectionData[0].ToAppointmentOtherParties)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(appointmentOtherPartiesData)

}

func (c *SAPAPICaller) callActivitySrvAPIRequirementAppointmentCollection(api, iD string) ([]sap_api_output_formatter.AppointmentCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithAppointmentCollection(req, iD)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAppointmentCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callAppointmentInvolvedParties(url string) ([]sap_api_output_formatter.AppointmentInvolvedParties, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAppointmentInvolvedParties(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callAppointmentOtherParties(url string) ([]sap_api_output_formatter.AppointmentOtherParties, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAppointmentOtherParties(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) EmailCollection(subject string) {
	emailCollectionData, err := c.callActivitySrvAPIRequirementEmailCollection("EMailCollection", subject)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(emailCollectionData)

	emailRecipientsData, err := c.callEmailRecipients(emailCollectionData[0].ToEMailRecipients)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(emailRecipientsData)

	emailSenderPartyData, err := c.callEmailSenderParty(emailCollectionData[0].ToEMailSenderParty)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(emailSenderPartyData)
}

func (c *SAPAPICaller) callActivitySrvAPIRequirementEmailCollection(api, subject string) ([]sap_api_output_formatter.EmailCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEmailCollection(req, subject)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmailCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEmailRecipients(url string) ([]sap_api_output_formatter.EmailRecipients, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmailRecipients(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEmailSenderParty(url string) ([]sap_api_output_formatter.EmailSenderParty, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmailSenderParty(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) TaskCollection(category string) {
	data, err := c.callActivitySrvAPIRequirementTaskCollection("TasksCollection", category)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callActivitySrvAPIRequirementTaskCollection(api, category string) ([]sap_api_output_formatter.TaskCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithTaskCollection(req, category)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToTaskCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}
func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithAppointmentCollection(req *http.Request, iD string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s'", iD))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithEmailCollection(req *http.Request, subject string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("substringof('%s', Subject)", subject))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithTaskCollection(req *http.Request, category string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Category eq '%s'", category))
	req.URL.RawQuery = params.Encode()
}
