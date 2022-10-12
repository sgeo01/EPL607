package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	gl "github.com/chsc/gogl/gl33"
	"github.com/veandco/go-sdl2/sdl"
)

func createprogram() gl.Uint {
	// VERTEX SHADER
	vs := gl.CreateShader(gl.VERTEX_SHADER)
	vs_source := gl.GLString(vertexShaderSource)
	gl.ShaderSource(vs, 1, &vs_source, nil)
	gl.CompileShader(vs)
	var vs_status gl.Int
	gl.GetShaderiv(vs, gl.COMPILE_STATUS, &vs_status)
	fmt.Printf("Compiled Vertex Shader: %v\n", vs_status)

	// FRAGMENT SHADER
	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	fs_source := gl.GLString(fragmentShaderSource)
	gl.ShaderSource(fs, 1, &fs_source, nil)
	gl.CompileShader(fs)
	var fstatus gl.Int
	gl.GetShaderiv(fs, gl.COMPILE_STATUS, &fstatus)
	fmt.Printf("Compiled Fragment Shader: %v\n", fstatus)

	// CREATE PROGRAM
	program := gl.CreateProgram()
	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)
	fragoutstring := gl.GLString("outColor")
	defer gl.GLStringFree(fragoutstring)
	gl.BindFragDataLocation(program, gl.Uint(0), fragoutstring)

	gl.LinkProgram(program)
	var linkstatus gl.Int
	gl.GetProgramiv(program, gl.LINK_STATUS, &linkstatus)
	fmt.Printf("Program Link: %v\n", linkstatus)

	return program
}

var uniRoll float32 = 0.0
var uniYaw float32 = 1.0
var uniPitch float32 = 0.0
var uniscale float32 = 0.3
var yrot float32 = 20.0
var zrot float32 = 0.0
var xrot float32 = 0.0
var UniScale gl.Int

var f = []int{}
var v = []float64{}
var triangle_vertices = []gl.Float{}
var triangle_colours = []gl.Float{}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	var window *sdl.Window
	var context sdl.GLContext
	var event sdl.Event
	var running bool
	var err error

	// open file for reading
	// read line by line
	lines, err := readLines(`C:\Users\Solonas\Desktop\EPL 607\file.obj`)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	id := []int{}
	for i, line := range lines {
		//read := strings.Split(line, "")
		read := strings.Fields(line)

		if line == "" {
			fmt.Println("ERROR")
			read = append(read, "\n")
		}
		//fmt.Println(line)

		if read[0] == "f" {
			fline1 := strings.Split(read[1], "/")
			fline2 := strings.Split(read[2], "/")
			fline3 := strings.Split(read[3], "/")
			intfline1, d1 := strconv.Atoi(fline1[0])
			intfline2, d2 := strconv.Atoi(fline2[0])
			intfline3, d3 := strconv.Atoi(fline3[0])
			f = append(f, intfline1, intfline2, intfline3)

			/*fmt.Println(i, intvline1)
			fmt.Println(i, intvline2)
			fmt.Println(i, intvline3)*/
			if d1 = sdl.Init(sdl.INIT_EVERYTHING); d1 != nil {
				panic(d1)
			}
			if d2 = sdl.Init(sdl.INIT_EVERYTHING); d2 != nil {
				panic(d2)
			}
			if d3 = sdl.Init(sdl.INIT_EVERYTHING); d3 != nil {
				panic(d3)
			}
		} else if read[0] == "v" {
			vline1 := strings.Fields(read[1])
			vline2 := strings.Fields(read[2])
			vline3 := strings.Fields(read[3])
			if floatvline1, err := strconv.ParseFloat(vline1[0], 64); err == nil {
				v = append(v, floatvline1)
			}
			if floatvline2, err := strconv.ParseFloat(vline2[0], 64); err == nil {
				v = append(v, floatvline2)
			}
			if floatvline3, err := strconv.ParseFloat(vline3[0], 64); err == nil {
				v = append(v, floatvline3)
			}

		} else {
			//fmt.Println(read)
		}

		id = append(id, i)

	}

	for i := 0; i < len(f); i++ {
		num1 := f[i]*3 - 3
		num2 := f[i]*3 - 2
		num3 := f[i]*3 - 1

		t1 := gl.Float(v[num1])
		t2 := gl.Float(v[num2])
		t3 := gl.Float(v[num3])
		random1 := gl.Float(rand.Float64())
		random2 := gl.Float(rand.Float64())
		random3 := gl.Float(rand.Float64())
		triangle_vertices = append(triangle_vertices, t1, t2, t3)
		triangle_colours = append(triangle_colours, random1, random2, random3)

	}

	runtime.LockOSThread()
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	gl.Init()
	gl.Viewport(0, 0, gl.Sizei(winWidth), gl.Sizei(winHeight))
	// OPENGL FLAGS
	gl.ClearColor(0.0, 0.1, 0.0, 1.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// VERTEX BUFFER
	var vertexbuffer gl.Uint
	gl.GenBuffers(1, &vertexbuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexbuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(triangle_vertices)*4), gl.Pointer(&triangle_vertices[0]), gl.STATIC_DRAW)

	// COLOUR BUFFER
	var colourbuffer gl.Uint
	gl.GenBuffers(1, &colourbuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, colourbuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(triangle_colours)*4), gl.Pointer(&triangle_colours[0]), gl.STATIC_DRAW)

	// GUESS WHAT
	program := createprogram()

	// VERTEX ARRAY
	var VertexArrayID gl.Uint
	gl.GenVertexArrays(1, &VertexArrayID)
	gl.BindVertexArray(VertexArrayID)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexbuffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, gl.FALSE, 0, nil)

	// VERTEX ARRAY HOOK COLOURS
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, colourbuffer)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, gl.FALSE, 0, nil)

	//UNIFORM HOOK
	unistring := gl.GLString("scaleMove")
	UniScale = gl.GetUniformLocation(program, unistring)
	fmt.Printf("Uniform Link: %v\n", UniScale+1)

	gl.UseProgram(program)

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event =
			sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:

				xrot = float32(t.Y) / 2
				yrot = float32(t.X) / 2
				fmt.Printf("[%dms]MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			}
		}
		drawgl()
		window.GLSwap()
	}

}

