package sonarqube

import (
	"context"
	"net/http"
	"testing"
)

func basicAuth(ctx context.Context, req *http.Request) error {
	req.SetBasicAuth("admin", "admin")
	return nil
}

func client() (*ClientWithResponses, error) {
	hc := http.Client{}

	return NewClientWithResponses("http://127.0.0.1:9000", WithHTTPClient(&hc))
}

func newClient(t *testing.T) *ClientWithResponses {
	t.Skip()

	c, err := client()
	if err != nil {
		t.Errorf("%v", err)
	}

	return c
}

func assertHTTPStatus(t *testing.T, resp *http.Response, err error, msg any) {
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		t.Errorf("HTTP Status: %v %v", resp.StatusCode, msg)
	}
}

func assertResponseBody(t *testing.T, resp any) {
	if resp == nil {
		t.Errorf("Error: No Content")
	}
}

func TestMain(m *testing.M) {
	//c, err := client()
	//if err != nil {
	//	panic(err)
	//}

	//cparams := ApiProjectsCreateParams{}
	//cparams.Name = "Test Project"
	//cparams.Project = "test_project"
	//_, err = c.ApiProjectsCreateWithResponse(context.TODO(), &cparams, basicAuth)
	//if err != nil {
	//	panic(err)
	//}

	m.Run()

	//dparams := ApiProjectsDeleteParams{}
	//dparams.Project = cparams.Project
	//_, err = c.ApiProjectsDeleteWithResponse(context.TODO(), &dparams, basicAuth)
	//if err != nil {
	//	panic(err)
	//}
}

