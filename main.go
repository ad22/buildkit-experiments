package main

import ("fmt"
        bkclient "github.com/moby/buildkit/client")

func main() {
	fmt.Println("moby/buildkit experiments")
	fmt.Println(bkclient.ExporterTar)
}