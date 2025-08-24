package cli

import (
	"context"
	"os"
	"strings"

	"github.com/libdns/libdns"
	"github.com/libdns/tencentcloud"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dnsCmd)
}

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "set dns for cert",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			log.Fatal().Msg("dns command requires exactly 3 arguments: name type value")
		}
		setDNS(args[0], args[1], args[2])
	},
}

func setDNS(name, kind, value string) {
	if strings.ToUpper(kind) != "TXT" {
		log.Fatal().Msg("only txt records are supported")
	}
	provider := &tencentcloud.Provider{
		SecretId:     os.Getenv("TENCENTCLOUD_SECRET_ID"),
		SecretKey:    os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	}
	zone := os.Getenv("TENCENTCLOUD_ZONE")
	log.Info().Msgf("Setting DNS record %s to %s, zone %s", name, value, zone)
	res, err := provider.SetRecords(context.Background(), zone, []libdns.Record{
		libdns.TXT{Name: name, Text: value},
	})
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("Error setting DNS record")
	}
	log.Info().Msgf("Set DNS record: %+v", res)
}
