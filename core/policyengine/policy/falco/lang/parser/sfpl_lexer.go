// Code generated from Sfpl.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 56, 709,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42,
	3, 43, 3, 43, 7, 43, 433, 10, 43, 12, 43, 14, 43, 436, 11, 43, 3, 43, 5,
	43, 439, 10, 43, 3, 44, 3, 44, 5, 44, 443, 10, 44, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 5, 45, 461, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 5, 46, 534, 10, 46, 3, 47, 3, 47, 3, 47, 5, 47, 539,
	10, 47, 3, 47, 3, 47, 3, 47, 5, 47, 544, 10, 47, 3, 47, 3, 47, 7, 47, 548,
	10, 47, 12, 47, 14, 47, 551, 11, 47, 3, 47, 3, 47, 3, 47, 7, 47, 556, 10,
	47, 12, 47, 14, 47, 559, 11, 47, 3, 48, 6, 48, 562, 10, 48, 13, 48, 14,
	48, 563, 3, 48, 3, 48, 6, 48, 568, 10, 48, 13, 48, 14, 48, 569, 5, 48,
	572, 10, 48, 3, 49, 3, 49, 7, 49, 576, 10, 49, 12, 49, 14, 49, 579, 11,
	49, 3, 50, 3, 50, 3, 50, 5, 50, 584, 10, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 5, 50, 591, 10, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 5, 50, 600, 10, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 5, 50, 610, 10, 50, 3, 50, 3, 50, 3, 50, 5, 50, 615, 10, 50, 3,
	51, 3, 51, 3, 51, 3, 51, 3, 52, 7, 52, 622, 10, 52, 12, 52, 14, 52, 625,
	11, 52, 3, 53, 3, 53, 3, 53, 3, 53, 5, 53, 631, 10, 53, 3, 54, 6, 54, 634,
	10, 54, 13, 54, 14, 54, 635, 3, 54, 3, 54, 3, 55, 5, 55, 641, 10, 55, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 7, 56, 649, 10, 56, 12, 56, 14,
	56, 652, 11, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59,
	3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3,
	65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70,
	3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3, 74, 3, 75, 3,
	75, 3, 76, 3, 76, 3, 77, 3, 77, 3, 78, 3, 78, 3, 79, 3, 79, 3, 80, 3, 80,
	3, 81, 3, 81, 3, 82, 3, 82, 3, 83, 3, 83, 3, 623, 2, 84, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81,
	42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99,
	51, 101, 52, 103, 2, 105, 2, 107, 53, 109, 54, 111, 55, 113, 56, 115, 2,
	117, 2, 119, 2, 121, 2, 123, 2, 125, 2, 127, 2, 129, 2, 131, 2, 133, 2,
	135, 2, 137, 2, 139, 2, 141, 2, 143, 2, 145, 2, 147, 2, 149, 2, 151, 2,
	153, 2, 155, 2, 157, 2, 159, 2, 161, 2, 163, 2, 165, 2, 3, 2, 34, 6, 2,
	50, 59, 67, 92, 97, 97, 99, 124, 7, 2, 47, 48, 50, 59, 67, 92, 97, 97,
	99, 124, 5, 2, 48, 59, 67, 92, 99, 124, 7, 2, 44, 44, 47, 59, 67, 92, 97,
	97, 99, 124, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 14, 15, 34, 34, 4, 2,
	67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70,
	70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73,
	73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76,
	76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79,
	79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82,
	82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85,
	85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88,
	88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91,
	91, 123, 123, 4, 2, 92, 92, 124, 124, 2, 715, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2,
	109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 3, 167, 3, 2,
	2, 2, 5, 172, 3, 2, 2, 2, 7, 179, 3, 2, 2, 2, 9, 184, 3, 2, 2, 2, 11, 190,
	3, 2, 2, 2, 13, 195, 3, 2, 2, 2, 15, 200, 3, 2, 2, 2, 17, 206, 3, 2, 2,
	2, 19, 216, 3, 2, 2, 2, 21, 221, 3, 2, 2, 2, 23, 229, 3, 2, 2, 2, 25, 236,
	3, 2, 2, 2, 27, 245, 3, 2, 2, 2, 29, 250, 3, 2, 2, 2, 31, 260, 3, 2, 2,
	2, 33, 268, 3, 2, 2, 2, 35, 282, 3, 2, 2, 2, 37, 305, 3, 2, 2, 2, 39, 312,
	3, 2, 2, 2, 41, 336, 3, 2, 2, 2, 43, 340, 3, 2, 2, 2, 45, 343, 3, 2, 2,
	2, 47, 347, 3, 2, 2, 2, 49, 349, 3, 2, 2, 2, 51, 352, 3, 2, 2, 2, 53, 354,
	3, 2, 2, 2, 55, 357, 3, 2, 2, 2, 57, 359, 3, 2, 2, 2, 59, 362, 3, 2, 2,
	2, 61, 365, 3, 2, 2, 2, 63, 374, 3, 2, 2, 2, 65, 384, 3, 2, 2, 2, 67, 395,
	3, 2, 2, 2, 69, 404, 3, 2, 2, 2, 71, 411, 3, 2, 2, 2, 73, 418, 3, 2, 2,
	2, 75, 420, 3, 2, 2, 2, 77, 422, 3, 2, 2, 2, 79, 424, 3, 2, 2, 2, 81, 426,
	3, 2, 2, 2, 83, 428, 3, 2, 2, 2, 85, 430, 3, 2, 2, 2, 87, 442, 3, 2, 2,
	2, 89, 460, 3, 2, 2, 2, 91, 533, 3, 2, 2, 2, 93, 535, 3, 2, 2, 2, 95, 561,
	3, 2, 2, 2, 97, 573, 3, 2, 2, 2, 99, 614, 3, 2, 2, 2, 101, 616, 3, 2, 2,
	2, 103, 623, 3, 2, 2, 2, 105, 630, 3, 2, 2, 2, 107, 633, 3, 2, 2, 2, 109,
	640, 3, 2, 2, 2, 111, 646, 3, 2, 2, 2, 113, 655, 3, 2, 2, 2, 115, 657,
	3, 2, 2, 2, 117, 659, 3, 2, 2, 2, 119, 661, 3, 2, 2, 2, 121, 663, 3, 2,
	2, 2, 123, 665, 3, 2, 2, 2, 125, 667, 3, 2, 2, 2, 127, 669, 3, 2, 2, 2,
	129, 671, 3, 2, 2, 2, 131, 673, 3, 2, 2, 2, 133, 675, 3, 2, 2, 2, 135,
	677, 3, 2, 2, 2, 137, 679, 3, 2, 2, 2, 139, 681, 3, 2, 2, 2, 141, 683,
	3, 2, 2, 2, 143, 685, 3, 2, 2, 2, 145, 687, 3, 2, 2, 2, 147, 689, 3, 2,
	2, 2, 149, 691, 3, 2, 2, 2, 151, 693, 3, 2, 2, 2, 153, 695, 3, 2, 2, 2,
	155, 697, 3, 2, 2, 2, 157, 699, 3, 2, 2, 2, 159, 701, 3, 2, 2, 2, 161,
	703, 3, 2, 2, 2, 163, 705, 3, 2, 2, 2, 165, 707, 3, 2, 2, 2, 167, 168,
	7, 116, 2, 2, 168, 169, 7, 119, 2, 2, 169, 170, 7, 110, 2, 2, 170, 171,
	7, 103, 2, 2, 171, 4, 3, 2, 2, 2, 172, 173, 7, 104, 2, 2, 173, 174, 7,
	107, 2, 2, 174, 175, 7, 110, 2, 2, 175, 176, 7, 118, 2, 2, 176, 177, 7,
	103, 2, 2, 177, 178, 7, 116, 2, 2, 178, 6, 3, 2, 2, 2, 179, 180, 7, 102,
	2, 2, 180, 181, 7, 116, 2, 2, 181, 182, 7, 113, 2, 2, 182, 183, 7, 114,
	2, 2, 183, 8, 3, 2, 2, 2, 184, 185, 7, 111, 2, 2, 185, 186, 7, 99, 2, 2,
	186, 187, 7, 101, 2, 2, 187, 188, 7, 116, 2, 2, 188, 189, 7, 113, 2, 2,
	189, 10, 3, 2, 2, 2, 190, 191, 7, 110, 2, 2, 191, 192, 7, 107, 2, 2, 192,
	193, 7, 117, 2, 2, 193, 194, 7, 118, 2, 2, 194, 12, 3, 2, 2, 2, 195, 196,
	7, 112, 2, 2, 196, 197, 7, 99, 2, 2, 197, 198, 7, 111, 2, 2, 198, 199,
	7, 103, 2, 2, 199, 14, 3, 2, 2, 2, 200, 201, 7, 107, 2, 2, 201, 202, 7,
	118, 2, 2, 202, 203, 7, 103, 2, 2, 203, 204, 7, 111, 2, 2, 204, 205, 7,
	117, 2, 2, 205, 16, 3, 2, 2, 2, 206, 207, 7, 101, 2, 2, 207, 208, 7, 113,
	2, 2, 208, 209, 7, 112, 2, 2, 209, 210, 7, 102, 2, 2, 210, 211, 7, 107,
	2, 2, 211, 212, 7, 118, 2, 2, 212, 213, 7, 107, 2, 2, 213, 214, 7, 113,
	2, 2, 214, 215, 7, 112, 2, 2, 215, 18, 3, 2, 2, 2, 216, 217, 7, 102, 2,
	2, 217, 218, 7, 103, 2, 2, 218, 219, 7, 117, 2, 2, 219, 220, 7, 101, 2,
	2, 220, 20, 3, 2, 2, 2, 221, 222, 7, 99, 2, 2, 222, 223, 7, 101, 2, 2,
	223, 224, 7, 118, 2, 2, 224, 225, 7, 107, 2, 2, 225, 226, 7, 113, 2, 2,
	226, 227, 7, 112, 2, 2, 227, 228, 7, 117, 2, 2, 228, 22, 3, 2, 2, 2, 229,
	230, 7, 113, 2, 2, 230, 231, 7, 119, 2, 2, 231, 232, 7, 118, 2, 2, 232,
	233, 7, 114, 2, 2, 233, 234, 7, 119, 2, 2, 234, 235, 7, 118, 2, 2, 235,
	24, 3, 2, 2, 2, 236, 237, 7, 114, 2, 2, 237, 238, 7, 116, 2, 2, 238, 239,
	7, 107, 2, 2, 239, 240, 7, 113, 2, 2, 240, 241, 7, 116, 2, 2, 241, 242,
	7, 107, 2, 2, 242, 243, 7, 118, 2, 2, 243, 244, 7, 123, 2, 2, 244, 26,
	3, 2, 2, 2, 245, 246, 7, 118, 2, 2, 246, 247, 7, 99, 2, 2, 247, 248, 7,
	105, 2, 2, 248, 249, 7, 117, 2, 2, 249, 28, 3, 2, 2, 2, 250, 251, 7, 114,
	2, 2, 251, 252, 7, 116, 2, 2, 252, 253, 7, 103, 2, 2, 253, 254, 7, 104,
	2, 2, 254, 255, 7, 107, 2, 2, 255, 256, 7, 110, 2, 2, 256, 257, 7, 118,
	2, 2, 257, 258, 7, 103, 2, 2, 258, 259, 7, 116, 2, 2, 259, 30, 3, 2, 2,
	2, 260, 261, 7, 103, 2, 2, 261, 262, 7, 112, 2, 2, 262, 263, 7, 99, 2,
	2, 263, 264, 7, 100, 2, 2, 264, 265, 7, 110, 2, 2, 265, 266, 7, 103, 2,
	2, 266, 267, 7, 102, 2, 2, 267, 32, 3, 2, 2, 2, 268, 269, 7, 121, 2, 2,
	269, 270, 7, 99, 2, 2, 270, 271, 7, 116, 2, 2, 271, 272, 7, 112, 2, 2,
	272, 273, 7, 97, 2, 2, 273, 274, 7, 103, 2, 2, 274, 275, 7, 120, 2, 2,
	275, 276, 7, 118, 2, 2, 276, 277, 7, 118, 2, 2, 277, 278, 7, 123, 2, 2,
	278, 279, 7, 114, 2, 2, 279, 280, 7, 103, 2, 2, 280, 281, 7, 117, 2, 2,
	281, 34, 3, 2, 2, 2, 282, 283, 7, 117, 2, 2, 283, 284, 7, 109, 2, 2, 284,
	285, 7, 107, 2, 2, 285, 286, 7, 114, 2, 2, 286, 287, 7, 47, 2, 2, 287,
	288, 7, 107, 2, 2, 288, 289, 7, 104, 2, 2, 289, 290, 7, 47, 2, 2, 290,
	291, 7, 119, 2, 2, 291, 292, 7, 112, 2, 2, 292, 293, 7, 109, 2, 2, 293,
	294, 7, 112, 2, 2, 294, 295, 7, 113, 2, 2, 295, 296, 7, 121, 2, 2, 296,
	297, 7, 112, 2, 2, 297, 298, 7, 47, 2, 2, 298, 299, 7, 104, 2, 2, 299,
	300, 7, 107, 2, 2, 300, 301, 7, 110, 2, 2, 301, 302, 7, 118, 2, 2, 302,
	303, 7, 103, 2, 2, 303, 304, 7, 116, 2, 2, 304, 36, 3, 2, 2, 2, 305, 306,
	7, 99, 2, 2, 306, 307, 7, 114, 2, 2, 307, 308, 7, 114, 2, 2, 308, 309,
	7, 103, 2, 2, 309, 310, 7, 112, 2, 2, 310, 311, 7, 102, 2, 2, 311, 38,
	3, 2, 2, 2, 312, 313, 7, 116, 2, 2, 313, 314, 7, 103, 2, 2, 314, 315, 7,
	115, 2, 2, 315, 316, 7, 119, 2, 2, 316, 317, 7, 107, 2, 2, 317, 318, 7,
	116, 2, 2, 318, 319, 7, 103, 2, 2, 319, 320, 7, 102, 2, 2, 320, 321, 7,
	97, 2, 2, 321, 322, 7, 103, 2, 2, 322, 323, 7, 112, 2, 2, 323, 324, 7,
	105, 2, 2, 324, 325, 7, 107, 2, 2, 325, 326, 7, 112, 2, 2, 326, 327, 7,
	103, 2, 2, 327, 328, 7, 97, 2, 2, 328, 329, 7, 120, 2, 2, 329, 330, 7,
	103, 2, 2, 330, 331, 7, 116, 2, 2, 331, 332, 7, 117, 2, 2, 332, 333, 7,
	107, 2, 2, 333, 334, 7, 113, 2, 2, 334, 335, 7, 112, 2, 2, 335, 40, 3,
	2, 2, 2, 336, 337, 7, 99, 2, 2, 337, 338, 7, 112, 2, 2, 338, 339, 7, 102,
	2, 2, 339, 42, 3, 2, 2, 2, 340, 341, 7, 113, 2, 2, 341, 342, 7, 116, 2,
	2, 342, 44, 3, 2, 2, 2, 343, 344, 7, 112, 2, 2, 344, 345, 7, 113, 2, 2,
	345, 346, 7, 118, 2, 2, 346, 46, 3, 2, 2, 2, 347, 348, 7, 62, 2, 2, 348,
	48, 3, 2, 2, 2, 349, 350, 7, 62, 2, 2, 350, 351, 7, 63, 2, 2, 351, 50,
	3, 2, 2, 2, 352, 353, 7, 64, 2, 2, 353, 52, 3, 2, 2, 2, 354, 355, 7, 64,
	2, 2, 355, 356, 7, 63, 2, 2, 356, 54, 3, 2, 2, 2, 357, 358, 7, 63, 2, 2,
	358, 56, 3, 2, 2, 2, 359, 360, 7, 35, 2, 2, 360, 361, 7, 63, 2, 2, 361,
	58, 3, 2, 2, 2, 362, 363, 7, 107, 2, 2, 363, 364, 7, 112, 2, 2, 364, 60,
	3, 2, 2, 2, 365, 366, 7, 101, 2, 2, 366, 367, 7, 113, 2, 2, 367, 368, 7,
	112, 2, 2, 368, 369, 7, 118, 2, 2, 369, 370, 7, 99, 2, 2, 370, 371, 7,
	107, 2, 2, 371, 372, 7, 112, 2, 2, 372, 373, 7, 117, 2, 2, 373, 62, 3,
	2, 2, 2, 374, 375, 7, 107, 2, 2, 375, 376, 7, 101, 2, 2, 376, 377, 7, 113,
	2, 2, 377, 378, 7, 112, 2, 2, 378, 379, 7, 118, 2, 2, 379, 380, 7, 99,
	2, 2, 380, 381, 7, 107, 2, 2, 381, 382, 7, 112, 2, 2, 382, 383, 7, 117,
	2, 2, 383, 64, 3, 2, 2, 2, 384, 385, 7, 117, 2, 2, 385, 386, 7, 118, 2,
	2, 386, 387, 7, 99, 2, 2, 387, 388, 7, 116, 2, 2, 388, 389, 7, 118, 2,
	2, 389, 390, 7, 117, 2, 2, 390, 391, 7, 121, 2, 2, 391, 392, 7, 107, 2,
	2, 392, 393, 7, 118, 2, 2, 393, 394, 7, 106, 2, 2, 394, 66, 3, 2, 2, 2,
	395, 396, 7, 103, 2, 2, 396, 397, 7, 112, 2, 2, 397, 398, 7, 102, 2, 2,
	398, 399, 7, 117, 2, 2, 399, 400, 7, 121, 2, 2, 400, 401, 7, 107, 2, 2,
	401, 402, 7, 118, 2, 2, 402, 403, 7, 106, 2, 2, 403, 68, 3, 2, 2, 2, 404,
	405, 7, 114, 2, 2, 405, 406, 7, 111, 2, 2, 406, 407, 7, 99, 2, 2, 407,
	408, 7, 118, 2, 2, 408, 409, 7, 101, 2, 2, 409, 410, 7, 106, 2, 2, 410,
	70, 3, 2, 2, 2, 411, 412, 7, 103, 2, 2, 412, 413, 7, 122, 2, 2, 413, 414,
	7, 107, 2, 2, 414, 415, 7, 117, 2, 2, 415, 416, 7, 118, 2, 2, 416, 417,
	7, 117, 2, 2, 417, 72, 3, 2, 2, 2, 418, 419, 7, 93, 2, 2, 419, 74, 3, 2,
	2, 2, 420, 421, 7, 95, 2, 2, 421, 76, 3, 2, 2, 2, 422, 423, 7, 42, 2, 2,
	423, 78, 3, 2, 2, 2, 424, 425, 7, 43, 2, 2, 425, 80, 3, 2, 2, 2, 426, 427,
	7, 46, 2, 2, 427, 82, 3, 2, 2, 2, 428, 429, 7, 47, 2, 2, 429, 84, 3, 2,
	2, 2, 430, 438, 7, 60, 2, 2, 431, 433, 7, 34, 2, 2, 432, 431, 3, 2, 2,
	2, 433, 436, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435,
	437, 3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 437, 439, 7, 64, 2, 2, 438, 434,
	3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 86, 3, 2, 2, 2, 440, 443, 5, 89,
	45, 2, 441, 443, 5, 91, 46, 2, 442, 440, 3, 2, 2, 2, 442, 441, 3, 2, 2,
	2, 443, 88, 3, 2, 2, 2, 444, 445, 5, 129, 65, 2, 445, 446, 5, 131, 66,
	2, 446, 447, 5, 127, 64, 2, 447, 448, 5, 129, 65, 2, 448, 461, 3, 2, 2,
	2, 449, 450, 5, 139, 70, 2, 450, 451, 5, 123, 62, 2, 451, 452, 5, 121,
	61, 2, 452, 453, 5, 131, 66, 2, 453, 454, 5, 155, 78, 2, 454, 455, 5, 139,
	70, 2, 455, 461, 3, 2, 2, 2, 456, 457, 5, 137, 69, 2, 457, 458, 5, 143,
	72, 2, 458, 459, 5, 159, 80, 2, 459, 461, 3, 2, 2, 2, 460, 444, 3, 2, 2,
	2, 460, 449, 3, 2, 2, 2, 460, 456, 3, 2, 2, 2, 461, 90, 3, 2, 2, 2, 462,
	463, 5, 123, 62, 2, 463, 464, 5, 139, 70, 2, 464, 465, 5, 123, 62, 2, 465,
	466, 5, 149, 75, 2, 466, 467, 5, 127, 64, 2, 467, 468, 5, 123, 62, 2, 468,
	469, 5, 141, 71, 2, 469, 470, 5, 119, 60, 2, 470, 471, 5, 163, 82, 2, 471,
	534, 3, 2, 2, 2, 472, 473, 5, 115, 58, 2, 473, 474, 5, 137, 69, 2, 474,
	475, 5, 123, 62, 2, 475, 476, 5, 149, 75, 2, 476, 477, 5, 153, 77, 2, 477,
	534, 3, 2, 2, 2, 478, 479, 5, 119, 60, 2, 479, 480, 5, 149, 75, 2, 480,
	481, 5, 131, 66, 2, 481, 482, 5, 153, 77, 2, 482, 483, 5, 131, 66, 2, 483,
	484, 5, 119, 60, 2, 484, 485, 5, 115, 58, 2, 485, 486, 5, 137, 69, 2, 486,
	534, 3, 2, 2, 2, 487, 488, 5, 123, 62, 2, 488, 489, 5, 149, 75, 2, 489,
	490, 5, 149, 75, 2, 490, 491, 5, 143, 72, 2, 491, 492, 5, 149, 75, 2, 492,
	534, 3, 2, 2, 2, 493, 494, 5, 159, 80, 2, 494, 495, 5, 115, 58, 2, 495,
	496, 5, 149, 75, 2, 496, 497, 5, 141, 71, 2, 497, 498, 5, 131, 66, 2, 498,
	499, 5, 141, 71, 2, 499, 500, 5, 127, 64, 2, 500, 534, 3, 2, 2, 2, 501,
	502, 5, 141, 71, 2, 502, 503, 5, 143, 72, 2, 503, 504, 5, 153, 77, 2, 504,
	505, 5, 131, 66, 2, 505, 506, 5, 119, 60, 2, 506, 507, 5, 123, 62, 2, 507,
	534, 3, 2, 2, 2, 508, 509, 5, 131, 66, 2, 509, 510, 5, 141, 71, 2, 510,
	511, 5, 125, 63, 2, 511, 512, 5, 143, 72, 2, 512, 534, 3, 2, 2, 2, 513,
	514, 5, 131, 66, 2, 514, 515, 5, 141, 71, 2, 515, 516, 5, 125, 63, 2, 516,
	517, 5, 143, 72, 2, 517, 518, 5, 149, 75, 2, 518, 519, 5, 139, 70, 2, 519,
	520, 5, 115, 58, 2, 520, 521, 5, 153, 77, 2, 521, 522, 5, 131, 66, 2, 522,
	523, 5, 143, 72, 2, 523, 524, 5, 141, 71, 2, 524, 525, 5, 115, 58, 2, 525,
	526, 5, 137, 69, 2, 526, 534, 3, 2, 2, 2, 527, 528, 5, 121, 61, 2, 528,
	529, 5, 123, 62, 2, 529, 530, 5, 117, 59, 2, 530, 531, 5, 155, 78, 2, 531,
	532, 5, 127, 64, 2, 532, 534, 3, 2, 2, 2, 533, 462, 3, 2, 2, 2, 533, 472,
	3, 2, 2, 2, 533, 478, 3, 2, 2, 2, 533, 487, 3, 2, 2, 2, 533, 493, 3, 2,
	2, 2, 533, 501, 3, 2, 2, 2, 533, 508, 3, 2, 2, 2, 533, 513, 3, 2, 2, 2,
	533, 527, 3, 2, 2, 2, 534, 92, 3, 2, 2, 2, 535, 557, 9, 2, 2, 2, 536, 556,
	9, 3, 2, 2, 537, 539, 7, 60, 2, 2, 538, 537, 3, 2, 2, 2, 538, 539, 3, 2,
	2, 2, 539, 540, 3, 2, 2, 2, 540, 543, 7, 93, 2, 2, 541, 544, 5, 95, 48,
	2, 542, 544, 5, 97, 49, 2, 543, 541, 3, 2, 2, 2, 543, 542, 3, 2, 2, 2,
	544, 549, 3, 2, 2, 2, 545, 546, 7, 60, 2, 2, 546, 548, 5, 97, 49, 2, 547,
	545, 3, 2, 2, 2, 548, 551, 3, 2, 2, 2, 549, 547, 3, 2, 2, 2, 549, 550,
	3, 2, 2, 2, 550, 552, 3, 2, 2, 2, 551, 549, 3, 2, 2, 2, 552, 553, 7, 95,
	2, 2, 553, 556, 3, 2, 2, 2, 554, 556, 7, 44, 2, 2, 555, 536, 3, 2, 2, 2,
	555, 538, 3, 2, 2, 2, 555, 554, 3, 2, 2, 2, 556, 559, 3, 2, 2, 2, 557,
	555, 3, 2, 2, 2, 557, 558, 3, 2, 2, 2, 558, 94, 3, 2, 2, 2, 559, 557, 3,
	2, 2, 2, 560, 562, 4, 50, 59, 2, 561, 560, 3, 2, 2, 2, 562, 563, 3, 2,
	2, 2, 563, 561, 3, 2, 2, 2, 563, 564, 3, 2, 2, 2, 564, 571, 3, 2, 2, 2,
	565, 567, 7, 48, 2, 2, 566, 568, 4, 50, 59, 2, 567, 566, 3, 2, 2, 2, 568,
	569, 3, 2, 2, 2, 569, 567, 3, 2, 2, 2, 569, 570, 3, 2, 2, 2, 570, 572,
	3, 2, 2, 2, 571, 565, 3, 2, 2, 2, 571, 572, 3, 2, 2, 2, 572, 96, 3, 2,
	2, 2, 573, 577, 9, 4, 2, 2, 574, 576, 9, 5, 2, 2, 575, 574, 3, 2, 2, 2,
	576, 579, 3, 2, 2, 2, 577, 575, 3, 2, 2, 2, 577, 578, 3, 2, 2, 2, 578,
	98, 3, 2, 2, 2, 579, 577, 3, 2, 2, 2, 580, 583, 7, 36, 2, 2, 581, 584,
	5, 99, 50, 2, 582, 584, 5, 103, 52, 2, 583, 581, 3, 2, 2, 2, 583, 582,
	3, 2, 2, 2, 584, 585, 3, 2, 2, 2, 585, 586, 7, 36, 2, 2, 586, 615, 3, 2,
	2, 2, 587, 590, 7, 41, 2, 2, 588, 591, 5, 99, 50, 2, 589, 591, 5, 103,
	52, 2, 590, 588, 3, 2, 2, 2, 590, 589, 3, 2, 2, 2, 591, 592, 3, 2, 2, 2,
	592, 593, 7, 41, 2, 2, 593, 615, 3, 2, 2, 2, 594, 595, 7, 94, 2, 2, 595,
	596, 7, 36, 2, 2, 596, 599, 3, 2, 2, 2, 597, 600, 5, 99, 50, 2, 598, 600,
	5, 103, 52, 2, 599, 597, 3, 2, 2, 2, 599, 598, 3, 2, 2, 2, 600, 601, 3,
	2, 2, 2, 601, 602, 7, 94, 2, 2, 602, 603, 7, 36, 2, 2, 603, 615, 3, 2,
	2, 2, 604, 605, 7, 41, 2, 2, 605, 606, 7, 41, 2, 2, 606, 609, 3, 2, 2,
	2, 607, 610, 5, 99, 50, 2, 608, 610, 5, 103, 52, 2, 609, 607, 3, 2, 2,
	2, 609, 608, 3, 2, 2, 2, 610, 611, 3, 2, 2, 2, 611, 612, 7, 41, 2, 2, 612,
	613, 7, 41, 2, 2, 613, 615, 3, 2, 2, 2, 614, 580, 3, 2, 2, 2, 614, 587,
	3, 2, 2, 2, 614, 594, 3, 2, 2, 2, 614, 604, 3, 2, 2, 2, 615, 100, 3, 2,
	2, 2, 616, 617, 5, 93, 47, 2, 617, 618, 7, 60, 2, 2, 618, 619, 5, 93, 47,
	2, 619, 102, 3, 2, 2, 2, 620, 622, 10, 6, 2, 2, 621, 620, 3, 2, 2, 2, 622,
	625, 3, 2, 2, 2, 623, 624, 3, 2, 2, 2, 623, 621, 3, 2, 2, 2, 624, 104,
	3, 2, 2, 2, 625, 623, 3, 2, 2, 2, 626, 627, 7, 94, 2, 2, 627, 631, 7, 36,
	2, 2, 628, 629, 7, 41, 2, 2, 629, 631, 7, 41, 2, 2, 630, 626, 3, 2, 2,
	2, 630, 628, 3, 2, 2, 2, 631, 106, 3, 2, 2, 2, 632, 634, 9, 7, 2, 2, 633,
	632, 3, 2, 2, 2, 634, 635, 3, 2, 2, 2, 635, 633, 3, 2, 2, 2, 635, 636,
	3, 2, 2, 2, 636, 637, 3, 2, 2, 2, 637, 638, 8, 54, 2, 2, 638, 108, 3, 2,
	2, 2, 639, 641, 7, 15, 2, 2, 640, 639, 3, 2, 2, 2, 640, 641, 3, 2, 2, 2,
	641, 642, 3, 2, 2, 2, 642, 643, 7, 12, 2, 2, 643, 644, 3, 2, 2, 2, 644,
	645, 8, 55, 2, 2, 645, 110, 3, 2, 2, 2, 646, 650, 7, 37, 2, 2, 647, 649,
	10, 6, 2, 2, 648, 647, 3, 2, 2, 2, 649, 652, 3, 2, 2, 2, 650, 648, 3, 2,
	2, 2, 650, 651, 3, 2, 2, 2, 651, 653, 3, 2, 2, 2, 652, 650, 3, 2, 2, 2,
	653, 654, 8, 56, 2, 2, 654, 112, 3, 2, 2, 2, 655, 656, 11, 2, 2, 2, 656,
	114, 3, 2, 2, 2, 657, 658, 9, 8, 2, 2, 658, 116, 3, 2, 2, 2, 659, 660,
	9, 9, 2, 2, 660, 118, 3, 2, 2, 2, 661, 662, 9, 10, 2, 2, 662, 120, 3, 2,
	2, 2, 663, 664, 9, 11, 2, 2, 664, 122, 3, 2, 2, 2, 665, 666, 9, 12, 2,
	2, 666, 124, 3, 2, 2, 2, 667, 668, 9, 13, 2, 2, 668, 126, 3, 2, 2, 2, 669,
	670, 9, 14, 2, 2, 670, 128, 3, 2, 2, 2, 671, 672, 9, 15, 2, 2, 672, 130,
	3, 2, 2, 2, 673, 674, 9, 16, 2, 2, 674, 132, 3, 2, 2, 2, 675, 676, 9, 17,
	2, 2, 676, 134, 3, 2, 2, 2, 677, 678, 9, 18, 2, 2, 678, 136, 3, 2, 2, 2,
	679, 680, 9, 19, 2, 2, 680, 138, 3, 2, 2, 2, 681, 682, 9, 20, 2, 2, 682,
	140, 3, 2, 2, 2, 683, 684, 9, 21, 2, 2, 684, 142, 3, 2, 2, 2, 685, 686,
	9, 22, 2, 2, 686, 144, 3, 2, 2, 2, 687, 688, 9, 23, 2, 2, 688, 146, 3,
	2, 2, 2, 689, 690, 9, 24, 2, 2, 690, 148, 3, 2, 2, 2, 691, 692, 9, 25,
	2, 2, 692, 150, 3, 2, 2, 2, 693, 694, 9, 26, 2, 2, 694, 152, 3, 2, 2, 2,
	695, 696, 9, 27, 2, 2, 696, 154, 3, 2, 2, 2, 697, 698, 9, 28, 2, 2, 698,
	156, 3, 2, 2, 2, 699, 700, 9, 29, 2, 2, 700, 158, 3, 2, 2, 2, 701, 702,
	9, 30, 2, 2, 702, 160, 3, 2, 2, 2, 703, 704, 9, 31, 2, 2, 704, 162, 3,
	2, 2, 2, 705, 706, 9, 32, 2, 2, 706, 164, 3, 2, 2, 2, 707, 708, 9, 33,
	2, 2, 708, 166, 3, 2, 2, 2, 27, 2, 434, 438, 442, 460, 533, 538, 543, 549,
	555, 557, 563, 569, 571, 577, 583, 590, 599, 609, 614, 623, 630, 635, 640,
	650, 3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'rule'", "'filter'", "'drop'", "'macro'", "'list'", "'name'", "'items'",
	"'condition'", "'desc'", "'actions'", "'output'", "'priority'", "'tags'",
	"'prefilter'", "'enabled'", "'warn_evttypes'", "'skip-if-unknown-filter'",
	"'append'", "'required_engine_version'", "'and'", "'or'", "'not'", "'<'",
	"'<='", "'>'", "'>='", "'='", "'!='", "'in'", "'contains'", "'icontains'",
	"'startswith'", "'endswith'", "'pmatch'", "'exists'", "'['", "']'", "'('",
	"')'", "','", "'-'",
}

