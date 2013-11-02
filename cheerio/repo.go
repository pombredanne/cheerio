package cheerio

import (
	"fmt"
	"regexp"
)

var repoPatterns = []*regexp.Regexp{
	regexp.MustCompile(`Home-page: (https?://github.com/(:?[^/\n]+)/(:?[^/\n]+))(:?/.*)?\n`),
	regexp.MustCompile(`Home-page: (https?://bitbucket.org/(:?[^/\n]+)/(:?[^/\n]+))(:?/.*)?\n`),
	regexp.MustCompile(`Home-page: (https?://code.google.com/p/(:?[^/\n]+))(:?/.*)?\n`),
}

var homepageRegexp = regexp.MustCompile(`Home-page: (.+)\n`)

func (p *PackageIndex) FetchSourceRepoURI(pkg string) (string, error) {
	pattern := "**/PKG-INFO"
	b, err := p.FetchRawMetadata(pkg, pattern, pattern, pattern)
	if err != nil {
		return "", nil
	}
	rawMetadata := string(b)

	// Check PyPI
	for _, pattern := range repoPatterns {
		if match := pattern.FindStringSubmatch(rawMetadata); len(match) >= 1 {
			return match[1], nil
		}
	}

	// Try to fall back to hard-coded URIs
	if hardURI, in := pypiRepos[NormalizedPkgName(pkg)]; in {
		return fmt.Sprintf("https://%s", hardURI), nil
	}

	// Return most informative error
	if match := homepageRegexp.FindStringSubmatch(rawMetadata); len(match) >= 1 {
		return "", fmt.Errorf("Could not parse repo URI from homepage: %s", match[1])
	}
	return "", fmt.Errorf("No homepage found in metadata: %s", rawMetadata)
}

var pypiRepos = map[string]string{
	"ansible":               "github.com/ansible/ansible",
	"apache-libcloud":       "github.com/apache/libcloud",
	"bottle":                "github.com/bottlepy/bottle",
	"celery":                "github.com/celery/celery",
	"chameleon":             "github.com/malthe/chameleon",
	"coverage":              "bitbucket.org/ned/coveragepy",
	"distribute":            "bitbucket.org/tarek/distribute",
	"django":                "github.com/django/django",
	"django-cms":            "github.com/divio/django-cms",
	"django-tastypie":       "github.com/toastdriven/django-tastypie",
	"djangocms-admin-style": "github.com/divio/djangocms-admin-style",
	"djangorestframework":   "github.com/tomchristie/django-rest-framework",
	"eve":               "github.com/nicolaiarocci/eve",
	"fabric":            "github.com/fabric/fabric",
	"flask":             "github.com/mitsuhiko/flask",
	"gevent":            "github.com/surfly/gevent",
	"gunicorn":          "github.com/benoitc/gunicorn",
	"httpie":            "github.com/jkbr/httpie",
	"httplib2":          "github.com/jcgregorio/httplib2",
	"itsdangerous":      "github.com/mitsuhiko/itsdangerous",
	"jinja2":            "github.com/mitsuhiko/jinja2",
	"kazoo":             "github.com/python-zk/kazoo",
	"kombu":             "github.com/celery/kombu",
	"lamson":            "github.com/zedshaw/lamson",
	"libcloud":          "github.com/apache/libcloud",
	"lxml":              "github.com/lxml/lxml",
	"mako":              "github.com/zzzeek/mako",
	"markupsafe":        "github.com/mitsuhiko/markupsafe",
	"matplotlib":        "github.com/matplotlib/matplotlib",
	"mimeparse":         "github.com/crosbymichael/mimeparse",
	"mock":              "github.com/beyang/mock",
	"nltk":              "github.com/nltk/nltk",
	"nose":              "github.com/nose-devs/nose",
	"nova":              "github.com/openstack/nova",
	"numpy":             "github.com/numpy/numpy",
	"pandas":            "github.com/pydata/pandas",
	"pastedeploy":       "bitbucket.org/ianb/pastedeploy",
	"pattern":           "github.com/clips/pattern",
	"psycopg2":          "github.com/beyang/psycopg2",
	"pyramid":           "github.com/Pylons/pyramid",
	"python-dateutil":   "github.com/paxan/python-dateutil",
	"python-lust":       "github.com/zedshaw/python-lust",
	"pyyaml":            "github.com/yaml/pyyaml",
	"repoze.lru":        "github.com/repoze/repoze.lru",
	"requests":          "github.com/kennethreitz/requests",
	"salt":              "github.com/saltstack/salt",
	"scikit-learn":      "github.com/scikit-learn/scikit-learn",
	"scipy":             "github.com/scipy/scipy",
	"sentry":            "github.com/getsentry/sentry",
	"setuptools":        "github.com/jaraco/setuptools",
	"sockjs-tornado":    "github.com/mrjoes/sockjs-tornado",
	"south":             "bitbucket.org/andrewgodwin/south",
	"sqlalchemy":        "github.com/zzzeek/sqlalchemy",
	"ssh":               "github.com/bitprophet/ssh",
	"tornado":           "github.com/facebook/tornado",
	"translationstring": "github.com/Pylons/translationstring",
	"tulip":             "github.com/sourcegraph/tulip",
	"venusian":          "github.com/Pylons/venusian",
	"webob":             "github.com/Pylons/webob",
	"webpy":             "github.com/webpy/webpy",
	"werkzeug":          "github.com/mitsuhiko/werkzeug",
	"zope.interface":    "github.com/zopefoundation/zope.interface",
}