func drawgl() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	uniYaw = yrot * (math.Pi / 180.0)
	yrot = yrot - 1.0
	uniPitch = zrot * (math.Pi / 180.0)
	zrot = zrot - 0.5
	uniRoll = xrot * (math.Pi / 180.0)
	xrot = xrot - 0.2

	gl.Uniform4f(UniScale, gl.Float(uniRoll), gl.Float(uniYaw), gl.Float(uniPitch), gl.Float(uniscale))
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.DrawArrays(gl.TRIANGLES, gl.Int(0), gl.Sizei(len(triangle_vertices)*4))

	time.Sleep(50 * time.Millisecond)

}

const (
	winTitle           = "OpenGL Shader"
	winWidth           = 640
	winHeight          = 480
	vertexShaderSource = `
#version 330
layout (location = 0) in vec3 Position;
layout(location = 1) in vec3 vertexColor;
uniform vec4 scaleMove;
out vec3 fragmentColor;
void main()
{ 
// YOU CAN OPTIMISE OUT cos(scaleMove.x) AND sin(scaleMove.y) AND UNIFORM THE VALUES IN
    vec3 scale = Position.xyz * scaleMove.w;
// rotate on z pole
   vec3 rotatez = vec3((scale.x * cos(scaleMove.x) - scale.y * sin(scaleMove.x)), (scale.x * sin(scaleMove.x) + scale.y * cos(scaleMove.x)), scale.z);
// rotate on y pole
    vec3 rotatey = vec3((rotatez.x * cos(scaleMove.y) - rotatez.z * sin(scaleMove.y)), rotatez.y, (rotatez.x * sin(scaleMove.y) + rotatez.z * cos(scaleMove.y)));
// rotate on x pole
    vec3 rotatex = vec3(rotatey.x, (rotatey.y * cos(scaleMove.z) - rotatey.z * sin(scaleMove.z)), (rotatey.y * sin(scaleMove.z) + rotatey.z * cos(scaleMove.z)));
// move
vec3 move = vec3(rotatex.xy, rotatex.z - 0.2);
// terrible perspective transform
vec3 persp = vec3( move.x  / ( (move.z + 2) / 3 ) ,
		   move.y  / ( (move.z + 2) / 3 ) ,
		     move.z);

    gl_Position = vec4(persp, 1.0);
    fragmentColor = vertexColor;
}
`
	fragmentShaderSource = `
#version 330
out vec4 outColor;
in vec3 fragmentColor;
void main()
{
	outColor = vec4(fragmentColor, 1.0);
}
`
)
