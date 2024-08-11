package controller

import (
	"victorzhou123/vicblog/article/app/dto"
	articleent "victorzhou123/vicblog/article/domain/article/entity"
	articlesvc "victorzhou123/vicblog/article/domain/article/service"
	"victorzhou123/vicblog/article/domain/category/entity"
	categorysvc "victorzhou123/vicblog/article/domain/category/service"
	tagent "victorzhou123/vicblog/article/domain/tag/entity"
	"victorzhou123/vicblog/article/domain/tag/repository"
	tagsvc "victorzhou123/vicblog/article/domain/tag/service"
	cmctl "victorzhou123/vicblog/common/controller"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type reqCategory struct {
	Name string `json:"name"`
}

func (req *reqCategory) toCategoryName() (entity.CategoryName, error) {
	return entity.NewCategoryName(req.Name)
}

type reqCategoryList struct {
	cmctl.ReqList
}

func (req *reqCategoryList) toCmd() (cmd categorysvc.CategoryListCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()
	if err != nil {
		return
	}

	cmd = categorysvc.CategoryListCmd{
		PaginationCmd: listCmd,
	}

	return cmd, cmd.Validate()
}

type reqTag struct {
	Names []string `json:"names"`
}

func (req *reqTag) toTagNames() (repository.TagNames, error) {
	tagNames := make([]tagent.TagName, len(req.Names))

	var err error
	for i := range req.Names {
		tagNames[i], err = tagent.NewTagName(req.Names[i])
		if err != nil {
			return repository.TagNames{}, err
		}
	}

	return repository.TagNames{Names: tagNames}, nil
}

type reqTagList struct {
	cmctl.ReqList
}

func (req *reqTagList) emptyValue() bool {
	return req.ReqList.EmptyValue()
}

func (req *reqTagList) toCmd() (cmd tagsvc.TagListCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()

	cmd = tagsvc.TagListCmd{
		PaginationCmd: listCmd,
	}

	return cmd, cmd.Validate()
}

type reqArticle struct {
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Cover    string `json:"cover"`
	Category uint   `json:"categoryId"`
	Tags     []uint `json:"tags"`
}

func (req *reqArticle) toCmd(user cmprimitive.Username) (cmd dto.AddArticleCmd, err error) {

	if cmd.Title, err = cmprimitive.NewTitle(req.Title); err != nil {
		return
	}

	if cmd.Summary, err = articleent.NewArticleSummary(req.Summary); err != nil {
		return
	}

	if cmd.Content, err = cmprimitive.NewArticleContent(req.Content); err != nil {
		return
	}

	if cmd.Cover, err = cmprimitive.NewUrlx(req.Cover); err != nil {
		return
	}

	cmd.Owner = user

	cmd.Category = cmprimitive.NewIdByUint(req.Category)

	cmd.Tags = make([]cmprimitive.Id, len(req.Tags))
	for i := range req.Tags {
		cmd.Tags[i] = cmprimitive.NewIdByUint(req.Tags[i])
	}

	return
}

type reqListArticle struct {
	cmctl.ReqList
}

func (req *reqListArticle) toCmd(user cmprimitive.Username) (cmd articlesvc.ArticleListCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()

	cmd = articlesvc.ArticleListCmd{
		PaginationCmd: listCmd,
		User:          user,
	}

	return cmd, cmd.Validate()
}
