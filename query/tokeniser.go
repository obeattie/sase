//line query/tokeniser.rl:1
package query

import (
	"fmt"
)

//line query/tokeniser.rl:8
//line query/tokeniser.go:13
const sase_start int = 1
const sase_first_final int = 434
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
		case 434:
			goto st_case_434
		case 435:
			goto st_case_435
		case 436:
			goto st_case_436
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
		case 437:
			goto st_case_437
		case 438:
			goto st_case_438
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
		case 439:
			goto st_case_439
		case 440:
			goto st_case_440
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
		case 441:
			goto st_case_441
		case 442:
			goto st_case_442
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 443:
			goto st_case_443
		case 45:
			goto st_case_45
		case 444:
			goto st_case_444
		case 46:
			goto st_case_46
		case 445:
			goto st_case_445
		case 446:
			goto st_case_446
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
		case 447:
			goto st_case_447
		case 56:
			goto st_case_56
		case 448:
			goto st_case_448
		case 449:
			goto st_case_449
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
		case 450:
			goto st_case_450
		case 451:
			goto st_case_451
		case 452:
			goto st_case_452
		case 65:
			goto st_case_65
		case 453:
			goto st_case_453
		case 66:
			goto st_case_66
		case 454:
			goto st_case_454
		case 455:
			goto st_case_455
		case 67:
			goto st_case_67
		case 456:
			goto st_case_456
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
		case 457:
			goto st_case_457
		case 458:
			goto st_case_458
		case 459:
			goto st_case_459
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 460:
			goto st_case_460
		case 461:
			goto st_case_461
		case 462:
			goto st_case_462
		case 95:
			goto st_case_95
		case 463:
			goto st_case_463
		case 96:
			goto st_case_96
		case 464:
			goto st_case_464
		case 465:
			goto st_case_465
		case 97:
			goto st_case_97
		case 466:
			goto st_case_466
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
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
		case 467:
			goto st_case_467
		case 468:
			goto st_case_468
		case 469:
			goto st_case_469
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
		case 470:
			goto st_case_470
		case 126:
			goto st_case_126
		case 471:
			goto st_case_471
		case 472:
			goto st_case_472
		case 127:
			goto st_case_127
		case 473:
			goto st_case_473
		case 128:
			goto st_case_128
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
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
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
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 474:
			goto st_case_474
		case 475:
			goto st_case_475
		case 476:
			goto st_case_476
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
		case 477:
			goto st_case_477
		case 157:
			goto st_case_157
		case 478:
			goto st_case_478
		case 479:
			goto st_case_479
		case 158:
			goto st_case_158
		case 480:
			goto st_case_480
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
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
		case 169:
			goto st_case_169
		case 170:
			goto st_case_170
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
		case 481:
			goto st_case_481
		case 482:
			goto st_case_482
		case 483:
			goto st_case_483
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 484:
			goto st_case_484
		case 485:
			goto st_case_485
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
		case 486:
			goto st_case_486
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 487:
			goto st_case_487
		case 488:
			goto st_case_488
		case 197:
			goto st_case_197
		case 489:
			goto st_case_489
		case 490:
			goto st_case_490
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
		case 203:
			goto st_case_203
		case 204:
			goto st_case_204
		case 205:
			goto st_case_205
		case 491:
			goto st_case_491
		case 492:
			goto st_case_492
		case 493:
			goto st_case_493
		case 206:
			goto st_case_206
		case 494:
			goto st_case_494
		case 207:
			goto st_case_207
		case 495:
			goto st_case_495
		case 496:
			goto st_case_496
		case 208:
			goto st_case_208
		case 497:
			goto st_case_497
		case 209:
			goto st_case_209
		case 210:
			goto st_case_210
		case 211:
			goto st_case_211
		case 212:
			goto st_case_212
		case 213:
			goto st_case_213
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
		case 222:
			goto st_case_222
		case 223:
			goto st_case_223
		case 224:
			goto st_case_224
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
		case 498:
			goto st_case_498
		case 499:
			goto st_case_499
		case 500:
			goto st_case_500
		case 232:
			goto st_case_232
		case 233:
			goto st_case_233
		case 234:
			goto st_case_234
		case 235:
			goto st_case_235
		case 501:
			goto st_case_501
		case 236:
			goto st_case_236
		case 502:
			goto st_case_502
		case 237:
			goto st_case_237
		case 503:
			goto st_case_503
		case 504:
			goto st_case_504
		case 238:
			goto st_case_238
		case 505:
			goto st_case_505
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
		case 252:
			goto st_case_252
		case 253:
			goto st_case_253
		case 254:
			goto st_case_254
		case 255:
			goto st_case_255
		case 256:
			goto st_case_256
		case 257:
			goto st_case_257
		case 258:
			goto st_case_258
		case 259:
			goto st_case_259
		case 260:
			goto st_case_260
		case 261:
			goto st_case_261
		case 506:
			goto st_case_506
		case 507:
			goto st_case_507
		case 508:
			goto st_case_508
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
		case 509:
			goto st_case_509
		case 510:
			goto st_case_510
		case 269:
			goto st_case_269
		case 270:
			goto st_case_270
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
		case 511:
			goto st_case_511
		case 512:
			goto st_case_512
		case 279:
			goto st_case_279
		case 513:
			goto st_case_513
		case 514:
			goto st_case_514
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
		case 286:
			goto st_case_286
		case 287:
			goto st_case_287
		case 515:
			goto st_case_515
		case 516:
			goto st_case_516
		case 517:
			goto st_case_517
		case 288:
			goto st_case_288
		case 518:
			goto st_case_518
		case 289:
			goto st_case_289
		case 519:
			goto st_case_519
		case 520:
			goto st_case_520
		case 290:
			goto st_case_290
		case 521:
			goto st_case_521
		case 291:
			goto st_case_291
		case 292:
			goto st_case_292
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
		case 303:
			goto st_case_303
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
		case 312:
			goto st_case_312
		case 313:
			goto st_case_313
		case 522:
			goto st_case_522
		case 523:
			goto st_case_523
		case 524:
			goto st_case_524
		case 314:
			goto st_case_314
		case 315:
			goto st_case_315
		case 316:
			goto st_case_316
		case 317:
			goto st_case_317
		case 525:
			goto st_case_525
		case 526:
			goto st_case_526
		case 318:
			goto st_case_318
		case 527:
			goto st_case_527
		case 319:
			goto st_case_319
		case 528:
			goto st_case_528
		case 529:
			goto st_case_529
		case 320:
			goto st_case_320
		case 530:
			goto st_case_530
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
		case 342:
			goto st_case_342
		case 343:
			goto st_case_343
		case 531:
			goto st_case_531
		case 532:
			goto st_case_532
		case 533:
			goto st_case_533
		case 344:
			goto st_case_344
		case 345:
			goto st_case_345
		case 346:
			goto st_case_346
		case 347:
			goto st_case_347
		case 348:
			goto st_case_348
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
		case 534:
			goto st_case_534
		case 374:
			goto st_case_374
		case 375:
			goto st_case_375
		case 376:
			goto st_case_376
		case 377:
			goto st_case_377
		case 378:
			goto st_case_378
		case 535:
			goto st_case_535
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
		case 536:
			goto st_case_536
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
//line query/tokeniser.go:1317
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
			goto st434
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
	tr846:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st434
	tr857:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st434
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
//line query/tokeniser.go:1384
		switch data[p] {
		case 32:
			goto tr971
		case 59:
			goto tr972
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr971
		}
		goto st0
	tr971:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st435
	tr1211:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st435
	tr1214:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st435
	tr1216:
//line query/tokeniser.rl:146
		commit(ttSeqDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st435
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
//line query/tokeniser.go:1428
		switch data[p] {
		case 32:
			goto st435
		case 59:
			goto st436
		case 87:
			goto st11
		case 119:
			goto st11
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st435
		}
		goto st0
	tr972:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st436
	tr977:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st436
	tr1002:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st436
	tr1079:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st436
	tr1083:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st436
	tr1086:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st436
	tr1089:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st436
	tr1213:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st436
	tr1215:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st436
	tr1217:
//line query/tokeniser.rl:146
		commit(ttSeqDecl)
//line query/tokeniser.rl:152
		commit(ttEventClause)
		goto st436
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
//line query/tokeniser.go:1522
		if data[p] == 32 {
			goto st436
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st436
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
			goto st172
		case 104:
			goto st12
		case 105:
			goto st172
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
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line query/tokeniser.go:1629
		switch data[p] {
		case 32:
			goto tr25
		case 33:
			goto tr26
		case 46:
			goto tr27
		case 60:
			goto tr29
		case 61:
			goto tr30
		case 62:
			goto tr31
		case 95:
			goto st17
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
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
	tr25:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st18
	tr834:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line query/tokeniser.go:1683
		switch data[p] {
		case 32:
			goto st18
		case 33:
			goto tr33
		case 60:
			goto tr34
		case 61:
			goto tr35
		case 62:
			goto tr36
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st18
		}
		goto st0
	tr26:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st19
	tr33:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st19
	tr835:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line query/tokeniser.go:1727
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
			goto tr38
		case 34:
			goto tr39
		case 39:
			goto tr40
		case 43:
			goto tr41
		case 45:
			goto tr41
		case 95:
			goto tr43
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr43
				}
			case data[p] >= 65:
				goto tr43
			}
		default:
			goto tr42
		}
		goto st0
	tr38:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st21
	tr815:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st21
	tr819:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st21
	tr823:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st21
	tr826:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st21
	tr830:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line query/tokeniser.go:1798
		switch data[p] {
		case 32:
			goto st21
		case 34:
			goto tr45
		case 39:
			goto tr46
		case 43:
			goto tr47
		case 45:
			goto tr47
		case 95:
			goto tr49
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
					goto tr49
				}
			case data[p] >= 65:
				goto tr49
			}
		default:
			goto tr48
		}
		goto st0
	tr39:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr45:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr816:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr820:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr824:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr827:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	tr831:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line query/tokeniser.go:1876
		switch data[p] {
		case 34:
			goto tr51
		case 92:
			goto tr52
		}
		goto tr50
	tr50:
//line query/tokeniser.rl:87
		mark = p
		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line query/tokeniser.go:1893
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		goto st23
	tr51:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st437
	tr54:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st437
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
//line query/tokeniser.go:1916
		switch data[p] {
		case 32:
			goto tr976
		case 59:
			goto tr977
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr976
		}
		goto st0
	tr976:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st438
	tr1001:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st438
	tr1077:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st438
	tr1080:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st438
	tr1084:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st438
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
//line query/tokeniser.go:1970
		switch data[p] {
		case 32:
			goto st438
		case 38:
			goto tr979
		case 59:
			goto st436
		case 65:
			goto tr980
		case 79:
			goto tr981
		case 87:
			goto st171
		case 94:
			goto tr983
		case 97:
			goto tr980
		case 111:
			goto tr981
		case 119:
			goto st171
		case 124:
			goto tr984
		case 226:
			goto tr985
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st438
		}
		goto st0
	tr979:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line query/tokeniser.go:2010
		if data[p] == 38 {
			goto st25
		}
		goto st0
	tr983:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line query/tokeniser.go:2024
		if data[p] == 32 {
			goto tr57
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr57
		}
		goto st0
	tr57:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st26
	tr417:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line query/tokeniser.go:2045
		switch data[p] {
		case 32:
			goto st26
		case 95:
			goto tr59
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr59
			}
		default:
			goto tr59
		}
		goto st0
	tr59:
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line query/tokeniser.go:2076
		switch data[p] {
		case 32:
			goto tr60
		case 33:
			goto tr61
		case 46:
			goto tr62
		case 60:
			goto tr64
		case 61:
			goto tr65
		case 62:
			goto tr66
		case 95:
			goto st27
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr60
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
	tr60:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st28
	tr409:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line query/tokeniser.go:2130
		switch data[p] {
		case 32:
			goto st28
		case 33:
			goto tr68
		case 60:
			goto tr69
		case 61:
			goto tr70
		case 62:
			goto tr71
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st28
		}
		goto st0
	tr61:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st29
	tr68:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st29
	tr410:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line query/tokeniser.go:2174
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
			goto tr73
		case 34:
			goto tr74
		case 39:
			goto tr75
		case 43:
			goto tr41
		case 45:
			goto tr41
		case 95:
			goto tr43
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr73
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr43
				}
			case data[p] >= 65:
				goto tr43
			}
		default:
			goto tr42
		}
		goto st0
	tr73:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st31
	tr375:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st31
	tr382:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st31
	tr389:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st31
	tr395:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st31
	tr402:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line query/tokeniser.go:2245
		switch data[p] {
		case 32:
			goto st31
		case 34:
			goto tr77
		case 39:
			goto tr78
		case 43:
			goto tr47
		case 45:
			goto tr47
		case 95:
			goto tr49
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
					goto tr49
				}
			case data[p] >= 65:
				goto tr49
			}
		default:
			goto tr48
		}
		goto st0
	tr74:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr77:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr376:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr383:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr390:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr396:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	tr403:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line query/tokeniser.go:2323
		switch data[p] {
		case 34:
			goto tr51
		case 92:
			goto tr80
		}
		goto tr79
	tr79:
//line query/tokeniser.rl:87
		mark = p
		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line query/tokeniser.go:2340
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		goto st33
	tr80:
//line query/tokeniser.rl:87
		mark = p
		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line query/tokeniser.go:2357
		switch data[p] {
		case 34:
			goto tr83
		case 92:
			goto st34
		}
		goto st33
	tr83:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st439
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
//line query/tokeniser.go:2374
		switch data[p] {
		case 32:
			goto tr986
		case 34:
			goto tr54
		case 59:
			goto tr987
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr986
		}
		goto st33
	tr986:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st440
	tr997:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st440
	tr1014:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st440
	tr1063:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st440
	tr1066:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st440
	tr1070:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st440
	tr1102:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
//line query/tokeniser.go:2448
		switch data[p] {
		case 32:
			goto st440
		case 34:
			goto tr54
		case 38:
			goto tr989
		case 59:
			goto st442
		case 65:
			goto tr991
		case 79:
			goto tr992
		case 87:
			goto st140
		case 92:
			goto st34
		case 94:
			goto tr994
		case 97:
			goto tr991
		case 111:
			goto tr992
		case 119:
			goto st140
		case 124:
			goto tr995
		case 226:
			goto tr996
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st440
		}
		goto st33
	tr989:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line query/tokeniser.go:2492
		switch data[p] {
		case 34:
			goto tr54
		case 38:
			goto st36
		case 92:
			goto st34
		}
		goto st33
	tr994:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line query/tokeniser.go:2511
		switch data[p] {
		case 32:
			goto tr85
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr85
		}
		goto st33
	tr85:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st37
	tr353:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line query/tokeniser.go:2537
		switch data[p] {
		case 32:
			goto st37
		case 34:
			goto tr54
		case 92:
			goto st34
		case 95:
			goto tr87
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr87
			}
		default:
			goto tr87
		}
		goto st33
	tr87:
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line query/tokeniser.go:2572
		switch data[p] {
		case 32:
			goto tr88
		case 33:
			goto tr89
		case 34:
			goto tr54
		case 46:
			goto tr90
		case 60:
			goto tr92
		case 61:
			goto tr93
		case 62:
			goto tr94
		case 92:
			goto st34
		case 95:
			goto st38
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr88
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
	tr88:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st39
	tr345:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line query/tokeniser.go:2630
		switch data[p] {
		case 32:
			goto st39
		case 33:
			goto tr96
		case 34:
			goto tr54
		case 60:
			goto tr97
		case 61:
			goto tr98
		case 62:
			goto tr99
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st39
		}
		goto st33
	tr89:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st40
	tr96:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st40
	tr346:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line query/tokeniser.go:2678
		switch data[p] {
		case 34:
			goto tr54
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
			goto tr101
		case 34:
			goto tr102
		case 39:
			goto tr103
		case 43:
			goto tr104
		case 45:
			goto tr104
		case 92:
			goto st34
		case 95:
			goto tr106
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr101
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr106
				}
			case data[p] >= 65:
				goto tr106
			}
		default:
			goto tr105
		}
		goto st33
	tr101:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st42
	tr311:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st42
	tr318:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st42
	tr325:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st42
	tr331:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st42
	tr338:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line query/tokeniser.go:2756
		switch data[p] {
		case 32:
			goto st42
		case 34:
			goto tr108
		case 39:
			goto tr109
		case 43:
			goto tr110
		case 45:
			goto tr110
		case 92:
			goto st34
		case 95:
			goto tr112
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
					goto tr112
				}
			case data[p] >= 65:
				goto tr112
			}
		default:
			goto tr111
		}
		goto st33
	tr102:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st441
	tr108:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st441
	tr312:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st441
	tr319:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st441
	tr326:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st441
	tr332:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st441
	tr339:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
//line query/tokeniser.go:2850
		switch data[p] {
		case 32:
			goto tr997
		case 34:
			goto tr51
		case 59:
			goto tr998
		case 92:
			goto tr80
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr997
		}
		goto tr79
	tr987:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st442
	tr998:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st442
	tr1015:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st442
	tr1065:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st442
	tr1069:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st442
	tr1072:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st442
	tr1075:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st442
	tr1103:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st442
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
//line query/tokeniser.go:2932
		switch data[p] {
		case 32:
			goto st442
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st442
		}
		goto st33
	tr103:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr109:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr313:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr320:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr327:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr333:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	tr340:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line query/tokeniser.go:2990
		switch data[p] {
		case 34:
			goto tr114
		case 39:
			goto tr115
		case 92:
			goto tr116
		}
		goto tr113
	tr113:
//line query/tokeniser.rl:87
		mark = p
		goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line query/tokeniser.go:3009
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		goto st44
	tr114:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st443
	tr118:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st443
	tr463:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:87
		mark = p
		goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line query/tokeniser.go:3040
		switch data[p] {
		case 32:
			goto tr999
		case 39:
			goto tr122
		case 59:
			goto tr1000
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr999
		}
		goto st45
	tr369:
//line query/tokeniser.rl:87
		mark = p
		goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line query/tokeniser.go:3064
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		goto st45
	tr370:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st444
	tr122:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st444
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
//line query/tokeniser.go:3087
		switch data[p] {
		case 32:
			goto tr1001
		case 59:
			goto tr1002
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1001
		}
		goto st0
	tr371:
//line query/tokeniser.rl:87
		mark = p
		goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line query/tokeniser.go:3107
		switch data[p] {
		case 39:
			goto tr124
		case 92:
			goto st46
		}
		goto st45
	tr124:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
