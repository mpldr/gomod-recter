package data

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/mod/semver"

	"git.sr.ht/~poldi1405/glog"
	"github.com/spf13/viper"
)

func (p *Project) GetData() error {
	glog.Infof("updating project info for %s", p.Name)
	var ustr strings.Builder

	repoaddr, err := url.Parse(p.Repo)
	if err != nil {
		return fmt.Errorf("could not parse repo address as URL: %w", err)
	}

	ustr.WriteString(viper.GetString("Proxy.Address"))
	ustr.WriteString(repoaddr.Host + repoaddr.Path)

	glog.Debugf("Parsing URL: %s/@v/list", ustr.String())
	u, err := url.Parse(ustr.String() + "/@v/list")
	if err != nil {
		return fmt.Errorf("unable to parse '%s' as URL: %w", ustr.String(), err)
	}

	glog.Debugf("Retrieving URL: %s/@v/list", ustr.String())
	res, err := http.Get(u.String())
	if err != nil {
		return fmt.Errorf("listing versions failed: %w", err)
	}
	defer res.Body.Close()

	glog.Debug(res.Status)
	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("could not read responsebody: %w", err)
		}

		bodystring := string(body)
		glog.Debugf("response body: %s", bodystring)

		versions := strings.Split(bodystring, "\n")

		semver.Sort(versions)

		for i, j := 0, len(versions)-1; i < j; i, j = i+1, j-1 {
			versions[i], versions[j] = versions[j], versions[i]
		}

		versions = versions[:len(versions)-1]

		glog.Debugf("sorted versions: %v", versions)
	}

	return nil
}
