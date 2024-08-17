package controller

import (
	"github.com/victorzhou123/vicblog/article/app/dto"
	articleent "github.com/victorzhou123/vicblog/article/domain/article/entity"
	"github.com/victorzhou123/vicblog/article/domain/category/entity"
	tagent "github.com/victorzhou123/vicblog/article/domain/tag/entity"
	cmctl "github.com/victorzhou123/vicblog/common/controller"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
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

func (req *reqCategoryList) toCmd() (cmd dto.ListCategoryCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()
	if err != nil {
		return
	}

	cmd = dto.ListCategoryCmd{
		PaginationCmd: listCmd,
	}

	return
}

type reqTag struct {
	Names []string `json:"names"`
}

func (req *reqTag) toTagNames() ([]tagent.TagName, error) {
	tagNames := make([]tagent.TagName, len(req.Names))

	var err error
	for i := range req.Names {
		tagNames[i], err = tagent.NewTagName(req.Names[i])
		if err != nil {
			return nil, err
		}
	}

	return tagNames, nil
}

type reqTagList struct {
	cmctl.ReqList
}

func (req *reqTagList) emptyValue() bool {
	return req.ReqList.EmptyValue()
}

func (req *reqTagList) toCmd() (cmd dto.ListTagCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()
	if err != nil {
		return
	}

	cmd = dto.ListTagCmd{
		PaginationCmd: listCmd,
	}

	return
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

func (req *reqListArticle) toCmd(user cmprimitive.Username) (cmd dto.GetArticleListCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()
	if err != nil {
		return
	}

	return dto.GetArticleListCmd{
		PaginationCmd: listCmd,
		User:          user,
	}, nil
}

type reqListAllArticle struct {
	cmctl.ReqList
}

func (req *reqListAllArticle) toCmd() (cmd dto.GetArticleListCmd, err error) {

	listCmd, err := req.ReqList.ToCmd()
	if err != nil {
		return
	}

	cmd = dto.GetArticleListCmd{
		PaginationCmd: listCmd,
	}

	return
}

type reqUpdateArticle struct {
	reqArticle

	Id uint `json:"id"`
}

func (req *reqUpdateArticle) toCmd(user cmprimitive.Username) (cmd dto.UpdateArticleCmd, err error) {

	if cmd.AddArticleCmd, err = req.reqArticle.toCmd(user); err != nil {
		return
	}

	cmd.Id = cmprimitive.NewIdByUint(req.Id)

	return
}

// list article cards
type reqListArticleCards struct {
	cmctl.ReqList

	CategoryId string `json:"categoryId"`
}

func (req *reqListArticleCards) toCmd() (cmd dto.GetArticleCardListByCateIdCmd, err error) {

	if cmd.PaginationCmd, err = req.ReqList.ToCmd(); err != nil {
		return
	}

	cmd.CategoryId = cmprimitive.NewId(req.CategoryId)

	return
}
