/*
 * Copyright contributors to the Galasa project
 *
 * SPDX-License-Identifier: EPL-2.0
 */
package tokensformatter

import (
	"strconv"
	"strings"

	galasaapi "github.com/jt-nti/galasa-api-go"
)

// -----------------------------------------------------
// Summary format.
const (
	SUMMARY_FORMATTER_NAME = "summary"
)

type TokenSummaryFormatter struct {
}

func NewTokenSummaryFormatter() TokenFormatter {
	return new(TokenSummaryFormatter)
}

func (*TokenSummaryFormatter) GetName() string {
	return SUMMARY_FORMATTER_NAME
}

func (*TokenSummaryFormatter) FormatTokens(authTokens []galasaapi.AuthToken) (string, error) {
	var result string = ""
	var err error = nil
	buff := strings.Builder{}
	totalTokens := len(authTokens)

	if totalTokens > 0 {
		var table [][]string

		var headers = []string{HEADER_TOKEN_ID, HEADER_TOKEN_CREATION_TIME, HEADER_TOKEN_USER, HEADER_TOKEN_DESCRIPTION}

		table = append(table, headers)
		for _, token := range authTokens {
			var line []string
			id := token.GetTokenId()
			creationTime := formatTimeToNearestDate(token.GetCreationTime())
			owner := token.GetOwner()
			ownerLoginId := owner.GetLoginId()
			description := token.GetDescription()

			line = append(line, id, creationTime, ownerLoginId, description)
			table = append(table, line)
		}

		columnLengths := calculateMaxLengthOfEachColumn(table)
		writeFormattedTableToStringBuilder(table, &buff, columnLengths)

		buff.WriteString("\n")

	}
	buff.WriteString("Total:" + strconv.Itoa(totalTokens) + "\n")

	result = buff.String()
	return result, err
}
