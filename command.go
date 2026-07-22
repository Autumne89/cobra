func (c *Command) execute(a []string) error {
	if c.preRun != nil {
		c.preRun(c, a)
	}

	for _, p := range c.parents() {
		if p.persistentPreRun != nil {
			p.persistentPreRun(c, a)
		}
	}

	err := c.ParseFlags(a)
	if err != nil {
		return err
	}

	// Validate required flags before running pre-run hooks
	if err := c.validateRequiredFlags(); err != nil {
		return err
	}

	if c.PersistentPreRunE != nil {
		if err := c.PersistentPreRunE(c, a); err != nil {
			return err
		}
	} else if c.PersistentPreRun != nil {
		c.PersistentPreRun(c, a)
	}

	for _, p := range c.parents() {
		if p.PersistentPreRunE != nil {
			if err := p.PersistentPreRunE(c, a); err != nil {
				return err
			}
		} else if p.PersistentPreRun != nil {
			p.PersistentPreRun(c, a)
		}
	}

	if c.PreRunE != nil {
		if err := c.PreRunE(c, a); err != nil {
			return err
		}
	} else if c.PreRun != nil {
		c.PreRun(c, a)
	}

	if c.RunE != nil {
		if err := c.RunE(c, a); err != nil {
			return err
		}
	} else {
		c.Run(c, a)
	}
	return nil
}