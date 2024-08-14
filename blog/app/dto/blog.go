package dto

import "github.com/victorzhou123/vicblog/blog/domain/entity"

type BlogInformationDto struct {
	Logo           string `json:"logo"`
	Name           string `json:"name"`
	Author         string `json:"author"`
	Introduction   string `json:"introduction"`
	Avatar         string `json:"avatar"`
	GithubHomepage string `json:"githubHomepage"`
	GiteeHomepage  string `json:"giteeHomepage"`
	CsdnHomepage   string `json:"csdnHomepage"`
	ZhihuHomepage  string `json:"zhihuHomepage"`
}

func ToBlogInformationDto(blog entity.Blog) BlogInformationDto {
	return BlogInformationDto{
		Logo:           blog.Logo.Urlx(),
		Name:           blog.Name.Text(),
		Author:         blog.Author.Username(),
		Introduction:   blog.Introduction.Text(),
		Avatar:         blog.Avatar.Urlx(),
		GithubHomepage: blog.GithubHomepage.Urlx(),
		GiteeHomepage:  blog.GiteeHomepage.Urlx(),
		CsdnHomepage:   blog.CsdnHomepage.Urlx(),
		ZhihuHomepage:  blog.ZhihuHomepage.Urlx(),
	}
}
