package channel

type Channel struct {
	name string
	quantity int
	cost float64	
}

func NewChannel(name string, quantity int, cost float64) (*Channel, error) {
	return &Channel{
		name: name,
		quantity: quantity,
		cost: cost,
	}, nil
}
