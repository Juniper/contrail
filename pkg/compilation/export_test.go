package compilation

func (ics *IntentCompilationService) Store() Store {
	return ics.store
}

func (ics *IntentCompilationService) SetStore(newStore Store) {
	ics.store = newStore
}
