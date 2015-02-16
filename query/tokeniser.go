//line query/tokeniser.rl:1
package query

import (
	"fmt"
)

//line query/tokeniser.rl:8
//line query/tokeniser.go:13
const sase_start int = 1
const sase_first_final int = 466
const sase_error int = 0

const sase_en_main int = 1

//line query/tokeniser.rl:9
func tokenize(data string) ([]*token, error) {
	var (
		cs        int             // current state
		p         int = 0         // data offset
		pe        int = len(data) // data end offset
		eof       int = pe        // eof offset
		mark      int = -1
		tokens        = make([]*token, 0, 10)
		proposals     = make([]*proposedToken, 0, 100)
	)

	// Add a single token onto the committed tokens list to collect the actual tokens
	root := &token{tt: ttGeneric}
	tokens = append(tokens, root)

	markedText := func() string {
		return data[mark:p]
	}

	propose := func(typ tt) {
		// log.Tracef("[Tokenizer] propose: %s", typ.String())
		proposals = append(proposals, &proposedToken{
			token: &token{
				tt: typ,
			},
			i:  len(tokens),
			pi: len(proposals),
		})
	}

	proposal := func(typ tt) *proposedToken {
		for i := len(proposals) - 1; i >= 0; i-- {
			pt := proposals[i]
			if pt.tt == typ {
				return pt
			}
		}
		return nil
	}

	setText := func(typ tt) string {
		text := markedText()
		proposal(typ).content = text
		return text
	}

	commit := func(typ tt) *token {
		// log.Tracef("[Tokenizer] commit: %s", typ.String())
		pt := proposal(typ)
		t := pt.token
		// Add everything on the tokens list that has been committed since this token was proposed as children of this
		// token
		children := tokens[pt.i:]
		t.children = append(make([]*token, 0, len(children)), children...)
		tokens = append(tokens[:pt.i], t)
		// Remove any "defunct" proposals (like this one)
		// log.Tracef("[Tokenizer] commit removing %d defunct proposals", len(proposals)-pt.pi)
		proposals = proposals[:pt.pi]
		return t
	}

	// Suppress "variable not used"
	var (
		_ = cs
		_ = p
		_ = pe
		_ = eof
		_ = mark
		_ = tokens
		_ = proposals
		_ = propose
		_ = proposal
		_ = setText
		_ = commit
	)

//line query/tokeniser.go:100
	{
		cs = sase_start
	}

//line query/tokeniser.go:105
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 466:
			goto st_case_466
		case 467:
			goto st_case_467
		case 468:
			goto st_case_468
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 469:
			goto st_case_469
		case 470:
			goto st_case_470
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 471:
			goto st_case_471
		case 472:
			goto st_case_472
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 473:
			goto st_case_473
		case 474:
			goto st_case_474
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 475:
			goto st_case_475
		case 45:
			goto st_case_45
		case 476:
			goto st_case_476
		case 46:
			goto st_case_46
		case 477:
			goto st_case_477
		case 478:
			goto st_case_478
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 479:
			goto st_case_479
		case 56:
			goto st_case_56
		case 480:
			goto st_case_480
		case 481:
			goto st_case_481
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 482:
			goto st_case_482
		case 483:
			goto st_case_483
		case 484:
			goto st_case_484
		case 65:
			goto st_case_65
		case 485:
			goto st_case_485
		case 66:
			goto st_case_66
		case 486:
			goto st_case_486
		case 487:
			goto st_case_487
		case 67:
			goto st_case_67
		case 488:
			goto st_case_488
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 489:
			goto st_case_489
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 490:
			goto st_case_490
		case 491:
			goto st_case_491
		case 492:
			goto st_case_492
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 493:
			goto st_case_493
		case 494:
			goto st_case_494
		case 495:
			goto st_case_495
		case 99:
			goto st_case_99
		case 496:
			goto st_case_496
		case 100:
			goto st_case_100
		case 497:
			goto st_case_497
		case 498:
			goto st_case_498
		case 101:
			goto st_case_101
		case 499:
			goto st_case_499
		case 102:
			goto st_case_102
		case 103:
			goto st_case_103
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 112:
			goto st_case_112
		case 113:
			goto st_case_113
		case 500:
			goto st_case_500
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 501:
			goto st_case_501
		case 502:
			goto st_case_502
		case 503:
			goto st_case_503
		case 129:
			goto st_case_129
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 504:
			goto st_case_504
		case 134:
			goto st_case_134
		case 505:
			goto st_case_505
		case 506:
			goto st_case_506
		case 135:
			goto st_case_135
		case 507:
			goto st_case_507
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 508:
			goto st_case_508
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 158:
			goto st_case_158
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
		case 509:
			goto st_case_509
		case 510:
			goto st_case_510
		case 511:
			goto st_case_511
		case 163:
			goto st_case_163
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 166:
			goto st_case_166
		case 167:
			goto st_case_167
		case 168:
			goto st_case_168
		case 512:
			goto st_case_512
		case 169:
			goto st_case_169
		case 513:
			goto st_case_513
		case 514:
			goto st_case_514
		case 170:
			goto st_case_170
		case 515:
			goto st_case_515
		case 171:
			goto st_case_171
		case 172:
			goto st_case_172
		case 173:
			goto st_case_173
		case 174:
			goto st_case_174
		case 175:
			goto st_case_175
		case 176:
			goto st_case_176
		case 177:
			goto st_case_177
		case 178:
			goto st_case_178
		case 179:
			goto st_case_179
		case 180:
			goto st_case_180
		case 181:
			goto st_case_181
		case 182:
			goto st_case_182
		case 516:
			goto st_case_516
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 190:
			goto st_case_190
		case 191:
			goto st_case_191
		case 192:
			goto st_case_192
		case 193:
			goto st_case_193
		case 194:
			goto st_case_194
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 197:
			goto st_case_197
		case 517:
			goto st_case_517
		case 518:
			goto st_case_518
		case 519:
			goto st_case_519
		case 198:
			goto st_case_198
		case 199:
			goto st_case_199
		case 200:
			goto st_case_200
		case 201:
			goto st_case_201
		case 202:
			goto st_case_202
		case 520:
			goto st_case_520
		case 521:
			goto st_case_521
		case 203:
			goto st_case_203
		case 204:
			goto st_case_204
		case 205:
			goto st_case_205
		case 206:
			goto st_case_206
		case 207:
			goto st_case_207
		case 208:
			goto st_case_208
		case 209:
			goto st_case_209
		case 210:
			goto st_case_210
		case 522:
			goto st_case_522
		case 211:
			goto st_case_211
		case 212:
			goto st_case_212
		case 523:
			goto st_case_523
		case 524:
			goto st_case_524
		case 213:
			goto st_case_213
		case 525:
			goto st_case_525
		case 526:
			goto st_case_526
		case 214:
			goto st_case_214
		case 215:
			goto st_case_215
		case 216:
			goto st_case_216
		case 217:
			goto st_case_217
		case 218:
			goto st_case_218
		case 219:
			goto st_case_219
		case 220:
			goto st_case_220
		case 221:
			goto st_case_221
		case 527:
			goto st_case_527
		case 528:
			goto st_case_528
		case 529:
			goto st_case_529
		case 222:
			goto st_case_222
		case 530:
			goto st_case_530
		case 223:
			goto st_case_223
		case 531:
			goto st_case_531
		case 532:
			goto st_case_532
		case 224:
			goto st_case_224
		case 533:
			goto st_case_533
		case 225:
			goto st_case_225
		case 226:
			goto st_case_226
		case 227:
			goto st_case_227
		case 228:
			goto st_case_228
		case 229:
			goto st_case_229
		case 230:
			goto st_case_230
		case 231:
			goto st_case_231
		case 232:
			goto st_case_232
		case 233:
			goto st_case_233
		case 234:
			goto st_case_234
		case 235:
			goto st_case_235
		case 236:
			goto st_case_236
		case 534:
			goto st_case_534
		case 237:
			goto st_case_237
		case 238:
			goto st_case_238
		case 239:
			goto st_case_239
		case 240:
			goto st_case_240
		case 241:
			goto st_case_241
		case 242:
			goto st_case_242
		case 243:
			goto st_case_243
		case 244:
			goto st_case_244
		case 245:
			goto st_case_245
		case 246:
			goto st_case_246
		case 247:
			goto st_case_247
		case 248:
			goto st_case_248
		case 249:
			goto st_case_249
		case 250:
			goto st_case_250
		case 251:
			goto st_case_251
		case 535:
			goto st_case_535
		case 536:
			goto st_case_536
		case 537:
			goto st_case_537
		case 252:
			goto st_case_252
		case 253:
			goto st_case_253
		case 254:
			goto st_case_254
		case 255:
			goto st_case_255
		case 538:
			goto st_case_538
		case 256:
			goto st_case_256
		case 539:
			goto st_case_539
		case 257:
			goto st_case_257
		case 540:
			goto st_case_540
		case 541:
			goto st_case_541
		case 258:
			goto st_case_258
		case 542:
			goto st_case_542
		case 259:
			goto st_case_259
		case 260:
			goto st_case_260
		case 261:
			goto st_case_261
		case 262:
			goto st_case_262
		case 263:
			goto st_case_263
		case 264:
			goto st_case_264
		case 265:
			goto st_case_265
		case 266:
			goto st_case_266
		case 267:
			goto st_case_267
		case 268:
			goto st_case_268
		case 269:
			goto st_case_269
		case 270:
			goto st_case_270
		case 543:
			goto st_case_543
		case 271:
			goto st_case_271
		case 272:
			goto st_case_272
		case 273:
			goto st_case_273
		case 274:
			goto st_case_274
		case 275:
			goto st_case_275
		case 276:
			goto st_case_276
		case 277:
			goto st_case_277
		case 278:
			goto st_case_278
		case 279:
			goto st_case_279
		case 280:
			goto st_case_280
		case 281:
			goto st_case_281
		case 282:
			goto st_case_282
		case 283:
			goto st_case_283
		case 284:
			goto st_case_284
		case 285:
			goto st_case_285
		case 544:
			goto st_case_544
		case 545:
			goto st_case_545
		case 546:
			goto st_case_546
		case 286:
			goto st_case_286
		case 287:
			goto st_case_287
		case 288:
			goto st_case_288
		case 289:
			goto st_case_289
		case 290:
			goto st_case_290
		case 291:
			goto st_case_291
		case 292:
			goto st_case_292
		case 547:
			goto st_case_547
		case 548:
			goto st_case_548
		case 293:
			goto st_case_293
		case 294:
			goto st_case_294
		case 295:
			goto st_case_295
		case 296:
			goto st_case_296
		case 297:
			goto st_case_297
		case 298:
			goto st_case_298
		case 299:
			goto st_case_299
		case 300:
			goto st_case_300
		case 301:
			goto st_case_301
		case 302:
			goto st_case_302
		case 549:
			goto st_case_549
		case 550:
			goto st_case_550
		case 303:
			goto st_case_303
		case 551:
			goto st_case_551
		case 552:
			goto st_case_552
		case 304:
			goto st_case_304
		case 305:
			goto st_case_305
		case 306:
			goto st_case_306
		case 307:
			goto st_case_307
		case 308:
			goto st_case_308
		case 309:
			goto st_case_309
		case 310:
			goto st_case_310
		case 311:
			goto st_case_311
		case 553:
			goto st_case_553
		case 554:
			goto st_case_554
		case 555:
			goto st_case_555
		case 312:
			goto st_case_312
		case 556:
			goto st_case_556
		case 313:
			goto st_case_313
		case 557:
			goto st_case_557
		case 558:
			goto st_case_558
		case 314:
			goto st_case_314
		case 559:
			goto st_case_559
		case 315:
			goto st_case_315
		case 316:
			goto st_case_316
		case 317:
			goto st_case_317
		case 318:
			goto st_case_318
		case 319:
			goto st_case_319
		case 320:
			goto st_case_320
		case 321:
			goto st_case_321
		case 322:
			goto st_case_322
		case 323:
			goto st_case_323
		case 324:
			goto st_case_324
		case 325:
			goto st_case_325
		case 326:
			goto st_case_326
		case 560:
			goto st_case_560
		case 327:
			goto st_case_327
		case 328:
			goto st_case_328
		case 329:
			goto st_case_329
		case 330:
			goto st_case_330
		case 331:
			goto st_case_331
		case 332:
			goto st_case_332
		case 333:
			goto st_case_333
		case 334:
			goto st_case_334
		case 335:
			goto st_case_335
		case 336:
			goto st_case_336
		case 337:
			goto st_case_337
		case 338:
			goto st_case_338
		case 339:
			goto st_case_339
		case 340:
			goto st_case_340
		case 341:
			goto st_case_341
		case 561:
			goto st_case_561
		case 562:
			goto st_case_562
		case 563:
			goto st_case_563
		case 342:
			goto st_case_342
		case 343:
			goto st_case_343
		case 344:
			goto st_case_344
		case 345:
			goto st_case_345
		case 564:
			goto st_case_564
		case 565:
			goto st_case_565
		case 346:
			goto st_case_346
		case 566:
			goto st_case_566
		case 347:
			goto st_case_347
		case 567:
			goto st_case_567
		case 568:
			goto st_case_568
		case 348:
			goto st_case_348
		case 569:
			goto st_case_569
		case 349:
			goto st_case_349
		case 350:
			goto st_case_350
		case 351:
			goto st_case_351
		case 352:
			goto st_case_352
		case 353:
			goto st_case_353
		case 354:
			goto st_case_354
		case 355:
			goto st_case_355
		case 356:
			goto st_case_356
		case 357:
			goto st_case_357
		case 358:
			goto st_case_358
		case 359:
			goto st_case_359
		case 360:
			goto st_case_360
		case 570:
			goto st_case_570
		case 361:
			goto st_case_361
		case 362:
			goto st_case_362
		case 363:
			goto st_case_363
		case 364:
			goto st_case_364
		case 365:
			goto st_case_365
		case 366:
			goto st_case_366
		case 367:
			goto st_case_367
		case 368:
			goto st_case_368
		case 369:
			goto st_case_369
		case 370:
			goto st_case_370
		case 371:
			goto st_case_371
		case 372:
			goto st_case_372
		case 373:
			goto st_case_373
		case 374:
			goto st_case_374
		case 375:
			goto st_case_375
		case 571:
			goto st_case_571
		case 572:
			goto st_case_572
		case 573:
			goto st_case_573
		case 376:
			goto st_case_376
		case 377:
			goto st_case_377
		case 378:
			goto st_case_378
		case 379:
			goto st_case_379
		case 380:
			goto st_case_380
		case 381:
			goto st_case_381
		case 382:
			goto st_case_382
		case 383:
			goto st_case_383
		case 384:
			goto st_case_384
		case 385:
			goto st_case_385
		case 386:
			goto st_case_386
		case 387:
			goto st_case_387
		case 388:
			goto st_case_388
		case 389:
			goto st_case_389
		case 390:
			goto st_case_390
		case 391:
			goto st_case_391
		case 392:
			goto st_case_392
		case 393:
			goto st_case_393
		case 394:
			goto st_case_394
		case 395:
			goto st_case_395
		case 396:
			goto st_case_396
		case 397:
			goto st_case_397
		case 398:
			goto st_case_398
		case 399:
			goto st_case_399
		case 400:
			goto st_case_400
		case 401:
			goto st_case_401
		case 402:
			goto st_case_402
		case 403:
			goto st_case_403
		case 404:
			goto st_case_404
		case 405:
			goto st_case_405
		case 574:
			goto st_case_574
		case 406:
			goto st_case_406
		case 407:
			goto st_case_407
		case 408:
			goto st_case_408
		case 409:
			goto st_case_409
		case 410:
			goto st_case_410
		case 575:
			goto st_case_575
		case 411:
			goto st_case_411
		case 412:
			goto st_case_412
		case 413:
			goto st_case_413
		case 414:
			goto st_case_414
		case 415:
			goto st_case_415
		case 416:
			goto st_case_416
		case 417:
			goto st_case_417
		case 418:
			goto st_case_418
		case 419:
			goto st_case_419
		case 420:
			goto st_case_420
		case 421:
			goto st_case_421
		case 422:
			goto st_case_422
		case 423:
			goto st_case_423
		case 424:
			goto st_case_424
		case 425:
			goto st_case_425
		case 576:
			goto st_case_576
		case 426:
			goto st_case_426
		case 427:
			goto st_case_427
		case 428:
			goto st_case_428
		case 429:
			goto st_case_429
		case 430:
			goto st_case_430
		case 431:
			goto st_case_431
		case 432:
			goto st_case_432
		case 433:
			goto st_case_433
		case 434:
			goto st_case_434
		case 435:
			goto st_case_435
		case 436:
			goto st_case_436
		case 437:
			goto st_case_437
		case 438:
			goto st_case_438
		case 439:
			goto st_case_439
		case 440:
			goto st_case_440
		case 441:
			goto st_case_441
		case 442:
			goto st_case_442
		case 443:
			goto st_case_443
		case 444:
			goto st_case_444
		case 445:
			goto st_case_445
		case 446:
			goto st_case_446
		case 447:
			goto st_case_447
		case 448:
			goto st_case_448
		case 449:
			goto st_case_449
		case 450:
			goto st_case_450
		case 451:
			goto st_case_451
		case 452:
			goto st_case_452
		case 453:
			goto st_case_453
		case 454:
			goto st_case_454
		case 455:
			goto st_case_455
		case 456:
			goto st_case_456
		case 457:
			goto st_case_457
		case 458:
			goto st_case_458
		case 459:
			goto st_case_459
		case 460:
			goto st_case_460
		case 461:
			goto st_case_461
		case 462:
			goto st_case_462
		case 463:
			goto st_case_463
		case 464:
			goto st_case_464
		case 465:
			goto st_case_465
		}
		goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st1
		case 69:
			goto st2
		case 101:
			goto st2
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 86:
			goto st3
		case 118:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 69:
			goto st4
		case 101:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 78:
			goto st5
		case 110:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 84:
			goto st6
		case 116:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 32 {
			goto st7
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 32:
			goto st7
		case 33:
			goto tr8
		case 65:
			goto tr9
		case 78:
			goto tr11
		case 83:
			goto tr12
		case 95:
			goto tr10
		case 97:
			goto tr9
		case 110:
			goto tr11
		case 115:
			goto tr12
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st7
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto st0
	tr8:
//line query/tokeniser.rl:151
		propose(ttEventClause)
//line query/tokeniser.rl:131
		propose(ttNegatedDecl)
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line query/tokeniser.go:1397
		switch data[p] {
		case 32:
			goto st9
		case 40:
			goto st10
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 40 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 32:
			goto st10
		case 41:
			goto st466
		case 65:
			goto tr16
		case 95:
			goto tr17
		case 97:
			goto tr16
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st10
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto st0
	tr918:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st466
	tr929:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st466
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
//line query/tokeniser.go:1464
		switch data[p] {
		case 32:
			goto tr1043
		case 59:
			goto tr1044
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1043
		}
		goto st0
	tr1043:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st467
	tr1299:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st467
	tr1302:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st467
	tr1304:
//line query/tokeniser.rl:146
		commit(ttSeqDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
//line query/tokeniser.go:1508
		switch data[p] {
		case 32:
			goto st467
		case 59:
			goto st468
		case 87:
			goto st11
		case 119:
			goto st11
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st467
		}
		goto st0
	tr1044:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st468
	tr1049:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st468
	tr1074:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st468
	tr1157:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st468
	tr1161:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st468
	tr1164:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st468
	tr1166:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st468
	tr1169:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st468
	tr1301:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st468
	tr1303:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st468
	tr1305:
//line query/tokeniser.rl:146
		commit(ttSeqDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line query/tokeniser.go:1606
		if data[p] == 32 {
			goto st468
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st468
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 72:
			goto st12
		case 73:
			goto st188
		case 104:
			goto st12
		case 105:
			goto st188
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 69:
			goto st13
		case 101:
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 82:
			goto st14
		case 114:
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 69:
			goto st15
		case 101:
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 32 {
			goto st16
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 32:
			goto st16
		case 91:
			goto st179
		case 95:
			goto tr24
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st16
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto st0
	tr24:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line query/tokeniser.go:1715
		switch data[p] {
		case 32:
			goto tr26
		case 33:
			goto tr27
		case 46:
			goto tr28
		case 60:
			goto tr30
		case 61:
			goto tr31
		case 62:
			goto tr32
		case 95:
			goto st17
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr26
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st17
				}
			case data[p] >= 65:
				goto st17
			}
		default:
			goto st17
		}
		goto st0
	tr26:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st18
	tr906:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line query/tokeniser.go:1769
		switch data[p] {
		case 32:
			goto st18
		case 33:
			goto tr34
		case 60:
			goto tr35
		case 61:
			goto tr36
		case 62:
			goto tr37
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st18
		}
		goto st0
	tr27:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st19
	tr34:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st19
	tr907:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line query/tokeniser.go:1813
		if data[p] == 61 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 32:
			goto tr39
		case 34:
			goto tr40
		case 39:
			goto tr41
		case 43:
			goto tr42
		case 45:
			goto tr42
		case 95:
			goto tr44
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr39
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr44
				}
			case data[p] >= 65:
				goto tr44
			}
		default:
			goto tr43
		}
		goto st0
	tr39:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st21
	tr887:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st21
	tr891:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st21
	tr895:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st21
	tr898:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st21
	tr902:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line query/tokeniser.go:1884
		switch data[p] {
		case 32:
			goto st21
		case 34:
			goto tr46
		case 39:
			goto tr47
		case 43:
			goto tr48
		case 45:
			goto tr48
		case 95:
			goto tr50
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st21
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr50
				}
			case data[p] >= 65:
				goto tr50
			}
		default:
			goto tr49
		}
		goto st0
	tr40:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr46:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr888:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr892:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr896:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr899:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr903:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line query/tokeniser.go:1962
		switch data[p] {
		case 34:
			goto tr52
		case 92:
			goto tr53
		}
		goto tr51
	tr51:
//line query/tokeniser.rl:87
		mark = p
		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line query/tokeniser.go:1979
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		goto st23
	tr52:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st469
	tr55:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st469
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
//line query/tokeniser.go:2002
		switch data[p] {
		case 32:
			goto tr1048
		case 59:
			goto tr1049
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1048
		}
		goto st0
	tr1048:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st470
	tr1073:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st470
	tr1155:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st470
	tr1158:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st470
	tr1162:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st470
	tr1165:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st470
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
//line query/tokeniser.go:2060
		switch data[p] {
		case 32:
			goto st470
		case 38:
			goto tr1051
		case 59:
			goto st468
		case 65:
			goto tr1052
		case 79:
			goto tr1053
		case 87:
			goto st187
		case 94:
			goto tr1055
		case 97:
			goto tr1052
		case 111:
			goto tr1053
		case 119:
			goto st187
		case 124:
			goto tr1056
		case 226:
			goto tr1057
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st470
		}
		goto st0
	tr1051:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line query/tokeniser.go:2100
		if data[p] == 38 {
			goto st25
		}
		goto st0
	tr1055:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line query/tokeniser.go:2114
		if data[p] == 32 {
			goto tr58
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr58
		}
		goto st0
	tr58:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st26
	tr453:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line query/tokeniser.go:2135
		switch data[p] {
		case 32:
			goto st26
		case 91:
			goto st179
		case 95:
			goto tr60
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr60
			}
		default:
			goto tr60
		}
		goto st0
	tr60:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line query/tokeniser.go:2168
		switch data[p] {
		case 32:
			goto tr61
		case 33:
			goto tr62
		case 46:
			goto tr63
		case 60:
			goto tr65
		case 61:
			goto tr66
		case 62:
			goto tr67
		case 95:
			goto st27
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr61
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st27
				}
			case data[p] >= 65:
				goto st27
			}
		default:
			goto st27
		}
		goto st0
	tr61:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st28
	tr437:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line query/tokeniser.go:2222
		switch data[p] {
		case 32:
			goto st28
		case 33:
			goto tr69
		case 60:
			goto tr70
		case 61:
			goto tr71
		case 62:
			goto tr72
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st28
		}
		goto st0
	tr62:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st29
	tr69:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st29
	tr438:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line query/tokeniser.go:2266
		if data[p] == 61 {
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 32:
			goto tr74
		case 34:
			goto tr75
		case 39:
			goto tr76
		case 43:
			goto tr42
		case 45:
			goto tr42
		case 95:
			goto tr44
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr74
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr44
				}
			case data[p] >= 65:
				goto tr44
			}
		default:
			goto tr43
		}
		goto st0
	tr74:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st31
	tr403:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st31
	tr410:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st31
	tr417:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st31
	tr423:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st31
	tr430:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line query/tokeniser.go:2337
		switch data[p] {
		case 32:
			goto st31
		case 34:
			goto tr78
		case 39:
			goto tr79
		case 43:
			goto tr48
		case 45:
			goto tr48
		case 95:
			goto tr50
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st31
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr50
				}
			case data[p] >= 65:
				goto tr50
			}
		default:
			goto tr49
		}
		goto st0
	tr75:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr78:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr404:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr411:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr418:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr424:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr431:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line query/tokeniser.go:2415
		switch data[p] {
		case 34:
			goto tr52
		case 92:
			goto tr81
		}
		goto tr80
	tr80:
