// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mtl_test

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"golang.org/x/image/math/f32"

	"github.com/duplicants-ai/ebiten/internal/graphicsdriver/metal/mtl"
)

func init() {
	// for these tests to pass Metal must be linked directly.
	// it is not needed for any of the others nor for Ebitengine to work properly.
	//go:cgo_import_dynamic _ _ "/System/Library/Frameworks/Metal.framework/Metal"

	// It is also necessary for CoreGraphics to be linked
	_, _ = purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
}

func Example_listDevices() {
	device, err := mtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}
	printJSON("preferred system default Metal device = ", device)

	// This test is not executed by `go test` command as this doesn't have an `Output:` comment.

	// Sample output:
	// all Metal devices in the system = [
	// 	{
	// 		"Headless": false,
	// 		"LowPower": true,
	// 		"Removable": false,
	// 		"RegistryID": 4294968287,
	// 		"Name": "Intel Iris Pro Graphics"
	// 	},
	// 	{
	// 		"Headless": false,
	// 		"LowPower": false,
	// 		"Removable": false,
	// 		"RegistryID": 4294968322,
	// 		"Name": "AMD Radeon R9 M370X"
	// 	}
	// ]
	// preferred system default Metal device = {
	// 	"Headless": false,
	// 	"LowPower": false,
	// 	"Removable": false,
	// 	"RegistryID": 4294968322,
	// 	"Name": "AMD Radeon R9 M370X"
	// }
}

// printJSON prints label, then v as JSON encoded with indent to stdout. It panics on any error.
// It's meant to be used by examples to print the output.
func printJSON(label string, v any) {
	fmt.Print(label)
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func Example_renderTriangle() {
	device, err := mtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}

	// Create a render pipeline state.
	const source = `#include <metal_stdlib>

using namespace metal;

struct Vertex {
	float4 position [[position]];
	float4 color;
};

vertex Vertex VertexShader(
	uint vertexID [[vertex_id]],
	device Vertex * vertices [[buffer(0)]]
) {
	return vertices[vertexID];
}

fragment float4 FragmentShader(Vertex in [[stage_in]]) {
	return in.color;
}
`
	lib, err := device.NewLibraryWithSource(source, mtl.CompileOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	vs, err := lib.NewFunctionWithName("VertexShader")
	if err != nil {
		log.Fatalln(err)
	}
	fs, err := lib.NewFunctionWithName("FragmentShader")
	if err != nil {
		log.Fatalln(err)
	}
	var rpld mtl.RenderPipelineDescriptor
	rpld.VertexFunction = vs
	rpld.FragmentFunction = fs
	rpld.ColorAttachments[0].PixelFormat = mtl.PixelFormatRGBA8UNorm
	rpld.ColorAttachments[0].WriteMask = mtl.ColorWriteMaskAll
	rps, err := device.NewRenderPipelineStateWithDescriptor(rpld)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a vertex buffer.
	type Vertex struct {
		Position f32.Vec4
		Color    f32.Vec4
	}
	vertexData := [...]Vertex{
		{f32.Vec4{+0.00, +0.75, 0, 1}, f32.Vec4{1, 1, 1, 1}},
		{f32.Vec4{-0.75, -0.75, 0, 1}, f32.Vec4{1, 1, 1, 1}},
		{f32.Vec4{+0.75, -0.75, 0, 1}, f32.Vec4{0, 0, 0, 1}},
	}
	vertexBuffer := device.NewBufferWithBytes(unsafe.Pointer(&vertexData[0]), unsafe.Sizeof(vertexData), mtl.ResourceStorageModeManaged)

	// Create an output texture to render into.
	td := mtl.TextureDescriptor{
		TextureType: mtl.TextureType2D,
		PixelFormat: mtl.PixelFormatRGBA8UNorm,
		Width:       80,
		Height:      20,
		StorageMode: mtl.StorageModeManaged,
	}
	texture := device.NewTextureWithDescriptor(td)

	cq := device.NewCommandQueue()
	cb := cq.CommandBuffer()

	// Encode all render commands.
	var rpd mtl.RenderPassDescriptor
	rpd.ColorAttachments[0].LoadAction = mtl.LoadActionClear
	rpd.ColorAttachments[0].StoreAction = mtl.StoreActionStore
	rpd.ColorAttachments[0].ClearColor = mtl.ClearColor{Red: 0, Green: 0, Blue: 0, Alpha: 1}
	rpd.ColorAttachments[0].Texture = texture
	rce := cb.RenderCommandEncoderWithDescriptor(rpd)
	rce.SetRenderPipelineState(rps)
	rce.SetVertexBuffer(vertexBuffer, 0, 0)
	rce.DrawPrimitives(mtl.PrimitiveTypeTriangle, 0, 3)
	rce.EndEncoding()

	// Encode all blit commands.
	bce := cb.BlitCommandEncoder()
	bce.Synchronize(texture)
	bce.EndEncoding()

	cb.Commit()
	cb.WaitUntilCompleted()

	// Read pixels from output texture into an image.
	img := image.NewNRGBA(image.Rect(0, 0, texture.Width(), texture.Height()))
	bytesPerRow := 4 * texture.Width()
	region := mtl.RegionMake2D(0, 0, texture.Width(), texture.Height())
	texture.GetBytes(&img.Pix[0], uintptr(bytesPerRow), region, 0)

	// Output image to stdout as grayscale ASCII art.
	levels := []struct {
		MinY  uint8
		Shade string
	}{{220, " "}, {170, "░"}, {85, "▒"}, {35, "▓"}, {0, "█"}}
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			for _, l := range levels {
				if c.Y >= l.MinY {
					fmt.Print(l.Shade)
					break
				}
			}
		}
		fmt.Println()
	}

	// Output:
	// ████████████████████████████████████████████████████████████████████████████████
	// ████████████████████████████████████████████████████████████████████████████████
	// ████████████████████████████████████████████████████████████████████████████████
	// ██████████████████████████████████████    ██████████████████████████████████████
	// ████████████████████████████████████        ████████████████████████████████████
	// ██████████████████████████████████        ░░░░██████████████████████████████████
	// ████████████████████████████████        ░░░░░░░░████████████████████████████████
	// ██████████████████████████████        ░░░░░░░░░░░░██████████████████████████████
	// ████████████████████████████        ░░░░░░░░░░░░▒▒▒▒████████████████████████████
	// ██████████████████████████        ░░░░░░░░░░░░▒▒▒▒▒▒▒▒██████████████████████████
	// ████████████████████████        ░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒████████████████████████
	// ██████████████████████        ░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██████████████████████
	// ████████████████████        ░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒████████████████████
	// ██████████████████        ░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓██████████████████
	// ████████████████        ░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓████████████████
	// ██████████████        ░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓██████████████
	// ████████████        ░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓████████████████
	// ████████████████████████████████████████████████████████████████████████████████
	// ████████████████████████████████████████████████████████████████████████████████
	// ████████████████████████████████████████████████████████████████████████████████
}
