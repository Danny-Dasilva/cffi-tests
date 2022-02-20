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



char* getRequest(GoString data);
void freeString(char* cs);
""")

lib = ffi.dlopen("./golang.so")


def create_payload_from_dict(data):
    json_data  = json.dumps(data).encode("utf-8")
    data = ffi.new("char[]", json_data)
    msg = ffi.new("GoString*", {'data':data, 'len':len(json_data)})
    return msg[0]
def get_dict_from_payload(data):
    output = ffi.string(data).decode()
    lib.freeString(data)
    return json.loads(output)

test_dict = {
		"url": "https://google.com",
		"ja3": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0",
		"userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36"
	}

payload = create_payload_from_dict(test_dict)
call = lib.getRequest(payload)


data = get_dict_from_payload(call)
print(data)