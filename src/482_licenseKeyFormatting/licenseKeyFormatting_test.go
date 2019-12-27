package _82_licenseKeyFormatting

import (
	"fmt"
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	fmt.Println(LicenseKeyFormatting("5F3Z-2e-9-w",4))
}

func TestLicenseKeyFormatting2(t *testing.T) {
	fmt.Println(LicenseKeyFormatting("2-5g-3-J",2))
}

func TestLicenseKeyFormatting3(t *testing.T) {
	fmt.Println(LicenseKeyFormatting("aaaa",2))
}

func TestLicenseKeyFormatting4(t *testing.T) {
	fmt.Println(LicenseKeyFormatting("---",3))
}