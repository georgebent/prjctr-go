package hw3_game

func initialize() (Scene, Player) {
	player := initPlayer()

	start := Scene{
		name: "Вітаємо у грі",
		description: "Ви прокидаєтесь посеред лісу, відчуваєте біль у голові, не можете пригадати останні події." +
			" У вас є залізний меч і ліхтар, здається вас звати Ґеральт",
		question: "На землі, видно сліди якогось монстра, стежку, яка ймовірно веде у поселення і" +
			" недалеко також видно печеру, ви вирішуєте піти у:",
	}

	start.getReward = func(player *Player) {
		ironSwordWeapon := createWeapon(ironSword, 1)
		player.inventory.addWeapon(&ironSwordWeapon)
	}

	forest := start.withName("Повернутися у ліс").withDescription("Ви знову у лісі")
	forest.getReward = nil

	village := Scene{
		name:        "Перейти у поселення",
		description: "Ви входите в поселення",
		question:    "Ваші дії",
	}

	camp := initCamp(&village)
	witcherHouse := initWitcherHouse(&village)
	doctor := initDoctor(&village)

	village = village.withOption(&witcherHouse).withOption(&camp).withOption(&forest).withOption(&doctor)

	cave := initCaveQuest(&forest)
	steps := initWolfBoss(&forest)

	start = start.withOption(&cave).withOption(&steps).withOption(&village)
	forest = forest.withOption(&cave).withOption(&steps).withOption(&village)

	return start, player
}

func initPlayer() Player {
	inventory := Inventory{
		bug:    nil,
		skills: nil,
	}

	player := Player{
		inventory: &inventory,
	}

	return player
}

func initDoctor(village *Scene) Scene {
	doctor := Scene{
		name: "Перейти у дім знахаря",
		description: "Ви заходите у дім знахаря, він бачить що ви поранені," +
			" дає вам цілющого зілля і пропонує відпочити, випивши зілля, ви вирішуєте скористатись гостинністю " +
			"і поспати у теплій хатині. Вночі вам сниться сон, ви згадуєте що ви відьмак," +
			" людина з надприродніми здібностями, яка полює на монстрів і інших істот які приносять людям біди, " +
			"ви пригадуєте магічні знаки і можете тепер ними користуватись, " +
			"Аард для телекінетичного удару, Ігні для удару вогнем." +
			" Зранку знахар розказує вам що ви взяли завдання вбити вовкулаку і жили в сусідньому будинку" +
			" вивчаючи справу, можливо вам пора навідатись у цю стару закинуту хатину.",
		question: "Подальші дії",
		isAllowed: func(player *Player) bool {
			return !player.hasProgress(doctorVisited)
		},
		getReward: func(player *Player) {
			player.addProgress(doctorVisited)
			signAardSkill := createSign(signAard, false)
			player.inventory.addSkill(&signAardSkill)

			signIgniSkill := createSign(signIgni, false)
			player.inventory.addSkill(&signIgniSkill)
		},
	}

	return doctor.withOption(village)
}

func initCamp(village *Scene) Scene {
	camp := Scene{
		name:        "Перейти у табір розбійників",
		description: "Ви входите у табір",
		question:    "Подальші дії",
	}

	campAttack := Scene{
		name:        "Напасти",
		description: "Ви нападаєте на розбійників, маючи надлюдські здібності ви з легкістю перемагаєте розбійників і отримуєте в винагороду арбалет",
		question:    "Подальші дії",
		isAllowed: func(player *Player) bool {
			return !player.hasProgress(isCampKilled)
		},
		getReward: func(player *Player) {
			player.addProgress(isCampKilled)

			arbalestWeapon := createWeapon(arbalest, 1)
			player.inventory.addWeapon(&arbalestWeapon)
		},
	}

	campAttack = campAttack.withOption(village)

	return camp.withOption(&campAttack).withOption(village)
}

func initWitcherHouse(village *Scene) Scene {
	witcherHouse := Scene{
		name:        "Перейти у зактнутий дім",
		description: "Ви підходите до дивного закинутого будинку, бачите велику купу предметів при вході",
		question:    "Напевно без магії тут не обійтись, ви пробуєте:",
		isAllowed: func(player *Player) bool {
			return !player.hasProgress(witcherHouseVisited)
		},
	}

	noRememberSighs := Scene{
		name:        "Спробувати розібрати завали",
		description: "Ви пробуєте розібрати завали, але це не дає результату, потрібно пригадати магічні знаки",
		question:    "Ваші дії:",
	}

	for _, sign := range getAllGameSigns() {
		if sign == signAard {
			houseSuccess := Scene{
				name: sign,
				description: "Ви входите в дім і знаходите свої записи, виявляється ви з'ясували що вовкулаком" +
					" став могутній чарівник, тому він виявився могутнішим за інших, ви " +
					"готували настоянки а також занотували що вбити його можна срібним мечем," +
					" ви пригадуєте відьмачий знак квен, який використовується для захисту, " +
					"він неодмінно вам пригодиться",
				question: "Ваші дії",
				isAllowed: func(player *Player) bool {
					return player.inventory.hasSkill(sign)
				},
				getReward: func(player *Player) {
					player.addProgress(witcherHouseVisited)
					signKvenSkill := createSign(signKven, true)
					player.inventory.addSkill(&signKvenSkill)
				},
			}

			houseSuccess = houseSuccess.withOption(village)
			witcherHouse = witcherHouse.withOption(&houseSuccess)

			continue
		}

		houseWrong := Scene{
			name:        sign,
			description: "Ви використали хибний знак, будинок закритий",
			question:    "Ваші дії",
			isAllowed: func(player *Player) bool {
				return player.inventory.hasSkill(sign)
			},
		}

		houseWrong = houseWrong.withOption(village)
		witcherHouse = witcherHouse.withOption(&houseWrong)
	}

	noRememberSighs = noRememberSighs.withOption(village)

	return witcherHouse.withOption(&noRememberSighs)
}

