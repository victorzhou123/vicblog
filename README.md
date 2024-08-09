# The Personal Blog, VicBlog
VicBlog sharing IT technology

## Transaction
about transaction of project: 
* if you want begin a transaction operation
    * new a transaction instance by `NewTransaction`
    * inject this transaction instance as dependency into all Repo builder you want operate under one transaction
    * carefull!
	    * use `tx.Begin` in first repository implement function
        * use `tx.Commit` in last function

code: new and inject dependency
```go
// new transaction instance
txAddArticle := cminframysql.NewTransaction()

// dependency injection
articleRepo := articlerepoimpl.NewArticleRepo(mysqlImpl, txAddArticle)

categoryArticleRepo := articlerepoimpl.NewCategoryArticleRepo(txAddArticle)

tagArticleRepo := articlerepoimpl.NewTagArticleRepo(txAddArticle)

// app service
articleAppService := articleappsvc.NewArticleAggService(articleService, categoryService, tagService)

```

code: things you should do while use transaction instance
```go
func (impl *articleRepoImpl) AddArticle(info *entity.ArticleInfo) (uint, error) {
    // ingnore unrelated code... 

	// transaction begin
	impl.tx.Begin()

	if err := impl.tx.Insert(&ArticleDO{}, &do); err != nil {
		return 0, err
	}

	return do.ID, nil
}


func (impl *tagArticleImpl) AddRelateWithArticle(
	articleId cmprimitive.Id, tagIds []cmprimitive.Id,
) error {
    // ingnore unrelated code... 

	return impl.tx.Insert(&TagArticleDO{}, &dos) // nothing to do here
}


func (impl *categoryArticleImpl) BindCategoryAndArticle(articleId, cateId cmprimitive.Id) error {
    // ingnore unrelated code... 

	if err := impl.tx.Insert(&CategoryArticleDO{}, &do); err != nil {
		return err
	}

	// transaction commit
	impl.tx.Commit()

	return nil
}


```
