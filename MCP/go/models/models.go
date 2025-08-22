package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// VerifySecurityCodeResponse represents the VerifySecurityCodeResponse schema from the OpenAPI specification
type VerifySecurityCodeResponse struct {
	Verified bool `json:"verified"` // True if the code is valid
}

// BrowserBotResponse represents the BrowserBotResponse schema from the OpenAPI specification
type BrowserBotResponse struct {
	Exec_results []string `json:"exec-results"` // If you executed any JavaScript this array holds the results as objects
	Http_status_message string `json:"http-status-message"` // The HTTP status message the URL returned
	Mime_type string `json:"mime-type"` // The document MIME type
	Title string `json:"title"` // The document title
	Content string `json:"content"` // The complete raw, decompressed and decoded page content. Usually will be either HTML, JSON or XML
	Language_code string `json:"language-code"` // The ISO 2-letter language code of the page. Extracted from either the HTML document or via HTTP headers
	Error_message string `json:"error-message"` // Contains the error message if an error has occurred ('is-error' will be true)
	Is_http_redirect bool `json:"is-http-redirect"` // True if the URL responded with an HTTP redirect
	Is_secure bool `json:"is-secure"` // True if the page is secured using TLS/SSL
	Security_details map[string]interface{} `json:"security-details"` // Map containing details of the TLS/SSL setup
	Http_status_code int `json:"http-status-code"` // The HTTP status code the URL returned
	Is_timeout bool `json:"is-timeout"` // True if a timeout occurred while loading the page. You can set the timeout with the request parameter 'timeout'
	Load_time float64 `json:"load-time"` // The number of seconds taken to load the page (from initial request until DOM ready)
	Response_headers map[string]interface{} `json:"response-headers"` // Map containing all the HTTP response headers the URL responded with
	Url string `json:"url"` // The page URL
	Elements []string `json:"elements"` // Array containing all the elements matching the supplied selector. <br>Each element object will contain the text content, HTML content and all current element attributes
	Is_http_ok bool `json:"is-http-ok"` // True if the HTTP status is OK (200)
	Http_redirect_url string `json:"http-redirect-url"` // The redirected URL if the URL responded with an HTTP redirect
	Server_ip string `json:"server-ip"` // The HTTP servers IP address
	Is_error bool `json:"is-error"` // True if an error has occurred loading the page. Check the 'error-message' field for details
}

// ConvertResponse represents the ConvertResponse schema from the OpenAPI specification
type ConvertResponse struct {
	Result string `json:"result"` // The result of the conversion in string format
	Result_float float64 `json:"result-float"` // The result of the conversion as a floating-point number
	To_type string `json:"to-type"` // The type being converted to
	Valid bool `json:"valid"` // True if the conversion was successful and produced a valid result
	From_type string `json:"from-type"` // The type of the value being converted from
	From_value string `json:"from-value"` // The value being converted from
}

