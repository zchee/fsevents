// go:build darwin && !ios
// +build darwin,!ios

#include "textflag.h"

// NOTE: The trampolines are ABIInternal as they are address-taken in Go code.

TEXT ·fsEventsGetCurrentEventID_trampoline<ABIInternal>(SB), NOSPLIT, $0-0
	JMP FSEventsGetCurrentEventId(SB)

// TEXT ·x509_CFArrayGetValueAtIndex_trampoline<ABIInternal>(SB),NOSPLIT,$0-0
// 	JMP	x509_CFArrayGetValueAtIndex(SB)

// TEXT ·x509_CFDataGetBytePtr_trampoline<ABIInternal>(SB),NOSPLIT,$0-0
// 	JMP	x509_CFDataGetBytePtr(SB)

// TEXT ·x509_CFDataGetLength_trampoline<ABIInternal>(SB),NOSPLIT,$0-0
// 	JMP	x509_CFDataGetLength(SB)

// TEXT ·x509_CFStringCreateWithBytes_trampoline<ABIInternal>(SB),NOSPLIT,$0-0
// 	JMP	x509_CFStringCreateWithBytes(SB)

// TEXT ·x509_CFRelease_trampoline<ABIInternal>(SB),NOSPLIT,$0-0
// 	JMP	x509_CFRelease(SB)

// TEXT ·x509_CFDictionaryGetValueIfPresent_trampoline<ABIInternal>(SB),NOSPLIT,$0-0
// 	JMP	x509_CFDictionaryGetValueIfPresent(SB)

// TEXT ·x509_CFNumberGetValue_trampoline<ABIInternal>(SB),NOSPLIT,$0-0
// 	JMP	x509_CFNumberGetValue(SB)

// TEXT ·x509_CFEqual_trampoline<ABIInternal>(SB),NOSPLIT,$0-0
// 	JMP	x509_CFEqual(SB)
