package controllers

import (
	"encoding/json"
	"fmt"
	"moduleab_server/models"
	"net/http"

	"github.com/astaxie/beego"
)

type HostsController struct {
	beego.Controller
}

// @Title registerHost
// @Description create Host
// @Param	body		body 	models.Hosts	true		"body for host content"
// @Success 201 {string} models.Hosts.Id
// @Failure 403 body is empty
// @router / [post]
func (h *HostsController) Post() {
	host := new(models.Hosts)
	err := json.Unmarshal(h.Ctx.Input.RequestBody, host)
	if err != nil {
		beego.Warn("[C] Got error:", err)
		h.Data["json"] = map[string]string{
			"message": "Bad request",
			"error":   err.Error(),
		}
		h.Ctx.Output.SetStatus(http.StatusBadRequest)
		h.ServeJSON()
	}
	beego.Debug("[C] Got data:", host)
	id, err := models.AddHost(host)
	if err != nil {
		beego.Warn("[C] Got error:", err)
		h.Data["json"] = map[string]string{
			"message": "Failed to add New host",
			"error":   err.Error(),
		}
		h.Ctx.Output.SetStatus(http.StatusInternalServerError)
		h.ServeJSON()
	}

	beego.Debug("[C] Got id:", id)
	h.Data["json"] = map[string]string{
		"id": id,
	}
	h.Ctx.Output.SetStatus(http.StatusCreated)
	h.ServeJSON()
}

// @Title getHost
// @Description get Host info
// @Param	body		body 	models.Hosts	true		"body for host content"
// @Success 200 {string} models.Hosts.Id
// @Failure 403 body is empty
// @router /:name [get]
func (h *HostsController) Get() {
	name := h.GetString(":name")
	beego.Debug("[C] Got name:", name)
	if name != "" {
		host := &models.Hosts{
			Name: name,
		}
		hosts, err := models.GetHosts(host)
		if err != nil {
			h.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to get with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			h.Ctx.Output.SetStatus(http.StatusInternalServerError)
			h.ServeJSON()
		}
		h.Data["json"] = hosts
		if len(hosts) == 0 {
			beego.Debug("[C] Got nothing with name:", name)
			h.Ctx.Output.SetStatus(http.StatusNotFound)
			h.ServeJSON()
		} else {
			h.Ctx.Output.SetStatus(http.StatusOK)
			h.ServeJSON()
		}
	}
}

// @Title listHosts
// @Description get all Host info
// @Success 200
// @router / [get]
func (h *HostsController) GetAll() {
	host := &models.Hosts{}
	hosts, err := models.GetHosts(host)
	if err != nil {
		h.Data["json"] = map[string]string{
			"message": fmt.Sprint("Failed to get"),
			"error":   err.Error(),
		}
		beego.Warn("[C] Got error:", err)
		h.Ctx.Output.SetStatus(http.StatusInternalServerError)
		h.ServeJSON()
	}
	h.Data["json"] = hosts
	if len(hosts) == 0 {
		beego.Debug("[C] Got nothing")
		h.Ctx.Output.SetStatus(http.StatusNotFound)
		h.ServeJSON()
	} else {
		h.Ctx.Output.SetStatus(http.StatusOK)
		h.ServeJSON()
	}
}

// @Title deleteHost
// @Description delete host
// @Success 204
// @Failure 404
// @router /:name [delete]
func (h *HostsController) Delete() {
	name := h.GetString(":name")
	beego.Debug("[C] Got name:", name)
	if name != "" {
		host := &models.Hosts{
			Name: name,
		}
		hosts, err := models.GetHosts(host)
		if err != nil {
			h.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to get with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			h.Ctx.Output.SetStatus(http.StatusInternalServerError)
			h.ServeJSON()
		}
		if len(hosts) == 0 {
			beego.Debug("[C] Got nothing with name:", name)
			h.Ctx.Output.SetStatus(http.StatusNotFound)
			h.ServeJSON()
		}
		err = models.DeleteHost(hosts[0])
		if err != nil {
			h.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to delete with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			h.Ctx.Output.SetStatus(http.StatusInternalServerError)
			h.ServeJSON()

		}
		h.Ctx.Output.SetStatus(http.StatusNoContent)
		h.ServeJSON()
	}
}

// @Title updateHost
// @Description update host
// @Success 204
// @Failure 404
// @router /:name [put]
func (h *HostsController) Put() {
	name := h.GetString(":name")
	beego.Debug("[C] Got host name:", name)
	if name != "" {
		host := &models.Hosts{
			Name: name,
		}
		hosts, err := models.GetHosts(host)
		if err != nil {
			h.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to get with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			h.Ctx.Output.SetStatus(http.StatusInternalServerError)
			h.ServeJSON()
		}
		if len(hosts) == 0 {
			beego.Debug("[C] Got nothing with name:", name)
			h.Ctx.Output.SetStatus(http.StatusNotFound)
			h.ServeJSON()
		}

		err = json.Unmarshal(h.Ctx.Input.RequestBody, host)
		host.Id = hosts[0].Id
		if err != nil {
			beego.Warn("[C] Got error:", err)
			h.Data["json"] = map[string]string{
				"message": "Bad request",
				"error":   err.Error(),
			}
			h.Ctx.Output.SetStatus(http.StatusBadRequest)
			h.ServeJSON()
		}
		beego.Debug("[C] Got host data:", host)
		err = models.UpdateHost(host)
		if err != nil {
			h.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to update with name:", name),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			h.Ctx.Output.SetStatus(http.StatusInternalServerError)
			h.ServeJSON()

		}
		h.Ctx.Output.SetStatus(http.StatusAccepted)
		h.ServeJSON()
	}
}