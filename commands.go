package main

import (
	"github.com/jamesacampbell/gotweed/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"add-draft": func() (cli.Command, error) {
			return &command.AddDraftCommand{
				Meta: *meta,
			}, nil
		},
		"list-drafts": func() (cli.Command, error) {
			return &command.ListDraftsCommand{
				Meta: *meta,
			}, nil
		},
		"authorize": func() (cli.Command, error) {
			return &command.AuthorizeCommand{
				Meta: *meta,
			}, nil
		},
		"post-draft": func() (cli.Command, error) {
			return &command.PostDraftCommand{
				Meta: *meta,
			}, nil
		},
		"archive-user": func() (cli.Command, error) {
			return &command.ArchiveUserCommand{
				Meta: *meta,
			}, nil
		},
		"search-user": func() (cli.Command, error) {
			return &command.SearchUserCommand{
				Meta: *meta,
			}, nil
		},
		"generate-markov": func() (cli.Command, error) {
			return &command.GenerateMarkovCommand{
				Meta: *meta,
			}, nil
		},
		"mentioner": func() (cli.Command, error) {
			return &command.MentionerCommand{
				Meta: *meta,
			}, nil
		},
		"add-to-mentioner-list": func() (cli.Command, error) {
			return &command.AddToMentionerListCommand{
				Meta: *meta,
			}, nil
		},
		"set-project": func() (cli.Command, error) {
			return &command.SetProjectCommand{
				Meta: *meta,
			}, nil
		},
		"load-project": func() (cli.Command, error) {
			return &command.LoadProjectCommand{
				Meta: *meta,
			}, nil
		},
		"delete-project": func() (cli.Command, error) {
			return &command.DeleteProjectCommand{
				Meta: *meta,
			}, nil
		},
		"delete-draft": func() (cli.Command, error) {
			return &command.DeleteDraftCommand{
				Meta: *meta,
			}, nil
		},
		"de-auth": func() (cli.Command, error) {
			return &command.DeAuthCommand{
				Meta: *meta,
			}, nil
		},
		"check-api": func() (cli.Command, error) {
			return &command.CheckApiCommand{
				Meta: *meta,
			}, nil
		},
		"get-random": func() (cli.Command, error) {
			return &command.GetRandomCommand{
				Meta: *meta,
			}, nil
		},
		"generate-random": func() (cli.Command, error) {
			return &command.GenerateRandomCommand{
				Meta: *meta,
			}, nil
		},
		"associate-user-with-project": func() (cli.Command, error) {
			return &command.AssociateUserWithProjectCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
