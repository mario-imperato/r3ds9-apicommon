package definitions

const (
	ApiKeyHeaderName   = "X-R3ds9-Api-Key"
	SidHeaderName      = "X-R3ds9-Sid"
	UserHeaderNickName = "X-R3ds9-Nickname"
	UserHeaderUserId   = "X-R3ds9-Userid"
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

	ApiCorePathGroupHost            = "/core/:hostDomain/:hostSite/:hostLang/host"
	ApiCorePathGroupUser            = "/core/:hostDomain/:hostSite/:hostLang/user"
	ApiCorePathGroupSite            = "/core/:hostDomain/:hostSite/:hostLang/site"
	ApiCorePathGroupDomain          = "/core/:hostDomain/:hostSite/:hostLang/domain"
	ApiCorePathGroupKeyValuePackage = "/core/:hostDomain/:hostSite/:hostLang/properties"

	ApiCmsPathGroupCard      = "/cms/:hostDomain/:hostSite/:hostLang/card"
	ApiCmsPathGroupUserCards = "/cms/:hostDomain/:hostSite/:hostLang/user-cards"
	ApiCmsPathGroupFiles     = "/cms/:hostDomain/:hostSite/:hostLang/files"
)

const (
	RootDomain   = "root"
	SiteWildCard = "*"
)
