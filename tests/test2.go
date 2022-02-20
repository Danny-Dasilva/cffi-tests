package main
/*

typedef long GoInt

typedef struct {
	int Y;
} Nested;

struct Vertex {
    int X;
    Nested Y;
};

*/
import "C"
import "fmt"

//export getVertex
func getVertex(X C.int, Y C.int) C.struct_Vertex {
	v := C.struct_Vertex{X, C.Nested{Y}}
    return v

}
func main() {
    fmt.Println(getVertex(1, 1))
}
