// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the entrance function for starting the web server.

package main

import (
	_ "ClaytonUniversityLibrary/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
