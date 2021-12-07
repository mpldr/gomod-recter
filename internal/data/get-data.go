package data

import (
	"crypto/tls"
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
		glog.Errorf("could not parse repo address as URL: %v", err)
		return fmt.Errorf("could not parse repo address as URL: %w", err)
	}

	ustr.WriteString(viper.GetString("Proxy.Address"))
	ustr.WriteString(repoaddr.Host + repoaddr.Path)

	glog.Debugf("Parsing URL: %s/@v/list", ustr.String())
	u, err := url.Parse(ustr.String() + "/@v/list")
	if err != nil {
		glog.Errorf("unable to parse '%s' as URL: %w", ustr.String(), err)
		return fmt.Errorf("unable to parse '%s' as URL: %w", ustr.String(), err)
	}

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: viper.GetBool("Proxy.IgnoreCert")}}}

	glog.Debugf("Retrieving URL: %s/@v/list", ustr.String())
	res, err := client.Get(u.String())
	if err != nil {
		glog.Errorf("listing versions failed: %v", err)
		return fmt.Errorf("listing versions failed: %w", err)
	}
	defer res.Body.Close()

	glog.Debug(res.Status)
	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			glog.Errorf("could not read responsebody: %v", err)
			return fmt.Errorf("could not read responsebody: %w", err)
		}

		bodystring := string(body)
		glog.Tracef("response body: %s", bodystring)

		versions := strings.Split(bodystring, "\n")

		semver.Sort(versions)

		for i, j := 0, len(versions)-1; i < j; i, j = i+1, j-1 {
			versions[i], versions[j] = versions[j], versions[i]
		}

		versions = versions[:len(versions)-1]

		glog.Debugf("sorted versions: %v", versions)

		p.Versions = versions
	}

	return nil
}

func (p *Project) GetGoSource() string {
	// I can't know every format
	if p.GoSourceFmt != "" {
		glog.Debug("custom go-source header defined.")
		return fmt.Sprintf(`<meta name="go-source" content="%s _ %s">`, p.RootPath, p.GoSourceFmt)
	}

	if p.DefaultBranch == "" {
		glog.Debug("no default branch is set. cannot generate go-source header.")
		return ""
	}

	u, err := url.Parse(strings.TrimRight(p.Repo, "/"))
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		glog.Errorf("failed to parse rootpath as URL: %v", err)
		return ""
	}
	glog.Debugf("Host from RootPath: %s", u.Host)
	dir := dirFromUrl(u, p.DefaultBranch)
	file := fileFromUrl(u, p.DefaultBranch)

	if dir == "" || file == "" {
		return ""
	}

	return fmt.Sprintf(`<meta name="go-source" content="%s _ %s %s">`, p.RootPath, dir, file)
}

func dirFromUrl(u *url.URL, DefaultBranch string) string {
	switch u.Host {
	// everything but sourcehut should be considered untested!
	case "git.sr.ht", "github.com":
		return fmt.Sprintf("%s/tree/%s{/dir}", u, DefaultBranch)
	case "gitlab.com":
		return fmt.Sprintf("%s/-/tree/%s{/dir}", u, DefaultBranch)
	case "codeberg.org", "git.disroot.org": // gitea
		return fmt.Sprintf("%s/src/branch/%s{/dir}", u, DefaultBranch)
	case "notabug.org": // gogs
		return fmt.Sprintf("%s/src/%s{/dir}", u, DefaultBranch)
	case "git.code.sf.net": // sourceforge.net, is someone still using that?
		return fmt.Sprintf("https://sourceforge.net%s/ci/%s/tree{/dir}", u.Path, DefaultBranch)
	case "repo.or.cz":
		return fmt.Sprintf("%s/tree/refs/heads/%s:{/dir}", u, DefaultBranch)

		// These are selfhosted instances. You can specify them in the GoSourceFmt key
		//
		// case "cgit":
		//	return "https://[cgit-instance]/[repo]/tree/{/dir}?h=[default-branch]"
		// case "gitea":
		//	return "https://[gitea-instance]/[user]/[repo]/src/branch/[default-branch]{/dir}"
		// case "gogs":
		//	return "https://[gogs-instance]/[user]/[repo]/src/[default-branch]{/dir}"
	}
	return ""
}

func fileFromUrl(u *url.URL, DefaultBranch string) string {
	switch u.Host {
	// everything but sourcehut should be considered untested!
	case "git.sr.ht", "github.com":
		return fmt.Sprintf("%s/tree/%s{/dir}/{file}#L{line}", u, DefaultBranch)
	case "gitlab.com":
		return fmt.Sprintf("%s/-/tree/%s{/dir}/{file}#L{line}", u, DefaultBranch)
	case "codeberg.org", "git.disroot.org": // gitea
		return fmt.Sprintf("%s/src/branch/%s{/dir}/{file}#L{line}", u, DefaultBranch)
	case "notabug.org": // gogs
		return fmt.Sprintf("%s/src/%s{/dir}/{file}#L{line}", u, DefaultBranch)
	case "git.code.sf.net": // sourceforge.net, is someone still using that?
		return fmt.Sprintf("https://sourceforge.net%s/ci/%s/tree{/dir}/{file}#l{line}", u.Path, DefaultBranch)
	case "repo.or.cz":
		return fmt.Sprintf("%s/tree/refs/heads/%s:{/dir}/{file}#l{line}", u, DefaultBranch)

		// These are selfhosted instances. You can specify them in the GoSourceFmt key
		//
		// case "cgit":
		//	return "https://[cgit-instance]/[repo]/tree/{/dir}/{file}?h=[default-branch]#n{line}"
		// case "gitea":
		//	return "https://[gitea-instance]/[user]/[repo]/src/branch/[default-branch]{/dir}/{file}#L{line}"
		// case "gogs":
		//	return "https://[gogs-instance]/[user]/[repo]/src/[default-branch]{/dir}/{file}#L{line}"
	}
	return ""
}
