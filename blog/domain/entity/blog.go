package entity

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type Blog struct {
	Logo           cmprimitive.Urlx
	Name           cmprimitive.Text
	Author         cmprimitive.Username
	Introduction   cmprimitive.Text
	Avatar         cmprimitive.Urlx
	GithubHomepage cmprimitive.Urlx
	GiteeHomepage  cmprimitive.Urlx
	CsdnHomepage   cmprimitive.Urlx
	ZhihuHomepage  cmprimitive.Urlx
}
