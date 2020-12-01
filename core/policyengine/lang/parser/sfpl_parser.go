// Code generated from Sfpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sfpl
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 310,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 6, 2, 60, 10, 2, 13, 2, 14, 2, 61, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 7, 3, 71, 10, 3, 12, 3, 14, 3, 74, 11, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 109, 10, 4, 12, 4, 14,
	4, 112, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 145, 10,
	5, 12, 5, 14, 5, 148, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 160, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 172, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 184, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 7, 12, 204, 10, 12, 12, 12, 14, 12, 207, 11, 12, 3,
	13, 3, 13, 3, 13, 7, 13, 212, 10, 13, 12, 13, 14, 13, 215, 11, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 5, 14, 232, 10, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	237, 10, 14, 7, 14, 239, 10, 14, 12, 14, 14, 14, 242, 11, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 250, 10, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 7, 15, 256, 10, 15, 12, 15, 14, 15, 259, 11, 15, 5, 15, 261, 10,
	15, 3, 15, 5, 15, 264, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	7, 16, 272, 10, 16, 12, 16, 14, 16, 275, 11, 16, 5, 16, 277, 10, 16, 3,
	16, 5, 16, 280, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 6, 25, 302, 10, 25, 13, 25, 14, 25, 303, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 2, 2, 28, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 2, 6, 3, 2,
	11, 12, 4, 2, 30, 30, 35, 35, 5, 2, 24, 24, 26, 26, 47, 51, 4, 2, 24, 29,
	31, 34, 2, 327, 2, 59, 3, 2, 2, 2, 4, 72, 3, 2, 2, 2, 6, 77, 3, 2, 2, 2,
	8, 113, 3, 2, 2, 2, 10, 149, 3, 2, 2, 2, 12, 161, 3, 2, 2, 2, 14, 173,
	3, 2, 2, 2, 16, 185, 3, 2, 2, 2, 18, 193, 3, 2, 2, 2, 20, 198, 3, 2, 2,
	2, 22, 200, 3, 2, 2, 2, 24, 208, 3, 2, 2, 2, 26, 249, 3, 2, 2, 2, 28, 251,
	3, 2, 2, 2, 30, 267, 3, 2, 2, 2, 32, 283, 3, 2, 2, 2, 34, 285, 3, 2, 2,
	2, 36, 287, 3, 2, 2, 2, 38, 289, 3, 2, 2, 2, 40, 291, 3, 2, 2, 2, 42, 293,
	3, 2, 2, 2, 44, 295, 3, 2, 2, 2, 46, 297, 3, 2, 2, 2, 48, 301, 3, 2, 2,
	2, 50, 305, 3, 2, 2, 2, 52, 307, 3, 2, 2, 2, 54, 60, 5, 6, 4, 2, 55, 60,
	5, 10, 6, 2, 56, 60, 5, 14, 8, 2, 57, 60, 5, 16, 9, 2, 58, 60, 5, 18, 10,
	2, 59, 54, 3, 2, 2, 2, 59, 55, 3, 2, 2, 2, 59, 56, 3, 2, 2, 2, 59, 57,
	3, 2, 2, 2, 59, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2,
	61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 7, 2, 2, 3, 64, 3, 3, 2,
	2, 2, 65, 71, 5, 8, 5, 2, 66, 71, 5, 12, 7, 2, 67, 71, 5, 14, 8, 2, 68,
	71, 5, 16, 9, 2, 69, 71, 5, 18, 10, 2, 70, 65, 3, 2, 2, 2, 70, 66, 3, 2,
	2, 2, 70, 67, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 69, 3, 2, 2, 2, 71, 74,
	3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 75, 76, 7, 2, 2, 3, 76, 5, 3, 2, 2, 2, 77, 78, 7, 42,
	2, 2, 78, 79, 7, 3, 2, 2, 79, 80, 7, 43, 2, 2, 80, 81, 5, 48, 25, 2, 81,
	82, 7, 10, 2, 2, 82, 83, 7, 43, 2, 2, 83, 84, 5, 48, 25, 2, 84, 85, 7,
	9, 2, 2, 85, 86, 7, 43, 2, 2, 86, 110, 5, 20, 11, 2, 87, 88, 9, 2, 2, 2,
	88, 89, 7, 43, 2, 2, 89, 109, 5, 48, 25, 2, 90, 91, 7, 13, 2, 2, 91, 92,
	7, 43, 2, 2, 92, 109, 5, 34, 18, 2, 93, 94, 7, 14, 2, 2, 94, 95, 7, 43,
	2, 2, 95, 109, 5, 30, 16, 2, 96, 97, 7, 15, 2, 2, 97, 98, 7, 43, 2, 2,
	98, 109, 5, 32, 17, 2, 99, 100, 7, 16, 2, 2, 100, 101, 7, 43, 2, 2, 101,
	109, 5, 36, 19, 2, 102, 103, 7, 17, 2, 2, 103, 104, 7, 43, 2, 2, 104, 109,
	5, 38, 20, 2, 105, 106, 7, 18, 2, 2, 106, 107, 7, 43, 2, 2, 107, 109, 5,
	40, 21, 2, 108, 87, 3, 2, 2, 2, 108, 90, 3, 2, 2, 2, 108, 93, 3, 2, 2,
	2, 108, 96, 3, 2, 2, 2, 108, 99, 3, 2, 2, 2, 108, 102, 3, 2, 2, 2, 108,
	105, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111,
	3, 2, 2, 2, 111, 7, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 114, 7, 42,
	2, 2, 114, 115, 7, 3, 2, 2, 115, 116, 7, 43, 2, 2, 116, 117, 5, 48, 25,
	2, 117, 118, 7, 10, 2, 2, 118, 119, 7, 43, 2, 2, 119, 120, 5, 48, 25, 2,
	120, 121, 7, 9, 2, 2, 121, 122, 7, 43, 2, 2, 122, 146, 5, 20, 11, 2, 123,
	124, 9, 2, 2, 2, 124, 125, 7, 43, 2, 2, 125, 145, 5, 48, 25, 2, 126, 127,
	7, 13, 2, 2, 127, 128, 7, 43, 2, 2, 128, 145, 5, 34, 18, 2, 129, 130, 7,
	14, 2, 2, 130, 131, 7, 43, 2, 2, 131, 145, 5, 30, 16, 2, 132, 133, 7, 15,
	2, 2, 133, 134, 7, 43, 2, 2, 134, 145, 5, 32, 17, 2, 135, 136, 7, 16, 2,
	2, 136, 137, 7, 43, 2, 2, 137, 145, 5, 36, 19, 2, 138, 139, 7, 17, 2, 2,
	139, 140, 7, 43, 2, 2, 140, 145, 5, 38, 20, 2, 141, 142, 7, 18, 2, 2, 142,
	143, 7, 43, 2, 2, 143, 145, 5, 40, 21, 2, 144, 123, 3, 2, 2, 2, 144, 126,
	3, 2, 2, 2, 144, 129, 3, 2, 2, 2, 144, 132, 3, 2, 2, 2, 144, 135, 3, 2,
	2, 2, 144, 138, 3, 2, 2, 2, 144, 141, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2,
	146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 9, 3, 2, 2, 2, 148, 146,
	3, 2, 2, 2, 149, 150, 7, 42, 2, 2, 150, 151, 7, 4, 2, 2, 151, 152, 7, 43,
	2, 2, 152, 153, 7, 47, 2, 2, 153, 154, 7, 9, 2, 2, 154, 155, 7, 43, 2,
	2, 155, 159, 5, 20, 11, 2, 156, 157, 7, 16, 2, 2, 157, 158, 7, 43, 2, 2,
	158, 160, 5, 36, 19, 2, 159, 156, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	11, 3, 2, 2, 2, 161, 162, 7, 42, 2, 2, 162, 163, 7, 4, 2, 2, 163, 164,
	7, 43, 2, 2, 164, 165, 7, 47, 2, 2, 165, 166, 7, 9, 2, 2, 166, 167, 7,
	43, 2, 2, 167, 171, 5, 20, 11, 2, 168, 169, 7, 16, 2, 2, 169, 170, 7, 43,
	2, 2, 170, 172, 5, 36, 19, 2, 171, 168, 3, 2, 2, 2, 171, 172, 3, 2, 2,
	2, 172, 13, 3, 2, 2, 2, 173, 174, 7, 42, 2, 2, 174, 175, 7, 5, 2, 2, 175,
	176, 7, 43, 2, 2, 176, 177, 7, 47, 2, 2, 177, 178, 7, 9, 2, 2, 178, 179,
	7, 43, 2, 2, 179, 183, 5, 20, 11, 2, 180, 181, 7, 19, 2, 2, 181, 182, 7,
	43, 2, 2, 182, 184, 5, 42, 22, 2, 183, 180, 3, 2, 2, 2, 183, 184, 3, 2,
	2, 2, 184, 15, 3, 2, 2, 2, 185, 186, 7, 42, 2, 2, 186, 187, 7, 6, 2, 2,
	187, 188, 7, 43, 2, 2, 188, 189, 7, 47, 2, 2, 189, 190, 7, 8, 2, 2, 190,
	191, 7, 43, 2, 2, 191, 192, 5, 28, 15, 2, 192, 17, 3, 2, 2, 2, 193, 194,
	7, 42, 2, 2, 194, 195, 7, 20, 2, 2, 195, 196, 7, 43, 2, 2, 196, 197, 5,
	46, 24, 2, 197, 19, 3, 2, 2, 2, 198, 199, 5, 22, 12, 2, 199, 21, 3, 2,
	2, 2, 200, 205, 5, 24, 13, 2, 201, 202, 7, 22, 2, 2, 202, 204, 5, 24, 13,
	2, 203, 201, 3, 2, 2, 2, 204, 207, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205,
	206, 3, 2, 2, 2, 206, 23, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 208, 213, 5,
	26, 14, 2, 209, 210, 7, 21, 2, 2, 210, 212, 5, 26, 14, 2, 211, 209, 3,
	2, 2, 2, 212, 215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2,
	2, 214, 25, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216, 250, 5, 44, 23, 2, 217,
	218, 7, 23, 2, 2, 218, 250, 5, 26, 14, 2, 219, 220, 5, 46, 24, 2, 220,
	221, 5, 52, 27, 2, 221, 250, 3, 2, 2, 2, 222, 223, 5, 46, 24, 2, 223, 224,
	5, 50, 26, 2, 224, 225, 5, 46, 24, 2, 225, 250, 3, 2, 2, 2, 226, 227, 5,
	46, 24, 2, 227, 228, 9, 3, 2, 2, 228, 231, 7, 39, 2, 2, 229, 232, 5, 46,
	24, 2, 230, 232, 5, 28, 15, 2, 231, 229, 3, 2, 2, 2, 231, 230, 3, 2, 2,
	2, 232, 240, 3, 2, 2, 2, 233, 236, 7, 41, 2, 2, 234, 237, 5, 46, 24, 2,
	235, 237, 5, 28, 15, 2, 236, 234, 3, 2, 2, 2, 236, 235, 3, 2, 2, 2, 237,
	239, 3, 2, 2, 2, 238, 233, 3, 2, 2, 2, 239, 242, 3, 2, 2, 2, 240, 238,
	3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2, 2, 2, 242, 240, 3, 2,
	2, 2, 243, 244, 7, 40, 2, 2, 244, 250, 3, 2, 2, 2, 245, 246, 7, 39, 2,
	2, 246, 247, 5, 20, 11, 2, 247, 248, 7, 40, 2, 2, 248, 250, 3, 2, 2, 2,
	249, 216, 3, 2, 2, 2, 249, 217, 3, 2, 2, 2, 249, 219, 3, 2, 2, 2, 249,
	222, 3, 2, 2, 2, 249, 226, 3, 2, 2, 2, 249, 245, 3, 2, 2, 2, 250, 27, 3,
	2, 2, 2, 251, 260, 7, 37, 2, 2, 252, 257, 5, 46, 24, 2, 253, 254, 7, 41,
	2, 2, 254, 256, 5, 46, 24, 2, 255, 253, 3, 2, 2, 2, 256, 259, 3, 2, 2,
	2, 257, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 261, 3, 2, 2, 2, 259,
	257, 3, 2, 2, 2, 260, 252, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 263,
	3, 2, 2, 2, 262, 264, 7, 41, 2, 2, 263, 262, 3, 2, 2, 2, 263, 264, 3, 2,
	2, 2, 264, 265, 3, 2, 2, 2, 265, 266, 7, 38, 2, 2, 266, 29, 3, 2, 2, 2,
	267, 276, 7, 37, 2, 2, 268, 273, 5, 46, 24, 2, 269, 270, 7, 41, 2, 2, 270,
	272, 5, 46, 24, 2, 271, 269, 3, 2, 2, 2, 272, 275, 3, 2, 2, 2, 273, 271,
	3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 277, 3, 2, 2, 2, 275, 273, 3, 2,
	2, 2, 276, 268, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 279, 3, 2, 2, 2,
	278, 280, 7, 41, 2, 2, 279, 278, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280,
	281, 3, 2, 2, 2, 281, 282, 7, 38, 2, 2, 282, 31, 3, 2, 2, 2, 283, 284,
	5, 28, 15, 2, 284, 33, 3, 2, 2, 2, 285, 286, 7, 44, 2, 2, 286, 35, 3, 2,
	2, 2, 287, 288, 5, 46, 24, 2, 288, 37, 3, 2, 2, 2, 289, 290, 5, 46, 24,
	2, 290, 39, 3, 2, 2, 2, 291, 292, 5, 46, 24, 2, 292, 41, 3, 2, 2, 2, 293,
	294, 5, 46, 24, 2, 294, 43, 3, 2, 2, 2, 295, 296, 7, 47, 2, 2, 296, 45,
	3, 2, 2, 2, 297, 298, 9, 4, 2, 2, 298, 47, 3, 2, 2, 2, 299, 300, 6, 25,
	2, 2, 300, 302, 11, 2, 2, 2, 301, 299, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2,
	303, 301, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 49, 3, 2, 2, 2, 305, 306,
	9, 5, 2, 2, 306, 51, 3, 2, 2, 2, 307, 308, 7, 36, 2, 2, 308, 53, 3, 2,
	2, 2, 26, 59, 61, 70, 72, 108, 110, 144, 146, 159, 171, 183, 205, 213,
	231, 236, 240, 249, 257, 260, 263, 273, 276, 279, 303,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'rule'", "'filter'", "'macro'", "'list'", "'name'", "'items'", "'condition'",
	"'desc'", "'action'", "'output'", "'priority'", "'tags'", "'prefilter'",
	"'enabled'", "'warn_evttypes'", "'skip-if-unknown-filter'", "'append'",
	"'required_engine_version'", "'and'", "'or'", "'not'", "'<'", "'<='", "'>'",
	"'>='", "'='", "'!='", "'in'", "'contains'", "'icontains'", "'startswith'",
	"'endswith'", "'pmatch'", "'exists'", "'['", "']'", "'('", "')'", "','",
	"'-'",
}
var symbolicNames = []string{
	"", "RULE", "FILTER", "MACRO", "LIST", "NAME", "ITEMS", "COND", "DESC",
	"ACTION", "OUTPUT", "PRIORITY", "TAGS", "PREFILTER", "ENABLED", "WARNEVTTYPE",
	"SKIPUNKNOWN", "FAPPEND", "REQ", "AND", "OR", "NOT", "LT", "LE", "GT",
	"GE", "EQ", "NEQ", "IN", "CONTAINS", "ICONTAINS", "STARTSWITH", "ENDSWITH",
	"PMATCH", "EXISTS", "LBRACK", "RBRACK", "LPAREN", "RPAREN", "LISTSEP",
	"DECL", "DEF", "SEVERITY", "SFSEVERITY", "FSEVERITY", "ID", "NUMBER", "PATH",
	"STRING", "TAG", "WS", "NL", "COMMENT", "ANY",
}

