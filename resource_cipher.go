package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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

func resourceCipherCreate(d *schema.ResourceData, m interface{}) error {
	filename := d.Get("filename").(string)
	if !isFileExist(filename) {
		return fmt.Errorf("File %s does not exist.", filename)
	}
	d.SetId(filename)
	return nil
}
func resourceCipherRead(d *schema.ResourceData, m interface{}) error {
	filename := d.Get("filename").(string)
	if d.Id() != filename {
		if !isFileExist(filename) {
			d.SetId("")
			return nil
		}
	}
	return d.Set("filename", filename)
}
func resourceCipherUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	if d.HasChange("filename") {
		filename := d.Get("filename").(string)
		if !isFileExist(filename) {
			return fmt.Errorf("File %s does not exist.", filename)
		}
	}
	d.Partial(false)
	return nil
}
func resourceCipherDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