//line query/tokeniser.rl:87
		mark = p
		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line query/tokeniser.go:2432
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		goto st33
	tr81:
//line query/tokeniser.rl:87
		mark = p
		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line query/tokeniser.go:2449
		switch data[p] {
		case 34:
			goto tr84
		case 92:
			goto st34
		}
		goto st33
	tr84:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st471
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
//line query/tokeniser.go:2466
		switch data[p] {
		case 32:
			goto tr1058
		case 34:
			goto tr55
		case 59:
			goto tr1059
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1058
		}
		goto st33
	tr1058:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st472
	tr1069:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st472
	tr1086:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st472
	tr1139:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st472
	tr1142:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st472
	tr1146:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st472
	tr1149:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st472
	tr1182:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st472
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
//line query/tokeniser.go:2544
		switch data[p] {
		case 32:
			goto st472
		case 34:
			goto tr55
		case 38:
			goto tr1061
		case 59:
			goto st474
		case 65:
			goto tr1063
		case 79:
			goto tr1064
		case 87:
			goto st152
		case 92:
			goto st34
		case 94:
			goto tr1066
		case 97:
			goto tr1063
		case 111:
			goto tr1064
		case 119:
			goto st152
		case 124:
			goto tr1067
		case 226:
			goto tr1068
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st472
		}
		goto st33
	tr1061:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line query/tokeniser.go:2588
		switch data[p] {
		case 34:
			goto tr55
		case 38:
			goto st36
		case 92:
			goto st34
		}
		goto st33
	tr1066:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line query/tokeniser.go:2607
		switch data[p] {
		case 32:
			goto tr86
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr86
		}
		goto st33
	tr86:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st37
	tr381:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line query/tokeniser.go:2633
		switch data[p] {
		case 32:
			goto st37
		case 34:
			goto tr55
		case 91:
			goto st144
		case 92:
			goto st34
		case 95:
			goto tr88
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr88
			}
		default:
			goto tr88
		}
		goto st33
	tr88:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line query/tokeniser.go:2670
		switch data[p] {
		case 32:
			goto tr90
		case 33:
			goto tr91
		case 34:
			goto tr55
		case 46:
			goto tr92
		case 60:
			goto tr94
		case 61:
			goto tr95
		case 62:
			goto tr96
		case 92:
			goto st34
		case 95:
			goto st38
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr90
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st38
				}
			case data[p] >= 65:
				goto st38
			}
		default:
			goto st38
		}
		goto st33
	tr90:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st39
	tr365:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line query/tokeniser.go:2728
		switch data[p] {
		case 32:
			goto st39
		case 33:
			goto tr98
		case 34:
			goto tr55
		case 60:
			goto tr99
		case 61:
			goto tr100
		case 62:
			goto tr101
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st39
		}
		goto st33
	tr91:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st40
	tr98:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st40
	tr366:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line query/tokeniser.go:2776
		switch data[p] {
		case 34:
			goto tr55
		case 61:
			goto st41
		case 92:
			goto st34
		}
		goto st33
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 32:
			goto tr103
		case 34:
			goto tr104
		case 39:
			goto tr105
		case 43:
			goto tr106
		case 45:
			goto tr106
		case 92:
			goto st34
		case 95:
			goto tr108
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr103
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr108
				}
			case data[p] >= 65:
				goto tr108
			}
		default:
			goto tr107
		}
		goto st33
	tr103:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st42
	tr331:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st42
	tr338:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st42
	tr345:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st42
	tr351:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st42
	tr358:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line query/tokeniser.go:2854
		switch data[p] {
		case 32:
			goto st42
		case 34:
			goto tr110
		case 39:
			goto tr111
		case 43:
			goto tr112
		case 45:
			goto tr112
		case 92:
			goto st34
		case 95:
			goto tr114
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st42
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr114
				}
			case data[p] >= 65:
				goto tr114
			}
		default:
			goto tr113
		}
		goto st33
	tr104:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st473
	tr110:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st473
	tr332:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st473
	tr339:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st473
	tr346:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st473
	tr352:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st473
	tr359:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st473
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
//line query/tokeniser.go:2948
		switch data[p] {
		case 32:
			goto tr1069
		case 34:
			goto tr52
		case 59:
			goto tr1070
		case 92:
			goto tr81
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1069
		}
		goto tr80
	tr1059:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st474
	tr1070:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st474
	tr1087:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st474
	tr1141:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st474
	tr1145:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st474
	tr1148:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st474
	tr1150:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st474
	tr1153:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st474
	tr1183:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st474
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
//line query/tokeniser.go:3034
		switch data[p] {
		case 32:
			goto st474
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st474
		}
		goto st33
	tr105:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr111:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr333:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr340:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr347:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr353:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr360:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line query/tokeniser.go:3092
		switch data[p] {
		case 34:
			goto tr116
		case 39:
			goto tr117
		case 92:
			goto tr118
		}
		goto tr115
	tr115:
//line query/tokeniser.rl:87
		mark = p
		goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line query/tokeniser.go:3111
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		goto st44
	tr116:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st475
	tr120:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st475
	tr500:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:87
		mark = p
		goto st475
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
//line query/tokeniser.go:3142
		switch data[p] {
		case 32:
			goto tr1071
		case 39:
			goto tr124
		case 59:
			goto tr1072
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1071
		}
		goto st45
	tr397:
//line query/tokeniser.rl:87
		mark = p
		goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line query/tokeniser.go:3166
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		goto st45
	tr398:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st476
	tr124:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st476
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
//line query/tokeniser.go:3189
		switch data[p] {
		case 32:
			goto tr1073
		case 59:
			goto tr1074
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1073
		}
		goto st0
	tr399:
//line query/tokeniser.rl:87
		mark = p
		goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line query/tokeniser.go:3209
		switch data[p] {
		case 39:
			goto tr126
		case 92:
			goto st46
		}
		goto st45
	tr126:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st477
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
//line query/tokeniser.go:3226
		switch data[p] {
		case 32:
			goto tr1075
		case 39:
			goto tr124
		case 59:
			goto tr1076
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1075
		}
		goto st45
	tr1071:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st478
	tr1075:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st478
	tr1121:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st478
	tr1123:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st478
	tr1126:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st478
	tr1130:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st478
	tr1133:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st478
	tr1281:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st478
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
//line query/tokeniser.go:3304
		switch data[p] {
		case 32:
			goto st478
		case 38:
			goto tr1078
		case 39:
			goto tr124
		case 59:
			goto st495
		case 65:
			goto tr1080
		case 79:
			goto tr1081
		case 87:
			goto st118
		case 92:
			goto st46
		case 94:
			goto tr1083
		case 97:
			goto tr1080
		case 111:
			goto tr1081
		case 119:
			goto st118
		case 124:
			goto tr1084
		case 226:
			goto tr1085
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st478
		}
		goto st45
	tr1078:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line query/tokeniser.go:3348
		switch data[p] {
		case 38:
			goto st48
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		goto st45
	tr1083:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line query/tokeniser.go:3367
		switch data[p] {
		case 32:
			goto tr128
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr128
		}
		goto st45
	tr128:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st49
	tr312:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line query/tokeniser.go:3393
		switch data[p] {
		case 32:
			goto st49
		case 39:
			goto tr124
		case 91:
			goto st110
		case 92:
			goto st46
		case 95:
			goto tr130
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st49
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr130
			}
		default:
			goto tr130
		}
		goto st45
	tr130:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line query/tokeniser.go:3430
		switch data[p] {
		case 32:
			goto tr132
		case 33:
			goto tr133
		case 39:
			goto tr124
		case 46:
			goto tr134
		case 60:
			goto tr136
		case 61:
			goto tr137
		case 62:
			goto tr138
		case 92:
			goto st46
		case 95:
			goto st50
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr132
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st50
				}
			case data[p] >= 65:
				goto st50
			}
		default:
			goto st50
		}
		goto st45
	tr132:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st51
	tr296:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line query/tokeniser.go:3488
		switch data[p] {
		case 32:
			goto st51
		case 33:
			goto tr140
		case 39:
			goto tr124
		case 60:
			goto tr141
		case 61:
			goto tr142
		case 62:
			goto tr143
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st51
		}
		goto st45
	tr133:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st52
	tr140:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st52
	tr297:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line query/tokeniser.go:3536
		switch data[p] {
		case 39:
			goto tr124
		case 61:
			goto st53
		case 92:
			goto st46
		}
		goto st45
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 32:
			goto tr145
		case 34:
			goto tr146
		case 39:
			goto tr147
		case 43:
			goto tr148
		case 45:
			goto tr148
		case 92:
			goto st46
		case 95:
			goto tr150
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr145
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr150
				}
			case data[p] >= 65:
				goto tr150
			}
		default:
			goto tr149
		}
		goto st45
	tr145:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st54
	tr262:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st54
	tr269:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st54
	tr276:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st54
	tr282:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st54
	tr289:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line query/tokeniser.go:3614
		switch data[p] {
		case 32:
			goto st54
		case 34:
			goto tr152
		case 39:
			goto tr153
		case 43:
			goto tr154
		case 45:
			goto tr154
		case 92:
			goto st46
		case 95:
			goto tr156
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st54
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr156
				}
			case data[p] >= 65:
				goto tr156
			}
		default:
			goto tr155
		}
		goto st45
	tr146:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr152:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr263:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr270:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr277:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr283:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr290:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line query/tokeniser.go:3694
		switch data[p] {
		case 34:
			goto tr116
		case 39:
			goto tr157
		case 92:
			goto tr118
		}
		goto tr115
	tr117:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st479
	tr121:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st479
	tr157:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:87
		mark = p
		goto st479
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
//line query/tokeniser.go:3725
		switch data[p] {
		case 32:
			goto tr1086
		case 34:
			goto tr55
		case 59:
			goto tr1087
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1086
		}
		goto st33
	tr118:
//line query/tokeniser.rl:87
		mark = p
		goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line query/tokeniser.go:3749
		switch data[p] {
		case 34:
			goto tr158
		case 39:
			goto tr159
		case 92:
			goto st56
		}
		goto st44
	tr158:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st480
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
//line query/tokeniser.go:3768
		switch data[p] {
		case 32:
			goto tr1088
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 59:
			goto tr1089
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1088
		}
		goto st44
	tr1088:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st481
	tr1099:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st481
	tr1119:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st481
	tr1101:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st481
	tr1103:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st481
	tr1106:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st481
	tr1110:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st481
	tr1113:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st481
	tr1197:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st481
	tr1261:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st481
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
//line query/tokeniser.go:3864
		switch data[p] {
		case 32:
			goto st481
		case 34:
			goto tr120
		case 38:
			goto tr1091
		case 39:
			goto tr121
		case 59:
			goto st483
		case 65:
			goto tr1093
		case 79:
			goto tr1094
		case 87:
			goto st84
		case 92:
			goto st56
		case 94:
			goto tr1096
		case 97:
			goto tr1093
		case 111:
			goto tr1094
		case 119:
			goto st84
		case 124:
			goto tr1097
		case 226:
			goto tr1098
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st481
		}
		goto st44
	tr1091:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line query/tokeniser.go:3910
		switch data[p] {
		case 34:
			goto tr120
		case 38:
			goto st58
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		goto st44
	tr1096:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line query/tokeniser.go:3931
		switch data[p] {
		case 32:
			goto tr161
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr161
		}
		goto st44
	tr161:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st59
	tr243:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line query/tokeniser.go:3959
		switch data[p] {
		case 32:
			goto st59
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 91:
			goto st76
		case 92:
			goto st56
		case 95:
			goto tr163
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st59
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr163
			}
		default:
			goto tr163
		}
		goto st44
	tr163:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line query/tokeniser.go:3998
		switch data[p] {
		case 32:
			goto tr165
		case 33:
			goto tr166
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 46:
			goto tr167
		case 60:
			goto tr169
		case 61:
			goto tr170
		case 62:
			goto tr171
		case 92:
			goto st56
		case 95:
			goto st60
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr165
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st60
				}
			case data[p] >= 65:
				goto st60
			}
		default:
			goto st60
		}
		goto st44
	tr165:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st61
	tr227:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line query/tokeniser.go:4058
		switch data[p] {
		case 32:
			goto st61
		case 33:
			goto tr173
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 60:
			goto tr174
		case 61:
			goto tr175
		case 62:
			goto tr176
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st61
		}
		goto st44
	tr166:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st62
	tr173:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st62
	tr228:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line query/tokeniser.go:4108
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 61:
			goto st63
		case 92:
			goto st56
		}
		goto st44
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 32:
			goto tr178
		case 34:
			goto tr179
		case 39:
			goto tr180
		case 43:
			goto tr181
		case 45:
			goto tr181
		case 92:
			goto st56
		case 95:
			goto tr183
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr178
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr183
				}
			case data[p] >= 65:
				goto tr183
			}
		default:
			goto tr182
		}
		goto st44
	tr178:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st64
	tr193:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st64
	tr200:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st64
	tr207:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st64
	tr213:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st64
	tr220:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line query/tokeniser.go:4188
		switch data[p] {
		case 32:
			goto st64
		case 34:
			goto tr185
		case 39:
			goto tr186
		case 43:
			goto tr187
		case 45:
			goto tr187
		case 92:
			goto st56
		case 95:
			goto tr189
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st64
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr189
				}
			case data[p] >= 65:
				goto tr189
			}
		default:
			goto tr188
		}
		goto st44
	tr179:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st482
	tr185:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st482
	tr194:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st482
	tr201:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st482
	tr208:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st482
	tr214:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st482
	tr221:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st482
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
//line query/tokeniser.go:4282
		switch data[p] {
		case 32:
			goto tr1099
		case 34:
			goto tr116
		case 39:
			goto tr157
		case 59:
			goto tr1100
		case 92:
			goto tr118
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1099
		}
		goto tr115
	tr1089:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st483
	tr1100:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st483
	tr1120:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st483
	tr1102:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st483
	tr1105:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st483
	tr1109:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st483
	tr1112:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st483
	tr1114:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st483
	tr1117:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st483
	tr1198:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st483
	tr1262:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st483
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
//line query/tokeniser.go:4386
		switch data[p] {
		case 32:
			goto st483
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st483
		}
		goto st44
	tr180:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st484
	tr186:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st484
	tr195:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st484
	tr202:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st484
	tr209:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st484
	tr215:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st484
	tr222:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st484
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
//line query/tokeniser.go:4460
		switch data[p] {
		case 32:
			goto tr1101
		case 34:
			goto tr116
		case 39:
			goto tr117
		case 59:
			goto tr1102
		case 92:
			goto tr118
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1101
		}
		goto tr115
	tr181:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr187:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr196:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr203:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr210:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr216:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr223:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line query/tokeniser.go:4536
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st485
		}
		goto st44
	tr182:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st485
	tr188:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st485
	tr197:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st485
	tr204:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st485
	tr211:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st485
	tr217:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st485
	tr224:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
//line query/tokeniser.go:4608
		switch data[p] {
		case 32:
			goto tr1103
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 46:
			goto st66
		case 59:
			goto tr1105
		case 92:
			goto st56
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st485
			}
		case data[p] >= 9:
			goto tr1103
		}
		goto st44
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st486
		}
		goto st44
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		switch data[p] {
		case 32:
			goto tr1103
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 59:
			goto tr1105
		case 92:
			goto st56
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st486
			}
		case data[p] >= 9:
			goto tr1103
		}
		goto st44
	tr183:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
		goto st487
	tr189:
//line query/tokeniser.rl:87
		mark = p
		goto st487
	tr199:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
		goto st487
	tr205:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
		goto st487
	tr212:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st487
	tr219:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
		goto st487
	tr225:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
		goto st487
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
//line query/tokeniser.go:4720
		switch data[p] {
		case 32:
			goto tr1106
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 46:
			goto tr1107
		case 59:
			goto tr1109
		case 92:
			goto st56
		case 95:
			goto st487
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1106
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st487
				}
			case data[p] >= 65:
				goto st487
			}
		default:
			goto st487
		}
		goto st44
	tr1107:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line query/tokeniser.go:4764
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 95:
			goto st488
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st488
			}
		case data[p] >= 65:
			goto st488
		}
		goto st44
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
		switch data[p] {
		case 32:
			goto tr1110
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 46:
			goto st67
		case 59:
			goto tr1112
		case 92:
			goto st56
		case 95:
			goto st488
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1110
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st488
				}
			case data[p] >= 65:
				goto st488
			}
		default:
			goto st488
		}
		goto st44
	tr169:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st68
	tr174:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st68
	tr230:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line query/tokeniser.go:4856
		switch data[p] {
		case 32:
			goto tr193
		case 34:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr196
		case 61:
			goto st69
		case 92:
			goto st56
		case 95:
			goto tr199
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr193
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr199
				}
			case data[p] >= 65:
				goto tr199
			}
		default:
			goto tr197
		}
		goto st44
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 32:
			goto tr200
		case 34:
			goto tr201
		case 39:
			goto tr202
		case 43:
			goto tr203
		case 45:
			goto tr203
		case 92:
			goto st56
		case 95:
			goto tr205
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr200
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr205
				}
			case data[p] >= 65:
				goto tr205
			}
		default:
			goto tr204
		}
		goto st44
	tr170:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st70
	tr175:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st70
	tr231:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line query/tokeniser.go:4959
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 61:
			goto st71
		case 92:
			goto st56
		}
		goto st44
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 32:
			goto tr207
		case 34:
			goto tr208
		case 39:
			goto tr209
		case 43:
			goto tr210
		case 45:
			goto tr210
		case 92:
			goto st56
		case 95:
			goto tr212
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr207
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr212
				}
			case data[p] >= 65:
				goto tr212
			}
		default:
			goto tr211
		}
		goto st44
	tr171:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st72
	tr176:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st72
	tr232:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line query/tokeniser.go:5043
		switch data[p] {
		case 32:
			goto tr213
		case 34:
			goto tr214
		case 39:
			goto tr215
		case 43:
			goto tr216
		case 45:
			goto tr216
		case 61:
			goto st73
		case 92:
			goto st56
		case 95:
			goto tr219
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr213
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr219
				}
			case data[p] >= 65:
				goto tr219
			}
		default:
			goto tr217
		}
		goto st44
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 32:
			goto tr220
		case 34:
			goto tr221
		case 39:
			goto tr222
		case 43:
			goto tr223
		case 45:
			goto tr223
		case 92:
			goto st56
		case 95:
			goto tr225
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr220
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr225
				}
			case data[p] >= 65:
				goto tr225
			}
		default:
			goto tr224
		}
		goto st44
	tr167:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line query/tokeniser.go:5128
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 95:
			goto st75
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st75
			}
		case data[p] >= 65:
			goto st75
		}
		goto st44
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 32:
			goto tr227
		case 33:
			goto tr228
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 46:
			goto st74
		case 60:
			goto tr230
		case 61:
			goto tr231
		case 62:
			goto tr232
		case 92:
			goto st56
		case 95:
			goto st75
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr227
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st75
				}
			case data[p] >= 65:
				goto st75
			}
		default:
			goto st75
		}
		goto st44
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 32:
			goto tr233
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 95:
			goto tr234
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr233
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr234
			}
		default:
			goto tr234
		}
		goto st44
	tr233:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line query/tokeniser.go:5234
		switch data[p] {
		case 32:
			goto st77
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 95:
			goto st78
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st77
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st78
			}
		default:
			goto st78
		}
		goto st44
	tr234:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line query/tokeniser.go:5271
		switch data[p] {
		case 32:
			goto tr237
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 93:
			goto tr238
		case 95:
			goto st78
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr237
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st78
				}
			case data[p] >= 65:
				goto st78
			}
		default:
			goto st78
		}
		goto st44
	tr237:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line query/tokeniser.go:5313
		switch data[p] {
		case 32:
			goto st79
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 93:
			goto st489
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st79
		}
		goto st44
	tr238:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st489
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
//line query/tokeniser.go:5339
		switch data[p] {
		case 32:
			goto tr1113
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 59:
			goto tr1114
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1113
		}
		goto st44
	tr1093:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line query/tokeniser.go:5365
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 78:
			goto st81
		case 92:
			goto st56
		case 110:
			goto st81
		}
		goto st44
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 68:
			goto st58
		case 92:
			goto st56
		case 100:
			goto st58
		}
		goto st44
	tr1094:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line query/tokeniser.go:5406
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 82:
			goto st83
		case 92:
			goto st56
		case 114:
			goto st83
		}
		goto st44
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 32:
			goto tr243
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr243
		}
		goto st44
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 73:
			goto st85
		case 92:
			goto st56
		case 105:
			goto st85
		}
		goto st44
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 84:
			goto st86
		case 92:
			goto st56
		case 116:
			goto st86
		}
		goto st44
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 72:
			goto st87
		case 92:
			goto st56
		case 104:
			goto st87
		}
		goto st44
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 73:
			goto st88
		case 92:
			goto st56
		case 105:
			goto st88
		}
		goto st44
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 78:
			goto st89
		case 92:
			goto st56
		case 110:
			goto st89
		}
		goto st44
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto st90
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st90
		}
		goto st44
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 32:
			goto st90
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 43:
			goto tr250
		case 45:
			goto tr250
		case 92:
			goto st56
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr251
			}
		case data[p] >= 9:
			goto st90
		}
		goto st44
	tr250:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line query/tokeniser.go:5589
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st92
		}
		goto st44
	tr251:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line query/tokeniser.go:5615
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 46:
			goto st93
		case 72:
			goto st490
		case 77:
			goto st492
		case 78:
			goto st95
		case 83:
			goto st490
		case 85:
			goto st95
		case 92:
			goto st56
		case 104:
			goto st490
		case 109:
			goto st492
		case 110:
			goto st95
		case 115:
			goto st490
		case 117:
			goto st95
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st92
		}
		goto st44
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st94
		}
		goto st44
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 72:
			goto st490
		case 77:
			goto st492
		case 78:
			goto st95
		case 83:
			goto st490
		case 85:
			goto st95
		case 92:
			goto st56
		case 104:
			goto st490
		case 109:
			goto st492
		case 110:
			goto st95
		case 115:
			goto st490
		case 117:
			goto st95
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st94
		}
		goto st44
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		switch data[p] {
		case 32:
			goto tr1115
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 43:
			goto st91
		case 45:
			goto st91
		case 59:
			goto tr1117
		case 92:
			goto st56
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st92
			}
		case data[p] >= 9:
			goto tr1115
		}
		goto st44
	tr1115:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st491
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
//line query/tokeniser.go:5747
		switch data[p] {
		case 32:
			goto st491
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 59:
			goto st483
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st491
		}
		goto st44
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		switch data[p] {
		case 32:
			goto tr1115
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 43:
			goto st91
		case 45:
			goto st91
		case 59:
			goto tr1117
		case 83:
			goto st490
		case 92:
			goto st56
		case 115:
			goto st490
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st92
			}
		case data[p] >= 9:
			goto tr1115
		}
		goto st44
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 83:
			goto st490
		case 92:
			goto st56
		case 115:
			goto st490
		}
		goto st44
	tr1097:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line query/tokeniser.go:5825
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 124:
			goto st83
		}
		goto st44
	tr1098:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line query/tokeniser.go:5846
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 136:
			goto st98
		}
		goto st44
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 92:
			goto st56
		case 168:
			goto st83
		}
		goto st44
	tr159:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st493
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
//line query/tokeniser.go:5883
		switch data[p] {
		case 32:
			goto tr1119
		case 34:
			goto tr120
		case 39:
			goto tr121
		case 59:
			goto tr1120
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1119
		}
		goto st44
	tr147:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st494
	tr153:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st494
	tr264:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st494
	tr271:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st494
	tr278:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st494
	tr284:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st494
	tr291:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st494
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
//line query/tokeniser.go:5959
		switch data[p] {
		case 32:
			goto tr1121
		case 39:
			goto tr398
		case 59:
			goto tr1122
		case 92:
			goto tr399
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1121
		}
		goto tr397
	tr1072:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st495
	tr1076:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st495
	tr1122:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st495
	tr1125:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st495
	tr1129:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st495
	tr1132:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st495
	tr1134:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st495
	tr1137:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st495
	tr1282:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st495
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
//line query/tokeniser.go:6045
		switch data[p] {
		case 32:
			goto st495
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st495
		}
		goto st45
	tr148:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st99
	tr154:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st99
	tr265:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st99
	tr272:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st99
	tr279:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st99
	tr285:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st99
	tr292:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line query/tokeniser.go:6117
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st496
		}
		goto st45
	tr149:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st496
	tr155:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st496
	tr266:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st496
	tr273:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st496
	tr280:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st496
	tr286:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st496
	tr293:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st496
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
//line query/tokeniser.go:6187
		switch data[p] {
		case 32:
			goto tr1123
		case 39:
			goto tr124
		case 46:
			goto st100
		case 59:
			goto tr1125
		case 92:
			goto st46
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st496
			}
		case data[p] >= 9:
			goto tr1123
		}
		goto st45
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st497
		}
		goto st45
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 32:
			goto tr1123
		case 39:
			goto tr124
		case 59:
			goto tr1125
		case 92:
			goto st46
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st497
			}
		case data[p] >= 9:
			goto tr1123
		}
		goto st45
	tr150:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
		goto st498
	tr156:
//line query/tokeniser.rl:87
		mark = p
		goto st498
	tr268:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
		goto st498
	tr274:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
		goto st498
	tr281:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st498
	tr288:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
		goto st498
	tr294:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
		goto st498
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
//line query/tokeniser.go:6293
		switch data[p] {
		case 32:
			goto tr1126
		case 39:
			goto tr124
		case 46:
			goto tr1127
		case 59:
			goto tr1129
		case 92:
			goto st46
		case 95:
			goto st498
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1126
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st498
				}
			case data[p] >= 65:
				goto st498
			}
		default:
			goto st498
		}
		goto st45
	tr1127:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
//line query/tokeniser.go:6335
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		case 95:
			goto st499
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st499
			}
		case data[p] >= 65:
			goto st499
		}
		goto st45
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		switch data[p] {
		case 32:
			goto tr1130
		case 39:
			goto tr124
		case 46:
			goto st101
		case 59:
			goto tr1132
		case 92:
			goto st46
		case 95:
			goto st499
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1130
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st499
				}
			case data[p] >= 65:
				goto st499
			}
		default:
			goto st499
		}
		goto st45
	tr136:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st102
	tr141:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st102
	tr299:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line query/tokeniser.go:6423
		switch data[p] {
		case 32:
			goto tr262
		case 34:
			goto tr263
		case 39:
			goto tr264
		case 43:
			goto tr265
		case 45:
			goto tr265
		case 61:
			goto st103
		case 92:
			goto st46
		case 95:
			goto tr268
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr262
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr268
				}
			case data[p] >= 65:
				goto tr268
			}
		default:
			goto tr266
		}
		goto st45
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto tr269
		case 34:
			goto tr270
		case 39:
			goto tr271
		case 43:
			goto tr272
		case 45:
			goto tr272
		case 92:
			goto st46
		case 95:
			goto tr274
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr269
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr274
				}
			case data[p] >= 65:
				goto tr274
			}
		default:
			goto tr273
		}
		goto st45
	tr137:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st104
	tr142:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st104
	tr300:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line query/tokeniser.go:6526
		switch data[p] {
		case 39:
			goto tr124
		case 61:
			goto st105
		case 92:
			goto st46
		}
		goto st45
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 32:
			goto tr276
		case 34:
			goto tr277
		case 39:
			goto tr278
		case 43:
			goto tr279
		case 45:
			goto tr279
		case 92:
			goto st46
		case 95:
			goto tr281
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr276
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr281
				}
			case data[p] >= 65:
				goto tr281
			}
		default:
			goto tr280
		}
		goto st45
	tr138:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st106
	tr143:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st106
	tr301:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line query/tokeniser.go:6608
		switch data[p] {
		case 32:
			goto tr282
		case 34:
			goto tr283
		case 39:
			goto tr284
		case 43:
			goto tr285
		case 45:
			goto tr285
		case 61:
			goto st107
		case 92:
			goto st46
		case 95:
			goto tr288
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr282
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr288
				}
			case data[p] >= 65:
				goto tr288
			}
		default:
			goto tr286
		}
		goto st45
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 32:
			goto tr289
		case 34:
			goto tr290
		case 39:
			goto tr291
		case 43:
			goto tr292
		case 45:
			goto tr292
		case 92:
			goto st46
		case 95:
			goto tr294
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr289
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr294
				}
			case data[p] >= 65:
				goto tr294
			}
		default:
			goto tr293
		}
		goto st45
	tr134:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line query/tokeniser.go:6693
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		case 95:
			goto st109
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st109
			}
		case data[p] >= 65:
			goto st109
		}
		goto st45
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 32:
			goto tr296
		case 33:
			goto tr297
		case 39:
			goto tr124
		case 46:
			goto st108
		case 60:
			goto tr299
		case 61:
			goto tr300
		case 62:
			goto tr301
		case 92:
			goto st46
		case 95:
			goto st109
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr296
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st109
				}
			case data[p] >= 65:
				goto st109
			}
		default:
			goto st109
		}
		goto st45
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 32:
			goto tr302
		case 39:
			goto tr124
		case 92:
			goto st46
		case 95:
			goto tr303
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr302
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr303
			}
		default:
			goto tr303
		}
		goto st45
	tr302:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line query/tokeniser.go:6793
		switch data[p] {
		case 32:
			goto st111
		case 39:
			goto tr124
		case 92:
			goto st46
		case 95:
			goto st112
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st112
			}
		default:
			goto st112
		}
		goto st45
	tr303:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line query/tokeniser.go:6828
		switch data[p] {
		case 32:
			goto tr306
		case 39:
			goto tr124
		case 92:
			goto st46
		case 93:
			goto tr307
		case 95:
			goto st112
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr306
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st112
				}
			case data[p] >= 65:
				goto st112
			}
		default:
			goto st112
		}
		goto st45
	tr306:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line query/tokeniser.go:6868
		switch data[p] {
		case 32:
			goto st113
		case 39:
			goto tr124
		case 92:
			goto st46
		case 93:
			goto st500
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st113
		}
		goto st45
	tr307:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st500
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
//line query/tokeniser.go:6892
		switch data[p] {
		case 32:
			goto tr1133
		case 39:
			goto tr124
		case 59:
			goto tr1134
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1133
		}
		goto st45
	tr1080:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line query/tokeniser.go:6916
		switch data[p] {
		case 39:
			goto tr124
		case 78:
			goto st115
		case 92:
			goto st46
		case 110:
			goto st115
		}
		goto st45
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 39:
			goto tr124
		case 68:
			goto st48
		case 92:
			goto st46
		case 100:
			goto st48
		}
		goto st45
	tr1081:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line query/tokeniser.go:6953
		switch data[p] {
		case 39:
			goto tr124
		case 82:
			goto st117
		case 92:
			goto st46
		case 114:
			goto st117
		}
		goto st45
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 32:
			goto tr312
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr312
		}
		goto st45
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 39:
			goto tr124
		case 73:
			goto st119
		case 92:
			goto st46
		case 105:
			goto st119
		}
		goto st45
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 39:
			goto tr124
		case 84:
			goto st120
		case 92:
			goto st46
		case 116:
			goto st120
		}
		goto st45
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 39:
			goto tr124
		case 72:
			goto st121
		case 92:
			goto st46
		case 104:
			goto st121
		}
		goto st45
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 39:
			goto tr124
		case 73:
			goto st122
		case 92:
			goto st46
		case 105:
			goto st122
		}
		goto st45
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 39:
			goto tr124
		case 78:
			goto st123
		case 92:
			goto st46
		case 110:
			goto st123
		}
		goto st45
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 32:
			goto st124
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st124
		}
		goto st45
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 32:
			goto st124
		case 39:
			goto tr124
		case 43:
			goto tr319
		case 45:
			goto tr319
		case 92:
			goto st46
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr320
			}
		case data[p] >= 9:
			goto st124
		}
		goto st45
	tr319:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
//line query/tokeniser.go:7118
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st126
		}
		goto st45
	tr320:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line query/tokeniser.go:7142
		switch data[p] {
		case 39:
			goto tr124
		case 46:
			goto st127
		case 72:
			goto st501
		case 77:
			goto st503
		case 78:
			goto st129
		case 83:
			goto st501
		case 85:
			goto st129
		case 92:
			goto st46
		case 104:
			goto st501
		case 109:
			goto st503
		case 110:
			goto st129
		case 115:
			goto st501
		case 117:
			goto st129
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st126
		}
		goto st45
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st128
		}
		goto st45
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 39:
			goto tr124
		case 72:
			goto st501
		case 77:
			goto st503
		case 78:
			goto st129
		case 83:
			goto st501
		case 85:
			goto st129
		case 92:
			goto st46
		case 104:
			goto st501
		case 109:
			goto st503
		case 110:
			goto st129
		case 115:
			goto st501
		case 117:
			goto st129
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st128
		}
		goto st45
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		switch data[p] {
		case 32:
			goto tr1135
		case 39:
			goto tr124
		case 43:
			goto st125
		case 45:
			goto st125
		case 59:
			goto tr1137
		case 92:
			goto st46
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st126
			}
		case data[p] >= 9:
			goto tr1135
		}
		goto st45
	tr1135:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st502
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
//line query/tokeniser.go:7266
		switch data[p] {
		case 32:
			goto st502
		case 39:
			goto tr124
		case 59:
			goto st495
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st502
		}
		goto st45
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 32:
			goto tr1135
		case 39:
			goto tr124
		case 43:
			goto st125
		case 45:
			goto st125
		case 59:
			goto tr1137
		case 83:
			goto st501
		case 92:
			goto st46
		case 115:
			goto st501
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st126
			}
		case data[p] >= 9:
			goto tr1135
		}
		goto st45
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 39:
			goto tr124
		case 83:
			goto st501
		case 92:
			goto st46
		case 115:
			goto st501
		}
		goto st45
	tr1084:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st130
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
//line query/tokeniser.go:7338
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		case 124:
			goto st117
		}
		goto st45
	tr1085:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line query/tokeniser.go:7357
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		case 136:
			goto st132
		}
		goto st45
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		case 168:
			goto st117
		}
		goto st45
	tr106:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st133
	tr112:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st133
	tr334:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st133
	tr341:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st133
	tr348:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st133
	tr354:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st133
	tr361:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line query/tokeniser.go:7440
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st504
		}
		goto st33
	tr107:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st504
	tr113:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st504
	tr335:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st504
	tr342:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st504
	tr349:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st504
	tr355:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st504
	tr362:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st504
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
//line query/tokeniser.go:7510
		switch data[p] {
		case 32:
			goto tr1139
		case 34:
			goto tr55
		case 46:
			goto st134
		case 59:
			goto tr1141
		case 92:
			goto st34
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st504
			}
		case data[p] >= 9:
			goto tr1139
		}
		goto st33
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st505
		}
		goto st33
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		switch data[p] {
		case 32:
			goto tr1139
		case 34:
			goto tr55
		case 59:
			goto tr1141
		case 92:
			goto st34
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st505
			}
		case data[p] >= 9:
			goto tr1139
		}
		goto st33
	tr108:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
		goto st506
	tr114:
//line query/tokeniser.rl:87
		mark = p
		goto st506
	tr337:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
		goto st506
	tr343:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
		goto st506
	tr350:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st506
	tr357:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
		goto st506
	tr363:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
		goto st506
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
//line query/tokeniser.go:7616
		switch data[p] {
		case 32:
			goto tr1142
		case 34:
			goto tr55
		case 46:
			goto tr1143
		case 59:
			goto tr1145
		case 92:
			goto st34
		case 95:
			goto st506
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1142
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st506
				}
			case data[p] >= 65:
				goto st506
			}
		default:
			goto st506
		}
		goto st33
	tr1143:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
//line query/tokeniser.go:7658
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		case 95:
			goto st507
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st507
			}
		case data[p] >= 65:
			goto st507
		}
		goto st33
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		switch data[p] {
		case 32:
			goto tr1146
		case 34:
			goto tr55
		case 46:
			goto st135
		case 59:
			goto tr1148
		case 92:
			goto st34
		case 95:
			goto st507
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1146
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st507
				}
			case data[p] >= 65:
				goto st507
			}
		default:
			goto st507
		}
		goto st33
	tr94:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st136
	tr99:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st136
	tr368:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line query/tokeniser.go:7746
		switch data[p] {
		case 32:
			goto tr331
		case 34:
			goto tr332
		case 39:
			goto tr333
		case 43:
			goto tr334
		case 45:
			goto tr334
		case 61:
			goto st137
		case 92:
			goto st34
		case 95:
			goto tr337
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr331
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr337
				}
			case data[p] >= 65:
				goto tr337
			}
		default:
			goto tr335
		}
		goto st33
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 32:
			goto tr338
		case 34:
			goto tr339
		case 39:
			goto tr340
		case 43:
			goto tr341
		case 45:
			goto tr341
		case 92:
			goto st34
		case 95:
			goto tr343
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr338
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr343
				}
			case data[p] >= 65:
				goto tr343
			}
		default:
			goto tr342
		}
		goto st33
	tr95:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st138
	tr100:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st138
	tr369:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line query/tokeniser.go:7849
		switch data[p] {
		case 34:
			goto tr55
		case 61:
			goto st139
		case 92:
			goto st34
		}
		goto st33
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 32:
			goto tr345
		case 34:
			goto tr346
		case 39:
			goto tr347
		case 43:
			goto tr348
		case 45:
			goto tr348
		case 92:
			goto st34
		case 95:
			goto tr350
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr345
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr350
				}
			case data[p] >= 65:
				goto tr350
			}
		default:
			goto tr349
		}
		goto st33
	tr96:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st140
	tr101:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st140
	tr370:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
//line query/tokeniser.go:7931
		switch data[p] {
		case 32:
			goto tr351
		case 34:
			goto tr352
		case 39:
			goto tr353
		case 43:
			goto tr354
		case 45:
			goto tr354
		case 61:
			goto st141
		case 92:
			goto st34
		case 95:
			goto tr357
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr351
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr357
				}
			case data[p] >= 65:
				goto tr357
			}
		default:
			goto tr355
		}
		goto st33
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 32:
			goto tr358
		case 34:
			goto tr359
		case 39:
			goto tr360
		case 43:
			goto tr361
		case 45:
			goto tr361
		case 92:
			goto st34
		case 95:
			goto tr363
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr358
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr363
				}
			case data[p] >= 65:
				goto tr363
			}
		default:
			goto tr362
		}
		goto st33
	tr92:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
//line query/tokeniser.go:8016
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		case 95:
			goto st143
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st143
			}
		case data[p] >= 65:
			goto st143
		}
		goto st33
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 32:
			goto tr365
		case 33:
			goto tr366
		case 34:
			goto tr55
		case 46:
			goto st142
		case 60:
			goto tr368
		case 61:
			goto tr369
		case 62:
			goto tr370
		case 92:
			goto st34
		case 95:
			goto st143
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr365
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st143
				}
			case data[p] >= 65:
				goto st143
			}
		default:
			goto st143
		}
		goto st33
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 32:
			goto tr371
		case 34:
			goto tr55
		case 92:
			goto st34
		case 95:
			goto tr372
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr371
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr372
			}
		default:
			goto tr372
		}
		goto st33
	tr371:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
//line query/tokeniser.go:8116
		switch data[p] {
		case 32:
			goto st145
		case 34:
			goto tr55
		case 92:
			goto st34
		case 95:
			goto st146
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st145
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st146
			}
		default:
			goto st146
		}
		goto st33
	tr372:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st146
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
//line query/tokeniser.go:8151
		switch data[p] {
		case 32:
			goto tr375
		case 34:
			goto tr55
		case 92:
			goto st34
		case 93:
			goto tr376
		case 95:
			goto st146
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr375
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st146
				}
			case data[p] >= 65:
				goto st146
			}
		default:
			goto st146
		}
		goto st33
	tr375:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line query/tokeniser.go:8191
		switch data[p] {
		case 32:
			goto st147
		case 34:
			goto tr55
		case 92:
			goto st34
		case 93:
			goto st508
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st147
		}
		goto st33
	tr376:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st508
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
//line query/tokeniser.go:8215
		switch data[p] {
		case 32:
			goto tr1149
		case 34:
			goto tr55
		case 59:
			goto tr1150
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1149
		}
		goto st33
	tr1063:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line query/tokeniser.go:8239
		switch data[p] {
		case 34:
			goto tr55
		case 78:
			goto st149
		case 92:
			goto st34
		case 110:
			goto st149
		}
		goto st33
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 34:
			goto tr55
		case 68:
			goto st36
		case 92:
			goto st34
		case 100:
			goto st36
		}
		goto st33
	tr1064:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line query/tokeniser.go:8276
		switch data[p] {
		case 34:
			goto tr55
		case 82:
			goto st151
		case 92:
			goto st34
		case 114:
			goto st151
		}
		goto st33
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		switch data[p] {
		case 32:
			goto tr381
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr381
		}
		goto st33
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 34:
			goto tr55
		case 73:
			goto st153
		case 92:
			goto st34
		case 105:
			goto st153
		}
		goto st33
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 34:
			goto tr55
		case 84:
			goto st154
		case 92:
			goto st34
		case 116:
			goto st154
		}
		goto st33
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 34:
			goto tr55
		case 72:
			goto st155
		case 92:
			goto st34
		case 104:
			goto st155
		}
		goto st33
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 34:
			goto tr55
		case 73:
			goto st156
		case 92:
			goto st34
		case 105:
			goto st156
		}
		goto st33
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 34:
			goto tr55
		case 78:
			goto st157
		case 92:
			goto st34
		case 110:
			goto st157
		}
		goto st33
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 32:
			goto st158
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st158
		}
		goto st33
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 32:
			goto st158
		case 34:
			goto tr55
		case 43:
			goto tr388
		case 45:
			goto tr388
		case 92:
			goto st34
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr389
			}
		case data[p] >= 9:
			goto st158
		}
		goto st33
	tr388:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st159
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
//line query/tokeniser.go:8441
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st160
		}
		goto st33
	tr389:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st160
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
//line query/tokeniser.go:8465
		switch data[p] {
		case 34:
			goto tr55
		case 46:
			goto st161
		case 72:
			goto st509
		case 77:
			goto st511
		case 78:
			goto st163
		case 83:
			goto st509
		case 85:
			goto st163
		case 92:
			goto st34
		case 104:
			goto st509
		case 109:
			goto st511
		case 110:
			goto st163
		case 115:
			goto st509
		case 117:
			goto st163
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st160
		}
		goto st33
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st162
		}
		goto st33
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 34:
			goto tr55
		case 72:
			goto st509
		case 77:
			goto st511
		case 78:
			goto st163
		case 83:
			goto st509
		case 85:
			goto st163
		case 92:
			goto st34
		case 104:
			goto st509
		case 109:
			goto st511
		case 110:
			goto st163
		case 115:
			goto st509
		case 117:
			goto st163
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st162
		}
		goto st33
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		switch data[p] {
		case 32:
			goto tr1151
		case 34:
			goto tr55
		case 43:
			goto st159
		case 45:
			goto st159
		case 59:
			goto tr1153
		case 92:
			goto st34
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st160
			}
		case data[p] >= 9:
			goto tr1151
		}
		goto st33
	tr1151:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st510
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
//line query/tokeniser.go:8589
		switch data[p] {
		case 32:
			goto st510
		case 34:
			goto tr55
		case 59:
			goto st474
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st510
		}
		goto st33
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		switch data[p] {
		case 32:
			goto tr1151
		case 34:
			goto tr55
		case 43:
			goto st159
		case 45:
			goto st159
		case 59:
			goto tr1153
		case 83:
			goto st509
		case 92:
			goto st34
		case 115:
			goto st509
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st160
			}
		case data[p] >= 9:
			goto tr1151
		}
		goto st33
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 34:
			goto tr55
		case 83:
			goto st509
		case 92:
			goto st34
		case 115:
			goto st509
		}
		goto st33
	tr1067:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st164
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
//line query/tokeniser.go:8661
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		case 124:
			goto st151
		}
		goto st33
	tr1068:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st165
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
//line query/tokeniser.go:8680
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		case 136:
			goto st166
		}
		goto st33
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st34
		case 168:
			goto st151
		}
		goto st33
	tr76:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st167
	tr79:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st167
	tr405:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st167
	tr412:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st167
	tr419:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st167
	tr425:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st167
	tr432:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st167
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
//line query/tokeniser.go:8749
		switch data[p] {
		case 39:
			goto tr398
		case 92:
			goto tr399
		}
		goto tr397
	tr42:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st168
	tr48:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st168
	tr406:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st168
	tr413:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st168
	tr420:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st168
	tr426:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st168
	tr433:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st168
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
//line query/tokeniser.go:8816
		if 48 <= data[p] && data[p] <= 57 {
			goto st512
		}
		goto st0
	tr43:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st512
	tr49:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st512
	tr407:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st512
	tr414:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st512
	tr421:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st512
	tr427:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st512
	tr434:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st512
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
//line query/tokeniser.go:8880
		switch data[p] {
		case 32:
			goto tr1155
		case 46:
			goto st169
		case 59:
			goto tr1157
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st512
			}
		case data[p] >= 9:
			goto tr1155
		}
		goto st0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		if 48 <= data[p] && data[p] <= 57 {
			goto st513
		}
		goto st0
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		switch data[p] {
		case 32:
			goto tr1155
		case 59:
			goto tr1157
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st513
			}
		case data[p] >= 9:
			goto tr1155
		}
		goto st0
	tr44:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
		goto st514
	tr50:
//line query/tokeniser.rl:87
		mark = p
		goto st514
	tr409:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
		goto st514
	tr415:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
		goto st514
	tr422:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st514
	tr429:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
		goto st514
	tr435:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
		goto st514
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
//line query/tokeniser.go:8972
		switch data[p] {
		case 32:
			goto tr1158
		case 46:
			goto tr1159
		case 59:
			goto tr1161
		case 95:
			goto st514
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1158
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st514
				}
			case data[p] >= 65:
				goto st514
			}
		default:
			goto st514
		}
		goto st0
	tr1159:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st170
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
//line query/tokeniser.go:9010
		if data[p] == 95 {
			goto st515
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st515
			}
		case data[p] >= 65:
			goto st515
		}
		goto st0
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		switch data[p] {
		case 32:
			goto tr1162
		case 46:
			goto st170
		case 59:
			goto tr1164
		case 95:
			goto st515
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1162
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st515
				}
			case data[p] >= 65:
				goto st515
			}
		default:
			goto st515
		}
		goto st0
	tr65:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st171
	tr70:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st171
	tr440:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st171
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
//line query/tokeniser.go:9089
		switch data[p] {
		case 32:
			goto tr403
		case 34:
			goto tr404
		case 39:
			goto tr405
		case 43:
			goto tr406
		case 45:
			goto tr406
		case 61:
			goto st172
		case 95:
			goto tr409
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr403
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr409
				}
			case data[p] >= 65:
				goto tr409
			}
		default:
			goto tr407
		}
		goto st0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 32:
			goto tr410
		case 34:
			goto tr411
		case 39:
			goto tr412
		case 43:
			goto tr413
		case 45:
			goto tr413
		case 95:
			goto tr415
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr410
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr415
				}
			case data[p] >= 65:
				goto tr415
			}
		default:
			goto tr414
		}
		goto st0
	tr66:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st173
	tr71:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st173
	tr441:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st173
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
//line query/tokeniser.go:9188
		if data[p] == 61 {
			goto st174
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 32:
			goto tr417
		case 34:
			goto tr418
		case 39:
			goto tr419
		case 43:
			goto tr420
		case 45:
			goto tr420
		case 95:
			goto tr422
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr417
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr422
				}
			case data[p] >= 65:
				goto tr422
			}
		default:
			goto tr421
		}
		goto st0
	tr67:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st175
	tr72:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st175
	tr442:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st175
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
//line query/tokeniser.go:9263
		switch data[p] {
		case 32:
			goto tr423
		case 34:
			goto tr424
		case 39:
			goto tr425
		case 43:
			goto tr426
		case 45:
			goto tr426
		case 61:
			goto st176
		case 95:
			goto tr429
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr423
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr429
				}
			case data[p] >= 65:
				goto tr429
			}
		default:
			goto tr427
		}
		goto st0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 32:
			goto tr430
		case 34:
			goto tr431
		case 39:
			goto tr432
		case 43:
			goto tr433
		case 45:
			goto tr433
		case 95:
			goto tr435
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr430
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr435
				}
			case data[p] >= 65:
				goto tr435
			}
		default:
			goto tr434
		}
		goto st0
	tr63:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st177
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
//line query/tokeniser.go:9344
		if data[p] == 95 {
			goto st178
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st178
			}
		case data[p] >= 65:
			goto st178
		}
		goto st0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 32:
			goto tr437
		case 33:
			goto tr438
		case 46:
			goto st177
		case 60:
			goto tr440
		case 61:
			goto tr441
		case 62:
			goto tr442
		case 95:
			goto st178
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr437
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st178
				}
			case data[p] >= 65:
				goto st178
			}
		default:
			goto st178
		}
		goto st0
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		switch data[p] {
		case 32:
			goto tr443
		case 95:
			goto tr444
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr443
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr444
			}
		default:
			goto tr444
		}
		goto st0
	tr443:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st180
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
//line query/tokeniser.go:9431
		switch data[p] {
		case 32:
			goto st180
		case 95:
			goto st181
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st180
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st181
			}
		default:
			goto st181
		}
		goto st0
	tr444:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