// BINLookupResponse represents the BINLookupResponse schema from the OpenAPI specification
type BINLookupResponse struct {
	Is_prepaid bool `json:"is-prepaid"` // Is this a prepaid or prepaid reloadable card
	Issuer_website string `json:"issuer-website"` // The card issuers website
	Card_type string `json:"card-type"` // The card type, will always be one of: DEBIT, CREDIT, CHARGE CARD
	Ip_matches_bin bool `json:"ip-matches-bin"` // True if the customers IP country matches the BIN country
	Ip_city string `json:"ip-city"` // The city of the customers IP (if detectable)
	Currency_code string `json:"currency-code"` // ISO 4217 currency code associated with the country of the issuer
	Country_code3 string `json:"country-code3"` // The ISO 3-letter country code of the issuer
	Ip_country_code3 string `json:"ip-country-code3"` // The ISO 3-letter country code of the customers IP
	Ip_blocklists []string `json:"ip-blocklists"` // An array of strings indicating which blocklists this IP is listed on
	Country string `json:"country"` // The full country name of the issuer
	Country_code string `json:"country-code"` // The ISO 2-letter country code of the issuer
	Issuer_phone string `json:"issuer-phone"` // The card issuers phone number
	Card_category string `json:"card-category"` // The card category. There are many different card categories the most common card categories are: CLASSIC, BUSINESS, CORPORATE, PLATINUM, PREPAID
	Ip_blocklisted bool `json:"ip-blocklisted"` // True if the customers IP is listed on one of our blocklists, see the <a href="http://www.neutrinoapi.com/api/ip-blocklist/">IP Blocklist API</a>
	Is_commercial bool `json:"is-commercial"` // Is this a commercial/business use card
	Ip_country string `json:"ip-country"` // The country of the customers IP
	Ip_country_code string `json:"ip-country-code"` // The ISO 2-letter country code of the customers IP
	Ip_region string `json:"ip-region"` // The region of the customers IP (if detectable)
	Valid bool `json:"valid"` // Is this a valid BIN or IIN number
	Issuer string `json:"issuer"` // The card issuer
	Card_brand string `json:"card-brand"` // The card brand (e.g. Visa or Mastercard)
	Bin_number string `json:"bin-number"` // The BIN or IIN number
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	Postal_code string `json:"postal-code"` // The postal code for the location
	Country_code string `json:"country-code"` // The ISO 2-letter country code of the location
	Timezone Timezone `json:"timezone"` // Map containing timezone details
	Location_type string `json:"location-type"` // The detected location type ordered roughly from most to least precise, possible values are: <br> <ul> <li>address - indicates a precise street address</li> <li>street - accurate to the street level but may not point to the exact location of the house/building number</li> <li>city - accurate to the city level, this includes villages, towns, suburbs, etc</li> <li>postal-code - indicates a postal code area (no house or street information present)</li> <li>railway - location is part of a rail network such as a station or railway track</li> <li>natural - indicates a natural feature, for example a mountain peak or a waterway</li> <li>island - location is an island or archipelago</li> <li>administrative - indicates an administrative boundary such as a country, state or province</li> </ul>
	Latitude float64 `json:"latitude"` // The location latitude
	Currency_code string `json:"currency-code"` // ISO 4217 currency code associated with the country
	Region_code string `json:"region-code"` // The ISO 3166-2 region code for the location
	Address_components map[string]interface{} `json:"address-components"` // The components which make up the address such as road, city, state, etc
	City string `json:"city"` // The city of the location
	Location_tags []string `json:"location-tags"` // Array of strings containing any location tags associated with the address. Tags are additional pieces of metadata about a specific location, there are thousands of different tags. Some examples of tags: shop, office, cafe, bank, pub
	Longitude float64 `json:"longitude"` // The location longitude
	State string `json:"state"` // The state of the location
	Country string `json:"country"` // The country of the location
	Postal_address string `json:"postal-address"` // The formatted address using local standards suitable for printing on an envelope
	Country_code3 string `json:"country-code3"` // The ISO 3-letter country code of the location
	Address string `json:"address"` // The complete address using comma-separated values
}

// APIError represents the APIError schema from the OpenAPI specification
type APIError struct {
	Api_error_msg string `json:"api-error-msg"` // API error message
	Api_error int `json:"api-error"` // API error code. If set and > 0 then an API error has occurred your request could not be completed
}

// IPInfoResponse represents the IPInfoResponse schema from the OpenAPI specification
type IPInfoResponse struct {
	Country_code3 string `json:"country-code3"` // ISO 3-letter country code
	Currency_code string `json:"currency-code"` // ISO 4217 currency code associated with the country
	Region_code string `json:"region-code"` // ISO 3166-2 region code (if detectable)
	Continent_code string `json:"continent-code"` // ISO 2-letter continent code
	Is_v6 bool `json:"is-v6"` // True if this is a IPv6 address. False if IPv4
	Timezone Timezone `json:"timezone"` // Map containing timezone details
	Host_domain string `json:"host-domain"` // The IPs host domain (only set if reverse-lookup has been used)
	Latitude float64 `json:"latitude"` // Location latitude
	Is_bogon bool `json:"is-bogon"` // True if this is a bogon IP address such as a private network, local network or reserved address
	Longitude float64 `json:"longitude"` // Location longitude
	Region string `json:"region"` // Name of the region (if detectable)
	Is_v4_mapped bool `json:"is-v4-mapped"` // True if this is a <a href="https://en.wikipedia.org/wiki/IPv6#IPv4-mapped_IPv6_addresses">IPv4 mapped IPv6 address</a>
	Country string `json:"country"` // Full country name
	Hostname string `json:"hostname"` // The IPs full hostname (only set if reverse-lookup has been used)
	Ip string `json:"ip"` // The IP address
	Country_code string `json:"country-code"` // ISO 2-letter country code
	Valid bool `json:"valid"` // True if this is a valid IPv4 or IPv6 address
	City string `json:"city"` // Name of the city (if detectable)
}

