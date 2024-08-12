package repositoryimpl

type Config struct {
	Logo           string `json:"logo"`
	Name           string `json:"name"`
	Author         string `json:"author"`
	Introduction   string `json:"introduction"`
	Avatar         string `json:"avatar"`
	GithubHomepage string `json:"github_homepage"`
	GiteeHomepage  string `json:"gitee_homepage"`
	CsdnHomepage   string `json:"csdn_homepage"`
	ZhihuHomepage  string `json:"zhihu_homepage"`
}
