package main

type GoStreamPlugin interface {
	Setup(gost *GoStream, r *Router) error
}