//line query/tokeniser.go:3124
		switch data[p] {
		case 32:
			goto tr1003
		case 39:
			goto tr122
		case 59:
			goto tr1004
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1003
		}
		goto st45
	tr999:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st446
	tr1003:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st446
	tr1047:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st446
	tr1049:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st446
	tr1052:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st446
	tr1056:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st446
	tr1195:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st446
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
//line query/tokeniser.go:3198
		switch data[p] {
		case 32:
			goto st446
		case 38:
			goto tr1006
		case 39:
			goto tr122
		case 59:
			goto st462
		case 65:
			goto tr1008
		case 79:
			goto tr1009
		case 87:
			goto st110
		case 92:
			goto st46
		case 94:
			goto tr1011
		case 97:
			goto tr1008
		case 111:
			goto tr1009
		case 119:
			goto st110
		case 124:
			goto tr1012
		case 226:
			goto tr1013
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st446
		}
		goto st45
	tr1006:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line query/tokeniser.go:3242
		switch data[p] {
		case 38:
			goto st48
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		goto st45
	tr1011:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line query/tokeniser.go:3261
		switch data[p] {
		case 32:
			goto tr126
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr126
		}
		goto st45
	tr126:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st49
	tr292:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line query/tokeniser.go:3287
		switch data[p] {
		case 32:
			goto st49
		case 39:
			goto tr122
		case 92:
			goto st46
		case 95:
			goto tr128
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st49
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr128
			}
		default:
			goto tr128
		}
		goto st45
	tr128:
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line query/tokeniser.go:3322
		switch data[p] {
		case 32:
			goto tr129
		case 33:
			goto tr130
		case 39:
			goto tr122
		case 46:
			goto tr131
		case 60:
			goto tr133
		case 61:
			goto tr134
		case 62:
			goto tr135
		case 92:
			goto st46
		case 95:
			goto st50
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr129
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
	tr129:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st51
	tr284:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line query/tokeniser.go:3380
		switch data[p] {
		case 32:
			goto st51
		case 33:
			goto tr137
		case 39:
			goto tr122
		case 60:
			goto tr138
		case 61:
			goto tr139
		case 62:
			goto tr140
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st51
		}
		goto st45
	tr130:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st52
	tr137:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st52
	tr285:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line query/tokeniser.go:3428
		switch data[p] {
		case 39:
			goto tr122
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
			goto tr142
		case 34:
			goto tr143
		case 39:
			goto tr144
		case 43:
			goto tr145
		case 45:
			goto tr145
		case 92:
			goto st46
		case 95:
			goto tr147
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr142
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr147
				}
			case data[p] >= 65:
				goto tr147
			}
		default:
			goto tr146
		}
		goto st45
	tr142:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st54
	tr250:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st54
	tr257:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st54
	tr264:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st54
	tr270:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st54
	tr277:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line query/tokeniser.go:3506
		switch data[p] {
		case 32:
			goto st54
		case 34:
			goto tr149
		case 39:
			goto tr150
		case 43:
			goto tr151
		case 45:
			goto tr151
		case 92:
			goto st46
		case 95:
			goto tr153
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
					goto tr153
				}
			case data[p] >= 65:
				goto tr153
			}
		default:
			goto tr152
		}
		goto st45
	tr143:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr149:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr251:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr258:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr265:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr271:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	tr278:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line query/tokeniser.go:3586
		switch data[p] {
		case 34:
			goto tr114
		case 39:
			goto tr154
		case 92:
			goto tr116
		}
		goto tr113
	tr115:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st447
	tr119:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st447
	tr154:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:87
		mark = p
		goto st447
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
//line query/tokeniser.go:3617
		switch data[p] {
		case 32:
			goto tr1014
		case 34:
			goto tr54
		case 59:
			goto tr1015
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1014
		}
		goto st33
	tr116:
//line query/tokeniser.rl:87
		mark = p
		goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line query/tokeniser.go:3641
		switch data[p] {
		case 34:
			goto tr155
		case 39:
			goto tr156
		case 92:
			goto st56
		}
		goto st44
	tr155:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st448
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
//line query/tokeniser.go:3660
		switch data[p] {
		case 32:
			goto tr1016
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 59:
			goto tr1017
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1016
		}
		goto st44
	tr1016:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st449
	tr1027:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st449
	tr1045:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st449
	tr1029:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st449
	tr1031:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st449
	tr1034:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st449
	tr1038:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st449
	tr1117:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st449
	tr1177:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st449
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
//line query/tokeniser.go:3752
		switch data[p] {
		case 32:
			goto st449
		case 34:
			goto tr118
		case 38:
			goto tr1019
		case 39:
			goto tr119
		case 59:
			goto st451
		case 65:
			goto tr1021
		case 79:
			goto tr1022
		case 87:
			goto st80
		case 92:
			goto st56
		case 94:
			goto tr1024
		case 97:
			goto tr1021
		case 111:
			goto tr1022
		case 119:
			goto st80
		case 124:
			goto tr1025
		case 226:
			goto tr1026
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st449
		}
		goto st44
	tr1019:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line query/tokeniser.go:3798
		switch data[p] {
		case 34:
			goto tr118
		case 38:
			goto st58
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		goto st44
	tr1024:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line query/tokeniser.go:3819
		switch data[p] {
		case 32:
			goto tr158
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr158
		}
		goto st44
	tr158:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st59
	tr231:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line query/tokeniser.go:3847
		switch data[p] {
		case 32:
			goto st59
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		case 95:
			goto tr160
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st59
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr160
			}
		default:
			goto tr160
		}
		goto st44
	tr160:
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line query/tokeniser.go:3884
		switch data[p] {
		case 32:
			goto tr161
		case 33:
			goto tr162
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 46:
			goto tr163
		case 60:
			goto tr165
		case 61:
			goto tr166
		case 62:
			goto tr167
		case 92:
			goto st56
		case 95:
			goto st60
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr161
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
	tr161:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st61
	tr223:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line query/tokeniser.go:3944
		switch data[p] {
		case 32:
			goto st61
		case 33:
			goto tr169
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 60:
			goto tr170
		case 61:
			goto tr171
		case 62:
			goto tr172
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st61
		}
		goto st44
	tr162:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st62
	tr169:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st62
	tr224:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line query/tokeniser.go:3994
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
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
			goto tr174
		case 34:
			goto tr175
		case 39:
			goto tr176
		case 43:
			goto tr177
		case 45:
			goto tr177
		case 92:
			goto st56
		case 95:
			goto tr179
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr174
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr179
				}
			case data[p] >= 65:
				goto tr179
			}
		default:
			goto tr178
		}
		goto st44
	tr174:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st64
	tr189:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st64
	tr196:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st64
	tr203:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st64
	tr209:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st64
	tr216:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line query/tokeniser.go:4074
		switch data[p] {
		case 32:
			goto st64
		case 34:
			goto tr181
		case 39:
			goto tr182
		case 43:
			goto tr183
		case 45:
			goto tr183
		case 92:
			goto st56
		case 95:
			goto tr185
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
					goto tr185
				}
			case data[p] >= 65:
				goto tr185
			}
		default:
			goto tr184
		}
		goto st44
	tr175:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st450
	tr181:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st450
	tr190:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st450
	tr197:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st450
	tr204:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st450
	tr210:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st450
	tr217:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st450
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
//line query/tokeniser.go:4168
		switch data[p] {
		case 32:
			goto tr1027
		case 34:
			goto tr114
		case 39:
			goto tr154
		case 59:
			goto tr1028
		case 92:
			goto tr116
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1027
		}
		goto tr113
	tr1017:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st451
	tr1028:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st451
	tr1046:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st451
	tr1030:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st451
	tr1033:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st451
	tr1037:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st451
	tr1040:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st451
	tr1043:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st451
	tr1118:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st451
	tr1178:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st451
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
//line query/tokeniser.go:4268
		switch data[p] {
		case 32:
			goto st451
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st451
		}
		goto st44
	tr176:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st452
	tr182:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st452
	tr191:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st452
	tr198:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st452
	tr205:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st452
	tr211:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st452
	tr218:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st452
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
//line query/tokeniser.go:4342
		switch data[p] {
		case 32:
			goto tr1029
		case 34:
			goto tr114
		case 39:
			goto tr115
		case 59:
			goto tr1030
		case 92:
			goto tr116
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1029
		}
		goto tr113
	tr177:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr183:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr192:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr199:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr206:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr212:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st65
	tr219:
//line query/tokeniser.rl:159
		commit(ttEq)
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
//line query/tokeniser.go:4418
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st453
		}
		goto st44
	tr178:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st453
	tr184:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st453
	tr193:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st453
	tr200:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st453
	tr207:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st453
	tr213:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st453
	tr220:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st453
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
//line query/tokeniser.go:4490
		switch data[p] {
		case 32:
			goto tr1031
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 46:
			goto st66
		case 59:
			goto tr1033
		case 92:
			goto st56
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st453
			}
		case data[p] >= 9:
			goto tr1031
		}
		goto st44
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st454
		}
		goto st44
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 32:
			goto tr1031
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 59:
			goto tr1033
		case 92:
			goto st56
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st454
			}
		case data[p] >= 9:
			goto tr1031
		}
		goto st44
	tr179:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st455
	tr185:
//line query/tokeniser.rl:87
		mark = p
		goto st455
	tr195:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st455
	tr201:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st455
	tr208:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st455
	tr215:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st455
	tr221:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st455
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
//line query/tokeniser.go:4602
		switch data[p] {
		case 32:
			goto tr1034
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 46:
			goto tr1035
		case 59:
			goto tr1037
		case 92:
			goto st56
		case 95:
			goto st455
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1034
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
		goto st44
	tr1035:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line query/tokeniser.go:4646
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		case 95:
			goto st456
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st456
			}
		case data[p] >= 65:
			goto st456
		}
		goto st44
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		switch data[p] {
		case 32:
			goto tr1038
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 46:
			goto st67
		case 59:
			goto tr1040
		case 92:
			goto st56
		case 95:
			goto st456
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1038
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st456
				}
			case data[p] >= 65:
				goto st456
			}
		default:
			goto st456
		}
		goto st44
	tr165:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st68
	tr170:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st68
	tr226:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line query/tokeniser.go:4738
		switch data[p] {
		case 32:
			goto tr189
		case 34:
			goto tr190
		case 39:
			goto tr191
		case 43:
			goto tr192
		case 45:
			goto tr192
		case 61:
			goto st69
		case 92:
			goto st56
		case 95:
			goto tr195
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr189
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr195
				}
			case data[p] >= 65:
				goto tr195
			}
		default:
			goto tr193
		}
		goto st44
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 32:
			goto tr196
		case 34:
			goto tr197
		case 39:
			goto tr198
		case 43:
			goto tr199
		case 45:
			goto tr199
		case 92:
			goto st56
		case 95:
			goto tr201
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr196
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr201
				}
			case data[p] >= 65:
				goto tr201
			}
		default:
			goto tr200
		}
		goto st44
	tr166:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st70
	tr171:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st70
	tr227:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line query/tokeniser.go:4841
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
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
			goto tr203
		case 34:
			goto tr204
		case 39:
			goto tr205
		case 43:
			goto tr206
		case 45:
			goto tr206
		case 92:
			goto st56
		case 95:
			goto tr208
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr203
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr208
				}
			case data[p] >= 65:
				goto tr208
			}
		default:
			goto tr207
		}
		goto st44
	tr167:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st72
	tr172:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st72
	tr228:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line query/tokeniser.go:4925
		switch data[p] {
		case 32:
			goto tr209
		case 34:
			goto tr210
		case 39:
			goto tr211
		case 43:
			goto tr212
		case 45:
			goto tr212
		case 61:
			goto st73
		case 92:
			goto st56
		case 95:
			goto tr215
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr209
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr215
				}
			case data[p] >= 65:
				goto tr215
			}
		default:
			goto tr213
		}
		goto st44
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 32:
			goto tr216
		case 34:
			goto tr217
		case 39:
			goto tr218
		case 43:
			goto tr219
		case 45:
			goto tr219
		case 92:
			goto st56
		case 95:
			goto tr221
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr216
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr221
				}
			case data[p] >= 65:
				goto tr221
			}
		default:
			goto tr220
		}
		goto st44
	tr163:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line query/tokeniser.go:5010
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
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
			goto tr223
		case 33:
			goto tr224
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 46:
			goto st74
		case 60:
			goto tr226
		case 61:
			goto tr227
		case 62:
			goto tr228
		case 92:
			goto st56
		case 95:
			goto st75
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr223
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
	tr1021:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line query/tokeniser.go:5084
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 78:
			goto st77
		case 92:
			goto st56
		case 110:
			goto st77
		}
		goto st44
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 68:
			goto st58
		case 92:
			goto st56
		case 100:
			goto st58
		}
		goto st44
	tr1022:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line query/tokeniser.go:5125
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 82:
			goto st79
		case 92:
			goto st56
		case 114:
			goto st79
		}
		goto st44
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 32:
			goto tr231
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr231
		}
		goto st44
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 73:
			goto st81
		case 92:
			goto st56
		case 105:
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
			goto tr118
		case 39:
			goto tr119
		case 84:
			goto st82
		case 92:
			goto st56
		case 116:
			goto st82
		}
		goto st44
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 72:
			goto st83
		case 92:
			goto st56
		case 104:
			goto st83
		}
		goto st44
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 73:
			goto st84
		case 92:
			goto st56
		case 105:
			goto st84
		}
		goto st44
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 78:
			goto st85
		case 92:
			goto st56
		case 110:
			goto st85
		}
		goto st44
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 32:
			goto st86
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st86
		}
		goto st44
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 32:
			goto st86
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 43:
			goto tr238
		case 45:
			goto tr238
		case 92:
			goto st56
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr239
			}
		case data[p] >= 9:
			goto st86
		}
		goto st44
	tr238:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line query/tokeniser.go:5308
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st88
		}
		goto st44
	tr239:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line query/tokeniser.go:5334
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 46:
			goto st89
		case 72:
			goto st457
		case 77:
			goto st459
		case 78:
			goto st91
		case 83:
			goto st457
		case 85:
			goto st91
		case 92:
			goto st56
		case 104:
			goto st457
		case 109:
			goto st459
		case 110:
			goto st91
		case 115:
			goto st457
		case 117:
			goto st91
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st88
		}
		goto st44
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st90
		}
		goto st44
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 72:
			goto st457
		case 77:
			goto st459
		case 78:
			goto st91
		case 83:
			goto st457
		case 85:
			goto st91
		case 92:
			goto st56
		case 104:
			goto st457
		case 109:
			goto st459
		case 110:
			goto st91
		case 115:
			goto st457
		case 117:
			goto st91
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st90
		}
		goto st44
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 32:
			goto tr1041
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 43:
			goto st87
		case 45:
			goto st87
		case 59:
			goto tr1043
		case 92:
			goto st56
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st88
			}
		case data[p] >= 9:
			goto tr1041
		}
		goto st44
	tr1041:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
//line query/tokeniser.go:5466
		switch data[p] {
		case 32:
			goto st458
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 59:
			goto st451
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st458
		}
		goto st44
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		switch data[p] {
		case 32:
			goto tr1041
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 43:
			goto st87
		case 45:
			goto st87
		case 59:
			goto tr1043
		case 83:
			goto st457
		case 92:
			goto st56
		case 115:
			goto st457
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st88
			}
		case data[p] >= 9:
			goto tr1041
		}
		goto st44
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 83:
			goto st457
		case 92:
			goto st56
		case 115:
			goto st457
		}
		goto st44
	tr1025:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line query/tokeniser.go:5544
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		case 124:
			goto st79
		}
		goto st44
	tr1026:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line query/tokeniser.go:5565
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		case 136:
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
			goto tr118
		case 39:
			goto tr119
		case 92:
			goto st56
		case 168:
			goto st79
		}
		goto st44
	tr156:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
//line query/tokeniser.go:5602
		switch data[p] {
		case 32:
			goto tr1045
		case 34:
			goto tr118
		case 39:
			goto tr119
		case 59:
			goto tr1046
		case 92:
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1045
		}
		goto st44
	tr144:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st461
	tr150:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st461
	tr252:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st461
	tr259:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st461
	tr266:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st461
	tr272:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st461
	tr279:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st461
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
//line query/tokeniser.go:5678
		switch data[p] {
		case 32:
			goto tr1047
		case 39:
			goto tr370
		case 59:
			goto tr1048
		case 92:
			goto tr371
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1047
		}
		goto tr369
	tr1000:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st462
	tr1004:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st462
	tr1048:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st462
	tr1051:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st462
	tr1055:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st462
	tr1058:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st462
	tr1061:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st462
	tr1196:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st462
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
//line query/tokeniser.go:5760
		switch data[p] {
		case 32:
			goto st462
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st462
		}
		goto st45
	tr145:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st95
	tr151:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st95
	tr253:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st95
	tr260:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st95
	tr267:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st95
	tr273:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st95
	tr280:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line query/tokeniser.go:5832
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st463
		}
		goto st45
	tr146:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st463
	tr152:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st463
	tr254:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st463
	tr261:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st463
	tr268:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st463
	tr274:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st463
	tr281:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st463
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
//line query/tokeniser.go:5902
		switch data[p] {
		case 32:
			goto tr1049
		case 39:
			goto tr122
		case 46:
			goto st96
		case 59:
			goto tr1051
		case 92:
			goto st46
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st463
			}
		case data[p] >= 9:
			goto tr1049
		}
		goto st45
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st464
		}
		goto st45
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		switch data[p] {
		case 32:
			goto tr1049
		case 39:
			goto tr122
		case 59:
			goto tr1051
		case 92:
			goto st46
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st464
			}
		case data[p] >= 9:
			goto tr1049
		}
		goto st45
	tr147:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st465
	tr153:
//line query/tokeniser.rl:87
		mark = p
		goto st465
	tr256:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st465
	tr262:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st465
	tr269:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st465
	tr276:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st465
	tr282:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st465
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
//line query/tokeniser.go:6008
		switch data[p] {
		case 32:
			goto tr1052
		case 39:
			goto tr122
		case 46:
			goto tr1053
		case 59:
			goto tr1055
		case 92:
			goto st46
		case 95:
			goto st465
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1052
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st465
				}
			case data[p] >= 65:
				goto st465
			}
		default:
			goto st465
		}
		goto st45
	tr1053:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line query/tokeniser.go:6050
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		case 95:
			goto st466
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st466
			}
		case data[p] >= 65:
			goto st466
		}
		goto st45
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		switch data[p] {
		case 32:
			goto tr1056
		case 39:
			goto tr122
		case 46:
			goto st97
		case 59:
			goto tr1058
		case 92:
			goto st46
		case 95:
			goto st466
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1056
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st466
				}
			case data[p] >= 65:
				goto st466
			}
		default:
			goto st466
		}
		goto st45
	tr133:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st98
	tr138:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st98
	tr287:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line query/tokeniser.go:6138
		switch data[p] {
		case 32:
			goto tr250
		case 34:
			goto tr251
		case 39:
			goto tr252
		case 43:
			goto tr253
		case 45:
			goto tr253
		case 61:
			goto st99
		case 92:
			goto st46
		case 95:
			goto tr256
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr250
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr256
				}
			case data[p] >= 65:
				goto tr256
			}
		default:
			goto tr254
		}
		goto st45
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 32:
			goto tr257
		case 34:
			goto tr258
		case 39:
			goto tr259
		case 43:
			goto tr260
		case 45:
			goto tr260
		case 92:
			goto st46
		case 95:
			goto tr262
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr257
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr262
				}
			case data[p] >= 65:
				goto tr262
			}
		default:
			goto tr261
		}
		goto st45
	tr134:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st100
	tr139:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st100
	tr288:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