// UALookupResponse represents the UALookupResponse schema from the OpenAPI specification
type UALookupResponse struct {
	Device_pixel_ratio float64 `json:"device-pixel-ratio"` // The device display pixel ratio (the ratio of the resolution in physical pixels to the resolution in CSS pixels)
	Device_width_px float64 `json:"device-width-px"` // The device display width in CSS 'px'
	Device_model string `json:"device-model"` // The device model
	Os_family string `json:"os-family"` // The operating system family. The major OS families are: Android, Windows, macOS, iOS, Linux
	Os_version_major string `json:"os-version-major"` // The operating system major version
	Browser_engine string `json:"browser-engine"` // If the client is a web browser which underlying browser engine does it use
	Device_resolution string `json:"device-resolution"` // The device display resolution in physical pixels (e.g. 720x1280)
	Name string `json:"name"` // The client software name
	Ua string `json:"ua"` // The user agent string
	Device_price float64 `json:"device-price"` // The average device price on release in USD
	Is_webview bool `json:"is-webview"` // Is this a WebView / embedded software client
	Os string `json:"os"` // The full operating system name
	Version_major string `json:"version-major"` // The client software major version
	Browser_release string `json:"browser-release"` // If the client is a web browser which year was this browser version released
	Version string `json:"version"` // The client software full version
	Os_version string `json:"os-version"` // The operating system full version
	TypeField string `json:"type"` // The user agent type, possible values are: <br> <ul> <li>desktop</li> <li>phone</li> <li>tablet</li> <li>wearable</li> <li>tv</li> <li>console</li> <li>email</li> <li>library</li> <li>robot</li> <li>unknown</li> </ul>
	Device_release string `json:"device-release"` // The year when this device model was released
	Device_height_px float64 `json:"device-height-px"` // The device display height in CSS 'px'
	Device_model_code string `json:"device-model-code"` // The device model code
	Device_ppi float64 `json:"device-ppi"` // The device display PPI (pixels per inch)
	Is_mobile bool `json:"is-mobile"` // Is this a mobile device (e.g. a phone or tablet)
	Device_brand string `json:"device-brand"` // The device brand / manufacturer
}

// HLRLookupResponse represents the HLRLookupResponse schema from the OpenAPI specification
type HLRLookupResponse struct {
	Number_type string `json:"number-type"` // The number type, possible values are: <br> <ul> <li>mobile</li> <li>fixed-line</li> <li>premium-rate</li> <li>toll-free</li> <li>voip</li> <li>unknown</li> </ul>
	Roaming_country_code string `json:"roaming-country-code"` // If the number is currently roaming, the ISO 2-letter country code of the roaming in country
	Country_code string `json:"country-code"` // The number location as an ISO 2-letter country code
	Country_code3 string `json:"country-code3"` // The number location as an ISO 3-letter country code
	Is_mobile bool `json:"is-mobile"` // True if this is a mobile number (only true with 100% certainty, if the number type is unknown this value will be false)
	Local_number string `json:"local-number"` // The number represented in local dialing format
	Msc string `json:"msc"` // The mobile MSC number (Mobile Switching Center)
	Msin string `json:"msin"` // The mobile MSIN number (Mobile Subscription Identification Number)
	Hlr_valid bool `json:"hlr-valid"` // Was the HLR lookup successful. If true then this is a working and registered cell-phone or mobile device (SMS and phone calls will be delivered)
	International_number string `json:"international-number"` // The number represented in full international format
	Mnc string `json:"mnc"` // The mobile MNC number (Mobile Network Code)
	International_calling_code string `json:"international-calling-code"` // The international calling code
	Number_valid bool `json:"number-valid"` // True if this a valid phone number
	Ported_network string `json:"ported-network"` // The ported to network/carrier name (only set if the number has been ported)
	Imsi string `json:"imsi"` // The mobile IMSI number (International Mobile Subscriber Identity)
	Is_ported bool `json:"is-ported"` // Has this number been ported to another network
	Origin_network string `json:"origin-network"` // The origin network/carrier name
	Mcc string `json:"mcc"` // The mobile MCC number (Mobile Country Code)
	Is_roaming bool `json:"is-roaming"` // Is this number currently roaming from its origin country
	Currency_code string `json:"currency-code"` // ISO 4217 currency code associated with the country
	Hlr_status string `json:"hlr-status"` // The HLR lookup status, possible values are: <br> <ul> <li>ok - the HLR lookup was successful and the device is connected</li> <li>absent - the number was once registered but the device has been switched off or out of network range for some time</li> <li>unknown - the number is not known by the mobile network</li> <li>invalid - the number is not a valid mobile MSISDN number</li> <li>fixed-line - the number is a registered fixed-line not mobile</li> <li>voip - the number has been detected as a VOIP line</li> <li>failed - the HLR lookup has failed, we could not determine the real status of this number</li> </ul>
	Current_network string `json:"current-network"` // The currently used network/carrier name
	Country string `json:"country"` // The phone number country
	Location string `json:"location"` // The number location. Could be a city, region or country depending on the type of number
}