var ruleNames = []string{
	"policy", "defs", "prule", "srule", "pfilter", "sfilter", "pmacro", "plist",
	"preq", "expression", "or_expression", "and_expression", "term", "items",
	"tags", "prefilter", "severity", "enabled", "warnevttype", "skipunknown",
	"fappend", "variable", "atom", "text", "binary_operator", "unary_operator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SfplParser struct {
	*antlr.BaseParser
}

func NewSfplParser(input antlr.TokenStream) *SfplParser {
	this := new(SfplParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sfpl.g4"

	return this
}

// SfplParser tokens.
const (
	SfplParserEOF         = antlr.TokenEOF
	SfplParserRULE        = 1
	SfplParserFILTER      = 2
	SfplParserMACRO       = 3
	SfplParserLIST        = 4
	SfplParserNAME        = 5
	SfplParserITEMS       = 6
	SfplParserCOND        = 7
	SfplParserDESC        = 8
	SfplParserACTION      = 9
	SfplParserOUTPUT      = 10
	SfplParserPRIORITY    = 11
	SfplParserTAGS        = 12
	SfplParserPREFILTER   = 13
	SfplParserENABLED     = 14
	SfplParserWARNEVTTYPE = 15
	SfplParserSKIPUNKNOWN = 16
	SfplParserFAPPEND     = 17
	SfplParserREQ         = 18
	SfplParserAND         = 19
	SfplParserOR          = 20
	SfplParserNOT         = 21
	SfplParserLT          = 22
	SfplParserLE          = 23
	SfplParserGT          = 24
	SfplParserGE          = 25
	SfplParserEQ          = 26
	SfplParserNEQ         = 27
	SfplParserIN          = 28
	SfplParserCONTAINS    = 29
	SfplParserICONTAINS   = 30
	SfplParserSTARTSWITH  = 31
	SfplParserENDSWITH    = 32
	SfplParserPMATCH      = 33
	SfplParserEXISTS      = 34
	SfplParserLBRACK      = 35
	SfplParserRBRACK      = 36
	SfplParserLPAREN      = 37
	SfplParserRPAREN      = 38
	SfplParserLISTSEP     = 39
	SfplParserDECL        = 40
	SfplParserDEF         = 41
	SfplParserSEVERITY    = 42
	SfplParserSFSEVERITY  = 43
	SfplParserFSEVERITY   = 44
	SfplParserID          = 45
	SfplParserNUMBER      = 46
	SfplParserPATH        = 47
	SfplParserSTRING      = 48
	SfplParserTAG         = 49
	SfplParserWS          = 50
	SfplParserNL          = 51
	SfplParserCOMMENT     = 52
	SfplParserANY         = 53
)

// SfplParser rules.
const (
	SfplParserRULE_policy          = 0
	SfplParserRULE_defs            = 1
	SfplParserRULE_prule           = 2
	SfplParserRULE_srule           = 3
	SfplParserRULE_pfilter         = 4
	SfplParserRULE_sfilter         = 5
	SfplParserRULE_pmacro          = 6
	SfplParserRULE_plist           = 7
	SfplParserRULE_preq            = 8
	SfplParserRULE_expression      = 9
	SfplParserRULE_or_expression   = 10
	SfplParserRULE_and_expression  = 11
	SfplParserRULE_term            = 12
	SfplParserRULE_items           = 13
	SfplParserRULE_tags            = 14
	SfplParserRULE_prefilter       = 15
	SfplParserRULE_severity        = 16
	SfplParserRULE_enabled         = 17
	SfplParserRULE_warnevttype     = 18
	SfplParserRULE_skipunknown     = 19
	SfplParserRULE_fappend         = 20
	SfplParserRULE_variable        = 21
	SfplParserRULE_atom            = 22
	SfplParserRULE_text            = 23
	SfplParserRULE_binary_operator = 24
	SfplParserRULE_unary_operator  = 25
)

// IPolicyContext is an interface to support dynamic dispatch.
type IPolicyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolicyContext differentiates from other interfaces.
	IsPolicyContext()
}

type PolicyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolicyContext() *PolicyContext {
	var p = new(PolicyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_policy
	return p
}

func (*PolicyContext) IsPolicyContext() {}

func NewPolicyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolicyContext {
	var p = new(PolicyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_policy

	return p
}

func (s *PolicyContext) GetParser() antlr.Parser { return s.parser }

func (s *PolicyContext) EOF() antlr.TerminalNode {
	return s.GetToken(SfplParserEOF, 0)
}

func (s *PolicyContext) AllPrule() []IPruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPruleContext)(nil)).Elem())
	var tst = make([]IPruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPruleContext)
		}
	}

	return tst
}

func (s *PolicyContext) Prule(i int) IPruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPruleContext)
}

func (s *PolicyContext) AllPfilter() []IPfilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPfilterContext)(nil)).Elem())
	var tst = make([]IPfilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPfilterContext)
		}
	}

	return tst
}

func (s *PolicyContext) Pfilter(i int) IPfilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPfilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPfilterContext)
}

func (s *PolicyContext) AllPmacro() []IPmacroContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPmacroContext)(nil)).Elem())
	var tst = make([]IPmacroContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPmacroContext)
		}
	}

	return tst
}

func (s *PolicyContext) Pmacro(i int) IPmacroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmacroContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPmacroContext)
}

func (s *PolicyContext) AllPlist() []IPlistContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPlistContext)(nil)).Elem())
	var tst = make([]IPlistContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPlistContext)
		}
	}

	return tst
}

func (s *PolicyContext) Plist(i int) IPlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlistContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPlistContext)
}

func (s *PolicyContext) AllPreq() []IPreqContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPreqContext)(nil)).Elem())
	var tst = make([]IPreqContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPreqContext)
		}
	}

	return tst
}

func (s *PolicyContext) Preq(i int) IPreqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreqContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPreqContext)
}

func (s *PolicyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolicyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolicyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPolicy(s)
	}
}

func (s *PolicyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPolicy(s)
	}
}

func (s *PolicyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPolicy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Policy() (localctx IPolicyContext) {
	localctx = NewPolicyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SfplParserRULE_policy)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SfplParserDECL {
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(52)
				p.Prule()
			}

		case 2:
			{
				p.SetState(53)
				p.Pfilter()
			}

		case 3:
			{
				p.SetState(54)
				p.Pmacro()
			}

		case 4:
			{
				p.SetState(55)
				p.Plist()
			}

		case 5:
			{
				p.SetState(56)
				p.Preq()
			}

		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
		p.Match(SfplParserEOF)
	}

	return localctx
}

// IDefsContext is an interface to support dynamic dispatch.
type IDefsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefsContext differentiates from other interfaces.
	IsDefsContext()
}

type DefsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefsContext() *DefsContext {
	var p = new(DefsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_defs
	return p
}

func (*DefsContext) IsDefsContext() {}

func NewDefsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefsContext {
	var p = new(DefsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_defs

	return p
}

func (s *DefsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefsContext) EOF() antlr.TerminalNode {
	return s.GetToken(SfplParserEOF, 0)
}

func (s *DefsContext) AllSrule() []ISruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISruleContext)(nil)).Elem())
	var tst = make([]ISruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISruleContext)
		}
	}

	return tst
}

func (s *DefsContext) Srule(i int) ISruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISruleContext)
}

func (s *DefsContext) AllSfilter() []ISfilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISfilterContext)(nil)).Elem())
	var tst = make([]ISfilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISfilterContext)
		}
	}

	return tst
}

func (s *DefsContext) Sfilter(i int) ISfilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISfilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISfilterContext)
}

func (s *DefsContext) AllPmacro() []IPmacroContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPmacroContext)(nil)).Elem())
	var tst = make([]IPmacroContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPmacroContext)
		}
	}

	return tst
}

func (s *DefsContext) Pmacro(i int) IPmacroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmacroContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPmacroContext)
}

func (s *DefsContext) AllPlist() []IPlistContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPlistContext)(nil)).Elem())
	var tst = make([]IPlistContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPlistContext)
		}
	}

	return tst
}

func (s *DefsContext) Plist(i int) IPlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlistContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPlistContext)
}

func (s *DefsContext) AllPreq() []IPreqContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPreqContext)(nil)).Elem())
	var tst = make([]IPreqContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPreqContext)
		}
	}

	return tst
}

func (s *DefsContext) Preq(i int) IPreqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreqContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPreqContext)
}

func (s *DefsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterDefs(s)
	}
}

func (s *DefsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitDefs(s)
	}
}

func (s *DefsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitDefs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Defs() (localctx IDefsContext) {
	localctx = NewDefsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SfplParserRULE_defs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserDECL {
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(63)
				p.Srule()
			}

		case 2:
			{
				p.SetState(64)
				p.Sfilter()
			}

		case 3:
			{
				p.SetState(65)
				p.Pmacro()
			}

		case 4:
			{
				p.SetState(66)
				p.Plist()
			}

		case 5:
			{
				p.SetState(67)
				p.Preq()
			}

		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
		p.Match(SfplParserEOF)
	}

	return localctx
}

