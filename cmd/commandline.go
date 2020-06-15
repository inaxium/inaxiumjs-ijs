package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"inaxium.com/ijs/public"
)

var meta public.Meta = public.Meta{}

var mainCmd = &cobra.Command{
	Use:   "ijs",
	Short: "InaxiumJS Command Line Interpreter",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Init() public.Meta {

	viper.SetEnvPrefix("INAXIUMJS")
	viper.AutomaticEnv()

	flags := mainCmd.PersistentFlags()

	mainCmd.PersistentFlags().StringVarP(&meta.Version,"version", "v", "2.0", "InaxiumJS Version")
	mainCmd.PersistentFlags().StringVarP(&meta.Destination,"destination", "d", ".", "Destination (OS Directory)")
	mainCmd.PersistentFlags().StringVarP(&meta.Type,"type", "t", "framework", "framework | demo")
	mainCmd.PersistentFlags().BoolP("show", "s", true, "Show Versions")
	mainCmd.PersistentFlags().BoolP("copy", "c", true, "Copy Version to local Disk")

	_ = viper.BindPFlag("version",  flags.Lookup("version"))
	_ = viper.BindPFlag("destination", flags.Lookup("destination"))
	_ = viper.BindPFlag("type", flags.Lookup("type"))
	_ = viper.BindPFlag("show", flags.Lookup("show"))
	_ = viper.BindPFlag("copy", flags.Lookup("copy"))

	viper.SetConfigName("inaxiumjs")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()

	_ = mainCmd.Execute()

	return meta
}