// IPProbeResponse represents the IPProbeResponse schema from the OpenAPI specification
type IPProbeResponse struct {
	City string `json:"city"` // Full city name (if detectable)
	Valid bool `json:"valid"` // True if this is a valid IPv4 or IPv6 address
	As_country_code3 string `json:"as-country-code3"` // The autonomous system (AS) ISO 3-letter country code
	Provider_type string `json:"provider-type"` // The detected provider type, possible values are: <br> <ul> <li>isp - IP belongs to an internet service provider. This includes both mobile, home and business internet providers</li> <li>hosting - IP belongs to a hosting company. This includes website hosting, cloud computing platforms and colocation facilities</li> <li>vpn - IP belongs to a VPN provider</li> <li>proxy - IP belongs to a proxy service. This includes HTTP/SOCKS proxies and browser based proxies</li> <li>university - IP belongs to a university/college/campus</li> <li>government - IP belongs to a government department. This includes military facilities</li> <li>commercial - IP belongs to a commercial entity such as a corporate headquarters or company office</li> <li>unknown - could not identify the provider type</li> </ul>
	Asn string `json:"asn"` // The autonomous system (AS) number
	Ip string `json:"ip"` // The IP address
	As_cidr string `json:"as-cidr"` // The autonomous system (AS) CIDR range
	Region string `json:"region"` // Full region name (if detectable)
	Region_code string `json:"region-code"` // ISO 3166-2 region code (if detectable)
	Is_hosting bool `json:"is-hosting"` // True if this IP belongs to a hosting company. Note that this can still be true even if the provider type is VPN/proxy, this occurs in the case that the IP is detected as both types
	Continent_code string `json:"continent-code"` // ISO 2-letter continent code
	Provider_domain string `json:"provider-domain"` // The domain name of the provider
	As_country_code string `json:"as-country-code"` // The autonomous system (AS) ISO 2-letter country code
	As_description string `json:"as-description"` // The autonomous system (AS) description / company name
	Host_domain string `json:"host-domain"` // The IPs host domain
	Vpn_domain string `json:"vpn-domain"` // The domain of the VPN provider (may be empty if the VPN domain is not detectable)
	Is_bogon bool `json:"is-bogon"` // True if this is a bogon IP address such as a private network, local network or reserved address
	Hostname string `json:"hostname"` // The IPs full hostname (PTR)
	Is_v4_mapped bool `json:"is-v4-mapped"` // True if this is a <a href="https://en.wikipedia.org/wiki/IPv6#IPv4-mapped_IPv6_addresses">IPv4 mapped IPv6 address</a>
	Provider_website string `json:"provider-website"` // The website URL for the provider
	As_age int `json:"as-age"` // The age of the autonomous system (AS) in number of years since registration
	Country_code string `json:"country-code"` // ISO 2-letter country code
	Is_proxy bool `json:"is-proxy"` // True if this IP ia a proxy
	Country string `json:"country"` // Full country name
	Provider_description string `json:"provider-description"` // A description of the provider (usually extracted from the providers website)
	Is_isp bool `json:"is-isp"` // True if this IP belongs to an internet service provider. Note that this can still be true even if the provider type is VPN/proxy, this occurs in the case that the IP is detected as both types
	Country_code3 string `json:"country-code3"` // ISO 3-letter country code
	Is_v6 bool `json:"is-v6"` // True if this is a IPv6 address. False if IPv4
	Is_vpn bool `json:"is-vpn"` // True if this IP ia a VPN
	Currency_code string `json:"currency-code"` // ISO 4217 currency code associated with the country
	As_domains []string `json:"as-domains"` // Array of all the domains associated with the autonomous system (AS)
}

