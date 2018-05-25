package models

func TypeValidatorWithFormat() TypeValidator {
	tv := TypeValidator{}

	// Register all format validators

	tv.SchemaValidator.addFormatValidator("hostname", func(hostname string) error {
		// Validate hostname
		return nil
	})

	return tv

}