//line query/tokeniser.go:6241
		switch data[p] {
		case 39:
			goto tr122
		case 61:
			goto st101
		case 92:
			goto st46
		}
		goto st45
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 32:
			goto tr264
		case 34:
			goto tr265
		case 39:
			goto tr266
		case 43:
			goto tr267
		case 45:
			goto tr267
		case 92:
			goto st46
		case 95:
			goto tr269
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr264
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr269
				}
			case data[p] >= 65:
				goto tr269
			}
		default:
			goto tr268
		}
		goto st45
	tr135:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st102
	tr140:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st102
	tr289:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line query/tokeniser.go:6323
		switch data[p] {
		case 32:
			goto tr270
		case 34:
			goto tr271
		case 39:
			goto tr272
		case 43:
			goto tr273
		case 45:
			goto tr273
		case 61:
			goto st103
		case 92:
			goto st46
		case 95:
			goto tr276
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr270
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr276
				}
			case data[p] >= 65:
				goto tr276
			}
		default:
			goto tr274
		}
		goto st45
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto tr277
		case 34:
			goto tr278
		case 39:
			goto tr279
		case 43:
			goto tr280
		case 45:
			goto tr280
		case 92:
			goto st46
		case 95:
			goto tr282
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr277
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr282
				}
			case data[p] >= 65:
				goto tr282
			}
		default:
			goto tr281
		}
		goto st45
	tr131:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line query/tokeniser.go:6408
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		case 95:
			goto st105
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st105
			}
		case data[p] >= 65:
			goto st105
		}
		goto st45
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 32:
			goto tr284
		case 33:
			goto tr285
		case 39:
			goto tr122
		case 46:
			goto st104
		case 60:
			goto tr287
		case 61:
			goto tr288
		case 62:
			goto tr289
		case 92:
			goto st46
		case 95:
			goto st105
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr284
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st105
				}
			case data[p] >= 65:
				goto st105
			}
		default:
			goto st105
		}
		goto st45
	tr1008:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line query/tokeniser.go:6478
		switch data[p] {
		case 39:
			goto tr122
		case 78:
			goto st107
		case 92:
			goto st46
		case 110:
			goto st107
		}
		goto st45
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 39:
			goto tr122
		case 68:
			goto st48
		case 92:
			goto st46
		case 100:
			goto st48
		}
		goto st45
	tr1009:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line query/tokeniser.go:6515
		switch data[p] {
		case 39:
			goto tr122
		case 82:
			goto st109
		case 92:
			goto st46
		case 114:
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
			goto tr292
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr292
		}
		goto st45
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 39:
			goto tr122
		case 73:
			goto st111
		case 92:
			goto st46
		case 105:
			goto st111
		}
		goto st45
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 39:
			goto tr122
		case 84:
			goto st112
		case 92:
			goto st46
		case 116:
			goto st112
		}
		goto st45
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 39:
			goto tr122
		case 72:
			goto st113
		case 92:
			goto st46
		case 104:
			goto st113
		}
		goto st45
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 39:
			goto tr122
		case 73:
			goto st114
		case 92:
			goto st46
		case 105:
			goto st114
		}
		goto st45
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 39:
			goto tr122
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
		case 32:
			goto st116
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st116
		}
		goto st45
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 32:
			goto st116
		case 39:
			goto tr122
		case 43:
			goto tr299
		case 45:
			goto tr299
		case 92:
			goto st46
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr300
			}
		case data[p] >= 9:
			goto st116
		}
		goto st45
	tr299:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
//line query/tokeniser.go:6680
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st118
		}
		goto st45
	tr300:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line query/tokeniser.go:6704
		switch data[p] {
		case 39:
			goto tr122
		case 46:
			goto st119
		case 72:
			goto st467
		case 77:
			goto st469
		case 78:
			goto st121
		case 83:
			goto st467
		case 85:
			goto st121
		case 92:
			goto st46
		case 104:
			goto st467
		case 109:
			goto st469
		case 110:
			goto st121
		case 115:
			goto st467
		case 117:
			goto st121
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st118
		}
		goto st45
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		}
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr122
		case 72:
			goto st467
		case 77:
			goto st469
		case 78:
			goto st121
		case 83:
			goto st467
		case 85:
			goto st121
		case 92:
			goto st46
		case 104:
			goto st467
		case 109:
			goto st469
		case 110:
			goto st121
		case 115:
			goto st467
		case 117:
			goto st121
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st120
		}
		goto st45
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		switch data[p] {
		case 32:
			goto tr1059
		case 39:
			goto tr122
		case 43:
			goto st117
		case 45:
			goto st117
		case 59:
			goto tr1061
		case 92:
			goto st46
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st118
			}
		case data[p] >= 9:
			goto tr1059
		}
		goto st45
	tr1059:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line query/tokeniser.go:6828
		switch data[p] {
		case 32:
			goto st468
		case 39:
			goto tr122
		case 59:
			goto st462
		case 92:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st468
		}
		goto st45
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		switch data[p] {
		case 32:
			goto tr1059
		case 39:
			goto tr122
		case 43:
			goto st117
		case 45:
			goto st117
		case 59:
			goto tr1061
		case 83:
			goto st467
		case 92:
			goto st46
		case 115:
			goto st467
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st118
			}
		case data[p] >= 9:
			goto tr1059
		}
		goto st45
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 39:
			goto tr122
		case 83:
			goto st467
		case 92:
			goto st46
		case 115:
			goto st467
		}
		goto st45
	tr1012:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
//line query/tokeniser.go:6900
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		case 124:
			goto st109
		}
		goto st45
	tr1013:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
//line query/tokeniser.go:6919
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		case 136:
			goto st124
		}
		goto st45
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st46
		case 168:
			goto st109
		}
		goto st45
	tr104:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st125
	tr110:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st125
	tr314:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st125
	tr321:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st125
	tr328:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st125
	tr334:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st125
	tr341:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
//line query/tokeniser.go:7002
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st470
		}
		goto st33
	tr105:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st470
	tr111:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st470
	tr315:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st470
	tr322:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st470
	tr329:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st470
	tr335:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st470
	tr342:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st470
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
//line query/tokeniser.go:7072
		switch data[p] {
		case 32:
			goto tr1063
		case 34:
			goto tr54
		case 46:
			goto st126
		case 59:
			goto tr1065
		case 92:
			goto st34
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st470
			}
		case data[p] >= 9:
			goto tr1063
		}
		goto st33
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st471
		}
		goto st33
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		switch data[p] {
		case 32:
			goto tr1063
		case 34:
			goto tr54
		case 59:
			goto tr1065
		case 92:
			goto st34
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st471
			}
		case data[p] >= 9:
			goto tr1063
		}
		goto st33
	tr106:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st472
	tr112:
//line query/tokeniser.rl:87
		mark = p
		goto st472
	tr317:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st472
	tr323:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st472
	tr330:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st472
	tr337:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st472
	tr343:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st472
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
//line query/tokeniser.go:7178
		switch data[p] {
		case 32:
			goto tr1066
		case 34:
			goto tr54
		case 46:
			goto tr1067
		case 59:
			goto tr1069
		case 92:
			goto st34
		case 95:
			goto st472
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1066
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st472
				}
			case data[p] >= 65:
				goto st472
			}
		default:
			goto st472
		}
		goto st33
	tr1067:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line query/tokeniser.go:7220
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		case 95:
			goto st473
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st473
			}
		case data[p] >= 65:
			goto st473
		}
		goto st33
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		switch data[p] {
		case 32:
			goto tr1070
		case 34:
			goto tr54
		case 46:
			goto st127
		case 59:
			goto tr1072
		case 92:
			goto st34
		case 95:
			goto st473
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1070
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st473
				}
			case data[p] >= 65:
				goto st473
			}
		default:
			goto st473
		}
		goto st33
	tr92:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st128
	tr97:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st128
	tr348:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line query/tokeniser.go:7308
		switch data[p] {
		case 32:
			goto tr311
		case 34:
			goto tr312
		case 39:
			goto tr313
		case 43:
			goto tr314
		case 45:
			goto tr314
		case 61:
			goto st129
		case 92:
			goto st34
		case 95:
			goto tr317
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr311
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr317
				}
			case data[p] >= 65:
				goto tr317
			}
		default:
			goto tr315
		}
		goto st33
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 32:
			goto tr318
		case 34:
			goto tr319
		case 39:
			goto tr320
		case 43:
			goto tr321
		case 45:
			goto tr321
		case 92:
			goto st34
		case 95:
			goto tr323
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr318
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr323
				}
			case data[p] >= 65:
				goto tr323
			}
		default:
			goto tr322
		}
		goto st33
	tr93:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st130
	tr98:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st130
	tr349:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st130
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
//line query/tokeniser.go:7411
		switch data[p] {
		case 34:
			goto tr54
		case 61:
			goto st131
		case 92:
			goto st34
		}
		goto st33
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 32:
			goto tr325
		case 34:
			goto tr326
		case 39:
			goto tr327
		case 43:
			goto tr328
		case 45:
			goto tr328
		case 92:
			goto st34
		case 95:
			goto tr330
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr325
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr330
				}
			case data[p] >= 65:
				goto tr330
			}
		default:
			goto tr329
		}
		goto st33
	tr94:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st132
	tr99:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st132
	tr350:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line query/tokeniser.go:7493
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
			goto st133
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
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
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
	tr90:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line query/tokeniser.go:7578
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		case 95:
			goto st135
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st135
			}
		case data[p] >= 65:
			goto st135
		}
		goto st33
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 32:
			goto tr345
		case 33:
			goto tr346
		case 34:
			goto tr54
		case 46:
			goto st134
		case 60:
			goto tr348
		case 61:
			goto tr349
		case 62:
			goto tr350
		case 92:
			goto st34
		case 95:
			goto st135
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
					goto st135
				}
			case data[p] >= 65:
				goto st135
			}
		default:
			goto st135
		}
		goto st33
	tr991:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line query/tokeniser.go:7648
		switch data[p] {
		case 34:
			goto tr54
		case 78:
			goto st137
		case 92:
			goto st34
		case 110:
			goto st137
		}
		goto st33
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 34:
			goto tr54
		case 68:
			goto st36
		case 92:
			goto st34
		case 100:
			goto st36
		}
		goto st33
	tr992:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line query/tokeniser.go:7685
		switch data[p] {
		case 34:
			goto tr54
		case 82:
			goto st139
		case 92:
			goto st34
		case 114:
			goto st139
		}
		goto st33
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 32:
			goto tr353
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr353
		}
		goto st33
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 34:
			goto tr54
		case 73:
			goto st141
		case 92:
			goto st34
		case 105:
			goto st141
		}
		goto st33
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 34:
			goto tr54
		case 84:
			goto st142
		case 92:
			goto st34
		case 116:
			goto st142
		}
		goto st33
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 34:
			goto tr54
		case 72:
			goto st143
		case 92:
			goto st34
		case 104:
			goto st143
		}
		goto st33
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 34:
			goto tr54
		case 73:
			goto st144
		case 92:
			goto st34
		case 105:
			goto st144
		}
		goto st33
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 34:
			goto tr54
		case 78:
			goto st145
		case 92:
			goto st34
		case 110:
			goto st145
		}
		goto st33
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 32:
			goto st146
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st146
		}
		goto st33
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 32:
			goto st146
		case 34:
			goto tr54
		case 43:
			goto tr360
		case 45:
			goto tr360
		case 92:
			goto st34
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr361
			}
		case data[p] >= 9:
			goto st146
		}
		goto st33
	tr360:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line query/tokeniser.go:7850
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st148
		}
		goto st33
	tr361:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line query/tokeniser.go:7874
		switch data[p] {
		case 34:
			goto tr54
		case 46:
			goto st149
		case 72:
			goto st474
		case 77:
			goto st476
		case 78:
			goto st151
		case 83:
			goto st474
		case 85:
			goto st151
		case 92:
			goto st34
		case 104:
			goto st474
		case 109:
			goto st476
		case 110:
			goto st151
		case 115:
			goto st474
		case 117:
			goto st151
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st148
		}
		goto st33
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st150
		}
		goto st33
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		switch data[p] {
		case 34:
			goto tr54
		case 72:
			goto st474
		case 77:
			goto st476
		case 78:
			goto st151
		case 83:
			goto st474
		case 85:
			goto st151
		case 92:
			goto st34
		case 104:
			goto st474
		case 109:
			goto st476
		case 110:
			goto st151
		case 115:
			goto st474
		case 117:
			goto st151
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st150
		}
		goto st33
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		switch data[p] {
		case 32:
			goto tr1073
		case 34:
			goto tr54
		case 43:
			goto st147
		case 45:
			goto st147
		case 59:
			goto tr1075
		case 92:
			goto st34
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] >= 9:
			goto tr1073
		}
		goto st33
	tr1073:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st475
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
//line query/tokeniser.go:7998
		switch data[p] {
		case 32:
			goto st475
		case 34:
			goto tr54
		case 59:
			goto st442
		case 92:
			goto st34
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st475
		}
		goto st33
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		switch data[p] {
		case 32:
			goto tr1073
		case 34:
			goto tr54
		case 43:
			goto st147
		case 45:
			goto st147
		case 59:
			goto tr1075
		case 83:
			goto st474
		case 92:
			goto st34
		case 115:
			goto st474
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st148
			}
		case data[p] >= 9:
			goto tr1073
		}
		goto st33
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		switch data[p] {
		case 34:
			goto tr54
		case 83:
			goto st474
		case 92:
			goto st34
		case 115:
			goto st474
		}
		goto st33
	tr995:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st152
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
//line query/tokeniser.go:8070
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		case 124:
			goto st139
		}
		goto st33
	tr996:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st153
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
//line query/tokeniser.go:8089
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st34
		case 136:
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
			goto tr54
		case 92:
			goto st34
		case 168:
			goto st139
		}
		goto st33
	tr75:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st155
	tr78:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st155
	tr377:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st155
	tr384:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st155
	tr391:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st155
	tr397:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st155
	tr404:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st155
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
//line query/tokeniser.go:8158
		switch data[p] {
		case 39:
			goto tr370
		case 92:
			goto tr371
		}
		goto tr369
	tr41:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st156
	tr47:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st156
	tr378:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st156
	tr385:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st156
	tr392:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st156
	tr398:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st156
	tr405:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st156
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
//line query/tokeniser.go:8225
		if 48 <= data[p] && data[p] <= 57 {
			goto st477
		}
		goto st0
	tr42:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st477
	tr48:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st477
	tr379:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st477
	tr386:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st477
	tr393:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st477
	tr399:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st477
	tr406:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st477
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
//line query/tokeniser.go:8289
		switch data[p] {
		case 32:
			goto tr1077
		case 46:
			goto st157
		case 59:
			goto tr1079
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st477
			}
		case data[p] >= 9:
			goto tr1077
		}
		goto st0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		if 48 <= data[p] && data[p] <= 57 {
			goto st478
		}
		goto st0
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		switch data[p] {
		case 32:
			goto tr1077
		case 59:
			goto tr1079
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st478
			}
		case data[p] >= 9:
			goto tr1077
		}
		goto st0
	tr43:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st479
	tr49:
//line query/tokeniser.rl:87
		mark = p
		goto st479
	tr381:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st479
	tr387:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st479
	tr394:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st479
	tr401:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st479
	tr407:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st479
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
//line query/tokeniser.go:8381
		switch data[p] {
		case 32:
			goto tr1080
		case 46:
			goto tr1081
		case 59:
			goto tr1083
		case 95:
			goto st479
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1080
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st479
				}
			case data[p] >= 65:
				goto st479
			}
		default:
			goto st479
		}
		goto st0
	tr1081:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st158
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
//line query/tokeniser.go:8419
		if data[p] == 95 {
			goto st480
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st480
			}
		case data[p] >= 65:
			goto st480
		}
		goto st0
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		switch data[p] {
		case 32:
			goto tr1084
		case 46:
			goto st158
		case 59:
			goto tr1086
		case 95:
			goto st480
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1084
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st480
				}
			case data[p] >= 65:
				goto st480
			}
		default:
			goto st480
		}
		goto st0
	tr64:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st159
	tr69:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st159
	tr412:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st159
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
//line query/tokeniser.go:8498
		switch data[p] {
		case 32:
			goto tr375
		case 34:
			goto tr376
		case 39:
			goto tr377
		case 43:
			goto tr378
		case 45:
			goto tr378
		case 61:
			goto st160
		case 95:
			goto tr381
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
					goto tr381
				}
			case data[p] >= 65:
				goto tr381
			}
		default:
			goto tr379
		}
		goto st0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 32:
			goto tr382
		case 34:
			goto tr383
		case 39:
			goto tr384
		case 43:
			goto tr385
		case 45:
			goto tr385
		case 95:
			goto tr387
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr382
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr387
				}
			case data[p] >= 65:
				goto tr387
			}
		default:
			goto tr386
		}
		goto st0
	tr65:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st161
	tr70:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st161
	tr413:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st161
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
//line query/tokeniser.go:8597
		if data[p] == 61 {
			goto st162
		}
		goto st0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 32:
			goto tr389
		case 34:
			goto tr390
		case 39:
			goto tr391
		case 43:
			goto tr392
		case 45:
			goto tr392
		case 95:
			goto tr394
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr389
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr394
				}
			case data[p] >= 65:
				goto tr394
			}
		default:
			goto tr393
		}
		goto st0
	tr66:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st163
	tr71:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st163
	tr414:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st163
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
//line query/tokeniser.go:8672
		switch data[p] {
		case 32:
			goto tr395
		case 34:
			goto tr396
		case 39:
			goto tr397
		case 43:
			goto tr398
		case 45:
			goto tr398
		case 61:
			goto st164
		case 95:
			goto tr401
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr395
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr401
				}
			case data[p] >= 65:
				goto tr401
			}
		default:
			goto tr399
		}
		goto st0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 32:
			goto tr402
		case 34:
			goto tr403
		case 39:
			goto tr404
		case 43:
			goto tr405
		case 45:
			goto tr405
		case 95:
			goto tr407
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr402
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr407
				}
			case data[p] >= 65:
				goto tr407
			}
		default:
			goto tr406
		}
		goto st0
	tr62:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st165
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
//line query/tokeniser.go:8753
		if data[p] == 95 {
			goto st166
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st166
			}
		case data[p] >= 65:
			goto st166
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 32:
			goto tr409
		case 33:
			goto tr410
		case 46:
			goto st165
		case 60:
			goto tr412
		case 61:
			goto tr413
		case 62:
			goto tr414
		case 95:
			goto st166
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr409
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st166
				}
			case data[p] >= 65:
				goto st166
			}
		default:
			goto st166
		}
		goto st0
	tr980:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st167
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
//line query/tokeniser.go:8814
		switch data[p] {
		case 78:
			goto st168
		case 110:
			goto st168
		}
		goto st0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 68:
			goto st25
		case 100:
			goto st25
		}
		goto st0
	tr981:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st169
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
//line query/tokeniser.go:8843
		switch data[p] {
		case 82:
			goto st170
		case 114:
			goto st170
		}
		goto st0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		if data[p] == 32 {
			goto tr417
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr417
		}
		goto st0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 73:
			goto st172
		case 105:
			goto st172
		}
		goto st0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 84:
			goto st173
		case 116:
			goto st173
		}
		goto st0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 72:
			goto st174
		case 104:
			goto st174
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 73:
			goto st175
		case 105:
			goto st175
		}
		goto st0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 78:
			goto st176
		case 110:
			goto st176
		}
		goto st0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		if data[p] == 32 {
			goto st177
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st177
		}
		goto st0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 32:
			goto st177
		case 43:
			goto tr423
		case 45:
			goto tr423
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr424
			}
		case data[p] >= 9:
			goto st177
		}
		goto st0
	tr423:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st178
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
//line query/tokeniser.go:8970
		if 48 <= data[p] && data[p] <= 57 {
			goto st179
		}
		goto st0
	tr424:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st179
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
//line query/tokeniser.go:8988
		switch data[p] {
		case 46:
			goto st180
		case 72:
			goto st481
		case 77:
			goto st483
		case 78:
			goto st182
		case 83:
			goto st481
		case 85:
			goto st182
		case 104:
			goto st481
		case 109:
			goto st483
		case 110:
			goto st182
		case 115:
			goto st481
		case 117:
			goto st182
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st179
		}
		goto st0
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		if 48 <= data[p] && data[p] <= 57 {
			goto st181
		}
		goto st0
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 72:
			goto st481
		case 77:
			goto st483
		case 78:
			goto st182
		case 83:
			goto st481
		case 85:
			goto st182
		case 104:
			goto st481
		case 109:
			goto st483
		case 110:
			goto st182
		case 115:
			goto st481
		case 117:
			goto st182
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st181
		}
		goto st0
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		switch data[p] {
		case 32:
			goto tr1087
		case 43:
			goto st178
		case 45:
			goto st178
		case 59:
			goto tr1089
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st179
			}
		case data[p] >= 9:
			goto tr1087
		}
		goto st0
	tr1087:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st482
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
//line query/tokeniser.go:9094
		switch data[p] {
		case 32:
			goto st482
		case 59:
			goto st436
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st482
		}
		goto st0
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
		switch data[p] {
		case 32:
			goto tr1087
		case 43:
			goto st178
		case 45:
			goto st178
		case 59:
			goto tr1089
		case 83:
			goto st481
		case 115:
			goto st481
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st179
			}
		case data[p] >= 9:
			goto tr1087
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch data[p] {
		case 83:
			goto st481
		case 115:
			goto st481
		}
		goto st0
	tr984:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st183
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
//line query/tokeniser.go:9154
		if data[p] == 124 {
			goto st170
		}
		goto st0
	tr985:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st184
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
//line query/tokeniser.go:9168
		if data[p] == 136 {
			goto st185
		}
		goto st0
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		if data[p] == 168 {
			goto st170
		}
		goto st0
	tr52:
//line query/tokeniser.rl:87
		mark = p
		goto st186
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
//line query/tokeniser.go:9191
		switch data[p] {
		case 34:
			goto tr432
		case 92:
			goto st186
		}
		goto st23
	tr432:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st484
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
//line query/tokeniser.go:9208
		switch data[p] {
		case 32:
			goto tr1091
		case 34:
			goto tr54
		case 59:
			goto tr1092
		case 92:
			goto st186
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1091
		}
		goto st23
	tr1091:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st485
	tr1104:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st485
	tr1137:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st485
	tr1140:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st485
	tr1144:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
//line query/tokeniser.go:9266
		switch data[p] {
		case 32:
			goto st485
		case 34:
			goto tr54
		case 38:
			goto tr1094
		case 59:
			goto st488
		case 65:
			goto tr1096
		case 79:
			goto tr1097
		case 87:
			goto st251
		case 92:
			goto st186
		case 94:
			goto tr1099
		case 97:
			goto tr1096
		case 111:
			goto tr1097
		case 119:
			goto st251
		case 124:
			goto tr1100
		case 226:
			goto tr1101
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st485
		}
		goto st23
	tr1094:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st187
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
//line query/tokeniser.go:9310
		switch data[p] {
		case 34:
			goto tr54
		case 38:
			goto st188
		case 92:
			goto st186
		}
		goto st23
	tr1099:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st188
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
//line query/tokeniser.go:9329
		switch data[p] {
		case 32:
			goto tr434
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr434
		}
		goto st23
	tr434:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st189
	tr606:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
//line query/tokeniser.go:9355
		switch data[p] {
		case 32:
			goto st189
		case 34:
			goto tr54
		case 92:
			goto st186
		case 95:
			goto tr436
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st189
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr436
			}
		default:
			goto tr436
		}
		goto st23
	tr436:
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
//line query/tokeniser.go:9390
		switch data[p] {
		case 32:
			goto tr437
		case 33:
			goto tr438
		case 34:
			goto tr54
		case 46:
			goto tr439
		case 60:
			goto tr441
		case 61:
			goto tr442
		case 62:
			goto tr443
		case 92:
			goto st186
		case 95:
			goto st190
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
					goto st190
				}
			case data[p] >= 65:
				goto st190
			}
		default:
			goto st190
		}
		goto st23
	tr437:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st191
	tr598:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st191
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
//line query/tokeniser.go:9448
		switch data[p] {
		case 32:
			goto st191
		case 33:
			goto tr445
		case 34:
			goto tr54
		case 60:
			goto tr446
		case 61:
			goto tr447
		case 62:
			goto tr448
		case 92:
			goto st186
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st191
		}
		goto st23
	tr438:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st192
	tr445:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st192
	tr599:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st192
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
//line query/tokeniser.go:9496
		switch data[p] {
		case 34:
			goto tr54
		case 61:
			goto st193
		case 92:
			goto st186
		}
		goto st23
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch data[p] {
		case 32:
			goto tr450
		case 34:
			goto tr451
		case 39:
			goto tr452
		case 43:
			goto tr453
		case 45:
			goto tr453
		case 92:
			goto st186
		case 95:
			goto tr455
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr450
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr455
				}
			case data[p] >= 65:
				goto tr455
			}
		default:
			goto tr454
		}
		goto st23
	tr450:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st194
	tr564:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st194
	tr571:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st194
	tr578:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st194
	tr584:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st194
	tr591:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line query/tokeniser.go:9574
		switch data[p] {
		case 32:
			goto st194
		case 34:
			goto tr457
		case 39:
			goto tr458
		case 43:
			goto tr459
		case 45:
			goto tr459
		case 92:
			goto st186
		case 95:
			goto tr461
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st194
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr461
				}
			case data[p] >= 65:
				goto tr461
			}
		default:
			goto tr460
		}
		goto st23
	tr451:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st486
	tr457:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st486
	tr565:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st486
	tr572:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st486
	tr579:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st486
	tr585:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st486
	tr592:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st486
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
//line query/tokeniser.go:9668
		switch data[p] {
		case 32:
			goto tr1102
		case 34:
			goto tr51
		case 59:
			goto tr1103
		case 92:
			goto tr80
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1102
		}
		goto tr79
	tr452:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st195
	tr458:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st195
	tr566:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st195
	tr573:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st195
	tr580:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st195
	tr586:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st195
	tr593:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
//line query/tokeniser.go:9728
		switch data[p] {
		case 34:
			goto tr463
		case 39:
			goto tr464
		case 92:
			goto tr465
		}
		goto tr462
	tr462:
//line query/tokeniser.rl:87
		mark = p
		goto st196
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
//line query/tokeniser.go:9747
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		goto st196
	tr464:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st487
	tr467:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st487
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
//line query/tokeniser.go:9772
		switch data[p] {
		case 32:
			goto tr1104
		case 34:
			goto tr54
		case 59:
			goto tr1105
		case 92:
			goto st186
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1104
		}
		goto st23
	tr1092:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st488
	tr1105:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st488
	tr1139:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st488
	tr1143:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st488
	tr1146:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st488
	tr1149:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st488
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
//line query/tokeniser.go:9838
		switch data[p] {
		case 32:
			goto st488
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st488
		}
		goto st23
	tr465:
//line query/tokeniser.rl:87
		mark = p
		goto st197
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
//line query/tokeniser.go:9860
		switch data[p] {
		case 34:
			goto tr469
		case 39:
			goto tr470
		case 92:
			goto st197
		}
		goto st196
	tr469:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st489
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
//line query/tokeniser.go:9879
		switch data[p] {
		case 32:
			goto tr1106
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 59:
			goto tr1107
		case 92:
			goto st197
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1106
		}
		goto st196
	tr1106:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st490
	tr1135:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st490
	tr1119:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st490
	tr1121:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st490
	tr1124:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st490
	tr1128:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st490
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
//line query/tokeniser.go:9947
		switch data[p] {
		case 32:
			goto st490
		case 34:
			goto tr118
		case 38:
			goto tr1109
		case 39:
			goto tr467
		case 59:
			goto st493
		case 65:
			goto tr1111
		case 79:
			goto tr1112
		case 87:
			goto st221
		case 92:
			goto st197
		case 94:
			goto tr1114
		case 97:
			goto tr1111
		case 111:
			goto tr1112
		case 119:
			goto st221
		case 124:
			goto tr1115
		case 226:
			goto tr1116
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st490
		}
		goto st196
	tr1109:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st198
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
//line query/tokeniser.go:9993
		switch data[p] {
		case 34:
			goto tr118
		case 38:
			goto st199
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		goto st196
	tr1114:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line query/tokeniser.go:10014
		switch data[p] {
		case 32:
			goto tr472
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr472
		}
		goto st196
	tr472:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st200
	tr545:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st200
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
//line query/tokeniser.go:10042
		switch data[p] {
		case 32:
			goto st200
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		case 95:
			goto tr474
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st200
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr474
			}
		default:
			goto tr474
		}
		goto st196
	tr474:
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st201
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
//line query/tokeniser.go:10079
		switch data[p] {
		case 32:
			goto tr475
		case 33:
			goto tr476
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 46:
			goto tr477
		case 60:
			goto tr479
		case 61:
			goto tr480
		case 62:
			goto tr481
		case 92:
			goto st197
		case 95:
			goto st201
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr475
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st201
				}
			case data[p] >= 65:
				goto st201
			}
		default:
			goto st201
		}
		goto st196
	tr475:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st202
	tr537:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st202
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
//line query/tokeniser.go:10139
		switch data[p] {
		case 32:
			goto st202
		case 33:
			goto tr483
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 60:
			goto tr484
		case 61:
			goto tr485
		case 62:
			goto tr486
		case 92:
			goto st197
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st202
		}
		goto st196
	tr476:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st203
	tr483:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st203
	tr538:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st203
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
//line query/tokeniser.go:10189
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 61:
			goto st204
		case 92:
			goto st197
		}
		goto st196
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		switch data[p] {
		case 32:
			goto tr488
		case 34:
			goto tr489
		case 39:
			goto tr490
		case 43:
			goto tr491
		case 45:
			goto tr491
		case 92:
			goto st197
		case 95:
			goto tr493
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr488
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr493
				}
			case data[p] >= 65:
				goto tr493
			}
		default:
			goto tr492
		}
		goto st196
	tr488:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st205
	tr503:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st205
	tr510:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st205
	tr517:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st205
	tr523:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st205
	tr530:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
//line query/tokeniser.go:10269
		switch data[p] {
		case 32:
			goto st205
		case 34:
			goto tr495
		case 39:
			goto tr496
		case 43:
			goto tr497
		case 45:
			goto tr497
		case 92:
			goto st197
		case 95:
			goto tr499
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st205
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr499
				}
			case data[p] >= 65:
				goto tr499
			}
		default:
			goto tr498
		}
		goto st196
	tr489:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st491
	tr495:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st491
	tr504:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st491
	tr511:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st491
	tr518:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st491
	tr524:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st491
	tr531:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st491
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
//line query/tokeniser.go:10363
		switch data[p] {
		case 32:
			goto tr1117
		case 34:
			goto tr114
		case 39:
			goto tr154
		case 59:
			goto tr1118
		case 92:
			goto tr116
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1117
		}
		goto tr113
	tr490:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st492
	tr496:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st492
	tr505:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st492
	tr512:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st492
	tr519:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st492
	tr525:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st492
	tr532:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st492
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
//line query/tokeniser.go:10439
		switch data[p] {
		case 32:
			goto tr1119
		case 34:
			goto tr463
		case 39:
			goto tr464
		case 59:
			goto tr1120
		case 92:
			goto tr465
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1119
		}
		goto tr462
	tr1107:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st493
	tr1136:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st493
	tr1120:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st493
	tr1123:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st493
	tr1127:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st493
	tr1130:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st493
	tr1133:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st493
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
//line query/tokeniser.go:10515
		switch data[p] {
		case 32:
			goto st493
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st493
		}
		goto st196
	tr491:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st206
	tr497:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st206
	tr506:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st206
	tr513:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st206
	tr520:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st206
	tr526:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st206
	tr533:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line query/tokeniser.go:10589
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st494
		}
		goto st196
	tr492:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st494
	tr498:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st494
	tr507:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st494
	tr514:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st494
	tr521:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st494
	tr527:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st494
	tr534:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st494
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
//line query/tokeniser.go:10661
		switch data[p] {
		case 32:
			goto tr1121
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 46:
			goto st207
		case 59:
			goto tr1123
		case 92:
			goto st197
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st494
			}
		case data[p] >= 9:
			goto tr1121
		}
		goto st196
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st495
		}
		goto st196
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		switch data[p] {
		case 32:
			goto tr1121
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 59:
			goto tr1123
		case 92:
			goto st197
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st495
			}
		case data[p] >= 9:
			goto tr1121
		}
		goto st196
	tr493:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st496
	tr499:
//line query/tokeniser.rl:87
		mark = p
		goto st496
	tr509:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st496
	tr515:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st496
	tr522:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st496
	tr529:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st496
	tr535:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st496
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
//line query/tokeniser.go:10773
		switch data[p] {
		case 32:
			goto tr1124
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 46:
			goto tr1125
		case 59:
			goto tr1127
		case 92:
			goto st197
		case 95:
			goto st496
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1124
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st496
				}
			case data[p] >= 65:
				goto st496
			}
		default:
			goto st496
		}
		goto st196
	tr1125:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
//line query/tokeniser.go:10817
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		case 95:
			goto st497
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st497
			}
		case data[p] >= 65:
			goto st497
		}
		goto st196
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 32:
			goto tr1128
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 46:
			goto st208
		case 59:
			goto tr1130
		case 92:
			goto st197
		case 95:
			goto st497
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1128
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st497
				}
			case data[p] >= 65:
				goto st497
			}
		default:
			goto st497
		}
		goto st196
	tr479:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st209
	tr484:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st209
	tr540:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st209
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
//line query/tokeniser.go:10909
		switch data[p] {
		case 32:
			goto tr503
		case 34:
			goto tr504
		case 39:
			goto tr505
		case 43:
			goto tr506
		case 45:
			goto tr506
		case 61:
			goto st210
		case 92:
			goto st197
		case 95:
			goto tr509
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr503
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr509
				}
			case data[p] >= 65:
				goto tr509
			}
		default:
			goto tr507
		}
		goto st196
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		switch data[p] {
		case 32:
			goto tr510
		case 34:
			goto tr511
		case 39:
			goto tr512
		case 43:
			goto tr513
		case 45:
			goto tr513
		case 92:
			goto st197
		case 95:
			goto tr515
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr510
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr515
				}
			case data[p] >= 65:
				goto tr515
			}
		default:
			goto tr514
		}
		goto st196
	tr480:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st211
	tr485:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st211
	tr541:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st211
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
//line query/tokeniser.go:11012
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 61:
			goto st212
		case 92:
			goto st197
		}
		goto st196
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch data[p] {
		case 32:
			goto tr517
		case 34:
			goto tr518
		case 39:
			goto tr519
		case 43:
			goto tr520
		case 45:
			goto tr520
		case 92:
			goto st197
		case 95:
			goto tr522
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr517
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr522
				}
			case data[p] >= 65:
				goto tr522
			}
		default:
			goto tr521
		}
		goto st196
	tr481:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st213
	tr486:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st213
	tr542:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st213
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
//line query/tokeniser.go:11096
		switch data[p] {
		case 32:
			goto tr523
		case 34:
			goto tr524
		case 39:
			goto tr525
		case 43:
			goto tr526
		case 45:
			goto tr526
		case 61:
			goto st214
		case 92:
			goto st197
		case 95:
			goto tr529
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr523
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr529
				}
			case data[p] >= 65:
				goto tr529
			}
		default:
			goto tr527
		}
		goto st196
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		switch data[p] {
		case 32:
			goto tr530
		case 34:
			goto tr531
		case 39:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr533
		case 92:
			goto st197
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr530
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr535
				}
			case data[p] >= 65:
				goto tr535
			}
		default:
			goto tr534
		}
		goto st196
	tr477:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st215
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
//line query/tokeniser.go:11181
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		case 95:
			goto st216
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st216
			}
		case data[p] >= 65:
			goto st216
		}
		goto st196
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 32:
			goto tr537
		case 33:
			goto tr538
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 46:
			goto st215
		case 60:
			goto tr540
		case 61:
			goto tr541
		case 62:
			goto tr542
		case 92:
			goto st197
		case 95:
			goto st216
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr537
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st216
				}
			case data[p] >= 65:
				goto st216
			}
		default:
			goto st216
		}
		goto st196
	tr1111:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
//line query/tokeniser.go:11255
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 78:
			goto st218
		case 92:
			goto st197
		case 110:
			goto st218
		}
		goto st196
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 68:
			goto st199
		case 92:
			goto st197
		case 100:
			goto st199
		}
		goto st196
	tr1112:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
//line query/tokeniser.go:11296
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 82:
			goto st220
		case 92:
			goto st197
		case 114:
			goto st220
		}
		goto st196
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 32:
			goto tr545
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr545
		}
		goto st196
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 73:
			goto st222
		case 92:
			goto st197
		case 105:
			goto st222
		}
		goto st196
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 84:
			goto st223
		case 92:
			goto st197
		case 116:
			goto st223
		}
		goto st196
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 72:
			goto st224
		case 92:
			goto st197
		case 104:
			goto st224
		}
		goto st196
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 73:
			goto st225
		case 92:
			goto st197
		case 105:
			goto st225
		}
		goto st196
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 78:
			goto st226
		case 92:
			goto st197
		case 110:
			goto st226
		}
		goto st196
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 32:
			goto st227
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st227
		}
		goto st196
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		switch data[p] {
		case 32:
			goto st227
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 43:
			goto tr552
		case 45:
			goto tr552
		case 92:
			goto st197
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr553
			}
		case data[p] >= 9:
			goto st227
		}
		goto st196
	tr552:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