var lexerSymbolicNames = []string{
	"", "RULE", "FILTER", "DROP", "MACRO", "LIST", "NAME", "ITEMS", "COND",
	"DESC", "ACTIONS", "OUTPUT", "PRIORITY", "TAGS", "PREFILTER", "ENABLED",
	"WARNEVTTYPE", "SKIPUNKNOWN", "FAPPEND", "REQ", "AND", "OR", "NOT", "LT",
	"LE", "GT", "GE", "EQ", "NEQ", "IN", "CONTAINS", "ICONTAINS", "STARTSWITH",
	"ENDSWITH", "PMATCH", "EXISTS", "LBRACK", "RBRACK", "LPAREN", "RPAREN",
	"LISTSEP", "DECL", "DEF", "SEVERITY", "SFSEVERITY", "FSEVERITY", "ID",
	"NUMBER", "PATH", "STRING", "TAG", "WS", "NL", "COMMENT", "ANY",
}

var lexerRuleNames = []string{
	"RULE", "FILTER", "DROP", "MACRO", "LIST", "NAME", "ITEMS", "COND", "DESC",
	"ACTIONS", "OUTPUT", "PRIORITY", "TAGS", "PREFILTER", "ENABLED", "WARNEVTTYPE",
	"SKIPUNKNOWN", "FAPPEND", "REQ", "AND", "OR", "NOT", "LT", "LE", "GT",
	"GE", "EQ", "NEQ", "IN", "CONTAINS", "ICONTAINS", "STARTSWITH", "ENDSWITH",
	"PMATCH", "EXISTS", "LBRACK", "RBRACK", "LPAREN", "RPAREN", "LISTSEP",
	"DECL", "DEF", "SEVERITY", "SFSEVERITY", "FSEVERITY", "ID", "NUMBER", "PATH",
	"STRING", "TAG", "STRLIT", "ESC", "WS", "NL", "COMMENT", "ANY", "A", "B",
	"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q",
	"R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

type SfplLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewSfplLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *SfplLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSfplLexer(input antlr.CharStream) *SfplLexer {
	l := new(SfplLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Sfpl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SfplLexer tokens.
const (
	SfplLexerRULE        = 1
	SfplLexerFILTER      = 2
	SfplLexerDROP        = 3
	SfplLexerMACRO       = 4
	SfplLexerLIST        = 5
	SfplLexerNAME        = 6
	SfplLexerITEMS       = 7
	SfplLexerCOND        = 8
	SfplLexerDESC        = 9
	SfplLexerACTIONS     = 10
	SfplLexerOUTPUT      = 11
	SfplLexerPRIORITY    = 12
	SfplLexerTAGS        = 13
	SfplLexerPREFILTER   = 14
	SfplLexerENABLED     = 15
	SfplLexerWARNEVTTYPE = 16
	SfplLexerSKIPUNKNOWN = 17
	SfplLexerFAPPEND     = 18
	SfplLexerREQ         = 19
	SfplLexerAND         = 20
	SfplLexerOR          = 21
	SfplLexerNOT         = 22
	SfplLexerLT          = 23
	SfplLexerLE          = 24
	SfplLexerGT          = 25
	SfplLexerGE          = 26
	SfplLexerEQ          = 27
	SfplLexerNEQ         = 28
	SfplLexerIN          = 29
	SfplLexerCONTAINS    = 30
	SfplLexerICONTAINS   = 31
	SfplLexerSTARTSWITH  = 32
	SfplLexerENDSWITH    = 33
	SfplLexerPMATCH      = 34
	SfplLexerEXISTS      = 35
	SfplLexerLBRACK      = 36
	SfplLexerRBRACK      = 37
	SfplLexerLPAREN      = 38
	SfplLexerRPAREN      = 39
	SfplLexerLISTSEP     = 40
	SfplLexerDECL        = 41
	SfplLexerDEF         = 42
	SfplLexerSEVERITY    = 43
	SfplLexerSFSEVERITY  = 44
	SfplLexerFSEVERITY   = 45
	SfplLexerID          = 46
	SfplLexerNUMBER      = 47
	SfplLexerPATH        = 48
	SfplLexerSTRING      = 49
	SfplLexerTAG         = 50
	SfplLexerWS          = 51
	SfplLexerNL          = 52
	SfplLexerCOMMENT     = 53
	SfplLexerANY         = 54
)