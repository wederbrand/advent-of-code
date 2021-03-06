package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day08Test {

	public static final String INPUT = "t inc 245 if xq != 0\n" +
		"hi inc 119 if gf >= -5\n" +
		"w dec 529 if s == 0\n" +
		"p inc 19 if fi > -2\n" +
		"kgp dec 133 if kbm > -6\n" +
		"pl inc -407 if pvo != 0\n" +
		"gf dec 348 if gf <= 7\n" +
		"e inc -277 if pdg == 0\n" +
		"pdg dec 480 if p == 19\n" +
		"pl dec 932 if kgp <= -127\n" +
		"gf inc 711 if xq <= 7\n" +
		"e dec -359 if esj < -2\n" +
		"l dec -543 if jcf <= 9\n" +
		"bh inc 958 if t > -10\n" +
		"h dec 87 if hsv == -6\n" +
		"e inc 290 if esj >= -4\n" +
		"e dec -378 if fi >= -2\n" +
		"fi inc 722 if esj == 0\n" +
		"w dec -350 if bh != 948\n" +
		"e dec 974 if fi > 716\n" +
		"hsv inc 852 if xq >= -6\n" +
		"gf dec -548 if bh >= 958\n" +
		"fi dec -234 if fi > 722\n" +
		"fi dec 983 if gf != 919\n" +
		"esj dec -155 if kgp <= -125\n" +
		"p inc -411 if gf <= 913\n" +
		"kgp dec 304 if kgp == -133\n" +
		"hi inc -488 if e == -583\n" +
		"pvo inc -524 if f > -1\n" +
		"gf inc 654 if s == 0\n" +
		"fw dec -872 if pdg > -478\n" +
		"kgp inc 627 if kbm != 0\n" +
		"u dec 294 if fw >= 8\n" +
		"esj dec -979 if esj <= 160\n" +
		"pl dec -194 if fw <= 4\n" +
		"ls dec 74 if fw <= 5\n" +
		"pvo inc 413 if f <= 6\n" +
		"ls inc 917 if bh != 951\n" +
		"jlg inc 866 if h == 0\n" +
		"hi inc -598 if gf <= 1567\n" +
		"p dec 138 if l > 549\n" +
		"hsv inc -530 if pl != -738\n" +
		"bh inc 160 if f <= 0\n" +
		"s dec -619 if p >= -396\n" +
		"p dec 243 if hi != -971\n" +
		"gf dec 614 if pdg < -470\n" +
		"s dec 804 if f == -2\n" +
		"xq inc 10 if esj != 1144\n" +
		"p inc 860 if gf > 943\n" +
		"pl dec 564 if fi >= -267\n" +
		"ls inc 736 if kbm >= -5\n" +
		"hsv dec 774 if gf != 946\n" +
		"f inc 301 if p == 232\n" +
		"w dec -553 if hi != -965\n" +
		"kbm inc -875 if pvo != -115\n" +
		"pvo dec 376 if esj > 1134\n" +
		"kbm dec -293 if u != 0\n" +
		"pdg dec -274 if pvo <= -111\n" +
		"xq dec -385 if f > 0\n" +
		"gf dec 202 if pl < -1293\n" +
		"s inc 782 if pvo <= -104\n" +
		"p dec -652 if w == 374\n" +
		"w inc 61 if bh < 1122\n" +
		"ls inc -978 if ls != 1570\n" +
		"p inc 504 if esj < 1137\n" +
		"xq inc 111 if fw == -10\n" +
		"h dec -183 if h > -3\n" +
		"fi dec 990 if fw < 1\n" +
		"kbm inc -534 if s == 1401\n" +
		"esj inc -619 if hsv > 70\n" +
		"p inc 704 if kgp != -437\n" +
		"hi inc 637 if esj > 511\n" +
		"hsv dec 653 if h > 179\n" +
		"s dec -866 if kgp < -431\n" +
		"pdg dec -769 if bh > 1111\n" +
		"s dec -796 if p >= 1390\n" +
		"s dec 137 if jcf == 0\n" +
		"w inc -818 if p <= 1381\n" +
		"l dec -601 if h >= 177\n" +
		"pdg dec 50 if f == 0\n" +
		"ls inc -607 if f <= 4\n" +
		"gf inc 894 if l >= 1136\n" +
		"kbm dec -775 if u < 7\n" +
		"pdg inc 151 if pvo != -109\n" +
		"hsv inc 366 if jlg != 861\n" +
		"t inc 960 if f > -7\n" +
		"pdg inc -397 if kbm < -632\n" +
		"xq inc -71 if fi > -1258\n" +
		"ls inc 211 if hsv > -208\n" +
		"pl dec 442 if fi == -1251\n" +
		"kbm inc -920 if pdg > 263\n" +
		"kgp dec 627 if t >= 956\n" +
		"kgp inc 447 if u < -6\n" +
		"s inc 815 if f != 3\n" +
		"w dec -201 if jlg < 876\n" +
		"s inc 630 if t > 967\n" +
		"h dec 4 if l != 1144\n" +
		"gf inc 346 if l < 1146\n" +
		"jcf dec -72 if pdg == 267\n" +
		"pl inc 804 if p <= 1384\n" +
		"pl dec -48 if e == -583\n" +
		"p inc 622 if s == 2945\n" +
		"fi inc -513 if ls != -14\n" +
		"pl dec -802 if fw > -2\n" +
		"bh dec -438 if f <= 6\n" +
		"fi dec -968 if gf > 1987\n" +
		"pdg inc -973 if xq <= -57\n" +
		"kbm inc -482 if pvo != -111\n" +
		"s dec -228 if ls >= -8\n" +
		"e inc 850 if hsv == -209\n" +
		"gf inc -540 if u != 9\n" +
		"hi inc 176 if kbm > -1555\n" +
		"p dec -868 if e > 265\n" +
		"s dec 141 if kgp > -1063\n" +
		"e inc 438 if w == -178\n" +
		"hsv inc 982 if fi <= -796\n" +
		"hsv inc 67 if hsv == 773\n" +
		"hi inc -633 if e == 267\n" +
		"hi dec -361 if kgp <= -1070\n" +
		"ls dec -474 if pdg < -704\n" +
		"kgp dec 650 if fw <= 9\n" +
		"pdg dec 355 if h == 183\n" +
		"e inc -298 if kgp == -1714\n" +
		"hsv dec -112 if f <= 0\n" +
		"p dec -491 if jcf == 72\n" +
		"pvo inc 749 if pvo >= -116\n" +
		"bh inc 165 if s <= 3178\n" +
		"p dec -598 if jlg < 872\n" +
		"jlg inc -72 if f >= -5\n" +
		"t inc 416 if l < 1146\n" +
		"w dec 343 if ls < 468\n" +
		"ls inc 405 if h >= 179\n" +
		"pl inc -984 if u >= 5\n" +
		"pvo inc 440 if jlg >= 785\n" +
		"fi inc -729 if jlg != 784\n" +
		"kgp inc 351 if s < 3182\n" +
		"bh inc 181 if h >= 189\n" +
		"pvo dec -445 if pl <= -90\n" +
		"pdg dec 658 if t < 1385\n" +
		"bh inc 64 if xq != -67\n" +
		"hi inc -88 if bh >= 1778\n" +
		"fi dec 413 if jlg < 795\n" +
		"ls dec -602 if fw > -9\n" +
		"h inc 228 if jlg <= 791\n" +
		"e inc -190 if e == -31\n" +
		"jcf inc -600 if xq == -61\n" +
		"ls dec -549 if pdg != -1723\n" +
		"u inc 348 if u != -7\n" +
		"hi dec 967 if h < 193\n" +
		"s inc -323 if l >= 1149\n" +
		"xq inc -252 if hi >= -1848\n" +
		"f inc 80 if t <= 1383\n" +
		"bh inc -890 if jcf > -520\n" +
		"pdg inc -267 if kgp < -1369\n" +
		"jcf dec 569 if gf <= 1453\n" +
		"h inc -632 if hsv == 952\n" +
		"jcf dec 917 if s == 3173\n" +
		"fw inc -403 if e == -221\n" +
		"kbm inc -763 if t >= 1376\n" +
		"fw dec 821 if e <= -219\n" +
		"esj inc 630 if xq != -313\n" +
		"e inc -982 if t <= 1382\n" +
		"kgp dec -802 if t == 1372\n" +
		"t inc 511 if fw < -1215\n" +
		"e dec 750 if w >= -178\n" +
		"bh dec 549 if hsv == 952\n" +
		"h dec 881 if kbm == -2317\n" +
		"fw dec -883 if jlg < 798\n" +
		"p inc -86 if kbm <= -2317\n" +
		"xq dec -277 if ls <= 2017\n" +
		"pdg inc 518 if e < -1209\n" +
		"l inc 933 if esj == 515\n" +
		"xq dec 962 if w <= -173\n" +
		"jlg dec -79 if w >= -178\n" +
		"pvo dec 200 if jlg == 794\n" +
		"f dec -31 if pvo != 1321\n" +
		"pl inc -585 if pvo == 1323\n" +
		"u inc -750 if pdg == -1719\n" +
		"jlg dec 445 if l <= 2082\n" +
		"l inc -373 if kgp <= -1362\n" +
		"w dec 812 if pl == -682\n" +
		"fi dec 760 if f <= 115\n" +
		"t inc 457 if w != -181\n" +
		"u inc -562 if ls >= 2017\n" +
		"hsv dec -666 if t > 2338\n" +
		"pvo dec 771 if gf != 1446\n" +
		"gf inc 665 if pl <= -674\n" +
		"xq inc -465 if t < 2339\n" +
		"e dec 519 if gf > 2112\n" +
		"esj inc -518 if f <= 114\n" +
		"gf inc -794 if s != 3183\n" +
		"p inc 254 if hsv == 1618\n" +
		"t inc 242 if fw < -331\n" +
		"fw inc -52 if kbm > -2327\n" +
		"gf inc -483 if l >= 1703\n" +
		"kgp dec -569 if hi > -1851\n" +
		"t inc 409 if kgp < -801\n" +
		"kgp dec 528 if u < -958\n" +
		"gf dec -961 if jlg < 348\n" +
		"jcf inc 713 if esj > -12\n" +
		"w inc -429 if hi != -1835\n" +
		"h dec -168 if f > 104\n" +
		"e dec -942 if jlg > 344\n" +
		"h dec -429 if esj != 6\n" +
		"e dec -481 if jcf <= -1295\n" +
		"e inc -676 if e > -307\n" +
		"p inc 855 if pdg == -1719\n" +
		"pvo inc 993 if fw > -398\n" +
		"kgp dec 196 if kgp >= -1325\n" +
		"s dec -576 if pl >= -679\n" +
		"xq inc 510 if p > 4979\n" +
		"jlg dec -898 if hsv != 1610\n" +
		"pl inc -153 if xq > -774\n" +
		"l dec -319 if ls == 2024\n" +
		"l inc -945 if s <= 3757\n" +
		"bh dec -937 if hi <= -1846\n" +
		"u inc -631 if xq <= -758\n" +
		"l dec 966 if h <= -734\n" +
		"gf dec -2 if w == -611\n" +
		"t dec 803 if w <= -619\n" +
		"hsv dec -167 if jcf < -1296\n" +
		"pdg inc 136 if pl <= -835\n" +
		"hsv inc 29 if t > 2582\n" +
		"fi dec 808 if jcf >= -1306\n" +
		"pl inc -910 if pvo > 1546\n" +
		"pdg inc 487 if pl <= -831\n" +
		"hsv inc 265 if fi < -3509\n" +
		"jlg dec 239 if pdg >= -1727\n" +
		"esj dec -811 if hsv < 1822\n" +
		"f inc -467 if xq == -765\n" +
		"fw dec -644 if u >= -1592\n" +
		"fw dec -767 if kbm > -2324\n" +
		"jlg dec 890 if xq <= -766\n" +
		"xq dec -855 if jlg < 1009\n" +
		"u inc 810 if hsv == 1819\n" +
		"s inc 379 if jcf < -1306\n" +
		"ls dec -475 if l <= 1078\n" +
		"kbm dec -339 if pdg <= -1717\n" +
		"hi dec -911 if l < 1083\n" +
		"hi inc 85 if xq != 85\n" +
		"hsv inc 439 if pl > -834\n" +
		"h dec -956 if bh < 1237\n" +
		"pdg inc -508 if kgp > -1522\n" +
		"t dec -63 if kbm >= -1974\n" +
		"kbm inc -619 if f < -354\n" +
		"s inc -110 if pvo < 1552\n" +
		"e dec -100 if jcf <= -1301\n" +
		"fi dec 369 if jcf > -1304\n" +
		"w dec -191 if kgp <= -1515\n" +
		"hsv inc 742 if w == -420\n" +
		"gf dec -445 if e >= -877\n" +
		"ls dec -582 if kgp >= -1519\n" +
		"pl inc 561 if jcf >= -1310\n" +
		"t dec -910 if fi <= -3881\n" +
		"ls inc -772 if e < -874\n" +
		"e dec -322 if bh <= 1237\n" +
		"pvo dec 189 if l == 1078\n" +
		"kgp inc -727 if l != 1071\n" +
		"ls inc -461 if esj == 808\n" +
		"p inc -781 if ls > 1843\n" +
		"fw dec -284 if hi == -853\n" +
		"pvo dec -946 if w <= -414\n" +
		"s dec 409 if pl < -259\n" +
		"pdg dec -490 if esj != 808\n" +
		"h inc 400 if kgp > -2251\n" +
		"ls inc 240 if fi > -3882\n" +
		"jcf inc -944 if u > -1602\n" +
		"u dec 315 if f >= -361\n" +
		"hi dec -353 if h >= 621\n" +
		"jlg inc 846 if f < -354\n" +
		"hi dec -49 if jcf > -2246\n" +
		"f dec 274 if p > 4208\n" +
		"h dec -772 if esj != 815\n" +
		"fi inc 656 if hi < -440\n" +
		"fi dec -49 if t > 2585\n" +
		"w dec -219 if fw != 368\n" +
		"fi inc 142 if fi == -3162\n" +
		"kgp dec 574 if pl > -262\n" +
		"u dec -855 if xq > 85\n" +
		"f dec -137 if f >= -347\n" +
		"bh inc 355 if e == -553\n" +
		"hsv inc 310 if pdg > -2230\n" +
		"hi dec -219 if hi != -445\n" +
		"u dec -915 if pdg > -2236\n" +
		"hi inc -4 if fw > 369\n" +
		"pdg dec -364 if fw >= 365\n" +
		"bh inc 768 if e <= -551\n" +
		"f inc 409 if s >= 3224\n" +
		"esj inc 968 if gf < 1275\n" +
		"jlg dec -21 if kgp >= -2245\n" +
		"f inc -206 if fi < -3162\n" +
		"u dec -829 if w == -201\n" +
		"s dec 878 if hi != -230\n" +
		"hsv inc 467 if fw >= 365\n" +
		"u dec 725 if fi <= -3166\n" +
		"gf dec 391 if kbm < -2589\n" +
		"bh dec 535 if fi == -3166\n" +
		"gf dec 540 if e <= -553\n" +
		"fw inc 645 if kbm <= -2592\n" +
		"pl inc 606 if f != -149\n" +
		"xq dec 421 if jlg < 1878\n" +
		"p dec 244 if esj == 808\n" +
		"gf inc 160 if jlg > 1869\n" +
		"pvo dec -243 if pvo >= 2295\n" +
		"jcf inc 100 if hsv >= 3776\n" +
		"s dec 43 if ls < 2093\n" +
		"fi dec -487 if hi == -229\n" +
		"pl inc -875 if pdg == -1863\n" +
		"fi dec 56 if bh == 2359\n" +
		"s inc -659 if hi > -237\n" +
		"gf dec 371 if gf >= 523\n" +
		"gf dec -431 if bh == 2359\n" +
		"jlg dec 650 if w <= -210\n" +
		"h inc 7 if t != 2576\n" +
		"pdg dec -147 if hi <= -230\n" +
		"fw inc 118 if pl > -531\n" +
		"fw dec -291 if f >= -160\n" +
		"fw inc -548 if s == 1650\n" +
		"bh inc 519 if pdg >= -1872\n" +
		"xq dec -192 if hi < -221\n" +
		"f inc 988 if kbm > -2604\n" +
		"pl dec -188 if esj >= 803\n" +
		"w dec -643 if u != -36\n" +
		"fw dec 215 if hi < -228\n" +
		"jcf inc -158 if fi != -2739\n" +
		"f dec 574 if s == 1650\n" +
		"f inc -960 if pvo != 2553\n" +
		"gf dec -877 if u == -35\n" +
		"h dec -456 if l == 1078\n" +
		"l inc 382 if l < 1085\n" +
		"fw dec -301 if jcf < -2241\n" +
		"u inc 273 if w < -195\n" +
		"jcf dec 801 if fw >= 848\n" +
		"h dec -679 if jcf != -3049\n" +
		"e dec 925 if fi != -2734\n" +
		"e inc -167 if t != 2583\n" +
		"xq dec 194 if jcf == -3046\n" +
		"w dec -367 if ls <= 2095\n" +
		"jlg inc 425 if pl == -347\n" +
		"p dec 594 if pdg > -1862\n" +
		"h dec 693 if kbm <= -2595\n" +
		"pl inc -951 if w > 164\n" +
		"pvo inc -781 if gf <= 953\n" +
		"w dec 903 if hsv <= 3777\n" +
		"e dec 679 if kbm != -2592\n" +
		"e inc 805 if hi != -220\n" +
		"l inc 93 if gf < 943\n" +
		"bh dec 26 if w >= -738\n" +
		"w dec 66 if jlg < 1883\n" +
		"esj dec 174 if gf > 945\n" +
		"e inc 954 if bh != 2853\n" +
		"ls dec 197 if hsv > 3778\n" +
		"kgp inc -129 if hi > -232\n" +
		"kbm inc 496 if f == -699\n" +
		"ls inc 466 if s > 1648\n" +
		"pl inc -942 if w == -803\n" +
		"s inc 211 if pl == -2241\n" +
		"xq dec 992 if e > -566\n" +
		"pl inc 644 if t >= 2580\n" +
		"hsv dec -309 if u <= 236\n" +
		"pdg dec -381 if bh > 2843\n" +
		"hsv dec -865 if w == -803\n" +
		"fw inc -456 if kgp == -2374\n" +
		"f inc 454 if hsv != 4635\n" +
		"jcf inc 902 if xq < -1321\n" +
		"t inc -496 if jlg <= 1880\n" +
		"kgp dec -228 if kgp >= -2375\n" +
		"xq dec -161 if e != -566\n" +
		"pdg dec -738 if l == 1460\n" +
		"f dec -965 if l < 1465\n" +
		"jlg dec 518 if f == 720\n" +
		"esj inc 358 if bh <= 2852\n" +
		"ls inc 916 if h > 1840\n" +
		"t dec -94 if bh < 2858\n" +
		"h inc -881 if kbm == -2101\n" +
		"e dec -741 if gf >= 950\n" +
		"ls inc 868 if fw >= 402\n" +
		"jcf inc 330 if p < 3964\n" +
		"l dec 428 if kbm >= -2103\n" +
		"p inc 699 if kgp < -2136\n" +
		"l dec 602 if w == -803\n" +
		"fi inc 843 if jcf != -1820\n" +
		"gf dec -860 if esj < 1175\n" +
		"l inc -184 if hsv > 4628\n" +
		"l dec -305 if fi <= -1906\n" +
		"pl dec -338 if f != 720\n" +
		"hi inc -202 if f == 720\n" +
		"s inc 342 if w == -811\n" +
		"kgp dec -684 if p != 4665\n" +
		"t inc 815 if pl >= -1598\n" +
		"fw dec 543 if hsv <= 4628\n" +
		"t dec -414 if xq >= -1167\n" +
		"p inc 275 if e == -565\n" +
		"pvo dec -896 if fw >= 400\n" +
		"pl inc 505 if pvo > 1771\n" +
		"xq dec 129 if esj <= 1170\n" +
		"bh dec -75 if l == 246\n" +
		"pl dec -397 if gf != 1812\n" +
		"p dec 681 if esj <= 1167\n" +
		"pl inc -224 if h == 963\n" +
		"gf dec 964 if pl <= -1426\n" +
		"f inc -298 if kgp >= -1471\n" +
		"bh inc 1000 if l <= 248\n" +
		"pvo dec -455 if p != 4257\n" +
		"s dec 183 if w != -797\n" +
		"f dec -232 if l < 247\n" +
		"hi inc 766 if hi == -431\n" +
		"fw dec 973 if ls < 3474\n" +
		"pvo dec 685 if kgp == -1462\n" +
		"p inc 781 if h < 965\n" +
		"kbm dec -179 if hsv < 4632\n" +
		"hsv dec -827 if w == -803\n" +
		"jlg dec 83 if jlg <= 1360\n" +
		"kbm dec 283 if fi == -1896\n" +
		"w dec -933 if xq <= -1288\n" +
		"jlg inc -566 if kbm <= -2382\n" +
		"jlg dec -652 if e > -575\n" +
		"hi dec 965 if ls > 3467\n" +
		"hsv dec 599 if f >= 663\n" +
		"xq inc 395 if hi > -631\n" +
		"hi inc 29 if kgp <= -1464\n" +
		"l inc -656 if ls <= 3471\n" +
		"l inc 475 if xq == -898\n" +
		"fw dec -997 if jlg >= 1361\n" +
		"xq dec 218 if pvo < 1532\n" +
		"gf inc -920 if pl == -1424\n" +
		"e inc -707 if f <= 657\n" +
		"p dec 438 if pl < -1422\n" +
		"fw dec -192 if h <= 972\n" +
		"p dec 153 if fi != -1890\n" +
		"hsv inc 614 if pvo != 1534\n" +
		"bh dec 508 if gf >= 882\n" +
		"pvo dec 670 if ls <= 3473\n" +
		"pvo inc 72 if ls <= 3463\n" +
		"kbm dec -783 if jlg != 1360\n" +
		"gf inc 309 if s <= 1684\n" +
		"u inc 250 if l <= 73\n" +
		"hi dec -128 if h < 971\n" +
		"jlg dec 588 if pvo > 861\n" +
		"h inc -690 if jlg <= 776\n" +
		"xq dec -56 if esj == 1166\n" +
		"t inc 762 if esj < 1168\n" +
		"fw dec 190 if w <= 138\n" +
		"jcf inc 305 if bh == 3419\n" +
		"t dec -35 if hsv == 5464\n" +
		"gf dec -464 if kbm < -2380\n" +
		"bh inc -315 if jlg < 782\n" +
		"e dec 436 if kbm <= -2379\n" +
		"u inc -135 if l < 66\n" +
		"gf inc -996 if p == 4441\n" +
		"s inc 69 if l == 65\n" +
		"kbm dec 583 if kgp < -1470\n" +
		"w inc -539 if p >= 4435\n" +
		"s dec 630 if fw >= -576\n" +
		"jcf dec -585 if w < -405\n" +
		"ls dec 803 if pvo != 864\n" +
		"u dec -704 if p >= 4448\n" +
		"t dec 567 if s < 1748\n" +
		"bh dec 973 if w >= -413\n" +
		"kgp dec 605 if p <= 4435\n" +
		"fw inc 502 if t != 3651\n" +
		"s dec 804 if jcf > -932\n" +
		"u inc 723 if jlg <= 776\n" +
		"f inc 476 if u > 1074\n" +
		"fi dec -345 if xq >= -851\n" +
		"ls inc 554 if fi <= -1559\n" +
		"ls inc 705 if pl <= -1421\n" +
		"kgp dec -439 if xq >= -843\n" +
		"xq inc -540 if xq > -847\n" +
		"e inc -969 if xq != -1377\n" +
		"pl dec -935 if p <= 4442\n" +
		"jlg inc 208 if jlg < 780\n" +
		"t dec 431 if xq > -1391\n" +
		"s inc -182 if kbm <= -2382\n" +
		"w inc -678 if h < 283\n" +
		"esj inc 662 if hi <= -507\n" +
		"jcf dec 267 if t != 3219\n" +
		"jcf inc -948 if fw == -77\n" +
		"s inc 358 if hi >= -507\n" +
		"f inc 364 if jcf <= -2139\n" +
		"u dec -178 if xq == -1382\n" +
		"fw inc 174 if kbm <= -2378\n" +
		"p inc 607 if t == 3213\n" +
		"jcf inc 448 if pl > -494\n" +
		"pl inc 217 if kgp != -1023\n" +
		"pl inc -395 if kbm < -2378\n" +
		"bh inc -422 if hi <= -507\n" +
		"kgp dec -845 if p != 4437\n" +
		"kbm dec -660 if pdg == -747\n" +
		"jcf dec 653 if e == -2677\n" +
		"h dec 103 if s < 1123\n" +
		"jlg inc -504 if pvo > 868\n" +
		"l dec 824 if fi < -1543\n" +
		"jcf inc -325 if ls <= 4166\n" +
		"esj inc -485 if fi >= -1546\n" +
		"pl dec 719 if l >= -761\n" +
		"kgp inc -257 if gf >= 655\n" +
		"s dec 612 if kgp >= -425\n" +
		"s inc 67 if s > 1111\n" +
		"f dec 967 if pdg > -745\n" +
		"kbm inc 827 if l > -765\n" +
		"h dec 423 if fw > 105\n" +
		"ls inc 431 if h > 166\n" +
		"f dec 592 if esj >= 1171\n" +
		"f dec 532 if jcf == -2338\n" +
		"kbm inc -162 if h > 165\n" +
		"w inc 180 if hi != -495\n" +
		"ls inc 603 if hi >= -496\n" +
		"xq inc -516 if fi != -1551\n" +
		"s inc -30 if hi != -495\n" +
		"jcf inc -342 if pdg < -737\n" +
		"pl dec -278 if esj > 1158\n" +
		"pdg dec -961 if hi == -502\n" +
		"bh inc 972 if esj >= 1157\n" +
		"l dec -854 if jlg <= 985\n" +
		"l dec -966 if h == 170\n" +
		"fw dec 262 if pdg < 223\n" +
		"esj dec 35 if ls != 4603\n" +
		"xq dec 497 if pdg == 217\n" +
		"t inc 998 if hsv < 5471\n" +
		"gf inc 283 if fi > -1553\n" +
		"s dec -592 if kgp < -427\n" +
		"hi inc 883 if l < 1064\n" +
		"s inc -388 if s <= 1751\n" +
		"p dec -950 if pdg == 218\n" +
		"kbm dec 127 if kgp == -435\n" +
		"l dec -217 if w >= -911\n" +
		"w dec -210 if p > 4435\n" +
		"hsv inc 726 if gf <= 946\n" +
		"hi dec -199 if h <= 170\n" +
		"hi inc -481 if fi >= -1541\n" +
		"w dec 295 if t <= 4205\n" +
		"gf inc 791 if jcf == -2692\n" +
		"u dec -988 if bh >= 3096\n" +
		"jlg inc -821 if esj < 1132\n" +
		"fi inc 893 if fw == -165\n" +
		"pvo inc 177 if esj > 1130\n" +
		"bh dec 475 if kgp == -435\n" +
		"ls inc 273 if pvo >= 1049\n" +
		"pvo dec -983 if fw >= -161\n" +
		"s inc 529 if jcf >= -2691\n" +
		"pl dec -154 if xq < -1872\n" +
		"bh inc 211 if pl <= -1164\n" +
		"gf inc 519 if jcf != -2695\n" +
		"esj dec -726 if kbm >= -1851\n" +
		"w inc -251 if jlg != 151\n" +
		"l dec -162 if fw < -163\n" +
		"w inc -25 if h != 174\n" +
		"s inc -303 if kgp == -435\n" +
		"l dec 439 if ls < 4600\n" +
		"fi dec -171 if pdg <= 211\n" +
		"pvo dec 4 if hi > 586\n" +
		"pvo dec 492 if pdg == 211\n" +
		"u inc -355 if gf == 1463\n" +
		"kgp dec 822 if kbm < -1841\n" +
		"bh inc 808 if s < 1592\n" +
		"pl inc -827 if t == 4210\n" +
		"l inc -748 if pdg <= 213\n" +
		"jcf inc 49 if p < 4446\n" +
		"f inc 781 if pdg >= 208\n" +
		"hsv dec -263 if t > 4201\n" +
		"p dec -424 if kgp < -1253\n" +
		"hi inc 528 if ls <= 4606\n" +
		"u inc 240 if ls != 4596\n" +
		"ls inc -2 if h == 170\n" +
		"esj dec -884 if l == 1440\n" +
		"w dec -883 if u <= 2130\n" +
		"kbm inc -450 if kgp <= -1250\n" +
		"jcf inc 38 if f < 1304\n" +
		"ls dec 121 if fw < -162\n" +
		"s inc 139 if jlg == 159\n" +
		"jlg dec 984 if ls < 4491\n" +
		"u inc -837 if bh != 3640\n" +
		"fi inc -56 if p <= 4873\n" +
		"jcf dec 198 if esj == 2741\n" +
		"hsv inc 731 if jcf > -2836\n" +
		"jlg inc 765 if hi > 1104\n" +
		"gf inc 630 if pvo < 1044\n" +
		"s dec 510 if pvo <= 1043\n" +
		"bh dec 168 if esj < 2742\n" +
		"pvo dec 135 if fi >= -713\n" +
		"kgp inc 578 if u > 1279\n" +
		"hi dec -842 if pdg < 213\n" +
		"pvo dec -908 if pdg < 221\n" +
		"p dec 152 if fi >= -718\n" +
		"fi dec -217 if t >= 4206\n" +
		"esj dec -237 if jlg == -60\n" +
		"kbm inc -637 if pvo <= 1946\n" +
		"t inc 324 if gf > 2090\n" +
		"xq inc -950 if h > 179\n" +
		"s inc -285 if u < 1292\n" +
		"h dec 469 if fi >= -503\n" +
		"jlg dec 810 if pvo >= 1948\n" +
		"hi dec 903 if fi != -493\n" +
		"h inc -548 if u < 1295\n" +
		"hi inc 269 if l >= 1438\n" +
		"fi dec 341 if hsv < 7193\n" +
		"xq inc -292 if u == 1289\n" +
		"s dec -183 if ls >= 4480\n" +
		"gf dec -654 if hsv > 7178\n" +
		"pvo dec 871 if l < 1445\n" +
		"kgp inc -251 if xq > -2179\n" +
		"f inc 7 if l >= 1439\n" +
		"pvo inc -116 if f > 1313\n" +
		"fw inc 199 if f > 1310\n" +
		"kbm inc -481 if ls == 4483\n" +
		"h dec -765 if pvo == 962\n" +
		"fw dec 109 if bh > 3475\n" +
		"jcf dec 691 if pl <= -1997\n" +
		"t inc -786 if fw <= -71\n" +
		"f inc 209 if l > 1434\n" +
		"xq inc -11 if h > -83\n" +
		"pdg dec 650 if kbm < -2771\n" +
		"w dec -914 if xq <= -2191\n" +
		"kgp inc 403 if pdg >= -442\n" +
		"f inc -274 if h >= -76\n" +
		"l dec -548 if gf >= 2742\n" +
		"e dec -969 if gf != 2739\n" +
		"hi dec -299 if e > -1705\n" +
		"fw inc 654 if jcf < -3521\n" +
		"bh inc 86 if esj >= 2969\n" +
		"l inc -671 if fw >= 578\n" +
		"f inc 44 if hi < 472\n" +
		"pdg inc 882 if xq <= -2175\n" +
		"fi dec 672 if bh == 3565\n" +
		"hsv dec 916 if f != 1528\n" +
		"fi dec 355 if fw >= 579\n" +
		"pvo dec -34 if kbm < -2776\n" +
		"bh inc 495 if t >= 3739\n" +
		"gf dec -974 if kgp >= -527\n" +
		"pdg inc -220 if xq < -2178\n" +
		"h inc -968 if t < 3750\n" +
		"s dec -868 if esj > 2977\n" +
		"ls inc -988 if h != -1041\n" +
		"kgp inc -730 if jcf > -3528\n" +
		"s inc 354 if esj > 2969\n" +
		"xq dec -358 if pdg > 226\n" +
		"pdg dec 183 if pdg != 238\n" +
		"fi dec -27 if l < 1320\n" +
		"p dec 930 if f <= 1531\n" +
		"esj inc -690 if kgp <= -1251\n" +
		"fw dec 855 if fi >= -1840\n" +
		"fi inc 395 if kgp > -1260\n" +
		"e inc 271 if kgp >= -1255\n" +
		"hi inc 453 if p > 3781\n" +
		"bh inc -111 if hi < 935\n" +
		"fw dec 868 if bh >= 3947\n" +
		"kbm inc 13 if l == 1317\n" +
		"w inc -530 if gf >= 3725\n" +
		"p inc -743 if l <= 1322\n" +
		"t dec 180 if gf >= 3716\n" +
		"hi dec 332 if f != 1532\n" +
		"kgp dec 330 if bh >= 3954\n" +
		"xq dec 709 if kbm != -2772\n" +
		"l inc -44 if hi > 590\n" +
		"u dec -442 if pl != -2006\n" +
		"t dec -2 if hi <= 593\n" +
		"p dec -846 if e >= -1710\n" +
		"e inc 692 if t >= 3565\n" +
		"s inc 548 if xq < -2539\n" +
		"kgp inc -40 if fi != -1450\n" +
		"jlg dec -560 if jlg >= -877\n" +
		"pvo dec 865 if w == -85\n" +
		"fi dec -879 if kgp >= -1304\n" +
		"hsv dec 470 if hsv >= 6265\n" +
		"e dec 240 if p >= 3882\n" +
		"hsv inc -213 if hi == 591\n" +
		"pvo dec 835 if e <= -1250\n" +
		"s inc -382 if l <= 1277\n" +
		"ls dec 322 if t == 3568\n" +
		"jlg inc -864 if bh != 3949\n" +
		"s dec -350 if s <= 1956\n" +
		"fi inc -772 if jlg > -304\n" +
		"ls dec -906 if jlg < -307\n" +
		"e dec -301 if fi > -566\n" +
		"pvo dec -421 if fw != -1148\n" +
		"l dec 159 if xq == -2533\n" +
		"e inc -426 if f < 1523\n" +
		"p inc -161 if jcf == -3519\n" +
		"p dec -180 if w == -86\n" +
		"u inc 164 if pdg >= 38\n" +
		"kgp dec 514 if u <= 1896\n" +
		"fw dec 43 if xq <= -2534\n" +
		"ls inc -549 if hsv >= 5790\n" +
		"jcf dec 999 if kgp <= -1807\n" +
		"u dec -265 if esj <= 2292\n" +
		"pvo dec -637 if pdg == 56\n" +
		"hsv dec 315 if ls < 3538\n" +
		"pl inc -158 if w < -88\n" +
		"jcf dec 832 if f < 1530\n" +
		"fi dec -25 if p < 3891\n" +
		"p inc 311 if pvo > 572\n" +
		"f dec -634 if h <= -1060\n" +
		"f dec 659 if hsv >= 5479\n" +
		"kgp dec 143 if gf <= 3729\n" +
		"hsv inc 16 if fi <= -539\n" +
		"f inc -13 if ls <= 3533\n" +
		"kbm dec 35 if kgp > -1957\n" +
		"w dec -392 if u == 2160\n" +
		"pl dec -367 if hsv < 5494\n" +
		"pdg dec 733 if xq > -2539\n" +
		"fi inc 772 if h <= -1055\n" +
		"jcf inc 925 if p >= 4192\n" +
		"jcf inc 806 if hi >= 587\n" +
		"l dec 267 if hi > 589\n" +
		"esj inc -42 if jlg > -314\n" +
		"f dec 75 if e <= -953\n" +
		"s dec 742 if l <= 848\n" +
		"hi inc 785 if kgp == -1954\n" +
		"s dec 562 if hsv < 5503\n" +
		"s dec -392 if esj <= 2248\n" +
		"fw dec 342 if jcf == -3630\n" +
		"u dec 379 if e >= -963\n" +
		"xq dec 510 if pdg != -693\n" +
		"e dec -436 if u > 1778\n" +
		"jcf dec -307 if hsv != 5499\n" +
		"pvo dec 695 if kbm == -2799\n" +
		"xq inc -425 if e < -513\n" +
		"bh inc -161 if pdg > -692\n" +
		"gf dec -656 if u < 1782\n" +
		"esj dec 6 if esj != 2250\n" +
		"hsv dec 153 if s > 1383\n" +
		"w dec 947 if p > 4187\n" +
		"t inc 616 if ls < 3540\n" +
		"ls dec -110 if w <= -646\n" +
		"f inc -813 if w >= -647\n" +
		"jcf inc 447 if kbm >= -2805\n" +
		"hsv inc 945 if pdg > -680\n" +
		"jlg dec 600 if f >= -45\n" +
		"s inc -138 if xq >= -3474\n" +
		"u inc -39 if kbm != -2805\n" +
		"t inc -161 if s >= 1252\n" +
		"kbm dec 685 if e >= -517\n" +
		"fw inc 431 if jcf == -3179\n" +
		"l dec 791 if h > -1044\n" +
		"esj dec 534 if fw >= -713\n" +
		"pl inc -500 if fi < -538\n" +
		"pdg inc -632 if gf >= 4385\n" +
		"hi inc -269 if fi < -538\n" +
		"ls dec 561 if pdg != -677\n" +
		"hi dec -611 if esj > 1701\n" +
		"gf dec -800 if pdg < -678\n" +
		"s inc 624 if w != -636\n" +
		"hsv inc -935 if h != -1052\n" +
		"jcf inc 343 if e == -519\n" +
		"xq dec -199 if f <= -28\n" +
		"p inc -511 if gf >= 5168\n" +
		"p dec 973 if t <= 4024\n" +
		"gf inc 197 if hi > 1730\n" +
		"jlg inc -239 if l <= 855\n" +
		"hi dec -402 if l >= 845\n" +
		"p dec -125 if bh == 3792\n" +
		"w inc -263 if gf != 5175\n" +
		"xq inc -658 if xq >= -3277\n" +
		"h inc 174 if jlg < -1152\n" +
		"gf dec -81 if t <= 4031\n" +
		"u inc -677 if kgp == -1954\n" +
		"u dec 647 if jlg == -1157\n" +
		"fi dec 783 if fw >= -714\n" +
		"hsv dec -829 if hsv <= 4418\n" +
		"jcf dec -524 if f == -36\n" +
		"ls dec -760 if w >= -911\n" +
		"e inc -690 if pvo <= -122\n" +
		"jcf dec -870 if hi < 2126\n" +
		"f inc 500 if s == 1877\n" +
		"jlg dec 2 if p != 2713\n" +
		"jlg inc 260 if p != 2705\n" +
		"bh inc -167 if pdg >= -677\n" +
		"fw inc -545 if s >= 1870\n" +
		"jcf dec 304 if w != -911\n" +
		"u dec -355 if pvo <= -113\n" +
		"jlg dec 425 if u == 1420\n" +
		"l dec 994 if kgp == -1954\n" +
		"l inc 985 if pl <= -2666\n" +
		"hi dec 870 if ls == 3729\n" +
		"esj dec 552 if f >= 474\n" +
		"bh inc 234 if bh < 3794\n" +
		"s dec -537 if pl <= -2655\n" +
		"t dec -838 if xq <= -3925\n" +
		"t inc 562 if jlg != -1317\n" +
		"hi inc -205 if e > -523\n" +
		"s inc 26 if fw > -1260\n" +
		"kbm dec 52 if pl < -2647\n" +
		"jcf dec -478 if h <= -1048\n" +
		"esj inc 982 if l >= -141\n" +
		"u dec -579 if h <= -1047\n" +
		"fi inc 757 if w < -903\n" +
		"fi dec 56 if hi == 1052\n" +
		"xq dec -443 if u <= 2006\n" +
		"kgp inc -883 if esj <= 1709\n" +
		"hsv dec -429 if xq < -3482\n" +
		"s inc 362 if h > -1059\n" +
		"u dec -45 if l >= -155\n" +
		"hi inc -870 if pdg > -694\n" +
		"fw inc -436 if l < -141\n" +
		"jlg dec 544 if p < 2709\n" +
		"t inc -345 if bh != 4021\n" +
		"h dec -67 if esj == 1705\n" +
		"jcf dec -458 if jlg < -1312\n" +
		"s inc 115 if ls < 3735\n" +
		"kgp dec 884 if p <= 2719\n" +
		"pvo dec 476 if esj <= 1706\n" +
		"f inc -41 if pvo < -583\n" +
		"fw dec 699 if h <= -1041\n" +
		"e inc -922 if gf > 5255\n" +
		"p dec 482 if s > 2912\n" +
		"w inc -240 if fi == -565\n" +
		"f inc -39 if kgp > -3731\n" +
		"ls dec 891 if xq >= -3487\n" +
		"esj inc -711 if fw == -2393\n" +
		"ls dec 653 if fw <= -2390\n" +
		"bh dec 447 if hsv > 5662\n" +
		"ls dec 207 if pdg < -685\n" +
		"hi dec -461 if bh == 3575\n" +
		"pdg inc -499 if w <= -1146\n" +
		"w inc 33 if kbm > -2858\n" +
		"pvo inc -830 if gf < 5265\n" +
		"xq dec -212 if s == 2917\n" +
		"esj inc -124 if hi < 645\n" +
		"hsv dec 661 if u < 2043\n" +
		"xq dec 473 if jlg == -1320\n" +
		"esj dec 18 if l <= -151\n" +
		"kbm dec 184 if kbm <= -2850\n" +
		"gf inc -706 if h < -1058\n" +
		"esj dec 772 if ls < 1982\n" +
		"hi inc 573 if pl > -2660\n" +
		"t dec 612 if h > -1059\n" +
		"u inc 946 if f != 381\n" +
		"t dec 775 if h == -1050\n" +
		"jlg inc -310 if l < -137\n" +
		"w dec -649 if hi == 1213\n" +
		"pl inc -561 if jlg != -1622\n" +
		"bh dec -845 if kgp >= -3718\n" +
		"fi dec 927 if h == -1050\n" +
		"fi inc 240 if s != 2922\n" +
		"pl dec 529 if l >= -153\n" +
		"bh dec -913 if jcf == -810\n" +
		"fw inc 208 if pvo < -1418\n" +
		"e dec -944 if hi > 1206\n" +
		"t inc -627 if e > -505\n" +
		"t inc -524 if s >= 2923\n" +
		"w dec -41 if pvo != -1419\n" +
		"jcf inc 387 if fi < -1250\n" +
		"kbm inc 878 if w == -466\n" +
		"ls inc -595 if kbm >= -2164\n" +
		"kbm dec 41 if bh == 4488\n" +
		"u dec -752 if h == -1050\n" +
		"fw dec 901 if fi > -1254\n" +
		"s inc -939 if h > -1051\n" +
		"l dec 888 if xq < -3268\n" +
		"f inc -13 if p < 2232\n" +
		"esj inc -479 if ls == 1383\n" +
		"pl dec -640 if kgp < -3713\n" +
		"l inc 999 if p == 2231\n" +
		"e inc -493 if s < 1978\n" +
		"p dec -7 if jcf < -416\n" +
		"u dec -496 if l > -40\n" +
		"jcf dec -648 if pdg < -1193\n" +
		"l dec -754 if h >= -1054\n" +
		"w dec 119 if ls >= 1383\n" +
		"t inc 403 if ls <= 1392\n" +
		"pl inc -232 if w == -585\n" +
		"gf inc 807 if w > -590\n" +
		"kbm dec 700 if hi != 1206\n" +
		"esj inc -67 if e < -487\n" +
		"e inc 736 if p >= 2233\n" +
		"e dec 346 if h <= -1043\n" +
		"kbm dec -791 if t != 3467\n" +
		"gf dec -100 if bh == 4488\n" +
		"fi inc 713 if h > -1056\n" +
		"hsv dec 744 if t >= 3466\n" +
		"xq inc -731 if jcf > -416\n" +
		"p dec -974 if l > 708\n" +
		"w inc -485 if ls >= 1381\n" +
		"u inc 982 if fi >= -543\n" +
		"esj dec -651 if jlg != -1624\n" +
		"esj inc -754 if ls <= 1387\n" +
		"bh dec 670 if p != 3203\n" +
		"w inc 121 if bh > 3814\n" +
		"kbm dec 696 if pl > -3342\n" +
		"l inc 800 if hsv >= 4926\n" +
		"hsv inc 232 if kbm > -3601\n" +
		"gf dec 100 if h != -1050\n" +
		"gf inc -403 if t >= 3465\n" +
		"p dec -642 if esj == -1199\n" +
		"jlg dec 807 if e != -99\n" +
		"w inc -428 if fw < -3082\n" +
		"ls inc -893 if hi >= 1210\n" +
		"jcf inc -412 if pvo <= -1411\n" +
		"fw inc 942 if bh > 3809\n" +
		"bh dec 996 if esj <= -1195\n" +
		"pvo dec 388 if h >= -1044\n" +
		"fw inc 354 if p <= 3212\n" +
		"f dec 667 if jcf > -841\n" +
		"f inc 842 if s == 1978\n" +
		"pvo inc 9 if ls == 490\n" +
		"hi dec 355 if ls == 490\n" +
		"pl inc 371 if jcf < -825\n" +
		"hsv dec -741 if t == 3458\n" +
		"xq inc -377 if fi > -537\n" +
		"xq dec 848 if fi <= -540\n" +
		"e inc -855 if t <= 3467\n" +
		"pl inc 207 if xq < -3273\n" +
		"jlg dec -270 if kgp <= -3716\n" +
		"jlg inc 909 if w > -1383\n" +
		"p dec -736 if fi <= -547\n" +
		"l dec 869 if fw == -1790\n" +
		"bh inc -398 if u != 5227\n" +
		"kgp dec -312 if s > 1976\n" +
		"e inc 494 if fi == -539\n" +
		"hi inc -942 if xq == -3272\n" +
		"hsv dec 698 if e >= -471\n" +
		"kbm dec 546 if h <= -1060\n" +
		"ls inc -782 if e > -473\n" +
		"fw dec -245 if p <= 3221\n" +
		"s inc 19 if esj > -1211\n" +
		"pdg dec 394 if esj < -1208\n" +
		"l dec -433 if w >= -1380\n" +
		"e dec -738 if l <= 283\n" +
		"kgp inc 383 if pl < -2962\n" +
		"ls inc -369 if hi > -75\n" +
		"bh dec -724 if bh <= 2429\n" +
		"fi inc -52 if bh != 3154\n" +
		"xq inc -145 if h >= -1053\n" +
		"fi dec 1000 if xq < -3413\n" +
		"pvo inc 591 if jcf < -838\n" +
		"p dec 693 if pl >= -2971\n" +
		"jcf dec 232 if pdg < -1178\n" +
		"s dec -669 if pvo != -1410\n" +
		"kgp inc -102 if l > 280\n" +
		"l dec -421 if w >= -1378\n" +
		"l inc -299 if hi <= -84\n" +
		"esj dec 707 if e <= 270\n" +
		"kgp dec -225 if f >= 538\n" +
		"f dec -845 if hi != -80\n" +
		"hsv dec 870 if l <= 409\n" +
		"e dec 213 if gf == 5762\n" +
		"jlg inc 210 if jcf > -1076\n" +
		"pl dec -436 if kbm > -3604\n" +
		"p dec 431 if jlg != -1037\n" +
		"xq inc -703 if pvo > -1411\n" +
		"fi inc -484 if e >= 54\n" +
		"xq dec -134 if w == -1377\n" +
		"hi dec 640 if e < 48\n" +
		"pdg dec 745 if kgp <= -2909\n" +
		"l inc 444 if s < 2004\n" +
		"hi dec 874 if w <= -1387\n" +
		"jlg dec -178 if f != 1387\n" +
		"t inc -216 if u < 5218\n" +
		"pl inc 601 if fi >= -2077\n" +
		"hi inc -818 if ls <= -286\n" +
		"pvo inc -31 if w < -1368\n" +
		"pvo inc -865 if f < 1395\n" +
		"s dec -626 if t < 3474\n" +
		"h inc 478 if jlg < -858\n" +
		"u inc -390 if jcf <= -1061\n" +
		"h inc -155 if pl != -1921\n" +
		"jlg inc -949 if esj < -1904\n" +
		"xq dec -437 if h > -731\n" +
		"hsv dec -812 if pl < -1925\n" +
		"t inc -940 if pl >= -1933\n" +
		"jcf inc -903 if pvo <= -2303\n" +
		"xq inc -19 if esj >= -1911\n" +
		"w inc 238 if pdg <= -1183\n" +
		"f inc -441 if hsv >= 4408\n" +
		"l dec 15 if s > 2614\n" +
		"hsv dec -238 if pvo == -2306\n" +
		"ls dec 641 if p > 2083\n" +
		"h dec -548 if kgp >= -2899\n" +
		"hsv inc 472 if hsv >= 4637\n" +
		"bh dec 123 if s <= 2632\n" +
		"pdg dec 314 if u >= 4827\n" +
		"pl dec -600 if pl >= -1934\n" +
		"h dec -697 if hsv <= 5104\n" +
		"u dec -458 if s >= 2623\n" +
		"gf dec -688 if w == -1139\n" +
		"jcf dec -739 if xq != -3567\n" +
		"l dec -886 if pdg > -1498\n" +
		"s dec -24 if pvo < -2300\n" +
		"kbm inc -76 if jlg >= -1821\n" +
		"fi inc 992 if pvo != -2299\n" +
		"fi inc 638 if fi >= -1091\n" +
		"u inc 452 if ls > -934\n" +
		"h inc 713 if gf <= 6443\n" +
		"p inc -60 if jlg > -1816\n" +
		"bh dec -100 if hsv < 5117\n" +
		"l inc -290 if kgp <= -2895\n" +
		"pvo dec 900 if gf == 6450\n" +
		"u dec 470 if jcf != -1222\n" +
		"bh inc -730 if fw >= -1547\n" +
		"jcf inc 584 if esj >= -1915\n" +
		"e inc -43 if e >= 57\n" +
		"bh inc -883 if l != 552\n" +
		"fi dec 306 if jcf < -645\n" +
		"pl dec -428 if l > 535\n" +
		"s dec 793 if esj == -1908\n" +
		"pl inc 78 if w == -1139\n" +
		"gf dec 107 if p <= 2031\n" +
		"jcf dec -150 if pl <= -820\n" +
		"w inc -854 if jlg < -1808";

	@Test
	public void testPart1() throws Exception {
		assertEquals(1  , new Day08().part1("b inc 5 if a > 1\n" +
			"a inc 1 if b < 5\n" +
			"c dec -10 if a >= 1\n" +
			"c inc -20 if c == 10"));
	}

	@Test
	public void actualPart1() throws Exception {
		System.out.println(new Day08().part1(INPUT));
	}

	@Test
	public void testPart2() throws Exception {
		assertEquals(10  , new Day08().part2("b inc 5 if a > 1\n" +
			"a inc 1 if b < 5\n" +
			"c dec -10 if a >= 1\n" +
			"c inc -20 if c == 10"));
	}

	@Test
	public void actualPart10() throws Exception {
		System.out.println(new Day08().part2(INPUT));
	}

}