//line query/tokeniser.go:9462
		switch data[p] {
		case 32:
			goto tr447
		case 93:
			goto tr448
		case 95:
			goto st181
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr447
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st181
				}
			case data[p] >= 65:
				goto st181
			}
		default:
			goto st181
		}
		goto st0
	tr447:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st182
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
//line query/tokeniser.go:9498
		switch data[p] {
		case 32:
			goto st182
		case 93:
			goto st516
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st182
		}
		goto st0
	tr448:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st516
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
//line query/tokeniser.go:9518
		switch data[p] {
		case 32:
			goto tr1165
		case 59:
			goto tr1166
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1165
		}
		goto st0
	tr1052:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st183
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
//line query/tokeniser.go:9538
		switch data[p] {
		case 78:
			goto st184
		case 110:
			goto st184
		}
		goto st0
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		switch data[p] {
		case 68:
			goto st25
		case 100:
			goto st25
		}
		goto st0
	tr1053:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st185
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
//line query/tokeniser.go:9567
		switch data[p] {
		case 82:
			goto st186
		case 114:
			goto st186
		}
		goto st0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		if data[p] == 32 {
			goto tr453
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr453
		}
		goto st0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		switch data[p] {
		case 73:
			goto st188
		case 105:
			goto st188
		}
		goto st0
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 84:
			goto st189
		case 116:
			goto st189
		}
		goto st0
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 72:
			goto st190
		case 104:
			goto st190
		}
		goto st0
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch data[p] {
		case 73:
			goto st191
		case 105:
			goto st191
		}
		goto st0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 78:
			goto st192
		case 110:
			goto st192
		}
		goto st0
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		if data[p] == 32 {
			goto st193
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st193
		}
		goto st0
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch data[p] {
		case 32:
			goto st193
		case 43:
			goto tr459
		case 45:
			goto tr459
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr460
			}
		case data[p] >= 9:
			goto st193
		}
		goto st0
	tr459:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line query/tokeniser.go:9694
		if 48 <= data[p] && data[p] <= 57 {
			goto st195
		}
		goto st0
	tr460:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
//line query/tokeniser.go:9712
		switch data[p] {
		case 46:
			goto st196
		case 72:
			goto st517
		case 77:
			goto st519
		case 78:
			goto st198
		case 83:
			goto st517
		case 85:
			goto st198
		case 104:
			goto st517
		case 109:
			goto st519
		case 110:
			goto st198
		case 115:
			goto st517
		case 117:
			goto st198
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st195
		}
		goto st0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		if 48 <= data[p] && data[p] <= 57 {
			goto st197
		}
		goto st0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 72:
			goto st517
		case 77:
			goto st519
		case 78:
			goto st198
		case 83:
			goto st517
		case 85:
			goto st198
		case 104:
			goto st517
		case 109:
			goto st519
		case 110:
			goto st198
		case 115:
			goto st517
		case 117:
			goto st198
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st197
		}
		goto st0
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		switch data[p] {
		case 32:
			goto tr1167
		case 43:
			goto st194
		case 45:
			goto st194
		case 59:
			goto tr1169
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st195
			}
		case data[p] >= 9:
			goto tr1167
		}
		goto st0
	tr1167:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st518
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
//line query/tokeniser.go:9818
		switch data[p] {
		case 32:
			goto st518
		case 59:
			goto st468
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st518
		}
		goto st0
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		switch data[p] {
		case 32:
			goto tr1167
		case 43:
			goto st194
		case 45:
			goto st194
		case 59:
			goto tr1169
		case 83:
			goto st517
		case 115:
			goto st517
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st195
			}
		case data[p] >= 9:
			goto tr1167
		}
		goto st0
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 83:
			goto st517
		case 115:
			goto st517
		}
		goto st0
	tr1056:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line query/tokeniser.go:9878
		if data[p] == 124 {
			goto st186
		}
		goto st0
	tr1057:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st200
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
//line query/tokeniser.go:9892
		if data[p] == 136 {
			goto st201
		}
		goto st0
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		if data[p] == 168 {
			goto st186
		}
		goto st0
	tr53:
//line query/tokeniser.rl:87
		mark = p
		goto st202
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
//line query/tokeniser.go:9915
		switch data[p] {
		case 34:
			goto tr468
		case 92:
			goto st202
		}
		goto st23
	tr468:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st520
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
//line query/tokeniser.go:9932
		switch data[p] {
		case 32:
			goto tr1171
		case 34:
			goto tr55
		case 59:
			goto tr1172
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1171
		}
		goto st23
	tr1171:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st521
	tr1184:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st521
	tr1219:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st521
	tr1222:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st521
	tr1226:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st521
	tr1229:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st521
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
//line query/tokeniser.go:9994
		switch data[p] {
		case 32:
			goto st521
		case 34:
			goto tr55
		case 38:
			goto tr1174
		case 59:
			goto st524
		case 65:
			goto tr1176
		case 79:
			goto tr1177
		case 87:
			goto st275
		case 92:
			goto st202
		case 94:
			goto tr1179
		case 97:
			goto tr1176
		case 111:
			goto tr1177
		case 119:
			goto st275
		case 124:
			goto tr1180
		case 226:
			goto tr1181
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st521
		}
		goto st23
	tr1174:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st203
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
//line query/tokeniser.go:10038
		switch data[p] {
		case 34:
			goto tr55
		case 38:
			goto st204
		case 92:
			goto st202
		}
		goto st23
	tr1179:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st204
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
//line query/tokeniser.go:10057
		switch data[p] {
		case 32:
			goto tr470
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr470
		}
		goto st23
	tr470:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st205
	tr660:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
//line query/tokeniser.go:10083
		switch data[p] {
		case 32:
			goto st205
		case 34:
			goto tr55
		case 91:
			goto st267
		case 92:
			goto st202
		case 95:
			goto tr472
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st205
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr472
			}
		default:
			goto tr472
		}
		goto st23
	tr472:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line query/tokeniser.go:10120
		switch data[p] {
		case 32:
			goto tr474
		case 33:
			goto tr475
		case 34:
			goto tr55
		case 46:
			goto tr476
		case 60:
			goto tr478
		case 61:
			goto tr479
		case 62:
			goto tr480
		case 92:
			goto st202
		case 95:
			goto st206
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr474
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st206
				}
			case data[p] >= 65:
				goto st206
			}
		default:
			goto st206
		}
		goto st23
	tr474:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st207
	tr644:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st207
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
//line query/tokeniser.go:10178
		switch data[p] {
		case 32:
			goto st207
		case 33:
			goto tr482
		case 34:
			goto tr55
		case 60:
			goto tr483
		case 61:
			goto tr484
		case 62:
			goto tr485
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st207
		}
		goto st23
	tr475:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st208
	tr482:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st208
	tr645:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
//line query/tokeniser.go:10226
		switch data[p] {
		case 34:
			goto tr55
		case 61:
			goto st209
		case 92:
			goto st202
		}
		goto st23
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		switch data[p] {
		case 32:
			goto tr487
		case 34:
			goto tr488
		case 39:
			goto tr489
		case 43:
			goto tr490
		case 45:
			goto tr490
		case 92:
			goto st202
		case 95:
			goto tr492
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr487
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr492
				}
			case data[p] >= 65:
				goto tr492
			}
		default:
			goto tr491
		}
		goto st23
	tr487:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st210
	tr610:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st210
	tr617:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st210
	tr624:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st210
	tr630:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st210
	tr637:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
//line query/tokeniser.go:10304
		switch data[p] {
		case 32:
			goto st210
		case 34:
			goto tr494
		case 39:
			goto tr495
		case 43:
			goto tr496
		case 45:
			goto tr496
		case 92:
			goto st202
		case 95:
			goto tr498
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st210
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr498
				}
			case data[p] >= 65:
				goto tr498
			}
		default:
			goto tr497
		}
		goto st23
	tr488:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st522
	tr494:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st522
	tr611:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st522
	tr618:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st522
	tr625:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st522
	tr631:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st522
	tr638:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st522
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
//line query/tokeniser.go:10398
		switch data[p] {
		case 32:
			goto tr1182
		case 34:
			goto tr52
		case 59:
			goto tr1183
		case 92:
			goto tr81
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1182
		}
		goto tr80
	tr489:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st211
	tr495:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st211
	tr612:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st211
	tr619:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st211
	tr626:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st211
	tr632:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st211
	tr639:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st211
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
//line query/tokeniser.go:10458
		switch data[p] {
		case 34:
			goto tr500
		case 39:
			goto tr501
		case 92:
			goto tr502
		}
		goto tr499
	tr499:
//line query/tokeniser.rl:87
		mark = p
		goto st212
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
//line query/tokeniser.go:10477
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		goto st212
	tr501:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st523
	tr504:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st523
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
//line query/tokeniser.go:10502
		switch data[p] {
		case 32:
			goto tr1184
		case 34:
			goto tr55
		case 59:
			goto tr1185
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1184
		}
		goto st23
	tr1172:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st524
	tr1185:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st524
	tr1221:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st524
	tr1225:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st524
	tr1228:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st524
	tr1230:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st524
	tr1233:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st524
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
//line query/tokeniser.go:10572
		switch data[p] {
		case 32:
			goto st524
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st524
		}
		goto st23
	tr502:
//line query/tokeniser.rl:87
		mark = p
		goto st213
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
//line query/tokeniser.go:10594
		switch data[p] {
		case 34:
			goto tr506
		case 39:
			goto tr507
		case 92:
			goto st213
		}
		goto st212
	tr506:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st525
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
//line query/tokeniser.go:10613
		switch data[p] {
		case 32:
			goto tr1186
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 59:
			goto tr1187
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1186
		}
		goto st212
	tr1186:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st526
	tr1217:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st526
	tr1199:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st526
	tr1201:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st526
	tr1204:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st526
	tr1208:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st526
	tr1211:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st526
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
//line query/tokeniser.go:10685
		switch data[p] {
		case 32:
			goto st526
		case 34:
			goto tr120
		case 38:
			goto tr1189
		case 39:
			goto tr504
		case 59:
			goto st529
		case 65:
			goto tr1191
		case 79:
			goto tr1192
		case 87:
			goto st241
		case 92:
			goto st213
		case 94:
			goto tr1194
		case 97:
			goto tr1191
		case 111:
			goto tr1192
		case 119:
			goto st241
		case 124:
			goto tr1195
		case 226:
			goto tr1196
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st526
		}
		goto st212
	tr1189:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st214
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
//line query/tokeniser.go:10731
		switch data[p] {
		case 34:
			goto tr120
		case 38:
			goto st215
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		goto st212
	tr1194:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st215
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
//line query/tokeniser.go:10752
		switch data[p] {
		case 32:
			goto tr509
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr509
		}
		goto st212
	tr509:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st216
	tr591:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st216
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
//line query/tokeniser.go:10780
		switch data[p] {
		case 32:
			goto st216
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 91:
			goto st233
		case 92:
			goto st213
		case 95:
			goto tr511
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st216
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr511
			}
		default:
			goto tr511
		}
		goto st212
	tr511:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
//line query/tokeniser.go:10819
		switch data[p] {
		case 32:
			goto tr513
		case 33:
			goto tr514
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 46:
			goto tr515
		case 60:
			goto tr517
		case 61:
			goto tr518
		case 62:
			goto tr519
		case 92:
			goto st213
		case 95:
			goto st217
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr513
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st217
				}
			case data[p] >= 65:
				goto st217
			}
		default:
			goto st217
		}
		goto st212
	tr513:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st218
	tr575:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st218
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
//line query/tokeniser.go:10879
		switch data[p] {
		case 32:
			goto st218
		case 33:
			goto tr521
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 60:
			goto tr522
		case 61:
			goto tr523
		case 62:
			goto tr524
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st218
		}
		goto st212
	tr514:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st219
	tr521:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st219
	tr576:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
//line query/tokeniser.go:10929
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 61:
			goto st220
		case 92:
			goto st213
		}
		goto st212
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 32:
			goto tr526
		case 34:
			goto tr527
		case 39:
			goto tr528
		case 43:
			goto tr529
		case 45:
			goto tr529
		case 92:
			goto st213
		case 95:
			goto tr531
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr526
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr531
				}
			case data[p] >= 65:
				goto tr531
			}
		default:
			goto tr530
		}
		goto st212
	tr526:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st221
	tr541:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st221
	tr548:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st221
	tr555:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st221
	tr561:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st221
	tr568:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
//line query/tokeniser.go:11009
		switch data[p] {
		case 32:
			goto st221
		case 34:
			goto tr533
		case 39:
			goto tr534
		case 43:
			goto tr535
		case 45:
			goto tr535
		case 92:
			goto st213
		case 95:
			goto tr537
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st221
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr537
				}
			case data[p] >= 65:
				goto tr537
			}
		default:
			goto tr536
		}
		goto st212
	tr527:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st527
	tr533:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st527
	tr542:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st527
	tr549:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st527
	tr556:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st527
	tr562:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st527
	tr569:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st527
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
//line query/tokeniser.go:11103
		switch data[p] {
		case 32:
			goto tr1197
		case 34:
			goto tr116
		case 39:
			goto tr157
		case 59:
			goto tr1198
		case 92:
			goto tr118
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1197
		}
		goto tr115
	tr528:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st528
	tr534:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st528
	tr543:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st528
	tr550:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st528
	tr557:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st528
	tr563:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st528
	tr570:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st528
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
//line query/tokeniser.go:11179
		switch data[p] {
		case 32:
			goto tr1199
		case 34:
			goto tr500
		case 39:
			goto tr501
		case 59:
			goto tr1200
		case 92:
			goto tr502
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1199
		}
		goto tr499
	tr1187:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st529
	tr1218:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st529
	tr1200:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st529
	tr1203:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st529
	tr1207:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st529
	tr1210:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st529
	tr1212:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st529
	tr1215:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st529
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
//line query/tokeniser.go:11259
		switch data[p] {
		case 32:
			goto st529
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st529
		}
		goto st212
	tr529:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st222
	tr535:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st222
	tr544:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st222
	tr551:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st222
	tr558:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st222
	tr564:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st222
	tr571:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st222
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
//line query/tokeniser.go:11333
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st530
		}
		goto st212
	tr530:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st530
	tr536:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st530
	tr545:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st530
	tr552:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st530
	tr559:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st530
	tr565:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st530
	tr572:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st530
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
//line query/tokeniser.go:11405
		switch data[p] {
		case 32:
			goto tr1201
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 46:
			goto st223
		case 59:
			goto tr1203
		case 92:
			goto st213
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st530
			}
		case data[p] >= 9:
			goto tr1201
		}
		goto st212
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st531
		}
		goto st212
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		switch data[p] {
		case 32:
			goto tr1201
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 59:
			goto tr1203
		case 92:
			goto st213
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st531
			}
		case data[p] >= 9:
			goto tr1201
		}
		goto st212
	tr531:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
		goto st532
	tr537:
//line query/tokeniser.rl:87
		mark = p
		goto st532
	tr547:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
		goto st532
	tr553:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
		goto st532
	tr560:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st532
	tr567:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
		goto st532
	tr573:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
		goto st532
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
//line query/tokeniser.go:11517
		switch data[p] {
		case 32:
			goto tr1204
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 46:
			goto tr1205
		case 59:
			goto tr1207
		case 92:
			goto st213
		case 95:
			goto st532
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1204
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st532
				}
			case data[p] >= 65:
				goto st532
			}
		default:
			goto st532
		}
		goto st212
	tr1205:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st224
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
//line query/tokeniser.go:11561
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 95:
			goto st533
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st533
			}
		case data[p] >= 65:
			goto st533
		}
		goto st212
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		switch data[p] {
		case 32:
			goto tr1208
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 46:
			goto st224
		case 59:
			goto tr1210
		case 92:
			goto st213
		case 95:
			goto st533
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1208
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st533
				}
			case data[p] >= 65:
				goto st533
			}
		default:
			goto st533
		}
		goto st212
	tr517:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st225
	tr522:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st225
	tr578:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st225
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
//line query/tokeniser.go:11653
		switch data[p] {
		case 32:
			goto tr541
		case 34:
			goto tr542
		case 39:
			goto tr543
		case 43:
			goto tr544
		case 45:
			goto tr544
		case 61:
			goto st226
		case 92:
			goto st213
		case 95:
			goto tr547
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr541
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr547
				}
			case data[p] >= 65:
				goto tr547
			}
		default:
			goto tr545
		}
		goto st212
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 32:
			goto tr548
		case 34:
			goto tr549
		case 39:
			goto tr550
		case 43:
			goto tr551
		case 45:
			goto tr551
		case 92:
			goto st213
		case 95:
			goto tr553
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr548
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr553
				}
			case data[p] >= 65:
				goto tr553
			}
		default:
			goto tr552
		}
		goto st212
	tr518:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st227
	tr523:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st227
	tr579:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st227
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
//line query/tokeniser.go:11756
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 61:
			goto st228
		case 92:
			goto st213
		}
		goto st212
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 32:
			goto tr555
		case 34:
			goto tr556
		case 39:
			goto tr557
		case 43:
			goto tr558
		case 45:
			goto tr558
		case 92:
			goto st213
		case 95:
			goto tr560
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr555
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr560
				}
			case data[p] >= 65:
				goto tr560
			}
		default:
			goto tr559
		}
		goto st212
	tr519:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st229
	tr524:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st229
	tr580:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line query/tokeniser.go:11840
		switch data[p] {
		case 32:
			goto tr561
		case 34:
			goto tr562
		case 39:
			goto tr563
		case 43:
			goto tr564
		case 45:
			goto tr564
		case 61:
			goto st230
		case 92:
			goto st213
		case 95:
			goto tr567
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr561
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr567
				}
			case data[p] >= 65:
				goto tr567
			}
		default:
			goto tr565
		}
		goto st212
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 32:
			goto tr568
		case 34:
			goto tr569
		case 39:
			goto tr570
		case 43:
			goto tr571
		case 45:
			goto tr571
		case 92:
			goto st213
		case 95:
			goto tr573
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr568
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr573
				}
			case data[p] >= 65:
				goto tr573
			}
		default:
			goto tr572
		}
		goto st212
	tr515:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st231
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
//line query/tokeniser.go:11925
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 95:
			goto st232
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st232
			}
		case data[p] >= 65:
			goto st232
		}
		goto st212
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		switch data[p] {
		case 32:
			goto tr575
		case 33:
			goto tr576
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 46:
			goto st231
		case 60:
			goto tr578
		case 61:
			goto tr579
		case 62:
			goto tr580
		case 92:
			goto st213
		case 95:
			goto st232
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr575
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st232
				}
			case data[p] >= 65:
				goto st232
			}
		default:
			goto st232
		}
		goto st212
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		switch data[p] {
		case 32:
			goto tr581
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 95:
			goto tr582
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr581
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr582
			}
		default:
			goto tr582
		}
		goto st212
	tr581:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st234
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
//line query/tokeniser.go:12031
		switch data[p] {
		case 32:
			goto st234
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 95:
			goto st235
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st234
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st235
			}
		default:
			goto st235
		}
		goto st212
	tr582:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st235
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
//line query/tokeniser.go:12068
		switch data[p] {
		case 32:
			goto tr585
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 93:
			goto tr586
		case 95:
			goto st235
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr585
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st235
				}
			case data[p] >= 65:
				goto st235
			}
		default:
			goto st235
		}
		goto st212
	tr585:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
//line query/tokeniser.go:12110
		switch data[p] {
		case 32:
			goto st236
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 93:
			goto st534
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st236
		}
		goto st212
	tr586:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st534
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
//line query/tokeniser.go:12136
		switch data[p] {
		case 32:
			goto tr1211
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 59:
			goto tr1212
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1211
		}
		goto st212
	tr1191:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st237
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
//line query/tokeniser.go:12162
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 78:
			goto st238
		case 92:
			goto st213
		case 110:
			goto st238
		}
		goto st212
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 68:
			goto st215
		case 92:
			goto st213
		case 100:
			goto st215
		}
		goto st212
	tr1192:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st239
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
//line query/tokeniser.go:12203
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 82:
			goto st240
		case 92:
			goto st213
		case 114:
			goto st240
		}
		goto st212
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 32:
			goto tr591
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr591
		}
		goto st212
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 73:
			goto st242
		case 92:
			goto st213
		case 105:
			goto st242
		}
		goto st212
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 84:
			goto st243
		case 92:
			goto st213
		case 116:
			goto st243
		}
		goto st212
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 72:
			goto st244
		case 92:
			goto st213
		case 104:
			goto st244
		}
		goto st212
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 73:
			goto st245
		case 92:
			goto st213
		case 105:
			goto st245
		}
		goto st212
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 78:
			goto st246
		case 92:
			goto st213
		case 110:
			goto st246
		}
		goto st212
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		switch data[p] {
		case 32:
			goto st247
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st247
		}
		goto st212
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		switch data[p] {
		case 32:
			goto st247
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 43:
			goto tr598
		case 45:
			goto tr598
		case 92:
			goto st213
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr599
			}
		case data[p] >= 9:
			goto st247
		}
		goto st212
	tr598:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st248
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
//line query/tokeniser.go:12386
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st249
		}
		goto st212
	tr599:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st249
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
//line query/tokeniser.go:12412
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 46:
			goto st250
		case 72:
			goto st535
		case 77:
			goto st537
		case 78:
			goto st252
		case 83:
			goto st535
		case 85:
			goto st252
		case 92:
			goto st213
		case 104:
			goto st535
		case 109:
			goto st537
		case 110:
			goto st252
		case 115:
			goto st535
		case 117:
			goto st252
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st249
		}
		goto st212
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st251
		}
		goto st212
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 72:
			goto st535
		case 77:
			goto st537
		case 78:
			goto st252
		case 83:
			goto st535
		case 85:
			goto st252
		case 92:
			goto st213
		case 104:
			goto st535
		case 109:
			goto st537
		case 110:
			goto st252
		case 115:
			goto st535
		case 117:
			goto st252
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st251
		}
		goto st212
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		switch data[p] {
		case 32:
			goto tr1213
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 43:
			goto st248
		case 45:
			goto st248
		case 59:
			goto tr1215
		case 92:
			goto st213
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st249
			}
		case data[p] >= 9:
			goto tr1213
		}
		goto st212
	tr1213:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st536
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
//line query/tokeniser.go:12544
		switch data[p] {
		case 32:
			goto st536
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 59:
			goto st529
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st536
		}
		goto st212
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		switch data[p] {
		case 32:
			goto tr1213
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 43:
			goto st248
		case 45:
			goto st248
		case 59:
			goto tr1215
		case 83:
			goto st535
		case 92:
			goto st213
		case 115:
			goto st535
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st249
			}
		case data[p] >= 9:
			goto tr1213
		}
		goto st212
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 83:
			goto st535
		case 92:
			goto st213
		case 115:
			goto st535
		}
		goto st212
	tr1195:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st253
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
//line query/tokeniser.go:12622
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 124:
			goto st240
		}
		goto st212
	tr1196:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st254
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
//line query/tokeniser.go:12643
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 136:
			goto st255
		}
		goto st212
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 92:
			goto st213
		case 168:
			goto st240
		}
		goto st212
	tr507:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st538
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
//line query/tokeniser.go:12680
		switch data[p] {
		case 32:
			goto tr1217
		case 34:
			goto tr120
		case 39:
			goto tr504
		case 59:
			goto tr1218
		case 92:
			goto st213
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1217
		}
		goto st212
	tr490:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st256
	tr496:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st256
	tr613:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st256
	tr620:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st256
	tr627:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st256
	tr633:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st256
	tr640:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st256
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
//line query/tokeniser.go:12756
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st539
		}
		goto st23
	tr491:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st539
	tr497:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st539
	tr614:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st539
	tr621:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st539
	tr628:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st539
	tr634:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st539
	tr641:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st539
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
//line query/tokeniser.go:12826
		switch data[p] {
		case 32:
			goto tr1219
		case 34:
			goto tr55
		case 46:
			goto st257
		case 59:
			goto tr1221
		case 92:
			goto st202
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st539
			}
		case data[p] >= 9:
			goto tr1219
		}
		goto st23
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st540
		}
		goto st23
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
		switch data[p] {
		case 32:
			goto tr1219
		case 34:
			goto tr55
		case 59:
			goto tr1221
		case 92:
			goto st202
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st540
			}
		case data[p] >= 9:
			goto tr1219
		}
		goto st23
	tr492:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
		goto st541
	tr498:
