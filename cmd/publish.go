package cmd

import (
	"github.com/smira/aptly/utils"
	"github.com/smira/commander"
	"github.com/smira/flag"
)

func getSigner(flags *flag.FlagSet) (utils.Signer, error) {
	if flags.Lookup("skip-signing").Value.Get().(bool) || context.Config().GpgDisableSign {
		return nil, nil
	}

	signer := &utils.GpgSigner{}
	signer.SetKey(flags.Lookup("gpg-key").Value.String())
	signer.SetKeyRing(flags.Lookup("keyring").Value.String(), flags.Lookup("secret-keyring").Value.String())

	err := signer.Init()
	if err != nil {
		return nil, err
	}

	return signer, nil

}

func makeCmdPublish() *commander.Command {
	return &commander.Command{
		UsageLine: "publish",
		Short:     "manage published repositories",
		Subcommands: []*commander.Command{
			makeCmdPublishDrop(),
			makeCmdPublishList(),
			makeCmdPublishRepo(),
			makeCmdPublishSnapshot(),
			makeCmdPublishSwitch(),
			makeCmdPublishUpdate(),
		},
	}
}
