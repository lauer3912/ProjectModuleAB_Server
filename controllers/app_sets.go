package controllers

import (
	"encoding/json"
	"fmt"
	"moduleab_server/common"
	"moduleab_server/models"
	"net/http"

	"github.com/astaxie/beego"
)

type AppSetsController struct {
	beego.Controller
}

// @Title createAppSet
// @router / [post]
func (a *AppSetsController) Post() {
	if a.Ctx.Input.Header("Signature") != "" {
		err := common.AuthWithKey(a.Ctx)
		if err != nil {
			a.Data["json"] = map[string]string{
				"error": err.Error(),
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	} else {
		if a.GetSession("id") == nil {
			a.Data["json"] = map[string]string{
				"error": "You need login first.",
			}
			a.Ctx.Output.SetStatus(http.StatusUnauthorized)
			a.ServeJSON()
		}
		if models.CheckPrivileges(
			a.GetSession("id").(string),
			models.RoleFlagOperator,
		) {
			a.Data["json"] = map[string]string{
				"error": "No privilege",
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	}

	appSet := new(models.AppSets)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, appSet)
	if err != nil {
		beego.Warn("[C] Got error:", err)
		a.Data["json"] = map[string]string{
			"message": "Bad request",
			"error":   err.Error(),
		}
		a.Ctx.Output.SetStatus(http.StatusBadRequest)
		a.ServeJSON()
		return
	}
	beego.Debug("[C] Got data:", appSet)
	id, err := models.AddAppSet(appSet)
	if err != nil {
		beego.Warn("[C] Got error:", err)
		a.Data["json"] = map[string]string{
			"message": "Failed to add New host",
			"error":   err.Error(),
		}
		a.Ctx.Output.SetStatus(http.StatusInternalServerError)
		a.ServeJSON()
		return
	}

	beego.Debug("[C] Got id:", id)
	a.Data["json"] = map[string]string{
		"id": id,
	}
	a.Ctx.Output.SetStatus(http.StatusCreated)
	a.ServeJSON()
	return
}

// @Title getAppSet
// @router /:name [get]
func (a *AppSetsController) Get() {
	if a.Ctx.Input.Header("Signature") != "" {
		err := common.AuthWithKey(a.Ctx)
		if err != nil {
			a.Data["json"] = map[string]string{
				"error": err.Error(),
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	} else {
		if a.GetSession("id") == nil {
			a.Data["json"] = map[string]string{
				"error": "You need login first.",
			}
			a.Ctx.Output.SetStatus(http.StatusUnauthorized)
			a.ServeJSON()
		}
		if models.CheckPrivileges(
			a.GetSession("id").(string),
			models.RoleFlagOperator,
		) {
			a.Data["json"] = map[string]string{
				"error": "No privilege",
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	}

	name := a.GetString(":name")
	beego.Debug("[C] Got name:", name)
	if name != "" {
		appSet := &models.AppSets{
			Name: name,
		}
		appSets, err := models.GetAppSets(appSet, 0, 0)
		if err != nil {
			a.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to get  with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			a.Ctx.Output.SetStatus(http.StatusInternalServerError)
			a.ServeJSON()
			return
		}
		a.Data["json"] = appSets
		if len(appSets) == 0 {
			beego.Debug("[C] Got nothing with name:", name)
			a.Ctx.Output.SetStatus(http.StatusNotFound)
			a.ServeJSON()
			return
		} else {
			a.Ctx.Output.SetStatus(http.StatusOK)
			a.ServeJSON()
			return
		}
	}
}

// @Title listAppSets
// @router / [get]
func (a *AppSetsController) GetAll() {
	if a.Ctx.Input.Header("Signature") != "" {
		err := common.AuthWithKey(a.Ctx)
		if err != nil {
			a.Data["json"] = map[string]string{
				"error": err.Error(),
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	} else {
		if a.GetSession("id") == nil {
			a.Data["json"] = map[string]string{
				"error": "You need login first.",
			}
			a.Ctx.Output.SetStatus(http.StatusUnauthorized)
			a.ServeJSON()
		}
		if models.CheckPrivileges(
			a.GetSession("id").(string),
			models.RoleFlagOperator,
		) {
			a.Data["json"] = map[string]string{
				"error": "No privilege",
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	}

	limit, _ := a.GetInt("limit", 0)
	index, _ := a.GetInt("index", 0)

	appSet := &models.AppSets{}
	appSets, err := models.GetAppSets(appSet, limit, index)
	if err != nil {
		a.Data["json"] = map[string]string{
			"message": fmt.Sprint("Failed to get"),
			"error":   err.Error(),
		}
		beego.Warn("[C] Got error:", err)
		a.Ctx.Output.SetStatus(http.StatusInternalServerError)
		a.ServeJSON()
		return
	}
	a.Data["json"] = appSets
	if len(appSets) == 0 {
		beego.Debug("[C] Got nothing")
		a.Ctx.Output.SetStatus(http.StatusNotFound)
		a.ServeJSON()
		return
	} else {
		a.Ctx.Output.SetStatus(http.StatusOK)
		a.ServeJSON()
		return
	}
}

// @Title deleteAppSet
// @router /:name [delete]
func (a *AppSetsController) Delete() {
	if a.Ctx.Input.Header("Signature") != "" {
		err := common.AuthWithKey(a.Ctx)
		if err != nil {
			a.Data["json"] = map[string]string{
				"error": err.Error(),
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	} else {
		if a.GetSession("id") == nil {
			a.Data["json"] = map[string]string{
				"error": "You need login first.",
			}
			a.Ctx.Output.SetStatus(http.StatusUnauthorized)
			a.ServeJSON()
		}
		if models.CheckPrivileges(
			a.GetSession("id").(string),
			models.RoleFlagOperator,
		) {
			a.Data["json"] = map[string]string{
				"error": "No privilege",
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	}

	name := a.GetString(":name")
	beego.Debug("[C] Got name:", name)
	if name != "" {
		appSet := &models.AppSets{
			Name: name,
		}
		appSets, err := models.GetAppSets(appSet, 0, 0)
		if err != nil {
			a.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to get with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			a.Ctx.Output.SetStatus(http.StatusInternalServerError)
			a.ServeJSON()
			return
		}
		if len(appSets) == 0 {
			beego.Debug("[C] Got nothing with name:", name)
			a.Ctx.Output.SetStatus(http.StatusNotFound)
			a.ServeJSON()
			return
		}
		err = models.DeleteAppSet(appSets[0])
		if err != nil {
			a.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to delete with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			a.Ctx.Output.SetStatus(http.StatusInternalServerError)
			a.ServeJSON()
			return

		}
		a.Ctx.Output.SetStatus(http.StatusNoContent)
		a.ServeJSON()
		return
	}
}

// @Title updateAppSet
// @router /:name [put]
func (a *AppSetsController) Put() {
	if a.Ctx.Input.Header("Signature") != "" {
		err := common.AuthWithKey(a.Ctx)
		if err != nil {
			a.Data["json"] = map[string]string{
				"error": err.Error(),
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	} else {
		if a.GetSession("id") == nil {
			a.Data["json"] = map[string]string{
				"error": "You need login first.",
			}
			a.Ctx.Output.SetStatus(http.StatusUnauthorized)
			a.ServeJSON()
		}
		if models.CheckPrivileges(
			a.GetSession("id").(string),
			models.RoleFlagOperator,
		) {
			a.Data["json"] = map[string]string{
				"error": "No privilege",
			}
			a.Ctx.Output.SetStatus(http.StatusForbidden)
			a.ServeJSON()
		}
	}

	name := a.GetString(":name")
	beego.Debug("[C] Got appSet name:", name)
	if name != "" {
		appSet := &models.AppSets{
			Name: name,
		}
		appSets, err := models.GetAppSets(appSet, 0, 0)
		if err != nil {
			a.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to get with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			a.Ctx.Output.SetStatus(http.StatusInternalServerError)
			a.ServeJSON()
			return
		}
		if len(appSets) == 0 {
			beego.Debug("[C] Got nothing with name:", name)
			a.Ctx.Output.SetStatus(http.StatusNotFound)
			a.ServeJSON()
			return
		}

		err = json.Unmarshal(a.Ctx.Input.RequestBody, appSet)
		appSet.Id = appSets[0].Id
		if err != nil {
			beego.Warn("[C] Got error:", err)
			a.Data["json"] = map[string]string{
				"message": "Bad request",
				"error":   err.Error(),
			}
			a.Ctx.Output.SetStatus(http.StatusBadRequest)
			a.ServeJSON()
			return
		}
		beego.Debug("[C] Got appSet data:", appSet)
		err = models.UpdateAppSet(appSet)
		if err != nil {
			a.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to update with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			a.Ctx.Output.SetStatus(http.StatusInternalServerError)
			a.ServeJSON()
			return

		}
		a.Ctx.Output.SetStatus(http.StatusAccepted)
		a.ServeJSON()
		return
	}
}