//line query/tokeniser.go:11479
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st229
		}
		goto st196
	tr553:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line query/tokeniser.go:11505
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 46:
			goto st230
		case 72:
			goto st498
		case 77:
			goto st500
		case 78:
			goto st232
		case 83:
			goto st498
		case 85:
			goto st232
		case 92:
			goto st197
		case 104:
			goto st498
		case 109:
			goto st500
		case 110:
			goto st232
		case 115:
			goto st498
		case 117:
			goto st232
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st229
		}
		goto st196
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st231
		}
		goto st196
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 72:
			goto st498
		case 77:
			goto st500
		case 78:
			goto st232
		case 83:
			goto st498
		case 85:
			goto st232
		case 92:
			goto st197
		case 104:
			goto st498
		case 109:
			goto st500
		case 110:
			goto st232
		case 115:
			goto st498
		case 117:
			goto st232
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st231
		}
		goto st196
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		switch data[p] {
		case 32:
			goto tr1131
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 43:
			goto st228
		case 45:
			goto st228
		case 59:
			goto tr1133
		case 92:
			goto st197
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st229
			}
		case data[p] >= 9:
			goto tr1131
		}
		goto st196
	tr1131:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st499
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
//line query/tokeniser.go:11637
		switch data[p] {
		case 32:
			goto st499
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 59:
			goto st493
		case 92:
			goto st197
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st499
		}
		goto st196
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		switch data[p] {
		case 32:
			goto tr1131
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 43:
			goto st228
		case 45:
			goto st228
		case 59:
			goto tr1133
		case 83:
			goto st498
		case 92:
			goto st197
		case 115:
			goto st498
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st229
			}
		case data[p] >= 9:
			goto tr1131
		}
		goto st196
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 83:
			goto st498
		case 92:
			goto st197
		case 115:
			goto st498
		}
		goto st196
	tr1115:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
//line query/tokeniser.go:11715
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		case 124:
			goto st220
		}
		goto st196
	tr1116:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st234
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
//line query/tokeniser.go:11736
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		case 136:
			goto st235
		}
		goto st196
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		switch data[p] {
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 92:
			goto st197
		case 168:
			goto st220
		}
		goto st196
	tr470:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st501
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
//line query/tokeniser.go:11773
		switch data[p] {
		case 32:
			goto tr1135
		case 34:
			goto tr118
		case 39:
			goto tr467
		case 59:
			goto tr1136
		case 92:
			goto st197
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1135
		}
		goto st196
	tr453:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st236
	tr459:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st236
	tr567:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st236
	tr574:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st236
	tr581:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st236
	tr587:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st236
	tr594:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
//line query/tokeniser.go:11849
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st502
		}
		goto st23
	tr454:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st502
	tr460:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st502
	tr568:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st502
	tr575:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st502
	tr582:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st502
	tr588:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st502
	tr595:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st502
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
//line query/tokeniser.go:11919
		switch data[p] {
		case 32:
			goto tr1137
		case 34:
			goto tr54
		case 46:
			goto st237
		case 59:
			goto tr1139
		case 92:
			goto st186
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st502
			}
		case data[p] >= 9:
			goto tr1137
		}
		goto st23
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st503
		}
		goto st23
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 32:
			goto tr1137
		case 34:
			goto tr54
		case 59:
			goto tr1139
		case 92:
			goto st186
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st503
			}
		case data[p] >= 9:
			goto tr1137
		}
		goto st23
	tr455:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st504
	tr461:
//line query/tokeniser.rl:87
		mark = p
		goto st504
	tr570:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st504
	tr576:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st504
	tr583:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st504
	tr590:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st504
	tr596:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st504
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
//line query/tokeniser.go:12025
		switch data[p] {
		case 32:
			goto tr1140
		case 34:
			goto tr54
		case 46:
			goto tr1141
		case 59:
			goto tr1143
		case 92:
			goto st186
		case 95:
			goto st504
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1140
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st504
				}
			case data[p] >= 65:
				goto st504
			}
		default:
			goto st504
		}
		goto st23
	tr1141:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st238
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
//line query/tokeniser.go:12067
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		case 95:
			goto st505
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st505
			}
		case data[p] >= 65:
			goto st505
		}
		goto st23
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		switch data[p] {
		case 32:
			goto tr1144
		case 34:
			goto tr54
		case 46:
			goto st238
		case 59:
			goto tr1146
		case 92:
			goto st186
		case 95:
			goto st505
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1144
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st505
				}
			case data[p] >= 65:
				goto st505
			}
		default:
			goto st505
		}
		goto st23
	tr441:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st239
	tr446:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st239
	tr601:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st239
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
//line query/tokeniser.go:12155
		switch data[p] {
		case 32:
			goto tr564
		case 34:
			goto tr565
		case 39:
			goto tr566
		case 43:
			goto tr567
		case 45:
			goto tr567
		case 61:
			goto st240
		case 92:
			goto st186
		case 95:
			goto tr570
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr564
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr570
				}
			case data[p] >= 65:
				goto tr570
			}
		default:
			goto tr568
		}
		goto st23
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 32:
			goto tr571
		case 34:
			goto tr572
		case 39:
			goto tr573
		case 43:
			goto tr574
		case 45:
			goto tr574
		case 92:
			goto st186
		case 95:
			goto tr576
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr571
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr576
				}
			case data[p] >= 65:
				goto tr576
			}
		default:
			goto tr575
		}
		goto st23
	tr442:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st241
	tr447:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st241
	tr602:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st241
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
//line query/tokeniser.go:12258
		switch data[p] {
		case 34:
			goto tr54
		case 61:
			goto st242
		case 92:
			goto st186
		}
		goto st23
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		switch data[p] {
		case 32:
			goto tr578
		case 34:
			goto tr579
		case 39:
			goto tr580
		case 43:
			goto tr581
		case 45:
			goto tr581
		case 92:
			goto st186
		case 95:
			goto tr583
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr578
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr583
				}
			case data[p] >= 65:
				goto tr583
			}
		default:
			goto tr582
		}
		goto st23
	tr443:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st243
	tr448:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st243
	tr603:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st243
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
//line query/tokeniser.go:12340
		switch data[p] {
		case 32:
			goto tr584
		case 34:
			goto tr585
		case 39:
			goto tr586
		case 43:
			goto tr587
		case 45:
			goto tr587
		case 61:
			goto st244
		case 92:
			goto st186
		case 95:
			goto tr590
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr584
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr590
				}
			case data[p] >= 65:
				goto tr590
			}
		default:
			goto tr588
		}
		goto st23
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 32:
			goto tr591
		case 34:
			goto tr592
		case 39:
			goto tr593
		case 43:
			goto tr594
		case 45:
			goto tr594
		case 92:
			goto st186
		case 95:
			goto tr596
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr591
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr596
				}
			case data[p] >= 65:
				goto tr596
			}
		default:
			goto tr595
		}
		goto st23
	tr439:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st245
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
//line query/tokeniser.go:12425
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		case 95:
			goto st246
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st246
			}
		case data[p] >= 65:
			goto st246
		}
		goto st23
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		switch data[p] {
		case 32:
			goto tr598
		case 33:
			goto tr599
		case 34:
			goto tr54
		case 46:
			goto st245
		case 60:
			goto tr601
		case 61:
			goto tr602
		case 62:
			goto tr603
		case 92:
			goto st186
		case 95:
			goto st246
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr598
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st246
				}
			case data[p] >= 65:
				goto st246
			}
		default:
			goto st246
		}
		goto st23
	tr1096:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st247
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
//line query/tokeniser.go:12495
		switch data[p] {
		case 34:
			goto tr54
		case 78:
			goto st248
		case 92:
			goto st186
		case 110:
			goto st248
		}
		goto st23
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		switch data[p] {
		case 34:
			goto tr54
		case 68:
			goto st188
		case 92:
			goto st186
		case 100:
			goto st188
		}
		goto st23
	tr1097:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st249
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
//line query/tokeniser.go:12532
		switch data[p] {
		case 34:
			goto tr54
		case 82:
			goto st250
		case 92:
			goto st186
		case 114:
			goto st250
		}
		goto st23
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 32:
			goto tr606
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr606
		}
		goto st23
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 34:
			goto tr54
		case 73:
			goto st252
		case 92:
			goto st186
		case 105:
			goto st252
		}
		goto st23
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		switch data[p] {
		case 34:
			goto tr54
		case 84:
			goto st253
		case 92:
			goto st186
		case 116:
			goto st253
		}
		goto st23
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		switch data[p] {
		case 34:
			goto tr54
		case 72:
			goto st254
		case 92:
			goto st186
		case 104:
			goto st254
		}
		goto st23
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		switch data[p] {
		case 34:
			goto tr54
		case 73:
			goto st255
		case 92:
			goto st186
		case 105:
			goto st255
		}
		goto st23
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 34:
			goto tr54
		case 78:
			goto st256
		case 92:
			goto st186
		case 110:
			goto st256
		}
		goto st23
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		switch data[p] {
		case 32:
			goto st257
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st257
		}
		goto st23
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		switch data[p] {
		case 32:
			goto st257
		case 34:
			goto tr54
		case 43:
			goto tr613
		case 45:
			goto tr613
		case 92:
			goto st186
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr614
			}
		case data[p] >= 9:
			goto st257
		}
		goto st23
	tr613:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st258
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
//line query/tokeniser.go:12697
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st259
		}
		goto st23
	tr614:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st259
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
//line query/tokeniser.go:12721
		switch data[p] {
		case 34:
			goto tr54
		case 46:
			goto st260
		case 72:
			goto st506
		case 77:
			goto st508
		case 78:
			goto st262
		case 83:
			goto st506
		case 85:
			goto st262
		case 92:
			goto st186
		case 104:
			goto st506
		case 109:
			goto st508
		case 110:
			goto st262
		case 115:
			goto st506
		case 117:
			goto st262
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st259
		}
		goto st23
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st261
		}
		goto st23
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 34:
			goto tr54
		case 72:
			goto st506
		case 77:
			goto st508
		case 78:
			goto st262
		case 83:
			goto st506
		case 85:
			goto st262
		case 92:
			goto st186
		case 104:
			goto st506
		case 109:
			goto st508
		case 110:
			goto st262
		case 115:
			goto st506
		case 117:
			goto st262
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st261
		}
		goto st23
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		switch data[p] {
		case 32:
			goto tr1147
		case 34:
			goto tr54
		case 43:
			goto st258
		case 45:
			goto st258
		case 59:
			goto tr1149
		case 92:
			goto st186
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st259
			}
		case data[p] >= 9:
			goto tr1147
		}
		goto st23
	tr1147:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st507
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
//line query/tokeniser.go:12845
		switch data[p] {
		case 32:
			goto st507
		case 34:
			goto tr54
		case 59:
			goto st488
		case 92:
			goto st186
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st507
		}
		goto st23
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		switch data[p] {
		case 32:
			goto tr1147
		case 34:
			goto tr54
		case 43:
			goto st258
		case 45:
			goto st258
		case 59:
			goto tr1149
		case 83:
			goto st506
		case 92:
			goto st186
		case 115:
			goto st506
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st259
			}
		case data[p] >= 9:
			goto tr1147
		}
		goto st23
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 34:
			goto tr54
		case 83:
			goto st506
		case 92:
			goto st186
		case 115:
			goto st506
		}
		goto st23
	tr1100:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st263
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
//line query/tokeniser.go:12917
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		case 124:
			goto st250
		}
		goto st23
	tr1101:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st264
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
//line query/tokeniser.go:12936
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		case 136:
			goto st265
		}
		goto st23
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		switch data[p] {
		case 34:
			goto tr54
		case 92:
			goto st186
		case 168:
			goto st250
		}
		goto st23
	tr40:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st266
	tr46:
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st266
	tr817:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st266
	tr821:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st266
	tr825:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st266
	tr828:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st266
	tr832:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line query/tokeniser.go:13005
		switch data[p] {
		case 39:
			goto tr370
		case 92:
			goto tr623
		}
		goto tr622
	tr622:
//line query/tokeniser.rl:87
		mark = p
		goto st267
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
//line query/tokeniser.go:13022
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		goto st267
	tr623:
//line query/tokeniser.rl:87
		mark = p
		goto st268
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
//line query/tokeniser.go:13039
		switch data[p] {
		case 39:
			goto tr626
		case 92:
			goto st268
		}
		goto st267
	tr626:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st509
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
//line query/tokeniser.go:13056
		switch data[p] {
		case 32:
			goto tr1151
		case 39:
			goto tr122
		case 59:
			goto tr1152
		case 92:
			goto st268
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1151
		}
		goto st267
	tr1162:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st510
	tr1151:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st510
	tr1197:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st510
	tr1200:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st510
	tr1204:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st510
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
//line query/tokeniser.go:13114
		switch data[p] {
		case 32:
			goto st510
		case 38:
			goto tr1154
		case 39:
			goto tr122
		case 59:
			goto st512
		case 65:
			goto tr1156
		case 79:
			goto tr1157
		case 87:
			goto st333
		case 92:
			goto st268
		case 94:
			goto tr1159
		case 97:
			goto tr1156
		case 111:
			goto tr1157
		case 119:
			goto st333
		case 124:
			goto tr1160
		case 226:
			goto tr1161
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st510
		}
		goto st267
	tr1154:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st269
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
//line query/tokeniser.go:13158
		switch data[p] {
		case 38:
			goto st270
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		goto st267
	tr1159:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st270
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
//line query/tokeniser.go:13177
		switch data[p] {
		case 32:
			goto tr628
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr628
		}
		goto st267
	tr628:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st271
	tr799:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st271
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
//line query/tokeniser.go:13203
		switch data[p] {
		case 32:
			goto st271
		case 39:
			goto tr122
		case 92:
			goto st268
		case 95:
			goto tr630
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st271
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr630
			}
		default:
			goto tr630
		}
		goto st267
	tr630:
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st272
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
//line query/tokeniser.go:13238
		switch data[p] {
		case 32:
			goto tr631
		case 33:
			goto tr632
		case 39:
			goto tr122
		case 46:
			goto tr633
		case 60:
			goto tr635
		case 61:
			goto tr636
		case 62:
			goto tr637
		case 92:
			goto st268
		case 95:
			goto st272
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr631
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st272
				}
			case data[p] >= 65:
				goto st272
			}
		default:
			goto st272
		}
		goto st267
	tr631:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st273
	tr791:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st273
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
//line query/tokeniser.go:13296
		switch data[p] {
		case 32:
			goto st273
		case 33:
			goto tr639
		case 39:
			goto tr122
		case 60:
			goto tr640
		case 61:
			goto tr641
		case 62:
			goto tr642
		case 92:
			goto st268
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st273
		}
		goto st267
	tr632:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st274
	tr639:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st274
	tr792:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st274
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
//line query/tokeniser.go:13344
		switch data[p] {
		case 39:
			goto tr122
		case 61:
			goto st275
		case 92:
			goto st268
		}
		goto st267
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 32:
			goto tr644
		case 34:
			goto tr645
		case 39:
			goto tr646
		case 43:
			goto tr647
		case 45:
			goto tr647
		case 92:
			goto st268
		case 95:
			goto tr649
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
					goto tr649
				}
			case data[p] >= 65:
				goto tr649
			}
		default:
			goto tr648
		}
		goto st267
	tr644:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st276
	tr757:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st276
	tr764:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st276
	tr771:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st276
	tr777:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st276
	tr784:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st276
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
//line query/tokeniser.go:13422
		switch data[p] {
		case 32:
			goto st276
		case 34:
			goto tr651
		case 39:
			goto tr652
		case 43:
			goto tr653
		case 45:
			goto tr653
		case 92:
			goto st268
		case 95:
			goto tr655
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st276
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr655
				}
			case data[p] >= 65:
				goto tr655
			}
		default:
			goto tr654
		}
		goto st267
	tr645:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st277
	tr651:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st277
	tr758:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st277
	tr765:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st277
	tr772:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st277
	tr778:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st277
	tr785:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
		goto st277
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
//line query/tokeniser.go:13502
		switch data[p] {
		case 34:
			goto tr657
		case 39:
			goto tr154
		case 92:
			goto tr658
		}
		goto tr656
	tr656:
//line query/tokeniser.rl:87
		mark = p
		goto st278
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
//line query/tokeniser.go:13521
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		goto st278
	tr657:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st511
	tr660:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st511
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
//line query/tokeniser.go:13546
		switch data[p] {
		case 32:
			goto tr1162
		case 39:
			goto tr122
		case 59:
			goto tr1163
		case 92:
			goto st268
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1162
		}
		goto st267
	tr1163:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st512
	tr1152:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st512
	tr1199:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st512
	tr1203:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st512
	tr1206:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st512
	tr1209:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st512
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
//line query/tokeniser.go:13612
		switch data[p] {
		case 32:
			goto st512
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st512
		}
		goto st267
	tr658:
//line query/tokeniser.rl:87
		mark = p
		goto st279
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
//line query/tokeniser.go:13634
		switch data[p] {
		case 34:
			goto tr662
		case 39:
			goto tr663
		case 92:
			goto st279
		}
		goto st278
	tr662:
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st513
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
//line query/tokeniser.go:13653
		switch data[p] {
		case 32:
			goto tr1164
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 59:
			goto tr1165
		case 92:
			goto st279
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1164
		}
		goto st278
	tr1164:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st514
	tr1175:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st514
	tr1193:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st514
	tr1179:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st514
	tr1182:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st514
	tr1186:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st514
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
//line query/tokeniser.go:13721
		switch data[p] {
		case 32:
			goto st514
		case 34:
			goto tr660
		case 38:
			goto tr1167
		case 39:
			goto tr119
		case 59:
			goto st516
		case 65:
			goto tr1169
		case 79:
			goto tr1170
		case 87:
			goto st303
		case 92:
			goto st279
		case 94:
			goto tr1172
		case 97:
			goto tr1169
		case 111:
			goto tr1170
		case 119:
			goto st303
		case 124:
			goto tr1173
		case 226:
			goto tr1174
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st514
		}
		goto st278
	tr1167:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st280
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
//line query/tokeniser.go:13767
		switch data[p] {
		case 34:
			goto tr660
		case 38:
			goto st281
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		goto st278
	tr1172:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st281
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
//line query/tokeniser.go:13788
		switch data[p] {
		case 32:
			goto tr665
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr665
		}
		goto st278
	tr665:
//line query/tokeniser.rl:166
		commit(ttConjunction)
		goto st282
	tr738:
//line query/tokeniser.rl:170
		commit(ttDisjunction)
		goto st282
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
//line query/tokeniser.go:13816
		switch data[p] {
		case 32:
			goto st282
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		case 95:
			goto tr667
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st282
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr667
			}
		default:
			goto tr667
		}
		goto st278
	tr667:
//line query/tokeniser.rl:206
		propose(ttPredicate)
//line query/tokeniser.rl:87
		mark = p
		goto st283
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
//line query/tokeniser.go:13853
		switch data[p] {
		case 32:
			goto tr668
		case 33:
			goto tr669
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 46:
			goto tr670
		case 60:
			goto tr672
		case 61:
			goto tr673
		case 62:
			goto tr674
		case 92:
			goto st279
		case 95:
			goto st283
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr668
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st283
				}
			case data[p] >= 65:
				goto st283
			}
		default:
			goto st283
		}
		goto st278
	tr668:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st284
	tr730:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
		goto st284
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
//line query/tokeniser.go:13913
		switch data[p] {
		case 32:
			goto st284
		case 33:
			goto tr676
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 60:
			goto tr677
		case 61:
			goto tr678
		case 62:
			goto tr679
		case 92:
			goto st279
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st284
		}
		goto st278
	tr669:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st285
	tr676:
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st285
	tr731:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:156
		propose(ttEq)
		goto st285
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
//line query/tokeniser.go:13963
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 61:
			goto st286
		case 92:
			goto st279
		}
		goto st278
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		switch data[p] {
		case 32:
			goto tr681
		case 34:
			goto tr682
		case 39:
			goto tr683
		case 43:
			goto tr684
		case 45:
			goto tr684
		case 92:
			goto st279
		case 95:
			goto tr686
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr681
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr686
				}
			case data[p] >= 65:
				goto tr686
			}
		default:
			goto tr685
		}
		goto st278
	tr681:
