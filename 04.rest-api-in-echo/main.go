package main

import (
	"crypto/subtle"
	"github.com/asdine/storm/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"rest-api-in-echo/cache"
	"rest-api-in-echo/user"

	"strings"
)

func serveCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if cache.Serve(c.Response(), c.Request()) {
			return nil
		}
		return next(c)
	}
}

func cacheResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Writer = cache.NewWriter(c.Response().Writer, c.Request())
		return next(c)
	}
}

type jsonResponse map[string]interface{}

func usersOptions(c echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	c.Response().Header().Set("Allow", strings.Join(methods, ","))
	return c.NoContent(http.StatusOK)
}

func userOptions(c echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	c.Response().Header().Set("Allow", strings.Join(methods, ","))
	return c.NoContent(http.StatusOK)
}

func usersGetAll(c echo.Context) error {
	users, err := user.All()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if c.Request().Method == http.MethodHead {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(http.StatusOK, jsonResponse{"users": users})
}

func usersPostOne(c echo.Context) error {
	u := new(user.User)
	err := c.Bind(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	u.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil {
		if err == user.ErrorRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}
	cache.Drop("/users")
	c.Response().Header().Set("Location", "/users/"+u.ID.Hex())
	c.NoContent(http.StatusCreated)
	return nil
}

func usersGetOne(c echo.Context) error {
	if cache.Serve(c.Response(), c.Request()) {
		return nil
	}
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	id := bson.ObjectIdHex(c.Param("id"))
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if c.Request().Method == http.MethodHead {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(http.StatusOK, jsonResponse{"user": u})
}

func usersPutOne(c echo.Context) error {
	u := new(user.User)
	err := c.Bind(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	id := bson.ObjectIdHex(c.Param("id"))
	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrorRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}
	cache.Drop("/users")
	return c.JSON(http.StatusOK, jsonResponse{"user": u})
}

func usersPatchOne(c echo.Context) error {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	id := bson.ObjectIdHex(c.Param("id"))
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = c.Bind(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrorRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}
	cache.Drop("/users")
	return c.JSON(http.StatusOK, jsonResponse{"user": u})
}

func usersDeleteOne(c echo.Context) error {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	id := bson.ObjectIdHex(c.Param("id"))
	err := user.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/users")
	cache.Drop(cache.MakeResource(c.Request()))
	c.NoContent(http.StatusOK)
	return nil
}

func auth(username, password string, c echo.Context) (bool, error) {
	// Be careful to use constant time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
		return true, nil
	}
	return false, nil
}

func root(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "API is running v1")
}

func main() {
	e := echo.New()
	// Middleware section
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	// do a user readable logging using logger middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, ${uri}, ${status} ${latency_human}\n",
	}))
	// Routes and handlers section
	e.GET("/", root)

	u := e.Group("/users")
	u.OPTIONS("", usersOptions)
	u.HEAD("", usersGetAll, serveCache)
	u.GET("", usersGetAll, serveCache, cacheResponse)
	u.POST("", usersPostOne, middleware.BasicAuth(auth))

	uid := u.Group("/:id")

	uid.OPTIONS("", userOptions)
	uid.HEAD("", usersGetOne, serveCache)
	uid.GET("", usersGetOne, serveCache, cacheResponse)
	uid.PUT("", usersPutOne, middleware.BasicAuth(auth), cacheResponse)
	uid.PATCH("", usersPatchOne, middleware.BasicAuth(auth), cacheResponse)
	uid.DELETE("", usersDeleteOne, middleware.BasicAuth(auth))

	e.Logger.Fatal(e.Start("127.0.0.1:12345"))
}
