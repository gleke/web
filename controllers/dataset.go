// Copyright 2016 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package controllers

import (
	"net/http"

	"github.com/gleke/hexya/src/actions"
	"github.com/gleke/hexya/src/server"
)

// CallKW executes the given method of the given model
func CallKW(c *server.Context) {
	uid := c.Session().Get("uid").(int64)
	var params CallParams
	c.BindRPCParams(&params)
	res, err := Execute(uid, params)
	c.RPC(http.StatusOK, res, err)
}

// CallButton executes the given method of the given model
// and returns the result only if it is an action
func CallButton(c *server.Context) {
	uid := c.Session().Get("uid").(int64)
	var params CallParams
	c.BindRPCParams(&params)
	res, err := Execute(uid, params)
	switch act := res.(type) {
	case actions.Action:
		log.Panic("Call button functions should return a pointer to action", "params", params, "received", "action")
	case *actions.Action:
		act.Sanitize()
		c.RPC(http.StatusOK, act, err)
	default:
		c.RPC(http.StatusOK, false, err)
	}
}

// SearchReadController returns Records from the database
func SearchReadController(c *server.Context) {
	uid := c.Session().Get("uid").(int64)
	var params SearchReadParams
	c.BindRPCParams(&params)
	res, err := SearchRead(uid, params)
	c.RPC(http.StatusOK, res, err)
}
