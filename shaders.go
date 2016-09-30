package main

var (
	vertexShader = `
#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 tex;

out vec2 fragTex;

void main() {
	fragTex = tex;
	gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00"

	fragmentShader = `
#version 330

uniform sampler2D tex;
in vec2 texCoord;

out vec4 outputColor;

void main() {
	// ignore tex for now
	// outputColor = texture(tex, texCoord);
	outputColor = vec4(1.0, 1.0, 0.0, 1.0);
}
` + "\x00"
)
