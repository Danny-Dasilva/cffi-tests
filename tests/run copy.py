from __future__ import print_function
import sys
from cffi import FFI

is_64b = sys.maxsize > 2**32

ffi = FFI()
if is_64b: ffi.cdef("typedef long GoInt;\n")
else:      ffi.cdef("typedef int GoInt;\n")

ffi.cdef("""
typedef struct {
    void* data;
    GoInt len;
    GoInt cap;
} GoSlice;

typedef struct {
    const char *data;
    GoInt len;
} GoString;

typedef struct {
    int X;
    GoString data;
} Vertex;

Vertex getVertex(int a,  char* b);
""")

lib = ffi.dlopen("./test.so")
def to_string(data):
    return ffi.string(data).decode()
# vertex = lib.getVertex(12,"Hello Python!".encode("utf-8"))
# breakpoint()
# breakpoint()
data = ffi.new("char*","b".encode('ascii')) 
# msg = ffi.new("GoString*", {'data':data, 'len':13})
output = to_string(data)
print(output)