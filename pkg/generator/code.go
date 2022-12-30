package codegenerator

import (
	"fmt"

	"github.com/AmirSoleimani/VoucherCodeGenerator/vcgen"
)

func Generate() (string, error) {
	vc, err := vcgen.NewWithOptions(
		vcgen.SetCount(1),
		vcgen.SetPattern("###-###-####"),
		vcgen.SetCharset("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"),
	)
	if err != nil {
		return "", fmt.Errorf("fail construct invitation code generator: %v", err)
	}
	codes, err := vc.Run()
	if err != nil {
		return "", fmt.Errorf("fail run invitation code generator: %v", err)
	}
	return codes[0], nil
}
