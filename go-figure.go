package main

import (
	"fmt"
	"os"
  "strings"
	"log"
	"path/filepath"
)

var digits    [10][]string
var bigAlphas [26][]string
var lilAlphas [26][]string

func main() {
  checkArgs()
  setAsciiMap()
  printChars()
}

func checkArgs() {
  if len(os.Args) == 1 {
    fmt.Printf("usage: %s <[a-zA-Z0-9]>\n", filepath.Base(os.Args[0]))
    os.Exit(1)
  }
}

func printChars() {
  input := strings.Join(os.Args[1:], " ")
  for r := 0 ; r < 6 ; r++ {
    printRow := ""
    for c := 0 ; c < len(input) ; c++ {
      char := input[c]
      digitIndex := char - '0'
      bigAlphaIndex := char - 'A'
      lilAlphaIndex := char - 'a'
      if digitIndex >= 0 && digitIndex < 10 {
        printRow = printRow + digits[digitIndex][r]
      } else if bigAlphaIndex >= 0 && bigAlphaIndex < 26 {
        printRow = printRow + bigAlphas[bigAlphaIndex][r]
      } else if lilAlphaIndex >= 0 && lilAlphaIndex < 26 {
        printRow = printRow + lilAlphas[lilAlphaIndex][r]
      } else if char == ' ' {
        printRow = printRow + "   "
      } else {
				log.Fatal("invalid input. only [a-zA-Z0-9]")
      }
    }
    fmt.Println(printRow)
  }
}

func setAsciiMap() {
  digits[0] = zero
  digits[1] = one
  digits[2] = two
  digits[3] = three
  digits[4] = four
  digits[5] = five
  digits[6] = six
  digits[7] = seven
  digits[8] = eight
  digits[9] = nine

  bigAlphas[0]  = big_a
  bigAlphas[1]  = big_b
  bigAlphas[2]  = big_c
  bigAlphas[3]  = big_d
  bigAlphas[4]  = big_e
  bigAlphas[5]  = big_f
  bigAlphas[6]  = big_g
  bigAlphas[7]  = big_h
  bigAlphas[8]  = big_i
  bigAlphas[9]  = big_j
  bigAlphas[10] = big_k
  bigAlphas[11] = big_l
  bigAlphas[12] = big_m
  bigAlphas[13] = big_n
  bigAlphas[14] = big_o
  bigAlphas[15] = big_p
  bigAlphas[16] = big_q
  bigAlphas[17] = big_r
  bigAlphas[18] = big_s
  bigAlphas[19] = big_t
  bigAlphas[20] = big_u
  bigAlphas[21] = big_v
  bigAlphas[22] = big_w
  bigAlphas[23] = big_x
  bigAlphas[24] = big_y
  bigAlphas[25] = big_z

  lilAlphas[0]  = lil_a
  lilAlphas[1]  = lil_b
  lilAlphas[2]  = lil_c
  lilAlphas[3]  = lil_d
  lilAlphas[4]  = lil_e
  lilAlphas[5]  = lil_f
  lilAlphas[6]  = lil_g
  lilAlphas[7]  = lil_h
  lilAlphas[8]  = lil_i
  lilAlphas[9]  = lil_j
  lilAlphas[10] = lil_k
  lilAlphas[11] = lil_l
  lilAlphas[12] = lil_m
  lilAlphas[13] = lil_n
  lilAlphas[14] = lil_o
  lilAlphas[15] = lil_p
  lilAlphas[16] = lil_q
  lilAlphas[17] = lil_r
  lilAlphas[18] = lil_s
  lilAlphas[19] = lil_t
  lilAlphas[20] = lil_u
  lilAlphas[21] = lil_v
  lilAlphas[22] = lil_w
  lilAlphas[23] = lil_x
  lilAlphas[24] = lil_y
  lilAlphas[25] = lil_z
}

