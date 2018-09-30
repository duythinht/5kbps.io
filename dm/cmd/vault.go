package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.5kbps.io/dm/vault"
)

var (
	VaultCommand = &cobra.Command{
		Use:   "vault",
		Short: "encrypt/decrypt secret",
	}
	vaultParams = struct {
		File      string
		String    string
		Key       string
		RawOutput bool
	}{}
)

func init() {

	encrypt := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a file or text",
		Run: func(cmd *cobra.Command, args []string) {

			switch {
			case vaultParams.Key == "":
				fmt.Fprint(os.Stderr, "You must enter encrypt key to --key or ENCRYPTION_KEY env")
			case vaultParams.File != "":
				result := vault.EncryptFile(vaultParams.Key, vaultParams.File)
				_ = result
			case vaultParams.String != "":
				result, err := vault.EncryptText(vaultParams.String, vaultParams.Key)
				if err != nil {
					fmt.Println("error", err)
				}
				fmt.Println(result)
			default:
				fmt.Fprint(os.Stderr, "You must specific to encrypt file or text")
			}
		},
	}
	encrypt.Flags().StringVarP(&vaultParams.File, "file", "f", "", "File path")
	encrypt.Flags().StringVarP(&vaultParams.String, "string", "s", "", "Text")
	encrypt.Flags().StringVarP(&vaultParams.Key, "key", "k", os.Getenv("ENCRYPTION_KEY"), "Encrypt key")

	decrypt := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt a file or text",
		Run: func(cmd *cobra.Command, args []string) {
			switch {
			case vaultParams.Key == "":
				fmt.Fprint(os.Stderr, "You must enter encrypt key to --key or ENCRYPTION_KEY env")
			case vaultParams.File != "":
				f, err := os.Open(vaultParams.File)
				if err != nil {
					fmt.Println("error", err)
				}
				result, err := vault.DecryptFile(f, vaultParams.Key)
				if err != nil {
					fmt.Println("error", err)
				}
				fmt.Println(result)
			case vaultParams.String != "":
				result, err := vault.DecryptText(vaultParams.String, vaultParams.Key)
				if err != nil {
					fmt.Println("error", err)
				}
				fmt.Println(result)
			default:
				fmt.Fprint(os.Stderr, "You must specific to decrypt file or text")
			}

		},
	}

	decrypt.Flags().StringVarP(&vaultParams.File, "file", "f", "", "File path")
	decrypt.Flags().StringVarP(&vaultParams.String, "string", "s", "", "Text")
	decrypt.Flags().StringVarP(&vaultParams.Key, "key", "k", os.Getenv("ENCRYPTION_KEY"), "Encrypt key")

	generateKey := &cobra.Command{
		Use:   "generate-key",
		Short: "Generate encryption key",
		Run: func(cmd *cobra.Command, args []string) {
			if vaultParams.RawOutput {
				fmt.Println(vault.GenerateEncryptionKey())
				return
			}
			fmt.Printf("export ENCRYPTION_KEY=\"%s\"\n", vault.GenerateEncryptionKey())
			fmt.Printf("# This command is meant to be used with your shell's eval function.\n")
			fmt.Printf("# Run 'eval $(dm vault generate-key)' to generate and use encryption key.\n")
			fmt.Printf("# If you wish to use the session token itself, pass the --raw flag value.\n")
		},
	}

	generateKey.Flags().BoolVarP(&vaultParams.RawOutput, "raw", "r", false, "Raw output")

	VaultCommand.AddCommand(encrypt, decrypt, generateKey)
}
