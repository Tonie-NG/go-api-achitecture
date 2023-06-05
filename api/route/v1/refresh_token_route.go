package route

import (
	"time"

	"github.com/Tonie-NG/go-api-achitecture/api/controller"
	"github.com/Tonie-NG/go-api-achitecture/bootstrap"
	"github.com/Tonie-NG/go-api-achitecture/domain"
	"github.com/Tonie-NG/go-api-achitecture/mongo"
	"github.com/Tonie-NG/go-api-achitecture/repository"
	"github.com/Tonie-NG/go-api-achitecture/usecase"
	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
