package main


type Phone struct {
	cpu       string
	screen    string
	memory    string
	mainboard string
}

type Builder struct {
	cpu       string
	screen    string
	memory    string
	mainboard string
}

func (p *Phone)phone(builder Builder) *Phone{
	p.cpu = builder.cpu
	p.screen = builder.screen
	p.memory = builder.memory
	p.mainboard = builder.mainboard
	return p
}

func (p Phone) toString() string{
	return "Phone [cpu=" + p.cpu + ", mainboard=" + p.mainboard + ", memory=" + p.memory + ", screen=" + p.screen + "]";
}

func (b *Builder) SetCpu(cpu string) *Builder{
	b.cpu = cpu
	return b
}

func (b *Builder) SetScreen(screen string) *Builder{
	b.screen = screen
	return b
}
func (b *Builder) SetMemory(memory string) *Builder{
	b.memory = memory
	return b
}
func (b *Builder) SetMainboard(mainboard string) *Builder{
	b.mainboard = mainboard
	return b
}
func (b Builder)build() *Phone{
	return new(Phone).phone(b);
}