//line query/tokeniser.rl:87
		mark = p
		goto st541
	tr616:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
		goto st541
	tr622:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
		goto st541
	tr629:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st541
	tr636:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
		goto st541
	tr642:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
		goto st541
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
//line query/tokeniser.go:12932
		switch data[p] {
		case 32:
			goto tr1222
		case 34:
			goto tr55
		case 46:
			goto tr1223
		case 59:
			goto tr1225
		case 92:
			goto st202
		case 95:
			goto st541
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1222
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st541
				}
			case data[p] >= 65:
				goto st541
			}
		default:
			goto st541
		}
		goto st23
	tr1223:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st258
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
//line query/tokeniser.go:12974
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		case 95:
			goto st542
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st542
			}
		case data[p] >= 65:
			goto st542
		}
		goto st23
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
		switch data[p] {
		case 32:
			goto tr1226
		case 34:
			goto tr55
		case 46:
			goto st258
		case 59:
			goto tr1228
		case 92:
			goto st202
		case 95:
			goto st542
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1226
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st542
				}
			case data[p] >= 65:
				goto st542
			}
		default:
			goto st542
		}
		goto st23
	tr478:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st259
	tr483:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st259
	tr647:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st259
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
//line query/tokeniser.go:13062
		switch data[p] {
		case 32:
			goto tr610
		case 34:
			goto tr611
		case 39:
			goto tr612
		case 43:
			goto tr613
		case 45:
			goto tr613
		case 61:
			goto st260
		case 92:
			goto st202
		case 95:
			goto tr616
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr610
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr616
				}
			case data[p] >= 65:
				goto tr616
			}
		default:
			goto tr614
		}
		goto st23
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 32:
			goto tr617
		case 34:
			goto tr618
		case 39:
			goto tr619
		case 43:
			goto tr620
		case 45:
			goto tr620
		case 92:
			goto st202
		case 95:
			goto tr622
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr617
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr622
				}
			case data[p] >= 65:
				goto tr622
			}
		default:
			goto tr621
		}
		goto st23
	tr479:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st261
	tr484:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st261
	tr648:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st261
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
//line query/tokeniser.go:13165
		switch data[p] {
		case 34:
			goto tr55
		case 61:
			goto st262
		case 92:
			goto st202
		}
		goto st23
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 32:
			goto tr624
		case 34:
			goto tr625
		case 39:
			goto tr626
		case 43:
			goto tr627
		case 45:
			goto tr627
		case 92:
			goto st202
		case 95:
			goto tr629
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr624
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr629
				}
			case data[p] >= 65:
				goto tr629
			}
		default:
			goto tr628
		}
		goto st23
	tr480:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st263
	tr485:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st263
	tr649:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st263
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
//line query/tokeniser.go:13247
		switch data[p] {
		case 32:
			goto tr630
		case 34:
			goto tr631
		case 39:
			goto tr632
		case 43:
			goto tr633
		case 45:
			goto tr633
		case 61:
			goto st264
		case 92:
			goto st202
		case 95:
			goto tr636
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr630
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr636
				}
			case data[p] >= 65:
				goto tr636
			}
		default:
			goto tr634
		}
		goto st23
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 32:
			goto tr637
		case 34:
			goto tr638
		case 39:
			goto tr639
		case 43:
			goto tr640
		case 45:
			goto tr640
		case 92:
			goto st202
		case 95:
			goto tr642
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr637
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr642
				}
			case data[p] >= 65:
				goto tr642
			}
		default:
			goto tr641
		}
		goto st23
	tr476:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st265
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
//line query/tokeniser.go:13332
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		case 95:
			goto st266
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st266
			}
		case data[p] >= 65:
			goto st266
		}
		goto st23
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		switch data[p] {
		case 32:
			goto tr644
		case 33:
			goto tr645
		case 34:
			goto tr55
		case 46:
			goto st265
		case 60:
			goto tr647
		case 61:
			goto tr648
		case 62:
			goto tr649
		case 92:
			goto st202
		case 95:
			goto st266
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr644
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st266
				}
			case data[p] >= 65:
				goto st266
			}
		default:
			goto st266
		}
		goto st23
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		switch data[p] {
		case 32:
			goto tr650
		case 34:
			goto tr55
		case 92:
			goto st202
		case 95:
			goto tr651
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr650
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr651
			}
		default:
			goto tr651
		}
		goto st23
	tr650:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st268
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
//line query/tokeniser.go:13432
		switch data[p] {
		case 32:
			goto st268
		case 34:
			goto tr55
		case 92:
			goto st202
		case 95:
			goto st269
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st268
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st269
			}
		default:
			goto st269
		}
		goto st23
	tr651:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st269
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
//line query/tokeniser.go:13467
		switch data[p] {
		case 32:
			goto tr654
		case 34:
			goto tr55
		case 92:
			goto st202
		case 93:
			goto tr655
		case 95:
			goto st269
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr654
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st269
				}
			case data[p] >= 65:
				goto st269
			}
		default:
			goto st269
		}
		goto st23
	tr654:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st270
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
//line query/tokeniser.go:13507
		switch data[p] {
		case 32:
			goto st270
		case 34:
			goto tr55
		case 92:
			goto st202
		case 93:
			goto st543
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st270
		}
		goto st23
	tr655:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st543
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
//line query/tokeniser.go:13531
		switch data[p] {
		case 32:
			goto tr1229
		case 34:
			goto tr55
		case 59:
			goto tr1230
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1229
		}
		goto st23
	tr1176:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st271
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
//line query/tokeniser.go:13555
		switch data[p] {
		case 34:
			goto tr55
		case 78:
			goto st272
		case 92:
			goto st202
		case 110:
			goto st272
		}
		goto st23
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		switch data[p] {
		case 34:
			goto tr55
		case 68:
			goto st204
		case 92:
			goto st202
		case 100:
			goto st204
		}
		goto st23
	tr1177:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st273
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
//line query/tokeniser.go:13592
		switch data[p] {
		case 34:
			goto tr55
		case 82:
			goto st274
		case 92:
			goto st202
		case 114:
			goto st274
		}
		goto st23
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		switch data[p] {
		case 32:
			goto tr660
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr660
		}
		goto st23
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 34:
			goto tr55
		case 73:
			goto st276
		case 92:
			goto st202
		case 105:
			goto st276
		}
		goto st23
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		switch data[p] {
		case 34:
			goto tr55
		case 84:
			goto st277
		case 92:
			goto st202
		case 116:
			goto st277
		}
		goto st23
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		switch data[p] {
		case 34:
			goto tr55
		case 72:
			goto st278
		case 92:
			goto st202
		case 104:
			goto st278
		}
		goto st23
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		switch data[p] {
		case 34:
			goto tr55
		case 73:
			goto st279
		case 92:
			goto st202
		case 105:
			goto st279
		}
		goto st23
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		switch data[p] {
		case 34:
			goto tr55
		case 78:
			goto st280
		case 92:
			goto st202
		case 110:
			goto st280
		}
		goto st23
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		switch data[p] {
		case 32:
			goto st281
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st281
		}
		goto st23
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		switch data[p] {
		case 32:
			goto st281
		case 34:
			goto tr55
		case 43:
			goto tr667
		case 45:
			goto tr667
		case 92:
			goto st202
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr668
			}
		case data[p] >= 9:
			goto st281
		}
		goto st23
	tr667:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st282
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
//line query/tokeniser.go:13757
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st283
		}
		goto st23
	tr668:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st283
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
//line query/tokeniser.go:13781
		switch data[p] {
		case 34:
			goto tr55
		case 46:
			goto st284
		case 72:
			goto st544
		case 77:
			goto st546
		case 78:
			goto st286
		case 83:
			goto st544
		case 85:
			goto st286
		case 92:
			goto st202
		case 104:
			goto st544
		case 109:
			goto st546
		case 110:
			goto st286
		case 115:
			goto st544
		case 117:
			goto st286
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st283
		}
		goto st23
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st285
		}
		goto st23
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		switch data[p] {
		case 34:
			goto tr55
		case 72:
			goto st544
		case 77:
			goto st546
		case 78:
			goto st286
		case 83:
			goto st544
		case 85:
			goto st286
		case 92:
			goto st202
		case 104:
			goto st544
		case 109:
			goto st546
		case 110:
			goto st286
		case 115:
			goto st544
		case 117:
			goto st286
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st285
		}
		goto st23
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		switch data[p] {
		case 32:
			goto tr1231
		case 34:
			goto tr55
		case 43:
			goto st282
		case 45:
			goto st282
		case 59:
			goto tr1233
		case 92:
			goto st202
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st283
			}
		case data[p] >= 9:
			goto tr1231
		}
		goto st23
	tr1231:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st545
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
//line query/tokeniser.go:13905
		switch data[p] {
		case 32:
			goto st545
		case 34:
			goto tr55
		case 59:
			goto st524
		case 92:
			goto st202
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st545
		}
		goto st23
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
		switch data[p] {
		case 32:
			goto tr1231
		case 34:
			goto tr55
		case 43:
			goto st282
		case 45:
			goto st282
		case 59:
			goto tr1233
		case 83:
			goto st544
		case 92:
			goto st202
		case 115:
			goto st544
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st283
			}
		case data[p] >= 9:
			goto tr1231
		}
		goto st23
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		switch data[p] {
		case 34:
			goto tr55
		case 83:
			goto st544
		case 92:
			goto st202
		case 115:
			goto st544
		}
		goto st23
	tr1180:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st287
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
//line query/tokeniser.go:13977
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		case 124:
			goto st274
		}
		goto st23
	tr1181:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st288
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
//line query/tokeniser.go:13996
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		case 136:
			goto st289
		}
		goto st23
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		switch data[p] {
		case 34:
			goto tr55
		case 92:
			goto st202
		case 168:
			goto st274
		}
		goto st23
	tr41:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st290
	tr47:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st290
	tr889:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st290
	tr893:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st290
	tr897:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st290
	tr900:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st290
	tr904:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line query/tokeniser.go:14065
		switch data[p] {
		case 39:
			goto tr398
		case 92:
			goto tr677
		}
		goto tr676
	tr676:
//line query/tokeniser.rl:87
		mark = p
		goto st291
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
//line query/tokeniser.go:14082
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		goto st291
	tr677:
//line query/tokeniser.rl:87
		mark = p
		goto st292
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
//line query/tokeniser.go:14099
		switch data[p] {
		case 39:
			goto tr680
		case 92:
			goto st292
		}
		goto st291
	tr680:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st547
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
//line query/tokeniser.go:14116
		switch data[p] {
		case 32:
			goto tr1235
		case 39:
			goto tr124
		case 59:
			goto tr1236
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1235
		}
		goto st291
	tr1246:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st548
	tr1235:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st548
	tr1283:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st548
	tr1286:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st548
	tr1290:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st548
	tr1293:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st548
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
//line query/tokeniser.go:14178
		switch data[p] {
		case 32:
			goto st548
		case 38:
			goto tr1238
		case 39:
			goto tr124
		case 59:
			goto st550
		case 65:
			goto tr1240
		case 79:
			goto tr1241
		case 87:
			goto st365
		case 92:
			goto st292
		case 94:
			goto tr1243
		case 97:
			goto tr1240
		case 111:
			goto tr1241
		case 119:
			goto st365
		case 124:
			goto tr1244
		case 226:
			goto tr1245
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st548
		}
		goto st291
	tr1238:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st293
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
//line query/tokeniser.go:14222
		switch data[p] {
		case 38:
			goto st294
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		goto st291
	tr1243:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st294
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
//line query/tokeniser.go:14241
		switch data[p] {
		case 32:
			goto tr682
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr682
		}
		goto st291
	tr682:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st295
	tr871:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st295
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
//line query/tokeniser.go:14267
		switch data[p] {
		case 32:
			goto st295
		case 39:
			goto tr124
		case 91:
			goto st357
		case 92:
			goto st292
		case 95:
			goto tr684
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st295
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr684
			}
		default:
			goto tr684
		}
		goto st291
	tr684:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st296
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
//line query/tokeniser.go:14304
		switch data[p] {
		case 32:
			goto tr686
		case 33:
			goto tr687
		case 39:
			goto tr124
		case 46:
			goto tr688
		case 60:
			goto tr690
		case 61:
			goto tr691
		case 62:
			goto tr692
		case 92:
			goto st292
		case 95:
			goto st296
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr686
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st296
				}
			case data[p] >= 65:
				goto st296
			}
		default:
			goto st296
		}
		goto st291
	tr686:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st297
	tr855:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st297
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
//line query/tokeniser.go:14362
		switch data[p] {
		case 32:
			goto st297
		case 33:
			goto tr694
		case 39:
			goto tr124
		case 60:
			goto tr695
		case 61:
			goto tr696
		case 62:
			goto tr697
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st297
		}
		goto st291
	tr687:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st298
	tr694:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st298
	tr856:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st298
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
//line query/tokeniser.go:14410
		switch data[p] {
		case 39:
			goto tr124
		case 61:
			goto st299
		case 92:
			goto st292
		}
		goto st291
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 32:
			goto tr699
		case 34:
			goto tr700
		case 39:
			goto tr701
		case 43:
			goto tr702
		case 45:
			goto tr702
		case 92:
			goto st292
		case 95:
			goto tr704
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr699
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr704
				}
			case data[p] >= 65:
				goto tr704
			}
		default:
			goto tr703
		}
		goto st291
	tr699:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st300
	tr821:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st300
	tr828:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st300
	tr835:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st300
	tr841:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st300
	tr848:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st300
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
//line query/tokeniser.go:14488
		switch data[p] {
		case 32:
			goto st300
		case 34:
			goto tr706
		case 39:
			goto tr707
		case 43:
			goto tr708
		case 45:
			goto tr708
		case 92:
			goto st292
		case 95:
			goto tr710
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st300
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr710
				}
			case data[p] >= 65:
				goto tr710
			}
		default:
			goto tr709
		}
		goto st291
	tr700:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st301
	tr706:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st301
	tr822:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st301
	tr829:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st301
	tr836:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st301
	tr842:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st301
	tr849:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st301
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
//line query/tokeniser.go:14568
		switch data[p] {
		case 34:
			goto tr712
		case 39:
			goto tr157
		case 92:
			goto tr713
		}
		goto tr711
	tr711:
//line query/tokeniser.rl:87
		mark = p
		goto st302
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
//line query/tokeniser.go:14587
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		goto st302
	tr712:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st549
	tr715:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st549
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
//line query/tokeniser.go:14612
		switch data[p] {
		case 32:
			goto tr1246
		case 39:
			goto tr124
		case 59:
			goto tr1247
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1246
		}
		goto st291
	tr1247:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st550
	tr1236:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st550
	tr1285:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st550
	tr1289:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st550
	tr1292:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st550
	tr1294:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st550
	tr1297:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st550
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
//line query/tokeniser.go:14682
		switch data[p] {
		case 32:
			goto st550
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st550
		}
		goto st291
	tr713:
//line query/tokeniser.rl:87
		mark = p
		goto st303
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
//line query/tokeniser.go:14704
		switch data[p] {
		case 34:
			goto tr717
		case 39:
			goto tr718
		case 92:
			goto st303
		}
		goto st302
	tr717:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st551
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
//line query/tokeniser.go:14723
		switch data[p] {
		case 32:
			goto tr1248
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 59:
			goto tr1249
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1248
		}
		goto st302
	tr1248:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st552
	tr1259:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st552
	tr1279:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st552
	tr1263:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st552
	tr1266:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st552
	tr1270:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st552
	tr1273:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st552
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
//line query/tokeniser.go:14795
		switch data[p] {
		case 32:
			goto st552
		case 34:
			goto tr715
		case 38:
			goto tr1251
		case 39:
			goto tr121
		case 59:
			goto st554
		case 65:
			goto tr1253
		case 79:
			goto tr1254
		case 87:
			goto st331
		case 92:
			goto st303
		case 94:
			goto tr1256
		case 97:
			goto tr1253
		case 111:
			goto tr1254
		case 119:
			goto st331
		case 124:
			goto tr1257
		case 226:
			goto tr1258
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st552
		}
		goto st302
	tr1251:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st304
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
//line query/tokeniser.go:14841
		switch data[p] {
		case 34:
			goto tr715
		case 38:
			goto st305
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		goto st302
	tr1256:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st305
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
//line query/tokeniser.go:14862
		switch data[p] {
		case 32:
			goto tr720
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr720
		}
		goto st302
	tr720:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st306
	tr802:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st306
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
//line query/tokeniser.go:14890
		switch data[p] {
		case 32:
			goto st306
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 91:
			goto st323
		case 92:
			goto st303
		case 95:
			goto tr722
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st306
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr722
			}
		default:
			goto tr722
		}
		goto st302
	tr722:
//line query/tokeniser.rl:215
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st307
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
//line query/tokeniser.go:14929
		switch data[p] {
		case 32:
			goto tr724
		case 33:
			goto tr725
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 46:
			goto tr726
		case 60:
			goto tr728
		case 61:
			goto tr729
		case 62:
			goto tr730
		case 92:
			goto st303
		case 95:
			goto st307
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr724
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st307
				}
			case data[p] >= 65:
				goto st307
			}
		default:
			goto st307
		}
		goto st302
	tr724:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st308
	tr786:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
		goto st308
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
//line query/tokeniser.go:14989
		switch data[p] {
		case 32:
			goto st308
		case 33:
			goto tr732
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 60:
			goto tr733
		case 61:
			goto tr734
		case 62:
			goto tr735
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st308
		}
		goto st302
	tr725:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st309
	tr732:
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st309
	tr787:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttNe)
		goto st309
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
//line query/tokeniser.go:15039
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 61:
			goto st310
		case 92:
			goto st303
		}
		goto st302
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		switch data[p] {
		case 32:
			goto tr737
		case 34:
			goto tr738
		case 39:
			goto tr739
		case 43:
			goto tr740
		case 45:
			goto tr740
		case 92:
			goto st303
		case 95:
			goto tr742
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr737
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr742
				}
			case data[p] >= 65:
				goto tr742
			}
		default:
			goto tr741
		}
		goto st302
	tr737:
//line query/tokeniser.rl:156
		commit(ttNe)
		goto st311
	tr752:
//line query/tokeniser.rl:158
		commit(ttLt)
		goto st311
	tr759:
//line query/tokeniser.rl:160
		commit(ttLe)
		goto st311
	tr766:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st311
	tr772:
//line query/tokeniser.rl:157
		commit(ttGt)
		goto st311
	tr779:
//line query/tokeniser.rl:159
		commit(ttGe)
		goto st311
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
//line query/tokeniser.go:15119
		switch data[p] {
		case 32:
			goto st311
		case 34:
			goto tr744
		case 39:
			goto tr745
		case 43:
			goto tr746
		case 45:
			goto tr746
		case 92:
			goto st303
		case 95:
			goto tr748
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st311
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr748
				}
			case data[p] >= 65:
				goto tr748
			}
		default:
			goto tr747
		}
		goto st302
	tr738:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st553
	tr744:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st553
	tr753:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st553
	tr760:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st553
	tr767:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st553
	tr773:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st553
	tr780:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st553
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
//line query/tokeniser.go:15213
		switch data[p] {
		case 32:
			goto tr1259
		case 34:
			goto tr712
		case 39:
			goto tr157
		case 59:
			goto tr1260
		case 92:
			goto tr713
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1259
		}
		goto tr711
	tr1249:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st554
	tr1260:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st554
	tr1280:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st554
	tr1265:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st554
	tr1269:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st554
	tr1272:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:218
		commit(ttPredicate)
		goto st554
	tr1274:
//line query/tokeniser.rl:204
		commit(ttEquivalenceTest)
		goto st554
	tr1277:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st554
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
//line query/tokeniser.go:15293
		switch data[p] {
		case 32:
			goto st554
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st554
		}
		goto st302
	tr739:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st555
	tr745:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st555
	tr754:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st555
	tr761:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st555
	tr768:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st555
	tr774:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st555
	tr781:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st555
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
//line query/tokeniser.go:15367
		switch data[p] {
		case 32:
			goto tr1261
		case 34:
			goto tr116
		case 39:
			goto tr117
		case 59:
			goto tr1262
		case 92:
			goto tr118
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1261
		}
		goto tr115
	tr740:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st312
	tr746:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st312
	tr755:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st312
	tr762:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st312
	tr769:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st312
	tr775:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st312
	tr782:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st312
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
//line query/tokeniser.go:15443
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st556
		}
		goto st302
	tr741:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st556
	tr747:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st556
	tr756:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st556
	tr763:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st556
	tr770:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st556
	tr776:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st556
	tr783:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st556
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
//line query/tokeniser.go:15515
		switch data[p] {
		case 32:
			goto tr1263
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 46:
			goto st313
		case 59:
			goto tr1265
		case 92:
			goto st303
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] >= 9:
			goto tr1263
		}
		goto st302
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st557
		}
		goto st302
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		switch data[p] {
		case 32:
			goto tr1263
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 59:
			goto tr1265
		case 92:
			goto st303
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st557
			}
		case data[p] >= 9:
			goto tr1263
		}
		goto st302
	tr742:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
		goto st558
	tr748:
