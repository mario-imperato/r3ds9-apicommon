package definitions

const (
	ApiKeyHeaderName = "X-R3ds9-Api-Key"
	SidHeaderName    = "X-R3ds9-Sid"
	UserHeaderName   = "X-R3ds9-User"
)

const (
	HostDomainUrlParam          = "hostDomain"
	HostSiteUrlParam            = "hostSite"
	HostLangUrlParam            = "hostLang"
	HostAppIdUrlParam           = "hostAppId"
	ApiContextUrlParam          = "apiCtx"
	ApiGtwExtraPathInfoUrlParam = "exPathInfo"

	ApiGtwUiGroup           = "/ui"
	ApiGtwUiConsoleGroup    = "/ui-console"
	ApiGtwUiPathPatternHome = ":hostDomain/:hostSite/:hostLang/:hostAppId"

	ApiGtwUiPathPatternNested = ":hostDomain/:hostSite/:hostLang/:hostAppId/*exPathInfo"
	ApiGtwApiGroupPattern     = ":apiCtx/:hostDomain/:hostSite/:hostLang"

	ApiCorePathGroupWhoAmI     = "/core/:hostDomain/:hostSite/:hostLang/host/user"
	ApiCorePathGroupHostSite   = "/core/:hostDomain/:hostSite/:hostLang/host/site"
	ApiCorePathGroupHostDomain = "/core/:hostDomain/:hostSite/:hostLang/host/domain"

	ApiCorePathGroupUser   = "/core/:hostDomain/:hostSite/:hostLang/user"
	ApiCorePathGroupSite   = "/core/:hostDomain/:hostSite/:hostLang/site"
	ApiCorePathGroupDomain = "/core/:hostDomain/:hostSite/:hostLang/domain"
)
