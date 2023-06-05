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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
