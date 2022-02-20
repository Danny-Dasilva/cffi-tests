from __future__ import print_function
import sys
from cffi import FFI
import json
is_64b = sys.maxsize > 2**32

ffi = FFI()
if is_64b: ffi.cdef("typedef long GoInt;\n")
else:      ffi.cdef("typedef int GoInt;\n")

ffi.cdef("""
typedef unsigned char GoUint8;

typedef struct {
    void* data;
    GoInt len;
    GoInt cap;
} GoSlice;

typedef struct {
    const char *data;
    GoInt len;
} GoString;



char* getVertex(GoString data);
void freeString(char* cs);
""")

lib = ffi.dlopen("./run.so")


def create_payload_from_dict(data):
    
    data = ffi.new("char[]", json.dumps(test_dict).encode("utf-8"))
    msg = ffi.new("GoString*", {'data':data, 'len':len(json.dumps(test_dict))})
    return msg[0]
def get_dict_from_payload(data):
    output = ffi.string(data).decode()
    lib.freeString(data)
    return json.loads(output)

test_dict = {
		"Name": "Standard",
		"Fruit": [
			"Apple",
			"Banana",
			"Orange"
		],
		"ref": 999
	}

payload = create_payload_from_dict(test_dict)
call = lib.getVertex(payload)


data = get_dict_from_payload(call)
breakpoint()