/*
 * Copyright 2019-2020 VMware, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package api

import (
	"github.com/FederatedAI/KubeFATE/k8s-deploy/pkg/modules"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Namespace struct {
}

func (j *Namespace) Router(r *gin.RouterGroup) {

	authMiddleware, _ := GetAuthMiddleware()
	job := r.Group("/namespace")
	job.Use(authMiddleware.MiddlewareFunc())
	{
		job.GET("/", j.getNamespaceList)
	}
}

func (_ *Namespace) getNamespaceList(c *gin.Context) {

	namespace := new(modules.Namespace)
	namespaceList, err := namespace.GetList()
	if err != nil {
		log.Error().Err(err).Msg("request error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Debug().Interface("data", namespaceList).Msg("getNamespaceList success")
	c.JSON(200, gin.H{"data": namespaceList, "msg": "getNamespaceList success"})
}
