## Task
> A friend sent me this pdf, I do not know what he wants me to do with this.

## Solution
We made a hexdump of the pdf because 
the content in it wasn't relevant.
Examining the hexdump we see that 
it contains stream object among which a `PNG` file

````
000048e0  3c 2f 54 69 74 6c 65 3c  46 45 46 46 30 30 35 35  |</Title<FEFF0055|
000048f0  30 30 36 45 30 30 37 34  30 30 36 39 30 30 37 34  |006E007400690074|
00004900  30 30 36 43 30 30 36 35  30 30 36 34 30 30 32 30  |006C006500640020|
00004910  30 30 33 31 3e 2f 43 72  65 61 74 6f 72 3c 46 45  |0031>/Creator<FE|
00004920  46 46 30 30 34 43 30 30  36 39 30 30 36 32 30 30  |FF004C0069006200|
00004930  37 32 30 30 36 35 30 30  34 46 30 30 36 36 30 30  |720065004F006600|
00004940  36 36 30 30 36 39 30 30  36 33 30 30 36 35 30 30  |6600690063006500|
00004950  32 30 30 30 33 37 30 30  32 45 30 30 33 33 30 30  |200037002E003300|
00004960  32 45 30 30 33 37 30 30  32 45 30 30 33 32 3e 2f  |2E0037002E0032>/|
00004970  50 72 6f 64 75 63 65 72  3c 46 45 46 46 30 30 34  |Producer<FEFF004|
00004980  43 30 30 36 39 30 30 36  32 30 30 37 32 30 30 36  |C006900620072006|
00004990  35 30 30 34 46 30 30 36  36 30 30 36 36 30 30 36  |5004F00660066006|
000049a0  39 30 30 36 33 30 30 36  35 30 30 32 30 30 30 33  |9006300650020003|
000049b0  37 30 30 32 45 30 30 33  33 30 30 32 45 30 30 33  |7002E0033002E003|
000049c0  37 30 30 32 45 30 30 33  32 3e 2f 43 72 65 61 74  |7002E0032>/Creat|
000049d0  69 6f 6e 44 61 74 65 28  44 3a 32 30 32 34 31 30  |ionDate(D:202410|
000049e0  32 31 31 39 35 35 32 30  2b 30 32 27 30 30 27 29  |21195520+02'00')|
000049f0  3e 3e 0a 65 6e 64 6f 62  6a 0a 0a 32 30 20 30 20  |>>.endobj..20 0 |
00004a00  6f 62 6a 0a 3c 3c 2f 4c  65 6e 67 74 68 20 31 30  |obj.<</Length 10|
00004a10  30 3e 3e 0a 73 74 72 65  61 6d 0a 89 50 4e 47 0d  |0>>.stream..PNG.|
00004a20  0a 1a 0a 00 00 00 0d 49  48 44 52 00 00 07 80 00  |.......IHDR.....|
00004a30  00 04 38 08 06 00 00 00  e8 d3 c1 43 00 00 01 84  |..8........C....|
00004a40  69 43 43 50 49 43 43 20  50 72 6f 66 69 6c 65 00  |iCCPICC Profile.|
00004a50  00 78 9c 7d 91 3d 48 c3  40 1c c5 5f 53 a5 52 2a  |.x.}.=H.@.._S.R*|
````

Wee see that the object has length 100 always but at the end it has 90

```
0001f6e0  0a 65 6e 64 73 74 72 65  61 6d 0a 65 6e 64 6f 62  |.endstream.endob|
0001f6f0  6a 0a 0a 37 34 33 20 30  20 6f 62 6a 0a 3c 3c 2f  |j..743 0 obj.<</|
0001f700  4c 65 6e 67 74 68 20 39  30 3e 3e 0a 73 74 72 65  |Length 90>>.stre|
0001f710  61 6d 0a 08 82 20 08 82  20 08 82 20 08 82 20 08  |am... .. .. .. .|
0001f720  82 20 88 01 02 39 80 09  82 20 08 82 20 08 82 20  |. ...9... .. .. |
0001f730  08 82 20 08 82 20 08 82  20 08 82 20 06 08 e4 00  |.. .. .. .. ....|
0001f740  26 08 82 20 08 82 20 08  82 20 08 82 20 08 82 20  |&.. .. .. .. .. |
0001f750  08 82 20 08 82 18 20 fc  7f 01 83 89 e5 6e d6 bd  |.. ... ......n..|
0001f760  55 00 00 00 00 49 45 4e  44 ae 42 60 82 0a 65 6e  |U....IEND.B`..en|
```

We need to extract form pdf the object with length 100 or 90

After writing to a png file we obtain the flag
![file.png](file.png)
