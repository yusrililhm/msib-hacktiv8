package main

import "belajar-gin/routers"

func main() {
	routers.Router().Run(":3000")
}
