/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/2/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package main

import "fmt"

type Obj struct {
	a int
}

func addObj(objs []*Obj) {
	objs = append(objs, &Obj{a:1})
}

func main() {

	objs := []*Obj{}

	addObj(objs)
	addObj(objs)
	fmt.Println(len(objs))

}
