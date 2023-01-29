package tools

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"golang.org/x/net/context"
	"log"
	"testing"
)

func TestVault(t *testing.T) {
	config := vault.DefaultConfig()
	config.Address = "http://cloud-ip:8200"
	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
		return
	}
	client.SetToken("VAULT_TOKEN")

	secretData := map[string]interface{}{
		"password": "Hashi123",
		"hello":    "hello world",
	}
	_, err = client.KVv2("kv").Put(context.Background(), "my-secret-password", secretData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
		return
	}
	fmt.Println("Secret written successfully.")

	secret, err := client.KVv2("kv").Get(context.Background(), "my-secret-password")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
		return
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}
	if value != "Hashi123" {
		log.Fatalf("unexpected password value %q retrived from vault", value)
		return
	}
	fmt.Println("Access granted!")

}
