package digest_auth_client

import "strings"

type wwwAuthenticate struct {
	Algorithm string // unquoted
	Domain    string // quoted
	Nonce     string // quoted
	Opaque    string // quoted
	Qop       string // quoted
	Realm     string // quoted
	Stale     bool   // unquoted
	Charset   string // quoted
	Userhash  bool   // quoted
}

func newWwwAuthenticate(s string) *wwwAuthenticate {

	var wa = wwwAuthenticate{}

	params := strings.Split(s, ",")
	for _, param := range params {
		if algorithmBegin := strings.Index(param, "algorithm"); algorithmBegin != -1 {
			wa.Algorithm = param[algorithmBegin+11 : strings.LastIndex(param, "\"")]
		}

		if domainBegin := strings.Index(param, "domain"); domainBegin != -1 {
			wa.Domain = param[domainBegin+8 : strings.LastIndex(param, "\"")]
		}

		if nonceBegin := strings.Index(param, "nonce"); nonceBegin != -1 {
			wa.Nonce = param[nonceBegin+7 : strings.LastIndex(param, "\"")]
		}

		if opaqueBegin := strings.Index(param, "opaque"); opaqueBegin != -1 {
			wa.Opaque = param[opaqueBegin+8 : strings.LastIndex(param, "\"")]
		}

		if qopBegin := strings.Index(param, "qop"); qopBegin != -1 {
			wa.Qop = param[qopBegin+5 : strings.LastIndex(param, "\"")]
		}

		if realmBegin := strings.Index(param, "realm"); realmBegin != -1 {
			wa.Realm = param[realmBegin+7 : strings.LastIndex(param, "\"")]
		}

		if realmBegin := strings.Index(param, "realm"); realmBegin != -1 {
			wa.Realm = param[realmBegin+7 : strings.LastIndex(param, "\"")]
		}

		if staleBegin := strings.Index(param, "stale"); staleBegin != -1 {
			if strings.ToUpper(param[staleBegin+7:strings.LastIndex(param, "\"")]) == "TRUE" {
				wa.Stale = true
			}
		}

		if charsetBegin := strings.Index(param, "charset"); charsetBegin != -1 {
			wa.Charset = param[charsetBegin+9 : strings.LastIndex(param, "\"")]
		}

		if userhashBegin := strings.Index(param, "userhash"); userhashBegin != -1 {
			if strings.ToUpper(param[userhashBegin+9:strings.LastIndex(param, "\"")]) == "TRUE" {
				wa.Userhash = true
			}
		}
	}
	// FIXME:正则表达匹配问题
	// algorithmRegex := regexp.MustCompile(`algorithm="([^ ,]+)"`)
	// algorithmMatch := algorithmRegex.FindStringSubmatch(s)
	// if algorithmMatch != nil {
	// 	wa.Algorithm = algorithmMatch[1]
	// }

	// domainRegex := regexp.MustCompile(`domain="(.+?)"`)
	// domainMatch := domainRegex.FindStringSubmatch(s)
	// if domainMatch != nil {
	// 	wa.Domain = domainMatch[1]
	// }

	// nonceRegex := regexp.MustCompile(`nonce="(.+?)"`)
	// nonceMatch := nonceRegex.FindStringSubmatch(s)
	// if nonceMatch != nil {
	// 	fmt.Printf("%+v\n", nonceMatch)
	// 	wa.Nonce = nonceMatch[1]
	// }

	// opaqueRegex := regexp.MustCompile(`opaque="(.+?)"`)
	// opaqueMatch := opaqueRegex.FindStringSubmatch(s)
	// fmt.Printf("%+v\n", opaqueMatch)
	// if opaqueMatch != nil {
	// 	wa.Opaque = opaqueMatch[1]
	// }

	// qopRegex := regexp.MustCompile(`qop="(.+?)"`)
	// qopMatch := qopRegex.FindStringSubmatch(s)
	// if qopMatch != nil {
	// 	wa.Qop = qopMatch[1]
	// }

	// realmRegex := regexp.MustCompile(`realm="(.+?)"`)
	// realmMatch := realmRegex.FindStringSubmatch(s)
	// if realmMatch != nil {
	// 	wa.Realm = realmMatch[1]
	// }

	// staleRegex := regexp.MustCompile(`stale=([^ ,])"`)
	// staleMatch := staleRegex.FindStringSubmatch(s)
	// if staleMatch != nil {
	// 	wa.Stale = (strings.ToLower(staleMatch[1]) == "true")
	// }

	// charsetRegex := regexp.MustCompile(`charset="(.+?)"`)
	// charsetMatch := charsetRegex.FindStringSubmatch(s)
	// if charsetMatch != nil {
	// 	wa.Charset = charsetMatch[1]
	// }

	// userhashRegex := regexp.MustCompile(`userhash=([^ ,])"`)
	// userhashMatch := userhashRegex.FindStringSubmatch(s)
	// if userhashMatch != nil {
	// 	wa.Userhash = (strings.ToLower(userhashMatch[1]) == "true")
	// }

	// fmt.Printf("%+v\n", wa)
	return &wa
}
