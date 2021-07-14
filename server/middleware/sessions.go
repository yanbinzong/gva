package middleware
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/context"
//)
//
//func Sessioins(name string, store Store ) gin.HandlerFunc  {
//	return func(c *gin.Context) {
//		s := &session{name, c.Request, Store, nil, false, c.Writer}
//		c.Set(DefaultKey,s)
//		defer context.Clear(c.Rquest)
//		c.Next()
//	}
//}