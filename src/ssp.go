package src

import (
	"math/rand"
	"time"
)

type SSP struct {
	ID                           int
	UUID                         string
	SystemName                   string
	Version                      string
	OtherIdentifier              string
	Confidentiality              string
	Integrity                    string
	Availability                 string
	CategorizationJustification  string
	Status                       string
	ExplanationForNonOperational string
	Description                  string
	SystemType                   string
	SystemURL                    string
	Purpose                      string
	ConditionsOfApproval         string
	Environment                  string
	LawsAndRegulations           string
	AuthorizationBoundary        string
	NetworkArchitecture          string
	DataFlow                     string
	OverallCategorization        string
	MaturityTier                 string
	WizProjectID                 string
	ServiceNowAssignmentGroup    string
	JiraProject                  string
	TenableGroup                 string
	ParentID                     int
	ParentModule                 string
	CreatedByID                  string
	DateCreated                  string
	LastUpdatedByID              string
	DateLastUpdated              string
	Users                        int
	PrivilegedUsers              int
	UsersMFA                     int
	PrivilegedUsersMFA           int
	InternalUsers                int
	ExternalUsers                int
	InternalUsersFuture          int
	ExternalUsersFuture          int
	HVA                          bool
	PracticeLevel                string
	ProcessLevel                 string
	CMMCLevel                    string
	CMMCStatus                   string
	IsPublic                     bool
	ExecutiveSummary             string
	Recommendations              string
	PrepOrgName                  string
	PrepAddress                  string
	PrepOffice                   string
	PrepCityState                string
	CSPOrgName                   string
	CSPAddress                   string
	CSPOffice                    string
	CSPCityState                 string
	BModelSaaS                   bool
	BModelPaaS                   bool
	BModelIaaS                   bool
	BModelOther                  bool
	OtherModelRemarks            string
	BDeployPublic                bool
	BDeployPrivate               bool
	BDeployGov                   bool
	BDeployHybrid                bool
	BDeployOther                 bool
	DeployOtherRemarks           string
	AuthenticationLevel          string
	IdentityAssuranceLevel       string
	AuthenticatorAssuranceLevel  string
	FederationAssuranceLevel     string
	FedrampAuthorizationType     string
	FedrampAuthorizationLevel    string
	FedrampAuthorizationStatus   string
	FedrampID                    string
	DefaultAssessmentDays        int
}

func NewSSP(systemName string) *SSP {
	rand.NewSource(time.Now().UnixNano())
	return &SSP{
		SystemName:            systemName,
		Status:                []string{"Operational", "Under Development", "Undergoing a Major Modification", "In Transition", "Other"}[rand.Intn(5)],
		SystemType:            []string{"Major Application", "General Support System"}[rand.Intn(2)],
		DefaultAssessmentDays: rand.Intn(365),
	}
}
