package main 

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
    team := c.QueryParam("team")
    member := c.QueryParam("member")
    return c.String(http.StatusOK, "team: " + team + ", member: " + member)
}

func save(c echo.Context) error {
    name := c.FormValue("name")
    email := c.FormValue("email")

    return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func main(){
    e := echo.New()


    // path parameters 
    e.GET("/users/:id", getUser)

    // query parameters
    e.GET("/show", show)

    // post
    e.POST("/save", save)

    e.GET("/", func (c echo.Context) error {
        return c.String(http.StatusOK, "Hello world")
    })

    e.Logger.Fatal(e.Start(":1323"))


}




