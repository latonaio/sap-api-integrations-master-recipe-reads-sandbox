package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-master-recipe-reads/SAP_API_Output_Formatter"
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

func (c *SAPAPICaller) AsyncGetMasterRecipe(masterRecipeGroup, masterRecipe string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(masterRecipeGroup, masterRecipe)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(masterRecipeGroup, masterRecipe string) {
	headerData, err := c.callMasterRecipeSrvAPIRequirementHeader("MasterRecipeHeader", masterRecipeGroup, masterRecipe)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	materialAssignmentData, err := c.callToMaterialAssignment(headerData[0].ToMaterialAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialAssignmentData)
	
	operationData, err := c.callToOperation(headerData[0].ToOperation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(operationData)
	
	phaseData, err := c.callToPhase(operationData[0].ToPhase)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(phaseData)
	
	phaseRelationshipData, err := c.callToPhaseRelationship(phaseData[0].ToPhaseRelationship)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(phaseRelationshipData)

}

func (c *SAPAPICaller) callMasterRecipeSrvAPIRequirementHeader(api, masterRecipeGroup, masterRecipe string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_MASTER_RECIPE", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, masterRecipeGroup, masterRecipe)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToMaterialAssignment(url string) ([]sap_api_output_formatter.ToMaterialAssignment, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToOperation(url string) ([]sap_api_output_formatter.ToOperation, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPhase(url string) ([]sap_api_output_formatter.ToPhase, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPhase(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPhaseRelationship(url string) ([]sap_api_output_formatter.ToPhaseRelationship, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPhaseRelationship(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, masterRecipeGroup, masterRecipe string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("MasterRecipeGroup eq '%s' and MasterRecipe eq '%s'", masterRecipeGroup, masterRecipe))
	req.URL.RawQuery = params.Encode()
}