// IPruleContext is an interface to support dynamic dispatch.
type IPruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPruleContext differentiates from other interfaces.
	IsPruleContext()
}

type PruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPruleContext() *PruleContext {
	var p = new(PruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_prule
	return p
}

func (*PruleContext) IsPruleContext() {}

func NewPruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PruleContext {
	var p = new(PruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_prule

	return p
}

func (s *PruleContext) GetParser() antlr.Parser { return s.parser }

func (s *PruleContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PruleContext) RULE() antlr.TerminalNode {
	return s.GetToken(SfplParserRULE, 0)
}

func (s *PruleContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *PruleContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *PruleContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *PruleContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *PruleContext) DESC() antlr.TerminalNode {
	return s.GetToken(SfplParserDESC, 0)
}

func (s *PruleContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *PruleContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PruleContext) AllPRIORITY() []antlr.TerminalNode {
	return s.GetTokens(SfplParserPRIORITY)
}

func (s *PruleContext) PRIORITY(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserPRIORITY, i)
}

func (s *PruleContext) AllSeverity() []ISeverityContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISeverityContext)(nil)).Elem())
	var tst = make([]ISeverityContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISeverityContext)
		}
	}

	return tst
}

func (s *PruleContext) Severity(i int) ISeverityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeverityContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISeverityContext)
}

func (s *PruleContext) AllTAGS() []antlr.TerminalNode {
	return s.GetTokens(SfplParserTAGS)
}

func (s *PruleContext) TAGS(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserTAGS, i)
}

func (s *PruleContext) AllTags() []ITagsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagsContext)(nil)).Elem())
	var tst = make([]ITagsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagsContext)
		}
	}

	return tst
}

func (s *PruleContext) Tags(i int) ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *PruleContext) AllPREFILTER() []antlr.TerminalNode {
	return s.GetTokens(SfplParserPREFILTER)
}

func (s *PruleContext) PREFILTER(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserPREFILTER, i)
}

func (s *PruleContext) AllPrefilter() []IPrefilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrefilterContext)(nil)).Elem())
	var tst = make([]IPrefilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrefilterContext)
		}
	}

	return tst
}

func (s *PruleContext) Prefilter(i int) IPrefilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrefilterContext)
}

func (s *PruleContext) AllENABLED() []antlr.TerminalNode {
	return s.GetTokens(SfplParserENABLED)
}

func (s *PruleContext) ENABLED(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserENABLED, i)
}

func (s *PruleContext) AllEnabled() []IEnabledContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnabledContext)(nil)).Elem())
	var tst = make([]IEnabledContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnabledContext)
		}
	}

	return tst
}

func (s *PruleContext) Enabled(i int) IEnabledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnabledContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnabledContext)
}

func (s *PruleContext) AllWARNEVTTYPE() []antlr.TerminalNode {
	return s.GetTokens(SfplParserWARNEVTTYPE)
}

func (s *PruleContext) WARNEVTTYPE(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserWARNEVTTYPE, i)
}

func (s *PruleContext) AllWarnevttype() []IWarnevttypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWarnevttypeContext)(nil)).Elem())
	var tst = make([]IWarnevttypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWarnevttypeContext)
		}
	}

	return tst
}

func (s *PruleContext) Warnevttype(i int) IWarnevttypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWarnevttypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWarnevttypeContext)
}

func (s *PruleContext) AllSKIPUNKNOWN() []antlr.TerminalNode {
	return s.GetTokens(SfplParserSKIPUNKNOWN)
}

func (s *PruleContext) SKIPUNKNOWN(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserSKIPUNKNOWN, i)
}

func (s *PruleContext) AllSkipunknown() []ISkipunknownContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISkipunknownContext)(nil)).Elem())
	var tst = make([]ISkipunknownContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISkipunknownContext)
		}
	}

	return tst
}

func (s *PruleContext) Skipunknown(i int) ISkipunknownContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISkipunknownContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISkipunknownContext)
}

func (s *PruleContext) AllACTION() []antlr.TerminalNode {
	return s.GetTokens(SfplParserACTION)
}

func (s *PruleContext) ACTION(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserACTION, i)
}

func (s *PruleContext) AllOUTPUT() []antlr.TerminalNode {
	return s.GetTokens(SfplParserOUTPUT)
}

func (s *PruleContext) OUTPUT(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserOUTPUT, i)
}

func (s *PruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPrule(s)
	}
}

func (s *PruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPrule(s)
	}
}

func (s *PruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPrule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Prule() (localctx IPruleContext) {
	localctx = NewPruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SfplParserRULE_prule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(76)
		p.Match(SfplParserRULE)
	}
	{
		p.SetState(77)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(78)
		p.Text()
	}
	{
		p.SetState(79)
		p.Match(SfplParserDESC)
	}
	{
		p.SetState(80)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(81)
		p.Text()
	}
	{
		p.SetState(82)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(83)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(84)
		p.Expression()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfplParserACTION)|(1<<SfplParserOUTPUT)|(1<<SfplParserPRIORITY)|(1<<SfplParserTAGS)|(1<<SfplParserPREFILTER)|(1<<SfplParserENABLED)|(1<<SfplParserWARNEVTTYPE)|(1<<SfplParserSKIPUNKNOWN))) != 0 {
		p.SetState(106)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SfplParserACTION, SfplParserOUTPUT:
			{
				p.SetState(85)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SfplParserACTION || _la == SfplParserOUTPUT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(86)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(87)
				p.Text()
			}

		case SfplParserPRIORITY:
			{
				p.SetState(88)
				p.Match(SfplParserPRIORITY)
			}
			{
				p.SetState(89)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(90)
				p.Severity()
			}

		case SfplParserTAGS:
			{
				p.SetState(91)
				p.Match(SfplParserTAGS)
			}
			{
				p.SetState(92)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(93)
				p.Tags()
			}

		case SfplParserPREFILTER:
			{
				p.SetState(94)
				p.Match(SfplParserPREFILTER)
			}
			{
				p.SetState(95)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(96)
				p.Prefilter()
			}

		case SfplParserENABLED:
			{
				p.SetState(97)
				p.Match(SfplParserENABLED)
			}
			{
				p.SetState(98)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(99)
				p.Enabled()
			}

		case SfplParserWARNEVTTYPE:
			{
				p.SetState(100)
				p.Match(SfplParserWARNEVTTYPE)
			}
			{
				p.SetState(101)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(102)
				p.Warnevttype()
			}

		case SfplParserSKIPUNKNOWN:
			{
				p.SetState(103)
				p.Match(SfplParserSKIPUNKNOWN)
			}
			{
				p.SetState(104)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(105)
				p.Skipunknown()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISruleContext is an interface to support dynamic dispatch.
type ISruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSruleContext differentiates from other interfaces.
	IsSruleContext()
}

type SruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySruleContext() *SruleContext {
	var p = new(SruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_srule
	return p
}

func (*SruleContext) IsSruleContext() {}

func NewSruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SruleContext {
	var p = new(SruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_srule

	return p
}

func (s *SruleContext) GetParser() antlr.Parser { return s.parser }

func (s *SruleContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *SruleContext) RULE() antlr.TerminalNode {
	return s.GetToken(SfplParserRULE, 0)
}

func (s *SruleContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *SruleContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *SruleContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *SruleContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *SruleContext) DESC() antlr.TerminalNode {
	return s.GetToken(SfplParserDESC, 0)
}

func (s *SruleContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *SruleContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SruleContext) AllPRIORITY() []antlr.TerminalNode {
	return s.GetTokens(SfplParserPRIORITY)
}

func (s *SruleContext) PRIORITY(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserPRIORITY, i)
}

func (s *SruleContext) AllSeverity() []ISeverityContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISeverityContext)(nil)).Elem())
	var tst = make([]ISeverityContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISeverityContext)
		}
	}

	return tst
}

func (s *SruleContext) Severity(i int) ISeverityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeverityContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISeverityContext)
}

func (s *SruleContext) AllTAGS() []antlr.TerminalNode {
	return s.GetTokens(SfplParserTAGS)
}

func (s *SruleContext) TAGS(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserTAGS, i)
}

