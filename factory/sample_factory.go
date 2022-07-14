package factory

type Factory interface {
	Create()
	Update()
	Delete()
	Select()
}

type FactoryA struct {}
func (a *FactoryA)Create()  {}
func (a *FactoryA)Update()  {}
func (a *FactoryA)Delete()  {}
func (a *FactoryA)Select()  {}

type FactoryB struct {}
func (b *FactoryB)Create()  {}
func (b *FactoryB)Update()  {}
func (b *FactoryB)Delete()  {}
func (b *FactoryB)Select()  {}

type SimpleFactory struct {}
//sample factory
func (s *SimpleFactory)Create(createType string)  Factory{
	switch createType {
	case "A":
		return &FactoryA{}
	case "B":
		return &FactoryA{}
	}
	return nil
}