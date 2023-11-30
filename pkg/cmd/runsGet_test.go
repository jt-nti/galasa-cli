/*
 * Copyright contributors to the Galasa project
 *
 * SPDX-License-Identifier: EPL-2.0
 */
package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunsGetCommandInCommandCollectionIsAsExpected(t *testing.T) {

	factory := NewMockFactory()
	commands, _ := NewCommandCollection(factory)

	runsGetCommand := commands.GetCommand(COMMAND_NAME_RUNS_GET)
	assert.Equal(t, COMMAND_NAME_RUNS_GET, runsGetCommand.GetName())
	assert.NotNil(t, runsGetCommand.GetValues())
	assert.IsType(t, &RunsGetCmdValues{}, runsGetCommand.GetValues())
	assert.NotNil(t, runsGetCommand.GetCobraCommand())
}
