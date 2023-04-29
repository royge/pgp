package age

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"

	"filippo.io/age"
	"filippo.io/age/armor"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func id(input string) string {
	hash := sha256.Sum256([]byte(input))

	return fmt.Sprintf("%x", hash)
}

func resourceCipher() *schema.Resource {
	return &schema.Resource{
		Create: resourceCipherCreate,
		Read:   resourceCipherRead,
		Update: resourceCipherUpdate,
		Delete: resourceCipherDelete,
		Schema: map[string]*schema.Schema{
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"result": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
				ForceNew:  true,
			},
		},
	}
}

func isFileExist(filename string) bool {
	dir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(dir + "/" + filename)
	return !os.IsNotExist(err)
}

func getEnv(env string) (string, error) {
	value := os.Getenv(env)
	if value == "" {
		return "", fmt.Errorf("You haven't set '%v' in the environment variable.", env)
	}

	return value, nil
}

func resourceCipherCreate(d *schema.ResourceData, m interface{}) error {
	filename := d.Get("filename").(string)
	if !isFileExist(filename) {
		return fmt.Errorf("File %s does not exist.", filename)
	}
	d.SetId(id(filename))

	privateKey, err := getEnv("AGE_PRIVATE_KEY")
	if err != nil {
		return err
	}

	res, err := decrypt(filename, privateKey)
	if err != nil {
		return fmt.Errorf("Unable to descrypt '%v' because '%v'", filename, err)
	}

	return d.Set("result", res)
}

func resourceCipherRead(d *schema.ResourceData, m interface{}) error {
	// filename := d.Get("filename").(string)
	// if !isFileExist(filename) {
	// 	return fmt.Errorf("File %s does not exist.", filename)
	// }
	// d.SetId(filename)
	// return d.Set("filename", filename)
	return nil
}

func resourceCipherUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	if d.HasChange("filename") {
		filename := d.Get("filename").(string)
		if !isFileExist(filename) {
			return fmt.Errorf("File %s does not exist.", filename)
		}
		d.SetId(id(filename))

		privateKey, err := getEnv("AGE_PRIVATE_KEY")
		if err != nil {
			return err
		}

		res, err := decrypt(filename, privateKey)
		if err != nil {
			return fmt.Errorf("Unable to descrypt '%v' because '%v'", filename, err)
		}

		return d.Set("result", res)
	}
	d.Partial(false)
	return nil
}

func resourceCipherDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func decrypt(
	cipherFile, privateKey string,
) (string, error) {
	keyFile, err := os.Open(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to open private keys file: %v", err)
	}
	identities, err := age.ParseIdentities(keyFile)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	encrypted, err := os.ReadFile(cipherFile)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}

	f := strings.NewReader(string(encrypted))
	armorReader := armor.NewReader(f)

	r, err := age.Decrypt(armorReader, identities...)
	if err != nil {
		return "", fmt.Errorf("failed to open encrypted file: %v", err)
	}
	out := &bytes.Buffer{}
	if _, err := io.Copy(out, r); err != nil {
		return "", fmt.Errorf("failed to read encrypted file: %v", err)
	}

	return out.String(), nil
}