var zero = []string{
" 000  ",
"0  00 ",
"0 0 0 ",
"00  0 ",
" 000  ",
"      ",
}
var one = []string{
" 11  ",
"111  ",
" 11  ",
" 11  ",
"1111 ",
"      ",
}
var two = []string{
" 22  ",
"2  2 ",
"  2  ",
" 2   ",
"2222 ",
"      ",
}
var three = []string{
"333  ",
"   3 ",
" 33  ",
"   3 ",
"333  ",
"      ",
}
var four = []string{
"4  4 ",
"4  4 ",
"4444 ",
"   4 ",
"   4 ",
"      ",
}
var five = []string{
"5555 ",
"5    ",
"555  ",
"   5 ",
"555  ",
"      ",
}
var six = []string{
"  6   ",
" 6    ",
"6666  ",
"6   6 ",
" 666  ",
"      ",
}
var seven = []string{
"77777 ",
"   7  ",
"  7   ",
"  7   ",
"  7   ",
"      ",
}
var eight = []string{
" 888  ",
"8   8 ",
" 888  ",
"8   8 ",
" 888  ",
"      ",
}
var nine = []string{
" 9999 ",
"9   9 ",
" 9999 ",
"   9  ",
"  9   ",
"      ",
}
var big_a = []string{
" AA   ",
"A  A  ",
"AAAA  ",
"A  A  ",
"A  A  ",
"      ",
}
var big_b = []string{
"BBBB  ",
"B   B ",
"BBBB  ",
"B   B ",
"BBBB  ",
"      ",
}
var big_c = []string{
" CCC ",
"C    ",
"C    ",
"C    ",
" CCC ",
"     ",
}
var big_d = []string{
"DDD  ",
"D  D ",
"D  D ",
"D  D ",
"DDD  ",
"     ",
}
var big_e = []string{
"EEEE ",
"E    ",
"EEE  ",
"E    ",
"EEEE ",
"     ",
}
var big_f = []string{
"FFFF ",
"F    ",
"FFF  ",
"F    ",
"F    ",
"     ",
}
var big_g = []string{
" GGG  ",
"G     ",
"G  GG ",
"G   G ",
" GGG  ",
"      ",
}
var big_h = []string{
"H  H ",
"H  H ",
"HHHH ",
"H  H ",
"H  H ",
"     ",
}
var big_i = []string{
"III ",
" I  ",
" I  ",
" I  ",
"III ",
"    ",
}
var big_j = []string{
"    J ",
"    J ",
"    J ",
"J   J ",
" JJJ  ",
"      ",
}
var big_k = []string{
"K  K ",
"K K  ",
"KK   ",
"K K  ",
"K  K ",
"     ",
}
var big_l = []string{
"L    ",
"L    ",
"L    ",
"L    ",
"LLLL ",
"     ",
}
var big_m = []string{
"M   M ",
"MM MM ",
"M M M ",
"M   M ",
"M   M ",
"      ",
}
var big_n = []string{
"N   N ",
"NN  N ",
"N N N ",
"N  NN ",
"N   N ",
"      ",
}
var big_o = []string{
" OOO  ",
"O   O ",
"O   O ",
"O   O ",
" OOO  ",
"      ",
}
var big_p = []string{
"PPPP  ",
"P   P ",
"PPPP  ",
"P     ",
"P     ",
"      ",
}
var big_q = []string{
" QQQ   ",
"Q   Q  ",
"Q   Q  ",
"Q  QQ  ",
" QQQQ  ",
"     Q ",
}
var big_r = []string{
"RRRR  ",
"R   R ",
"RRRR  ",
"R R   ",
"R  RR ",
"      ",
}
var big_s = []string{
" SSS  ",
"S     ",
" SSS  ",
"    S ",
"SSSS  ",
"      ",
}
var big_t = []string{
"TTTTTT ",
"  TT   ",
"  TT   ",
"  TT   ",
"  TT   ",
"       ",
}
var big_u = []string{
"U   U ",
"U   U ",
"U   U ",
"U   U ",
" UUU  ",
"      ",
}
var big_v = []string{
"V     V ",
"V     V ",
" V   V  ",
"  V V   ",
"   V    ",
"        ",
}
var big_w = []string{
"W     W ",
"W     W ",
"W  W  W ",
" W W W  ",
"  W W   ",
"      ",
}
var big_x = []string{
"X   X ",
" X X  ",
"  X   ",
" X X  ",
"X   X ",
"      ",
}
var big_y = []string{
"Y   Y ",
" Y Y  ",
"  Y   ",
"  Y   ",
"  Y   ",
"      ",
}
var big_z = []string{
"ZZZZZ ",
"   Z  ",
"  Z   ",
" Z    ",
"ZZZZZ ",
"      ",
}
var lil_a = []string{
"    ",
"    ",
" aa ",
"a a ",
"aaa ",
"    ",
}
var lil_b = []string{
"b    ",
"b    ",
"bbb  ",
"b  b ",
"bbb  ",
"     ",
}
var lil_c = []string{
"     ",
"     ",
" ccc ",
"c    ",
" ccc ",
"     ",
}
var lil_d = []string{
"   d ",
"   d ",
" ddd ",
"d  d ",
" ddd ",
"     ",
}
var lil_e = []string{
"    ",
"    ",
"eee ",
"e e ",
"ee  ",
"    ",
}
var lil_f = []string{
" ff ",
" f  ",
"fff ",
" f  ",
" f  ",
"    ",
}
var lil_g = []string{
"    ",
"ggg ",
"g g ",
"ggg ",
"  g ",
"ggg ",
}
var lil_h = []string{
"h    ",
"h    ",
"hhh  ",
"h  h ",
"h  h ",
"     ",
}
var lil_i = []string{
"   ",
"ii ",
"   ",
"ii ",
"ii ",
"   ",
}
var lil_j = []string{
"   j ",
"     ",
"   j ",
"   j ",
"j  j ",
" jj  ",
}
var lil_k = []string{
"k    ",
"k k  ",
"kk   ",
"k k  ",
"k  k ",
"     ",
}
var lil_l = []string{
"l ",
"l ",
"l ",
"l ",
"l ",
"  ",
}
var lil_m = []string{
"      ",
"      ",
"mmmm  ",
"m m m ",
"m m m ",
"      ",
}
var lil_n = []string{
"     ",
"     ",
"nnn  ",
"n  n ",
"n  n ",
"     ",
}
var lil_o = []string{
"    ",
"    ",
"ooo ",
"o o ",
"ooo ",
"    ",
}
var lil_p = []string{
"     ",
"ppp  ",
"p  p ",
"ppp  ",
"p    ",
"p    ",
}
var lil_q = []string{
"      ",
" qqq  ",
"q  q  ",
" qqq  ",
"   q  ",
"   qq ",
}
var lil_r = []string{
"    ",
"    ",
"rrr ",
"r   ",
"r   ",
"    ",
}
var lil_s = []string{
"    ",
"    ",
" ss ",
" s  ",
"ss  ",
"    ",
}
var lil_t = []string{
" t  ",
" t  ",
"ttt ",
" t  ",
" tt ",
"    ",
}
var lil_u = []string{
"     ",
"     ",
"u  u ",
"u  u ",
" uuu ",
"     ",
}
var lil_v = []string{
"    ",
"    ",
"v v ",
"v v ",
" v  ",
"    ",
}
var lil_w = []string{
"      ",
"      ",
"w   w ",
"w w w ",
" w w  ",
"      ",
}
var lil_x = []string{
"    ",
"    ",
"x x ",
" x  ",
"x x ",
"    ",
}
var lil_y = []string{
"     ",
"y  y ",
"y  y ",
" yyy ",
"   y ",
"yyy  ",
}
var lil_z = []string{
"    ",
"    ",
"zz  ",
" z  ",
" zz ",
"    ",
}
