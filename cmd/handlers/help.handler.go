package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func HelpHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	// Get the comments from this source file
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "./cmd/handlers", nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing directory: %v", err)
	}

	// Create the help message by iterating over the declarations in this file and looking for handlers.
	var helpText strings.Builder
	helpText.WriteString("List of available commands:\n")
	for _, p := range pkgs {
		for _, f := range p.Files {
			for _, decl := range f.Decls {
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					if funcDecl.Doc != nil {
						for _, comment := range funcDecl.Doc.List {
							if strings.HasPrefix(comment.Text, "// func") {
								helpText.WriteString(strings.TrimPrefix(comment.Text, "// func") + "\n")
							}
						}
					}
				}
			}
		}
	}

	// Send help message
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText.String())
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = update.Message.MessageID
	_, err = bot.Send(msg)
	return err
}
