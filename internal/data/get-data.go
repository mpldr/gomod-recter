package data

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

func (p *Project) GetData() error {
	var ustr strings.Builder

	ustr.WriteString(viper.GetString("Proxy.Address"))
	ustr.WriteString("/")
	ustr.WriteString(p.Repo)
	ustr.WriteString("/@v/list")

	u, err := url.Parse(ustr.String())
	if err != nil {
		return fmt.Errorf("unable to parse '%s' as URL: %w", ustr.String(), err)
	}

	res, err := http.Get(u.String())
	if err != nil {
		return fmt.Errorf("listing versions failed: %w", err)
	}
}
