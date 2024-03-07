	auth_model "code.gitea.io/gitea/models/auth"
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/git/commits/12345", user.Name).
		AddTokenAuth(token)
	req = NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/git/commits/..", user.Name).
		AddTokenAuth(token)
	req = NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/git/commits/branch-not-exist", user.Name).
		AddTokenAuth(token)
		req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo1/git/commits/%s", user.Name, ref).
			AddTokenAuth(token)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo20/commits?not=master&sha=remove-files-a", user.Name).
		AddTokenAuth(token)
	resp := MakeRequest(t, req, http.StatusOK)

	var apiData []api.Commit
	DecodeJSON(t, resp, &apiData)

	assert.Len(t, apiData, 2)
	assert.EqualValues(t, "cfe3b3c1fd36fba04f9183287b106497e1afe986", apiData[0].CommitMeta.SHA)
	compareCommitFiles(t, []string{"link_hi", "test.csv"}, apiData[0].Files)
	assert.EqualValues(t, "c8e31bc7688741a5287fcde4fbb8fc129ca07027", apiData[1].CommitMeta.SHA)
	compareCommitFiles(t, []string{"test.csv"}, apiData[1].Files)

	assert.EqualValues(t, resp.Header().Get("X-Total"), "2")
}

func TestAPIReposGitCommitListNotMaster(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)

	// Test getting commits (Page 1)
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/commits", user.Name).
		AddTokenAuth(token)

	assert.EqualValues(t, resp.Header().Get("X-Total"), "3")
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/commits?page=2", user.Name).
		AddTokenAuth(token)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/commits?sha=good-sign", user.Name).
		AddTokenAuth(token)
func TestAPIReposGitCommitListWithoutSelectFields(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)

	// Test getting commits without files, verification, and stats
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/commits?sha=good-sign&stat=false&files=false&verification=false", user.Name).
		AddTokenAuth(token)
	resp := MakeRequest(t, req, http.StatusOK)

	var apiData []api.Commit
	DecodeJSON(t, resp, &apiData)

	assert.Len(t, apiData, 1)
	assert.Equal(t, "f27c2b2b03dcab38beaf89b0ab4ff61f6de63441", apiData[0].CommitMeta.SHA)
	assert.Equal(t, (*api.CommitStats)(nil), apiData[0].Stats)
	assert.Equal(t, (*api.PayloadCommitVerification)(nil), apiData[0].RepoCommit.Verification)
	assert.Equal(t, ([]*api.CommitAffectedFiles)(nil), apiData[0].Files)
}

	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	reqDiff := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/git/commits/f27c2b2b03dcab38beaf89b0ab4ff61f6de63441.diff", user.Name).
		AddTokenAuth(token)
	reqPatch := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/git/commits/f27c2b2b03dcab38beaf89b0ab4ff61f6de63441.patch", user.Name).
		AddTokenAuth(token)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/commits?path=readme.md&sha=good-sign", user.Name).
		AddTokenAuth(token)

	assert.EqualValues(t, resp.Header().Get("X-Total"), "1")
}

func TestGetFileHistoryNotOnMaster(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)

	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo20/commits?path=test.csv&sha=add-csv&not=master", user.Name).
		AddTokenAuth(token)
	resp := MakeRequest(t, req, http.StatusOK)

	var apiData []api.Commit
	DecodeJSON(t, resp, &apiData)

	assert.Len(t, apiData, 1)
	assert.Equal(t, "c8e31bc7688741a5287fcde4fbb8fc129ca07027", apiData[0].CommitMeta.SHA)
	compareCommitFiles(t, []string{"test.csv"}, apiData[0].Files)

	assert.EqualValues(t, resp.Header().Get("X-Total"), "1")