// URLInfoResponse represents the URLInfoResponse schema from the OpenAPI specification
type URLInfoResponse struct {
	Is_timeout bool `json:"is-timeout"` // True if a timeout occurred while loading the URL. You can set the timeout with the request parameter 'timeout'
	Load_time float64 `json:"load-time"` // The time taken to load the URL content in seconds
	Server_city string `json:"server-city"` // The servers IP geo-location: full city name (if detectable)
	Server_name string `json:"server-name"` // The name of the server software hosting this URL
	Language_code string `json:"language-code"` // The ISO 2-letter language code of the page. Extracted from either the HTML document or via HTTP headers
	Server_country string `json:"server-country"` // The servers IP geo-location: full country name
	Url string `json:"url"` // The fully qualified URL. This may be different to the URL requested if http-redirect is true
	Server_country_code string `json:"server-country-code"` // The servers IP geo-location: ISO 2-letter country code
	Server_ip string `json:"server-ip"` // The IP address of the server hosting this URL
	Url_port int `json:"url-port"` // The URL port
	Content_encoding string `json:"content-encoding"` // The encoding format the URL uses
	Content_size int `json:"content-size"` // The size of the URL content in bytes
	Http_status_message int `json:"http-status-message"` // The HTTP status message assoicated with the status code
	Query map[string]interface{} `json:"query"` // A key-value map of the URL query paramaters
	Title string `json:"title"` // The document title
	Content_type string `json:"content-type"` // The content-type this URL serves
	Http_ok bool `json:"http-ok"` // True if this URL responded with an HTTP OK (200) status
	Server_region string `json:"server-region"` // The servers IP geo-location: full region name (if detectable)
	Url_path string `json:"url-path"` // The URL path
	Url_protocol string `json:"url-protocol"` // The URL protocol, usually http or https
	Server_hostname string `json:"server-hostname"` // The servers hostname (PTR record)
	Valid bool `json:"valid"` // Is this a valid well-formed URL
	RealField bool `json:"real"` // Is this URL actually serving real content
	Content string `json:"content"` // The actual content this URL responded with. Only set if the 'fetch-content' option was used
	Http_redirect bool `json:"http-redirect"` // True if this URL responded with an HTTP redirect
	Http_status int `json:"http-status"` // The HTTP status code this URL responded with. An HTTP status of 0 indicates a network level issue
	Is_error bool `json:"is-error"` // True if an error occurred while loading the URL. This includes network errors, TLS errors and timeouts
}

// IPBlocklistResponse represents the IPBlocklistResponse schema from the OpenAPI specification
type IPBlocklistResponse struct {
	Is_proxy bool `json:"is-proxy"` // IP has been detected as an anonymous web proxy or anonymous HTTP proxy
	Cidr string `json:"cidr"` // The CIDR address for this listing (only set if the IP is listed)
	Is_spam_bot bool `json:"is-spam-bot"` // IP address is hosting a spam bot, comment spamming or any other spamming type software
	Ip string `json:"ip"` // The IP address
	Is_dshield bool `json:"is-dshield"` // IP has been flagged as a significant attack source by DShield (dshield.org)
	Is_hijacked bool `json:"is-hijacked"` // IP is part of a hijacked netblock or a netblock controlled by a criminal organization
	Is_exploit_bot bool `json:"is-exploit-bot"` // IP is hosting an exploit finding bot or is running exploit scanning software
	Is_listed bool `json:"is-listed"` // Is this IP on a blocklist
	Is_spyware bool `json:"is-spyware"` // IP is involved in distributing or is running spyware
	Last_seen int `json:"last-seen"` // The unix time when this IP was last seen on any blocklist. IPs are automatically removed after 7 days therefor this value will never be older than 7 days
	List_count int `json:"list-count"` // The number of blocklists the IP is listed on
	Is_tor bool `json:"is-tor"` // IP is a Tor node or running a Tor related service
	Sensors []BlocklistSensor `json:"sensors"` // An array of objects containing details on which specific sensors detected the IP
	Blocklists []string `json:"blocklists"` // An array of strings indicating which blocklist categories this IP is listed on
	Is_bot bool `json:"is-bot"` // IP is hosting a malicious bot or is part of a botnet. This is a broad category which includes brute-force crackers
	Is_malware bool `json:"is-malware"` // IP is involved in distributing or is running malware
	Is_spider bool `json:"is-spider"` // IP is running a hostile web spider / web crawler
	Is_vpn bool `json:"is-vpn"` // IP belongs to a public VPN provider (only set if the 'vpn-lookup' option is enabled)
}

