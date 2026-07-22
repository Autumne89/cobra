func TestRequiredFlagsBeforePreRun(t *testing.T) {
	preRunExecuted := false

	cmd := &Command{
		Use: "testcmd",
		PersistentPreRunE: func(cmd *Command, args []string) error {
			preRunExecuted = true
			return nil
		},
		Run: func(cmd *Command, args []string) {},
	}

	cmd.Flags().String("required-flag", "", "A required flag")
	cmd.MarkFlagRequired("required-flag")

	// Execute without the required flag
	cmd.SetArgs([]string{})
	err := cmd.Execute()

	if err == nil {
		t.Fatal("Expected error due to missing required flag, got nil")
	}

	if preRunExecuted {
		t.Error("Expected PersistentPreRunE NOT to execute when required flags are missing")
	}
}