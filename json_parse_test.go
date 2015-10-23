package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
)

var (
	body = "{\"ref\":\"refs/heads/master\",\"before\":\"1ca6580870f36b153388e3fe2e22a9d5d1e0aa76\",\"after\":\"bb77d44c49be3928a88f9ebf1a05047137f3a967\",\"created\":false,\"deleted\":false,\"forced\":false,\"base_ref\":null,\"compare\":\"https://github.com/stevenzeiler/rippled/compare/1ca6580870f3...bb77d44c49be\",\"commits\":[{\"id\":\"bb77d44c49be3928a88f9ebf1a05047137f3a967\",\"distinct\":true,\"message\":\"ipple\",\"timestamp\":\"2015-10-22T16:11:22-07:00\",\"url\":\"https://github.com/stevenzeiler/rippled/commit/bb77d44c49be3928a88f9ebf1a05047137f3a967\",\"author\":{\"name\":\"Steven Zeiler\",\"email\":\"stevenzeiler@Stevens-MacBook-Air.local\"},\"committer\":{\"name\":\"Steven Zeiler\",\"email\":\"stevenzeiler@Stevens-MacBook-Air.local\"},\"added\":[],\"removed\":[],\"modified\":[\"README.md\"]}],\"head_commit\":{\"id\":\"bb77d44c49be3928a88f9ebf1a05047137f3a967\",\"distinct\":true,\"message\":\"ipple\",\"timestamp\":\"2015-10-22T16:11:22-07:00\",\"url\":\"https://github.com/stevenzeiler/rippled/commit/bb77d44c49be3928a88f9ebf1a05047137f3a967\",\"author\":{\"name\":\"Steven Zeiler\",\"email\":\"stevenzeiler@Stevens-MacBook-Air.local\"},\"committer\":{\"name\":\"Steven Zeiler\",\"email\":\"stevenzeiler@Stevens-MacBook-Air.local\"},\"added\":[],\"removed\":[],\"modified\":[\"README.md\"]},\"repository\":{\"id\":42969039,\"name\":\"rippled\",\"full_name\":\"stevenzeiler/rippled\",\"owner\":{\"name\":\"stevenzeiler\",\"email\":\"zeiler.steven@gmail.com\"},\"private\":false,\"html_url\":\"https://github.com/stevenzeiler/rippled\",\"description\":\"Ripple peer-to-peer network daemon\",\"fork\":true,\"url\":\"https://github.com/stevenzeiler/rippled\",\"forks_url\":\"https://api.github.com/repos/stevenzeiler/rippled/forks\",\"keys_url\":\"https://api.github.com/repos/stevenzeiler/rippled/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/stevenzeiler/rippled/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/stevenzeiler/rippled/teams\",\"hooks_url\":\"https://api.github.com/repos/stevenzeiler/rippled/hooks\",\"issue_events_url\":\"https://api.github.com/repos/stevenzeiler/rippled/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/stevenzeiler/rippled/events\",\"assignees_url\":\"https://api.github.com/repos/stevenzeiler/rippled/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/stevenzeiler/rippled/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/stevenzeiler/rippled/tags\",\"blobs_url\":\"https://api.github.com/repos/stevenzeiler/rippled/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/stevenzeiler/rippled/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/stevenzeiler/rippled/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/stevenzeiler/rippled/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/stevenzeiler/rippled/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/stevenzeiler/rippled/languages\",\"stargazers_url\":\"https://api.github.com/repos/stevenzeiler/rippled/stargazers\",\"contributors_url\":\"https://api.github.com/repos/stevenzeiler/rippled/contributors\",\"subscribers_url\":\"https://api.github.com/repos/stevenzeiler/rippled/subscribers\",\"subscription_url\":\"https://api.github.com/repos/stevenzeiler/rippled/subscription\",\"commits_url\":\"https://api.github.com/repos/stevenzeiler/rippled/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/stevenzeiler/rippled/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/stevenzeiler/rippled/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/stevenzeiler/rippled/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/stevenzeiler/rippled/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/stevenzeiler/rippled/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/stevenzeiler/rippled/merges\",\"archive_url\":\"https://api.github.com/repos/stevenzeiler/rippled/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/stevenzeiler/rippled/downloads\",\"issues_url\":\"https://api.github.com/repos/stevenzeiler/rippled/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/stevenzeiler/rippled/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/stevenzeiler/rippled/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/stevenzeiler/rippled/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/stevenzeiler/rippled/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/stevenzeiler/rippled/releases{/id}\",\"created_at\":1442968461,\"updated_at\":\"2015-09-23T00:34:37Z\",\"pushed_at\":1445555492,\"git_url\":\"git://github.com/stevenzeiler/rippled.git\",\"ssh_url\":\"git@github.com:stevenzeiler/rippled.git\",\"clone_url\":\"https://github.com/stevenzeiler/rippled.git\",\"svn_url\":\"https://github.com/stevenzeiler/rippled\",\"homepage\":\"https://ripple.com\",\"size\":57368,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"C++\",\"has_issues\":false,\"has_downloads\":true,\"has_wiki\":false,\"has_pages\":false,\"forks_count\":0,\"mirror_url\":null,\"open_issues_count\":0,\"forks\":0,\"open_issues\":0,\"watchers\":0,\"default_branch\":\"develop\",\"stargazers\":0,\"master_branch\":\"develop\"},\"pusher\":{\"name\":\"stevenzeiler\",\"email\":\"zeiler.steven@gmail.com\"},\"sender\":{\"login\":\"stevenzeiler\",\"id\":508282,\"avatar_url\":\"https://avatars.githubusercontent.com/u/508282?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/stevenzeiler\",\"html_url\":\"https://github.com/stevenzeiler\",\"followers_url\":\"https://api.github.com/users/stevenzeiler/followers\",\"following_url\":\"https://api.github.com/users/stevenzeiler/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/stevenzeiler/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/stevenzeiler/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/stevenzeiler/subscriptions\",\"organizations_url\":\"https://api.github.com/users/stevenzeiler/orgs\",\"repos_url\":\"https://api.github.com/users/stevenzeiler/repos\",\"events_url\":\"https://api.github.com/users/stevenzeiler/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/stevenzeiler/received_events\",\"type\":\"User\",\"site_admin\":false}}"
)

type Body struct {
	Compare string `json:"compare"`
}

func TestParsingJson(t *testing.T) {

	fromSteven := regexp.MustCompile(`stevenzeiler\/rippled`)
	fromRipple := regexp.MustCompile(`ripple\/rippled`)

	dat := new(Body)

	byt := []byte(body)

	if err := json.Unmarshal(byt, dat); err != nil {
		panic(err)
	}

	if fromSteven.MatchString(dat.Compare) {
		fmt.Println("repository is stevenzeiler/rippled")
	}

	if fromRipple.MatchString(dat.Compare) {
		fmt.Println("repository is ripple/rippled")
	}
}
