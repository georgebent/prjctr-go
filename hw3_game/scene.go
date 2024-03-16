package hw3_game

type Scene struct {
	name        string
	description string
	question    string
	options     []*Scene
	getReward   func(player *Player)
	isAllowed   func(player *Player) bool
}

func (s Scene) withOption(scene *Scene) Scene {
	s.options = append(s.options, scene)

	return s
}

func (s Scene) withName(name string) Scene {
	s.name = name

	return s
}

func (s Scene) withDescription(description string) Scene {
	s.description = description

	return s
}

func (s Scene) getChosenScene(chosen int, player *Player) Scene {
	options := getAvailableOptions(s.options, player)

	return *options[chosen]
}

func (s Scene) run(player *Player) int {
	options := getAvailableOptions(s.options, player)
	answer := ask(s.description, s.question, options)

	if s.getReward != nil {
		s.getReward(player)
	}

	return answer
}

func getAvailableOptions(sceneOptions []*Scene, player *Player) []*Scene {
	var options []*Scene
	for _, option := range sceneOptions {
		if option.isAllowed == nil || option.isAllowed(player) {
			options = append(options, option)
		}
	}

	return options
}
