package Unify

import "GoWhen/Unify/cmdConfig"


func PopArg(args []string) (string, []string) {
	if len(args) == 0 {
		return "", args
	}
	return (args)[0], (args)[1:]
}

func PopArgs(cull int, args []string) ([]string, []string) {
	if cull > len(args) {
		args = cmdConfig.FillArray(cull, args)
		return args, []string{}
	}
	if len(args) == 0 {
		return []string{}, args
	}
	return (args)[:cull], (args)[cull:]
}

func IsLastArg(args []string) bool {
	if len(args) == 0 {
		return true
	}
	return false
}