func initCaveQuest(forest *Scene) Scene {
	cave := Scene{
		name:        "Зайти у печеру",
		description: "Вхід в печеру закритий дивним замком, помітно що це магія холоду, ви можете використати магічне заклинання, ваші дії:",
		question:    "Ваші дії",
		isAllowed: func(player *Player) bool {
			return !player.hasProgress(caveVisited)
		},
	}

	noRememberSighs := Scene{
		name:        "Спробувати виламати замок",
		description: "Ви пробуєте виломати замок, але це не дає результату, потрібно пригадати магічні знаки",
		question:    "Ваші дії:",
	}

	for _, sign := range getAllGameSigns() {
		if sign == signIgni {
			caveSuccess := Scene{
				name: sign,
				description: "Ви входите в печеру і знаходите свій срібний меч який втратили, " +
					"також ви згадуєте що вистежували чарівника який перетворюється на вовкулаку, і між вами був бій " +
					"він використав проти вас смертоносне закляття, але завдяки знаку Квен ви змогли себе захистити" +
					" хоча отримали сильний удар і тимчасово втратили память. Чарівник подумав що вбив вас і відібрав срібний меч, " +
					"тепер потрібно знайти кривдника, зараз якраз північ, отже цього разу вас чекає бій з вовкулаком",
				question: "Ваші дії",
				getReward: func(player *Player) {
					player.addProgress(caveVisited)
					silverSwordWeapon := createWeapon(silverSword, 2)
					player.inventory.addWeapon(&silverSwordWeapon)
				},
				isAllowed: func(player *Player) bool {
					return player.inventory.hasSkill(sign)
				},
			}

			caveSuccess = caveSuccess.withOption(forest)

			cave = cave.withOption(&caveSuccess)

			continue
		}

		caveWrong := Scene{
			name:        sign,
			description: "Ви використали хибний знак, печера закрита",
			question:    "Ваші дії",
			isAllowed: func(player *Player) bool {
				return player.inventory.hasSkill(sign)
			},
		}

		caveWrong = caveWrong.withOption(forest)
		cave = cave.withOption(&caveWrong)
	}

	noRememberSighs = noRememberSighs.withOption(forest)

	return cave.withOption(&noRememberSighs)
}

func initWolfBoss(forest *Scene) Scene {
	steps := Scene{
		name:        "Піти по дивних слідах",
		description: "Ви бачите вовкулака",
		question:    "Ваші дії:",
	}

	wolf := Scene{
		name:        "Напасти",
		description: "Вовкулак почув вас і напав першим",
		question:    "Для захисту вам потрібно використати магічний знак, ви пробуєте:",
	}

	var finalOptions []*Scene
	for _, weapon := range getAllGameWeapon() {
		switch weapon {
		case silverSword:
			finalOptions = append(finalOptions, &Scene{
				name:        weapon,
				description: "Ви наносите удари сріблом, вовкулака падає від болю, ви перемогли, поселення врятоване, кінець гри!",
				question:    "ЯКінець гри",
				isAllowed: func(player *Player) bool {
					return player.inventory.hasWeapon(weapon)
				},
			})
		case arbalest:
			finalOptions = append(finalOptions, &Scene{
				name:        weapon,
				description: "Ви вистрілили з арбелату, але вовкулака ухилився від удару і розриває вас на шматки, кінець гри!",
				question:    "Кінець гри",
				isAllowed: func(player *Player) bool {
					return player.inventory.hasWeapon(weapon)
				},
			})
		case ironSword:
			finalOptions = append(finalOptions, &Scene{
				name:        weapon,
				description: "Ви наносите удари залізним мечем, але рани одразу заживають, ви помилились, залізо не діє на вовкулаку і він розриває вас на шматки, кінець гри!",
				question:    "Кінець гри",
				isAllowed: func(player *Player) bool {
					return player.inventory.hasWeapon(weapon)
				},
			})
		}
	}

	for _, sign := range getAllGameSigns() {
		if sign == signKven {
			wolfSuccessBlock := Scene{
				name:        sign,
				description: "Ви вдало блокували напад, тепер час вам нанести удар",
				question:    "Яку зброю ви оберете ?",
				isAllowed: func(player *Player) bool {
					return player.inventory.hasSkill(sign)
				},
				options: finalOptions,
			}

			wolf = wolf.withOption(&wolfSuccessBlock)

			continue
		}

		wolfBlockWrong := Scene{
			name:        sign,
			description: "Ви не використали знак захисту, вовкулака уник вашого знаку, і розірвав вас на шматки, кінць гри",
			question:    "Кінець гри",
			isAllowed: func(player *Player) bool {
				return player.inventory.hasSkill(sign)
			},
		}

		wolf = wolf.withOption(&wolfBlockWrong)
	}

	wolf = wolf.withOption(&Scene{
		name:        "Відхилитися від нападу",
		description: "Ви не використали знак захисту, вовкулака розірвав вас на шматки, кінць гри",
		question:    "Кінець гри",
	})

	steps = steps.withOption(forest).withOption(&wolf)

	return steps
}