// GeocodeReverseResponse represents the GeocodeReverseResponse schema from the OpenAPI specification
type GeocodeReverseResponse struct {
	Currency_code string `json:"currency-code"` // ISO 4217 currency code associated with the country
	Latitude float64 `json:"latitude"` // The location latitude
	Location_type string `json:"location-type"` // The detected location type ordered roughly from most to least precise, possible values are: <br> <ul> <li>address - indicates a precise street address</li> <li>street - accurate to the street level but may not point to the exact location of the house/building number</li> <li>city - accurate to the city level, this includes villages, towns, suburbs, etc</li> <li>postal-code - indicates a postal code area (no house or street information present)</li> <li>railway - location is part of a rail network such as a station or railway track</li> <li>natural - indicates a natural feature, for example a mountain peak or a waterway</li> <li>island - location is an island or archipelago</li> <li>administrative - indicates an administrative boundary such as a country, state or province</li> </ul>
	Region_code string `json:"region-code"` // The ISO 3166-2 region code for the location
	Country string `json:"country"` // The country of the location
	Postal_code string `json:"postal-code"` // The postal code for the location
	State string `json:"state"` // The state of the location
	Address_components map[string]interface{} `json:"address-components"` // The components which make up the address such as road, city, state, etc
	City string `json:"city"` // The city of the location
	Location_tags []string `json:"location-tags"` // Array of strings containing any location tags associated with the address. Tags are additional pieces of metadata about a specific location, there are thousands of different tags. Some examples of tags: shop, office, cafe, bank, pub
	Postal_address string `json:"postal-address"` // The formatted address using local standards suitable for printing on an envelope
	Timezone map[string]interface{} `json:"timezone"` // Map containing timezone details for the location
	Longitude float64 `json:"longitude"` // The location longitude
	Found bool `json:"found"` // True if these coordinates map to a real location
	Country_code string `json:"country-code"` // The ISO 2-letter country code of the location
	Address string `json:"address"` // The complete address using comma-separated values
	Country_code3 string `json:"country-code3"` // The ISO 3-letter country code of the location
}

// PhonePlaybackResponse represents the PhonePlaybackResponse schema from the OpenAPI specification
type PhonePlaybackResponse struct {
	Number_valid bool `json:"number-valid"` // True if this a valid phone number
	Calling bool `json:"calling"` // True if the call is being made now
}

// EmailValidateResponse represents the EmailValidateResponse schema from the OpenAPI specification
type EmailValidateResponse struct {
	Email string `json:"email"` // The email address. If you have used the fix-typos option then this will be the fixed address
	Typos_fixed bool `json:"typos-fixed"` // True if typos have been fixed
	Is_freemail bool `json:"is-freemail"` // True if this address is a free-mail address
	Is_personal bool `json:"is-personal"` // True if this address belongs to a person. False if this is a role based address, e.g. admin@, help@, office@, etc.
	Domain string `json:"domain"` // The email domain
	Domain_error bool `json:"domain-error"` // True if this address has a domain error (e.g. no valid mail server records)
	Is_disposable bool `json:"is-disposable"` // True if this address is a disposable, temporary or darknet related email address
	Syntax_error bool `json:"syntax-error"` // True if this address has a syntax error
	Valid bool `json:"valid"` // Is this a valid email
	Provider string `json:"provider"` // The email service provider domain
}

// Timezone represents the Timezone schema from the OpenAPI specification
type Timezone struct {
	Date string `json:"date"` // The current date at the time zone (ISO 8601 format 'YYYY-MM-DD')
	Id string `json:"id"` // The time zone ID as per the IANA time zone database (tzdata)
	Name string `json:"name"` // The full time zone name
	Offset string `json:"offset"` // The UTC offset for the time zone (ISO 8601 format '±hh:mm')
	Time string `json:"time"` // The current time at the time zone (ISO 8601 format 'hh:mm:ss.sss')
	Abbr string `json:"abbr"` // The time zone abbreviation
}

// BadWordFilterResponse represents the BadWordFilterResponse schema from the OpenAPI specification
type BadWordFilterResponse struct {
	Bad_words_list []string `json:"bad-words-list"` // An array of the bad words found
	Bad_words_total int `json:"bad-words-total"` // Total number of bad words detected
	Censored_content string `json:"censored-content"` // The censored content (only set if censor-character has been set)
	Is_bad bool `json:"is-bad"` // Does the text contain bad words
}