//line query/tokeniser.rl:87
		mark = p
		goto st558
	tr758:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
		goto st558
	tr764:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
		goto st558
	tr771:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st558
	tr778:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
		goto st558
	tr784:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
		goto st558
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
//line query/tokeniser.go:15627
		switch data[p] {
		case 32:
			goto tr1266
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 46:
			goto tr1267
		case 59:
			goto tr1269
		case 92:
			goto st303
		case 95:
			goto st558
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1266
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st558
				}
			case data[p] >= 65:
				goto st558
			}
		default:
			goto st558
		}
		goto st302
	tr1267:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st314
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
//line query/tokeniser.go:15671
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 95:
			goto st559
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st559
			}
		case data[p] >= 65:
			goto st559
		}
		goto st302
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
		switch data[p] {
		case 32:
			goto tr1270
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 46:
			goto st314
		case 59:
			goto tr1272
		case 92:
			goto st303
		case 95:
			goto st559
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1270
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st559
				}
			case data[p] >= 65:
				goto st559
			}
		default:
			goto st559
		}
		goto st302
	tr728:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st315
	tr733:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st315
	tr789:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st315
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
//line query/tokeniser.go:15763
		switch data[p] {
		case 32:
			goto tr752
		case 34:
			goto tr753
		case 39:
			goto tr754
		case 43:
			goto tr755
		case 45:
			goto tr755
		case 61:
			goto st316
		case 92:
			goto st303
		case 95:
			goto tr758
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr752
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr758
				}
			case data[p] >= 65:
				goto tr758
			}
		default:
			goto tr756
		}
		goto st302
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		switch data[p] {
		case 32:
			goto tr759
		case 34:
			goto tr760
		case 39:
			goto tr761
		case 43:
			goto tr762
		case 45:
			goto tr762
		case 92:
			goto st303
		case 95:
			goto tr764
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr759
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr764
				}
			case data[p] >= 65:
				goto tr764
			}
		default:
			goto tr763
		}
		goto st302
	tr729:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st317
	tr734:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st317
	tr790:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st317
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
//line query/tokeniser.go:15866
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 61:
			goto st318
		case 92:
			goto st303
		}
		goto st302
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		switch data[p] {
		case 32:
			goto tr766
		case 34:
			goto tr767
		case 39:
			goto tr768
		case 43:
			goto tr769
		case 45:
			goto tr769
		case 92:
			goto st303
		case 95:
			goto tr771
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr766
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr771
				}
			case data[p] >= 65:
				goto tr771
			}
		default:
			goto tr770
		}
		goto st302
	tr730:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st319
	tr735:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st319
	tr791:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st319
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
//line query/tokeniser.go:15950
		switch data[p] {
		case 32:
			goto tr772
		case 34:
			goto tr773
		case 39:
			goto tr774
		case 43:
			goto tr775
		case 45:
			goto tr775
		case 61:
			goto st320
		case 92:
			goto st303
		case 95:
			goto tr778
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr772
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr778
				}
			case data[p] >= 65:
				goto tr778
			}
		default:
			goto tr776
		}
		goto st302
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		switch data[p] {
		case 32:
			goto tr779
		case 34:
			goto tr780
		case 39:
			goto tr781
		case 43:
			goto tr782
		case 45:
			goto tr782
		case 92:
			goto st303
		case 95:
			goto tr784
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr779
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr784
				}
			case data[p] >= 65:
				goto tr784
			}
		default:
			goto tr783
		}
		goto st302
	tr726:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st321
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
//line query/tokeniser.go:16035
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 95:
			goto st322
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st322
			}
		case data[p] >= 65:
			goto st322
		}
		goto st302
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		switch data[p] {
		case 32:
			goto tr786
		case 33:
			goto tr787
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 46:
			goto st321
		case 60:
			goto tr789
		case 61:
			goto tr790
		case 62:
			goto tr791
		case 92:
			goto st303
		case 95:
			goto st322
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr786
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st322
				}
			case data[p] >= 65:
				goto st322
			}
		default:
			goto st322
		}
		goto st302
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 32:
			goto tr792
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 95:
			goto tr793
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr792
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr793
			}
		default:
			goto tr793
		}
		goto st302
	tr792:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st324
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
//line query/tokeniser.go:16141
		switch data[p] {
		case 32:
			goto st324
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 95:
			goto st325
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st324
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st325
			}
		default:
			goto st325
		}
		goto st302
	tr793:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st325
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
//line query/tokeniser.go:16178
		switch data[p] {
		case 32:
			goto tr796
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 93:
			goto tr797
		case 95:
			goto st325
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr796
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st325
				}
			case data[p] >= 65:
				goto st325
			}
		default:
			goto st325
		}
		goto st302
	tr796:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st326
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
//line query/tokeniser.go:16220
		switch data[p] {
		case 32:
			goto st326
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 93:
			goto st560
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st326
		}
		goto st302
	tr797:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st560
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
//line query/tokeniser.go:16246
		switch data[p] {
		case 32:
			goto tr1273
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 59:
			goto tr1274
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1273
		}
		goto st302
	tr1253:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st327
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
//line query/tokeniser.go:16272
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 78:
			goto st328
		case 92:
			goto st303
		case 110:
			goto st328
		}
		goto st302
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 68:
			goto st305
		case 92:
			goto st303
		case 100:
			goto st305
		}
		goto st302
	tr1254:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st329
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
//line query/tokeniser.go:16313
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 82:
			goto st330
		case 92:
			goto st303
		case 114:
			goto st330
		}
		goto st302
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		switch data[p] {
		case 32:
			goto tr802
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr802
		}
		goto st302
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 73:
			goto st332
		case 92:
			goto st303
		case 105:
			goto st332
		}
		goto st302
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 84:
			goto st333
		case 92:
			goto st303
		case 116:
			goto st333
		}
		goto st302
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 72:
			goto st334
		case 92:
			goto st303
		case 104:
			goto st334
		}
		goto st302
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 73:
			goto st335
		case 92:
			goto st303
		case 105:
			goto st335
		}
		goto st302
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 78:
			goto st336
		case 92:
			goto st303
		case 110:
			goto st336
		}
		goto st302
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		switch data[p] {
		case 32:
			goto st337
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st337
		}
		goto st302
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		switch data[p] {
		case 32:
			goto st337
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 43:
			goto tr809
		case 45:
			goto tr809
		case 92:
			goto st303
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr810
			}
		case data[p] >= 9:
			goto st337
		}
		goto st302
	tr809:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st338
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
//line query/tokeniser.go:16496
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st339
		}
		goto st302
	tr810:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st339
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
//line query/tokeniser.go:16522
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 46:
			goto st340
		case 72:
			goto st561
		case 77:
			goto st563
		case 78:
			goto st342
		case 83:
			goto st561
		case 85:
			goto st342
		case 92:
			goto st303
		case 104:
			goto st561
		case 109:
			goto st563
		case 110:
			goto st342
		case 115:
			goto st561
		case 117:
			goto st342
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st339
		}
		goto st302
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st341
		}
		goto st302
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 72:
			goto st561
		case 77:
			goto st563
		case 78:
			goto st342
		case 83:
			goto st561
		case 85:
			goto st342
		case 92:
			goto st303
		case 104:
			goto st561
		case 109:
			goto st563
		case 110:
			goto st342
		case 115:
			goto st561
		case 117:
			goto st342
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st341
		}
		goto st302
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
		switch data[p] {
		case 32:
			goto tr1275
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 43:
			goto st338
		case 45:
			goto st338
		case 59:
			goto tr1277
		case 92:
			goto st303
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st339
			}
		case data[p] >= 9:
			goto tr1275
		}
		goto st302
	tr1275:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st562
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
//line query/tokeniser.go:16654
		switch data[p] {
		case 32:
			goto st562
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 59:
			goto st554
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st562
		}
		goto st302
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		switch data[p] {
		case 32:
			goto tr1275
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 43:
			goto st338
		case 45:
			goto st338
		case 59:
			goto tr1277
		case 83:
			goto st561
		case 92:
			goto st303
		case 115:
			goto st561
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st339
			}
		case data[p] >= 9:
			goto tr1275
		}
		goto st302
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 83:
			goto st561
		case 92:
			goto st303
		case 115:
			goto st561
		}
		goto st302
	tr1257:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
//line query/tokeniser.go:16732
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 124:
			goto st330
		}
		goto st302
	tr1258:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st344
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
//line query/tokeniser.go:16753
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 136:
			goto st345
		}
		goto st302
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		switch data[p] {
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 92:
			goto st303
		case 168:
			goto st330
		}
		goto st302
	tr718:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st564
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
//line query/tokeniser.go:16790
		switch data[p] {
		case 32:
			goto tr1279
		case 34:
			goto tr715
		case 39:
			goto tr121
		case 59:
			goto tr1280
		case 92:
			goto st303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1279
		}
		goto st302
	tr701:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st565
	tr707:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st565
	tr823:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st565
	tr830:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st565
	tr837:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st565
	tr843:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st565
	tr850:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st565
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
//line query/tokeniser.go:16866
		switch data[p] {
		case 32:
			goto tr1281
		case 39:
			goto tr398
		case 59:
			goto tr1282
		case 92:
			goto tr399
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1281
		}
		goto tr397
	tr702:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st346
	tr708:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st346
	tr824:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st346
	tr831:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st346
	tr838:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st346
	tr844:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st346
	tr851:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st346
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
//line query/tokeniser.go:16940
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st566
		}
		goto st291
	tr703:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st566
	tr709:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st566
	tr825:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st566
	tr832:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st566
	tr839:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st566
	tr845:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st566
	tr852:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st566
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
//line query/tokeniser.go:17010
		switch data[p] {
		case 32:
			goto tr1283
		case 39:
			goto tr124
		case 46:
			goto st347
		case 59:
			goto tr1285
		case 92:
			goto st292
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st566
			}
		case data[p] >= 9:
			goto tr1283
		}
		goto st291
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st567
		}
		goto st291
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		switch data[p] {
		case 32:
			goto tr1283
		case 39:
			goto tr124
		case 59:
			goto tr1285
		case 92:
			goto st292
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st567
			}
		case data[p] >= 9:
			goto tr1283
		}
		goto st291
	tr704:
//line query/tokeniser.rl:156
		commit(ttNe)
//line query/tokeniser.rl:87
		mark = p
		goto st568
	tr710:
//line query/tokeniser.rl:87
		mark = p
		goto st568
	tr827:
//line query/tokeniser.rl:158
		commit(ttLt)
//line query/tokeniser.rl:87
		mark = p
		goto st568
	tr833:
//line query/tokeniser.rl:160
		commit(ttLe)
//line query/tokeniser.rl:87
		mark = p
		goto st568
	tr840:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st568
	tr847:
//line query/tokeniser.rl:157
		commit(ttGt)
//line query/tokeniser.rl:87
		mark = p
		goto st568
	tr853:
//line query/tokeniser.rl:159
		commit(ttGe)
//line query/tokeniser.rl:87
		mark = p
		goto st568
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
//line query/tokeniser.go:17116
		switch data[p] {
		case 32:
			goto tr1286
		case 39:
			goto tr124
		case 46:
			goto tr1287
		case 59:
			goto tr1289
		case 92:
			goto st292
		case 95:
			goto st568
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1286
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st568
				}
			case data[p] >= 65:
				goto st568
			}
		default:
			goto st568
		}
		goto st291
	tr1287:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st348
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
//line query/tokeniser.go:17158
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		case 95:
			goto st569
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st569
			}
		case data[p] >= 65:
			goto st569
		}
		goto st291
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		switch data[p] {
		case 32:
			goto tr1290
		case 39:
			goto tr124
		case 46:
			goto st348
		case 59:
			goto tr1292
		case 92:
			goto st292
		case 95:
			goto st569
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1290
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st569
				}
			case data[p] >= 65:
				goto st569
			}
		default:
			goto st569
		}
		goto st291
	tr690:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st349
	tr695:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st349
	tr858:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st349
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
//line query/tokeniser.go:17246
		switch data[p] {
		case 32:
			goto tr821
		case 34:
			goto tr822
		case 39:
			goto tr823
		case 43:
			goto tr824
		case 45:
			goto tr824
		case 61:
			goto st350
		case 92:
			goto st292
		case 95:
			goto tr827
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr821
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr827
				}
			case data[p] >= 65:
				goto tr827
			}
		default:
			goto tr825
		}
		goto st291
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 32:
			goto tr828
		case 34:
			goto tr829
		case 39:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr831
		case 92:
			goto st292
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr828
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr833
				}
			case data[p] >= 65:
				goto tr833
			}
		default:
			goto tr832
		}
		goto st291
	tr691:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st351
	tr696:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st351
	tr859:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st351
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
//line query/tokeniser.go:17349
		switch data[p] {
		case 39:
			goto tr124
		case 61:
			goto st352
		case 92:
			goto st292
		}
		goto st291
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		switch data[p] {
		case 32:
			goto tr835
		case 34:
			goto tr836
		case 39:
			goto tr837
		case 43:
			goto tr838
		case 45:
			goto tr838
		case 92:
			goto st292
		case 95:
			goto tr840
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr835
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr840
				}
			case data[p] >= 65:
				goto tr840
			}
		default:
			goto tr839
		}
		goto st291
	tr692:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st353
	tr697:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st353
	tr860:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st353
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
//line query/tokeniser.go:17431
		switch data[p] {
		case 32:
			goto tr841
		case 34:
			goto tr842
		case 39:
			goto tr843
		case 43:
			goto tr844
		case 45:
			goto tr844
		case 61:
			goto st354
		case 92:
			goto st292
		case 95:
			goto tr847
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr841
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr847
				}
			case data[p] >= 65:
				goto tr847
			}
		default:
			goto tr845
		}
		goto st291
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		switch data[p] {
		case 32:
			goto tr848
		case 34:
			goto tr849
		case 39:
			goto tr850
		case 43:
			goto tr851
		case 45:
			goto tr851
		case 92:
			goto st292
		case 95:
			goto tr853
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr848
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr853
				}
			case data[p] >= 65:
				goto tr853
			}
		default:
			goto tr852
		}
		goto st291
	tr688:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st355
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
//line query/tokeniser.go:17516
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		case 95:
			goto st356
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st356
			}
		case data[p] >= 65:
			goto st356
		}
		goto st291
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		switch data[p] {
		case 32:
			goto tr855
		case 33:
			goto tr856
		case 39:
			goto tr124
		case 46:
			goto st355
		case 60:
			goto tr858
		case 61:
			goto tr859
		case 62:
			goto tr860
		case 92:
			goto st292
		case 95:
			goto st356
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr855
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st356
				}
			case data[p] >= 65:
				goto st356
			}
		default:
			goto st356
		}
		goto st291
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		switch data[p] {
		case 32:
			goto tr861
		case 39:
			goto tr124
		case 92:
			goto st292
		case 95:
			goto tr862
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr861
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr862
			}
		default:
			goto tr862
		}
		goto st291
	tr861:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st358
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
//line query/tokeniser.go:17616
		switch data[p] {
		case 32:
			goto st358
		case 39:
			goto tr124
		case 92:
			goto st292
		case 95:
			goto st359
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st358
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st359
			}
		default:
			goto st359
		}
		goto st291
	tr862:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:200
		propose(ttEquivalenceTest)
		goto st359
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
//line query/tokeniser.go:17651
		switch data[p] {
		case 32:
			goto tr865
		case 39:
			goto tr124
		case 92:
			goto st292
		case 93:
			goto tr866
		case 95:
			goto st359
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr865
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st359
				}
			case data[p] >= 65:
				goto st359
			}
		default:
			goto st359
		}
		goto st291
	tr865:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st360
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
//line query/tokeniser.go:17691
		switch data[p] {
		case 32:
			goto st360
		case 39:
			goto tr124
		case 92:
			goto st292
		case 93:
			goto st570
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st360
		}
		goto st291
	tr866:
//line query/tokeniser.rl:202
		setText(ttEquivalenceTest)
		goto st570
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
//line query/tokeniser.go:17715
		switch data[p] {
		case 32:
			goto tr1293
		case 39:
			goto tr124
		case 59:
			goto tr1294
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1293
		}
		goto st291
	tr1240:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st361
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
//line query/tokeniser.go:17739
		switch data[p] {
		case 39:
			goto tr124
		case 78:
			goto st362
		case 92:
			goto st292
		case 110:
			goto st362
		}
		goto st291
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 39:
			goto tr124
		case 68:
			goto st294
		case 92:
			goto st292
		case 100:
			goto st294
		}
		goto st291
	tr1241:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st363
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
//line query/tokeniser.go:17776
		switch data[p] {
		case 39:
			goto tr124
		case 82:
			goto st364
		case 92:
			goto st292
		case 114:
			goto st364
		}
		goto st291
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		switch data[p] {
		case 32:
			goto tr871
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr871
		}
		goto st291
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 39:
			goto tr124
		case 73:
			goto st366
		case 92:
			goto st292
		case 105:
			goto st366
		}
		goto st291
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		switch data[p] {
		case 39:
			goto tr124
		case 84:
			goto st367
		case 92:
			goto st292
		case 116:
			goto st367
		}
		goto st291
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		switch data[p] {
		case 39:
			goto tr124
		case 72:
			goto st368
		case 92:
			goto st292
		case 104:
			goto st368
		}
		goto st291
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 39:
			goto tr124
		case 73:
			goto st369
		case 92:
			goto st292
		case 105:
			goto st369
		}
		goto st291
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		switch data[p] {
		case 39:
			goto tr124
		case 78:
			goto st370
		case 92:
			goto st292
		case 110:
			goto st370
		}
		goto st291
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		switch data[p] {
		case 32:
			goto st371
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st371
		}
		goto st291
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		switch data[p] {
		case 32:
			goto st371
		case 39:
			goto tr124
		case 43:
			goto tr878
		case 45:
			goto tr878
		case 92:
			goto st292
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr879
			}
		case data[p] >= 9:
			goto st371
		}
		goto st291
	tr878:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st372
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
//line query/tokeniser.go:17941
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st373
		}
		goto st291
	tr879:
//line query/tokeniser.rl:232
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:227
		propose(ttDuration)
		goto st373
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
//line query/tokeniser.go:17965
		switch data[p] {
		case 39:
			goto tr124
		case 46:
			goto st374
		case 72:
			goto st571
		case 77:
			goto st573
		case 78:
			goto st376
		case 83:
			goto st571
		case 85:
			goto st376
		case 92:
			goto st292
		case 104:
			goto st571
		case 109:
			goto st573
		case 110:
			goto st376
		case 115:
			goto st571
		case 117:
			goto st376
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st373
		}
		goto st291
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st375
		}
		goto st291
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		switch data[p] {
		case 39:
			goto tr124
		case 72:
			goto st571
		case 77:
			goto st573
		case 78:
			goto st376
		case 83:
			goto st571
		case 85:
			goto st376
		case 92:
			goto st292
		case 104:
			goto st571
		case 109:
			goto st573
		case 110:
			goto st376
		case 115:
			goto st571
		case 117:
			goto st376
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st375
		}
		goto st291
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		switch data[p] {
		case 32:
			goto tr1295
		case 39:
			goto tr124
		case 43:
			goto st372
		case 45:
			goto st372
		case 59:
			goto tr1297
		case 92:
			goto st292
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st373
			}
		case data[p] >= 9:
			goto tr1295
		}
		goto st291
	tr1295:
//line query/tokeniser.rl:228
		setText(ttDuration)
//line query/tokeniser.rl:229
		commit(ttDuration)
//line query/tokeniser.rl:233
		commit(ttWithinClause)
		goto st572
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
//line query/tokeniser.go:18089
		switch data[p] {
		case 32:
			goto st572
		case 39:
			goto tr124
		case 59:
			goto st550
		case 92:
			goto st292
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st572
		}
		goto st291
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		switch data[p] {
		case 32:
			goto tr1295
		case 39:
			goto tr124
		case 43:
			goto st372
		case 45:
			goto st372
		case 59:
			goto tr1297
		case 83:
			goto st571
		case 92:
			goto st292
		case 115:
			goto st571
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st373
			}
		case data[p] >= 9:
			goto tr1295
		}
		goto st291
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		switch data[p] {
		case 39:
			goto tr124
		case 83:
			goto st571
		case 92:
			goto st292
		case 115:
			goto st571
		}
		goto st291
	tr1244:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st377
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
//line query/tokeniser.go:18161
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		case 124:
			goto st364
		}
		goto st291
	tr1245:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st378
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
//line query/tokeniser.go:18180
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		case 136:
			goto st379
		}
		goto st291
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st292
		case 168:
			goto st364
		}
		goto st291
	tr30:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st380
	tr35:
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st380
	tr909:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttLt)
//line query/tokeniser.rl:160
		propose(ttLe)
		goto st380
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
//line query/tokeniser.go:18237
		switch data[p] {
		case 32:
			goto tr887
		case 34:
			goto tr888
		case 39:
			goto tr889
		case 43:
			goto tr406
		case 45:
			goto tr406
		case 61:
			goto st381
		case 95:
			goto tr409
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr887
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr409
				}
			case data[p] >= 65:
				goto tr409
			}
		default:
			goto tr407
		}
		goto st0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 32:
			goto tr891
		case 34:
			goto tr892
		case 39:
			goto tr893
		case 43:
			goto tr413
		case 45:
			goto tr413
		case 95:
			goto tr415
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr891
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr415
				}
			case data[p] >= 65:
				goto tr415
			}
		default:
			goto tr414
		}
		goto st0
	tr31:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st382
	tr36:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st382
	tr910:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st382
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
//line query/tokeniser.go:18336
		if data[p] == 61 {
			goto st383
		}
		goto st0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 32:
			goto tr895
		case 34:
			goto tr896
		case 39:
			goto tr897
		case 43:
			goto tr420
		case 45:
			goto tr420
		case 95:
			goto tr422
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr895
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr422
				}
			case data[p] >= 65:
				goto tr422
			}
		default:
			goto tr421
		}
		goto st0
	tr32:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st384
	tr37:
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st384
	tr911:
//line query/tokeniser.rl:211
		setText(ttAttributeSelector)
//line query/tokeniser.rl:212
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttGt)
//line query/tokeniser.rl:159
		propose(ttGe)
		goto st384
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
//line query/tokeniser.go:18411
		switch data[p] {
		case 32:
			goto tr898
		case 34:
			goto tr899
		case 39:
			goto tr900
		case 43:
			goto tr426
		case 45:
			goto tr426
		case 61:
			goto st385
		case 95:
			goto tr429
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr898
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr429
				}
			case data[p] >= 65:
				goto tr429
			}
		default:
			goto tr427
		}
		goto st0
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 32:
			goto tr902
		case 34:
			goto tr903
		case 39:
			goto tr904
		case 43:
			goto tr433
		case 45:
			goto tr433
		case 95:
			goto tr435
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr902
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr435
				}
			case data[p] >= 65:
				goto tr435
			}
		default:
			goto tr434
		}
		goto st0
	tr28:
//line query/tokeniser.rl:210
		propose(ttAttributeSelector)
		goto st386
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
//line query/tokeniser.go:18492
		if data[p] == 95 {
			goto st387
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st387
			}
		case data[p] >= 65:
			goto st387
		}
		goto st0
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		switch data[p] {
		case 32:
			goto tr906
		case 33:
			goto tr907
		case 46:
			goto st386
		case 60:
			goto tr909
		case 61:
			goto tr910
		case 62:
			goto tr911
		case 95:
			goto st387
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr906
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st387
				}
			case data[p] >= 65:
				goto st387
			}
		default:
			goto st387
		}
		goto st0
	tr16:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
//line query/tokeniser.rl:119
		propose(ttAnyDecl)
		goto st388
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
//line query/tokeniser.go:18559
		switch data[p] {
		case 32:
			goto tr912
		case 78:
			goto st394
		case 95:
			goto st393
		case 110:
			goto st394
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr912
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st393
				}
			case data[p] >= 65:
				goto st393
			}
		default:
			goto st393
		}
		goto st0
	tr912:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st389
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
//line query/tokeniser.go:18599
		switch data[p] {
		case 32:
			goto st389
		case 95:
			goto tr916
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st389
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr916
			}
		default:
			goto tr916
		}
		goto st0
	tr916:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st390
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
//line query/tokeniser.go:18630
		switch data[p] {
		case 32:
			goto tr917
		case 41:
			goto tr918
		case 44:
			goto tr919
		case 95:
			goto st390
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr917
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st390
				}
			case data[p] >= 65:
				goto st390
			}
		default:
			goto st390
		}
		goto st0
	tr917:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st391
	tr928:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st391
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
//line query/tokeniser.go:18676
		switch data[p] {
		case 32:
			goto st391
		case 41:
			goto st466
		case 44:
			goto st392
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st391
		}
		goto st0
	tr919:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st392
	tr930:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st392
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
//line query/tokeniser.go:18706
		switch data[p] {
		case 32:
			goto st392
		case 65:
			goto tr16
		case 95:
			goto tr17
		case 97:
			goto tr16
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st392
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto st0
	tr17:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st393
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
//line query/tokeniser.go:18743
		switch data[p] {
		case 32:
			goto tr912
		case 95:
			goto st393
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr912
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st393
				}
			case data[p] >= 65:
				goto st393
			}
		default:
			goto st393
		}
		goto st0
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		switch data[p] {
		case 32:
			goto tr912
		case 89:
			goto st395
		case 95:
			goto st393
		case 121:
			goto st395
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr912
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st393
				}
			case data[p] >= 65:
				goto st393
			}
		default:
			goto st393
		}
		goto st0
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		switch data[p] {
		case 32:
			goto tr924
		case 40:
			goto st397
		case 95:
			goto st393
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr924
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st393
				}
			case data[p] >= 65:
				goto st393
			}
		default:
			goto st393
		}
		goto st0
	tr924:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st396
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
//line query/tokeniser.go:18843
		switch data[p] {
		case 32:
			goto st389
		case 40:
			goto st397
		case 95:
			goto tr916
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st389
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr916
			}
		default:
			goto tr916
		}
		goto st0
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		switch data[p] {
		case 32:
			goto st397
		case 41:
			goto st398
		case 95:
			goto tr927
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st397
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr927
			}
		default:
			goto tr927
		}
		goto st0
	tr936:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st398
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
//line query/tokeniser.go:18904
		switch data[p] {
		case 32:
			goto tr928
		case 41:
			goto tr929
		case 44:
			goto tr930
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr928
		}
		goto st0
	tr927:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st399
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
//line query/tokeniser.go:18930
		switch data[p] {
		case 32:
			goto tr931
		case 95:
			goto st399
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr931
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st399
				}
			case data[p] >= 65:
				goto st399
			}
		default:
			goto st399
		}
		goto st0
	tr931:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st400
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
//line query/tokeniser.go:18966
		switch data[p] {
		case 32:
			goto st400
		case 95:
			goto tr934
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st400
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr934
			}
		default:
			goto tr934
		}
		goto st0
	tr934:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st401
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
//line query/tokeniser.go:18997
		switch data[p] {
		case 32:
			goto tr935
		case 41:
			goto tr936
		case 44:
			goto tr937
		case 95:
			goto st401
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr935
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st401
				}
			case data[p] >= 65:
				goto st401
			}
		default:
			goto st401
		}
		goto st0
	tr935:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st402
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
//line query/tokeniser.go:19039
		switch data[p] {
		case 32:
			goto st402
		case 41:
			goto st398
		case 44:
			goto st403
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st402
		}
		goto st0
	tr937:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st403
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
//line query/tokeniser.go:19065
		switch data[p] {
		case 32:
			goto st403
		case 95:
			goto tr927
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st403
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr927
			}
		default:
			goto tr927
		}
		goto st0
	tr9:
