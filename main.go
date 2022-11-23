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
var UniScale1 gl.Int
var UniScale2 gl.Int
var UniScalelamb gl.Int

var f = []int{}
var fn = []int{}
var v = []float64{}
var vn = []float64{}
var triangle_vertices = []gl.Float{}
var triangle_colours = []gl.Float{}
var triangle_normals = []gl.Float{}

var color1 gl.Float
var color2 gl.Float
var color3 gl.Float

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
var random1 gl.Float
var random2 gl.Float
var random3 gl.Float
func main() {
	var window *sdl.Window
	var context sdl.GLContext
	var event sdl.Event
	var running bool
	var err error

	// open file for reading
	// read line by line
	myProgramName := os.Args[1]
	cmdcolor1:= os.Args[2]
	cmdcolor2:= os.Args[3]
	cmdcolor3:= os.Args[4]
      
    // it will display 
    // the program name

	lines, err := readLines(myProgramName)
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

		
		if read[0] == "f" {
			fline1 := strings.Split(read[1], "/")
			fline2 := strings.Split(read[2], "/")
			fline3 := strings.Split(read[3], "/")
			intfline1, d1 := strconv.Atoi(fline1[0])
			intfline2, d2 := strconv.Atoi(fline2[0])
			intfline3, d3 := strconv.Atoi(fline3[0])
			intfvline1, d1 := strconv.Atoi(fline1[2])
			intfvline2, d2 := strconv.Atoi(fline2[2])
			intfvline3, d3 := strconv.Atoi(fline3[2])
			f = append(f, intfline1, intfline2, intfline3)
			fn = append(fn, intfvline1, intfvline2, intfvline3)
			
			d1t := []error{}
			d2t := []error{}
			d3t := []error{}

			d1t = append(d1t, d1)
			d2t = append(d2t, d2)
			d3t = append(d3t, d3)


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

		} else if read[0] == "vn" {
			vnline1 := strings.Fields(read[1])
			vnline2 := strings.Fields(read[2])
			vnline3 := strings.Fields(read[3])

			if floatvnline1, err := strconv.ParseFloat(vnline1[0], 64); err == nil {
				vn = append(vn, floatvnline1)
			}
			if floatvnline2, err := strconv.ParseFloat(vnline2[0], 64); err == nil {
				vn = append(vn, floatvnline2)
			}
			if floatvnline3, err := strconv.ParseFloat(vnline3[0], 64); err == nil {
				vn = append(vn, floatvnline3)
			}
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
		//random1 := gl.Float(rand.Float64())
		/*random2 := gl.Float(rand.Float64())
		random3 := gl.Float(rand.Float64())*/
		
		col1fl, err := strconv.ParseFloat(cmdcolor1, 1)
		col2fl, err := strconv.ParseFloat(cmdcolor2, 1)
		col3fl, err := strconv.ParseFloat(cmdcolor3, 1)

		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		color1= gl.Float(col1fl)
		color2= gl.Float(col2fl)
		color3= gl.Float(col3fl) 
		triangle_vertices = append(triangle_vertices, t1, t2, t3)
		triangle_colours = append(triangle_colours,color1, color2, color3)

	}

	random1=color1
	random2=color2
	random3=color3
	for i := 0; i < len(fn); i++ {
		num1vn := fn[i]*3 - 3
		num2vn := fn[i]*3 - 2
		num3vn := fn[i]*3 - 1

		t1n := gl.Float(vn[num1vn])
		t2n := gl.Float(vn[num2vn])
		t3n := gl.Float(vn[num3vn])

		triangle_normals = append(triangle_normals, t1n, t2n, t3n)
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

	
	// NORMAL BUFFER
	var normalbuffer gl.Uint
	gl.GenBuffers(1, &normalbuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalbuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(triangle_normals)*4),
	gl.Pointer(&triangle_normals[0]), gl.STATIC_DRAW)

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
	
	gl.EnableVertexAttribArray(2)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalbuffer)	
	gl.VertexAttribPointer(2, 3, gl.FLOAT, gl.FALSE, 0, nil)	

	//UNIFORM HOOK
	unistring := gl.GLString("scaleMove")
	UniScale = gl.GetUniformLocation(program, unistring)
	fmt.Printf("Uniform Link: %v\n", UniScale+1)

	//UNIFORM HOOK
	unistring1 := gl.GLString("lightSource")
	UniScale1 = gl.GetUniformLocation(program, unistring1)
	fmt.Printf("Uniform Link: %v\n", UniScale1+1)

	unistring2 := gl.GLString("colourChange")
	UniScale2 = gl.GetUniformLocation(program, unistring2)
	fmt.Printf("Uniform Link: %v\n", UniScale2+1)

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
var counter=0

func drawgl() {
	counter++
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	var colourbuffer gl.Uint
	gl.GenBuffers(1, &colourbuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, colourbuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(triangle_colours)*4), gl.Pointer(&triangle_colours[0]), gl.STATIC_DRAW)

	// VERTEX ARRAY HOOK COLOURS
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, colourbuffer)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, gl.FALSE, 0, nil)

	uniYaw = yrot * (math.Pi / 180.0)
	yrot = yrot - 1.0
	uniPitch = zrot * (math.Pi / 180.0)
	zrot = zrot - 0.5
	uniRoll = xrot * (math.Pi / 180.0)
	xrot = xrot - 0.2
	
	gl.Uniform4f(UniScale, gl.Float(uniRoll), gl.Float(uniYaw), gl.Float(uniPitch), gl.Float(uniscale))
	gl.Uniform3f(UniScale1, gl.Float(10), gl.Float(0), gl.Float(30))

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.DrawArrays(gl.TRIANGLES, gl.Int(0), gl.Sizei(len(triangle_vertices)*4))
	

	if (counter==60){
		counter=0
		random1 = gl.Float(rand.Float64())
		time.Sleep(40 * time.Millisecond)
		random2 = gl.Float(rand.Float64())
		time.Sleep(40 * time.Millisecond)
		random3 = gl.Float(rand.Float64())
		gl.Uniform3f(UniScale2, gl.Float(random1), gl.Float(random2), gl.Float(random3))
	}else{	gl.Uniform3f(UniScale2, gl.Float(random1), gl.Float(random2), gl.Float(random3))
	}
	


	time.Sleep(50 * time.Millisecond)
	
}

const (
	winTitle           = "OpenGL Shader"
	winWidth           = 760
	winHeight          = 600
	vertexShaderSource = `
#version 330
layout (location = 0) in vec3 Position;
layout(location = 1) in vec3 vertexColor;
layout(location = 2) in vec3 normal;


uniform vec4 scaleMove;
uniform vec3 lightSource;
uniform vec3 colourChange;
out vec3 fragmentColor;

struct Lights
{
  vec3 position;
  vec3 diffuse; // Colour
};

float lambert(vec3 N, vec3 L)
{
  vec3 nrmN = normalize(N);
  vec3 nrmL = normalize(L);
  float result = dot(nrmN, nrmL);
  return max(result, 0.0);
}

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

	Lights light;
	light.diffuse = vec3(1.0, 1.0, 1.0);
  
	fragmentColor = colourChange*light.diffuse* lambert(normal, lightSource);;

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
