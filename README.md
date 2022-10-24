# mmm-go
Manual memory management in Go. There are allocators based on C's malloc and allocators based on the Win32 APIs:\
https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc \
https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heaprealloc \
https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapfree \

These APIs are used through https://github.com/junyupL/sys which is a modified version of https://cs.opensource.google/go/x/sys with HeapAlloc, HeapReAlloc and HeapFree added.

There are also minimal versions of dynamic array, linked list, hash map, and AVL tree using these allocators.
