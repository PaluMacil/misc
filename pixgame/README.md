# Pixgame

This repo is an expiriment for playing with [Pixel](https://github.com/faiface/pixel), a 2D game engine in Go. It supports lots of graphic effects, a game loop, font rendering, optimized sprites and more. If I wind up settling on a genre, graphic style, and storyline, I will probably make a dedicated repo instead of continuing to work on this sub-repo.

# Compile

You need GCC installed and in your path (check in a terminal with `gcc -v`) for the OpenGL bindings. If you don't have this, [use these instructions](./docs/gcc.md). Otherwise, simply run:

```
go get -u github.com/go-gl/glfw/v3.2/glfw
go get -u github.com/palumacil/misc/pixgame...
```

[&#x2190; Back to Project List](../README.md)