//line query/tokeniser.rl:156
		commit(ttEq)
		goto st287
	tr696:
//line query/tokeniser.rl:158
		commit(ttEq)
		goto st287
	tr703:
//line query/tokeniser.rl:160
		commit(ttEq)
		goto st287
	tr710:
//line query/tokeniser.rl:155
		commit(ttEq)
		goto st287
	tr716:
//line query/tokeniser.rl:157
		commit(ttEq)
		goto st287
	tr723:
//line query/tokeniser.rl:159
		commit(ttEq)
		goto st287
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
//line query/tokeniser.go:14043
		switch data[p] {
		case 32:
			goto st287
		case 34:
			goto tr688
		case 39:
			goto tr689
		case 43:
			goto tr690
		case 45:
			goto tr690
		case 92:
			goto st279
		case 95:
			goto tr692
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st287
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr692
				}
			case data[p] >= 65:
				goto tr692
			}
		default:
			goto tr691
		}
		goto st278
	tr682:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st515
	tr688:
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st515
	tr697:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st515
	tr704:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st515
	tr711:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st515
	tr717:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st515
	tr724:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:189
		propose(ttStringLiteral)
//line query/tokeniser.rl:192
		setText(ttStringLiteral)
		goto st515
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
//line query/tokeniser.go:14137
		switch data[p] {
		case 32:
			goto tr1175
		case 34:
			goto tr657
		case 39:
			goto tr154
		case 59:
			goto tr1176
		case 92:
			goto tr658
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1175
		}
		goto tr656
	tr1165:
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st516
	tr1176:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:194
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st516
	tr1194:
//line query/tokeniser.rl:186
		commit(ttStringLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st516
	tr1181:
//line query/tokeniser.rl:177
		setText(ttNumericLiteral)
//line query/tokeniser.rl:178
		commit(ttNumericLiteral)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st516
	tr1185:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st516
	tr1188:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:209
		commit(ttPredicate)
		goto st516
	tr1191:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st516
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
//line query/tokeniser.go:14213
		switch data[p] {
		case 32:
			goto st516
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st516
		}
		goto st278
	tr683:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st517
	tr689:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st517
	tr698:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st517
	tr705:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st517
	tr712:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st517
	tr718:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st517
	tr725:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st517
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
//line query/tokeniser.go:14287
		switch data[p] {
		case 32:
			goto tr1177
		case 34:
			goto tr114
		case 39:
			goto tr115
		case 59:
			goto tr1178
		case 92:
			goto tr116
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1177
		}
		goto tr113
	tr684:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st288
	tr690:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st288
	tr699:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st288
	tr706:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st288
	tr713:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st288
	tr719:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st288
	tr726:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st288
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
//line query/tokeniser.go:14363
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st518
		}
		goto st278
	tr685:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st518
	tr691:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st518
	tr700:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st518
	tr707:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st518
	tr714:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st518
	tr720:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st518
	tr727:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st518
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
//line query/tokeniser.go:14435
		switch data[p] {
		case 32:
			goto tr1179
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 46:
			goto st289
		case 59:
			goto tr1181
		case 92:
			goto st279
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st518
			}
		case data[p] >= 9:
			goto tr1179
		}
		goto st278
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st519
		}
		goto st278
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		switch data[p] {
		case 32:
			goto tr1179
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 59:
			goto tr1181
		case 92:
			goto st279
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st519
			}
		case data[p] >= 9:
			goto tr1179
		}
		goto st278
	tr686:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st520
	tr692:
//line query/tokeniser.rl:87
		mark = p
		goto st520
	tr702:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st520
	tr708:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st520
	tr715:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st520
	tr722:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st520
	tr728:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st520
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
//line query/tokeniser.go:14547
		switch data[p] {
		case 32:
			goto tr1182
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 46:
			goto tr1183
		case 59:
			goto tr1185
		case 92:
			goto st279
		case 95:
			goto st520
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1182
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st520
				}
			case data[p] >= 65:
				goto st520
			}
		default:
			goto st520
		}
		goto st278
	tr1183:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line query/tokeniser.go:14591
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		case 95:
			goto st521
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st521
			}
		case data[p] >= 65:
			goto st521
		}
		goto st278
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		switch data[p] {
		case 32:
			goto tr1186
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 46:
			goto st290
		case 59:
			goto tr1188
		case 92:
			goto st279
		case 95:
			goto st521
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1186
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st521
				}
			case data[p] >= 65:
				goto st521
			}
		default:
			goto st521
		}
		goto st278
	tr672:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st291
	tr677:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st291
	tr733:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st291
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
//line query/tokeniser.go:14683
		switch data[p] {
		case 32:
			goto tr696
		case 34:
			goto tr697
		case 39:
			goto tr698
		case 43:
			goto tr699
		case 45:
			goto tr699
		case 61:
			goto st292
		case 92:
			goto st279
		case 95:
			goto tr702
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr696
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr702
				}
			case data[p] >= 65:
				goto tr702
			}
		default:
			goto tr700
		}
		goto st278
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		switch data[p] {
		case 32:
			goto tr703
		case 34:
			goto tr704
		case 39:
			goto tr705
		case 43:
			goto tr706
		case 45:
			goto tr706
		case 92:
			goto st279
		case 95:
			goto tr708
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr703
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr708
				}
			case data[p] >= 65:
				goto tr708
			}
		default:
			goto tr707
		}
		goto st278
	tr673:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st293
	tr678:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st293
	tr734:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st293
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
//line query/tokeniser.go:14786
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 61:
			goto st294
		case 92:
			goto st279
		}
		goto st278
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		switch data[p] {
		case 32:
			goto tr710
		case 34:
			goto tr711
		case 39:
			goto tr712
		case 43:
			goto tr713
		case 45:
			goto tr713
		case 92:
			goto st279
		case 95:
			goto tr715
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr710
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr715
				}
			case data[p] >= 65:
				goto tr715
			}
		default:
			goto tr714
		}
		goto st278
	tr674:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st295
	tr679:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st295
	tr735:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st295
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
//line query/tokeniser.go:14870
		switch data[p] {
		case 32:
			goto tr716
		case 34:
			goto tr717
		case 39:
			goto tr718
		case 43:
			goto tr719
		case 45:
			goto tr719
		case 61:
			goto st296
		case 92:
			goto st279
		case 95:
			goto tr722
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr716
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr722
				}
			case data[p] >= 65:
				goto tr722
			}
		default:
			goto tr720
		}
		goto st278
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		switch data[p] {
		case 32:
			goto tr723
		case 34:
			goto tr724
		case 39:
			goto tr725
		case 43:
			goto tr726
		case 45:
			goto tr726
		case 92:
			goto st279
		case 95:
			goto tr728
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr723
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr728
				}
			case data[p] >= 65:
				goto tr728
			}
		default:
			goto tr727
		}
		goto st278
	tr670:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st297
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
//line query/tokeniser.go:14955
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		case 95:
			goto st298
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st298
			}
		case data[p] >= 65:
			goto st298
		}
		goto st278
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 32:
			goto tr730
		case 33:
			goto tr731
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 46:
			goto st297
		case 60:
			goto tr733
		case 61:
			goto tr734
		case 62:
			goto tr735
		case 92:
			goto st279
		case 95:
			goto st298
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr730
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st298
				}
			case data[p] >= 65:
				goto st298
			}
		default:
			goto st298
		}
		goto st278
	tr1169:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st299
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
//line query/tokeniser.go:15029
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 78:
			goto st300
		case 92:
			goto st279
		case 110:
			goto st300
		}
		goto st278
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 68:
			goto st281
		case 92:
			goto st279
		case 100:
			goto st281
		}
		goto st278
	tr1170:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st301
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
//line query/tokeniser.go:15070
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 82:
			goto st302
		case 92:
			goto st279
		case 114:
			goto st302
		}
		goto st278
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		switch data[p] {
		case 32:
			goto tr738
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr738
		}
		goto st278
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 73:
			goto st304
		case 92:
			goto st279
		case 105:
			goto st304
		}
		goto st278
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 84:
			goto st305
		case 92:
			goto st279
		case 116:
			goto st305
		}
		goto st278
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 72:
			goto st306
		case 92:
			goto st279
		case 104:
			goto st306
		}
		goto st278
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 73:
			goto st307
		case 92:
			goto st279
		case 105:
			goto st307
		}
		goto st278
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 78:
			goto st308
		case 92:
			goto st279
		case 110:
			goto st308
		}
		goto st278
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		switch data[p] {
		case 32:
			goto st309
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st309
		}
		goto st278
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		switch data[p] {
		case 32:
			goto st309
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 43:
			goto tr745
		case 45:
			goto tr745
		case 92:
			goto st279
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr746
			}
		case data[p] >= 9:
			goto st309
		}
		goto st278
	tr745:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st310
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
//line query/tokeniser.go:15253
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st311
		}
		goto st278
	tr746:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st311
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
//line query/tokeniser.go:15279
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 46:
			goto st312
		case 72:
			goto st522
		case 77:
			goto st524
		case 78:
			goto st314
		case 83:
			goto st522
		case 85:
			goto st314
		case 92:
			goto st279
		case 104:
			goto st522
		case 109:
			goto st524
		case 110:
			goto st314
		case 115:
			goto st522
		case 117:
			goto st314
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st311
		}
		goto st278
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st313
		}
		goto st278
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 72:
			goto st522
		case 77:
			goto st524
		case 78:
			goto st314
		case 83:
			goto st522
		case 85:
			goto st314
		case 92:
			goto st279
		case 104:
			goto st522
		case 109:
			goto st524
		case 110:
			goto st314
		case 115:
			goto st522
		case 117:
			goto st314
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st313
		}
		goto st278
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		switch data[p] {
		case 32:
			goto tr1189
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 43:
			goto st310
		case 45:
			goto st310
		case 59:
			goto tr1191
		case 92:
			goto st279
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st311
			}
		case data[p] >= 9:
			goto tr1189
		}
		goto st278
	tr1189:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st523
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
//line query/tokeniser.go:15411
		switch data[p] {
		case 32:
			goto st523
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 59:
			goto st516
		case 92:
			goto st279
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st523
		}
		goto st278
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
		switch data[p] {
		case 32:
			goto tr1189
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 43:
			goto st310
		case 45:
			goto st310
		case 59:
			goto tr1191
		case 83:
			goto st522
		case 92:
			goto st279
		case 115:
			goto st522
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st311
			}
		case data[p] >= 9:
			goto tr1189
		}
		goto st278
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 83:
			goto st522
		case 92:
			goto st279
		case 115:
			goto st522
		}
		goto st278
	tr1173:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st315
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
//line query/tokeniser.go:15489
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		case 124:
			goto st302
		}
		goto st278
	tr1174:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st316
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
//line query/tokeniser.go:15510
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		case 136:
			goto st317
		}
		goto st278
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		switch data[p] {
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 92:
			goto st279
		case 168:
			goto st302
		}
		goto st278
	tr663:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
		goto st525
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
//line query/tokeniser.go:15547
		switch data[p] {
		case 32:
			goto tr1193
		case 34:
			goto tr660
		case 39:
			goto tr119
		case 59:
			goto tr1194
		case 92:
			goto st279
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1193
		}
		goto st278
	tr646:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st526
	tr652:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st526
	tr759:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st526
	tr766:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st526
	tr773:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st526
	tr779:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st526
	tr786:
//line query/tokeniser.rl:184
		setText(ttStringLiteral)
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:181
		propose(ttStringLiteral)
		goto st526
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
//line query/tokeniser.go:15623
		switch data[p] {
		case 32:
			goto tr1195
		case 39:
			goto tr370
		case 59:
			goto tr1196
		case 92:
			goto tr371
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1195
		}
		goto tr369
	tr647:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st318
	tr653:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st318
	tr760:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st318
	tr767:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st318
	tr774:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st318
	tr780:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st318
	tr787:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st318
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
//line query/tokeniser.go:15697
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st527
		}
		goto st267
	tr648:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st527
	tr654:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st527
	tr761:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st527
	tr768:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st527
	tr775:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st527
	tr781:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st527
	tr788:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:176
		propose(ttNumericLiteral)
		goto st527
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
//line query/tokeniser.go:15767
		switch data[p] {
		case 32:
			goto tr1197
		case 39:
			goto tr122
		case 46:
			goto st319
		case 59:
			goto tr1199
		case 92:
			goto st268
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st527
			}
		case data[p] >= 9:
			goto tr1197
		}
		goto st267
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st528
		}
		goto st267
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
		switch data[p] {
		case 32:
			goto tr1197
		case 39:
			goto tr122
		case 59:
			goto tr1199
		case 92:
			goto st268
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st528
			}
		case data[p] >= 9:
			goto tr1197
		}
		goto st267
	tr649:
//line query/tokeniser.rl:156
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st529
	tr655:
//line query/tokeniser.rl:87
		mark = p
		goto st529
	tr763:
//line query/tokeniser.rl:158
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st529
	tr769:
//line query/tokeniser.rl:160
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st529
	tr776:
//line query/tokeniser.rl:155
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st529
	tr783:
//line query/tokeniser.rl:157
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st529
	tr789:
//line query/tokeniser.rl:159
		commit(ttEq)
//line query/tokeniser.rl:87
		mark = p
		goto st529
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
//line query/tokeniser.go:15873
		switch data[p] {
		case 32:
			goto tr1200
		case 39:
			goto tr122
		case 46:
			goto tr1201
		case 59:
			goto tr1203
		case 92:
			goto st268
		case 95:
			goto st529
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1200
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st529
				}
			case data[p] >= 65:
				goto st529
			}
		default:
			goto st529
		}
		goto st267
	tr1201:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st320
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
//line query/tokeniser.go:15915
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		case 95:
			goto st530
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st530
			}
		case data[p] >= 65:
			goto st530
		}
		goto st267
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
		switch data[p] {
		case 32:
			goto tr1204
		case 39:
			goto tr122
		case 46:
			goto st320
		case 59:
			goto tr1206
		case 92:
			goto st268
		case 95:
			goto st530
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
					goto st530
				}
			case data[p] >= 65:
				goto st530
			}
		default:
			goto st530
		}
		goto st267
	tr635:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st321
	tr640:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st321
	tr794:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st321
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
//line query/tokeniser.go:16003
		switch data[p] {
		case 32:
			goto tr757
		case 34:
			goto tr758
		case 39:
			goto tr759
		case 43:
			goto tr760
		case 45:
			goto tr760
		case 61:
			goto st322
		case 92:
			goto st268
		case 95:
			goto tr763
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr757
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr763
				}
			case data[p] >= 65:
				goto tr763
			}
		default:
			goto tr761
		}
		goto st267
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		switch data[p] {
		case 32:
			goto tr764
		case 34:
			goto tr765
		case 39:
			goto tr766
		case 43:
			goto tr767
		case 45:
			goto tr767
		case 92:
			goto st268
		case 95:
			goto tr769
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr764
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr769
				}
			case data[p] >= 65:
				goto tr769
			}
		default:
			goto tr768
		}
		goto st267
	tr636:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st323
	tr641:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st323
	tr795:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st323
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
//line query/tokeniser.go:16106
		switch data[p] {
		case 39:
			goto tr122
		case 61:
			goto st324
		case 92:
			goto st268
		}
		goto st267
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 32:
			goto tr771
		case 34:
			goto tr772
		case 39:
			goto tr773
		case 43:
			goto tr774
		case 45:
			goto tr774
		case 92:
			goto st268
		case 95:
			goto tr776
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr771
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr776
				}
			case data[p] >= 65:
				goto tr776
			}
		default:
			goto tr775
		}
		goto st267
	tr637:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st325
	tr642:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st325
	tr796:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st325
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
//line query/tokeniser.go:16188
		switch data[p] {
		case 32:
			goto tr777
		case 34:
			goto tr778
		case 39:
			goto tr779
		case 43:
			goto tr780
		case 45:
			goto tr780
		case 61:
			goto st326
		case 92:
			goto st268
		case 95:
			goto tr783
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr777
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr783
				}
			case data[p] >= 65:
				goto tr783
			}
		default:
			goto tr781
		}
		goto st267
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		switch data[p] {
		case 32:
			goto tr784
		case 34:
			goto tr785
		case 39:
			goto tr786
		case 43:
			goto tr787
		case 45:
			goto tr787
		case 92:
			goto st268
		case 95:
			goto tr789
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr784
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr789
				}
			case data[p] >= 65:
				goto tr789
			}
		default:
			goto tr788
		}
		goto st267
	tr633:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st327
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
//line query/tokeniser.go:16273
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		case 95:
			goto st328
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st328
			}
		case data[p] >= 65:
			goto st328
		}
		goto st267
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		switch data[p] {
		case 32:
			goto tr791
		case 33:
			goto tr792
		case 39:
			goto tr122
		case 46:
			goto st327
		case 60:
			goto tr794
		case 61:
			goto tr795
		case 62:
			goto tr796
		case 92:
			goto st268
		case 95:
			goto st328
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr791
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st328
				}
			case data[p] >= 65:
				goto st328
			}
		default:
			goto st328
		}
		goto st267
	tr1156:
//line query/tokeniser.rl:165
		propose(ttConjunction)
		goto st329
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
//line query/tokeniser.go:16343
		switch data[p] {
		case 39:
			goto tr122
		case 78:
			goto st330
		case 92:
			goto st268
		case 110:
			goto st330
		}
		goto st267
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		switch data[p] {
		case 39:
			goto tr122
		case 68:
			goto st270
		case 92:
			goto st268
		case 100:
			goto st270
		}
		goto st267
	tr1157:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st331
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
//line query/tokeniser.go:16380
		switch data[p] {
		case 39:
			goto tr122
		case 82:
			goto st332
		case 92:
			goto st268
		case 114:
			goto st332
		}
		goto st267
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 32:
			goto tr799
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr799
		}
		goto st267
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 39:
			goto tr122
		case 73:
			goto st334
		case 92:
			goto st268
		case 105:
			goto st334
		}
		goto st267
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		switch data[p] {
		case 39:
			goto tr122
		case 84:
			goto st335
		case 92:
			goto st268
		case 116:
			goto st335
		}
		goto st267
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		switch data[p] {
		case 39:
			goto tr122
		case 72:
			goto st336
		case 92:
			goto st268
		case 104:
			goto st336
		}
		goto st267
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		switch data[p] {
		case 39:
			goto tr122
		case 73:
			goto st337
		case 92:
			goto st268
		case 105:
			goto st337
		}
		goto st267
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		switch data[p] {
		case 39:
			goto tr122
		case 78:
			goto st338
		case 92:
			goto st268
		case 110:
			goto st338
		}
		goto st267
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		switch data[p] {
		case 32:
			goto st339
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st339
		}
		goto st267
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		switch data[p] {
		case 32:
			goto st339
		case 39:
			goto tr122
		case 43:
			goto tr806
		case 45:
			goto tr806
		case 92:
			goto st268
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr807
			}
		case data[p] >= 9:
			goto st339
		}
		goto st267
	tr806:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st340
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
//line query/tokeniser.go:16545
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st341
		}
		goto st267
	tr807:
//line query/tokeniser.rl:223
		propose(ttWithinClause)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:218
		propose(ttDuration)
		goto st341
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
//line query/tokeniser.go:16569
		switch data[p] {
		case 39:
			goto tr122
		case 46:
			goto st342
		case 72:
			goto st531
		case 77:
			goto st533
		case 78:
			goto st344
		case 83:
			goto st531
		case 85:
			goto st344
		case 92:
			goto st268
		case 104:
			goto st531
		case 109:
			goto st533
		case 110:
			goto st344
		case 115:
			goto st531
		case 117:
			goto st344
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st341
		}
		goto st267
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st343
		}
		goto st267
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		switch data[p] {
		case 39:
			goto tr122
		case 72:
			goto st531
		case 77:
			goto st533
		case 78:
			goto st344
		case 83:
			goto st531
		case 85:
			goto st344
		case 92:
			goto st268
		case 104:
			goto st531
		case 109:
			goto st533
		case 110:
			goto st344
		case 115:
			goto st531
		case 117:
			goto st344
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st343
		}
		goto st267
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		switch data[p] {
		case 32:
			goto tr1207
		case 39:
			goto tr122
		case 43:
			goto st340
		case 45:
			goto st340
		case 59:
			goto tr1209
		case 92:
			goto st268
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st341
			}
		case data[p] >= 9:
			goto tr1207
		}
		goto st267
	tr1207:
//line query/tokeniser.rl:219
		setText(ttDuration)
//line query/tokeniser.rl:220
		commit(ttDuration)
//line query/tokeniser.rl:224
		commit(ttWithinClause)
		goto st532
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
//line query/tokeniser.go:16693
		switch data[p] {
		case 32:
			goto st532
		case 39:
			goto tr122
		case 59:
			goto st512
		case 92:
			goto st268
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st532
		}
		goto st267
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		switch data[p] {
		case 32:
			goto tr1207
		case 39:
			goto tr122
		case 43:
			goto st340
		case 45:
			goto st340
		case 59:
			goto tr1209
		case 83:
			goto st531
		case 92:
			goto st268
		case 115:
			goto st531
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st341
			}
		case data[p] >= 9:
			goto tr1207
		}
		goto st267
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		switch data[p] {
		case 39:
			goto tr122
		case 83:
			goto st531
		case 92:
			goto st268
		case 115:
			goto st531
		}
		goto st267
	tr1160:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st345
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
//line query/tokeniser.go:16765
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		case 124:
			goto st332
		}
		goto st267
	tr1161:
//line query/tokeniser.rl:169
		propose(ttDisjunction)
		goto st346
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
//line query/tokeniser.go:16784
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		case 136:
			goto st347
		}
		goto st267
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		switch data[p] {
		case 39:
			goto tr122
		case 92:
			goto st268
		case 168:
			goto st332
		}
		goto st267
	tr29:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st348
	tr34:
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st348
	tr837:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:158
		propose(ttEq)
//line query/tokeniser.rl:160
		propose(ttEq)
		goto st348
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
//line query/tokeniser.go:16841
		switch data[p] {
		case 32:
			goto tr815
		case 34:
			goto tr816
		case 39:
			goto tr817
		case 43:
			goto tr378
		case 45:
			goto tr378
		case 61:
			goto st349
		case 95:
			goto tr381
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr815
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr381
				}
			case data[p] >= 65:
				goto tr381
			}
		default:
			goto tr379
		}
		goto st0
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		switch data[p] {
		case 32:
			goto tr819
		case 34:
			goto tr820
		case 39:
			goto tr821
		case 43:
			goto tr385
		case 45:
			goto tr385
		case 95:
			goto tr387
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr819
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr387
				}
			case data[p] >= 65:
				goto tr387
			}
		default:
			goto tr386
		}
		goto st0
	tr30:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st350
	tr35:
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st350
	tr838:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:155
		propose(ttEq)
		goto st350
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
//line query/tokeniser.go:16940
		if data[p] == 61 {
			goto st351
		}
		goto st0
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		switch data[p] {
		case 32:
			goto tr823
		case 34:
			goto tr824
		case 39:
			goto tr825
		case 43:
			goto tr392
		case 45:
			goto tr392
		case 95:
			goto tr394
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr823
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr394
				}
			case data[p] >= 65:
				goto tr394
			}
		default:
			goto tr393
		}
		goto st0
	tr31:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st352
	tr36:
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st352
	tr839:
//line query/tokeniser.rl:202
		setText(ttAttributeSelector)
//line query/tokeniser.rl:203
		commit(ttAttributeSelector)
//line query/tokeniser.rl:157
		propose(ttEq)
//line query/tokeniser.rl:159
		propose(ttEq)
		goto st352
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
//line query/tokeniser.go:17015
		switch data[p] {
		case 32:
			goto tr826
		case 34:
			goto tr827
		case 39:
			goto tr828
		case 43:
			goto tr398
		case 45:
			goto tr398
		case 61:
			goto st353
		case 95:
			goto tr401
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr826
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr401
				}
			case data[p] >= 65:
				goto tr401
			}
		default:
			goto tr399
		}
		goto st0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		switch data[p] {
		case 32:
			goto tr830
		case 34:
			goto tr831
		case 39:
			goto tr832
		case 43:
			goto tr405
		case 45:
			goto tr405
		case 95:
			goto tr407
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr830
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr407
				}
			case data[p] >= 65:
				goto tr407
			}
		default:
			goto tr406
		}
		goto st0
	tr27:
//line query/tokeniser.rl:201
		propose(ttAttributeSelector)
		goto st354
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
//line query/tokeniser.go:17096
		if data[p] == 95 {
			goto st355
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st355
			}
		case data[p] >= 65:
			goto st355
		}
		goto st0
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		switch data[p] {
		case 32:
			goto tr834
		case 33:
			goto tr835
		case 46:
			goto st354
		case 60:
			goto tr837
		case 61:
			goto tr838
		case 62:
			goto tr839
		case 95:
			goto st355
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr834
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st355
				}
			case data[p] >= 65:
				goto st355
			}
		default:
			goto st355
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
		goto st356
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
//line query/tokeniser.go:17163
		switch data[p] {
		case 32:
			goto tr840
		case 78:
			goto st362
		case 95:
			goto st361
		case 110:
			goto st362
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr840
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st361
				}
			case data[p] >= 65:
				goto st361
			}
		default:
			goto st361
		}
		goto st0
	tr840:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st357
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
//line query/tokeniser.go:17203
		switch data[p] {
		case 32:
			goto st357
		case 95:
			goto tr844
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st357
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr844
			}
		default:
			goto tr844
		}
		goto st0
	tr844:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st358
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
//line query/tokeniser.go:17234
		switch data[p] {
		case 32:
			goto tr845
		case 41:
			goto tr846
		case 44:
			goto tr847
		case 95:
			goto st358
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr845
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st358
				}
			case data[p] >= 65:
				goto st358
			}
		default:
			goto st358
		}
		goto st0
	tr845:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st359
	tr856:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st359
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
//line query/tokeniser.go:17280
		switch data[p] {
		case 32:
			goto st359
		case 41:
			goto st434
		case 44:
			goto st360
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st359
		}
		goto st0
	tr847:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st360
	tr858:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st360
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
//line query/tokeniser.go:17310
		switch data[p] {
		case 32:
			goto st360
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
				goto st360
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
		goto st361
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
//line query/tokeniser.go:17347
		switch data[p] {
		case 32:
			goto tr840
		case 95:
			goto st361
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr840
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st361
				}
			case data[p] >= 65:
				goto st361
			}
		default:
			goto st361
		}
		goto st0
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 32:
			goto tr840
		case 89:
			goto st363
		case 95:
			goto st361
		case 121:
			goto st363
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr840
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st361
				}
			case data[p] >= 65:
				goto st361
			}
		default:
			goto st361
		}
		goto st0
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		switch data[p] {
		case 32:
			goto tr852
		case 40:
			goto st365
		case 95:
			goto st361
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr852
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st361
				}
			case data[p] >= 65:
				goto st361
			}
		default:
			goto st361
		}
		goto st0
	tr852:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st364
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
//line query/tokeniser.go:17447
		switch data[p] {
		case 32:
			goto st357
		case 40:
			goto st365
		case 95:
			goto tr844
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st357
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr844
			}
		default:
			goto tr844
		}
		goto st0
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 32:
			goto st365
		case 41:
			goto st366
		case 95:
			goto tr855
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st365
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr855
			}
		default:
			goto tr855
		}
		goto st0
	tr864:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st366
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
//line query/tokeniser.go:17508
		switch data[p] {
		case 32:
			goto tr856
		case 41:
			goto tr857
		case 44:
			goto tr858
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr856
		}
		goto st0
	tr855:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st367
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
//line query/tokeniser.go:17534
		switch data[p] {
		case 32:
			goto tr859
		case 95:
			goto st367
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr859
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st367
				}
			case data[p] >= 65:
				goto st367
			}
		default:
			goto st367
		}
		goto st0
	tr859:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st368
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
//line query/tokeniser.go:17570
		switch data[p] {
		case 32:
			goto st368
		case 95:
			goto tr862
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st368
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr862
			}
		default:
			goto tr862
		}
		goto st0
	tr862:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st369
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
//line query/tokeniser.go:17601
		switch data[p] {
		case 32:
			goto tr863
		case 41:
			goto tr864
		case 44:
			goto tr865
		case 95:
			goto st369
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr863
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st369
				}
			case data[p] >= 65:
				goto st369
			}
		default:
			goto st369
		}
		goto st0
	tr863:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st370
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
//line query/tokeniser.go:17643
		switch data[p] {
		case 32:
			goto st370
		case 41:
			goto st366
		case 44:
			goto st371
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st370
		}
		goto st0
	tr865:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st371
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
//line query/tokeniser.go:17669
		switch data[p] {
		case 32:
			goto st371
		case 95:
			goto tr855
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st371
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr855
			}
		default:
			goto tr855
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
		goto st372
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
//line query/tokeniser.go:17706
		switch data[p] {
		case 32:
			goto tr869
		case 78:
			goto st375
		case 95:
			goto st374
		case 110:
			goto st375
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr869
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	tr869:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st373
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
//line query/tokeniser.go:17746
		switch data[p] {
		case 32:
			goto st373
		case 95:
			goto tr873
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st373
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr873
			}
		default:
			goto tr873
		}
		goto st0
	tr873:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st534
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
//line query/tokeniser.go:17777
		switch data[p] {
		case 32:
			goto tr1211
		case 59:
			goto tr1213
		case 95:
			goto st534
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1211
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st534
				}
			case data[p] >= 65:
				goto st534
			}
		default:
			goto st534
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
		goto st374
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
//line query/tokeniser.go:17819
		switch data[p] {
		case 32:
			goto tr869
		case 95:
			goto st374
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr869
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		switch data[p] {
		case 32:
			goto tr869
		case 89:
			goto st376
		case 95:
			goto st374
		case 121:
			goto st376
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr869
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		switch data[p] {
		case 32:
			goto tr875
		case 40:
			goto st378
		case 95:
			goto st374
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr875
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	tr875:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st377
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
//line query/tokeniser.go:17919
		switch data[p] {
		case 32:
			goto st373
		case 40:
			goto st378
		case 95:
			goto tr873
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st373
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr873
			}
		default:
			goto tr873
		}
		goto st0
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		switch data[p] {
		case 32:
			goto st378
		case 41:
			goto st535
		case 95:
			goto tr878
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st378
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr878
			}
		default:
			goto tr878
		}
		goto st0
	tr884:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st535
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
//line query/tokeniser.go:17980
		switch data[p] {
		case 32:
			goto tr1214
		case 59:
			goto tr1215
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1214
		}
		goto st0
	tr878:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st379
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
//line query/tokeniser.go:18004
		switch data[p] {
		case 32:
			goto tr879
		case 95:
			goto st379
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr879
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st379
				}
			case data[p] >= 65:
				goto st379
			}
		default:
			goto st379
		}
		goto st0
	tr879:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st380
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
//line query/tokeniser.go:18040
		switch data[p] {
		case 32:
			goto st380
		case 95:
			goto tr882
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st380
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr882
			}
		default:
			goto tr882
		}
		goto st0
	tr882:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st381
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
//line query/tokeniser.go:18071
		switch data[p] {
		case 32:
			goto tr883
		case 41:
			goto tr884
		case 44:
			goto tr885
		case 95:
			goto st381
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr883
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st381
				}
			case data[p] >= 65:
				goto st381
			}
		default:
			goto st381
		}
		goto st0
	tr883:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st382
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
//line query/tokeniser.go:18113
		switch data[p] {
		case 32:
			goto st382
		case 41:
			goto st535
		case 44:
			goto st383
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st382
		}
		goto st0
	tr885:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st383
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
//line query/tokeniser.go:18139
		switch data[p] {
		case 32:
			goto st383
		case 95:
			goto tr878
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st383
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr878
			}
		default:
			goto tr878
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
		goto st384
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
//line query/tokeniser.go:18176
		switch data[p] {
		case 32:
			goto tr869
		case 79:
			goto st385
		case 95:
			goto st374
		case 111:
			goto st385
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr869
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 32:
			goto tr869
		case 84:
			goto st386
		case 95:
			goto st374
		case 116:
			goto st386
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr869
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		switch data[p] {
		case 32:
			goto tr891
		case 40:
			goto st10
		case 95:
			goto st374
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
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	tr891:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st387
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
//line query/tokeniser.go:18280
		switch data[p] {
		case 32:
			goto st373
		case 40:
			goto st10
		case 95:
			goto tr873
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st373
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr873
			}
		default:
			goto tr873
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
		goto st388
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
//line query/tokeniser.go:18319
		switch data[p] {
		case 32:
			goto tr869
		case 69:
			goto st389
		case 95:
			goto st374
		case 101:
			goto st389
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr869
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 32:
			goto tr869
		case 81:
			goto st390
		case 95:
			goto st374
		case 113:
			goto st390
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr869
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		switch data[p] {
		case 32:
			goto tr894
		case 40:
			goto st392
		case 95:
			goto st374
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr894
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st374
				}
			case data[p] >= 65:
				goto st374
			}
		default:
			goto st374
		}
		goto st0
	tr894:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st391
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
//line query/tokeniser.go:18423
		switch data[p] {
		case 32:
			goto st373
		case 40:
			goto st392
		case 95:
			goto tr873
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st373
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr873
			}
		default:
			goto tr873
		}
		goto st0
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		switch data[p] {
		case 32:
			goto st393
		case 33:
			goto tr897
		case 41:
			goto st536
		case 65:
			goto tr899
		case 78:
			goto tr901
		case 95:
			goto tr900
		case 97:
			goto tr899
		case 110:
			goto tr901
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st393
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr900
			}
		default:
			goto tr900
		}
		goto st0
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		switch data[p] {
		case 32:
			goto st393
		case 41:
			goto st536
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st393
		}
		goto st0
	tr918:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st536
	tr927:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st536
	tr908:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
		goto st536
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
//line query/tokeniser.go:18517
		switch data[p] {
		case 32:
			goto tr1216
		case 59:
			goto tr1217
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1216
		}
		goto st0
	tr897:
//line query/tokeniser.rl:131
		propose(ttNegatedDecl)
		goto st394
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
//line query/tokeniser.go:18537
		switch data[p] {
		case 32:
			goto st395
		case 40:
			goto st396
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st395
		}
		goto st0
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		if data[p] == 40 {
			goto st396
		}
		goto st0
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		switch data[p] {
		case 32:
			goto st396
		case 41:
			goto st397
		case 65:
			goto tr905
		case 95:
			goto tr906
		case 97:
			goto tr905
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st396
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr906
			}
		default:
			goto tr906
		}
		goto st0
	tr948:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st397
	tr959:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st397
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
//line query/tokeniser.go:18604
		switch data[p] {
		case 32:
			goto tr907
		case 41:
			goto tr908
		case 44:
			goto tr909
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr907
		}
		goto st0
	tr917:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st398
	tr926:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st398
	tr907:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
		goto st398
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
//line query/tokeniser.go:18638
		switch data[p] {
		case 32:
			goto st398
		case 41:
			goto st536
		case 44:
			goto st399
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st398
		}
		goto st0
	tr919:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st399
	tr928:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st399
	tr909:
//line query/tokeniser.rl:135
		commit(ttNegatedDecl)
		goto st399
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
//line query/tokeniser.go:18672
		switch data[p] {
		case 32:
			goto st399
		case 33:
			goto tr897
		case 65:
			goto tr899
		case 78:
			goto tr901
		case 95:
			goto tr900
		case 97:
			goto tr899
		case 110:
			goto tr901
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st399
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr900
			}
		default:
			goto tr900
		}
		goto st0
	tr899:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