func (s *SruleContext) AllTags() []ITagsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagsContext)(nil)).Elem())
	var tst = make([]ITagsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagsContext)
		}
	}

	return tst
}

func (s *SruleContext) Tags(i int) ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *SruleContext) AllPREFILTER() []antlr.TerminalNode {
	return s.GetTokens(SfplParserPREFILTER)
}

func (s *SruleContext) PREFILTER(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserPREFILTER, i)
}

func (s *SruleContext) AllPrefilter() []IPrefilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrefilterContext)(nil)).Elem())
	var tst = make([]IPrefilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrefilterContext)
		}
	}

	return tst
}

func (s *SruleContext) Prefilter(i int) IPrefilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrefilterContext)
}

func (s *SruleContext) AllENABLED() []antlr.TerminalNode {
	return s.GetTokens(SfplParserENABLED)
}

func (s *SruleContext) ENABLED(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserENABLED, i)
}

func (s *SruleContext) AllEnabled() []IEnabledContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnabledContext)(nil)).Elem())
	var tst = make([]IEnabledContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnabledContext)
		}
	}

	return tst
}

func (s *SruleContext) Enabled(i int) IEnabledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnabledContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnabledContext)
}

func (s *SruleContext) AllWARNEVTTYPE() []antlr.TerminalNode {
	return s.GetTokens(SfplParserWARNEVTTYPE)
}

func (s *SruleContext) WARNEVTTYPE(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserWARNEVTTYPE, i)
}

func (s *SruleContext) AllWarnevttype() []IWarnevttypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWarnevttypeContext)(nil)).Elem())
	var tst = make([]IWarnevttypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWarnevttypeContext)
		}
	}

	return tst
}

func (s *SruleContext) Warnevttype(i int) IWarnevttypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWarnevttypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWarnevttypeContext)
}

func (s *SruleContext) AllSKIPUNKNOWN() []antlr.TerminalNode {
	return s.GetTokens(SfplParserSKIPUNKNOWN)
}

func (s *SruleContext) SKIPUNKNOWN(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserSKIPUNKNOWN, i)
}

func (s *SruleContext) AllSkipunknown() []ISkipunknownContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISkipunknownContext)(nil)).Elem())
	var tst = make([]ISkipunknownContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISkipunknownContext)
		}
	}

	return tst
}

func (s *SruleContext) Skipunknown(i int) ISkipunknownContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISkipunknownContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISkipunknownContext)
}

func (s *SruleContext) AllACTION() []antlr.TerminalNode {
	return s.GetTokens(SfplParserACTION)
}

func (s *SruleContext) ACTION(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserACTION, i)
}

func (s *SruleContext) AllOUTPUT() []antlr.TerminalNode {
	return s.GetTokens(SfplParserOUTPUT)
}

func (s *SruleContext) OUTPUT(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserOUTPUT, i)
}

func (s *SruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterSrule(s)
	}
}

func (s *SruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitSrule(s)
	}
}

func (s *SruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitSrule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Srule() (localctx ISruleContext) {
	localctx = NewSruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SfplParserRULE_srule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(112)
		p.Match(SfplParserRULE)
	}
	{
		p.SetState(113)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(114)
		p.Text()
	}
	{
		p.SetState(115)
		p.Match(SfplParserDESC)
	}
	{
		p.SetState(116)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(117)
		p.Text()
	}
	{
		p.SetState(118)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(119)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(120)
		p.Expression()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfplParserACTION)|(1<<SfplParserOUTPUT)|(1<<SfplParserPRIORITY)|(1<<SfplParserTAGS)|(1<<SfplParserPREFILTER)|(1<<SfplParserENABLED)|(1<<SfplParserWARNEVTTYPE)|(1<<SfplParserSKIPUNKNOWN))) != 0 {
		p.SetState(142)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SfplParserACTION, SfplParserOUTPUT:
			{
				p.SetState(121)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SfplParserACTION || _la == SfplParserOUTPUT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(122)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(123)
				p.Text()
			}

		case SfplParserPRIORITY:
			{
				p.SetState(124)
				p.Match(SfplParserPRIORITY)
			}
			{
				p.SetState(125)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(126)
				p.Severity()
			}

		case SfplParserTAGS:
			{
				p.SetState(127)
				p.Match(SfplParserTAGS)
			}
			{
				p.SetState(128)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(129)
				p.Tags()
			}

		case SfplParserPREFILTER:
			{
				p.SetState(130)
				p.Match(SfplParserPREFILTER)
			}
			{
				p.SetState(131)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(132)
				p.Prefilter()
			}

		case SfplParserENABLED:
			{
				p.SetState(133)
				p.Match(SfplParserENABLED)
			}
			{
				p.SetState(134)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(135)
				p.Enabled()
			}

		case SfplParserWARNEVTTYPE:
			{
				p.SetState(136)
				p.Match(SfplParserWARNEVTTYPE)
			}
			{
				p.SetState(137)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(138)
				p.Warnevttype()
			}

		case SfplParserSKIPUNKNOWN:
			{
				p.SetState(139)
				p.Match(SfplParserSKIPUNKNOWN)
			}
			{
				p.SetState(140)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(141)
				p.Skipunknown()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPfilterContext is an interface to support dynamic dispatch.
type IPfilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPfilterContext differentiates from other interfaces.
	IsPfilterContext()
}

type PfilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPfilterContext() *PfilterContext {
	var p = new(PfilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_pfilter
	return p
}

func (*PfilterContext) IsPfilterContext() {}

func NewPfilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PfilterContext {
	var p = new(PfilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_pfilter

	return p
}

func (s *PfilterContext) GetParser() antlr.Parser { return s.parser }

func (s *PfilterContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PfilterContext) FILTER() antlr.TerminalNode {
	return s.GetToken(SfplParserFILTER, 0)
}

func (s *PfilterContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *PfilterContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *PfilterContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *PfilterContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *PfilterContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PfilterContext) ENABLED() antlr.TerminalNode {
	return s.GetToken(SfplParserENABLED, 0)
}

func (s *PfilterContext) Enabled() IEnabledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnabledContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnabledContext)
}

func (s *PfilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PfilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PfilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPfilter(s)
	}
}

func (s *PfilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPfilter(s)
	}
}

func (s *PfilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPfilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Pfilter() (localctx IPfilterContext) {
	localctx = NewPfilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SfplParserRULE_pfilter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(148)
		p.Match(SfplParserFILTER)
	}
	{
		p.SetState(149)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(150)
		p.Match(SfplParserID)
	}
	{
		p.SetState(151)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(152)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(153)
		p.Expression()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfplParserENABLED {
		{
			p.SetState(154)
			p.Match(SfplParserENABLED)
		}
		{
			p.SetState(155)
			p.Match(SfplParserDEF)
		}
		{
			p.SetState(156)
			p.Enabled()
		}

	}

	return localctx
}

// ISfilterContext is an interface to support dynamic dispatch.
type ISfilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSfilterContext differentiates from other interfaces.
	IsSfilterContext()
}

type SfilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySfilterContext() *SfilterContext {
	var p = new(SfilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_sfilter
	return p
}

func (*SfilterContext) IsSfilterContext() {}

func NewSfilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SfilterContext {
	var p = new(SfilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_sfilter

	return p
}

func (s *SfilterContext) GetParser() antlr.Parser { return s.parser }

func (s *SfilterContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *SfilterContext) FILTER() antlr.TerminalNode {
	return s.GetToken(SfplParserFILTER, 0)
}

func (s *SfilterContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *SfilterContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *SfilterContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *SfilterContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *SfilterContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SfilterContext) ENABLED() antlr.TerminalNode {
	return s.GetToken(SfplParserENABLED, 0)
}

func (s *SfilterContext) Enabled() IEnabledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnabledContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnabledContext)
}

func (s *SfilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SfilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SfilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterSfilter(s)
	}
}

func (s *SfilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitSfilter(s)
	}
}

