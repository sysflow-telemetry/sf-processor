package commons

// Configuration keys.
const (
	FindingsApiKeyConfigKey      string = "findings.apikey"
	FindingsUrlConfigKey         string = "findings.url"
	FindingsAccountIDConfigKey   string = "findings.accountid"
	FindingsProviderIDConfigKey  string = "findings.provider"
	FindingsNoteIDConfigKey      string = "findings.note"
	FindingsSqlQueryUrlConfigKey string = "findings.sqlqueryurl"
	FindingsSqlQueryCrnConfigKey string = "findings.sqlquerycrn"
	FindingsRegionConfigKey      string = "findings.region"
)

// FindingsConfig holds IBM Findings API specific configuration.
type FindingsConfig struct {
	FindingsApiKey      string
	FindingsUrl         string
	FindingsAccountID   string
	FindingsProviderID  string
	FindingsNoteID      string
	FindingsSqlQueryUrl string
	FindingsSqlQueryCrn string
	FindingsRegion      string
}

// CreateFindingsConfig creates a new config object from config dictionary.
func CreateFindingsConfig(bc Config, conf map[string]interface{}) (c FindingsConfig, err error) {
	// default values
	c = FindingsConfig{
		FindingsUrl:         "https://us-south.secadvisor.cloud.ibm.com/findings",
		FindingsSqlQueryUrl: "https://us.sql-query.cloud.ibm.com/sqlquery"}

	// parse config map
	if v, ok := conf[FindingsApiKeyConfigKey].(string); ok {
		c.FindingsApiKey = v
	} else if bc.VaultEnabled {
		s, err := bc.secrets.GetDecoded(FindingsApiKeyConfigKey)
		if err != nil {
			return c, err
		}
		c.FindingsApiKey = string(s)
	}
	if v, ok := conf[FindingsAccountIDConfigKey].(string); ok {
		c.FindingsAccountID = v
	}
	if v, ok := conf[FindingsUrlConfigKey].(string); ok {
		c.FindingsUrl = v
	}
	if v, ok := conf[FindingsProviderIDConfigKey].(string); ok {
		c.FindingsProviderID = v
	}
	if v, ok := conf[FindingsNoteIDConfigKey].(string); ok {
		c.FindingsNoteID = v
	}
	if v, ok := conf[FindingsSqlQueryUrlConfigKey].(string); ok {
		c.FindingsSqlQueryUrl = v
	}
	if v, ok := conf[FindingsSqlQueryCrnConfigKey].(string); ok {
		c.FindingsSqlQueryCrn = v
	}
	if v, ok := conf[FindingsRegionConfigKey].(string); ok {
		c.FindingsRegion = v
	}
	return
}
