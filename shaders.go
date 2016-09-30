package main

var (
	vertexShader = `
#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;

void main() {
	gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00"

	fragmentShader = `
#version 330

uniform vec4 col;

out vec4 outputColor;

void main() {
	outputColor = col;
}
` + "\x00"
)