//line query/tokeniser.rl:151
		propose(ttEventClause)
//line query/tokeniser.rl:119
		propose(ttAnyDecl)
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st404
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
//line query/tokeniser.go:19102
		switch data[p] {
		case 32:
			goto tr941
		case 78:
			goto st407
		case 95:
			goto st406
		case 110:
			goto st407
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr941
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	tr941:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st405
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
//line query/tokeniser.go:19142
		switch data[p] {
		case 32:
			goto st405
		case 95:
			goto tr945
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st405
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr945
			}
		default:
			goto tr945
		}
		goto st0
	tr945:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st574
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
//line query/tokeniser.go:19173
		switch data[p] {
		case 32:
			goto tr1299
		case 59:
			goto tr1301
		case 95:
			goto st574
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1299
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st574
				}
			case data[p] >= 65:
				goto st574
			}
		default:
			goto st574
		}
		goto st0
	tr10:
//line query/tokeniser.rl:151
		propose(ttEventClause)
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st406
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
//line query/tokeniser.go:19215
		switch data[p] {
		case 32:
			goto tr941
		case 95:
			goto st406
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr941
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 32:
			goto tr941
		case 89:
			goto st408
		case 95:
			goto st406
		case 121:
			goto st408
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr941
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		switch data[p] {
		case 32:
			goto tr947
		case 40:
			goto st410
		case 95:
			goto st406
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr947
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	tr947:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st409
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
//line query/tokeniser.go:19315
		switch data[p] {
		case 32:
			goto st405
		case 40:
			goto st410
		case 95:
			goto tr945
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st405
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr945
			}
		default:
			goto tr945
		}
		goto st0
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		switch data[p] {
		case 32:
			goto st410
		case 41:
			goto st575
		case 95:
			goto tr950
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st410
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr950
			}
		default:
			goto tr950
		}
		goto st0
	tr956:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st575
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
//line query/tokeniser.go:19376
		switch data[p] {
		case 32:
			goto tr1302
		case 59:
			goto tr1303
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1302
		}
		goto st0
	tr950:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st411
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
//line query/tokeniser.go:19400
		switch data[p] {
		case 32:
			goto tr951
		case 95:
			goto st411
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr951
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st411
				}
			case data[p] >= 65:
				goto st411
			}
		default:
			goto st411
		}
		goto st0
	tr951:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st412
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
//line query/tokeniser.go:19436
		switch data[p] {
		case 32:
			goto st412
		case 95:
			goto tr954
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st412
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr954
			}
		default:
			goto tr954
		}
		goto st0
	tr954:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st413
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
//line query/tokeniser.go:19467
		switch data[p] {
		case 32:
			goto tr955
		case 41:
			goto tr956
		case 44:
			goto tr957
		case 95:
			goto st413
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr955
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st413
				}
			case data[p] >= 65:
				goto st413
			}
		default:
			goto st413
		}
		goto st0
	tr955:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st414
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
//line query/tokeniser.go:19509
		switch data[p] {
		case 32:
			goto st414
		case 41:
			goto st575
		case 44:
			goto st415
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st414
		}
		goto st0
	tr957:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st415
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
//line query/tokeniser.go:19535
		switch data[p] {
		case 32:
			goto st415
		case 95:
			goto tr950
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st415
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr950
			}
		default:
			goto tr950
		}
		goto st0
	tr11:
//line query/tokeniser.rl:151
		propose(ttEventClause)
//line query/tokeniser.rl:131
		propose(ttNegatedDecl)
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st416
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
//line query/tokeniser.go:19572
		switch data[p] {
		case 32:
			goto tr941
		case 79:
			goto st417
		case 95:
			goto st406
		case 111:
			goto st417
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr941
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		switch data[p] {
		case 32:
			goto tr941
		case 84:
			goto st418
		case 95:
			goto st406
		case 116:
			goto st418
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr941
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		switch data[p] {
		case 32:
			goto tr963
		case 40:
			goto st10
		case 95:
			goto st406
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr963
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	tr963:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st419
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
//line query/tokeniser.go:19676
		switch data[p] {
		case 32:
			goto st405
		case 40:
			goto st10
		case 95:
			goto tr945
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st405
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr945
			}
		default:
			goto tr945
		}
		goto st0
	tr12:
//line query/tokeniser.rl:151
		propose(ttEventClause)
//line query/tokeniser.rl:142
		propose(ttSeqDecl)
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
//line query/tokeniser.go:19715
		switch data[p] {
		case 32:
			goto tr941
		case 69:
			goto st421
		case 95:
			goto st406
		case 101:
			goto st421
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr941
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		switch data[p] {
		case 32:
			goto tr941
		case 81:
			goto st422
		case 95:
			goto st406
		case 113:
			goto st422
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr941
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		switch data[p] {
		case 32:
			goto tr966
		case 40:
			goto st424
		case 95:
			goto st406
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr966
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st406
				}
			case data[p] >= 65:
				goto st406
			}
		default:
			goto st406
		}
		goto st0
	tr966:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st423
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
//line query/tokeniser.go:19819
		switch data[p] {
		case 32:
			goto st405
		case 40:
			goto st424
		case 95:
			goto tr945
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st405
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr945
			}
		default:
			goto tr945
		}
		goto st0
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 32:
			goto st425
		case 33:
			goto tr969
		case 41:
			goto st576
		case 65:
			goto tr971
		case 78:
			goto tr973
		case 95:
			goto tr972
		case 97:
			goto tr971
		case 110:
			goto tr973
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st425
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr972
			}
		default:
			goto tr972
		}
		goto st0
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		switch data[p] {
		case 32:
			goto st425
		case 41:
			goto st576
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st425
		}
		goto st0
	tr990:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st576
	tr999:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st576
	tr980:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
		goto st576
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
//line query/tokeniser.go:19913
		switch data[p] {
		case 32:
			goto tr1304
		case 59:
			goto tr1305
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1304
		}
		goto st0
	tr969:
//line query/tokeniser.rl:131
		propose(ttNegatedDecl)
		goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line query/tokeniser.go:19933
		switch data[p] {
		case 32:
			goto st427
		case 40:
			goto st428
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st427
		}
		goto st0
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		if data[p] == 40 {
			goto st428
		}
		goto st0
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		switch data[p] {
		case 32:
			goto st428
		case 41:
			goto st429
		case 65:
			goto tr977
		case 95:
			goto tr978
		case 97:
			goto tr977
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st428
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr978
			}
		default:
			goto tr978
		}
		goto st0
	tr1020:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st429
	tr1031:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
//line query/tokeniser.go:20000
		switch data[p] {
		case 32:
			goto tr979
		case 41:
			goto tr980
		case 44:
			goto tr981
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr979
		}
		goto st0
	tr989:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st430
	tr998:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st430
	tr979:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
		goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
//line query/tokeniser.go:20034
		switch data[p] {
		case 32:
			goto st430
		case 41:
			goto st576
		case 44:
			goto st431
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st430
		}
		goto st0
	tr991:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st431
	tr1000:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st431
	tr981:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
		goto st431
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
//line query/tokeniser.go:20068
		switch data[p] {
		case 32:
			goto st431
		case 33:
			goto tr969
		case 65:
			goto tr971
		case 78:
			goto tr973
		case 95:
			goto tr972
		case 97:
			goto tr971
		case 110:
			goto tr973
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st431
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr972
			}
		default:
			goto tr972
		}
		goto st0
	tr971:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
//line query/tokeniser.rl:119
		propose(ttAnyDecl)
		goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
//line query/tokeniser.go:20113
		switch data[p] {
		case 32:
			goto tr984
		case 78:
			goto st436
		case 95:
			goto st435
		case 110:
			goto st436
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr984
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st435
				}
			case data[p] >= 65:
				goto st435
			}
		default:
			goto st435
		}
		goto st0
	tr984:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st433
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
//line query/tokeniser.go:20153
		switch data[p] {
		case 32:
			goto st433
		case 95:
			goto tr988
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st433
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr988
			}
		default:
			goto tr988
		}
		goto st0
	tr988:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st434
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
//line query/tokeniser.go:20184
		switch data[p] {
		case 32:
			goto tr989
		case 41:
			goto tr990
		case 44:
			goto tr991
		case 95:
			goto st434
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr989
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st434
				}
			case data[p] >= 65:
				goto st434
			}
		default:
			goto st434
		}
		goto st0
	tr972:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st435
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
//line query/tokeniser.go:20226
		switch data[p] {
		case 32:
			goto tr984
		case 95:
			goto st435
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr984
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st435
				}
			case data[p] >= 65:
				goto st435
			}
		default:
			goto st435
		}
		goto st0
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		switch data[p] {
		case 32:
			goto tr984
		case 89:
			goto st437
		case 95:
			goto st435
		case 121:
			goto st437
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr984
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st435
				}
			case data[p] >= 65:
				goto st435
			}
		default:
			goto st435
		}
		goto st0
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		switch data[p] {
		case 32:
			goto tr994
		case 40:
			goto st439
		case 95:
			goto st435
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr994
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st435
				}
			case data[p] >= 65:
				goto st435
			}
		default:
			goto st435
		}
		goto st0
	tr994:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st438
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
//line query/tokeniser.go:20326
		switch data[p] {
		case 32:
			goto st433
		case 40:
			goto st439
		case 95:
			goto tr988
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st433
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr988
			}
		default:
			goto tr988
		}
		goto st0
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 32:
			goto st439
		case 41:
			goto st440
		case 95:
			goto tr997
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st439
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr997
			}
		default:
			goto tr997
		}
		goto st0
	tr1006:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
//line query/tokeniser.go:20387
		switch data[p] {
		case 32:
			goto tr998
		case 41:
			goto tr999
		case 44:
			goto tr1000
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr998
		}
		goto st0
	tr997:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
//line query/tokeniser.go:20413
		switch data[p] {
		case 32:
			goto tr1001
		case 95:
			goto st441
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1001
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st441
				}
			case data[p] >= 65:
				goto st441
			}
		default:
			goto st441
		}
		goto st0
	tr1001:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st442
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
//line query/tokeniser.go:20449
		switch data[p] {
		case 32:
			goto st442
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st442
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	tr1004:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line query/tokeniser.go:20480
		switch data[p] {
		case 32:
			goto tr1005
		case 41:
			goto tr1006
		case 44:
			goto tr1007
		case 95:
			goto st443
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1005
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st443
				}
			case data[p] >= 65:
				goto st443
			}
		default:
			goto st443
		}
		goto st0
	tr1005:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st444
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
//line query/tokeniser.go:20522
		switch data[p] {
		case 32:
			goto st444
		case 41:
			goto st440
		case 44:
			goto st445
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st444
		}
		goto st0
	tr1007:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
//line query/tokeniser.go:20548
		switch data[p] {
		case 32:
			goto st445
		case 95:
			goto tr997
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st445
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr997
			}
		default:
			goto tr997
		}
		goto st0
	tr973:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
//line query/tokeniser.rl:131
		propose(ttNegatedDecl)
		goto st446
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
//line query/tokeniser.go:20583
		switch data[p] {
		case 32:
			goto tr984
		case 79:
			goto st447
		case 95:
			goto st435
		case 111:
			goto st447
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr984
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st435
				}
			case data[p] >= 65:
				goto st435
			}
		default:
			goto st435
		}
		goto st0
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		switch data[p] {
		case 32:
			goto tr984
		case 84:
			goto st448
		case 95:
			goto st435
		case 116:
			goto st448
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr984
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st435
				}
			case data[p] >= 65:
				goto st435
			}
		default:
			goto st435
		}
		goto st0
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		switch data[p] {
		case 32:
			goto tr1013
		case 40:
			goto st428
		case 95:
			goto st435
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1013
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st435
				}
			case data[p] >= 65:
				goto st435
			}
		default:
			goto st435
		}
		goto st0
	tr1013:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st449
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
//line query/tokeniser.go:20687
		switch data[p] {
		case 32:
			goto st433
		case 40:
			goto st428
		case 95:
			goto tr988
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st433
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr988
			}
		default:
			goto tr988
		}
		goto st0
	tr977:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
//line query/tokeniser.rl:119
		propose(ttAnyDecl)
		goto st450
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
//line query/tokeniser.go:20724
		switch data[p] {
		case 32:
			goto tr1014
		case 78:
			goto st456
		case 95:
			goto st455
		case 110:
			goto st456
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1014
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st455
				}
			case data[p] >= 65:
				goto st455
			}
		default:
			goto st455
		}
		goto st0
	tr1014:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st451
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
//line query/tokeniser.go:20764
		switch data[p] {
		case 32:
			goto st451
		case 95:
			goto tr1018
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st451
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1018
			}
		default:
			goto tr1018
		}
		goto st0
	tr1018:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st452
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
//line query/tokeniser.go:20795
		switch data[p] {
		case 32:
			goto tr1019
		case 41:
			goto tr1020
		case 44:
			goto tr1021
		case 95:
			goto st452
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1019
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st452
				}
			case data[p] >= 65:
				goto st452
			}
		default:
			goto st452
		}
		goto st0
	tr1019:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st453
	tr1030:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st453
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
//line query/tokeniser.go:20841
		switch data[p] {
		case 32:
			goto st453
		case 41:
			goto st429
		case 44:
			goto st454
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st453
		}
		goto st0
	tr1021:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st454
	tr1032:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st454
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
//line query/tokeniser.go:20871
		switch data[p] {
		case 32:
			goto st454
		case 65:
			goto tr977
		case 95:
			goto tr978
		case 97:
			goto tr977
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st454
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr978
			}
		default:
			goto tr978
		}
		goto st0
	tr978:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st455
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
//line query/tokeniser.go:20908
		switch data[p] {
		case 32:
			goto tr1014
		case 95:
			goto st455
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1014
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st455
				}
			case data[p] >= 65:
				goto st455
			}
		default:
			goto st455
		}
		goto st0
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		switch data[p] {
		case 32:
			goto tr1014
		case 89:
			goto st457
		case 95:
			goto st455
		case 121:
			goto st457
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1014
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st455
				}
			case data[p] >= 65:
				goto st455
			}
		default:
			goto st455
		}
		goto st0
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 32:
			goto tr1026
		case 40:
			goto st459
		case 95:
			goto st455
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1026
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st455
				}
			case data[p] >= 65:
				goto st455
			}
		default:
			goto st455
		}
		goto st0
	tr1026:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
//line query/tokeniser.go:21008
		switch data[p] {
		case 32:
			goto st451
		case 40:
			goto st459
		case 95:
			goto tr1018
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st451
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1018
			}
		default:
			goto tr1018
		}
		goto st0
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		switch data[p] {
		case 32:
			goto st459
		case 41:
			goto st460
		case 95:
			goto tr1029
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st459
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1029
			}
		default:
			goto tr1029
		}
		goto st0
	tr1038:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
