package main

type GoStreamPlugin interface {
	Setup(g *GoStream, r *Router) error
}
