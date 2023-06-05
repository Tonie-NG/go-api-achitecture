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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