func TestApiAlmIntegrationsImportGitlabProjectWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmIntegrationsImportGitlabProjectParams{}
	params.AlmSetting = "settings_key_gitlab"
	params.GitlabProjectId = "gitlab_project_id"
	resp, err := c.ApiAlmIntegrationsImportGitlabProjectWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmIntegrationsListAzureProjectsWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmIntegrationsListAzureProjectsParams{}
	params.AlmSetting = "settings_key_azure"
	resp, err := c.ApiAlmIntegrationsListAzureProjectsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmIntegrationsListBitbucketserverProjectsWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmIntegrationsListBitbucketserverProjectsParams{}
	params.AlmSetting = "settings_key_bitbucket"
	resp, err := c.ApiAlmIntegrationsListBitbucketserverProjectsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmIntegrationsSearchAzureReposWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmIntegrationsSearchAzureReposParams{}
	params.AlmSetting = "settings_key_azure"
	resp, err := c.ApiAlmIntegrationsSearchAzureReposWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmIntegrationsSearchBitbucketcloudReposWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmIntegrationsSearchBitbucketcloudReposParams{}
	params.AlmSetting = "settings_key_bitbucketcloud"
	resp, err := c.ApiAlmIntegrationsSearchBitbucketcloudReposWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmIntegrationsSearchBitbucketserverReposWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmIntegrationsSearchBitbucketserverReposParams{}
	params.AlmSetting = "settings_key_bitbucket"
	resp, err := c.ApiAlmIntegrationsSearchBitbucketserverReposWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmIntegrationsSearchGitlabReposWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmIntegrationsSearchGitlabReposParams{}
	params.AlmSetting = "settings_key_gitlab"
	resp, err := c.ApiAlmIntegrationsSearchGitlabReposWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmIntegrationsSetPatWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmIntegrationsSetPatParams{}
	params.AlmSetting = "settings_key"
	params.Pat = "pat"
	resp, err := c.ApiAlmIntegrationsSetPatWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsCountBindingWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsCountBindingParams{}
	params.AlmSetting = "settings_key"
	resp, err := c.ApiAlmSettingsCountBindingWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmSettingsCreateAzureWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsCreateAzureParams{}
	params.Key = "settings_key_azure"
	params.PersonalAccessToken = "azure_pat"
	params.Url = "http://example.com"
	resp, err := c.ApiAlmSettingsCreateAzureWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsCreateBitbucketWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsCreateBitbucketParams{}
	params.Key = "settings_key_bitbucket"
	params.PersonalAccessToken = "bitbucket_pat"
	params.Url = "http://example.com"
	resp, err := c.ApiAlmSettingsCreateBitbucketWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsCreateBitbucketcloudWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsCreateBitbucketcloudParams{}
	params.ClientId = "bitbucket_client_id"
	params.ClientSecret = "bitbucket_client_secret"
	params.Key = "settings_key_bitbucketcloud"
	params.Workspace = "test_workspace"
	resp, err := c.ApiAlmSettingsCreateBitbucketcloudWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsCreateGithubWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsCreateGithubParams{}
	params.AppId = "github_app_id"
	params.ClientId = "github_client_id"
	params.ClientSecret = "github_client_secret"
	params.Key = "settings_key_github"
	params.PrivateKey = "github_app_private_key"
	params.Url = "http://example.com"
	resp, err := c.ApiAlmSettingsCreateGithubWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsCreateGitlabWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsCreateGitlabParams{}
	params.Key = "settings_key_gitlab"
	params.PersonalAccessToken = "gitlab_pat"
	params.Url = "http://example.com"
	resp, err := c.ApiAlmSettingsCreateGitlabWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsDeleteWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsDeleteParams{}
	params.Key = "settings_key"
	resp, err := c.ApiAlmSettingsDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsGetBindingWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsGetBindingParams{}
	params.Project = "test_project"
	resp, err := c.ApiAlmSettingsGetBindingWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmSettingsListWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiAlmSettingsListParams{}
	resp, err := c.ApiAlmSettingsListWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmSettingsListDefinitionsWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiAlmSettingsListDefinitionsWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAlmSettingsUpdateAzureWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsUpdateAzureParams{}
	params.Key = "settings_key_azure"
	params.Url = "http://example.com"
	resp, err := c.ApiAlmSettingsUpdateAzureWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsUpdateBitbucketWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsUpdateBitbucketParams{}
	params.Key = "settings_key_bitbucket"
	params.Url = "http://example.com"
	resp, err := c.ApiAlmSettingsUpdateBitbucketWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsUpdateBitbucketcloudWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsUpdateBitbucketcloudParams{}
	params.ClientId = "bitbucket_client_id"
	params.Key = "settings_key_bitbucketcloud"
	params.Workspace = "test_workspace"
	resp, err := c.ApiAlmSettingsUpdateBitbucketcloudWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsUpdateGithubWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsUpdateGithubParams{}
	params.AppId = "github_app_id"
	params.ClientId = "github_client_id"
	params.Key = "settings_key_github"
	params.Url = "http://example.com"
	resp, err := c.ApiAlmSettingsUpdateGithubWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsUpdateGitlabWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsUpdateGitlabParams{}
	params.Key = "settings_key_gitlab"
	params.Url = "http://example.com"
	resp, err := c.ApiAlmSettingsUpdateGitlabWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAlmSettingsValidateWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiAlmSettingsValidateParams{}
	params.Key = "settings_key"
	resp, err := c.ApiAlmSettingsValidateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiAnalysisCacheGetWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiAnalysisCacheGetParams{}
	params.Project = "test_project"
	resp, err := c.ApiAnalysisCacheGetWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAuthenticationLoginWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiAuthenticationLoginParams{}
	params.Login = "operator"
	params.Password = "password"
	resp, err := c.ApiAuthenticationLoginWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAuthenticationLogoutWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiAuthenticationLogoutWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiAuthenticationValidateWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiAuthenticationValidateWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiCeActivityWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiCeActivityParams{}
	resp, err := c.ApiCeActivityWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiCeActivityStatusWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiCeActivityStatusParams{}
	resp, err := c.ApiCeActivityStatusWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiCeComponentWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiCeComponentParams{}
	params.Component = "test_project"
	resp, err := c.ApiCeComponentWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiCeTaskWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiCeTaskParams{}
	params.Id = "test_task_id"
	resp, err := c.ApiCeTaskWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiComponentsSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiComponentsSearchParams{}
	resp, err := c.ApiComponentsSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiComponentsShowWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiComponentsShowParams{}
	params.Component = "test_project"
	resp, err := c.ApiComponentsShowWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiComponentsTreeWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiComponentsTreeParams{}
	params.Component = "test_project"
	resp, err := c.ApiComponentsTreeWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiDuplicationsShowWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiDuplicationsShowParams{}
	params.Key = "test_project:/src/main.go"
	resp, err := c.ApiDuplicationsShowWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiFavoritesAddWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiFavoritesAddParams{}
	params.Component = "test_project"
	resp, err := c.ApiFavoritesAddWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiFavoritesRemoveWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiFavoritesRemoveParams{}
	params.Component = "test_project"
	resp, err := c.ApiFavoritesRemoveWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiFavoritesSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiFavoritesSearchParams{}
	resp, err := c.ApiFavoritesSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiHotspotsSearchWithResponse(t *testing.T) {
	c := newClient(t)

	projectKey := "test_project"
	params := ApiHotspotsSearchParams{}
	params.ProjectKey = &projectKey
	resp, err := c.ApiHotspotsSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiHotspotsShowWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiHotspotsShowParams{}
	params.Hotspot = "test_hotspot_id"
	resp, err := c.ApiHotspotsShowWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesAddCommentWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesAddCommentParams{}
	params.Issue = "test_issue_key"
	params.Text = "Test Comment"
	resp, err := c.ApiIssuesAddCommentWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesAssignWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesAssignParams{}
	params.Issue = "test_issue_key"
	resp, err := c.ApiIssuesAssignWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesAuthorsWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesAuthorsParams{}
	resp, err := c.ApiIssuesAuthorsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesBulkChangeWithResponse(t *testing.T) {
	c := newClient(t)

	setType := "BUG"
	params := ApiIssuesBulkChangeParams{}
	params.Issues = "test_issue_key"
	params.SetType = &setType
	resp, err := c.ApiIssuesBulkChangeWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesChangelogWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesChangelogParams{}
	params.Issue = "test_issue_key"
	resp, err := c.ApiIssuesChangelogWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesDeleteCommentWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesDeleteCommentParams{}
	params.Comment = "test_comment_id"
	resp, err := c.ApiIssuesDeleteCommentWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesDoTransitionWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesDoTransitionParams{}
	params.Issue = "test_issue_key"
	params.Transition = "resolve"
	resp, err := c.ApiIssuesDoTransitionWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesEditCommentWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesEditCommentParams{}
	params.Comment = "test_comment_id"
	params.Text = "Test Comment"
	resp, err := c.ApiIssuesEditCommentWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesReindexWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesReindexParams{}
	params.Project = "test_project"
	resp, err := c.ApiIssuesReindexWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiIssuesSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesSearchParams{}
	resp, err := c.ApiIssuesSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesSetSeverityWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesSetSeverityParams{}
	params.Issue = "test_issue_key"
	params.Severity = "BLOCKER"
	resp, err := c.ApiIssuesSetSeverityWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesSetTagsWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesSetTagsParams{}
	params.Issue = "test_issue_key"
	resp, err := c.ApiIssuesSetTagsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesSetTypeWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesSetTypeParams{}
	params.Issue = "test_issue_key"
	params.Type = "BUG"
	resp, err := c.ApiIssuesSetTypeWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiIssuesTagsWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiIssuesTagsParams{}
	resp, err := c.ApiIssuesTagsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiLanguagesListWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiLanguagesListParams{}
	resp, err := c.ApiLanguagesListWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiMeasuresComponentWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiMeasuresComponentParams{}
	params.Component = "test_project"
	params.MetricKeys = "ncloc,complexity"
	resp, err := c.ApiMeasuresComponentWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiMeasuresComponentTreeWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiMeasuresComponentTreeParams{}
	params.Component = "test_project"
	params.MetricKeys = "ncloc,complexity"
	resp, err := c.ApiMeasuresComponentTreeWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiMeasuresSearchHistoryWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiMeasuresSearchHistoryParams{}
	params.Component = "test_project"
	params.Metrics = "ncloc"
	resp, err := c.ApiMeasuresSearchHistoryWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiMetricsSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiMetricsSearchParams{}
	resp, err := c.ApiMetricsSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiMetricsTypesWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiMetricsTypesWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiMonitoringMetricsWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiMonitoringMetricsWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.Body)
}

func TestApiNewCodePeriodsListWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiNewCodePeriodsListParams{}
	params.Project = "test_project"
	resp, err := c.ApiNewCodePeriodsListWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiNewCodePeriodsSetWithResponse(t *testing.T) {
	c := newClient(t)

	project := "test_project"
	value := "30"
	params := ApiNewCodePeriodsSetParams{}
	params.Project = &project
	params.Type = "NUMBER_OF_DAYS"
	params.Value = &value
	resp, err := c.ApiNewCodePeriodsSetWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiNewCodePeriodsShowWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiNewCodePeriodsShowParams{}
	resp, err := c.ApiNewCodePeriodsShowWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiNewCodePeriodsUnsetWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiNewCodePeriodsUnsetParams{}
	resp, err := c.ApiNewCodePeriodsUnsetWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiNotificationsAddWithResponse(t *testing.T) {
	c := newClient(t)

	project := "test_project"
	params := ApiNotificationsAddParams{}
	params.Project = &project
	params.Type = "NewIssues"
	resp, err := c.ApiNotificationsAddWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiNotificationsListWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiNotificationsListParams{}
	resp, err := c.ApiNotificationsListWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiNotificationsRemoveWithResponse(t *testing.T) {
	c := newClient(t)

	project := "test_project"
	params := ApiNotificationsRemoveParams{}
	params.Project = &project
	params.Type = "NewIssues"
	resp, err := c.ApiNotificationsRemoveWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsAddGroupWithResponse(t *testing.T) {
	c := newClient(t)

	groupId := "test_group_id"
	projectKey := "test_project"
	params := ApiPermissionsAddGroupParams{}
	params.GroupId = &groupId
	params.Permission = "user"
	params.ProjectKey = &projectKey
	resp, err := c.ApiPermissionsAddGroupWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsAddGroupToTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	groupId := "test_group_id"
	params := ApiPermissionsAddGroupToTemplateParams{}
	params.GroupId = &groupId
	params.Permission = "user"
	resp, err := c.ApiPermissionsAddGroupToTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsAddProjectCreatorToTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	templateId := "test_template_id"
	params := ApiPermissionsAddProjectCreatorToTemplateParams{}
	params.Permission = "user"
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsAddProjectCreatorToTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsAddUserWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiPermissionsAddUserParams{}
	params.Login = "operator"
	params.Login = "user"
	resp, err := c.ApiPermissionsAddUserWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsAddUserToTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	templateId := "test_template_id"
	params := ApiPermissionsAddUserToTemplateParams{}
	params.Login = "operator"
	params.Permission = "user"
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsAddUserToTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsApplyTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	projectKey := "test_project"
	templateId := "test_template_id"
	params := ApiPermissionsApplyTemplateParams{}
	params.ProjectKey = &projectKey
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsApplyTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsBulkApplyTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	templateId := "test_template_id"
	params := ApiPermissionsBulkApplyTemplateParams{}
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsBulkApplyTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsCreateTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiPermissionsCreateTemplateParams{}
	params.Name = "Test Template"
	resp, err := c.ApiPermissionsCreateTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiPermissionsDeleteTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	templateId := "test_template_id"
	params := ApiPermissionsDeleteTemplateParams{}
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsDeleteTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsRemoveGroupWithResponse(t *testing.T) {
	c := newClient(t)

	groupId := "test_group_id"
	projectKey := "test_project"
	params := ApiPermissionsRemoveGroupParams{}
	params.GroupId = &groupId
	params.Permission = "user"
	params.ProjectKey = &projectKey
	resp, err := c.ApiPermissionsRemoveGroupWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsRemoveGroupFromTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	groupId := "test_group_id"
	templateId := "test_template_id"
	params := ApiPermissionsRemoveGroupFromTemplateParams{}
	params.GroupId = &groupId
	params.Permission = "user"
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsRemoveGroupFromTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsRemoveProjectCreatorFromTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	templateId := "test_template_id"
	params := ApiPermissionsRemoveProjectCreatorFromTemplateParams{}
	params.Permission = "user"
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsRemoveProjectCreatorFromTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsRemoveUserWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiPermissionsRemoveUserParams{}
	params.Login = "operator"
	params.Permission = "user"
	resp, err := c.ApiPermissionsRemoveUserWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsRemoveUserFromTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	templateId := "test_template_id"
	params := ApiPermissionsRemoveUserFromTemplateParams{}
	params.Login = "operator"
	params.Permission = "user"
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsRemoveUserFromTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsSearchTemplatesWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiPermissionsSearchTemplatesParams{}
	resp, err := c.ApiPermissionsSearchTemplatesWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiPermissionsSetDefaultTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	templateId := "test_template_id"
	params := ApiPermissionsSetDefaultTemplateParams{}
	params.TemplateId = &templateId
	resp, err := c.ApiPermissionsSetDefaultTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPermissionsUpdateTemplateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiPermissionsUpdateTemplateParams{}
	params.Id = "test_template_id"
	resp, err := c.ApiPermissionsUpdateTemplateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiPluginsAvailableWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiPluginsAvailableWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiPluginsCancelAllWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiPluginsCancelAllWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPluginsInstallWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiPluginsInstallParams{}
	params.Key = "test_plugin"
	resp, err := c.ApiPluginsInstallWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPluginsInstalledWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiPluginsInstalledParams{}
	resp, err := c.ApiPluginsInstalledWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiPluginsPendingWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiPluginsPendingWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiPluginsUninstallWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiPluginsUninstallParams{}
	params.Key = "test_plugin"
	resp, err := c.ApiPluginsUninstallWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPluginsUpdateWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiPluginsUpdateParams{}
	params.Key = "test_plugin"
	resp, err := c.ApiPluginsUpdateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiPluginsUpdatesWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiPluginsUpdatesWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectAnalysesCreateEventWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectAnalysesCreateEventParams{}
	params.Analysis = "test_analysis"
	params.Name = "Test Event"
	resp, err := c.ApiProjectAnalysesCreateEventWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectAnalysesDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectAnalysesDeleteParams{}
	params.Analysis = "test_analysis"
	resp, err := c.ApiProjectAnalysesDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectAnalysesDeleteEventWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectAnalysesDeleteEventParams{}
	params.Event = "test_event"
	resp, err := c.ApiProjectAnalysesDeleteEventWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectAnalysesSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectAnalysesSearchParams{}
	params.Project = "test_project"
	resp, err := c.ApiProjectAnalysesSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectAnalysesSetBaselineWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectAnalysesSetBaselineParams{}
	params.Analysis = "test_analysis"
	params.Project = "test_project"
	resp, err := c.ApiProjectAnalysesSetBaselineWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectAnalysesUnsetBaselineWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectAnalysesUnsetBaselineParams{}
	params.Project = "test_project"
	resp, err := c.ApiProjectAnalysesUnsetBaselineWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectAnalysesUpdateEventWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectAnalysesUpdateEventParams{}
	params.Event = "test_event"
	params.Name = "Test Event"
	resp, err := c.ApiProjectAnalysesUpdateEventWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectBadgesMeasureWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectBadgesMeasureParams{}
	params.Metric = "bugs"
	params.Project = "test_project"
	resp, err := c.ApiProjectBadgesMeasureWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.Body)
}

func TestApiProjectBadgesQualityGateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectBadgesQualityGateParams{}
	params.Project = "test_project"
	resp, err := c.ApiProjectBadgesQualityGateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.Body)
}

func TestApiProjectBadgesRenewTokenWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectBadgesRenewTokenParams{}
	params.Project = "test_project"
	resp, err := c.ApiProjectBadgesRenewTokenWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectBadgesTokenWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectBadgesTokenParams{}
	params.Project = "test_project"
	resp, err := c.ApiProjectBadgesTokenWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectBranchesDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectBranchesDeleteParams{}
	params.Branch = "feature/test"
	params.Project = "test_project"
	resp, err := c.ApiProjectBranchesDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectBranchesListWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectBranchesListParams{}
	params.Project = "test_project"
	resp, err := c.ApiProjectBranchesListWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectBranchesRenameWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectBranchesRenameParams{}
	params.Name = "feature/test"
	params.Project = "test_project"
	resp, err := c.ApiProjectBranchesRenameWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectBranchesSetAutomaticDeletionProtectionWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectBranchesSetAutomaticDeletionProtectionParams{}
	params.Branch = "feature/test"
	params.Project = "test_project"
	params.Value = "true"
	resp, err := c.ApiProjectBranchesSetAutomaticDeletionProtectionWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectDumpExportWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectDumpExportParams{}
	params.Key = "test_project"
	resp, err := c.ApiProjectDumpExportWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectLinksCreateWithResponse(t *testing.T) {
	c := newClient(t)

	projectKey := "test_project"
	params := ApiProjectLinksCreateParams{}
	params.Name = "Test Link"
	params.ProjectKey = &projectKey
	params.Url = "http://example.com"
	resp, err := c.ApiProjectLinksCreateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectLinksDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectLinksDeleteParams{}
	params.Id = "test_link_id"
	resp, err := c.ApiProjectLinksDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectLinksSearchWithResponse(t *testing.T) {
	c := newClient(t)

	projectKey := "test_project"
	params := ApiProjectLinksSearchParams{}
	params.ProjectKey = &projectKey
	resp, err := c.ApiProjectLinksSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectTagsSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectTagsSearchParams{}
	resp, err := c.ApiProjectTagsSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectTagsSetWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectTagsSetParams{}
	params.Project = "test_project"
	params.Tags = "tag1,tag2"
	resp, err := c.ApiProjectTagsSetWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectsBulkDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	projects := "test_project_bulkdeleted"
	params := ApiProjectsBulkDeleteParams{}
	params.Projects = &projects
	resp, err := c.ApiProjectsBulkDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectsCreateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectsCreateParams{}
	params.Name = "Test Project Created"
	params.Project = "test_project_created"
	resp, err := c.ApiProjectsCreateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectsDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectsDeleteParams{}
	params.Project = "test_project_deleted"
	resp, err := c.ApiProjectsDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectsSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectsSearchParams{}
	resp, err := c.ApiProjectsSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiProjectsUpdateKeyWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectsUpdateKeyParams{}
	params.From = "old_project_key"
	params.To = "new_project_key"
	resp, err := c.ApiProjectsUpdateKeyWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiProjectsUpdateVisibilityWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiProjectsUpdateVisibilityParams{}
	params.Project = "test_project"
	params.Visibility = "public"
	resp, err := c.ApiProjectsUpdateVisibilityWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualitygatesCopyWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_gate_id"
	params := ApiQualitygatesCopyParams{}
	params.Id = &id
	params.Name = "Test Gate"
	resp, err := c.ApiQualitygatesCopyWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualitygatesCreateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualitygatesCreateParams{}
	params.Name = "Test Gate"
	resp, err := c.ApiQualitygatesCreateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualitygatesCreateConditionWithResponse(t *testing.T) {
	c := newClient(t)

	gateId := "test_gate_id"
	op := "GT"
	params := ApiQualitygatesCreateConditionParams{}
	params.Error = "0"
	params.GateId = &gateId
	params.Metric = "vulnerabilities"
	params.Op = &op
	resp, err := c.ApiQualitygatesCreateConditionWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualitygatesDeleteConditionWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualitygatesDeleteConditionParams{}
	params.Id = "test_condition_id"
	resp, err := c.ApiQualitygatesDeleteConditionWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualitygatesDeselectWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualitygatesDeselectParams{}
	params.ProjectKey = "test_project"
	resp, err := c.ApiQualitygatesDeselectWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualitygatesDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_gate_id"
	params := ApiQualitygatesDestroyParams{}
	params.Id = &id
	resp, err := c.ApiQualitygatesDestroyWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualitygatesGetByProjectWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualitygatesGetByProjectParams{}
	params.Project = "test_project"
	resp, err := c.ApiQualitygatesGetByProjectWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualitygatesListWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiQualitygatesListWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualitygatesProjectStatusWithResponse(t *testing.T) {
	c := newClient(t)

	projectKey := "test_project"
	params := ApiQualitygatesProjectStatusParams{}
	params.ProjectKey = &projectKey
	resp, err := c.ApiQualitygatesProjectStatusWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualitygatesRenameWithResponse(t *testing.T) {
	c := newClient(t)

	gateId := "test_gate_id"
	params := ApiQualitygatesRenameParams{}
	params.Id = &gateId
	params.Name = "Test GAte"
	resp, err := c.ApiQualitygatesRenameWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualitygatesSearchWithResponse(t *testing.T) {
	c := newClient(t)

	gateId := "test_gate_id"
	params := ApiQualitygatesSearchParams{}
	params.GateId = &gateId
	resp, err := c.ApiQualitygatesSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualitygatesSelectWithResponse(t *testing.T) {
	c := newClient(t)

	gateId := "test_gate_id"
	params := ApiQualitygatesSelectParams{}
	params.GateId = &gateId
	params.ProjectKey = "test_project"
	resp, err := c.ApiQualitygatesSelectWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualitygatesSetAsDefaultWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_gate_id"
	params := ApiQualitygatesSetAsDefaultParams{}
	params.Id = &id
	resp, err := c.ApiQualitygatesSetAsDefaultWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualitygatesShowWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_gate_id"
	params := ApiQualitygatesShowParams{}
	params.Id = &id
	resp, err := c.ApiQualitygatesShowWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualitygatesUpdateConditionWithResponse(t *testing.T) {
	c := newClient(t)

	op := "GT"
	params := ApiQualitygatesUpdateConditionParams{}
	params.Error = "0"
	params.Id = "test_condition_id"
	params.Metric = "vulnerabilities"
	params.Op = &op
	resp, err := c.ApiQualitygatesUpdateConditionWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesActivateRuleWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesActivateRuleParams{}
	params.Key = "test_profile_id"
	params.Rule = "test_rule_id"
	resp, err := c.ApiQualityprofilesActivateRuleWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesActivateRulesWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesActivateRulesParams{}
	params.TargetKey = "test_profile_id"
	resp, err := c.ApiQualityprofilesActivateRulesWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesAddProjectWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesAddProjectParams{}
	params.Language = "go"
	params.Project = "test_project"
	params.QualityProfile = "Test Profile"
	resp, err := c.ApiQualityprofilesAddProjectWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesBackupWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesBackupParams{}
	params.Language = "go"
	params.QualityProfile = "Test Profile"
	resp, err := c.ApiQualityprofilesBackupWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.XML2XX)
}

func TestApiQualityprofilesChangeParentWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesChangeParentParams{}
	params.Language = "go"
	params.QualityProfile = "Test Profile"
	resp, err := c.ApiQualityprofilesChangeParentWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesChangelogWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesChangelogParams{}
	params.Language = "go"
	params.QualityProfile = "Test Profile"
	resp, err := c.ApiQualityprofilesChangelogWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualityprofilesCopyWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesCopyParams{}
	params.FromKey = "test_profile_id"
	params.ToName = "Test Profile Copied"
	resp, err := c.ApiQualityprofilesCopyWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualityprofilesCreateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesCreateParams{}
	params.Language = "go"
	params.Name = "Test Profile"
	resp, err := c.ApiQualityprofilesCreateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualityprofilesDeactivateRuleWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesDeactivateRuleParams{}
	params.Key = "test_profile_id"
	params.Rule = "go:test_rule_id"
	resp, err := c.ApiQualityprofilesDeactivateRuleWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesDeactivateRulesWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesDeactivateRulesParams{}
	params.TargetKey = "test_profile_id"
	resp, err := c.ApiQualityprofilesDeactivateRulesWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesDeleteParams{}
	params.Language = "go"
	params.QualityProfile = "Test Profile"
	resp, err := c.ApiQualityprofilesDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesExportWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesExportParams{}
	params.Language = "go"
	resp, err := c.ApiQualityprofilesExportWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.XML2XX)
}

func TestApiQualityprofilesExportersWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiQualityprofilesExportersWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualityprofilesImportersWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiQualityprofilesImportersWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualityprofilesInheritanceWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesInheritanceParams{}
	params.Language = "go"
	params.QualityProfile = "Test Profile"
	resp, err := c.ApiQualityprofilesInheritanceWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualityprofilesProjectsWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesProjectsParams{}
	params.Key = "test_profile_id"
	resp, err := c.ApiQualityprofilesProjectsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualityprofilesRemoveProjectWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesRemoveProjectParams{}
	params.Language = "go"
	params.Project = "test_project"
	params.QualityProfile = "Test Profile"
	resp, err := c.ApiQualityprofilesRemoveProjectWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesRenameWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesRenameParams{}
	params.Key = "test_profile_id"
	params.Name = "Test Profile Renamed"
	resp, err := c.ApiQualityprofilesRenameWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesRestoreWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	params := ApiQualityprofilesRestoreParams{}
	params.Backup = ""
	resp, err := c.ApiQualityprofilesRestoreWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiQualityprofilesSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesSearchParams{}
	resp, err := c.ApiQualityprofilesSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiQualityprofilesSetDefaultWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiQualityprofilesSetDefaultParams{}
	params.Language = "go"
	params.QualityProfile = "Test Profile"
	resp, err := c.ApiQualityprofilesSetDefaultWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiRulesCreateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiRulesCreateParams{}
	params.CustomKey = "go:test_rule_id"
	params.MarkdownDescription = "# Test Custom Rule"
	params.Name = "Test Custom Rule"
	params.TemplateKey = "test_template_id"
	resp, err := c.ApiRulesCreateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiRulesDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiRulesDeleteParams{}
	params.Key = "go:test_rule_id"
	resp, err := c.ApiRulesDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiRulesRepositoriesWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiRulesRepositoriesParams{}
	resp, err := c.ApiRulesRepositoriesWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiRulesSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiRulesSearchParams{}
	resp, err := c.ApiRulesSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiRulesShowWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiRulesShowParams{}
	params.Key = "go:test_rule_id"
	resp, err := c.ApiRulesShowWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiRulesTagsWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiRulesTagsParams{}
	resp, err := c.ApiRulesTagsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiRulesUpdateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiRulesUpdateParams{}
	params.Key = "go:test_rule_id"
	resp, err := c.ApiRulesUpdateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiServerVersionWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiServerVersionWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.Body)
}

func TestApiSettingsListDefinitionsWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiSettingsListDefinitionsParams{}
	resp, err := c.ApiSettingsListDefinitionsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSettingsResetWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiSettingsResetParams{}
	params.Keys = "settings_key"
	resp, err := c.ApiSettingsResetWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiSettingsSetWithResponse(t *testing.T) {
	c := newClient(t)

	value := "settings_value"
	params := ApiSettingsSetParams{}
	params.Key = "settings_key"
	params.Value = &value
	resp, err := c.ApiSettingsSetWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiSettingsValuesWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiSettingsValuesParams{}
	resp, err := c.ApiSettingsValuesWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSourcesRawWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiSourcesRawParams{}
	params.Key = "test_project:/src/main.go"
	resp, err := c.ApiSourcesRawWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.Body)
}

func TestApiSourcesScmWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiSourcesScmParams{}
	params.Key = "test_project:/src/main.go"
	resp, err := c.ApiSourcesScmWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSourcesShowWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiSourcesShowParams{}
	params.Key = "test_project:/src/main.go"
	resp, err := c.ApiSourcesShowWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSystemChangeLogLevelWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiSystemChangeLogLevelParams{}
	params.Level = "INFO"
	resp, err := c.ApiSystemChangeLogLevelWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiSystemDbMigrationStatusWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiSystemDbMigrationStatusWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSystemHealthWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiSystemHealthWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSystemInfoWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiSystemInfoWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSystemLogsWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiSystemLogsParams{}
	resp, err := c.ApiSystemLogsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.Body)
}

func TestApiSystemMigrateDbWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiSystemMigrateDbWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSystemPingWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiSystemPingWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.Body)
}

func TestApiSystemRestartWithResponse(t *testing.T) {
	t.Skip()

	c := newClient(t)

	resp, err := c.ApiSystemRestartWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiSystemStatusWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiSystemStatusWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiSystemUpgradesWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiSystemUpgradesWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUserGroupsAddUserWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_group_id"
	params := ApiUserGroupsAddUserParams{}
	params.Id = &id
	resp, err := c.ApiUserGroupsAddUserWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUserGroupsCreateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUserGroupsCreateParams{}
	params.Name = "Test Group"
	resp, err := c.ApiUserGroupsCreateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUserGroupsDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_group_id"
	params := ApiUserGroupsDeleteParams{}
	params.Id = &id
	resp, err := c.ApiUserGroupsDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUserGroupsRemoveUserWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_group_id"
	params := ApiUserGroupsRemoveUserParams{}
	params.Id = &id
	resp, err := c.ApiUserGroupsRemoveUserWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUserGroupsSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUserGroupsSearchParams{}
	resp, err := c.ApiUserGroupsSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUserGroupsUpdateWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_group_id"
	params := ApiUserGroupsUpdateParams{}
	params.Id = &id
	resp, err := c.ApiUserGroupsUpdateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUserGroupsUsersWithResponse(t *testing.T) {
	c := newClient(t)

	id := "test_group_id"
	params := ApiUserGroupsUsersParams{}
	params.Id = &id
	resp, err := c.ApiUserGroupsUsersWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUserTokensGenerateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUserTokensGenerateParams{}
	params.Name = "Test Token"
	resp, err := c.ApiUserTokensGenerateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUserTokensRevokeWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUserTokensRevokeParams{}
	params.Name = "Test Token"
	resp, err := c.ApiUserTokensRevokeWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUserTokensSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUserTokensSearchParams{}
	resp, err := c.ApiUserTokensSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUsersAnonymizeWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUsersAnonymizeParams{}
	params.Login = "operator"
	resp, err := c.ApiUsersAnonymizeWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUsersChangePasswordWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUsersChangePasswordParams{}
	params.Login = "operator"
	params.Password = "operatorpassword"
	resp, err := c.ApiUsersChangePasswordWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUsersCreateWithResponse(t *testing.T) {
	c := newClient(t)

	local := "true"
	password := "operatorpassword"
	params := ApiUsersCreateParams{}
	params.Local = &local
	params.Login = "operator"
	params.Name = "Operator"
	params.Password = &password
	resp, err := c.ApiUsersCreateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUsersDeactivateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUsersDeactivateParams{}
	params.Login = "operator"
	resp, err := c.ApiUsersDeactivateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUsersDismissSonarlintAdWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ApiUsersDismissSonarlintAdWithResponse(context.TODO(), basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUsersGroupsWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUsersGroupsParams{}
	params.Login = "operator"
	resp, err := c.ApiUsersGroupsWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUsersSearchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUsersSearchParams{}
	resp, err := c.ApiUsersSearchWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUsersUpdateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUsersUpdateParams{}
	params.Login = "operator"
	resp, err := c.ApiUsersUpdateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiUsersUpdateIdentityProviderWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUsersUpdateIdentityProviderParams{}
	params.Login = "operator"
	params.NewExternalProvider = "LDAP_serverKey"
	resp, err := c.ApiUsersUpdateIdentityProviderWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiUsersUpdateLoginWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiUsersUpdateLoginParams{}
	params.Login = "operator"
	params.NewLogin = "new_operator"
	resp, err := c.ApiUsersUpdateLoginWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiWebhooksCreateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiWebhooksCreateParams{}
	params.Name = "Test Webhook"
	params.Url = "http://example.com/"
	resp, err := c.ApiWebhooksCreateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiWebhooksDeleteWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiWebhooksDeleteParams{}
	params.Webhook = "test_webhook_id"
	resp, err := c.ApiWebhooksDeleteWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiWebhooksDeliveriesWithResponse(t *testing.T) {
	c := newClient(t)

	componentKey := "test_project"
	params := ApiWebhooksDeliveriesParams{}
	params.ComponentKey = &componentKey
	resp, err := c.ApiWebhooksDeliveriesWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiWebhooksDeliveryWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiWebhooksDeliveryParams{}
	params.DeliveryId = "test_delivery_id"
	resp, err := c.ApiWebhooksDeliveryWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiWebhooksListWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiWebhooksListParams{}
	resp, err := c.ApiWebhooksListWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiWebhooksUpdateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiWebhooksUpdateParams{}
	params.Name = "Test Webhook Updated"
	params.Url = "http://example.com/"
	params.Webhook = "test_webhook_id"
	resp, err := c.ApiWebhooksUpdateWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
}

func TestApiWebservicesListWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiWebservicesListParams{}
	resp, err := c.ApiWebservicesListWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}

func TestApiWebservicesResponseExampleWithResponse(t *testing.T) {
	c := newClient(t)

	params := ApiWebservicesResponseExampleParams{}
	params.Action = "version"
	params.Controller = "api/server"
	resp, err := c.ApiWebservicesResponseExampleWithResponse(context.TODO(), &params, basicAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, resp.JSONDefault)
	assertResponseBody(t, &resp.JSON2XX)
}
