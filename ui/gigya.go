package ui

import (
	"gigya-module-go/accounts"

	"github.com/evertras/bubble-table/table"
)

/* ╭──────────────────────────────────────────╮ */
/* │                 LIVGOLF                  │ */
/* ╰──────────────────────────────────────────╯ */
const (
	columnKeyUID                  = "uid"
	columnKeyEmail                = "email"
	columnKeyFirstName            = "firstName"
	columnKeyLastName             = "lastName"
	columnKeyHasLiteAccount       = "hasLiteAccount"
	columnKeyHasFullAccount       = "hasFullAccount"
	columnKeyLastUpdatedTimestamp = "lastUpdatedTimestamp"
	columnKeyEventsName           = "eventsName"
	columnKeyEventsWhen           = "eventsWhen"
	columnKeyCreatedTimestamp     = "createdTimestamp"
)

func makeRow(account accounts.Account) table.Row {
	return table.NewRow(table.RowData{
		columnKeyUID:            account.UID,
		columnKeyEmail:          account.Email,
		columnKeyFirstName:      account.Profile.FirstName,
		columnKeyLastName:       account.Profile.LastName,
		columnKeyHasLiteAccount: account.HasLiteAccount,
		columnKeyHasFullAccount: account.HasFullAccount,

		columnKeyLastUpdatedTimestamp: account.LastUpdatedTimestamp,
		columnKeyEventsName:           account.Data.Competition.Name,
		columnKeyEventsWhen:           account.Data.Competition.When,
		columnKeyCreatedTimestamp:     account.CreatedTimestamp,
	})
}
func GetSearchNullsTable(accounts accounts.Accounts) table.Model {
	var rows []table.Row
	for _, account := range accounts {
		rows = append(rows, makeRow(account))
	}

	table := table.New([]table.Column{
		table.NewColumn(columnKeyUID, "UID", 40),
		table.NewColumn(columnKeyEmail, "Email", 20),
		table.NewColumn(columnKeyFirstName, "First Name", 15),
		table.NewColumn(columnKeyLastName, "Last Name", 15),
		table.NewColumn(columnKeyHasLiteAccount, "Has Lite Account", 15),
		table.NewColumn(columnKeyHasFullAccount, "Has Full Account", 15),

		table.NewColumn(columnKeyLastUpdatedTimestamp, "Last Updated Timestamp", 15),
		table.NewColumn(columnKeyEventsName, "Events Name", 15),
		table.NewColumn(columnKeyEventsWhen, "Events When", 15),
		table.NewColumn(columnKeyCreatedTimestamp, "Created Timestamp", 15),
	}).WithRows(rows).
		BorderRounded().
		WithBaseStyle(StyleBase).
		WithPageSize(6).
		SortByDesc(columnKeyUID).
		Focused(true)

	return table
}
