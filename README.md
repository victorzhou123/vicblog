# The Personal Blog, VicBlog
VicBlog sharing IT technology

## Architecture Diagram
[![architecture diagram](https://i0.hdslb.com/bfs/article/25109453a74462b32b1f9e25702bdff6391490864.png)](https://github.com/victorzhou123/vicblog)

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
transactionImpl := cminframysql.NewTransaction()

// dependency injection
articleRepo := articlerepoimpl.NewArticleRepo(mysqlImpl, transactionImpl)

categoryArticleRepo := articlerepoimpl.NewCategoryArticleRepo(transactionImpl)

tagArticleRepo := articlerepoimpl.NewTagArticleRepo(transactionImpl)

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

## Development List

| task                                    | done | note                             |
| ----------------------------------------- | :----: | ---------------------------------- |
| kuberentes cluster                      |  √  |                                  |
| vicblog                                 |  √  |                                  |
| micro service framework                 |  √  |                                  |
| micro service (user, category, tag)     |  √  |                                  |
| mysql master-slave replication          |  √  |                                  |
| cyclic mysql data backup                |  √  | backup nfs data of mysql         |
| GPU scheduler of kubernetes             |  √  | allocate GPU power in kubernetes |
| deepseek in kubernetes                  |  √  | based ollama                     |
| kafka cluster                           |      |                                  |
| upgrade of sensitive-world-check-server |      |                                  |
| replace oss with own object storage     |      |                                  |
| own message queue implement             |      |                                  |