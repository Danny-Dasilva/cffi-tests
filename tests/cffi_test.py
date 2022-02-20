from inspect import getmembers
from cffi import FFI
ffi = FFI()
from pprint import pprint

def cdata_dict(cd):
    if isinstance(cd, ffi.CData):
        try:
            return ffi.string(cd)
        except TypeError:
            try:
                return [cdata_dict(x) for x in cd]
            except TypeError:
                return {k: cdata_dict(v) for k, v in getmembers(cd)}
    else:
        return cd
        
ffi.cdef("""

typedef struct {
    char RequestID[6];
    int Status;
    char Body[40];
    struct {
        int a, b[3];
    } item;
} Foo;

""")
foo = ffi.new("Foo *",{
    'RequestID': b"Foo",
    "Status": 204,
    "Body": b"Foo",
    'item': {'a': 3, 'b': [1, 2, 3]}
})
pprint(cdata_dict(foo))
