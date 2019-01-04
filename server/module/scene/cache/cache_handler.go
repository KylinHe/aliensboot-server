/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import (
	"github.com/KylinHe/aliensboot-core/cluster/center"
)

const (
	spaceNodeKey = "space:"
	entityNodeKey = "entity:"
)

//设置空间所在的服务节点信息
func (this *cacheManager) SetSpaceNode(spaceID string) error {
	return this.redisClient.SetData(spaceNodeKey + spaceID, center.ClusterCenter.GetNodeID())
}

//获取空间所在的服务节点信息
func (this *cacheManager) GetSpaceNode(spaceID string) (string, error) {
	return this.redisClient.GetData(spaceNodeKey + spaceID)
}

//设置entity所在的服务节点信息
func (this *cacheManager) SetEntityNode(entityID string) error {
	return this.redisClient.SetData(entityNodeKey + entityID, center.ClusterCenter.GetNodeID())
}

//获取entity所在的服务节点信息
func (this *cacheManager) GetEntityNode(entityID string) (string, error) {
	return this.redisClient.GetData(entityNodeKey + entityID)
}