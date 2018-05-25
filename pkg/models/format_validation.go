package models

func TypeValidatorWithFormat() TypeValidator {
	tv := TypeValidator{}
	tv.SchemaValidator = SchemaValidator{}

	tv.SchemaValidator.validators = map[string]func(string) error{}

	// Register all format validators

	tv.SchemaValidator.addFormatValidator("hostname", func(hostname string) error {
		// Validate hostname
		return nil
	})

	tv.SchemaValidator.addFormatValidator("date-time", func(hostname string) error {
		// Validate date-time
		return nil
	})

	return tv

}