// PhoneValidateResponse represents the PhoneValidateResponse schema from the OpenAPI specification
type PhoneValidateResponse struct {
	Currency_code string `json:"currency-code"` // ISO 4217 currency code associated with the country
	Is_mobile bool `json:"is-mobile"` // True if this is a mobile number. If the number type is unknown this value will be false
	Valid bool `json:"valid"` // Is this a valid phone number
	International_calling_code string `json:"international-calling-code"` // The international calling code
	International_number string `json:"international-number"` // The number represented in full international format (E.164)
	Local_number string `json:"local-number"` // The number represented in local dialing format
	TypeField string `json:"type"` // The number type based on the number prefix. <br>Possible values are: <br> <ul> <li>mobile</li> <li>fixed-line</li> <li>premium-rate</li> <li>toll-free</li> <li>voip</li> <li>unknown (use HLR lookup)</li> </ul>
	Country_code string `json:"country-code"` // The phone number country as an ISO 2-letter country code
	Country_code3 string `json:"country-code3"` // The phone number country as an ISO 3-letter country code
	Location string `json:"location"` // The phone number location. Could be the city, region or country depending on the type of number
	Prefix_network string `json:"prefix-network"` // The network/carrier who owns the prefix (this only works for some countries, use HLR lookup for global network detection)
	Country string `json:"country"` // The phone number country
}

// EmailVerifyResponse represents the EmailVerifyResponse schema from the OpenAPI specification
type EmailVerifyResponse struct {
	Verified bool `json:"verified"` // True if this address has passed SMTP verification. Check the smtp-status and smtp-response fields for specific verification details
	Is_deferred bool `json:"is-deferred"` // True if the mail server responded with a temporary failure (either a 4xx response code or unresponsive server). You can retry this address later, we recommend waiting at least 15 minutes before retrying
	Provider string `json:"provider"` // The email service provider domain
	Is_personal bool `json:"is-personal"` // True if this address is for a person. False if this is a role based address, e.g. admin@, help@, office@, etc.
	Syntax_error bool `json:"syntax-error"` // True if this address has a syntax error
	Smtp_status string `json:"smtp-status"` // The SMTP verification status for the address: <br> <ul> <li>ok - SMTP verification was successful, this is a real address that can receive mail</li> <li>invalid - this is not a valid email address (has either a domain or syntax error)</li> <li>absent - this address is not registered with the email service provider</li> <li>unresponsive - the mail server(s) for this address timed-out or refused to open an SMTP connection</li> <li>unknown - sorry, we could not reliably determine the real status of this address (this address may or may not exist)</li> </ul>
	Is_disposable bool `json:"is-disposable"` // True if this address is a disposable, temporary or darknet related email address
	Typos_fixed bool `json:"typos-fixed"` // True if typos have been fixed
	Domain_error bool `json:"domain-error"` // True if this address has a domain error (e.g. no valid mail server records)
	Is_catch_all bool `json:"is-catch-all"` // True if this email domain has a catch-all policy (it will accept mail for any username)
	Is_freemail bool `json:"is-freemail"` // True if this address is a free-mail address
	Valid bool `json:"valid"` // Is this a valid email address (syntax and domain is valid)
	Domain string `json:"domain"` // The email domain
	Smtp_response string `json:"smtp-response"` // The raw SMTP response message received during verification
	Email string `json:"email"` // The email address. If you have used the fix-typos option then this will be the fixed address
}

// GeocodeAddressResponse represents the GeocodeAddressResponse schema from the OpenAPI specification
type GeocodeAddressResponse struct {
	Locations []Location `json:"locations"` // Array of matching location objects
	Found int `json:"found"` // The number of possible matching locations found
}

// HostReputationResponse represents the HostReputationResponse schema from the OpenAPI specification
type HostReputationResponse struct {
	List_count int `json:"list-count"` // The number of DNSBLs the host is listed on
	Lists []Blacklist `json:"lists"` // Array of objects for each DNSBL checked
	Host string `json:"host"` // The IP address or host name
	Is_listed bool `json:"is-listed"` // Is this host blacklisted
}