func (s *SfilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitSfilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Sfilter() (localctx ISfilterContext) {
	localctx = NewSfilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SfplParserRULE_sfilter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(160)
		p.Match(SfplParserFILTER)
	}
	{
		p.SetState(161)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(162)
		p.Match(SfplParserID)
	}
	{
		p.SetState(163)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(164)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(165)
		p.Expression()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfplParserENABLED {
		{
			p.SetState(166)
			p.Match(SfplParserENABLED)
		}
		{
			p.SetState(167)
			p.Match(SfplParserDEF)
		}
		{
			p.SetState(168)
			p.Enabled()
		}

	}

	return localctx
}

// IPmacroContext is an interface to support dynamic dispatch.
type IPmacroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmacroContext differentiates from other interfaces.
	IsPmacroContext()
}

type PmacroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmacroContext() *PmacroContext {
	var p = new(PmacroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_pmacro
	return p
}

func (*PmacroContext) IsPmacroContext() {}

func NewPmacroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmacroContext {
	var p = new(PmacroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_pmacro

	return p
}

func (s *PmacroContext) GetParser() antlr.Parser { return s.parser }

func (s *PmacroContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PmacroContext) MACRO() antlr.TerminalNode {
	return s.GetToken(SfplParserMACRO, 0)
}

func (s *PmacroContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *PmacroContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *PmacroContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *PmacroContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *PmacroContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PmacroContext) FAPPEND() antlr.TerminalNode {
	return s.GetToken(SfplParserFAPPEND, 0)
}

func (s *PmacroContext) Fappend() IFappendContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFappendContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFappendContext)
}

func (s *PmacroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmacroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmacroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPmacro(s)
	}
}

func (s *PmacroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPmacro(s)
	}
}

func (s *PmacroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPmacro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Pmacro() (localctx IPmacroContext) {
	localctx = NewPmacroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SfplParserRULE_pmacro)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(172)
		p.Match(SfplParserMACRO)
	}
	{
		p.SetState(173)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(174)
		p.Match(SfplParserID)
	}
	{
		p.SetState(175)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(176)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(177)
		p.Expression()
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfplParserFAPPEND {
		{
			p.SetState(178)
			p.Match(SfplParserFAPPEND)
		}
		{
			p.SetState(179)
			p.Match(SfplParserDEF)
		}
		{
			p.SetState(180)
			p.Fappend()
		}

	}

	return localctx
}

// IPlistContext is an interface to support dynamic dispatch.
type IPlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlistContext differentiates from other interfaces.
	IsPlistContext()
}

type PlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlistContext() *PlistContext {
	var p = new(PlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_plist
	return p
}

func (*PlistContext) IsPlistContext() {}

func NewPlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlistContext {
	var p = new(PlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_plist

	return p
}

func (s *PlistContext) GetParser() antlr.Parser { return s.parser }

func (s *PlistContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PlistContext) LIST() antlr.TerminalNode {
	return s.GetToken(SfplParserLIST, 0)
}

func (s *PlistContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *PlistContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *PlistContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *PlistContext) ITEMS() antlr.TerminalNode {
	return s.GetToken(SfplParserITEMS, 0)
}

func (s *PlistContext) Items() IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *PlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPlist(s)
	}
}

func (s *PlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPlist(s)
	}
}

func (s *PlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Plist() (localctx IPlistContext) {
	localctx = NewPlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SfplParserRULE_plist)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(184)
		p.Match(SfplParserLIST)
	}
	{
		p.SetState(185)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(186)
		p.Match(SfplParserID)
	}
	{
		p.SetState(187)
		p.Match(SfplParserITEMS)
	}
	{
		p.SetState(188)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(189)
		p.Items()
	}

	return localctx
}

// IPreqContext is an interface to support dynamic dispatch.
type IPreqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPreqContext differentiates from other interfaces.
	IsPreqContext()
}

type PreqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreqContext() *PreqContext {
	var p = new(PreqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_preq
	return p
}

func (*PreqContext) IsPreqContext() {}

func NewPreqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreqContext {
	var p = new(PreqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_preq

	return p
}

func (s *PreqContext) GetParser() antlr.Parser { return s.parser }

func (s *PreqContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PreqContext) REQ() antlr.TerminalNode {
	return s.GetToken(SfplParserREQ, 0)
}

func (s *PreqContext) DEF() antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, 0)
}

func (s *PreqContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *PreqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPreq(s)
	}
}

func (s *PreqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPreq(s)
	}
}

func (s *PreqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPreq(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Preq() (localctx IPreqContext) {
	localctx = NewPreqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SfplParserRULE_preq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(192)
		p.Match(SfplParserREQ)
	}
	{
		p.SetState(193)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(194)
		p.Atom()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Or_expression() IOr_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SfplParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Or_expression()
	}

	return localctx
}

// IOr_expressionContext is an interface to support dynamic dispatch.
type IOr_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_expressionContext differentiates from other interfaces.
	IsOr_expressionContext()
}

type Or_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_expressionContext() *Or_expressionContext {
	var p = new(Or_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_or_expression
	return p
}

func (*Or_expressionContext) IsOr_expressionContext() {}

func NewOr_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_expressionContext {
	var p = new(Or_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_or_expression

	return p
}

func (s *Or_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_expressionContext) AllAnd_expression() []IAnd_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnd_expressionContext)(nil)).Elem())
	var tst = make([]IAnd_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnd_expressionContext)
		}
	}

	return tst
}

func (s *Or_expressionContext) And_expression(i int) IAnd_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnd_expressionContext)
}

func (s *Or_expressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(SfplParserOR)
}

func (s *Or_expressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserOR, i)
}

func (s *Or_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterOr_expression(s)
	}
}

func (s *Or_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitOr_expression(s)
	}
}

func (s *Or_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitOr_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Or_expression() (localctx IOr_expressionContext) {
	localctx = NewOr_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SfplParserRULE_or_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.And_expression()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserOR {
		{
			p.SetState(199)
			p.Match(SfplParserOR)
		}
		{
			p.SetState(200)
			p.And_expression()
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAnd_expressionContext is an interface to support dynamic dispatch.
type IAnd_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_expressionContext differentiates from other interfaces.
	IsAnd_expressionContext()
}

type And_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_expressionContext() *And_expressionContext {
	var p = new(And_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_and_expression
	return p
}

func (*And_expressionContext) IsAnd_expressionContext() {}

func NewAnd_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_expressionContext {
	var p = new(And_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_and_expression

	return p
}

func (s *And_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *And_expressionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *And_expressionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *And_expressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(SfplParserAND)
}

func (s *And_expressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserAND, i)
}

func (s *And_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterAnd_expression(s)
	}
}

func (s *And_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitAnd_expression(s)
	}
}

func (s *And_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitAnd_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) And_expression() (localctx IAnd_expressionContext) {
	localctx = NewAnd_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SfplParserRULE_and_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Term()
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserAND {
		{
			p.SetState(207)
			p.Match(SfplParserAND)
		}
		{
			p.SetState(208)
			p.Term()
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *TermContext) NOT() antlr.TerminalNode {
	return s.GetToken(SfplParserNOT, 0)
}

func (s *TermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *TermContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *TermContext) Unary_operator() IUnary_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_operatorContext)
}

func (s *TermContext) Binary_operator() IBinary_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_operatorContext)
}

func (s *TermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SfplParserLPAREN, 0)
}

func (s *TermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SfplParserRPAREN, 0)
}

func (s *TermContext) IN() antlr.TerminalNode {
	return s.GetToken(SfplParserIN, 0)
}

func (s *TermContext) PMATCH() antlr.TerminalNode {
	return s.GetToken(SfplParserPMATCH, 0)
}

func (s *TermContext) AllItems() []IItemsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IItemsContext)(nil)).Elem())
	var tst = make([]IItemsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IItemsContext)
		}
	}

	return tst
}

func (s *TermContext) Items(i int) IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *TermContext) AllLISTSEP() []antlr.TerminalNode {
	return s.GetTokens(SfplParserLISTSEP)
}

func (s *TermContext) LISTSEP(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserLISTSEP, i)
}

