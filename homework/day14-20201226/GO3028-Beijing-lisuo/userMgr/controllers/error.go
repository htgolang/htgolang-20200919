package controllers

// HandleError wrap err handle code
func HandleError(c *UserController, err error) {
	r := c.Ctx.Request.Referer()
	if err != nil {
		c.Data["msg"] = err.Error()
		c.Data["refer"] = r
		c.TplName = "error/error.html"
		return
	}
}

// HandleError wrap err handle code
func HandleAuthError(c *AuthController, err error) {
	r := c.Ctx.Request.Referer()
	if err != nil {
		c.Data["msg"] = err.Error()
		c.Data["refer"] = r
		c.TplName = "error/error.html"
		return
	}
}