// PhoneVerifyResponse represents the PhoneVerifyResponse schema from the OpenAPI specification
type PhoneVerifyResponse struct {
	Number_valid bool `json:"number-valid"` // True if this a valid phone number
	Security_code string `json:"security-code"` // The security code generated, you can save this code to perform your own verification or you can use the <a href="https://www.neutrinoapi.com/api/verify-security-code/">Verify Security Code API</a>
	Calling bool `json:"calling"` // True if the call is being made now
}

// SMSVerifyResponse represents the SMSVerifyResponse schema from the OpenAPI specification
type SMSVerifyResponse struct {
	Number_valid bool `json:"number-valid"` // True if this a valid phone number
	Security_code string `json:"security-code"` // The security code generated, you can save this code to perform your own verification or you can use the <a href="https://www.neutrinoapi.com/api/verify-security-code/">Verify Security Code API</a>
	Sent bool `json:"sent"` // True if the SMS has been sent
}

// DomainLookupResponse represents the DomainLookupResponse schema from the OpenAPI specification
type DomainLookupResponse struct {
	Tld string `json:"tld"` // The top-level domain (TLD)
	Age int `json:"age"` // The number of days since the domain was registered. A domain age of under 90 days is generally considered to be potentially risky. A value of 0 indicates no registration date was found for this domain
	Valid bool `json:"valid"` // True if a valid domain was found. For a domain to be considered valid it must be registered and have valid DNS NS records
	Blocklists []string `json:"blocklists"` // An array of strings indicating which blocklist categories this domain is listed on. Current categories are: phishing, malware, spam, anonymizer, nefarious
	Sensors []BlocklistSensor `json:"sensors"` // An array of objects containing details on which specific blocklist sensors have detected this domain
	Registrar_id int `json:"registrar-id"` // The IANA registrar ID (0 if no registrar ID was found)
	Rank int `json:"rank"` // The domains estimated global traffic rank with the highest rank being 1. A value of 0 indicates the domain is currently ranked outside of the top 1M of domains
	Registered_date string `json:"registered-date"` // The ISO date this domain was registered or first seen on the internet. An empty value indicates we could not reliably determine the date
	Mail_provider string `json:"mail-provider"` // The primary domain of the email provider for this domain. An empty value indicates the domain has no valid MX records
	Dns_provider string `json:"dns-provider"` // The primary domain of the DNS provider for this domain
	Fqdn string `json:"fqdn"` // The fully qualified domain name (FQDN)
	Tld_cc string `json:"tld-cc"` // For a country code top-level domain (ccTLD) this will contain the associated ISO 2-letter country code
	Is_malicious bool `json:"is-malicious"` // Consider this domain malicious as it is currently listed on at least 1 blocklist
	Registrar_name string `json:"registrar-name"` // The name of the domain registrar owning this domain
	Domain string `json:"domain"` // The primary domain name excluding any subdomains. This is also referred to as the second-level domain (SLD)
	Is_opennic bool `json:"is-opennic"` // Is this domain under an OpenNIC TLD
	Is_subdomain bool `json:"is-subdomain"` // Is the FQDN a subdomain of the primary domain
	Is_pending bool `json:"is-pending"` // True if this domain is unseen and is currently being processed in the background. This field only matters when the 'live' lookup setting has been explicitly disabled and indicates that not all domain data my be present yet
	Is_adult bool `json:"is-adult"` // This domain is hosting adult content such as porn, webcams, escorts, etc
	Is_gov bool `json:"is-gov"` // Is this domain under a government or military TLD
}

// BlocklistSensor represents the BlocklistSensor schema from the OpenAPI specification
type BlocklistSensor struct {
	Blocklist string `json:"blocklist"` // The primary blocklist category this sensor belongs to
	Description string `json:"description"` // Contains details about the sensor source and what type of malicious activity was detected
	Id int `json:"id"` // The sensor ID. This is a permanent and unique ID for each sensor
}

// Blacklist represents the Blacklist schema from the OpenAPI specification
type Blacklist struct {
	List_host string `json:"list-host"` // The hostname of the DNSBL
	List_name string `json:"list-name"` // The name of the DNSBL
	List_rating int `json:"list-rating"` // The list rating [1-3] with 1 being the best rating and 3 the lowest rating
	Response_time int `json:"response-time"` // The DNSBL server response time in milliseconds
	Return_code string `json:"return-code"` // The specific return code for this listing (only set if listed)
	Txt_record string `json:"txt-record"` // The TXT record returned for this listing (only set if listed)
	Is_listed bool `json:"is-listed"` // True if the host is currently black-listed
}