func (s *TermContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SfplParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(SfplParserNOT)
		}
		{
			p.SetState(216)
			p.Term()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(217)
			p.Atom()
		}
		{
			p.SetState(218)
			p.Unary_operator()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.Atom()
		}
		{
			p.SetState(221)
			p.Binary_operator()
		}
		{
			p.SetState(222)
			p.Atom()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(224)
			p.Atom()
		}
		{
			p.SetState(225)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SfplParserIN || _la == SfplParserPMATCH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(226)
			p.Match(SfplParserLPAREN)
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SfplParserLT, SfplParserGT, SfplParserID, SfplParserNUMBER, SfplParserPATH, SfplParserSTRING, SfplParserTAG:
			{
				p.SetState(227)
				p.Atom()
			}

		case SfplParserLBRACK:
			{
				p.SetState(228)
				p.Items()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SfplParserLISTSEP {
			{
				p.SetState(231)
				p.Match(SfplParserLISTSEP)
			}
			p.SetState(234)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case SfplParserLT, SfplParserGT, SfplParserID, SfplParserNUMBER, SfplParserPATH, SfplParserSTRING, SfplParserTAG:
				{
					p.SetState(232)
					p.Atom()
				}

			case SfplParserLBRACK:
				{
					p.SetState(233)
					p.Items()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(241)
			p.Match(SfplParserRPAREN)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(243)
			p.Match(SfplParserLPAREN)
		}
		{
			p.SetState(244)
			p.Expression()
		}
		{
			p.SetState(245)
			p.Match(SfplParserRPAREN)
		}

	}

	return localctx
}

// IItemsContext is an interface to support dynamic dispatch.
type IItemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemsContext differentiates from other interfaces.
	IsItemsContext()
}

type ItemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemsContext() *ItemsContext {
	var p = new(ItemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_items
	return p
}

func (*ItemsContext) IsItemsContext() {}

func NewItemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemsContext {
	var p = new(ItemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_items

	return p
}

func (s *ItemsContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemsContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(SfplParserLBRACK, 0)
}

func (s *ItemsContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(SfplParserRBRACK, 0)
}

func (s *ItemsContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *ItemsContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ItemsContext) AllLISTSEP() []antlr.TerminalNode {
	return s.GetTokens(SfplParserLISTSEP)
}

func (s *ItemsContext) LISTSEP(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserLISTSEP, i)
}

func (s *ItemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterItems(s)
	}
}

func (s *ItemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitItems(s)
	}
}

func (s *ItemsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitItems(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Items() (localctx IItemsContext) {
	localctx = NewItemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SfplParserRULE_items)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(SfplParserLBRACK)
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(SfplParserLT-22))|(1<<(SfplParserGT-22))|(1<<(SfplParserID-22))|(1<<(SfplParserNUMBER-22))|(1<<(SfplParserPATH-22))|(1<<(SfplParserSTRING-22))|(1<<(SfplParserTAG-22)))) != 0 {
		{
			p.SetState(250)
			p.Atom()
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(251)
					p.Match(SfplParserLISTSEP)
				}
				{
					p.SetState(252)
					p.Atom()
				}

			}
			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}

	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfplParserLISTSEP {
		{
			p.SetState(260)
			p.Match(SfplParserLISTSEP)
		}

	}
	{
		p.SetState(263)
		p.Match(SfplParserRBRACK)
	}

	return localctx
}

// ITagsContext is an interface to support dynamic dispatch.
type ITagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagsContext differentiates from other interfaces.
	IsTagsContext()
}

type TagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagsContext() *TagsContext {
	var p = new(TagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_tags
	return p
}

func (*TagsContext) IsTagsContext() {}

func NewTagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagsContext {
	var p = new(TagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_tags

	return p
}

func (s *TagsContext) GetParser() antlr.Parser { return s.parser }

func (s *TagsContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(SfplParserLBRACK, 0)
}

func (s *TagsContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(SfplParserRBRACK, 0)
}

func (s *TagsContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *TagsContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *TagsContext) AllLISTSEP() []antlr.TerminalNode {
	return s.GetTokens(SfplParserLISTSEP)
}

func (s *TagsContext) LISTSEP(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserLISTSEP, i)
}

func (s *TagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterTags(s)
	}
}

func (s *TagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitTags(s)
	}
}

func (s *TagsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitTags(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Tags() (localctx ITagsContext) {
	localctx = NewTagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SfplParserRULE_tags)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(SfplParserLBRACK)
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(SfplParserLT-22))|(1<<(SfplParserGT-22))|(1<<(SfplParserID-22))|(1<<(SfplParserNUMBER-22))|(1<<(SfplParserPATH-22))|(1<<(SfplParserSTRING-22))|(1<<(SfplParserTAG-22)))) != 0 {
		{
			p.SetState(266)
			p.Atom()
		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(267)
					p.Match(SfplParserLISTSEP)
				}
				{
					p.SetState(268)
					p.Atom()
				}

			}
			p.SetState(273)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
		}

	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfplParserLISTSEP {
		{
			p.SetState(276)
			p.Match(SfplParserLISTSEP)
		}

	}
	{
		p.SetState(279)
		p.Match(SfplParserRBRACK)
	}

	return localctx
}

// IPrefilterContext is an interface to support dynamic dispatch.
type IPrefilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefilterContext differentiates from other interfaces.
	IsPrefilterContext()
}

type PrefilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefilterContext() *PrefilterContext {
	var p = new(PrefilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_prefilter
	return p
}

func (*PrefilterContext) IsPrefilterContext() {}

func NewPrefilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefilterContext {
	var p = new(PrefilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_prefilter

	return p
}

func (s *PrefilterContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefilterContext) Items() IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *PrefilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPrefilter(s)
	}
}

func (s *PrefilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPrefilter(s)
	}
}

func (s *PrefilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPrefilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Prefilter() (localctx IPrefilterContext) {
	localctx = NewPrefilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SfplParserRULE_prefilter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Items()
	}

	return localctx
}

// ISeverityContext is an interface to support dynamic dispatch.
type ISeverityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeverityContext differentiates from other interfaces.
	IsSeverityContext()
}

type SeverityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeverityContext() *SeverityContext {
	var p = new(SeverityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_severity
	return p
}

func (*SeverityContext) IsSeverityContext() {}

func NewSeverityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeverityContext {
	var p = new(SeverityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_severity

	return p
}

func (s *SeverityContext) GetParser() antlr.Parser { return s.parser }

func (s *SeverityContext) SEVERITY() antlr.TerminalNode {
	return s.GetToken(SfplParserSEVERITY, 0)
}

func (s *SeverityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeverityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeverityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterSeverity(s)
	}
}

func (s *SeverityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitSeverity(s)
	}
}

func (s *SeverityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitSeverity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Severity() (localctx ISeverityContext) {
	localctx = NewSeverityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SfplParserRULE_severity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(SfplParserSEVERITY)
	}

	return localctx
}

// IEnabledContext is an interface to support dynamic dispatch.
type IEnabledContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnabledContext differentiates from other interfaces.
	IsEnabledContext()
}

type EnabledContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnabledContext() *EnabledContext {
	var p = new(EnabledContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_enabled
	return p
}

func (*EnabledContext) IsEnabledContext() {}

func NewEnabledContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnabledContext {
	var p = new(EnabledContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_enabled

	return p
}

func (s *EnabledContext) GetParser() antlr.Parser { return s.parser }

func (s *EnabledContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *EnabledContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnabledContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnabledContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterEnabled(s)
	}
}

func (s *EnabledContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitEnabled(s)
	}
}

func (s *EnabledContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitEnabled(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Enabled() (localctx IEnabledContext) {
	localctx = NewEnabledContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SfplParserRULE_enabled)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Atom()
	}

	return localctx
}

// IWarnevttypeContext is an interface to support dynamic dispatch.
type IWarnevttypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWarnevttypeContext differentiates from other interfaces.
	IsWarnevttypeContext()
}

type WarnevttypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWarnevttypeContext() *WarnevttypeContext {
	var p = new(WarnevttypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_warnevttype
	return p
}

func (*WarnevttypeContext) IsWarnevttypeContext() {}

func NewWarnevttypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WarnevttypeContext {
	var p = new(WarnevttypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_warnevttype

	return p
}

func (s *WarnevttypeContext) GetParser() antlr.Parser { return s.parser }

func (s *WarnevttypeContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *WarnevttypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WarnevttypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WarnevttypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterWarnevttype(s)
	}
}

func (s *WarnevttypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitWarnevttype(s)
	}
}

func (s *WarnevttypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitWarnevttype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Warnevttype() (localctx IWarnevttypeContext) {
	localctx = NewWarnevttypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SfplParserRULE_warnevttype)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Atom()
	}

	return localctx
}

// ISkipunknownContext is an interface to support dynamic dispatch.
type ISkipunknownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSkipunknownContext differentiates from other interfaces.
	IsSkipunknownContext()
}

type SkipunknownContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkipunknownContext() *SkipunknownContext {
	var p = new(SkipunknownContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_skipunknown
	return p
}

func (*SkipunknownContext) IsSkipunknownContext() {}

func NewSkipunknownContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipunknownContext {
	var p = new(SkipunknownContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_skipunknown

	return p
}

func (s *SkipunknownContext) GetParser() antlr.Parser { return s.parser }

func (s *SkipunknownContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *SkipunknownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipunknownContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SkipunknownContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterSkipunknown(s)
	}
}

func (s *SkipunknownContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitSkipunknown(s)
	}
}

func (s *SkipunknownContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitSkipunknown(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Skipunknown() (localctx ISkipunknownContext) {
	localctx = NewSkipunknownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SfplParserRULE_skipunknown)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Atom()
	}

	return localctx
}

// IFappendContext is an interface to support dynamic dispatch.
type IFappendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFappendContext differentiates from other interfaces.
	IsFappendContext()
}

type FappendContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFappendContext() *FappendContext {
	var p = new(FappendContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_fappend
	return p
}

func (*FappendContext) IsFappendContext() {}

func NewFappendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FappendContext {
	var p = new(FappendContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_fappend

	return p
}

func (s *FappendContext) GetParser() antlr.Parser { return s.parser }

func (s *FappendContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *FappendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FappendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FappendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterFappend(s)
	}
}

func (s *FappendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitFappend(s)
	}
}

func (s *FappendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitFappend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Fappend() (localctx IFappendContext) {
	localctx = NewFappendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SfplParserRULE_fappend)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Atom()
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SfplParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(SfplParserID)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *AtomContext) PATH() antlr.TerminalNode {
	return s.GetToken(SfplParserPATH, 0)
}

func (s *AtomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SfplParserNUMBER, 0)
}

func (s *AtomContext) TAG() antlr.TerminalNode {
	return s.GetToken(SfplParserTAG, 0)
}

func (s *AtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(SfplParserSTRING, 0)
}

func (s *AtomContext) LT() antlr.TerminalNode {
	return s.GetToken(SfplParserLT, 0)
}

func (s *AtomContext) GT() antlr.TerminalNode {
	return s.GetToken(SfplParserGT, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SfplParserRULE_atom)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(SfplParserLT-22))|(1<<(SfplParserGT-22))|(1<<(SfplParserID-22))|(1<<(SfplParserNUMBER-22))|(1<<(SfplParserPATH-22))|(1<<(SfplParserSTRING-22))|(1<<(SfplParserTAG-22)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }
func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitText(s)
	}
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SfplParserRULE_text)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(297)

			if !(!(p.GetCurrentToken().GetText() == "desc" ||
				p.GetCurrentToken().GetText() == "condition" ||
				p.GetCurrentToken().GetText() == "action" ||
				p.GetCurrentToken().GetText() == "output" ||
				p.GetCurrentToken().GetText() == "priority" ||
				p.GetCurrentToken().GetText() == "tags" ||
				p.GetCurrentToken().GetText() == "prefilter" ||
				p.GetCurrentToken().GetText() == "enabled" ||
				p.GetCurrentToken().GetText() == "warn_evttypes" ||
				p.GetCurrentToken().GetText() == "skip-if-unknown-filter" ||
				p.GetCurrentToken().GetText() == "append")) {
				panic(antlr.NewFailedPredicateException(p, "!(p.GetCurrentToken().GetText() == \"desc\" ||\n\t      p.GetCurrentToken().GetText() == \"condition\" ||\n\t      p.GetCurrentToken().GetText() == \"action\" ||\n\t      p.GetCurrentToken().GetText() == \"output\" ||\n\t      p.GetCurrentToken().GetText() == \"priority\" ||\n\t      p.GetCurrentToken().GetText() == \"tags\" ||\n\t\t  p.GetCurrentToken().GetText() == \"prefilter\" ||\n\t\t  p.GetCurrentToken().GetText() == \"enabled\" ||\n\t\t  p.GetCurrentToken().GetText() == \"warn_evttypes\" ||\n\t\t  p.GetCurrentToken().GetText() == \"skip-if-unknown-filter\" ||\n\t\t  p.GetCurrentToken().GetText() == \"append\")", ""))
			}
			p.SetState(298)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IBinary_operatorContext is an interface to support dynamic dispatch.
type IBinary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_operatorContext differentiates from other interfaces.
	IsBinary_operatorContext()
}

type Binary_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_operatorContext() *Binary_operatorContext {
	var p = new(Binary_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_binary_operator
	return p
}

func (*Binary_operatorContext) IsBinary_operatorContext() {}

func NewBinary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_operatorContext {
	var p = new(Binary_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_binary_operator

	return p
}

func (s *Binary_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Binary_operatorContext) LT() antlr.TerminalNode {
	return s.GetToken(SfplParserLT, 0)
}

func (s *Binary_operatorContext) LE() antlr.TerminalNode {
	return s.GetToken(SfplParserLE, 0)
}

func (s *Binary_operatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SfplParserGT, 0)
}

func (s *Binary_operatorContext) GE() antlr.TerminalNode {
	return s.GetToken(SfplParserGE, 0)
}

func (s *Binary_operatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(SfplParserEQ, 0)
}

func (s *Binary_operatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SfplParserNEQ, 0)
}

func (s *Binary_operatorContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(SfplParserCONTAINS, 0)
}

func (s *Binary_operatorContext) ICONTAINS() antlr.TerminalNode {
	return s.GetToken(SfplParserICONTAINS, 0)
}

func (s *Binary_operatorContext) STARTSWITH() antlr.TerminalNode {
	return s.GetToken(SfplParserSTARTSWITH, 0)
}

func (s *Binary_operatorContext) ENDSWITH() antlr.TerminalNode {
	return s.GetToken(SfplParserENDSWITH, 0)
}

func (s *Binary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Binary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterBinary_operator(s)
	}
}

func (s *Binary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitBinary_operator(s)
	}
}

func (s *Binary_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitBinary_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Binary_operator() (localctx IBinary_operatorContext) {
	localctx = NewBinary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SfplParserRULE_binary_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(SfplParserLT-22))|(1<<(SfplParserLE-22))|(1<<(SfplParserGT-22))|(1<<(SfplParserGE-22))|(1<<(SfplParserEQ-22))|(1<<(SfplParserNEQ-22))|(1<<(SfplParserCONTAINS-22))|(1<<(SfplParserICONTAINS-22))|(1<<(SfplParserSTARTSWITH-22))|(1<<(SfplParserENDSWITH-22)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnary_operatorContext is an interface to support dynamic dispatch.
type IUnary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnary_operatorContext differentiates from other interfaces.
	IsUnary_operatorContext()
}

type Unary_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_operatorContext() *Unary_operatorContext {
	var p = new(Unary_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_unary_operator
	return p
}

func (*Unary_operatorContext) IsUnary_operatorContext() {}

func NewUnary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_operatorContext {
	var p = new(Unary_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_unary_operator

	return p
}

func (s *Unary_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_operatorContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(SfplParserEXISTS, 0)
}

func (s *Unary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterUnary_operator(s)
	}
}

func (s *Unary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitUnary_operator(s)
	}
}

func (s *Unary_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitUnary_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Unary_operator() (localctx IUnary_operatorContext) {
	localctx = NewUnary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SfplParserRULE_unary_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(SfplParserEXISTS)
	}

	return localctx
}

func (p *SfplParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 23:
		var t *TextContext = nil
		if localctx != nil {
			t = localctx.(*TextContext)
		}
		return p.Text_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SfplParser) Text_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return !(p.GetCurrentToken().GetText() == "desc" ||
			p.GetCurrentToken().GetText() == "condition" ||
			p.GetCurrentToken().GetText() == "action" ||
			p.GetCurrentToken().GetText() == "output" ||
			p.GetCurrentToken().GetText() == "priority" ||
			p.GetCurrentToken().GetText() == "tags" ||
			p.GetCurrentToken().GetText() == "prefilter" ||
			p.GetCurrentToken().GetText() == "enabled" ||
			p.GetCurrentToken().GetText() == "warn_evttypes" ||
			p.GetCurrentToken().GetText() == "skip-if-unknown-filter" ||
			p.GetCurrentToken().GetText() == "append")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
