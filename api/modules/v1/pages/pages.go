package pages

import (
	"fmt"
	"net/http"
	"okapi-public-api/lib/aws"
	"okapi-public-api/lib/env"
	"okapi-public-api/lib/redis"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/protsack-stephan/dev-toolkit/lib/s3"
	"github.com/protsack-stephan/gin-toolkit/httpmod"
	"github.com/protsack-stephan/gin-toolkit/httpmw"
)

// Init for page endpoints
func Init() httpmod.Module {
	expire := time.Second * time.Duration(env.PagesExpire)
	store := s3.NewStorage(aws.Session(), env.AWSBucket)
	cmdable := redis.Client()

	return httpmod.Module{
		Path:       "/v1/pages",
		Middleware: []gin.HandlerFunc{},
		Routes: []httpmod.Route{
			{
				Path:   "/meta/:project/*name",
				Method: http.MethodGet,
				Handler: httpmw.Cache(&httpmw.CacheParams{
					Cache:       cmdable,
					Expire:      expire,
					Handle:      Meta(store),
					ContentType: fmt.Sprintf("%s; charset=UTF-8", gin.MIMEJSON),
				}),
				Middleware: []gin.HandlerFunc{
					httpmw.LimitPerUser(cmdable, env.GroupLimit, env.Group, 0, env.Group),
				},
			},
		},
	}
}