//line query/tokeniser.rl:119
		propose(ttAnyDecl)
		goto st400
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
//line query/tokeniser.go:18717
		switch data[p] {
		case 32:
			goto tr912
		case 78:
			goto st404
		case 95:
			goto st403
		case 110:
			goto st404
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
					goto st403
				}
			case data[p] >= 65:
				goto st403
			}
		default:
			goto st403
		}
		goto st0
	tr912:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st401
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
//line query/tokeniser.go:18757
		switch data[p] {
		case 32:
			goto st401
		case 95:
			goto tr916
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st401
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
		goto st402
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
//line query/tokeniser.go:18788
		switch data[p] {
		case 32:
			goto tr917
		case 41:
			goto tr918
		case 44:
			goto tr919
		case 95:
			goto st402
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
					goto st402
				}
			case data[p] >= 65:
				goto st402
			}
		default:
			goto st402
		}
		goto st0
	tr900:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st403
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
//line query/tokeniser.go:18830
		switch data[p] {
		case 32:
			goto tr912
		case 95:
			goto st403
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
					goto st403
				}
			case data[p] >= 65:
				goto st403
			}
		default:
			goto st403
		}
		goto st0
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		switch data[p] {
		case 32:
			goto tr912
		case 89:
			goto st405
		case 95:
			goto st403
		case 121:
			goto st405
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
					goto st403
				}
			case data[p] >= 65:
				goto st403
			}
		default:
			goto st403
		}
		goto st0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 32:
			goto tr922
		case 40:
			goto st407
		case 95:
			goto st403
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr922
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st403
				}
			case data[p] >= 65:
				goto st403
			}
		default:
			goto st403
		}
		goto st0
	tr922:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st406
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
//line query/tokeniser.go:18930
		switch data[p] {
		case 32:
			goto st401
		case 40:
			goto st407
		case 95:
			goto tr916
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st401
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr916
			}
		default:
			goto tr916
		}
		goto st0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 32:
			goto st407
		case 41:
			goto st408
		case 95:
			goto tr925
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st407
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr925
			}
		default:
			goto tr925
		}
		goto st0
	tr934:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st408
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
//line query/tokeniser.go:18991
		switch data[p] {
		case 32:
			goto tr926
		case 41:
			goto tr927
		case 44:
			goto tr928
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr926
		}
		goto st0
	tr925:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st409
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
//line query/tokeniser.go:19017
		switch data[p] {
		case 32:
			goto tr929
		case 95:
			goto st409
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr929
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st409
				}
			case data[p] >= 65:
				goto st409
			}
		default:
			goto st409
		}
		goto st0
	tr929:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st410
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
//line query/tokeniser.go:19053
		switch data[p] {
		case 32:
			goto st410
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st410
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	tr932:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st411
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
//line query/tokeniser.go:19084
		switch data[p] {
		case 32:
			goto tr933
		case 41:
			goto tr934
		case 44:
			goto tr935
		case 95:
			goto st411
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr933
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
	tr933:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st412
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
//line query/tokeniser.go:19126
		switch data[p] {
		case 32:
			goto st412
		case 41:
			goto st408
		case 44:
			goto st413
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st412
		}
		goto st0
	tr935:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st413
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
//line query/tokeniser.go:19152
		switch data[p] {
		case 32:
			goto st413
		case 95:
			goto tr925
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st413
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr925
			}
		default:
			goto tr925
		}
		goto st0
	tr901:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
//line query/tokeniser.rl:131
		propose(ttNegatedDecl)
		goto st414
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
//line query/tokeniser.go:19187
		switch data[p] {
		case 32:
			goto tr912
		case 79:
			goto st415
		case 95:
			goto st403
		case 111:
			goto st415
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
					goto st403
				}
			case data[p] >= 65:
				goto st403
			}
		default:
			goto st403
		}
		goto st0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		switch data[p] {
		case 32:
			goto tr912
		case 84:
			goto st416
		case 95:
			goto st403
		case 116:
			goto st416
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
					goto st403
				}
			case data[p] >= 65:
				goto st403
			}
		default:
			goto st403
		}
		goto st0
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 32:
			goto tr941
		case 40:
			goto st396
		case 95:
			goto st403
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
					goto st403
				}
			case data[p] >= 65:
				goto st403
			}
		default:
			goto st403
		}
		goto st0
	tr941:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st417
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
//line query/tokeniser.go:19291
		switch data[p] {
		case 32:
			goto st401
		case 40:
			goto st396
		case 95:
			goto tr916
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st401
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr916
			}
		default:
			goto tr916
		}
		goto st0
	tr905:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
//line query/tokeniser.rl:119
		propose(ttAnyDecl)
		goto st418
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
//line query/tokeniser.go:19328
		switch data[p] {
		case 32:
			goto tr942
		case 78:
			goto st424
		case 95:
			goto st423
		case 110:
			goto st424
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr942
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st423
				}
			case data[p] >= 65:
				goto st423
			}
		default:
			goto st423
		}
		goto st0
	tr942:
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
//line query/tokeniser.go:19368
		switch data[p] {
		case 32:
			goto st419
		case 95:
			goto tr946
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st419
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr946
			}
		default:
			goto tr946
		}
		goto st0
	tr946:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
//line query/tokeniser.go:19399
		switch data[p] {
		case 32:
			goto tr947
		case 41:
			goto tr948
		case 44:
			goto tr949
		case 95:
			goto st420
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
					goto st420
				}
			case data[p] >= 65:
				goto st420
			}
		default:
			goto st420
		}
		goto st0
	tr947:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st421
	tr958:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st421
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
//line query/tokeniser.go:19445
		switch data[p] {
		case 32:
			goto st421
		case 41:
			goto st397
		case 44:
			goto st422
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st421
		}
		goto st0
	tr949:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st422
	tr960:
//line query/tokeniser.rl:123
		commit(ttAnyDecl)
		goto st422
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
//line query/tokeniser.go:19475
		switch data[p] {
		case 32:
			goto st422
		case 65:
			goto tr905
		case 95:
			goto tr906
		case 97:
			goto tr905
		}
		switch {
		case data[p] < 66:
			if 9 <= data[p] && data[p] <= 13 {
				goto st422
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr906
			}
		default:
			goto tr906
		}
		goto st0
	tr906:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st423
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
//line query/tokeniser.go:19512
		switch data[p] {
		case 32:
			goto tr942
		case 95:
			goto st423
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr942
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st423
				}
			case data[p] >= 65:
				goto st423
			}
		default:
			goto st423
		}
		goto st0
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 32:
			goto tr942
		case 89:
			goto st425
		case 95:
			goto st423
		case 121:
			goto st425
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr942
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st423
				}
			case data[p] >= 65:
				goto st423
			}
		default:
			goto st423
		}
		goto st0
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		switch data[p] {
		case 32:
			goto tr954
		case 40:
			goto st427
		case 95:
			goto st423
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr954
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st423
				}
			case data[p] >= 65:
				goto st423
			}
		default:
			goto st423
		}
		goto st0
	tr954:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line query/tokeniser.go:19612
		switch data[p] {
		case 32:
			goto st419
		case 40:
			goto st427
		case 95:
			goto tr946
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st419
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr946
			}
		default:
			goto tr946
		}
		goto st0
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 32:
			goto st427
		case 41:
			goto st428
		case 95:
			goto tr957
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st427
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr957
			}
		default:
			goto tr957
		}
		goto st0
	tr966:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st428
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
//line query/tokeniser.go:19673
		switch data[p] {
		case 32:
			goto tr958
		case 41:
			goto tr959
		case 44:
			goto tr960
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr958
		}
		goto st0
	tr957:
//line query/tokeniser.rl:111
		propose(ttEventDecl)
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:100
		propose(ttEventDeclType)
		goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
//line query/tokeniser.go:19699
		switch data[p] {
		case 32:
			goto tr961
		case 95:
			goto st429
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr961
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st429
				}
			case data[p] >= 65:
				goto st429
			}
		default:
			goto st429
		}
		goto st0
	tr961:
//line query/tokeniser.rl:101
		setText(ttEventDeclType)
//line query/tokeniser.rl:102
		commit(ttEventDeclType)
		goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
//line query/tokeniser.go:19735
		switch data[p] {
		case 32:
			goto st430
		case 95:
			goto tr964
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st430
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr964
			}
		default:
			goto tr964
		}
		goto st0
	tr964:
//line query/tokeniser.rl:87
		mark = p
//line query/tokeniser.rl:106
		propose(ttEventDeclAlias)
		goto st431
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
//line query/tokeniser.go:19766
		switch data[p] {
		case 32:
			goto tr965
		case 41:
			goto tr966
		case 44:
			goto tr967
		case 95:
			goto st431
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr965
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st431
				}
			case data[p] >= 65:
				goto st431
			}
		default:
			goto st431
		}
		goto st0
	tr965:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
//line query/tokeniser.go:19808
		switch data[p] {
		case 32:
			goto st432
		case 41:
			goto st428
		case 44:
			goto st433
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st432
		}
		goto st0
	tr967:
//line query/tokeniser.rl:107
		setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
		commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
		commit(ttEventDecl)
		goto st433
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
//line query/tokeniser.go:19834
		switch data[p] {
		case 32:
			goto st433
		case 95:
			goto tr957
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st433
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr957
			}
		default:
			goto tr957
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
	_test_eof434:
		cs = 434
		goto _test_eof
	_test_eof435:
		cs = 435
		goto _test_eof
	_test_eof436:
		cs = 436
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
	_test_eof437:
		cs = 437
		goto _test_eof
	_test_eof438:
		cs = 438
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
	_test_eof439:
		cs = 439
		goto _test_eof
	_test_eof440:
		cs = 440
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
	_test_eof441:
		cs = 441
		goto _test_eof
	_test_eof442:
		cs = 442
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof443:
		cs = 443
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof444:
		cs = 444
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof445:
		cs = 445
		goto _test_eof
	_test_eof446:
		cs = 446
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
	_test_eof447:
		cs = 447
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof448:
		cs = 448
		goto _test_eof
	_test_eof449:
		cs = 449
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
	_test_eof450:
		cs = 450
		goto _test_eof
	_test_eof451:
		cs = 451
		goto _test_eof
	_test_eof452:
		cs = 452
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof453:
		cs = 453
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof454:
		cs = 454
		goto _test_eof
	_test_eof455:
		cs = 455
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof456:
		cs = 456
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
	_test_eof457:
		cs = 457
		goto _test_eof
	_test_eof458:
		cs = 458
		goto _test_eof
	_test_eof459:
		cs = 459
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
	_test_eof460:
		cs = 460
		goto _test_eof
	_test_eof461:
		cs = 461
		goto _test_eof
	_test_eof462:
		cs = 462
		goto _test_eof
	_test_eof95:
		cs = 95
		goto _test_eof
	_test_eof463:
		cs = 463
		goto _test_eof
	_test_eof96:
		cs = 96
		goto _test_eof
	_test_eof464:
		cs = 464
		goto _test_eof
	_test_eof465:
		cs = 465
		goto _test_eof
	_test_eof97:
		cs = 97
		goto _test_eof
	_test_eof466:
		cs = 466
		goto _test_eof
	_test_eof98:
		cs = 98
		goto _test_eof
	_test_eof99:
		cs = 99
		goto _test_eof
	_test_eof100:
		cs = 100
		goto _test_eof
	_test_eof101:
		cs = 101
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
	_test_eof467:
		cs = 467
		goto _test_eof
	_test_eof468:
		cs = 468
		goto _test_eof
	_test_eof469:
		cs = 469
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
	_test_eof470:
		cs = 470
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof471:
		cs = 471
		goto _test_eof
	_test_eof472:
		cs = 472
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof473:
		cs = 473
		goto _test_eof
	_test_eof128:
		cs = 128
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
	_test_eof134:
		cs = 134
		goto _test_eof
	_test_eof135:
		cs = 135
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
	_test_eof148:
		cs = 148
		goto _test_eof
	_test_eof149:
		cs = 149
		goto _test_eof
	_test_eof150:
		cs = 150
		goto _test_eof
	_test_eof474:
		cs = 474
		goto _test_eof
	_test_eof475:
		cs = 475
		goto _test_eof
	_test_eof476:
		cs = 476
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
	_test_eof477:
		cs = 477
		goto _test_eof
	_test_eof157:
		cs = 157
		goto _test_eof
	_test_eof478:
		cs = 478
		goto _test_eof
	_test_eof479:
		cs = 479
		goto _test_eof
	_test_eof158:
		cs = 158
		goto _test_eof
	_test_eof480:
		cs = 480
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
	_test_eof169:
		cs = 169
		goto _test_eof
	_test_eof170:
		cs = 170
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
	_test_eof481:
		cs = 481
		goto _test_eof
	_test_eof482:
		cs = 482
		goto _test_eof
	_test_eof483:
		cs = 483
		goto _test_eof
	_test_eof182:
		cs = 182
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
	_test_eof484:
		cs = 484
		goto _test_eof
	_test_eof485:
		cs = 485
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
	_test_eof486:
		cs = 486
		goto _test_eof
	_test_eof195:
		cs = 195
		goto _test_eof
	_test_eof196:
		cs = 196
		goto _test_eof
	_test_eof487:
		cs = 487
		goto _test_eof
	_test_eof488:
		cs = 488
		goto _test_eof
	_test_eof197:
		cs = 197
		goto _test_eof
	_test_eof489:
		cs = 489
		goto _test_eof
	_test_eof490:
		cs = 490
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
	_test_eof203:
		cs = 203
		goto _test_eof
	_test_eof204:
		cs = 204
		goto _test_eof
	_test_eof205:
		cs = 205
		goto _test_eof
	_test_eof491:
		cs = 491
		goto _test_eof
	_test_eof492:
		cs = 492
		goto _test_eof
	_test_eof493:
		cs = 493
		goto _test_eof
	_test_eof206:
		cs = 206
		goto _test_eof
	_test_eof494:
		cs = 494
		goto _test_eof
	_test_eof207:
		cs = 207
		goto _test_eof
	_test_eof495:
		cs = 495
		goto _test_eof
	_test_eof496:
		cs = 496
		goto _test_eof
	_test_eof208:
		cs = 208
		goto _test_eof
	_test_eof497:
		cs = 497
		goto _test_eof
	_test_eof209:
		cs = 209
		goto _test_eof
	_test_eof210:
		cs = 210
		goto _test_eof
	_test_eof211:
		cs = 211
		goto _test_eof
	_test_eof212:
		cs = 212
		goto _test_eof
	_test_eof213:
		cs = 213
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
	_test_eof222:
		cs = 222
		goto _test_eof
	_test_eof223:
		cs = 223
		goto _test_eof
	_test_eof224:
		cs = 224
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
	_test_eof498:
		cs = 498
		goto _test_eof
	_test_eof499:
		cs = 499
		goto _test_eof
	_test_eof500:
		cs = 500
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
	_test_eof501:
		cs = 501
		goto _test_eof
	_test_eof236:
		cs = 236
		goto _test_eof
	_test_eof502:
		cs = 502
		goto _test_eof
	_test_eof237:
		cs = 237
		goto _test_eof
	_test_eof503:
		cs = 503
		goto _test_eof
	_test_eof504:
		cs = 504
		goto _test_eof
	_test_eof238:
		cs = 238
		goto _test_eof
	_test_eof505:
		cs = 505
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
	_test_eof256:
		cs = 256
		goto _test_eof
	_test_eof257:
		cs = 257
		goto _test_eof
	_test_eof258:
		cs = 258
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
	_test_eof506:
		cs = 506
		goto _test_eof
	_test_eof507:
		cs = 507
		goto _test_eof
	_test_eof508:
		cs = 508
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
	_test_eof509:
		cs = 509
		goto _test_eof
	_test_eof510:
		cs = 510
		goto _test_eof
	_test_eof269:
		cs = 269
		goto _test_eof
	_test_eof270:
		cs = 270
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
	_test_eof511:
		cs = 511
		goto _test_eof
	_test_eof512:
		cs = 512
		goto _test_eof
	_test_eof279:
		cs = 279
		goto _test_eof
	_test_eof513:
		cs = 513
		goto _test_eof
	_test_eof514:
		cs = 514
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
	_test_eof286:
		cs = 286
		goto _test_eof
	_test_eof287:
		cs = 287
		goto _test_eof
	_test_eof515:
		cs = 515
		goto _test_eof
	_test_eof516:
		cs = 516
		goto _test_eof
	_test_eof517:
		cs = 517
		goto _test_eof
	_test_eof288:
		cs = 288
		goto _test_eof
	_test_eof518:
		cs = 518
		goto _test_eof
	_test_eof289:
		cs = 289
		goto _test_eof
	_test_eof519:
		cs = 519
		goto _test_eof
	_test_eof520:
		cs = 520
		goto _test_eof
	_test_eof290:
		cs = 290
		goto _test_eof
	_test_eof521:
		cs = 521
		goto _test_eof
	_test_eof291:
		cs = 291
		goto _test_eof
	_test_eof292:
		cs = 292
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
	_test_eof303:
		cs = 303
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
	_test_eof312:
		cs = 312
		goto _test_eof
	_test_eof313:
		cs = 313
		goto _test_eof
	_test_eof522:
		cs = 522
		goto _test_eof
	_test_eof523:
		cs = 523
		goto _test_eof
	_test_eof524:
		cs = 524
		goto _test_eof
	_test_eof314:
		cs = 314
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
	_test_eof525:
		cs = 525
		goto _test_eof
	_test_eof526:
		cs = 526
		goto _test_eof
	_test_eof318:
		cs = 318
		goto _test_eof
	_test_eof527:
		cs = 527
		goto _test_eof
	_test_eof319:
		cs = 319
		goto _test_eof
	_test_eof528:
		cs = 528
		goto _test_eof
	_test_eof529:
		cs = 529
		goto _test_eof
	_test_eof320:
		cs = 320
		goto _test_eof
	_test_eof530:
		cs = 530
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
	_test_eof342:
		cs = 342
		goto _test_eof
	_test_eof343:
		cs = 343
		goto _test_eof
	_test_eof531:
		cs = 531
		goto _test_eof
	_test_eof532:
		cs = 532
		goto _test_eof
	_test_eof533:
		cs = 533
		goto _test_eof
	_test_eof344:
		cs = 344
		goto _test_eof
	_test_eof345:
		cs = 345
		goto _test_eof
	_test_eof346:
		cs = 346
		goto _test_eof
	_test_eof347:
		cs = 347
		goto _test_eof
	_test_eof348:
		cs = 348
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
	_test_eof534:
		cs = 534
		goto _test_eof
	_test_eof374:
		cs = 374
		goto _test_eof
	_test_eof375:
		cs = 375
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
	_test_eof535:
		cs = 535
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
	_test_eof536:
		cs = 536
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

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 535:
//line query/tokeniser.rl:123
				commit(ttAnyDecl)
//line query/tokeniser.rl:152
				commit(ttEventClause)
			case 434:
//line query/tokeniser.rl:135
				commit(ttNegatedDecl)
//line query/tokeniser.rl:152
				commit(ttEventClause)
			case 536:
//line query/tokeniser.rl:146
				commit(ttSeqDecl)
//line query/tokeniser.rl:152
				commit(ttEventClause)
			case 444, 445, 447, 452, 460, 461, 487, 492, 501, 509, 517, 525, 526:
//line query/tokeniser.rl:186
				commit(ttStringLiteral)
//line query/tokeniser.rl:209
				commit(ttPredicate)
			case 437, 439, 441, 443, 448, 450, 484, 486, 489, 491, 511, 513, 515:
//line query/tokeniser.rl:194
				commit(ttStringLiteral)
//line query/tokeniser.rl:209
				commit(ttPredicate)
			case 453, 454, 463, 464, 470, 471, 477, 478, 494, 495, 502, 503, 518, 519, 527, 528:
//line query/tokeniser.rl:177
				setText(ttNumericLiteral)
//line query/tokeniser.rl:178
				commit(ttNumericLiteral)
//line query/tokeniser.rl:209
				commit(ttPredicate)
			case 456, 466, 473, 480, 497, 505, 521, 530:
//line query/tokeniser.rl:202
				setText(ttAttributeSelector)
//line query/tokeniser.rl:203
				commit(ttAttributeSelector)
//line query/tokeniser.rl:209
				commit(ttPredicate)
			case 457, 459, 467, 469, 474, 476, 481, 483, 498, 500, 506, 508, 522, 524, 531, 533:
//line query/tokeniser.rl:219
				setText(ttDuration)
//line query/tokeniser.rl:220
				commit(ttDuration)
//line query/tokeniser.rl:224
				commit(ttWithinClause)
			case 534:
//line query/tokeniser.rl:107
				setText(ttEventDeclAlias)
//line query/tokeniser.rl:108
				commit(ttEventDeclAlias)
//line query/tokeniser.rl:112
				commit(ttEventDecl)
//line query/tokeniser.rl:152
				commit(ttEventClause)
			case 455, 465, 472, 479, 496, 504, 520, 529:
//line query/tokeniser.rl:201
				propose(ttAttributeSelector)
//line query/tokeniser.rl:202
				setText(ttAttributeSelector)
//line query/tokeniser.rl:203
				commit(ttAttributeSelector)
//line query/tokeniser.rl:209
				commit(ttPredicate)
//line query/tokeniser.go:20459
			}
		}

	_out:
		{
		}
	}

//line query/tokeniser.rl:236

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
