package catalogue

func CreateUseCase() InputPort {
	return &UseCase{
		CategoryRepo: &CategoryPsqlAdapter{},
	}
}
