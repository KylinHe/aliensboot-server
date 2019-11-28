/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2019/11/6
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package util

type String2Int64 map[string][]int64

func (this String2Int64) Put(key string, value int64) {
	data := this[key]
	if data == nil {
		data = []int64{}
	}
	data = append(data, value)
	this[key] = data
}
