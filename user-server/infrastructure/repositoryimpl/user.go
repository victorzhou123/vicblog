package repositoryimpl

import (
	cmerror "github.com/victorzhou123/vicblog/common/domain/error"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/user-server/domain"
	"github.com/victorzhou123/vicblog/user-server/domain/repository"
)

func NewUserRepo(db mysql.Impl) repository.User {

	if err := mysql.AutoMigrate(&UserDO{}); err != nil {
		return nil
	}

	return &userRepoImpl{db}
}

type userRepoImpl struct {
	mysql.Impl
}

func (impl *userRepoImpl) GetUserInfo(account *repository.Account) (domain.User, error) {
	userDo := &UserDO{}
	userDo.Username = account.Username.Username()
	userDo.Password = account.Password.Password()

	if err := impl.GetRecord(&UserDO{}, &userDo, &userDo); err != nil {
		if cmerror.IsNotFound(err) {
			return domain.User{}, err
		}

		return domain.User{}, cmerror.New(cmerror.ErrorCodeInternalError, err.Error())
	}

	return userDo.toUser()
}