//line query/tokeniser.go:21069
		switch data[p] {
		case 32:
			goto tr1030
		case 41:
			goto tr1031
		case 44:
			goto tr1032
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1030
		}
		goto st0
	tr1029:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st461
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
//line query/tokeniser.go:21095
		switch data[p] {
		case 32:
			goto tr1033
		case 95:
			goto st461
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1033
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st461
				}
			case data[p] >= 65:
				goto st461
			}
		default:
			goto st461
		}
		goto st0
	tr1033:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st462
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
//line query/tokeniser.go:21131
		switch data[p] {
		case 32:
			goto st462
		case 95:
			goto tr1036
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st462
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1036
			}
		default:
			goto tr1036
		}
		goto st0
	tr1036:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st463
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
//line query/tokeniser.go:21162
		switch data[p] {
		case 32:
			goto tr1037
		case 41:
			goto tr1038
		case 44:
			goto tr1039
		case 95:
			goto st463
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1037
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st463
				}
			case data[p] >= 65:
				goto st463
			}
		default:
			goto st463
		}
		goto st0
	tr1037:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st464
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
//line query/tokeniser.go:21204
		switch data[p] {
		case 32:
			goto st464
		case 41:
			goto st460
		case 44:
			goto st465
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st464
		}
		goto st0
	tr1039:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st465
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
//line query/tokeniser.go:21230
		switch data[p] {
		case 32:
			goto st465
		case 95:
			goto tr1029
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st465
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1029
			}
		default:
			goto tr1029
		}
		goto st0
	st_out:
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof466:
		cs = 466
		goto _test_eof
	_test_eof467:
		cs = 467
		goto _test_eof
	_test_eof468:
		cs = 468
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof469:
		cs = 469
		goto _test_eof
	_test_eof470:
		cs = 470
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof471:
		cs = 471
		goto _test_eof
	_test_eof472:
		cs = 472
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof473:
		cs = 473
		goto _test_eof
	_test_eof474:
		cs = 474
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof475:
		cs = 475
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof476:
		cs = 476
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof477:
		cs = 477
		goto _test_eof
	_test_eof478:
		cs = 478
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof479:
		cs = 479
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof480:
		cs = 480
		goto _test_eof
	_test_eof481:
		cs = 481
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof482:
		cs = 482
		goto _test_eof
	_test_eof483:
		cs = 483
		goto _test_eof
	_test_eof484:
		cs = 484
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof485:
		cs = 485
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof486:
		cs = 486
		goto _test_eof
	_test_eof487:
		cs = 487
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof488:
		cs = 488
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof78:
		cs = 78
		goto _test_eof
	_test_eof79:
		cs = 79
		goto _test_eof
	_test_eof489:
		cs = 489
		goto _test_eof
	_test_eof80:
		cs = 80
		goto _test_eof
	_test_eof81:
		cs = 81
		goto _test_eof
	_test_eof82:
		cs = 82
		goto _test_eof
	_test_eof83:
		cs = 83
		goto _test_eof
	_test_eof84:
		cs = 84
		goto _test_eof
	_test_eof85:
		cs = 85
		goto _test_eof
	_test_eof86:
		cs = 86
		goto _test_eof
	_test_eof87:
		cs = 87
		goto _test_eof
	_test_eof88:
		cs = 88
		goto _test_eof
	_test_eof89:
		cs = 89
		goto _test_eof
	_test_eof90:
		cs = 90
		goto _test_eof
	_test_eof91:
		cs = 91
		goto _test_eof
	_test_eof92:
		cs = 92
		goto _test_eof
	_test_eof93:
		cs = 93
		goto _test_eof
	_test_eof94:
		cs = 94
		goto _test_eof
	_test_eof490:
		cs = 490
		goto _test_eof
	_test_eof491:
		cs = 491
		goto _test_eof
	_test_eof492:
		cs = 492
		goto _test_eof
	_test_eof95:
		cs = 95
		goto _test_eof
	_test_eof96:
		cs = 96
		goto _test_eof
	_test_eof97:
		cs = 97
		goto _test_eof
	_test_eof98:
		cs = 98
		goto _test_eof
	_test_eof493:
		cs = 493
		goto _test_eof
	_test_eof494:
		cs = 494
		goto _test_eof
	_test_eof495:
		cs = 495
		goto _test_eof
	_test_eof99:
		cs = 99
		goto _test_eof
	_test_eof496:
		cs = 496
		goto _test_eof
	_test_eof100:
		cs = 100
		goto _test_eof
	_test_eof497:
		cs = 497
		goto _test_eof
	_test_eof498:
		cs = 498
		goto _test_eof
	_test_eof101:
		cs = 101
		goto _test_eof
	_test_eof499:
		cs = 499
		goto _test_eof
	_test_eof102:
		cs = 102
		goto _test_eof
	_test_eof103:
		cs = 103
		goto _test_eof
	_test_eof104:
		cs = 104
		goto _test_eof
	_test_eof105:
		cs = 105
		goto _test_eof
	_test_eof106:
		cs = 106
		goto _test_eof
	_test_eof107:
		cs = 107
		goto _test_eof
	_test_eof108:
		cs = 108
		goto _test_eof
	_test_eof109:
		cs = 109
		goto _test_eof
	_test_eof110:
		cs = 110
		goto _test_eof
	_test_eof111:
		cs = 111
		goto _test_eof
	_test_eof112:
		cs = 112
		goto _test_eof
	_test_eof113:
		cs = 113
		goto _test_eof
	_test_eof500:
		cs = 500
		goto _test_eof
	_test_eof114:
		cs = 114
		goto _test_eof
	_test_eof115:
		cs = 115
		goto _test_eof
	_test_eof116:
		cs = 116
		goto _test_eof
	_test_eof117:
		cs = 117
		goto _test_eof
	_test_eof118:
		cs = 118
		goto _test_eof
	_test_eof119:
		cs = 119
		goto _test_eof
	_test_eof120:
		cs = 120
		goto _test_eof
	_test_eof121:
		cs = 121
		goto _test_eof
	_test_eof122:
		cs = 122
		goto _test_eof
	_test_eof123:
		cs = 123
		goto _test_eof
	_test_eof124:
		cs = 124
		goto _test_eof
	_test_eof125:
		cs = 125
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof128:
		cs = 128
		goto _test_eof
	_test_eof501:
		cs = 501
		goto _test_eof
	_test_eof502:
		cs = 502
		goto _test_eof
	_test_eof503:
		cs = 503
		goto _test_eof
	_test_eof129:
		cs = 129
		goto _test_eof
	_test_eof130:
		cs = 130
		goto _test_eof
	_test_eof131:
		cs = 131
		goto _test_eof
	_test_eof132:
		cs = 132
		goto _test_eof
	_test_eof133:
		cs = 133
		goto _test_eof
	_test_eof504:
		cs = 504
		goto _test_eof
	_test_eof134:
		cs = 134
		goto _test_eof
	_test_eof505:
		cs = 505
		goto _test_eof
	_test_eof506:
		cs = 506
		goto _test_eof
	_test_eof135:
		cs = 135
		goto _test_eof
	_test_eof507:
		cs = 507
		goto _test_eof
	_test_eof136:
		cs = 136
		goto _test_eof
	_test_eof137:
		cs = 137
		goto _test_eof
	_test_eof138:
		cs = 138
		goto _test_eof
	_test_eof139:
		cs = 139
		goto _test_eof
	_test_eof140:
		cs = 140
		goto _test_eof
	_test_eof141:
		cs = 141
		goto _test_eof
	_test_eof142:
		cs = 142
		goto _test_eof
	_test_eof143:
		cs = 143
		goto _test_eof
	_test_eof144:
		cs = 144
		goto _test_eof
	_test_eof145:
		cs = 145
		goto _test_eof
	_test_eof146:
		cs = 146
		goto _test_eof
	_test_eof147:
		cs = 147
		goto _test_eof
	_test_eof508:
		cs = 508
		goto _test_eof
	_test_eof148:
		cs = 148
		goto _test_eof
	_test_eof149:
		cs = 149
		goto _test_eof
	_test_eof150:
		cs = 150
		goto _test_eof
	_test_eof151:
		cs = 151
		goto _test_eof
	_test_eof152:
		cs = 152
		goto _test_eof
	_test_eof153:
		cs = 153
		goto _test_eof
	_test_eof154:
		cs = 154
		goto _test_eof
	_test_eof155:
		cs = 155
		goto _test_eof
	_test_eof156:
		cs = 156
		goto _test_eof
	_test_eof157:
		cs = 157
		goto _test_eof
	_test_eof158:
		cs = 158
		goto _test_eof
	_test_eof159:
		cs = 159
		goto _test_eof
	_test_eof160:
		cs = 160
		goto _test_eof
	_test_eof161:
		cs = 161
		goto _test_eof
	_test_eof162:
		cs = 162
		goto _test_eof
	_test_eof509:
		cs = 509
		goto _test_eof
	_test_eof510:
		cs = 510
		goto _test_eof
	_test_eof511:
		cs = 511
		goto _test_eof
	_test_eof163:
		cs = 163
		goto _test_eof
	_test_eof164:
		cs = 164
		goto _test_eof
	_test_eof165:
		cs = 165
		goto _test_eof
	_test_eof166:
		cs = 166
		goto _test_eof
	_test_eof167:
		cs = 167
		goto _test_eof
	_test_eof168:
		cs = 168
		goto _test_eof
	_test_eof512:
		cs = 512
		goto _test_eof
	_test_eof169:
		cs = 169
		goto _test_eof
	_test_eof513:
		cs = 513
		goto _test_eof
	_test_eof514:
		cs = 514
		goto _test_eof
	_test_eof170:
		cs = 170
		goto _test_eof
	_test_eof515:
		cs = 515
		goto _test_eof
	_test_eof171:
		cs = 171
		goto _test_eof
	_test_eof172:
		cs = 172
		goto _test_eof
	_test_eof173:
		cs = 173
		goto _test_eof
	_test_eof174:
		cs = 174
		goto _test_eof
	_test_eof175:
		cs = 175
		goto _test_eof
	_test_eof176:
		cs = 176
		goto _test_eof
	_test_eof177:
		cs = 177
		goto _test_eof
	_test_eof178:
		cs = 178
		goto _test_eof
	_test_eof179:
		cs = 179
		goto _test_eof
	_test_eof180:
		cs = 180
		goto _test_eof
	_test_eof181:
		cs = 181
		goto _test_eof
	_test_eof182:
		cs = 182
		goto _test_eof
	_test_eof516:
		cs = 516
		goto _test_eof
	_test_eof183:
		cs = 183
		goto _test_eof
	_test_eof184:
		cs = 184
		goto _test_eof
	_test_eof185:
		cs = 185
		goto _test_eof
	_test_eof186:
		cs = 186
		goto _test_eof
	_test_eof187:
		cs = 187
		goto _test_eof
	_test_eof188:
		cs = 188
		goto _test_eof
	_test_eof189:
		cs = 189
		goto _test_eof
	_test_eof190:
		cs = 190
		goto _test_eof
	_test_eof191:
		cs = 191
		goto _test_eof
	_test_eof192:
		cs = 192
		goto _test_eof
	_test_eof193:
		cs = 193
		goto _test_eof
	_test_eof194:
		cs = 194
		goto _test_eof
	_test_eof195:
		cs = 195
		goto _test_eof
	_test_eof196:
		cs = 196
		goto _test_eof
	_test_eof197:
		cs = 197
		goto _test_eof
	_test_eof517:
		cs = 517
		goto _test_eof
	_test_eof518:
		cs = 518
		goto _test_eof
	_test_eof519:
		cs = 519
		goto _test_eof
	_test_eof198:
		cs = 198
		goto _test_eof
	_test_eof199:
		cs = 199
		goto _test_eof
	_test_eof200:
		cs = 200
		goto _test_eof
	_test_eof201:
		cs = 201
		goto _test_eof
	_test_eof202:
		cs = 202
		goto _test_eof
	_test_eof520:
		cs = 520
		goto _test_eof
	_test_eof521:
		cs = 521
		goto _test_eof
	_test_eof203:
		cs = 203
		goto _test_eof
	_test_eof204:
		cs = 204
		goto _test_eof
	_test_eof205:
		cs = 205
		goto _test_eof
	_test_eof206:
		cs = 206
		goto _test_eof
	_test_eof207:
		cs = 207
		goto _test_eof
	_test_eof208:
		cs = 208
		goto _test_eof
	_test_eof209:
		cs = 209
		goto _test_eof
	_test_eof210:
		cs = 210
		goto _test_eof
	_test_eof522:
		cs = 522
		goto _test_eof
	_test_eof211:
		cs = 211
		goto _test_eof
	_test_eof212:
		cs = 212
		goto _test_eof
	_test_eof523:
		cs = 523
		goto _test_eof
	_test_eof524:
		cs = 524
		goto _test_eof
	_test_eof213:
		cs = 213
		goto _test_eof
	_test_eof525:
		cs = 525
		goto _test_eof
	_test_eof526:
		cs = 526
		goto _test_eof
	_test_eof214:
		cs = 214
		goto _test_eof
	_test_eof215:
		cs = 215
		goto _test_eof
	_test_eof216:
		cs = 216
		goto _test_eof
	_test_eof217:
		cs = 217
		goto _test_eof
	_test_eof218:
		cs = 218
		goto _test_eof
	_test_eof219:
		cs = 219
		goto _test_eof
	_test_eof220:
		cs = 220
		goto _test_eof
	_test_eof221:
		cs = 221
		goto _test_eof
	_test_eof527:
		cs = 527
		goto _test_eof
	_test_eof528:
		cs = 528
		goto _test_eof
	_test_eof529:
		cs = 529
		goto _test_eof
	_test_eof222:
		cs = 222
		goto _test_eof
	_test_eof530:
		cs = 530
		goto _test_eof
	_test_eof223:
		cs = 223
		goto _test_eof
	_test_eof531:
		cs = 531
		goto _test_eof
	_test_eof532:
		cs = 532
		goto _test_eof
	_test_eof224:
		cs = 224
		goto _test_eof
	_test_eof533:
		cs = 533
		goto _test_eof
	_test_eof225:
		cs = 225
		goto _test_eof
	_test_eof226:
		cs = 226
		goto _test_eof
	_test_eof227:
		cs = 227
		goto _test_eof
	_test_eof228:
		cs = 228
		goto _test_eof
	_test_eof229:
		cs = 229
		goto _test_eof
	_test_eof230:
		cs = 230
		goto _test_eof
	_test_eof231:
		cs = 231
		goto _test_eof
	_test_eof232:
		cs = 232
		goto _test_eof
	_test_eof233:
		cs = 233
		goto _test_eof
	_test_eof234:
		cs = 234
		goto _test_eof
	_test_eof235:
		cs = 235
		goto _test_eof
	_test_eof236:
		cs = 236
		goto _test_eof
	_test_eof534:
		cs = 534
		goto _test_eof
	_test_eof237:
		cs = 237
		goto _test_eof
	_test_eof238:
		cs = 238
		goto _test_eof
	_test_eof239:
		cs = 239
		goto _test_eof
	_test_eof240:
		cs = 240
		goto _test_eof
	_test_eof241:
		cs = 241
		goto _test_eof
	_test_eof242:
		cs = 242
		goto _test_eof
	_test_eof243:
		cs = 243
		goto _test_eof
	_test_eof244:
		cs = 244
		goto _test_eof
	_test_eof245:
		cs = 245
		goto _test_eof
	_test_eof246:
		cs = 246
		goto _test_eof
	_test_eof247:
		cs = 247
		goto _test_eof
	_test_eof248:
		cs = 248
		goto _test_eof
	_test_eof249:
		cs = 249
		goto _test_eof
	_test_eof250:
		cs = 250
		goto _test_eof
	_test_eof251:
		cs = 251
		goto _test_eof
	_test_eof535:
		cs = 535
		goto _test_eof
	_test_eof536:
		cs = 536
		goto _test_eof
	_test_eof537:
		cs = 537
		goto _test_eof
	_test_eof252:
		cs = 252
		goto _test_eof
	_test_eof253:
		cs = 253
		goto _test_eof
	_test_eof254:
		cs = 254
		goto _test_eof
	_test_eof255:
		cs = 255
		goto _test_eof
	_test_eof538:
		cs = 538
		goto _test_eof
	_test_eof256:
		cs = 256
		goto _test_eof
	_test_eof539:
		cs = 539
		goto _test_eof
	_test_eof257:
		cs = 257
		goto _test_eof
	_test_eof540:
		cs = 540
		goto _test_eof
	_test_eof541:
		cs = 541
		goto _test_eof
	_test_eof258:
		cs = 258
		goto _test_eof
	_test_eof542:
		cs = 542
		goto _test_eof
	_test_eof259:
		cs = 259
		goto _test_eof
	_test_eof260:
		cs = 260
		goto _test_eof
	_test_eof261:
		cs = 261
		goto _test_eof
	_test_eof262:
		cs = 262
		goto _test_eof
	_test_eof263:
		cs = 263
		goto _test_eof
	_test_eof264:
		cs = 264
		goto _test_eof
	_test_eof265:
		cs = 265
		goto _test_eof
	_test_eof266:
		cs = 266
		goto _test_eof
	_test_eof267:
		cs = 267
		goto _test_eof
	_test_eof268:
		cs = 268
		goto _test_eof
	_test_eof269:
		cs = 269
		goto _test_eof
	_test_eof270:
		cs = 270
		goto _test_eof
	_test_eof543:
		cs = 543
		goto _test_eof
	_test_eof271:
		cs = 271
		goto _test_eof
	_test_eof272:
		cs = 272
		goto _test_eof
	_test_eof273:
		cs = 273
		goto _test_eof
	_test_eof274:
		cs = 274
		goto _test_eof
	_test_eof275:
		cs = 275
		goto _test_eof
	_test_eof276:
		cs = 276
		goto _test_eof
	_test_eof277:
		cs = 277
		goto _test_eof
	_test_eof278:
		cs = 278
		goto _test_eof
	_test_eof279:
		cs = 279
		goto _test_eof
	_test_eof280:
		cs = 280
		goto _test_eof
	_test_eof281:
		cs = 281
		goto _test_eof
	_test_eof282:
		cs = 282
		goto _test_eof
	_test_eof283:
		cs = 283
		goto _test_eof
	_test_eof284:
		cs = 284
		goto _test_eof
	_test_eof285:
		cs = 285
		goto _test_eof
	_test_eof544:
		cs = 544
		goto _test_eof
	_test_eof545:
		cs = 545
		goto _test_eof
	_test_eof546:
		cs = 546
		goto _test_eof
	_test_eof286:
		cs = 286
		goto _test_eof
	_test_eof287:
		cs = 287
		goto _test_eof
	_test_eof288:
		cs = 288
		goto _test_eof
	_test_eof289:
		cs = 289
		goto _test_eof
	_test_eof290:
		cs = 290
		goto _test_eof
	_test_eof291:
		cs = 291
		goto _test_eof
	_test_eof292:
		cs = 292
		goto _test_eof
	_test_eof547:
		cs = 547
		goto _test_eof
	_test_eof548:
		cs = 548
		goto _test_eof
	_test_eof293:
		cs = 293
		goto _test_eof
	_test_eof294:
		cs = 294
		goto _test_eof
	_test_eof295:
		cs = 295
		goto _test_eof
	_test_eof296:
		cs = 296
		goto _test_eof
	_test_eof297:
		cs = 297
		goto _test_eof
	_test_eof298:
		cs = 298
		goto _test_eof
	_test_eof299:
		cs = 299
		goto _test_eof
	_test_eof300:
		cs = 300
		goto _test_eof
	_test_eof301:
		cs = 301
		goto _test_eof
	_test_eof302:
		cs = 302
		goto _test_eof
	_test_eof549:
		cs = 549
		goto _test_eof
	_test_eof550:
		cs = 550
		goto _test_eof
	_test_eof303:
		cs = 303
		goto _test_eof
	_test_eof551:
		cs = 551
		goto _test_eof
	_test_eof552:
		cs = 552
		goto _test_eof
	_test_eof304:
		cs = 304
		goto _test_eof
	_test_eof305:
		cs = 305
		goto _test_eof
	_test_eof306:
		cs = 306
		goto _test_eof
	_test_eof307:
		cs = 307
		goto _test_eof
	_test_eof308:
		cs = 308
		goto _test_eof
	_test_eof309:
		cs = 309
		goto _test_eof
	_test_eof310:
		cs = 310
		goto _test_eof
	_test_eof311:
		cs = 311
		goto _test_eof
	_test_eof553:
		cs = 553
		goto _test_eof
	_test_eof554:
		cs = 554
		goto _test_eof
	_test_eof555:
		cs = 555
		goto _test_eof
	_test_eof312:
		cs = 312
		goto _test_eof
	_test_eof556:
		cs = 556
		goto _test_eof
	_test_eof313:
		cs = 313
		goto _test_eof
	_test_eof557:
		cs = 557
		goto _test_eof
	_test_eof558:
		cs = 558
		goto _test_eof
	_test_eof314:
		cs = 314
		goto _test_eof
	_test_eof559:
		cs = 559
		goto _test_eof
	_test_eof315:
		cs = 315
		goto _test_eof
	_test_eof316:
		cs = 316
		goto _test_eof
	_test_eof317:
		cs = 317
		goto _test_eof
	_test_eof318:
		cs = 318
		goto _test_eof
	_test_eof319:
		cs = 319
		goto _test_eof
	_test_eof320:
		cs = 320
		goto _test_eof
	_test_eof321:
		cs = 321
		goto _test_eof
	_test_eof322:
		cs = 322
		goto _test_eof
	_test_eof323:
		cs = 323
		goto _test_eof
	_test_eof324:
		cs = 324
		goto _test_eof
	_test_eof325:
		cs = 325
		goto _test_eof
	_test_eof326:
		cs = 326
		goto _test_eof
	_test_eof560:
		cs = 560
		goto _test_eof
	_test_eof327:
		cs = 327
		goto _test_eof
	_test_eof328:
		cs = 328
		goto _test_eof
	_test_eof329:
		cs = 329
		goto _test_eof
	_test_eof330:
		cs = 330
		goto _test_eof
	_test_eof331:
		cs = 331
		goto _test_eof
	_test_eof332:
		cs = 332
		goto _test_eof
	_test_eof333:
		cs = 333
		goto _test_eof
	_test_eof334:
		cs = 334
		goto _test_eof
	_test_eof335:
		cs = 335
		goto _test_eof
	_test_eof336:
		cs = 336
		goto _test_eof
	_test_eof337:
		cs = 337
		goto _test_eof
	_test_eof338:
		cs = 338
		goto _test_eof
	_test_eof339:
		cs = 339
		goto _test_eof
	_test_eof340:
		cs = 340
		goto _test_eof
	_test_eof341:
		cs = 341
		goto _test_eof
	_test_eof561:
		cs = 561
		goto _test_eof
	_test_eof562:
		cs = 562
		goto _test_eof
	_test_eof563:
		cs = 563
		goto _test_eof
	_test_eof342:
		cs = 342
		goto _test_eof
	_test_eof343:
		cs = 343
		goto _test_eof
	_test_eof344:
		cs = 344
		goto _test_eof
	_test_eof345:
		cs = 345
		goto _test_eof
	_test_eof564:
		cs = 564
		goto _test_eof
	_test_eof565:
		cs = 565
		goto _test_eof
	_test_eof346:
		cs = 346
		goto _test_eof
	_test_eof566:
		cs = 566
		goto _test_eof
	_test_eof347:
		cs = 347
		goto _test_eof
	_test_eof567:
		cs = 567
		goto _test_eof
	_test_eof568:
		cs = 568
		goto _test_eof
	_test_eof348:
		cs = 348
		goto _test_eof
	_test_eof569:
		cs = 569
		goto _test_eof
	_test_eof349:
		cs = 349
		goto _test_eof
	_test_eof350:
		cs = 350
		goto _test_eof
	_test_eof351:
		cs = 351
		goto _test_eof
	_test_eof352:
		cs = 352
		goto _test_eof
	_test_eof353:
		cs = 353
		goto _test_eof
	_test_eof354:
		cs = 354
		goto _test_eof
	_test_eof355:
		cs = 355
		goto _test_eof
	_test_eof356:
		cs = 356
		goto _test_eof
	_test_eof357:
		cs = 357
		goto _test_eof
	_test_eof358:
		cs = 358
		goto _test_eof
	_test_eof359:
		cs = 359
		goto _test_eof
	_test_eof360:
		cs = 360
		goto _test_eof
	_test_eof570:
		cs = 570
		goto _test_eof
	_test_eof361:
		cs = 361
		goto _test_eof
	_test_eof362:
		cs = 362
		goto _test_eof
	_test_eof363:
		cs = 363
		goto _test_eof
	_test_eof364:
		cs = 364
		goto _test_eof
	_test_eof365:
		cs = 365
		goto _test_eof
	_test_eof366:
		cs = 366
		goto _test_eof
	_test_eof367:
		cs = 367
		goto _test_eof
	_test_eof368:
		cs = 368
		goto _test_eof
	_test_eof369:
		cs = 369
		goto _test_eof
	_test_eof370:
		cs = 370
		goto _test_eof
	_test_eof371:
		cs = 371
		goto _test_eof
	_test_eof372:
		cs = 372
		goto _test_eof
	_test_eof373:
		cs = 373
		goto _test_eof
	_test_eof374:
		cs = 374
		goto _test_eof
	_test_eof375:
		cs = 375
		goto _test_eof
	_test_eof571:
		cs = 571
		goto _test_eof
	_test_eof572:
		cs = 572
		goto _test_eof
	_test_eof573:
		cs = 573
		goto _test_eof
	_test_eof376:
		cs = 376
		goto _test_eof
	_test_eof377:
		cs = 377
		goto _test_eof
	_test_eof378:
		cs = 378
		goto _test_eof
	_test_eof379:
		cs = 379
		goto _test_eof
	_test_eof380:
		cs = 380
		goto _test_eof
	_test_eof381:
		cs = 381
		goto _test_eof
	_test_eof382:
		cs = 382
		goto _test_eof
	_test_eof383:
		cs = 383
		goto _test_eof
	_test_eof384:
		cs = 384
		goto _test_eof
	_test_eof385:
		cs = 385
		goto _test_eof
	_test_eof386:
		cs = 386
		goto _test_eof
	_test_eof387:
		cs = 387
		goto _test_eof
	_test_eof388:
		cs = 388
		goto _test_eof
	_test_eof389:
		cs = 389
		goto _test_eof
	_test_eof390:
		cs = 390
		goto _test_eof
	_test_eof391:
		cs = 391
		goto _test_eof
	_test_eof392:
		cs = 392
		goto _test_eof
	_test_eof393:
		cs = 393
		goto _test_eof
	_test_eof394:
		cs = 394
		goto _test_eof
	_test_eof395:
		cs = 395
		goto _test_eof
	_test_eof396:
		cs = 396
		goto _test_eof
	_test_eof397:
		cs = 397
		goto _test_eof
	_test_eof398:
		cs = 398
		goto _test_eof
	_test_eof399:
		cs = 399
		goto _test_eof
	_test_eof400:
		cs = 400
		goto _test_eof
	_test_eof401:
		cs = 401
		goto _test_eof
	_test_eof402:
		cs = 402
		goto _test_eof
	_test_eof403:
		cs = 403
		goto _test_eof
	_test_eof404:
		cs = 404
		goto _test_eof
	_test_eof405:
		cs = 405
		goto _test_eof
	_test_eof574:
		cs = 574
		goto _test_eof
	_test_eof406:
		cs = 406
		goto _test_eof
	_test_eof407:
		cs = 407
		goto _test_eof
	_test_eof408:
		cs = 408
		goto _test_eof
	_test_eof409:
		cs = 409
		goto _test_eof
	_test_eof410:
		cs = 410
		goto _test_eof
	_test_eof575:
		cs = 575
		goto _test_eof
	_test_eof411:
		cs = 411
		goto _test_eof
	_test_eof412:
		cs = 412
		goto _test_eof
	_test_eof413:
		cs = 413
		goto _test_eof
	_test_eof414:
		cs = 414
		goto _test_eof
	_test_eof415:
		cs = 415
		goto _test_eof
	_test_eof416:
		cs = 416
		goto _test_eof
	_test_eof417:
		cs = 417
		goto _test_eof
	_test_eof418:
		cs = 418
		goto _test_eof
	_test_eof419:
		cs = 419
		goto _test_eof
	_test_eof420:
		cs = 420
		goto _test_eof
	_test_eof421:
		cs = 421
		goto _test_eof
	_test_eof422:
		cs = 422
		goto _test_eof
	_test_eof423:
		cs = 423
		goto _test_eof
	_test_eof424:
		cs = 424
		goto _test_eof
	_test_eof425:
		cs = 425
		goto _test_eof
	_test_eof576:
		cs = 576
		goto _test_eof
	_test_eof426:
		cs = 426
		goto _test_eof
	_test_eof427:
		cs = 427
		goto _test_eof
	_test_eof428:
		cs = 428
		goto _test_eof
	_test_eof429:
		cs = 429
		goto _test_eof
	_test_eof430:
		cs = 430
		goto _test_eof
	_test_eof431:
		cs = 431
		goto _test_eof
	_test_eof432:
		cs = 432
		goto _test_eof
	_test_eof433:
		cs = 433
		goto _test_eof
	_test_eof434:
		cs = 434
		goto _test_eof
	_test_eof435:
		cs = 435
		goto _test_eof
	_test_eof436:
		cs = 436
		goto _test_eof
	_test_eof437:
		cs = 437
		goto _test_eof
	_test_eof438:
		cs = 438
		goto _test_eof
	_test_eof439:
		cs = 439
		goto _test_eof
	_test_eof440:
		cs = 440
		goto _test_eof
	_test_eof441:
		cs = 441
		goto _test_eof
	_test_eof442:
		cs = 442
		goto _test_eof
	_test_eof443:
		cs = 443
		goto _test_eof
	_test_eof444:
		cs = 444
		goto _test_eof
	_test_eof445:
		cs = 445
		goto _test_eof
	_test_eof446:
		cs = 446
		goto _test_eof
	_test_eof447:
		cs = 447
		goto _test_eof
	_test_eof448:
		cs = 448
		goto _test_eof
	_test_eof449:
		cs = 449
		goto _test_eof
	_test_eof450:
		cs = 450
		goto _test_eof
	_test_eof451:
		cs = 451
		goto _test_eof
	_test_eof452:
		cs = 452
		goto _test_eof
	_test_eof453:
		cs = 453
		goto _test_eof
	_test_eof454:
		cs = 454
		goto _test_eof
	_test_eof455:
		cs = 455
		goto _test_eof
	_test_eof456:
		cs = 456
		goto _test_eof
	_test_eof457:
		cs = 457
		goto _test_eof
	_test_eof458:
		cs = 458
		goto _test_eof
	_test_eof459:
		cs = 459
		goto _test_eof
	_test_eof460:
		cs = 460
		goto _test_eof
	_test_eof461:
		cs = 461
		goto _test_eof
	_test_eof462:
		cs = 462
		goto _test_eof
	_test_eof463:
		cs = 463
		goto _test_eof
	_test_eof464:
		cs = 464
		goto _test_eof
	_test_eof465:
		cs = 465
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 489, 500, 508, 516, 534, 543, 560, 570:
//line query/tokeniser.rl:204
				commit(ttEquivalenceTest)
			case 575:
//line query/tokeniser.rl:123
				commit(ttAnyDecl)
//line query/tokeniser.rl:152
				commit(ttEventClause)
			case 466:
//line query/tokeniser.rl:135
				commit(ttNegatedDecl)
//line query/tokeniser.rl:152
				commit(ttEventClause)
			case 576:
//line query/tokeniser.rl:146
				commit(ttSeqDecl)
//line query/tokeniser.rl:152
				commit(ttEventClause)
			case 476, 477, 479, 484, 493, 494, 523, 528, 538, 547, 555, 564, 565:
//line query/tokeniser.rl:186
				commit(ttStringLiteral)
//line query/tokeniser.rl:218
				commit(ttPredicate)
			case 469, 471, 473, 475, 480, 482, 520, 522, 525, 527, 549, 551, 553:
//line query/tokeniser.rl:194
				commit(ttStringLiteral)
//line query/tokeniser.rl:218
				commit(ttPredicate)
			case 485, 486, 496, 497, 504, 505, 512, 513, 530, 531, 539, 540, 556, 557, 566, 567:
//line query/tokeniser.rl:177
				setText(ttNumericLiteral)
//line query/tokeniser.rl:178
				commit(ttNumericLiteral)
//line query/tokeniser.rl:218
				commit(ttPredicate)
			case 488, 499, 507, 515, 533, 542, 559, 569:
//line query/tokeniser.rl:211
				setText(ttAttributeSelector)
//line query/tokeniser.rl:212
				commit(ttAttributeSelector)
//line query/tokeniser.rl:218
				commit(ttPredicate)
			case 490, 492, 501, 503, 509, 511, 517, 519, 535, 537, 544, 546, 561, 563, 571, 573:
//line query/tokeniser.rl:228
				setText(ttDuration)
//line query/tokeniser.rl:229
				commit(ttDuration)
//line query/tokeniser.rl:233
				commit(ttWithinClause)
			case 574:
//line query/tokeniser.rl:107
				setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
				commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
				commit(ttEventDecl)
//line query/tokeniser.rl:152
				commit(ttEventClause)
			case 487, 498, 506, 514, 532, 541, 558, 568:
//line query/tokeniser.rl:210
				propose(ttAttributeSelector)
//line query/tokeniser.rl:211
				setText(ttAttributeSelector)
//line query/tokeniser.rl:212
				commit(ttAttributeSelector)
//line query/tokeniser.rl:218
				commit(ttPredicate)
//line query/tokeniser.go:21898
			}
		}

	_out:
		{
		}
	}

//line query/tokeniser.rl:245

	if cs < sase_first_final {
		if p == pe {
			return nil, fmt.Errorf("Unexpected EOF")
		} else {
			end := p + 30
			if end > len(data) {
				end = len(data)
			}
			return nil, fmt.Errorf("Error at position %d (\"%s\"…)", p, data[p:end])
		}
	}

	tokens = tokens[1:] // Remove root node
	return tokens, nil
}
