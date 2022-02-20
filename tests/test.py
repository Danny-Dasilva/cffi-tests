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
    const char *data;
} Vertex;

Vertex getVertex(int a,  char* b);
""")

lib = ffi.dlopen("./test.so")

# vertex = lib.getVertex(12,bytes("data", "utf-8"))
# breakpoint()
data = ffi.new("char[]", b"Hello Python!")
msg = ffi.new("GoString*", {'data':data, 'len':13})
print(msg[0])
breakpoint()
