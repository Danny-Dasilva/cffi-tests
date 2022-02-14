package main
/*
#include <stdio.h>
#include <stdlib.h>

static void myprint(char* s) {
  printf("%s\n", s);
}
struct Vertex {
    int X;
    char* Y;
};

static void myyprint(char* s) {
  printf("%s\n", s);
}
*/
import "C"
import "unsafe"
import "fmt"

//export getVertex
func getVertex(X C.int, Y *C.char) C.struct_Vertex {
	v := C.struct_Vertex{X, Y}
	fmt.Println(v.X)
    return C.struct_Vertex{X, Y}

}



func main() {
	cs := C.CString("Hello from stdio")
    fmt.Println(getVertex(1, cs))
	// cs := C.CString("Hello from stdio")
	// C.myyprint(cs)
	C.free(unsafe.Pointer(cs))
}
