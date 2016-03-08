Collision2D
======

[![Build Status](https://travis-ci.org/Tarliton/collision2d.svg?branch=master)](https://travis-ci.org/Tarliton/collision2d)
[![GoDoc](https://godoc.org/github.com/Tarliton/collision2d?status.svg)](https://godoc.org/github.com/Tarliton/collision2d)
[![Coverage Status](https://coveralls.io/repos/github/Tarliton/collision2d/badge.svg?branch=master)](https://coveralls.io/github/Tarliton/collision2d?branch=master)

About
-----

Collision2D is a simple Go library for performing collision detection (and projection-based collision response) of simple 2D shapes.  It uses the [Separating Axis Theorem](http://en.wikipedia.org/wiki/Hyperplane_separation_theorem). It was based on the JavaScript library [SAT-js](https://github.com/jriecken/sat-js) by [jriecken](https://github.com/jriecken).

It supports detecting collisions between:
 - Circles (using Voronoi Regions.)
 - Convex Polygons (and simple Axis-Aligned Boxes, which are of course, convex polygons.)

It also supports checking whether a point is inside a circle or polygon.

It's released under the [MIT](http://en.wikipedia.org/wiki/MIT_License) license.

To import in your Go project you can use:

    `import "github.com/Tarliton/collision2d"`