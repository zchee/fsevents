//go:build ignore
// +build ignore

package fsevents

/*
#cgo LDFLAGS: -framework CoreServices
#include <CoreServices/CoreServices.h>
#include <sys/stat.h>

// #define __DARWIN_UNIX03 0
// #define KERNEL
// #define _DARWIN_USE_64_BIT_INODE
// #include <dirent.h>
// #include <fcntl.h>
// #include <poll.h>
// #include <signal.h>
// #include <termios.h>
// #include <unistd.h>
// #include <mach/mach.h>
// #include <mach/message.h>
// #include <sys/event.h>
// #include <sys/kern_control.h>
// #include <sys/mman.h>
// #include <sys/mount.h>
// #include <sys/param.h>
// #include <sys/ptrace.h>
// #include <sys/resource.h>
// #include <sys/select.h>
// #include <sys/signal.h>
// #include <sys/socket.h>
// #include <sys/stat.h>
// #include <sys/time.h>
// #include <sys/types.h>
// #include <sys/ucred.h>
// #include <sys/uio.h>
// #include <sys/un.h>
// #include <sys/utsname.h>
// #include <sys/wait.h>
// #include <net/bpf.h>
// #include <net/if.h>
// #include <net/if_dl.h>
// #include <net/if_var.h>
// #include <net/route.h>
// #include <netinet/in.h>
// #include <netinet/icmp6.h>
// #include <netinet/tcp.h>

enum {
	sizeofPtr = sizeof(void*),
};

// static CFArrayRef ArrayCreateMutable(int len) {
// 	return CFArrayCreateMutable(NULL, len, &kCFTypeArrayCallBacks);
// }
//
// extern void fsevtCallback(FSEventStreamRef p0, uintptr_t info, size_t p1, char** p2, FSEventStreamEventFlags* p3, FSEventStreamEventId* p4);
//
// static FSEventStreamRef EventStreamCreateRelativeToDevice(FSEventStreamContext * context, uintptr_t info, dev_t dev, CFArrayRef paths, FSEventStreamEventId since, CFTimeInterval latency, FSEventStreamCreateFlags flags) {
// 	context->info = (void*) info;
// 	return FSEventStreamCreateRelativeToDevice(NULL, (FSEventStreamCallback) fsevtCallback, context, dev, paths, since, latency, flags);
// }
//
// static FSEventStreamRef EventStreamCreate(FSEventStreamContext * context, uintptr_t info, CFArrayRef paths, FSEventStreamEventId since, CFTimeInterval latency, FSEventStreamCreateFlags flags) {
// 	context->info = (void*) info;
// 	return FSEventStreamCreate(NULL, (FSEventStreamCallback) fsevtCallback, context, paths, since, latency, flags);
// }
*/
import "C"

// Machine characteristics

const (
	SizeofPtr      = C.sizeofPtr
	SizeofShort    = C.sizeof_short
	SizeofInt      = C.sizeof_int
	SizeofLong     = C.sizeof_long
	SizeofLongLong = C.sizeof_longlong
)

// Basic types

type (
	_C_short     C.short
	_C_int       C.int
	_C_long      C.long
	_C_long_long C.longlong
)

type FSEventStream C.struct_FSEventStream
